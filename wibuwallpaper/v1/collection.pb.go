// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuwallpaper/v1/collection.proto

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

type Collection_Visibility int32

const (
	Collection_VISIBILITY_PRIVATE_UNSPECIFIED Collection_Visibility = 0
	Collection_VISIBILITY_PUBLIC              Collection_Visibility = 1
	Collection_VISIBILITY_UNLISTED            Collection_Visibility = 2
)

// Enum value maps for Collection_Visibility.
var (
	Collection_Visibility_name = map[int32]string{
		0: "VISIBILITY_PRIVATE_UNSPECIFIED",
		1: "VISIBILITY_PUBLIC",
		2: "VISIBILITY_UNLISTED",
	}
	Collection_Visibility_value = map[string]int32{
		"VISIBILITY_PRIVATE_UNSPECIFIED": 0,
		"VISIBILITY_PUBLIC":              1,
		"VISIBILITY_UNLISTED":            2,
	}
)

func (x Collection_Visibility) Enum() *Collection_Visibility {
	p := new(Collection_Visibility)
	*p = x
	return p
}

func (x Collection_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Collection_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_wibuwallpaper_v1_collection_proto_enumTypes[0].Descriptor()
}

func (Collection_Visibility) Type() protoreflect.EnumType {
	return &file_wibuwallpaper_v1_collection_proto_enumTypes[0]
}

func (x Collection_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Collection_Visibility.Descriptor instead.
func (Collection_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_collection_proto_rawDescGZIP(), []int{0, 0}
}

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Name        string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CoverUrl    string                 `protobuf:"bytes,6,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Publisher   *Profile               `protobuf:"bytes,7,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Visibility  Collection_Visibility  `protobuf:"varint,9,opt,name=visibility,proto3,enum=wibuwallpaper.v1.Collection_Visibility" json:"visibility,omitempty"`
	Tags        []*Tag                 `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_collection_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Collection) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Collection) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Collection) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Collection) GetPublisher() *Profile {
	if x != nil {
		return x.Publisher
	}
	return nil
}

func (x *Collection) GetVisibility() Collection_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Collection_VISIBILITY_PRIVATE_UNSPECIFIED
}

func (x *Collection) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CollectionStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPhotos int64 `protobuf:"varint,1,opt,name=total_photos,json=totalPhotos,proto3" json:"total_photos,omitempty"`
}

func (x *CollectionStats) Reset() {
	*x = CollectionStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionStats) ProtoMessage() {}

func (x *CollectionStats) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionStats.ProtoReflect.Descriptor instead.
func (*CollectionStats) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionStats) GetTotalPhotos() int64 {
	if x != nil {
		return x.TotalPhotos
	}
	return 0
}

var File_wibuwallpaper_v1_collection_proto protoreflect.FileDescriptor

var file_wibuwallpaper_v1_collection_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf8, 0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x37,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x77, 0x69,
	0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x60, 0x0a, 0x0a, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x1e, 0x56, 0x49, 0x53,
	0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x34, 0x0a,
	0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x73, 0x42, 0xc4, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
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
	file_wibuwallpaper_v1_collection_proto_rawDescOnce sync.Once
	file_wibuwallpaper_v1_collection_proto_rawDescData = file_wibuwallpaper_v1_collection_proto_rawDesc
)

func file_wibuwallpaper_v1_collection_proto_rawDescGZIP() []byte {
	file_wibuwallpaper_v1_collection_proto_rawDescOnce.Do(func() {
		file_wibuwallpaper_v1_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuwallpaper_v1_collection_proto_rawDescData)
	})
	return file_wibuwallpaper_v1_collection_proto_rawDescData
}

var file_wibuwallpaper_v1_collection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wibuwallpaper_v1_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wibuwallpaper_v1_collection_proto_goTypes = []interface{}{
	(Collection_Visibility)(0),    // 0: wibuwallpaper.v1.Collection.Visibility
	(*Collection)(nil),            // 1: wibuwallpaper.v1.Collection
	(*CollectionStats)(nil),       // 2: wibuwallpaper.v1.CollectionStats
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Profile)(nil),               // 4: wibuwallpaper.v1.Profile
	(*Tag)(nil),                   // 5: wibuwallpaper.v1.Tag
}
var file_wibuwallpaper_v1_collection_proto_depIdxs = []int32{
	3, // 0: wibuwallpaper.v1.Collection.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: wibuwallpaper.v1.Collection.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: wibuwallpaper.v1.Collection.publisher:type_name -> wibuwallpaper.v1.Profile
	0, // 3: wibuwallpaper.v1.Collection.visibility:type_name -> wibuwallpaper.v1.Collection.Visibility
	5, // 4: wibuwallpaper.v1.Collection.tags:type_name -> wibuwallpaper.v1.Tag
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_wibuwallpaper_v1_collection_proto_init() }
func file_wibuwallpaper_v1_collection_proto_init() {
	if File_wibuwallpaper_v1_collection_proto != nil {
		return
	}
	file_wibuwallpaper_v1_profile_proto_init()
	file_wibuwallpaper_v1_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuwallpaper_v1_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_wibuwallpaper_v1_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionStats); i {
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
			RawDescriptor: file_wibuwallpaper_v1_collection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuwallpaper_v1_collection_proto_goTypes,
		DependencyIndexes: file_wibuwallpaper_v1_collection_proto_depIdxs,
		EnumInfos:         file_wibuwallpaper_v1_collection_proto_enumTypes,
		MessageInfos:      file_wibuwallpaper_v1_collection_proto_msgTypes,
	}.Build()
	File_wibuwallpaper_v1_collection_proto = out.File
	file_wibuwallpaper_v1_collection_proto_rawDesc = nil
	file_wibuwallpaper_v1_collection_proto_goTypes = nil
	file_wibuwallpaper_v1_collection_proto_depIdxs = nil
}
