package render_test

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	gen_multicluster "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/controller/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
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

	Context("kube clientsests", func() {
		It("uses the generated clientsets to crud", func() {

			paint := &Paint{
				ObjectMeta: v1.ObjectMeta{
					Name:      "paint-1",
					Namespace: ns,
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

			written, err = clientSet.Paints().GetPaint(ctx, client.ObjectKey{
				Namespace: paint.Namespace,
				Name:      paint.Name,
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(written.Status).To(Equal(status))
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

			paint := &Paint{
				ObjectMeta: v1.ObjectMeta{
					Name:      "paint-2",
					Namespace: ns,
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

	Context("multicluster kube reconciler", func() {
		var (
			ctx      context.Context
			cancel   context.CancelFunc
			mgr      manager.Manager
			cw       multicluster.ClusterWatcher
			cluster2 = "foo"
			kcSecret *corev1.Secret
			paint    *Paint
		)
		BeforeEach(func() {
			ctx, cancel = context.WithCancel(context.Background())
			mgr = test.MustManagerNotStarted(ns)
			cw = multicluster.NewClusterWatcher(ctx, manager.Options{Namespace: ns})

			localKubeConfig, err := kubeutils.GetKubeConfig("", "")
			Expect(err).NotTo(HaveOccurred())
			kcSecret, err = kubeconfig.ToSecret(ns, cluster2, *localKubeConfig)
			Expect(err).NotTo(HaveOccurred())
			kcSecret, err = kubehelp.MustKubeClient().CoreV1().Secrets(ns).Create(kcSecret)
			Expect(err).NotTo(HaveOccurred())

			paint = &Paint{
				ObjectMeta: v1.ObjectMeta{
					Name:      "paint-2",
					Namespace: ns,
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

			err = clientSet.Paints().CreatePaint(ctx, paint)
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			cancel()
			err := kubehelp.MustKubeClient().CoreV1().Secrets(ns).Delete(kcSecret.Name, &v1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred())
		})

		It("works when a loop is registered before the watcher is started", func() {
			loop := gen_multicluster.NewMulticlusterPaintReconcileLoop("pre-run-paint", cw)

			preRunReconciledPaint := newConcurrentPaintMap()
			preRunReconciledDeleteRequests := newConcurrentRequestMap()

			loop.AddMulticlusterPaintReconciler(ctx, &gen_multicluster.MulticlusterPaintReconcilerFuncs{
				OnReconcilePaint: func(clusterName string, obj *Paint) (result reconcile.Result, e error) {
					preRunReconciledPaint.add(clusterName, obj)
					return result, e
				},
				OnReconcilePaintDeletion: func(clusterName string, req reconcile.Request) {
					preRunReconciledDeleteRequests.add(clusterName, req)
				},
			})

			err := cw.Run(mgr)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() *Paint {
				return preRunReconciledPaint.get(multicluster.MasterCluster)
			}, time.Second).ShouldNot(BeNil())

			Eventually(func() *Paint {
				return preRunReconciledPaint.get(cluster2)
			}, time.Second).ShouldNot(BeNil())

			// update
			paint.Spec.Color = &PaintColor{Value: 0.7}

			err = clientSet.Paints().UpdatePaint(ctx, paint)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() PaintSpec {
				return preRunReconciledPaint.get(multicluster.MasterCluster).Spec
			}, time.Second).Should(Equal(paint.Spec))

			Eventually(func() PaintSpec {
				return preRunReconciledPaint.get(cluster2).Spec
			}, time.Second).Should(Equal(paint.Spec))

			// delete
			err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			})
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() reconcile.Request {
				return preRunReconciledDeleteRequests.get(multicluster.MasterCluster)
			}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			}}))

			Eventually(func() reconcile.Request {
				return preRunReconciledDeleteRequests.get(cluster2)
			}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			}}))

		})

		It("works when a loop is registered after the watcher is started", func() {
			err := cw.Run(mgr)
			Expect(err).NotTo(HaveOccurred())

			loop := gen_multicluster.NewMulticlusterPaintReconcileLoop("mid-run-paint", cw)

			midRunReconciledPaint := newConcurrentPaintMap()
			midRunReconciledDeleteRequests := newConcurrentRequestMap()

			loop.AddMulticlusterPaintReconciler(ctx, &gen_multicluster.MulticlusterPaintReconcilerFuncs{
				OnReconcilePaint: func(clusterName string, obj *Paint) (result reconcile.Result, e error) {
					midRunReconciledPaint.add(clusterName, obj)
					return result, e
				},
				OnReconcilePaintDeletion: func(clusterName string, req reconcile.Request) {
					midRunReconciledDeleteRequests.add(clusterName, req)
				},
			})

			Eventually(func() *Paint {
				return midRunReconciledPaint.get(multicluster.MasterCluster)
			}, time.Second).ShouldNot(BeNil())

			Eventually(func() *Paint {
				return midRunReconciledPaint.get(cluster2)
			}, time.Second).ShouldNot(BeNil())

			err = clientSet.Paints().DeletePaint(ctx, client.ObjectKey{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			})
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() reconcile.Request {
				return midRunReconciledDeleteRequests.get(multicluster.MasterCluster)
			}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			}}))

			Eventually(func() reconcile.Request {
				return midRunReconciledDeleteRequests.get(cluster2)
			}, time.Second).Should(Equal(reconcile.Request{NamespacedName: types.NamespacedName{
				Name:      paint.Name,
				Namespace: paint.Namespace,
			}}))
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
	c.m[cluster] = paint
}

func (c *concurrentPaintMap) get(cluster string) *Paint {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	paint := c.m[cluster]
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
	c.m[cluster] = req
}

func (c *concurrentRequestMap) get(cluster string) reconcile.Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	req := c.m[cluster]
	return req
}
