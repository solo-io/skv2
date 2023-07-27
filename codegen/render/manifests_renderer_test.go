package render_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

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

	Describe("Generate non-alpha versioned CRD", func() {
		var (
			grps []*model.Group
		)

		BeforeEach(func() {
			grps = []*model.Group{{
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
						Stored: true,
					},
				}},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renderse manifests with chart and spec hash", func() {

			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(
				"appName", "manifestDir", "protoDir",
				nil,
				model.GroupOptions{},
				grps,
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			Expect(outFiles[0].Content).To(ContainSubstring(crdutils.CRDVersionKey + ": 1.0.0"))
			Expect(outFiles[0].Content).To(ContainSubstring(crdutils.CRDSpecHashKey + ": b6ec737002f7d02e"))
			// only alpha versioned CRDs contain logic to conditionally render templates
			Expect(outFiles[0].Content).To(Equal(outFiles[0].Content))
		})
	})

	Describe("Generate alpha versioned CRD", func() {
		var (
			grps []*model.Group
		)

		BeforeEach(func() {
			grps = []*model.Group{{
				GroupVersion: schema.GroupVersion{
					Group:   "things.test.io",
					Version: "v1alpha1",
				},
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
						Stored: true,
					},
				}},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renderse manifests with chart and spec hash", func() {

			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(
				"appName", "manifestDir", "protoDir",
				nil,
				model.GroupOptions{},
				grps,
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			// only alpha versioned CRDs contain logic to conditionally render templates
			Expect(outFiles[1].Content).To(HavePrefix("\n{{- if has \"kinds.things.test.io/v1alpha1\" $.Values.enabledExperimentalApi }}"))
			Expect(outFiles[1].Content).To(HaveSuffix("{{- end  }}\n---\n"))
			Expect(outFiles[1].Content).To(ContainSubstring(crdutils.CRDVersionKey + ": 1.0.0"))
			Expect(outFiles[1].Content).To(ContainSubstring(crdutils.CRDSpecHashKey + ": 80c06d3e2484e4c8"))
		})
	})

	Describe("Generate combined alpha and non-alpha versioned CRD", func() {
		var (
			grps []*model.Group
		)

		BeforeEach(func() {
			grps = []*model.Group{{
				GroupVersion: schema.GroupVersion{
					Group:   "things.test.io",
					Version: "v2alpha1",
				},
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
						Stored: false,
					},
				}},
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
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
							Stored: true,
						},
					}},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renderse manifests with chart and spec hash", func() {

			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(
				"appName", "manifestDir", "protoDir",
				nil,
				model.GroupOptions{},
				grps,
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			// only v2alpha1 version of the CRDs is conditionally rendered, v1 has no conditions surrounding it
			Expect(outFiles[1].Content).To(ContainSubstring("subresources: {}\n  {{- if has \"kinds.things.test.io/v2alpha1\" $.Values.enabledExperimentalApi }}\n  - name: v2alpha1"))
			Expect(outFiles[1].Content).To(ContainSubstring("{{- end }}\n---\n"))
		})
	})

})
