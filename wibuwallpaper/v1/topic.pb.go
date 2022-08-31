// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuwallpaper/v1/topic.proto

package wibuwallpaperv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	ThumbnailUrl string                 `protobuf:"bytes,6,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_topic_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Topic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topic) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Topic) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Topic) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Topic) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

type TopicStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPhotos    int64 `protobuf:"varint,1,opt,name=total_photos,json=totalPhotos,proto3" json:"total_photos,omitempty"`
	TotalFollowers int64 `protobuf:"varint,2,opt,name=total_followers,json=totalFollowers,proto3" json:"total_followers,omitempty"`
}

func (x *TopicStats) Reset() {
	*x = TopicStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicStats) ProtoMessage() {}

func (x *TopicStats) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicStats.ProtoReflect.Descriptor instead.
func (*TopicStats) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_topic_proto_rawDescGZIP(), []int{1}
}

func (x *TopicStats) GetTotalPhotos() int64 {
	if x != nil {
		return x.TotalPhotos
	}
	return 0
}

func (x *TopicStats) GetTotalFollowers() int64 {
	if x != nil {
		return x.TotalFollowers
	}
	return 0
}

var File_wibuwallpaper_v1_topic_proto protoreflect.FileDescriptor

var file_wibuwallpaper_v1_topic_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xec, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c,
	0x22, 0x58, 0x0a, 0x0a, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x42, 0xbf, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69,
	0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x77, 0x69, 0x62,
	0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69,
	0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x57, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x57, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x57, 0x69, 0x62, 0x75, 0x77,
	0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wibuwallpaper_v1_topic_proto_rawDescOnce sync.Once
	file_wibuwallpaper_v1_topic_proto_rawDescData = file_wibuwallpaper_v1_topic_proto_rawDesc
)

func file_wibuwallpaper_v1_topic_proto_rawDescGZIP() []byte {
	file_wibuwallpaper_v1_topic_proto_rawDescOnce.Do(func() {
		file_wibuwallpaper_v1_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuwallpaper_v1_topic_proto_rawDescData)
	})
	return file_wibuwallpaper_v1_topic_proto_rawDescData
}

var file_wibuwallpaper_v1_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wibuwallpaper_v1_topic_proto_goTypes = []interface{}{
	(*Topic)(nil),                 // 0: wibuwallpaper.v1.Topic
	(*TopicStats)(nil),            // 1: wibuwallpaper.v1.TopicStats
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_wibuwallpaper_v1_topic_proto_depIdxs = []int32{
	2, // 0: wibuwallpaper.v1.Topic.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: wibuwallpaper.v1.Topic.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wibuwallpaper_v1_topic_proto_init() }
func file_wibuwallpaper_v1_topic_proto_init() {
	if File_wibuwallpaper_v1_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wibuwallpaper_v1_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
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
		file_wibuwallpaper_v1_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicStats); i {
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
			RawDescriptor: file_wibuwallpaper_v1_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuwallpaper_v1_topic_proto_goTypes,
		DependencyIndexes: file_wibuwallpaper_v1_topic_proto_depIdxs,
		MessageInfos:      file_wibuwallpaper_v1_topic_proto_msgTypes,
	}.Build()
	File_wibuwallpaper_v1_topic_proto = out.File
	file_wibuwallpaper_v1_topic_proto_rawDesc = nil
	file_wibuwallpaper_v1_topic_proto_goTypes = nil
	file_wibuwallpaper_v1_topic_proto_depIdxs = nil
}