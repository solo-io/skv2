// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The State of a reconciled object
type Status_State int32

const (
	// Waiting to be processed.
	Status_PENDING Status_State = 0
	// Currently processing.
	Status_PROCESSING Status_State = 1
	// Invalid parameters supplied, will not continue.
	Status_INVALID Status_State = 2
	// Failed during processing.
	Status_FAILED Status_State = 3
	// Finished processing successfully.
	Status_ACCEPTED Status_State = 4
)

var Status_State_name = map[int32]string{
	0: "PENDING",
	1: "PROCESSING",
	2: "INVALID",
	3: "FAILED",
	4: "ACCEPTED",
}

var Status_State_value = map[string]int32{
	"PENDING":    0,
	"PROCESSING": 1,
	"INVALID":    2,
	"FAILED":     3,
	"ACCEPTED":   4,
}

func (x Status_State) String() string {
	return proto.EnumName(Status_State_name, int32(x))
}

func (Status_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b01d80c02a5b697, []int{4, 0}
}

// Resource reference for an object
type ObjectRef struct {
	// name of the resource being referenced
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectRef) Reset()         { *m = ObjectRef{} }
func (m *ObjectRef) String() string { return proto.CompactTextString(m) }
func (*ObjectRef) ProtoMessage()    {}
func (*ObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b01d80c02a5b697, []int{0}
}

func (m *ObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectRef.Unmarshal(m, b)
}
func (m *ObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectRef.Marshal(b, m, deterministic)
}
func (m *ObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectRef.Merge(m, src)
}
func (m *ObjectRef) XXX_Size() int {
	return xxx_messageInfo_ObjectRef.Size(m)
}
func (m *ObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectRef proto.InternalMessageInfo

func (m *ObjectRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Resource reference for a cross-cluster-scoped object
type ClusterObjectRef struct {
	// name of the resource being referenced
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// name of the cluster in which the resource exists
	ClusterName          string   `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterObjectRef) Reset()         { *m = ClusterObjectRef{} }
func (m *ClusterObjectRef) String() string { return proto.CompactTextString(m) }
func (*ClusterObjectRef) ProtoMessage()    {}
func (*ClusterObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b01d80c02a5b697, []int{1}
}

func (m *ClusterObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterObjectRef.Unmarshal(m, b)
}
func (m *ClusterObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterObjectRef.Marshal(b, m, deterministic)
}
func (m *ClusterObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterObjectRef.Merge(m, src)
}
func (m *ClusterObjectRef) XXX_Size() int {
	return xxx_messageInfo_ClusterObjectRef.Size(m)
}
func (m *ClusterObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterObjectRef proto.InternalMessageInfo

func (m *ClusterObjectRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterObjectRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ClusterObjectRef) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

// Resource reference for a typed object
type TypedObjectRef struct {
	// API group of the resource being referenced
	ApiGroup *wrappers.StringValue `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Kind of the resource being referenced
	Kind *wrappers.StringValue `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// name of the resource being referenced
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace            string   `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypedObjectRef) Reset()         { *m = TypedObjectRef{} }
func (m *TypedObjectRef) String() string { return proto.CompactTextString(m) }
func (*TypedObjectRef) ProtoMessage()    {}
func (*TypedObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b01d80c02a5b697, []int{2}
}

func (m *TypedObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedObjectRef.Unmarshal(m, b)
}
func (m *TypedObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedObjectRef.Marshal(b, m, deterministic)
}
func (m *TypedObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedObjectRef.Merge(m, src)
}
func (m *TypedObjectRef) XXX_Size() int {
	return xxx_messageInfo_TypedObjectRef.Size(m)
}
func (m *TypedObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_TypedObjectRef proto.InternalMessageInfo

func (m *TypedObjectRef) GetApiGroup() *wrappers.StringValue {
	if m != nil {
		return m.ApiGroup
	}
	return nil
}

func (m *TypedObjectRef) GetKind() *wrappers.StringValue {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *TypedObjectRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypedObjectRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Resource reference for a typed, cross-cluster-scoped object
type TypedClusterObjectRef struct {
	// API group of the resource being referenced
	ApiGroup *wrappers.StringValue `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Kind of the resource being referenced
	Kind *wrappers.StringValue `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// name of the resource being referenced
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// name of the cluster in which the resource exists
	ClusterName          string   `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypedClusterObjectRef) Reset()         { *m = TypedClusterObjectRef{} }
func (m *TypedClusterObjectRef) String() string { return proto.CompactTextString(m) }
func (*TypedClusterObjectRef) ProtoMessage()    {}
func (*TypedClusterObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b01d80c02a5b697, []int{3}
}

func (m *TypedClusterObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedClusterObjectRef.Unmarshal(m, b)
}
func (m *TypedClusterObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedClusterObjectRef.Marshal(b, m, deterministic)
}
func (m *TypedClusterObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedClusterObjectRef.Merge(m, src)
}
func (m *TypedClusterObjectRef) XXX_Size() int {
	return xxx_messageInfo_TypedClusterObjectRef.Size(m)
}
func (m *TypedClusterObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedClusterObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_TypedClusterObjectRef proto.InternalMessageInfo

func (m *TypedClusterObjectRef) GetApiGroup() *wrappers.StringValue {
	if m != nil {
		return m.ApiGroup
	}
	return nil
}

func (m *TypedClusterObjectRef) GetKind() *wrappers.StringValue {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *TypedClusterObjectRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypedClusterObjectRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TypedClusterObjectRef) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

// A generic status
type Status struct {
	// The current state of the resource
	State Status_State `protobuf:"varint,1,opt,name=state,proto3,enum=core.skv2.solo.io.Status_State" json:"state,omitempty"`
	// A human readable message about the current state of the object
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The most recently observed generation of the resource. This value corresponds to the `metadata.generation` of
	// a kubernetes resource
	ObservedGeneration int64 `protobuf:"varint,3,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The time at which this status was recorded
	ProcessingTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
	// (optional) The owner of the status, this value can be used to identify the entity which wrote this status.
	// This is useful in situations where a given resource may have multiple owners.
	Owner                *wrappers.StringValue `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b01d80c02a5b697, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetState() Status_State {
	if m != nil {
		return m.State
	}
	return Status_PENDING
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *Status) GetProcessingTime() *timestamp.Timestamp {
	if m != nil {
		return m.ProcessingTime
	}
	return nil
}

func (m *Status) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterEnum("core.skv2.solo.io.Status_State", Status_State_name, Status_State_value)
	proto.RegisterType((*ObjectRef)(nil), "core.skv2.solo.io.ObjectRef")
	proto.RegisterType((*ClusterObjectRef)(nil), "core.skv2.solo.io.ClusterObjectRef")
	proto.RegisterType((*TypedObjectRef)(nil), "core.skv2.solo.io.TypedObjectRef")
	proto.RegisterType((*TypedClusterObjectRef)(nil), "core.skv2.solo.io.TypedClusterObjectRef")
	proto.RegisterType((*Status)(nil), "core.skv2.solo.io.Status")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/api/core/v1/core.proto", fileDescriptor_0b01d80c02a5b697)
}

var fileDescriptor_0b01d80c02a5b697 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0x76, 0x61, 0xa1, 0xe5, 0xd1, 0x20, 0x8e, 0x31, 0x21, 0xa4, 0xb1, 0xca, 0xc9, 0x8b, 0xbb,
	0x16, 0xf5, 0xe0, 0x41, 0x13, 0x04, 0x24, 0xc4, 0x86, 0x92, 0x85, 0xf4, 0xe0, 0x85, 0x0c, 0xcb,
	0xeb, 0x38, 0xc2, 0xee, 0x4c, 0x66, 0x66, 0x69, 0xfd, 0x55, 0x5e, 0x3d, 0xfa, 0x13, 0xbc, 0xfa,
	0x43, 0x7a, 0x37, 0x3b, 0x5b, 0x20, 0x76, 0xd3, 0xf4, 0xe0, 0xa5, 0xa7, 0xf7, 0xf6, 0xbd, 0xef,
	0xcd, 0xbc, 0xef, 0x7b, 0x6f, 0x07, 0x3c, 0xc6, 0xcd, 0xd7, 0x64, 0xee, 0x85, 0x22, 0xf2, 0xb5,
	0x58, 0x89, 0x97, 0x5c, 0xf8, 0x7a, 0xb9, 0x6e, 0xfb, 0x54, 0x72, 0x3f, 0x14, 0x0a, 0xfd, 0xf5,
	0xb1, 0xb5, 0x9e, 0x54, 0xc2, 0x08, 0xf2, 0xc8, 0xfa, 0x29, 0xc2, 0x4b, 0xe1, 0x1e, 0x17, 0xcd,
	0xa7, 0x4c, 0x08, 0xb6, 0x42, 0xdf, 0x02, 0xe6, 0xc9, 0xb9, 0x7f, 0xa1, 0xa8, 0x94, 0xa8, 0x74,
	0x56, 0xd2, 0x3c, 0xba, 0x99, 0x37, 0x3c, 0x42, 0x6d, 0x68, 0x24, 0xaf, 0x01, 0x04, 0x2f, 0x8d,
	0xf5, 0x7c, 0xbc, 0x34, 0x59, 0xac, 0xf5, 0x1e, 0x2a, 0xa7, 0xf3, 0x6f, 0x18, 0x9a, 0x00, 0xcf,
	0x09, 0x01, 0x37, 0xa6, 0x11, 0x36, 0x9c, 0x67, 0xce, 0x8b, 0x4a, 0x60, 0x7d, 0x72, 0x08, 0x95,
	0xd4, 0x6a, 0x49, 0x43, 0x6c, 0x14, 0x6c, 0x62, 0x17, 0x68, 0x31, 0xa8, 0x77, 0x57, 0x89, 0x36,
	0xa8, 0xfe, 0xe3, 0x14, 0xf2, 0x1c, 0x0e, 0xc2, 0xec, 0x94, 0x99, 0xad, 0x2c, 0x5a, 0x40, 0xf5,
	0x3a, 0x36, 0xa2, 0x11, 0xb6, 0x7e, 0x38, 0x50, 0x9b, 0x7e, 0x97, 0xb8, 0xd8, 0xdd, 0xf3, 0x0e,
	0x2a, 0x54, 0xf2, 0x19, 0x53, 0x22, 0x91, 0xf6, 0xb2, 0x6a, 0xfb, 0xd0, 0xcb, 0x34, 0xf0, 0x36,
	0x1a, 0x78, 0x13, 0xa3, 0x78, 0xcc, 0xce, 0xe8, 0x2a, 0xc1, 0x60, 0x9f, 0x4a, 0x3e, 0x48, 0xd1,
	0xe4, 0x15, 0xb8, 0x4b, 0x1e, 0x2f, 0x6c, 0x27, 0x77, 0x55, 0x59, 0xe4, 0x96, 0x54, 0xf1, 0x36,
	0x52, 0xee, 0x4d, 0x69, 0xfe, 0x38, 0xf0, 0xc4, 0x76, 0x9c, 0x13, 0xe8, 0x7e, 0x37, 0x9e, 0x9b,
	0x46, 0x29, 0x3f, 0x8d, 0xdf, 0x05, 0x28, 0x4f, 0x0c, 0x35, 0x89, 0x26, 0x6f, 0xa1, 0xa4, 0x0d,
	0x35, 0xd9, 0xb8, 0x6b, 0xed, 0x23, 0x2f, 0xb7, 0xb8, 0x5e, 0x86, 0xb4, 0x06, 0x83, 0x0c, 0x4d,
	0x1a, 0xb0, 0x17, 0xa1, 0xd6, 0x94, 0x6d, 0xd6, 0x61, 0xf3, 0x49, 0x7c, 0x78, 0x2c, 0xe6, 0x1a,
	0xd5, 0x1a, 0x17, 0x33, 0x86, 0x31, 0x2a, 0x6a, 0xb8, 0x88, 0x6d, 0xff, 0xc5, 0x80, 0x6c, 0x52,
	0x83, 0x6d, 0x86, 0x74, 0xe1, 0xa1, 0x54, 0x22, 0x44, 0xad, 0x79, 0xcc, 0x66, 0xe9, 0xd2, 0x5b,
	0x4e, 0xd5, 0x76, 0x33, 0x27, 0xcf, 0x74, 0xf3, 0x47, 0x04, 0xb5, 0x5d, 0x49, 0x1a, 0x24, 0x6d,
	0x28, 0x89, 0x8b, 0x18, 0x95, 0x65, 0x7b, 0x97, 0xb2, 0x19, 0xb4, 0xf5, 0x19, 0x4a, 0x96, 0x13,
	0xa9, 0xc2, 0xde, 0xb8, 0x3f, 0xea, 0x0d, 0x47, 0x83, 0xfa, 0x03, 0x52, 0x03, 0x18, 0x07, 0xa7,
	0xdd, 0xfe, 0x64, 0x92, 0x7e, 0x3b, 0x69, 0x72, 0x38, 0x3a, 0xeb, 0x9c, 0x0c, 0x7b, 0xf5, 0x02,
	0x01, 0x28, 0x7f, 0xea, 0x0c, 0x4f, 0xfa, 0xbd, 0x7a, 0x91, 0x1c, 0xc0, 0x7e, 0xa7, 0xdb, 0xed,
	0x8f, 0xa7, 0xfd, 0x5e, 0xdd, 0xfd, 0xf8, 0xe1, 0xe7, 0x95, 0xeb, 0xfc, 0xba, 0x72, 0x9d, 0x2f,
	0x6f, 0x6e, 0x7b, 0x2a, 0xe4, 0x92, 0x6d, 0x9f, 0x8b, 0x7f, 0x14, 0xf6, 0xd7, 0xc7, 0xf3, 0xb2,
	0xed, 0xf4, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0xa5, 0x09, 0xfe, 0x69, 0x04, 0x00,
	0x00,
}
