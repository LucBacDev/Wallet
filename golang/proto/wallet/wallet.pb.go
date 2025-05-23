// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: golang/proto/wallet/wallet.proto

package wallet

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

type GetUserByAccountNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber string `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
}

func (x *GetUserByAccountNumberRequest) Reset() {
	*x = GetUserByAccountNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByAccountNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByAccountNumberRequest) ProtoMessage() {}

func (x *GetUserByAccountNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByAccountNumberRequest.ProtoReflect.Descriptor instead.
func (*GetUserByAccountNumberRequest) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByAccountNumberRequest) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

type GetUserByAccountNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId        int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AccountNumber string `protobuf:"bytes,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
}

func (x *GetUserByAccountNumberResponse) Reset() {
	*x = GetUserByAccountNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByAccountNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByAccountNumberResponse) ProtoMessage() {}

func (x *GetUserByAccountNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByAccountNumberResponse.ProtoReflect.Descriptor instead.
func (*GetUserByAccountNumberResponse) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByAccountNumberResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetUserByAccountNumberResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserByAccountNumberResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserByAccountNumberResponse) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

type DebitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DebitRequest) Reset() {
	*x = DebitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebitRequest) ProtoMessage() {}

func (x *DebitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebitRequest.ProtoReflect.Descriptor instead.
func (*DebitRequest) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *DebitRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DebitRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreditRequest) Reset() {
	*x = CreditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditRequest) ProtoMessage() {}

func (x *CreditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditRequest.ProtoReflect.Descriptor instead.
func (*CreditRequest) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *CreditRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreditRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type RefundDebitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *RefundDebitRequest) Reset() {
	*x = RefundDebitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundDebitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundDebitRequest) ProtoMessage() {}

func (x *RefundDebitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundDebitRequest.ProtoReflect.Descriptor instead.
func (*RefundDebitRequest) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *RefundDebitRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RefundDebitRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RefundDebitRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type UndoCreditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *UndoCreditRequest) Reset() {
	*x = UndoCreditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UndoCreditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UndoCreditRequest) ProtoMessage() {}

func (x *UndoCreditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UndoCreditRequest.ProtoReflect.Descriptor instead.
func (*UndoCreditRequest) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *UndoCreditRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UndoCreditRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UndoCreditRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type DebitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DebitResponse) Reset() {
	*x = DebitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebitResponse) ProtoMessage() {}

func (x *DebitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebitResponse.ProtoReflect.Descriptor instead.
func (*DebitResponse) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *DebitResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DebitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreditResponse) Reset() {
	*x = CreditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditResponse) ProtoMessage() {}

func (x *CreditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditResponse.ProtoReflect.Descriptor instead.
func (*CreditResponse) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *CreditResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreditResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RefundDebitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RefundDebitResponse) Reset() {
	*x = RefundDebitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundDebitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundDebitResponse) ProtoMessage() {}

func (x *RefundDebitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundDebitResponse.ProtoReflect.Descriptor instead.
func (*RefundDebitResponse) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *RefundDebitResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RefundDebitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UndoCreditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UndoCreditResponse) Reset() {
	*x = UndoCreditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_proto_wallet_wallet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UndoCreditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UndoCreditResponse) ProtoMessage() {}

func (x *UndoCreditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_golang_proto_wallet_wallet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UndoCreditResponse.ProtoReflect.Descriptor instead.
func (*UndoCreditResponse) Descriptor() ([]byte, []int) {
	return file_golang_proto_wallet_wallet_proto_rawDescGZIP(), []int{9}
}

func (x *UndoCreditResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UndoCreditResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_golang_proto_wallet_wallet_proto protoreflect.FileDescriptor

var file_golang_proto_wallet_wallet_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x46, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x65,
	0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x6b, 0x0a, 0x11, 0x55, 0x6e, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x43, 0x0a, 0x0d, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x6e, 0x64, 0x6f, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x82, 0x03, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x65,
	0x62, 0x69, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x44, 0x65, 0x62, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x0a, 0x55, 0x6e, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x19, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x55, 0x6e, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x55, 0x6e, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_golang_proto_wallet_wallet_proto_rawDescOnce sync.Once
	file_golang_proto_wallet_wallet_proto_rawDescData = file_golang_proto_wallet_wallet_proto_rawDesc
)

func file_golang_proto_wallet_wallet_proto_rawDescGZIP() []byte {
	file_golang_proto_wallet_wallet_proto_rawDescOnce.Do(func() {
		file_golang_proto_wallet_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_golang_proto_wallet_wallet_proto_rawDescData)
	})
	return file_golang_proto_wallet_wallet_proto_rawDescData
}

var file_golang_proto_wallet_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_golang_proto_wallet_wallet_proto_goTypes = []interface{}{
	(*GetUserByAccountNumberRequest)(nil),  // 0: wallet.GetUserByAccountNumberRequest
	(*GetUserByAccountNumberResponse)(nil), // 1: wallet.GetUserByAccountNumberResponse
	(*DebitRequest)(nil),                   // 2: wallet.DebitRequest
	(*CreditRequest)(nil),                  // 3: wallet.CreditRequest
	(*RefundDebitRequest)(nil),             // 4: wallet.RefundDebitRequest
	(*UndoCreditRequest)(nil),              // 5: wallet.UndoCreditRequest
	(*DebitResponse)(nil),                  // 6: wallet.DebitResponse
	(*CreditResponse)(nil),                 // 7: wallet.CreditResponse
	(*RefundDebitResponse)(nil),            // 8: wallet.RefundDebitResponse
	(*UndoCreditResponse)(nil),             // 9: wallet.UndoCreditResponse
}
var file_golang_proto_wallet_wallet_proto_depIdxs = []int32{
	0, // 0: wallet.WalletService.GetUserByAccountNumber:input_type -> wallet.GetUserByAccountNumberRequest
	2, // 1: wallet.WalletService.DebitBalance:input_type -> wallet.DebitRequest
	3, // 2: wallet.WalletService.CreditBalance:input_type -> wallet.CreditRequest
	4, // 3: wallet.WalletService.RefundDebit:input_type -> wallet.RefundDebitRequest
	5, // 4: wallet.WalletService.UndoCredit:input_type -> wallet.UndoCreditRequest
	1, // 5: wallet.WalletService.GetUserByAccountNumber:output_type -> wallet.GetUserByAccountNumberResponse
	6, // 6: wallet.WalletService.DebitBalance:output_type -> wallet.DebitResponse
	7, // 7: wallet.WalletService.CreditBalance:output_type -> wallet.CreditResponse
	8, // 8: wallet.WalletService.RefundDebit:output_type -> wallet.RefundDebitResponse
	9, // 9: wallet.WalletService.UndoCredit:output_type -> wallet.UndoCreditResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_golang_proto_wallet_wallet_proto_init() }
func file_golang_proto_wallet_wallet_proto_init() {
	if File_golang_proto_wallet_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_golang_proto_wallet_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByAccountNumberRequest); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByAccountNumberResponse); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebitRequest); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditRequest); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundDebitRequest); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UndoCreditRequest); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebitResponse); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditResponse); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundDebitResponse); i {
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
		file_golang_proto_wallet_wallet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UndoCreditResponse); i {
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
			RawDescriptor: file_golang_proto_wallet_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_golang_proto_wallet_wallet_proto_goTypes,
		DependencyIndexes: file_golang_proto_wallet_wallet_proto_depIdxs,
		MessageInfos:      file_golang_proto_wallet_wallet_proto_msgTypes,
	}.Build()
	File_golang_proto_wallet_wallet_proto = out.File
	file_golang_proto_wallet_wallet_proto_rawDesc = nil
	file_golang_proto_wallet_wallet_proto_goTypes = nil
	file_golang_proto_wallet_wallet_proto_depIdxs = nil
}
