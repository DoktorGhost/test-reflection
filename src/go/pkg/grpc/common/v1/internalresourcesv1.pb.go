// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.29.3
// source: internalresourcesv1.proto

package internalcommonv1

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

type ResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PrivilegeUsed int64 `protobuf:"varint,2,opt,name=privilege_used,json=privilegeUsed,proto3" json:"privilege_used,omitempty"`
}

func (x *ResourcesRequest) Reset() {
	*x = ResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalresourcesv1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcesRequest) ProtoMessage() {}

func (x *ResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalresourcesv1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcesRequest.ProtoReflect.Descriptor instead.
func (*ResourcesRequest) Descriptor() ([]byte, []int) {
	return file_internalresourcesv1_proto_rawDescGZIP(), []int{0}
}

func (x *ResourcesRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ResourcesRequest) GetPrivilegeUsed() int64 {
	if x != nil {
		return x.PrivilegeUsed
	}
	return 0
}

type ResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PrivilegeUsed int64 `protobuf:"varint,2,opt,name=privilege_used,json=privilegeUsed,proto3" json:"privilege_used,omitempty"`
}

func (x *ResourcesResponse) Reset() {
	*x = ResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalresourcesv1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcesResponse) ProtoMessage() {}

func (x *ResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalresourcesv1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcesResponse.ProtoReflect.Descriptor instead.
func (*ResourcesResponse) Descriptor() ([]byte, []int) {
	return file_internalresourcesv1_proto_rawDescGZIP(), []int{1}
}

func (x *ResourcesResponse) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ResourcesResponse) GetPrivilegeUsed() int64 {
	if x != nil {
		return x.PrivilegeUsed
	}
	return 0
}

var File_internalresourcesv1_proto protoreflect.FileDescriptor

var file_internalresourcesv1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x10, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x55, 0x73, 0x65, 0x64, 0x42, 0x53,
	0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6b,
	0x74, 0x6f, 0x72, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalresourcesv1_proto_rawDescOnce sync.Once
	file_internalresourcesv1_proto_rawDescData = file_internalresourcesv1_proto_rawDesc
)

func file_internalresourcesv1_proto_rawDescGZIP() []byte {
	file_internalresourcesv1_proto_rawDescOnce.Do(func() {
		file_internalresourcesv1_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalresourcesv1_proto_rawDescData)
	})
	return file_internalresourcesv1_proto_rawDescData
}

var file_internalresourcesv1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internalresourcesv1_proto_goTypes = []interface{}{
	(*ResourcesRequest)(nil),  // 0: internal.common.v1.resources.ResourcesRequest
	(*ResourcesResponse)(nil), // 1: internal.common.v1.resources.ResourcesResponse
}
var file_internalresourcesv1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internalresourcesv1_proto_init() }
func file_internalresourcesv1_proto_init() {
	if File_internalresourcesv1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalresourcesv1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcesRequest); i {
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
		file_internalresourcesv1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcesResponse); i {
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
			RawDescriptor: file_internalresourcesv1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalresourcesv1_proto_goTypes,
		DependencyIndexes: file_internalresourcesv1_proto_depIdxs,
		MessageInfos:      file_internalresourcesv1_proto_msgTypes,
	}.Build()
	File_internalresourcesv1_proto = out.File
	file_internalresourcesv1_proto_rawDesc = nil
	file_internalresourcesv1_proto_goTypes = nil
	file_internalresourcesv1_proto_depIdxs = nil
}
