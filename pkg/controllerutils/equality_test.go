package controllerutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/solo-io/skv2/pkg/controllerutils"
)

var _ = Describe("Equality", func() {
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
