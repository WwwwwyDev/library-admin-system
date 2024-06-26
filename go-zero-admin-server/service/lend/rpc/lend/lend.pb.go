// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: lend.proto

package lend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{0}
}

type BookIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId uint64 `protobuf:"varint,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
}

func (x *BookIdReq) Reset() {
	*x = BookIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookIdReq) ProtoMessage() {}

func (x *BookIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookIdReq.ProtoReflect.Descriptor instead.
func (*BookIdReq) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{1}
}

func (x *BookIdReq) GetBookId() uint64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{2}
}

func (x *IdReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LendsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	BookName      string `protobuf:"bytes,3,opt,name=bookName,proto3" json:"bookName,omitempty"`
	VipCardNumber string `protobuf:"bytes,4,opt,name=vipCardNumber,proto3" json:"vipCardNumber,omitempty"`
}

func (x *LendsReq) Reset() {
	*x = LendsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LendsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LendsReq) ProtoMessage() {}

func (x *LendsReq) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LendsReq.ProtoReflect.Descriptor instead.
func (*LendsReq) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{3}
}

func (x *LendsReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *LendsReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LendsReq) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *LendsReq) GetVipCardNumber() string {
	if x != nil {
		return x.VipCardNumber
	}
	return ""
}

type LendAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VipId         uint64                 `protobuf:"varint,1,opt,name=vipId,proto3" json:"vipId,omitempty"`
	BookId        uint64                 `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	BookName      string                 `protobuf:"bytes,3,opt,name=bookName,proto3" json:"bookName,omitempty"`
	VipCardNumber string                 `protobuf:"bytes,4,opt,name=vipCardNumber,proto3" json:"vipCardNumber,omitempty"`
	LendTime      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=lendTime,proto3" json:"lendTime,omitempty"`
}

func (x *LendAddReq) Reset() {
	*x = LendAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LendAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LendAddReq) ProtoMessage() {}

func (x *LendAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LendAddReq.ProtoReflect.Descriptor instead.
func (*LendAddReq) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{4}
}

func (x *LendAddReq) GetVipId() uint64 {
	if x != nil {
		return x.VipId
	}
	return 0
}

func (x *LendAddReq) GetBookId() uint64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *LendAddReq) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *LendAddReq) GetVipCardNumber() string {
	if x != nil {
		return x.VipCardNumber
	}
	return ""
}

func (x *LendAddReq) GetLendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LendTime
	}
	return nil
}

type LendInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VipId         uint64                 `protobuf:"varint,2,opt,name=vipId,proto3" json:"vipId,omitempty"`
	BookId        uint64                 `protobuf:"varint,3,opt,name=bookId,proto3" json:"bookId,omitempty"`
	BookName      string                 `protobuf:"bytes,4,opt,name=bookName,proto3" json:"bookName,omitempty"`
	VipCardNumber string                 `protobuf:"bytes,5,opt,name=vipCardNumber,proto3" json:"vipCardNumber,omitempty"`
	LendTime      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=lendTime,proto3" json:"lendTime,omitempty"`
}

func (x *LendInfoReply) Reset() {
	*x = LendInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LendInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LendInfoReply) ProtoMessage() {}

func (x *LendInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LendInfoReply.ProtoReflect.Descriptor instead.
func (*LendInfoReply) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{5}
}

func (x *LendInfoReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LendInfoReply) GetVipId() uint64 {
	if x != nil {
		return x.VipId
	}
	return 0
}

func (x *LendInfoReply) GetBookId() uint64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *LendInfoReply) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *LendInfoReply) GetVipCardNumber() string {
	if x != nil {
		return x.VipCardNumber
	}
	return ""
}

func (x *LendInfoReply) GetLendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LendTime
	}
	return nil
}

type LendsInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int64            `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	LendsInfo []*LendInfoReply `protobuf:"bytes,2,rep,name=lendsInfo,proto3" json:"lendsInfo,omitempty"`
}

func (x *LendsInfoReply) Reset() {
	*x = LendsInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LendsInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LendsInfoReply) ProtoMessage() {}

func (x *LendsInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LendsInfoReply.ProtoReflect.Descriptor instead.
func (*LendsInfoReply) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{6}
}

func (x *LendsInfoReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *LendsInfoReply) GetLendsInfo() []*LendInfoReply {
	if x != nil {
		return x.LendsInfo
	}
	return nil
}

type IsLendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLend bool `protobuf:"varint,1,opt,name=isLend,proto3" json:"isLend,omitempty"`
}

func (x *IsLendReply) Reset() {
	*x = IsLendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLendReply) ProtoMessage() {}

func (x *IsLendReply) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLendReply.ProtoReflect.Descriptor instead.
func (*IsLendReply) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{7}
}

func (x *IsLendReply) GetIsLend() bool {
	if x != nil {
		return x.IsLend
	}
	return false
}

type IsExistReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExist bool `protobuf:"varint,1,opt,name=isExist,proto3" json:"isExist,omitempty"`
}

func (x *IsExistReply) Reset() {
	*x = IsExistReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsExistReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsExistReply) ProtoMessage() {}

func (x *IsExistReply) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsExistReply.ProtoReflect.Descriptor instead.
func (*IsExistReply) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{8}
}

func (x *IsExistReply) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

type IsSuccessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
}

func (x *IsSuccessReply) Reset() {
	*x = IsSuccessReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lend_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsSuccessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsSuccessReply) ProtoMessage() {}

func (x *IsSuccessReply) ProtoReflect() protoreflect.Message {
	mi := &file_lend_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsSuccessReply.ProtoReflect.Descriptor instead.
func (*IsSuccessReply) Descriptor() ([]byte, []int) {
	return file_lend_proto_rawDescGZIP(), []int{9}
}

func (x *IsSuccessReply) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_lend_proto protoreflect.FileDescriptor

var file_lend_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x65,
	0x6e, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x22,
	0x23, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x62, 0x6f,
	0x6f, 0x6b, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x76, 0x0a,
	0x08, 0x4c, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x76, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x4c, 0x65, 0x6e, 0x64, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x69, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f,
	0x6f, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x76, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x6c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc7, 0x01, 0x0a,
	0x0d, 0x4c, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x69, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x69, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x69, 0x70, 0x43,
	0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x76, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x36,
	0x0a, 0x08, 0x6c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x0e, 0x4c, 0x65, 0x6e, 0x64, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x31,
	0x0a, 0x09, 0x6c, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x6c, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x69, 0x73, 0x4c, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4c, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x4c, 0x65, 0x6e, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x69, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x04, 0x6c, 0x65, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x0f, 0x49,
	0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b,
	0x2e, 0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x65,
	0x6e, 0x64, 0x2e, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2c, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x6c, 0x65, 0x6e, 0x64,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6c, 0x65, 0x6e,
	0x64, 0x2e, 0x69, 0x73, 0x4c, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a,
	0x08, 0x67, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x0e, 0x2e, 0x6c, 0x65, 0x6e, 0x64,
	0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6c, 0x65, 0x6e, 0x64,
	0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x31, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x6c, 0x65, 0x6e,
	0x64, 0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6c,
	0x65, 0x6e, 0x64, 0x2e, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x4c, 0x65, 0x6e, 0x64, 0x12, 0x0b, 0x2e,
	0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6c, 0x65, 0x6e,
	0x64, 0x2e, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lend_proto_rawDescOnce sync.Once
	file_lend_proto_rawDescData = file_lend_proto_rawDesc
)

func file_lend_proto_rawDescGZIP() []byte {
	file_lend_proto_rawDescOnce.Do(func() {
		file_lend_proto_rawDescData = protoimpl.X.CompressGZIP(file_lend_proto_rawDescData)
	})
	return file_lend_proto_rawDescData
}

var file_lend_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_lend_proto_goTypes = []interface{}{
	(*EmptyReq)(nil),              // 0: lend.EmptyReq
	(*BookIdReq)(nil),             // 1: lend.BookIdReq
	(*IdReq)(nil),                 // 2: lend.IdReq
	(*LendsReq)(nil),              // 3: lend.LendsReq
	(*LendAddReq)(nil),            // 4: lend.LendAddReq
	(*LendInfoReply)(nil),         // 5: lend.LendInfoReply
	(*LendsInfoReply)(nil),        // 6: lend.LendsInfoReply
	(*IsLendReply)(nil),           // 7: lend.isLendReply
	(*IsExistReply)(nil),          // 8: lend.isExistReply
	(*IsSuccessReply)(nil),        // 9: lend.isSuccessReply
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_lend_proto_depIdxs = []int32{
	10, // 0: lend.LendAddReq.lendTime:type_name -> google.protobuf.Timestamp
	10, // 1: lend.LendInfoReply.lendTime:type_name -> google.protobuf.Timestamp
	5,  // 2: lend.LendsInfoReply.lendsInfo:type_name -> lend.LendInfoReply
	2,  // 3: lend.lend.IsExistLendById:input_type -> lend.IdReq
	1,  // 4: lend.lend.IsLend:input_type -> lend.BookIdReq
	3,  // 5: lend.lend.getLends:input_type -> lend.LendsReq
	4,  // 6: lend.lend.AddLend:input_type -> lend.LendAddReq
	2,  // 7: lend.lend.DelLend:input_type -> lend.IdReq
	8,  // 8: lend.lend.IsExistLendById:output_type -> lend.isExistReply
	7,  // 9: lend.lend.IsLend:output_type -> lend.isLendReply
	6,  // 10: lend.lend.getLends:output_type -> lend.LendsInfoReply
	9,  // 11: lend.lend.AddLend:output_type -> lend.isSuccessReply
	9,  // 12: lend.lend.DelLend:output_type -> lend.isSuccessReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_lend_proto_init() }
func file_lend_proto_init() {
	if File_lend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_lend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookIdReq); i {
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
		file_lend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_lend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LendsReq); i {
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
		file_lend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LendAddReq); i {
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
		file_lend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LendInfoReply); i {
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
		file_lend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LendsInfoReply); i {
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
		file_lend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLendReply); i {
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
		file_lend_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsExistReply); i {
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
		file_lend_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsSuccessReply); i {
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
			RawDescriptor: file_lend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lend_proto_goTypes,
		DependencyIndexes: file_lend_proto_depIdxs,
		MessageInfos:      file_lend_proto_msgTypes,
	}.Build()
	File_lend_proto = out.File
	file_lend_proto_rawDesc = nil
	file_lend_proto_goTypes = nil
	file_lend_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LendClient is the client API for Lend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LendClient interface {
	IsExistLendById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IsExistReply, error)
	IsLend(ctx context.Context, in *BookIdReq, opts ...grpc.CallOption) (*IsLendReply, error)
	GetLends(ctx context.Context, in *LendsReq, opts ...grpc.CallOption) (*LendsInfoReply, error)
	AddLend(ctx context.Context, in *LendAddReq, opts ...grpc.CallOption) (*IsSuccessReply, error)
	DelLend(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IsSuccessReply, error)
}

type lendClient struct {
	cc grpc.ClientConnInterface
}

func NewLendClient(cc grpc.ClientConnInterface) LendClient {
	return &lendClient{cc}
}

func (c *lendClient) IsExistLendById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IsExistReply, error) {
	out := new(IsExistReply)
	err := c.cc.Invoke(ctx, "/lend.lend/IsExistLendById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) IsLend(ctx context.Context, in *BookIdReq, opts ...grpc.CallOption) (*IsLendReply, error) {
	out := new(IsLendReply)
	err := c.cc.Invoke(ctx, "/lend.lend/IsLend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) GetLends(ctx context.Context, in *LendsReq, opts ...grpc.CallOption) (*LendsInfoReply, error) {
	out := new(LendsInfoReply)
	err := c.cc.Invoke(ctx, "/lend.lend/getLends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) AddLend(ctx context.Context, in *LendAddReq, opts ...grpc.CallOption) (*IsSuccessReply, error) {
	out := new(IsSuccessReply)
	err := c.cc.Invoke(ctx, "/lend.lend/AddLend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) DelLend(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IsSuccessReply, error) {
	out := new(IsSuccessReply)
	err := c.cc.Invoke(ctx, "/lend.lend/DelLend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LendServer is the server API for Lend service.
type LendServer interface {
	IsExistLendById(context.Context, *IdReq) (*IsExistReply, error)
	IsLend(context.Context, *BookIdReq) (*IsLendReply, error)
	GetLends(context.Context, *LendsReq) (*LendsInfoReply, error)
	AddLend(context.Context, *LendAddReq) (*IsSuccessReply, error)
	DelLend(context.Context, *IdReq) (*IsSuccessReply, error)
}

// UnimplementedLendServer can be embedded to have forward compatible implementations.
type UnimplementedLendServer struct {
}

func (*UnimplementedLendServer) IsExistLendById(context.Context, *IdReq) (*IsExistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExistLendById not implemented")
}
func (*UnimplementedLendServer) IsLend(context.Context, *BookIdReq) (*IsLendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLend not implemented")
}
func (*UnimplementedLendServer) GetLends(context.Context, *LendsReq) (*LendsInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLends not implemented")
}
func (*UnimplementedLendServer) AddLend(context.Context, *LendAddReq) (*IsSuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLend not implemented")
}
func (*UnimplementedLendServer) DelLend(context.Context, *IdReq) (*IsSuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelLend not implemented")
}

func RegisterLendServer(s *grpc.Server, srv LendServer) {
	s.RegisterService(&_Lend_serviceDesc, srv)
}

func _Lend_IsExistLendById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendServer).IsExistLendById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lend.lend/IsExistLendById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendServer).IsExistLendById(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lend_IsLend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendServer).IsLend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lend.lend/IsLend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendServer).IsLend(ctx, req.(*BookIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lend_GetLends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LendsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendServer).GetLends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lend.lend/GetLends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendServer).GetLends(ctx, req.(*LendsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lend_AddLend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LendAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendServer).AddLend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lend.lend/AddLend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendServer).AddLend(ctx, req.(*LendAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lend_DelLend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendServer).DelLend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lend.lend/DelLend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendServer).DelLend(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lend.lend",
	HandlerType: (*LendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsExistLendById",
			Handler:    _Lend_IsExistLendById_Handler,
		},
		{
			MethodName: "IsLend",
			Handler:    _Lend_IsLend_Handler,
		},
		{
			MethodName: "getLends",
			Handler:    _Lend_GetLends_Handler,
		},
		{
			MethodName: "AddLend",
			Handler:    _Lend_AddLend_Handler,
		},
		{
			MethodName: "DelLend",
			Handler:    _Lend_DelLend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lend.proto",
}
