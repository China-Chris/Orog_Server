// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.20.3
// source: chains/bsc.proto

package bsc

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

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ping          string                 `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_chains_bsc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_chains_bsc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_chains_bsc_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pong          string                 `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_chains_bsc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_chains_bsc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_chains_bsc_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_chains_bsc_proto protoreflect.FileDescriptor

var file_chains_bsc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x62, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x62, 0x73, 0x63, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x32, 0x2a, 0x0a, 0x03, 0x42, 0x73, 0x63, 0x12, 0x23, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x62, 0x73, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chains_bsc_proto_rawDescOnce sync.Once
	file_chains_bsc_proto_rawDescData = file_chains_bsc_proto_rawDesc
)

func file_chains_bsc_proto_rawDescGZIP() []byte {
	file_chains_bsc_proto_rawDescOnce.Do(func() {
		file_chains_bsc_proto_rawDescData = protoimpl.X.CompressGZIP(file_chains_bsc_proto_rawDescData)
	})
	return file_chains_bsc_proto_rawDescData
}

var file_chains_bsc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chains_bsc_proto_goTypes = []any{
	(*Request)(nil),  // 0: bsc.Request
	(*Response)(nil), // 1: bsc.Response
}
var file_chains_bsc_proto_depIdxs = []int32{
	0, // 0: bsc.Bsc.Ping:input_type -> bsc.Request
	1, // 1: bsc.Bsc.Ping:output_type -> bsc.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chains_bsc_proto_init() }
func file_chains_bsc_proto_init() {
	if File_chains_bsc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chains_bsc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chains_bsc_proto_goTypes,
		DependencyIndexes: file_chains_bsc_proto_depIdxs,
		MessageInfos:      file_chains_bsc_proto_msgTypes,
	}.Build()
	File_chains_bsc_proto = out.File
	file_chains_bsc_proto_rawDesc = nil
	file_chains_bsc_proto_goTypes = nil
	file_chains_bsc_proto_depIdxs = nil
}
