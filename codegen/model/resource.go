package model

import (
	"text/template"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GeneratorType string

const (
	GeneratorType_Deepcopy  GeneratorType = "deepcopy"
	GeneratorType_Defaulter GeneratorType = "defaulter"
	GeneratorType_Client    GeneratorType = "client"
	GeneratorType_Lister    GeneratorType = "lister"
	GeneratorType_Informer  GeneratorType = "informer"
	GeneratorType_All       GeneratorType = "all"
)

type GeneratorTypes []GeneratorType

func (g GeneratorTypes) Strings() []string {
	var strs []string
	for _, generatorType := range g {
		strs = append(strs, string(generatorType))
	}
	return strs
}

// returns true if the 'deepcopy' or 'all' generator is present
func (g GeneratorTypes) HasDeepcopy() bool {
	for _, generatorType := range g {
		if generatorType == GeneratorType_Deepcopy || generatorType == GeneratorType_All {
			return true
		}
	}
	return false
}

type Group struct {
	// the group version of the group
	schema.GroupVersion

	// the go  module this group belongs to
	Module string

	// the root directory for generated API code
	ApiRoot string

	// search protos recursively starting from this directory.
	// will default vendor_any if empty
	ProtoDir string

	// the kinds in the group
	Resources []Resource

	// Should we compile protos?
	RenderProtos bool

	// Should we generate kubernetes manifests?
	RenderManifests bool

	// Should we generate kubernetes Go structs?
	RenderTypes bool

	// Should we generate kubernetes Go clients?
	RenderClients bool

	// Deprecated: use generated deepcopy methods instead
	// Should we run kubernetes code generators? (see https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh)
	// Note: if this field is nil and RenderTypes is true,
	// skv2 will run the 'deepcopy' generator by default.
	Generators GeneratorTypes

	// Should we generate kubernetes Go controllers?
	RenderController bool

	// custom import path to the package
	// containing the Go types
	// use this if you are generating controllers
	// for types in an external project
	CustomTypesImportPath string

	// proto descriptors will be available to the templates if the group was compiled with them.
	Descriptors []*model.DescriptorWithPath

	// data for providing custom templates to generate custom code for groups
	CustomTemplates CustomTemplates

	RenderContrib RenderContrib
}

type CustomTemplates struct {
	// the custom templates to run generation on.
	// maps output filename to template text
	Templates map[string]string

	// custom data that can be provided for use with custom templates
	Data interface{}

	// custom template funcs which will be inserted into the
	// default template funcmap at rendering time
	Funcs template.FuncMap
}

type RenderContrib struct {
	// set data structure
	Sets bool

	// gomega matchers for set data structures
	SetMatchers bool
}

// ensures the resources point to this group
func (g *Group) Init() {
	for i, resource := range g.Resources {
		resource.Group = *g
		g.Resources[i] = resource
	}
}

type Resource struct {
	Group  // the group I belong to
	Kind   string
	Spec   Field
	Status *Field

	// Whether or not the resource is cluster-scoped.
	// This is important when rendering the CustomResourceDefinition manifest.
	ClusterScoped bool

	// The set of additional printer columns to apply to the CustomResourceDefinition
	AdditionalPrinterColumns []apiextv1beta1.CustomResourceColumnDefinition
}

type Field struct {
	Type Type
}

type Type struct {
	// name of the type.
	Name string

	// proto message for the type, if the proto message is compiled with skv2
	Message proto.Message

	/*

		The go package containing the type, if different than group root api directory (where the resource itself lives).
		Will be set automatically for proto-based types.

		If unset, SKv2 uses the default types package for the type.
	*/
	GoPackage string
}
