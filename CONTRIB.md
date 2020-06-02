## Contributions

Contributions to skv2 are welcome!

Please follow the guidelines below:

For contributing **library code**:
 
1. Should be placed in a new directory under `pkg/contrib`
2. Should include test cases covering core functionality

For contributing **generated code templates**:

1. Should be placed under a new directory in `codegen/templates/contrib`
2. Add a new function variable in `codegen/templates/contrib/custom_templates.go`
with the type signature `func() model.CustomTemplates`
3. Update `codegen/cmd_test.go` with an entry for the new template under some `CustomTemplate` struct,
then execute the code generation.
4. Add test cases for the generated code under <TODO>
