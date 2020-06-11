package contrib

import (
	"github.com/gobuffalo/packr"
	"github.com/solo-io/skv2/codegen/model"
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
