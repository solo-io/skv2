package contrib

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib/codegen/funcs"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// use to get all cross-group templates in contrib
func AllCrossGroupTemplates(params CrossGroupTemplateParameters) []model.CustomTemplates {
	var templates []model.CustomTemplates
	for _, tmpl := range registeredCrossGroupTemplates {
		templates = append(templates, tmpl(params))
	}
	return templates
}

var registeredCrossGroupTemplates []func(params CrossGroupTemplateParameters) model.CustomTemplates

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
		Templates: map[string]string{params.OutputFilename: templateContents},
		Funcs:     funcs.MakeTopLevelFuncs(params.SelectFromGroups, params.ResourcesToSelect),
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

// register InputSnapshot
var _ = func() bool {
	registeredCrossGroupTemplates = append(registeredCrossGroupTemplates, InputSnapshot)
	return false
}()

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

// register InputReconciler
var _ = func() bool {
	registeredCrossGroupTemplates = append(registeredCrossGroupTemplates, InputReconciler)
	return false
}()

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

// register OutputSnapshot
var _ = func() bool {
	registeredCrossGroupTemplates = append(registeredCrossGroupTemplates, OutputSnapshot)
	return false
}()
