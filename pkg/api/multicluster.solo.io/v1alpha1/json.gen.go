// Code generated by skv2. DO NOT EDIT.

// Generated json marshal and unmarshal functions

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	skv2jsonpb "github.com/solo-io/skv2/pkg/jsonpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var (
	marshaller   = &skv2jsonpb.Marshaler{}
	unmarshaller = &skv2jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler for KubernetesClusterSpec
func (this *KubernetesClusterSpec) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesClusterSpec
func (this *KubernetesClusterSpec) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubernetesClusterStatus
func (this *KubernetesClusterStatus) MarshalJSON() ([]byte, error) {
	str, err := marshaller.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesClusterStatus
func (this *KubernetesClusterStatus) UnmarshalJSON(b []byte) error {
	return unmarshaller.Unmarshal(bytes.NewReader(b), this)
}
