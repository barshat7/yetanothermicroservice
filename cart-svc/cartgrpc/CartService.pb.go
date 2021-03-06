// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: CartService.proto

package cartgrpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string             `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	CartItems []*CartItemRequest `protobuf:"bytes,2,rep,name=cartItems,proto3" json:"cartItems,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCartRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateCartRequest) GetCartItems() []*CartItemRequest {
	if x != nil {
		return x.CartItems
	}
	return nil
}

type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID      string              `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	CreatedDate string              `protobuf:"bytes,3,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	CartItems   []*CartItemResponse `protobuf:"bytes,4,rep,name=cartItems,proto3" json:"cartItems,omitempty"`
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{1}
}

func (x *CartResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CartResponse) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *CartResponse) GetCartItems() []*CartItemResponse {
	if x != nil {
		return x.CartItems
	}
	return nil
}

type CartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID    int64   `protobuf:"varint,1,opt,name=itemID,proto3" json:"itemID,omitempty"`
	ItemName  string  `protobuf:"bytes,2,opt,name=itemName,proto3" json:"itemName,omitempty"`
	ItemPrice float64 `protobuf:"fixed64,3,opt,name=itemPrice,proto3" json:"itemPrice,omitempty"`
	Quantity  int64   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItemRequest) Reset() {
	*x = CartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemRequest) ProtoMessage() {}

func (x *CartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemRequest.ProtoReflect.Descriptor instead.
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{2}
}

func (x *CartItemRequest) GetItemID() int64 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *CartItemRequest) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *CartItemRequest) GetItemPrice() float64 {
	if x != nil {
		return x.ItemPrice
	}
	return 0
}

func (x *CartItemRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CartItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CartID    int64   `protobuf:"varint,2,opt,name=cartID,proto3" json:"cartID,omitempty"`
	ItemID    int64   `protobuf:"varint,3,opt,name=itemID,proto3" json:"itemID,omitempty"`
	ItemName  string  `protobuf:"bytes,4,opt,name=itemName,proto3" json:"itemName,omitempty"`
	ItemPrice float64 `protobuf:"fixed64,5,opt,name=itemPrice,proto3" json:"itemPrice,omitempty"`
	Quantity  int64   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItemResponse) Reset() {
	*x = CartItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemResponse) ProtoMessage() {}

func (x *CartItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemResponse.ProtoReflect.Descriptor instead.
func (*CartItemResponse) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{3}
}

func (x *CartItemResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartItemResponse) GetCartID() int64 {
	if x != nil {
		return x.CartID
	}
	return 0
}

func (x *CartItemResponse) GetItemID() int64 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *CartItemResponse) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *CartItemResponse) GetItemPrice() float64 {
	if x != nil {
		return x.ItemPrice
	}
	return 0
}

func (x *CartItemResponse) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CartID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CartID) Reset() {
	*x = CartID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartID) ProtoMessage() {}

func (x *CartID) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartID.ProtoReflect.Descriptor instead.
func (*CartID) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{4}
}

func (x *CartID) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{5}
}

func (x *UserID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AddCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string             `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	CartItems []*CartItemRequest `protobuf:"bytes,2,rep,name=cartItems,proto3" json:"cartItems,omitempty"`
}

func (x *AddCartItemRequest) Reset() {
	*x = AddCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CartService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartItemRequest) ProtoMessage() {}

func (x *AddCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CartService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartItemRequest.ProtoReflect.Descriptor instead.
func (*AddCartItemRequest) Descriptor() ([]byte, []int) {
	return file_CartService_proto_rawDescGZIP(), []int{6}
}

func (x *AddCartItemRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AddCartItemRequest) GetCartItems() []*CartItemRequest {
	if x != nil {
		return x.CartItems
	}
	return nil
}

var File_CartService_proto protoreflect.FileDescriptor

var file_CartService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70, 0x63, 0x22, 0x64, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7f, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xa8, 0x01, 0x0a, 0x10, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x1e, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xcd, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x61, 0x64,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_CartService_proto_rawDescOnce sync.Once
	file_CartService_proto_rawDescData = file_CartService_proto_rawDesc
)

func file_CartService_proto_rawDescGZIP() []byte {
	file_CartService_proto_rawDescOnce.Do(func() {
		file_CartService_proto_rawDescData = protoimpl.X.CompressGZIP(file_CartService_proto_rawDescData)
	})
	return file_CartService_proto_rawDescData
}

var file_CartService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_CartService_proto_goTypes = []interface{}{
	(*CreateCartRequest)(nil),  // 0: cartgrpc.CreateCartRequest
	(*CartResponse)(nil),       // 1: cartgrpc.CartResponse
	(*CartItemRequest)(nil),    // 2: cartgrpc.CartItemRequest
	(*CartItemResponse)(nil),   // 3: cartgrpc.CartItemResponse
	(*CartID)(nil),             // 4: cartgrpc.CartID
	(*UserID)(nil),             // 5: cartgrpc.UserID
	(*AddCartItemRequest)(nil), // 6: cartgrpc.AddCartItemRequest
}
var file_CartService_proto_depIdxs = []int32{
	2, // 0: cartgrpc.CreateCartRequest.cartItems:type_name -> cartgrpc.CartItemRequest
	3, // 1: cartgrpc.CartResponse.cartItems:type_name -> cartgrpc.CartItemResponse
	2, // 2: cartgrpc.AddCartItemRequest.cartItems:type_name -> cartgrpc.CartItemRequest
	0, // 3: cartgrpc.CartService.createCart:input_type -> cartgrpc.CreateCartRequest
	5, // 4: cartgrpc.CartService.retrieveCartByUserID:input_type -> cartgrpc.UserID
	6, // 5: cartgrpc.CartService.addItemToCart:input_type -> cartgrpc.AddCartItemRequest
	4, // 6: cartgrpc.CartService.createCart:output_type -> cartgrpc.CartID
	1, // 7: cartgrpc.CartService.retrieveCartByUserID:output_type -> cartgrpc.CartResponse
	4, // 8: cartgrpc.CartService.addItemToCart:output_type -> cartgrpc.CartID
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_CartService_proto_init() }
func file_CartService_proto_init() {
	if File_CartService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CartService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartRequest); i {
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
		file_CartService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartResponse); i {
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
		file_CartService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemRequest); i {
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
		file_CartService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemResponse); i {
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
		file_CartService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartID); i {
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
		file_CartService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_CartService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartItemRequest); i {
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
			RawDescriptor: file_CartService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_CartService_proto_goTypes,
		DependencyIndexes: file_CartService_proto_depIdxs,
		MessageInfos:      file_CartService_proto_msgTypes,
	}.Build()
	File_CartService_proto = out.File
	file_CartService_proto_rawDesc = nil
	file_CartService_proto_goTypes = nil
	file_CartService_proto_depIdxs = nil
}
