// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuwallpaper/v1/topic_api.proto

package wibuwallpaperv1

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

type ListTopicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTopicsRequest) Reset() {
	*x = ListTopicsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_topic_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopicsRequest) ProtoMessage() {}

func (x *ListTopicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_topic_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopicsRequest.ProtoReflect.Descriptor instead.
func (*ListTopicsRequest) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_topic_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListTopicsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListTopicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics        []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTopicsResponse) Reset() {
	*x = ListTopicsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_topic_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopicsResponse) ProtoMessage() {}

func (x *ListTopicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_topic_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopicsResponse.ProtoReflect.Descriptor instead.
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_topic_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListTopicsResponse) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *ListTopicsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_wibuwallpaper_v1_topic_api_proto protoreflect.FileDescriptor

var file_wibuwallpaper_v1_topic_api_proto_rawDesc = []byte{
	0x0a, 0x20, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77,
	0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0xc2, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69,
	0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x77, 0x69, 0x62, 0x75, 0x77,
	0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58,
	0x58, 0xaa, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61,
	0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wibuwallpaper_v1_topic_api_proto_rawDescOnce sync.Once
	file_wibuwallpaper_v1_topic_api_proto_rawDescData = file_wibuwallpaper_v1_topic_api_proto_rawDesc
)

func file_wibuwallpaper_v1_topic_api_proto_rawDescGZIP() []byte {
	file_wibuwallpaper_v1_topic_api_proto_rawDescOnce.Do(func() {
		file_wibuwallpaper_v1_topic_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuwallpaper_v1_topic_api_proto_rawDescData)
	})
	return file_wibuwallpaper_v1_topic_api_proto_rawDescData
}

var file_wibuwallpaper_v1_topic_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wibuwallpaper_v1_topic_api_proto_goTypes = []interface{}{
	(*ListTopicsRequest)(nil),  // 0: wibuwallpaper.v1.ListTopicsRequest
	(*ListTopicsResponse)(nil), // 1: wibuwallpaper.v1.ListTopicsResponse
	(*Topic)(nil),              // 2: wibuwallpaper.v1.Topic
}
var file_wibuwallpaper_v1_topic_api_proto_depIdxs = []int32{
	2, // 0: wibuwallpaper.v1.ListTopicsResponse.topics:type_name -> wibuwallpaper.v1.Topic
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wibuwallpaper_v1_topic_api_proto_init() }
func file_wibuwallpaper_v1_topic_api_proto_init() {
	if File_wibuwallpaper_v1_topic_api_proto != nil {
		return
	}
	file_wibuwallpaper_v1_topic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuwallpaper_v1_topic_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopicsRequest); i {
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
		file_wibuwallpaper_v1_topic_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopicsResponse); i {
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
			RawDescriptor: file_wibuwallpaper_v1_topic_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuwallpaper_v1_topic_api_proto_goTypes,
		DependencyIndexes: file_wibuwallpaper_v1_topic_api_proto_depIdxs,
		MessageInfos:      file_wibuwallpaper_v1_topic_api_proto_msgTypes,
	}.Build()
	File_wibuwallpaper_v1_topic_api_proto = out.File
	file_wibuwallpaper_v1_topic_api_proto_rawDesc = nil
	file_wibuwallpaper_v1_topic_api_proto_goTypes = nil
	file_wibuwallpaper_v1_topic_api_proto_depIdxs = nil
}
