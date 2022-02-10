package controllerutils_test

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	controller_client "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/resource"
	mock_resource "github.com/solo-io/skv2/pkg/resource/mocks"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		client.EXPECT().Scheme().Return(scheme.Scheme).Times(2)
		client.EXPECT().Get(ctx, resource.ToClientKey(desired), &v1.ConfigMap{}).Return(makeErr(metav1.StatusReasonNotFound))
		client.EXPECT().Create(ctx, desired).Return(nil)

		result, err := Upsert(ctx, client, desired)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(controllerutil.OperationResultCreated))
	})
	It("updates + calls tx funcs when resource is found", func() {
		var called bool

		client.EXPECT().Scheme().Return(scheme.Scheme).Times(2)
		client.EXPECT().Get(ctx, resource.ToClientKey(desired), &v1.ConfigMap{}).Return(nil)
		client.EXPECT().Update(ctx, desired).Return(nil)

		existingTest := desired.DeepCopyObject().(*v1.ConfigMap)

		result, err := Upsert(ctx, client, desired, func(existing, desired runtime.Object) error {
			called = true

			// necessary to ensure there is a diff between existing and desired
			desired.(*v1.ConfigMap).Data = map[string]string{"some": "otherdata"}
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(controllerutil.OperationResultUpdated))
		Expect(called).To(BeTrue())
		// object gets updated by transition function correctly
		Expect(existingTest).ToNot(Equal(desired))
		Expect(desired.Data).To(Equal(map[string]string{"some": "otherdata"}))
	})
})

var _ = Describe("Update Status", func() {
	var (
		client controller_client.Client
		ctx    = context.Background()
		pv     *v1.PersistentVolume
		cl     controller_client.Client
	)
	BeforeEach(func() {
		var err error

		client = fake.NewClientBuilder().Build()
		Expect(err).NotTo(HaveOccurred())
		Expect(client).NotTo(BeNil())
		pv = &v1.PersistentVolume{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "PersistentVolume",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pv",
				Namespace: "ns",
			},
			Status: v1.PersistentVolumeStatus{
				Message: "Test1",
			},
		}
		// This seems to default the resource version to "999"...
		cl = fake.NewClientBuilder().
			WithObjects(pv).
			Build()
		pv.Status = v1.PersistentVolumeStatus{
			Message: "Test2",
		}
	})
	It("updates status when resource is found", func() {
		// update status
		result, err := UpdateStatus(ctx, cl, pv)
		Expect(err).To(BeNil())
		Expect(result).To(Equal(controllerutil.OperationResultUpdated))
		// make sure object was updated
		Expect(pv.ResourceVersion).NotTo(Equal("1000"))
	})
	It("updates status when resource is found; but not the object", func() {
		// update status
		result, err := UpdateStatusImmutable(ctx, cl, pv)
		Expect(err).To(BeNil())
		Expect(result).To(Equal(controllerutil.OperationResultUpdated))
		// make sure object was not updated
		Expect(pv.ResourceVersion).To(Equal("999"))
	})
})

func makeErr(reason metav1.StatusReason) error {
	return &errors.StatusError{
		ErrStatus: metav1.Status{Reason: reason},
	}
}
