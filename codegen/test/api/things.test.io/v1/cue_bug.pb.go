// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/skv2/codegen/test/cue_bug.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CueBugSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MyOneOf:
	//
	//	*CueBugSpec_MyParent
	MyOneOf isCueBugSpec_MyOneOf `protobuf_oneof:"my_one_of"`
}

func (x *CueBugSpec) Reset() {
	*x = CueBugSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CueBugSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CueBugSpec) ProtoMessage() {}

func (x *CueBugSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CueBugSpec.ProtoReflect.Descriptor instead.
func (*CueBugSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescGZIP(), []int{0}
}

func (m *CueBugSpec) GetMyOneOf() isCueBugSpec_MyOneOf {
	if m != nil {
		return m.MyOneOf
	}
	return nil
}

func (x *CueBugSpec) GetMyParent() *CueBugSpec_ParentMessage {
	if x, ok := x.GetMyOneOf().(*CueBugSpec_MyParent); ok {
		return x.MyParent
	}
	return nil
}

type isCueBugSpec_MyOneOf interface {
	isCueBugSpec_MyOneOf()
}

type CueBugSpec_MyParent struct {
	MyParent *CueBugSpec_ParentMessage `protobuf:"bytes,1,opt,name=my_parent,json=myParent,proto3,oneof"`
}

func (*CueBugSpec_MyParent) isCueBugSpec_MyOneOf() {}

type CueBugStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CueBugStatus) Reset() {
	*x = CueBugStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CueBugStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CueBugStatus) ProtoMessage() {}

func (x *CueBugStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CueBugStatus.ProtoReflect.Descriptor instead.
func (*CueBugStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescGZIP(), []int{1}
}

func (x *CueBugStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CueBugSpec_ChildMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CueBugSpec_ChildMessage) Reset() {
	*x = CueBugSpec_ChildMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CueBugSpec_ChildMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CueBugSpec_ChildMessage) ProtoMessage() {}

func (x *CueBugSpec_ChildMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CueBugSpec_ChildMessage.ProtoReflect.Descriptor instead.
func (*CueBugSpec_ChildMessage) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CueBugSpec_ChildMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CueBugSpec_ParentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyChild *CueBugSpec_ChildMessage `protobuf:"bytes,2,opt,name=my_child,json=myChild,proto3" json:"my_child,omitempty"`
}

func (x *CueBugSpec_ParentMessage) Reset() {
	*x = CueBugSpec_ParentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CueBugSpec_ParentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CueBugSpec_ParentMessage) ProtoMessage() {}

func (x *CueBugSpec_ParentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CueBugSpec_ParentMessage.ProtoReflect.Descriptor instead.
func (*CueBugSpec_ParentMessage) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CueBugSpec_ParentMessage) GetMyChild() *CueBugSpec_ChildMessage {
	if x != nil {
		return x.MyChild
	}
	return nil
}

var File_github_com_solo_io_skv2_codegen_test_cue_bug_proto protoreflect.FileDescriptor

var file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65,
	0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x69, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x43, 0x75, 0x65, 0x42, 0x75, 0x67, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x47, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x75, 0x65, 0x42, 0x75, 0x67, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x6d, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x1a, 0x22, 0x0a, 0x0c,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x53, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x79, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x75, 0x65, 0x42, 0x75, 0x67, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x79,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x6f, 0x6e, 0x65, 0x5f,
	0x6f, 0x66, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x65, 0x42, 0x75, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescOnce sync.Once
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescData = file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDesc
)

func file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescData)
	})
	return file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDescData
}

var file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_goTypes = []interface{}{
	(*CueBugSpec)(nil),               // 0: things.test.io.CueBugSpec
	(*CueBugStatus)(nil),             // 1: things.test.io.CueBugStatus
	(*CueBugSpec_ChildMessage)(nil),  // 2: things.test.io.CueBugSpec.ChildMessage
	(*CueBugSpec_ParentMessage)(nil), // 3: things.test.io.CueBugSpec.ParentMessage
}
var file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_depIdxs = []int32{
	3, // 0: things.test.io.CueBugSpec.my_parent:type_name -> things.test.io.CueBugSpec.ParentMessage
	2, // 1: things.test.io.CueBugSpec.ParentMessage.my_child:type_name -> things.test.io.CueBugSpec.ChildMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_init() }
func file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_init() {
	if File_github_com_solo_io_skv2_codegen_test_cue_bug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CueBugSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CueBugStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CueBugSpec_ChildMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CueBugSpec_ParentMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CueBugSpec_MyParent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_skv2_codegen_test_cue_bug_proto = out.File
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_rawDesc = nil
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_goTypes = nil
	file_github_com_solo_io_skv2_codegen_test_cue_bug_proto_depIdxs = nil
}
