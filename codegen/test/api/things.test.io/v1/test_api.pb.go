// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/solo-io/skv2/codegen/test/test_api.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type AcrylicType_Body int32

const (
	AcrylicType_Light  AcrylicType_Body = 0
	AcrylicType_Medium AcrylicType_Body = 1
	AcrylicType_Heavy  AcrylicType_Body = 2
)

var AcrylicType_Body_name = map[int32]string{
	0: "Light",
	1: "Medium",
	2: "Heavy",
}

var AcrylicType_Body_value = map[string]int32{
	"Light":  0,
	"Medium": 1,
	"Heavy":  2,
}

func (x AcrylicType_Body) String() string {
	return proto.EnumName(AcrylicType_Body_name, int32(x))
}

func (AcrylicType_Body) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{2, 0}
}

type PaintSpec struct {
	Color *PaintColor `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	// Types that are valid to be assigned to PaintType:
	//	*PaintSpec_Acrylic
	//	*PaintSpec_Oil
	PaintType            isPaintSpec_PaintType `protobuf_oneof:"paintType"`
	MyFavorite           *any.Any              `protobuf:"bytes,4,opt,name=my_favorite,json=myFavorite,proto3" json:"my_favorite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaintSpec) Reset()         { *m = PaintSpec{} }
func (m *PaintSpec) String() string { return proto.CompactTextString(m) }
func (*PaintSpec) ProtoMessage()    {}
func (*PaintSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{0}
}

func (m *PaintSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintSpec.Unmarshal(m, b)
}
func (m *PaintSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintSpec.Marshal(b, m, deterministic)
}
func (m *PaintSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintSpec.Merge(m, src)
}
func (m *PaintSpec) XXX_Size() int {
	return xxx_messageInfo_PaintSpec.Size(m)
}
func (m *PaintSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PaintSpec proto.InternalMessageInfo

func (m *PaintSpec) GetColor() *PaintColor {
	if m != nil {
		return m.Color
	}
	return nil
}

type isPaintSpec_PaintType interface {
	isPaintSpec_PaintType()
}

type PaintSpec_Acrylic struct {
	Acrylic *AcrylicType `protobuf:"bytes,2,opt,name=acrylic,proto3,oneof"`
}

type PaintSpec_Oil struct {
	Oil *OilType `protobuf:"bytes,3,opt,name=oil,proto3,oneof"`
}

func (*PaintSpec_Acrylic) isPaintSpec_PaintType() {}

func (*PaintSpec_Oil) isPaintSpec_PaintType() {}

func (m *PaintSpec) GetPaintType() isPaintSpec_PaintType {
	if m != nil {
		return m.PaintType
	}
	return nil
}

func (m *PaintSpec) GetAcrylic() *AcrylicType {
	if x, ok := m.GetPaintType().(*PaintSpec_Acrylic); ok {
		return x.Acrylic
	}
	return nil
}

func (m *PaintSpec) GetOil() *OilType {
	if x, ok := m.GetPaintType().(*PaintSpec_Oil); ok {
		return x.Oil
	}
	return nil
}

func (m *PaintSpec) GetMyFavorite() *any.Any {
	if m != nil {
		return m.MyFavorite
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PaintSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PaintSpec_Acrylic)(nil),
		(*PaintSpec_Oil)(nil),
	}
}

type PaintColor struct {
	Hue                  string   `protobuf:"bytes,1,opt,name=hue,proto3" json:"hue,omitempty"`
	Value                float32  `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaintColor) Reset()         { *m = PaintColor{} }
func (m *PaintColor) String() string { return proto.CompactTextString(m) }
func (*PaintColor) ProtoMessage()    {}
func (*PaintColor) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{1}
}

func (m *PaintColor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintColor.Unmarshal(m, b)
}
func (m *PaintColor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintColor.Marshal(b, m, deterministic)
}
func (m *PaintColor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintColor.Merge(m, src)
}
func (m *PaintColor) XXX_Size() int {
	return xxx_messageInfo_PaintColor.Size(m)
}
func (m *PaintColor) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintColor.DiscardUnknown(m)
}

var xxx_messageInfo_PaintColor proto.InternalMessageInfo

func (m *PaintColor) GetHue() string {
	if m != nil {
		return m.Hue
	}
	return ""
}

func (m *PaintColor) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AcrylicType struct {
	Body                 AcrylicType_Body `protobuf:"varint,3,opt,name=body,proto3,enum=things.test.io.AcrylicType_Body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AcrylicType) Reset()         { *m = AcrylicType{} }
func (m *AcrylicType) String() string { return proto.CompactTextString(m) }
func (*AcrylicType) ProtoMessage()    {}
func (*AcrylicType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{2}
}

func (m *AcrylicType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcrylicType.Unmarshal(m, b)
}
func (m *AcrylicType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcrylicType.Marshal(b, m, deterministic)
}
func (m *AcrylicType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcrylicType.Merge(m, src)
}
func (m *AcrylicType) XXX_Size() int {
	return xxx_messageInfo_AcrylicType.Size(m)
}
func (m *AcrylicType) XXX_DiscardUnknown() {
	xxx_messageInfo_AcrylicType.DiscardUnknown(m)
}

var xxx_messageInfo_AcrylicType proto.InternalMessageInfo

func (m *AcrylicType) GetBody() AcrylicType_Body {
	if m != nil {
		return m.Body
	}
	return AcrylicType_Light
}

type OilType struct {
	WaterMixable bool `protobuf:"varint,1,opt,name=waterMixable,proto3" json:"waterMixable,omitempty"`
	// Types that are valid to be assigned to PigmentType:
	//	*OilType_Powder
	//	*OilType_Fluid
	PigmentType          isOilType_PigmentType `protobuf_oneof:"pigmentType"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OilType) Reset()         { *m = OilType{} }
func (m *OilType) String() string { return proto.CompactTextString(m) }
func (*OilType) ProtoMessage()    {}
func (*OilType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{3}
}

func (m *OilType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OilType.Unmarshal(m, b)
}
func (m *OilType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OilType.Marshal(b, m, deterministic)
}
func (m *OilType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OilType.Merge(m, src)
}
func (m *OilType) XXX_Size() int {
	return xxx_messageInfo_OilType.Size(m)
}
func (m *OilType) XXX_DiscardUnknown() {
	xxx_messageInfo_OilType.DiscardUnknown(m)
}

var xxx_messageInfo_OilType proto.InternalMessageInfo

func (m *OilType) GetWaterMixable() bool {
	if m != nil {
		return m.WaterMixable
	}
	return false
}

type isOilType_PigmentType interface {
	isOilType_PigmentType()
}

type OilType_Powder struct {
	Powder string `protobuf:"bytes,2,opt,name=powder,proto3,oneof"`
}

type OilType_Fluid struct {
	Fluid string `protobuf:"bytes,3,opt,name=fluid,proto3,oneof"`
}

func (*OilType_Powder) isOilType_PigmentType() {}

func (*OilType_Fluid) isOilType_PigmentType() {}

func (m *OilType) GetPigmentType() isOilType_PigmentType {
	if m != nil {
		return m.PigmentType
	}
	return nil
}

func (m *OilType) GetPowder() string {
	if x, ok := m.GetPigmentType().(*OilType_Powder); ok {
		return x.Powder
	}
	return ""
}

func (m *OilType) GetFluid() string {
	if x, ok := m.GetPigmentType().(*OilType_Fluid); ok {
		return x.Fluid
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OilType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OilType_Powder)(nil),
		(*OilType_Fluid)(nil),
	}
}

type PaintStatus struct {
	ObservedGeneration   int64    `protobuf:"varint,1,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	PercentRemaining     int64    `protobuf:"varint,2,opt,name=percentRemaining,proto3" json:"percentRemaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaintStatus) Reset()         { *m = PaintStatus{} }
func (m *PaintStatus) String() string { return proto.CompactTextString(m) }
func (*PaintStatus) ProtoMessage()    {}
func (*PaintStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{4}
}

func (m *PaintStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintStatus.Unmarshal(m, b)
}
func (m *PaintStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintStatus.Marshal(b, m, deterministic)
}
func (m *PaintStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintStatus.Merge(m, src)
}
func (m *PaintStatus) XXX_Size() int {
	return xxx_messageInfo_PaintStatus.Size(m)
}
func (m *PaintStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PaintStatus proto.InternalMessageInfo

func (m *PaintStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *PaintStatus) GetPercentRemaining() int64 {
	if m != nil {
		return m.PercentRemaining
	}
	return 0
}

type ClusterResourceSpec struct {
	Imported             *wrappers.StringValue `protobuf:"bytes,1,opt,name=imported,proto3" json:"imported,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterResourceSpec) Reset()         { *m = ClusterResourceSpec{} }
func (m *ClusterResourceSpec) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceSpec) ProtoMessage()    {}
func (*ClusterResourceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0310f6dbe7aaf4b, []int{5}
}

func (m *ClusterResourceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceSpec.Unmarshal(m, b)
}
func (m *ClusterResourceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceSpec.Marshal(b, m, deterministic)
}
func (m *ClusterResourceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceSpec.Merge(m, src)
}
func (m *ClusterResourceSpec) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceSpec.Size(m)
}
func (m *ClusterResourceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceSpec proto.InternalMessageInfo

func (m *ClusterResourceSpec) GetImported() *wrappers.StringValue {
	if m != nil {
		return m.Imported
	}
	return nil
}

func init() {
	proto.RegisterEnum("things.test.io.AcrylicType_Body", AcrylicType_Body_name, AcrylicType_Body_value)
	proto.RegisterType((*PaintSpec)(nil), "things.test.io.PaintSpec")
	proto.RegisterType((*PaintColor)(nil), "things.test.io.PaintColor")
	proto.RegisterType((*AcrylicType)(nil), "things.test.io.AcrylicType")
	proto.RegisterType((*OilType)(nil), "things.test.io.OilType")
	proto.RegisterType((*PaintStatus)(nil), "things.test.io.PaintStatus")
	proto.RegisterType((*ClusterResourceSpec)(nil), "things.test.io.ClusterResourceSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/codegen/test/test_api.proto", fileDescriptor_d0310f6dbe7aaf4b)
}

var fileDescriptor_d0310f6dbe7aaf4b = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xed, 0xf7, 0xd6, 0x1b, 0x98, 0x2a, 0x33, 0x41, 0x29, 0x08, 0x4d, 0x79, 0x9a, 0x40, 0x38,
	0xd0, 0x0d, 0x81, 0x10, 0x2f, 0xeb, 0x24, 0xe8, 0x03, 0xd3, 0x90, 0x87, 0x78, 0xe0, 0x65, 0x72,
	0x92, 0xdb, 0xd4, 0x22, 0xb1, 0x2d, 0xc7, 0x49, 0xc9, 0x2f, 0xe6, 0x6f, 0xa0, 0x38, 0x19, 0xd0,
	0x0e, 0x21, 0x5e, 0xa2, 0xdc, 0x7b, 0xce, 0xb1, 0xcf, 0xc9, 0xbd, 0x81, 0x93, 0x44, 0xd8, 0x75,
	0x11, 0xd2, 0x48, 0x65, 0x41, 0xae, 0x52, 0xf5, 0x5c, 0xa8, 0x20, 0xff, 0x56, 0xce, 0x83, 0x48,
	0xc5, 0x98, 0xa0, 0x0c, 0x2c, 0xe6, 0xd6, 0x3d, 0xae, 0xb9, 0x16, 0x54, 0x1b, 0x65, 0x15, 0x39,
	0xb0, 0x6b, 0x21, 0x93, 0x9c, 0xd6, 0x6d, 0x2a, 0xd4, 0xec, 0x49, 0xa2, 0x54, 0x92, 0x62, 0xe0,
	0xd0, 0xb0, 0x58, 0x05, 0x1b, 0xc3, 0xb5, 0x46, 0x93, 0x37, 0xfc, 0xd9, 0xc3, 0x5d, 0x9c, 0xcb,
	0xaa, 0x81, 0xfc, 0x1f, 0x5d, 0x18, 0x7f, 0xe2, 0x42, 0xda, 0x2b, 0x8d, 0x11, 0x79, 0x01, 0xc3,
	0x48, 0xa5, 0xca, 0x4c, 0xbb, 0x47, 0xdd, 0x63, 0x6f, 0x3e, 0xa3, 0xdb, 0x17, 0x51, 0xc7, 0x3c,
	0xaf, 0x19, 0xac, 0x21, 0x92, 0xd7, 0xb0, 0xc7, 0x23, 0x53, 0xa5, 0x22, 0x9a, 0xf6, 0x9c, 0xe6,
	0xd1, 0xae, 0xe6, 0xac, 0x81, 0x3f, 0x57, 0x1a, 0x97, 0x1d, 0x76, 0xc3, 0x26, 0xcf, 0xa0, 0xaf,
	0x44, 0x3a, 0xed, 0x3b, 0xd1, 0x83, 0x5d, 0xd1, 0xa5, 0x48, 0x5b, 0x41, 0xcd, 0x22, 0xaf, 0xc0,
	0xcb, 0xaa, 0xeb, 0x15, 0x2f, 0x95, 0x11, 0x16, 0xa7, 0x03, 0x27, 0x3a, 0xa4, 0x4d, 0x2c, 0x7a,
	0x13, 0x8b, 0x9e, 0xc9, 0x8a, 0x41, 0x56, 0xbd, 0x6f, 0x79, 0x0b, 0x0f, 0xc6, 0xba, 0x76, 0x5c,
	0x1f, 0xe5, 0x9f, 0x02, 0xfc, 0xb6, 0x4f, 0x26, 0xd0, 0x5f, 0x17, 0xe8, 0x72, 0x8e, 0x59, 0xfd,
	0x4a, 0x0e, 0x61, 0x58, 0xf2, 0xb4, 0x40, 0x97, 0xa3, 0xc7, 0x9a, 0xc2, 0xcf, 0xc0, 0xfb, 0x23,
	0x00, 0x39, 0x85, 0x41, 0xa8, 0xe2, 0xca, 0xd9, 0x3e, 0x98, 0x1f, 0xfd, 0x23, 0x2b, 0x5d, 0xa8,
	0xb8, 0x62, 0x8e, 0xed, 0x1f, 0xc3, 0xa0, 0xae, 0xc8, 0x18, 0x86, 0x1f, 0x45, 0xb2, 0xb6, 0x93,
	0x0e, 0x01, 0x18, 0x5d, 0x60, 0x2c, 0x8a, 0x6c, 0xd2, 0xad, 0xdb, 0x4b, 0xe4, 0x65, 0x35, 0xe9,
	0xf9, 0x12, 0xf6, 0xda, 0xe8, 0xc4, 0x87, 0x3b, 0x1b, 0x6e, 0xd1, 0x5c, 0x88, 0xef, 0x3c, 0x4c,
	0x1b, 0xab, 0xfb, 0x6c, 0xab, 0x47, 0xa6, 0x30, 0xd2, 0x6a, 0x13, 0xa3, 0x71, 0xa6, 0xc7, 0xcb,
	0x0e, 0x6b, 0x6b, 0x72, 0x1f, 0x86, 0xab, 0xb4, 0x10, 0xb1, 0x73, 0x5a, 0x03, 0x4d, 0xb9, 0xb8,
	0x0b, 0x9e, 0x16, 0x49, 0x86, 0xed, 0x47, 0x11, 0xe0, 0x35, 0xd3, 0xb7, 0xdc, 0x16, 0x39, 0xa1,
	0x40, 0x54, 0x98, 0xa3, 0x29, 0x31, 0xfe, 0x80, 0x12, 0x0d, 0xb7, 0x42, 0x49, 0x77, 0x73, 0x9f,
	0xfd, 0x05, 0x21, 0x4f, 0x61, 0xa2, 0xd1, 0x44, 0x28, 0x2d, 0xc3, 0x8c, 0x0b, 0x29, 0x64, 0xe2,
	0x9c, 0xf4, 0xd9, 0xad, 0xbe, 0x7f, 0x09, 0xf7, 0xce, 0xd3, 0x22, 0xb7, 0x68, 0x18, 0xe6, 0xaa,
	0x30, 0x11, 0xba, 0x95, 0x7b, 0x03, 0xfb, 0x22, 0xd3, 0xca, 0x58, 0x8c, 0xdb, 0xad, 0x7b, 0x7c,
	0x6b, 0xae, 0x57, 0xd6, 0x08, 0x99, 0x7c, 0xa9, 0x07, 0xc2, 0x7e, 0xb1, 0x17, 0xef, 0xbe, 0xbe,
	0xfd, 0xaf, 0x9f, 0x87, 0x6b, 0x11, 0x6c, 0xcf, 0x29, 0x28, 0x5f, 0x86, 0x23, 0x77, 0xfa, 0xc9,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x47, 0xea, 0xb5, 0x81, 0x03, 0x00, 0x00,
}
