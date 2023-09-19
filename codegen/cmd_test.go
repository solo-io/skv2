package codegen_test

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	goyaml "gopkg.in/yaml.v3"
	rbacv1 "k8s.io/api/rbac/v1"
	v12 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/codegen"
	. "github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/skv2_anyvendor"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/skv2/contrib"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubeyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"
)

var _ = Describe("Cmd", func() {
	skv2Imports := skv2_anyvendor.CreateDefaultMatchOptions(
		[]string{"codegen/test/*.proto"},
	)

	// used for a proto option which disables openapi validation on fields
	skv2Imports.External["github.com/solo-io/cue"] = []string{
		"encoding/protobuf/cue/cue.proto",
	}

	It("generates controller code and manifests for a proto file", func() {
		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2Imports,
			RenderProtos:    true,

			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},
								Args: []string{"foo"},
							},
						},
					},
				},
				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		err = exec.Command("goimports", "-w", "codegen/test/api").Run()
		Expect(err).NotTo(HaveOccurred())
	})

	// note that there is no .proto file this is generated from; it simply pulls in field definitions
	// from other packages
	It("allows fields from other packages", func() {
		// note that there is no .proto file this is generated from; it simply pulls in field definitions
		// from other packages
		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "other.things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind: "KubernetesCluster",
							Spec: Field{Type: Type{
								Name:      "KubernetesCluster",
								GoPackage: "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1",
							}},
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2Imports,
			RenderProtos:    true,
			ManifestRoot:    "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		// Make sure all generated code compiles
		err = exec.Command("go", "build", "codegen/test/api/other.things.test.io/v1/...").Run()
		Expect(err).NotTo(HaveOccurred())
	})

	It("generates controller code and manifests for a proto file", func() {
		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2Imports,
			RenderProtos:    true,

			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},
								Args: []string{"foo"},
							},
						},
					},
				},
				Values: nil,
				ValuesInlineDocs: &ValuesInlineDocs{
					LineLengthLimit: 80,
				},
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		fileContents, err := os.ReadFile("codegen/test/chart/values.yaml")
		Expect(err).NotTo(HaveOccurred())

		node := goyaml.Node{}
		Expect(goyaml.Unmarshal(fileContents, &node)).NotTo(HaveOccurred())

		painterNode := node.Content[0].Content[1]
		enabledMapField := painterNode.Content[0]
		Expect(enabledMapField.HeadComment).To(Equal("# Arbitrary overrides for the component's [deployment\n# template](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/)"))
		envMapField := painterNode.Content[2]
		Expect(envMapField.HeadComment).To(Equal("# Enable creation of the deployment/service."))
	})

	It("generates from templates using a name override", func() {
		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2Imports,
			RenderProtos:    true,
			Chart: &Chart{
				Operators: []Operator{
					{
						Name:                   "painter-original-name",
						ValuesFileNameOverride: "override-name",
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},
								Args: []string{"foo"},
							},
						},
					},
				},
				Values: nil,
				ValuesInlineDocs: &ValuesInlineDocs{
					LineLengthLimit: 80,
				},
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
				ValuesReferenceDocs: ValuesReferenceDocs{
					Title:    "Name Override Chart",
					Filename: "name_override_chart_reference.md",
				},
			},

			ManifestRoot: "codegen/test/name_override_chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		fileContents, err := os.ReadFile("codegen/test/name_override_chart/values.yaml")
		Expect(err).NotTo(HaveOccurred())
		valuesString := string(fileContents)

		Expect(valuesString).NotTo(ContainSubstring("painterOriginalName"))
		Expect(valuesString).To(ContainSubstring("overrideName"))

		fileContents, err = os.ReadFile("codegen/test/name_override_chart/templates/deployment.yaml")
		Expect(err).NotTo(HaveOccurred())
		deploymentString := string(fileContents)

		Expect(deploymentString).NotTo(ContainSubstring("$.Values.painterOriginalName"))
		Expect(deploymentString).To(ContainSubstring("$.Values.overrideName"))

		fileContents, err = os.ReadFile("codegen/test/name_override_chart/name_override_chart_reference.md")
		Expect(err).NotTo(HaveOccurred())
		docString := string(fileContents)

		Expect(docString).NotTo(ContainSubstring("painterOriginalName"))
		Expect(docString).To(ContainSubstring("overrideName"))
	})

	It("generates json schema for the values file", func() {
		type CustomType1 struct {
			Field1 string `json:"customField1"`
		}
		type CustomType2 struct {
			Field2 string `json:"customField2"`
		}
		typeMapper := func(t reflect.Type, s map[string]interface{}) interface{} {
			if t == reflect.TypeOf(CustomType2{}) {
				return map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"customField2_renamed": map[string]interface{}{
							"type": "number",
						},
					},
				}
			}

			return nil
		}

		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2Imports,
			RenderProtos:    true,

			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},
								Args: []string{"foo"},
							},
						},
						Values: &CustomType2{},
					},
				},
				Values: &CustomType1{},
				ValuesInlineDocs: &ValuesInlineDocs{
					LineLengthLimit: 80,
				},
				// TODO here we should also test the custom type mappings ...
				JsonSchema: &JsonSchema{
					CustomTypeMapper: typeMapper,
				},
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		fileContents, err := os.ReadFile("codegen/test/chart/values.schema.json")
		Expect(err).NotTo(HaveOccurred())

		type jsonSchema struct {
			Properties *struct {
				CustomField1 map[string]interface{} `json:"customField1"`
				Painter      *struct {
					Properties map[string]interface{} `json:"properties"`
				} `json:"painter"`
			} `json:"properties"`
		}

		schema := jsonSchema{}
		Expect(json.Unmarshal(fileContents, &schema)).NotTo(HaveOccurred())

		// expect that the custom values are in the schema
		Expect(schema.Properties.CustomField1).NotTo(BeNil())
		Expect(schema.Properties.CustomField1["type"]).To(Equal("string"))

		// expect that the values from the painter opeartor are in the schema
		painter := schema.Properties.Painter
		Expect(painter).NotTo(BeNil())

		// expect painterSchema has some of the base properities
		painterProperties := painter.Properties
		Expect(painterProperties).To(HaveKey("image"))
		Expect(painterProperties).To(HaveKey("env"))
		Expect(painterProperties).To(HaveKey("sidecars"))
		Expect(painterProperties).To(HaveKey("securityContext"))

		// expect painter schema also contains properties from the custom values
		Expect(painterProperties).To(HaveKey("customField2_renamed"))
	})

	DescribeTable("configuring the runAsUser value",
		func(floatingUserId bool, runAsUser, expectedRunAsUser int) {
			cmd := &Command{
				Chart: &Chart{
					Operators: []Operator{
						{
							Name: "painter",
							Deployment: Deployment{
								Container: Container{
									Image: Image{
										Tag:        "v0.0.0",
										Repository: "painter",
										Registry:   "quay.io/solo-io",
										PullPolicy: "IfNotPresent",
									},
								},
							},
						},
					},

					Values: nil,
					Data: Data{
						ApiVersion:  "v1",
						Description: "",
						Name:        "Painting Operator",
						Version:     "v0.0.1",
						Home:        "https://docs.solo.io/skv2/latest",
						Sources: []string{
							"https://github.com/solo-io/skv2",
						},
					},
				},

				ManifestRoot: "codegen/test/chart",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			painterValues := map[string]interface{}{"floatingUserId": floatingUserId}
			if runAsUser > 0 {
				painterValues["runAsUser"] = runAsUser
			}
			helmValues := map[string]interface{}{"painter": painterValues}

			renderedManifests := helmTemplate("codegen/test/chart", helmValues)

			var renderedDeployment *appsv1.Deployment
			decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
			for {
				obj := &unstructured.Unstructured{}
				err := decoder.Decode(obj)
				if err != nil {
					break
				}
				if obj.GetName() != "painter" || obj.GetKind() != "Deployment" {
					continue
				}

				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedDeployment = &appsv1.Deployment{}
				err = json.Unmarshal(bytes, renderedDeployment)
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(renderedDeployment).NotTo(BeNil())
			renderedRunAsUser := renderedDeployment.Spec.Template.Spec.Containers[0].SecurityContext.RunAsUser
			if expectedRunAsUser == 0 {
				Expect(renderedRunAsUser).To(BeNil())
			} else {
				Expect(renderedRunAsUser).ToNot(BeNil())
				Expect(*renderedRunAsUser).To(BeEquivalentTo(expectedRunAsUser))
			}
		},
		Entry("default values", false, 0, 10101),
		Entry("set runAsUser value", false, 20202, 20202),
		Entry("floatingUserId enabled", true, 10101, 0),
	)

	DescribeTable("supports overriding the container security context",
		func(securityContext *v1.SecurityContext, omitSecurityContext bool) {
			cmd := &Command{
				Chart: &Chart{
					Operators: []Operator{
						{
							Name: "painter",
							Deployment: Deployment{
								Container: Container{
									Image: Image{
										Tag:        "v0.0.0",
										Repository: "painter",
										Registry:   "quay.io/solo-io",
										PullPolicy: "IfNotPresent",
									},
									SecurityContext: securityContext,
								},
							},
						},
					},

					Values: nil,
					Data: Data{
						ApiVersion:  "v1",
						Description: "",
						Name:        "Painting Operator",
						Version:     "v0.0.1",
						Home:        "https://docs.solo.io/skv2/latest",
						Sources: []string{
							"https://github.com/solo-io/skv2",
						},
					},
				},

				ManifestRoot: "codegen/test/chart",
			}
			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			helmValues := map[string]interface{}{}
			if omitSecurityContext {
				helmValues["painter"] = map[string]interface{}{"securityContext": false}
			}

			renderedManifests := helmTemplate("codegen/test/chart", helmValues)

			var renderedDeployment *appsv1.Deployment
			decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
			for {
				obj := &unstructured.Unstructured{}
				err := decoder.Decode(obj)
				if err != nil {
					break
				}
				if obj.GetName() != "painter" || obj.GetKind() != "Deployment" {
					continue
				}

				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedDeployment = &appsv1.Deployment{}
				err = json.Unmarshal(bytes, renderedDeployment)
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(renderedDeployment).NotTo(BeNil())

			pointerBool := func(b bool) *bool { return &b }
			pointerInt64 := func(i int64) *int64 { return &i }
			defaultSecurityContext := v1.SecurityContext{
				RunAsNonRoot:             pointerBool(true),
				RunAsUser:                pointerInt64(10101),
				ReadOnlyRootFilesystem:   pointerBool(true),
				AllowPrivilegeEscalation: pointerBool(false),
				Capabilities: &v1.Capabilities{
					Drop: []v1.Capability{"ALL"},
				},
			}

			renderedSecurityContext := renderedDeployment.Spec.Template.Spec.Containers[0].SecurityContext

			if securityContext == nil && !omitSecurityContext {
				Expect(*renderedSecurityContext).To(Equal(defaultSecurityContext))
			} else if securityContext == nil && omitSecurityContext {
				Expect(*renderedSecurityContext).To(Equal(v1.SecurityContext{}))
			} else {
				Expect(*renderedSecurityContext).To(Equal(*securityContext))
			}
		},
		Entry("renders default container security context", nil, false),
		Entry("renders empty map for container security context when set as false via helm cli", nil, true),
		Entry("overrides container security context with empty map", &v1.SecurityContext{}, false),
		Entry("overrides container security context", &v1.SecurityContext{
			RunAsNonRoot: func(b bool) *bool { return &b }(true),
			RunAsUser:    func(i int64) *int64 { return &i }(20202),
		}, false),
	)

	It("supports disabling the deployment and service specs via helm values", func() {
		cmd := &Command{
			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},

								Args: []string{"foo"},
								Env: []v1.EnvVar{
									{
										Name:  "FOO",
										Value: "BAR",
									},
								},
								ReadinessProbe: &ReadinessProbe{
									Path:                "/",
									Port:                "8080",
									PeriodSeconds:       10,
									InitialDelaySeconds: 5,
								},
							},

							Sidecars: []Sidecar{
								{
									Name: "palette",
									Container: Container{
										Image: Image{
											Tag:        "v0.0.0",
											Repository: "palette",
											Registry:   "quay.io/solo-io",
											PullPolicy: "IfNotPresent",
										},
										Args: []string{"bar", "baz"},
										VolumeMounts: []v1.VolumeMount{
											{
												Name:      "paint",
												MountPath: "/etc/paint",
											},
										},
										LivenessProbe: &v1.Probe{
											ProbeHandler: v1.ProbeHandler{
												HTTPGet: &v1.HTTPGetAction{
													Path: "/",
													Port: intstr.FromInt(8080),
												},
											},
											PeriodSeconds:       60,
											InitialDelaySeconds: 30,
										},
									},
								},
							},

							Volumes: []v1.Volume{
								{
									Name: "paint",
									VolumeSource: v1.VolumeSource{
										EmptyDir: &v1.EmptyDirVolumeSource{},
									},
								},
							},

							CustomDeploymentAnnotations: map[string]string{
								"deployment": "annotation",
							},
							CustomDeploymentLabels: map[string]string{
								"deployment": "labels",
							},
							CustomPodAnnotations: map[string]string{
								"pod": "annotations",
							},
							CustomPodLabels: map[string]string{
								"pod": "labels",
							},
						},

						Service: Service{
							Ports: []ServicePort{{
								Name:        "http",
								DefaultPort: 1234,
							}},
							CustomAnnotations: map[string]string{
								"service": "annotations",
							},
							CustomLabels: map[string]string{
								"service": "labels",
							},
						},
					},
				},

				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		helmValues := map[string]interface{}{
			"painter": map[string]interface{}{
				"enabled": false,
			},
		}
		renderedManifests := helmTemplate("codegen/test/chart", helmValues)

		var (
			renderedService    *v1.Service
			renderedDeployment *appsv1.Deployment
		)
		decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
		for {
			obj := &unstructured.Unstructured{}
			err := decoder.Decode(obj)
			if err != nil {
				break
			}
			if obj.GetName() != "painter" {
				continue
			}

			switch obj.GetKind() {
			case "Service":
				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedService = &v1.Service{}
				err = json.Unmarshal(bytes, renderedService)
				Expect(err).NotTo(HaveOccurred())
			case "Deployment":
				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedDeployment = &appsv1.Deployment{}
				err = json.Unmarshal(bytes, renderedDeployment)
				Expect(err).NotTo(HaveOccurred())
			}
		}
		Expect(renderedDeployment).To(BeNil())
		Expect(renderedService).To(BeNil())
	})

	It("supports overriding the deployment and service specs via helm values", func() {
		cmd := &Command{
			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},

								Args: []string{"foo"},
								Env: []v1.EnvVar{
									{
										Name:  "FOO",
										Value: "BAR",
									},
								},
								ReadinessProbe: &ReadinessProbe{
									Path:                "/",
									Port:                "8080",
									PeriodSeconds:       10,
									InitialDelaySeconds: 5,
								},
							},

							Sidecars: []Sidecar{
								{
									Name: "palette",
									Container: Container{
										Image: Image{
											Tag:        "v0.0.0",
											Repository: "palette",
											Registry:   "quay.io/solo-io",
											PullPolicy: "IfNotPresent",
										},
										Args: []string{"bar", "baz"},
										VolumeMounts: []v1.VolumeMount{
											{
												Name:      "paint",
												MountPath: "/etc/paint",
											},
										},
										LivenessProbe: &v1.Probe{
											ProbeHandler: v1.ProbeHandler{
												HTTPGet: &v1.HTTPGetAction{
													Path: "/",
													Port: intstr.FromInt(8080),
												},
											},
											PeriodSeconds:       60,
											InitialDelaySeconds: 30,
										},
									},
								},
							},

							Volumes: []v1.Volume{
								{
									Name: "paint",
									VolumeSource: v1.VolumeSource{
										EmptyDir: &v1.EmptyDirVolumeSource{},
									},
								},
							},

							CustomDeploymentAnnotations: map[string]string{
								"deployment": "annotation",
							},
							CustomDeploymentLabels: map[string]string{
								"deployment": "labels",
							},
							CustomPodAnnotations: map[string]string{
								"pod": "annotations",
							},
							CustomPodLabels: map[string]string{
								"pod": "labels",
							},
						},

						Service: Service{
							Ports: []ServicePort{{
								Name:        "http",
								DefaultPort: 1234,
							}},
							CustomAnnotations: map[string]string{
								"service": "annotations",
							},
							CustomLabels: map[string]string{
								"service": "labels",
							},
						},
					},
				},

				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		var (
			// overrides used in test
			loadBalancerIp       = "1.2.3.4"
			replicas       int32 = 3
		)

		marshalMap := func(v interface{}) map[string]interface{} {
			data, err := json.Marshal(v)
			Expect(err).NotTo(HaveOccurred())

			var m map[string]interface{}
			err = json.Unmarshal(data, &m)
			Expect(err).NotTo(HaveOccurred())

			return m
		}

		helmValues := map[string]interface{}{
			"painter": map[string]interface{}{
				"serviceOverrides": marshalMap(&v1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							// override labels
							"extraservice": "labels",
						},
						Annotations: map[string]string{
							// override annotations
							"extraservice": "annotations",
						},
					},
					// override load balancer ip
					Spec: v1.ServiceSpec{
						LoadBalancerIP: loadBalancerIp,
					},
				}),
				"deploymentOverrides": marshalMap(&appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							// override labels
							"extradeployment": "labels",
						},
						Annotations: map[string]string{
							// override annotations
							"extradeployment": "annotations",
						},
					},
					// override replicas
					Spec: appsv1.DeploymentSpec{
						Replicas: &replicas,
						Template: v1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									// override labels
									"extrapod": "labels",
								},
								Annotations: map[string]string{
									// override annotations
									"extrapod": "annotations",
								},
							},
						},
					},
				}),
			},
		}

		renderedManifests := helmTemplate("codegen/test/chart", helmValues)

		var (
			renderedService    *v1.Service
			renderedDeployment *appsv1.Deployment
		)
		decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
		for {
			obj := &unstructured.Unstructured{}
			err := decoder.Decode(obj)
			if err != nil {
				break
			}
			if obj.GetName() != "painter" {
				continue
			}

			switch obj.GetKind() {
			case "Service":
				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedService = &v1.Service{}
				err = json.Unmarshal(bytes, renderedService)
				Expect(err).NotTo(HaveOccurred())
			case "Deployment":
				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedDeployment = &appsv1.Deployment{}
				err = json.Unmarshal(bytes, renderedDeployment)
				Expect(err).NotTo(HaveOccurred())
			}
		}
		Expect(renderedDeployment).NotTo(BeNil())
		Expect(renderedService).NotTo(BeNil())

		Expect(renderedDeployment.Spec.Template.Labels).To(Equal(map[string]string{
			"app":      "painter",
			"pod":      "labels",
			"extrapod": "labels",
		}))
		Expect(renderedDeployment.Spec.Template.Annotations).To(Equal(map[string]string{
			"prometheus.io/port":     "9091",
			"prometheus.io/scrape":   "true",
			"app.kubernetes.io/name": "painter",
			"extrapod":               "annotations",
			"pod":                    "annotations",
			"prometheus.io/path":     "/metrics",
		}))
		Expect(renderedService.Labels).To(Equal(map[string]string{
			"app":          "painter",
			"extraservice": "labels",
			"service":      "labels",
		}))
		Expect(renderedService.Annotations).To(Equal(map[string]string{
			"app.kubernetes.io/name": "painter",
			"extraservice":           "annotations",
			"service":                "annotations",
		}))
		Expect(renderedDeployment.Labels).To(Equal(map[string]string{
			"app":             "painter",
			"deployment":      "labels",
			"extradeployment": "labels",
		}))
		Expect(renderedDeployment.Annotations).To(Equal(map[string]string{
			"extradeployment":        "annotations",
			"app.kubernetes.io/name": "painter",
			"deployment":             "annotation",
		}))
		Expect(renderedService.Spec.LoadBalancerIP).To(Equal(loadBalancerIp))
		Expect(*renderedDeployment.Spec.Replicas).To(Equal(replicas))
	})

	It("supports enabling one service depending on whether another is enabled", func() {
		cmd := &Command{
			Groups: []Group{
				{
					GroupVersion: schema.GroupVersion{
						Group:   "things.test.io",
						Version: "v1",
					},
					Module: "github.com/solo-io/skv2",
					Resources: []Resource{
						{
							Kind:   "Paint",
							Spec:   Field{Type: Type{Name: "PaintSpec"}},
							Status: &Field{Type: Type{Name: "PaintStatus"}},
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
						},
					},
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllGroupCustomTemplates,
				},
			},
			AnyVendorConfig: skv2Imports,
			RenderProtos:    true,

			Chart: &Chart{
				Operators: []Operator{
					{
						Name:             "painter",
						EnabledDependsOn: []string{"test1", "test2"},
						ClusterRbac: []rbacv1.PolicyRule{
							{
								Verbs: []string{"GET"},
							},
						},
						Deployment: Deployment{
							Container: Container{
								Image: Image{
									Tag:        "v0.0.0",
									Repository: "painter",
									Registry:   "quay.io/solo-io",
									PullPolicy: "IfNotPresent",
								},
								Args: []string{"foo"},
							},
						},
					},
				},
				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},

			ManifestRoot: "codegen/test/chart",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		fileContents, err := os.ReadFile("codegen/test/chart/templates/deployment.yaml")
		Expect(err).NotTo(HaveOccurred())

		Expect(string(fileContents)).To(ContainSubstring("{{- if and $painter.enabled $.Values.test1.enabled $.Values.test2.enabled }}"))

		expectedSA := "kind: ServiceAccount\nmetadata:\n  labels:\n    app: painter\n  name: painter\n"
		expectedCR := "kind: ClusterRole\napiVersion: rbac.authorization.k8s.io/v1\nmetadata:\n  name: painter"

		type enabledThing struct {
			Enabled bool `json:"enabled"`
		}
		helmValues := map[string]*enabledThing{
			"painter": {Enabled: true},
			"test1":   {Enabled: false},
			"test2":   {Enabled: false},
		}

		renderedManifests := helmTemplate("codegen/test/chart", helmValues)
		Expect(renderedManifests).NotTo(ContainSubstring(expectedSA))
		Expect(renderedManifests).NotTo(ContainSubstring(expectedCR))

		helmValues["test1"].Enabled = true
		renderedManifests = helmTemplate("codegen/test/chart", helmValues)
		Expect(renderedManifests).NotTo(ContainSubstring(expectedSA))
		Expect(renderedManifests).NotTo(ContainSubstring(expectedCR))

		helmValues["test2"].Enabled = true
		renderedManifests = helmTemplate("codegen/test/chart", helmValues)
		Expect(string(renderedManifests)).To(ContainSubstring(expectedSA))
		Expect(renderedManifests).To(ContainSubstring(expectedCR))
	})

	Context("generates CRD with validation schema for a proto file", func() {
		var cmd *Command

		BeforeEach(func() {
			cmd = &Command{
				Groups: []Group{
					{
						GroupVersion: schema.GroupVersion{
							Group:   "things.test.io",
							Version: "v1",
						},
						Module: "github.com/solo-io/skv2",
						Resources: []Resource{
							{
								Kind:   "Paint",
								Spec:   Field{Type: Type{Name: "PaintSpec"}},
								Status: &Field{Type: Type{Name: "PaintStatus"}},
							},
							{
								Kind:          "ClusterResource",
								Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
								ClusterScoped: true,
							},
						},
						RenderManifests:         true,
						RenderValidationSchemas: true,
						ApiRoot:                 "codegen/test/api",
						PointerSlices:           true,
					},
				},
				AnyVendorConfig: skv2Imports,
				RenderProtos:    true,

				Chart: &Chart{
					Operators: []Operator{
						{
							Name: "painter",
							Deployment: Deployment{
								Container: Container{
									Image: Image{
										Tag:        "v0.0.0",
										Repository: "painter",
										Registry:   "quay.io/solo-io",
										PullPolicy: "IfNotPresent",
									},
									Args: []string{"foo"},
									Env: []v1.EnvVar{
										{
											Name:  "FOO",
											Value: "BAR",
										},
									},
									ReadinessProbe: &ReadinessProbe{
										Exec:                []string{"redis-cli", "ping"},
										PeriodSeconds:       10,
										InitialDelaySeconds: 5,
									},
								},

								Sidecars: []Sidecar{
									{
										Name: "palette",
										Container: Container{
											Image: Image{
												Tag:        "v0.0.0",
												Repository: "palette",
												Registry:   "quay.io/solo-io",
												PullPolicy: "IfNotPresent",
											},
											Args: []string{"bar", "baz"},
											VolumeMounts: []v1.VolumeMount{
												{
													Name:      "paint",
													MountPath: "/etc/paint",
												},
											},
											LivenessProbe: &v1.Probe{
												ProbeHandler: v1.ProbeHandler{
													HTTPGet: &v1.HTTPGetAction{
														Path: "/",
														Port: intstr.FromInt(8080),
													},
												},
												PeriodSeconds:       60,
												InitialDelaySeconds: 30,
											},
										},
									},
								},

								Volumes: []v1.Volume{
									{
										Name: "paint",
										VolumeSource: v1.VolumeSource{
											EmptyDir: &v1.EmptyDirVolumeSource{},
										},
									},
								},

								CustomDeploymentAnnotations: map[string]string{
									"deployment": "annotation",
								},
								CustomDeploymentLabels: map[string]string{
									"deployment": "labels",
								},
								CustomPodAnnotations: map[string]string{
									"pod": "annotations",
								},
								CustomPodLabels: map[string]string{
									"pod": "labels",
								},
							},
						},
					},
					Values: nil,
					Data: Data{
						ApiVersion:  "v1",
						Description: "",
						Name:        "Painting Operator",
						Version:     "v0.0.1",
						Home:        "https://docs.solo.io/skv2/latest",
						Sources: []string{
							"https://github.com/solo-io/skv2",
						},
					},
				},

				ManifestRoot: "codegen/test/chart",
			}
		})

		It("can include field descriptions", func() {
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_v1_crds.yaml")

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			bytes, err := os.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bytes)).To(ContainSubstring("description: OpenAPI gen test for recursive fields"))
		})

		It("generates google.protobuf.Value with no type", func() {
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_v1_crds.yaml")

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			bytes, err := os.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			generatedCrd := &v12.CustomResourceDefinition{}
			Expect(yaml.Unmarshal(bytes, generatedCrd)).NotTo(HaveOccurred())
			protobufValueField := generatedCrd.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties["spec"].Properties["recursiveType"].Properties["protobufValue"]
			// access the field to make sure it's not nil
			Expect(protobufValueField.XPreserveUnknownFields).ToNot(BeNil())
			Expect(*protobufValueField.XPreserveUnknownFields).To(BeTrue())
			Expect(protobufValueField.Type).To(BeEmpty())
		})

		It("can exclude field descriptions", func() {
			// write this manifest to a different dir to avoid modifying the crd file from the
			// above test, which other tests seem to depend on
			cmd.ManifestRoot = "codegen/test/chart-no-desc"
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_v1_crds.yaml")

			cmd.Groups[0].SkipSchemaDescriptions = true

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			bytes, err := os.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bytes)).NotTo(ContainSubstring("description:"))
		})
	})

	DescribeTable("rendering the sidecar values",
		func(sidecarName string, expectedSidecarParentKey string) {
			cmd := &Command{
				Chart: &Chart{
					Operators: []Operator{
						{
							Name: "painter",
							Deployment: Deployment{
								Container: Container{
									Image: Image{
										Tag:        "v0.0.0",
										Repository: "painter",
										Registry:   "quay.io/solo-io",
										PullPolicy: "IfNotPresent",
									},
								},
								Sidecars: []Sidecar{
									{
										Name: sidecarName,
										Container: Container{
											Image: Image{
												Tag:        "v0.0.0",
												Repository: "painter",
												Registry:   "quay.io/solo-io",
												PullPolicy: "IfNotPresent",
											},
										},
									},
								},
							},
						},
					},

					Values: nil,
					Data: Data{
						ApiVersion:  "v1",
						Description: "",
						Name:        "Painting Operator",
						Version:     "v0.0.1",
						Home:        "https://docs.solo.io/skv2/latest",
						Sources: []string{
							"https://github.com/solo-io/skv2",
						},
					},
				},

				ManifestRoot: "codegen/test/chart-sidecar",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			values := helmValuesFromFile("codegen/test/chart-sidecar/values.yaml")
			Expect(values["painter"].(map[string]interface{})["sidecars"].(map[string]interface{})[expectedSidecarParentKey]).ToNot(BeNil())
		},

		Entry("camelCase sidecar name", "fooBar", "fooBar"),
		Entry("hyphened sidecar name", "foo-bar", "fooBar"),
	)

	DescribeTable("rendering extra env vars",
		func(extraEnvs map[string]interface{}, expectedEnvVars []v1.EnvVar) {
			cmd := &Command{
				Chart: &Chart{
					Operators: []Operator{
						{
							Name: "painter",
							Deployment: Deployment{
								Container: Container{
									Image: Image{
										Tag:        "v0.0.0",
										Repository: "painter",
										Registry:   "quay.io/solo-io",
										PullPolicy: "IfNotPresent",
									},
								},
							},
						},
					},

					Values: nil,
					Data: Data{
						ApiVersion:  "v1",
						Description: "",
						Name:        "Painting Operator",
						Version:     "v0.0.1",
						Home:        "https://docs.solo.io/skv2/latest",
						Sources: []string{
							"https://github.com/solo-io/skv2",
						},
					},
				},

				ManifestRoot: "codegen/test/chart-envvars",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			painterValues := map[string]interface{}{"extraEnvs": extraEnvs}
			helmValues := map[string]interface{}{"painter": painterValues}

			renderedManifests := helmTemplate("codegen/test/chart-envvars", helmValues)

			var renderedDeployment *appsv1.Deployment
			decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
			for {
				obj := &unstructured.Unstructured{}
				err := decoder.Decode(obj)
				if err != nil {
					break
				}
				if obj.GetName() != "painter" || obj.GetKind() != "Deployment" {
					continue
				}

				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedDeployment = &appsv1.Deployment{}
				err = json.Unmarshal(bytes, renderedDeployment)
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(renderedDeployment).NotTo(BeNil())
			renderedEnvVars := renderedDeployment.Spec.Template.Spec.Containers[0].Env
			Expect(renderedEnvVars).To(ConsistOf(expectedEnvVars))
		},
		Entry("no env vars",
			nil, nil),
		Entry("value env var",
			map[string]interface{}{"FOO": map[string]interface{}{"value": "bar"}}, []v1.EnvVar{{Name: "FOO", Value: "bar"}}),
		Entry("valueFrom env var",
			map[string]interface{}{"FOO": map[string]interface{}{"valueFrom": map[string]interface{}{"secretKeyRef": map[string]interface{}{"name": "bar", "key": "baz"}}}},
			[]v1.EnvVar{{Name: "FOO", ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{LocalObjectReference: v1.LocalObjectReference{Name: "bar"}, Key: "baz"}}}}),
	)

	It("can configure cluster-scoped and namespace-scoped RBAC", func() {
		cmd := &Command{
			RenderProtos: false,
			Chart: &Chart{
				Operators: []Operator{
					{
						Name:             "painter",
						EnabledDependsOn: []string{"$painter.enabled"},
						ClusterRbac: []rbacv1.PolicyRule{
							{
								Verbs: []string{"GET"},
							},
						},
						NamespaceRbac: map[string][]rbacv1.PolicyRule{
							"secrets": {
								rbacv1.PolicyRule{
									Verbs:     []string{"GET", "LIST", "WATCH"},
									APIGroups: []string{""},
									Resources: []string{"secrets"},
								},
							},
						},
					},
				},
				Values: nil,
				Data: Data{
					ApiVersion:  "v1",
					Description: "",
					Name:        "Painting Operator",
					Version:     "v0.0.1",
					Home:        "https://docs.solo.io/skv2/latest",
					Sources: []string{
						"https://github.com/solo-io/skv2",
					},
				},
			},
			ManifestRoot: "codegen/test/chart",
		}

		Expect(cmd.Execute()).NotTo(HaveOccurred(), "failed to execute command")

		absPath, err := filepath.Abs("./codegen/test/chart/templates/rbac.yaml")
		Expect(err).NotTo(HaveOccurred(), "failed to get abs path")

		rbac, err := os.ReadFile(absPath)
		Expect(err).NotTo(HaveOccurred(), "failed to read rbac.yaml")
		clusterRole1Tmpl := `
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ .Release.Namespace }}
  labels:
    app: painter
rules:
- verbs:
  - GET`
		clusterRoleBinding1Tmpl := `
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ .Release.Namespace }}
  labels:
    app: painter
subjects:
- kind: ServiceAccount
  name: painter
  namespace: {{ .Release.Namespace }}
roleRef:
  kind: ClusterRole
  name: painter-{{ .Release.Namespace }}
  apiGroup: rbac.authorization.k8s.io`
		clusterRole2Tmpl := `
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ .Release.Name }}-{{ .Release.Namespace }}-namespaced
  labels:
    app: painter
rules:
{{- if not (has "secrets" $painterNamespacedResources) }}
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - GET
  - LIST
  - WATCH
{{- end }}`
		clusterRoleBinding2Tmpl := `
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ .Release.Name }}-{{ .Release.Namespace }}-namespaced
  labels:
    app: painter
subjects:
- kind: ServiceAccount
  name: painter
  namespace: {{ .Release.Namespace }}
roleRef:
  kind: ClusterRole
  name: painter-{{ .Release.Name }}-{{ .Release.Namespace }}-namespaced
  apiGroup: rbac.authorization.k8s.io`
		roleTmpl := `
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ $.Release.Name }}-{{ $.Release.Namespace }}-namespaced
  namespace: {{ $ns }}
  labels:
    app: painter
rules:
{{- if (has "secrets" $resources) }}
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - GET
  - LIST
  - WATCH
{{- end }}`
		roleBindingTmpl := `
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ $.Release.Name }}-{{ $.Release.Namespace }}-namespaced
  namespace: {{ $ns }}
  labels:
    app: painter
subjects:
- kind: ServiceAccount
  name: painter
  namespace: {{ $.Release.Namespace }}
roleRef:
  kind: Role
  name: painter-{{ $.Release.Name }}-{{ $.Release.Namespace }}-namespaced
  apiGroup: rbac.authorization.k8s.io`
		Expect(string(rbac)).To(ContainSubstring(clusterRole1Tmpl))
		Expect(string(rbac)).To(ContainSubstring(clusterRoleBinding1Tmpl))
		Expect(string(rbac)).To(ContainSubstring(clusterRole2Tmpl))
		Expect(string(rbac)).To(ContainSubstring(clusterRoleBinding2Tmpl))
		Expect(string(rbac)).To(ContainSubstring(roleTmpl))
		Expect(string(rbac)).To(ContainSubstring(roleBindingTmpl))
	})
})

func helmTemplate(path string, values interface{}) []byte {
	raw, err := yaml.Marshal(values)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	helmValuesFile, err := os.CreateTemp("", "-helm-values-skv2-test")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	_, err = helmValuesFile.Write(raw)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	err = helmValuesFile.Close()
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	defer os.RemoveAll(helmValuesFile.Name())

	out, err := exec.Command("helm", "template",
		path,
		"--values", helmValuesFile.Name(),
	).CombinedOutput()
	ExpectWithOffset(1, err).NotTo(HaveOccurred(), string(out))
	return out
}

func helmValuesFromFile(path string) map[string]interface{} {
	data, err := os.ReadFile(path)
	Expect(err).NotTo(HaveOccurred())

	out := make(map[string]interface{})
	err = yaml.Unmarshal(data, &out)
	Expect(err).NotTo(HaveOccurred())

	return out
}
