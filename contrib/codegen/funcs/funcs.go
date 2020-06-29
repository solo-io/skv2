package funcs

import (
	"sort"
	"strings"
	"text/template"

	"github.com/solo-io/skv2/codegen/model"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type importedGroup struct {
	model.Group
	GoModule string // the module where the group is defined, if it differs from the group module itself. e.g. for external type imports such as k8s.io/api
}

// make funcs for "Top Level" templates, i.e.: templates which
// combine resources from multiple (including externally defined) codegen Groups.
//
// selectFromGroups = a map of Go modules to (a superset of) the imported codegen Groups. only required if the codegen group is defined in a different go module than the types (i.e. it is using a CustomTypesImportPath)
// resourcesToSelect = the GVKs of the resources which we want to select from the provided groups
func MakeTopLevelFuncs(selectFromGroups map[string][]model.Group, resourcesToSelect map[schema.GroupVersion][]string) template.FuncMap {
	importedGroups := selectResources(selectFromGroups, resourcesToSelect)
	var groups []model.Group
	groupImports := map[schema.GroupVersion]importedGroup{}

	for _, grp := range importedGroups {
		groups = append(groups, grp.Group)
		groupImports[grp.GroupVersion] = grp
	}

	return template.FuncMap{
		"imported_groups": func() []model.Group { return groups },
		"client_import_path": func(group model.Group) string {
			grp, ok := groupImports[group.GroupVersion]
			if !ok {
				panic("group not found " + grp.String())
			}
			return clientImportPath(grp)
		},
		"set_import_path": func(group model.Group) string {
			grp, ok := groupImports[group.GroupVersion]
			if !ok {
				panic("group not found " + grp.String())
			}
			return clientImportPath(grp) + "/sets"
		},
		"controller_import_path": func(group model.Group) string {
			grp, ok := groupImports[group.GroupVersion]
			if !ok {
				panic("group not found " + grp.String())
			}
			return clientImportPath(grp) + "/controller"
		},
	}
}

// gets the go package for an imported group's clients
func clientImportPath(grp importedGroup) string {

	grp.ApiRoot = strings.Trim(grp.ApiRoot, "/")

	module := grp.GoModule
	if module == "" {
		// import should be our local module, which comes from the imported group
		module = grp.Group.Module
	}

	s := strings.ReplaceAll(
		strings.Join([]string{
			module,
			grp.ApiRoot,
			grp.Group.Group,
			grp.Version,
		}, "/"),
		"//", "/",
	)

	return s
}

// pass empty string if clients live in the same go module as the type definitions
func selectResources(groups map[string][]model.Group, resourcesToSelect map[schema.GroupVersion][]string) []importedGroup {
	var selectedResources []importedGroup
	for clientModule, groups := range groups {
	for _, group := range groups {
		resources := resourcesToSelect[group.GroupVersion]
		if len(resources) == 0 {
			continue
		}
		filteredGroup := group
		filteredGroup.Resources = nil

		isResourceSelected := func(kind string) bool {
			for _, resource := range resources {
				if resource == kind {
					return true
				}
			}
			return false
		}

		for _, resource := range group.Resources {
			if !isResourceSelected(resource.Kind) {
				continue
			}
			filteredGroup.Resources = append(filteredGroup.Resources, resource)
		}

		selectedResources = append(selectedResources, importedGroup{
			Group:    filteredGroup,
			GoModule: clientModule,
		})
	}
	}

	sort.SliceStable(selectedResources, func(i, j int) bool {
		return selectedResources[i].GoModule < selectedResources[j].GoModule
	})

	return selectedResources
}
