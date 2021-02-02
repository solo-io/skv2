package equalityutils

import (
	"reflect"

	protoV2 "google.golang.org/protobuf/proto"
)

// DeepEqual should be used in place of reflect.DeepEqual when the type of an object is unknown and may be a proto message.
// see https://github.com/golang/protobuf/issues/1173 for details on why reflect.DeepEqual no longer works for proto messages
func DeepEqual(val1, val2 interface{}) bool {
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
