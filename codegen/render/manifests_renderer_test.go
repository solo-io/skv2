package render_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/render"
	"github.com/solo-io/skv2/pkg/crdutils"
)

var _ = Describe("ManifestsRenderer", func() {

	It("should set version without patch", func() {
		m := new(metav1.ObjectMeta)
		render.SetVersionForObject(m, "1.0.0-patch1")
		Expect(m.Annotations[crdutils.CRDVersionKey]).To(Equal("1.0.0"))
	})

	It("should set version without metadata", func() {
		m := new(metav1.ObjectMeta)
		render.SetVersionForObject(m, "1.0.0+metadata")
		Expect(m.Annotations[crdutils.CRDVersionKey]).To(Equal("1.0.0"))
	})
	It("should set version with v", func() {
		m := new(metav1.ObjectMeta)
		render.SetVersionForObject(m, "v1.2.3+metadata")
		Expect(m.Annotations[crdutils.CRDVersionKey]).To(Equal("1.2.3"))
	})

	It("should sanitize crd descriptions", func() {
		obj := map[string]interface{}{
			"description": "{{ some string }} foo",
			"bool":        true,
			"properties": map[string]interface{}{
				"description": "{{ some other string }} foo",
				"bool":        true,
			},
		}
		expectedObj := map[string]interface{}{
			"description": `{{"{{"}} some string {{"}}"}} foo`,
			"bool":        true,
			"properties": map[string]interface{}{
				"description": `{{"{{"}} some other string {{"}}"}} foo`,
				"bool":        true,
			},
		}
		render.EscapeGoTemplateOperators(obj)
		Expect(obj).To(Equal(expectedObj))
	})

	Describe("CRD gen", func() {
		var (
			grp model.Group
		)

		BeforeEach(func() {
			grp = model.Group{
				RenderManifests: true,
				AddChartVersion: "1.0.0",
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
		It("Renderse manifests with chart and spec hash", func() {

			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(
				"appName", "manifestDir", "protoDir",
				nil,
				model.GroupOptions{},
				grp,
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(1))
			Expect(outFiles[0].Content).To(ContainSubstring(crdutils.CRDVersionKey + ": 1.0.0"))
			Expect(outFiles[0].Content).To(ContainSubstring(crdutils.CRDSpecHashKey + ": b6ec737002f7d02e"))

		})
	})
})
