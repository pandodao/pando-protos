// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: sweeper/v1/sweeper.proto

package pigeon

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

type Transfer_Status int32

const (
	Transfer_NOT_SET Transfer_Status = 0
	Transfer_PENDING Transfer_Status = 1
	Transfer_HANDLED Transfer_Status = 2
	Transfer_SPENT   Transfer_Status = 3
)

// Enum value maps for Transfer_Status.
var (
	Transfer_Status_name = map[int32]string{
		0: "NOT_SET",
		1: "PENDING",
		2: "HANDLED",
		3: "SPENT",
	}
	Transfer_Status_value = map[string]int32{
		"NOT_SET": 0,
		"PENDING": 1,
		"HANDLED": 2,
		"SPENT":   3,
	}
)

func (x Transfer_Status) Enum() *Transfer_Status {
	p := new(Transfer_Status)
	*p = x
	return p
}

func (x Transfer_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Transfer_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_sweeper_v1_sweeper_proto_enumTypes[0].Descriptor()
}

func (Transfer_Status) Type() protoreflect.EnumType {
	return &file_sweeper_v1_sweeper_proto_enumTypes[0]
}

func (x Transfer_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Transfer_Status.Descriptor instead.
func (Transfer_Status) EnumDescriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{0, 0}
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId   string          `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	AssetId   string          `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount    string          `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Opponents []string        `protobuf:"bytes,4,rep,name=opponents,proto3" json:"opponents,omitempty"`
	Threshold uint32          `protobuf:"varint,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Memo      string          `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	Status    Transfer_Status `protobuf:"varint,7,opt,name=status,proto3,enum=sweeper.v1.Transfer_Status" json:"status,omitempty"`
	TxHash    string          `protobuf:"bytes,8,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{0}
}

func (x *Transfer) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *Transfer) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *Transfer) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transfer) GetOpponents() []string {
	if x != nil {
		return x.Opponents
	}
	return nil
}

func (x *Transfer) GetThreshold() uint32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *Transfer) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Transfer) GetStatus() Transfer_Status {
	if x != nil {
		return x.Status
	}
	return Transfer_NOT_SET
}

func (x *Transfer) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount  string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Count   int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{1}
}

func (x *Balance) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *Balance) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Balance) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SendTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId   string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	AssetId   string   `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount    string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Opponents []string `protobuf:"bytes,4,rep,name=opponents,proto3" json:"opponents,omitempty"`
	Threshold uint32   `protobuf:"varint,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Memo      string   `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *SendTransferRequest) Reset() {
	*x = SendTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransferRequest) ProtoMessage() {}

func (x *SendTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransferRequest.ProtoReflect.Descriptor instead.
func (*SendTransferRequest) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{2}
}

func (x *SendTransferRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *SendTransferRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *SendTransferRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *SendTransferRequest) GetOpponents() []string {
	if x != nil {
		return x.Opponents
	}
	return nil
}

func (x *SendTransferRequest) GetThreshold() uint32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *SendTransferRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type SendTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendTransferResponse) Reset() {
	*x = SendTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransferResponse) ProtoMessage() {}

func (x *SendTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransferResponse.ProtoReflect.Descriptor instead.
func (*SendTransferResponse) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{3}
}

type SendTransfersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfers []*SendTransferRequest `protobuf:"bytes,1,rep,name=transfers,proto3" json:"transfers,omitempty"`
}

func (x *SendTransfersRequest) Reset() {
	*x = SendTransfersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransfersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransfersRequest) ProtoMessage() {}

func (x *SendTransfersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransfersRequest.ProtoReflect.Descriptor instead.
func (*SendTransfersRequest) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{4}
}

func (x *SendTransfersRequest) GetTransfers() []*SendTransferRequest {
	if x != nil {
		return x.Transfers
	}
	return nil
}

type SendTransfersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendTransfersResponse) Reset() {
	*x = SendTransfersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransfersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransfersResponse) ProtoMessage() {}

func (x *SendTransfersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransfersResponse.ProtoReflect.Descriptor instead.
func (*SendTransfersResponse) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{5}
}

type GetTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *GetTransferRequest) Reset() {
	*x = GetTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransferRequest) ProtoMessage() {}

func (x *GetTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransferRequest.ProtoReflect.Descriptor instead.
func (*GetTransferRequest) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransferRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type GetTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfer *Transfer `protobuf:"bytes,1,opt,name=transfer,proto3" json:"transfer,omitempty"`
}

func (x *GetTransferResponse) Reset() {
	*x = GetTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransferResponse) ProtoMessage() {}

func (x *GetTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransferResponse.ProtoReflect.Descriptor instead.
func (*GetTransferResponse) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{7}
}

func (x *GetTransferResponse) GetTransfer() *Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

type GetBalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBalancesRequest) Reset() {
	*x = GetBalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancesRequest) ProtoMessage() {}

func (x *GetBalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancesRequest.ProtoReflect.Descriptor instead.
func (*GetBalancesRequest) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{8}
}

type GetBalancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balances []*Balance `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *GetBalancesResponse) Reset() {
	*x = GetBalancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sweeper_v1_sweeper_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancesResponse) ProtoMessage() {}

func (x *GetBalancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sweeper_v1_sweeper_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancesResponse.ProtoReflect.Descriptor instead.
func (*GetBalancesResponse) Descriptor() ([]byte, []int) {
	return file_sweeper_v1_sweeper_proto_rawDescGZIP(), []int{9}
}

func (x *GetBalancesResponse) GetBalances() []*Balance {
	if x != nil {
		return x.Balances
	}
	return nil
}

var File_sweeper_v1_sweeper_proto protoreflect.FileDescriptor

var file_sweeper_v1_sweeper_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x77, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x77, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xb2, 0x02, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22,
	0x3a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x22, 0x52, 0x0a, 0x07, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xb3, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a,
	0x14, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x47,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x08, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x32, 0xd9, 0x02, 0x0a, 0x0e, 0x53, 0x77, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x77, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x73,
	0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x1e, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x6f, 0x78, 0x2d, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x61, 0x6e, 0x64, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x69, 0x67, 0x65, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sweeper_v1_sweeper_proto_rawDescOnce sync.Once
	file_sweeper_v1_sweeper_proto_rawDescData = file_sweeper_v1_sweeper_proto_rawDesc
)

func file_sweeper_v1_sweeper_proto_rawDescGZIP() []byte {
	file_sweeper_v1_sweeper_proto_rawDescOnce.Do(func() {
		file_sweeper_v1_sweeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_sweeper_v1_sweeper_proto_rawDescData)
	})
	return file_sweeper_v1_sweeper_proto_rawDescData
}

var file_sweeper_v1_sweeper_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sweeper_v1_sweeper_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_sweeper_v1_sweeper_proto_goTypes = []interface{}{
	(Transfer_Status)(0),          // 0: sweeper.v1.Transfer.Status
	(*Transfer)(nil),              // 1: sweeper.v1.Transfer
	(*Balance)(nil),               // 2: sweeper.v1.Balance
	(*SendTransferRequest)(nil),   // 3: sweeper.v1.SendTransferRequest
	(*SendTransferResponse)(nil),  // 4: sweeper.v1.SendTransferResponse
	(*SendTransfersRequest)(nil),  // 5: sweeper.v1.SendTransfersRequest
	(*SendTransfersResponse)(nil), // 6: sweeper.v1.SendTransfersResponse
	(*GetTransferRequest)(nil),    // 7: sweeper.v1.GetTransferRequest
	(*GetTransferResponse)(nil),   // 8: sweeper.v1.GetTransferResponse
	(*GetBalancesRequest)(nil),    // 9: sweeper.v1.GetBalancesRequest
	(*GetBalancesResponse)(nil),   // 10: sweeper.v1.GetBalancesResponse
}
var file_sweeper_v1_sweeper_proto_depIdxs = []int32{
	0,  // 0: sweeper.v1.Transfer.status:type_name -> sweeper.v1.Transfer.Status
	3,  // 1: sweeper.v1.SendTransfersRequest.transfers:type_name -> sweeper.v1.SendTransferRequest
	1,  // 2: sweeper.v1.GetTransferResponse.transfer:type_name -> sweeper.v1.Transfer
	2,  // 3: sweeper.v1.GetBalancesResponse.balances:type_name -> sweeper.v1.Balance
	3,  // 4: sweeper.v1.SweeperService.SendTransfer:input_type -> sweeper.v1.SendTransferRequest
	5,  // 5: sweeper.v1.SweeperService.SendTransfers:input_type -> sweeper.v1.SendTransfersRequest
	7,  // 6: sweeper.v1.SweeperService.GetTransfer:input_type -> sweeper.v1.GetTransferRequest
	9,  // 7: sweeper.v1.SweeperService.GetBalances:input_type -> sweeper.v1.GetBalancesRequest
	4,  // 8: sweeper.v1.SweeperService.SendTransfer:output_type -> sweeper.v1.SendTransferResponse
	6,  // 9: sweeper.v1.SweeperService.SendTransfers:output_type -> sweeper.v1.SendTransfersResponse
	8,  // 10: sweeper.v1.SweeperService.GetTransfer:output_type -> sweeper.v1.GetTransferResponse
	10, // 11: sweeper.v1.SweeperService.GetBalances:output_type -> sweeper.v1.GetBalancesResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_sweeper_v1_sweeper_proto_init() }
func file_sweeper_v1_sweeper_proto_init() {
	if File_sweeper_v1_sweeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sweeper_v1_sweeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransferRequest); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransferResponse); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransfersRequest); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransfersResponse); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransferRequest); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransferResponse); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalancesRequest); i {
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
		file_sweeper_v1_sweeper_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalancesResponse); i {
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
			RawDescriptor: file_sweeper_v1_sweeper_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sweeper_v1_sweeper_proto_goTypes,
		DependencyIndexes: file_sweeper_v1_sweeper_proto_depIdxs,
		EnumInfos:         file_sweeper_v1_sweeper_proto_enumTypes,
		MessageInfos:      file_sweeper_v1_sweeper_proto_msgTypes,
	}.Build()
	File_sweeper_v1_sweeper_proto = out.File
	file_sweeper_v1_sweeper_proto_rawDesc = nil
	file_sweeper_v1_sweeper_proto_goTypes = nil
	file_sweeper_v1_sweeper_proto_depIdxs = nil
}
