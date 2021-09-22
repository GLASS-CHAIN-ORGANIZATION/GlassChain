// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.1
// source: qbftNode.proto

package types

import (
	context "context"
	reflect "reflect"
	sync "sync"

	types "github.com/33cn/chain33/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type QbftNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Power  int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *QbftNode) Reset() {
	*x = QbftNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QbftNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QbftNode) ProtoMessage() {}

func (x *QbftNode) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QbftNode.ProtoReflect.Descriptor instead.
func (*QbftNode) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{0}
}

func (x *QbftNode) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *QbftNode) GetPower() int64 {
	if x != nil {
		return x.Power
	}
	return 0
}

type QbftNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*QbftNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *QbftNodes) Reset() {
	*x = QbftNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QbftNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QbftNodes) ProtoMessage() {}

func (x *QbftNodes) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QbftNodes.ProtoReflect.Descriptor instead.
func (*QbftNodes) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{1}
}

func (x *QbftNodes) GetNodes() []*QbftNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type QbftNodeAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*QbftNodeAction_Node
	//	*QbftNodeAction_BlockInfo
	Value isQbftNodeAction_Value `protobuf_oneof:"value"`
	Ty    int32                  `protobuf:"varint,3,opt,name=Ty,proto3" json:"Ty,omitempty"`
}

func (x *QbftNodeAction) Reset() {
	*x = QbftNodeAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QbftNodeAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QbftNodeAction) ProtoMessage() {}

func (x *QbftNodeAction) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QbftNodeAction.ProtoReflect.Descriptor instead.
func (*QbftNodeAction) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{2}
}

func (m *QbftNodeAction) GetValue() isQbftNodeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *QbftNodeAction) GetNode() *QbftNode {
	if x, ok := x.GetValue().(*QbftNodeAction_Node); ok {
		return x.Node
	}
	return nil
}

func (x *QbftNodeAction) GetBlockInfo() *QbftBlockInfo {
	if x, ok := x.GetValue().(*QbftNodeAction_BlockInfo); ok {
		return x.BlockInfo
	}
	return nil
}

func (x *QbftNodeAction) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type isQbftNodeAction_Value interface {
	isQbftNodeAction_Value()
}

type QbftNodeAction_Node struct {
	Node *QbftNode `protobuf:"bytes,1,opt,name=node,proto3,oneof"`
}

type QbftNodeAction_BlockInfo struct {
	BlockInfo *QbftBlockInfo `protobuf:"bytes,2,opt,name=blockInfo,proto3,oneof"`
}

func (*QbftNodeAction_Node) isQbftNodeAction_Value() {}

func (*QbftNodeAction_BlockInfo) isQbftNodeAction_Value() {}

type ReqQbftNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ReqQbftNodes) Reset() {
	*x = ReqQbftNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQbftNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQbftNodes) ProtoMessage() {}

func (x *ReqQbftNodes) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQbftNodes.ProtoReflect.Descriptor instead.
func (*ReqQbftNodes) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{3}
}

func (x *ReqQbftNodes) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type ReqQbftBlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ReqQbftBlockInfo) Reset() {
	*x = ReqQbftBlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQbftBlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQbftBlockInfo) ProtoMessage() {}

func (x *ReqQbftBlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQbftBlockInfo.ProtoReflect.Descriptor instead.
func (*ReqQbftBlockInfo) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{4}
}

func (x *ReqQbftBlockInfo) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type QbftNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIP      string `protobuf:"bytes,1,opt,name=nodeIP,proto3" json:"nodeIP,omitempty"`
	NodeID      string `protobuf:"bytes,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PubKey      string `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	VotingPower int64  `protobuf:"varint,5,opt,name=votingPower,proto3" json:"votingPower,omitempty"`
	Accum       int64  `protobuf:"varint,6,opt,name=accum,proto3" json:"accum,omitempty"`
}

func (x *QbftNodeInfo) Reset() {
	*x = QbftNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QbftNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QbftNodeInfo) ProtoMessage() {}

func (x *QbftNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QbftNodeInfo.ProtoReflect.Descriptor instead.
func (*QbftNodeInfo) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{5}
}

func (x *QbftNodeInfo) GetNodeIP() string {
	if x != nil {
		return x.NodeIP
	}
	return ""
}

func (x *QbftNodeInfo) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *QbftNodeInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QbftNodeInfo) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *QbftNodeInfo) GetVotingPower() int64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

func (x *QbftNodeInfo) GetAccum() int64 {
	if x != nil {
		return x.Accum
	}
	return 0
}

type QbftNodeInfoSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*QbftNodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *QbftNodeInfoSet) Reset() {
	*x = QbftNodeInfoSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QbftNodeInfoSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QbftNodeInfoSet) ProtoMessage() {}

func (x *QbftNodeInfoSet) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QbftNodeInfoSet.ProtoReflect.Descriptor instead.
func (*QbftNodeInfoSet) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{6}
}

func (x *QbftNodeInfoSet) GetNodes() []*QbftNodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type QbftPerfStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalTx     int64 `protobuf:"varint,1,opt,name=totalTx,proto3" json:"totalTx,omitempty"`
	TotalBlock  int64 `protobuf:"varint,2,opt,name=totalBlock,proto3" json:"totalBlock,omitempty"`
	TxPerBlock  int64 `protobuf:"varint,3,opt,name=txPerBlock,proto3" json:"txPerBlock,omitempty"`
	TotalSecond int64 `protobuf:"varint,4,opt,name=totalSecond,proto3" json:"totalSecond,omitempty"`
	TxPerSecond int64 `protobuf:"varint,5,opt,name=txPerSecond,proto3" json:"txPerSecond,omitempty"`
}

func (x *QbftPerfStat) Reset() {
	*x = QbftPerfStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QbftPerfStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QbftPerfStat) ProtoMessage() {}

func (x *QbftPerfStat) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QbftPerfStat.ProtoReflect.Descriptor instead.
func (*QbftPerfStat) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{7}
}

func (x *QbftPerfStat) GetTotalTx() int64 {
	if x != nil {
		return x.TotalTx
	}
	return 0
}

func (x *QbftPerfStat) GetTotalBlock() int64 {
	if x != nil {
		return x.TotalBlock
	}
	return 0
}

func (x *QbftPerfStat) GetTxPerBlock() int64 {
	if x != nil {
		return x.TxPerBlock
	}
	return 0
}

func (x *QbftPerfStat) GetTotalSecond() int64 {
	if x != nil {
		return x.TotalSecond
	}
	return 0
}

func (x *QbftPerfStat) GetTxPerSecond() int64 {
	if x != nil {
		return x.TxPerSecond
	}
	return 0
}

type ReqQbftPerfStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ReqQbftPerfStat) Reset() {
	*x = ReqQbftPerfStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qbftNode_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQbftPerfStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQbftPerfStat) ProtoMessage() {}

func (x *ReqQbftPerfStat) ProtoReflect() protoreflect.Message {
	mi := &file_qbftNode_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQbftPerfStat.ProtoReflect.Descriptor instead.
func (*ReqQbftPerfStat) Descriptor() ([]byte, []int) {
	return file_qbftNode_proto_rawDescGZIP(), []int{8}
}

func (x *ReqQbftPerfStat) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ReqQbftPerfStat) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_qbftNode_proto protoreflect.FileDescriptor

var file_qbftNode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x71, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x71, 0x62, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x38, 0x0a, 0x08, 0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x09, 0x51,
	0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22,
	0x86, 0x01, 0x0a, 0x0e, 0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x62, 0x66, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x54, 0x79, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x51,
	0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x51, 0x62, 0x66, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xa8, 0x01, 0x0a,
	0x0c, 0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x22, 0x3c, 0x0a, 0x0f, 0x51, 0x62, 0x66, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x51, 0x62, 0x66, 0x74, 0x50, 0x65,
	0x72, 0x66, 0x53, 0x74, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x78,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x78, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x78, 0x50, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x78, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x78, 0x50, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x22, 0x39, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x51, 0x62, 0x66, 0x74, 0x50,
	0x65, 0x72, 0x66, 0x53, 0x74, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x32,
	0x73, 0x0a, 0x08, 0x71, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x49,
	0x73, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x4e, 0x69, 0x6c, 0x1a, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x62, 0x66,
	0x74, 0x49, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x4e, 0x69, 0x6c, 0x1a, 0x16, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x51, 0x62, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x74, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qbftNode_proto_rawDescOnce sync.Once
	file_qbftNode_proto_rawDescData = file_qbftNode_proto_rawDesc
)

func file_qbftNode_proto_rawDescGZIP() []byte {
	file_qbftNode_proto_rawDescOnce.Do(func() {
		file_qbftNode_proto_rawDescData = protoimpl.X.CompressGZIP(file_qbftNode_proto_rawDescData)
	})
	return file_qbftNode_proto_rawDescData
}

var file_qbftNode_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_qbftNode_proto_goTypes = []interface{}{
	(*QbftNode)(nil),         // 0: types.QbftNode
	(*QbftNodes)(nil),        // 1: types.QbftNodes
	(*QbftNodeAction)(nil),   // 2: types.QbftNodeAction
	(*ReqQbftNodes)(nil),     // 3: types.ReqQbftNodes
	(*ReqQbftBlockInfo)(nil), // 4: types.ReqQbftBlockInfo
	(*QbftNodeInfo)(nil),     // 5: types.QbftNodeInfo
	(*QbftNodeInfoSet)(nil),  // 6: types.QbftNodeInfoSet
	(*QbftPerfStat)(nil),     // 7: types.QbftPerfStat
	(*ReqQbftPerfStat)(nil),  // 8: types.ReqQbftPerfStat
	(*QbftBlockInfo)(nil),    // 9: types.QbftBlockInfo
	(*types.ReqNil)(nil),     // 10: types.ReqNil
	(*QbftIsHealthy)(nil),    // 11: types.QbftIsHealthy
}
var file_qbftNode_proto_depIdxs = []int32{
	0,  // 0: types.QbftNodes.nodes:type_name -> types.QbftNode
	0,  // 1: types.QbftNodeAction.node:type_name -> types.QbftNode
	9,  // 2: types.QbftNodeAction.blockInfo:type_name -> types.QbftBlockInfo
	5,  // 3: types.QbftNodeInfoSet.nodes:type_name -> types.QbftNodeInfo
	10, // 4: types.qbftNode.IsSync:input_type -> types.ReqNil
	10, // 5: types.qbftNode.GetNodeInfo:input_type -> types.ReqNil
	11, // 6: types.qbftNode.IsSync:output_type -> types.QbftIsHealthy
	6,  // 7: types.qbftNode.GetNodeInfo:output_type -> types.QbftNodeInfoSet
	6,  // [6:8] is the sub-list for method output_type
	4,  // [4:6] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_qbftNode_proto_init() }
func file_qbftNode_proto_init() {
	if File_qbftNode_proto != nil {
		return
	}
	file_qbft_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qbftNode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QbftNode); i {
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
		file_qbftNode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QbftNodes); i {
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
		file_qbftNode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QbftNodeAction); i {
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
		file_qbftNode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQbftNodes); i {
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
		file_qbftNode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQbftBlockInfo); i {
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
		file_qbftNode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QbftNodeInfo); i {
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
		file_qbftNode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QbftNodeInfoSet); i {
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
		file_qbftNode_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QbftPerfStat); i {
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
		file_qbftNode_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQbftPerfStat); i {
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
	file_qbftNode_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QbftNodeAction_Node)(nil),
		(*QbftNodeAction_BlockInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qbftNode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qbftNode_proto_goTypes,
		DependencyIndexes: file_qbftNode_proto_depIdxs,
		MessageInfos:      file_qbftNode_proto_msgTypes,
	}.Build()
	File_qbftNode_proto = out.File
	file_qbftNode_proto_rawDesc = nil
	file_qbftNode_proto_goTypes = nil
	file_qbftNode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QbftNodeClient is the client API for QbftNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QbftNodeClient interface {
	IsSync(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftIsHealthy, error)
	GetNodeInfo(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftNodeInfoSet, error)
}

type qbftNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewQbftNodeClient(cc grpc.ClientConnInterface) QbftNodeClient {
	return &qbftNodeClient{cc}
}

func (c *qbftNodeClient) IsSync(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftIsHealthy, error) {
	out := new(QbftIsHealthy)
	err := c.cc.Invoke(ctx, "/types.qbftNode/IsSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qbftNodeClient) GetNodeInfo(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftNodeInfoSet, error) {
	out := new(QbftNodeInfoSet)
	err := c.cc.Invoke(ctx, "/types.qbftNode/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QbftNodeServer is the server API for QbftNode service.
type QbftNodeServer interface {
	IsSync(context.Context, *types.ReqNil) (*QbftIsHealthy, error)
	GetNodeInfo(context.Context, *types.ReqNil) (*QbftNodeInfoSet, error)
}

// UnimplementedQbftNodeServer can be embedded to have forward compatible implementations.
type UnimplementedQbftNodeServer struct {
}

func (*UnimplementedQbftNodeServer) IsSync(context.Context, *types.ReqNil) (*QbftIsHealthy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSync not implemented")
}
func (*UnimplementedQbftNodeServer) GetNodeInfo(context.Context, *types.ReqNil) (*QbftNodeInfoSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}

func RegisterQbftNodeServer(s *grpc.Server, srv QbftNodeServer) {
	s.RegisterService(&_QbftNode_serviceDesc, srv)
}

func _QbftNode_IsSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QbftNodeServer).IsSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.qbftNode/IsSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QbftNodeServer).IsSync(ctx, req.(*types.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _QbftNode_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QbftNodeServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.qbftNode/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QbftNodeServer).GetNodeInfo(ctx, req.(*types.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

var _QbftNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.qbftNode",
	HandlerType: (*QbftNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsSync",
			Handler:    _QbftNode_IsSync_Handler,
		},
		{
			MethodName: "GetNodeInfo",
			Handler:    _QbftNode_GetNodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qbftNode.proto",
}
