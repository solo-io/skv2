// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Representation of a Kubernetes cluster that has been registered.
type KubernetesClusterSpec struct {
	// name of the secret which contains the kubeconfig with information to connect to the remote cluster.
	SecretName           string   `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesClusterSpec) Reset()         { *m = KubernetesClusterSpec{} }
func (m *KubernetesClusterSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec) ProtoMessage()    {}
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{0}
}
func (m *KubernetesClusterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec.Merge(m, src)
}
func (m *KubernetesClusterSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec.Size(m)
}
func (m *KubernetesClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec proto.InternalMessageInfo

func (m *KubernetesClusterSpec) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

type KubernetesClusterStatus struct {
	// List of statuses about the kubernetes cluster.
	// This list allows for multiple applications/pods to record their connection status.
	Status               []*v1.Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KubernetesClusterStatus) Reset()         { *m = KubernetesClusterStatus{} }
func (m *KubernetesClusterStatus) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterStatus) ProtoMessage()    {}
func (*KubernetesClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{1}
}
func (m *KubernetesClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterStatus.Unmarshal(m, b)
}
func (m *KubernetesClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterStatus.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterStatus.Merge(m, src)
}
func (m *KubernetesClusterStatus) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterStatus.Size(m)
}
func (m *KubernetesClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterStatus proto.InternalMessageInfo

func (m *KubernetesClusterStatus) GetStatus() []*v1.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*KubernetesClusterSpec)(nil), "multicluster.solo.io.KubernetesClusterSpec")
	proto.RegisterType((*KubernetesClusterStatus)(nil), "multicluster.solo.io.KubernetesClusterStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto", fileDescriptor_533baf1d8d7a56a7)
}

var fileDescriptor_533baf1d8d7a56a7 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x2f,
	0xce, 0x2e, 0x33, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x2d, 0xcd, 0x29, 0xc9, 0x4c, 0xce, 0x29,
	0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x87, 0x0a,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x20, 0x2b, 0xd2, 0x03, 0x19, 0xa1, 0x97, 0x99,
	0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa0, 0x0f, 0x62, 0x41, 0xd4, 0x4a, 0x09, 0xa5,
	0x56, 0x94, 0x40, 0x04, 0x53, 0x2b, 0x4a, 0xa0, 0x62, 0xd2, 0x70, 0xdb, 0x92, 0xf3, 0x8b, 0x52,
	0xf5, 0xcb, 0x0c, 0xc1, 0x34, 0x44, 0x52, 0xc9, 0x82, 0x4b, 0xd4, 0xbb, 0x34, 0x29, 0xb5, 0x28,
	0x2f, 0xb5, 0x24, 0xb5, 0xd8, 0x19, 0x62, 0x47, 0x70, 0x41, 0x6a, 0xb2, 0x90, 0x3c, 0x17, 0x77,
	0x71, 0x6a, 0x72, 0x51, 0x6a, 0x49, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x17, 0x44, 0xc8, 0x2f, 0x31, 0x37, 0x55, 0xc9, 0x87, 0x4b, 0x1c, 0x53, 0x67, 0x49,
	0x62, 0x49, 0x69, 0xb1, 0x90, 0x21, 0x17, 0x5b, 0x31, 0x98, 0x25, 0xc1, 0xa8, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xa9, 0x07, 0xb6, 0x11, 0xe4, 0x0e, 0x98, 0xfb, 0xf5, 0x20, 0x4a, 0x83, 0xa0, 0x0a,
	0x9d, 0xdc, 0x77, 0x7c, 0x65, 0x61, 0x5c, 0xf1, 0x48, 0x8e, 0x31, 0xca, 0x16, 0x57, 0x78, 0x15,
	0x64, 0xa7, 0x63, 0x84, 0x19, 0xcc, 0x38, 0x78, 0xd8, 0x25, 0xb1, 0x81, 0xfd, 0x65, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x17, 0x71, 0x0f, 0x77, 0x01, 0x00, 0x00,
}

func (this *KubernetesClusterSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SecretName != that1.SecretName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *KubernetesClusterStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterStatus)
	if !ok {
		that2, ok := that.(KubernetesClusterStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Status) != len(that1.Status) {
		return false
	}
	for i := range this.Status {
		if !this.Status[i].Equal(that1.Status[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
