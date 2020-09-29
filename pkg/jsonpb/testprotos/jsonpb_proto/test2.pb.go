package jsonpb_proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type Widget_Color int32

const (
	Widget_RED   Widget_Color = 0
	Widget_GREEN Widget_Color = 1
	Widget_BLUE  Widget_Color = 2
)

var Widget_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

var Widget_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Widget_Color) Enum() *Widget_Color {
	p := new(Widget_Color)
	*p = x
	return p
}

func (x Widget_Color) String() string {
	return proto.EnumName(Widget_Color_name, int32(x))
}

func (x *Widget_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Widget_Color_value, data, "Widget_Color")
	if err != nil {
		return err
	}
	*x = Widget_Color(value)
	return nil
}

func (Widget_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{3, 0}
}

// Test message for holding primitive types.
type Simple struct {
	OBool                *bool    `protobuf:"varint,1,opt,name=o_bool,json=oBool" json:"o_bool,omitempty"`
	OInt32               *int32   `protobuf:"varint,2,opt,name=o_int32,json=oInt32" json:"o_int32,omitempty"`
	OInt32Str            *int32   `protobuf:"varint,3,opt,name=o_int32_str,json=oInt32Str" json:"o_int32_str,omitempty"`
	OInt64               *int64   `protobuf:"varint,4,opt,name=o_int64,json=oInt64" json:"o_int64,omitempty"`
	OInt64Str            *int64   `protobuf:"varint,5,opt,name=o_int64_str,json=oInt64Str" json:"o_int64_str,omitempty"`
	OUint32              *uint32  `protobuf:"varint,6,opt,name=o_uint32,json=oUint32" json:"o_uint32,omitempty"`
	OUint32Str           *uint32  `protobuf:"varint,7,opt,name=o_uint32_str,json=oUint32Str" json:"o_uint32_str,omitempty"`
	OUint64              *uint64  `protobuf:"varint,8,opt,name=o_uint64,json=oUint64" json:"o_uint64,omitempty"`
	OUint64Str           *uint64  `protobuf:"varint,9,opt,name=o_uint64_str,json=oUint64Str" json:"o_uint64_str,omitempty"`
	OSint32              *int32   `protobuf:"zigzag32,10,opt,name=o_sint32,json=oSint32" json:"o_sint32,omitempty"`
	OSint32Str           *int32   `protobuf:"zigzag32,11,opt,name=o_sint32_str,json=oSint32Str" json:"o_sint32_str,omitempty"`
	OSint64              *int64   `protobuf:"zigzag64,12,opt,name=o_sint64,json=oSint64" json:"o_sint64,omitempty"`
	OSint64Str           *int64   `protobuf:"zigzag64,13,opt,name=o_sint64_str,json=oSint64Str" json:"o_sint64_str,omitempty"`
	OFloat               *float32 `protobuf:"fixed32,14,opt,name=o_float,json=oFloat" json:"o_float,omitempty"`
	OFloatStr            *float32 `protobuf:"fixed32,15,opt,name=o_float_str,json=oFloatStr" json:"o_float_str,omitempty"`
	ODouble              *float64 `protobuf:"fixed64,16,opt,name=o_double,json=oDouble" json:"o_double,omitempty"`
	ODoubleStr           *float64 `protobuf:"fixed64,17,opt,name=o_double_str,json=oDoubleStr" json:"o_double_str,omitempty"`
	OString              *string  `protobuf:"bytes,18,opt,name=o_string,json=oString" json:"o_string,omitempty"`
	OBytes               []byte   `protobuf:"bytes,19,opt,name=o_bytes,json=oBytes" json:"o_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Simple) Reset()         { *m = Simple{} }
func (m *Simple) String() string { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()    {}
func (*Simple) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{0}
}

func (m *Simple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple.Unmarshal(m, b)
}
func (m *Simple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple.Marshal(b, m, deterministic)
}
func (m *Simple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple.Merge(m, src)
}
func (m *Simple) XXX_Size() int {
	return xxx_messageInfo_Simple.Size(m)
}
func (m *Simple) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple.DiscardUnknown(m)
}

var xxx_messageInfo_Simple proto.InternalMessageInfo

func (m *Simple) GetOBool() bool {
	if m != nil && m.OBool != nil {
		return *m.OBool
	}
	return false
}

func (m *Simple) GetOInt32() int32 {
	if m != nil && m.OInt32 != nil {
		return *m.OInt32
	}
	return 0
}

func (m *Simple) GetOInt32Str() int32 {
	if m != nil && m.OInt32Str != nil {
		return *m.OInt32Str
	}
	return 0
}

func (m *Simple) GetOInt64() int64 {
	if m != nil && m.OInt64 != nil {
		return *m.OInt64
	}
	return 0
}

func (m *Simple) GetOInt64Str() int64 {
	if m != nil && m.OInt64Str != nil {
		return *m.OInt64Str
	}
	return 0
}

func (m *Simple) GetOUint32() uint32 {
	if m != nil && m.OUint32 != nil {
		return *m.OUint32
	}
	return 0
}

func (m *Simple) GetOUint32Str() uint32 {
	if m != nil && m.OUint32Str != nil {
		return *m.OUint32Str
	}
	return 0
}

func (m *Simple) GetOUint64() uint64 {
	if m != nil && m.OUint64 != nil {
		return *m.OUint64
	}
	return 0
}

func (m *Simple) GetOUint64Str() uint64 {
	if m != nil && m.OUint64Str != nil {
		return *m.OUint64Str
	}
	return 0
}

func (m *Simple) GetOSint32() int32 {
	if m != nil && m.OSint32 != nil {
		return *m.OSint32
	}
	return 0
}

func (m *Simple) GetOSint32Str() int32 {
	if m != nil && m.OSint32Str != nil {
		return *m.OSint32Str
	}
	return 0
}

func (m *Simple) GetOSint64() int64 {
	if m != nil && m.OSint64 != nil {
		return *m.OSint64
	}
	return 0
}

func (m *Simple) GetOSint64Str() int64 {
	if m != nil && m.OSint64Str != nil {
		return *m.OSint64Str
	}
	return 0
}

func (m *Simple) GetOFloat() float32 {
	if m != nil && m.OFloat != nil {
		return *m.OFloat
	}
	return 0
}

func (m *Simple) GetOFloatStr() float32 {
	if m != nil && m.OFloatStr != nil {
		return *m.OFloatStr
	}
	return 0
}

func (m *Simple) GetODouble() float64 {
	if m != nil && m.ODouble != nil {
		return *m.ODouble
	}
	return 0
}

func (m *Simple) GetODoubleStr() float64 {
	if m != nil && m.ODoubleStr != nil {
		return *m.ODoubleStr
	}
	return 0
}

func (m *Simple) GetOString() string {
	if m != nil && m.OString != nil {
		return *m.OString
	}
	return ""
}

func (m *Simple) GetOBytes() []byte {
	if m != nil {
		return m.OBytes
	}
	return nil
}

// Test message for holding special non-finites primitives.
type NonFinites struct {
	FNan                 *float32 `protobuf:"fixed32,1,opt,name=f_nan,json=fNan" json:"f_nan,omitempty"`
	FPinf                *float32 `protobuf:"fixed32,2,opt,name=f_pinf,json=fPinf" json:"f_pinf,omitempty"`
	FNinf                *float32 `protobuf:"fixed32,3,opt,name=f_ninf,json=fNinf" json:"f_ninf,omitempty"`
	DNan                 *float64 `protobuf:"fixed64,4,opt,name=d_nan,json=dNan" json:"d_nan,omitempty"`
	DPinf                *float64 `protobuf:"fixed64,5,opt,name=d_pinf,json=dPinf" json:"d_pinf,omitempty"`
	DNinf                *float64 `protobuf:"fixed64,6,opt,name=d_ninf,json=dNinf" json:"d_ninf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonFinites) Reset()         { *m = NonFinites{} }
func (m *NonFinites) String() string { return proto.CompactTextString(m) }
func (*NonFinites) ProtoMessage()    {}
func (*NonFinites) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{1}
}

func (m *NonFinites) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonFinites.Unmarshal(m, b)
}
func (m *NonFinites) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonFinites.Marshal(b, m, deterministic)
}
func (m *NonFinites) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonFinites.Merge(m, src)
}
func (m *NonFinites) XXX_Size() int {
	return xxx_messageInfo_NonFinites.Size(m)
}
func (m *NonFinites) XXX_DiscardUnknown() {
	xxx_messageInfo_NonFinites.DiscardUnknown(m)
}

var xxx_messageInfo_NonFinites proto.InternalMessageInfo

func (m *NonFinites) GetFNan() float32 {
	if m != nil && m.FNan != nil {
		return *m.FNan
	}
	return 0
}

func (m *NonFinites) GetFPinf() float32 {
	if m != nil && m.FPinf != nil {
		return *m.FPinf
	}
	return 0
}

func (m *NonFinites) GetFNinf() float32 {
	if m != nil && m.FNinf != nil {
		return *m.FNinf
	}
	return 0
}

func (m *NonFinites) GetDNan() float64 {
	if m != nil && m.DNan != nil {
		return *m.DNan
	}
	return 0
}

func (m *NonFinites) GetDPinf() float64 {
	if m != nil && m.DPinf != nil {
		return *m.DPinf
	}
	return 0
}

func (m *NonFinites) GetDNinf() float64 {
	if m != nil && m.DNinf != nil {
		return *m.DNinf
	}
	return 0
}

// Test message for holding repeated primitives.
type Repeats struct {
	RBool                []bool    `protobuf:"varint,1,rep,name=r_bool,json=rBool" json:"r_bool,omitempty"`
	RInt32               []int32   `protobuf:"varint,2,rep,name=r_int32,json=rInt32" json:"r_int32,omitempty"`
	RInt64               []int64   `protobuf:"varint,3,rep,name=r_int64,json=rInt64" json:"r_int64,omitempty"`
	RUint32              []uint32  `protobuf:"varint,4,rep,name=r_uint32,json=rUint32" json:"r_uint32,omitempty"`
	RUint64              []uint64  `protobuf:"varint,5,rep,name=r_uint64,json=rUint64" json:"r_uint64,omitempty"`
	RSint32              []int32   `protobuf:"zigzag32,6,rep,name=r_sint32,json=rSint32" json:"r_sint32,omitempty"`
	RSint64              []int64   `protobuf:"zigzag64,7,rep,name=r_sint64,json=rSint64" json:"r_sint64,omitempty"`
	RFloat               []float32 `protobuf:"fixed32,8,rep,name=r_float,json=rFloat" json:"r_float,omitempty"`
	RDouble              []float64 `protobuf:"fixed64,9,rep,name=r_double,json=rDouble" json:"r_double,omitempty"`
	RString              []string  `protobuf:"bytes,10,rep,name=r_string,json=rString" json:"r_string,omitempty"`
	RBytes               [][]byte  `protobuf:"bytes,11,rep,name=r_bytes,json=rBytes" json:"r_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Repeats) Reset()         { *m = Repeats{} }
func (m *Repeats) String() string { return proto.CompactTextString(m) }
func (*Repeats) ProtoMessage()    {}
func (*Repeats) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{2}
}

func (m *Repeats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repeats.Unmarshal(m, b)
}
func (m *Repeats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repeats.Marshal(b, m, deterministic)
}
func (m *Repeats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repeats.Merge(m, src)
}
func (m *Repeats) XXX_Size() int {
	return xxx_messageInfo_Repeats.Size(m)
}
func (m *Repeats) XXX_DiscardUnknown() {
	xxx_messageInfo_Repeats.DiscardUnknown(m)
}

var xxx_messageInfo_Repeats proto.InternalMessageInfo

func (m *Repeats) GetRBool() []bool {
	if m != nil {
		return m.RBool
	}
	return nil
}

func (m *Repeats) GetRInt32() []int32 {
	if m != nil {
		return m.RInt32
	}
	return nil
}

func (m *Repeats) GetRInt64() []int64 {
	if m != nil {
		return m.RInt64
	}
	return nil
}

func (m *Repeats) GetRUint32() []uint32 {
	if m != nil {
		return m.RUint32
	}
	return nil
}

func (m *Repeats) GetRUint64() []uint64 {
	if m != nil {
		return m.RUint64
	}
	return nil
}

func (m *Repeats) GetRSint32() []int32 {
	if m != nil {
		return m.RSint32
	}
	return nil
}

func (m *Repeats) GetRSint64() []int64 {
	if m != nil {
		return m.RSint64
	}
	return nil
}

func (m *Repeats) GetRFloat() []float32 {
	if m != nil {
		return m.RFloat
	}
	return nil
}

func (m *Repeats) GetRDouble() []float64 {
	if m != nil {
		return m.RDouble
	}
	return nil
}

func (m *Repeats) GetRString() []string {
	if m != nil {
		return m.RString
	}
	return nil
}

func (m *Repeats) GetRBytes() [][]byte {
	if m != nil {
		return m.RBytes
	}
	return nil
}

// Test message for holding enums and nested messages.
type Widget struct {
	Color                *Widget_Color  `protobuf:"varint,1,opt,name=color,enum=jsonpb_test.Widget_Color" json:"color,omitempty"`
	RColor               []Widget_Color `protobuf:"varint,2,rep,name=r_color,json=rColor,enum=jsonpb_test.Widget_Color" json:"r_color,omitempty"`
	Simple               *Simple        `protobuf:"bytes,10,opt,name=simple" json:"simple,omitempty"`
	RSimple              []*Simple      `protobuf:"bytes,11,rep,name=r_simple,json=rSimple" json:"r_simple,omitempty"`
	Repeats              *Repeats       `protobuf:"bytes,20,opt,name=repeats" json:"repeats,omitempty"`
	RRepeats             []*Repeats     `protobuf:"bytes,21,rep,name=r_repeats,json=rRepeats" json:"r_repeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Widget) Reset()         { *m = Widget{} }
func (m *Widget) String() string { return proto.CompactTextString(m) }
func (*Widget) ProtoMessage()    {}
func (*Widget) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{3}
}

func (m *Widget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Widget.Unmarshal(m, b)
}
func (m *Widget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Widget.Marshal(b, m, deterministic)
}
func (m *Widget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Widget.Merge(m, src)
}
func (m *Widget) XXX_Size() int {
	return xxx_messageInfo_Widget.Size(m)
}
func (m *Widget) XXX_DiscardUnknown() {
	xxx_messageInfo_Widget.DiscardUnknown(m)
}

var xxx_messageInfo_Widget proto.InternalMessageInfo

func (m *Widget) GetColor() Widget_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Widget_RED
}

func (m *Widget) GetRColor() []Widget_Color {
	if m != nil {
		return m.RColor
	}
	return nil
}

func (m *Widget) GetSimple() *Simple {
	if m != nil {
		return m.Simple
	}
	return nil
}

func (m *Widget) GetRSimple() []*Simple {
	if m != nil {
		return m.RSimple
	}
	return nil
}

func (m *Widget) GetRepeats() *Repeats {
	if m != nil {
		return m.Repeats
	}
	return nil
}

func (m *Widget) GetRRepeats() []*Repeats {
	if m != nil {
		return m.RRepeats
	}
	return nil
}

type Maps struct {
	MInt64Str            map[int64]string `protobuf:"bytes,1,rep,name=m_int64_str,json=mInt64Str" json:"m_int64_str,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MBoolSimple          map[bool]*Simple `protobuf:"bytes,2,rep,name=m_bool_simple,json=mBoolSimple" json:"m_bool_simple,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Maps) Reset()         { *m = Maps{} }
func (m *Maps) String() string { return proto.CompactTextString(m) }
func (*Maps) ProtoMessage()    {}
func (*Maps) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{4}
}

func (m *Maps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maps.Unmarshal(m, b)
}
func (m *Maps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maps.Marshal(b, m, deterministic)
}
func (m *Maps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maps.Merge(m, src)
}
func (m *Maps) XXX_Size() int {
	return xxx_messageInfo_Maps.Size(m)
}
func (m *Maps) XXX_DiscardUnknown() {
	xxx_messageInfo_Maps.DiscardUnknown(m)
}

var xxx_messageInfo_Maps proto.InternalMessageInfo

func (m *Maps) GetMInt64Str() map[int64]string {
	if m != nil {
		return m.MInt64Str
	}
	return nil
}

func (m *Maps) GetMBoolSimple() map[bool]*Simple {
	if m != nil {
		return m.MBoolSimple
	}
	return nil
}

type MsgWithOneof struct {
	// Types that are valid to be assigned to Union:
	//	*MsgWithOneof_Title
	//	*MsgWithOneof_Salary
	//	*MsgWithOneof_Country
	//	*MsgWithOneof_HomeAddress
	//	*MsgWithOneof_MsgWithRequired
	Union                isMsgWithOneof_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MsgWithOneof) Reset()         { *m = MsgWithOneof{} }
func (m *MsgWithOneof) String() string { return proto.CompactTextString(m) }
func (*MsgWithOneof) ProtoMessage()    {}
func (*MsgWithOneof) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{5}
}

func (m *MsgWithOneof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithOneof.Unmarshal(m, b)
}
func (m *MsgWithOneof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithOneof.Marshal(b, m, deterministic)
}
func (m *MsgWithOneof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithOneof.Merge(m, src)
}
func (m *MsgWithOneof) XXX_Size() int {
	return xxx_messageInfo_MsgWithOneof.Size(m)
}
func (m *MsgWithOneof) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithOneof.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithOneof proto.InternalMessageInfo

type isMsgWithOneof_Union interface {
	isMsgWithOneof_Union()
}

type MsgWithOneof_Title struct {
	Title string `protobuf:"bytes,1,opt,name=title,oneof"`
}

type MsgWithOneof_Salary struct {
	Salary int64 `protobuf:"varint,2,opt,name=salary,oneof"`
}

type MsgWithOneof_Country struct {
	Country string `protobuf:"bytes,3,opt,name=Country,oneof"`
}

type MsgWithOneof_HomeAddress struct {
	HomeAddress string `protobuf:"bytes,4,opt,name=home_address,json=homeAddress,oneof"`
}

type MsgWithOneof_MsgWithRequired struct {
	MsgWithRequired *MsgWithRequired `protobuf:"bytes,5,opt,name=msg_with_required,json=msgWithRequired,oneof"`
}

func (*MsgWithOneof_Title) isMsgWithOneof_Union() {}

func (*MsgWithOneof_Salary) isMsgWithOneof_Union() {}

func (*MsgWithOneof_Country) isMsgWithOneof_Union() {}

func (*MsgWithOneof_HomeAddress) isMsgWithOneof_Union() {}

func (*MsgWithOneof_MsgWithRequired) isMsgWithOneof_Union() {}

func (m *MsgWithOneof) GetUnion() isMsgWithOneof_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *MsgWithOneof) GetTitle() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Title); ok {
		return x.Title
	}
	return ""
}

func (m *MsgWithOneof) GetSalary() int64 {
	if x, ok := m.GetUnion().(*MsgWithOneof_Salary); ok {
		return x.Salary
	}
	return 0
}

func (m *MsgWithOneof) GetCountry() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Country); ok {
		return x.Country
	}
	return ""
}

func (m *MsgWithOneof) GetHomeAddress() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_HomeAddress); ok {
		return x.HomeAddress
	}
	return ""
}

func (m *MsgWithOneof) GetMsgWithRequired() *MsgWithRequired {
	if x, ok := m.GetUnion().(*MsgWithOneof_MsgWithRequired); ok {
		return x.MsgWithRequired
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MsgWithOneof) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MsgWithOneof_Title)(nil),
		(*MsgWithOneof_Salary)(nil),
		(*MsgWithOneof_Country)(nil),
		(*MsgWithOneof_HomeAddress)(nil),
		(*MsgWithOneof_MsgWithRequired)(nil),
	}
}

type Real struct {
	Value                        *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Real) Reset()         { *m = Real{} }
func (m *Real) String() string { return proto.CompactTextString(m) }
func (*Real) ProtoMessage()    {}
func (*Real) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{6}
}

var extRange_Real = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Real) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Real
}

func (m *Real) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Real.Unmarshal(m, b)
}
func (m *Real) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Real.Marshal(b, m, deterministic)
}
func (m *Real) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Real.Merge(m, src)
}
func (m *Real) XXX_Size() int {
	return xxx_messageInfo_Real.Size(m)
}
func (m *Real) XXX_DiscardUnknown() {
	xxx_messageInfo_Real.DiscardUnknown(m)
}

var xxx_messageInfo_Real proto.InternalMessageInfo

func (m *Real) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Complex struct {
	Imaginary                    *float64 `protobuf:"fixed64,1,opt,name=imaginary" json:"imaginary,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Complex) Reset()         { *m = Complex{} }
func (m *Complex) String() string { return proto.CompactTextString(m) }
func (*Complex) ProtoMessage()    {}
func (*Complex) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{7}
}

var extRange_Complex = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Complex) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Complex
}

func (m *Complex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Complex.Unmarshal(m, b)
}
func (m *Complex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Complex.Marshal(b, m, deterministic)
}
func (m *Complex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Complex.Merge(m, src)
}
func (m *Complex) XXX_Size() int {
	return xxx_messageInfo_Complex.Size(m)
}
func (m *Complex) XXX_DiscardUnknown() {
	xxx_messageInfo_Complex.DiscardUnknown(m)
}

var xxx_messageInfo_Complex proto.InternalMessageInfo

func (m *Complex) GetImaginary() float64 {
	if m != nil && m.Imaginary != nil {
		return *m.Imaginary
	}
	return 0
}

var E_Complex_RealExtension = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*Complex)(nil),
	Field:         123,
	Name:          "jsonpb_test.Complex.real_extension",
	Tag:           "bytes,123,opt,name=real_extension",
	Filename:      "jsonpb_proto/test2.proto",
}

type KnownTypes struct {
	An                   *anypb.Any              `protobuf:"bytes,14,opt,name=an" json:"an,omitempty"`
	Dur                  *durationpb.Duration    `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	St                   *structpb.Struct        `protobuf:"bytes,12,opt,name=st" json:"st,omitempty"`
	Ts                   *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Lv                   *structpb.ListValue     `protobuf:"bytes,15,opt,name=lv" json:"lv,omitempty"`
	Val                  *structpb.Value         `protobuf:"bytes,16,opt,name=val" json:"val,omitempty"`
	Dbl                  *wrapperspb.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt                  *wrapperspb.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64                  *wrapperspb.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64                  *wrapperspb.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32                  *wrapperspb.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32                  *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool                 *wrapperspb.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str                  *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes                *wrapperspb.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KnownTypes) Reset()         { *m = KnownTypes{} }
func (m *KnownTypes) String() string { return proto.CompactTextString(m) }
func (*KnownTypes) ProtoMessage()    {}
func (*KnownTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{8}
}

func (m *KnownTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KnownTypes.Unmarshal(m, b)
}
func (m *KnownTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KnownTypes.Marshal(b, m, deterministic)
}
func (m *KnownTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KnownTypes.Merge(m, src)
}
func (m *KnownTypes) XXX_Size() int {
	return xxx_messageInfo_KnownTypes.Size(m)
}
func (m *KnownTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_KnownTypes.DiscardUnknown(m)
}

var xxx_messageInfo_KnownTypes proto.InternalMessageInfo

func (m *KnownTypes) GetAn() *anypb.Any {
	if m != nil {
		return m.An
	}
	return nil
}

func (m *KnownTypes) GetDur() *durationpb.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *KnownTypes) GetSt() *structpb.Struct {
	if m != nil {
		return m.St
	}
	return nil
}

func (m *KnownTypes) GetTs() *timestamppb.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *KnownTypes) GetLv() *structpb.ListValue {
	if m != nil {
		return m.Lv
	}
	return nil
}

func (m *KnownTypes) GetVal() *structpb.Value {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *KnownTypes) GetDbl() *wrapperspb.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *KnownTypes) GetFlt() *wrapperspb.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *KnownTypes) GetI64() *wrapperspb.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *KnownTypes) GetU64() *wrapperspb.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *KnownTypes) GetI32() *wrapperspb.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *KnownTypes) GetU32() *wrapperspb.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *KnownTypes) GetBool() *wrapperspb.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *KnownTypes) GetStr() *wrapperspb.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *KnownTypes) GetBytes() *wrapperspb.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

// Test messages for marshaling/unmarshaling required fields.
type MsgWithRequired struct {
	Str                  *string  `protobuf:"bytes,1,req,name=str" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgWithRequired) Reset()         { *m = MsgWithRequired{} }
func (m *MsgWithRequired) String() string { return proto.CompactTextString(m) }
func (*MsgWithRequired) ProtoMessage()    {}
func (*MsgWithRequired) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{9}
}

func (m *MsgWithRequired) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithRequired.Unmarshal(m, b)
}
func (m *MsgWithRequired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithRequired.Marshal(b, m, deterministic)
}
func (m *MsgWithRequired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithRequired.Merge(m, src)
}
func (m *MsgWithRequired) XXX_Size() int {
	return xxx_messageInfo_MsgWithRequired.Size(m)
}
func (m *MsgWithRequired) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithRequired.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithRequired proto.InternalMessageInfo

func (m *MsgWithRequired) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

type MsgWithIndirectRequired struct {
	Subm                 *MsgWithRequired            `protobuf:"bytes,1,opt,name=subm" json:"subm,omitempty"`
	MapField             map[string]*MsgWithRequired `protobuf:"bytes,2,rep,name=map_field,json=mapField" json:"map_field,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SliceField           []*MsgWithRequired          `protobuf:"bytes,3,rep,name=slice_field,json=sliceField" json:"slice_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MsgWithIndirectRequired) Reset()         { *m = MsgWithIndirectRequired{} }
func (m *MsgWithIndirectRequired) String() string { return proto.CompactTextString(m) }
func (*MsgWithIndirectRequired) ProtoMessage()    {}
func (*MsgWithIndirectRequired) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{10}
}

func (m *MsgWithIndirectRequired) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithIndirectRequired.Unmarshal(m, b)
}
func (m *MsgWithIndirectRequired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithIndirectRequired.Marshal(b, m, deterministic)
}
func (m *MsgWithIndirectRequired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithIndirectRequired.Merge(m, src)
}
func (m *MsgWithIndirectRequired) XXX_Size() int {
	return xxx_messageInfo_MsgWithIndirectRequired.Size(m)
}
func (m *MsgWithIndirectRequired) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithIndirectRequired.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithIndirectRequired proto.InternalMessageInfo

func (m *MsgWithIndirectRequired) GetSubm() *MsgWithRequired {
	if m != nil {
		return m.Subm
	}
	return nil
}

func (m *MsgWithIndirectRequired) GetMapField() map[string]*MsgWithRequired {
	if m != nil {
		return m.MapField
	}
	return nil
}

func (m *MsgWithIndirectRequired) GetSliceField() []*MsgWithRequired {
	if m != nil {
		return m.SliceField
	}
	return nil
}

type MsgWithRequiredBytes struct {
	Byts                 []byte   `protobuf:"bytes,1,req,name=byts" json:"byts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgWithRequiredBytes) Reset()         { *m = MsgWithRequiredBytes{} }
func (m *MsgWithRequiredBytes) String() string { return proto.CompactTextString(m) }
func (*MsgWithRequiredBytes) ProtoMessage()    {}
func (*MsgWithRequiredBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{11}
}

func (m *MsgWithRequiredBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithRequiredBytes.Unmarshal(m, b)
}
func (m *MsgWithRequiredBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithRequiredBytes.Marshal(b, m, deterministic)
}
func (m *MsgWithRequiredBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithRequiredBytes.Merge(m, src)
}
func (m *MsgWithRequiredBytes) XXX_Size() int {
	return xxx_messageInfo_MsgWithRequiredBytes.Size(m)
}
func (m *MsgWithRequiredBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithRequiredBytes.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithRequiredBytes proto.InternalMessageInfo

func (m *MsgWithRequiredBytes) GetByts() []byte {
	if m != nil {
		return m.Byts
	}
	return nil
}

type MsgWithRequiredWKT struct {
	Str                  *wrapperspb.StringValue `protobuf:"bytes,1,req,name=str" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MsgWithRequiredWKT) Reset()         { *m = MsgWithRequiredWKT{} }
func (m *MsgWithRequiredWKT) String() string { return proto.CompactTextString(m) }
func (*MsgWithRequiredWKT) ProtoMessage()    {}
func (*MsgWithRequiredWKT) Descriptor() ([]byte, []int) {
	return fileDescriptor_50cab1d8463dea41, []int{12}
}

func (m *MsgWithRequiredWKT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithRequiredWKT.Unmarshal(m, b)
}
func (m *MsgWithRequiredWKT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithRequiredWKT.Marshal(b, m, deterministic)
}
func (m *MsgWithRequiredWKT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithRequiredWKT.Merge(m, src)
}
func (m *MsgWithRequiredWKT) XXX_Size() int {
	return xxx_messageInfo_MsgWithRequiredWKT.Size(m)
}
func (m *MsgWithRequiredWKT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithRequiredWKT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithRequiredWKT proto.InternalMessageInfo

func (m *MsgWithRequiredWKT) GetStr() *wrapperspb.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*string)(nil),
	Field:         124,
	Name:          "jsonpb_test.name",
	Tag:           "bytes,124,opt,name=name",
	Filename:      "jsonpb_proto/test2.proto",
}

var E_Extm = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*MsgWithRequired)(nil),
	Field:         125,
	Name:          "jsonpb_test.extm",
	Tag:           "bytes,125,opt,name=extm",
	Filename:      "jsonpb_proto/test2.proto",
}

func init() {
	proto.RegisterEnum("jsonpb_test.Widget_Color", Widget_Color_name, Widget_Color_value)
	proto.RegisterType((*Simple)(nil), "jsonpb_test.Simple")
	proto.RegisterType((*NonFinites)(nil), "jsonpb_test.NonFinites")
	proto.RegisterType((*Repeats)(nil), "jsonpb_test.Repeats")
	proto.RegisterType((*Widget)(nil), "jsonpb_test.Widget")
	proto.RegisterType((*Maps)(nil), "jsonpb_test.Maps")
	proto.RegisterMapType((map[bool]*Simple)(nil), "jsonpb_test.Maps.MBoolSimpleEntry")
	proto.RegisterMapType((map[int64]string)(nil), "jsonpb_test.Maps.MInt64StrEntry")
	proto.RegisterType((*MsgWithOneof)(nil), "jsonpb_test.MsgWithOneof")
	proto.RegisterType((*Real)(nil), "jsonpb_test.Real")
	proto.RegisterExtension(E_Complex_RealExtension)
	proto.RegisterType((*Complex)(nil), "jsonpb_test.Complex")
	proto.RegisterType((*KnownTypes)(nil), "jsonpb_test.KnownTypes")
	proto.RegisterType((*MsgWithRequired)(nil), "jsonpb_test.MsgWithRequired")
	proto.RegisterType((*MsgWithIndirectRequired)(nil), "jsonpb_test.MsgWithIndirectRequired")
	proto.RegisterMapType((map[string]*MsgWithRequired)(nil), "jsonpb_test.MsgWithIndirectRequired.MapFieldEntry")
	proto.RegisterType((*MsgWithRequiredBytes)(nil), "jsonpb_test.MsgWithRequiredBytes")
	proto.RegisterType((*MsgWithRequiredWKT)(nil), "jsonpb_test.MsgWithRequiredWKT")
	proto.RegisterExtension(E_Name)
	proto.RegisterExtension(E_Extm)
}

func init() { proto.RegisterFile("jsonpb_proto/test2.proto", fileDescriptor_50cab1d8463dea41) }

var fileDescriptor_50cab1d8463dea41 = []byte{
	// 1510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xdd, 0x6e, 0xdb, 0x46,
	0x16, 0x36, 0x49, 0x51, 0x12, 0x8f, 0x6c, 0xc7, 0x1e, 0x3b, 0x09, 0xed, 0x0d, 0xb2, 0x84, 0xb2,
	0xd9, 0xd5, 0x26, 0x58, 0x79, 0x43, 0x0b, 0x42, 0x91, 0x36, 0x40, 0xe3, 0xd8, 0x6e, 0xd2, 0x24,
	0x4e, 0x41, 0x27, 0x0d, 0xda, 0x1b, 0x81, 0x32, 0x29, 0x99, 0x2d, 0xc9, 0x51, 0x67, 0x46, 0x4e,
	0x84, 0xb6, 0x80, 0xfb, 0x0a, 0x7d, 0x85, 0x02, 0xbd, 0xed, 0x5d, 0x2f, 0xfa, 0x1c, 0x45, 0x9f,
	0xa7, 0x98, 0x33, 0x43, 0xfd, 0x59, 0x36, 0x72, 0x65, 0xcd, 0x7c, 0x3f, 0x33, 0x9c, 0xf3, 0xf1,
	0x0c, 0x0d, 0xee, 0x37, 0x9c, 0xe6, 0x83, 0x6e, 0x67, 0xc0, 0xa8, 0xa0, 0x3b, 0x22, 0xe6, 0xc2,
	0x6f, 0xe2, 0x6f, 0x52, 0xd3, 0x88, 0x9c, 0xdb, 0xde, 0xea, 0x53, 0xda, 0x4f, 0xe3, 0x1d, 0x84,
	0xba, 0xc3, 0xde, 0x4e, 0x98, 0x8f, 0x14, 0x6f, 0xfb, 0xf6, 0x3c, 0x14, 0x0d, 0x59, 0x28, 0x12,
	0x9a, 0x6b, 0xfc, 0xd6, 0x3c, 0xce, 0x05, 0x1b, 0x9e, 0x08, 0x8d, 0xfe, 0x73, 0x1e, 0x15, 0x49,
	0x16, 0x73, 0x11, 0x66, 0x83, 0xcb, 0xec, 0xdf, 0xb1, 0x70, 0x30, 0x88, 0x19, 0x57, 0x78, 0xfd,
	0xb7, 0x12, 0x94, 0x8f, 0x93, 0x6c, 0x90, 0xc6, 0xe4, 0x3a, 0x94, 0x69, 0xa7, 0x4b, 0x69, 0xea,
	0x1a, 0x9e, 0xd1, 0xa8, 0x06, 0x36, 0xdd, 0xa3, 0x34, 0x25, 0x37, 0xa1, 0x42, 0x3b, 0x49, 0x2e,
	0x76, 0x7d, 0xd7, 0xf4, 0x8c, 0x86, 0x1d, 0x94, 0xe9, 0x33, 0x39, 0x22, 0xb7, 0xa1, 0xa6, 0x81,
	0x0e, 0x17, 0xcc, 0xb5, 0x10, 0x74, 0x14, 0x78, 0x2c, 0xd8, 0x58, 0xd8, 0x6e, 0xb9, 0x25, 0xcf,
	0x68, 0x58, 0x4a, 0xd8, 0x6e, 0x8d, 0x85, 0xed, 0x16, 0x0a, 0x6d, 0x04, 0x1d, 0x05, 0x4a, 0xe1,
	0x16, 0x54, 0x69, 0x67, 0xa8, 0x96, 0x2c, 0x7b, 0x46, 0x63, 0x25, 0xa8, 0xd0, 0x37, 0x38, 0x24,
	0x1e, 0x2c, 0x17, 0x10, 0x6a, 0x2b, 0x08, 0x83, 0x86, 0x67, 0xc4, 0xed, 0x96, 0x5b, 0xf5, 0x8c,
	0x46, 0x49, 0x8b, 0xdb, 0xad, 0x89, 0x58, 0x2f, 0xec, 0x20, 0x0c, 0x1a, 0x1e, 0x8b, 0xb9, 0x5a,
	0x19, 0x3c, 0xa3, 0xb1, 0x1e, 0x54, 0xe8, 0xf1, 0xd4, 0xca, 0x7c, 0xb2, 0x72, 0x0d, 0x61, 0xd0,
	0xf0, 0x8c, 0xb8, 0xdd, 0x72, 0x97, 0x3d, 0xa3, 0x41, 0xb4, 0xb8, 0x58, 0x99, 0x4f, 0x56, 0x5e,
	0x41, 0x18, 0x34, 0x3c, 0x3e, 0xac, 0x5e, 0x4a, 0x43, 0xe1, 0xae, 0x7a, 0x46, 0xc3, 0x0c, 0xca,
	0xf4, 0x50, 0x8e, 0xd4, 0x61, 0x21, 0x80, 0xca, 0x6b, 0x08, 0x3a, 0x0a, 0x1c, 0xaf, 0x1a, 0xd1,
	0x61, 0x37, 0x8d, 0xdd, 0x35, 0xcf, 0x68, 0x18, 0x41, 0x85, 0xee, 0xe3, 0x50, 0xad, 0xaa, 0x20,
	0xd4, 0xae, 0x23, 0x0c, 0x1a, 0x9e, 0x6c, 0x59, 0xb0, 0x24, 0xef, 0xbb, 0xc4, 0x33, 0x1a, 0x8e,
	0xdc, 0x32, 0x0e, 0xd5, 0x86, 0xba, 0x23, 0x11, 0x73, 0x77, 0xc3, 0x33, 0x1a, 0xcb, 0x41, 0x99,
	0xee, 0xc9, 0x51, 0xfd, 0x67, 0x03, 0xe0, 0x88, 0xe6, 0x87, 0x49, 0x9e, 0x88, 0x98, 0x93, 0x0d,
	0xb0, 0x7b, 0x9d, 0x3c, 0xcc, 0x31, 0x34, 0x66, 0x50, 0xea, 0x1d, 0x85, 0xb9, 0x8c, 0x52, 0xaf,
	0x33, 0x48, 0xf2, 0x1e, 0x46, 0xc6, 0x0c, 0xec, 0xde, 0x17, 0x49, 0xde, 0x53, 0xd3, 0xb9, 0x9c,
	0xb6, 0xf4, 0xf4, 0x91, 0x9c, 0xde, 0x00, 0x3b, 0x42, 0x8b, 0x12, 0x6e, 0xb0, 0x14, 0x69, 0x8b,
	0x48, 0x59, 0xd8, 0x38, 0x6b, 0x47, 0x85, 0x45, 0xa4, 0x2c, 0xca, 0x7a, 0x5a, 0x5a, 0xd4, 0x7f,
	0x35, 0xa1, 0x12, 0xc4, 0x83, 0x38, 0x14, 0x5c, 0x52, 0x58, 0x91, 0x63, 0x4b, 0xe6, 0x98, 0x15,
	0x39, 0x66, 0xe3, 0x1c, 0x5b, 0x32, 0xc7, 0x4c, 0xe5, 0xb8, 0x00, 0xda, 0x2d, 0xd7, 0xf2, 0x2c,
	0x99, 0x53, 0xa6, 0x72, 0xba, 0x05, 0x55, 0x56, 0xe4, 0xb0, 0xe4, 0x59, 0x32, 0x87, 0x4c, 0xe7,
	0x70, 0x0c, 0xb5, 0x5b, 0xae, 0xed, 0x59, 0x32, 0x65, 0x4c, 0xa7, 0x0c, 0x21, 0x5e, 0xa4, 0xd7,
	0x92, 0x19, 0x62, 0xc7, 0x53, 0x2a, 0x9d, 0x90, 0x8a, 0x67, 0xc9, 0x84, 0x30, 0x9d, 0x10, 0xdc,
	0x84, 0xaa, 0x7f, 0xd5, 0xb3, 0x64, 0xfd, 0x99, 0xaa, 0x3f, 0x6a, 0x74, 0x7d, 0x1d, 0xcf, 0x92,
	0xf5, 0x65, 0xba, 0xbe, 0xca, 0x4e, 0x55, 0x0f, 0x3c, 0x4b, 0x56, 0x8f, 0x4d, 0xaa, 0xc7, 0x74,
	0xf5, 0x6a, 0x9e, 0x25, 0xab, 0xc7, 0x54, 0xf5, 0xfe, 0x34, 0xa1, 0xfc, 0x36, 0x89, 0xfa, 0xb1,
	0x20, 0x3b, 0x60, 0x9f, 0xd0, 0x94, 0x32, 0xac, 0xdc, 0xaa, 0xbf, 0xd5, 0x9c, 0xea, 0x58, 0x4d,
	0xc5, 0x69, 0x3e, 0x91, 0x84, 0x40, 0xf1, 0x88, 0x2f, 0x4d, 0x95, 0x44, 0x9e, 0xe0, 0x95, 0x92,
	0x32, 0xc3, 0xbf, 0xe4, 0x3e, 0x94, 0x39, 0xb6, 0x17, 0x7c, 0x9f, 0x6a, 0xfe, 0xc6, 0x8c, 0x44,
	0x75, 0x9e, 0x40, 0x53, 0x48, 0x53, 0x9d, 0x0f, 0xd2, 0xe5, 0xb6, 0x2f, 0xa1, 0xcb, 0x43, 0xd3,
	0xfc, 0x0a, 0x53, 0x45, 0x77, 0x37, 0xd1, 0x7d, 0x73, 0x86, 0xae, 0x03, 0x11, 0x14, 0x24, 0xf2,
	0x00, 0x1c, 0xd6, 0x29, 0x14, 0xd7, 0x71, 0x81, 0xc5, 0x8a, 0x2a, 0xd3, 0xbf, 0xea, 0x77, 0xc1,
	0x56, 0x0f, 0x52, 0x01, 0x2b, 0x38, 0xd8, 0x5f, 0x5b, 0x22, 0x0e, 0xd8, 0x9f, 0x05, 0x07, 0x07,
	0x47, 0x6b, 0x06, 0xa9, 0x42, 0x69, 0xef, 0xc5, 0x9b, 0x83, 0x35, 0xb3, 0xfe, 0x8b, 0x09, 0xa5,
	0x97, 0xe1, 0x80, 0x93, 0x4f, 0xa1, 0x96, 0x4d, 0xf5, 0x36, 0x03, 0x17, 0xf1, 0x66, 0x16, 0x91,
	0xbc, 0xe6, 0xcb, 0xa2, 0xdb, 0x1d, 0xe4, 0x82, 0x8d, 0x02, 0x27, 0x1b, 0x77, 0xbf, 0x43, 0x58,
	0xc9, 0x30, 0xbe, 0xc5, 0x49, 0x98, 0xe8, 0x51, 0x5f, 0xe0, 0x21, 0x73, 0xad, 0x8e, 0x42, 0xb9,
	0xd4, 0xb2, 0xc9, 0xcc, 0xf6, 0x27, 0xb0, 0x3a, 0xbb, 0x08, 0x59, 0x03, 0xeb, 0xdb, 0x78, 0x84,
	0xe5, 0xb6, 0x02, 0xf9, 0x93, 0x6c, 0x82, 0x7d, 0x16, 0xa6, 0xc3, 0x18, 0x5f, 0x53, 0x27, 0x50,
	0x83, 0x87, 0xe6, 0x47, 0xc6, 0xf6, 0x31, 0xac, 0xcd, 0xdb, 0x4f, 0xeb, 0xab, 0x4a, 0xff, 0xdf,
	0x69, 0xfd, 0x25, 0xd5, 0x9a, 0x98, 0xd6, 0xff, 0x32, 0x60, 0xf9, 0x25, 0xef, 0xbf, 0x4d, 0xc4,
	0xe9, 0xab, 0x3c, 0xa6, 0x3d, 0x72, 0x03, 0x6c, 0x91, 0x88, 0x34, 0x46, 0x4f, 0xe7, 0xe9, 0x52,
	0xa0, 0x86, 0xc4, 0x85, 0x32, 0x0f, 0xd3, 0x90, 0x8d, 0xd0, 0xd8, 0x7a, 0xba, 0x14, 0xe8, 0x31,
	0xd9, 0x86, 0xca, 0x13, 0x3a, 0x94, 0xdb, 0xc1, 0x1e, 0x22, 0x35, 0xc5, 0x04, 0xb9, 0x03, 0xcb,
	0xa7, 0x34, 0x8b, 0x3b, 0x61, 0x14, 0xb1, 0x98, 0x73, 0x6c, 0x27, 0x92, 0x50, 0x93, 0xb3, 0x8f,
	0xd5, 0x24, 0xf9, 0x1c, 0xd6, 0x33, 0xde, 0xef, 0xbc, 0x4b, 0xc4, 0x69, 0x87, 0xc5, 0xdf, 0x0d,
	0x13, 0x16, 0x47, 0xd8, 0x62, 0x6a, 0xfe, 0xad, 0xd9, 0x23, 0x56, 0x1b, 0x0d, 0x34, 0xe7, 0xe9,
	0x52, 0x70, 0x2d, 0x9b, 0x9d, 0xda, 0xab, 0x80, 0x3d, 0xcc, 0x13, 0x9a, 0xd7, 0xff, 0x0d, 0xa5,
	0x20, 0x0e, 0xd3, 0xc9, 0x79, 0x1a, 0xaa, 0x39, 0xe1, 0xe0, 0x5e, 0xb5, 0x1a, 0xad, 0x9d, 0x9f,
	0x9f, 0x9f, 0x9b, 0xf5, 0x9f, 0x0c, 0xb9, 0x7d, 0x79, 0x2c, 0xef, 0xc9, 0x2d, 0x70, 0x92, 0x2c,
	0xec, 0x27, 0xb9, 0x7c, 0x4c, 0xc5, 0x9f, 0x4c, 0x4c, 0x34, 0xfe, 0x11, 0xac, 0xb2, 0x38, 0x4c,
	0x3b, 0xf1, 0x7b, 0x11, 0xe7, 0x3c, 0xa1, 0x39, 0x59, 0x9f, 0xcb, 0x6c, 0x98, 0xba, 0xdf, 0x2f,
	0x88, 0xbf, 0x5e, 0x28, 0x58, 0x91, 0xf2, 0x83, 0x42, 0x5d, 0xff, 0xc3, 0x06, 0x78, 0x9e, 0xd3,
	0x77, 0xf9, 0xeb, 0xd1, 0x20, 0xe6, 0xe4, 0x5f, 0x60, 0x86, 0x39, 0xde, 0x39, 0x52, 0xaf, 0xbe,
	0x16, 0x9a, 0xc5, 0xd7, 0x42, 0xf3, 0x71, 0x3e, 0x0a, 0xcc, 0x30, 0x27, 0xf7, 0xc1, 0x8a, 0x86,
	0xaa, 0x53, 0xd4, 0xfc, 0xad, 0x0b, 0xb4, 0x7d, 0xfd, 0xcd, 0x12, 0x48, 0x16, 0xf9, 0x0f, 0x98,
	0x5c, 0xe0, 0x15, 0x58, 0xf3, 0x6f, 0x5e, 0xe0, 0x1e, 0xe3, 0xf7, 0x4b, 0x60, 0x72, 0x41, 0xee,
	0x81, 0x29, 0xb8, 0xce, 0xce, 0xf6, 0x05, 0xe2, 0xeb, 0xe2, 0x53, 0x26, 0x30, 0x05, 0x97, 0xdc,
	0xf4, 0x0c, 0xaf, 0xbf, 0x45, 0xdc, 0x17, 0x09, 0x17, 0x5f, 0xca, 0xc3, 0x0e, 0xcc, 0xf4, 0x8c,
	0x34, 0xc0, 0x3a, 0x0b, 0x53, 0xbc, 0x0e, 0x6b, 0xfe, 0x8d, 0x0b, 0x64, 0x45, 0x94, 0x14, 0xd2,
	0x04, 0x2b, 0xea, 0xa6, 0x18, 0x25, 0x59, 0xff, 0x0b, 0xcf, 0x85, 0x8d, 0x56, 0xf3, 0xa3, 0x6e,
	0x4a, 0xfe, 0x07, 0x56, 0x2f, 0x15, 0x98, 0xac, 0x9a, 0xff, 0x8f, 0x0b, 0x7c, 0x6c, 0xd9, 0x9a,
	0xde, 0x4b, 0x85, 0xa4, 0x27, 0x78, 0x43, 0x2c, 0xa6, 0xe3, 0xeb, 0xa9, 0xe9, 0x49, 0xbb, 0x25,
	0x77, 0x33, 0x6c, 0xb7, 0xf0, 0x66, 0x5b, 0xb4, 0x9b, 0x37, 0xd3, 0xfc, 0x61, 0xbb, 0x85, 0xf6,
	0xbb, 0x3e, 0x7e, 0x04, 0x5d, 0x62, 0xbf, 0xeb, 0x17, 0xf6, 0xbb, 0x3e, 0xda, 0xef, 0xfa, 0xf8,
	0x55, 0x74, 0x99, 0xfd, 0x98, 0x3f, 0x44, 0x7e, 0x09, 0xaf, 0x51, 0xe7, 0x92, 0x43, 0x97, 0xfd,
	0x41, 0xd1, 0x91, 0x27, 0xfd, 0x65, 0xcf, 0x83, 0x4b, 0xfc, 0xd5, 0xd5, 0xa4, 0xfd, 0xb9, 0x60,
	0xe4, 0x01, 0xd8, 0xc5, 0x15, 0xb5, 0xf8, 0x01, 0xf0, 0xca, 0x52, 0x02, 0xc5, 0xac, 0xdf, 0x81,
	0x6b, 0x73, 0xef, 0xa5, 0xec, 0x4a, 0xaa, 0xd3, 0x9a, 0x0d, 0x07, 0x7d, 0xeb, 0xbf, 0x9b, 0x70,
	0x53, 0xb3, 0x9e, 0xe5, 0x51, 0xc2, 0xe2, 0x13, 0x31, 0x66, 0xff, 0x1f, 0x4a, 0x7c, 0xd8, 0xcd,
	0x74, 0x92, 0xaf, 0x7c, 0xe3, 0x03, 0x64, 0x92, 0x57, 0xe0, 0x64, 0xe1, 0xa0, 0xd3, 0x4b, 0xe2,
	0x34, 0xd2, 0xbd, 0xd8, 0x5f, 0x24, 0x9b, 0x5f, 0x4a, 0xf6, 0xe8, 0x43, 0x29, 0x52, 0xbd, 0xb9,
	0x9a, 0xe9, 0x21, 0x79, 0x04, 0x35, 0x9e, 0x26, 0x27, 0xb1, 0xb6, 0xb4, 0xd0, 0xf2, 0xea, 0x9d,
	0x00, 0x0a, 0x50, 0xbe, 0xfd, 0x15, 0xac, 0xcc, 0x38, 0x4f, 0xb7, 0x65, 0x47, 0xb5, 0x65, 0x7f,
	0xb6, 0x2d, 0x5f, 0xed, 0x3d, 0xd5, 0x9f, 0xef, 0xc1, 0xe6, 0x1c, 0x8a, 0x15, 0x20, 0x04, 0x4a,
	0xdd, 0x91, 0xe0, 0x78, 0xc6, 0xcb, 0x01, 0xfe, 0xae, 0xef, 0x03, 0x99, 0xe3, 0xbe, 0x7d, 0xfe,
	0xba, 0x88, 0x80, 0x24, 0x7e, 0x48, 0x04, 0x1e, 0xde, 0x85, 0x52, 0x1e, 0x66, 0xf1, 0xa2, 0x96,
	0xf6, 0x03, 0x3e, 0x0f, 0xc2, 0x0f, 0x9f, 0x40, 0x29, 0x7e, 0x2f, 0xb2, 0x45, 0xb4, 0x1f, 0x3f,
	0xa4, 0x90, 0x52, 0xbc, 0xf7, 0xe8, 0xeb, 0x8f, 0xfb, 0x89, 0x38, 0x1d, 0x76, 0x9b, 0x27, 0x34,
	0xdb, 0xe9, 0xd3, 0x34, 0xcc, 0xfb, 0x93, 0xff, 0x8b, 0x92, 0x5c, 0xc4, 0x2c, 0x0f, 0x53, 0xfc,
	0x27, 0x0e, 0x67, 0xf9, 0xce, 0xf4, 0x3f, 0x77, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6b,
	0x0b, 0xa0, 0xeb, 0x0d, 0x00, 0x00,
}
