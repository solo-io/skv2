package controllerutils_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	v1 "k8s.io/api/core/v1"
	rbac_authorization_k8s_io_v1 "k8s.io/api/rbac/v1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/solo-io/skv2/pkg/controllerutils"
)

var _ = Describe("ObjectsEqual", func() {
	It("asserts equality on two objects even if their resource versions differ", func() {
		obj1 := &v1.ConfigMap{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:            "name",
				Namespace:       "ns",
				Labels:          map[string]string{"a": "b"},
				Annotations:     map[string]string{"c": "d"},
				ResourceVersion: "should-be",
			},
			Data: map[string]string{"some": "data"},
		}
		obj2 := &v1.ConfigMap{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:            "name",
				Namespace:       "ns",
				Labels:          map[string]string{"a": "b"},
				Annotations:     map[string]string{"c": "d"},
				ResourceVersion: "ignored",
			},
			Data: map[string]string{"some": "data"},
		}
		equal := ObjectsEqual(obj1, obj2)
		Expect(equal).To(BeTrue())
	})
	It("asserts difference on two objects if their labels differ", func() {

		obj1 := &v1.ConfigMap{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:        "name",
				Namespace:   "ns",
				Labels:      map[string]string{"a": "b"},
				Annotations: map[string]string{"c": "d"},
			},
			Data: map[string]string{"some": "data"},
		}
		obj2 := &v1.ConfigMap{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:        "name",
				Namespace:   "ns",
				Labels:      map[string]string{"a": "different"},
				Annotations: map[string]string{"c": "d"},
			},
			Data: map[string]string{"some": "data"},
		}
		equal := ObjectsEqual(obj1, obj2)
		Expect(equal).To(BeFalse())
	})
	It("asserts difference on two objects if their non-metadata fields", func() {

		obj1 := &v1.ConfigMap{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:        "name",
				Namespace:   "ns",
				Labels:      map[string]string{"a": "b"},
				Annotations: map[string]string{"c": "d"},
			},
			Data: map[string]string{"some": "data"},
		}
		obj2 := &v1.ConfigMap{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:        "name",
				Namespace:   "ns",
				Labels:      map[string]string{"a": "b"},
				Annotations: map[string]string{"c": "d"},
			},
			Data: map[string]string{"some": "thingelse"},
		}
		equal := ObjectsEqual(obj1, obj2)
		Expect(equal).To(BeFalse())
	})
	It("asserts equality on two proto.Message objects (protov2, with unknown fields) even if they are passed in by reference", func() {
		obj1 := &things_test_io_v1.Paint{
			Spec: things_test_io_v1.PaintSpec{
				Color: &things_test_io_v1.PaintColor{
					Hue: "red",
				},
				PaintType: &things_test_io_v1.PaintSpec_Oil{
					Oil: nil,
				},
			},
		}
		obj1.Spec.ProtoReflect().SetUnknown([]byte(""))
		obj2 := &things_test_io_v1.Paint{
			Spec: things_test_io_v1.PaintSpec{
				Color: &things_test_io_v1.PaintColor{
					Hue: "red",
				},
				PaintType: &things_test_io_v1.PaintSpec_Oil{
					Oil: nil,
				},
			},
		}
		equal := ObjectsEqual(obj1, obj2)
		Expect(equal).To(BeTrue())
	})
	It("asserts equality on two proto.Message objects even if they are 'fake' (i.e., k8s protos implementing github protov1, not google protov2 protos)", func() {
		obj1 := &rbac_authorization_k8s_io_v1.RoleBinding{
			TypeMeta: k8s_meta.TypeMeta{},
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:      "sample-role-binding",
				Namespace: "default",
				Labels:    map[string]string{"k": "v"},
			},
			Subjects: []rbac_authorization_k8s_io_v1.Subject{{
				Kind:      "ServiceAccount",
				Name:      "sample-target",
				Namespace: "default",
			}},
			RoleRef: rbac_authorization_k8s_io_v1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
				Name:     "sample-controller",
			},
		}
		obj2 := &rbac_authorization_k8s_io_v1.RoleBinding{
			TypeMeta: k8s_meta.TypeMeta{},
			ObjectMeta: k8s_meta.ObjectMeta{
				Name:      "sample-role-binding",
				Namespace: "default",
				Labels:    map[string]string{"k": "v"},
			},
			Subjects: []rbac_authorization_k8s_io_v1.Subject{{
				Kind:      "ServiceAccount",
				Name:      "sample-target",
				Namespace: "default",
			}},
			RoleRef: rbac_authorization_k8s_io_v1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
				Name:     "sample-controller",
			},
		}
		equal := ObjectsEqual(obj1, obj2)
		Expect(equal).To(BeTrue())
	})
})

var _ = Describe("ObjectStatusesEqual", func() {
	It("asserts equality on two objects based on their status", func() {
		obj1 := &things_test_io_v1.Paint{
			Status: things_test_io_v1.PaintStatus{
				PercentRemaining: 50,
			},
		}
		obj2 := &things_test_io_v1.Paint{
			Status: things_test_io_v1.PaintStatus{
				PercentRemaining: 50,
			},
		}
		equal := ObjectStatusesEqual(obj1, obj2)
		Expect(equal).To(BeTrue())

		obj2.Status.PercentRemaining = 5
		equal = ObjectStatusesEqual(obj1, obj2)
		Expect(equal).To(BeFalse())
	})

	It("doesnt fail with deep copy in paralel", func() {
		// this was discovered by race detector
		// and this test will only fail with race detector on.
		obj1 := &things_test_io_v1.Paint{
			Status: things_test_io_v1.PaintStatus{
				PercentRemaining: 50,
			},
		}
		obj2 := &things_test_io_v1.Paint{
			Status: things_test_io_v1.PaintStatus{
				PercentRemaining: 50,
			},
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() {
			for ctx.Err() == nil {
				ObjectsEqual(obj1, obj2)
			}
		}()
		go func() {
			for ctx.Err() == nil {
				obj1.DeepCopyObject()
				obj2.DeepCopyObject()
			}
		}()

		time.Sleep(time.Second)
		// if we got here without race detector crash, we're good
	})
})
