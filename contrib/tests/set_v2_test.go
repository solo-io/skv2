package tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	sets_v2 "github.com/solo-io/skv2/contrib/pkg/sets/v2"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = FDescribe("PaintSetV2", func() {
	var (
		setA, setB             sets_v2.ResourceSet[*v1.Paint]
		paintA, paintB, paintC *v1.Paint
	)

	BeforeEach(func() {
		setA = sets_v2.NewResourceSet[*v1.Paint](ezkube.ObjectsAscending, ezkube.CompareObjects)
		setB = sets_v2.NewResourceSet[*v1.Paint](ezkube.ObjectsAscending, ezkube.CompareObjects)
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
		Expect(setA.Has(paintA)).To(BeTrue())
		setA.Insert(paintB, paintC)
		Expect(setA.Has(paintB)).To(BeTrue())
		Expect(setA.Has(paintC)).To(BeTrue())
		Expect(setA.Length()).To(Equal(3))
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

	It("should union two sets and return new set", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB, paintC)
		unionSet := setA.Union(setB)
		Expect(unionSet.Length()).To(Equal(3))
		Expect(unionSet.Has(paintA)).To(BeTrue())
		Expect(unionSet.Has(paintB)).To(BeTrue())
		Expect(unionSet.Has(paintC)).To(BeTrue())
		Expect(unionSet).ToNot(BeIdenticalTo(setA))
		Expect(unionSet).ToNot(BeIdenticalTo(setB))
	})

	It("should take the difference of two sets and return new set", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB, paintC)
		differenceA := setA.Difference(setB)
		Expect(differenceA.Length()).To(Equal(0))
		Expect(differenceA.Map()).To(BeEmpty())
		Expect(differenceA).ToNot(BeIdenticalTo(setA))

		differenceB := setB.Difference(setA)
		Expect(differenceB.Has(paintC)).To(BeTrue())
		Expect(differenceB.Map()).To(HaveKeyWithValue(sets.Key(paintC), paintC))
		Expect(differenceB).ToNot(BeIdenticalTo(setB))
	})

	It("should take the intersection of two sets and return new set", func() {
		setA.Insert(paintA, paintB)
		setB.Insert(paintA, paintB, paintC)
		intersectionA := setA.Intersection(setB)
		Expect(intersectionA.Has(paintA)).To(BeTrue())
		Expect(intersectionA.Has(paintB)).To(BeTrue())
		Expect(intersectionA.Length()).To(Equal(2))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintA), paintA))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintB), paintB))
		Expect(intersectionA).ToNot(BeIdenticalTo(setA))
	})

	// It("should correctly match two sets", func() {
	// 	setA.Insert(paintA, paintB)
	// 	setB.Insert(paintA, paintB)
	// 	Expect(setA).To(Equal(setB))
	// 	setB.Insert(paintC)
	// 	Expect(setA).ToNot(Equal(setB))
	// })

	It("should return corrent length", func() {
		setA.Insert(paintA, paintB)
		Expect(setA.Length()).To(Equal(2))
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

		expectedDelta := sets_v2.ResourceDelta[*v1.Paint]{
			Inserted: sets_v2.NewResourceSet(ezkube.ObjectsAscending, ezkube.CompareObjects, newPaintA, newPaintC),
			Removed:  sets_v2.NewResourceSet(ezkube.ObjectsAscending, ezkube.CompareObjects, oldPaintB),
		}

		actualDelta := setA.Delta(setB)
		Expect(actualDelta.Inserted.Length()).To(Equal(expectedDelta.DeltaV1().Inserted.Length()))
	})
})
