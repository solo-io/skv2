package kuberesource_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	. "github.com/solo-io/skv2/codegen/kuberesource"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/pkg/crdutils"
)

var _ = Describe("Crd", func() {

	Describe("CRD gen", func() {
		var (
			grps []*model.Group
		)

		BeforeEach(func() {
			grps = []*model.Group{{
				Resources: []model.Resource{
					{
						Kind: "kind",
						Spec: model.Field{
							Type: model.Type{
								Name:    "test",
								Message: &v1.AcrylicType{},
							},
						},
						Stored:     true,
						Deprecated: false,
					},
					{
						Kind: "kind-1",
						Spec: model.Field{
							Type: model.Type{
								Name:    "test",
								Message: &v1.PaintColor{},
							},
						},
						Stored:     false,
						Deprecated: true,
					},
				}},
			}
			for i := range grps {
				grps[i].Init()
			}
		})

		It("should generate spec hash", func() {
			grps[0].SkipSpecHash = false
			o, err := CustomResourceDefinitions(grps)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(2))
			// note: we intentionally provide the "b6ec737002f7d02e" hash in the test, as it shouldn't change
			// between runs.
			Expect(o[0].GetAnnotations()).To(HaveKeyWithValue(crdutils.CRDSpecHashKey, "b6ec737002f7d02e"))
		})
		It("should not generate spec hash", func() {
			grps[0].SkipSpecHash = true
			o, err := CustomResourceDefinitions(grps)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(2))
			// note: we intentionally provide the "d18828e563010e32" hash in the test, as it shouldn't change
			// between runs.
			Expect(o[0].GetAnnotations()).NotTo(HaveKey(crdutils.CRDSpecHashKey))
		})
		It("should set 'Stored' and 'Deprecated' fields", func() {
			grps[0].SkipSpecHash = false
			o, err := CustomResourceDefinitions(grps)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(2))
			Expect(o[0].Spec.Versions).To(HaveLen(1))
			Expect(o[0].Spec.Versions[0].Storage).To(BeTrue())
			Expect(o[0].Spec.Versions[0].Deprecated).To(BeFalse())
			Expect(o[1].Spec.Versions).To(HaveLen(1))
			Expect(o[1].Spec.Versions[0].Storage).To(BeFalse())
			Expect(o[1].Spec.Versions[0].Deprecated).To(BeTrue())
		})

	})

	Describe("CRD gen with errors", func() {
		It("should return an error when scopes are mismatched", func() {
			resources := []model.Resource{
				{
					Kind: "kind",
					Spec: model.Field{
						Type: model.Type{
							Name:    "test",
							Message: &v1.AcrylicType{},
						},
					},
					ClusterScoped: true,
				},
				{
					Kind: "kind",
					Spec: model.Field{
						Type: model.Type{
							Name:    "test",
							Message: &v1.PaintColor{},
						},
					},
					ClusterScoped: false,
				},
			}

			_, err := CustomResourceDefinition(resources, map[string]*apiextv1.CustomResourceValidation{}, false)
			Expect(err).To(Equal(fmt.Errorf("mismatched 'currentScope' in versions of CRD for resource kind kind")))
		})
		It("should return an error when ShortNames are mismatched", func() {
			resources := []model.Resource{
				{
					Kind: "kind",
					Spec: model.Field{
						Type: model.Type{
							Name:    "test",
							Message: &v1.AcrylicType{},
						},
					},
					ShortNames: []string{"name1", "name2"},
				},
				{
					Kind: "kind",
					Spec: model.Field{
						Type: model.Type{
							Name:    "test",
							Message: &v1.PaintColor{},
						},
					},
					ShortNames: []string{"name2", "name3"},
				},
			}

			_, err := CustomResourceDefinition(resources, map[string]*apiextv1.CustomResourceValidation{}, false)
			Expect(err).To(Equal(fmt.Errorf("mismatched 'ShortNames' in versions of CRD for resource kind kind")))
		})
		It("should return an error when Categories are mismatched", func() {
			resources := []model.Resource{
				{
					Kind: "kind",
					Spec: model.Field{
						Type: model.Type{
							Name:    "test",
							Message: &v1.AcrylicType{},
						},
					},
					Categories: []string{"cat1", "cat2"},
				},
				{
					Kind: "kind",
					Spec: model.Field{
						Type: model.Type{
							Name:    "test",
							Message: &v1.PaintColor{},
						},
					},
					Categories: []string{"cat2", "cat3"},
				},
			}

			_, err := CustomResourceDefinition(resources, map[string]*apiextv1.CustomResourceValidation{}, false)
			Expect(err).To(Equal(fmt.Errorf("mismatched 'Categories' in versions of CRD for resource kind kind")))
		})
	})
})
