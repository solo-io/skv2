package render_test

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/randutils"
	. "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	mc_client "github.com/solo-io/skv2/pkg/multicluster/client"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/solo-io/go-utils/kubeutils"
	kubehelp "github.com/solo-io/go-utils/testutils/kube"
	"github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/controller"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/watch"
	"github.com/solo-io/skv2/pkg/reconcile"
	"github.com/solo-io/skv2/test"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	zaputil "sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func WithRemoteClusterContextDescribe(text string, body func()) bool {
	if os.Getenv("REMOTE_CLUSTER_CONTEXT") == "" {
		return PDescribe("This test depends on a second cluster with context REMOTE_CLUSTER_CONTEXT. "+text, body)
	}
	return Describe(text, body)
}

var _ = WithRemoteClusterContextDescribe("Multicluster", func() {
	var (
		ns            string
		clientSet     Clientset
		logLevel      = zap.NewAtomicLevel()
		ctx           = context.TODO()
		remoteContext = os.Getenv("REMOTE_CLUSTER_CONTEXT")
	)

	BeforeEach(func() {
		logLevel.SetLevel(zap.DebugLevel)
		log.SetLogger(zaputil.New(
			zaputil.Level(&logLevel),
		))
		ns = randutils.RandString(4)
		for _, kubeContext := range []string{"", remoteContext} {
			err := applyFile("things.test.io_v1_crds.yaml", "--context", kubeContext)
			Expect(err).NotTo(HaveOccurred())
			cfg := test.MustConfig(kubeContext)
			kube := kubernetes.NewForConfigOrDie(cfg)
			err = kubeutils.CreateNamespacesInParallel(kube, ns)
			Expect(err).NotTo(HaveOccurred())
			clientSet, err = NewClientsetFromConfig(test.MustConfig(""))
			Expect(err).NotTo(HaveOccurred())
		}
	})

	AfterEach(func() {
		for _, kubeContext := range []string{"", remoteContext} {
			cfg := test.MustConfig(kubeContext)
			kube := kubernetes.NewForConfigOrDie(cfg)
			err := kubeutils.DeleteNamespacesInParallelBlocking(kube, ns)
			Expect(err).NotTo(HaveOccurred())
		}
	})

	Context("multicluster", func() {
		var (
			cancel        = func() {}
			masterManager manager.Manager
			kcSecret      *corev1.Secret
			cluster2      = "cluster-two"
			err           error
		)

		mustNewKubeConfigSecret := func() *corev1.Secret {
			cfg, err := kubeutils.GetKubeConfigWithContext("", os.Getenv("KUBECONFIG"), remoteContext)
			Expect(err).NotTo(HaveOccurred())
			kcSecret, err = kubeconfig.ToSecret(ns, cluster2, *cfg)
			Expect(err).NotTo(HaveOccurred())
			return kcSecret
		}

		BeforeEach(func() {
			ctx, cancel = context.WithCancel(context.Background())
			masterManager = test.MustManagerNotStarted(ns)
			kcSecret, err = kubehelp.MustKubeClient().CoreV1().Secrets(ns).Create(mustNewKubeConfigSecret())
		})
		AfterEach(func() {
			cancel()
			for _, kubeContext := range []string{"", remoteContext} {
				cfg, err := kubeutils.GetConfigWithContext("", os.Getenv("KUBECONFIG"), kubeContext)
				Expect(err).NotTo(HaveOccurred())
				kube := kubernetes.NewForConfigOrDie(cfg)
				kube.CoreV1().Secrets(ns).Delete(kcSecret.Name, &v1.DeleteOptions{})
			}
		})

		Describe("clientset", func() {
			It("works", func() {
				cw := watch.NewClusterWatcher(ctx, manager.Options{Namespace: ns})
				err := cw.Run(masterManager)
				Expect(err).NotTo(HaveOccurred())
				mcClientset := mc_client.NewClient(cw)
				clientset := NewMulticlusterClientset(mcClientset)

				masterPaint := newPaint(ns, "paint-mc-clientset-master")
				masterClusterClientset, err := clientset.Cluster(multicluster.MasterCluster)
				Expect(err).NotTo(HaveOccurred())
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
				cw    multicluster.ClusterWatcher
				paint *Paint
			)

			mustReconcile := func(paint *Paint, paints *concurrentPaintMap, reqs *concurrentRequestMap) {
				Eventually(func() *Paint {
					return paints.get(multicluster.MasterCluster, paint.Name)
				}, time.Second).ShouldNot(BeNil())

				Eventually(func() *Paint {
					return paints.get(cluster2, paint.Name)
				}, time.Second).ShouldNot(BeNil())

				// update
				paint.Spec.Color = &PaintColor{Value: 0.7}

				err = clientSet.Paints().UpdatePaint(ctx, paint)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() PaintSpec {
					return paints.get(multicluster.MasterCluster, paint.Name).Spec
				}, time.Second).Should(Equal(paint.Spec))

				Eventually(func() PaintSpec {
					return paints.get(cluster2, paint.Name).Spec
				}, time.Second).Should(Equal(paint.Spec))

				// delete
				err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				})
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() reconcile.Request {
					return reqs.get(multicluster.MasterCluster, paint.Name)
				}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				}}))

				Eventually(func() reconcile.Request {
					return reqs.get(cluster2, paint.Name)
				}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				}}))
			}

			BeforeEach(func() {
				cw = watch.NewClusterWatcher(ctx, manager.Options{Namespace: ns})

				Expect(err).NotTo(HaveOccurred())

				paint = newPaint(ns, "paint-mc-reonciler")

				err = clientSet.Paints().CreatePaint(ctx, paint)
				Expect(err).NotTo(HaveOccurred())
			})

			It("works when a loop is registered before the watcher is started", func() {
				loop := controller.NewMulticlusterPaintReconcileLoop("pre-run-paint", cw)

				paintMap := newConcurrentPaintMap()
				paintDeletes := newConcurrentRequestMap()

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

				mustReconcile(paint, paintMap, paintDeletes)
			})

			It("works when a kubeconfig secret is deleted and recreated", func() {
				loop := controller.NewMulticlusterPaintReconcileLoop("paint", cw)

				paintMap := newConcurrentPaintMap()
				paintDeletes := newConcurrentRequestMap()

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

				mustReconcile(paint, paintMap, paintDeletes)

				/**
				When a KubeConfig secret is deleted, paint is no longer reconciled for that cluster
				*/

				err = kubehelp.MustKubeClient().CoreV1().Secrets(ns).Delete(kcSecret.Name, &v1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred())

				// Sleep to be sure that the secret deletion propagates.
				time.Sleep(100 * time.Millisecond)

				paint2 := newPaint(ns, "paint-mc-reconciler-2")
				err = clientSet.Paints().CreatePaint(ctx, paint2)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() *Paint {
					return paintMap.get(multicluster.MasterCluster, paint2.Name)
				}, time.Second).ShouldNot(BeNil())

				// Sleep the amount of time we typically allow for a reconcile to happen.
				time.Sleep(1 * time.Second)
				Expect(paintMap.get(cluster2, paint2.Name)).To(BeNil())

				// update
				paint2.Spec.Color = &PaintColor{Value: 0.7}

				err = clientSet.Paints().UpdatePaint(ctx, paint2)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() PaintSpec {
					return paintMap.get(multicluster.MasterCluster, paint2.Name).Spec
				}, time.Second).Should(Equal(paint2.Spec))

				// delete
				err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
					Name:      paint2.Name,
					Namespace: paint2.Namespace,
				})
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() reconcile.Request {
					return paintDeletes.get(multicluster.MasterCluster, paint2.Name)
				}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      paint2.Name,
					Namespace: paint2.Namespace,
				}}))

				// Sleep the amount of time we typically allow for a reconcile to happen.
				time.Sleep(1 * time.Second)
				Expect(paintDeletes.get(cluster2, paint2.Name)).To(Equal(reconcile.Request{}))

				/**
				When a KubeConfig secret is restored, paint is again reconciled for that cluster
				*/

				_, err = kubehelp.MustKubeClient().CoreV1().Secrets(ns).Create(mustNewKubeConfigSecret())
				Expect(err).NotTo(HaveOccurred())

				paint3 := newPaint(ns, "paint-mc-reconciler-3")
				err = clientSet.Paints().CreatePaint(ctx, paint3)
				Expect(err).NotTo(HaveOccurred())

				mustReconcile(paint3, paintMap, paintDeletes)
			})

			It("works when a loop is registered after the watcher is started", func() {
				err := cw.Run(masterManager)
				Expect(err).NotTo(HaveOccurred())

				loop := controller.NewMulticlusterPaintReconcileLoop("mid-run-paint", cw)

				midRunReconciledPaint := newConcurrentPaintMap()
				midRunReconciledDeleteRequests := newConcurrentRequestMap()

				loop.AddMulticlusterPaintReconciler(ctx, &controller.MulticlusterPaintReconcilerFuncs{
					OnReconcilePaint: func(clusterName string, obj *Paint) (result reconcile.Result, e error) {
						midRunReconciledPaint.add(clusterName, obj)
						return result, e
					},
					OnReconcilePaintDeletion: func(clusterName string, req reconcile.Request) {
						midRunReconciledDeleteRequests.add(clusterName, req)
					},
				})

				Eventually(func() *Paint {
					return midRunReconciledPaint.get(multicluster.MasterCluster, paint.Name)
				}, time.Second).ShouldNot(BeNil())

				Eventually(func() *Paint {
					return midRunReconciledPaint.get(cluster2, paint.Name)
				}, time.Second).ShouldNot(BeNil())

				err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				})
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() reconcile.Request {
					return midRunReconciledDeleteRequests.get(multicluster.MasterCluster, paint.Name)
				}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				}}))

				Eventually(func() reconcile.Request {
					return midRunReconciledDeleteRequests.get(cluster2, paint.Name)
				}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      paint.Name,
					Namespace: paint.Namespace,
				}}))
			})
		})
	})
})
