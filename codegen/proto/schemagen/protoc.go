package schemagen

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/log"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/solo-io/skv2/codegen/collector"
)

// Implementation of JsonSchemaGenerator that uses a plugin for the protocol buffer compiler
type protocGenerator struct {
	validationSchemaOptions *ValidationSchemaOptions
}

func NewProtocGenerator(validationSchemaOptions *ValidationSchemaOptions) *protocGenerator {
	return &protocGenerator{
		validationSchemaOptions: validationSchemaOptions,
	}
}

func (p *protocGenerator) GetJSONSchemas(protoFiles []string, imports []string, gv schema.GroupVersion) (map[schema.GroupVersionKind]*apiextv1.JSONSchemaProps, error) {
	// Use a tmp directory as the output of schemas
	// The schemas will then be matched with the appropriate CRD
	tmpOutputDir, err := os.MkdirTemp("", "skv2-schema-gen-")
	if err != nil {
		return nil, err
	}
	_ = os.MkdirAll(tmpOutputDir, os.ModePerm)
	defer os.Remove(tmpOutputDir)

	// The Executor used to compile protos
	protocExecutor := &collector.OpenApiProtocExecutor{
		OutputDir:                   tmpOutputDir,
		EnumAsIntOrString:           p.validationSchemaOptions.EnumAsIntOrString,
		MessagesWithEmptySchema:     p.validationSchemaOptions.MessagesWithEmptySchema,
		DisableKubeMarkers:          p.validationSchemaOptions.DisableKubeMarkers,
		IncludeDescriptionsInSchema: p.validationSchemaOptions.IncludeDescriptionsInSchema,
		IgnoredKubeMarkerSubstrings: p.validationSchemaOptions.IgnoredKubeMarkerSubstrings,
	}

	// 1. Generate the openApiSchemas for the project, writing them to a temp directory (schemaOutputDir)
	for _, f := range protoFiles {
		if err := p.generateSchemasForProjectProto(protocExecutor, f, imports); err != nil {
			return nil, err
		}
	}

	// 2. Walk the schemaOutputDir and convert the open api schemas into JSONSchemaProps
	return p.processGeneratedSchemas(gv, tmpOutputDir)
}

func (p *protocGenerator) generateSchemasForProjectProto(
	protocExecutor collector.ProtocExecutor,
	projectProtoFile string,
	imports []string,
) error {
	log.Printf("Generating schema for proto file: %s", projectProtoFile)

	// we don't use the output of protoc so use a temp file
	tmpFile, err := os.CreateTemp("", "sv2-schema-gen-")
	if err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	return protocExecutor.Execute(projectProtoFile, tmpFile.Name(), imports)
}

func (p *protocGenerator) processGeneratedSchemas(gv schema.GroupVersion, schemaOutputDir string) (map[schema.GroupVersionKind]*apiextv1.JSONSchemaProps, error) {
	jsonSchemasByGVK := make(map[schema.GroupVersionKind]*apiextv1.JSONSchemaProps)
	err := filepath.Walk(schemaOutputDir, func(schemaFile string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(schemaFile, ".yaml") {
			return nil
		}

		log.Debugf("Generated Schema File: %s", schemaFile)
		doc, err := p.readOpenApiDocumentFromFile(schemaFile)
		if err != nil {
			// Stop traversing the output directory
			return err
		}

		schemas := doc.Components.Schemas
		if schemas == nil {
			// Continue traversing the output directory
			return nil
		}

		for schemaKey, schemaValue := range schemas {
			schemaGVK := p.getGVKForSchemaKey(gv, schemaKey)

			// Spec validation schema
			specJsonSchema, err := p.getJsonSchema(schemaKey, schemaValue)
			if err != nil {
				return err
			}

			jsonSchemasByGVK[schemaGVK] = specJsonSchema
		}
		// Continue traversing the output directory
		return nil
	})

	return jsonSchemasByGVK, err
}

func (p *protocGenerator) readOpenApiDocumentFromFile(file string) (*openapi3.T, error) {
	var openApiDocument *openapi3.T
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "reading file")
	}
	if err := yaml.Unmarshal(bytes, &openApiDocument); err != nil {
		return nil, errors.Wrapf(err, "unmarshalling tmp file as schemas")
	}
	return openApiDocument, nil
}

func (p *protocGenerator) getGVKForSchemaKey(gv schema.GroupVersion, schemaKey string) schema.GroupVersionKind {
	// The generated keys look like testing.solo.io.MockResource
	// The kind is the `MockResource` portion
	ss := strings.Split(schemaKey, ".")
	kind := ss[len(ss)-1]

	return schema.GroupVersionKind{
		Group:   gv.Group,
		Version: gv.Version,
		Kind:    kind,
	}
}

func (p *protocGenerator) getJsonSchema(schemaKey string, schema *openapi3.SchemaRef) (*apiextv1.JSONSchemaProps, error) {
	if schema == nil {
		return nil, eris.Errorf("no open api schema for %s", schemaKey)
	}

	oApiJson, err := schema.MarshalJSON()
	if err != nil {
		return nil, eris.Errorf("Cannot marshal OpenAPI schema for %v: %v", schemaKey, err)
	}

	var obj map[string]interface{}
	if err = json.Unmarshal(oApiJson, &obj); err != nil {
		return nil, err
	}

	// detect proto.Any field from presence of "typeUrl" as field under "properties"
	removeProtoAnyValidation(obj, "typeUrl")

	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	jsonSchema := &apiextv1.JSONSchemaProps{}
	if err = json.Unmarshal(bytes, jsonSchema); err != nil {
		return nil, eris.Errorf("Cannot unmarshal raw OpenAPI schema to JSONSchemaProps for %v: %v", schemaKey, err)
	}

	return jsonSchema, nil
}
