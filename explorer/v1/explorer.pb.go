// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: explorer/v1/explorer.proto

package explorer

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

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt       int32  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AssetId         string `protobuf:"bytes,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount          string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender          string `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	Memo            string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	TransactionHash string `protobuf:"bytes,7,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	OutputIndex     int32  `protobuf:"varint,8,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
	HeaderAction    int32  `protobuf:"varint,9,opt,name=header_action,json=headerAction,proto3" json:"header_action,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explorer_v1_explorer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_explorer_v1_explorer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_explorer_v1_explorer_proto_rawDescGZIP(), []int{0}
}

func (x *Output) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Output) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Output) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *Output) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Output) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Output) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Output) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *Output) GetOutputIndex() int32 {
	if x != nil {
		return x.OutputIndex
	}
	return 0
}

func (x *Output) GetHeaderAction() int32 {
	if x != nil {
		return x.HeaderAction
	}
	return 0
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int32  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AssetId   string `protobuf:"bytes,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount    string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Threshold int32  `protobuf:"varint,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
	SignedTx  string `protobuf:"bytes,6,opt,name=signed_tx,json=signedTx,proto3" json:"signed_tx,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explorer_v1_explorer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_explorer_v1_explorer_proto_msgTypes[1]
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
	return file_explorer_v1_explorer_proto_rawDescGZIP(), []int{1}
}

func (x *Transfer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transfer) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
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

func (x *Transfer) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *Transfer) GetSignedTx() string {
	if x != nil {
		return x.SignedTx
	}
	return ""
}

type ListOutputsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListOutputsRequest) Reset() {
	*x = ListOutputsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explorer_v1_explorer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOutputsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOutputsRequest) ProtoMessage() {}

func (x *ListOutputsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explorer_v1_explorer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOutputsRequest.ProtoReflect.Descriptor instead.
func (*ListOutputsRequest) Descriptor() ([]byte, []int) {
	return file_explorer_v1_explorer_proto_rawDescGZIP(), []int{2}
}

func (x *ListOutputsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *ListOutputsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListOutputsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outputs []*Output `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ListOutputsResponse) Reset() {
	*x = ListOutputsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explorer_v1_explorer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOutputsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOutputsResponse) ProtoMessage() {}

func (x *ListOutputsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explorer_v1_explorer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOutputsResponse.ProtoReflect.Descriptor instead.
func (*ListOutputsResponse) Descriptor() ([]byte, []int) {
	return file_explorer_v1_explorer_proto_rawDescGZIP(), []int{3}
}

func (x *ListOutputsResponse) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	OutputIndex     int32  `protobuf:"varint,2,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explorer_v1_explorer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explorer_v1_explorer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_explorer_v1_explorer_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionRequest) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *GetTransactionRequest) GetOutputIndex() int32 {
	if x != nil {
		return x.OutputIndex
	}
	return 0
}

type GetTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outputs   []*Output   `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Transfers []*Transfer `protobuf:"bytes,2,rep,name=transfers,proto3" json:"transfers,omitempty"`
}

func (x *GetTransactionResponse) Reset() {
	*x = GetTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explorer_v1_explorer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResponse) ProtoMessage() {}

func (x *GetTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explorer_v1_explorer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return file_explorer_v1_explorer_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransactionResponse) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *GetTransactionResponse) GetTransfers() []*Transfer {
	if x != nil {
		return x.Transfers
	}
	return nil
}

var File_explorer_v1_explorer_proto protoreflect.FileDescriptor

var file_explorer_v1_explorer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x89, 0x02, 0x0a, 0x06, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x78, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x22,
	0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x44, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x65, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x7c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x32, 0xbe,
	0x01, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f,
	0x78, 0x2d, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x61, 0x6e, 0x64, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_explorer_v1_explorer_proto_rawDescOnce sync.Once
	file_explorer_v1_explorer_proto_rawDescData = file_explorer_v1_explorer_proto_rawDesc
)

func file_explorer_v1_explorer_proto_rawDescGZIP() []byte {
	file_explorer_v1_explorer_proto_rawDescOnce.Do(func() {
		file_explorer_v1_explorer_proto_rawDescData = protoimpl.X.CompressGZIP(file_explorer_v1_explorer_proto_rawDescData)
	})
	return file_explorer_v1_explorer_proto_rawDescData
}

var file_explorer_v1_explorer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_explorer_v1_explorer_proto_goTypes = []interface{}{
	(*Output)(nil),                 // 0: explorer.v1.Output
	(*Transfer)(nil),               // 1: explorer.v1.Transfer
	(*ListOutputsRequest)(nil),     // 2: explorer.v1.ListOutputsRequest
	(*ListOutputsResponse)(nil),    // 3: explorer.v1.ListOutputsResponse
	(*GetTransactionRequest)(nil),  // 4: explorer.v1.GetTransactionRequest
	(*GetTransactionResponse)(nil), // 5: explorer.v1.GetTransactionResponse
}
var file_explorer_v1_explorer_proto_depIdxs = []int32{
	0, // 0: explorer.v1.ListOutputsResponse.outputs:type_name -> explorer.v1.Output
	0, // 1: explorer.v1.GetTransactionResponse.outputs:type_name -> explorer.v1.Output
	1, // 2: explorer.v1.GetTransactionResponse.transfers:type_name -> explorer.v1.Transfer
	2, // 3: explorer.v1.ExplorerService.ListOutputs:input_type -> explorer.v1.ListOutputsRequest
	4, // 4: explorer.v1.ExplorerService.GetTransaction:input_type -> explorer.v1.GetTransactionRequest
	3, // 5: explorer.v1.ExplorerService.ListOutputs:output_type -> explorer.v1.ListOutputsResponse
	5, // 6: explorer.v1.ExplorerService.GetTransaction:output_type -> explorer.v1.GetTransactionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_explorer_v1_explorer_proto_init() }
func file_explorer_v1_explorer_proto_init() {
	if File_explorer_v1_explorer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_explorer_v1_explorer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_explorer_v1_explorer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_explorer_v1_explorer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOutputsRequest); i {
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
		file_explorer_v1_explorer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOutputsResponse); i {
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
		file_explorer_v1_explorer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRequest); i {
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
		file_explorer_v1_explorer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionResponse); i {
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
			RawDescriptor: file_explorer_v1_explorer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_explorer_v1_explorer_proto_goTypes,
		DependencyIndexes: file_explorer_v1_explorer_proto_depIdxs,
		MessageInfos:      file_explorer_v1_explorer_proto_msgTypes,
	}.Build()
	File_explorer_v1_explorer_proto = out.File
	file_explorer_v1_explorer_proto_rawDesc = nil
	file_explorer_v1_explorer_proto_goTypes = nil
	file_explorer_v1_explorer_proto_depIdxs = nil
}
