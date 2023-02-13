package render

import goyaml "gopkg.in/yaml.v3"

// this file makes some private package members visible for testing

func ToNode(v interface{}, commentsConfig yamlCommentsConfig) goyaml.Node {
	return toNode(v, commentsConfig)
}

func FromNode(n goyaml.Node) string {
	return fromNode(n)
}
