// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pinterest/v1/pinterest.proto

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

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName      string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Id             string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Username       string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	FullName       string `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	ImageMediumUrl string `protobuf:"bytes,5,opt,name=image_medium_url,json=imageMediumUrl,proto3" json:"image_medium_url,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinterest_v1_pinterest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_pinterest_v1_pinterest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_pinterest_v1_pinterest_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Profile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Profile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Profile) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Profile) GetImageMediumUrl() string {
	if x != nil {
		return x.ImageMediumUrl
	}
	return ""
}

type VideoPin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url           string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ThumbnailUrl  string   `protobuf:"bytes,5,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	VideoUrl      string   `protobuf:"bytes,6,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	StaticUrl     string   `protobuf:"bytes,7,opt,name=static_url,json=staticUrl,proto3" json:"static_url,omitempty"`
	Width         int32    `protobuf:"varint,8,opt,name=width,proto3" json:"width,omitempty"`
	Height        int32    `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	Duration      int32    `protobuf:"varint,10,opt,name=duration,proto3" json:"duration,omitempty"`
	OriginPinner  *Profile `protobuf:"bytes,11,opt,name=origin_pinner,json=originPinner,proto3" json:"origin_pinner,omitempty"`
	DominantColor string   `protobuf:"bytes,12,opt,name=dominant_color,json=dominantColor,proto3" json:"dominant_color,omitempty"`
}

func (x *VideoPin) Reset() {
	*x = VideoPin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinterest_v1_pinterest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPin) ProtoMessage() {}

func (x *VideoPin) ProtoReflect() protoreflect.Message {
	mi := &file_pinterest_v1_pinterest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPin.ProtoReflect.Descriptor instead.
func (*VideoPin) Descriptor() ([]byte, []int) {
	return file_pinterest_v1_pinterest_proto_rawDescGZIP(), []int{1}
}

func (x *VideoPin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VideoPin) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoPin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VideoPin) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoPin) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *VideoPin) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *VideoPin) GetStaticUrl() string {
	if x != nil {
		return x.StaticUrl
	}
	return ""
}

func (x *VideoPin) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VideoPin) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VideoPin) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *VideoPin) GetOriginPinner() *Profile {
	if x != nil {
		return x.OriginPinner
	}
	return nil
}

func (x *VideoPin) GetDominantColor() string {
	if x != nil {
		return x.DominantColor
	}
	return ""
}

var File_pinterest_v1_pinterest_proto protoreflect.FileDescriptor

var file_pinterest_v1_pinterest_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x9b, 0x01, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x55, 0x72, 0x6c, 0x22, 0xf2, 0x02, 0x0a, 0x08, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x50, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0d,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x42,
	0xa7, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x0c, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0c, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x18, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x50, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pinterest_v1_pinterest_proto_rawDescOnce sync.Once
	file_pinterest_v1_pinterest_proto_rawDescData = file_pinterest_v1_pinterest_proto_rawDesc
)

func file_pinterest_v1_pinterest_proto_rawDescGZIP() []byte {
	file_pinterest_v1_pinterest_proto_rawDescOnce.Do(func() {
		file_pinterest_v1_pinterest_proto_rawDescData = protoimpl.X.CompressGZIP(file_pinterest_v1_pinterest_proto_rawDescData)
	})
	return file_pinterest_v1_pinterest_proto_rawDescData
}

var file_pinterest_v1_pinterest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pinterest_v1_pinterest_proto_goTypes = []interface{}{
	(*Profile)(nil),  // 0: pinterest.v1.Profile
	(*VideoPin)(nil), // 1: pinterest.v1.VideoPin
}
var file_pinterest_v1_pinterest_proto_depIdxs = []int32{
	0, // 0: pinterest.v1.VideoPin.origin_pinner:type_name -> pinterest.v1.Profile
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pinterest_v1_pinterest_proto_init() }
func file_pinterest_v1_pinterest_proto_init() {
	if File_pinterest_v1_pinterest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pinterest_v1_pinterest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_pinterest_v1_pinterest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPin); i {
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
			RawDescriptor: file_pinterest_v1_pinterest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pinterest_v1_pinterest_proto_goTypes,
		DependencyIndexes: file_pinterest_v1_pinterest_proto_depIdxs,
		MessageInfos:      file_pinterest_v1_pinterest_proto_msgTypes,
	}.Build()
	File_pinterest_v1_pinterest_proto = out.File
	file_pinterest_v1_pinterest_proto_rawDesc = nil
	file_pinterest_v1_pinterest_proto_goTypes = nil
	file_pinterest_v1_pinterest_proto_depIdxs = nil
}
