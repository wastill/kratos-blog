// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: pagination.proto

package pagination

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 排序类型
type SortOrder int32

const (
	// 不排序
	SortOrder_UNSORTED SortOrder = 0
	// 升序
	SortOrder_ASCENDING SortOrder = 1
	// 降序
	SortOrder_DESCENDING SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "UNSORTED",
		1: "ASCENDING",
		2: "DESCENDING",
	}
	SortOrder_value = map[string]int32{
		"UNSORTED":   0,
		"ASCENDING":  1,
		"DESCENDING": 2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_pagination_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_pagination_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_pagination_proto_rawDescGZIP(), []int{0}
}

// 分页通用请求
type PagingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前页
	Page *int32 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	// 每一页的行数
	PageSize *int32 `protobuf:"varint,2,opt,name=pageSize,proto3,oneof" json:"pageSize,omitempty"`
	// 查询参数
	Query map[string]string `protobuf:"bytes,3,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 排序
	OrderBy map[string]string `protobuf:"bytes,4,rep,name=orderBy,proto3" json:"orderBy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 是否不分页
	Nopaging *bool `protobuf:"varint,5,opt,name=nopaging,proto3,oneof" json:"nopaging,omitempty"`
}

func (x *PagingRequest) Reset() {
	*x = PagingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingRequest) ProtoMessage() {}

func (x *PagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingRequest.ProtoReflect.Descriptor instead.
func (*PagingRequest) Descriptor() ([]byte, []int) {
	return file_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PagingRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *PagingRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *PagingRequest) GetQuery() map[string]string {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PagingRequest) GetOrderBy() map[string]string {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *PagingRequest) GetNopaging() bool {
	if x != nil && x.Nopaging != nil {
		return *x.Nopaging
	}
	return false
}

// 分页通用结果
type PagingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 分页数据
	Items []*anypb.Any `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PagingResponse) Reset() {
	*x = PagingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingResponse) ProtoMessage() {}

func (x *PagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingResponse.ProtoReflect.Descriptor instead.
func (*PagingResponse) Descriptor() ([]byte, []int) {
	return file_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PagingResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagingResponse) GetItems() []*anypb.Any {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pagination_proto protoreflect.FileDescriptor

var file_pagination_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x0d, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x40, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x08, 0x6e, 0x6f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x08, 0x6e, 0x6f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x52, 0x0a,
	0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x2a, 0x38, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x4e, 0x53, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44,
	0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x2f, 0x5a, 0x2d, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pagination_proto_rawDescOnce sync.Once
	file_pagination_proto_rawDescData = file_pagination_proto_rawDesc
)

func file_pagination_proto_rawDescGZIP() []byte {
	file_pagination_proto_rawDescOnce.Do(func() {
		file_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_pagination_proto_rawDescData)
	})
	return file_pagination_proto_rawDescData
}

var file_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pagination_proto_goTypes = []interface{}{
	(SortOrder)(0),         // 0: pagination.SortOrder
	(*PagingRequest)(nil),  // 1: pagination.PagingRequest
	(*PagingResponse)(nil), // 2: pagination.PagingResponse
	nil,                    // 3: pagination.PagingRequest.QueryEntry
	nil,                    // 4: pagination.PagingRequest.OrderByEntry
	(*anypb.Any)(nil),      // 5: google.protobuf.Any
}
var file_pagination_proto_depIdxs = []int32{
	3, // 0: pagination.PagingRequest.query:type_name -> pagination.PagingRequest.QueryEntry
	4, // 1: pagination.PagingRequest.orderBy:type_name -> pagination.PagingRequest.OrderByEntry
	5, // 2: pagination.PagingResponse.items:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pagination_proto_init() }
func file_pagination_proto_init() {
	if File_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingRequest); i {
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
		file_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingResponse); i {
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
	file_pagination_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagination_proto_goTypes,
		DependencyIndexes: file_pagination_proto_depIdxs,
		EnumInfos:         file_pagination_proto_enumTypes,
		MessageInfos:      file_pagination_proto_msgTypes,
	}.Build()
	File_pagination_proto = out.File
	file_pagination_proto_rawDesc = nil
	file_pagination_proto_goTypes = nil
	file_pagination_proto_depIdxs = nil
}
