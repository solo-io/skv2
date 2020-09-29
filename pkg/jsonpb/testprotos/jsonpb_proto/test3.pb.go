package jsonpb_proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Numeral int32

const (
	Numeral_UNKNOWN Numeral = 0
	Numeral_ARABIC  Numeral = 1
	Numeral_ROMAN   Numeral = 2
)

var Numeral_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARABIC",
	2: "ROMAN",
}

var Numeral_value = map[string]int32{
	"UNKNOWN": 0,
	"ARABIC":  1,
	"ROMAN":   2,
}

func (x Numeral) String() string {
	return proto.EnumName(Numeral_name, int32(x))
}

func (Numeral) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_813baf511b225405, []int{0}
}

type Simple3 struct {
	Dub                  float64  `protobuf:"fixed64,1,opt,name=dub,proto3" json:"dub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Simple3) Reset()         { *m = Simple3{} }
func (m *Simple3) String() string { return proto.CompactTextString(m) }
func (*Simple3) ProtoMessage()    {}
func (*Simple3) Descriptor() ([]byte, []int) {
	return fileDescriptor_813baf511b225405, []int{0}
}

func (m *Simple3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple3.Unmarshal(m, b)
}
func (m *Simple3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple3.Marshal(b, m, deterministic)
}
func (m *Simple3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple3.Merge(m, src)
}
func (m *Simple3) XXX_Size() int {
	return xxx_messageInfo_Simple3.Size(m)
}
func (m *Simple3) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple3.DiscardUnknown(m)
}

var xxx_messageInfo_Simple3 proto.InternalMessageInfo

func (m *Simple3) GetDub() float64 {
	if m != nil {
		return m.Dub
	}
	return 0
}

type SimpleSlice3 struct {
	Slices               []string `protobuf:"bytes,1,rep,name=slices,proto3" json:"slices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleSlice3) Reset()         { *m = SimpleSlice3{} }
func (m *SimpleSlice3) String() string { return proto.CompactTextString(m) }
func (*SimpleSlice3) ProtoMessage()    {}
func (*SimpleSlice3) Descriptor() ([]byte, []int) {
	return fileDescriptor_813baf511b225405, []int{1}
}

func (m *SimpleSlice3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleSlice3.Unmarshal(m, b)
}
func (m *SimpleSlice3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleSlice3.Marshal(b, m, deterministic)
}
func (m *SimpleSlice3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleSlice3.Merge(m, src)
}
func (m *SimpleSlice3) XXX_Size() int {
	return xxx_messageInfo_SimpleSlice3.Size(m)
}
func (m *SimpleSlice3) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleSlice3.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleSlice3 proto.InternalMessageInfo

func (m *SimpleSlice3) GetSlices() []string {
	if m != nil {
		return m.Slices
	}
	return nil
}

type SimpleMap3 struct {
	Stringy              map[string]string `protobuf:"bytes,1,rep,name=stringy,proto3" json:"stringy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SimpleMap3) Reset()         { *m = SimpleMap3{} }
func (m *SimpleMap3) String() string { return proto.CompactTextString(m) }
func (*SimpleMap3) ProtoMessage()    {}
func (*SimpleMap3) Descriptor() ([]byte, []int) {
	return fileDescriptor_813baf511b225405, []int{2}
}

func (m *SimpleMap3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMap3.Unmarshal(m, b)
}
func (m *SimpleMap3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMap3.Marshal(b, m, deterministic)
}
func (m *SimpleMap3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMap3.Merge(m, src)
}
func (m *SimpleMap3) XXX_Size() int {
	return xxx_messageInfo_SimpleMap3.Size(m)
}
func (m *SimpleMap3) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMap3.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMap3 proto.InternalMessageInfo

func (m *SimpleMap3) GetStringy() map[string]string {
	if m != nil {
		return m.Stringy
	}
	return nil
}

type SimpleNull3 struct {
	Simple               *Simple3 `protobuf:"bytes,1,opt,name=simple,proto3" json:"simple,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleNull3) Reset()         { *m = SimpleNull3{} }
func (m *SimpleNull3) String() string { return proto.CompactTextString(m) }
func (*SimpleNull3) ProtoMessage()    {}
func (*SimpleNull3) Descriptor() ([]byte, []int) {
	return fileDescriptor_813baf511b225405, []int{3}
}

func (m *SimpleNull3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleNull3.Unmarshal(m, b)
}
func (m *SimpleNull3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleNull3.Marshal(b, m, deterministic)
}
func (m *SimpleNull3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleNull3.Merge(m, src)
}
func (m *SimpleNull3) XXX_Size() int {
	return xxx_messageInfo_SimpleNull3.Size(m)
}
func (m *SimpleNull3) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleNull3.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleNull3 proto.InternalMessageInfo

func (m *SimpleNull3) GetSimple() *Simple3 {
	if m != nil {
		return m.Simple
	}
	return nil
}

type Mappy struct {
	Nummy                map[int64]int32    `protobuf:"bytes,1,rep,name=nummy,proto3" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Strry                map[string]string  `protobuf:"bytes,2,rep,name=strry,proto3" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Objjy                map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy,proto3" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Buggy                map[int64]string   `protobuf:"bytes,4,rep,name=buggy,proto3" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Booly                map[bool]bool      `protobuf:"bytes,5,rep,name=booly,proto3" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Enumy                map[string]Numeral `protobuf:"bytes,6,rep,name=enumy,proto3" json:"enumy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=jsonpb_test.Numeral"`
	S32Booly             map[int32]bool     `protobuf:"bytes,7,rep,name=s32booly,proto3" json:"s32booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	S64Booly             map[int64]bool     `protobuf:"bytes,8,rep,name=s64booly,proto3" json:"s64booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	U32Booly             map[uint32]bool    `protobuf:"bytes,9,rep,name=u32booly,proto3" json:"u32booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	U64Booly             map[uint64]bool    `protobuf:"bytes,10,rep,name=u64booly,proto3" json:"u64booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Mappy) Reset()         { *m = Mappy{} }
func (m *Mappy) String() string { return proto.CompactTextString(m) }
func (*Mappy) ProtoMessage()    {}
func (*Mappy) Descriptor() ([]byte, []int) {
	return fileDescriptor_813baf511b225405, []int{4}
}

func (m *Mappy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mappy.Unmarshal(m, b)
}
func (m *Mappy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mappy.Marshal(b, m, deterministic)
}
func (m *Mappy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mappy.Merge(m, src)
}
func (m *Mappy) XXX_Size() int {
	return xxx_messageInfo_Mappy.Size(m)
}
func (m *Mappy) XXX_DiscardUnknown() {
	xxx_messageInfo_Mappy.DiscardUnknown(m)
}

var xxx_messageInfo_Mappy proto.InternalMessageInfo

func (m *Mappy) GetNummy() map[int64]int32 {
	if m != nil {
		return m.Nummy
	}
	return nil
}

func (m *Mappy) GetStrry() map[string]string {
	if m != nil {
		return m.Strry
	}
	return nil
}

func (m *Mappy) GetObjjy() map[int32]*Simple3 {
	if m != nil {
		return m.Objjy
	}
	return nil
}

func (m *Mappy) GetBuggy() map[int64]string {
	if m != nil {
		return m.Buggy
	}
	return nil
}

func (m *Mappy) GetBooly() map[bool]bool {
	if m != nil {
		return m.Booly
	}
	return nil
}

func (m *Mappy) GetEnumy() map[string]Numeral {
	if m != nil {
		return m.Enumy
	}
	return nil
}

func (m *Mappy) GetS32Booly() map[int32]bool {
	if m != nil {
		return m.S32Booly
	}
	return nil
}

func (m *Mappy) GetS64Booly() map[int64]bool {
	if m != nil {
		return m.S64Booly
	}
	return nil
}

func (m *Mappy) GetU32Booly() map[uint32]bool {
	if m != nil {
		return m.U32Booly
	}
	return nil
}

func (m *Mappy) GetU64Booly() map[uint64]bool {
	if m != nil {
		return m.U64Booly
	}
	return nil
}

func init() {
	proto.RegisterEnum("jsonpb_test.Numeral", Numeral_name, Numeral_value)
	proto.RegisterType((*Simple3)(nil), "jsonpb_test.Simple3")
	proto.RegisterType((*SimpleSlice3)(nil), "jsonpb_test.SimpleSlice3")
	proto.RegisterType((*SimpleMap3)(nil), "jsonpb_test.SimpleMap3")
	proto.RegisterMapType((map[string]string)(nil), "jsonpb_test.SimpleMap3.StringyEntry")
	proto.RegisterType((*SimpleNull3)(nil), "jsonpb_test.SimpleNull3")
	proto.RegisterType((*Mappy)(nil), "jsonpb_test.Mappy")
	proto.RegisterMapType((map[bool]bool)(nil), "jsonpb_test.Mappy.BoolyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "jsonpb_test.Mappy.BuggyEntry")
	proto.RegisterMapType((map[string]Numeral)(nil), "jsonpb_test.Mappy.EnumyEntry")
	proto.RegisterMapType((map[int64]int32)(nil), "jsonpb_test.Mappy.NummyEntry")
	proto.RegisterMapType((map[int32]*Simple3)(nil), "jsonpb_test.Mappy.ObjjyEntry")
	proto.RegisterMapType((map[int32]bool)(nil), "jsonpb_test.Mappy.S32boolyEntry")
	proto.RegisterMapType((map[int64]bool)(nil), "jsonpb_test.Mappy.S64boolyEntry")
	proto.RegisterMapType((map[string]string)(nil), "jsonpb_test.Mappy.StrryEntry")
	proto.RegisterMapType((map[uint32]bool)(nil), "jsonpb_test.Mappy.U32boolyEntry")
	proto.RegisterMapType((map[uint64]bool)(nil), "jsonpb_test.Mappy.U64boolyEntry")
}

func init() { proto.RegisterFile("jsonpb_proto/test3.proto", fileDescriptor_813baf511b225405) }

var fileDescriptor_813baf511b225405 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0x4d, 0xbb, 0x69, 0x9b, 0xd3, 0x5d, 0x29, 0xc3, 0x22, 0xa1, 0x22, 0x94, 0x22, 0xb2,
	0x2c, 0x9a, 0x40, 0x23, 0xb2, 0x6c, 0x55, 0x68, 0x65, 0x2f, 0x44, 0x9a, 0x42, 0x4a, 0x11, 0xbc,
	0x91, 0x64, 0x8d, 0x31, 0x35, 0x5f, 0x24, 0x19, 0x21, 0x6f, 0xe0, 0x2b, 0xf9, 0x76, 0x32, 0x1f,
	0xd9, 0x4c, 0x4a, 0x86, 0xea, 0xd5, 0x9e, 0x99, 0xf3, 0xff, 0xe5, 0x7c, 0xec, 0xbf, 0x0c, 0xe8,
	0x87, 0x22, 0x4d, 0x32, 0xef, 0x6b, 0x96, 0xa7, 0x65, 0x6a, 0x96, 0x7e, 0x51, 0x5a, 0x06, 0x8d,
	0xd1, 0x98, 0x67, 0xc8, 0xdd, 0xfc, 0x29, 0x0c, 0x77, 0x61, 0x9c, 0x45, 0xbe, 0x85, 0x26, 0xd0,
	0xff, 0x86, 0x3d, 0x5d, 0x99, 0x29, 0x57, 0x8a, 0x43, 0xc2, 0xf9, 0x0b, 0x38, 0x67, 0xc9, 0x5d,
	0x14, 0xde, 0xfb, 0x16, 0x7a, 0x02, 0x83, 0x82, 0x44, 0x85, 0xae, 0xcc, 0xfa, 0x57, 0x9a, 0xc3,
	0x4f, 0xf3, 0xdf, 0x0a, 0x00, 0x13, 0x6e, 0xdc, 0xcc, 0x42, 0xef, 0x61, 0x58, 0x94, 0x79, 0x98,
	0x04, 0x15, 0xd5, 0x8d, 0x17, 0xcf, 0x0d, 0xa1, 0xa4, 0xd1, 0x28, 0x8d, 0x1d, 0x93, 0xdd, 0x25,
	0x65, 0x5e, 0x39, 0x35, 0x34, 0xbd, 0x85, 0x73, 0x31, 0x41, 0x1a, 0xfb, 0xe9, 0x57, 0xb4, 0x31,
	0xcd, 0x21, 0x21, 0xba, 0x04, 0xf5, 0x97, 0x1b, 0x61, 0x5f, 0xef, 0xd1, 0x3b, 0x76, 0xb8, 0xed,
	0xdd, 0x28, 0xf3, 0x25, 0x8c, 0xd9, 0xf7, 0x6d, 0x1c, 0x45, 0x16, 0x7a, 0x09, 0x83, 0x82, 0x1e,
	0x29, 0x3d, 0x5e, 0x5c, 0x76, 0x74, 0x62, 0x39, 0x5c, 0x33, 0xff, 0xa3, 0x81, 0xba, 0x71, 0xb3,
	0xac, 0x42, 0x16, 0xa8, 0x09, 0x8e, 0xe3, 0x7a, 0x80, 0x67, 0x2d, 0x8c, 0x4a, 0x0c, 0x9b, 0xe4,
	0x59, 0xe7, 0x4c, 0x4b, 0xa0, 0xa2, 0xcc, 0xf3, 0x4a, 0xef, 0x49, 0xa1, 0x1d, 0xc9, 0x73, 0x88,
	0x6a, 0x09, 0x94, 0x7a, 0x87, 0x43, 0xa5, 0xf7, 0xa5, 0xd0, 0x96, 0xe4, 0x39, 0x44, 0xb5, 0x04,
	0xf2, 0x70, 0x10, 0x54, 0xfa, 0x99, 0x14, 0x5a, 0x93, 0x3c, 0x87, 0xa8, 0x96, 0x42, 0x69, 0x1a,
	0x55, 0xba, 0x2a, 0x87, 0x48, 0xbe, 0x86, 0x48, 0x4c, 0x20, 0x3f, 0xc1, 0x71, 0xa5, 0x0f, 0xa4,
	0xd0, 0x1d, 0xc9, 0x73, 0x88, 0x6a, 0xd1, 0x5b, 0x18, 0x15, 0xd6, 0x82, 0x15, 0x1b, 0x52, 0x6e,
	0xd6, 0xb5, 0x0b, 0x2e, 0x61, 0xe8, 0x03, 0x41, 0xe9, 0x37, 0xaf, 0x19, 0x3d, 0x92, 0xd3, 0x5c,
	0x52, 0xd3, 0xfc, 0x48, 0x68, 0x5c, 0xd7, 0xd6, 0xa4, 0xf4, 0xbe, 0x5d, 0x1b, 0x0b, 0xb5, 0x71,
	0x5d, 0x1b, 0xe4, 0x74, 0xbb, 0x76, 0x4d, 0x4c, 0x6f, 0x00, 0x1a, 0x57, 0x88, 0xb6, 0xed, 0x77,
	0xd8, 0x56, 0x15, 0x6c, 0x4b, 0xc8, 0xc6, 0x1a, 0xff, 0x63, 0xf8, 0xa9, 0x0d, 0xd0, 0xf8, 0x43,
	0x24, 0x55, 0x46, 0x5e, 0x8b, 0xa4, 0xec, 0x07, 0xd0, 0xee, 0xa4, 0xb1, 0xce, 0xa9, 0x19, 0xb4,
	0x63, 0xf2, 0x61, 0x2b, 0x22, 0x39, 0xea, 0x20, 0x47, 0x47, 0x33, 0x34, 0x26, 0xea, 0x98, 0xbe,
	0x35, 0xc3, 0xe3, 0xa3, 0x19, 0x6c, 0x1c, 0xfb, 0xb9, 0x1b, 0x89, 0xdf, 0x5b, 0xc2, 0x45, 0xcb,
	0x5c, 0x1d, 0x6b, 0x91, 0x37, 0x43, 0x60, 0xf1, 0xff, 0x7b, 0x6a, 0x07, 0xc7, 0xf0, 0x5e, 0x56,
	0xf9, 0xe2, 0x5f, 0x60, 0x59, 0xe5, 0xb3, 0x13, 0xf0, 0xf5, 0x2b, 0x18, 0xf2, 0x4d, 0xa0, 0x31,
	0x0c, 0xf7, 0xf6, 0x27, 0x7b, 0xfb, 0xd9, 0x9e, 0x3c, 0x42, 0x00, 0x83, 0x95, 0xb3, 0x5a, 0x7f,
	0xfc, 0x30, 0x51, 0x90, 0x06, 0xaa, 0xb3, 0xdd, 0xac, 0xec, 0x49, 0x6f, 0xfd, 0xee, 0xcb, 0x32,
	0x08, 0xcb, 0x1f, 0xd8, 0x33, 0xee, 0xd3, 0xd8, 0x0c, 0xd2, 0xc8, 0x4d, 0x02, 0x93, 0xbe, 0x0f,
	0x1e, 0xfe, 0x6e, 0x86, 0x49, 0xe9, 0xe7, 0x89, 0x1b, 0xd1, 0x77, 0x83, 0xde, 0x16, 0xa6, 0xf8,
	0x9e, 0x78, 0x03, 0xfa, 0xc7, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xda, 0x44, 0x24, 0x66,
	0x06, 0x00, 0x00,
}
