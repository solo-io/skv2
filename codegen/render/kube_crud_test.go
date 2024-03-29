package render_test

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/solo-io/skv2/pkg/reconcile"
	"github.com/solo-io/skv2/test/matchers"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/randutils"
	"github.com/solo-io/k8s-utils/kubeutils"
	kubehelp "github.com/solo-io/k8s-utils/testutils/kube"
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

// TODO (dmitri-d): Remove dependency of these tests on crds generated in cmd_test
func applyFile(file string, extraArgs ...string) error {
	path := filepath.Join(util.GetModuleRoot(), "codegen/test/chart/crds", file)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return util.KubectlApply(b, extraArgs...)
}

func newPaint(namespace, name string) *Paint {
	serializedAnyValue, _ := proto.Marshal(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"my": {Kind: &structpb.Value_StringValue{StringValue: "favorite-color"}},
		},
	})

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
			MyFavorite: &any.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
				Value:   serializedAnyValue,
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

	Expect(&written.Spec).To(matchers.MatchProto(&paint.Spec))

	status := PaintStatus{
		ObservedGeneration: written.Generation,
		PercentRemaining:   22,
	}

	written.Status = status

	err = clientSet.Paints().UpdatePaintStatus(ctx, written)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() *PaintStatus {
		written, err = clientSet.Paints().GetPaint(ctx, client.ObjectKey{
			Namespace: paint.Namespace,
			Name:      paint.Name,
		})
		Expect(err).NotTo(HaveOccurred())
		return &written.Status
	}, time.Second).Should(matchers.MatchProto(&status))
}

var _ = Describe("Generated Code", func() {
	var (
		ctx       context.Context
		cancel    context.CancelFunc
		ns        string
		kube      kubernetes.Interface
		clientSet Clientset
		logLevel  = zap.NewAtomicLevel()
	)
	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		logLevel.SetLevel(zap.DebugLevel)
		log.SetLogger(zaputil.New(
			zaputil.Level(&logLevel),
		))
		err := applyFile("things.test.io_crds.yaml")

		Expect(err).NotTo(HaveOccurred(), "failed to apply crds")
		ns = randutils.RandString(4)
		kube = kubehelp.MustKubeClient()
		err = kubeutils.CreateNamespacesInParallel(ctx, kube, ns)
		Expect(err).NotTo(HaveOccurred())
		clientSet, err = NewClientsetFromConfig(test.MustConfig(""))
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		err := kubeutils.DeleteNamespacesInParallelBlocking(ctx, kube, ns)
		Expect(err).NotTo(HaveOccurred())
		cancel()
	})

	Context("kube clientsets", func() {
		It("uses the generated clientsets to crud", func() {

			paint := newPaint(ns, "paint-kube-clientsets")

			mustCrud(ctx, clientSet, paint)
		})
	})

	Context("kube reconciler", func() {
		var (
			mgr manager.Manager
		)
		BeforeEach(func() {
			mgr = test.MustManager(ctx, ns)
		})

		It("uses the generated controller to reconcile", func() {

			paint := newPaint(ns, "paint-kube-reconciler")

			err := clientSet.Paints().CreatePaint(ctx, paint)
			Expect(err).NotTo(HaveOccurred())

			paint.GetObjectKind().GroupVersionKind()

			loop := controller.NewPaintReconcileLoop("blick", mgr, reconcile.Options{})

			var reconciled *Paint
			var deleted reconcile.Request
			reconciler := &controller.PaintReconcilerFuncs{
				OnReconcilePaint: func(obj *Paint) (result reconcile.Result, err error) {
					reconciled = obj
					return
				},
				OnReconcilePaintDeletion: func(req reconcile.Request) error {
					deleted = req
					return nil
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

			Eventually(func() *PaintSpec {
				return &reconciled.Spec
			}, time.Second).Should(matchers.MatchProto(&paint.Spec))

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

	Context("recursive fields in kube api", func() {

		It("uses the generated controller to reconcile", func() {

			paint := newPaint(ns, "paint-kube-reconciler")

			// make the paint recursive
			recursiveType := &PaintSpec_RecursiveType{
				RecursiveField: &PaintSpec_RecursiveType{
					RecursiveField: &PaintSpec_RecursiveType{},
					RepeatedRecursiveField: []*PaintSpec_RecursiveType{
						{},
						{},
					},
				},
				RepeatedRecursiveField: []*PaintSpec_RecursiveType{
					{
						RecursiveField: &PaintSpec_RecursiveType{},
						RepeatedRecursiveField: []*PaintSpec_RecursiveType{
							{},
							{},
						},
					},
					{
						RecursiveField: &PaintSpec_RecursiveType{},
						RepeatedRecursiveField: []*PaintSpec_RecursiveType{
							{},
							{},
						},
					},
				},
			}
			paint.Spec.RecursiveType = recursiveType

			err := clientSet.Paints().CreatePaint(ctx, paint)
			Expect(err).NotTo(HaveOccurred())

			// expect the server accepted our spec
			Expect(paint.Spec.RecursiveType).To(Equal(recursiveType))
		})
	})
})
