// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/plugin/server/federation/v1/federation.proto

package federationv1

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

type PushBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundle string `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *PushBundleRequest) Reset() {
	*x = PushBundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBundleRequest) ProtoMessage() {}

func (x *PushBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBundleRequest.ProtoReflect.Descriptor instead.
func (*PushBundleRequest) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_federation_v1_federation_proto_rawDescGZIP(), []int{0}
}

func (x *PushBundleRequest) GetBundle() string {
	if x != nil {
		return x.Bundle
	}
	return ""
}

type PushBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushBundleResponse) Reset() {
	*x = PushBundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBundleResponse) ProtoMessage() {}

func (x *PushBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBundleResponse.ProtoReflect.Descriptor instead.
func (*PushBundleResponse) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_federation_v1_federation_proto_rawDescGZIP(), []int{1}
}

type RelatioshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestField string `protobuf:"bytes,1,opt,name=request_field,json=requestField,proto3" json:"request_field,omitempty"`
}

func (x *RelatioshipRequest) Reset() {
	*x = RelatioshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatioshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatioshipRequest) ProtoMessage() {}

func (x *RelatioshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatioshipRequest.ProtoReflect.Descriptor instead.
func (*RelatioshipRequest) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_federation_v1_federation_proto_rawDescGZIP(), []int{2}
}

func (x *RelatioshipRequest) GetRequestField() string {
	if x != nil {
		return x.RequestField
	}
	return ""
}

type RelatioshipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseField string `protobuf:"bytes,1,opt,name=response_field,json=responseField,proto3" json:"response_field,omitempty"`
}

func (x *RelatioshipResponse) Reset() {
	*x = RelatioshipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatioshipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatioshipResponse) ProtoMessage() {}

func (x *RelatioshipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_federation_v1_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatioshipResponse.ProtoReflect.Descriptor instead.
func (*RelatioshipResponse) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_federation_v1_federation_proto_rawDescGZIP(), []int{3}
}

func (x *RelatioshipResponse) GetResponseField() string {
	if x != nil {
		return x.ResponseField
	}
	return ""
}

var File_spire_plugin_server_federation_v1_federation_proto protoreflect.FileDescriptor

var file_spire_plugin_server_federation_v1_federation_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x2b, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x75, 0x73, 0x68, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x12, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x8e, 0x02, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x79, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x34, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84, 0x01,
	0x0a, 0x13, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x35, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_plugin_server_federation_v1_federation_proto_rawDescOnce sync.Once
	file_spire_plugin_server_federation_v1_federation_proto_rawDescData = file_spire_plugin_server_federation_v1_federation_proto_rawDesc
)

func file_spire_plugin_server_federation_v1_federation_proto_rawDescGZIP() []byte {
	file_spire_plugin_server_federation_v1_federation_proto_rawDescOnce.Do(func() {
		file_spire_plugin_server_federation_v1_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_plugin_server_federation_v1_federation_proto_rawDescData)
	})
	return file_spire_plugin_server_federation_v1_federation_proto_rawDescData
}

var file_spire_plugin_server_federation_v1_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spire_plugin_server_federation_v1_federation_proto_goTypes = []interface{}{
	(*PushBundleRequest)(nil),   // 0: spire.plugin.server.federation.v1.PushBundleRequest
	(*PushBundleResponse)(nil),  // 1: spire.plugin.server.federation.v1.PushBundleResponse
	(*RelatioshipRequest)(nil),  // 2: spire.plugin.server.federation.v1.RelatioshipRequest
	(*RelatioshipResponse)(nil), // 3: spire.plugin.server.federation.v1.RelatioshipResponse
}
var file_spire_plugin_server_federation_v1_federation_proto_depIdxs = []int32{
	0, // 0: spire.plugin.server.federation.v1.Federation.PushBundle:input_type -> spire.plugin.server.federation.v1.PushBundleRequest
	2, // 1: spire.plugin.server.federation.v1.Federation.ApproveRelationship:input_type -> spire.plugin.server.federation.v1.RelatioshipRequest
	1, // 2: spire.plugin.server.federation.v1.Federation.PushBundle:output_type -> spire.plugin.server.federation.v1.PushBundleResponse
	3, // 3: spire.plugin.server.federation.v1.Federation.ApproveRelationship:output_type -> spire.plugin.server.federation.v1.RelatioshipResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spire_plugin_server_federation_v1_federation_proto_init() }
func file_spire_plugin_server_federation_v1_federation_proto_init() {
	if File_spire_plugin_server_federation_v1_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_plugin_server_federation_v1_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBundleRequest); i {
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
		file_spire_plugin_server_federation_v1_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBundleResponse); i {
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
		file_spire_plugin_server_federation_v1_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatioshipRequest); i {
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
		file_spire_plugin_server_federation_v1_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatioshipResponse); i {
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
			RawDescriptor: file_spire_plugin_server_federation_v1_federation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_plugin_server_federation_v1_federation_proto_goTypes,
		DependencyIndexes: file_spire_plugin_server_federation_v1_federation_proto_depIdxs,
		MessageInfos:      file_spire_plugin_server_federation_v1_federation_proto_msgTypes,
	}.Build()
	File_spire_plugin_server_federation_v1_federation_proto = out.File
	file_spire_plugin_server_federation_v1_federation_proto_rawDesc = nil
	file_spire_plugin_server_federation_v1_federation_proto_goTypes = nil
	file_spire_plugin_server_federation_v1_federation_proto_depIdxs = nil
}