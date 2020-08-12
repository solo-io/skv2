package contrib

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib/codegen/funcs"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

/*
Define custom templates that live on the group level here.
*/

// Parameters for constructing templates that span across multiple Groups.
type CrossGroupTemplateParameters struct {
	// the path of the output file produced from this template. relative to project root.
	OutputFilename string

	// a map of Go modules to (a superset of) the imported codegen Groups. only required if the codegen group is defined in a different go module than the types (i.e. it is using a CustomTypesImportPath)
	SelectFromGroups map[string][]model.Group

	// a map of the GVKs to the resources which we want to include in the input snapshot.
	ResourcesToSelect map[schema.GroupVersion][]string
}

func (p CrossGroupTemplateParameters) constructTemplate(params CrossGroupTemplateParameters, templatePath string) model.CustomTemplates {
	templateContents, err := templatesBox.FindString(templatePath)
	if err != nil {
		panic(err)
	}
	inputSnapshotTemplate := model.CustomTemplates{
		Templates:        map[string]string{params.OutputFilename: templateContents},
		MockgenDirective: true,
		Funcs:            funcs.MakeTopLevelFuncs(params.OutputFilename, params.SelectFromGroups, params.ResourcesToSelect),
	}

	return inputSnapshotTemplate
}

// NOTE(ilackarms): to add your template, copy paste XXX custom template below

/*
InputSnapshot custom template
*/
const (
	InputSnapshotCustomTemplatePath = "input/input_snapshot.gotmpl"
)

// Returns the template for generating input snapshots.
func InputSnapshot(params CrossGroupTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, InputSnapshotCustomTemplatePath)
}

/*
InputSnapshot test builder custom template
*/
const (
	InputSnapshotManualBuilderCustomTemplatePath = "input/input_snapshot_manual_builder.gotmpl"
)

// Returns the template for generating input snapshots.
func InputSnapshotManualBuilder(params CrossGroupTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, InputSnapshotManualBuilderCustomTemplatePath)
}

/*
InputReconciler custom template
*/
const (
	InputReconcilerCustomTemplatePath = "input/input_reconciler.gotmpl"
)

// Returns the template for generating input reconcilers.
func InputReconciler(params CrossGroupTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, InputReconcilerCustomTemplatePath)
}

/*
OutputSnapshot custom template
*/
const (
	OutputSnapshotCustomTemplatePath = "output/output_snapshot.gotmpl"
)

// Returns the template for generating output snapshots.
func OutputSnapshot(params CrossGroupTemplateParameters) model.CustomTemplates {
	return params.constructTemplate(params, OutputSnapshotCustomTemplatePath)
}
