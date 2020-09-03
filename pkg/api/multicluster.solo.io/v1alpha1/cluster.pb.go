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
	// name local DNS suffix used by the cluster.
	// used for building FQDNs for in-cluster services
	// defaults to 'cluster.local'
	ClusterDomain string `protobuf:"bytes,2,opt,name=cluster_domain,json=clusterDomain,proto3" json:"cluster_domain,omitempty"`
	// Metadata for clusters provisioned from cloud providers.
	ProviderInfo         *KubernetesClusterSpec_ProviderInfo `protobuf:"bytes,3,opt,name=provider_info,json=providerInfo,proto3" json:"provider_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
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

func (m *KubernetesClusterSpec) GetClusterDomain() string {
	if m != nil {
		return m.ClusterDomain
	}
	return ""
}

func (m *KubernetesClusterSpec) GetProviderInfo() *KubernetesClusterSpec_ProviderInfo {
	if m != nil {
		return m.ProviderInfo
	}
	return nil
}

// Metadata for clusters provisioned from cloud providers.
type KubernetesClusterSpec_ProviderInfo struct {
	// Types that are valid to be assigned to ProviderInfoType:
	//	*KubernetesClusterSpec_ProviderInfo_Eks
	ProviderInfoType     isKubernetesClusterSpec_ProviderInfo_ProviderInfoType `protobuf_oneof:"ProviderInfoType"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *KubernetesClusterSpec_ProviderInfo) Reset()         { *m = KubernetesClusterSpec_ProviderInfo{} }
func (m *KubernetesClusterSpec_ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec_ProviderInfo) ProtoMessage()    {}
func (*KubernetesClusterSpec_ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{0, 0}
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Merge(m, src)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Size(m)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec_ProviderInfo proto.InternalMessageInfo

type isKubernetesClusterSpec_ProviderInfo_ProviderInfoType interface {
	isKubernetesClusterSpec_ProviderInfo_ProviderInfoType()
	Equal(interface{}) bool
}

type KubernetesClusterSpec_ProviderInfo_Eks struct {
	Eks *KubernetesClusterSpec_Eks `protobuf:"bytes,1,opt,name=eks,proto3,oneof" json:"eks,omitempty"`
}

func (*KubernetesClusterSpec_ProviderInfo_Eks) isKubernetesClusterSpec_ProviderInfo_ProviderInfoType() {
}

func (m *KubernetesClusterSpec_ProviderInfo) GetProviderInfoType() isKubernetesClusterSpec_ProviderInfo_ProviderInfoType {
	if m != nil {
		return m.ProviderInfoType
	}
	return nil
}

func (m *KubernetesClusterSpec_ProviderInfo) GetEks() *KubernetesClusterSpec_Eks {
	if x, ok := m.GetProviderInfoType().(*KubernetesClusterSpec_ProviderInfo_Eks); ok {
		return x.Eks
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KubernetesClusterSpec_ProviderInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KubernetesClusterSpec_ProviderInfo_Eks)(nil),
	}
}

// AWS metadata associated with an EKS provisioned cluster.
type KubernetesClusterSpec_Eks struct {
	// AWS ARN.
	Arn string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	// AWS 12 digit account ID.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// AWS region.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// EKS resource name.
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesClusterSpec_Eks) Reset()         { *m = KubernetesClusterSpec_Eks{} }
func (m *KubernetesClusterSpec_Eks) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec_Eks) ProtoMessage()    {}
func (*KubernetesClusterSpec_Eks) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{0, 1}
}
func (m *KubernetesClusterSpec_Eks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec_Eks.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec_Eks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec_Eks.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec_Eks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec_Eks.Merge(m, src)
}
func (m *KubernetesClusterSpec_Eks) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec_Eks.Size(m)
}
func (m *KubernetesClusterSpec_Eks) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec_Eks.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec_Eks proto.InternalMessageInfo

func (m *KubernetesClusterSpec_Eks) GetArn() string {
	if m != nil {
		return m.Arn
	}
	return ""
}

func (m *KubernetesClusterSpec_Eks) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *KubernetesClusterSpec_Eks) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *KubernetesClusterSpec_Eks) GetName() string {
	if m != nil {
		return m.Name
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
	proto.RegisterType((*KubernetesClusterSpec_ProviderInfo)(nil), "multicluster.solo.io.KubernetesClusterSpec.ProviderInfo")
	proto.RegisterType((*KubernetesClusterSpec_Eks)(nil), "multicluster.solo.io.KubernetesClusterSpec.Eks")
	proto.RegisterType((*KubernetesClusterStatus)(nil), "multicluster.solo.io.KubernetesClusterStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto", fileDescriptor_533baf1d8d7a56a7)
}

var fileDescriptor_533baf1d8d7a56a7 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0x1d, 0xa7, 0x14, 0x7a, 0xba, 0x2b, 0x4b, 0x58, 0xb5, 0x56, 0xd4, 0xb2, 0x20, 0xf4,
	0xc6, 0x84, 0xd6, 0x1b, 0x6f, 0x44, 0xd8, 0x75, 0xd1, 0x45, 0x11, 0x19, 0xbd, 0x12, 0xa4, 0xa4,
	0xd3, 0xb3, 0xb3, 0x61, 0x3a, 0x39, 0x21, 0xc9, 0x0c, 0xeb, 0x33, 0xf8, 0x22, 0x3e, 0x82, 0xcf,
	0xe3, 0x3b, 0x78, 0x2f, 0x93, 0x89, 0x65, 0x60, 0xeb, 0x85, 0x57, 0x73, 0xf2, 0xcb, 0x37, 0xe7,
	0x5f, 0x3e, 0x78, 0x55, 0x28, 0x7f, 0x55, 0xaf, 0x79, 0x4e, 0x95, 0x70, 0xb4, 0xa5, 0x67, 0x8a,
	0x84, 0x2b, 0x9b, 0xa5, 0x90, 0x46, 0x89, 0xaa, 0xde, 0x7a, 0x95, 0x6f, 0x6b, 0xe7, 0xd1, 0x8a,
	0x66, 0x21, 0xb7, 0xe6, 0x4a, 0x2e, 0x44, 0x04, 0xdc, 0x58, 0xf2, 0xc4, 0x8e, 0xfb, 0x22, 0xde,
	0xa6, 0xe0, 0x8a, 0xa6, 0xc7, 0x05, 0x15, 0x14, 0x04, 0xa2, 0x8d, 0x3a, 0xed, 0x94, 0xe1, 0xb5,
	0xef, 0x20, 0x5e, 0xfb, 0xc8, 0x1e, 0xee, 0xaa, 0xe5, 0x64, 0x51, 0x34, 0x8b, 0xf0, 0xed, 0x2e,
	0x4f, 0xbe, 0xa7, 0x70, 0xf7, 0x5d, 0xbd, 0x46, 0xab, 0xd1, 0xa3, 0x3b, 0xeb, 0x8a, 0x7c, 0x32,
	0x98, 0xb3, 0x27, 0x30, 0x76, 0x98, 0x5b, 0xf4, 0x2b, 0x2d, 0x2b, 0x9c, 0x24, 0xb3, 0x64, 0x3e,
	0xca, 0xa0, 0x43, 0x1f, 0x64, 0x85, 0xec, 0x29, 0xdc, 0x89, 0x4d, 0xad, 0x36, 0x54, 0x49, 0xa5,
	0x27, 0xb7, 0x83, 0xe6, 0x30, 0xd2, 0xd7, 0x01, 0xb2, 0xaf, 0x70, 0x68, 0x2c, 0x35, 0x6a, 0x83,
	0x76, 0xa5, 0xf4, 0x25, 0x4d, 0xd2, 0x59, 0x32, 0x1f, 0x2f, 0x5f, 0xf0, 0x7d, 0x63, 0xf1, 0xbd,
	0xbd, 0xf0, 0x8f, 0x31, 0xc1, 0x85, 0xbe, 0xa4, 0xec, 0xc0, 0xf4, 0x4e, 0xd3, 0x02, 0x0e, 0xfa,
	0xb7, 0xec, 0x0c, 0x52, 0x2c, 0x5d, 0x68, 0x77, 0xbc, 0x14, 0xff, 0x53, 0xe4, 0xbc, 0x74, 0x6f,
	0x6f, 0x65, 0xed, 0xdf, 0xa7, 0x0c, 0x8e, 0xfa, 0x49, 0x3f, 0x7f, 0x33, 0x38, 0x5d, 0x43, 0x7a,
	0x5e, 0x3a, 0x76, 0x04, 0xa9, 0xb4, 0x3a, 0xae, 0xa3, 0x0d, 0xd9, 0x23, 0x00, 0x99, 0xe7, 0x54,
	0x6b, 0xbf, 0x52, 0x9b, 0xb8, 0x83, 0x51, 0x24, 0x17, 0x1b, 0x76, 0x0f, 0x86, 0x16, 0x0b, 0x45,
	0x3a, 0x0c, 0x3e, 0xca, 0xe2, 0x89, 0x31, 0x18, 0x84, 0xc5, 0x0e, 0x02, 0x0d, 0xf1, 0xc9, 0x7b,
	0xb8, 0x7f, 0xb3, 0x37, 0x2f, 0x7d, 0xed, 0xd8, 0x02, 0x86, 0x2e, 0x44, 0x93, 0x64, 0x96, 0xce,
	0xc7, 0xcb, 0x07, 0x3c, 0xbc, 0x62, 0xfb, 0xb6, 0xbb, 0xb9, 0x3a, 0x69, 0x16, 0x85, 0xa7, 0x6f,
	0x7e, 0xfe, 0x1e, 0x24, 0x3f, 0x7e, 0x3d, 0x4e, 0xbe, 0xbc, 0xfc, 0x97, 0x07, 0x4d, 0x59, 0xdc,
	0xf0, 0xe1, 0xdf, 0x74, 0x3b, 0x3f, 0xae, 0x87, 0xc1, 0x2b, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x90, 0xe8, 0xa8, 0xe5, 0xcb, 0x02, 0x00, 0x00,
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
	if this.ClusterDomain != that1.ClusterDomain {
		return false
	}
	if !this.ProviderInfo.Equal(that1.ProviderInfo) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *KubernetesClusterSpec_ProviderInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec_ProviderInfo)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_ProviderInfo)
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
	if that1.ProviderInfoType == nil {
		if this.ProviderInfoType != nil {
			return false
		}
	} else if this.ProviderInfoType == nil {
		return false
	} else if !this.ProviderInfoType.Equal(that1.ProviderInfoType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *KubernetesClusterSpec_ProviderInfo_Eks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec_ProviderInfo_Eks)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_ProviderInfo_Eks)
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
	if !this.Eks.Equal(that1.Eks) {
		return false
	}
	return true
}
func (this *KubernetesClusterSpec_Eks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec_Eks)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_Eks)
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
	if this.Arn != that1.Arn {
		return false
	}
	if this.AccountId != that1.AccountId {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.Name != that1.Name {
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
