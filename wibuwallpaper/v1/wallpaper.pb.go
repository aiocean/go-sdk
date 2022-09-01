// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuwallpaper/v1/wallpaper.proto

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
	return file_wibuwallpaper_v1_wallpaper_proto_enumTypes[0].Descriptor()
}

func (Wallpaper_Visibility) Type() protoreflect.EnumType {
	return &file_wibuwallpaper_v1_wallpaper_proto_enumTypes[0]
}

func (x Wallpaper_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Wallpaper_Visibility.Descriptor instead.
func (Wallpaper_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_proto_rawDescGZIP(), []int{0, 0}
}

type Wallpaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description    string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TotalViews     int64                  `protobuf:"varint,3,opt,name=total_views,json=totalViews,proto3" json:"total_views,omitempty"`
	TotalLikes     int64                  `protobuf:"varint,4,opt,name=total_likes,json=totalLikes,proto3" json:"total_likes,omitempty"`
	TotalDownloads int64                  `protobuf:"varint,5,opt,name=total_downloads,json=totalDownloads,proto3" json:"total_downloads,omitempty"`
	SizeMb         float32                `protobuf:"fixed32,6,opt,name=size_mb,json=sizeMb,proto3" json:"size_mb,omitempty"`
	PublishedTime  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=published_time,json=publishedTime,proto3" json:"published_time,omitempty"`
	CreateTime     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	ThumbnailUrl   string                 `protobuf:"bytes,10,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	HdUrl          string                 `protobuf:"bytes,11,opt,name=hd_url,json=hdUrl,proto3" json:"hd_url,omitempty"`
	Publisher      *Profile               `protobuf:"bytes,12,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Topics         []*Topic               `protobuf:"bytes,15,rep,name=topics,proto3" json:"topics,omitempty"`
	Visibility     Wallpaper_Visibility   `protobuf:"varint,16,opt,name=visibility,proto3,enum=wibuwallpaper.v1.Wallpaper_Visibility" json:"visibility,omitempty"`
	SourceUrl      string                 `protobuf:"bytes,17,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	Height         float32                `protobuf:"fixed32,18,opt,name=height,proto3" json:"height,omitempty"`
	Width          float32                `protobuf:"fixed32,19,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *Wallpaper) Reset() {
	*x = Wallpaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallpaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallpaper) ProtoMessage() {}

func (x *Wallpaper) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_proto_msgTypes[0]
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
	return file_wibuwallpaper_v1_wallpaper_proto_rawDescGZIP(), []int{0}
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

func (x *Wallpaper) GetTotalViews() int64 {
	if x != nil {
		return x.TotalViews
	}
	return 0
}

func (x *Wallpaper) GetTotalLikes() int64 {
	if x != nil {
		return x.TotalLikes
	}
	return 0
}

func (x *Wallpaper) GetTotalDownloads() int64 {
	if x != nil {
		return x.TotalDownloads
	}
	return 0
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

func (x *Wallpaper) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
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

var File_wibuwallpaper_v1_wallpaper_proto protoreflect.FileDescriptor

var file_wibuwallpaper_v1_wallpaper_proto_rawDesc = []byte{
	0x0a, 0x20, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x06, 0x0a, 0x09, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56,
	0x69, 0x65, 0x77, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x62, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x68, 0x64, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x64, 0x55, 0x72, 0x6c,
	0x12, 0x37, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x13, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22,
	0x60, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a,
	0x1e, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49, 0x53, 0x49,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x42, 0xc3, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61,
	0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x57, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02,
	0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wibuwallpaper_v1_wallpaper_proto_rawDescOnce sync.Once
	file_wibuwallpaper_v1_wallpaper_proto_rawDescData = file_wibuwallpaper_v1_wallpaper_proto_rawDesc
)

func file_wibuwallpaper_v1_wallpaper_proto_rawDescGZIP() []byte {
	file_wibuwallpaper_v1_wallpaper_proto_rawDescOnce.Do(func() {
		file_wibuwallpaper_v1_wallpaper_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuwallpaper_v1_wallpaper_proto_rawDescData)
	})
	return file_wibuwallpaper_v1_wallpaper_proto_rawDescData
}

var file_wibuwallpaper_v1_wallpaper_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wibuwallpaper_v1_wallpaper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wibuwallpaper_v1_wallpaper_proto_goTypes = []interface{}{
	(Wallpaper_Visibility)(0),     // 0: wibuwallpaper.v1.Wallpaper.Visibility
	(*Wallpaper)(nil),             // 1: wibuwallpaper.v1.Wallpaper
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Profile)(nil),               // 3: wibuwallpaper.v1.Profile
	(*Topic)(nil),                 // 4: wibuwallpaper.v1.Topic
}
var file_wibuwallpaper_v1_wallpaper_proto_depIdxs = []int32{
	2, // 0: wibuwallpaper.v1.Wallpaper.published_time:type_name -> google.protobuf.Timestamp
	2, // 1: wibuwallpaper.v1.Wallpaper.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: wibuwallpaper.v1.Wallpaper.update_time:type_name -> google.protobuf.Timestamp
	3, // 3: wibuwallpaper.v1.Wallpaper.publisher:type_name -> wibuwallpaper.v1.Profile
	4, // 4: wibuwallpaper.v1.Wallpaper.topics:type_name -> wibuwallpaper.v1.Topic
	0, // 5: wibuwallpaper.v1.Wallpaper.visibility:type_name -> wibuwallpaper.v1.Wallpaper.Visibility
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_wibuwallpaper_v1_wallpaper_proto_init() }
func file_wibuwallpaper_v1_wallpaper_proto_init() {
	if File_wibuwallpaper_v1_wallpaper_proto != nil {
		return
	}
	file_wibuwallpaper_v1_topic_proto_init()
	file_wibuwallpaper_v1_profile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuwallpaper_v1_wallpaper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wibuwallpaper_v1_wallpaper_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuwallpaper_v1_wallpaper_proto_goTypes,
		DependencyIndexes: file_wibuwallpaper_v1_wallpaper_proto_depIdxs,
		EnumInfos:         file_wibuwallpaper_v1_wallpaper_proto_enumTypes,
		MessageInfos:      file_wibuwallpaper_v1_wallpaper_proto_msgTypes,
	}.Build()
	File_wibuwallpaper_v1_wallpaper_proto = out.File
	file_wibuwallpaper_v1_wallpaper_proto_rawDesc = nil
	file_wibuwallpaper_v1_wallpaper_proto_goTypes = nil
	file_wibuwallpaper_v1_wallpaper_proto_depIdxs = nil
}
