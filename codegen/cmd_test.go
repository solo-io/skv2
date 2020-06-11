package codegen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
	"k8s.io/apimachinery/pkg/runtime/schema"

	. "github.com/solo-io/skv2/codegen"
)

var _ = Describe("Cmd", func() {
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
					RenderProtos:     true,
					RenderManifests:  true,
					RenderTypes:      true,
					RenderClients:    true,
					RenderController: true,
					MockgenDirective: true,
					ApiRoot:          "codegen/test/api",
					CustomTemplates:  contrib.AllCustomTemplates,
				},
			},
			AnyVendorConfig: &sk_anyvendor.Imports{
				Local: []string{"codegen/test/*.proto"},
			},

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
