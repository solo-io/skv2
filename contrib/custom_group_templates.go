package contrib

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/util"
	"io/ioutil"
)

/*
Define custom templates that live on the group level here.
*/

// use to get all group-level templates in contrib
var AllGroupCustomTemplates []model.CustomTemplates

var templatesDir =  util.MustGetThisDir() + "/codegen/templates/"

// NOTE(ilackarms): to add your template, copy paste XXX custom template below

/*
Sets custom template
*/
const (
	SetOutputFilename     = "sets/sets.go"
	SetCustomTemplatePath = "sets/sets.gotmpl"
)

var Sets = func() model.CustomTemplates {
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + SetCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
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
	templateContentsBytes, err := ioutil.ReadFile(templatesDir + ClientProvidersCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	templateContents := string(templateContentsBytes)
	clientProvidersTemplate := model.CustomTemplates{
		Templates: map[string]string{ClientProvidersOutputFilename: templateContents},
	}
	// register sets
	AllGroupCustomTemplates = append(AllGroupCustomTemplates, clientProvidersTemplate)

	return clientProvidersTemplate
}()
