package render

import (
	"github.com/solo-io/skv2/codegen/util"
)

// runs kubernetes code-generator.sh
// cannot be used to write output to memory
// also generates deecopy code
func KubeCodegen(group Group) error {
	return util.KubeCodegen(group.Group, group.Version, group.ApiRoot, group.Generators.Strings())
}
