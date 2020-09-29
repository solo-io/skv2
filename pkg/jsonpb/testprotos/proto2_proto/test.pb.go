package proto2_proto

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

type FOO int32

const (
	FOO_FOO1 FOO = 1
)

var FOO_name = map[int32]string{
	1: "FOO1",
}

var FOO_value = map[string]int32{
	"FOO1": 1,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}

func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}

func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{0}
}

// An enum, for completeness.
type GoTest_KIND int32

const (
	GoTest_VOID GoTest_KIND = 0
	// Basic types
	GoTest_BOOL        GoTest_KIND = 1
	GoTest_BYTES       GoTest_KIND = 2
	GoTest_FINGERPRINT GoTest_KIND = 3
	GoTest_FLOAT       GoTest_KIND = 4
	GoTest_INT         GoTest_KIND = 5
	GoTest_STRING      GoTest_KIND = 6
	GoTest_TIME        GoTest_KIND = 7
	// Groupings
	GoTest_TUPLE GoTest_KIND = 8
	GoTest_ARRAY GoTest_KIND = 9
	GoTest_MAP   GoTest_KIND = 10
	// Table types
	GoTest_TABLE GoTest_KIND = 11
	// Functions
	GoTest_FUNCTION GoTest_KIND = 12
)

var GoTest_KIND_name = map[int32]string{
	0:  "VOID",
	1:  "BOOL",
	2:  "BYTES",
	3:  "FINGERPRINT",
	4:  "FLOAT",
	5:  "INT",
	6:  "STRING",
	7:  "TIME",
	8:  "TUPLE",
	9:  "ARRAY",
	10: "MAP",
	11: "TABLE",
	12: "FUNCTION",
}

var GoTest_KIND_value = map[string]int32{
	"VOID":        0,
	"BOOL":        1,
	"BYTES":       2,
	"FINGERPRINT": 3,
	"FLOAT":       4,
	"INT":         5,
	"STRING":      6,
	"TIME":        7,
	"TUPLE":       8,
	"ARRAY":       9,
	"MAP":         10,
	"TABLE":       11,
	"FUNCTION":    12,
}

func (x GoTest_KIND) Enum() *GoTest_KIND {
	p := new(GoTest_KIND)
	*p = x
	return p
}

func (x GoTest_KIND) String() string {
	return proto.EnumName(GoTest_KIND_name, int32(x))
}

func (x *GoTest_KIND) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GoTest_KIND_value, data, "GoTest_KIND")
	if err != nil {
		return err
	}
	*x = GoTest_KIND(value)
	return nil
}

func (GoTest_KIND) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{2, 0}
}

type MyMessage_Color int32

const (
	MyMessage_RED   MyMessage_Color = 0
	MyMessage_GREEN MyMessage_Color = 1
	MyMessage_BLUE  MyMessage_Color = 2
)

var MyMessage_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

var MyMessage_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x MyMessage_Color) Enum() *MyMessage_Color {
	p := new(MyMessage_Color)
	*p = x
	return p
}

func (x MyMessage_Color) String() string {
	return proto.EnumName(MyMessage_Color_name, int32(x))
}

func (x *MyMessage_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyMessage_Color_value, data, "MyMessage_Color")
	if err != nil {
		return err
	}
	*x = MyMessage_Color(value)
	return nil
}

func (MyMessage_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{13, 0}
}

type DefaultsMessage_DefaultsEnum int32

const (
	DefaultsMessage_ZERO DefaultsMessage_DefaultsEnum = 0
	DefaultsMessage_ONE  DefaultsMessage_DefaultsEnum = 1
	DefaultsMessage_TWO  DefaultsMessage_DefaultsEnum = 2
)

var DefaultsMessage_DefaultsEnum_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
	2: "TWO",
}

var DefaultsMessage_DefaultsEnum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
	"TWO":  2,
}

func (x DefaultsMessage_DefaultsEnum) Enum() *DefaultsMessage_DefaultsEnum {
	p := new(DefaultsMessage_DefaultsEnum)
	*p = x
	return p
}

func (x DefaultsMessage_DefaultsEnum) String() string {
	return proto.EnumName(DefaultsMessage_DefaultsEnum_name, int32(x))
}

func (x *DefaultsMessage_DefaultsEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DefaultsMessage_DefaultsEnum_value, data, "DefaultsMessage_DefaultsEnum")
	if err != nil {
		return err
	}
	*x = DefaultsMessage_DefaultsEnum(value)
	return nil
}

func (DefaultsMessage_DefaultsEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{16, 0}
}

type Defaults_Color int32

const (
	Defaults_RED   Defaults_Color = 0
	Defaults_GREEN Defaults_Color = 1
	Defaults_BLUE  Defaults_Color = 2
)

var Defaults_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

var Defaults_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Defaults_Color) Enum() *Defaults_Color {
	p := new(Defaults_Color)
	*p = x
	return p
}

func (x Defaults_Color) String() string {
	return proto.EnumName(Defaults_Color_name, int32(x))
}

func (x *Defaults_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Defaults_Color_value, data, "Defaults_Color")
	if err != nil {
		return err
	}
	*x = Defaults_Color(value)
	return nil
}

func (Defaults_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{20, 0}
}

type RepeatedEnum_Color int32

const (
	RepeatedEnum_RED RepeatedEnum_Color = 1
)

var RepeatedEnum_Color_name = map[int32]string{
	1: "RED",
}

var RepeatedEnum_Color_value = map[string]int32{
	"RED": 1,
}

func (x RepeatedEnum_Color) Enum() *RepeatedEnum_Color {
	p := new(RepeatedEnum_Color)
	*p = x
	return p
}

func (x RepeatedEnum_Color) String() string {
	return proto.EnumName(RepeatedEnum_Color_name, int32(x))
}

func (x *RepeatedEnum_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RepeatedEnum_Color_value, data, "RepeatedEnum_Color")
	if err != nil {
		return err
	}
	*x = RepeatedEnum_Color(value)
	return nil
}

func (RepeatedEnum_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{22, 0}
}

type GoEnum struct {
	Foo                  *FOO     `protobuf:"varint,1,req,name=foo,enum=proto2_test.FOO" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoEnum) Reset()         { *m = GoEnum{} }
func (m *GoEnum) String() string { return proto.CompactTextString(m) }
func (*GoEnum) ProtoMessage()    {}
func (*GoEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{0}
}

func (m *GoEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoEnum.Unmarshal(m, b)
}
func (m *GoEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoEnum.Marshal(b, m, deterministic)
}
func (m *GoEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoEnum.Merge(m, src)
}
func (m *GoEnum) XXX_Size() int {
	return xxx_messageInfo_GoEnum.Size(m)
}
func (m *GoEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GoEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GoEnum proto.InternalMessageInfo

func (m *GoEnum) GetFoo() FOO {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return FOO_FOO1
}

type GoTestField struct {
	Label                *string  `protobuf:"bytes,1,req,name=Label" json:"Label,omitempty"`
	Type                 *string  `protobuf:"bytes,2,req,name=Type" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoTestField) Reset()         { *m = GoTestField{} }
func (m *GoTestField) String() string { return proto.CompactTextString(m) }
func (*GoTestField) ProtoMessage()    {}
func (*GoTestField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{1}
}

func (m *GoTestField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTestField.Unmarshal(m, b)
}
func (m *GoTestField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTestField.Marshal(b, m, deterministic)
}
func (m *GoTestField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTestField.Merge(m, src)
}
func (m *GoTestField) XXX_Size() int {
	return xxx_messageInfo_GoTestField.Size(m)
}
func (m *GoTestField) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTestField.DiscardUnknown(m)
}

var xxx_messageInfo_GoTestField proto.InternalMessageInfo

func (m *GoTestField) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *GoTestField) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

type GoTest struct {
	// Some typical parameters
	Kind  *GoTest_KIND `protobuf:"varint,1,req,name=Kind,enum=proto2_test.GoTest_KIND" json:"Kind,omitempty"`
	Table *string      `protobuf:"bytes,2,opt,name=Table" json:"Table,omitempty"`
	Param *int32       `protobuf:"varint,3,opt,name=Param" json:"Param,omitempty"`
	// Required, repeated and optional foreign fields.
	RequiredField *GoTestField   `protobuf:"bytes,4,req,name=RequiredField" json:"RequiredField,omitempty"`
	RepeatedField []*GoTestField `protobuf:"bytes,5,rep,name=RepeatedField" json:"RepeatedField,omitempty"`
	OptionalField *GoTestField   `protobuf:"bytes,6,opt,name=OptionalField" json:"OptionalField,omitempty"`
	// Required fields of all basic types
	F_BoolRequired     *bool    `protobuf:"varint,10,req,name=F_Bool_required,json=FBoolRequired" json:"F_Bool_required,omitempty"`
	F_Int32Required    *int32   `protobuf:"varint,11,req,name=F_Int32_required,json=FInt32Required" json:"F_Int32_required,omitempty"`
	F_Int64Required    *int64   `protobuf:"varint,12,req,name=F_Int64_required,json=FInt64Required" json:"F_Int64_required,omitempty"`
	F_Fixed32Required  *uint32  `protobuf:"fixed32,13,req,name=F_Fixed32_required,json=FFixed32Required" json:"F_Fixed32_required,omitempty"`
	F_Fixed64Required  *uint64  `protobuf:"fixed64,14,req,name=F_Fixed64_required,json=FFixed64Required" json:"F_Fixed64_required,omitempty"`
	F_Uint32Required   *uint32  `protobuf:"varint,15,req,name=F_Uint32_required,json=FUint32Required" json:"F_Uint32_required,omitempty"`
	F_Uint64Required   *uint64  `protobuf:"varint,16,req,name=F_Uint64_required,json=FUint64Required" json:"F_Uint64_required,omitempty"`
	F_FloatRequired    *float32 `protobuf:"fixed32,17,req,name=F_Float_required,json=FFloatRequired" json:"F_Float_required,omitempty"`
	F_DoubleRequired   *float64 `protobuf:"fixed64,18,req,name=F_Double_required,json=FDoubleRequired" json:"F_Double_required,omitempty"`
	F_StringRequired   *string  `protobuf:"bytes,19,req,name=F_String_required,json=FStringRequired" json:"F_String_required,omitempty"`
	F_BytesRequired    []byte   `protobuf:"bytes,101,req,name=F_Bytes_required,json=FBytesRequired" json:"F_Bytes_required,omitempty"`
	F_Sint32Required   *int32   `protobuf:"zigzag32,102,req,name=F_Sint32_required,json=FSint32Required" json:"F_Sint32_required,omitempty"`
	F_Sint64Required   *int64   `protobuf:"zigzag64,103,req,name=F_Sint64_required,json=FSint64Required" json:"F_Sint64_required,omitempty"`
	F_Sfixed32Required *int32   `protobuf:"fixed32,104,req,name=F_Sfixed32_required,json=FSfixed32Required" json:"F_Sfixed32_required,omitempty"`
	F_Sfixed64Required *int64   `protobuf:"fixed64,105,req,name=F_Sfixed64_required,json=FSfixed64Required" json:"F_Sfixed64_required,omitempty"`
	// Repeated fields of all basic types
	F_BoolRepeated     []bool    `protobuf:"varint,20,rep,name=F_Bool_repeated,json=FBoolRepeated" json:"F_Bool_repeated,omitempty"`
	F_Int32Repeated    []int32   `protobuf:"varint,21,rep,name=F_Int32_repeated,json=FInt32Repeated" json:"F_Int32_repeated,omitempty"`
	F_Int64Repeated    []int64   `protobuf:"varint,22,rep,name=F_Int64_repeated,json=FInt64Repeated" json:"F_Int64_repeated,omitempty"`
	F_Fixed32Repeated  []uint32  `protobuf:"fixed32,23,rep,name=F_Fixed32_repeated,json=FFixed32Repeated" json:"F_Fixed32_repeated,omitempty"`
	F_Fixed64Repeated  []uint64  `protobuf:"fixed64,24,rep,name=F_Fixed64_repeated,json=FFixed64Repeated" json:"F_Fixed64_repeated,omitempty"`
	F_Uint32Repeated   []uint32  `protobuf:"varint,25,rep,name=F_Uint32_repeated,json=FUint32Repeated" json:"F_Uint32_repeated,omitempty"`
	F_Uint64Repeated   []uint64  `protobuf:"varint,26,rep,name=F_Uint64_repeated,json=FUint64Repeated" json:"F_Uint64_repeated,omitempty"`
	F_FloatRepeated    []float32 `protobuf:"fixed32,27,rep,name=F_Float_repeated,json=FFloatRepeated" json:"F_Float_repeated,omitempty"`
	F_DoubleRepeated   []float64 `protobuf:"fixed64,28,rep,name=F_Double_repeated,json=FDoubleRepeated" json:"F_Double_repeated,omitempty"`
	F_StringRepeated   []string  `protobuf:"bytes,29,rep,name=F_String_repeated,json=FStringRepeated" json:"F_String_repeated,omitempty"`
	F_BytesRepeated    [][]byte  `protobuf:"bytes,201,rep,name=F_Bytes_repeated,json=FBytesRepeated" json:"F_Bytes_repeated,omitempty"`
	F_Sint32Repeated   []int32   `protobuf:"zigzag32,202,rep,name=F_Sint32_repeated,json=FSint32Repeated" json:"F_Sint32_repeated,omitempty"`
	F_Sint64Repeated   []int64   `protobuf:"zigzag64,203,rep,name=F_Sint64_repeated,json=FSint64Repeated" json:"F_Sint64_repeated,omitempty"`
	F_Sfixed32Repeated []int32   `protobuf:"fixed32,204,rep,name=F_Sfixed32_repeated,json=FSfixed32Repeated" json:"F_Sfixed32_repeated,omitempty"`
	F_Sfixed64Repeated []int64   `protobuf:"fixed64,205,rep,name=F_Sfixed64_repeated,json=FSfixed64Repeated" json:"F_Sfixed64_repeated,omitempty"`
	// Optional fields of all basic types
	F_BoolOptional     *bool    `protobuf:"varint,30,opt,name=F_Bool_optional,json=FBoolOptional" json:"F_Bool_optional,omitempty"`
	F_Int32Optional    *int32   `protobuf:"varint,31,opt,name=F_Int32_optional,json=FInt32Optional" json:"F_Int32_optional,omitempty"`
	F_Int64Optional    *int64   `protobuf:"varint,32,opt,name=F_Int64_optional,json=FInt64Optional" json:"F_Int64_optional,omitempty"`
	F_Fixed32Optional  *uint32  `protobuf:"fixed32,33,opt,name=F_Fixed32_optional,json=FFixed32Optional" json:"F_Fixed32_optional,omitempty"`
	F_Fixed64Optional  *uint64  `protobuf:"fixed64,34,opt,name=F_Fixed64_optional,json=FFixed64Optional" json:"F_Fixed64_optional,omitempty"`
	F_Uint32Optional   *uint32  `protobuf:"varint,35,opt,name=F_Uint32_optional,json=FUint32Optional" json:"F_Uint32_optional,omitempty"`
	F_Uint64Optional   *uint64  `protobuf:"varint,36,opt,name=F_Uint64_optional,json=FUint64Optional" json:"F_Uint64_optional,omitempty"`
	F_FloatOptional    *float32 `protobuf:"fixed32,37,opt,name=F_Float_optional,json=FFloatOptional" json:"F_Float_optional,omitempty"`
	F_DoubleOptional   *float64 `protobuf:"fixed64,38,opt,name=F_Double_optional,json=FDoubleOptional" json:"F_Double_optional,omitempty"`
	F_StringOptional   *string  `protobuf:"bytes,39,opt,name=F_String_optional,json=FStringOptional" json:"F_String_optional,omitempty"`
	F_BytesOptional    []byte   `protobuf:"bytes,301,opt,name=F_Bytes_optional,json=FBytesOptional" json:"F_Bytes_optional,omitempty"`
	F_Sint32Optional   *int32   `protobuf:"zigzag32,302,opt,name=F_Sint32_optional,json=FSint32Optional" json:"F_Sint32_optional,omitempty"`
	F_Sint64Optional   *int64   `protobuf:"zigzag64,303,opt,name=F_Sint64_optional,json=FSint64Optional" json:"F_Sint64_optional,omitempty"`
	F_Sfixed32Optional *int32   `protobuf:"fixed32,304,opt,name=F_Sfixed32_optional,json=FSfixed32Optional" json:"F_Sfixed32_optional,omitempty"`
	F_Sfixed64Optional *int64   `protobuf:"fixed64,305,opt,name=F_Sfixed64_optional,json=FSfixed64Optional" json:"F_Sfixed64_optional,omitempty"`
	// Default-valued fields of all basic types
	F_BoolDefaulted     *bool    `protobuf:"varint,40,opt,name=F_Bool_defaulted,json=FBoolDefaulted,def=1" json:"F_Bool_defaulted,omitempty"`
	F_Int32Defaulted    *int32   `protobuf:"varint,41,opt,name=F_Int32_defaulted,json=FInt32Defaulted,def=32" json:"F_Int32_defaulted,omitempty"`
	F_Int64Defaulted    *int64   `protobuf:"varint,42,opt,name=F_Int64_defaulted,json=FInt64Defaulted,def=64" json:"F_Int64_defaulted,omitempty"`
	F_Fixed32Defaulted  *uint32  `protobuf:"fixed32,43,opt,name=F_Fixed32_defaulted,json=FFixed32Defaulted,def=320" json:"F_Fixed32_defaulted,omitempty"`
	F_Fixed64Defaulted  *uint64  `protobuf:"fixed64,44,opt,name=F_Fixed64_defaulted,json=FFixed64Defaulted,def=640" json:"F_Fixed64_defaulted,omitempty"`
	F_Uint32Defaulted   *uint32  `protobuf:"varint,45,opt,name=F_Uint32_defaulted,json=FUint32Defaulted,def=3200" json:"F_Uint32_defaulted,omitempty"`
	F_Uint64Defaulted   *uint64  `protobuf:"varint,46,opt,name=F_Uint64_defaulted,json=FUint64Defaulted,def=6400" json:"F_Uint64_defaulted,omitempty"`
	F_FloatDefaulted    *float32 `protobuf:"fixed32,47,opt,name=F_Float_defaulted,json=FFloatDefaulted,def=314159" json:"F_Float_defaulted,omitempty"`
	F_DoubleDefaulted   *float64 `protobuf:"fixed64,48,opt,name=F_Double_defaulted,json=FDoubleDefaulted,def=271828" json:"F_Double_defaulted,omitempty"`
	F_StringDefaulted   *string  `protobuf:"bytes,49,opt,name=F_String_defaulted,json=FStringDefaulted,def=hello, \"world!\"\n" json:"F_String_defaulted,omitempty"`
	F_BytesDefaulted    []byte   `protobuf:"bytes,401,opt,name=F_Bytes_defaulted,json=FBytesDefaulted,def=Bignose" json:"F_Bytes_defaulted,omitempty"`
	F_Sint32Defaulted   *int32   `protobuf:"zigzag32,402,opt,name=F_Sint32_defaulted,json=FSint32Defaulted,def=-32" json:"F_Sint32_defaulted,omitempty"`
	F_Sint64Defaulted   *int64   `protobuf:"zigzag64,403,opt,name=F_Sint64_defaulted,json=FSint64Defaulted,def=-64" json:"F_Sint64_defaulted,omitempty"`
	F_Sfixed32Defaulted *int32   `protobuf:"fixed32,404,opt,name=F_Sfixed32_defaulted,json=FSfixed32Defaulted,def=-32" json:"F_Sfixed32_defaulted,omitempty"`
	F_Sfixed64Defaulted *int64   `protobuf:"fixed64,405,opt,name=F_Sfixed64_defaulted,json=FSfixed64Defaulted,def=-64" json:"F_Sfixed64_defaulted,omitempty"`
	// Packed repeated fields (no string or bytes).
	F_BoolRepeatedPacked     []bool                  `protobuf:"varint,50,rep,packed,name=F_Bool_repeated_packed,json=FBoolRepeatedPacked" json:"F_Bool_repeated_packed,omitempty"`
	F_Int32RepeatedPacked    []int32                 `protobuf:"varint,51,rep,packed,name=F_Int32_repeated_packed,json=FInt32RepeatedPacked" json:"F_Int32_repeated_packed,omitempty"`
	F_Int64RepeatedPacked    []int64                 `protobuf:"varint,52,rep,packed,name=F_Int64_repeated_packed,json=FInt64RepeatedPacked" json:"F_Int64_repeated_packed,omitempty"`
	F_Fixed32RepeatedPacked  []uint32                `protobuf:"fixed32,53,rep,packed,name=F_Fixed32_repeated_packed,json=FFixed32RepeatedPacked" json:"F_Fixed32_repeated_packed,omitempty"`
	F_Fixed64RepeatedPacked  []uint64                `protobuf:"fixed64,54,rep,packed,name=F_Fixed64_repeated_packed,json=FFixed64RepeatedPacked" json:"F_Fixed64_repeated_packed,omitempty"`
	F_Uint32RepeatedPacked   []uint32                `protobuf:"varint,55,rep,packed,name=F_Uint32_repeated_packed,json=FUint32RepeatedPacked" json:"F_Uint32_repeated_packed,omitempty"`
	F_Uint64RepeatedPacked   []uint64                `protobuf:"varint,56,rep,packed,name=F_Uint64_repeated_packed,json=FUint64RepeatedPacked" json:"F_Uint64_repeated_packed,omitempty"`
	F_FloatRepeatedPacked    []float32               `protobuf:"fixed32,57,rep,packed,name=F_Float_repeated_packed,json=FFloatRepeatedPacked" json:"F_Float_repeated_packed,omitempty"`
	F_DoubleRepeatedPacked   []float64               `protobuf:"fixed64,58,rep,packed,name=F_Double_repeated_packed,json=FDoubleRepeatedPacked" json:"F_Double_repeated_packed,omitempty"`
	F_Sint32RepeatedPacked   []int32                 `protobuf:"zigzag32,502,rep,packed,name=F_Sint32_repeated_packed,json=FSint32RepeatedPacked" json:"F_Sint32_repeated_packed,omitempty"`
	F_Sint64RepeatedPacked   []int64                 `protobuf:"zigzag64,503,rep,packed,name=F_Sint64_repeated_packed,json=FSint64RepeatedPacked" json:"F_Sint64_repeated_packed,omitempty"`
	F_Sfixed32RepeatedPacked []int32                 `protobuf:"fixed32,504,rep,packed,name=F_Sfixed32_repeated_packed,json=FSfixed32RepeatedPacked" json:"F_Sfixed32_repeated_packed,omitempty"`
	F_Sfixed64RepeatedPacked []int64                 `protobuf:"fixed64,505,rep,packed,name=F_Sfixed64_repeated_packed,json=FSfixed64RepeatedPacked" json:"F_Sfixed64_repeated_packed,omitempty"`
	Requiredgroup            *GoTest_RequiredGroup   `protobuf:"group,70,req,name=RequiredGroup,json=requiredgroup" json:"requiredgroup,omitempty"`
	Repeatedgroup            []*GoTest_RepeatedGroup `protobuf:"group,80,rep,name=RepeatedGroup,json=repeatedgroup" json:"repeatedgroup,omitempty"`
	Optionalgroup            *GoTest_OptionalGroup   `protobuf:"group,90,opt,name=OptionalGroup,json=optionalgroup" json:"optionalgroup,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                `json:"-"`
	XXX_unrecognized         []byte                  `json:"-"`
	XXX_sizecache            int32                   `json:"-"`
}

func (m *GoTest) Reset()         { *m = GoTest{} }
func (m *GoTest) String() string { return proto.CompactTextString(m) }
func (*GoTest) ProtoMessage()    {}
func (*GoTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{2}
}

func (m *GoTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTest.Unmarshal(m, b)
}
func (m *GoTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTest.Marshal(b, m, deterministic)
}
func (m *GoTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTest.Merge(m, src)
}
func (m *GoTest) XXX_Size() int {
	return xxx_messageInfo_GoTest.Size(m)
}
func (m *GoTest) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTest.DiscardUnknown(m)
}

var xxx_messageInfo_GoTest proto.InternalMessageInfo

const Default_GoTest_F_BoolDefaulted bool = true
const Default_GoTest_F_Int32Defaulted int32 = 32
const Default_GoTest_F_Int64Defaulted int64 = 64
const Default_GoTest_F_Fixed32Defaulted uint32 = 320
const Default_GoTest_F_Fixed64Defaulted uint64 = 640
const Default_GoTest_F_Uint32Defaulted uint32 = 3200
const Default_GoTest_F_Uint64Defaulted uint64 = 6400
const Default_GoTest_F_FloatDefaulted float32 = 314159
const Default_GoTest_F_DoubleDefaulted float64 = 271828
const Default_GoTest_F_StringDefaulted string = "hello, \"world!\"\n"

var Default_GoTest_F_BytesDefaulted []byte = []byte("Bignose")

const Default_GoTest_F_Sint32Defaulted int32 = -32
const Default_GoTest_F_Sint64Defaulted int64 = -64
const Default_GoTest_F_Sfixed32Defaulted int32 = -32
const Default_GoTest_F_Sfixed64Defaulted int64 = -64

func (m *GoTest) GetKind() GoTest_KIND {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return GoTest_VOID
}

func (m *GoTest) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *GoTest) GetParam() int32 {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return 0
}

func (m *GoTest) GetRequiredField() *GoTestField {
	if m != nil {
		return m.RequiredField
	}
	return nil
}

func (m *GoTest) GetRepeatedField() []*GoTestField {
	if m != nil {
		return m.RepeatedField
	}
	return nil
}

func (m *GoTest) GetOptionalField() *GoTestField {
	if m != nil {
		return m.OptionalField
	}
	return nil
}

func (m *GoTest) GetF_BoolRequired() bool {
	if m != nil && m.F_BoolRequired != nil {
		return *m.F_BoolRequired
	}
	return false
}

func (m *GoTest) GetF_Int32Required() int32 {
	if m != nil && m.F_Int32Required != nil {
		return *m.F_Int32Required
	}
	return 0
}

func (m *GoTest) GetF_Int64Required() int64 {
	if m != nil && m.F_Int64Required != nil {
		return *m.F_Int64Required
	}
	return 0
}

func (m *GoTest) GetF_Fixed32Required() uint32 {
	if m != nil && m.F_Fixed32Required != nil {
		return *m.F_Fixed32Required
	}
	return 0
}

func (m *GoTest) GetF_Fixed64Required() uint64 {
	if m != nil && m.F_Fixed64Required != nil {
		return *m.F_Fixed64Required
	}
	return 0
}

func (m *GoTest) GetF_Uint32Required() uint32 {
	if m != nil && m.F_Uint32Required != nil {
		return *m.F_Uint32Required
	}
	return 0
}

func (m *GoTest) GetF_Uint64Required() uint64 {
	if m != nil && m.F_Uint64Required != nil {
		return *m.F_Uint64Required
	}
	return 0
}

func (m *GoTest) GetF_FloatRequired() float32 {
	if m != nil && m.F_FloatRequired != nil {
		return *m.F_FloatRequired
	}
	return 0
}

func (m *GoTest) GetF_DoubleRequired() float64 {
	if m != nil && m.F_DoubleRequired != nil {
		return *m.F_DoubleRequired
	}
	return 0
}

func (m *GoTest) GetF_StringRequired() string {
	if m != nil && m.F_StringRequired != nil {
		return *m.F_StringRequired
	}
	return ""
}

func (m *GoTest) GetF_BytesRequired() []byte {
	if m != nil {
		return m.F_BytesRequired
	}
	return nil
}

func (m *GoTest) GetF_Sint32Required() int32 {
	if m != nil && m.F_Sint32Required != nil {
		return *m.F_Sint32Required
	}
	return 0
}

func (m *GoTest) GetF_Sint64Required() int64 {
	if m != nil && m.F_Sint64Required != nil {
		return *m.F_Sint64Required
	}
	return 0
}

func (m *GoTest) GetF_Sfixed32Required() int32 {
	if m != nil && m.F_Sfixed32Required != nil {
		return *m.F_Sfixed32Required
	}
	return 0
}

func (m *GoTest) GetF_Sfixed64Required() int64 {
	if m != nil && m.F_Sfixed64Required != nil {
		return *m.F_Sfixed64Required
	}
	return 0
}

func (m *GoTest) GetF_BoolRepeated() []bool {
	if m != nil {
		return m.F_BoolRepeated
	}
	return nil
}

func (m *GoTest) GetF_Int32Repeated() []int32 {
	if m != nil {
		return m.F_Int32Repeated
	}
	return nil
}

func (m *GoTest) GetF_Int64Repeated() []int64 {
	if m != nil {
		return m.F_Int64Repeated
	}
	return nil
}

func (m *GoTest) GetF_Fixed32Repeated() []uint32 {
	if m != nil {
		return m.F_Fixed32Repeated
	}
	return nil
}

func (m *GoTest) GetF_Fixed64Repeated() []uint64 {
	if m != nil {
		return m.F_Fixed64Repeated
	}
	return nil
}

func (m *GoTest) GetF_Uint32Repeated() []uint32 {
	if m != nil {
		return m.F_Uint32Repeated
	}
	return nil
}

func (m *GoTest) GetF_Uint64Repeated() []uint64 {
	if m != nil {
		return m.F_Uint64Repeated
	}
	return nil
}

func (m *GoTest) GetF_FloatRepeated() []float32 {
	if m != nil {
		return m.F_FloatRepeated
	}
	return nil
}

func (m *GoTest) GetF_DoubleRepeated() []float64 {
	if m != nil {
		return m.F_DoubleRepeated
	}
	return nil
}

func (m *GoTest) GetF_StringRepeated() []string {
	if m != nil {
		return m.F_StringRepeated
	}
	return nil
}

func (m *GoTest) GetF_BytesRepeated() [][]byte {
	if m != nil {
		return m.F_BytesRepeated
	}
	return nil
}

func (m *GoTest) GetF_Sint32Repeated() []int32 {
	if m != nil {
		return m.F_Sint32Repeated
	}
	return nil
}

func (m *GoTest) GetF_Sint64Repeated() []int64 {
	if m != nil {
		return m.F_Sint64Repeated
	}
	return nil
}

func (m *GoTest) GetF_Sfixed32Repeated() []int32 {
	if m != nil {
		return m.F_Sfixed32Repeated
	}
	return nil
}

func (m *GoTest) GetF_Sfixed64Repeated() []int64 {
	if m != nil {
		return m.F_Sfixed64Repeated
	}
	return nil
}

func (m *GoTest) GetF_BoolOptional() bool {
	if m != nil && m.F_BoolOptional != nil {
		return *m.F_BoolOptional
	}
	return false
}

func (m *GoTest) GetF_Int32Optional() int32 {
	if m != nil && m.F_Int32Optional != nil {
		return *m.F_Int32Optional
	}
	return 0
}

func (m *GoTest) GetF_Int64Optional() int64 {
	if m != nil && m.F_Int64Optional != nil {
		return *m.F_Int64Optional
	}
	return 0
}

func (m *GoTest) GetF_Fixed32Optional() uint32 {
	if m != nil && m.F_Fixed32Optional != nil {
		return *m.F_Fixed32Optional
	}
	return 0
}

func (m *GoTest) GetF_Fixed64Optional() uint64 {
	if m != nil && m.F_Fixed64Optional != nil {
		return *m.F_Fixed64Optional
	}
	return 0
}

func (m *GoTest) GetF_Uint32Optional() uint32 {
	if m != nil && m.F_Uint32Optional != nil {
		return *m.F_Uint32Optional
	}
	return 0
}

func (m *GoTest) GetF_Uint64Optional() uint64 {
	if m != nil && m.F_Uint64Optional != nil {
		return *m.F_Uint64Optional
	}
	return 0
}

func (m *GoTest) GetF_FloatOptional() float32 {
	if m != nil && m.F_FloatOptional != nil {
		return *m.F_FloatOptional
	}
	return 0
}

func (m *GoTest) GetF_DoubleOptional() float64 {
	if m != nil && m.F_DoubleOptional != nil {
		return *m.F_DoubleOptional
	}
	return 0
}

func (m *GoTest) GetF_StringOptional() string {
	if m != nil && m.F_StringOptional != nil {
		return *m.F_StringOptional
	}
	return ""
}

func (m *GoTest) GetF_BytesOptional() []byte {
	if m != nil {
		return m.F_BytesOptional
	}
	return nil
}

func (m *GoTest) GetF_Sint32Optional() int32 {
	if m != nil && m.F_Sint32Optional != nil {
		return *m.F_Sint32Optional
	}
	return 0
}

func (m *GoTest) GetF_Sint64Optional() int64 {
	if m != nil && m.F_Sint64Optional != nil {
		return *m.F_Sint64Optional
	}
	return 0
}

func (m *GoTest) GetF_Sfixed32Optional() int32 {
	if m != nil && m.F_Sfixed32Optional != nil {
		return *m.F_Sfixed32Optional
	}
	return 0
}

func (m *GoTest) GetF_Sfixed64Optional() int64 {
	if m != nil && m.F_Sfixed64Optional != nil {
		return *m.F_Sfixed64Optional
	}
	return 0
}

func (m *GoTest) GetF_BoolDefaulted() bool {
	if m != nil && m.F_BoolDefaulted != nil {
		return *m.F_BoolDefaulted
	}
	return Default_GoTest_F_BoolDefaulted
}

func (m *GoTest) GetF_Int32Defaulted() int32 {
	if m != nil && m.F_Int32Defaulted != nil {
		return *m.F_Int32Defaulted
	}
	return Default_GoTest_F_Int32Defaulted
}

func (m *GoTest) GetF_Int64Defaulted() int64 {
	if m != nil && m.F_Int64Defaulted != nil {
		return *m.F_Int64Defaulted
	}
	return Default_GoTest_F_Int64Defaulted
}

func (m *GoTest) GetF_Fixed32Defaulted() uint32 {
	if m != nil && m.F_Fixed32Defaulted != nil {
		return *m.F_Fixed32Defaulted
	}
	return Default_GoTest_F_Fixed32Defaulted
}

func (m *GoTest) GetF_Fixed64Defaulted() uint64 {
	if m != nil && m.F_Fixed64Defaulted != nil {
		return *m.F_Fixed64Defaulted
	}
	return Default_GoTest_F_Fixed64Defaulted
}

func (m *GoTest) GetF_Uint32Defaulted() uint32 {
	if m != nil && m.F_Uint32Defaulted != nil {
		return *m.F_Uint32Defaulted
	}
	return Default_GoTest_F_Uint32Defaulted
}

func (m *GoTest) GetF_Uint64Defaulted() uint64 {
	if m != nil && m.F_Uint64Defaulted != nil {
		return *m.F_Uint64Defaulted
	}
	return Default_GoTest_F_Uint64Defaulted
}

func (m *GoTest) GetF_FloatDefaulted() float32 {
	if m != nil && m.F_FloatDefaulted != nil {
		return *m.F_FloatDefaulted
	}
	return Default_GoTest_F_FloatDefaulted
}

func (m *GoTest) GetF_DoubleDefaulted() float64 {
	if m != nil && m.F_DoubleDefaulted != nil {
		return *m.F_DoubleDefaulted
	}
	return Default_GoTest_F_DoubleDefaulted
}

func (m *GoTest) GetF_StringDefaulted() string {
	if m != nil && m.F_StringDefaulted != nil {
		return *m.F_StringDefaulted
	}
	return Default_GoTest_F_StringDefaulted
}

func (m *GoTest) GetF_BytesDefaulted() []byte {
	if m != nil && m.F_BytesDefaulted != nil {
		return m.F_BytesDefaulted
	}
	return append([]byte(nil), Default_GoTest_F_BytesDefaulted...)
}

func (m *GoTest) GetF_Sint32Defaulted() int32 {
	if m != nil && m.F_Sint32Defaulted != nil {
		return *m.F_Sint32Defaulted
	}
	return Default_GoTest_F_Sint32Defaulted
}

func (m *GoTest) GetF_Sint64Defaulted() int64 {
	if m != nil && m.F_Sint64Defaulted != nil {
		return *m.F_Sint64Defaulted
	}
	return Default_GoTest_F_Sint64Defaulted
}

func (m *GoTest) GetF_Sfixed32Defaulted() int32 {
	if m != nil && m.F_Sfixed32Defaulted != nil {
		return *m.F_Sfixed32Defaulted
	}
	return Default_GoTest_F_Sfixed32Defaulted
}

func (m *GoTest) GetF_Sfixed64Defaulted() int64 {
	if m != nil && m.F_Sfixed64Defaulted != nil {
		return *m.F_Sfixed64Defaulted
	}
	return Default_GoTest_F_Sfixed64Defaulted
}

func (m *GoTest) GetF_BoolRepeatedPacked() []bool {
	if m != nil {
		return m.F_BoolRepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Int32RepeatedPacked() []int32 {
	if m != nil {
		return m.F_Int32RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Int64RepeatedPacked() []int64 {
	if m != nil {
		return m.F_Int64RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Fixed32RepeatedPacked() []uint32 {
	if m != nil {
		return m.F_Fixed32RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Fixed64RepeatedPacked() []uint64 {
	if m != nil {
		return m.F_Fixed64RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Uint32RepeatedPacked() []uint32 {
	if m != nil {
		return m.F_Uint32RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Uint64RepeatedPacked() []uint64 {
	if m != nil {
		return m.F_Uint64RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_FloatRepeatedPacked() []float32 {
	if m != nil {
		return m.F_FloatRepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_DoubleRepeatedPacked() []float64 {
	if m != nil {
		return m.F_DoubleRepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Sint32RepeatedPacked() []int32 {
	if m != nil {
		return m.F_Sint32RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Sint64RepeatedPacked() []int64 {
	if m != nil {
		return m.F_Sint64RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Sfixed32RepeatedPacked() []int32 {
	if m != nil {
		return m.F_Sfixed32RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetF_Sfixed64RepeatedPacked() []int64 {
	if m != nil {
		return m.F_Sfixed64RepeatedPacked
	}
	return nil
}

func (m *GoTest) GetRequiredgroup() *GoTest_RequiredGroup {
	if m != nil {
		return m.Requiredgroup
	}
	return nil
}

func (m *GoTest) GetRepeatedgroup() []*GoTest_RepeatedGroup {
	if m != nil {
		return m.Repeatedgroup
	}
	return nil
}

func (m *GoTest) GetOptionalgroup() *GoTest_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

// Required, repeated, and optional groups.
type GoTest_RequiredGroup struct {
	RequiredField        *string  `protobuf:"bytes,71,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoTest_RequiredGroup) Reset()         { *m = GoTest_RequiredGroup{} }
func (m *GoTest_RequiredGroup) String() string { return proto.CompactTextString(m) }
func (*GoTest_RequiredGroup) ProtoMessage()    {}
func (*GoTest_RequiredGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{2, 0}
}

func (m *GoTest_RequiredGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTest_RequiredGroup.Unmarshal(m, b)
}
func (m *GoTest_RequiredGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTest_RequiredGroup.Marshal(b, m, deterministic)
}
func (m *GoTest_RequiredGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTest_RequiredGroup.Merge(m, src)
}
func (m *GoTest_RequiredGroup) XXX_Size() int {
	return xxx_messageInfo_GoTest_RequiredGroup.Size(m)
}
func (m *GoTest_RequiredGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTest_RequiredGroup.DiscardUnknown(m)
}

var xxx_messageInfo_GoTest_RequiredGroup proto.InternalMessageInfo

func (m *GoTest_RequiredGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

type GoTest_RepeatedGroup struct {
	RequiredField        *string  `protobuf:"bytes,81,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoTest_RepeatedGroup) Reset()         { *m = GoTest_RepeatedGroup{} }
func (m *GoTest_RepeatedGroup) String() string { return proto.CompactTextString(m) }
func (*GoTest_RepeatedGroup) ProtoMessage()    {}
func (*GoTest_RepeatedGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{2, 1}
}

func (m *GoTest_RepeatedGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTest_RepeatedGroup.Unmarshal(m, b)
}
func (m *GoTest_RepeatedGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTest_RepeatedGroup.Marshal(b, m, deterministic)
}
func (m *GoTest_RepeatedGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTest_RepeatedGroup.Merge(m, src)
}
func (m *GoTest_RepeatedGroup) XXX_Size() int {
	return xxx_messageInfo_GoTest_RepeatedGroup.Size(m)
}
func (m *GoTest_RepeatedGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTest_RepeatedGroup.DiscardUnknown(m)
}

var xxx_messageInfo_GoTest_RepeatedGroup proto.InternalMessageInfo

func (m *GoTest_RepeatedGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

type GoTest_OptionalGroup struct {
	RequiredField        *string  `protobuf:"bytes,91,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoTest_OptionalGroup) Reset()         { *m = GoTest_OptionalGroup{} }
func (m *GoTest_OptionalGroup) String() string { return proto.CompactTextString(m) }
func (*GoTest_OptionalGroup) ProtoMessage()    {}
func (*GoTest_OptionalGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{2, 2}
}

func (m *GoTest_OptionalGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTest_OptionalGroup.Unmarshal(m, b)
}
func (m *GoTest_OptionalGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTest_OptionalGroup.Marshal(b, m, deterministic)
}
func (m *GoTest_OptionalGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTest_OptionalGroup.Merge(m, src)
}
func (m *GoTest_OptionalGroup) XXX_Size() int {
	return xxx_messageInfo_GoTest_OptionalGroup.Size(m)
}
func (m *GoTest_OptionalGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTest_OptionalGroup.DiscardUnknown(m)
}

var xxx_messageInfo_GoTest_OptionalGroup proto.InternalMessageInfo

func (m *GoTest_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

// For testing a group containing a required field.
type GoTestRequiredGroupField struct {
	Group                *GoTestRequiredGroupField_Group `protobuf:"group,1,req,name=Group,json=group" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GoTestRequiredGroupField) Reset()         { *m = GoTestRequiredGroupField{} }
func (m *GoTestRequiredGroupField) String() string { return proto.CompactTextString(m) }
func (*GoTestRequiredGroupField) ProtoMessage()    {}
func (*GoTestRequiredGroupField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{3}
}

func (m *GoTestRequiredGroupField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTestRequiredGroupField.Unmarshal(m, b)
}
func (m *GoTestRequiredGroupField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTestRequiredGroupField.Marshal(b, m, deterministic)
}
func (m *GoTestRequiredGroupField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTestRequiredGroupField.Merge(m, src)
}
func (m *GoTestRequiredGroupField) XXX_Size() int {
	return xxx_messageInfo_GoTestRequiredGroupField.Size(m)
}
func (m *GoTestRequiredGroupField) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTestRequiredGroupField.DiscardUnknown(m)
}

var xxx_messageInfo_GoTestRequiredGroupField proto.InternalMessageInfo

func (m *GoTestRequiredGroupField) GetGroup() *GoTestRequiredGroupField_Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GoTestRequiredGroupField_Group struct {
	Field                *int32   `protobuf:"varint,2,req,name=Field" json:"Field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoTestRequiredGroupField_Group) Reset()         { *m = GoTestRequiredGroupField_Group{} }
func (m *GoTestRequiredGroupField_Group) String() string { return proto.CompactTextString(m) }
func (*GoTestRequiredGroupField_Group) ProtoMessage()    {}
func (*GoTestRequiredGroupField_Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{3, 0}
}

func (m *GoTestRequiredGroupField_Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoTestRequiredGroupField_Group.Unmarshal(m, b)
}
func (m *GoTestRequiredGroupField_Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoTestRequiredGroupField_Group.Marshal(b, m, deterministic)
}
func (m *GoTestRequiredGroupField_Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoTestRequiredGroupField_Group.Merge(m, src)
}
func (m *GoTestRequiredGroupField_Group) XXX_Size() int {
	return xxx_messageInfo_GoTestRequiredGroupField_Group.Size(m)
}
func (m *GoTestRequiredGroupField_Group) XXX_DiscardUnknown() {
	xxx_messageInfo_GoTestRequiredGroupField_Group.DiscardUnknown(m)
}

var xxx_messageInfo_GoTestRequiredGroupField_Group proto.InternalMessageInfo

func (m *GoTestRequiredGroupField_Group) GetField() int32 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

// For testing skipping of unrecognized fields.
// Numbers are all big, larger than tag numbers in GoTestField,
// the message used in the corresponding test.
type GoSkipTest struct {
	SkipInt32            *int32                `protobuf:"varint,11,req,name=skip_int32,json=skipInt32" json:"skip_int32,omitempty"`
	SkipFixed32          *uint32               `protobuf:"fixed32,12,req,name=skip_fixed32,json=skipFixed32" json:"skip_fixed32,omitempty"`
	SkipFixed64          *uint64               `protobuf:"fixed64,13,req,name=skip_fixed64,json=skipFixed64" json:"skip_fixed64,omitempty"`
	SkipString           *string               `protobuf:"bytes,14,req,name=skip_string,json=skipString" json:"skip_string,omitempty"`
	Skipgroup            *GoSkipTest_SkipGroup `protobuf:"group,15,req,name=SkipGroup,json=skipgroup" json:"skipgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GoSkipTest) Reset()         { *m = GoSkipTest{} }
func (m *GoSkipTest) String() string { return proto.CompactTextString(m) }
func (*GoSkipTest) ProtoMessage()    {}
func (*GoSkipTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{4}
}

func (m *GoSkipTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoSkipTest.Unmarshal(m, b)
}
func (m *GoSkipTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoSkipTest.Marshal(b, m, deterministic)
}
func (m *GoSkipTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoSkipTest.Merge(m, src)
}
func (m *GoSkipTest) XXX_Size() int {
	return xxx_messageInfo_GoSkipTest.Size(m)
}
func (m *GoSkipTest) XXX_DiscardUnknown() {
	xxx_messageInfo_GoSkipTest.DiscardUnknown(m)
}

var xxx_messageInfo_GoSkipTest proto.InternalMessageInfo

func (m *GoSkipTest) GetSkipInt32() int32 {
	if m != nil && m.SkipInt32 != nil {
		return *m.SkipInt32
	}
	return 0
}

func (m *GoSkipTest) GetSkipFixed32() uint32 {
	if m != nil && m.SkipFixed32 != nil {
		return *m.SkipFixed32
	}
	return 0
}

func (m *GoSkipTest) GetSkipFixed64() uint64 {
	if m != nil && m.SkipFixed64 != nil {
		return *m.SkipFixed64
	}
	return 0
}

func (m *GoSkipTest) GetSkipString() string {
	if m != nil && m.SkipString != nil {
		return *m.SkipString
	}
	return ""
}

func (m *GoSkipTest) GetSkipgroup() *GoSkipTest_SkipGroup {
	if m != nil {
		return m.Skipgroup
	}
	return nil
}

type GoSkipTest_SkipGroup struct {
	GroupInt32           *int32   `protobuf:"varint,16,req,name=group_int32,json=groupInt32" json:"group_int32,omitempty"`
	GroupString          *string  `protobuf:"bytes,17,req,name=group_string,json=groupString" json:"group_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoSkipTest_SkipGroup) Reset()         { *m = GoSkipTest_SkipGroup{} }
func (m *GoSkipTest_SkipGroup) String() string { return proto.CompactTextString(m) }
func (*GoSkipTest_SkipGroup) ProtoMessage()    {}
func (*GoSkipTest_SkipGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{4, 0}
}

func (m *GoSkipTest_SkipGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoSkipTest_SkipGroup.Unmarshal(m, b)
}
func (m *GoSkipTest_SkipGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoSkipTest_SkipGroup.Marshal(b, m, deterministic)
}
func (m *GoSkipTest_SkipGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoSkipTest_SkipGroup.Merge(m, src)
}
func (m *GoSkipTest_SkipGroup) XXX_Size() int {
	return xxx_messageInfo_GoSkipTest_SkipGroup.Size(m)
}
func (m *GoSkipTest_SkipGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_GoSkipTest_SkipGroup.DiscardUnknown(m)
}

var xxx_messageInfo_GoSkipTest_SkipGroup proto.InternalMessageInfo

func (m *GoSkipTest_SkipGroup) GetGroupInt32() int32 {
	if m != nil && m.GroupInt32 != nil {
		return *m.GroupInt32
	}
	return 0
}

func (m *GoSkipTest_SkipGroup) GetGroupString() string {
	if m != nil && m.GroupString != nil {
		return *m.GroupString
	}
	return ""
}

// For testing packed/non-packed decoder switching.
// A serialized instance of one should be deserializable as the other.
type NonPackedTest struct {
	A                    []int32  `protobuf:"varint,1,rep,name=a" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonPackedTest) Reset()         { *m = NonPackedTest{} }
func (m *NonPackedTest) String() string { return proto.CompactTextString(m) }
func (*NonPackedTest) ProtoMessage()    {}
func (*NonPackedTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{5}
}

func (m *NonPackedTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonPackedTest.Unmarshal(m, b)
}
func (m *NonPackedTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonPackedTest.Marshal(b, m, deterministic)
}
func (m *NonPackedTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonPackedTest.Merge(m, src)
}
func (m *NonPackedTest) XXX_Size() int {
	return xxx_messageInfo_NonPackedTest.Size(m)
}
func (m *NonPackedTest) XXX_DiscardUnknown() {
	xxx_messageInfo_NonPackedTest.DiscardUnknown(m)
}

var xxx_messageInfo_NonPackedTest proto.InternalMessageInfo

func (m *NonPackedTest) GetA() []int32 {
	if m != nil {
		return m.A
	}
	return nil
}

type PackedTest struct {
	B                    []int32  `protobuf:"varint,1,rep,packed,name=b" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackedTest) Reset()         { *m = PackedTest{} }
func (m *PackedTest) String() string { return proto.CompactTextString(m) }
func (*PackedTest) ProtoMessage()    {}
func (*PackedTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{6}
}

func (m *PackedTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackedTest.Unmarshal(m, b)
}
func (m *PackedTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackedTest.Marshal(b, m, deterministic)
}
func (m *PackedTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackedTest.Merge(m, src)
}
func (m *PackedTest) XXX_Size() int {
	return xxx_messageInfo_PackedTest.Size(m)
}
func (m *PackedTest) XXX_DiscardUnknown() {
	xxx_messageInfo_PackedTest.DiscardUnknown(m)
}

var xxx_messageInfo_PackedTest proto.InternalMessageInfo

func (m *PackedTest) GetB() []int32 {
	if m != nil {
		return m.B
	}
	return nil
}

type MaxTag struct {
	// Maximum possible tag number.
	LastField            *string  `protobuf:"bytes,536870911,opt,name=last_field,json=lastField" json:"last_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxTag) Reset()         { *m = MaxTag{} }
func (m *MaxTag) String() string { return proto.CompactTextString(m) }
func (*MaxTag) ProtoMessage()    {}
func (*MaxTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{7}
}

func (m *MaxTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxTag.Unmarshal(m, b)
}
func (m *MaxTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxTag.Marshal(b, m, deterministic)
}
func (m *MaxTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxTag.Merge(m, src)
}
func (m *MaxTag) XXX_Size() int {
	return xxx_messageInfo_MaxTag.Size(m)
}
func (m *MaxTag) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxTag.DiscardUnknown(m)
}

var xxx_messageInfo_MaxTag proto.InternalMessageInfo

func (m *MaxTag) GetLastField() string {
	if m != nil && m.LastField != nil {
		return *m.LastField
	}
	return ""
}

type OldMessage struct {
	Nested               *OldMessage_Nested `protobuf:"bytes,1,opt,name=nested" json:"nested,omitempty"`
	Num                  *int32             `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OldMessage) Reset()         { *m = OldMessage{} }
func (m *OldMessage) String() string { return proto.CompactTextString(m) }
func (*OldMessage) ProtoMessage()    {}
func (*OldMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{8}
}

func (m *OldMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldMessage.Unmarshal(m, b)
}
func (m *OldMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldMessage.Marshal(b, m, deterministic)
}
func (m *OldMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldMessage.Merge(m, src)
}
func (m *OldMessage) XXX_Size() int {
	return xxx_messageInfo_OldMessage.Size(m)
}
func (m *OldMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OldMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OldMessage proto.InternalMessageInfo

func (m *OldMessage) GetNested() *OldMessage_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *OldMessage) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type OldMessage_Nested struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OldMessage_Nested) Reset()         { *m = OldMessage_Nested{} }
func (m *OldMessage_Nested) String() string { return proto.CompactTextString(m) }
func (*OldMessage_Nested) ProtoMessage()    {}
func (*OldMessage_Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{8, 0}
}

func (m *OldMessage_Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldMessage_Nested.Unmarshal(m, b)
}
func (m *OldMessage_Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldMessage_Nested.Marshal(b, m, deterministic)
}
func (m *OldMessage_Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldMessage_Nested.Merge(m, src)
}
func (m *OldMessage_Nested) XXX_Size() int {
	return xxx_messageInfo_OldMessage_Nested.Size(m)
}
func (m *OldMessage_Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_OldMessage_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_OldMessage_Nested proto.InternalMessageInfo

func (m *OldMessage_Nested) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// NewMessage is wire compatible with OldMessage;
// imagine it as a future version.
type NewMessage struct {
	Nested *NewMessage_Nested `protobuf:"bytes,1,opt,name=nested" json:"nested,omitempty"`
	// This is an int32 in OldMessage.
	Num                  *int64   `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMessage) Reset()         { *m = NewMessage{} }
func (m *NewMessage) String() string { return proto.CompactTextString(m) }
func (*NewMessage) ProtoMessage()    {}
func (*NewMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{9}
}

func (m *NewMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMessage.Unmarshal(m, b)
}
func (m *NewMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMessage.Marshal(b, m, deterministic)
}
func (m *NewMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMessage.Merge(m, src)
}
func (m *NewMessage) XXX_Size() int {
	return xxx_messageInfo_NewMessage.Size(m)
}
func (m *NewMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NewMessage proto.InternalMessageInfo

func (m *NewMessage) GetNested() *NewMessage_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *NewMessage) GetNum() int64 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type NewMessage_Nested struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	FoodGroup            *string  `protobuf:"bytes,2,opt,name=food_group,json=foodGroup" json:"food_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewMessage_Nested) Reset()         { *m = NewMessage_Nested{} }
func (m *NewMessage_Nested) String() string { return proto.CompactTextString(m) }
func (*NewMessage_Nested) ProtoMessage()    {}
func (*NewMessage_Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{9, 0}
}

func (m *NewMessage_Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewMessage_Nested.Unmarshal(m, b)
}
func (m *NewMessage_Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewMessage_Nested.Marshal(b, m, deterministic)
}
func (m *NewMessage_Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewMessage_Nested.Merge(m, src)
}
func (m *NewMessage_Nested) XXX_Size() int {
	return xxx_messageInfo_NewMessage_Nested.Size(m)
}
func (m *NewMessage_Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_NewMessage_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_NewMessage_Nested proto.InternalMessageInfo

func (m *NewMessage_Nested) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NewMessage_Nested) GetFoodGroup() string {
	if m != nil && m.FoodGroup != nil {
		return *m.FoodGroup
	}
	return ""
}

type InnerMessage struct {
	Host                 *string  `protobuf:"bytes,1,req,name=host" json:"host,omitempty"`
	Port                 *int32   `protobuf:"varint,2,opt,name=port,def=4000" json:"port,omitempty"`
	Connected            *bool    `protobuf:"varint,3,opt,name=connected" json:"connected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerMessage) Reset()         { *m = InnerMessage{} }
func (m *InnerMessage) String() string { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()    {}
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{10}
}

func (m *InnerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerMessage.Unmarshal(m, b)
}
func (m *InnerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerMessage.Marshal(b, m, deterministic)
}
func (m *InnerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMessage.Merge(m, src)
}
func (m *InnerMessage) XXX_Size() int {
	return xxx_messageInfo_InnerMessage.Size(m)
}
func (m *InnerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMessage proto.InternalMessageInfo

const Default_InnerMessage_Port int32 = 4000

func (m *InnerMessage) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *InnerMessage) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_InnerMessage_Port
}

func (m *InnerMessage) GetConnected() bool {
	if m != nil && m.Connected != nil {
		return *m.Connected
	}
	return false
}

type OtherMessage struct {
	Key                          *int64        `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Value                        []byte        `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Weight                       *float32      `protobuf:"fixed32,3,opt,name=weight" json:"weight,omitempty"`
	Inner                        *InnerMessage `protobuf:"bytes,4,opt,name=inner" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}      `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *OtherMessage) Reset()         { *m = OtherMessage{} }
func (m *OtherMessage) String() string { return proto.CompactTextString(m) }
func (*OtherMessage) ProtoMessage()    {}
func (*OtherMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{11}
}

var extRange_OtherMessage = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*OtherMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OtherMessage
}

func (m *OtherMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherMessage.Unmarshal(m, b)
}
func (m *OtherMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherMessage.Marshal(b, m, deterministic)
}
func (m *OtherMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherMessage.Merge(m, src)
}
func (m *OtherMessage) XXX_Size() int {
	return xxx_messageInfo_OtherMessage.Size(m)
}
func (m *OtherMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OtherMessage proto.InternalMessageInfo

func (m *OtherMessage) GetKey() int64 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *OtherMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *OtherMessage) GetWeight() float32 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return 0
}

func (m *OtherMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

type RequiredInnerMessage struct {
	LeoFinallyWonAnOscar *InnerMessage `protobuf:"bytes,1,req,name=leo_finally_won_an_oscar,json=leoFinallyWonAnOscar" json:"leo_finally_won_an_oscar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RequiredInnerMessage) Reset()         { *m = RequiredInnerMessage{} }
func (m *RequiredInnerMessage) String() string { return proto.CompactTextString(m) }
func (*RequiredInnerMessage) ProtoMessage()    {}
func (*RequiredInnerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{12}
}

func (m *RequiredInnerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequiredInnerMessage.Unmarshal(m, b)
}
func (m *RequiredInnerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequiredInnerMessage.Marshal(b, m, deterministic)
}
func (m *RequiredInnerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredInnerMessage.Merge(m, src)
}
func (m *RequiredInnerMessage) XXX_Size() int {
	return xxx_messageInfo_RequiredInnerMessage.Size(m)
}
func (m *RequiredInnerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredInnerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredInnerMessage proto.InternalMessageInfo

func (m *RequiredInnerMessage) GetLeoFinallyWonAnOscar() *InnerMessage {
	if m != nil {
		return m.LeoFinallyWonAnOscar
	}
	return nil
}

type MyMessage struct {
	Count          *int32                `protobuf:"varint,1,req,name=count" json:"count,omitempty"`
	Name           *string               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Quote          *string               `protobuf:"bytes,3,opt,name=quote" json:"quote,omitempty"`
	Pet            []string              `protobuf:"bytes,4,rep,name=pet" json:"pet,omitempty"`
	Inner          *InnerMessage         `protobuf:"bytes,5,opt,name=inner" json:"inner,omitempty"`
	Others         []*OtherMessage       `protobuf:"bytes,6,rep,name=others" json:"others,omitempty"`
	WeMustGoDeeper *RequiredInnerMessage `protobuf:"bytes,13,opt,name=we_must_go_deeper,json=weMustGoDeeper" json:"we_must_go_deeper,omitempty"`
	RepInner       []*InnerMessage       `protobuf:"bytes,12,rep,name=rep_inner,json=repInner" json:"rep_inner,omitempty"`
	Bikeshed       *MyMessage_Color      `protobuf:"varint,7,opt,name=bikeshed,enum=proto2_test.MyMessage_Color" json:"bikeshed,omitempty"`
	Somegroup      *MyMessage_SomeGroup  `protobuf:"group,8,opt,name=SomeGroup,json=somegroup" json:"somegroup,omitempty"`
	// This field becomes [][]byte in the generated code.
	RepBytes                     [][]byte `protobuf:"bytes,10,rep,name=rep_bytes,json=repBytes" json:"rep_bytes,omitempty"`
	Bigfloat                     *float64 `protobuf:"fixed64,11,opt,name=bigfloat" json:"bigfloat,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{13}
}

var extRange_MyMessage = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*MyMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MyMessage
}

func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage.Unmarshal(m, b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
}
func (m *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(m, src)
}
func (m *MyMessage) XXX_Size() int {
	return xxx_messageInfo_MyMessage.Size(m)
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *MyMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MyMessage) GetQuote() string {
	if m != nil && m.Quote != nil {
		return *m.Quote
	}
	return ""
}

func (m *MyMessage) GetPet() []string {
	if m != nil {
		return m.Pet
	}
	return nil
}

func (m *MyMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (m *MyMessage) GetOthers() []*OtherMessage {
	if m != nil {
		return m.Others
	}
	return nil
}

func (m *MyMessage) GetWeMustGoDeeper() *RequiredInnerMessage {
	if m != nil {
		return m.WeMustGoDeeper
	}
	return nil
}

func (m *MyMessage) GetRepInner() []*InnerMessage {
	if m != nil {
		return m.RepInner
	}
	return nil
}

func (m *MyMessage) GetBikeshed() MyMessage_Color {
	if m != nil && m.Bikeshed != nil {
		return *m.Bikeshed
	}
	return MyMessage_RED
}

func (m *MyMessage) GetSomegroup() *MyMessage_SomeGroup {
	if m != nil {
		return m.Somegroup
	}
	return nil
}

func (m *MyMessage) GetRepBytes() [][]byte {
	if m != nil {
		return m.RepBytes
	}
	return nil
}

func (m *MyMessage) GetBigfloat() float64 {
	if m != nil && m.Bigfloat != nil {
		return *m.Bigfloat
	}
	return 0
}

type MyMessage_SomeGroup struct {
	GroupField           *int32   `protobuf:"varint,9,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMessage_SomeGroup) Reset()         { *m = MyMessage_SomeGroup{} }
func (m *MyMessage_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*MyMessage_SomeGroup) ProtoMessage()    {}
func (*MyMessage_SomeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{13, 0}
}

func (m *MyMessage_SomeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage_SomeGroup.Unmarshal(m, b)
}
func (m *MyMessage_SomeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage_SomeGroup.Marshal(b, m, deterministic)
}
func (m *MyMessage_SomeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage_SomeGroup.Merge(m, src)
}
func (m *MyMessage_SomeGroup) XXX_Size() int {
	return xxx_messageInfo_MyMessage_SomeGroup.Size(m)
}
func (m *MyMessage_SomeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage_SomeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage_SomeGroup proto.InternalMessageInfo

func (m *MyMessage_SomeGroup) GetGroupField() int32 {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return 0
}

type Ext struct {
	Data                 *string         `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	MapField             map[int32]int32 `protobuf:"bytes,2,rep,name=map_field,json=mapField" json:"map_field,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Ext) Reset()         { *m = Ext{} }
func (m *Ext) String() string { return proto.CompactTextString(m) }
func (*Ext) ProtoMessage()    {}
func (*Ext) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{14}
}

func (m *Ext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ext.Unmarshal(m, b)
}
func (m *Ext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ext.Marshal(b, m, deterministic)
}
func (m *Ext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ext.Merge(m, src)
}
func (m *Ext) XXX_Size() int {
	return xxx_messageInfo_Ext.Size(m)
}
func (m *Ext) XXX_DiscardUnknown() {
	xxx_messageInfo_Ext.DiscardUnknown(m)
}

var xxx_messageInfo_Ext proto.InternalMessageInfo

func (m *Ext) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *Ext) GetMapField() map[int32]int32 {
	if m != nil {
		return m.MapField
	}
	return nil
}

var E_Ext_More = &proto.ExtensionDesc{
	ExtendedType:  (*MyMessage)(nil),
	ExtensionType: (*Ext)(nil),
	Field:         103,
	Name:          "proto2_test.Ext.more",
	Tag:           "bytes,103,opt,name=more",
	Filename:      "proto2_proto/test.proto",
}

var E_Ext_Text = &proto.ExtensionDesc{
	ExtendedType:  (*MyMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         104,
	Name:          "proto2_test.Ext.text",
	Tag:           "bytes,104,opt,name=text",
	Filename:      "proto2_proto/test.proto",
}

var E_Ext_Number = &proto.ExtensionDesc{
	ExtendedType:  (*MyMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         105,
	Name:          "proto2_test.Ext.number",
	Tag:           "varint,105,opt,name=number",
	Filename:      "proto2_proto/test.proto",
}

type ComplexExtension struct {
	First                *int32   `protobuf:"varint,1,opt,name=first" json:"first,omitempty"`
	Second               *int32   `protobuf:"varint,2,opt,name=second" json:"second,omitempty"`
	Third                []int32  `protobuf:"varint,3,rep,name=third" json:"third,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplexExtension) Reset()         { *m = ComplexExtension{} }
func (m *ComplexExtension) String() string { return proto.CompactTextString(m) }
func (*ComplexExtension) ProtoMessage()    {}
func (*ComplexExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{15}
}

func (m *ComplexExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexExtension.Unmarshal(m, b)
}
func (m *ComplexExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexExtension.Marshal(b, m, deterministic)
}
func (m *ComplexExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexExtension.Merge(m, src)
}
func (m *ComplexExtension) XXX_Size() int {
	return xxx_messageInfo_ComplexExtension.Size(m)
}
func (m *ComplexExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexExtension.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexExtension proto.InternalMessageInfo

func (m *ComplexExtension) GetFirst() int32 {
	if m != nil && m.First != nil {
		return *m.First
	}
	return 0
}

func (m *ComplexExtension) GetSecond() int32 {
	if m != nil && m.Second != nil {
		return *m.Second
	}
	return 0
}

func (m *ComplexExtension) GetThird() []int32 {
	if m != nil {
		return m.Third
	}
	return nil
}

type DefaultsMessage struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *DefaultsMessage) Reset()         { *m = DefaultsMessage{} }
func (m *DefaultsMessage) String() string { return proto.CompactTextString(m) }
func (*DefaultsMessage) ProtoMessage()    {}
func (*DefaultsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{16}
}

var extRange_DefaultsMessage = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*DefaultsMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_DefaultsMessage
}

func (m *DefaultsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultsMessage.Unmarshal(m, b)
}
func (m *DefaultsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultsMessage.Marshal(b, m, deterministic)
}
func (m *DefaultsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultsMessage.Merge(m, src)
}
func (m *DefaultsMessage) XXX_Size() int {
	return xxx_messageInfo_DefaultsMessage.Size(m)
}
func (m *DefaultsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultsMessage proto.InternalMessageInfo

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{17}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type MessageList struct {
	Message              []*MessageList_Message `protobuf:"group,1,rep,name=Message,json=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MessageList) Reset()         { *m = MessageList{} }
func (m *MessageList) String() string { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()    {}
func (*MessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{18}
}

func (m *MessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList.Unmarshal(m, b)
}
func (m *MessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList.Marshal(b, m, deterministic)
}
func (m *MessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList.Merge(m, src)
}
func (m *MessageList) XXX_Size() int {
	return xxx_messageInfo_MessageList.Size(m)
}
func (m *MessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList proto.InternalMessageInfo

func (m *MessageList) GetMessage() []*MessageList_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type MessageList_Message struct {
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Count                *int32   `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageList_Message) Reset()         { *m = MessageList_Message{} }
func (m *MessageList_Message) String() string { return proto.CompactTextString(m) }
func (*MessageList_Message) ProtoMessage()    {}
func (*MessageList_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{18, 0}
}

func (m *MessageList_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList_Message.Unmarshal(m, b)
}
func (m *MessageList_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList_Message.Marshal(b, m, deterministic)
}
func (m *MessageList_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList_Message.Merge(m, src)
}
func (m *MessageList_Message) XXX_Size() int {
	return xxx_messageInfo_MessageList_Message.Size(m)
}
func (m *MessageList_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList_Message.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList_Message proto.InternalMessageInfo

func (m *MessageList_Message) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MessageList_Message) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type Strings struct {
	StringField          *string  `protobuf:"bytes,1,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	BytesField           []byte   `protobuf:"bytes,2,opt,name=bytes_field,json=bytesField" json:"bytes_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Strings) Reset()         { *m = Strings{} }
func (m *Strings) String() string { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()    {}
func (*Strings) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{19}
}

func (m *Strings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strings.Unmarshal(m, b)
}
func (m *Strings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strings.Marshal(b, m, deterministic)
}
func (m *Strings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strings.Merge(m, src)
}
func (m *Strings) XXX_Size() int {
	return xxx_messageInfo_Strings.Size(m)
}
func (m *Strings) XXX_DiscardUnknown() {
	xxx_messageInfo_Strings.DiscardUnknown(m)
}

var xxx_messageInfo_Strings proto.InternalMessageInfo

func (m *Strings) GetStringField() string {
	if m != nil && m.StringField != nil {
		return *m.StringField
	}
	return ""
}

func (m *Strings) GetBytesField() []byte {
	if m != nil {
		return m.BytesField
	}
	return nil
}

type Defaults struct {
	// Default-valued fields of all basic types.
	// Same as GoTest, but copied here to make testing easier.
	F_Bool    *bool           `protobuf:"varint,1,opt,name=F_Bool,json=FBool,def=1" json:"F_Bool,omitempty"`
	F_Int32   *int32          `protobuf:"varint,2,opt,name=F_Int32,json=FInt32,def=32" json:"F_Int32,omitempty"`
	F_Int64   *int64          `protobuf:"varint,3,opt,name=F_Int64,json=FInt64,def=64" json:"F_Int64,omitempty"`
	F_Fixed32 *uint32         `protobuf:"fixed32,4,opt,name=F_Fixed32,json=FFixed32,def=320" json:"F_Fixed32,omitempty"`
	F_Fixed64 *uint64         `protobuf:"fixed64,5,opt,name=F_Fixed64,json=FFixed64,def=640" json:"F_Fixed64,omitempty"`
	F_Uint32  *uint32         `protobuf:"varint,6,opt,name=F_Uint32,json=FUint32,def=3200" json:"F_Uint32,omitempty"`
	F_Uint64  *uint64         `protobuf:"varint,7,opt,name=F_Uint64,json=FUint64,def=6400" json:"F_Uint64,omitempty"`
	F_Float   *float32        `protobuf:"fixed32,8,opt,name=F_Float,json=FFloat,def=314159" json:"F_Float,omitempty"`
	F_Double  *float64        `protobuf:"fixed64,9,opt,name=F_Double,json=FDouble,def=271828" json:"F_Double,omitempty"`
	F_String  *string         `protobuf:"bytes,10,opt,name=F_String,json=FString,def=hello, \"world!\"\n" json:"F_String,omitempty"`
	F_Bytes   []byte          `protobuf:"bytes,11,opt,name=F_Bytes,json=FBytes,def=Bignose" json:"F_Bytes,omitempty"`
	F_Sint32  *int32          `protobuf:"zigzag32,12,opt,name=F_Sint32,json=FSint32,def=-32" json:"F_Sint32,omitempty"`
	F_Sint64  *int64          `protobuf:"zigzag64,13,opt,name=F_Sint64,json=FSint64,def=-64" json:"F_Sint64,omitempty"`
	F_Enum    *Defaults_Color `protobuf:"varint,14,opt,name=F_Enum,json=FEnum,enum=proto2_test.Defaults_Color,def=1" json:"F_Enum,omitempty"`
	// More fields with crazy defaults.
	F_Pinf *float32 `protobuf:"fixed32,15,opt,name=F_Pinf,json=FPinf,def=inf" json:"F_Pinf,omitempty"`
	F_Ninf *float32 `protobuf:"fixed32,16,opt,name=F_Ninf,json=FNinf,def=-inf" json:"F_Ninf,omitempty"`
	F_Nan  *float32 `protobuf:"fixed32,17,opt,name=F_Nan,json=FNan,def=nan" json:"F_Nan,omitempty"`
	// Sub-message.
	Sub *SubDefaults `protobuf:"bytes,18,opt,name=sub" json:"sub,omitempty"`
	// Redundant but explicit defaults.
	StrZero              *string  `protobuf:"bytes,19,opt,name=str_zero,json=strZero,def=" json:"str_zero,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Defaults) Reset()         { *m = Defaults{} }
func (m *Defaults) String() string { return proto.CompactTextString(m) }
func (*Defaults) ProtoMessage()    {}
func (*Defaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{20}
}

func (m *Defaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Defaults.Unmarshal(m, b)
}
func (m *Defaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Defaults.Marshal(b, m, deterministic)
}
func (m *Defaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Defaults.Merge(m, src)
}
func (m *Defaults) XXX_Size() int {
	return xxx_messageInfo_Defaults.Size(m)
}
func (m *Defaults) XXX_DiscardUnknown() {
	xxx_messageInfo_Defaults.DiscardUnknown(m)
}

var xxx_messageInfo_Defaults proto.InternalMessageInfo

const Default_Defaults_F_Bool bool = true
const Default_Defaults_F_Int32 int32 = 32
const Default_Defaults_F_Int64 int64 = 64
const Default_Defaults_F_Fixed32 uint32 = 320
const Default_Defaults_F_Fixed64 uint64 = 640
const Default_Defaults_F_Uint32 uint32 = 3200
const Default_Defaults_F_Uint64 uint64 = 6400
const Default_Defaults_F_Float float32 = 314159
const Default_Defaults_F_Double float64 = 271828
const Default_Defaults_F_String string = "hello, \"world!\"\n"

var Default_Defaults_F_Bytes []byte = []byte("Bignose")

const Default_Defaults_F_Sint32 int32 = -32
const Default_Defaults_F_Sint64 int64 = -64
const Default_Defaults_F_Enum Defaults_Color = Defaults_GREEN

var Default_Defaults_F_Pinf float32 = float32(math.Inf(1))
var Default_Defaults_F_Ninf float32 = float32(math.Inf(-1))
var Default_Defaults_F_Nan float32 = float32(math.NaN())

func (m *Defaults) GetF_Bool() bool {
	if m != nil && m.F_Bool != nil {
		return *m.F_Bool
	}
	return Default_Defaults_F_Bool
}

func (m *Defaults) GetF_Int32() int32 {
	if m != nil && m.F_Int32 != nil {
		return *m.F_Int32
	}
	return Default_Defaults_F_Int32
}

func (m *Defaults) GetF_Int64() int64 {
	if m != nil && m.F_Int64 != nil {
		return *m.F_Int64
	}
	return Default_Defaults_F_Int64
}

func (m *Defaults) GetF_Fixed32() uint32 {
	if m != nil && m.F_Fixed32 != nil {
		return *m.F_Fixed32
	}
	return Default_Defaults_F_Fixed32
}

func (m *Defaults) GetF_Fixed64() uint64 {
	if m != nil && m.F_Fixed64 != nil {
		return *m.F_Fixed64
	}
	return Default_Defaults_F_Fixed64
}

func (m *Defaults) GetF_Uint32() uint32 {
	if m != nil && m.F_Uint32 != nil {
		return *m.F_Uint32
	}
	return Default_Defaults_F_Uint32
}

func (m *Defaults) GetF_Uint64() uint64 {
	if m != nil && m.F_Uint64 != nil {
		return *m.F_Uint64
	}
	return Default_Defaults_F_Uint64
}

func (m *Defaults) GetF_Float() float32 {
	if m != nil && m.F_Float != nil {
		return *m.F_Float
	}
	return Default_Defaults_F_Float
}

func (m *Defaults) GetF_Double() float64 {
	if m != nil && m.F_Double != nil {
		return *m.F_Double
	}
	return Default_Defaults_F_Double
}

func (m *Defaults) GetF_String() string {
	if m != nil && m.F_String != nil {
		return *m.F_String
	}
	return Default_Defaults_F_String
}

func (m *Defaults) GetF_Bytes() []byte {
	if m != nil && m.F_Bytes != nil {
		return m.F_Bytes
	}
	return append([]byte(nil), Default_Defaults_F_Bytes...)
}

func (m *Defaults) GetF_Sint32() int32 {
	if m != nil && m.F_Sint32 != nil {
		return *m.F_Sint32
	}
	return Default_Defaults_F_Sint32
}

func (m *Defaults) GetF_Sint64() int64 {
	if m != nil && m.F_Sint64 != nil {
		return *m.F_Sint64
	}
	return Default_Defaults_F_Sint64
}

func (m *Defaults) GetF_Enum() Defaults_Color {
	if m != nil && m.F_Enum != nil {
		return *m.F_Enum
	}
	return Default_Defaults_F_Enum
}

func (m *Defaults) GetF_Pinf() float32 {
	if m != nil && m.F_Pinf != nil {
		return *m.F_Pinf
	}
	return Default_Defaults_F_Pinf
}

func (m *Defaults) GetF_Ninf() float32 {
	if m != nil && m.F_Ninf != nil {
		return *m.F_Ninf
	}
	return Default_Defaults_F_Ninf
}

func (m *Defaults) GetF_Nan() float32 {
	if m != nil && m.F_Nan != nil {
		return *m.F_Nan
	}
	return Default_Defaults_F_Nan
}

func (m *Defaults) GetSub() *SubDefaults {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *Defaults) GetStrZero() string {
	if m != nil && m.StrZero != nil {
		return *m.StrZero
	}
	return ""
}

type SubDefaults struct {
	N                    *int64   `protobuf:"varint,1,opt,name=n,def=7" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubDefaults) Reset()         { *m = SubDefaults{} }
func (m *SubDefaults) String() string { return proto.CompactTextString(m) }
func (*SubDefaults) ProtoMessage()    {}
func (*SubDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{21}
}

func (m *SubDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubDefaults.Unmarshal(m, b)
}
func (m *SubDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubDefaults.Marshal(b, m, deterministic)
}
func (m *SubDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubDefaults.Merge(m, src)
}
func (m *SubDefaults) XXX_Size() int {
	return xxx_messageInfo_SubDefaults.Size(m)
}
func (m *SubDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_SubDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_SubDefaults proto.InternalMessageInfo

const Default_SubDefaults_N int64 = 7

func (m *SubDefaults) GetN() int64 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return Default_SubDefaults_N
}

type RepeatedEnum struct {
	Color                []RepeatedEnum_Color `protobuf:"varint,1,rep,name=color,enum=proto2_test.RepeatedEnum_Color" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RepeatedEnum) Reset()         { *m = RepeatedEnum{} }
func (m *RepeatedEnum) String() string { return proto.CompactTextString(m) }
func (*RepeatedEnum) ProtoMessage()    {}
func (*RepeatedEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{22}
}

func (m *RepeatedEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedEnum.Unmarshal(m, b)
}
func (m *RepeatedEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedEnum.Marshal(b, m, deterministic)
}
func (m *RepeatedEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedEnum.Merge(m, src)
}
func (m *RepeatedEnum) XXX_Size() int {
	return xxx_messageInfo_RepeatedEnum.Size(m)
}
func (m *RepeatedEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedEnum proto.InternalMessageInfo

func (m *RepeatedEnum) GetColor() []RepeatedEnum_Color {
	if m != nil {
		return m.Color
	}
	return nil
}

type MoreRepeated struct {
	Bools                []bool   `protobuf:"varint,1,rep,name=bools" json:"bools,omitempty"`
	BoolsPacked          []bool   `protobuf:"varint,2,rep,packed,name=bools_packed,json=boolsPacked" json:"bools_packed,omitempty"`
	Ints                 []int32  `protobuf:"varint,3,rep,name=ints" json:"ints,omitempty"`
	IntsPacked           []int32  `protobuf:"varint,4,rep,packed,name=ints_packed,json=intsPacked" json:"ints_packed,omitempty"`
	Int64SPacked         []int64  `protobuf:"varint,7,rep,packed,name=int64s_packed,json=int64sPacked" json:"int64s_packed,omitempty"`
	Strings              []string `protobuf:"bytes,5,rep,name=strings" json:"strings,omitempty"`
	Fixeds               []uint32 `protobuf:"fixed32,6,rep,name=fixeds" json:"fixeds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoreRepeated) Reset()         { *m = MoreRepeated{} }
func (m *MoreRepeated) String() string { return proto.CompactTextString(m) }
func (*MoreRepeated) ProtoMessage()    {}
func (*MoreRepeated) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{23}
}

func (m *MoreRepeated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoreRepeated.Unmarshal(m, b)
}
func (m *MoreRepeated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoreRepeated.Marshal(b, m, deterministic)
}
func (m *MoreRepeated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoreRepeated.Merge(m, src)
}
func (m *MoreRepeated) XXX_Size() int {
	return xxx_messageInfo_MoreRepeated.Size(m)
}
func (m *MoreRepeated) XXX_DiscardUnknown() {
	xxx_messageInfo_MoreRepeated.DiscardUnknown(m)
}

var xxx_messageInfo_MoreRepeated proto.InternalMessageInfo

func (m *MoreRepeated) GetBools() []bool {
	if m != nil {
		return m.Bools
	}
	return nil
}

func (m *MoreRepeated) GetBoolsPacked() []bool {
	if m != nil {
		return m.BoolsPacked
	}
	return nil
}

func (m *MoreRepeated) GetInts() []int32 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *MoreRepeated) GetIntsPacked() []int32 {
	if m != nil {
		return m.IntsPacked
	}
	return nil
}

func (m *MoreRepeated) GetInt64SPacked() []int64 {
	if m != nil {
		return m.Int64SPacked
	}
	return nil
}

func (m *MoreRepeated) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *MoreRepeated) GetFixeds() []uint32 {
	if m != nil {
		return m.Fixeds
	}
	return nil
}

type GroupOld struct {
	G                    *GroupOld_G `protobuf:"group,101,opt,name=G,json=g" json:"g,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GroupOld) Reset()         { *m = GroupOld{} }
func (m *GroupOld) String() string { return proto.CompactTextString(m) }
func (*GroupOld) ProtoMessage()    {}
func (*GroupOld) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{24}
}

func (m *GroupOld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupOld.Unmarshal(m, b)
}
func (m *GroupOld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupOld.Marshal(b, m, deterministic)
}
func (m *GroupOld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupOld.Merge(m, src)
}
func (m *GroupOld) XXX_Size() int {
	return xxx_messageInfo_GroupOld.Size(m)
}
func (m *GroupOld) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupOld.DiscardUnknown(m)
}

var xxx_messageInfo_GroupOld proto.InternalMessageInfo

func (m *GroupOld) GetG() *GroupOld_G {
	if m != nil {
		return m.G
	}
	return nil
}

type GroupOld_G struct {
	X                    *int32   `protobuf:"varint,2,opt,name=x" json:"x,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupOld_G) Reset()         { *m = GroupOld_G{} }
func (m *GroupOld_G) String() string { return proto.CompactTextString(m) }
func (*GroupOld_G) ProtoMessage()    {}
func (*GroupOld_G) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{24, 0}
}

func (m *GroupOld_G) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupOld_G.Unmarshal(m, b)
}
func (m *GroupOld_G) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupOld_G.Marshal(b, m, deterministic)
}
func (m *GroupOld_G) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupOld_G.Merge(m, src)
}
func (m *GroupOld_G) XXX_Size() int {
	return xxx_messageInfo_GroupOld_G.Size(m)
}
func (m *GroupOld_G) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupOld_G.DiscardUnknown(m)
}

var xxx_messageInfo_GroupOld_G proto.InternalMessageInfo

func (m *GroupOld_G) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

type GroupNew struct {
	G                    *GroupNew_G `protobuf:"group,101,opt,name=G,json=g" json:"g,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GroupNew) Reset()         { *m = GroupNew{} }
func (m *GroupNew) String() string { return proto.CompactTextString(m) }
func (*GroupNew) ProtoMessage()    {}
func (*GroupNew) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{25}
}

func (m *GroupNew) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupNew.Unmarshal(m, b)
}
func (m *GroupNew) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupNew.Marshal(b, m, deterministic)
}
func (m *GroupNew) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupNew.Merge(m, src)
}
func (m *GroupNew) XXX_Size() int {
	return xxx_messageInfo_GroupNew.Size(m)
}
func (m *GroupNew) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupNew.DiscardUnknown(m)
}

var xxx_messageInfo_GroupNew proto.InternalMessageInfo

func (m *GroupNew) GetG() *GroupNew_G {
	if m != nil {
		return m.G
	}
	return nil
}

type GroupNew_G struct {
	X                    *int32   `protobuf:"varint,2,opt,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,3,opt,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupNew_G) Reset()         { *m = GroupNew_G{} }
func (m *GroupNew_G) String() string { return proto.CompactTextString(m) }
func (*GroupNew_G) ProtoMessage()    {}
func (*GroupNew_G) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{25, 0}
}

func (m *GroupNew_G) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupNew_G.Unmarshal(m, b)
}
func (m *GroupNew_G) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupNew_G.Marshal(b, m, deterministic)
}
func (m *GroupNew_G) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupNew_G.Merge(m, src)
}
func (m *GroupNew_G) XXX_Size() int {
	return xxx_messageInfo_GroupNew_G.Size(m)
}
func (m *GroupNew_G) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupNew_G.DiscardUnknown(m)
}

var xxx_messageInfo_GroupNew_G proto.InternalMessageInfo

func (m *GroupNew_G) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *GroupNew_G) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type FloatingPoint struct {
	F                    *float64 `protobuf:"fixed64,1,req,name=f" json:"f,omitempty"`
	Exact                *bool    `protobuf:"varint,2,opt,name=exact" json:"exact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatingPoint) Reset()         { *m = FloatingPoint{} }
func (m *FloatingPoint) String() string { return proto.CompactTextString(m) }
func (*FloatingPoint) ProtoMessage()    {}
func (*FloatingPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{26}
}

func (m *FloatingPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatingPoint.Unmarshal(m, b)
}
func (m *FloatingPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatingPoint.Marshal(b, m, deterministic)
}
func (m *FloatingPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatingPoint.Merge(m, src)
}
func (m *FloatingPoint) XXX_Size() int {
	return xxx_messageInfo_FloatingPoint.Size(m)
}
func (m *FloatingPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatingPoint.DiscardUnknown(m)
}

var xxx_messageInfo_FloatingPoint proto.InternalMessageInfo

func (m *FloatingPoint) GetF() float64 {
	if m != nil && m.F != nil {
		return *m.F
	}
	return 0
}

func (m *FloatingPoint) GetExact() bool {
	if m != nil && m.Exact != nil {
		return *m.Exact
	}
	return false
}

type MessageWithMap struct {
	NameMapping          map[int32]string         `protobuf:"bytes,1,rep,name=name_mapping,json=nameMapping" json:"name_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MsgMapping           map[int64]*FloatingPoint `protobuf:"bytes,2,rep,name=msg_mapping,json=msgMapping" json:"msg_mapping,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ByteMapping          map[bool][]byte          `protobuf:"bytes,3,rep,name=byte_mapping,json=byteMapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	StrToStr             map[string]string        `protobuf:"bytes,4,rep,name=str_to_str,json=strToStr" json:"str_to_str,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{27}
}

func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (m *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(m, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetNameMapping() map[int32]string {
	if m != nil {
		return m.NameMapping
	}
	return nil
}

func (m *MessageWithMap) GetMsgMapping() map[int64]*FloatingPoint {
	if m != nil {
		return m.MsgMapping
	}
	return nil
}

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

func (m *MessageWithMap) GetStrToStr() map[string]string {
	if m != nil {
		return m.StrToStr
	}
	return nil
}

type Oneof struct {
	// Types that are valid to be assigned to Union:
	//	*Oneof_F_Bool
	//	*Oneof_F_Int32
	//	*Oneof_F_Int64
	//	*Oneof_F_Fixed32
	//	*Oneof_F_Fixed64
	//	*Oneof_F_Uint32
	//	*Oneof_F_Uint64
	//	*Oneof_F_Float
	//	*Oneof_F_Double
	//	*Oneof_F_String
	//	*Oneof_F_Bytes
	//	*Oneof_F_Sint32
	//	*Oneof_F_Sint64
	//	*Oneof_F_Enum
	//	*Oneof_F_Message
	//	*Oneof_FGroup
	//	*Oneof_F_Largest_Tag
	Union isOneof_Union `protobuf_oneof:"union"`
	// Types that are valid to be assigned to Tormato:
	//	*Oneof_Value
	Tormato              isOneof_Tormato `protobuf_oneof:"tormato"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Oneof) Reset()         { *m = Oneof{} }
func (m *Oneof) String() string { return proto.CompactTextString(m) }
func (*Oneof) ProtoMessage()    {}
func (*Oneof) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{28}
}

func (m *Oneof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Oneof.Unmarshal(m, b)
}
func (m *Oneof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Oneof.Marshal(b, m, deterministic)
}
func (m *Oneof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oneof.Merge(m, src)
}
func (m *Oneof) XXX_Size() int {
	return xxx_messageInfo_Oneof.Size(m)
}
func (m *Oneof) XXX_DiscardUnknown() {
	xxx_messageInfo_Oneof.DiscardUnknown(m)
}

var xxx_messageInfo_Oneof proto.InternalMessageInfo

type isOneof_Union interface {
	isOneof_Union()
}

type Oneof_F_Bool struct {
	F_Bool bool `protobuf:"varint,1,opt,name=F_Bool,json=FBool,oneof"`
}

type Oneof_F_Int32 struct {
	F_Int32 int32 `protobuf:"varint,2,opt,name=F_Int32,json=FInt32,oneof"`
}

type Oneof_F_Int64 struct {
	F_Int64 int64 `protobuf:"varint,3,opt,name=F_Int64,json=FInt64,oneof"`
}

type Oneof_F_Fixed32 struct {
	F_Fixed32 uint32 `protobuf:"fixed32,4,opt,name=F_Fixed32,json=FFixed32,oneof"`
}

type Oneof_F_Fixed64 struct {
	F_Fixed64 uint64 `protobuf:"fixed64,5,opt,name=F_Fixed64,json=FFixed64,oneof"`
}

type Oneof_F_Uint32 struct {
	F_Uint32 uint32 `protobuf:"varint,6,opt,name=F_Uint32,json=FUint32,oneof"`
}

type Oneof_F_Uint64 struct {
	F_Uint64 uint64 `protobuf:"varint,7,opt,name=F_Uint64,json=FUint64,oneof"`
}

type Oneof_F_Float struct {
	F_Float float32 `protobuf:"fixed32,8,opt,name=F_Float,json=FFloat,oneof"`
}

type Oneof_F_Double struct {
	F_Double float64 `protobuf:"fixed64,9,opt,name=F_Double,json=FDouble,oneof"`
}

type Oneof_F_String struct {
	F_String string `protobuf:"bytes,10,opt,name=F_String,json=FString,oneof"`
}

type Oneof_F_Bytes struct {
	F_Bytes []byte `protobuf:"bytes,11,opt,name=F_Bytes,json=FBytes,oneof"`
}

type Oneof_F_Sint32 struct {
	F_Sint32 int32 `protobuf:"zigzag32,12,opt,name=F_Sint32,json=FSint32,oneof"`
}

type Oneof_F_Sint64 struct {
	F_Sint64 int64 `protobuf:"zigzag64,13,opt,name=F_Sint64,json=FSint64,oneof"`
}

type Oneof_F_Enum struct {
	F_Enum MyMessage_Color `protobuf:"varint,14,opt,name=F_Enum,json=FEnum,enum=proto2_test.MyMessage_Color,oneof"`
}

type Oneof_F_Message struct {
	F_Message *GoTestField `protobuf:"bytes,15,opt,name=F_Message,json=FMessage,oneof"`
}

type Oneof_FGroup struct {
	FGroup *Oneof_F_Group `protobuf:"group,16,opt,name=F_Group,json=fGroup,oneof"`
}

type Oneof_F_Largest_Tag struct {
	F_Largest_Tag int32 `protobuf:"varint,536870911,opt,name=F_Largest_Tag,json=FLargestTag,oneof"`
}

func (*Oneof_F_Bool) isOneof_Union() {}

func (*Oneof_F_Int32) isOneof_Union() {}

func (*Oneof_F_Int64) isOneof_Union() {}

func (*Oneof_F_Fixed32) isOneof_Union() {}

func (*Oneof_F_Fixed64) isOneof_Union() {}

func (*Oneof_F_Uint32) isOneof_Union() {}

func (*Oneof_F_Uint64) isOneof_Union() {}

func (*Oneof_F_Float) isOneof_Union() {}

func (*Oneof_F_Double) isOneof_Union() {}

func (*Oneof_F_String) isOneof_Union() {}

func (*Oneof_F_Bytes) isOneof_Union() {}

func (*Oneof_F_Sint32) isOneof_Union() {}

func (*Oneof_F_Sint64) isOneof_Union() {}

func (*Oneof_F_Enum) isOneof_Union() {}

func (*Oneof_F_Message) isOneof_Union() {}

func (*Oneof_FGroup) isOneof_Union() {}

func (*Oneof_F_Largest_Tag) isOneof_Union() {}

func (m *Oneof) GetUnion() isOneof_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Oneof) GetF_Bool() bool {
	if x, ok := m.GetUnion().(*Oneof_F_Bool); ok {
		return x.F_Bool
	}
	return false
}

func (m *Oneof) GetF_Int32() int32 {
	if x, ok := m.GetUnion().(*Oneof_F_Int32); ok {
		return x.F_Int32
	}
	return 0
}

func (m *Oneof) GetF_Int64() int64 {
	if x, ok := m.GetUnion().(*Oneof_F_Int64); ok {
		return x.F_Int64
	}
	return 0
}

func (m *Oneof) GetF_Fixed32() uint32 {
	if x, ok := m.GetUnion().(*Oneof_F_Fixed32); ok {
		return x.F_Fixed32
	}
	return 0
}

func (m *Oneof) GetF_Fixed64() uint64 {
	if x, ok := m.GetUnion().(*Oneof_F_Fixed64); ok {
		return x.F_Fixed64
	}
	return 0
}

func (m *Oneof) GetF_Uint32() uint32 {
	if x, ok := m.GetUnion().(*Oneof_F_Uint32); ok {
		return x.F_Uint32
	}
	return 0
}

func (m *Oneof) GetF_Uint64() uint64 {
	if x, ok := m.GetUnion().(*Oneof_F_Uint64); ok {
		return x.F_Uint64
	}
	return 0
}

func (m *Oneof) GetF_Float() float32 {
	if x, ok := m.GetUnion().(*Oneof_F_Float); ok {
		return x.F_Float
	}
	return 0
}

func (m *Oneof) GetF_Double() float64 {
	if x, ok := m.GetUnion().(*Oneof_F_Double); ok {
		return x.F_Double
	}
	return 0
}

func (m *Oneof) GetF_String() string {
	if x, ok := m.GetUnion().(*Oneof_F_String); ok {
		return x.F_String
	}
	return ""
}

func (m *Oneof) GetF_Bytes() []byte {
	if x, ok := m.GetUnion().(*Oneof_F_Bytes); ok {
		return x.F_Bytes
	}
	return nil
}

func (m *Oneof) GetF_Sint32() int32 {
	if x, ok := m.GetUnion().(*Oneof_F_Sint32); ok {
		return x.F_Sint32
	}
	return 0
}

func (m *Oneof) GetF_Sint64() int64 {
	if x, ok := m.GetUnion().(*Oneof_F_Sint64); ok {
		return x.F_Sint64
	}
	return 0
}

func (m *Oneof) GetF_Enum() MyMessage_Color {
	if x, ok := m.GetUnion().(*Oneof_F_Enum); ok {
		return x.F_Enum
	}
	return MyMessage_RED
}

func (m *Oneof) GetF_Message() *GoTestField {
	if x, ok := m.GetUnion().(*Oneof_F_Message); ok {
		return x.F_Message
	}
	return nil
}

func (m *Oneof) GetFGroup() *Oneof_F_Group {
	if x, ok := m.GetUnion().(*Oneof_FGroup); ok {
		return x.FGroup
	}
	return nil
}

func (m *Oneof) GetF_Largest_Tag() int32 {
	if x, ok := m.GetUnion().(*Oneof_F_Largest_Tag); ok {
		return x.F_Largest_Tag
	}
	return 0
}

type isOneof_Tormato interface {
	isOneof_Tormato()
}

type Oneof_Value struct {
	Value int32 `protobuf:"varint,100,opt,name=value,oneof"`
}

func (*Oneof_Value) isOneof_Tormato() {}

func (m *Oneof) GetTormato() isOneof_Tormato {
	if m != nil {
		return m.Tormato
	}
	return nil
}

func (m *Oneof) GetValue() int32 {
	if x, ok := m.GetTormato().(*Oneof_Value); ok {
		return x.Value
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Oneof) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Oneof_F_Bool)(nil),
		(*Oneof_F_Int32)(nil),
		(*Oneof_F_Int64)(nil),
		(*Oneof_F_Fixed32)(nil),
		(*Oneof_F_Fixed64)(nil),
		(*Oneof_F_Uint32)(nil),
		(*Oneof_F_Uint64)(nil),
		(*Oneof_F_Float)(nil),
		(*Oneof_F_Double)(nil),
		(*Oneof_F_String)(nil),
		(*Oneof_F_Bytes)(nil),
		(*Oneof_F_Sint32)(nil),
		(*Oneof_F_Sint64)(nil),
		(*Oneof_F_Enum)(nil),
		(*Oneof_F_Message)(nil),
		(*Oneof_FGroup)(nil),
		(*Oneof_F_Largest_Tag)(nil),
		(*Oneof_Value)(nil),
	}
}

type Oneof_F_Group struct {
	X                    *int32   `protobuf:"varint,17,opt,name=x" json:"x,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Oneof_F_Group) Reset()         { *m = Oneof_F_Group{} }
func (m *Oneof_F_Group) String() string { return proto.CompactTextString(m) }
func (*Oneof_F_Group) ProtoMessage()    {}
func (*Oneof_F_Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{28, 0}
}

func (m *Oneof_F_Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Oneof_F_Group.Unmarshal(m, b)
}
func (m *Oneof_F_Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Oneof_F_Group.Marshal(b, m, deterministic)
}
func (m *Oneof_F_Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oneof_F_Group.Merge(m, src)
}
func (m *Oneof_F_Group) XXX_Size() int {
	return xxx_messageInfo_Oneof_F_Group.Size(m)
}
func (m *Oneof_F_Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Oneof_F_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Oneof_F_Group proto.InternalMessageInfo

func (m *Oneof_F_Group) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

type Communique struct {
	MakeMeCry *bool `protobuf:"varint,1,opt,name=make_me_cry,json=makeMeCry" json:"make_me_cry,omitempty"`
	// This is a oneof, called "union".
	//
	// Types that are valid to be assigned to Union:
	//	*Communique_Number
	//	*Communique_Name
	//	*Communique_Data
	//	*Communique_TempC
	//	*Communique_Col
	//	*Communique_Msg
	Union                isCommunique_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Communique) Reset()         { *m = Communique{} }
func (m *Communique) String() string { return proto.CompactTextString(m) }
func (*Communique) ProtoMessage()    {}
func (*Communique) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{29}
}

func (m *Communique) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique.Unmarshal(m, b)
}
func (m *Communique) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique.Marshal(b, m, deterministic)
}
func (m *Communique) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique.Merge(m, src)
}
func (m *Communique) XXX_Size() int {
	return xxx_messageInfo_Communique.Size(m)
}
func (m *Communique) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique.DiscardUnknown(m)
}

var xxx_messageInfo_Communique proto.InternalMessageInfo

func (m *Communique) GetMakeMeCry() bool {
	if m != nil && m.MakeMeCry != nil {
		return *m.MakeMeCry
	}
	return false
}

type isCommunique_Union interface {
	isCommunique_Union()
}

type Communique_Number struct {
	Number int32 `protobuf:"varint,5,opt,name=number,oneof"`
}

type Communique_Name struct {
	Name string `protobuf:"bytes,6,opt,name=name,oneof"`
}

type Communique_Data struct {
	Data []byte `protobuf:"bytes,7,opt,name=data,oneof"`
}

type Communique_TempC struct {
	TempC float64 `protobuf:"fixed64,8,opt,name=temp_c,json=tempC,oneof"`
}

type Communique_Col struct {
	Col MyMessage_Color `protobuf:"varint,9,opt,name=col,enum=proto2_test.MyMessage_Color,oneof"`
}

type Communique_Msg struct {
	Msg *Strings `protobuf:"bytes,10,opt,name=msg,oneof"`
}

func (*Communique_Number) isCommunique_Union() {}

func (*Communique_Name) isCommunique_Union() {}

func (*Communique_Data) isCommunique_Union() {}

func (*Communique_TempC) isCommunique_Union() {}

func (*Communique_Col) isCommunique_Union() {}

func (*Communique_Msg) isCommunique_Union() {}

func (m *Communique) GetUnion() isCommunique_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Communique) GetNumber() int32 {
	if x, ok := m.GetUnion().(*Communique_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Communique) GetName() string {
	if x, ok := m.GetUnion().(*Communique_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Communique) GetData() []byte {
	if x, ok := m.GetUnion().(*Communique_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Communique) GetTempC() float64 {
	if x, ok := m.GetUnion().(*Communique_TempC); ok {
		return x.TempC
	}
	return 0
}

func (m *Communique) GetCol() MyMessage_Color {
	if x, ok := m.GetUnion().(*Communique_Col); ok {
		return x.Col
	}
	return MyMessage_RED
}

func (m *Communique) GetMsg() *Strings {
	if x, ok := m.GetUnion().(*Communique_Msg); ok {
		return x.Msg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Communique) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Communique_Number)(nil),
		(*Communique_Name)(nil),
		(*Communique_Data)(nil),
		(*Communique_TempC)(nil),
		(*Communique_Col)(nil),
		(*Communique_Msg)(nil),
	}
}

type TestUTF8 struct {
	Scalar *string  `protobuf:"bytes,1,opt,name=scalar" json:"scalar,omitempty"`
	Vector []string `protobuf:"bytes,2,rep,name=vector" json:"vector,omitempty"`
	// Types that are valid to be assigned to Oneof:
	//	*TestUTF8_Field
	Oneof                isTestUTF8_Oneof `protobuf_oneof:"oneof"`
	MapKey               map[string]int64 `protobuf:"bytes,4,rep,name=map_key,json=mapKey" json:"map_key,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	MapValue             map[int64]string `protobuf:"bytes,5,rep,name=map_value,json=mapValue" json:"map_value,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestUTF8) Reset()         { *m = TestUTF8{} }
func (m *TestUTF8) String() string { return proto.CompactTextString(m) }
func (*TestUTF8) ProtoMessage()    {}
func (*TestUTF8) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b3e7ca68f98362, []int{30}
}

func (m *TestUTF8) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUTF8.Unmarshal(m, b)
}
func (m *TestUTF8) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUTF8.Marshal(b, m, deterministic)
}
func (m *TestUTF8) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUTF8.Merge(m, src)
}
func (m *TestUTF8) XXX_Size() int {
	return xxx_messageInfo_TestUTF8.Size(m)
}
func (m *TestUTF8) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUTF8.DiscardUnknown(m)
}

var xxx_messageInfo_TestUTF8 proto.InternalMessageInfo

func (m *TestUTF8) GetScalar() string {
	if m != nil && m.Scalar != nil {
		return *m.Scalar
	}
	return ""
}

func (m *TestUTF8) GetVector() []string {
	if m != nil {
		return m.Vector
	}
	return nil
}

type isTestUTF8_Oneof interface {
	isTestUTF8_Oneof()
}

type TestUTF8_Field struct {
	Field string `protobuf:"bytes,3,opt,name=field,oneof"`
}

func (*TestUTF8_Field) isTestUTF8_Oneof() {}

func (m *TestUTF8) GetOneof() isTestUTF8_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (m *TestUTF8) GetField() string {
	if x, ok := m.GetOneof().(*TestUTF8_Field); ok {
		return x.Field
	}
	return ""
}

func (m *TestUTF8) GetMapKey() map[string]int64 {
	if m != nil {
		return m.MapKey
	}
	return nil
}

func (m *TestUTF8) GetMapValue() map[int64]string {
	if m != nil {
		return m.MapValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestUTF8) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestUTF8_Field)(nil),
	}
}

var E_Greeting = &proto.ExtensionDesc{
	ExtendedType:  (*MyMessage)(nil),
	ExtensionType: ([]string)(nil),
	Field:         106,
	Name:          "proto2_test.greeting",
	Tag:           "bytes,106,rep,name=greeting",
	Filename:      "proto2_proto/test.proto",
}

var E_Complex = &proto.ExtensionDesc{
	ExtendedType:  (*OtherMessage)(nil),
	ExtensionType: (*ComplexExtension)(nil),
	Field:         200,
	Name:          "proto2_test.complex",
	Tag:           "bytes,200,opt,name=complex",
	Filename:      "proto2_proto/test.proto",
}

var E_RComplex = &proto.ExtensionDesc{
	ExtendedType:  (*OtherMessage)(nil),
	ExtensionType: ([]*ComplexExtension)(nil),
	Field:         201,
	Name:          "proto2_test.r_complex",
	Tag:           "bytes,201,rep,name=r_complex",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultDouble = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "proto2_test.no_default_double",
	Tag:           "fixed64,101,opt,name=no_default_double",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultFloat = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*float32)(nil),
	Field:         102,
	Name:          "proto2_test.no_default_float",
	Tag:           "fixed32,102,opt,name=no_default_float",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultInt32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         103,
	Name:          "proto2_test.no_default_int32",
	Tag:           "varint,103,opt,name=no_default_int32",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultInt64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         104,
	Name:          "proto2_test.no_default_int64",
	Tag:           "varint,104,opt,name=no_default_int64",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultUint32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         105,
	Name:          "proto2_test.no_default_uint32",
	Tag:           "varint,105,opt,name=no_default_uint32",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultUint64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         106,
	Name:          "proto2_test.no_default_uint64",
	Tag:           "varint,106,opt,name=no_default_uint64",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultSint32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         107,
	Name:          "proto2_test.no_default_sint32",
	Tag:           "zigzag32,107,opt,name=no_default_sint32",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultSint64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         108,
	Name:          "proto2_test.no_default_sint64",
	Tag:           "zigzag64,108,opt,name=no_default_sint64",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultFixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         109,
	Name:          "proto2_test.no_default_fixed32",
	Tag:           "fixed32,109,opt,name=no_default_fixed32",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultFixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         110,
	Name:          "proto2_test.no_default_fixed64",
	Tag:           "fixed64,110,opt,name=no_default_fixed64",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultSfixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         111,
	Name:          "proto2_test.no_default_sfixed32",
	Tag:           "fixed32,111,opt,name=no_default_sfixed32",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultSfixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         112,
	Name:          "proto2_test.no_default_sfixed64",
	Tag:           "fixed64,112,opt,name=no_default_sfixed64",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultBool = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*bool)(nil),
	Field:         113,
	Name:          "proto2_test.no_default_bool",
	Tag:           "varint,113,opt,name=no_default_bool",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultString = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         114,
	Name:          "proto2_test.no_default_string",
	Tag:           "bytes,114,opt,name=no_default_string",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultBytes = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: ([]byte)(nil),
	Field:         115,
	Name:          "proto2_test.no_default_bytes",
	Tag:           "bytes,115,opt,name=no_default_bytes",
	Filename:      "proto2_proto/test.proto",
}

var E_NoDefaultEnum = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*DefaultsMessage_DefaultsEnum)(nil),
	Field:         116,
	Name:          "proto2_test.no_default_enum",
	Tag:           "varint,116,opt,name=no_default_enum,enum=proto2_test.DefaultsMessage_DefaultsEnum",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultDouble = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*float64)(nil),
	Field:         201,
	Name:          "proto2_test.default_double",
	Tag:           "fixed64,201,opt,name=default_double,def=3.1415",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultFloat = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*float32)(nil),
	Field:         202,
	Name:          "proto2_test.default_float",
	Tag:           "fixed32,202,opt,name=default_float,def=3.14",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultInt32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         203,
	Name:          "proto2_test.default_int32",
	Tag:           "varint,203,opt,name=default_int32,def=42",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultInt64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         204,
	Name:          "proto2_test.default_int64",
	Tag:           "varint,204,opt,name=default_int64,def=43",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultUint32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         205,
	Name:          "proto2_test.default_uint32",
	Tag:           "varint,205,opt,name=default_uint32,def=44",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultUint64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         206,
	Name:          "proto2_test.default_uint64",
	Tag:           "varint,206,opt,name=default_uint64,def=45",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultSint32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         207,
	Name:          "proto2_test.default_sint32",
	Tag:           "zigzag32,207,opt,name=default_sint32,def=46",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultSint64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         208,
	Name:          "proto2_test.default_sint64",
	Tag:           "zigzag64,208,opt,name=default_sint64,def=47",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultFixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         209,
	Name:          "proto2_test.default_fixed32",
	Tag:           "fixed32,209,opt,name=default_fixed32,def=48",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultFixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         210,
	Name:          "proto2_test.default_fixed64",
	Tag:           "fixed64,210,opt,name=default_fixed64,def=49",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultSfixed32 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         211,
	Name:          "proto2_test.default_sfixed32",
	Tag:           "fixed32,211,opt,name=default_sfixed32,def=50",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultSfixed64 = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         212,
	Name:          "proto2_test.default_sfixed64",
	Tag:           "fixed64,212,opt,name=default_sfixed64,def=51",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultBool = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*bool)(nil),
	Field:         213,
	Name:          "proto2_test.default_bool",
	Tag:           "varint,213,opt,name=default_bool,def=1",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultString = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         214,
	Name:          "proto2_test.default_string",
	Tag:           "bytes,214,opt,name=default_string,def=Hello, string,def=foo",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultBytes = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: ([]byte)(nil),
	Field:         215,
	Name:          "proto2_test.default_bytes",
	Tag:           "bytes,215,opt,name=default_bytes,def=Hello, bytes",
	Filename:      "proto2_proto/test.proto",
}

var E_DefaultEnum = &proto.ExtensionDesc{
	ExtendedType:  (*DefaultsMessage)(nil),
	ExtensionType: (*DefaultsMessage_DefaultsEnum)(nil),
	Field:         216,
	Name:          "proto2_test.default_enum",
	Tag:           "varint,216,opt,name=default_enum,enum=proto2_test.DefaultsMessage_DefaultsEnum,def=1",
	Filename:      "proto2_proto/test.proto",
}

func init() {
	proto.RegisterEnum("proto2_test.FOO", FOO_name, FOO_value)
	proto.RegisterEnum("proto2_test.GoTest_KIND", GoTest_KIND_name, GoTest_KIND_value)
	proto.RegisterEnum("proto2_test.MyMessage_Color", MyMessage_Color_name, MyMessage_Color_value)
	proto.RegisterEnum("proto2_test.DefaultsMessage_DefaultsEnum", DefaultsMessage_DefaultsEnum_name, DefaultsMessage_DefaultsEnum_value)
	proto.RegisterEnum("proto2_test.Defaults_Color", Defaults_Color_name, Defaults_Color_value)
	proto.RegisterEnum("proto2_test.RepeatedEnum_Color", RepeatedEnum_Color_name, RepeatedEnum_Color_value)
	proto.RegisterType((*GoEnum)(nil), "proto2_test.GoEnum")
	proto.RegisterType((*GoTestField)(nil), "proto2_test.GoTestField")
	proto.RegisterType((*GoTest)(nil), "proto2_test.GoTest")
	proto.RegisterType((*GoTest_RequiredGroup)(nil), "proto2_test.GoTest.RequiredGroup")
	proto.RegisterType((*GoTest_RepeatedGroup)(nil), "proto2_test.GoTest.RepeatedGroup")
	proto.RegisterType((*GoTest_OptionalGroup)(nil), "proto2_test.GoTest.OptionalGroup")
	proto.RegisterType((*GoTestRequiredGroupField)(nil), "proto2_test.GoTestRequiredGroupField")
	proto.RegisterType((*GoTestRequiredGroupField_Group)(nil), "proto2_test.GoTestRequiredGroupField.Group")
	proto.RegisterType((*GoSkipTest)(nil), "proto2_test.GoSkipTest")
	proto.RegisterType((*GoSkipTest_SkipGroup)(nil), "proto2_test.GoSkipTest.SkipGroup")
	proto.RegisterType((*NonPackedTest)(nil), "proto2_test.NonPackedTest")
	proto.RegisterType((*PackedTest)(nil), "proto2_test.PackedTest")
	proto.RegisterType((*MaxTag)(nil), "proto2_test.MaxTag")
	proto.RegisterType((*OldMessage)(nil), "proto2_test.OldMessage")
	proto.RegisterType((*OldMessage_Nested)(nil), "proto2_test.OldMessage.Nested")
	proto.RegisterType((*NewMessage)(nil), "proto2_test.NewMessage")
	proto.RegisterType((*NewMessage_Nested)(nil), "proto2_test.NewMessage.Nested")
	proto.RegisterType((*InnerMessage)(nil), "proto2_test.InnerMessage")
	proto.RegisterType((*OtherMessage)(nil), "proto2_test.OtherMessage")
	proto.RegisterType((*RequiredInnerMessage)(nil), "proto2_test.RequiredInnerMessage")
	proto.RegisterType((*MyMessage)(nil), "proto2_test.MyMessage")
	proto.RegisterType((*MyMessage_SomeGroup)(nil), "proto2_test.MyMessage.SomeGroup")
	proto.RegisterExtension(E_Ext_More)
	proto.RegisterExtension(E_Ext_Text)
	proto.RegisterExtension(E_Ext_Number)
	proto.RegisterType((*Ext)(nil), "proto2_test.Ext")
	proto.RegisterMapType((map[int32]int32)(nil), "proto2_test.Ext.MapFieldEntry")
	proto.RegisterType((*ComplexExtension)(nil), "proto2_test.ComplexExtension")
	proto.RegisterType((*DefaultsMessage)(nil), "proto2_test.DefaultsMessage")
	proto.RegisterType((*Empty)(nil), "proto2_test.Empty")
	proto.RegisterType((*MessageList)(nil), "proto2_test.MessageList")
	proto.RegisterType((*MessageList_Message)(nil), "proto2_test.MessageList.Message")
	proto.RegisterType((*Strings)(nil), "proto2_test.Strings")
	proto.RegisterType((*Defaults)(nil), "proto2_test.Defaults")
	proto.RegisterType((*SubDefaults)(nil), "proto2_test.SubDefaults")
	proto.RegisterType((*RepeatedEnum)(nil), "proto2_test.RepeatedEnum")
	proto.RegisterType((*MoreRepeated)(nil), "proto2_test.MoreRepeated")
	proto.RegisterType((*GroupOld)(nil), "proto2_test.GroupOld")
	proto.RegisterType((*GroupOld_G)(nil), "proto2_test.GroupOld.G")
	proto.RegisterType((*GroupNew)(nil), "proto2_test.GroupNew")
	proto.RegisterType((*GroupNew_G)(nil), "proto2_test.GroupNew.G")
	proto.RegisterType((*FloatingPoint)(nil), "proto2_test.FloatingPoint")
	proto.RegisterType((*MessageWithMap)(nil), "proto2_test.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto2_test.MessageWithMap.ByteMappingEntry")
	proto.RegisterMapType((map[int64]*FloatingPoint)(nil), "proto2_test.MessageWithMap.MsgMappingEntry")
	proto.RegisterMapType((map[int32]string)(nil), "proto2_test.MessageWithMap.NameMappingEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto2_test.MessageWithMap.StrToStrEntry")
	proto.RegisterType((*Oneof)(nil), "proto2_test.Oneof")
	proto.RegisterType((*Oneof_F_Group)(nil), "proto2_test.Oneof.F_Group")
	proto.RegisterType((*Communique)(nil), "proto2_test.Communique")
	proto.RegisterType((*TestUTF8)(nil), "proto2_test.TestUTF8")
	proto.RegisterMapType((map[string]int64)(nil), "proto2_test.TestUTF8.MapKeyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "proto2_test.TestUTF8.MapValueEntry")
	proto.RegisterExtension(E_Greeting)
	proto.RegisterExtension(E_Complex)
	proto.RegisterExtension(E_RComplex)
	proto.RegisterExtension(E_NoDefaultDouble)
	proto.RegisterExtension(E_NoDefaultFloat)
	proto.RegisterExtension(E_NoDefaultInt32)
	proto.RegisterExtension(E_NoDefaultInt64)
	proto.RegisterExtension(E_NoDefaultUint32)
	proto.RegisterExtension(E_NoDefaultUint64)
	proto.RegisterExtension(E_NoDefaultSint32)
	proto.RegisterExtension(E_NoDefaultSint64)
	proto.RegisterExtension(E_NoDefaultFixed32)
	proto.RegisterExtension(E_NoDefaultFixed64)
	proto.RegisterExtension(E_NoDefaultSfixed32)
	proto.RegisterExtension(E_NoDefaultSfixed64)
	proto.RegisterExtension(E_NoDefaultBool)
	proto.RegisterExtension(E_NoDefaultString)
	proto.RegisterExtension(E_NoDefaultBytes)
	proto.RegisterExtension(E_NoDefaultEnum)
	proto.RegisterExtension(E_DefaultDouble)
	proto.RegisterExtension(E_DefaultFloat)
	proto.RegisterExtension(E_DefaultInt32)
	proto.RegisterExtension(E_DefaultInt64)
	proto.RegisterExtension(E_DefaultUint32)
	proto.RegisterExtension(E_DefaultUint64)
	proto.RegisterExtension(E_DefaultSint32)
	proto.RegisterExtension(E_DefaultSint64)
	proto.RegisterExtension(E_DefaultFixed32)
	proto.RegisterExtension(E_DefaultFixed64)
	proto.RegisterExtension(E_DefaultSfixed32)
	proto.RegisterExtension(E_DefaultSfixed64)
	proto.RegisterExtension(E_DefaultBool)
	proto.RegisterExtension(E_DefaultString)
	proto.RegisterExtension(E_DefaultBytes)
	proto.RegisterExtension(E_DefaultEnum)
}

func init() { proto.RegisterFile("proto2_proto/test.proto", fileDescriptor_e5b3e7ca68f98362) }

var fileDescriptor_e5b3e7ca68f98362 = []byte{
	// 4330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x5b, 0x4b, 0x73, 0xdb, 0x58,
	0x76, 0x36, 0xc0, 0xf7, 0x21, 0x25, 0x42, 0xb7, 0xd5, 0x36, 0x2d, 0xb5, 0x6d, 0x98, 0x3d, 0x3d,
	0xc3, 0xb6, 0xdd, 0xb4, 0x4c, 0x51, 0xb4, 0x4d, 0x4f, 0x3b, 0x63, 0xd9, 0xa2, 0xac, 0xb4, 0x24,
	0x7a, 0x20, 0xb9, 0xbb, 0xda, 0xb3, 0x60, 0x41, 0x22, 0x48, 0x71, 0x4c, 0x02, 0x6c, 0x12, 0x1c,
	0x4b, 0xa9, 0x54, 0xaa, 0xb7, 0xd9, 0xa4, 0x2a, 0x99, 0xa4, 0x2a, 0x3f, 0x20, 0xdb, 0xc9, 0x63,
	0x97, 0x45, 0x7e, 0x40, 0x7a, 0x1e, 0x49, 0x26, 0xef, 0xac, 0x52, 0xf9, 0x07, 0x59, 0xe5, 0xb1,
	0xea, 0xa9, 0x73, 0xee, 0x05, 0x70, 0x01, 0x52, 0xaf, 0x95, 0x70, 0xef, 0xfd, 0xbe, 0x73, 0x5f,
	0x1f, 0xce, 0x39, 0xf7, 0x12, 0x82, 0x6b, 0xc3, 0x91, 0xe3, 0x3a, 0x95, 0x16, 0xfd, 0xb9, 0xef,
	0x5a, 0x63, 0xb7, 0x4c, 0x8f, 0x2c, 0x2b, 0x1a, 0xb0, 0xaa, 0x78, 0x0f, 0x92, 0x9b, 0xce, 0x86,
	0x3d, 0x19, 0xb0, 0x22, 0xc4, 0x3a, 0x8e, 0x53, 0x50, 0x74, 0xb5, 0x34, 0x5f, 0xd1, 0xca, 0x12,
	0xa8, 0xdc, 0x68, 0x36, 0x0d, 0x6c, 0x2c, 0x3e, 0x84, 0xec, 0xa6, 0xb3, 0x6f, 0x8d, 0xdd, 0x46,
	0xcf, 0xea, 0xb7, 0xd9, 0x22, 0x24, 0xb6, 0xcd, 0x03, 0xab, 0x4f, 0xa4, 0x8c, 0xc1, 0x0b, 0x8c,
	0x41, 0x7c, 0xff, 0x64, 0x68, 0x15, 0x54, 0xaa, 0xa4, 0xe7, 0xe2, 0xdf, 0x14, 0xb1, 0x1f, 0x64,
	0xb2, 0x7b, 0x10, 0xff, 0xac, 0x67, 0xb7, 0x45, 0x47, 0x85, 0x50, 0x47, 0x1c, 0x52, 0xfe, 0x6c,
	0x6b, 0xf7, 0x85, 0x41, 0x28, 0xec, 0x62, 0xdf, 0x3c, 0xe8, 0xa3, 0x35, 0x05, 0xbb, 0xa0, 0x02,
	0xd6, 0xbe, 0x32, 0x47, 0xe6, 0xa0, 0x10, 0xd3, 0x95, 0x52, 0xc2, 0xe0, 0x05, 0xf6, 0x14, 0xe6,
	0x0c, 0xeb, 0xab, 0x49, 0x6f, 0x64, 0xb5, 0x69, 0x7c, 0x85, 0xb8, 0xae, 0x96, 0xb2, 0x33, 0xbb,
	0xa0, 0x76, 0x23, 0x0c, 0xe7, 0xfc, 0xa1, 0x65, 0xba, 0x1e, 0x3f, 0xa1, 0xc7, 0xce, 0xe3, 0x4b,
	0x70, 0xe4, 0x37, 0x87, 0x6e, 0xcf, 0xb1, 0xcd, 0x3e, 0xe7, 0x27, 0x75, 0xe5, 0x6c, 0x7e, 0x08,
	0xce, 0xbe, 0x0b, 0xf9, 0x46, 0x6b, 0xdd, 0x71, 0xfa, 0xad, 0x91, 0x18, 0x57, 0x01, 0x74, 0xb5,
	0x94, 0x36, 0xe6, 0x1a, 0x58, 0xeb, 0x0d, 0x96, 0x95, 0x40, 0x6b, 0xb4, 0xb6, 0x6c, 0x77, 0xb5,
	0x12, 0x00, 0xb3, 0xba, 0x5a, 0x4a, 0x18, 0xf3, 0x0d, 0xaa, 0x9e, 0x42, 0xd6, 0xaa, 0x01, 0x32,
	0xa7, 0xab, 0xa5, 0x18, 0x47, 0xd6, 0xaa, 0x3e, 0xf2, 0x1e, 0xb0, 0x46, 0xab, 0xd1, 0x3b, 0xb6,
	0xda, 0xb2, 0xd5, 0x39, 0x5d, 0x2d, 0xa5, 0x0c, 0xad, 0x21, 0x1a, 0x66, 0xa0, 0x65, 0xcb, 0xf3,
	0xba, 0x5a, 0x4a, 0x7a, 0x68, 0xc9, 0xf6, 0x1d, 0x58, 0x68, 0xb4, 0x5e, 0xf7, 0xc2, 0x03, 0xce,
	0xeb, 0x6a, 0x69, 0xce, 0xc8, 0x37, 0x78, 0xfd, 0x34, 0x56, 0x36, 0xac, 0xe9, 0x6a, 0x29, 0x2e,
	0xb0, 0x92, 0x5d, 0x9a, 0x5d, 0xa3, 0xef, 0x98, 0x6e, 0x00, 0x5d, 0xd0, 0xd5, 0x92, 0x6a, 0xcc,
	0x37, 0xa8, 0x3a, 0x6c, 0xf5, 0x85, 0x33, 0x39, 0xe8, 0x5b, 0x01, 0x94, 0xe9, 0x6a, 0x49, 0x31,
	0xf2, 0x0d, 0x5e, 0x1f, 0xc6, 0xee, 0xb9, 0xa3, 0x9e, 0xdd, 0x0d, 0xb0, 0xef, 0x91, 0x96, 0xf3,
	0x0d, 0x5e, 0x1f, 0x1e, 0xc1, 0xfa, 0x89, 0x6b, 0x8d, 0x03, 0xa8, 0xa5, 0xab, 0xa5, 0x9c, 0x31,
	0xdf, 0xa0, 0xea, 0x88, 0xd5, 0xc8, 0x1a, 0x74, 0x74, 0xb5, 0xb4, 0x80, 0x56, 0x67, 0xac, 0xc1,
	0x5e, 0x64, 0x0d, 0xba, 0xba, 0x5a, 0x62, 0x02, 0x2b, 0xad, 0x41, 0x19, 0xde, 0x6b, 0xb4, 0xf6,
	0x3a, 0xd1, 0x8d, 0x3b, 0xd2, 0xd5, 0x52, 0xde, 0x58, 0x68, 0x78, 0x2d, 0xb3, 0xf0, 0xb2, 0xf5,
	0x9e, 0xae, 0x96, 0x34, 0x1f, 0x2f, 0xd9, 0x97, 0x35, 0xc9, 0xb5, 0x5e, 0x58, 0xd4, 0x63, 0x92,
	0x26, 0x79, 0x65, 0x58, 0x93, 0x02, 0xf8, 0xbe, 0x1e, 0x93, 0x35, 0x19, 0x41, 0x52, 0xf7, 0x02,
	0x79, 0x55, 0x8f, 0xc9, 0x9a, 0x14, 0xc8, 0x88, 0x26, 0x05, 0xf6, 0x9a, 0x1e, 0x0b, 0x6b, 0x72,
	0x0a, 0x2d, 0x5b, 0x2e, 0xe8, 0xb1, 0xb0, 0x26, 0x05, 0x3a, 0xac, 0x49, 0x01, 0xbe, 0xae, 0xc7,
	0x42, 0x9a, 0x8c, 0x62, 0x65, 0xc3, 0x4b, 0x7a, 0x2c, 0xa4, 0x49, 0x79, 0x76, 0x9e, 0x26, 0x05,
	0x74, 0x59, 0x8f, 0xc9, 0x9a, 0x94, 0xad, 0xfa, 0x9a, 0x14, 0xd0, 0x0f, 0xf4, 0x58, 0x48, 0x93,
	0x32, 0xd6, 0xd7, 0xa4, 0xc0, 0xde, 0xd0, 0x63, 0x21, 0x4d, 0x0a, 0xec, 0xc7, 0xb2, 0x26, 0x05,
	0xf4, 0x1b, 0x45, 0x8f, 0xc9, 0xa2, 0x14, 0xd0, 0xbb, 0x21, 0x51, 0x0a, 0xec, 0xcf, 0x11, 0x2b,
	0xab, 0x32, 0x0a, 0x96, 0x57, 0xe1, 0x17, 0x08, 0x96, 0x65, 0x29, 0xc0, 0xf7, 0x23, 0xb2, 0x14,
	0xf0, 0x5f, 0x22, 0x3c, 0xac, 0xcb, 0x69, 0x82, 0x6c, 0xff, 0x57, 0x48, 0x08, 0x0b, 0x53, 0x10,
	0x02, 0x61, 0x3a, 0xc2, 0x89, 0x16, 0x6e, 0xea, 0x8a, 0x2f, 0x4c, 0xcf, 0xb3, 0xca, 0xc2, 0xf4,
	0x81, 0xb7, 0x28, 0x6a, 0x08, 0x61, 0x4e, 0x21, 0x6b, 0xd5, 0x00, 0xa9, 0xeb, 0x4a, 0x20, 0x4c,
	0x1f, 0x19, 0x12, 0xa6, 0x8f, 0xbd, 0xad, 0x2b, 0xb2, 0x30, 0x67, 0xa0, 0x65, 0xcb, 0x45, 0x5d,
	0x91, 0x85, 0xe9, 0xa3, 0x65, 0x61, 0xfa, 0xe0, 0x0f, 0x75, 0x45, 0x12, 0xe6, 0x34, 0x56, 0x36,
	0xfc, 0x1d, 0x5d, 0x91, 0x84, 0x19, 0x9e, 0x1d, 0x17, 0xa6, 0x0f, 0xfd, 0x48, 0x57, 0x02, 0x61,
	0x86, 0xad, 0x0a, 0x61, 0xfa, 0xd0, 0xef, 0xea, 0x8a, 0x24, 0xcc, 0x30, 0x56, 0x08, 0xd3, 0xc7,
	0x7e, 0x8f, 0x42, 0xb5, 0x27, 0x4c, 0x1f, 0x2b, 0x09, 0xd3, 0x87, 0xfe, 0x0c, 0xc3, 0xba, 0x2f,
	0x4c, 0x1f, 0x2a, 0x0b, 0xd3, 0xc7, 0xfe, 0x39, 0x62, 0x03, 0x61, 0x4e, 0x83, 0xe5, 0x55, 0xf8,
	0x0b, 0x04, 0x07, 0xc2, 0xf4, 0xc1, 0x61, 0x61, 0xfa, 0xf0, 0xbf, 0x44, 0xb8, 0x2c, 0xcc, 0x59,
	0x04, 0xd9, 0xfe, 0x5f, 0x21, 0x41, 0x16, 0xa6, 0x4f, 0x28, 0xd3, 0x34, 0x51, 0x98, 0x6d, 0xab,
	0x63, 0x4e, 0xfa, 0x28, 0xe3, 0x12, 0x2a, 0xb3, 0x1e, 0x77, 0x47, 0x13, 0x0b, 0xe7, 0xea, 0x38,
	0xfd, 0x17, 0x5e, 0x1b, 0x2b, 0xe3, 0xf0, 0xb9, 0x40, 0x03, 0xc2, 0xc7, 0xa8, 0xd0, 0xba, 0xba,
	0x5a, 0x31, 0xf2, 0x5c, 0xa5, 0xd3, 0xf8, 0x5a, 0x55, 0xc2, 0xdf, 0x41, 0x9d, 0xd6, 0xd5, 0x5a,
	0x95, 0xe3, 0x6b, 0xd5, 0x00, 0xbf, 0x8a, 0x13, 0xf0, 0xc4, 0x1a, 0x30, 0xee, 0xa2, 0x5a, 0xeb,
	0xb1, 0xd5, 0xca, 0x8a, 0xb1, 0xe0, 0x49, 0x76, 0x16, 0x29, 0xd4, 0xcd, 0x3d, 0x14, 0x6d, 0x3d,
	0x56, 0xab, 0xfa, 0x24, 0xb9, 0xa7, 0x0a, 0x0a, 0x5d, 0x48, 0x37, 0xe0, 0x7c, 0x82, 0xda, 0xad,
	0xc7, 0x57, 0x2b, 0x2b, 0x2b, 0x86, 0x26, 0x14, 0x3c, 0x83, 0x13, 0xea, 0xa7, 0x8c, 0x1a, 0xae,
	0xc7, 0x6b, 0x55, 0x9f, 0x13, 0xee, 0x67, 0xc1, 0x93, 0x72, 0x40, 0xb9, 0x8f, 0x5a, 0xae, 0x27,
	0x57, 0x1f, 0x54, 0x1f, 0xac, 0x3d, 0x36, 0xf2, 0x5c, 0xd3, 0x01, 0xa7, 0x8a, 0xfd, 0x08, 0x51,
	0x07, 0xa4, 0x15, 0x54, 0x75, 0x3d, 0x59, 0x79, 0xf8, 0xe0, 0x51, 0xe5, 0x91, 0xa1, 0x09, 0x75,
	0x07, 0xac, 0xa7, 0xc8, 0x12, 0xf2, 0x0e, 0x58, 0x0f, 0x50, 0xdf, 0x75, 0xed, 0xc8, 0xea, 0xf7,
	0x9d, 0x7b, 0x7a, 0xf1, 0x9d, 0x33, 0xea, 0xb7, 0x6f, 0x17, 0xc1, 0xd0, 0x84, 0xe2, 0xe5, 0x5e,
	0x17, 0x3c, 0xc9, 0x07, 0xf4, 0x3f, 0xc4, 0xa4, 0x35, 0x57, 0x4f, 0xad, 0xf7, 0xba, 0xb6, 0x33,
	0xb6, 0x8c, 0x3c, 0x17, 0x7f, 0x64, 0x4d, 0xf6, 0xa2, 0xeb, 0xf8, 0x47, 0x48, 0x5b, 0xa8, 0xc7,
	0x3e, 0x59, 0xad, 0x60, 0x4f, 0xb3, 0xd6, 0x71, 0x2f, 0xba, 0x8e, 0x3f, 0x45, 0x0e, 0xab, 0xc7,
	0x3e, 0xa9, 0x55, 0x05, 0x47, 0x5e, 0xc7, 0x1a, 0x2c, 0x4a, 0xef, 0x42, 0xc0, 0xfa, 0x63, 0x64,
	0xe5, 0x79, 0x4f, 0xcc, 0x7f, 0x23, 0x66, 0xf2, 0x42, 0xbd, 0xfd, 0x09, 0xf2, 0x34, 0xde, 0x1b,
	0xf3, 0x5f, 0x8c, 0x80, 0xf7, 0x10, 0xae, 0x46, 0x72, 0x89, 0xd6, 0xd0, 0x3c, 0x7c, 0x6b, 0xb5,
	0x0b, 0x15, 0x4c, 0x29, 0xd6, 0x55, 0x4d, 0x31, 0xde, 0x0b, 0xa5, 0x15, 0xaf, 0xa8, 0x99, 0x3d,
	0x86, 0x6b, 0xd1, 0xe4, 0xc2, 0x63, 0xae, 0x62, 0x8e, 0x41, 0xcc, 0xc5, 0x70, 0x9e, 0x11, 0xa1,
	0x4a, 0x41, 0xc5, 0xa3, 0x56, 0x31, 0xe9, 0x08, 0xa8, 0x41, 0x6c, 0x11, 0xd4, 0x4f, 0xe1, 0xfa,
	0x74, 0xfa, 0xe1, 0x91, 0xd7, 0x30, 0x0b, 0x21, 0xf2, 0xd5, 0x68, 0x26, 0x32, 0x45, 0x9f, 0xd1,
	0x77, 0x0d, 0xd3, 0x12, 0x99, 0x3e, 0xd5, 0xfb, 0x13, 0x28, 0x4c, 0x25, 0x28, 0x1e, 0xfb, 0x21,
	0xe6, 0x29, 0xc4, 0x7e, 0x3f, 0x92, 0xab, 0x44, 0xc9, 0x33, 0xba, 0x7e, 0x84, 0x89, 0x8b, 0x44,
	0x9e, 0xea, 0x99, 0x96, 0x2c, 0x9c, 0xc2, 0x78, 0xdc, 0xc7, 0x98, 0xc9, 0x88, 0x25, 0x0b, 0x65,
	0x33, 0x72, 0xbf, 0x91, 0x9c, 0xc6, 0xe3, 0xd6, 0x31, 0xb5, 0x11, 0xfd, 0x86, 0xd3, 0x1b, 0x41,
	0xfe, 0x3e, 0x92, 0xf7, 0x66, 0xcf, 0xf8, 0x7f, 0x62, 0x98, 0x94, 0x08, 0xf6, 0xde, 0xac, 0x29,
	0xfb, 0xec, 0x19, 0x53, 0xfe, 0x5f, 0x64, 0x33, 0x89, 0x3d, 0x35, 0xe7, 0x1f, 0xc0, 0xd2, 0x8c,
	0x7c, 0xc5, 0xe3, 0xff, 0x1f, 0xf2, 0xf3, 0xc4, 0xbf, 0x36, 0x95, 0xba, 0x4c, 0x5b, 0x98, 0x31,
	0x82, 0xff, 0x47, 0x0b, 0x5a, 0xc8, 0xc2, 0xd4, 0x18, 0x36, 0x61, 0xce, 0xcb, 0xc7, 0xbb, 0x23,
	0x67, 0x32, 0x2c, 0x34, 0x74, 0xb5, 0x04, 0x95, 0xdb, 0xb3, 0x4e, 0xc8, 0x5e, 0x7e, 0xbe, 0x89,
	0x40, 0x23, 0xcc, 0xe3, 0x86, 0xb8, 0x69, 0x6e, 0xe8, 0x95, 0x1e, 0x3b, 0xdd, 0x10, 0x07, 0xfa,
	0x86, 0x24, 0x1e, 0x1a, 0xf2, 0x02, 0x1e, 0x37, 0xf4, 0x46, 0x57, 0x4e, 0x33, 0xe4, 0xc5, 0x3f,
	0x61, 0x28, 0xc4, 0x5b, 0x5a, 0x0b, 0x4e, 0xe6, 0xd4, 0xce, 0xbe, 0x13, 0x3d, 0xaa, 0x6f, 0xd2,
	0x01, 0x2b, 0x5c, 0xc9, 0x69, 0xd2, 0xf8, 0xa6, 0x69, 0x3f, 0x3c, 0x85, 0x16, 0x1a, 0xcd, 0x34,
	0xed, 0x47, 0x33, 0x68, 0xc5, 0x3f, 0x55, 0x20, 0xfe, 0xd9, 0xd6, 0xee, 0x0b, 0x96, 0x86, 0xf8,
	0xe7, 0xcd, 0xad, 0x17, 0xda, 0x15, 0x7c, 0x5a, 0x6f, 0x36, 0xb7, 0x35, 0x85, 0x65, 0x20, 0xb1,
	0xfe, 0xe5, 0xfe, 0xc6, 0x9e, 0xa6, 0xb2, 0x3c, 0x64, 0x1b, 0x5b, 0xbb, 0x9b, 0x1b, 0xc6, 0x2b,
	0x63, 0x6b, 0x77, 0x5f, 0x8b, 0x61, 0x5b, 0x63, 0xbb, 0xf9, 0x6c, 0x5f, 0x8b, 0xb3, 0x14, 0xc4,
	0xb0, 0x2e, 0xc1, 0x00, 0x92, 0x7b, 0xfb, 0xc6, 0xd6, 0xee, 0xa6, 0x96, 0x44, 0x2b, 0xfb, 0x5b,
	0x3b, 0x1b, 0x5a, 0x0a, 0x91, 0xfb, 0xaf, 0x5f, 0x6d, 0x6f, 0x68, 0x69, 0x7c, 0x7c, 0x66, 0x18,
	0xcf, 0xbe, 0xd4, 0x32, 0x48, 0xda, 0x79, 0xf6, 0x4a, 0x03, 0x6a, 0x7e, 0xb6, 0xbe, 0xbd, 0xa1,
	0x65, 0x59, 0x0e, 0xd2, 0x8d, 0xd7, 0xbb, 0xcf, 0xf7, 0xb7, 0x9a, 0xbb, 0x5a, 0xae, 0xf8, 0xbb,
	0x50, 0xe0, 0xcb, 0x1c, 0x5a, 0x45, 0x7e, 0x6b, 0xf0, 0x0c, 0x12, 0x7c, 0x73, 0x14, 0x92, 0xcb,
	0xdd, 0x19, 0x9b, 0x33, 0xcd, 0x2a, 0xf3, 0x6d, 0xe2, 0xcc, 0xa5, 0x1b, 0x90, 0xe0, 0x0b, 0xb5,
	0x08, 0x09, 0xbe, 0x40, 0x2a, 0x5d, 0x27, 0xf0, 0x42, 0xf1, 0xcf, 0x54, 0x80, 0x4d, 0x67, 0xef,
	0x6d, 0x6f, 0x48, 0x17, 0x38, 0x37, 0x00, 0xc6, 0x6f, 0x7b, 0xc3, 0x16, 0xbd, 0x85, 0xe2, 0xe2,
	0x21, 0x83, 0x35, 0xe4, 0x7f, 0xd9, 0x6d, 0xc8, 0x51, 0xb3, 0x78, 0x4d, 0xe8, 0xbe, 0x21, 0x65,
	0x64, 0xb1, 0x4e, 0x38, 0xca, 0x30, 0xa4, 0x56, 0xa5, 0x6b, 0x86, 0xa4, 0x04, 0xa9, 0x55, 0xd9,
	0x2d, 0xa0, 0x62, 0x6b, 0x4c, 0x11, 0x95, 0xae, 0x16, 0x32, 0x06, 0xf5, 0xcb, 0x63, 0x2c, 0xfb,
	0x2d, 0xa0, 0x3e, 0xf9, 0xd4, 0xf3, 0x33, 0xdf, 0x14, 0x6f, 0xc4, 0x65, 0x7c, 0xe0, 0x13, 0x0e,
	0x38, 0x4b, 0x4d, 0xc8, 0xf8, 0xf5, 0xd8, 0x1d, 0xd5, 0x8a, 0x49, 0x69, 0x34, 0x29, 0xa0, 0x2a,
	0x7f, 0x56, 0x1c, 0x20, 0x06, 0xb4, 0x40, 0x03, 0xe2, 0x24, 0x3e, 0xa2, 0xe2, 0x0d, 0x98, 0xdb,
	0x75, 0x6c, 0xfe, 0x32, 0xd3, 0x42, 0xe5, 0x40, 0x31, 0x0b, 0x0a, 0x1d, 0x82, 0x15, 0xb3, 0x78,
	0x13, 0x40, 0x6a, 0xd3, 0x40, 0x39, 0xe0, 0x6d, 0xe4, 0x14, 0x94, 0x83, 0xe2, 0x5d, 0x48, 0xee,
	0x98, 0xc7, 0xfb, 0x66, 0x97, 0xdd, 0x06, 0xe8, 0x9b, 0x63, 0xb7, 0xd5, 0xa1, 0xad, 0xf8, 0xf6,
	0xdb, 0x6f, 0xbf, 0x55, 0x28, 0xa3, 0xce, 0x60, 0x2d, 0xdf, 0x12, 0x17, 0xa0, 0xd9, 0x6f, 0xef,
	0x58, 0xe3, 0xb1, 0xd9, 0xb5, 0x58, 0x0d, 0x92, 0xb6, 0x35, 0xc6, 0x10, 0xac, 0xd0, 0x8d, 0xd3,
	0xcd, 0xd0, 0x42, 0x04, 0xc0, 0xf2, 0x2e, 0xa1, 0x0c, 0x81, 0x66, 0x1a, 0xc4, 0xec, 0xc9, 0x80,
	0xae, 0xd6, 0x12, 0x06, 0x3e, 0x2e, 0x7d, 0x00, 0x49, 0x8e, 0x61, 0x0c, 0xe2, 0xb6, 0x39, 0xb0,
	0x0a, 0xbc, 0x6b, 0x7a, 0x2e, 0xfe, 0x54, 0x01, 0xd8, 0xb5, 0xde, 0x5d, 0xac, 0xdb, 0x00, 0x78,
	0x46, 0xb7, 0x31, 0xde, 0xed, 0x93, 0xb3, 0xba, 0x45, 0xc1, 0x75, 0x1c, 0xa7, 0xdd, 0xe2, 0x7b,
	0xcd, 0x2f, 0x02, 0x33, 0x58, 0x43, 0x7b, 0x57, 0x7c, 0x03, 0xb9, 0x2d, 0xdb, 0xb6, 0x46, 0xde,
	0xb0, 0x18, 0xc4, 0x8f, 0x9c, 0xb1, 0x2b, 0x2e, 0x25, 0xe9, 0x99, 0x15, 0x20, 0x3e, 0x74, 0x46,
	0x2e, 0x9f, 0x6a, 0x3d, 0x5e, 0x5d, 0x59, 0x59, 0x31, 0xa8, 0x86, 0x7d, 0x00, 0x99, 0x43, 0xc7,
	0xb6, 0xad, 0x43, 0x9c, 0x47, 0x8c, 0x4e, 0x90, 0x41, 0x45, 0xf1, 0xf7, 0x15, 0xc8, 0x35, 0xdd,
	0xa3, 0xc0, 0xb8, 0x06, 0xb1, 0xb7, 0xd6, 0x09, 0x0d, 0x2f, 0x66, 0xe0, 0x23, 0xbe, 0x33, 0x3f,
	0x31, 0xfb, 0x13, 0x7e, 0x43, 0x99, 0x33, 0x78, 0x81, 0x5d, 0x85, 0xe4, 0x3b, 0xab, 0xd7, 0x3d,
	0x72, 0xc9, 0xa6, 0x6a, 0x88, 0x12, 0xbb, 0x0f, 0x89, 0x1e, 0x0e, 0xb6, 0x10, 0xa7, 0x25, 0xbb,
	0x1e, 0x5a, 0x32, 0x79, 0x1a, 0x06, 0xc7, 0xdd, 0x49, 0xa7, 0xdb, 0xda, 0xd7, 0x5f, 0x7f, 0xfd,
	0xb5, 0x5a, 0xec, 0xc1, 0xa2, 0xf7, 0x22, 0x87, 0xe6, 0xfb, 0x43, 0x28, 0xf4, 0x2d, 0xa7, 0xd5,
	0xe9, 0xd9, 0x66, 0xbf, 0x7f, 0xd2, 0x7a, 0xe7, 0xd8, 0x2d, 0xd3, 0x6e, 0x39, 0xe3, 0x43, 0x73,
	0x44, 0x6b, 0x70, 0x66, 0x2f, 0x8b, 0x7d, 0xcb, 0x69, 0x70, 0xe6, 0x17, 0x8e, 0xfd, 0xcc, 0x6e,
	0x22, 0xad, 0xf8, 0x1f, 0x71, 0xc8, 0xec, 0x9c, 0x78, 0x1d, 0x2c, 0x42, 0xe2, 0xd0, 0x99, 0xd8,
	0x7c, 0x45, 0x13, 0x06, 0x2f, 0xf8, 0x3b, 0xa5, 0x4a, 0x3b, 0xb5, 0x08, 0x89, 0xaf, 0x26, 0x8e,
	0x6b, 0xd1, 0xa4, 0x33, 0x06, 0x2f, 0xe0, 0x9a, 0x0d, 0x2d, 0xb7, 0x10, 0xa7, 0xfb, 0x0a, 0x7c,
	0x0c, 0x56, 0x21, 0x71, 0xb1, 0x55, 0x60, 0x0f, 0x20, 0xe9, 0xe0, 0x36, 0x8c, 0x0b, 0x49, 0xba,
	0x93, 0x0d, 0x33, 0xe4, 0x1d, 0x32, 0x04, 0x90, 0x6d, 0xc3, 0xc2, 0x3b, 0xab, 0x35, 0x98, 0x8c,
	0xdd, 0x56, 0xd7, 0x69, 0xb5, 0x2d, 0x6b, 0x68, 0x8d, 0x0a, 0x73, 0xd4, 0x5f, 0xd8, 0x51, 0xcc,
	0x5a, 0x54, 0x63, 0xfe, 0x9d, 0xb5, 0x33, 0x19, 0xbb, 0x9b, 0xce, 0x0b, 0x22, 0xb2, 0x1a, 0x64,
	0x46, 0x16, 0xba, 0x07, 0x1c, 0x75, 0x6e, 0xc6, 0x18, 0x42, 0xec, 0xf4, 0xc8, 0x1a, 0x52, 0x05,
	0x7b, 0x04, 0xe9, 0x83, 0xde, 0x5b, 0x6b, 0x7c, 0x64, 0xb5, 0x0b, 0x29, 0x5d, 0x29, 0xcd, 0x57,
	0x3e, 0x08, 0xd1, 0xfc, 0x55, 0x2e, 0x3f, 0x77, 0xfa, 0xce, 0xc8, 0xf0, 0xd1, 0xec, 0x29, 0x64,
	0xc6, 0xce, 0xc0, 0xe2, 0xa2, 0x4f, 0x53, 0xe0, 0xd5, 0x4f, 0xa1, 0xee, 0x39, 0x03, 0xcb, 0xf3,
	0x6f, 0x1e, 0x85, 0x2d, 0xf3, 0x11, 0x1f, 0xe0, 0xd9, 0xa2, 0x00, 0x74, 0xff, 0x83, 0xc3, 0xa2,
	0xb3, 0x06, 0x5b, 0xc2, 0x61, 0x75, 0x3b, 0x98, 0xc2, 0x15, 0xb2, 0x74, 0xb4, 0xf7, 0xcb, 0x4b,
	0xf7, 0x20, 0xe3, 0x1b, 0x0c, 0x1c, 0x23, 0x77, 0x46, 0x19, 0x72, 0x15, 0xdc, 0x31, 0x72, 0x4f,
	0xf4, 0x11, 0x24, 0x68, 0xe4, 0x18, 0xc5, 0x8c, 0x0d, 0x0c, 0x9a, 0x19, 0x48, 0x6c, 0x1a, 0x1b,
	0x1b, 0xbb, 0x9a, 0x42, 0xf1, 0x73, 0xfb, 0xf5, 0x86, 0xa6, 0x4a, 0x32, 0xfe, 0x99, 0x0a, 0xb1,
	0x8d, 0x63, 0xd2, 0x4f, 0xdb, 0x74, 0x4d, 0xef, 0x4d, 0xc7, 0x67, 0xf6, 0x04, 0x32, 0x03, 0xd3,
	0xeb, 0x4b, 0xa5, 0x55, 0x0e, 0x3b, 0x95, 0x8d, 0x63, 0xb7, 0xbc, 0x63, 0xf2, 0xae, 0x37, 0x6c,
	0x77, 0x74, 0x62, 0xa4, 0x07, 0xa2, 0xb8, 0xf4, 0x04, 0xe6, 0x42, 0x4d, 0xf2, 0xbb, 0x9a, 0x98,
	0xf1, 0xae, 0x26, 0xc4, 0xbb, 0x5a, 0x57, 0x1f, 0x29, 0x95, 0xef, 0x43, 0x7c, 0xe0, 0x8c, 0x2c,
	0x76, 0x75, 0xf6, 0x12, 0x17, 0xba, 0x24, 0x1c, 0x2d, 0x3a, 0x18, 0x83, 0x58, 0x95, 0x3b, 0x10,
	0x77, 0xad, 0x63, 0xf7, 0x54, 0xf6, 0x11, 0x9f, 0x23, 0x62, 0x2a, 0x65, 0x48, 0xda, 0x93, 0xc1,
	0x81, 0x35, 0x3a, 0x15, 0xdd, 0xa3, 0xc1, 0x09, 0x54, 0xf1, 0x73, 0xd0, 0x9e, 0x3b, 0x83, 0x61,
	0xdf, 0x3a, 0xde, 0x38, 0x76, 0x2d, 0x7b, 0xdc, 0x73, 0x6c, 0x9c, 0x47, 0xa7, 0x37, 0x22, 0x1f,
	0x47, 0xf3, 0xa0, 0x02, 0xfa, 0x9c, 0xb1, 0x75, 0xe8, 0xd8, 0x6d, 0x31, 0x3d, 0x51, 0x42, 0xb4,
	0x7b, 0xd4, 0x1b, 0xa1, 0x7b, 0xc3, 0x58, 0xc4, 0x0b, 0xc5, 0x4d, 0xc8, 0x8b, 0xa3, 0xd9, 0x58,
	0x74, 0x5c, 0xbc, 0x03, 0x39, 0xaf, 0x8a, 0x7e, 0x12, 0x4a, 0x43, 0xfc, 0xcd, 0x86, 0xd1, 0xd4,
	0xae, 0xe0, 0xe6, 0x36, 0x77, 0x37, 0x34, 0x05, 0x1f, 0xf6, 0xbf, 0x68, 0x86, 0x36, 0x34, 0x05,
	0x89, 0x8d, 0xc1, 0xd0, 0x3d, 0x29, 0xfe, 0x1e, 0x64, 0x85, 0xa5, 0xed, 0xde, 0xd8, 0x65, 0x75,
	0x48, 0x0d, 0xc4, 0x8c, 0x14, 0x4a, 0x40, 0x23, 0xf2, 0x0d, 0xa0, 0xde, 0xb3, 0xe1, 0x11, 0x96,
	0x56, 0x21, 0x25, 0xb9, 0x73, 0xe1, 0x67, 0x54, 0xd9, 0xcf, 0x70, 0x8f, 0x14, 0x93, 0x3c, 0x52,
	0x71, 0x07, 0x52, 0x3c, 0x14, 0x8f, 0x29, 0xc3, 0xe0, 0xc7, 0x76, 0xae, 0x25, 0x2e, 0xb2, 0x2c,
	0xaf, 0xe3, 0x79, 0xd3, 0x2d, 0xc8, 0xd2, 0xbb, 0xe1, 0xab, 0x0d, 0xbd, 0x37, 0x50, 0x15, 0x57,
	0xf6, 0x5f, 0x27, 0x20, 0xed, 0x2d, 0x07, 0x5b, 0x86, 0x24, 0x3f, 0xbb, 0x92, 0x29, 0xef, 0x2e,
	0x27, 0x41, 0xa7, 0x55, 0xb6, 0x0c, 0x29, 0x71, 0x3e, 0x15, 0x01, 0x46, 0x5d, 0xad, 0x18, 0x49,
	0x7e, 0x1e, 0xf5, 0x1b, 0x6b, 0x55, 0xf2, 0x8a, 0xfc, 0x96, 0x26, 0xc9, 0x4f, 0x9c, 0x4c, 0x87,
	0x8c, 0x7f, 0xc6, 0xa4, 0x90, 0x20, 0xae, 0x64, 0xd2, 0xde, 0xa1, 0x52, 0x42, 0xd4, 0xaa, 0xe4,
	0x2e, 0xc5, 0xfd, 0x4b, 0xba, 0x11, 0xa4, 0x4a, 0x69, 0xef, 0xa4, 0x48, 0xbf, 0x38, 0x79, 0x97,
	0x2d, 0x29, 0x71, 0x36, 0x0c, 0x00, 0xb5, 0x2a, 0xf9, 0x20, 0xef, 0x66, 0x25, 0x25, 0xce, 0x7f,
	0xec, 0x16, 0x0e, 0x91, 0xce, 0x73, 0xe4, 0x68, 0x82, 0x6b, 0x94, 0x24, 0x3f, 0xe5, 0xb1, 0xdb,
	0x68, 0x81, 0x1f, 0xda, 0xc8, 0x05, 0x04, 0x77, 0x26, 0x29, 0x71, 0x96, 0x63, 0x77, 0x11, 0xc2,
	0x97, 0xbf, 0x00, 0xa7, 0x5c, 0x90, 0xa4, 0xc4, 0x05, 0x09, 0xd3, 0xb1, 0x43, 0xf2, 0x44, 0xe4,
	0x7d, 0xa4, 0xcb, 0x90, 0x24, 0xbf, 0x0c, 0x61, 0x37, 0xc9, 0x1c, 0x9f, 0x54, 0x2e, 0xb8, 0xf8,
	0x48, 0x89, 0xc3, 0x5f, 0xd0, 0x4e, 0xe9, 0xa3, 0x7f, 0xc9, 0x91, 0x12, 0xc7, 0x3b, 0xf6, 0x18,
	0xf7, 0x0b, 0x45, 0x5c, 0x98, 0x27, 0xaf, 0xbb, 0x1c, 0xd2, 0x9e, 0xb7, 0xad, 0xdc, 0xe9, 0xd6,
	0xb9, 0xbf, 0x32, 0x12, 0x0d, 0x52, 0xfd, 0x12, 0x52, 0x5f, 0xf5, 0xec, 0x4e, 0x21, 0x4f, 0x8b,
	0x11, 0xeb, 0xd9, 0x1d, 0x23, 0xd1, 0xc0, 0x1a, 0x2e, 0x83, 0x5d, 0x6c, 0xd3, 0xa8, 0x2d, 0xfe,
	0x09, 0x6f, 0xc4, 0x2a, 0x56, 0x80, 0x44, 0xa3, 0xb5, 0x6b, 0xda, 0x85, 0x05, 0xce, 0xb3, 0x4d,
	0xdb, 0x88, 0x37, 0x76, 0x4d, 0x9b, 0xdd, 0x81, 0xd8, 0x78, 0x72, 0x50, 0x60, 0x33, 0x7e, 0x0f,
	0xdc, 0x9b, 0x1c, 0x78, 0xa3, 0x31, 0x10, 0xc4, 0x96, 0x21, 0x3d, 0x76, 0x47, 0xad, 0xdf, 0xb1,
	0x46, 0x4e, 0xe1, 0x3d, 0x5a, 0xc8, 0x2b, 0x46, 0x6a, 0xec, 0x8e, 0xde, 0x58, 0x23, 0xe7, 0x82,
	0xde, 0xb6, 0x78, 0x13, 0xb2, 0x92, 0x5d, 0x96, 0x07, 0xc5, 0xe6, 0x29, 0x4b, 0x5d, 0x79, 0x68,
	0x28, 0x76, 0xf1, 0x0b, 0xc8, 0x79, 0x07, 0x2b, 0x9a, 0xf2, 0x1a, 0xbe, 0x4f, 0x7d, 0x67, 0x44,
	0x2f, 0xea, 0x7c, 0xe5, 0x56, 0x24, 0x3e, 0x06, 0x48, 0x11, 0xa5, 0x38, 0xba, 0xa8, 0x45, 0x46,
	0xa3, 0x14, 0xff, 0x53, 0x81, 0xdc, 0x8e, 0x33, 0x0a, 0x7e, 0xb9, 0x58, 0x84, 0xc4, 0x81, 0xe3,
	0xf4, 0xc7, 0x64, 0x39, 0x6d, 0xf0, 0x02, 0xfb, 0x08, 0x72, 0xf4, 0xe0, 0x1d, 0x8f, 0x55, 0xff,
	0xfe, 0x27, 0x4b, 0xf5, 0xe2, 0x44, 0xcc, 0x20, 0xde, 0xb3, 0xdd, 0xb1, 0xf0, 0x5b, 0xf4, 0xcc,
	0x3e, 0x84, 0x2c, 0xfe, 0xf5, 0x98, 0x71, 0x3f, 0x85, 0x06, 0xac, 0x16, 0xc4, 0xef, 0xc1, 0x1c,
	0xc9, 0xc0, 0x87, 0xa5, 0xfc, 0xbb, 0x9e, 0x1c, 0x6f, 0x10, 0xc0, 0x02, 0xa4, 0xb8, 0x4f, 0x18,
	0xd3, 0x8f, 0xbd, 0x19, 0xc3, 0x2b, 0xa2, 0x33, 0xa5, 0xe3, 0x09, 0xcf, 0x38, 0x52, 0x86, 0x28,
	0x15, 0x5f, 0x40, 0x9a, 0x22, 0x63, 0xb3, 0xdf, 0x66, 0x1f, 0x81, 0xd2, 0x2d, 0x58, 0x14, 0x9a,
	0xaf, 0x85, 0xcf, 0x1e, 0x02, 0x51, 0xde, 0x34, 0x94, 0xee, 0xd2, 0x02, 0x28, 0x9b, 0x78, 0x18,
	0x38, 0x16, 0x7e, 0x59, 0x39, 0x2e, 0x1a, 0xc2, 0xca, 0xae, 0xf5, 0xee, 0x1c, 0x2b, 0xbb, 0xd6,
	0x3b, 0x6e, 0xe5, 0xd6, 0x94, 0x15, 0x2c, 0x9d, 0x88, 0x9f, 0xc0, 0x95, 0x93, 0xe2, 0x2a, 0xcc,
	0xd1, 0xdb, 0xda, 0xb3, 0xbb, 0xaf, 0x9c, 0x9e, 0x4d, 0xe7, 0x8f, 0x0e, 0xe5, 0x6c, 0x8a, 0xa1,
	0x74, 0x70, 0x27, 0xac, 0x63, 0xf3, 0x90, 0xe7, 0xc0, 0x69, 0x83, 0x17, 0x8a, 0xff, 0x15, 0x87,
	0x79, 0xe1, 0x69, 0xbf, 0xe8, 0xb9, 0x47, 0x3b, 0xe6, 0x90, 0x35, 0x21, 0x87, 0x4e, 0xb6, 0x35,
	0x30, 0x87, 0x43, 0x7c, 0x9b, 0x15, 0x8a, 0xc3, 0xf7, 0x66, 0x39, 0x6f, 0x41, 0x29, 0xef, 0x9a,
	0x03, 0x6b, 0x87, 0xc3, 0x79, 0x54, 0xce, 0xda, 0x41, 0x0d, 0xdb, 0x86, 0xec, 0x60, 0xdc, 0xf5,
	0xed, 0xf1, 0xb8, 0x7e, 0xf7, 0x2c, 0x7b, 0x3b, 0xe3, 0x6e, 0xc8, 0x1c, 0x0c, 0xfc, 0x0a, 0x1c,
	0x1e, 0x3a, 0x69, 0xdf, 0x5c, 0xec, 0xfc, 0xe1, 0xa1, 0x4b, 0x09, 0x0f, 0xef, 0x20, 0xa8, 0x61,
	0x9b, 0x00, 0xf8, 0xc2, 0xb9, 0x0e, 0x9e, 0xed, 0x48, 0x50, 0xd9, 0xca, 0xc7, 0x67, 0x99, 0xdb,
	0x73, 0x47, 0xfb, 0xce, 0x9e, 0x3b, 0x12, 0x09, 0xc8, 0x58, 0x14, 0x97, 0x9e, 0x82, 0x16, 0x5d,
	0x88, 0xf3, 0x72, 0x90, 0x8c, 0x94, 0x83, 0x2c, 0x7d, 0x09, 0xf9, 0xc8, 0xc4, 0x65, 0x3a, 0xe3,
	0xf4, 0x15, 0x99, 0x9e, 0xad, 0x2c, 0x85, 0x3f, 0xd4, 0x90, 0xf7, 0x5f, 0x36, 0xfd, 0x14, 0xb4,
	0xe8, 0x22, 0xc8, 0xb6, 0xd3, 0x67, 0x1c, 0x65, 0x88, 0xff, 0x04, 0xe6, 0x42, 0xb3, 0x96, 0xc9,
	0x99, 0x73, 0xe6, 0x55, 0xfc, 0x83, 0x04, 0x24, 0x9a, 0xb6, 0xe5, 0x74, 0xd8, 0xb5, 0x70, 0x14,
	0x7d, 0x79, 0xc5, 0x8b, 0xa0, 0xd7, 0x23, 0x11, 0xf4, 0xe5, 0x15, 0x3f, 0x7e, 0x5e, 0x8f, 0xc4,
	0x4f, 0xaf, 0xa9, 0x56, 0x65, 0x37, 0xa6, 0xa2, 0xe7, 0xcb, 0x2b, 0x52, 0xe8, 0xbc, 0x31, 0x15,
	0x3a, 0x83, 0xe6, 0x5a, 0x15, 0x1d, 0x6d, 0x38, 0x6e, 0xbe, 0xbc, 0x12, 0xc4, 0xcc, 0xe5, 0x68,
	0xcc, 0xf4, 0x1b, 0x6b, 0x55, 0x3e, 0x24, 0x29, 0x5e, 0xd2, 0x90, 0x78, 0xa4, 0x5c, 0x8e, 0x46,
	0x4a, 0xe2, 0x89, 0x18, 0xb9, 0x1c, 0x8d, 0x91, 0xd4, 0x28, 0x62, 0xe2, 0xf5, 0x48, 0x4c, 0x24,
	0xa3, 0x3c, 0x18, 0x2e, 0x47, 0x83, 0x21, 0xe7, 0x49, 0x23, 0x95, 0x23, 0xa1, 0xdf, 0x58, 0xab,
	0xb2, 0xb5, 0x48, 0x18, 0x3c, 0xf3, 0xf0, 0x41, 0xdb, 0x41, 0xf1, 0xe0, 0x21, 0xae, 0x9c, 0x97,
	0x8e, 0xe6, 0xcf, 0xfe, 0x8a, 0x85, 0xd6, 0xd4, 0x4b, 0xd6, 0xd6, 0x20, 0xd5, 0x11, 0xe7, 0x74,
	0x8d, 0x3c, 0x5a, 0x58, 0x9f, 0xa4, 0x82, 0x72, 0xa3, 0x45, 0x9e, 0x0d, 0x27, 0xd8, 0xe1, 0xa7,
	0x8c, 0x12, 0xcc, 0x35, 0x5a, 0xdb, 0xe6, 0xa8, 0x6b, 0x8d, 0xdd, 0xd6, 0xbe, 0xd9, 0xf5, 0x2f,
	0x3d, 0x50, 0x08, 0xd9, 0x86, 0x68, 0xd9, 0x37, 0xbb, 0xec, 0xaa, 0xa7, 0xb2, 0x36, 0xb5, 0x2a,
	0x42, 0x67, 0x4b, 0xd7, 0x70, 0xf5, 0xb8, 0x31, 0xf2, 0x91, 0x0b, 0xc2, 0x47, 0xae, 0xa7, 0x20,
	0x31, 0xb1, 0x7b, 0x8e, 0xbd, 0x9e, 0x81, 0x94, 0xeb, 0x8c, 0x06, 0xa6, 0xeb, 0x14, 0xff, 0x5b,
	0x01, 0x78, 0xee, 0x0c, 0x06, 0x13, 0xbb, 0xf7, 0xd5, 0xc4, 0x62, 0x37, 0x21, 0x3b, 0x30, 0xdf,
	0x5a, 0xad, 0x81, 0xd5, 0x3a, 0x1c, 0x79, 0x2f, 0x44, 0x06, 0xab, 0x76, 0xac, 0xe7, 0xa3, 0x13,
	0x56, 0xf0, 0x32, 0x76, 0x12, 0x11, 0x69, 0x53, 0x64, 0xf0, 0x8b, 0x22, 0x37, 0x4d, 0x8a, 0xcd,
	0xf4, 0xb2, 0x53, 0x7e, 0xb2, 0x49, 0x89, 0x6d, 0xe4, 0x67, 0x9b, 0x6b, 0x90, 0x74, 0xad, 0xc1,
	0xb0, 0x75, 0x48, 0x9a, 0x41, 0x5d, 0x24, 0xb0, 0xfc, 0x9c, 0xad, 0x40, 0xec, 0xd0, 0xe9, 0x93,
	0x5a, 0xce, 0xdf, 0x20, 0x84, 0xb2, 0x12, 0xc4, 0x06, 0x63, 0x2e, 0xa1, 0x6c, 0x65, 0x31, 0x9c,
	0x4e, 0xf0, 0xf0, 0x85, 0xc8, 0xc1, 0xb8, 0xeb, 0xcf, 0xbe, 0xf8, 0x6b, 0x15, 0xd2, 0xb8, 0x65,
	0xaf, 0xf7, 0x1b, 0x8f, 0xe8, 0xa0, 0x70, 0x68, 0xf6, 0xe9, 0x7e, 0x00, 0xdf, 0x55, 0x51, 0xc2,
	0xfa, 0x9f, 0x58, 0x87, 0xae, 0x33, 0x22, 0x1f, 0x9d, 0x31, 0x44, 0x09, 0x17, 0x9d, 0x27, 0xc9,
	0x31, 0x31, 0x4f, 0x5e, 0xa4, 0x0c, 0xdf, 0x1c, 0xb6, 0xd0, 0x11, 0x70, 0xb7, 0x19, 0x3e, 0x58,
	0x7b, 0xfd, 0xe1, 0x89, 0xed, 0x33, 0xeb, 0x84, 0xbb, 0xcb, 0xe4, 0x80, 0x0a, 0xec, 0x07, 0xfc,
	0xa8, 0xc7, 0x37, 0x93, 0x7f, 0x68, 0xf5, 0xe1, 0xa9, 0xec, 0xcf, 0x11, 0x15, 0x9c, 0xf7, 0xa8,
	0xb8, 0xf4, 0x18, 0xb2, 0x92, 0xe1, 0xf3, 0x3c, 0x52, 0x2c, 0xe2, 0xce, 0x42, 0x56, 0xcf, 0xbb,
	0xd6, 0x91, 0xdd, 0x19, 0xae, 0xa9, 0x83, 0x3a, 0xbe, 0x93, 0x87, 0x58, 0xa3, 0xd9, 0xc4, 0xb4,
	0xab, 0xd1, 0x6c, 0x3e, 0xd0, 0x94, 0x7a, 0x05, 0xd2, 0xdd, 0x91, 0x65, 0xa1, 0x07, 0x3e, 0xf5,
	0x70, 0xf7, 0x63, 0x5a, 0x59, 0x1f, 0x57, 0xdf, 0x83, 0xd4, 0x21, 0x3f, 0xde, 0xb1, 0xd3, 0x2f,
	0x35, 0x0a, 0x7f, 0xcb, 0x6f, 0xd8, 0x6e, 0x84, 0x10, 0xd1, 0x63, 0xa1, 0xe1, 0x59, 0xaa, 0x7f,
	0x0e, 0x99, 0x51, 0xeb, 0x02, 0x66, 0xbf, 0xe1, 0xb1, 0xfd, 0x1c, 0xb3, 0xe9, 0x91, 0xa8, 0xaa,
	0xbf, 0x84, 0x05, 0xdb, 0xf1, 0x7e, 0xf1, 0x6b, 0xb5, 0xb9, 0x57, 0xfb, 0x60, 0x66, 0x6a, 0xed,
	0x75, 0x61, 0xf1, 0x0f, 0x07, 0x6c, 0x47, 0x34, 0x70, 0x57, 0x58, 0x6f, 0x80, 0x26, 0x59, 0xa2,
	0x8b, 0x87, 0x73, 0x0c, 0x75, 0xf8, 0xc7, 0x0a, 0xbe, 0x21, 0xf2, 0xb7, 0x11, 0x3b, 0xdc, 0x23,
	0x9e, 0x6d, 0xa7, 0xcb, 0x3f, 0xfe, 0xf0, 0xed, 0x50, 0x94, 0x99, 0xb6, 0x53, 0xab, 0x9e, 0x63,
	0xe7, 0x88, 0x7f, 0x1a, 0x22, 0xdb, 0xa9, 0x55, 0x23, 0x2b, 0x34, 0xb9, 0xc8, 0x80, 0x7a, 0xfc,
	0xe3, 0x0e, 0xdf, 0x10, 0x8f, 0x40, 0x33, 0x2c, 0x9d, 0x3b, 0xa4, 0x1f, 0xf3, 0x4f, 0x3f, 0x42,
	0x96, 0xa6, 0xc6, 0x34, 0xbe, 0xc8, 0x98, 0xde, 0xf2, 0x4f, 0x2d, 0x7c, 0x4b, 0x7b, 0xb3, 0xc6,
	0x34, 0xbe, 0xc8, 0x98, 0xfa, 0xfc, 0x3b, 0x8c, 0x90, 0xa5, 0x5a, 0xb5, 0xfe, 0xdb, 0xc0, 0xe4,
	0xfd, 0x17, 0x11, 0xfb, 0x6c, 0x53, 0x03, 0xfe, 0x81, 0x4d, 0xa0, 0x00, 0xce, 0x9a, 0x65, 0xeb,
	0xdc, 0x61, 0xd9, 0xfc, 0xf3, 0x9b, 0xb0, 0xad, 0x5a, 0xb5, 0xbe, 0x0d, 0xef, 0xc9, 0x33, 0xbc,
	0xd8, 0xc0, 0x1c, 0xfe, 0xf1, 0x48, 0x30, 0x47, 0x41, 0x9b, 0x69, 0xed, 0xdc, 0xa1, 0x0d, 0xf9,
	0x97, 0x25, 0x11, 0x6b, 0xb5, 0x6a, 0xfd, 0x05, 0xe4, 0x25, 0x6b, 0x78, 0x50, 0x3a, 0xc7, 0xd2,
	0x57, 0xfc, 0x83, 0x28, 0xdf, 0x12, 0xa6, 0x5a, 0xd1, 0x3d, 0xe4, 0xc9, 0xc7, 0xd9, 0x76, 0x46,
	0xfc, 0x83, 0x9e, 0x60, 0x44, 0x44, 0x8a, 0xbc, 0x33, 0x74, 0x73, 0x72, 0x8e, 0xa1, 0x31, 0xff,
	0xda, 0x27, 0x18, 0x10, 0x72, 0xea, 0x4e, 0x68, 0x5e, 0x16, 0x26, 0x20, 0x67, 0x9b, 0x71, 0x29,
	0x54, 0x7e, 0x7c, 0x16, 0xa6, 0x2c, 0x5f, 0x64, 0x49, 0x4b, 0x80, 0xc5, 0xfa, 0x36, 0xcc, 0x5f,
	0xca, 0x87, 0x7d, 0xa3, 0xf0, 0x3b, 0x8f, 0xd5, 0xf2, 0x83, 0xea, 0x83, 0x35, 0x63, 0xae, 0x1d,
	0x72, 0x65, 0x2f, 0x61, 0xee, 0x32, 0x7e, 0xec, 0xe7, 0x0a, 0xbf, 0x39, 0x40, 0x63, 0x46, 0xae,
	0x1d, 0x76, 0x66, 0x73, 0x97, 0xf1, 0x64, 0xbf, 0x50, 0xf8, 0x65, 0x53, 0xb5, 0xe2, 0xdb, 0xf1,
	0x9c, 0xd9, 0xdc, 0x65, 0x3c, 0xd9, 0x2f, 0xf9, 0xd5, 0x80, 0x5a, 0x5d, 0x95, 0xed, 0x90, 0xe3,
	0x98, 0xbf, 0x94, 0x27, 0xfb, 0x95, 0x42, 0xd7, 0x4f, 0x6a, 0xb5, 0xea, 0xaf, 0x91, 0xef, 0xcc,
	0xe6, 0x2f, 0xe5, 0xc9, 0xfe, 0x4e, 0xa1, 0x7b, 0x2a, 0xb5, 0xba, 0x16, 0xb2, 0x14, 0x1e, 0xd3,
	0x85, 0x3c, 0xd9, 0xdf, 0x2b, 0x74, 0x7b, 0xa4, 0x56, 0x6b, 0xbe, 0xa5, 0xbd, 0xa9, 0x31, 0x5d,
	0xc8, 0x93, 0xfd, 0x03, 0x9d, 0xc2, 0xea, 0x6a, 0xf5, 0x61, 0xc8, 0x12, 0x39, 0xb3, 0xfc, 0xe5,
	0x3c, 0xd9, 0xaf, 0x15, 0xba, 0xea, 0x53, 0xab, 0x8f, 0x0c, 0x6f, 0x0c, 0x81, 0x33, 0xcb, 0x5f,
	0xce, 0x93, 0xfd, 0xa3, 0x42, 0x97, 0x82, 0x6a, 0xf5, 0x71, 0xd8, 0x16, 0x39, 0x33, 0xed, 0x92,
	0x9e, 0xec, 0x9f, 0x14, 0xfa, 0xf4, 0x47, 0x5d, 0x5b, 0x31, 0xbc, 0x61, 0x48, 0xce, 0x4c, 0xbb,
	0xa4, 0x27, 0xfb, 0x67, 0x85, 0x3e, 0x08, 0x52, 0xd7, 0x1e, 0x44, 0xac, 0xd5, 0xaa, 0xf5, 0x06,
	0xe4, 0x2e, 0xe1, 0xc9, 0xfe, 0x45, 0xbe, 0x75, 0xcd, 0xb6, 0x25, 0x77, 0xf6, 0x23, 0x69, 0x17,
	0x2f, 0xe2, 0xcb, 0xfe, 0x95, 0x12, 0xc4, 0xfa, 0xfb, 0x2f, 0xf9, 0xe5, 0x24, 0xe7, 0xdc, 0x6b,
	0x5b, 0x9d, 0x4f, 0x3b, 0x8e, 0x13, 0x6c, 0x2c, 0xf7, 0x70, 0xaf, 0x82, 0x17, 0xe9, 0x22, 0xee,
	0xed, 0xdf, 0x14, 0xba, 0xcc, 0xcc, 0x09, 0xdb, 0x44, 0xf1, 0x5f, 0x29, 0xee, 0xeb, 0x86, 0xc1,
	0xb4, 0x2f, 0xe0, 0xe8, 0xfe, 0x5d, 0xb9, 0xa4, 0xa7, 0xab, 0xc7, 0x9a, 0xbb, 0x1b, 0xfe, 0x02,
	0x61, 0xcd, 0xfa, 0xa7, 0x6f, 0x9e, 0x74, 0x7b, 0xee, 0xd1, 0xe4, 0xa0, 0x7c, 0xe8, 0x0c, 0xee,
	0x77, 0x9d, 0xbe, 0x69, 0x77, 0xef, 0x93, 0xbd, 0x83, 0x49, 0xe7, 0x7e, 0xcf, 0x76, 0xad, 0x91,
	0x6d, 0xf6, 0xe9, 0x1f, 0x44, 0xa8, 0x76, 0x7c, 0x5f, 0xfe, 0xc7, 0x91, 0xdf, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x65, 0x66, 0xff, 0x47, 0x32, 0x00, 0x00,
}
