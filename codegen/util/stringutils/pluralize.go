package stringutils

import "github.com/gertd/go-pluralize"

// Define cases for pluralizing which pluralize library does not handle
var SpecialCases = map[string]string{
	"schema": "schemas",
}

// Pluralize is the canonical pluralization function for SKv2. It should be used to special case
// when we want a different result than the underlying pluralize library
func Pluralize(s string) string {
	c := pluralize.NewClient()
	for singular, plural := range SpecialCases {
		c.AddIrregularRule(singular, plural)
	}
	return c.Plural(s)
}
