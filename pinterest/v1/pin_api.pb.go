// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pinterest/v1/pin_api.proto

package pinterestv1

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

type GetVideoPinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVideoPinRequest) Reset() {
	*x = GetVideoPinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinterest_v1_pin_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoPinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoPinRequest) ProtoMessage() {}

func (x *GetVideoPinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pinterest_v1_pin_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoPinRequest.ProtoReflect.Descriptor instead.
func (*GetVideoPinRequest) Descriptor() ([]byte, []int) {
	return file_pinterest_v1_pin_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetVideoPinRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetVideoPinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pin *VideoPin `protobuf:"bytes,1,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *GetVideoPinResponse) Reset() {
	*x = GetVideoPinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinterest_v1_pin_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoPinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoPinResponse) ProtoMessage() {}

func (x *GetVideoPinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pinterest_v1_pin_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoPinResponse.ProtoReflect.Descriptor instead.
func (*GetVideoPinResponse) Descriptor() ([]byte, []int) {
	return file_pinterest_v1_pin_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetVideoPinResponse) GetPin() *VideoPin {
	if x != nil {
		return x.Pin
	}
	return nil
}

var File_pinterest_v1_pin_api_proto protoreflect.FileDescriptor

var file_pinterest_v1_pin_api_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x70, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x50, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x50, 0x69, 0x6e, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x42, 0xa4, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x50, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61,
	0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0d, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pinterest_v1_pin_api_proto_rawDescOnce sync.Once
	file_pinterest_v1_pin_api_proto_rawDescData = file_pinterest_v1_pin_api_proto_rawDesc
)

func file_pinterest_v1_pin_api_proto_rawDescGZIP() []byte {
	file_pinterest_v1_pin_api_proto_rawDescOnce.Do(func() {
		file_pinterest_v1_pin_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pinterest_v1_pin_api_proto_rawDescData)
	})
	return file_pinterest_v1_pin_api_proto_rawDescData
}

var file_pinterest_v1_pin_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pinterest_v1_pin_api_proto_goTypes = []interface{}{
	(*GetVideoPinRequest)(nil),  // 0: pinterest.v1.GetVideoPinRequest
	(*GetVideoPinResponse)(nil), // 1: pinterest.v1.GetVideoPinResponse
	(*VideoPin)(nil),            // 2: pinterest.v1.VideoPin
}
var file_pinterest_v1_pin_api_proto_depIdxs = []int32{
	2, // 0: pinterest.v1.GetVideoPinResponse.pin:type_name -> pinterest.v1.VideoPin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pinterest_v1_pin_api_proto_init() }
func file_pinterest_v1_pin_api_proto_init() {
	if File_pinterest_v1_pin_api_proto != nil {
		return
	}
	file_pinterest_v1_pin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pinterest_v1_pin_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoPinRequest); i {
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
		file_pinterest_v1_pin_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoPinResponse); i {
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
			RawDescriptor: file_pinterest_v1_pin_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pinterest_v1_pin_api_proto_goTypes,
		DependencyIndexes: file_pinterest_v1_pin_api_proto_depIdxs,
		MessageInfos:      file_pinterest_v1_pin_api_proto_msgTypes,
	}.Build()
	File_pinterest_v1_pin_api_proto = out.File
	file_pinterest_v1_pin_api_proto_rawDesc = nil
	file_pinterest_v1_pin_api_proto_goTypes = nil
	file_pinterest_v1_pin_api_proto_depIdxs = nil
}
