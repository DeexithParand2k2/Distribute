// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: protos/proto/book-inventory-service.proto

package book_inventory_service

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

// request (Messages)
type BookAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
}

func (x *BookAvailabilityRequest) Reset() {
	*x = BookAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookAvailabilityRequest) ProtoMessage() {}

func (x *BookAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*BookAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{0}
}

func (x *BookAvailabilityRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type AddBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Stock  int32  `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *AddBookRequest) Reset() {
	*x = AddBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookRequest) ProtoMessage() {}

func (x *AddBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookRequest.ProtoReflect.Descriptor instead.
func (*AddBookRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *AddBookRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type BookAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available    bool   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Rented       bool   `protobuf:"varint,2,opt,name=rented,proto3" json:"rented,omitempty"`   // rented by someone
	InStock      bool   `protobuf:"varint,3,opt,name=inStock,proto3" json:"inStock,omitempty"` // no yet at library
	DateOfRent   string `protobuf:"bytes,4,opt,name=dateOfRent,proto3" json:"dateOfRent,omitempty"`
	DateOfReturn string `protobuf:"bytes,5,opt,name=dateOfReturn,proto3" json:"dateOfReturn,omitempty"` // depends on `rented` variables
}

func (x *BookAvailabilityResponse) Reset() {
	*x = BookAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookAvailabilityResponse) ProtoMessage() {}

func (x *BookAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*BookAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{2}
}

func (x *BookAvailabilityResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *BookAvailabilityResponse) GetRented() bool {
	if x != nil {
		return x.Rented
	}
	return false
}

func (x *BookAvailabilityResponse) GetInStock() bool {
	if x != nil {
		return x.InStock
	}
	return false
}

func (x *BookAvailabilityResponse) GetDateOfRent() string {
	if x != nil {
		return x.DateOfRent
	}
	return ""
}

func (x *BookAvailabilityResponse) GetDateOfReturn() string {
	if x != nil {
		return x.DateOfReturn
	}
	return ""
}

type AddBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type GetActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionId string `protobuf:"bytes,1,opt,name=actionId,proto3" json:"actionId,omitempty"`
}

func (x *GetActionRequest) Reset() {
	*x = GetActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionRequest) ProtoMessage() {}

func (x *GetActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionRequest.ProtoReflect.Descriptor instead.
func (*GetActionRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetActionRequest) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

// list the corresponding action to perform
type GetActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	ActionId string `protobuf:"bytes,2,opt,name=actionId,proto3" json:"actionId,omitempty"`
}

func (x *GetActionResponse) Reset() {
	*x = GetActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionResponse) ProtoMessage() {}

func (x *GetActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionResponse.ProtoReflect.Descriptor instead.
func (*GetActionResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetActionResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *GetActionResponse) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

type UpdateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId   string `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	ActionId string `protobuf:"bytes,2,opt,name=actionId,proto3" json:"actionId,omitempty"`
}

func (x *UpdateStockRequest) Reset() {
	*x = UpdateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockRequest) ProtoMessage() {}

func (x *UpdateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockRequest) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStockRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *UpdateStockRequest) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

type UpdateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId    string `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	ActionId  string `protobuf:"bytes,2,opt,name=actionId,proto3" json:"actionId,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *UpdateStockResponse) Reset() {
	*x = UpdateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_proto_book_inventory_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockResponse) ProtoMessage() {}

func (x *UpdateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_proto_book_inventory_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockResponse.ProtoReflect.Descriptor instead.
func (*UpdateStockResponse) Descriptor() ([]byte, []int) {
	return file_protos_proto_book_inventory_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStockResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *UpdateStockResponse) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *UpdateStockResponse) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

var File_protos_proto_book_inventory_service_proto protoreflect.FileDescriptor

var file_protos_proto_book_inventory_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2d, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x17, 0x42,
	0x6f, 0x6f, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0xae, 0x01, 0x0a, 0x18, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x29, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f,
	0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xfb, 0x01, 0x0a,
	0x14, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f,
	0x6f, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0f,
	0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_proto_book_inventory_service_proto_rawDescOnce sync.Once
	file_protos_proto_book_inventory_service_proto_rawDescData = file_protos_proto_book_inventory_service_proto_rawDesc
)

func file_protos_proto_book_inventory_service_proto_rawDescGZIP() []byte {
	file_protos_proto_book_inventory_service_proto_rawDescOnce.Do(func() {
		file_protos_proto_book_inventory_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_proto_book_inventory_service_proto_rawDescData)
	})
	return file_protos_proto_book_inventory_service_proto_rawDescData
}

var file_protos_proto_book_inventory_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_proto_book_inventory_service_proto_goTypes = []any{
	(*BookAvailabilityRequest)(nil),  // 0: BookAvailabilityRequest
	(*AddBookRequest)(nil),           // 1: AddBookRequest
	(*BookAvailabilityResponse)(nil), // 2: BookAvailabilityResponse
	(*AddBookResponse)(nil),          // 3: AddBookResponse
	(*GetActionRequest)(nil),         // 4: GetActionRequest
	(*GetActionResponse)(nil),        // 5: GetActionResponse
	(*UpdateStockRequest)(nil),       // 6: UpdateStockRequest
	(*UpdateStockResponse)(nil),      // 7: UpdateStockResponse
}
var file_protos_proto_book_inventory_service_proto_depIdxs = []int32{
	0, // 0: BookInventoryService.CheckBookAvailability:input_type -> BookAvailabilityRequest
	1, // 1: BookInventoryService.AddBook:input_type -> AddBookRequest
	4, // 2: BookInventoryService.GetActionId:input_type -> GetActionRequest
	3, // 3: BookInventoryService.UpdateStock:input_type -> AddBookResponse
	2, // 4: BookInventoryService.CheckBookAvailability:output_type -> BookAvailabilityResponse
	3, // 5: BookInventoryService.AddBook:output_type -> AddBookResponse
	5, // 6: BookInventoryService.GetActionId:output_type -> GetActionResponse
	3, // 7: BookInventoryService.UpdateStock:output_type -> AddBookResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_proto_book_inventory_service_proto_init() }
func file_protos_proto_book_inventory_service_proto_init() {
	if File_protos_proto_book_inventory_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_proto_book_inventory_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BookAvailabilityRequest); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddBookRequest); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BookAvailabilityResponse); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddBookResponse); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetActionRequest); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetActionResponse); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateStockRequest); i {
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
		file_protos_proto_book_inventory_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateStockResponse); i {
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
			RawDescriptor: file_protos_proto_book_inventory_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_proto_book_inventory_service_proto_goTypes,
		DependencyIndexes: file_protos_proto_book_inventory_service_proto_depIdxs,
		MessageInfos:      file_protos_proto_book_inventory_service_proto_msgTypes,
	}.Build()
	File_protos_proto_book_inventory_service_proto = out.File
	file_protos_proto_book_inventory_service_proto_rawDesc = nil
	file_protos_proto_book_inventory_service_proto_goTypes = nil
	file_protos_proto_book_inventory_service_proto_depIdxs = nil
}
