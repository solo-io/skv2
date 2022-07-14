package util

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/solo-io/skv2/codegen/model"
)

// gets the go package for the group's types
func GoPackage(grp model.Group) string {
	if grp.CustomTypesImportPath != "" {
		return grp.CustomTypesImportPath
	}

	grp.ApiRoot = strings.Trim(grp.ApiRoot, "/")

	s := strings.ReplaceAll(
		strings.Join([]string{
			grp.Module,
			grp.ApiRoot,
			grp.Group,
			grp.Version,
		}, "/"),
		"//", "/",
	)

	return s
}

// gets the go package for the group's generated code.
// same as GoPackage if the types do not come from custom imports
func GeneratedGoPackage(grp model.Group) string {
	s := strings.ReplaceAll(
		strings.Join([]string{
			GetGoModule(),
			grp.ApiRoot,
			grp.Group,
			grp.Version,
		}, "/"),
		"//", "/",
	)

	return s
}

// Generate a package alias, eg 'import alias github.com/my/package'
// TODO:  Do something prettier if this works
func AliasFor(pkg string) string {
	return fmt.Sprintf("i%x", md5.Sum([]byte(pkg)))
}
