package model

import (
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/solo-io/skv2/codegen/collector"
	"github.com/solo-io/skv2/codegen/proto/schemagen"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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

// Mapping from protobuf message name to OpenApi schema
type OpenApiSchemas map[string]*apiextv1.JSONSchemaProps

type Group struct {
	// the group version of the group
	schema.GroupVersion

	// the go  module this group belongs to
	Module string

	// the root directory for generated API code
	ApiRoot string

	// the kinds in the group
	Resources []Resource

	// Should we generate kubernetes manifests?
	RenderManifests bool

	// Should we add the chart version to the generated manifests?
	AddChartVersion string

	// Should we not add the spec hash to the annotation of the generated manifests?
	SkipSpecHash bool

	// Should we generate validating schemas for CRDs?
	RenderValidationSchemas bool

	// Should we exclude descriptions in the validation schemas?
	// Set this to true to make the CRD size smaller.
	SkipSchemaDescriptions bool

	// Should we generate deepcopy functions for non-proto Spec/Status fields?
	RenderFieldJsonDeepcopy bool

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

	// Enable to add //go:generate mockgen directive to the top of generated Go files.
	MockgenDirective bool

	// custom import path to the package
	// containing the Go types
	// use this if you are generating controllers
	// for types in an external project
	CustomTypesImportPath string

	// proto descriptors will be available to the templates if the group was compiled with them.
	Descriptors []*collector.DescriptorWithPath

	// data for providing custom templates to generate custom code for groups
	CustomTemplates []CustomTemplates

	// Mapping from protobuf message name to generated open api structural schema
	// This is populated during skv2 generation by the manifests renderer.
	OpenApiSchemas OpenApiSchemas

	// Custom properties used for custom templates
	Properties

	// Some resources use pointer slices for the Items field.
	PointerSlices bool `default:"false"`

	// Set to true to skip rendering of conditional loading logic
	// for CRDs containing alpha-versioned resources.
	// Used by codegen/templates/manifests/crd.yamltmpl
	SkipConditionalCRDLoading bool

	// Skip generation of crd manifests that live in crd/ directory of a chart
	SkipCRDManifest bool

	// Skip generation of templated crd manifests that live in templates/ dir of a chart
	SkipTemplatedCRDManifest bool
}

type GroupOptions struct {
	// Required when using crds in the templates directory
	EscapeGoTemplateOperators bool

	// Options for generating validation schemas
	SchemaValidationOpts schemagen.ValidationSchemaOptions

	SchemaGenerator schemagen.GeneratorKind
}

func (g Group) HasProtos() bool {
	return len(g.Descriptors) > 0
}

type CustomTemplates struct {
	// the custom templates to run generation on.
	// maps output filename to template text
	Templates map[string]string

	// Enable to add //go:generate mockgen directive to the top of generated Go files.
	MockgenDirective bool

	// custom template funcs which will be inserted into the
	// default template funcmap at rendering time
	Funcs template.FuncMap
}

func (g *Group) InitDescriptors(descriptors []*collector.DescriptorWithPath) {
	g.Descriptors = descriptors
	g.Init()
}

// ensures the resources point to this group
func (g *Group) Init() {
	for i, resource := range g.Resources {
		resource.Group = g
		g.Resources[i] = resource
	}
}

type Resource struct {
	*Group // the group I belong to
	Kind   string
	Spec   Field
	Status *Field

	// Whether or not the resource is cluster-scoped.
	// This is important when rendering the CustomResourceDefinition manifest and RBAC policies.
	ClusterScoped bool

	// Set the short name of the resource
	ShortNames []string

	// Set the categories of the resource
	Categories []string

	// The set of additional printer columns to apply to the CustomResourceDefinition
	AdditionalPrinterColumns []apiextv1.CustomResourceColumnDefinition

	// If enabled, the unmarshal will NOT allow unknown fields.
	StrictUnmarshal bool

	// Corresponds to CRD's versions.storage field
	// Only one version of a resource can be marked as "stored"
	// Set to false by default
	// See https://kubernetes.io/docs/reference/kubernetes-api/extend-resources/custom-resource-definition-v1/#CustomResourceDefinitionSpec
	Stored bool

	// Corresponds to CRD's versions.deprecated field
	// Set to false by default
	// See https://kubernetes.io/docs/reference/kubernetes-api/extend-resources/custom-resource-definition-v1/#CustomResourceDefinitionSpec
	Deprecated bool

	// Optional: if specified, this crd resource definition will be wrapped in the given conditional
	//
	// E.g: `and (.Values.customValueA) (.Values.customValueB)`
	CustomEnableCondition string

	// If the resouce is designated as codegen-only, it will not be rendered in the CRD manifest or included in snapshot generation.
	CodegenOnly bool
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

	/*
		The proto package containing the type, if different than the Group name of the Resource.
		If unset, SKv2 uses the Group name of the Resource that specifies this Type.
	*/
	ProtoPackage string
}

// Properties is an arbitrary set of KV properties which can be associated to an skv2 resource. Can be used in conjunction with custom templates
type Properties map[string]string

func (p Properties) Property(key string) string {
	if p == nil {
		return ""
	}
	return p[key]
}
