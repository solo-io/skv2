package render

import (
	"os"

	"github.com/solo-io/skv2/codegen/model"
)

type Group = model.Group

type Resource = model.Resource

type Field = model.Field

type HeaderOverride interface {
	Generate() string
}

type OutFile struct {
	Path       string
	Permission os.FileMode
	Content    string // set by Renderer
}
