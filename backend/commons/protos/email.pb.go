// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: commons/protos/email.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	mi := &file_commons_protos_email_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commons_protos_email_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_commons_protos_email_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type EmailResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	SuccessMessage string                 `protobuf:"bytes,1,opt,name=successMessage,proto3" json:"successMessage,omitempty"`
	Code           string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	mi := &file_commons_protos_email_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commons_protos_email_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_commons_protos_email_proto_rawDescGZIP(), []int{1}
}

func (x *EmailResponse) GetSuccessMessage() string {
	if x != nil {
		return x.SuccessMessage
	}
	return ""
}

func (x *EmailResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_commons_protos_email_proto protoreflect.FileDescriptor

var file_commons_protos_email_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4b, 0x0a, 0x0d, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x4b, 0x0a, 0x0d, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x4d, 0x41, 0x41, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_commons_protos_email_proto_rawDescOnce sync.Once
	file_commons_protos_email_proto_rawDescData []byte
)

func file_commons_protos_email_proto_rawDescGZIP() []byte {
	file_commons_protos_email_proto_rawDescOnce.Do(func() {
		file_commons_protos_email_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_commons_protos_email_proto_rawDesc), len(file_commons_protos_email_proto_rawDesc)))
	})
	return file_commons_protos_email_proto_rawDescData
}

var file_commons_protos_email_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commons_protos_email_proto_goTypes = []any{
	(*EmailRequest)(nil),  // 0: protos.EmailRequest
	(*EmailResponse)(nil), // 1: protos.EmailResponse
}
var file_commons_protos_email_proto_depIdxs = []int32{
	0, // 0: protos.EmailServices.SendEmail:input_type -> protos.EmailRequest
	1, // 1: protos.EmailServices.SendEmail:output_type -> protos.EmailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commons_protos_email_proto_init() }
func file_commons_protos_email_proto_init() {
	if File_commons_protos_email_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_commons_protos_email_proto_rawDesc), len(file_commons_protos_email_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commons_protos_email_proto_goTypes,
		DependencyIndexes: file_commons_protos_email_proto_depIdxs,
		MessageInfos:      file_commons_protos_email_proto_msgTypes,
	}.Build()
	File_commons_protos_email_proto = out.File
	file_commons_protos_email_proto_goTypes = nil
	file_commons_protos_email_proto_depIdxs = nil
}
