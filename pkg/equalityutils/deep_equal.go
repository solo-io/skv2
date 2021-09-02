package equalityutils

import (
	"reflect"

	"github.com/solo-io/protoc-gen-ext/pkg/equality"
	protoV2 "google.golang.org/protobuf/proto"
)

// DeepEqual should be used in place of reflect.DeepEqual when the type of an object is unknown and may be a proto message.
// see https://github.com/golang/protobuf/issues/1173 for details on why reflect.DeepEqual no longer works for proto messages
func DeepEqual(val1, val2 interface{}) bool {
	// Check if one of our types has an equal function, and use that
	if protoVal1, protoVal1IsEqualizer := val1.(equality.Equalizer); protoVal1IsEqualizer {
		return protoVal1.Equal(val2)
	} else if protoVal2, protoVal2IsEqualizer := val2.(equality.Equalizer); protoVal2IsEqualizer {
		return protoVal2.Equal(val1)
	}

	protoVal1, isProto := val1.(protoV2.Message)
	if isProto {
		protoVal2, isProto := val2.(protoV2.Message)
		if !isProto {
			return false // different types
		}
		return protoV2.Equal(protoVal1, protoVal2)
	}
	return reflect.DeepEqual(val1, val2)
}
