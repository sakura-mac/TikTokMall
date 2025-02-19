// idl/api.proto; 注解拓展

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.0
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_api_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50101,
		Name:          "api.raw_body",
		Tag:           "bytes,50101,opt,name=raw_body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50102,
		Name:          "api.query",
		Tag:           "bytes,50102,opt,name=query",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50103,
		Name:          "api.header",
		Tag:           "bytes,50103,opt,name=header",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50104,
		Name:          "api.cookie",
		Tag:           "bytes,50104,opt,name=cookie",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50105,
		Name:          "api.body",
		Tag:           "bytes,50105,opt,name=body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50106,
		Name:          "api.path",
		Tag:           "bytes,50106,opt,name=path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50107,
		Name:          "api.vd",
		Tag:           "bytes,50107,opt,name=vd",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50108,
		Name:          "api.form",
		Tag:           "bytes,50108,opt,name=form",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50109,
		Name:          "api.js_conv",
		Tag:           "bytes,50109,opt,name=js_conv",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50110,
		Name:          "api.file_name",
		Tag:           "bytes,50110,opt,name=file_name",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50111,
		Name:          "api.none",
		Tag:           "bytes,50111,opt,name=none",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50131,
		Name:          "api.form_compatible",
		Tag:           "bytes,50131,opt,name=form_compatible",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50132,
		Name:          "api.js_conv_compatible",
		Tag:           "bytes,50132,opt,name=js_conv_compatible",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50133,
		Name:          "api.file_name_compatible",
		Tag:           "bytes,50133,opt,name=file_name_compatible",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50134,
		Name:          "api.none_compatible",
		Tag:           "bytes,50134,opt,name=none_compatible",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51001,
		Name:          "api.go_tag",
		Tag:           "bytes,51001,opt,name=go_tag",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50201,
		Name:          "api.get",
		Tag:           "bytes,50201,opt,name=get",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50202,
		Name:          "api.post",
		Tag:           "bytes,50202,opt,name=post",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50203,
		Name:          "api.put",
		Tag:           "bytes,50203,opt,name=put",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50204,
		Name:          "api.delete",
		Tag:           "bytes,50204,opt,name=delete",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50205,
		Name:          "api.patch",
		Tag:           "bytes,50205,opt,name=patch",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50206,
		Name:          "api.options",
		Tag:           "bytes,50206,opt,name=options",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50207,
		Name:          "api.head",
		Tag:           "bytes,50207,opt,name=head",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50208,
		Name:          "api.any",
		Tag:           "bytes,50208,opt,name=any",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50301,
		Name:          "api.gen_path",
		Tag:           "bytes,50301,opt,name=gen_path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50302,
		Name:          "api.api_version",
		Tag:           "bytes,50302,opt,name=api_version",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50303,
		Name:          "api.tag",
		Tag:           "bytes,50303,opt,name=tag",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50304,
		Name:          "api.name",
		Tag:           "bytes,50304,opt,name=name",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50305,
		Name:          "api.api_level",
		Tag:           "bytes,50305,opt,name=api_level",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50306,
		Name:          "api.serializer",
		Tag:           "bytes,50306,opt,name=serializer",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50307,
		Name:          "api.param",
		Tag:           "bytes,50307,opt,name=param",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50308,
		Name:          "api.baseurl",
		Tag:           "bytes,50308,opt,name=baseurl",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50309,
		Name:          "api.handler_path",
		Tag:           "bytes,50309,opt,name=handler_path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50331,
		Name:          "api.handler_path_compatible",
		Tag:           "bytes,50331,opt,name=handler_path_compatible",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50401,
		Name:          "api.http_code",
		Tag:           "varint,50401,opt,name=http_code",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50402,
		Name:          "api.base_domain",
		Tag:           "bytes,50402,opt,name=base_domain",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50731,
		Name:          "api.base_domain_compatible",
		Tag:           "bytes,50731,opt,name=base_domain_compatible",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50830,
		Name:          "api.reserve",
		Tag:           "bytes,50830,opt,name=reserve",
		Filename:      "api.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string raw_body = 50101;
	E_RawBody = &file_api_proto_extTypes[0]
	// optional string query = 50102;
	E_Query = &file_api_proto_extTypes[1]
	// optional string header = 50103;
	E_Header = &file_api_proto_extTypes[2]
	// optional string cookie = 50104;
	E_Cookie = &file_api_proto_extTypes[3]
	// optional string body = 50105;
	E_Body = &file_api_proto_extTypes[4]
	// optional string path = 50106;
	E_Path = &file_api_proto_extTypes[5]
	// optional string vd = 50107;
	E_Vd = &file_api_proto_extTypes[6]
	// optional string form = 50108;
	E_Form = &file_api_proto_extTypes[7]
	// optional string js_conv = 50109;
	E_JsConv = &file_api_proto_extTypes[8]
	// optional string file_name = 50110;
	E_FileName = &file_api_proto_extTypes[9]
	// optional string none = 50111;
	E_None = &file_api_proto_extTypes[10]
	// 50131~50160 used to extend field option by hz
	//
	// optional string form_compatible = 50131;
	E_FormCompatible = &file_api_proto_extTypes[11]
	// optional string js_conv_compatible = 50132;
	E_JsConvCompatible = &file_api_proto_extTypes[12]
	// optional string file_name_compatible = 50133;
	E_FileNameCompatible = &file_api_proto_extTypes[13]
	// optional string none_compatible = 50134;
	E_NoneCompatible = &file_api_proto_extTypes[14]
	// optional string go_tag = 51001;
	E_GoTag = &file_api_proto_extTypes[15]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string get = 50201;
	E_Get = &file_api_proto_extTypes[16]
	// optional string post = 50202;
	E_Post = &file_api_proto_extTypes[17]
	// optional string put = 50203;
	E_Put = &file_api_proto_extTypes[18]
	// optional string delete = 50204;
	E_Delete = &file_api_proto_extTypes[19]
	// optional string patch = 50205;
	E_Patch = &file_api_proto_extTypes[20]
	// optional string options = 50206;
	E_Options = &file_api_proto_extTypes[21]
	// optional string head = 50207;
	E_Head = &file_api_proto_extTypes[22]
	// optional string any = 50208;
	E_Any = &file_api_proto_extTypes[23]
	// optional string gen_path = 50301;
	E_GenPath = &file_api_proto_extTypes[24] // The path specified by the user when the client code is generated, with a higher priority than api_version
	// optional string api_version = 50302;
	E_ApiVersion = &file_api_proto_extTypes[25] // Specify the value of the :version variable in path when the client code is generated
	// optional string tag = 50303;
	E_Tag = &file_api_proto_extTypes[26] // rpc tag, can be multiple, separated by commas
	// optional string name = 50304;
	E_Name = &file_api_proto_extTypes[27] // Name of rpc
	// optional string api_level = 50305;
	E_ApiLevel = &file_api_proto_extTypes[28] // Interface Level
	// optional string serializer = 50306;
	E_Serializer = &file_api_proto_extTypes[29] // Serialization method
	// optional string param = 50307;
	E_Param = &file_api_proto_extTypes[30] // Whether client requests take public parameters
	// optional string baseurl = 50308;
	E_Baseurl = &file_api_proto_extTypes[31] // Baseurl used in ttnet routing
	// optional string handler_path = 50309;
	E_HandlerPath = &file_api_proto_extTypes[32] // handler_path specifies the path to generate the method
	// 50331~50360 used to extend method option by hz
	//
	// optional string handler_path_compatible = 50331;
	E_HandlerPathCompatible = &file_api_proto_extTypes[33] // handler_path specifies the path to generate the method
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional int32 http_code = 50401;
	E_HttpCode = &file_api_proto_extTypes[34]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional string base_domain = 50402;
	E_BaseDomain = &file_api_proto_extTypes[35]
	// 50731~50760 used to extend service option by hz
	//
	// optional string base_domain_compatible = 50731;
	E_BaseDomainCompatible = &file_api_proto_extTypes[36]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string reserve = 50830;
	E_Reserve = &file_api_proto_extTypes[37]
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x3a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x87,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x42, 0x6f, 0x64, 0x79, 0x3a, 0x35,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb6, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x3a, 0x37,
	0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x3a, 0x33, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x33, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xba, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x3a, 0x2f, 0x0a, 0x02, 0x76, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbb, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x76, 0x64, 0x3a, 0x33, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbc, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x38, 0x0a, 0x07, 0x6a, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x76, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xbd, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x73, 0x43, 0x6f, 0x6e,
	0x76, 0x3a, 0x3c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0x87,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x3a,
	0x33, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x6e, 0x65, 0x3a, 0x48, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3a, 0x4d,
	0x0a, 0x12, 0x6a, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74,
	0x69, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd4, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a, 0x73, 0x43,
	0x6f, 0x6e, 0x76, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3a, 0x51, 0x0a,
	0x14, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd5, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65,
	0x3a, 0x48, 0x0a, 0x0f, 0x6e, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd6, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x6e, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3a, 0x36, 0x0a, 0x06, 0x67, 0x6f,
	0x5f, 0x74, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6f, 0x54,
	0x61, 0x67, 0x3a, 0x32, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x99, 0x88, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x67, 0x65, 0x74, 0x3a, 0x34, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a,
	0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x32, 0x0a, 0x03,
	0x70, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x9b, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x75, 0x74,
	0x3a, 0x38, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9c, 0x88, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x36, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x9d, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x3a, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9e, 0x88,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x34,
	0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9f, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x3a, 0x32, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0x88, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x3a, 0x3b, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x3a, 0x41, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x32, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xff, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x3a, 0x34, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x80, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x3a, 0x3d, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x81, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x3a, 0x40, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x82, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x72, 0x3a, 0x36, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x83, 0x89, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x3a, 0x3a, 0x0a, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x73, 0x65, 0x75, 0x72, 0x6c, 0x3a, 0x43, 0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x85, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x3a, 0x58, 0x0a, 0x17,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9b, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3a, 0x40, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe1, 0x89, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x42, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe2, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x57, 0x0a, 0x16,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xab, 0x8c, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x62, 0x61, 0x73, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x6c, 0x65, 0x3a, 0x3b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x8e, 0x8d, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
}

var file_api_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil),     // 0: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),    // 1: google.protobuf.MethodOptions
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
	(*descriptorpb.ServiceOptions)(nil),   // 3: google.protobuf.ServiceOptions
	(*descriptorpb.MessageOptions)(nil),   // 4: google.protobuf.MessageOptions
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.raw_body:extendee -> google.protobuf.FieldOptions
	0,  // 1: api.query:extendee -> google.protobuf.FieldOptions
	0,  // 2: api.header:extendee -> google.protobuf.FieldOptions
	0,  // 3: api.cookie:extendee -> google.protobuf.FieldOptions
	0,  // 4: api.body:extendee -> google.protobuf.FieldOptions
	0,  // 5: api.path:extendee -> google.protobuf.FieldOptions
	0,  // 6: api.vd:extendee -> google.protobuf.FieldOptions
	0,  // 7: api.form:extendee -> google.protobuf.FieldOptions
	0,  // 8: api.js_conv:extendee -> google.protobuf.FieldOptions
	0,  // 9: api.file_name:extendee -> google.protobuf.FieldOptions
	0,  // 10: api.none:extendee -> google.protobuf.FieldOptions
	0,  // 11: api.form_compatible:extendee -> google.protobuf.FieldOptions
	0,  // 12: api.js_conv_compatible:extendee -> google.protobuf.FieldOptions
	0,  // 13: api.file_name_compatible:extendee -> google.protobuf.FieldOptions
	0,  // 14: api.none_compatible:extendee -> google.protobuf.FieldOptions
	0,  // 15: api.go_tag:extendee -> google.protobuf.FieldOptions
	1,  // 16: api.get:extendee -> google.protobuf.MethodOptions
	1,  // 17: api.post:extendee -> google.protobuf.MethodOptions
	1,  // 18: api.put:extendee -> google.protobuf.MethodOptions
	1,  // 19: api.delete:extendee -> google.protobuf.MethodOptions
	1,  // 20: api.patch:extendee -> google.protobuf.MethodOptions
	1,  // 21: api.options:extendee -> google.protobuf.MethodOptions
	1,  // 22: api.head:extendee -> google.protobuf.MethodOptions
	1,  // 23: api.any:extendee -> google.protobuf.MethodOptions
	1,  // 24: api.gen_path:extendee -> google.protobuf.MethodOptions
	1,  // 25: api.api_version:extendee -> google.protobuf.MethodOptions
	1,  // 26: api.tag:extendee -> google.protobuf.MethodOptions
	1,  // 27: api.name:extendee -> google.protobuf.MethodOptions
	1,  // 28: api.api_level:extendee -> google.protobuf.MethodOptions
	1,  // 29: api.serializer:extendee -> google.protobuf.MethodOptions
	1,  // 30: api.param:extendee -> google.protobuf.MethodOptions
	1,  // 31: api.baseurl:extendee -> google.protobuf.MethodOptions
	1,  // 32: api.handler_path:extendee -> google.protobuf.MethodOptions
	1,  // 33: api.handler_path_compatible:extendee -> google.protobuf.MethodOptions
	2,  // 34: api.http_code:extendee -> google.protobuf.EnumValueOptions
	3,  // 35: api.base_domain:extendee -> google.protobuf.ServiceOptions
	3,  // 36: api.base_domain_compatible:extendee -> google.protobuf.ServiceOptions
	4,  // 37: api.reserve:extendee -> google.protobuf.MessageOptions
	38, // [38:38] is the sub-list for method output_type
	38, // [38:38] is the sub-list for method input_type
	38, // [38:38] is the sub-list for extension type_name
	0,  // [0:38] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 38,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		ExtensionInfos:    file_api_proto_extTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
