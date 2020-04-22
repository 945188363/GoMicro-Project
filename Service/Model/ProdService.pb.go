// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: ProdService

package Model

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProdRequest1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ProdRequest1) Reset() {
	*x = ProdRequest1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdService_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRequest1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRequest1) ProtoMessage() {}

func (x *ProdRequest1) ProtoReflect() protoreflect.Message {
	mi := &file_ProdService_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest1.ProtoReflect.Descriptor instead.
func (*ProdRequest1) Descriptor() ([]byte, []int) {
	return file_ProdService_rawDescGZIP(), []int{0}
}

func (x *ProdRequest1) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ProdResponse1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProdResponse1) Reset() {
	*x = ProdResponse1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdService_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponse1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponse1) ProtoMessage() {}

func (x *ProdResponse1) ProtoReflect() protoreflect.Message {
	mi := &file_ProdService_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse1.ProtoReflect.Descriptor instead.
func (*ProdResponse1) Descriptor() ([]byte, []int) {
	return file_ProdService_rawDescGZIP(), []int{1}
}

func (x *ProdResponse1) GetData() []*ProdModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ProdService protoreflect.FileDescriptor

var file_ProdService_rawDesc = []byte{
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x06, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x36, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4a,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x12, 0x3a,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x31, 0x1a, 0x15, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProdService_rawDescOnce sync.Once
	file_ProdService_rawDescData = file_ProdService_rawDesc
)

func file_ProdService_rawDescGZIP() []byte {
	file_ProdService_rawDescOnce.Do(func() {
		file_ProdService_rawDescData = protoimpl.X.CompressGZIP(file_ProdService_rawDescData)
	})
	return file_ProdService_rawDescData
}

var file_ProdService_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ProdService_goTypes = []interface{}{
	(*ProdRequest1)(nil),  // 0: Models.ProdRequest1
	(*ProdResponse1)(nil), // 1: Models.ProdResponse1
	(*ProdModel)(nil),     // 2: Models.ProdModel
}
var file_ProdService_depIdxs = []int32{
	2, // 0: Models.ProdResponse1.data:type_name -> Models.ProdModel
	0, // 1: Models.ProdService1.GetProdList:input_type -> Models.ProdRequest1
	1, // 2: Models.ProdService1.GetProdList:output_type -> Models.ProdResponse1
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ProdService_init() }
func file_ProdService_init() {
	if File_ProdService != nil {
		return
	}
	file_Prods_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ProdService_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRequest1); i {
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
		file_ProdService_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponse1); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ProdService_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ProdService_goTypes,
		DependencyIndexes: file_ProdService_depIdxs,
		MessageInfos:      file_ProdService_msgTypes,
	}.Build()
	File_ProdService = out.File
	file_ProdService_rawDesc = nil
	file_ProdService_goTypes = nil
	file_ProdService_depIdxs = nil
}
