// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuwallpaper/v1/wallpaper_api.proto

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

type GetWallpaperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WallpaperId string `protobuf:"bytes,1,opt,name=wallpaper_id,json=wallpaperId,proto3" json:"wallpaper_id,omitempty"`
}

func (x *GetWallpaperRequest) Reset() {
	*x = GetWallpaperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWallpaperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWallpaperRequest) ProtoMessage() {}

func (x *GetWallpaperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWallpaperRequest.ProtoReflect.Descriptor instead.
func (*GetWallpaperRequest) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetWallpaperRequest) GetWallpaperId() string {
	if x != nil {
		return x.WallpaperId
	}
	return ""
}

type GetWallpaperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallpaper *Wallpaper `protobuf:"bytes,1,opt,name=wallpaper,proto3" json:"wallpaper,omitempty"`
}

func (x *GetWallpaperResponse) Reset() {
	*x = GetWallpaperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWallpaperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWallpaperResponse) ProtoMessage() {}

func (x *GetWallpaperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWallpaperResponse.ProtoReflect.Descriptor instead.
func (*GetWallpaperResponse) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetWallpaperResponse) GetWallpaper() *Wallpaper {
	if x != nil {
		return x.Wallpaper
	}
	return nil
}

type ListWallpapersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListWallpapersRequest) Reset() {
	*x = ListWallpapersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWallpapersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWallpapersRequest) ProtoMessage() {}

func (x *ListWallpapersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWallpapersRequest.ProtoReflect.Descriptor instead.
func (*ListWallpapersRequest) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListWallpapersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListWallpapersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallpapers    []*Wallpaper `protobuf:"bytes,1,rep,name=wallpapers,proto3" json:"wallpapers,omitempty"`
	NextPageToken string       `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListWallpapersResponse) Reset() {
	*x = ListWallpapersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWallpapersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWallpapersResponse) ProtoMessage() {}

func (x *ListWallpapersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWallpapersResponse.ProtoReflect.Descriptor instead.
func (*ListWallpapersResponse) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListWallpapersResponse) GetWallpapers() []*Wallpaper {
	if x != nil {
		return x.Wallpapers
	}
	return nil
}

func (x *ListWallpapersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type SaveWallpaperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description  string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SizeMb       float32 `protobuf:"fixed32,6,opt,name=size_mb,json=sizeMb,proto3" json:"size_mb,omitempty"`
	ThumbnailUrl string  `protobuf:"bytes,10,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	HdUrl        string  `protobuf:"bytes,11,opt,name=hd_url,json=hdUrl,proto3" json:"hd_url,omitempty"`
	PublisherId  string  `protobuf:"bytes,12,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	Tags         []*Tag  `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	SourceUrl    string  `protobuf:"bytes,17,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	Height       float32 `protobuf:"fixed32,18,opt,name=height,proto3" json:"height,omitempty"`
	Width        float32 `protobuf:"fixed32,19,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *SaveWallpaperRequest) Reset() {
	*x = SaveWallpaperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveWallpaperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveWallpaperRequest) ProtoMessage() {}

func (x *SaveWallpaperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveWallpaperRequest.ProtoReflect.Descriptor instead.
func (*SaveWallpaperRequest) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{4}
}

func (x *SaveWallpaperRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SaveWallpaperRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SaveWallpaperRequest) GetSizeMb() float32 {
	if x != nil {
		return x.SizeMb
	}
	return 0
}

func (x *SaveWallpaperRequest) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *SaveWallpaperRequest) GetHdUrl() string {
	if x != nil {
		return x.HdUrl
	}
	return ""
}

func (x *SaveWallpaperRequest) GetPublisherId() string {
	if x != nil {
		return x.PublisherId
	}
	return ""
}

func (x *SaveWallpaperRequest) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *SaveWallpaperRequest) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *SaveWallpaperRequest) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SaveWallpaperRequest) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

type SaveWallpaperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallpaper *Wallpaper `protobuf:"bytes,1,opt,name=wallpaper,proto3" json:"wallpaper,omitempty"`
}

func (x *SaveWallpaperResponse) Reset() {
	*x = SaveWallpaperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveWallpaperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveWallpaperResponse) ProtoMessage() {}

func (x *SaveWallpaperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveWallpaperResponse.ProtoReflect.Descriptor instead.
func (*SaveWallpaperResponse) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{5}
}

func (x *SaveWallpaperResponse) GetWallpaper() *Wallpaper {
	if x != nil {
		return x.Wallpaper
	}
	return nil
}

var File_wibuwallpaper_v1_wallpaper_api_proto protoreflect.FileDescriptor

var file_wibuwallpaper_v1_wallpaper_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61,
	0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x77, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x51, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69,
	0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb8, 0x02, 0x0a, 0x14, 0x53,
	0x61, 0x76, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x62, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x68, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x69,
	0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x13, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x52, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x57, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x09,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x42, 0xc6, 0x01, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x11, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x57,
	0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1c, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wibuwallpaper_v1_wallpaper_api_proto_rawDescOnce sync.Once
	file_wibuwallpaper_v1_wallpaper_api_proto_rawDescData = file_wibuwallpaper_v1_wallpaper_api_proto_rawDesc
)

func file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP() []byte {
	file_wibuwallpaper_v1_wallpaper_api_proto_rawDescOnce.Do(func() {
		file_wibuwallpaper_v1_wallpaper_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuwallpaper_v1_wallpaper_api_proto_rawDescData)
	})
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescData
}

var file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wibuwallpaper_v1_wallpaper_api_proto_goTypes = []interface{}{
	(*GetWallpaperRequest)(nil),    // 0: wibuwallpaper.v1.GetWallpaperRequest
	(*GetWallpaperResponse)(nil),   // 1: wibuwallpaper.v1.GetWallpaperResponse
	(*ListWallpapersRequest)(nil),  // 2: wibuwallpaper.v1.ListWallpapersRequest
	(*ListWallpapersResponse)(nil), // 3: wibuwallpaper.v1.ListWallpapersResponse
	(*SaveWallpaperRequest)(nil),   // 4: wibuwallpaper.v1.SaveWallpaperRequest
	(*SaveWallpaperResponse)(nil),  // 5: wibuwallpaper.v1.SaveWallpaperResponse
	(*Wallpaper)(nil),              // 6: wibuwallpaper.v1.Wallpaper
	(*Tag)(nil),                    // 7: wibuwallpaper.v1.Tag
}
var file_wibuwallpaper_v1_wallpaper_api_proto_depIdxs = []int32{
	6, // 0: wibuwallpaper.v1.GetWallpaperResponse.wallpaper:type_name -> wibuwallpaper.v1.Wallpaper
	6, // 1: wibuwallpaper.v1.ListWallpapersResponse.wallpapers:type_name -> wibuwallpaper.v1.Wallpaper
	7, // 2: wibuwallpaper.v1.SaveWallpaperRequest.tags:type_name -> wibuwallpaper.v1.Tag
	6, // 3: wibuwallpaper.v1.SaveWallpaperResponse.wallpaper:type_name -> wibuwallpaper.v1.Wallpaper
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wibuwallpaper_v1_wallpaper_api_proto_init() }
func file_wibuwallpaper_v1_wallpaper_api_proto_init() {
	if File_wibuwallpaper_v1_wallpaper_api_proto != nil {
		return
	}
	file_wibuwallpaper_v1_wallpaper_proto_init()
	file_wibuwallpaper_v1_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWallpaperRequest); i {
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
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWallpaperResponse); i {
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
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWallpapersRequest); i {
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
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWallpapersResponse); i {
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
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveWallpaperRequest); i {
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
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveWallpaperResponse); i {
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
			RawDescriptor: file_wibuwallpaper_v1_wallpaper_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuwallpaper_v1_wallpaper_api_proto_goTypes,
		DependencyIndexes: file_wibuwallpaper_v1_wallpaper_api_proto_depIdxs,
		MessageInfos:      file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes,
	}.Build()
	File_wibuwallpaper_v1_wallpaper_api_proto = out.File
	file_wibuwallpaper_v1_wallpaper_api_proto_rawDesc = nil
	file_wibuwallpaper_v1_wallpaper_api_proto_goTypes = nil
	file_wibuwallpaper_v1_wallpaper_api_proto_depIdxs = nil
}
