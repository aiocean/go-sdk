// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuzone/v1/wallpaper.proto

package wibuzonev1

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

type Wallpaper_Visibility int32

const (
	Wallpaper_VISIBILITY_PRIVATE_UNSPECIFIED Wallpaper_Visibility = 0
	Wallpaper_VISIBILITY_PUBLIC              Wallpaper_Visibility = 1
	Wallpaper_VISIBILITY_UNLISTED            Wallpaper_Visibility = 2
)

// Enum value maps for Wallpaper_Visibility.
var (
	Wallpaper_Visibility_name = map[int32]string{
		0: "VISIBILITY_PRIVATE_UNSPECIFIED",
		1: "VISIBILITY_PUBLIC",
		2: "VISIBILITY_UNLISTED",
	}
	Wallpaper_Visibility_value = map[string]int32{
		"VISIBILITY_PRIVATE_UNSPECIFIED": 0,
		"VISIBILITY_PUBLIC":              1,
		"VISIBILITY_UNLISTED":            2,
	}
)

func (x Wallpaper_Visibility) Enum() *Wallpaper_Visibility {
	p := new(Wallpaper_Visibility)
	*p = x
	return p
}

func (x Wallpaper_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Wallpaper_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_wibuzone_v1_wallpaper_proto_enumTypes[0].Descriptor()
}

func (Wallpaper_Visibility) Type() protoreflect.EnumType {
	return &file_wibuzone_v1_wallpaper_proto_enumTypes[0]
}

func (x Wallpaper_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Wallpaper_Visibility.Descriptor instead.
func (Wallpaper_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{0, 0}
}

type Wallpaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SizeMb        float32                `protobuf:"fixed32,6,opt,name=size_mb,json=sizeMb,proto3" json:"size_mb,omitempty"`
	PublishedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=published_time,json=publishedTime,proto3" json:"published_time,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	ThumbnailUrl  string                 `protobuf:"bytes,10,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	HdUrl         string                 `protobuf:"bytes,11,opt,name=hd_url,json=hdUrl,proto3" json:"hd_url,omitempty"`
	Publisher     *Profile               `protobuf:"bytes,12,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Tags          []*Tag                 `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Visibility    Wallpaper_Visibility   `protobuf:"varint,16,opt,name=visibility,proto3,enum=wibuzone.v1.Wallpaper_Visibility" json:"visibility,omitempty"`
	SourceUrl     string                 `protobuf:"bytes,17,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	Height        float32                `protobuf:"fixed32,18,opt,name=height,proto3" json:"height,omitempty"`
	Width         float32                `protobuf:"fixed32,19,opt,name=width,proto3" json:"width,omitempty"`
	Stats         *WallpaperStats        `protobuf:"bytes,20,opt,name=stats,proto3" json:"stats,omitempty"`
	AgeRating     uint32                 `protobuf:"varint,21,opt,name=age_rating,json=ageRating,proto3" json:"age_rating,omitempty"`
}

func (x *Wallpaper) Reset() {
	*x = Wallpaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallpaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallpaper) ProtoMessage() {}

func (x *Wallpaper) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallpaper.ProtoReflect.Descriptor instead.
func (*Wallpaper) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{0}
}

func (x *Wallpaper) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wallpaper) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Wallpaper) GetSizeMb() float32 {
	if x != nil {
		return x.SizeMb
	}
	return 0
}

func (x *Wallpaper) GetPublishedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedTime
	}
	return nil
}

func (x *Wallpaper) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Wallpaper) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Wallpaper) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *Wallpaper) GetHdUrl() string {
	if x != nil {
		return x.HdUrl
	}
	return ""
}

func (x *Wallpaper) GetPublisher() *Profile {
	if x != nil {
		return x.Publisher
	}
	return nil
}

func (x *Wallpaper) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Wallpaper) GetVisibility() Wallpaper_Visibility {
	if x != nil {
		return x.Visibility
	}
	return Wallpaper_VISIBILITY_PRIVATE_UNSPECIFIED
}

func (x *Wallpaper) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *Wallpaper) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Wallpaper) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Wallpaper) GetStats() *WallpaperStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Wallpaper) GetAgeRating() uint32 {
	if x != nil {
		return x.AgeRating
	}
	return 0
}

type WallpaperStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TotalViews     int64  `protobuf:"varint,2,opt,name=total_views,json=totalViews,proto3" json:"total_views,omitempty"`
	TotalLikes     int64  `protobuf:"varint,3,opt,name=total_likes,json=totalLikes,proto3" json:"total_likes,omitempty"`
	TotalDownloads int64  `protobuf:"varint,4,opt,name=total_downloads,json=totalDownloads,proto3" json:"total_downloads,omitempty"`
}

func (x *WallpaperStats) Reset() {
	*x = WallpaperStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WallpaperStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WallpaperStats) ProtoMessage() {}

func (x *WallpaperStats) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WallpaperStats.ProtoReflect.Descriptor instead.
func (*WallpaperStats) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{1}
}

func (x *WallpaperStats) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WallpaperStats) GetTotalViews() int64 {
	if x != nil {
		return x.TotalViews
	}
	return 0
}

func (x *WallpaperStats) GetTotalLikes() int64 {
	if x != nil {
		return x.TotalLikes
	}
	return 0
}

func (x *WallpaperStats) GetTotalDownloads() int64 {
	if x != nil {
		return x.TotalDownloads
	}
	return 0
}

type UnlikeWallpaperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WallpaperId string `protobuf:"bytes,1,opt,name=wallpaper_id,json=wallpaperId,proto3" json:"wallpaper_id,omitempty"`
}

func (x *UnlikeWallpaperRequest) Reset() {
	*x = UnlikeWallpaperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlikeWallpaperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlikeWallpaperRequest) ProtoMessage() {}

func (x *UnlikeWallpaperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlikeWallpaperRequest.ProtoReflect.Descriptor instead.
func (*UnlikeWallpaperRequest) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{2}
}

func (x *UnlikeWallpaperRequest) GetWallpaperId() string {
	if x != nil {
		return x.WallpaperId
	}
	return ""
}

type UnlikeWallpaperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlikeWallpaperResponse) Reset() {
	*x = UnlikeWallpaperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlikeWallpaperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlikeWallpaperResponse) ProtoMessage() {}

func (x *UnlikeWallpaperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlikeWallpaperResponse.ProtoReflect.Descriptor instead.
func (*UnlikeWallpaperResponse) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{3}
}

type LikeWallpaperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WallpaperId string `protobuf:"bytes,1,opt,name=wallpaper_id,json=wallpaperId,proto3" json:"wallpaper_id,omitempty"`
}

func (x *LikeWallpaperRequest) Reset() {
	*x = LikeWallpaperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeWallpaperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeWallpaperRequest) ProtoMessage() {}

func (x *LikeWallpaperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeWallpaperRequest.ProtoReflect.Descriptor instead.
func (*LikeWallpaperRequest) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{4}
}

func (x *LikeWallpaperRequest) GetWallpaperId() string {
	if x != nil {
		return x.WallpaperId
	}
	return ""
}

type LikeWallpaperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LikeWallpaperResponse) Reset() {
	*x = LikeWallpaperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeWallpaperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeWallpaperResponse) ProtoMessage() {}

func (x *LikeWallpaperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_wallpaper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeWallpaperResponse.ProtoReflect.Descriptor instead.
func (*LikeWallpaperResponse) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_wallpaper_proto_rawDescGZIP(), []int{5}
}

var File_wibuzone_v1_wallpaper_proto protoreflect.FileDescriptor

var file_wibuzone_v1_wallpaper_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77,
	0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x77, 0x69, 0x62,
	0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x05,
	0x0a, 0x09, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x73, 0x69, 0x7a, 0x65, 0x4d, 0x62, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x68, 0x64, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x32, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x7a,
	0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x60, 0x0a, 0x0a, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x1e, 0x56, 0x49, 0x53,
	0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x8b, 0x01,
	0x0a, 0x0e, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x69, 0x65, 0x77,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x3b, 0x0a, 0x16, 0x55,
	0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x6e, 0x6c, 0x69,
	0x6b, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x14, 0x4c, 0x69, 0x6b, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77,
	0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17,
	0x0a, 0x15, 0x4c, 0x69, 0x6b, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa0, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x57, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61,
	0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x17, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x57, 0x69,
	0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wibuzone_v1_wallpaper_proto_rawDescOnce sync.Once
	file_wibuzone_v1_wallpaper_proto_rawDescData = file_wibuzone_v1_wallpaper_proto_rawDesc
)

func file_wibuzone_v1_wallpaper_proto_rawDescGZIP() []byte {
	file_wibuzone_v1_wallpaper_proto_rawDescOnce.Do(func() {
		file_wibuzone_v1_wallpaper_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuzone_v1_wallpaper_proto_rawDescData)
	})
	return file_wibuzone_v1_wallpaper_proto_rawDescData
}

var file_wibuzone_v1_wallpaper_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wibuzone_v1_wallpaper_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wibuzone_v1_wallpaper_proto_goTypes = []interface{}{
	(Wallpaper_Visibility)(0),       // 0: wibuzone.v1.Wallpaper.Visibility
	(*Wallpaper)(nil),               // 1: wibuzone.v1.Wallpaper
	(*WallpaperStats)(nil),          // 2: wibuzone.v1.WallpaperStats
	(*UnlikeWallpaperRequest)(nil),  // 3: wibuzone.v1.UnlikeWallpaperRequest
	(*UnlikeWallpaperResponse)(nil), // 4: wibuzone.v1.UnlikeWallpaperResponse
	(*LikeWallpaperRequest)(nil),    // 5: wibuzone.v1.LikeWallpaperRequest
	(*LikeWallpaperResponse)(nil),   // 6: wibuzone.v1.LikeWallpaperResponse
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
	(*Profile)(nil),                 // 8: wibuzone.v1.Profile
	(*Tag)(nil),                     // 9: wibuzone.v1.Tag
}
var file_wibuzone_v1_wallpaper_proto_depIdxs = []int32{
	7, // 0: wibuzone.v1.Wallpaper.published_time:type_name -> google.protobuf.Timestamp
	7, // 1: wibuzone.v1.Wallpaper.create_time:type_name -> google.protobuf.Timestamp
	7, // 2: wibuzone.v1.Wallpaper.update_time:type_name -> google.protobuf.Timestamp
	8, // 3: wibuzone.v1.Wallpaper.publisher:type_name -> wibuzone.v1.Profile
	9, // 4: wibuzone.v1.Wallpaper.tags:type_name -> wibuzone.v1.Tag
	0, // 5: wibuzone.v1.Wallpaper.visibility:type_name -> wibuzone.v1.Wallpaper.Visibility
	2, // 6: wibuzone.v1.Wallpaper.stats:type_name -> wibuzone.v1.WallpaperStats
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_wibuzone_v1_wallpaper_proto_init() }
func file_wibuzone_v1_wallpaper_proto_init() {
	if File_wibuzone_v1_wallpaper_proto != nil {
		return
	}
	file_wibuzone_v1_tag_proto_init()
	file_wibuzone_v1_profile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuzone_v1_wallpaper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallpaper); i {
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
		file_wibuzone_v1_wallpaper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WallpaperStats); i {
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
		file_wibuzone_v1_wallpaper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlikeWallpaperRequest); i {
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
		file_wibuzone_v1_wallpaper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlikeWallpaperResponse); i {
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
		file_wibuzone_v1_wallpaper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeWallpaperRequest); i {
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
		file_wibuzone_v1_wallpaper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeWallpaperResponse); i {
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
			RawDescriptor: file_wibuzone_v1_wallpaper_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuzone_v1_wallpaper_proto_goTypes,
		DependencyIndexes: file_wibuzone_v1_wallpaper_proto_depIdxs,
		EnumInfos:         file_wibuzone_v1_wallpaper_proto_enumTypes,
		MessageInfos:      file_wibuzone_v1_wallpaper_proto_msgTypes,
	}.Build()
	File_wibuzone_v1_wallpaper_proto = out.File
	file_wibuzone_v1_wallpaper_proto_rawDesc = nil
	file_wibuzone_v1_wallpaper_proto_goTypes = nil
	file_wibuzone_v1_wallpaper_proto_depIdxs = nil
}
