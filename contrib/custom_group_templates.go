package contrib

import (
	"github.com/gobuffalo/packr"
	"github.com/solo-io/skv2/codegen/model"
)

/*
Define custom templates that live on the group level here.
*/

// use to get all group-level templates in contrib
var AllGroupCustomTemplates []model.CustomTemplates

var templatesBox = packr.NewBox("./codegen/templates")

// NOTE(ilackarms): to add your template, copy paste XXX custom template below

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
	AllGroupCustomTemplates = append(AllGroupCustomTemplates, setsTemplates)

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
	AllGroupCustomTemplates = append(AllGroupCustomTemplates, clientProvidersTemplate)

	return clientProvidersTemplate
}()
