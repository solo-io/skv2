// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

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
	ApiGroup *types.StringValue `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Kind of the resource being referenced
	Kind *types.StringValue `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
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

func (m *TypedObjectRef) GetApiGroup() *types.StringValue {
	if m != nil {
		return m.ApiGroup
	}
	return nil
}

func (m *TypedObjectRef) GetKind() *types.StringValue {
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
	ApiGroup *types.StringValue `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Kind of the resource being referenced
	Kind *types.StringValue `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
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

func (m *TypedClusterObjectRef) GetApiGroup() *types.StringValue {
	if m != nil {
		return m.ApiGroup
	}
	return nil
}

func (m *TypedClusterObjectRef) GetKind() *types.StringValue {
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
	ProcessingTime *types.Timestamp `protobuf:"bytes,4,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
	// (optional) The owner of the status, this value can be used to identify the entity which wrote this status.
	// This is useful in situations where a given resource may have multiple owners.
	Owner                *types.StringValue `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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

func (m *Status) GetProcessingTime() *types.Timestamp {
	if m != nil {
		return m.ProcessingTime
	}
	return nil
}

func (m *Status) GetOwner() *types.StringValue {
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
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x66, 0xf3, 0xd7, 0x66, 0x52, 0x85, 0x60, 0x40, 0x8a, 0xa2, 0xaa, 0x85, 0x9c, 0xb8, 0xb0,
	0x4b, 0x03, 0x1c, 0x38, 0x80, 0x14, 0x92, 0x10, 0x45, 0x54, 0x69, 0xb4, 0x89, 0x7a, 0xe0, 0x12,
	0x39, 0x9b, 0xa9, 0x31, 0xc9, 0xae, 0x2d, 0xdb, 0x9b, 0x96, 0x37, 0xe2, 0xc6, 0x95, 0xc7, 0xe0,
	0xca, 0x95, 0x77, 0xe0, 0x8e, 0xd6, 0x9b, 0x1f, 0xd1, 0x55, 0xd5, 0x03, 0x97, 0x9e, 0x66, 0xfc,
	0xcd, 0x37, 0xf6, 0xcc, 0x37, 0xb6, 0xc1, 0x65, 0xdc, 0x7c, 0x8e, 0x67, 0x6e, 0x20, 0x42, 0x4f,
	0x8b, 0xa5, 0x78, 0xce, 0x85, 0xa7, 0x17, 0xab, 0x96, 0x47, 0x25, 0xf7, 0x02, 0xa1, 0xd0, 0x5b,
	0x9d, 0x58, 0xeb, 0x4a, 0x25, 0x8c, 0x20, 0x0f, 0xac, 0x9f, 0x30, 0xdc, 0x84, 0xee, 0x72, 0xd1,
	0x38, 0x62, 0x42, 0xb0, 0x25, 0x7a, 0x96, 0x30, 0x8b, 0x2f, 0xbc, 0x4b, 0x45, 0xa5, 0x44, 0xa5,
	0xd3, 0x94, 0xc6, 0xf1, 0xf5, 0xb8, 0xe1, 0x21, 0x6a, 0x43, 0x43, 0xb9, 0x26, 0x3c, 0x62, 0x82,
	0x09, 0xeb, 0x7a, 0x89, 0xb7, 0x46, 0x09, 0x5e, 0x99, 0x14, 0xc4, 0x2b, 0x93, 0x62, 0xcd, 0xb7,
	0x50, 0x3e, 0x9b, 0x7d, 0xc1, 0xc0, 0xf8, 0x78, 0x41, 0x08, 0x14, 0x22, 0x1a, 0x62, 0xdd, 0x79,
	0xe2, 0x3c, 0x2b, 0xfb, 0xd6, 0x27, 0x87, 0x50, 0x4e, 0xac, 0x96, 0x34, 0xc0, 0x7a, 0xce, 0x06,
	0x76, 0x40, 0x93, 0x41, 0xad, 0xb3, 0x8c, 0xb5, 0x41, 0xf5, 0x1f, 0xbb, 0x90, 0xa7, 0x70, 0x10,
	0xa4, 0xbb, 0x4c, 0x6d, 0x66, 0xde, 0x12, 0x2a, 0x6b, 0x6c, 0x48, 0x43, 0x6c, 0x7e, 0x77, 0xa0,
	0x3a, 0xf9, 0x2a, 0x71, 0xbe, 0x3b, 0xe7, 0x0d, 0x94, 0xa9, 0xe4, 0x53, 0xa6, 0x44, 0x2c, 0xed,
	0x61, 0x95, 0xd6, 0xa1, 0x9b, 0x2a, 0xe3, 0x6e, 0x94, 0x71, 0xc7, 0x46, 0xf1, 0x88, 0x9d, 0xd3,
	0x65, 0x8c, 0xfe, 0x3e, 0x95, 0xbc, 0x9f, 0xb0, 0xc9, 0x0b, 0x28, 0x2c, 0x78, 0x34, 0xb7, 0x95,
	0xdc, 0x96, 0x65, 0x99, 0xdb, 0xa6, 0xf2, 0x37, 0x35, 0x55, 0xb8, 0x2e, 0xcd, 0x2f, 0x07, 0x1e,
	0xdb, 0x8a, 0x33, 0x02, 0xdd, 0xed, 0xc2, 0x33, 0xd3, 0x28, 0x66, 0xa7, 0xf1, 0x33, 0x07, 0xa5,
	0xb1, 0xa1, 0x26, 0xd6, 0xe4, 0x35, 0x14, 0xb5, 0xa1, 0x26, 0x1d, 0x77, 0xb5, 0x75, 0xec, 0x66,
	0xae, 0xb3, 0x9b, 0x32, 0xad, 0x41, 0x3f, 0x65, 0x93, 0x3a, 0xec, 0x85, 0xa8, 0x35, 0x65, 0x9b,
	0xeb, 0xb0, 0x59, 0x12, 0x0f, 0x1e, 0x8a, 0x99, 0x46, 0xb5, 0xc2, 0xf9, 0x94, 0x61, 0x84, 0x8a,
	0x1a, 0x2e, 0x22, 0x5b, 0x7f, 0xde, 0x27, 0x9b, 0x50, 0x7f, 0x1b, 0x21, 0x1d, 0xb8, 0x2f, 0x95,
	0x08, 0x50, 0x6b, 0x1e, 0xb1, 0x69, 0xf2, 0x14, 0x6c, 0x4f, 0x95, 0x56, 0x23, 0x23, 0xcf, 0x64,
	0xf3, 0x4e, 0xfc, 0xea, 0x2e, 0x25, 0x01, 0x49, 0x0b, 0x8a, 0xe2, 0x32, 0x42, 0x65, 0xbb, 0xbd,
	0x4d, 0xd9, 0x94, 0xda, 0xfc, 0x08, 0x45, 0xdb, 0x13, 0xa9, 0xc0, 0xde, 0xa8, 0x37, 0xec, 0x0e,
	0x86, 0xfd, 0xda, 0x3d, 0x52, 0x05, 0x18, 0xf9, 0x67, 0x9d, 0xde, 0x78, 0x9c, 0xac, 0x9d, 0x24,
	0x38, 0x18, 0x9e, 0xb7, 0x4f, 0x07, 0xdd, 0x5a, 0x8e, 0x00, 0x94, 0x3e, 0xb4, 0x07, 0xa7, 0xbd,
	0x6e, 0x2d, 0x4f, 0x0e, 0x60, 0xbf, 0xdd, 0xe9, 0xf4, 0x46, 0x93, 0x5e, 0xb7, 0x56, 0x78, 0xff,
	0xee, 0xc7, 0x9f, 0x82, 0xf3, 0xed, 0xf7, 0x91, 0xf3, 0xe9, 0xd5, 0x4d, 0x1f, 0x88, 0x5c, 0xb0,
	0xed, 0x27, 0xf2, 0x8f, 0xc2, 0xde, 0xea, 0x64, 0x56, 0xb2, 0x95, 0xbe, 0xfc, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x29, 0x0a, 0x7d, 0xb2, 0x7f, 0x04, 0x00, 0x00,
}

func (this *ObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ObjectRef)
	if !ok {
		that2, ok := that.(ObjectRef)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ClusterObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterObjectRef)
	if !ok {
		that2, ok := that.(ClusterObjectRef)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.ClusterName != that1.ClusterName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TypedObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TypedObjectRef)
	if !ok {
		that2, ok := that.(TypedObjectRef)
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
	if !this.ApiGroup.Equal(that1.ApiGroup) {
		return false
	}
	if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TypedClusterObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TypedClusterObjectRef)
	if !ok {
		that2, ok := that.(TypedClusterObjectRef)
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
	if !this.ApiGroup.Equal(that1.ApiGroup) {
		return false
	}
	if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.ClusterName != that1.ClusterName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Status) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
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
	if this.State != that1.State {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !this.ProcessingTime.Equal(that1.ProcessingTime) {
		return false
	}
	if !this.Owner.Equal(that1.Owner) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
