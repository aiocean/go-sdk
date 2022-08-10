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
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWallpapersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWallpapersResponse) ProtoMessage() {}

func (x *ListWallpapersResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListWallpapersResponse.ProtoReflect.Descriptor instead.
func (*ListWallpapersResponse) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{0}
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

type CreateWallpaperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallpaper *Wallpaper `protobuf:"bytes,1,opt,name=wallpaper,proto3" json:"wallpaper,omitempty"`
}

func (x *CreateWallpaperRequest) Reset() {
	*x = CreateWallpaperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWallpaperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWallpaperRequest) ProtoMessage() {}

func (x *CreateWallpaperRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateWallpaperRequest.ProtoReflect.Descriptor instead.
func (*CreateWallpaperRequest) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWallpaperRequest) GetWallpaper() *Wallpaper {
	if x != nil {
		return x.Wallpaper
	}
	return nil
}

type CreateWallpaperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallpaper *Wallpaper `protobuf:"bytes,1,opt,name=wallpaper,proto3" json:"wallpaper,omitempty"`
}

func (x *CreateWallpaperResponse) Reset() {
	*x = CreateWallpaperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWallpaperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWallpaperResponse) ProtoMessage() {}

func (x *CreateWallpaperResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateWallpaperResponse.ProtoReflect.Descriptor instead.
func (*CreateWallpaperResponse) Descriptor() ([]byte, []int) {
	return file_wibuwallpaper_v1_wallpaper_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWallpaperResponse) GetWallpaper() *Wallpaper {
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
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77,
	0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x53, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x22, 0x54,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x77, 0x61, 0x6c,
	0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77,
	0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x42, 0xc6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x62,
	0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x57,
	0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x77, 0x69,
	0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77,
	0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x57, 0x69, 0x62, 0x75, 0x77, 0x61,
	0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x57, 0x69, 0x62,
	0x75, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x57, 0x69, 0x62, 0x75,
	0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wibuwallpaper_v1_wallpaper_api_proto_goTypes = []interface{}{
	(*ListWallpapersResponse)(nil),  // 0: wibuwallpaper.v1.ListWallpapersResponse
	(*CreateWallpaperRequest)(nil),  // 1: wibuwallpaper.v1.CreateWallpaperRequest
	(*CreateWallpaperResponse)(nil), // 2: wibuwallpaper.v1.CreateWallpaperResponse
	(*Wallpaper)(nil),               // 3: wibuwallpaper.v1.Wallpaper
}
var file_wibuwallpaper_v1_wallpaper_api_proto_depIdxs = []int32{
	3, // 0: wibuwallpaper.v1.ListWallpapersResponse.wallpapers:type_name -> wibuwallpaper.v1.Wallpaper
	3, // 1: wibuwallpaper.v1.CreateWallpaperRequest.wallpaper:type_name -> wibuwallpaper.v1.Wallpaper
	3, // 2: wibuwallpaper.v1.CreateWallpaperResponse.wallpaper:type_name -> wibuwallpaper.v1.Wallpaper
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wibuwallpaper_v1_wallpaper_api_proto_init() }
func file_wibuwallpaper_v1_wallpaper_api_proto_init() {
	if File_wibuwallpaper_v1_wallpaper_api_proto != nil {
		return
	}
	file_wibuwallpaper_v1_wallpaper_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_wibuwallpaper_v1_wallpaper_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWallpaperRequest); i {
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
			switch v := v.(*CreateWallpaperResponse); i {
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
			NumMessages:   3,
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
