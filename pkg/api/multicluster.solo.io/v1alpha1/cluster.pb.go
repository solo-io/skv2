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
	SecretName string `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	// key within the secret in which to find the kubeconfig data, defaults to `kubeconfig`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
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

func (m *KubernetesClusterSpec) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type KubernetesClusterStatus struct {
	// List of statuses about the kubernetes cluster.
	// This list allows for multiple applications/pods to record their connection status.
	Status []*v1.Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	// version of kubernetes
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// cloud provider, empty if unknown
	Cloud string `protobuf:"bytes,3,opt,name=cloud,proto3" json:"cloud,omitempty"`
	// Discovered information can be found here: https://kubernetes.io/docs/reference/kubernetes-api/labels-annotations-taints/
	// Geographic location in which a kubernetes cluster resides
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	// The specific zone within a region in which a the cluster resides
	Zone                 string   `protobuf:"bytes,5,opt,name=zone,proto3" json:"zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *KubernetesClusterStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *KubernetesClusterStatus) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *KubernetesClusterStatus) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *KubernetesClusterStatus) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*KubernetesClusterSpec)(nil), "multicluster.solo.io.KubernetesClusterSpec")
	proto.RegisterType((*KubernetesClusterStatus)(nil), "multicluster.solo.io.KubernetesClusterStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto", fileDescriptor_533baf1d8d7a56a7)
}

var fileDescriptor_533baf1d8d7a56a7 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0xbf, 0x2f, 0x7f, 0xba, 0xbd, 0x2d, 0x55, 0x63, 0x45, 0x2d, 0x3d, 0xf5, 0x62,
	0x96, 0xd4, 0xb3, 0x08, 0x7a, 0xf0, 0x50, 0xf0, 0x50, 0x0f, 0x82, 0x97, 0xb2, 0x5d, 0x87, 0x74,
	0x69, 0x92, 0x09, 0xfb, 0x12, 0x5a, 0x3f, 0x51, 0x3f, 0x82, 0x9f, 0xc7, 0xef, 0xe0, 0x5d, 0xb2,
	0xbb, 0x2d, 0x42, 0xf1, 0x94, 0x79, 0x7e, 0xf3, 0x30, 0x33, 0xd9, 0x87, 0xdc, 0x67, 0xd2, 0xac,
	0xec, 0x32, 0x11, 0x58, 0x30, 0x8d, 0x39, 0xde, 0x48, 0x64, 0x7a, 0x5d, 0x4f, 0x19, 0xaf, 0x24,
	0x2b, 0x6c, 0x6e, 0xa4, 0xc8, 0xad, 0x36, 0xa0, 0x58, 0x9d, 0xf2, 0xbc, 0x5a, 0xf1, 0x94, 0x05,
	0x90, 0x54, 0x0a, 0x0d, 0xd2, 0xc1, 0x6f, 0x53, 0xd2, 0x8c, 0x48, 0x24, 0x0e, 0x07, 0x19, 0x66,
	0xe8, 0x0c, 0xac, 0xa9, 0xbc, 0x77, 0x48, 0x61, 0x63, 0x3c, 0x84, 0x8d, 0x09, 0xec, 0xe2, 0xb0,
	0x4d, 0xa0, 0x02, 0x56, 0xa7, 0xee, 0xeb, 0x9b, 0xe3, 0x57, 0x72, 0x32, 0xb3, 0x4b, 0x50, 0x25,
	0x18, 0xd0, 0x8f, 0x7e, 0xc7, 0x4b, 0x05, 0x82, 0x5e, 0x93, 0xbe, 0x06, 0xa1, 0xc0, 0x2c, 0x4a,
	0x5e, 0x40, 0x1c, 0x8d, 0xa2, 0x49, 0x6f, 0x4e, 0x3c, 0x7a, 0xe6, 0x05, 0xd0, 0x4b, 0x12, 0xd4,
	0x62, 0x0d, 0xdb, 0xf8, 0x9f, 0xeb, 0xf7, 0x3c, 0x99, 0xc1, 0x76, 0xbc, 0x8b, 0xc8, 0xd9, 0xf1,
	0x64, 0xc3, 0x8d, 0xd5, 0x34, 0x25, 0x5d, 0xed, 0xaa, 0x38, 0x1a, 0xb5, 0x26, 0xfd, 0xe9, 0x79,
	0xe2, 0x2e, 0x6a, 0xee, 0xdc, 0xff, 0x5f, 0xe2, 0xad, 0xf3, 0x60, 0xa4, 0x31, 0xf9, 0x5f, 0x83,
	0xd2, 0x12, 0xcb, 0xb0, 0x6a, 0x2f, 0xe9, 0x80, 0x74, 0x44, 0x8e, 0xf6, 0x3d, 0x6e, 0x39, 0xee,
	0x05, 0x3d, 0x25, 0x5d, 0x05, 0x59, 0x63, 0x6f, 0x3b, 0x1c, 0x14, 0xa5, 0xa4, 0xfd, 0x81, 0x25,
	0xc4, 0x1d, 0x47, 0x5d, 0xfd, 0xf0, 0xf4, 0xf9, 0xdd, 0x8e, 0x76, 0x5f, 0x57, 0xd1, 0xdb, 0xdd,
	0x5f, 0x59, 0x55, 0xeb, 0xec, 0x28, 0xaf, 0xfd, 0xa9, 0x87, 0xdc, 0x96, 0x5d, 0xf7, 0xa6, 0xb7,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0xf8, 0x3c, 0xe1, 0xf3, 0x01, 0x00, 0x00,
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
	if this.SecretKey != that1.SecretKey {
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
	if this.Version != that1.Version {
		return false
	}
	if this.Cloud != that1.Cloud {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.Zone != that1.Zone {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
