package kuberesource_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	. "github.com/solo-io/skv2/codegen/kuberesource"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/pkg/crdutils"
)

var _ = Describe("Crd", func() {

	Describe("CRD gen", func() {
		var (
			grp model.Group
		)

		BeforeEach(func() {
			grp = model.Group{
				Resources: []model.Resource{
					{
						Kind: "kind",
						Spec: model.Field{
							Type: model.Type{
								Name:    "test",
								Message: &v1.AcrylicType{},
							},
						},
					},
				},
			}
			grp.Init()
		})

		It("should generate spec hash", func() {
			grp.SkipSpecHash = false
			o, err := CustomResourceDefinitions(grp)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(1))
			// note: we intentionally provide the "b6ec737002f7d02e" hash in the test, as it shouldn't change
			// between runs.
			Expect(o[0].GetAnnotations()).To(HaveKeyWithValue(crdutils.CRDSpecHashKey, "b6ec737002f7d02e"))

		})
		It("should not generate spec hash", func() {
			grp.SkipSpecHash = true
			o, err := CustomResourceDefinitions(grp)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(1))
			// note: we intentionally provide the "d18828e563010e32" hash in the test, as it shouldn't change
			// between runs.
			Expect(o[0].GetAnnotations()).NotTo(HaveKey(crdutils.CRDSpecHashKey))
		})
	})

})
