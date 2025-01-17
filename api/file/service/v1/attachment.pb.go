// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: attachment.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/third_party/pagination"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Path      *string `protobuf:"bytes,3,opt,name=path,proto3,oneof" json:"path,omitempty"`
	FileKey   *string `protobuf:"bytes,4,opt,name=fileKey,proto3,oneof" json:"fileKey,omitempty"`
	ThumbPath *string `protobuf:"bytes,5,opt,name=thumbPath,proto3,oneof" json:"thumbPath,omitempty"`
	MediaType *string `protobuf:"bytes,6,opt,name=mediaType,proto3,oneof" json:"mediaType,omitempty"`
	Suffix    *string `protobuf:"bytes,7,opt,name=suffix,proto3,oneof" json:"suffix,omitempty"`
	Width     *int32  `protobuf:"varint,8,opt,name=width,proto3,oneof" json:"width,omitempty"`
	Height    *int32  `protobuf:"varint,9,opt,name=height,proto3,oneof" json:"height,omitempty"`
	Size      *int64  `protobuf:"varint,10,opt,name=size,proto3,oneof" json:"size,omitempty"`
	Type      *int32  `protobuf:"varint,11,opt,name=type,proto3,oneof" json:"type,omitempty"`
	CreatedAt *string `protobuf:"bytes,12,opt,name=createdAt,proto3,oneof" json:"createdAt,omitempty"`
	UpdatedAt *string `protobuf:"bytes,13,opt,name=updatedAt,proto3,oneof" json:"updatedAt,omitempty"`
	DeletedAt *string `protobuf:"bytes,14,opt,name=deletedAt,proto3,oneof" json:"deletedAt,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *Attachment) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Attachment) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Attachment) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *Attachment) GetFileKey() string {
	if x != nil && x.FileKey != nil {
		return *x.FileKey
	}
	return ""
}

func (x *Attachment) GetThumbPath() string {
	if x != nil && x.ThumbPath != nil {
		return *x.ThumbPath
	}
	return ""
}

func (x *Attachment) GetMediaType() string {
	if x != nil && x.MediaType != nil {
		return *x.MediaType
	}
	return ""
}

func (x *Attachment) GetSuffix() string {
	if x != nil && x.Suffix != nil {
		return *x.Suffix
	}
	return ""
}

func (x *Attachment) GetWidth() int32 {
	if x != nil && x.Width != nil {
		return *x.Width
	}
	return 0
}

func (x *Attachment) GetHeight() int32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

func (x *Attachment) GetSize() int64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *Attachment) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *Attachment) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

func (x *Attachment) GetUpdatedAt() string {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return ""
}

func (x *Attachment) GetDeletedAt() string {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return ""
}

// 回应 - 附件列表
type ListAttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Attachment `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListAttachmentResponse) Reset() {
	*x = ListAttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentResponse) ProtoMessage() {}

func (x *ListAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentResponse.ProtoReflect.Descriptor instead.
func (*ListAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{1}
}

func (x *ListAttachmentResponse) GetItems() []*Attachment {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListAttachmentResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 请求 - 附件数据
type GetAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAttachmentRequest) Reset() {
	*x = GetAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentRequest) ProtoMessage() {}

func (x *GetAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentRequest.ProtoReflect.Descriptor instead.
func (*GetAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{2}
}

func (x *GetAttachmentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 请求 - 创建附件
type CreateAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachment *Attachment `protobuf:"bytes,1,opt,name=attachment,proto3" json:"attachment,omitempty"`
	OperatorId *uint32     `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *CreateAttachmentRequest) Reset() {
	*x = CreateAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttachmentRequest) ProtoMessage() {}

func (x *CreateAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttachmentRequest.ProtoReflect.Descriptor instead.
func (*CreateAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAttachmentRequest) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *CreateAttachmentRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 更新附件
type UpdateAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Attachment *Attachment `protobuf:"bytes,2,opt,name=attachment,proto3" json:"attachment,omitempty"`
	OperatorId *uint32     `protobuf:"varint,3,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *UpdateAttachmentRequest) Reset() {
	*x = UpdateAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttachmentRequest) ProtoMessage() {}

func (x *UpdateAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttachmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAttachmentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAttachmentRequest) GetAttachment() *Attachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *UpdateAttachmentRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 删除附件
type DeleteAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *DeleteAttachmentRequest) Reset() {
	*x = DeleteAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAttachmentRequest) ProtoMessage() {}

func (x *DeleteAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAttachmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAttachmentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteAttachmentRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_attachment_proto protoreflect.FileDescriptor

var file_attachment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb9, 0x04, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x50, 0x61,
	0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x07, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0a, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0c, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x50, 0x61, 0x74, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x61,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x32, 0xd4, 0x03, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attachment_proto_rawDescOnce sync.Once
	file_attachment_proto_rawDescData = file_attachment_proto_rawDesc
)

func file_attachment_proto_rawDescGZIP() []byte {
	file_attachment_proto_rawDescOnce.Do(func() {
		file_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_attachment_proto_rawDescData)
	})
	return file_attachment_proto_rawDescData
}

var file_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_attachment_proto_goTypes = []interface{}{
	(*Attachment)(nil),               // 0: file.service.v1.Attachment
	(*ListAttachmentResponse)(nil),   // 1: file.service.v1.ListAttachmentResponse
	(*GetAttachmentRequest)(nil),     // 2: file.service.v1.GetAttachmentRequest
	(*CreateAttachmentRequest)(nil),  // 3: file.service.v1.CreateAttachmentRequest
	(*UpdateAttachmentRequest)(nil),  // 4: file.service.v1.UpdateAttachmentRequest
	(*DeleteAttachmentRequest)(nil),  // 5: file.service.v1.DeleteAttachmentRequest
	(*pagination.PagingRequest)(nil), // 6: pagination.PagingRequest
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_attachment_proto_depIdxs = []int32{
	0, // 0: file.service.v1.ListAttachmentResponse.items:type_name -> file.service.v1.Attachment
	0, // 1: file.service.v1.CreateAttachmentRequest.attachment:type_name -> file.service.v1.Attachment
	0, // 2: file.service.v1.UpdateAttachmentRequest.attachment:type_name -> file.service.v1.Attachment
	6, // 3: file.service.v1.AttachmentService.ListAttachment:input_type -> pagination.PagingRequest
	2, // 4: file.service.v1.AttachmentService.GetAttachment:input_type -> file.service.v1.GetAttachmentRequest
	3, // 5: file.service.v1.AttachmentService.CreateAttachment:input_type -> file.service.v1.CreateAttachmentRequest
	4, // 6: file.service.v1.AttachmentService.UpdateAttachment:input_type -> file.service.v1.UpdateAttachmentRequest
	5, // 7: file.service.v1.AttachmentService.DeleteAttachment:input_type -> file.service.v1.DeleteAttachmentRequest
	1, // 8: file.service.v1.AttachmentService.ListAttachment:output_type -> file.service.v1.ListAttachmentResponse
	0, // 9: file.service.v1.AttachmentService.GetAttachment:output_type -> file.service.v1.Attachment
	0, // 10: file.service.v1.AttachmentService.CreateAttachment:output_type -> file.service.v1.Attachment
	0, // 11: file.service.v1.AttachmentService.UpdateAttachment:output_type -> file.service.v1.Attachment
	7, // 12: file.service.v1.AttachmentService.DeleteAttachment:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_attachment_proto_init() }
func file_attachment_proto_init() {
	if File_attachment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attachment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_attachment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentResponse); i {
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
		file_attachment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentRequest); i {
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
		file_attachment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttachmentRequest); i {
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
		file_attachment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttachmentRequest); i {
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
		file_attachment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAttachmentRequest); i {
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
	file_attachment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_attachment_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_attachment_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_attachment_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attachment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attachment_proto_goTypes,
		DependencyIndexes: file_attachment_proto_depIdxs,
		MessageInfos:      file_attachment_proto_msgTypes,
	}.Build()
	File_attachment_proto = out.File
	file_attachment_proto_rawDesc = nil
	file_attachment_proto_goTypes = nil
	file_attachment_proto_depIdxs = nil
}
