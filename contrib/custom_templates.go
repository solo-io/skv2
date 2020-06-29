package contrib

import (
	"github.com/gobuffalo/packr"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib/codegen/funcs"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// AllCustomTemplates will be generated for the
// test API in codegen/test
var AllCustomTemplates []model.CustomTemplates

var templatesBox = packr.NewBox("./codegen/templates")

/*
Sets custom template
*/
const (
	SetOutputFilename     = "sets/sets.go"
	SetCustomTemplatePath = "sets/sets.gotmpl"
)

var Sets = func() model.CustomTemplates {
	templateContents, err := templatesBox.FindString(SetCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	setsTemplates := model.CustomTemplates{
		Templates: map[string]string{SetOutputFilename: templateContents},
	}
	// register sets
	AllCustomTemplates = append(AllCustomTemplates, setsTemplates)

	return setsTemplates
}()

/*
ClientProviders custom template
*/
const (
	ClientProvidersOutputFilename     = "providers/client_providers.go"
	ClientProvidersCustomTemplatePath = "providers/client_providers.gotmpl"
)

var ClientProviders = func() model.CustomTemplates {
	templateContents, err := templatesBox.FindString(ClientProvidersCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	clientProvidersTemplate := model.CustomTemplates{
		Templates: map[string]string{ClientProvidersOutputFilename: templateContents},
	}
	// register sets
	AllCustomTemplates = append(AllCustomTemplates, clientProvidersTemplate)

	return clientProvidersTemplate
}()

/*
InputSnapshot custom template
*/
const (
	InputSnapshotCustomTemplatePath = "input/input_snapshot.gotmpl"
)

// Returns the template for generating input snapshots. Requires some inputs, and as such,
// is not included in AllCustomTemplates.
// Input parameters:
// outputFilename    = the path of the output file produced from this template. relative to project root.
// groupModule       = the module of the project containing the codegen Group. only required if the codegen group is defined in a different go module than the types (i.e. it is using a CustomTypesImportPath, such as k8s.io/api types)
// selectFromGroups  = a set of groups containing all resources to include in the snapshot. these can be imported from other skv2 projects.
// resourcesToSelect = a map of the GVKs to the resources which we want to include in the input snapshot.
var InputSnapshot = func(outputFilename, groupModule string, selectFromGroups []model.Group, resourcesToSelect map[schema.GroupVersion][]string) model.CustomTemplates {
	templateContents, err := templatesBox.FindString(InputSnapshotCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	inputSnapshotTemplate := model.CustomTemplates{
		Templates: map[string]string{outputFilename: templateContents},
		Funcs:     funcs.MakeTopLevelFuncs(groupModule, selectFromGroups, resourcesToSelect),
	}

	return inputSnapshotTemplate
}
