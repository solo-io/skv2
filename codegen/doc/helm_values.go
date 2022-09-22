package doc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/solo-io/go-utils/contextutils"
)

const (

	// if specified, ignore all descendants
	omitChildrenTag = "omitChildren"

	// if specified, hide the value of the field
	hideValueTag = "hideValue"

	// if specified, only render this field if the value of the env variable PRODUCT matches
	productSpecificTag = "product"
	productEnvVar      = "PRODUCT"

	// the field name
	jsonTag = "json"

	// the field description
	descTag = "desc"

	// markdown file header
	header = `
---
title: "%s"
description: Reference for Helm values.
weight: 2
---
`
)

type HelmValue struct {
	Key          string
	Type         string
	DefaultValue string
	Description  string
}

type HelmValues []HelmValue

func (v HelmValues) ToMarkdown(title string) string {
	result := new(strings.Builder)
	fmt.Fprintln(result, fmt.Sprintf(header, title))
	fmt.Fprintln(result, "|Option|Type|Default Value|Description|")
	fmt.Fprintln(result, "|------|----|-----------|-------------|")
	for _, value := range v {
		fmt.Fprintf(result, "|%s|%s|%s|%s|\n", value.Key, value.Type, value.DefaultValue, value.Description)
	}
	return result.String()
}

type addValue func(HelmValue)

func GenerateHelmValuesDoc(s interface{}, topLevelKey string, topLevelDesc string) HelmValues {
	var values []HelmValue
	cfgT := reflect.ValueOf(s)
	// If s is nil, we need to return early
	if reflect.DeepEqual(cfgT, reflect.Value{}) {
		return nil
	}
	addValue := func(v HelmValue) { values = append(values, v) }

	var path []string
	if topLevelKey != "" {
		path = []string{topLevelKey}
	}

	docReflect(addValue, path, topLevelDesc, cfgT.Type(), cfgT)

	return values
}

func docReflect(addValue addValue, path []string, desc string, typ reflect.Type, val reflect.Value) {
	switch typ.Kind() {
	case reflect.Ptr:
		var elemVal reflect.Value
		if elemVal != val {
			elemVal = val.Elem()
		}
		docReflect(addValue, path, desc, typ.Elem(), elemVal)
	case reflect.Map:

		// add entry for map itself
		addValue(HelmValue{Key: strings.Join(path, "."), Type: getMapType(typ), DefaultValue: valToString(val), Description: desc})

		if typ.Key().Kind() == reflect.String {
			docReflect(addValue, append(path, "<MAP_KEY>"), desc, typ.Elem(), reflect.Value{})

			if (val != reflect.Value{}) {
				// sort map keys for deterministic generation
				sortedKeys := val.MapKeys()
				sort.Slice(sortedKeys, func(i, j int) bool {
					return sortedKeys[i].String() < sortedKeys[j].String()
				})

				for _, k := range sortedKeys {
					elemPath := append(path, k.String())
					defaultVal := val.MapIndex(k)
					if typ.Elem().Kind() <= reflect.Float64 || typ.Elem().Kind() == reflect.String {
						// primitive type, print it as default value
						valStr := valToString(defaultVal)
						addValue(HelmValue{Key: strings.Join(elemPath, "."), Type: typ.Elem().Kind().String(), DefaultValue: valStr, Description: desc})
					} else {
						// non primitive type, descend
						docReflect(addValue, elemPath, desc, typ.Elem(), defaultVal)
					}
				}
			}
		}
	case reflect.Slice:
		lst := len(path) - 1
		path[lst] = path[lst] + "[]"

		// add entry for slice field itself
		addValue(HelmValue{Key: strings.Join(path, "."), Type: "[]" + typ.Elem().Kind().String(), DefaultValue: valToString(val), Description: desc})

		docReflect(addValue, path, desc, typ.Elem(), reflect.Value{})
	case reflect.Struct:

		// add entry for struct field itself, ignoring the top level struct
		if len(path) > 0 {
			addValue(HelmValue{Key: strings.Join(path, "."), Type: typ.Kind().String(), DefaultValue: " ", Description: desc})
		}

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)

			// ignore any private fields
			// specifically, this avoids an issue where internal proto struct fields cause infinite recursion
			if len(field.Name) > 0 && unicode.IsLower(rune(field.Name[0])) {
				continue
			}

			jsonTag := field.Tag.Get(jsonTag)
			parts := strings.Split(jsonTag, ",")
			jsonName := parts[0]
			desc := field.Tag.Get(descTag)
			fieldPath := path
			if jsonName != "" {
				fieldPath = append(fieldPath, jsonName)
			}
			var fieldVal reflect.Value
			if val != fieldVal {
				fieldVal = val.Field(i)
			}

			if _, ok := field.Tag.Lookup(hideValueTag); ok {
				fieldVal.SetString("Default value obscured")
			}

			if product, ok := field.Tag.Lookup(productSpecificTag); ok {
				if product != os.Getenv(productEnvVar) {
					continue
				}
			}

			// ignore the children of fields that are marked as such (i.e. don't recurse down)
			if _, ok := field.Tag.Lookup(omitChildrenTag); ok {
				path := strings.Join(append(path, field.Name), ".")
				kind := field.Type.Kind()
				if kind == reflect.Slice {
					path += "[]"
				} else if kind == reflect.Ptr {
					// get the underlying type of pointer types
					kind = fieldVal.Elem().Kind()
				}
				// add a HelmValue for the struct field whose children are ignored
				addValue(HelmValue{Key: path, Type: kind.String(), DefaultValue: valToString(fieldVal), Description: desc})
				continue
			}

			docReflect(addValue, fieldPath, desc, field.Type, fieldVal)
		}
	default:
		addValue(HelmValue{Key: strings.Join(path, "."), Type: typ.Kind().String(), DefaultValue: valToString(val), Description: desc})
	}
}

func valToString(val reflect.Value) string {
	var valStr string
	if val.IsValid() {
		switch val.Kind() {
		case reflect.Bool:
			valStr = fmt.Sprint(val.Bool())
		case reflect.String:
			valStr = fmt.Sprint(val.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valStr = fmt.Sprint(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			valStr = fmt.Sprint(val.Uint())
		case reflect.Ptr:
			return valToString(val.Elem())
		case reflect.Slice, reflect.Struct, reflect.Map:
			j, err := json.Marshal(val.Interface())
			if err != nil {
				contextutils.LoggerFrom(context.Background()).Warnf("failed to marshal value to json: %v", err)
			}
			valStr = string(j)
		}
	}
	// needed for correct markdown table formatting (can't have "||")
	if len(valStr) == 0 {
		valStr = " "
	}
	return valStr
}

// get string representation of map type in the form "map[keyType, valType]"
func getMapType(val reflect.Type) string {
	elemType := val.Elem().Kind().String()
	if val.Elem().Kind() == reflect.Ptr {
		elemType = val.Elem().Elem().Kind().String()
	}

	return fmt.Sprintf(
		"map[%s, %s]",
		val.Key().Kind().String(),
		elemType,
	)
}
