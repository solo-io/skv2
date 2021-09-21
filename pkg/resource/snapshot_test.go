package resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testv1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/pkg/resource"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Snapshot", func() {

	Context("ClusterSnapshot", func() {

		It("will do nothing if GVK doesn't exist", func() {
			paint := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{},
			}
			name := types.NamespacedName{
				Namespace: "a",
				Name:      "b",
			}
			clusterName := "cluster"
			snapshot := resource.ClusterSnapshot{
				clusterName: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]resource.TypedObject{
						name: paint,
					},
				},
			}

			snapshot.Delete(clusterName, schema.GroupVersionKind{}, name)

			paintMap, _ := snapshot[clusterName][testv1.PaintGVK]
			Expect(paintMap[name]).NotTo(BeNil())
		})

		It("will delete item if exists in GVK map", func() {
			paint := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{},
			}
			name := types.NamespacedName{
				Namespace: "a",
				Name:      "b",
			}
			clusterName := "cluster"
			snapshot := resource.ClusterSnapshot{
				clusterName: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]resource.TypedObject{
						name: paint,
					},
				},
			}

			snapshot.Delete(clusterName, testv1.PaintGVK, name)

			paintMap, _ := snapshot[clusterName][testv1.PaintGVK]
			Expect(paintMap[name]).To(BeNil())
		})

		It("will insert properly", func() {

			paint := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "a",
					Name:      "b",
				},
			}
			name := types.NamespacedName{
				Namespace: "a",
				Name:      "b",
			}

			clusterName := "cluster"
			snapshot := resource.ClusterSnapshot{}

			snapshot.Insert(clusterName, testv1.PaintGVK, paint)

			paintMap, _ := snapshot[clusterName][testv1.PaintGVK]
			Expect(paintMap[name]).To(Equal(paint))

		})

		It("will list each object", func() {
			name := types.NamespacedName{
				Namespace: "a",
				Name:      "b",
			}
			paint := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: name.Namespace,
					Name:      name.Name,
				},
			}
			secret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: name.Namespace,
					Name:      name.Name,
				},
			}
			cluster1 := "cluster1"
			cluster2 := "cluster2"
			snapshot := resource.ClusterSnapshot{
				cluster1: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]resource.TypedObject{
						name: paint,
					},
				},
				cluster2: resource.Snapshot{
					schema.GroupVersionKind{}: map[types.NamespacedName]resource.TypedObject{
						name: secret,
					},
				},
			}

			var (
				seenSecret, seenPaint bool
			)
			snapshot.ForEachObject(func(
				cluster string,
				gvk schema.GroupVersionKind,
				obj resource.TypedObject,
			) {
				if expected, ok := obj.(*v1.Secret); ok {
					Expect(secret).To(Equal(expected))
					seenSecret = true
				}
				if expected, ok := obj.(*testv1.Paint); ok {
					Expect(paint).To(Equal(expected))
					seenPaint = true
				}
			})
			Expect(seenPaint).To(BeTrue())
			Expect(seenSecret).To(BeTrue())
		})
	})

	// It("will list clone properly", func() {
	// 	name := types.NamespacedName{
	// 		Namespace: "a",
	// 		Name:      "b",
	// 	}
	// 	paint := &testv1.Paint{
	// 		ObjectMeta: metav1.ObjectMeta{
	// 			Namespace: name.Namespace,
	// 			Name:      name.Name,
	// 		},
	// 	}
	// 	snapshot := resource.Snapshot{
	// 		testv1.PaintGVK: map[types.NamespacedName]resource.TypedObject{
	// 			name: paint,
	// 		},
	// 	}
	//
	// 	emptySnap := snapshot.Clone()
	//
	// 	var (
	// 		seenSecret, seenPaint bool
	// 	)
	// 	snapshot.ForEachObject(func(gvk schema.GroupVersionKind, obj resource.TypedObject) {
	// 		if expected, ok := obj.(*v1.Secret); ok {
	// 			Expect(secret).To(Equal(expected))
	// 			seenSecret = true
	// 		}
	// 		if expected, ok := obj.(*testv1.Paint); ok {
	// 			Expect(paint).To(Equal(expected))
	// 			seenPaint = true
	// 		}
	// 	})
	// 	Expect(seenPaint).To(BeTrue())
	// 	Expect(seenSecret).To(BeTrue())
	// })

})
