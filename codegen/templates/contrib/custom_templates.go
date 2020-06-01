package contrib

import (
	"github.com/gobuffalo/packr"
	"github.com/solo-io/skv2/codegen/model"
)

const (
	SetOutputFilename     = "sets/sets.go"
	SetCustomTemplatePath = "sets/sets.gotmpl"
)

func Sets(customTemplate *model.CustomTemplates) error {
	packrBox := packr.NewBox("./")
	templateContents, err := packrBox.FindString(SetCustomTemplatePath)
	if err != nil {
		return err
	}
	return customTemplate.Merge(
		map[string]string{
			SetOutputFilename: templateContents,
		})
}
