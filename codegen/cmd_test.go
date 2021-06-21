package codegen_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/solo-io/skv2/codegen/render"

	"sigs.k8s.io/yaml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/skv2_anyvendor"
	"github.com/solo-io/skv2/contrib"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubeyaml "k8s.io/apimachinery/pkg/util/yaml"

	. "github.com/solo-io/skv2/codegen"
)

var _ = Describe("Cmd", func() {
	render.ProjectUnstructuredFields = map[string][][]string{
		"github.com/solo-io/skv2/codegen/test/api/things.test.io/v1": {
			{"PaintSpec", "recursiveType", "recursiveField"},
			{"PaintSpec", "recursiveType", "repeatedRecursiveField"},
		},
	}

	skv2Imports := skv2_anyvendor.CreateDefaultMatchOptions(
		[]string{"codegen/test/*.proto"},
	)

	// used for a proto option which disables openapi validation on fields
	skv2Imports.External["cuelang.org/go"] = []string{
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
							Image: Image{
								Tag:        "v0.0.0",
								Repository: "painter",
								Registry:   "quay.io/solo-io",
								PullPolicy: "IfNotPresent",
							},
						},
						Args: []string{"foo"},
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

	It("supports overriding the deployment and service specs via helm values", func() {

		cmd := &Command{
			Chart: &Chart{
				Operators: []Operator{
					{
						Name: "painter",
						Deployment: Deployment{
							Image: Image{
								Tag:        "v0.0.0",
								Repository: "painter",
								Registry:   "quay.io/solo-io",
								PullPolicy: "IfNotPresent",
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
						Args: []string{"foo"},
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

		helmValues := map[string]interface{}{
			"painter": map[string]interface{}{
				"extraPodLabels": map[string]string{
					"extrapod": "labels",
				},
				"extraPodAnnotations": map[string]string{
					"extrapod": "annotations",
				},
				"extraDeploymentLabels": map[string]string{
					"extradeployment": "labels",
				},
				"extraDeploymentAnnotations": map[string]string{
					"extradeployment": "annotations",
				},
				"extraServiceLabels": map[string]string{
					"extraservice": "labels",
				},
				"extraServiceAnnotations": map[string]string{
					"extraservice": "annotations",
				},
				"serviceOverrides": map[string]interface{}{
					"spec": map[string]interface{}{
						"loadBalancerIp": loadBalancerIp,
					},
				},
				"deploymentOverrides": map[string]interface{}{
					"spec": map[string]interface{}{
						"replicas": replicas,
					},
				},
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
	It("generates CRD with validaton schema for a proto file", func() {

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
							Image: Image{
								Tag:        "v0.0.0",
								Repository: "painter",
								Registry:   "quay.io/solo-io",
								PullPolicy: "IfNotPresent",
							},
						},
						Args: []string{"foo"},
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
