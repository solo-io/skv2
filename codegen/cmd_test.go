package codegen_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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
								ReadinessProbe: &v1.Probe{
									Handler: v1.Handler{
										HTTPGet: &v1.HTTPGetAction{
											Path: "/",
											Port: intstr.FromInt(8080),
										},
									},
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
											Handler: v1.Handler{
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
								ReadinessProbe: &v1.Probe{
									Handler: v1.Handler{
										HTTPGet: &v1.HTTPGetAction{
											Path: "/",
											Port: intstr.FromInt(8080),
										},
									},
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
											Handler: v1.Handler{
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
									ReadinessProbe: &v1.Probe{
										Handler: v1.Handler{
											HTTPGet: &v1.HTTPGetAction{
												Path: "/",
												Port: intstr.FromInt(8080),
											},
										},
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
												Handler: v1.Handler{
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

			bytes, err := ioutil.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bytes)).To(ContainSubstring("description: OpenAPI gen test for recursive fields"))
		})

		It("can exclude field descriptions", func() {
			// write this manifest to a different dir to avoid modifying the crd file from the
			// above test, which other tests seem to depend on
			cmd.ManifestRoot = "codegen/test/chart-no-desc"
			crdFilePath := filepath.Join(util.GetModuleRoot(), cmd.ManifestRoot, "/crds/things.test.io_v1_crds.yaml")

			cmd.Groups[0].SkipSchemaDescriptions = true

			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())

			bytes, err := ioutil.ReadFile(crdFilePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bytes)).NotTo(ContainSubstring("description:"))
		})
	})
})

func helmTemplate(path string, values interface{}) []byte {
	raw, err := yaml.Marshal(values)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	helmValuesFile, err := ioutil.TempFile("", "-helm-values-skv2-test")
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
