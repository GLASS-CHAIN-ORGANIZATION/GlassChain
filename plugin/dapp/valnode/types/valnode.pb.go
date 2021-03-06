// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valnode.proto

package types

import (
	context "context"
	fmt "fmt"
	math "math"

	types "github.com/33cn/chain33/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ValNode struct {
	PubKey               []byte   `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Power                int64    `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValNode) Reset()         { *m = ValNode{} }
func (m *ValNode) String() string { return proto.CompactTextString(m) }
func (*ValNode) ProtoMessage()    {}
func (*ValNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{0}
}

func (m *ValNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValNode.Unmarshal(m, b)
}
func (m *ValNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValNode.Marshal(b, m, deterministic)
}
func (m *ValNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValNode.Merge(m, src)
}
func (m *ValNode) XXX_Size() int {
	return xxx_messageInfo_ValNode.Size(m)
}
func (m *ValNode) XXX_DiscardUnknown() {
	xxx_messageInfo_ValNode.DiscardUnknown(m)
}

var xxx_messageInfo_ValNode proto.InternalMessageInfo

func (m *ValNode) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *ValNode) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

type ValNodes struct {
	Nodes                []*ValNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ValNodes) Reset()         { *m = ValNodes{} }
func (m *ValNodes) String() string { return proto.CompactTextString(m) }
func (*ValNodes) ProtoMessage()    {}
func (*ValNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{1}
}

func (m *ValNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValNodes.Unmarshal(m, b)
}
func (m *ValNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValNodes.Marshal(b, m, deterministic)
}
func (m *ValNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValNodes.Merge(m, src)
}
func (m *ValNodes) XXX_Size() int {
	return xxx_messageInfo_ValNodes.Size(m)
}
func (m *ValNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_ValNodes.DiscardUnknown(m)
}

var xxx_messageInfo_ValNodes proto.InternalMessageInfo

func (m *ValNodes) GetNodes() []*ValNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ValNodeAction struct {
	// Types that are valid to be assigned to Value:
	//	*ValNodeAction_Node
	//	*ValNodeAction_BlockInfo
	Value                isValNodeAction_Value `protobuf_oneof:"value"`
	Ty                   int32                 `protobuf:"varint,3,opt,name=Ty,proto3" json:"Ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ValNodeAction) Reset()         { *m = ValNodeAction{} }
func (m *ValNodeAction) String() string { return proto.CompactTextString(m) }
func (*ValNodeAction) ProtoMessage()    {}
func (*ValNodeAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{2}
}

func (m *ValNodeAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValNodeAction.Unmarshal(m, b)
}
func (m *ValNodeAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValNodeAction.Marshal(b, m, deterministic)
}
func (m *ValNodeAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValNodeAction.Merge(m, src)
}
func (m *ValNodeAction) XXX_Size() int {
	return xxx_messageInfo_ValNodeAction.Size(m)
}
func (m *ValNodeAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ValNodeAction.DiscardUnknown(m)
}

var xxx_messageInfo_ValNodeAction proto.InternalMessageInfo

type isValNodeAction_Value interface {
	isValNodeAction_Value()
}

type ValNodeAction_Node struct {
	Node *ValNode `protobuf:"bytes,1,opt,name=node,proto3,oneof"`
}

type ValNodeAction_BlockInfo struct {
	BlockInfo *TendermintBlockInfo `protobuf:"bytes,2,opt,name=blockInfo,proto3,oneof"`
}

func (*ValNodeAction_Node) isValNodeAction_Value() {}

func (*ValNodeAction_BlockInfo) isValNodeAction_Value() {}

func (m *ValNodeAction) GetValue() isValNodeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValNodeAction) GetNode() *ValNode {
	if x, ok := m.GetValue().(*ValNodeAction_Node); ok {
		return x.Node
	}
	return nil
}

func (m *ValNodeAction) GetBlockInfo() *TendermintBlockInfo {
	if x, ok := m.GetValue().(*ValNodeAction_BlockInfo); ok {
		return x.BlockInfo
	}
	return nil
}

func (m *ValNodeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValNodeAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValNodeAction_Node)(nil),
		(*ValNodeAction_BlockInfo)(nil),
	}
}

type ReqValNodes struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqValNodes) Reset()         { *m = ReqValNodes{} }
func (m *ReqValNodes) String() string { return proto.CompactTextString(m) }
func (*ReqValNodes) ProtoMessage()    {}
func (*ReqValNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{3}
}

func (m *ReqValNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqValNodes.Unmarshal(m, b)
}
func (m *ReqValNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqValNodes.Marshal(b, m, deterministic)
}
func (m *ReqValNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqValNodes.Merge(m, src)
}
func (m *ReqValNodes) XXX_Size() int {
	return xxx_messageInfo_ReqValNodes.Size(m)
}
func (m *ReqValNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqValNodes.DiscardUnknown(m)
}

var xxx_messageInfo_ReqValNodes proto.InternalMessageInfo

func (m *ReqValNodes) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ReqBlockInfo struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBlockInfo) Reset()         { *m = ReqBlockInfo{} }
func (m *ReqBlockInfo) String() string { return proto.CompactTextString(m) }
func (*ReqBlockInfo) ProtoMessage()    {}
func (*ReqBlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{4}
}

func (m *ReqBlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBlockInfo.Unmarshal(m, b)
}
func (m *ReqBlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBlockInfo.Marshal(b, m, deterministic)
}
func (m *ReqBlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBlockInfo.Merge(m, src)
}
func (m *ReqBlockInfo) XXX_Size() int {
	return xxx_messageInfo_ReqBlockInfo.Size(m)
}
func (m *ReqBlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBlockInfo proto.InternalMessageInfo

func (m *ReqBlockInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ValNodeInfo struct {
	NodeIP               string   `protobuf:"bytes,1,opt,name=nodeIP,proto3" json:"nodeIP,omitempty"`
	NodeID               string   `protobuf:"bytes,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PubKey               string   `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	VotingPower          int64    `protobuf:"varint,5,opt,name=votingPower,proto3" json:"votingPower,omitempty"`
	Accum                int64    `protobuf:"varint,6,opt,name=accum,proto3" json:"accum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValNodeInfo) Reset()         { *m = ValNodeInfo{} }
func (m *ValNodeInfo) String() string { return proto.CompactTextString(m) }
func (*ValNodeInfo) ProtoMessage()    {}
func (*ValNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{5}
}

func (m *ValNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValNodeInfo.Unmarshal(m, b)
}
func (m *ValNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValNodeInfo.Marshal(b, m, deterministic)
}
func (m *ValNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValNodeInfo.Merge(m, src)
}
func (m *ValNodeInfo) XXX_Size() int {
	return xxx_messageInfo_ValNodeInfo.Size(m)
}
func (m *ValNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValNodeInfo proto.InternalMessageInfo

func (m *ValNodeInfo) GetNodeIP() string {
	if m != nil {
		return m.NodeIP
	}
	return ""
}

func (m *ValNodeInfo) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *ValNodeInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValNodeInfo) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *ValNodeInfo) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *ValNodeInfo) GetAccum() int64 {
	if m != nil {
		return m.Accum
	}
	return 0
}

type ValNodeInfoSet struct {
	Nodes                []*ValNodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ValNodeInfoSet) Reset()         { *m = ValNodeInfoSet{} }
func (m *ValNodeInfoSet) String() string { return proto.CompactTextString(m) }
func (*ValNodeInfoSet) ProtoMessage()    {}
func (*ValNodeInfoSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{6}
}

func (m *ValNodeInfoSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValNodeInfoSet.Unmarshal(m, b)
}
func (m *ValNodeInfoSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValNodeInfoSet.Marshal(b, m, deterministic)
}
func (m *ValNodeInfoSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValNodeInfoSet.Merge(m, src)
}
func (m *ValNodeInfoSet) XXX_Size() int {
	return xxx_messageInfo_ValNodeInfoSet.Size(m)
}
func (m *ValNodeInfoSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValNodeInfoSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValNodeInfoSet proto.InternalMessageInfo

func (m *ValNodeInfoSet) GetNodes() []*ValNodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type PerfStat struct {
	TotalTx              int64    `protobuf:"varint,1,opt,name=totalTx,proto3" json:"totalTx,omitempty"`
	TotalBlock           int64    `protobuf:"varint,2,opt,name=totalBlock,proto3" json:"totalBlock,omitempty"`
	TxPerBlock           int64    `protobuf:"varint,3,opt,name=txPerBlock,proto3" json:"txPerBlock,omitempty"`
	TotalSecond          int64    `protobuf:"varint,4,opt,name=totalSecond,proto3" json:"totalSecond,omitempty"`
	TxPerSecond          int64    `protobuf:"varint,5,opt,name=txPerSecond,proto3" json:"txPerSecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerfStat) Reset()         { *m = PerfStat{} }
func (m *PerfStat) String() string { return proto.CompactTextString(m) }
func (*PerfStat) ProtoMessage()    {}
func (*PerfStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{7}
}

func (m *PerfStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerfStat.Unmarshal(m, b)
}
func (m *PerfStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerfStat.Marshal(b, m, deterministic)
}
func (m *PerfStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerfStat.Merge(m, src)
}
func (m *PerfStat) XXX_Size() int {
	return xxx_messageInfo_PerfStat.Size(m)
}
func (m *PerfStat) XXX_DiscardUnknown() {
	xxx_messageInfo_PerfStat.DiscardUnknown(m)
}

var xxx_messageInfo_PerfStat proto.InternalMessageInfo

func (m *PerfStat) GetTotalTx() int64 {
	if m != nil {
		return m.TotalTx
	}
	return 0
}

func (m *PerfStat) GetTotalBlock() int64 {
	if m != nil {
		return m.TotalBlock
	}
	return 0
}

func (m *PerfStat) GetTxPerBlock() int64 {
	if m != nil {
		return m.TxPerBlock
	}
	return 0
}

func (m *PerfStat) GetTotalSecond() int64 {
	if m != nil {
		return m.TotalSecond
	}
	return 0
}

func (m *PerfStat) GetTxPerSecond() int64 {
	if m != nil {
		return m.TxPerSecond
	}
	return 0
}

type ReqPerfStat struct {
	Start                int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqPerfStat) Reset()         { *m = ReqPerfStat{} }
func (m *ReqPerfStat) String() string { return proto.CompactTextString(m) }
func (*ReqPerfStat) ProtoMessage()    {}
func (*ReqPerfStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e9a3523ca7e0ea, []int{8}
}

func (m *ReqPerfStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqPerfStat.Unmarshal(m, b)
}
func (m *ReqPerfStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqPerfStat.Marshal(b, m, deterministic)
}
func (m *ReqPerfStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqPerfStat.Merge(m, src)
}
func (m *ReqPerfStat) XXX_Size() int {
	return xxx_messageInfo_ReqPerfStat.Size(m)
}
func (m *ReqPerfStat) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqPerfStat.DiscardUnknown(m)
}

var xxx_messageInfo_ReqPerfStat proto.InternalMessageInfo

func (m *ReqPerfStat) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ReqPerfStat) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*ValNode)(nil), "types.ValNode")
	proto.RegisterType((*ValNodes)(nil), "types.ValNodes")
	proto.RegisterType((*ValNodeAction)(nil), "types.ValNodeAction")
	proto.RegisterType((*ReqValNodes)(nil), "types.ReqValNodes")
	proto.RegisterType((*ReqBlockInfo)(nil), "types.ReqBlockInfo")
	proto.RegisterType((*ValNodeInfo)(nil), "types.ValNodeInfo")
	proto.RegisterType((*ValNodeInfoSet)(nil), "types.ValNodeInfoSet")
	proto.RegisterType((*PerfStat)(nil), "types.PerfStat")
	proto.RegisterType((*ReqPerfStat)(nil), "types.ReqPerfStat")
}

func init() {
	proto.RegisterFile("valnode.proto", fileDescriptor_38e9a3523ca7e0ea)
}

var fileDescriptor_38e9a3523ca7e0ea = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xd1, 0x6a, 0xdb, 0x40,
	0x10, 0xb4, 0xac, 0xc8, 0x8e, 0x57, 0xb6, 0x31, 0x47, 0x1a, 0x84, 0x1f, 0x8a, 0x11, 0x69, 0x11,
	0x14, 0x4c, 0x71, 0x09, 0x85, 0xbc, 0x35, 0x04, 0x6a, 0x53, 0x08, 0xe6, 0x64, 0xfa, 0x2e, 0x4b,
	0x9b, 0x58, 0x54, 0xba, 0xb3, 0xa5, 0xb3, 0x1b, 0xfd, 0x42, 0x7f, 0xa4, 0xfd, 0xcc, 0xa2, 0xd5,
	0x49, 0x55, 0x1a, 0xf2, 0xa6, 0x99, 0x9d, 0xbd, 0x9b, 0xdd, 0xd1, 0xc1, 0xe8, 0x14, 0x24, 0x42,
	0x46, 0x38, 0xdf, 0x67, 0x52, 0x49, 0x66, 0xa9, 0x62, 0x8f, 0xf9, 0x74, 0x18, 0xca, 0x34, 0x95,
	0xa2, 0x22, 0xa7, 0x13, 0x85, 0x22, 0xc2, 0x2c, 0x8d, 0x85, 0xaa, 0x18, 0xf7, 0x33, 0xf4, 0xbf,
	0x07, 0xc9, 0xbd, 0x8c, 0x90, 0x5d, 0x42, 0x6f, 0x7f, 0xdc, 0x7e, 0xc3, 0xc2, 0x31, 0x66, 0x86,
	0x37, 0xe4, 0x1a, 0xb1, 0x0b, 0xb0, 0xf6, 0xf2, 0x27, 0x66, 0x4e, 0x77, 0x66, 0x78, 0x26, 0xaf,
	0x80, 0xfb, 0x11, 0xce, 0x75, 0x63, 0xce, 0xae, 0xc0, 0x2a, 0x6f, 0xce, 0x1d, 0x63, 0x66, 0x7a,
	0xf6, 0x62, 0x3c, 0xa7, 0xbb, 0xe7, 0xba, 0xce, 0xab, 0xa2, 0xfb, 0xcb, 0x80, 0x91, 0xa6, 0xbe,
	0x84, 0x2a, 0x96, 0x82, 0x5d, 0xc1, 0x59, 0x59, 0xa2, 0xfb, 0x5e, 0xb4, 0x2d, 0x3b, 0x9c, 0xaa,
	0xec, 0x06, 0x06, 0xdb, 0x44, 0x86, 0x3f, 0x56, 0xe2, 0x41, 0x92, 0x07, 0x7b, 0x31, 0xd5, 0xd2,
	0x4d, 0x33, 0xce, 0x6d, 0xad, 0x58, 0x76, 0xf8, 0x3f, 0x39, 0x1b, 0x43, 0x77, 0x53, 0x38, 0xe6,
	0xcc, 0xf0, 0x2c, 0xde, 0xdd, 0x14, 0xb7, 0x7d, 0xb0, 0x4e, 0x41, 0x72, 0x44, 0xf7, 0x1d, 0xd8,
	0x1c, 0x0f, 0xcd, 0x04, 0x97, 0xd0, 0xdb, 0x61, 0xfc, 0xb8, 0x53, 0xe4, 0xc5, 0xe4, 0x1a, 0xb9,
	0xef, 0x61, 0xc8, 0xf1, 0xd0, 0x1c, 0xfe, 0xaa, 0xee, 0xb7, 0x01, 0xb6, 0x3e, 0xac, 0xd6, 0x95,
	0xde, 0x57, 0x6b, 0xd2, 0x0d, 0xb8, 0x46, 0x0d, 0x7f, 0x47, 0x83, 0xd4, 0xfc, 0x1d, 0x73, 0xa0,
	0x1f, 0x44, 0x51, 0x86, 0x79, 0x4e, 0x66, 0x07, 0xbc, 0x86, 0xad, 0x54, 0xce, 0xaa, 0x0e, 0x9d,
	0xca, 0x0c, 0xec, 0x93, 0x54, 0xb1, 0x78, 0x5c, 0x53, 0x36, 0x16, 0xd9, 0x69, 0x53, 0x65, 0x6e,
	0x41, 0x18, 0x1e, 0x53, 0xa7, 0x57, 0xe5, 0x46, 0xc0, 0xbd, 0x81, 0x71, 0xcb, 0xa8, 0x8f, 0x8a,
	0x79, 0xcf, 0xd3, 0x63, 0xcf, 0x63, 0x28, 0x55, 0x75, 0x82, 0x7f, 0x0c, 0x38, 0x5f, 0x63, 0xf6,
	0xe0, 0xab, 0x40, 0x95, 0x96, 0x95, 0x54, 0x41, 0xb2, 0x79, 0xd2, 0xbb, 0xa8, 0x21, 0x7b, 0x0b,
	0x40, 0x9f, 0xb4, 0x36, 0xfd, 0xd7, 0xb4, 0x18, 0xaa, 0x3f, 0xad, 0x31, 0xab, 0xea, 0xa6, 0xae,
	0x37, 0x4c, 0x39, 0x1a, 0xa9, 0x7d, 0x0c, 0xa5, 0x88, 0x68, 0x6e, 0x93, 0xb7, 0x29, 0x52, 0x94,
	0x7a, 0xad, 0xd0, 0xc3, 0xb7, 0x28, 0xf7, 0x9a, 0xf2, 0x6d, 0xcc, 0x5e, 0x80, 0x95, 0xab, 0x20,
	0xab, 0x63, 0xab, 0x00, 0x9b, 0x80, 0x89, 0x22, 0xd2, 0x0e, 0xcb, 0xcf, 0x45, 0x0a, 0x7d, 0xfd,
	0x8c, 0xd8, 0x07, 0xe8, 0xad, 0x72, 0xbf, 0x10, 0x21, 0x1b, 0xe9, 0x8d, 0x70, 0x3c, 0xdc, 0xc7,
	0xc9, 0x74, 0xa2, 0xe1, 0x2a, 0x5f, 0x62, 0x90, 0xa8, 0x5d, 0xe1, 0x76, 0xd8, 0x35, 0xd8, 0x5f,
	0x51, 0x35, 0xf1, 0xff, 0xd7, 0xf1, 0xe6, 0xe5, 0x4a, 0x7d, 0x54, 0x6e, 0x67, 0xdb, 0xa3, 0x47,
	0xf8, 0xe9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x95, 0xc7, 0x52, 0xbc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ValnodeClient is the client API for Valnode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValnodeClient interface {
	IsSync(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*IsHealthy, error)
	GetNodeInfo(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*ValNodeInfoSet, error)
}

type valnodeClient struct {
	cc grpc.ClientConnInterface
}

func NewValnodeClient(cc grpc.ClientConnInterface) ValnodeClient {
	return &valnodeClient{cc}
}

func (c *valnodeClient) IsSync(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*IsHealthy, error) {
	out := new(IsHealthy)
	err := c.cc.Invoke(ctx, "/types.valnode/IsSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valnodeClient) GetNodeInfo(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*ValNodeInfoSet, error) {
	out := new(ValNodeInfoSet)
	err := c.cc.Invoke(ctx, "/types.valnode/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValnodeServer is the server API for Valnode service.
type ValnodeServer interface {
	IsSync(context.Context, *types.ReqNil) (*IsHealthy, error)
	GetNodeInfo(context.Context, *types.ReqNil) (*ValNodeInfoSet, error)
}

// UnimplementedValnodeServer can be embedded to have forward compatible implementations.
type UnimplementedValnodeServer struct {
}

func (*UnimplementedValnodeServer) IsSync(ctx context.Context, req *types.ReqNil) (*IsHealthy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSync not implemented")
}
func (*UnimplementedValnodeServer) GetNodeInfo(ctx context.Context, req *types.ReqNil) (*ValNodeInfoSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}

func RegisterValnodeServer(s *grpc.Server, srv ValnodeServer) {
	s.RegisterService(&_Valnode_serviceDesc, srv)
}

func _Valnode_IsSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValnodeServer).IsSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.valnode/IsSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValnodeServer).IsSync(ctx, req.(*types.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Valnode_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValnodeServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.valnode/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValnodeServer).GetNodeInfo(ctx, req.(*types.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

var _Valnode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.valnode",
	HandlerType: (*ValnodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsSync",
			Handler:    _Valnode_IsSync_Handler,
		},
		{
			MethodName: "GetNodeInfo",
			Handler:    _Valnode_GetNodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "valnode.proto",
}
