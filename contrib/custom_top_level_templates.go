package contrib

import (
	"text/template"

	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib/codegen/funcs"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

/*
Define custom templates that live on the group level here.
*/

// Parameters for constructing templates that span across multiple Groups.
type SnapshotTemplateParameters struct {
	// the path of the output file produced from this template. relative to project root.
	OutputFilename string

	// a map of Go modules to (a superset of) the imported codegen Groups. only required if the codegen group is defined in a different go module than the types (i.e. it is using a CustomTypesImportPath)
	SelectFromGroups map[string][]model.Group

	SnapshotResources SnapshotResources
}

// SnapshotResources acts as a "oneof" to encapsulate
// HybridSnapshot or a HomogenousSnapshot
type SnapshotResources interface {
	makeTemplateFuncs(outputFilename string, selectFromGroups map[string][]model.Group) template.FuncMap
}

// HomogenousSnapshotResources represents a set of snapshot resources read from a single source (either remote clusters or local cluster)
type HomogenousSnapshotResources struct {
	// a map of the GVKs to the resources which we want to include in the input snapshot.
	ResourcesToSelect map[schema.GroupVersion][]string
}

func (r HomogenousSnapshotResources) makeTemplateFuncs(outputFilename string, selectFromGroups map[string][]model.Group) template.FuncMap {
	return funcs.MakeHomogenousSnapshotFuncs(
		outputFilename,
		selectFromGroups,
		r.ResourcesToSelect,
	)
}

// HybridSnapshotResources represents a set of snapshot resources read from a both source remote clusters and the local cluster
type HybridSnapshotResources struct {
	// a map of the GVKs to the resources which we want to include in the input snapshot.
	// these resourecs will be stored in the the local cluster where the controller is running.
	LocalResourcesToSelect map[schema.GroupVersion][]string

	// a map of the GVKs to the resources which we want to include in the input snapshot.
	// these resourecs will be stored in the the remote cluster which the controller manages.
	RemoteResourcesToSelect map[schema.GroupVersion][]string
}

func (r HybridSnapshotResources) makeTemplateFuncs(outputFilename string, selectFromGroups map[string][]model.Group) template.FuncMap {
	return funcs.MakeHybridSnapshotFuncs(
		outputFilename,
		selectFromGroups,
		r.LocalResourcesToSelect,
		r.RemoteResourcesToSelect,
	)
}

func (p SnapshotTemplateParameters) constructTemplate(params SnapshotTemplateParameters, templatePath string) model.CustomTemplates {
	templateContents, err := templatesBox.FindString(templatePath)
	if err != nil {
		panic(err)
	}

	crossGroupTemplate := model.CustomTemplates{
		Templates:        map[string]string{params.OutputFilename: templateContents},
		MockgenDirective: true,
		Funcs:            p.SnapshotResources.makeTemplateFuncs(p.OutputFilename, p.SelectFromGroups),
	}

	return crossGroupTemplate
}

// NOTE(ilackarms): to add your template, copy paste XXX custom template below

/*
InputSnapshot custom template
*/
const (
	InputSnapshotCustomTemplatePath = "input/input_snapshot.gotmpl"
)

// Returns the template for generating input snapshots.
func InputSnapshot(params SnapshotTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, InputSnapshotCustomTemplatePath)
}

/*
InputSnapshot test builder custom template
*/
const (
	InputSnapshotManualBuilderCustomTemplatePath = "input/input_snapshot_manual_builder.gotmpl"
)

// Returns the template for generating input snapshots.
func InputSnapshotManualBuilder(params SnapshotTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, InputSnapshotManualBuilderCustomTemplatePath)
}

/*
InputReconciler custom templates
*/
const (
	HomogenousInputReconcilerCustomTemplatePath = "input/input_reconciler.gotmpl"
	HybridInputReconcilerCustomTemplatePath     = "input/hybrid_input_reconciler.gotmpl"
)

// Returns the template for generating input reconcilers.
func InputReconciler(params SnapshotTemplateParameters) model.CustomTemplates {
	templatePath := HomogenousInputReconcilerCustomTemplatePath
	if _, isHybrid := params.SnapshotResources.(HybridSnapshotResources); isHybrid {
		templatePath = HybridInputReconcilerCustomTemplatePath
	}
	return params.constructTemplate(params, templatePath)
}

/*
OutputSnapshot custom template
*/
const (
	OutputSnapshotCustomTemplatePath = "output/output_snapshot.gotmpl"
)

// Returns the template for generating output snapshots.
func OutputSnapshot(params SnapshotTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, OutputSnapshotCustomTemplatePath)
}
