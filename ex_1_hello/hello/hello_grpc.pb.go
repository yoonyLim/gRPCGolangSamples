// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: hello_grpc.proto

package hello

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

type MyNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MyNumber) Reset() {
	*x = MyNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyNumber) ProtoMessage() {}

func (x *MyNumber) ProtoReflect() protoreflect.Message {
	mi := &file_hello_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyNumber.ProtoReflect.Descriptor instead.
func (*MyNumber) Descriptor() ([]byte, []int) {
	return file_hello_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *MyNumber) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_hello_grpc_proto protoreflect.FileDescriptor

var file_hello_grpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x20, 0x0a, 0x08, 0x4d, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x31, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x24, 0x0a, 0x0a, 0x4d, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x09, 0x2e, 0x4d, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x4d, 0x79, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_grpc_proto_rawDescOnce sync.Once
	file_hello_grpc_proto_rawDescData = file_hello_grpc_proto_rawDesc
)

func file_hello_grpc_proto_rawDescGZIP() []byte {
	file_hello_grpc_proto_rawDescOnce.Do(func() {
		file_hello_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_grpc_proto_rawDescData)
	})
	return file_hello_grpc_proto_rawDescData
}

var file_hello_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hello_grpc_proto_goTypes = []interface{}{
	(*MyNumber)(nil), // 0: MyNumber
}
var file_hello_grpc_proto_depIdxs = []int32{
	0, // 0: MyService.MyFunction:input_type -> MyNumber
	0, // 1: MyService.MyFunction:output_type -> MyNumber
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_grpc_proto_init() }
func file_hello_grpc_proto_init() {
	if File_hello_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyNumber); i {
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
			RawDescriptor: file_hello_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_grpc_proto_goTypes,
		DependencyIndexes: file_hello_grpc_proto_depIdxs,
		MessageInfos:      file_hello_grpc_proto_msgTypes,
	}.Build()
	File_hello_grpc_proto = out.File
	file_hello_grpc_proto_rawDesc = nil
	file_hello_grpc_proto_goTypes = nil
	file_hello_grpc_proto_depIdxs = nil
}
