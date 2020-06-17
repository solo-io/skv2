package render

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// renders kubernetes from templates
type KubeCodeRenderer struct {
	templateRenderer

	// the templates to use for rendering kube kypes
	TypesTemplates inputTemplates

	// the templates to use for rendering typed kube clients which use the underlying cache
	ClientsTemplates inputTemplates

	// the templates to use for rendering kube controllers
	ControllerTemplates inputTemplates

	// the go module of the project
	GoModule string

	// the relative path to the api dir
	// types will render in the package <module>/<apiRoot>/<group>/<version>
	ApiRoot string
}

var TypesTemplates = func(skipDeepCopy bool) inputTemplates {
	tmpl := inputTemplates{
		"code/types/types.gotmpl": {
			Path: "types.go",
		},
		"code/types/register.gotmpl": {
			Path: "register.go",
		},
		"code/types/doc.gotmpl": {
			Path: "doc.go",
		},
		"code/types/zz_generated.deepcopy.gotmpl": {
			Path: "zz_generated.deepcopy.go",
		},
	}

	// remove deepcopy template if using deprecated Deepcopy generator
	if skipDeepCopy {
		delete(tmpl, "code/types/zz_generated.deepcopy.gotmpl")
	}

	return tmpl
}

var ClientsTemplates = inputTemplates{
	"code/types/clients.gotmpl": {
		Path: "clients.go",
	},
	"code/types/type_helpers.gotmpl": {
		Path: "type_helpers.go",
	},
}

var ControllerTemplates = inputTemplates{
	"code/controller/event_handlers.gotmpl": {
		Path: "controller/event_handlers.go",
	},
	"code/controller/reconcilers.gotmpl": {
		Path: "controller/reconcilers.go",
	},
	"code/controller/multicluster_reconcilers.gotmpl": {
		Path: "controller/multicluster_reconcilers.go",
	},
}

func RenderApiTypes(grp Group) ([]OutFile, error) {
	defaultKubeCodeRenderer := KubeCodeRenderer{
		templateRenderer:    DefaultTemplateRenderer,
		TypesTemplates:      TypesTemplates(grp.Generators.HasDeepcopy()),
		ClientsTemplates:    ClientsTemplates,
		ControllerTemplates: ControllerTemplates,
		GoModule:            grp.Module,
		ApiRoot:             grp.ApiRoot,
	}

	return defaultKubeCodeRenderer.RenderKubeCode(grp)
}

func (r KubeCodeRenderer) RenderKubeCode(grp Group) ([]OutFile, error) {
	templatesToRender := make(inputTemplates)
	if grp.RenderTypes {
		templatesToRender.add(r.TypesTemplates)
	}
	if grp.RenderClients {
		templatesToRender.add(r.ClientsTemplates)
	}
	if grp.RenderController {
		templatesToRender.add(r.ControllerTemplates)
	}

	files, err := r.renderCoreTemplates(templatesToRender, grp)
	if err != nil {
		return nil, err
	}

	for _, customTemplates := range grp.CustomTemplates {
		customFiles, err := r.renderCustomTemplates(customTemplates.Templates, customTemplates.Funcs, grp)
		if err != nil {
			return nil, err
		}
		files = append(files, customFiles...)
	}

	// prepend output file paths with path to api dir
	for i, out := range files {
		out.Path = filepath.Join(r.ApiRoot, grp.Group, grp.Version, out.Path)
		files[i] = out
	}

	if grp.MockgenDirective {
		// prepend each .go file with a go:generate directive to run mockgen on the file
		for i, file := range files {
			if !strings.HasSuffix(file.Path, ".go") {
				continue
			}
			// only add the directive if the file contains interfaces
			if !containsInterface(file.Content) {
				continue
			}

			baseFile := filepath.Base(file.Path)
			mockgenComment := fmt.Sprintf("//go:generate mockgen -source ./%v -destination mocks/%v", baseFile, baseFile)

			file.Content = fmt.Sprintf("%v\n\n%v", mockgenComment, file.Content)
			files[i] = file
		}
	}

	return files, nil
}

var interfaceRegex = regexp.MustCompile("type .* interface")

func containsInterface(content string) bool {
	return interfaceRegex.MatchString(content)
}
