package controllerutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	v1 "k8s.io/api/core/v1"
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
})
