package ezkube

import (
	"strings"

	"github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// parses a string returned by a gvk.String() back into a gvk
func ParseGroupVersionKindString(gvk string) (schema.GroupVersionKind, error) {
	parts := strings.Split(gvk, ", Kind=")
	if len(parts) != 2 {
		return schema.GroupVersionKind{}, eris.Errorf("invalid gkv string (missing ', Kind='): %v", gvk)
	}
	gv, kind := parts[0], parts[1]
	groupVersion, err := schema.ParseGroupVersion(gv)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	return schema.GroupVersionKind{
		Group:   groupVersion.Group,
		Version: groupVersion.Version,
		Kind:    kind,
	}, nil
}
