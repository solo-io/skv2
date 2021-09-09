package controllerutils

import (
	"reflect"

	"github.com/solo-io/skv2/pkg/equalityutils"
	"istio.io/api/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type EqualityFunc = func(val1, val2 interface{}) bool

// returns true if "relevant" parts of obj1 and obj2 have equal:
// - labels,
// - annotations,
// - namespace+name,
// - non-metadata, non-status fields
// Note that Status fields are not compared.
// To compare status fields, use ObjectStatusesEqual
func ObjectsEqual(obj1, obj2 runtime.Object, equalityFunc ...EqualityFunc) bool {
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
		field1Name := value1.Type().Field(i).Name
		if field1Name == "ObjectMeta" {
			// skip ObjectMeta field, as we already asserted relevant fields are equal
			continue
		}
		if field1Name == "TypeMeta" {
			// skip TypeMeta field, as it is set by the server and not relevant for object comparison
			continue
		}
		if field1Name == "Status" {
			// skip Status field, as it is considered a separate and not relevant for object comparison
			continue
		}

		field1 := mkPointer(value1.Field(i))
		field2 := mkPointer(value2.Field(i))

		// for _, v := range

		field1ProtoVal, field1IsProto := field1.(*v1alpha3.ServiceEntry)
		if field1IsProto {
			field2ProtoVal, field2IsProto := field2.(*v1alpha3.ServiceEntry)
			if field2IsProto {
				if !serviceEntryEqual(field1ProtoVal, field2ProtoVal) {
					return false
				}
				continue
			}
		}

		// assert DeepEquality any other fields
		if !equalityutils.DeepEqual(field1, field2) {
			return false
		}
	}

	return true
}
//
// func destinationRuleEqual(val1, val2 *v1alpha3.DestinationRule) bool {
// 	if val1 == nil && val2 == nil {
// 		return true
// 	}
//
// 	if val1.GetHost() != val2.GetHost() {
// 		return false
// 	}
//
// 	if len(val1.GetExportTo()) != len(val2.GetExportTo()) {
// 		return false
// 	} else {
// 		for idx, val := range val1.GetExportTo() {
// 				if val != val2.GetExportTo()[idx] {
// 					return false
// 				}
// 		}
// 	}
//
// 	if len(val1.GetSubsets()) != len(val2.GetSubsets()) {
// 		return false
// 	} else {
//
// 	}
//
// 	if
//
// }

func serviceEntryEqual(val1, val2 *v1alpha3.ServiceEntry) bool {
	if val1 == nil && val2 == nil {
		return true
	}

	if len(val1.GetHosts()) != len(val2.GetHosts()) {
		return false
	} else {
		for idx, val := range val1.GetHosts() {
			if val != val2.GetHosts()[idx] {
				return false
			}
		}
	}

	if len(val1.GetAddresses()) != len(val2.GetAddresses()) {
		return false
	} else {
		for idx, val := range val1.GetAddresses() {
			if val != val2.GetAddresses()[idx] {
				return false
			}
		}
	}

	if len(val1.GetExportTo()) != len(val2.GetExportTo()) {
		return false
	} else {
		for idx, val := range val1.GetExportTo() {
			if val != val2.GetExportTo()[idx] {
				return false
			}
		}
	}

	if len(val1.GetSubjectAltNames()) != len(val2.GetSubjectAltNames()) {
		return false
	} else {
		for idx, val := range val1.GetSubjectAltNames() {
			if val != val2.GetSubjectAltNames()[idx] {
				return false
			}
		}
	}

	if len(val1.GetEndpoints()) != len(val2.GetEndpoints()) {
		return false
	} else {
		for idx, ep := range val1.GetEndpoints() {
			ep2 := val2.GetEndpoints()[idx]
			if !mapStringEqual(ep.GetLabels(), ep2.GetLabels()) {
				return false
			}

			if !reflect.DeepEqual(ep.GetPorts(), ep2.GetPorts()) {
				return false
			}
			if ep.GetAddress() != ep2.GetAddress() {
				return false
			}
			if ep.GetLocality() != ep2.GetLocality() {
				return false
			}
			if ep.GetServiceAccount() != ep2.GetServiceAccount() {
				return false
			}
			if ep.GetWeight() != ep2.GetWeight() {
				return false
			}
			if ep.GetNetwork() != ep2.GetNetwork() {
				return false
			}
		}
	}

	if len(val1.GetPorts()) != len(val2.GetPorts()) {
		return false
	} else {
		for idx, port := range val1.GetPorts() {
			port2 := val2.GetPorts()[idx]
			if port.GetName() != port2.GetName() {
				return false
			}
			if port.GetNumber() != port2.GetNumber() {
				return false
			}
			if port.GetTargetPort() != port2.GetTargetPort() {
				return false
			}
			if port.GetProtocol() != port2.GetProtocol() {
				return false
			}
		}
	}

	if !reflect.DeepEqual(val1.GetWorkloadSelector().GetLabels(), val2.GetWorkloadSelector().GetLabels()) {
		return false
	}

	if val1.GetResolution() != val2.GetResolution() {
		return false
	}

	if val1.GetLocation() != val2.GetLocation() {
		return false
	}
	return true
}

// if i is a pointer, just return the value.
// if i is addressable, return that.
// if i is a struct passed in by value, make a new instance of the type and copy the contents to that and return
// the pointer to that.
func mkPointer(val reflect.Value) interface{} {
	if val.Kind() == reflect.Ptr {
		return val.Interface()
	}
	if val.CanAddr() {
		return val.Addr().Interface()
	}
	if val.Kind() == reflect.Struct {
		nv := reflect.New(val.Type())
		nv.Elem().Set(val)
		return nv.Interface()
	}
	return val.Interface()
}

// returns true if "relevant" parts of obj1 and obj2 have equal:
// -labels
// -annotations
// -namespace+name
// or if the objects are not metav1.Objects
func ObjectMetasEqual(obj1, obj2 metav1.Object) bool {
	return obj1.GetNamespace() == obj2.GetNamespace() &&
		obj1.GetName() == obj2.GetName() &&
		mapStringEqual(obj1.GetLabels(), obj2.GetLabels()) &&
		mapStringEqual(obj1.GetAnnotations(), obj2.GetAnnotations())
}

func mapStringEqual(map1, map2 map[string]string) bool {
	if map1 == nil && map2 == nil {
		return false
	}

	if len(map1) != len(map2) {
		return false
	}

	for key1, val1 := range map1 {
		val2, ok := map2[key1]
		if !ok {
			return false
		}
		if val1 != val2 {
			return false
		}
	}
	return true
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

	statusField1 := value1.FieldByName("Status")
	statusField2 := value2.FieldByName("Status")

	if !statusField1.IsValid() && !statusField2.IsValid() {
		// no status
		return true
	}

	status1 := mkPointer(statusField1)
	status2 := mkPointer(statusField2)

	return equalityutils.DeepEqual(status1, status2)
}
