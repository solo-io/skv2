package tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	. "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
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
		setA = NewPaintSet(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds)
		setB = NewPaintSet(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds)
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
		for _, l := range setA.List() {
			Expect(l).To(Equal(paintA))
		}
		setA.Insert(paintB, paintC)
		for i, l := range setA.List() {
			if i == 0 {
				Expect(l).To(Equal(paintA))
			}
			if i == 1 {
				Expect(l).To(Equal(paintB))
			}
			if i == 2 {
				Expect(l).To(Equal(paintC))
			}
		}
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

	// It("should return set equality", func() {
	// 	setB.Insert(paintA, paintB, paintC)
	// 	setA.Insert(paintA)
	// 	Expect(setA.Equal(setB)).To(BeFalse())
	// 	setA.Insert(paintC, paintB)
	// 	Expect(setA.Equal(setB)).To(BeTrue())
	// })

	// It("should delete", func() {
	// 	setA.Insert(paintA, paintB, paintC)
	// 	Expect(setA.Has(paintA)).To(BeTrue())
	// 	setA.Delete(paintA)
	// 	Expect(setA.Has(paintA)).To(BeFalse())
	// })

	// It("should filter UnsortedList", func() {
	// 	setA.Insert(paintA, paintB, paintC)
	// 	Expect(setA.Has(paintA)).To(BeTrue())
	// 	filtered := setA.UnsortedList(func(p *v1.Paint) bool {
	// 		return p.Name == "nameA"
	// 	})
	// 	Expect(filtered).To(HaveLen(2))
	// 	Expect(filtered).To(ContainElements(paintB, paintC))
	// })

	// It("should double filter UnsortedList", func() {
	// 	setA.Insert(paintA, paintB, paintC)
	// 	Expect(setA.Has(paintA)).To(BeTrue())
	// 	filtered := setA.UnsortedList(func(p *v1.Paint) bool {
	// 		return p.Name == "nameA"
	// 	}, func(p *v1.Paint) bool {
	// 		return p.Name == "nameB"
	// 	})
	// 	Expect(filtered).To(HaveLen(1))
	// 	Expect(filtered).To(ContainElement(paintC))
	// })

	It("should filter List", func() {
		setA.Insert(paintA, paintB, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())

		for i, filtered := range setA.List(func(p *v1.Paint) bool { return p.GetName() == "nameA" }) {
			if i == 1 {
				Expect(filtered).To(Equal(paintB))
			}
			if i == 2 {
				Expect(filtered).To(Equal(paintC))
			}
		}
	})

	It("should double filter List", func() {
		setA.Insert(paintA, paintB, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		for _, filtered := range setA.List(func(p *v1.Paint) bool {
			return p.Name == "nameA"
		}, func(p *v1.Paint) bool {
			return p.Name == "nameB"
		}) {
			Expect(filtered).To(Equal(paintC))
		}
	})

	// It("should union two sets and return new set", func() {
	// 	setA.Insert(paintA, paintB)
	// 	setB.Insert(paintA, paintB, paintC)
	// 	unionSet := setA.Union(setB)
	// 	Expect(unionSet.List()).To(ConsistOf(paintA, paintB, paintC))
	// 	Expect(unionSet).ToNot(BeIdenticalTo(setA))
	// 	Expect(unionSet).ToNot(BeIdenticalTo(setB))
	// })

	// It("should take the difference of two sets and return new set", func() {
	// 	setA.Insert(paintA, paintB)
	// 	setB.Insert(paintA, paintB, paintC)
	// 	differenceA := setA.Difference(setB)
	// 	Expect(differenceA.List()).To(BeEmpty())
	// 	Expect(differenceA.Map()).To(BeEmpty())
	// 	Expect(differenceA).ToNot(BeIdenticalTo(setA))

	// 	differenceB := setB.Difference(setA)
	// 	Expect(differenceB.List()).To(ConsistOf(paintC))
	// 	Expect(differenceB.Map()).To(HaveKeyWithValue(sets.Key(paintC), paintC))
	// 	Expect(differenceB).ToNot(BeIdenticalTo(setB))
	// })

	// It("should take the intersection of two sets and return new set", func() {
	// 	setA.Insert(paintA, paintB)
	// 	setB.Insert(paintA, paintB, paintC)
	// 	intersectionA := setA.Intersection(setB)
	// 	Expect(intersectionA.List()).To(ConsistOf(paintA, paintB))
	// 	Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintA), paintA))
	// 	Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintB), paintB))
	// 	Expect(intersectionA).ToNot(BeIdenticalTo(setA))
	// })

	FIt("should correctly match two sets", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB)
		for _, a := range setA.List() {
			Expect(setB.Has(a)).To(BeTrue())
		}
		setB.Insert(paintC)
		Expect(setA.Length()).ToNot(Equal(setB.Length()))
	})

	It("should return corrent length", func() {
		setA.Insert(paintA, paintB)
		Expect(setA.Length()).To(Equal(2))
	})

	// It("can create a set from a kube list", func() {
	// 	setA.Insert(paintA, paintB)
	// 	setC := NewPaintSetFromList(
	// 		ezkube.ResourceIdsAscending,
	// 		ezkube.CompareResourceIds,
	// 		&v1.PaintList{
	// 			Items: []v1.Paint{*paintA, *paintB},
	// 		},
	// 	)
	// 	Expect(setA.Length()).To(Equal(setC.Length()))
	// 	for i := 0; i < setA.Length(); i++ {
	// 		Expect(setA.List()[i]).To(Equal(setC.List()[i]))
	// 	}
	// })

	// It("should return set deltas", func() {
	// 	// update
	// 	oldPaintA := &v1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{Name: "background", Namespace: "color"},
	// 		Spec:       v1.PaintSpec{Color: &v1.PaintColor{Hue: "orange"}},
	// 	}
	// 	newPaintA := &v1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{Name: "background", Namespace: "color"},
	// 		Spec:       v1.PaintSpec{Color: &v1.PaintColor{Hue: "green"}},
	// 	}
	// 	// remove
	// 	oldPaintB := &v1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{Name: "ugly", Namespace: "color"},
	// 	}
	// 	// add
	// 	newPaintC := &v1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{Name: "beautiful", Namespace: "color"},
	// 	}
	// 	// no change
	// 	oldPaintD := &v1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{Name: "decent", Namespace: "color"},
	// 	}
	// 	newPaintD := &v1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{Name: "decent", Namespace: "color"},
	// 	}
	// 	setA.Insert(oldPaintA, oldPaintB, oldPaintD)
	// 	setB.Insert(newPaintA, newPaintC, newPaintD)

	// 	delta := setA.Delta(setB)
	// 	expectedDelta := sets.ResourceDelta{
	// 		Inserted: sets.NewResourceSet(
	// 			ezkube.ResourceIdsAscending,
	// 			ezkube.CompareResourceIds,
	// 			newPaintA,
	// 			newPaintC,
	// 		),
	// 		Removed: sets.NewResourceSet(
	// 			ezkube.ResourceIdsAscending,
	// 			ezkube.CompareResourceIds,
	// 			oldPaintB,
	// 		),
	// 	}

	// 	// validate removed
	// 	Expect(delta.Removed.Length()).To(Equal(expectedDelta.Removed.Length()))
	// 	for i := 0; i < delta.Removed.Length(); i++ {
	// 		Expect(delta.Removed.List()[i]).To(Equal(expectedDelta.Removed.List()[i]))
	// 	}

	// 	// validate inserted
	// 	Expect(delta.Inserted.Length()).To(Equal(expectedDelta.Inserted.Length()))
	// 	for i := 0; i < delta.Inserted.Length(); i++ {
	// 		Expect(delta.Inserted.List()[i]).To(Equal(expectedDelta.Inserted.List()[i]))
	// 	}
	// })
})
