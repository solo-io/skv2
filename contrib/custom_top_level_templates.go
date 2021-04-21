package contrib

import (
	"io/ioutil"
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

	// the resources contained in the snapshot
	SnapshotResources SnapshotResources

	// name of the snapshot
	SnapshotName string
}

// SnapshotResources acts as a "oneof" to encapsulate
// HybridSnapshot or a HomogenousSnapshot
type SnapshotResources interface {
	makeTemplateFuncs(
		snapshotName, outputFilename string,
		selectFromGroups map[string][]model.Group,
	) template.FuncMap
}

// HomogenousSnapshotResources represents a set of snapshot resources read from a single source (either remote clusters or local cluster)
type HomogenousSnapshotResources struct {
	// a map of the GVKs to the resources which we want to include in the input snapshot.
	ResourcesToSelect map[schema.GroupVersion][]string
}

func (r HomogenousSnapshotResources) makeTemplateFuncs(snapshotName, outputFilename string, selectFromGroups map[string][]model.Group) template.FuncMap {
	return funcs.MakeHomogenousSnapshotFuncs(
		snapshotName,
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

func (r HybridSnapshotResources) makeTemplateFuncs(snapshotName, outputFilename string, selectFromGroups map[string][]model.Group) template.FuncMap {
	return funcs.MakeHybridSnapshotFuncs(
		snapshotName,
		outputFilename,
		selectFromGroups,
		r.LocalResourcesToSelect,
		r.RemoteResourcesToSelect,
	)
}

// NOTE(awang): to use your template in a separate repo, use this function and pass in your own mockgenDirective and templateContents
func (p SnapshotTemplateParameters) ConstructTemplate(params SnapshotTemplateParameters, templateContents string, mockgenDirective bool) model.CustomTemplates {
	crossGroupTemplate := model.CustomTemplates{
		Templates:        map[string]string{params.OutputFilename: templateContents},
		MockgenDirective: mockgenDirective,
		Funcs:            p.SnapshotResources.makeTemplateFuncs(p.SnapshotName, p.OutputFilename, p.SelectFromGroups),
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
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + InputSnapshotCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
	return params.ConstructTemplate(params, templateContents, true)
}

/*
InputSnapshot test builder custom template
*/
const (
	InputSnapshotManualBuilderCustomTemplatePath = "input/input_snapshot_manual_builder.gotmpl"
)

// Returns the template for generating input snapshots.
func InputSnapshotManualBuilder(params SnapshotTemplateParameters) model.CustomTemplates {
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + InputSnapshotManualBuilderCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
	return params.ConstructTemplate(params, templateContents, true)
}

/*
InputReconciler custom templates
*/
const (
	HomogenousInputReconcilerCustomTemplatePath       = "input/input_reconciler.gotmpl"
	HybridInputReconcilerCustomTemplatePath           = "input/hybrid_input_reconciler.gotmpl"
	HybridEventBasedInputReconcilerCustomTemplatePath = "input/hybrid_event_input_reconciler.gotmpl"
)

// Returns the template for generating input reconcilers.
func InputReconciler(params SnapshotTemplateParameters) model.CustomTemplates {
	templatePath := HomogenousInputReconcilerCustomTemplatePath
	if _, isHybrid := params.SnapshotResources.(HybridSnapshotResources); isHybrid {
		templatePath = HybridInputReconcilerCustomTemplatePath
	}
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + templatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)

	return params.ConstructTemplate(params, templateContents, true)
}

// Returns the template for generating input reconcilers.
func HybridEventBasedInputReconciler(params SnapshotTemplateParameters) model.CustomTemplates {
	templatePath := HybridEventBasedInputReconcilerCustomTemplatePath
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + templatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)

	return params.ConstructTemplate(params, templateContents, true)
}

/*
OutputSnapshot custom template
*/
const (
	OutputSnapshotCustomTemplatePath = "output/output_snapshot.gotmpl"
)

// Returns the template for generating output snapshots.
func OutputSnapshot(params SnapshotTemplateParameters) model.CustomTemplates {
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + OutputSnapshotCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
	return params.ConstructTemplate(params, templateContents, true)
}
