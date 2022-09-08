// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pixiv/v1/artwork.proto

package pixivv1

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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixiv_v1_artwork_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_pixiv_v1_artwork_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_pixiv_v1_artwork_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Artwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string      `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tags          []*Tag      `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	AgeLimit      string      `protobuf:"bytes,5,opt,name=age_limit,json=ageLimit,proto3" json:"age_limit,omitempty"`
	UserId        string      `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ImageUrls     []*ImageUrl `protobuf:"bytes,7,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	LikeCount     int64       `protobuf:"varint,8,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	ViewCount     int64       `protobuf:"varint,9,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	CommentCount  int64       `protobuf:"varint,10,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	BookmarkCount int64       `protobuf:"varint,11,opt,name=bookmark_count,json=bookmarkCount,proto3" json:"bookmark_count,omitempty"`
}

func (x *Artwork) Reset() {
	*x = Artwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixiv_v1_artwork_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artwork) ProtoMessage() {}

func (x *Artwork) ProtoReflect() protoreflect.Message {
	mi := &file_pixiv_v1_artwork_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artwork.ProtoReflect.Descriptor instead.
func (*Artwork) Descriptor() ([]byte, []int) {
	return file_pixiv_v1_artwork_proto_rawDescGZIP(), []int{1}
}

func (x *Artwork) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Artwork) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Artwork) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Artwork) GetAgeLimit() string {
	if x != nil {
		return x.AgeLimit
	}
	return ""
}

func (x *Artwork) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Artwork) GetImageUrls() []*ImageUrl {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *Artwork) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *Artwork) GetViewCount() int64 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *Artwork) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Artwork) GetBookmarkCount() int64 {
	if x != nil {
		return x.BookmarkCount
	}
	return 0
}

type ImageUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thumb        string  `protobuf:"bytes,1,opt,name=thumb,proto3" json:"thumb,omitempty"`
	Small        string  `protobuf:"bytes,2,opt,name=small,proto3" json:"small,omitempty"`
	Regular      string  `protobuf:"bytes,3,opt,name=regular,proto3" json:"regular,omitempty"`
	Original     string  `protobuf:"bytes,4,opt,name=original,proto3" json:"original,omitempty"`
	Width        uint32  `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height       uint32  `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	Mini         string  `protobuf:"bytes,7,opt,name=mini,proto3" json:"mini,omitempty"`
	OriginalSize float32 `protobuf:"fixed32,9,opt,name=original_size,json=originalSize,proto3" json:"original_size,omitempty"`
}

func (x *ImageUrl) Reset() {
	*x = ImageUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pixiv_v1_artwork_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUrl) ProtoMessage() {}

func (x *ImageUrl) ProtoReflect() protoreflect.Message {
	mi := &file_pixiv_v1_artwork_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUrl.ProtoReflect.Descriptor instead.
func (*ImageUrl) Descriptor() ([]byte, []int) {
	return file_pixiv_v1_artwork_proto_rawDescGZIP(), []int{2}
}

func (x *ImageUrl) GetThumb() string {
	if x != nil {
		return x.Thumb
	}
	return ""
}

func (x *ImageUrl) GetSmall() string {
	if x != nil {
		return x.Small
	}
	return ""
}

func (x *ImageUrl) GetRegular() string {
	if x != nil {
		return x.Regular
	}
	return ""
}

func (x *ImageUrl) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

func (x *ImageUrl) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageUrl) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageUrl) GetMini() string {
	if x != nil {
		return x.Mini
	}
	return ""
}

func (x *ImageUrl) GetOriginalSize() float32 {
	if x != nil {
		return x.OriginalSize
	}
	return 0
}

var File_pixiv_v1_artwork_proto protoreflect.FileDescriptor

var file_pixiv_v1_artwork_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x69, 0x78, 0x69, 0x76, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x69, 0x78, 0x69, 0x76, 0x2e,
	0x76, 0x31, 0x22, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc5, 0x02,
	0x0a, 0x07, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x69, 0x78, 0x69, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x69, 0x78, 0x69, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x69, 0x6e, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x89, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x76, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x69, 0x78, 0x69, 0x76, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x69, 0x78, 0x69, 0x76, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x50, 0x69, 0x78, 0x69, 0x76, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x50, 0x69, 0x78, 0x69,
	0x76, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x69, 0x78, 0x69, 0x76, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x50, 0x69,
	0x78, 0x69, 0x76, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pixiv_v1_artwork_proto_rawDescOnce sync.Once
	file_pixiv_v1_artwork_proto_rawDescData = file_pixiv_v1_artwork_proto_rawDesc
)

func file_pixiv_v1_artwork_proto_rawDescGZIP() []byte {
	file_pixiv_v1_artwork_proto_rawDescOnce.Do(func() {
		file_pixiv_v1_artwork_proto_rawDescData = protoimpl.X.CompressGZIP(file_pixiv_v1_artwork_proto_rawDescData)
	})
	return file_pixiv_v1_artwork_proto_rawDescData
}

var file_pixiv_v1_artwork_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pixiv_v1_artwork_proto_goTypes = []interface{}{
	(*Tag)(nil),      // 0: pixiv.v1.Tag
	(*Artwork)(nil),  // 1: pixiv.v1.Artwork
	(*ImageUrl)(nil), // 2: pixiv.v1.ImageUrl
}
var file_pixiv_v1_artwork_proto_depIdxs = []int32{
	0, // 0: pixiv.v1.Artwork.tags:type_name -> pixiv.v1.Tag
	2, // 1: pixiv.v1.Artwork.image_urls:type_name -> pixiv.v1.ImageUrl
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pixiv_v1_artwork_proto_init() }
func file_pixiv_v1_artwork_proto_init() {
	if File_pixiv_v1_artwork_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pixiv_v1_artwork_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_pixiv_v1_artwork_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artwork); i {
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
		file_pixiv_v1_artwork_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUrl); i {
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
			RawDescriptor: file_pixiv_v1_artwork_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pixiv_v1_artwork_proto_goTypes,
		DependencyIndexes: file_pixiv_v1_artwork_proto_depIdxs,
		MessageInfos:      file_pixiv_v1_artwork_proto_msgTypes,
	}.Build()
	File_pixiv_v1_artwork_proto = out.File
	file_pixiv_v1_artwork_proto_rawDesc = nil
	file_pixiv_v1_artwork_proto_goTypes = nil
	file_pixiv_v1_artwork_proto_depIdxs = nil
}
