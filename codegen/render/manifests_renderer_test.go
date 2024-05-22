package render_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/codegen/proto/schemagen"
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

	Describe("should respect skip manifest flags", func() {
		It("skips crd manifest if options.SkipCrdsManifest is true", func() {
			grps := []*model.Group{
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
							Stored: false,
						},
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:          "appName",
				ManifestRoot:     "manifestDir",
				ProtoDir:         "protoDir",
				Groups:           grps,
				SkipCrdsManifest: true,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(1)) // should only contain template manifest
			Expect(outFiles[0].Path).To(Equal("manifestDir/templates/things.test.io_crds.yaml"))
		})

		It("skips crd manifest if group.SkipCRDManifest is true", func() {
			grps := []*model.Group{
				{
					SkipCRDManifest: true,
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
							Stored: false,
						},
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:      "appName",
				ManifestRoot: "manifestDir",
				ProtoDir:     "protoDir",
				Groups:       grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(1)) // should only contain template manifest
			Expect(outFiles[0].Path).To(Equal("manifestDir/templates/things.test.io_crds.yaml"))
		})

		It("skips templated crd manifest if group.SkipTemplatedCRDManifest is true", func() {
			grps := []*model.Group{
				{
					SkipTemplatedCRDManifest: true,
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
							Stored: false,
						},
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:      "appName",
				ManifestRoot: "manifestDir",
				ProtoDir:     "protoDir",
				Groups:       grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(1)) // should only contain crd manifest
			Expect(outFiles[0].Path).To(Equal("manifestDir/crds/things.test.io_crds.yaml"))
		})

		It("skips both manifests if both skip flags are true", func() {
			grps := []*model.Group{
				{
					SkipCRDManifest:          true,
					SkipTemplatedCRDManifest: true,
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
							Stored: false,
						},
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:      "appName",
				ManifestRoot: "manifestDir",
				ProtoDir:     "protoDir",
				Groups:       grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(0))
		})
	})

	Describe("Generate non-alpha versioned CRD", func() {
		var grps []*model.Group

		BeforeEach(func() {
			grps = []*model.Group{
				{
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renders manifests with chart and spec hash", func() {
			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:                 "appName",
				ManifestRoot:            "manifestDir",
				ProtoDir:                "protoDir",
				EnabledAlphaApiFlagName: "enabledExperimentalApi",
				Groups:                  grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			Expect(outFiles[0].Content).To(ContainSubstring(crdutils.CRDVersionKey + ": 1.0.0"))
			Expect(outFiles[0].Content).To(ContainSubstring(crdutils.CRDSpecHashKey + ": 13a71ec8b0049565"))
			// only alpha versioned CRDs contain logic to conditionally render templates
			Expect(outFiles[0].Content).To(Equal(outFiles[0].Content))
		})
	})

	Describe("Generate alpha versioned CRD", func() {
		var grps []*model.Group

		BeforeEach(func() {
			grps = []*model.Group{
				{
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renders manifests with template and spec hash", func() {
			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:                 "appName",
				ManifestRoot:            "manifestDir",
				ProtoDir:                "protoDir",
				EnabledAlphaApiFlagName: "enabledExperimentalApi",
				Groups:                  grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			// only alpha versioned CRDs contain logic to conditionally render templates
			Expect(outFiles[1].Content).To(HavePrefix("\n{{- if has \"kinds.things.test.io/v1alpha1\" $.Values.enabledExperimentalApi }}"))
			Expect(outFiles[1].Content).To(ContainSubstring("{{- end  }}"))
			Expect(outFiles[1].Content).To(ContainSubstring(crdutils.CRDVersionKey + ": 1.0.0"))
			Expect(outFiles[1].Content).To(ContainSubstring(crdutils.CRDSpecHashKey + ": 105d1e4c78cac3c"))
		})
	})

	Describe("Skip template for grandfathered alpha versioned CRD", func() {
		var grps []*model.Group

		BeforeEach(func() {
			grps = []*model.Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1alpha1",
					},
					RenderManifests:           true,
					AddChartVersion:           "1.0.0",
					SkipConditionalCRDLoading: true,
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renders manifests without template", func() {
			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:                 "appName",
				ManifestRoot:            "manifestDir",
				ProtoDir:                "protoDir",
				EnabledAlphaApiFlagName: "enabledExperimentalApi",
				Groups:                  grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			// only alpha versioned CRDs contain logic to conditionally render templates
			Expect(outFiles[1].Content).ToNot(ContainSubstring("{{- if has \"kinds.things.test.io/v1alpha1\" $.Values.enabledExperimentalApi }}"))
			Expect(outFiles[1].Content).ToNot(ContainSubstring("{{- end  }}"))
		})
	})

	Describe("Generate combined alpha, grandfathered alpha, and non-alpha versioned CRD", func() {
		var grps []*model.Group

		BeforeEach(func() {
			grps = []*model.Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v3alpha1",
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
					},
				},
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v2",
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
					},
				},
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1alpha1",
					},
					RenderManifests:           true,
					AddChartVersion:           "1.0.0",
					SkipConditionalCRDLoading: true,
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renders manifests with chart and spec hash", func() {
			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:                 "appName",
				ManifestRoot:            "manifestDir",
				ProtoDir:                "protoDir",
				EnabledAlphaApiFlagName: "enabledExperimentalApi",
				Groups:                  grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			// only v3alpha1 version of the CRDs is conditionally rendered, v2 and v1alpha1 have no conditions surrounding them
			Expect(outFiles[1].Content).To(ContainSubstring("subresources: {}\n  {{- if has \"kinds.things.test.io/v3alpha1\" $.Values.enabledExperimentalApi }}"))
			Expect(outFiles[1].Content).ToNot(ContainSubstring("{{- if has \"kinds.things.test.io/v1alpha1\" $.Values.enabledExperimentalApi }}\n  - name: v2alpha1"))
			Expect(outFiles[1].Content).To(ContainSubstring("{{- end }}\n---\n"))
			Expect(strings.Count(outFiles[1].Content, "{{- end }}")).To(Equal(1))
		})
	})

	Describe("Generate combined alpha, grandfathered alpha, and non-alpha versioned CRD", func() {
		var grps []*model.Group

		BeforeEach(func() {
			grps = []*model.Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v3alpha1",
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
					},
				},
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v2",
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
					},
				},
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1alpha1",
					},
					RenderManifests:           true,
					AddChartVersion:           "1.0.0",
					SkipConditionalCRDLoading: true,
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}
		})
		It("Renders manifests with chart and spec hash", func() {
			// get api-level code gen options from descriptors
			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:                 "appName",
				ManifestRoot:            "manifestDir",
				ProtoDir:                "protoDir",
				EnabledAlphaApiFlagName: "enabledExperimentalApi",
				Groups:                  grps,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // legacy and templated manifests
			// only v3alpha1 version of the CRDs is conditionally rendered, v2 and v1alpha1 have no conditions surrounding them
			expectedTemplateString := "subresources: {}\n  {{- if has \"kinds.things.test.io/v3alpha1\" $.Values.enabledExperimentalApi }}"
			Expect(outFiles[1].Content).To(ContainSubstring(expectedTemplateString))
			Expect(strings.Count(outFiles[1].Content, expectedTemplateString)).To(Equal(1))
			Expect(outFiles[1].Content).ToNot(ContainSubstring("{{- if has \"kinds.things.test.io/v1alpha1\" $.Values.enabledExperimentalApi }}\n  - name: v2alpha1"))
			Expect(outFiles[1].Content).To(ContainSubstring("{{- end }}\n---\n"))
			Expect(strings.Count(outFiles[1].Content, "{{- end }}\n---\n")).To(Equal(1))
			Expect(strings.Count(outFiles[1].Content, "{{- end }}")).To(Equal(1))
		})
	})

	Describe("Render CRD template when 'EnabledAlphaApiFlagName' isn't set", func() {
		It("and resource contains an alpha version that should not be skipped", func() {
			grps := []*model.Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v3alpha1",
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			_, err := render.RenderManifests(render.RenderOptions{
				AppName:      "appName",
				ManifestRoot: "manifestDir",
				ProtoDir:     "protoDir",
				Groups:       grps,
			})
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(fmt.Errorf("error rendering CRD template for kind kind: 'EnabledAlphaApiFlagName' is not defined")))
		})
		It("and resource contains an alpha version that should be skipped", func() {
			grps := []*model.Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v3alpha1",
					},
					RenderManifests:           true,
					AddChartVersion:           "1.0.0",
					SkipConditionalCRDLoading: true,
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
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			_, err := render.RenderManifests(render.RenderOptions{
				AppName:      "appName",
				ManifestRoot: "manifestDir",
				ProtoDir:     "protoDir",
				Groups:       grps,
			})
			Expect(err).To(BeNil())
		})
	})

	Describe("using protoc-gen-openapi", func() {
		It("should generate CRD", func() {
			grps := []*model.Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "protoc-gen-openapi.test.io",
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
							Stored: false,
						},
					},
				},
			}
			for i := range grps {
				grps[i].Init()
			}

			outFiles, err := render.RenderManifests(render.RenderOptions{
				AppName:      "appName",
				ManifestRoot: "manifestDir",
				ProtoDir:     "protoDir",
				Groups:       grps,
				GroupOptions: model.GroupOptions{
					SchemaGenerator: schemagen.ProtocGenOpenAPI,
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(outFiles).To(HaveLen(2)) // CRD and templated CRD
		})
	})
})
