package render_test

import (
	"context"
	"os"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/randutils"
	kubehelp "github.com/solo-io/go-utils/testutils/kube"
	. "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/controller"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_client "github.com/solo-io/skv2/pkg/multicluster/client"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	"github.com/solo-io/skv2/pkg/multicluster/watch"
	"github.com/solo-io/skv2/pkg/reconcile"
	"github.com/solo-io/skv2/test"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	zaputil "sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func WithRemoteClusterContextDescribe(text string, body func()) bool {
	if os.Getenv("REMOTE_CLUSTER_CONTEXT") == "" {
		return PDescribe("[This test depends on a second cluster with context REMOTE_CLUSTER_CONTEXT] "+text, body)
	}
	return Describe(text, body)
}

var _ = WithRemoteClusterContextDescribe("Multicluster", func() {
	var (
		ns              string
		masterClientSet Clientset
		remoteClientSet Clientset
		logLevel        = zap.NewAtomicLevel()
		ctx             = context.TODO()
		remoteContext   = os.Getenv("REMOTE_CLUSTER_CONTEXT")

		cancel        context.CancelFunc
		masterManager manager.Manager
		cluster2      = "cluster-two"
		cluster1      = "cluster-one"
	)

	BeforeEach(func() {
		logLevel.SetLevel(zap.DebugLevel)
		log.SetLogger(zaputil.New(
			zaputil.Level(&logLevel),
		))
		ns = randutils.RandString(4)
		var err error
		for _, kubeContext := range []string{"", remoteContext} {
			err = applyFile("things.test.io_v1_crds.yaml", "--context", kubeContext)
			Expect(err).NotTo(HaveOccurred())
			cfg := test.MustConfig(kubeContext)
			kube := kubernetes.NewForConfigOrDie(cfg)
			err = kubeutils.CreateNamespacesInParallel(kube, ns)
			Expect(err).NotTo(HaveOccurred())
		}
		masterClientSet, err = NewClientsetFromConfig(test.MustConfig(""))
		Expect(err).NotTo(HaveOccurred())
		remoteClientSet, err = NewClientsetFromConfig(test.MustConfig(remoteContext))
		Expect(err).NotTo(HaveOccurred())

		ctx, cancel = context.WithCancel(context.Background())
		masterManager = test.MustManager(ctx, ns)
		remoteCfg := test.ClientConfigWithContext(remoteContext)
		registrant, err := register.DefaultRegistrant("")
		Expect(err).NotTo(HaveOccurred())
		err = register.RegisterClusterFromConfig(ctx, remoteCfg, register.RbacOptions{
			Options: register.Options{
				ClusterName: cluster2,
				Namespace:   ns,
				RemoteCtx:   remoteContext,
			},
			ClusterRoleBindings: test.ServiceAccountClusterAdminRoles,
		}, registrant)
		Expect(err).NotTo(HaveOccurred())
		cfg := test.ClientConfigWithContext("")
		err = register.RegisterClusterFromConfig(ctx, cfg, register.RbacOptions{
			Options: register.Options{
				ClusterName: cluster1,
				Namespace:   ns,
			},
			ClusterRoleBindings: test.ServiceAccountClusterAdminRoles,
		}, registrant)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if cancel != nil {
			cancel()
		}
		for _, kubeContext := range []string{"", remoteContext} {
			cfg := test.MustConfig(kubeContext)
			kube := kubernetes.NewForConfigOrDie(cfg)
			err := kubeutils.DeleteNamespacesInParallelBlocking(kube, ns)
			Expect(err).NotTo(HaveOccurred())
		}
	})

	Context("multicluster", func() {

		FDescribe("clientset", func() {
			It("works", func() {
				cw := watch.NewClusterWatcher(ctx, manager.Options{Namespace: ns})
				err := cw.Run(masterManager)
				Expect(err).NotTo(HaveOccurred())
				mcClientset := mc_client.NewClient(cw)
				clientset := NewMulticlusterClientset(mcClientset)

				masterPaint := newPaint(ns, "paint-mc-clientset-master")
				var masterClusterClientset Clientset
				Eventually(func() error {
					masterClusterClientset, err = clientset.Cluster(cluster1)
					return err
				}, "5s").Should(Not(HaveOccurred()))
				mustCrud(context.TODO(), masterClusterClientset, masterPaint)

				cluster2Paint := newPaint(ns, "paint-mc-clientset-cluster2")
				var cluster2ClusterClientset Clientset
				Eventually(func() Clientset {
					cluster2ClusterClientset, _ = clientset.Cluster(cluster2)
					return cluster2ClusterClientset
				}, time.Second).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
				mustCrud(context.TODO(), cluster2ClusterClientset, cluster2Paint)

				_, err = clientset.Cluster("nonexistent-cluster")
				Expect(err).To(HaveOccurred())
			})
		})

		Describe("kube reconciler", func() {
			var (
				cw multicluster.ClusterWatcher
			)

			mustReconcile := func(clientSet Clientset, cluster string, paint *Paint, paints *concurrentPaintMap, reqs *concurrentRequestMap) {
				err := clientSet.Paints().CreatePaint(ctx, paint)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() *Paint {
					return paints.get(cluster, paint.Name)
				}, time.Second*20).ShouldNot(BeNil())

				// update
				paint.Spec.Color = &PaintColor{Value: 0.7}

				err = clientSet.Paints().UpdatePaint(ctx, paint)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() PaintSpec {
					return paints.get(cluster, paint.Name).Spec
				}, time.Second).Should(Equal(paint.Spec))

				// delete
				err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				})
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() reconcile.Request {
					return reqs.get(cluster, paint.Name)
				}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				}}))
			}

			BeforeEach(func() {
				cw = watch.NewClusterWatcher(ctx, manager.Options{Namespace: ns})
			})

			It("works when a loop is registered before the watcher is started", func() {
				loop := controller.NewMulticlusterPaintReconcileLoop("pre-run-paint", cw)

				paintMap := &concurrentPaintMap{}
				paintDeletes := &concurrentRequestMap{}

				loop.AddMulticlusterPaintReconciler(ctx, &controller.MulticlusterPaintReconcilerFuncs{
					OnReconcilePaint: func(clusterName string, obj *Paint) (result reconcile.Result, e error) {
						paintMap.add(clusterName, obj)
						return result, e
					},
					OnReconcilePaintDeletion: func(clusterName string, req reconcile.Request) {
						paintDeletes.add(clusterName, req)
					},
				})

				err := cw.Run(masterManager)
				Expect(err).NotTo(HaveOccurred())

				mustReconcile(masterClientSet, cluster1, newPaint(ns, "paint-mc-pre-start"), paintMap, paintDeletes)
				mustReconcile(remoteClientSet, cluster2, newPaint(ns, "paint-mc-pre-start"), paintMap, paintDeletes)
			})

			It("works when a kubeconfig secret is deleted and recreated", func() {
				loop := controller.NewMulticlusterPaintReconcileLoop("paint", cw)

				paintMap := &concurrentPaintMap{}
				paintDeletes := &concurrentRequestMap{}

				loop.AddMulticlusterPaintReconciler(ctx, &controller.MulticlusterPaintReconcilerFuncs{
					OnReconcilePaint: func(clusterName string, obj *Paint) (result reconcile.Result, e error) {
						paintMap.add(clusterName, obj)
						return result, e
					},
					OnReconcilePaintDeletion: func(clusterName string, req reconcile.Request) {
						paintDeletes.add(clusterName, req)
					},
				})

				err := cw.Run(masterManager)
				Expect(err).NotTo(HaveOccurred())

				mustReconcile(masterClientSet, cluster1, newPaint(ns, "paint-mc-delete-recreate"), paintMap, paintDeletes)
				mustReconcile(remoteClientSet, cluster2, newPaint(ns, "paint-mc-delete-recreate"), paintMap, paintDeletes)
				/**
				When a KubeConfig secret is deleted, paint is no longer reconciled for that cluster
				Save that secret so it can be recreated later
				*/
				remoteKcSecret, err := kubehelp.MustKubeClient().CoreV1().Secrets(ns).Get(cluster2, metav1.GetOptions{})
				Expect(err).NotTo(HaveOccurred())
				err = kubehelp.MustKubeClient().CoreV1().Secrets(ns).Delete(remoteKcSecret.Name, &metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())

				// Sleep to be sure that the secret deletion propagates.
				// need to allow
				time.Sleep(time.Second)

				// Expect a new paint to reconcile as usual on the master cluster
				mustReconcile(masterClientSet, cluster1, newPaint(ns, "paint-mc-delete-recreate"), paintMap, paintDeletes)

				// Expect a new paint to NOT reconcile on the remote cluster, as we've deleted the kubeconfig
				paint2 := newPaint(ns, "paint-mc-reconciler-2")
				err = remoteClientSet.Paints().CreatePaint(ctx, paint2)
				Expect(err).NotTo(HaveOccurred())

				// Sleep the amount of time we typically allow for a reconcile to happen.
				Eventually(paintMap.get(cluster2, paint2.Name)).Should(BeNil())
				// delete
				err = remoteClientSet.Paints().DeletePaint(ctx, client.ObjectKey{
					Name:      paint2.Name,
					Namespace: paint2.Namespace,
				})
				Expect(err).NotTo(HaveOccurred())

				// Sleep the amount of time we typically allow for a reconcile to happen.
				Eventually(paintDeletes.get(cluster2, paint2.Name), "5s").Should(Equal(reconcile.Request{}))

				/**
				When a KubeConfig secret is restored, paint is again reconciled for that cluster
				*/
				secretCopy := &v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      remoteKcSecret.GetName(),
						Namespace: remoteKcSecret.GetNamespace(),
						Labels:    remoteKcSecret.GetLabels(),
					},
					Data: remoteKcSecret.DeepCopy().Data,
					Type: kubeconfig.SecretType,
				}
				_, err = kubehelp.MustKubeClient().CoreV1().Secrets(ns).Create(secretCopy)
				Expect(err).NotTo(HaveOccurred())

				mustReconcile(masterClientSet, cluster1, newPaint(ns, "paint-mc-reconciler-3"), paintMap, paintDeletes)
				mustReconcile(remoteClientSet, cluster2, newPaint(ns, "paint-mc-reconciler-3"), paintMap, paintDeletes)
			})

			It("works when a loop is registered after the watcher is started", func() {
				err := cw.Run(masterManager)
				Expect(err).NotTo(HaveOccurred())

				loop := controller.NewMulticlusterPaintReconcileLoop("mid-run-paint", cw)

				paintMap := &concurrentPaintMap{}
				deleteMap := &concurrentRequestMap{}

				loop.AddMulticlusterPaintReconciler(ctx, &controller.MulticlusterPaintReconcilerFuncs{
					OnReconcilePaint: func(clusterName string, obj *Paint) (result reconcile.Result, e error) {
						paintMap.add(clusterName, obj)
						return result, e
					},
					OnReconcilePaintDeletion: func(clusterName string, req reconcile.Request) {
						deleteMap.add(clusterName, req)
					},
				})

				mustReconcile(masterClientSet, cluster1, newPaint(ns, "late-registration"), paintMap, deleteMap)
				mustReconcile(remoteClientSet, cluster2, newPaint(ns, "late-registration"), paintMap, deleteMap)
			})
		})
	})
})

type concurrentPaintMap struct {
	sync.Map
}

func (c *concurrentPaintMap) add(cluster string, paint *Paint) {
	c.Store(cluster+paint.Name, paint)
}

func (c *concurrentPaintMap) get(cluster, name string) *Paint {
	obj, _ := c.Load(cluster + name)
	if paint, ok := obj.(*Paint); ok {
		return paint
	}
	return nil
}

type concurrentRequestMap struct {
	sync.Map
}

func (c *concurrentRequestMap) add(cluster string, req reconcile.Request) {
	c.Store(cluster+req.Name, req)
}

func (c *concurrentRequestMap) get(cluster, name string) reconcile.Request {
	obj, _ := c.Load(cluster + name)
	if req, ok := obj.(reconcile.Request); ok {
		return req
	}
	return reconcile.Request{}
}
