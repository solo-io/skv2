package resource_test

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testv1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/resource"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
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
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
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
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
						name: paint,
					},
				},
				cluster2: resource.Snapshot{
					schema.GroupVersionKind{}: map[types.NamespacedName]client.Object{
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
				obj client.Object,
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
				Spec: testv1.PaintSpec{
					Color: &testv1.PaintColor{
						Hue:   "hello",
						Value: 2,
					},
				},
			}
			secret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: name.Namespace,
					Name:      name.Name,
				},
				Data: map[string][]byte{
					"hello": []byte("world"),
				},
			}
			cluster1 := "cluster1"
			cluster2 := "cluster2"
			secretGVK := schema.GroupVersionKind{
				Kind:    "Secret",
				Version: "v1",
			}
			snapshot := resource.ClusterSnapshot{
				cluster1: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
						name: paint,
					},
				},
				cluster2: resource.Snapshot{
					secretGVK: map[types.NamespacedName]client.Object{
						name: secret,
					},
				},
			}

			fullClone := snapshot.Clone()
			paintClone := snapshot.Clone(func(GVK schema.GroupVersionKind) bool {
				return GVK == testv1.PaintGVK
			})
			secretClone := snapshot.Clone(func(GVK schema.GroupVersionKind) bool {
				return GVK == secretGVK
			})

			Expect(deep.Equal(fullClone, snapshot)).To(HaveLen(0))
			Expect(deep.Equal(paintClone, snapshot)).NotTo(HaveLen(0))
			Expect(deep.Equal(secretClone, snapshot)).NotTo(HaveLen(0))

			fullClonePaint, _ := paintClone[cluster1][testv1.PaintGVK][name]
			Expect(fullClonePaint).NotTo(BeNil())
			Expect(controllerutils.ObjectsEqual(fullClonePaint, paint)).To(BeTrue())

			fullCopiedSecret, _ := secretClone[cluster2][secretGVK][name]
			Expect(fullCopiedSecret).NotTo(BeNil())
			Expect(controllerutils.ObjectsEqual(fullCopiedSecret, secret)).To(BeTrue())

			copiedPaint, _ := paintClone[cluster1][testv1.PaintGVK][name]
			Expect(copiedPaint).NotTo(BeNil())
			Expect(controllerutils.ObjectsEqual(copiedPaint, paint)).To(BeTrue())

			copiedPaintNil, _ := paintClone[cluster2][testv1.PaintGVK][name]
			Expect(copiedPaintNil).To(BeNil())

			copiedSecret, _ := secretClone[cluster2][secretGVK][name]
			Expect(copiedSecret).NotTo(BeNil())
			Expect(controllerutils.ObjectsEqual(copiedSecret, secret)).To(BeTrue())

			copiedSecretNil, _ := secretClone[cluster1][secretGVK][name]
			Expect(copiedSecretNil).To(BeNil())

			// Check that object was deep copied
			copiedSecret.(*v1.Secret).Data["hello"] = []byte("seven")
			Expect(controllerutils.ObjectsEqual(copiedSecret, secret)).To(BeFalse())

			copiedPaint.(*testv1.Paint).Spec.Color.Hue = "seven"
			Expect(controllerutils.ObjectsEqual(copiedPaint, paint)).To(BeFalse())
		})

		It("will merge two snapshots properly", func() {
			paint := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{Name: "paint", Namespace: "x"},
			}
			paintOverride := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{Name: "override", Namespace: "x"},
			}
			paint2 := &testv1.Paint{
				ObjectMeta: metav1.ObjectMeta{Name: "paint2", Namespace: "y"},
			}
			name := types.NamespacedName{
				Namespace: "x",
				Name:      "paint",
			}
			name2 := types.NamespacedName{
				Namespace: "y",
				Name:      "paint2",
			}
			cluster1Name := "cluster1"
			cluster2Name := "cluster2"
			leftSnapshot := resource.ClusterSnapshot{
				cluster1Name: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
						name:  paint,
						name2: paint2,
					},
				},
			}
			rightSnapshot := resource.ClusterSnapshot{
				cluster1Name: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
						name: paintOverride,
					},
				},
				cluster2Name: resource.Snapshot{
					testv1.PaintGVK: map[types.NamespacedName]client.Object{
						name: paint,
					},
				},
			}

			snap := leftSnapshot.Merge(rightSnapshot)
			Expect(snap[cluster1Name][testv1.PaintGVK][name].GetName()).
				To(Equal(paintOverride.Name))
			Expect(snap[cluster1Name][testv1.PaintGVK][name2].GetName()).
				To(Equal(paint2.Name))
			Expect(snap[cluster2Name][testv1.PaintGVK][name].GetName()).
				To(Equal(paint.Name))
		})
	})

})
