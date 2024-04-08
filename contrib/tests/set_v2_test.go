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
		setA, setB                     sets_v2.ResourceSet[*v1.Paint]
		paintA, paintBCluster2, paintC *v1.Paint
	)

	BeforeEach(func() {
		setA = sets_v2.NewResourceSet[*v1.Paint]()
		setB = sets_v2.NewResourceSet[*v1.Paint]()
		paintA = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "nameA", Namespace: "nsA"},
		}
		paintBCluster2 = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "nameB", Namespace: "nsB"},
		}
		paintC = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "nameC", Namespace: "nsC"},
		}
	})

	It("should insert", func() {
		setA.Insert(paintA)
		Expect(setA.Has(paintA)).To(BeTrue())
		setA.Insert(paintBCluster2, paintC)
		Expect(setA.Has(paintBCluster2)).To(BeTrue())
		Expect(setA.Has(paintC)).To(BeTrue())
		Expect(setA.Len()).To(Equal(3))
	})

	When("inserting an existing resource", func() {
		It("should overwrite the existing resource", func() {
			setA.Insert(paintA)
			setA.Insert(paintBCluster2)
			setA.Insert(paintA)
			Expect(setA.Len()).To(Equal(2))
		})
	})

	It("should return set existence", func() {
		setA.Insert(paintA)
		Expect(setA.Has(paintA)).To(BeTrue())
		Expect(setA.Has(paintBCluster2)).To(BeFalse())
		setA.Insert(paintBCluster2, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		Expect(setA.Has(paintBCluster2)).To(BeTrue())
		Expect(setA.Has(paintC)).To(BeTrue())
	})

	It("should return set equality", func() {
		setB.Insert(paintA, paintBCluster2, paintC)
		setA.Insert(paintA)
		Expect(setA.Equal(setB)).To(BeFalse())
		setA.Insert(paintC, paintBCluster2)
		Expect(setA.Equal(setB)).To(BeTrue())
	})

	It("should delete", func() {
		setA.Insert(paintA, paintBCluster2, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		setA.Delete(paintA)
		Expect(setA.Len()).To(Equal(2))
		Expect(setA.Has(paintA)).To(BeFalse())
	})

	It("should filter out", func() {
		setA.Insert(paintA, paintBCluster2, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		Expect(setA.Len()).To(Equal(3))

		filterFn := func(p *v1.Paint) bool { return p.GetName() == "nameA" }

		for i, filtered := range setA.FilterOutAndCreateList(filterFn) {
			if i == 0 {
				Expect(filtered).To(Equal(paintBCluster2))
			}
			if i == 1 {
				Expect(filtered).To(Equal(paintC))
			}
		}

	})

	It("should filter", func() {
		setA.Insert(paintA, paintBCluster2, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		Expect(setA.Len()).To(Equal(3))

		filterFn := func(p *v1.Paint) bool { return p.GetName() == "nameA" }

		setA.Filter(filterFn)(func(i int, p *v1.Paint) bool {
			if i == 0 {
				Expect(p).To(Equal(paintA))
			}
			if i != 0 {
				Fail("only one item should be in the filtered set")
			}
			return true
		})
	})

	It("should double filter List", func() {
		setA.Insert(paintA, paintBCluster2, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())
		for _, filtered := range setA.FilterOutAndCreateList(func(p *v1.Paint) bool {
			return p.Name == "nameA"
		}, func(p *v1.Paint) bool {
			return p.Name == "nameB"
		}) {
			Expect(filtered).To(Equal(paintC))
		}
	})

	It("should union two sets and return new set", func() {
		setA.Insert(paintA, paintBCluster2)
		setB.Insert(paintA, paintBCluster2, paintC)
		unionSet := setA.Union(setB)
		Expect(unionSet.Len()).To(Equal(3))
		Expect(unionSet.Has(paintA)).To(BeTrue())
		Expect(unionSet.Has(paintBCluster2)).To(BeTrue())
		Expect(unionSet.Has(paintC)).To(BeTrue())
		Expect(unionSet).ToNot(BeIdenticalTo(setA))
		Expect(unionSet).ToNot(BeIdenticalTo(setB))
	})

	It("should take the difference of two sets and return new set", func() {
		setA.Insert(paintA, paintBCluster2)
		setB.Insert(paintA, paintBCluster2, paintC)
		differenceA := setA.Difference(setB)
		Expect(differenceA.Len()).To(Equal(0))
		Expect(differenceA.Map()).To(BeEmpty())
		Expect(differenceA).ToNot(BeIdenticalTo(setA))

		differenceB := setB.Difference(setA)
		Expect(differenceB.Has(paintC)).To(BeTrue())
		Expect(differenceB.Map()).To(HaveKeyWithValue(sets.Key(paintC), paintC))
		Expect(differenceB).ToNot(BeIdenticalTo(setB))
	})

	It("should take the intersection of two sets and return new set", func() {
		setA.Insert(paintA, paintBCluster2)
		setB.Insert(paintA, paintBCluster2, paintC)
		intersectionA := setA.Intersection(setB)
		Expect(intersectionA.Has(paintA)).To(BeTrue())
		Expect(intersectionA.Has(paintBCluster2)).To(BeTrue())
		Expect(intersectionA.Len()).To(Equal(2))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintA), paintA))
		Expect(intersectionA.Map()).To(HaveKeyWithValue(sets.Key(paintBCluster2), paintBCluster2))
		Expect(intersectionA).ToNot(BeIdenticalTo(setA))
	})

	It("should correctly match two sets", func() {
		setA.Insert(paintA, paintBCluster2)
		setB.Insert(paintA, paintBCluster2)
		Expect(setA.Equal(setB)).To(BeTrue())
		setB.Insert(paintC)
		Expect(setA.Equal(setB)).To(BeFalse())
	})

	It("should return corrent length", func() {
		setA.Insert(paintA, paintBCluster2)
		Expect(setA.Len()).To(Equal(2))
	})

	It("should find resources", func() {
		oldPaintA := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "background",
				Namespace:   "color",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "orange"},
			},
		}
		newPaintA := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "background",
				Namespace:   "color",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "orange"},
			},
		}
		oldPaintB := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "ugly", Namespace: "color"},
		}
		newPaintC := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "beautiful", Namespace: "color"},
		}
		paintList := []*v1.Paint{oldPaintA, newPaintA, oldPaintB, newPaintC, newPaintC}
		expectedPaintList := []*v1.Paint{newPaintC, oldPaintB, newPaintA}
		setA.Insert(paintList...)

		Expect(setA.Len()).To(Equal(3))

		for _, paint := range expectedPaintList {
			found, err := setA.Find(paint)
			Expect(Expect(err).NotTo(HaveOccurred()))
			Expect(found).To(Equal(paint))
		}
		// find based on something that implemented ClusterObjectRef
		found, err := setA.Find(ezkube.MakeClusterObjectRef(newPaintC))
		Expect(Expect(err).NotTo(HaveOccurred()))
		Expect(found).To(Equal(newPaintC))
	})

	It("should sort resources by name.ns.cluster", func() {
		paintAA1 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "a",
				Namespace:   "a",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "1"},
			},
		}
		paintAB1 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "a",
				Namespace:   "b",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "1"},
			},
		}
		paintBB1 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "b",
				Namespace:   "b",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "1"},
			},
		}
		paintAC2 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "a",
				Namespace:   "c",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "2"},
			},
		}
		paintBC2 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "b",
				Namespace:   "c",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "2"},
			},
		}
		paintAC := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "a",
				Namespace: "c",
			},
		}
		paintBC := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "b",
				Namespace: "c",
			},
		}
		expectedOrder := []*v1.Paint{paintAA1, paintAB1, paintAC, paintAC2, paintBB1, paintBC, paintBC2}
		setA.Insert(expectedOrder...)

		var paintList []*v1.Paint
		var paintList2 []*v1.Paint
		for _, paint := range setA.FilterOutAndCreateList() {
			paintList = append(paintList, paint)
		}

		setA.Iter(func(_ int, paint *v1.Paint) bool {
			paintList2 = append(paintList2, paint)
			return true
		})

		for i, paint := range expectedOrder {
			Expect(paintList[i]).To(Equal(paint))
			Expect(paintList2[i]).To(Equal(paint))
		}
	})
})
