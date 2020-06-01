package contrib

import (
	"github.com/gobuffalo/packr"
	"github.com/solo-io/skv2/codegen/model"
)

const (
	SetOutputFilename     = "sets/sets.go"
	SetCustomTemplatePath = "sets/sets.gotmpl"
)

var Sets = func() model.CustomTemplates {
	packrBox := packr.NewBox("./")
	templateContents, err := packrBox.FindString(SetCustomTemplatePath)
	if err != nil {
		panic(err)
	}
	return model.CustomTemplates{
		Templates: map[string]string{SetOutputFilename: templateContents},
	}
}()
