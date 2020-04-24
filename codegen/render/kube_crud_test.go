package render_test

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"github.com/solo-io/skv2/pkg/multicluster"
	mc_client "github.com/solo-io/skv2/pkg/multicluster/client"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/watch"
	"github.com/solo-io/skv2/pkg/reconcile"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/randutils"
	kubehelp "github.com/solo-io/go-utils/testutils/kube"
	. "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/controller"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/skv2/test"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/log"
	zaputil "sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func applyFile(file string) error {
	path := filepath.Join(util.GetModuleRoot(), "codegen/test/chart/crds", file)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return util.KubectlApply(b)
}

func newPaint(namespace, name string) *Paint {
	return &Paint{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: PaintSpec{
			Color: &PaintColor{
				Hue:   "prussian blue",
				Value: 0.5,
			},
			PaintType: &PaintSpec_Acrylic{
				Acrylic: &AcrylicType{
					Body: AcrylicType_Heavy,
				},
			},
		},
	}
}

func mustCrud(ctx context.Context, clientSet Clientset, paint *Paint) {
	err := clientSet.Paints().CreatePaint(ctx, paint)
	Expect(err).NotTo(HaveOccurred())

	written, err := clientSet.Paints().GetPaint(ctx, client.ObjectKey{
		Namespace: paint.Namespace,
		Name:      paint.Name,
	})
	Expect(err).NotTo(HaveOccurred())

	Expect(written.Spec).To(Equal(paint.Spec))

	status := PaintStatus{
		ObservedGeneration: written.Generation,
		PercentRemaining:   22,
	}

	written.Status = status

	err = clientSet.Paints().UpdatePaintStatus(ctx, written)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() PaintStatus {
		written, err = clientSet.Paints().GetPaint(ctx, client.ObjectKey{
			Namespace: paint.Namespace,
			Name:      paint.Name,
		})
		Expect(err).NotTo(HaveOccurred())
		return written.Status
	}, time.Second).Should(Equal(status))
}

var _ = Describe("Generated Code", func() {
	var (
		ns        string
		kube      kubernetes.Interface
		clientSet Clientset
		logLevel  = zap.NewAtomicLevel()
		ctx       = context.TODO()
	)
	BeforeEach(func() {
		logLevel.SetLevel(zap.DebugLevel)
		log.SetLogger(zaputil.New(
			zaputil.Level(&logLevel),
		))
		log.Log.Info("test")
		err := applyFile("things.test.io_v1_crds.yaml")
		Expect(err).NotTo(HaveOccurred())
		ns = randutils.RandString(4)
		kube = kubehelp.MustKubeClient()
		err = kubeutils.CreateNamespacesInParallel(kube, ns)
		Expect(err).NotTo(HaveOccurred())
		clientSet, err = NewClientsetFromConfig(test.MustConfig())
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		err := kubeutils.DeleteNamespacesInParallelBlocking(kube, ns)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("kube clientsets", func() {
		It("uses the generated clientsets to crud", func() {

			paint := newPaint(ns, "paint-kube-clientsets")

			mustCrud(ctx, clientSet, paint)
		})
	})

	Context("kube reconciler", func() {
		var (
			mgr    manager.Manager
			cancel = func() {}
		)
		BeforeEach(func() {
			mgr, cancel = test.MustManager(ns)
		})
		AfterEach(cancel)

		It("uses the generated controller to reconcile", func() {

			paint := newPaint(ns, "paint-kube-reconciler")

			err := clientSet.Paints().CreatePaint(ctx, paint)
			Expect(err).NotTo(HaveOccurred())

			paint.GetObjectKind().GroupVersionKind()

			loop := controller.NewPaintReconcileLoop("blick", mgr)

			var reconciled *Paint
			var deleted reconcile.Request
			reconciler := &controller.PaintReconcilerFuncs{
				OnReconcilePaint: func(obj *Paint) (result reconcile.Result, err error) {
					reconciled = obj
					return
				},
				OnReconcilePaintDeletion: func(req reconcile.Request) {
					deleted = req
					return
				},
			}
			err = loop.RunPaintReconciler(ctx, reconciler)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() *Paint {
				return reconciled
			}, time.Second).ShouldNot(BeNil())

			// update
			paint.Spec.Color = &PaintColor{Value: 0.7}

			err = clientSet.Paints().UpdatePaint(ctx, paint)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() PaintSpec {
				return reconciled.Spec
			}, time.Second).Should(Equal(paint.Spec))

			// delete
			err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			})
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() reconcile.Request {
				return deleted
			}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			}}))
		})
	})

	Context("multicluster", func() {
		var (
			ctx           context.Context
			cancel        = func() {}
			masterManager manager.Manager
			kcSecret      *corev1.Secret
			cluster2      = "cluster-two"
			err           error
		)

		mustNewKubeConfigSecret := func() *corev1.Secret {
			localKubeConfig, err := kubeutils.GetKubeConfig("", "")
			Expect(err).NotTo(HaveOccurred())
			kcSecret, err = kubeconfig.ToSecret(ns, cluster2, *localKubeConfig)
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
			kubehelp.MustKubeClient().CoreV1().Secrets(ns).Delete(kcSecret.Name, &v1.DeleteOptions{})
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

type concurrentPaintMap struct {
	m     map[string]*Paint
	mutex sync.Mutex
}

func newConcurrentPaintMap() *concurrentPaintMap {
	return &concurrentPaintMap{
		m:     make(map[string]*Paint),
		mutex: sync.Mutex{},
	}
}

func (c *concurrentPaintMap) add(cluster string, paint *Paint) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m[cluster+paint.Name] = paint
}

func (c *concurrentPaintMap) get(cluster, name string) *Paint {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	paint := c.m[cluster+name]
	return paint
}

type concurrentRequestMap struct {
	m     map[string]reconcile.Request
	mutex sync.Mutex
}

func newConcurrentRequestMap() *concurrentRequestMap {
	return &concurrentRequestMap{
		m:     make(map[string]reconcile.Request),
		mutex: sync.Mutex{},
	}
}

func (c *concurrentRequestMap) add(cluster string, req reconcile.Request) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m[cluster+req.Name] = req
}

func (c *concurrentRequestMap) get(cluster, name string) reconcile.Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	req := c.m[cluster+name]
	return req
}
