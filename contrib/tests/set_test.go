package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	. "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/sets"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("PaintSet", func() {
	var (
		setA   PaintSet
		setB   PaintSet
		paintA *v1.Paint
		paintB *v1.Paint
		paintC *v1.Paint
	)

	BeforeEach(func() {
		setA = NewPaintSet()
		setB = NewPaintSet()
		paintA = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "nameA", Namespace: "nsA"},
		}
		paintB = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "nameB", Namespace: "nsB"},
		}
		paintC = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "nameC", Namespace: "nsC"},
		}
	})

	It("should insert", func() {
		setA.Insert(paintA)
		list := setA.List()
		Expect(list).To(ConsistOf(paintA))
		setA.Insert(paintB, paintC)
		list = setA.List()
		Expect(list).To(ConsistOf(paintA, paintB, paintC))
	})

	It("should return set existence", func() {
		setA.Insert(paintA)
		Expect(setA.Has(paintA)).To(BeTrue())
		Expect(setA.Has(paintB)).To(BeFalse())
		setA.Insert(paintB, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		Expect(setA.Has(paintB)).To(BeTrue())
		Expect(setA.Has(paintC)).To(BeTrue())
	})

	It("should return set equality", func() {
		setB.Insert(paintA, paintB, paintC)
		setA.Insert(paintA)
		Expect(setA.Equal(setB)).To(BeFalse())
		setA.Insert(paintC, paintB)
		Expect(setA.Equal(setB)).To(BeTrue())
	})

	It("should delete", func() {
		setA.Insert(paintA, paintB, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		setA.Delete(paintA)
		Expect(setA.Has(paintA)).To(BeFalse())
	})

	It("should union two sets and return new set", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB, paintC)
		unionSet := setA.Union(setB)
		Expect(unionSet.List()).To(ConsistOf(paintA, paintB, paintC))
		Expect(unionSet).ToNot(BeIdenticalTo(setA))
		Expect(unionSet).ToNot(BeIdenticalTo(setB))
	})

	It("should take the difference of two sets and return new set", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB, paintC)
		differenceA := setA.Difference(setB)
		Expect(differenceA.List()).To(BeEmpty())
		Expect(differenceA.Map()).To(BeEmpty())
		Expect(differenceA).ToNot(BeIdenticalTo(setA))

		differenceB := setB.Difference(setA)
		Expect(differenceB.List()).To(ConsistOf(paintC))
		Expect(differenceB.Map()).To(HaveKeyWithValue(sets.Key(paintC), paintC))
		Expect(differenceB).ToNot(BeIdenticalTo(setB))
	})

	It("should take the intersection of two sets and return new set", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB, paintC)
		intersectionA := setA.Intersection(setB)
		Expect(intersectionA.List()).To(ConsistOf(paintA, paintB))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintA), paintA))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintB), paintB))
		Expect(intersectionA).ToNot(BeIdenticalTo(setA))
	})

	It("should correctly match two sets", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB)
		Expect(setA).To(Equal(setB))
		setB.Insert(paintC)
		Expect(setA).ToNot(Equal(setB))
	})

	It("should return corrent length", func() {
		setA.Insert(paintA, paintB)
		Expect(setA.Length()).To(Equal(2))
	})

	It("can create a set from a kube list", func() {
		setA.Insert(paintA, paintB)
		setC := NewPaintSetFromList(&v1.PaintList{
			Items: []v1.Paint{*paintA, *paintB},
		})
		Expect(setA).To(Equal(setC))

	})
	It("should return set deltas", func() {
		// update
		oldPaintA := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "background", Namespace: "color"},
			Spec:       v1.PaintSpec{Color: &v1.PaintColor{Hue: "orange"}},
		}
		newPaintA := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "background", Namespace: "color"},
			Spec:       v1.PaintSpec{Color: &v1.PaintColor{Hue: "green"}},
		}
		// remove
		oldPaintB := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "ugly", Namespace: "color"},
		}
		// add
		newPaintC := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "beautiful", Namespace: "color"},
		}
		// no change
		oldPaintD := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "decent", Namespace: "color"},
		}
		newPaintD := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "decent", Namespace: "color"},
		}
		setA.Insert(oldPaintA, oldPaintB, oldPaintD)
		setB.Insert(newPaintA, newPaintC, newPaintD)
		Expect(setA.Delta(setB)).To(Equal(sets.ResourceDelta{
			Inserted: sets.NewResourceSet(newPaintA, newPaintC),
			Removed:  sets.NewResourceSet(oldPaintB),
		}))
	})
})
