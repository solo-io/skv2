package codegen_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	goyaml "gopkg.in/yaml.v3"
	rbacv1 "k8s.io/api/rbac/v1"
	v12 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/utils/pointer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
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
	It("install conditional sidecars", func() {
		agentConditional := "and ($.Values.glooAgent.enabled) ($.Values.glooAgent.runAsSidecar)"

		cmd := &Command{
			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "gloo-mgmt-server",
						Service: Service{
							Ports: []ServicePort{{
								Name:        "grpc",
								DefaultPort: 9900,
							}},
						},
						ClusterRbac: []rbacv1.PolicyRule{{
							Verbs:     []string{"*"},
							APIGroups: []string{"coordination.k8s.io"},
							Resources: []string{"leases"},
						}},
						Deployment: Deployment{
							Sidecars: []Sidecar{
								{
									Name: "gloo-agent",
									Volumes: []v1.Volume{
										{
											Name: "agent-volume",
											VolumeSource: v1.VolumeSource{
												Secret: &v1.SecretVolumeSource{
													SecretName: "agent-volume",
												},
											},
										},
										{
											Name: "agent-volume-2",
											VolumeSource: v1.VolumeSource{
												Secret: &v1.SecretVolumeSource{
													SecretName: "agent-volume",
												},
											},
										},
									},
									ClusterRbac: []rbacv1.PolicyRule{{
										Verbs:     []string{"*"},
										APIGroups: []string{"apiextensions.k8s.io"},
										Resources: []string{"customresourcedefinitions"},
									}},
									Container: Container{
										Image: Image{
											Registry:   "gcr.io/gloo-mesh",
											Repository: "gloo-mesh-agent",
											Tag:        "0.0.1",
										},
									},
									Service: Service{
										Ports: []ServicePort{{
											Name:        "grpc",
											DefaultPort: 9977,
										}},
									},
									EnableStatement: agentConditional,
									ValuesPath:      "$.Values.glooAgent",
								},
							},
							Container: Container{
								Image: Image{
									Registry:   "gcr.io/gloo-mesh",
									Repository: "gloo-mesh-mgmt-server",
									Tag:        "0.0.1",
								},
								ContainerPorts: []ContainerPort{{
									Name: "stats",
									Port: "{{ $Values.glooMgmtServer.statsPort }}",
								}},
								VolumeMounts: []v1.VolumeMount{{
									Name:      "license-keys",
									MountPath: "/etc/gloo-mesh/license-keys",
									ReadOnly:  true,
								}},
							},
							Volumes: []v1.Volume{
								{
									Name: "license-keys",
									VolumeSource: v1.VolumeSource{
										Secret: &v1.SecretVolumeSource{
											SecretName: "license-keys",
										},
									},
								},
							},
						},
					},
					{
						Name:                  "gloo-agent",
						CustomEnableCondition: `and ($.Values.glooAgent.enabled) (not $.Values.glooAgent.runAsSidecar)`,
					},
				},
			},
			ManifestRoot: "codegen/test/chart/conditional-sidecar",
		}

		Expect(cmd.Execute()).NotTo(HaveOccurred(), "failed to execute command")

		absPath, err := filepath.Abs("./test/chart/conditional-sidecar/templates/deployment.yaml")
		Expect(err).NotTo(HaveOccurred(), "failed to get abs path")

		deployment, err := os.ReadFile(absPath)
		Expect(err).NotTo(HaveOccurred(), "failed to read deployment.yaml")

		Expect(deployment).To(ContainSubstring(fmt.Sprintf("{{ if %s }}", agentConditional)))
		Expect(deployment).To(ContainSubstring(fmt.Sprintf("{{ if %s }}", "and ($.Values.glooAgent.enabled) (not $.Values.glooAgent.runAsSidecar)")))
		Expect(deployment).To(ContainSubstring("name: agent-volume"))
		Expect(deployment).To(ContainSubstring(`{{ index $glooAgent "ports" "grpc" }}`))
		Expect(deployment).To(ContainSubstring("{{ $Values.glooMgmtServer.statsPort }}"))
	})
	It("generates conditional crds", func() {
		cmd := &Command{
			Groups: []Group{{
				GroupVersion: schema.GroupVersion{
					Group:   "things.test.io",
					Version: "v1",
				},
				Resources: []Resource{{
					Kind:   "Paint",
					Spec:   Field{Type: Type{Name: "PaintSpec"}},
					Status: &Field{Type: Type{Name: "PaintStatus"}},
					// Stored:                true,
					CustomEnableCondition: "$.Values.installValue",
				}},
				RenderManifests: true,
			}},
			SkipCrdsManifest: true, // Use templates folder to install crds conditionally
			Chart: &Chart{
				Values: map[string]interface{}{
					"installValue": true,
				},
			},
			ManifestRoot: "codegen/test/chart/conditional-crds",
		}
		Expect(cmd.Execute()).NotTo(HaveOccurred(), "failed to execute command")

		crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/templates/things.test.io_crds.yaml")

		bytes, err := os.ReadFile(crdFilePath)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bytes)).To(ContainSubstring("{{- if $.Values.installValue }}"))
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
							Stored: true,
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
							Stored:        true,
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
							Stored: true,
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
							Stored: true,
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
							Stored:        true,
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
		Expect(enabledMapField.HeadComment).To(Equal("# Arbitrary overrides for the component's [deployment\n# template](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/)."))
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
							Stored: true,
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
							Stored:        true,
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
							Stored: true,
						},
						{
							Kind:          "ClusterResource",
							Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
							ClusterScoped: true,
							Stored:        true,
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
			"app.kubernetes.io/name": "painter",
			"extrapod":               "annotations",
			"pod":                    "annotations",
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
						Name:                  "painter",
						CustomEnableCondition: "and $painter.enabled $.Values.test1.enabled $.Values.test2.enabled",
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

		Expect(string(fileContents)).To(ContainSubstring("{{ if and $painter.enabled $.Values.test1.enabled $.Values.test2.enabled }}"))

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

	It("supports rendering namespace template with custom value if provided", func() {
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
						Name:                   "painter",
						NamespaceFromValuePath: "$.Values.common.namespace",
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

		// Test that the deployment, SA, and CR templates are rendered correctly
		deployment, err := os.ReadFile("codegen/test/chart/templates/deployment.yaml")
		Expect(err).NotTo(HaveOccurred())

		expectedDeploymentTmpl := `
kind: Deployment
metadata:
  labels:
    app: painter
  annotations:
    app.kubernetes.io/name: painter
  name: painter
  namespace: {{ $.Values.common.namespace | default $.Release.Namespace }}`

		expectedSATmpl := `
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: painter
  {{- if $painter.serviceAccount}}
  {{- if $painter.serviceAccount.extraAnnotations }}
  annotations:
    {{- range $key, $value := $painter.serviceAccount.extraAnnotations }}
    {{ $key }}: {{ $value }}
    {{- end }}
  {{- end }}
  {{- end}}
  name: painter
  namespace: {{ $.Values.common.namespace | default $.Release.Namespace }}`

		Expect(string(deployment)).To(ContainSubstring(expectedDeploymentTmpl))
		Expect(string(deployment)).To(ContainSubstring(expectedSATmpl))

		rbac, err := os.ReadFile("codegen/test/chart/templates/rbac.yaml")
		Expect(err).NotTo(HaveOccurred())

		expectedClusterRoleTmpl := `
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ $.Values.common.namespace | default $.Release.Namespace }}`

		expectedClusterRoleBindingTmpl := `
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ $.Values.common.namespace | default $.Release.Namespace }}
  labels:
    app: painter
subjects:
- kind: ServiceAccount
  name: painter
  namespace: {{ $.Values.common.namespace | default $.Release.Namespace }}
roleRef:
  kind: ClusterRole
  name: painter-{{ $.Values.common.namespace | default $.Release.Namespace }}
  apiGroup: rbac.authorization.k8s.io`

		Expect(string(rbac)).To(ContainSubstring(expectedClusterRoleTmpl))
		Expect(string(rbac)).To(ContainSubstring(expectedClusterRoleBindingTmpl))

		// Test that the deployment, SA, and CR are rendered with the custom namespace value
		helmValues := map[string]interface{}{
			"common": map[string]interface{}{
				"namespace": "test-namespace",
			},
		}

		expectedDeployment := `
kind: Deployment
metadata:
  labels:
    app: painter
  annotations:
    app.kubernetes.io/name: painter
  name: painter
  namespace: test-namespace`

		expectedSA := `
kind: ServiceAccount
metadata:
  labels:
    app: painter
  name: painter
  namespace: test-namespace`

		expectedClusterRole := `
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-test-namespace`

		expectedClusterRoleBinding := `
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-test-namespace
  labels:
    app: painter
subjects:
- kind: ServiceAccount
  name: painter
  namespace: test-namespace
roleRef:
  kind: ClusterRole
  name: painter-test-namespace
  apiGroup: rbac.authorization.k8s.io`

		renderedManifests := helmTemplate("codegen/test/chart", helmValues)
		Expect(renderedManifests).To(ContainSubstring(expectedSA))
		Expect(renderedManifests).To(ContainSubstring(expectedDeployment))
		Expect(renderedManifests).To(ContainSubstring(expectedClusterRole))
		Expect(renderedManifests).To(ContainSubstring(expectedClusterRoleBinding))
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
								Stored: true,
							},
							{
								Kind:          "ClusterResource",
								Spec:          Field{Type: Type{Name: "ClusterResourceSpec"}},
								ClusterScoped: true,
								Stored:        true,
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
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_crds.yaml")

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			bytes, err := os.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bytes)).To(ContainSubstring("description: OpenAPI gen test for recursive fields"))
		})

		// TODO (dmitri-d): kube_crud_test and kube_multicluster_test depend on crds in this suite.
		It("generates google.protobuf.Value with no type", func() {
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_crds.yaml")

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			bytes, err := os.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			paintCrdYaml := ""
			for _, crd := range strings.Split(string(bytes), "---") {
				if strings.Contains(crd, "kind: Paint") {
					paintCrdYaml = crd
				}
			}
			Expect(paintCrdYaml).ToNot(BeEmpty())

			generatedCrd := &v12.CustomResourceDefinition{}
			Expect(yaml.Unmarshal([]byte(paintCrdYaml), &generatedCrd)).NotTo(HaveOccurred())
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
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_crds.yaml")

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

	DescribeTable("rendering deployment strategy",
		func(deploymentStrategy *appsv1.DeploymentStrategy) {
			cmd := &Command{
				Chart: &Chart{
					Operators: []Operator{
						{
							Name: "painter",
							Deployment: Deployment{
								Strategy: deploymentStrategy,
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

				ManifestRoot: "codegen/test/chart-deployment-strategy",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			values := map[string]interface{}{}
			helmValues := map[string]interface{}{"painter": values}

			renderedManifests := helmTemplate("codegen/test/chart-deployment-strategy", helmValues)

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
			renderedDeploymentStrategy := renderedDeployment.Spec.Strategy
			if deploymentStrategy == nil {
				Expect(renderedDeploymentStrategy).To(Equal(appsv1.DeploymentStrategy{}))
			} else {
				Expect(renderedDeploymentStrategy).To(Equal(*deploymentStrategy))
			}
		},
		Entry("when the deployment strategy is not defined",
			nil),
		Entry("when the deployment strategy is configured to recreate",
			&appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			}),
		Entry("when the deployment strategy is configured to rolling update",
			&appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						IntVal: 1,
					},
				},
			}),
	)

	DescribeTable("rendering pod security context",
		func(podSecurityContextValues map[string]interface{}, podSecurityContext *v1.PodSecurityContext, expectedPodSecurityContext *v1.PodSecurityContext) {
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
								PodSecurityContext: podSecurityContext,
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

				ManifestRoot: "codegen/test/chart-pod-security-context",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			values := map[string]interface{}{"podSecurityContext": podSecurityContextValues}
			helmValues := map[string]interface{}{"painter": values}

			renderedManifests := helmTemplate("codegen/test/chart-pod-security-context", helmValues)

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
			renderedPodSecurityContext := renderedDeployment.Spec.Template.Spec.SecurityContext
			Expect(renderedPodSecurityContext).To(Equal(expectedPodSecurityContext))
		},
		Entry("when PodSecurityContext is neither defined in values nor in the operator",
			nil,
			nil,
			nil),
		Entry("when PodSecurityContext is defined only in values",
			map[string]interface{}{"fsGroup": 1000},
			nil,
			&v1.PodSecurityContext{
				FSGroup: pointer.Int64(1000),
			}),
		Entry("when PodSecurityContext is defined only in the operator",
			nil,
			&v1.PodSecurityContext{
				FSGroup: pointer.Int64(1000),
			},
			&v1.PodSecurityContext{
				FSGroup: pointer.Int64(1000),
			}),
		Entry("when PodSecurityContext is defined in both values and the operator",
			map[string]interface{}{"fsGroup": 1024},
			&v1.PodSecurityContext{
				FSGroup: pointer.Int64(1000),
			},
			&v1.PodSecurityContext{
				FSGroup: pointer.Int64(1024), // should override the value defined in the operator
			}),
	)

	Describe("rendering template env vars", func() {
		var tmpDir string

		BeforeEach(func() {
			tmpDir = "codegen/test/chart-featureenv"
		})

		AfterEach(func() {
			Expect(os.RemoveAll(tmpDir)).NotTo(HaveOccurred())
		})

		DescribeTable("validation",
			func(featureGatesVals map[string]string, envs []v1.EnvVar, featureEnvs []TemplateEnvVar, expectedEnvVars []v1.EnvVar) {
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
										Env:             envs,
										TemplateEnvVars: featureEnvs,
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

					ManifestRoot: tmpDir,
				}

				err := cmd.Execute()
				Expect(err).NotTo(HaveOccurred())

				painterValues := map[string]interface{}{}
				// featureGates := map[string]interface{}{"Foo": true}
				helmValues := map[string]interface{}{"painter": painterValues, "featureGates": featureGatesVals}

				renderedManifests := helmTemplate(tmpDir, helmValues)

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

			Entry("when neither Env nor TemplateEnvVar is specified",
				map[string]string{"Foo": "true"},
				nil,
				nil,
				nil),
			Entry("when Env is not specified and TemplateEnvVar is specified",
				map[string]string{"Foo": "true"},
				nil,
				[]TemplateEnvVar{
					{
						Name:  "FEATURE_ENABLE_FOO",
						Value: "{{ $.Values.featureGates.Foo | quote }}",
					},
				},
				nil),
			Entry("when Env and TemplateEnvVar are specified, true value",
				map[string]string{"Foo": "true"},
				[]v1.EnvVar{
					{
						Name:  "FOO",
						Value: "bar",
					},
				},
				[]TemplateEnvVar{
					{
						Name:  "FEATURE_ENABLE_FOO",
						Value: "{{ $.Values.featureGates.Foo | quote }}",
					},
				},
				[]v1.EnvVar{
					{Name: "FOO", Value: "bar"},
					{Name: "FEATURE_ENABLE_FOO", Value: "true"},
				}),
			Entry("when Env and TemplateEnvVar are specified, false value",
				map[string]string{"Foo": "false"},
				[]v1.EnvVar{
					{
						Name:  "FOO",
						Value: "bar",
					},
				},
				[]TemplateEnvVar{
					{
						Name:  "FEATURE_ENABLE_FOO",
						Value: "{{ $.Values.featureGates.Foo | quote }}",
					},
				},
				[]v1.EnvVar{
					{Name: "FOO", Value: "bar"},
					{Name: "FEATURE_ENABLE_FOO", Value: "false"},
				}),
			Entry("when Env and Conditional TemplateEnvVar are specified, and condition is true",
				map[string]string{"Foo": "false"},
				[]v1.EnvVar{
					{
						Name:  "FOO",
						Value: "bar",
					},
				},
				[]TemplateEnvVar{
					{
						Condition: "$.Values.featureGates.Foo",
						Name:      "FEATURE_ENABLE_FOO",
						Value:     "{{ $.Values.featureGates.Foo | quote }}",
					},
				},
				[]v1.EnvVar{
					{Name: "FOO", Value: "bar"},
					{Name: "FEATURE_ENABLE_FOO", Value: "false"},
				}),
			Entry("when Env and Conditional TemplateEnvVar are specified, and condition is false",
				map[string]string{"Foo": "false"},
				[]v1.EnvVar{
					{
						Name:  "FOO",
						Value: "bar",
					},
				},
				[]TemplateEnvVar{
					{
						Condition: "$.Values.featureGates.InvalidCondition",
						Name:      "FEATURE_ENABLE_FOO",
						Value:     "{{ $.Values.featureGates.Foo | quote }}",
					},
				},
				[]v1.EnvVar{
					{Name: "FOO", Value: "bar"},
				}),
		)
	})

	Describe("rendering conditional volumes", func() {
		var tmpDir string

		BeforeEach(func() {
			tmpDir = "codegen/test/chart-volumes"
		})

		AfterEach(func() {
			Expect(os.RemoveAll(tmpDir)).NotTo(HaveOccurred())
		})

		DescribeTable("validation",
			func(values map[string]string, defaultVolumes []v1.Volume, conditionalVolumes []model.ConditionalVolume, expected []v1.Volume) {
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
									Volumes:            defaultVolumes,
									ConditionalVolumes: conditionalVolumes,
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

					ManifestRoot: tmpDir,
				}

				err := cmd.Execute()
				Expect(err).NotTo(HaveOccurred())

				// featureGates := map[string]interface{}{"Foo": true}
				helmValues := map[string]interface{}{"painter": values}

				renderedManifests := helmTemplate(tmpDir, helmValues)

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
					break
				}
				Expect(renderedDeployment).NotTo(BeNil())
				Expect(renderedDeployment.Spec.Template.Spec.Volumes).To(ConsistOf(expected))
			},

			Entry("no volumes or conditional volumes",
				map[string]string{},
				nil,
				nil,
				nil,
			),
			Entry("with default volume",
				map[string]string{},
				[]v1.Volume{
					{
						Name: "vol-1",
					},
				},
				nil,
				[]v1.Volume{
					{
						Name: "vol-1",
					},
				},
			),
			Entry("with conditional volume when condition is true",
				map[string]string{
					"condition": "true",
				},
				nil,
				[]model.ConditionalVolume{
					{
						Condition: "$.Values.painter.condition",
						Volume: v1.Volume{
							Name: "vol-1",
						},
					},
				},
				[]v1.Volume{
					{
						Name: "vol-1",
					},
				},
			),
			Entry("with conditional volume when condition is false",
				map[string]string{
					"condition": "true",
				},
				nil,
				[]model.ConditionalVolume{
					{
						Condition: "$.Values.painter.invalidCondition",
						Volume: v1.Volume{
							Name: "vol-1",
						},
					},
				},
				nil,
			),
			Entry("with default and conditional volume when condition is true",
				map[string]string{
					"condition": "true",
				},
				[]v1.Volume{
					{
						Name: "vol-1",
					},
				},
				[]model.ConditionalVolume{
					{
						Condition: "$.Values.painter.condition",
						Volume: v1.Volume{
							Name: "vol-2",
						},
					},
				},
				[]v1.Volume{
					{
						Name: "vol-1",
					},
					{
						Name: "vol-2",
					},
				},
			),
		)
	})

	Describe("rendering conditional volumeMounts", func() {
		var tmpDir string

		BeforeEach(func() {
			tmpDir = "codegen/test/chart-volumeMounts"
		})

		AfterEach(func() {
			Expect(os.RemoveAll(tmpDir)).NotTo(HaveOccurred())
		})

		DescribeTable("validation",
			func(values map[string]string, defaultMounts []v1.VolumeMount, conditionalMounts []model.ConditionalVolumeMount, expected []v1.VolumeMount) {
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
										VolumeMounts:            defaultMounts,
										ConditionalVolumeMounts: conditionalMounts,
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

					ManifestRoot: tmpDir,
				}

				err := cmd.Execute()
				Expect(err).NotTo(HaveOccurred())

				// featureGates := map[string]interface{}{"Foo": true}
				helmValues := map[string]interface{}{"painter": values}

				renderedManifests := helmTemplate(tmpDir, helmValues)

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
					break
				}
				Expect(renderedDeployment).NotTo(BeNil())
				containers := renderedDeployment.Spec.Template.Spec.Containers
				Expect(containers).To(HaveLen(1))
				Expect(containers[0].VolumeMounts).To(ConsistOf(expected))
			},

			Entry("no volumes or conditional mounts",
				map[string]string{},
				nil,
				nil,
				nil,
			),
			Entry("with default volume mount",
				map[string]string{},
				[]v1.VolumeMount{
					{
						Name: "vol-1",
					},
				},
				nil,
				[]v1.VolumeMount{
					{
						Name: "vol-1",
					},
				},
			),
			Entry("with conditional volume mount when condition is true",
				map[string]string{
					"condition": "true",
				},
				nil,
				[]model.ConditionalVolumeMount{
					{
						Condition: "$.Values.painter.condition",
						VolumeMount: v1.VolumeMount{
							Name: "vol-1",
						},
					},
				},
				[]v1.VolumeMount{
					{
						Name: "vol-1",
					},
				},
			),
			Entry("with conditional volume mount when condition is false",
				map[string]string{
					"condition": "true",
				},
				nil,
				[]model.ConditionalVolumeMount{
					{
						Condition: "$.Values.painter.invalidCondition",
						VolumeMount: v1.VolumeMount{
							Name: "vol-1",
						},
					},
				},
				nil,
			),
			Entry("with default and conditional volume mounts when condition is true",
				map[string]string{
					"condition": "true",
				},
				[]v1.VolumeMount{
					{
						Name: "vol-1",
					},
				},
				[]model.ConditionalVolumeMount{
					{
						Condition: "$.Values.painter.condition",
						VolumeMount: v1.VolumeMount{
							Name: "vol-2",
						},
					},
				},
				[]v1.VolumeMount{
					{
						Name: "vol-1",
					},
					{
						Name: "vol-2",
					},
				},
			),
		)
	})

	DescribeTable("rendering service ports",
		func(portName string) {
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
							Service: Service{
								Ports: []ServicePort{{
									Name:        portName,
									DefaultPort: 9900,
								}},
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

				ManifestRoot: "codegen/test/chart-svcport",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			helmValues := map[string]interface{}{}

			renderedManifests := helmTemplate("codegen/test/chart-svcport", helmValues)

			var renderedSvc *v1.Service
			decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
			for {
				obj := &unstructured.Unstructured{}
				err := decoder.Decode(obj)
				if err != nil {
					break
				}
				if obj.GetName() != "painter" || obj.GetKind() != "Service" {
					continue
				}

				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedSvc = &v1.Service{}
				err = json.Unmarshal(bytes, renderedSvc)
				Expect(err).NotTo(HaveOccurred())
				break
			}
			Expect(renderedSvc).NotTo(BeNil())
			Expect(renderedSvc.Spec.Ports[0].Name).To(Equal(portName))
		},
		Entry("port name without hyphen", "foo"),
		Entry("port name with hyphen", "foo-bar"),
	)

	DescribeTable("rendering conditional sidecars with service ports",
		func(portName string) {
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
										Name:            "sidecar-painter",
										EnableStatement: "true",
										Container: Container{
											Image: Image{
												Tag:        "v0.0.0",
												Repository: "painter",
												Registry:   "quay.io/solo-io",
												PullPolicy: "IfNotPresent",
											},
										},
										Service: Service{
											Ports: []ServicePort{
												{
													Name:        portName,
													DefaultPort: 1337,
												},
											},
										},
									},
								},
							},
							Service: Service{
								Ports: []ServicePort{{
									Name:        portName,
									DefaultPort: 9900,
								}},
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

				ManifestRoot: "codegen/test/chart-sidecar-svcport",
			}

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			// alternatively could be custom values file; if this needs to expanded upon consider moving out of code
			helmValues := map[string]interface{}{
				"painter": map[string]interface{}{
					"sidecars": map[string]interface{}{
						"sidecarPainter": map[string]interface{}{
							"ports": map[string]interface{}{
								portName: 1337,
							},
						},
					},
				},
			}

			renderedManifests := helmTemplate("codegen/test/chart-sidecar-svcport", helmValues)

			var renderedSvc *v1.Service
			decoder := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewBuffer(renderedManifests), 4096)
			for {
				obj := &unstructured.Unstructured{}
				err := decoder.Decode(obj)
				if err != nil {
					break
				}
				if obj.GetName() != "sidecar-painter" || obj.GetKind() != "Service" {
					continue
				}

				bytes, err := obj.MarshalJSON()
				Expect(err).NotTo(HaveOccurred())
				renderedSvc = &v1.Service{}
				err = json.Unmarshal(bytes, renderedSvc)
				Expect(err).NotTo(HaveOccurred())
				break
			}
			Expect(renderedSvc).NotTo(BeNil())
			Expect(renderedSvc.Spec.Ports[0].Name).To(Equal(portName))
			Expect(renderedSvc.Spec.Ports[0].Port).To(Equal(int32(1337)))
		},
		Entry("sidecar service port name without hyphen", "foo"),
		Entry("sidecar service port name with hyphen", "foo-bar"),
	)

	It("render readiness probe when scheme is specified", func() {
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
								ReadinessProbe: &ReadinessProbe{
									Path:                "/",
									Port:                "8080",
									Scheme:              "HTTPS",
									PeriodSeconds:       10,
									InitialDelaySeconds: 5,
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

			ManifestRoot: "codegen/test/chart-readiness",
		}

		err := cmd.Execute()
		Expect(err).NotTo(HaveOccurred())

		helmValues := map[string]interface{}{}

		renderedManifests := helmTemplate("codegen/test/chart-readiness", helmValues)

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
		renderedReadinessProbe := renderedDeployment.Spec.Template.Spec.Containers[0].ReadinessProbe.HTTPGet
		Expect(string(renderedReadinessProbe.Scheme)).To(Equal("HTTPS"))
		Expect(int(renderedReadinessProbe.Port.IntVal)).To(Equal(8080))
	})

	It("can configure cluster-scoped and namespace-scoped RBAC", func() {
		cmd := &Command{
			RenderProtos: false,
			Chart: &Chart{
				Operators: []Operator{
					{
						Name:                  "painter",
						CustomEnableCondition: "$painter.enabled",
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
  name: painter-{{ default .Release.Namespace $painter.namespace }}
  labels:
    app: painter
rules:
- verbs:
  - GET`
		clusterRoleBinding1Tmpl := `
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: painter-{{ default .Release.Namespace $painter.namespace }}
  labels:
    app: painter
subjects:
- kind: ServiceAccount
  name: painter
  namespace: {{ default .Release.Namespace $painter.namespace }}
roleRef:
  kind: ClusterRole
  name: painter-{{ default .Release.Namespace $painter.namespace }}
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
  namespace: {{ default .Release.Namespace $painter.namespace }}
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
  namespace: {{ default $.Release.Namespace $painter.namespace }}
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
