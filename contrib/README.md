## Contributions

Contributions to skv2 are welcome! Extensions to skv2 which are common across projects but not intended to be universally applied should be placed int the `contrib` directory.

* Templates should go in `contrib/codegen/templates/` (e.g. `contrib/codegen/templates/sets.go.tmpl`) 
* Libraries imported by contrib templates should go in `contrib/pkg/` (e.g. `contrib/pkg/sets.go`)
* Test code should be added to `contrib/tests/*_test.go` where `*` is the name of the template (e.g. `contrib/tests/sets_test.go`)
* A `CustomTemplate` should be added to `contrib/custom_templates.go` like so:

```go
/*
Sets custom template
 */
const (
	SetOutputFilename     = "sets/sets.go"
	SetCustomTemplatePath = "sets/sets.gotmpl"
)

var Sets = func() model.CustomTemplates {
	templateContents, err := templatesBox.FindString(SetCustomTemplatePath)
	if err != nil {
		contextUtils.LoggerFrom(nil).DPanic(err)
		return ""
	}
	setsTemplates := model.CustomTemplates{
		Templates: map[string]string{SetOutputFilename: templateContents},
	}
	// register sets
	AllCustomTemplates = append(AllCustomTemplates, setsTemplates)

	return setsTemplates
}()
```

Note: to test generated 