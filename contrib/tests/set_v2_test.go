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
		setA = sets_v2.NewResourceSet[*v1.Paint](ezkube.ResourceIdsCompare)
		setB = sets_v2.NewResourceSet[*v1.Paint](ezkube.ResourceIdsCompare)
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
		Expect(setA.Has(paintA)).To(BeFalse())
	})

	It("should filter List", func() {
		setA.Insert(paintA, paintBCluster2, paintC)
		Expect(setA.Has(paintA)).To(BeTrue())

		for i, filtered := range setA.List(func(p *v1.Paint) bool { return p.GetName() == "nameA" }) {
			if i == 1 {
				Expect(filtered).To(Equal(paintBCluster2))
			}
			if i == 2 {
				Expect(filtered).To(Equal(paintC))
			}
		}
	})

	It("should double filter List", func() {
		setA.Insert(paintA, paintBCluster2, paintC)
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

	// It("should correctly match two sets", func() {
	// 	setA.Insert(paintA, paintBCluster2)
	// 	setB.Insert(paintA, paintBCluster2)
	// 	Expect(setA).To(Equal(setB))
	// 	setB.Insert(paintC)
	// 	Expect(setA).ToNot(Equal(setB))
	// })

	It("should return corrent length", func() {
		setA.Insert(paintA, paintBCluster2)
		Expect(setA.Len()).To(Equal(2))
	})

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

	// 	expectedDelta := sets_v2.ResourceDelta[*v1.Paint]{
	// 		Inserted: sets_v2.NewResourceSet(ezkube.ObjectsAscending, ezkube.CompareObjects, newPaintA, newPaintC),
	// 		Removed:  sets_v2.NewResourceSet(ezkube.ObjectsAscending, ezkube.CompareObjects, oldPaintB),
	// 	}

	// 	actualDelta := setA.Delta(setB)

	// 	Expect(actualDelta.Removed.Length()).To(Equal(expectedDelta.DeltaV1().Removed.Length()))
	// 	for _, removed := range actualDelta.Removed.List() {
	// 		Expect(removed).To(Equal(oldPaintB))
	// 	}

	// 	Expect(actualDelta.Inserted.Length()).To(Equal(expectedDelta.DeltaV1().Inserted.Length()))
	// 	for i, inserted := range actualDelta.Inserted.List() {
	// 		if i == 0 {
	// 			Expect(inserted).To(Equal(newPaintA))
	// 		}
	// 		if i == 1 {
	// 			Expect(inserted).To(Equal(newPaintC))
	// 		}
	// 	}
	// })

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
		// remove
		oldPaintB := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "ugly", Namespace: "color"},
		}
		// add
		newPaintC := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{Name: "beautiful", Namespace: "color"},
		}
		paintList := []*v1.Paint{oldPaintA, newPaintA, oldPaintB, newPaintC}
		setA.Insert(paintList...)

		for _, paint := range paintList {
			found, err := setA.Find(paint)
			Expect(Expect(err).NotTo(HaveOccurred()))
			Expect(found).To(Equal(paint))
		}
	})

	It("should sort resources first by cluster, then by namespace, then by name", func() {
		paintAAcluster1 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "a",
				Namespace:   "a",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "cluster-1"},
			},
		}
		paintABcluster1 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "a",
				Namespace:   "b",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "cluster-1"},
			},
		}
		paintBBcluster1 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "b",
				Namespace:   "b",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "cluster-1"},
			},
		}
		paintACluster2 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "a",
				Namespace:   "c",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "cluster-2"},
			},
		}
		paintBCluster2 := &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "b",
				Namespace:   "c",
				Annotations: map[string]string{ezkube.ClusterAnnotation: "cluster-2"},
			},
		}
		expectedOrder := []*v1.Paint{
			paintAAcluster1, paintABcluster1, paintBBcluster1,
			paintACluster2, paintBCluster2,
		}
		setA.Insert(expectedOrder...)

		var paintList []*v1.Paint
		for _, paint := range setA.List() {
			paintList = append(paintList, paint)
		}

		for i, paint := range expectedOrder {
			Expect(paintList[i]).To(Equal(paint))
		}
	})
})
