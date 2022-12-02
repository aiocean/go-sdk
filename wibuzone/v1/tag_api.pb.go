// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wibuzone/v1/tag_api.proto

package wibuzonev1

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

type ListTagsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term string `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *ListTagsFilter) Reset() {
	*x = ListTagsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_tag_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsFilter) ProtoMessage() {}

func (x *ListTagsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_tag_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsFilter.ProtoReflect.Descriptor instead.
func (*ListTagsFilter) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_tag_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListTagsFilter) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type ListTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter    *ListTagsFilter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	PageToken string          `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  uint64          `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListTagsRequest) Reset() {
	*x = ListTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_tag_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsRequest) ProtoMessage() {}

func (x *ListTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_tag_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsRequest.ProtoReflect.Descriptor instead.
func (*ListTagsRequest) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_tag_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListTagsRequest) GetFilter() *ListTagsFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListTagsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTagsRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags          []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTagsResponse) Reset() {
	*x = ListTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_tag_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsResponse) ProtoMessage() {}

func (x *ListTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_tag_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsResponse.ProtoReflect.Descriptor instead.
func (*ListTagsResponse) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_tag_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListTagsResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ListTagsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_tag_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_tag_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_tag_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetTagResponse) Reset() {
	*x = GetTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wibuzone_v1_tag_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagResponse) ProtoMessage() {}

func (x *GetTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wibuzone_v1_tag_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagResponse.ProtoReflect.Descriptor instead.
func (*GetTagResponse) Descriptor() ([]byte, []int) {
	return file_wibuzone_v1_tag_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

var File_wibuzone_v1_tag_api_proto protoreflect.FileDescriptor

var file_wibuzone_v1_tag_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x69, 0x62,
	0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x24, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x62, 0x75,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x60, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77,
	0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77,
	0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x42, 0x9d, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x62, 0x75,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x54, 0x61, 0x67, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x77, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69,
	0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02,
	0x0b, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x57,
	0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x57, 0x69, 0x62,
	0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x57, 0x69, 0x62, 0x75, 0x7a, 0x6f, 0x6e, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wibuzone_v1_tag_api_proto_rawDescOnce sync.Once
	file_wibuzone_v1_tag_api_proto_rawDescData = file_wibuzone_v1_tag_api_proto_rawDesc
)

func file_wibuzone_v1_tag_api_proto_rawDescGZIP() []byte {
	file_wibuzone_v1_tag_api_proto_rawDescOnce.Do(func() {
		file_wibuzone_v1_tag_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_wibuzone_v1_tag_api_proto_rawDescData)
	})
	return file_wibuzone_v1_tag_api_proto_rawDescData
}

var file_wibuzone_v1_tag_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wibuzone_v1_tag_api_proto_goTypes = []interface{}{
	(*ListTagsFilter)(nil),   // 0: wibuzone.v1.ListTagsFilter
	(*ListTagsRequest)(nil),  // 1: wibuzone.v1.ListTagsRequest
	(*ListTagsResponse)(nil), // 2: wibuzone.v1.ListTagsResponse
	(*GetTagRequest)(nil),    // 3: wibuzone.v1.GetTagRequest
	(*GetTagResponse)(nil),   // 4: wibuzone.v1.GetTagResponse
	(*Tag)(nil),              // 5: wibuzone.v1.Tag
}
var file_wibuzone_v1_tag_api_proto_depIdxs = []int32{
	0, // 0: wibuzone.v1.ListTagsRequest.filter:type_name -> wibuzone.v1.ListTagsFilter
	5, // 1: wibuzone.v1.ListTagsResponse.tags:type_name -> wibuzone.v1.Tag
	5, // 2: wibuzone.v1.GetTagResponse.tag:type_name -> wibuzone.v1.Tag
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wibuzone_v1_tag_api_proto_init() }
func file_wibuzone_v1_tag_api_proto_init() {
	if File_wibuzone_v1_tag_api_proto != nil {
		return
	}
	file_wibuzone_v1_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wibuzone_v1_tag_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsFilter); i {
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
		file_wibuzone_v1_tag_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsRequest); i {
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
		file_wibuzone_v1_tag_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsResponse); i {
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
		file_wibuzone_v1_tag_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_wibuzone_v1_tag_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagResponse); i {
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
			RawDescriptor: file_wibuzone_v1_tag_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wibuzone_v1_tag_api_proto_goTypes,
		DependencyIndexes: file_wibuzone_v1_tag_api_proto_depIdxs,
		MessageInfos:      file_wibuzone_v1_tag_api_proto_msgTypes,
	}.Build()
	File_wibuzone_v1_tag_api_proto = out.File
	file_wibuzone_v1_tag_api_proto_rawDesc = nil
	file_wibuzone_v1_tag_api_proto_goTypes = nil
	file_wibuzone_v1_tag_api_proto_depIdxs = nil
}