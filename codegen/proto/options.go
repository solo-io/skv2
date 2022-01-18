package proto

import (
	"strings"
	"unicode"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/rotisserie/eris"
	"github.com/solo-io/cue/encoding/protobuf/cue"
	"github.com/solo-io/skv2/codegen/collector"
)

// a representation of set of parsed options for a given File's descriptor
type Options []FileOptions

/* GetUnstructuredFields gets the full list of unstructured fields contained within a given message. This function returns each field as an array of path elements, which is the fully expanded index of the field as measured from the root message. For example, given a message with the given structure:

// root level message
message MyCRDSpec {
    Options options = 1;
}

message Options {
    UnstructuredType unstructured_option = 1;
}

message RecursiveType {
    RecursiveType recursive_field = 1 [(solo.io.cue.opt).disable_openapi_validation = true];
    repeated RecursiveType repeated_recursive_field = 2 [(solo.io.cue.opt).disable_openapi_validation = true];
}

The unstructured fields of `MyCRDSpec` would be returned as:
- ["MyCRDSpec", "options", "recursiveField"]
- ["MyCRDSpec", "options", "repeatedRecursiveField"]
*/
func (o Options) GetUnstructuredFields(protoPkg string, rootMessage string) ([][]string, error) {
	rootFields, err := o.getUnstructuredFields(protoPkg, []string{rootMessage})
	if err != nil {
		return nil, eris.Wrapf(err, "getting root message")
	}
	// prepend the root message name
	for i, fieldPath := range rootFields {
		rootFields[i] = append([]string{rootMessage}, fieldPath...)
	}
	return rootFields, nil
}

func (o Options) getUnstructuredFields(protoPkg string, rootMessage []string) ([][]string, error) {
	root, err := o.getMessage(protoPkg, rootMessage)
	if err != nil {
		return nil, eris.Wrapf(err, "getting message")
	}
	var unstructuredFields [][]string
	for _, field := range root.Fields {
		rawFieldPath := []string{strcase.ToLowerCamel(field.Field.GetName())}
		if field.Field.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			// arrays become the path element '*' in the cue openapi builder
			rawFieldPath = append(rawFieldPath, "*")
		}

		// Cue does not include "value" in it's path so we have to remove it from
		// the fieldPath when it's included in map messages
		var fieldPath []string
		for _, fieldStr := range rawFieldPath {
			if fieldStr != "value" {
				fieldPath = append(fieldPath, fieldStr)
			}
		}

		if field.OpenAPIValidationDisabled || isUnstructuredMessage(protoPkg, root) {
			// we are in a leaf, return a single-level path
			unstructuredFields = append(unstructuredFields, fieldPath)
			continue
		}

		switch field.Field.GetType() {
		default:
			// the field is a primitive type
			continue
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		}

		// check if this field has any unstructured fields in its children
		fieldPkg, fieldType := splitTypeName(field.Field.GetTypeName())
		if fieldPkg == "" {
			fieldPkg = protoPkg // indicates they are the same package
		}
		// recurse into children
		childUnstructuredFields, err := o.getUnstructuredFields(fieldPkg, fieldType)
		if err != nil {
			return nil, eris.Wrapf(err, "getting child fields")
		}
		for _, childField := range childUnstructuredFields {
			// prepend the parent field name
			childPath := fieldPath
			childPath = append(childPath, childField...)
			unstructuredFields = append(unstructuredFields, childPath)
		}
	}
	return unstructuredFields, nil
}

// returns true for proto struct type, which is recursive
func isUnstructuredMessage(protoPkg string, m MessageOptions) bool {
	switch {
	case m.Message.GetName() == "Struct" && protoPkg == "google.protobuf":
		return true
	}
	return false
}

func findNestedMsg(parent MessageOptions, nestedName string) (MessageOptions, error) {
	for _, msg := range parent.NestedMessages {
		if msg.Message.GetName() == nestedName {
			return msg, nil
		}
	}
	return MessageOptions{}, eris.Errorf("nested message %v not found", nestedName)
}

func (o Options) getMessage(protoPkg string, rootMessage []string) (MessageOptions, error) {
	for _, file := range o {
		if file.File.GetPackage() != protoPkg {
			continue
		}
		for _, msg := range file.Messages {
			// traverse for nested messages
			msgName := rootMessage[0]
			if msg.Message.GetName() != msgName {
				continue
			}
			if len(rootMessage) > 1 {
				for _, nestedMsg := range rootMessage[1:] {
					var err error
					// recurse to find the nested message
					msg, err = findNestedMsg(msg, nestedMsg)
					if err != nil {
						return MessageOptions{}, err
					}
				}
			}
			return msg, nil
		}
	}
	return MessageOptions{}, eris.Errorf("message %s/%v not found in provided protos", protoPkg, rootMessage)
}

// splits a field type name into its proto package and message name constituents
func splitTypeName(fieldTypeName string) (string, []string) {
	parts := strings.Split(strings.TrimPrefix(fieldTypeName, "."), ".")
	var (
		protoPackageParts []string
		messageNameParts  []string
	)
	for i, part := range parts {
		if part == "" {
			continue
		}
		if unicode.IsUpper(rune(part[0])) {
			// first upper case indicates the beginning of the message name
			protoPackageParts = parts[:i]
			messageNameParts = parts[i:]
			break
		}
	}
	return strings.Join(protoPackageParts, "."), messageNameParts
}

// a representation of set of parsed options for a given File's descriptor
type FileOptions struct {
	File     *collector.DescriptorWithPath
	Messages []MessageOptions
}

// options for a message
type MessageOptions struct {
	Message        *descriptor.DescriptorProto
	Fields         []FieldOptions
	NestedMessages []MessageOptions
}

// options for a field
type FieldOptions struct {
	Field *descriptor.FieldDescriptorProto

	OpenAPIValidationDisabled bool
}

func parseMessageOptions(msg *descriptor.DescriptorProto) (MessageOptions, error) {
	var fields []FieldOptions
	for _, fieldDescriptorProto := range msg.GetField() {
		fieldOpts, err := getFieldOptions(fieldDescriptorProto)
		if err != nil {
			return MessageOptions{}, err
		}
		fields = append(fields, fieldOpts)
	}
	var nestedMessages []MessageOptions
	for _, nestedMsg := range msg.GetNestedType() {
		fieldOpts, err := parseMessageOptions(nestedMsg)
		if err != nil {
			return MessageOptions{}, err
		}
		nestedMessages = append(nestedMessages, fieldOpts)
	}
	return MessageOptions{
		Message:        msg,
		NestedMessages: nestedMessages,
		Fields:         fields,
	}, nil
}

func ParseOptions(fileDescriptors []*collector.DescriptorWithPath) (Options, error) {
	if len(fileDescriptors) == 0 {
		return nil, nil
	}
	var files Options
	// map of filename to the set of paths ([]string) to treat as unstructured. read from custom proto option
	for _, fileDescriptor := range fileDescriptors {
		var messages []MessageOptions
		for _, message := range fileDescriptor.GetMessageType() {
			msgOpts, err := parseMessageOptions(message)
			if err != nil {
				return nil, err
			}
			messages = append(messages, msgOpts)
		}
		files = append(files, FileOptions{
			File:     fileDescriptor,
			Messages: messages,
		})
	}
	return files, nil
}

func getFieldOptions(field *descriptor.FieldDescriptorProto) (FieldOptions, error) {
	validationDisabled, err := getFieldOptionOpenAPIValidationDisabled(field)
	if err != nil {
		return FieldOptions{}, err
	}
	return FieldOptions{
		Field:                     field,
		OpenAPIValidationDisabled: validationDisabled,
	}, nil
}
func getFieldOptionOpenAPIValidationDisabled(field *descriptor.FieldDescriptorProto) (bool, error) {
	cueOptRaw, err := proto.GetExtension(field.Options, cue.E_Opt)
	if err == nil {
		cueOpt, ok := cueOptRaw.(*cue.FieldOptions)
		if !ok {
			return false, eris.Errorf("internal error: invalid option type %T expecting *cueproto.FieldOptions", cueOpt)
		}
		return cueOpt.DisableOpenapiValidation, nil
	}
	return false, nil
}
