package controllerutils

import (
	"fmt"
	"reflect"

	"github.com/solo-io/skv2/pkg/equalityutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// returns true if "relevant" parts of obj1 and obj2 have equal:
// - labels,
// - annotations,
// - namespace+name,
// - non-metadata, non-status fields
// Note that Status fields are not compared.
// To compare status fields, use ObjectStatusesEqual
func ObjectsEqual(obj1, obj2 runtime.Object) bool {
	value1, value2 := reflect.ValueOf(obj1), reflect.ValueOf(obj2)

	if value1.Type() != value2.Type() {
		return false
	}

	if value1.Kind() == reflect.Ptr {
		value1 = value1.Elem()
		value2 = value2.Elem()
	}

	if meta1, hasMeta := obj1.(metav1.Object); hasMeta {
		if !ObjectMetasEqual(meta1, obj2.(metav1.Object)) {
			return false
		}
	}

	// recurse through fields of both, comparing each:
	for i := 0; i < value1.NumField(); i++ {
		if value1.Type().Field(i).Name == "ObjectMeta" {
			// skip ObjectMeta field, as we already asserted relevant fields are equal
			continue
		}
		if value1.Type().Field(i).Name == "TypeMeta" {
			// skip TypeMeta field, as it is set by the server and not relevant for object comparison
			continue
		}
		if value1.Type().Field(i).Name == "Status" {
			// skip Status field, as it is considered a separate and not relevant for object comparison
			continue
		}

		field1 := value1.Field(i).Interface()
		field2 := value2.Field(i).Interface()

		// assert DeepEquality any other fields
		if !equalityutils.DeepEqual(field1, field2) {
			return false
		}
	}

	return true
}

// returns true if "relevant" parts of obj1 and obj2 have equal:
// -labels
// -annotations
// -namespace+name
// or if the objects are not metav1.Objects
func ObjectMetasEqual(obj1, obj2 metav1.Object) bool {
	return obj1.GetNamespace() == obj2.GetNamespace() &&
		obj1.GetName() == obj2.GetName() &&
		reflect.DeepEqual(obj1.GetLabels(), obj2.GetLabels()) &&
		reflect.DeepEqual(obj1.GetAnnotations(), obj2.GetAnnotations())
}

// returns true if the Status of obj1 and obj2 are equal.
// The objects should have a field named Status or this function will panic.
func ObjectStatusesEqual(obj1, obj2 runtime.Object) bool {
	value1, value2 := reflect.ValueOf(obj1), reflect.ValueOf(obj2)

	if value1.Type() != value2.Type() {
		return false
	}

	if value1.Kind() == reflect.Ptr {
		value1 = value1.Elem()
		value2 = value2.Elem()
	}

	status1 := value1.FieldByName("Status")

	if !status1.IsValid() {
		panic(fmt.Sprintf("cannot compare statuses of object type %T, missing status field", obj1))
	}

	status2 := value2.FieldByName("Status")

	return reflect.DeepEqual(status1.Interface(), status2.Interface())
}
