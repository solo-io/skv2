package v1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("ServiceSet", func() {
	var (
		setA     ServiceSet
		setB     ServiceSet
		serviceA *corev1.Service
		serviceB *corev1.Service
		serviceC *corev1.Service
	)

	BeforeEach(func() {
		setA = NewServiceSet()
		setB = NewServiceSet()
		serviceA = &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "nameA", Namespace: "nsA"},
		}
		serviceB = &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "nameB", Namespace: "nsB"},
		}
		serviceC = &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "nameC", Namespace: "nsC"},
		}
	})

	It("should insert", func() {
		setA.Insert(serviceA)
		list := setA.List()
		Expect(list).To(ConsistOf(serviceA))
		setA.Insert(serviceB, serviceC)
		list = setA.List()
		Expect(list).To(ConsistOf(serviceA, serviceB, serviceC))
	})

	It("should return set existence", func() {
		setA.Insert(serviceA)
		Expect(setA.Has(serviceA)).To(BeTrue())
		Expect(setA.Has(serviceB)).To(BeFalse())
		setA.Insert(serviceB, serviceC)
		Expect(setA.Has(serviceA)).To(BeTrue())
		Expect(setA.Has(serviceB)).To(BeTrue())
		Expect(setA.Has(serviceC)).To(BeTrue())
	})

	It("should return set equality", func() {
		setB.Insert(serviceA, serviceB, serviceC)
		setA.Insert(serviceA)
		Expect(setA.Equal(setB)).To(BeFalse())
		setA.Insert(serviceC, serviceB)
		Expect(setA.Equal(setB)).To(BeTrue())
	})

	It("should delete", func() {
		setA.Insert(serviceA, serviceB, serviceC)
		Expect(setA.Has(serviceA)).To(BeTrue())
		setA.Delete(serviceA)
		Expect(setA.Has(serviceA)).To(BeFalse())
	})

	It("should union two sets and return new set", func() {
		setA.Insert(serviceA, serviceB)
		setB.Insert(serviceA, serviceB, serviceC)
		unionSet := setA.Union(setB)
		Expect(unionSet.List()).To(ConsistOf(serviceA, serviceB, serviceC))
		Expect(unionSet).ToNot(BeIdenticalTo(setA))
		Expect(unionSet).ToNot(BeIdenticalTo(setB))
	})

	It("should take the difference of two sets and return new set", func() {
		setA.Insert(serviceA, serviceB)
		setB.Insert(serviceA, serviceB, serviceC)
		differenceA := setA.Difference(setB)
		Expect(differenceA.List()).To(BeEmpty())
		Expect(differenceA.Map()).To(BeEmpty())
		Expect(differenceA).ToNot(BeIdenticalTo(setA))

		differenceB := setB.Difference(setA)
		Expect(differenceB.List()).To(ConsistOf(serviceC))
		Expect(differenceB.Map()).To(HaveKeyWithValue(key(serviceC.ObjectMeta), serviceC))
		Expect(differenceB).ToNot(BeIdenticalTo(setB))
	})

	It("should take the intersection of two sets and return new set", func() {
		setA.Insert(serviceA, serviceB)
		setB.Insert(serviceA, serviceB, serviceC)
		intersectionA := setA.Intersection(setB)
		Expect(intersectionA.List()).To(ConsistOf(serviceA, serviceB))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(key(serviceA.ObjectMeta), serviceA))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(key(serviceB.ObjectMeta), serviceB))
		Expect(intersectionA).ToNot(BeIdenticalTo(setA))
	})

	It("should correctly match two sets", func() {
		setA.Insert(serviceA, serviceB)
		setB.Insert(serviceA, serviceB)
		Expect(setA).To(Equal(setB))
		setB.Insert(serviceC)
		Expect(setA).ToNot(Equal(setB))
	})
})
