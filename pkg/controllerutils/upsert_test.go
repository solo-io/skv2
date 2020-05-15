package controllerutils_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/resource"
	mock_resource "github.com/solo-io/skv2/pkg/resource/mocks"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

var _ = Describe("Upsert", func() {
	var (
		desired *v1.ConfigMap
		ctl     *gomock.Controller
		client  *mock_resource.MockClient
		ctx     = context.Background()
	)
	BeforeEach(func() {
		desired = &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "ns",
				Name:      "name",
			},
			Data: map[string]string{"some": "data"},
		}
		ctl = gomock.NewController(GinkgoT())
		client = mock_resource.NewMockClient(ctl)
	})
	AfterEach(func() {
		ctl.Finish()
	})
	It("creates when resource is not found", func() {
		client.EXPECT().Get(ctx, resource.ToClientKey(desired), desired).Return(makeErr(metav1.StatusReasonNotFound))
		client.EXPECT().Create(ctx, desired).Return(nil)

		result, err := Upsert(ctx, client, desired)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(controllerutil.OperationResultCreated))
	})
	It("updates + calls tx funcs when resource is found", func() {
		var called bool

		client.EXPECT().Get(ctx, resource.ToClientKey(desired), desired).Return(nil)
		client.EXPECT().Update(ctx, desired).Return(nil)

		result, err := Upsert(ctx, client, desired, func(existing, desired runtime.Object) error {
			called = true

			// necessary to ensure there is a diff between existing and desired
			existing.(*v1.ConfigMap).Data = nil
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(controllerutil.OperationResultUpdated))
		Expect(called).To(BeTrue())
	})
})

func makeErr(reason metav1.StatusReason) error {
	return &errors.StatusError{
		ErrStatus: metav1.Status{Reason: reason},
	}
}
