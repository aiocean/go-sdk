// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuzone/v1/distiller.proto

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

type DistilledWebpage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author      string                 `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Url         string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	PublishedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	TextContent string                 `protobuf:"bytes,4,opt,name=text_content,json=textContent,proto3" json:"text_content,omitempty"`
	HtmlContent string                 `protobuf:"bytes,5,opt,name=html_content,json=htmlContent,proto3" json:"html_content,omitempty"`
}

func (x *DistilledWebpage) Reset() {
	*x = DistilledWebpage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_distiller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistilledWebpage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistilledWebpage) ProtoMessage() {}

func (x *DistilledWebpage) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_distiller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistilledWebpage.ProtoReflect.Descriptor instead.
func (*DistilledWebpage) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_distiller_proto_rawDescGZIP(), []int{0}
}

func (x *DistilledWebpage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DistilledWebpage) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *DistilledWebpage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DistilledWebpage) GetPublishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

func (x *DistilledWebpage) GetTextContent() string {
	if x != nil {
		return x.TextContent
	}
	return ""
}

func (x *DistilledWebpage) GetHtmlContent() string {
	if x != nil {
		return x.HtmlContent
	}
	return ""
}

type DistilledWebpageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DistilledWebpageRequest) Reset() {
	*x = DistilledWebpageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_distiller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistilledWebpageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistilledWebpageRequest) ProtoMessage() {}

func (x *DistilledWebpageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_distiller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistilledWebpageRequest.ProtoReflect.Descriptor instead.
func (*DistilledWebpageRequest) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_distiller_proto_rawDescGZIP(), []int{1}
}

func (x *DistilledWebpageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DistilledWebpageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistilledWebpage *DistilledWebpage `protobuf:"bytes,1,opt,name=distilled_webpage,json=distilledWebpage,proto3" json:"distilled_webpage,omitempty"`
}

func (x *DistilledWebpageResponse) Reset() {
	*x = DistilledWebpageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_distiller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistilledWebpageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistilledWebpageResponse) ProtoMessage() {}

func (x *DistilledWebpageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_distiller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistilledWebpageResponse.ProtoReflect.Descriptor instead.
func (*DistilledWebpageResponse) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_distiller_proto_rawDescGZIP(), []int{2}
}

func (x *DistilledWebpageResponse) GetDistilledWebpage() *DistilledWebpage {
	if x != nil {
		return x.DistilledWebpage
	}
	return nil
}

var File_wibuzone_v1_distiller_proto protoreflect.FileDescriptor

var file_wibuzone_v1_distiller_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77,
	0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x10,
	0x44, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x57, 0x65, 0x62, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x74, 0x6d, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x57, 0x65, 0x62, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x66, 0x0a, 0x18, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x57,
	0x65, 0x62, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x11, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x77, 0x65, 0x62, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x69, 0x62, 0x75,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65,
	0x64, 0x57, 0x65, 0x62, 0x70, 0x61, 0x67, 0x65, 0x52, 0x10, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x57, 0x65, 0x62, 0x70, 0x61, 0x67, 0x65, 0x42, 0xa0, 0x01, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e,
	0x44, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f,
	0x63, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x77, 0x69, 0x62, 0x75,
	0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x57, 0x69, 0x62, 0x75, 0x7a,
	0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wibuzone_v1_distiller_proto_rawDescOnce sync.Once
	file_wibuzone_v1_distiller_proto_rawDescData = file_wibuzone_v1_distiller_proto_rawDesc
)

func file_wibuzone_v1_distiller_proto_rawDescGZIP() []byte {
	file_wibuzone_v1_distiller_proto_rawDescOnce.Do(func() {
		file_wibuzone_v1_distiller_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuzone_v1_distiller_proto_rawDescData)
	})
	return file_wibuzone_v1_distiller_proto_rawDescData
}

var file_wibuzone_v1_distiller_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wibuzone_v1_distiller_proto_goTypes = []interface{}{
	(*DistilledWebpage)(nil),         // 0: wibuzone.v1.DistilledWebpage
	(*DistilledWebpageRequest)(nil),  // 1: wibuzone.v1.DistilledWebpageRequest
	(*DistilledWebpageResponse)(nil), // 2: wibuzone.v1.DistilledWebpageResponse
	(*timestamppb.Timestamp)(nil),    // 3: google.protobuf.Timestamp
}
var file_wibuzone_v1_distiller_proto_depIdxs = []int32{
	3, // 0: wibuzone.v1.DistilledWebpage.published_at:type_name -> google.protobuf.Timestamp
	0, // 1: wibuzone.v1.DistilledWebpageResponse.distilled_webpage:type_name -> wibuzone.v1.DistilledWebpage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wibuzone_v1_distiller_proto_init() }
func file_wibuzone_v1_distiller_proto_init() {
	if File_wibuzone_v1_distiller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wibuzone_v1_distiller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistilledWebpage); i {
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
		file_wibuzone_v1_distiller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistilledWebpageRequest); i {
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
		file_wibuzone_v1_distiller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistilledWebpageResponse); i {
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
			RawDescriptor: file_wibuzone_v1_distiller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuzone_v1_distiller_proto_goTypes,
		DependencyIndexes: file_wibuzone_v1_distiller_proto_depIdxs,
		MessageInfos:      file_wibuzone_v1_distiller_proto_msgTypes,
	}.Build()
	File_wibuzone_v1_distiller_proto = out.File
	file_wibuzone_v1_distiller_proto_rawDesc = nil
	file_wibuzone_v1_distiller_proto_goTypes = nil
	file_wibuzone_v1_distiller_proto_depIdxs = nil
}
