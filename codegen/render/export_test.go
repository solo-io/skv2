package render

import (
	"github.com/solo-io/skv2/codegen/model/values"
	goyaml "gopkg.in/yaml.v3"
)

// this file makes some private package members visible for testing

func ToNode(v interface{}, commentsConfig yamlCommentsConfig) goyaml.Node {
	return toNode(v, commentsConfig)
}

func FromNode(n goyaml.Node) string {
	return fromNode(n)
}

func ToJSONSchema(values values.UserHelmValues) string {
	return toJSONSchema(values)
}

func MergeNodes(nodes ...goyaml.Node) goyaml.Node {
	return mergeNodes(nodes...)
}
