// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.1
// source: project.proto

package types

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

type AutonomyProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropProject *ProposalProject `protobuf:"bytes,1,opt,name=propProject,proto3" json:"propProject,omitempty"`

	CurRule *RuleConfig `protobuf:"bytes,2,opt,name=curRule,proto3" json:"curRule,omitempty"`

	Boards []string `protobuf:"bytes,3,rep,name=boards,proto3" json:"boards,omitempty"`

	BoardVoteRes *VoteResult `protobuf:"bytes,4,opt,name=boardVoteRes,proto3" json:"boardVoteRes,omitempty"`

	PubVote *PublicVote `protobuf:"bytes,5,opt,name=pubVote,proto3" json:"pubVote,omitempty"`

	Status     int32  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Address    string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Height     int64  `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`
	Index      int32  `protobuf:"varint,9,opt,name=index,proto3" json:"index,omitempty"`
	ProposalID string `protobuf:"bytes,10,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (x *AutonomyProposalProject) Reset() {
	*x = AutonomyProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutonomyProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutonomyProposalProject) ProtoMessage() {}

func (x *AutonomyProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutonomyProposalProject.ProtoReflect.Descriptor instead.
func (*AutonomyProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *AutonomyProposalProject) GetPropProject() *ProposalProject {
	if x != nil {
		return x.PropProject
	}
	return nil
}

func (x *AutonomyProposalProject) GetCurRule() *RuleConfig {
	if x != nil {
		return x.CurRule
	}
	return nil
}

func (x *AutonomyProposalProject) GetBoards() []string {
	if x != nil {
		return x.Boards
	}
	return nil
}

func (x *AutonomyProposalProject) GetBoardVoteRes() *VoteResult {
	if x != nil {
		return x.BoardVoteRes
	}
	return nil
}

func (x *AutonomyProposalProject) GetPubVote() *PublicVote {
	if x != nil {
		return x.PubVote
	}
	return nil
}

func (x *AutonomyProposalProject) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AutonomyProposalProject) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AutonomyProposalProject) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *AutonomyProposalProject) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *AutonomyProposalProject) GetProposalID() string {
	if x != nil {
		return x.ProposalID
	}
	return ""
}

type ProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields


	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`

	FirstStage   string `protobuf:"bytes,4,opt,name=firstStage,proto3" json:"firstStage,omitempty"`      
	LastStage    string `protobuf:"bytes,5,opt,name=lastStage,proto3" json:"lastStage,omitempty"`        
	Production   string `protobuf:"bytes,6,opt,name=production,proto3" json:"production,omitempty"`      
	Description  string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`    
	Contractor   string `protobuf:"bytes,8,opt,name=contractor,proto3" json:"contractor,omitempty"`      
	Amount       int64  `protobuf:"varint,9,opt,name=amount,proto3" json:"amount,omitempty"`             
	AmountDetail string `protobuf:"bytes,10,opt,name=amountDetail,proto3" json:"amountDetail,omitempty"` 

	ToAddr string `protobuf:"bytes,11,opt,name=toAddr,proto3" json:"toAddr,omitempty"` 

	StartBlockHeight    int64 `protobuf:"varint,12,opt,name=startBlockHeight,proto3" json:"startBlockHeight,omitempty"`       
	EndBlockHeight      int64 `protobuf:"varint,13,opt,name=endBlockHeight,proto3" json:"endBlockHeight,omitempty"`           
	RealEndBlockHeight  int64 `protobuf:"varint,14,opt,name=realEndBlockHeight,proto3" json:"realEndBlockHeight,omitempty"`   
	ProjectNeedBlockNum int32 `protobuf:"varint,15,opt,name=projectNeedBlockNum,proto3" json:"projectNeedBlockNum,omitempty"` 
}

func (x *ProposalProject) Reset() {
	*x = ProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalProject) ProtoMessage() {}

func (x *ProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalProject.ProtoReflect.Descriptor instead.
func (*ProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

func (x *ProposalProject) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *ProposalProject) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *ProposalProject) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *ProposalProject) GetFirstStage() string {
	if x != nil {
		return x.FirstStage
	}
	return ""
}

func (x *ProposalProject) GetLastStage() string {
	if x != nil {
		return x.LastStage
	}
	return ""
}

func (x *ProposalProject) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *ProposalProject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProposalProject) GetContractor() string {
	if x != nil {
		return x.Contractor
	}
	return ""
}

func (x *ProposalProject) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ProposalProject) GetAmountDetail() string {
	if x != nil {
		return x.AmountDetail
	}
	return ""
}

func (x *ProposalProject) GetToAddr() string {
	if x != nil {
		return x.ToAddr
	}
	return ""
}

func (x *ProposalProject) GetStartBlockHeight() int64 {
	if x != nil {
		return x.StartBlockHeight
	}
	return 0
}

func (x *ProposalProject) GetEndBlockHeight() int64 {
	if x != nil {
		return x.EndBlockHeight
	}
	return 0
}

func (x *ProposalProject) GetRealEndBlockHeight() int64 {
	if x != nil {
		return x.RealEndBlockHeight
	}
	return 0
}

func (x *ProposalProject) GetProjectNeedBlockNum() int32 {
	if x != nil {
		return x.ProjectNeedBlockNum
	}
	return 0
}

type RevokeProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalID string `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (x *RevokeProposalProject) Reset() {
	*x = RevokeProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeProposalProject) ProtoMessage() {}

func (x *RevokeProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeProposalProject.ProtoReflect.Descriptor instead.
func (*RevokeProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

func (x *RevokeProposalProject) GetProposalID() string {
	if x != nil {
		return x.ProposalID
	}
	return ""
}

type VoteProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalID string `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	Approve    bool   `protobuf:"varint,2,opt,name=approve,proto3" json:"approve,omitempty"`
}

func (x *VoteProposalProject) Reset() {
	*x = VoteProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteProposalProject) ProtoMessage() {}

func (x *VoteProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteProposalProject.ProtoReflect.Descriptor instead.
func (*VoteProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{3}
}

func (x *VoteProposalProject) GetProposalID() string {
	if x != nil {
		return x.ProposalID
	}
	return ""
}

func (x *VoteProposalProject) GetApprove() bool {
	if x != nil {
		return x.Approve
	}
	return false
}

type PubVoteProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalID string   `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
	Oppose     bool     `protobuf:"varint,2,opt,name=oppose,proto3" json:"oppose,omitempty"`
	OriginAddr []string `protobuf:"bytes,3,rep,name=originAddr,proto3" json:"originAddr,omitempty"`
}

func (x *PubVoteProposalProject) Reset() {
	*x = PubVoteProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubVoteProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubVoteProposalProject) ProtoMessage() {}

func (x *PubVoteProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubVoteProposalProject.ProtoReflect.Descriptor instead.
func (*PubVoteProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{4}
}

func (x *PubVoteProposalProject) GetProposalID() string {
	if x != nil {
		return x.ProposalID
	}
	return ""
}

func (x *PubVoteProposalProject) GetOppose() bool {
	if x != nil {
		return x.Oppose
	}
	return false
}

func (x *PubVoteProposalProject) GetOriginAddr() []string {
	if x != nil {
		return x.OriginAddr
	}
	return nil
}

type TerminateProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalID string `protobuf:"bytes,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (x *TerminateProposalProject) Reset() {
	*x = TerminateProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateProposalProject) ProtoMessage() {}

func (x *TerminateProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateProposalProject.ProtoReflect.Descriptor instead.
func (*TerminateProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{5}
}

func (x *TerminateProposalProject) GetProposalID() string {
	if x != nil {
		return x.ProposalID
	}
	return ""
}

// receipt
type ReceiptProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prev    *AutonomyProposalProject `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current *AutonomyProposalProject `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *ReceiptProposalProject) Reset() {
	*x = ReceiptProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptProposalProject) ProtoMessage() {}

func (x *ReceiptProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptProposalProject.ProtoReflect.Descriptor instead.
func (*ReceiptProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{6}
}

func (x *ReceiptProposalProject) GetPrev() *AutonomyProposalProject {
	if x != nil {
		return x.Prev
	}
	return nil
}

func (x *ReceiptProposalProject) GetCurrent() *AutonomyProposalProject {
	if x != nil {
		return x.Current
	}
	return nil
}

type LocalProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropPrj  *AutonomyProposalProject `protobuf:"bytes,1,opt,name=propPrj,proto3" json:"propPrj,omitempty"`
	Comments []string                 `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *LocalProposalProject) Reset() {
	*x = LocalProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalProposalProject) ProtoMessage() {}

func (x *LocalProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalProposalProject.ProtoReflect.Descriptor instead.
func (*LocalProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{7}
}

func (x *LocalProposalProject) GetPropPrj() *AutonomyProposalProject {
	if x != nil {
		return x.PropPrj
	}
	return nil
}

func (x *LocalProposalProject) GetComments() []string {
	if x != nil {
		return x.Comments
	}
	return nil
}

// query
type ReqQueryProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Count     int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,4,opt,name=direction,proto3" json:"direction,omitempty"`
	Height    int64  `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Index     int32  `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *ReqQueryProposalProject) Reset() {
	*x = ReqQueryProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQueryProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQueryProposalProject) ProtoMessage() {}

func (x *ReqQueryProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQueryProposalProject.ProtoReflect.Descriptor instead.
func (*ReqQueryProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{8}
}

func (x *ReqQueryProposalProject) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ReqQueryProposalProject) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ReqQueryProposalProject) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ReqQueryProposalProject) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *ReqQueryProposalProject) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ReqQueryProposalProject) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type ReplyQueryProposalProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropProjects []*AutonomyProposalProject `protobuf:"bytes,1,rep,name=propProjects,proto3" json:"propProjects,omitempty"`
}

func (x *ReplyQueryProposalProject) Reset() {
	*x = ReplyQueryProposalProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyQueryProposalProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyQueryProposalProject) ProtoMessage() {}

func (x *ReplyQueryProposalProject) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyQueryProposalProject.ProtoReflect.Descriptor instead.
func (*ReplyQueryProposalProject) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{9}
}

func (x *ReplyQueryProposalProject) GetPropProjects() []*AutonomyProposalProject {
	if x != nil {
		return x.PropProjects
	}
	return nil
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0d, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x38, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x07, 0x63, 0x75, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x35, 0x0a, 0x0c, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x56, 0x6f,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x07, 0x70, 0x75, 0x62,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x49, 0x44, 0x22, 0xf7, 0x03, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x64, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2e, 0x0a, 0x12,
	0x72, 0x65, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x72, 0x65, 0x61, 0x6c, 0x45, 0x6e,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x13,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x22, 0x37,
	0x0a, 0x15, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x13, 0x56, 0x6f, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x22, 0x70, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x56,
	0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x6f, 0x70, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x22, 0x3a, 0x0a, 0x18, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x49, 0x44, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22,
	0x6c, 0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x50,
	0x72, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x50, 0x72,
	0x6a, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa7, 0x01,
	0x0a, 0x17, 0x52, 0x65, 0x71, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x5f, 0x0a, 0x19, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData = file_project_proto_rawDesc
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_proto_rawDescData)
	})
	return file_project_proto_rawDescData
}

var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_project_proto_goTypes = []interface{}{
	(*AutonomyProposalProject)(nil),   // 0: types.AutonomyProposalProject
	(*ProposalProject)(nil),           // 1: types.ProposalProject
	(*RevokeProposalProject)(nil),     // 2: types.RevokeProposalProject
	(*VoteProposalProject)(nil),       // 3: types.VoteProposalProject
	(*PubVoteProposalProject)(nil),    // 4: types.PubVoteProposalProject
	(*TerminateProposalProject)(nil),  // 5: types.TerminateProposalProject
	(*ReceiptProposalProject)(nil),    // 6: types.ReceiptProposalProject
	(*LocalProposalProject)(nil),      // 7: types.LocalProposalProject
	(*ReqQueryProposalProject)(nil),   // 8: types.ReqQueryProposalProject
	(*ReplyQueryProposalProject)(nil), // 9: types.ReplyQueryProposalProject
	(*RuleConfig)(nil),                // 10: types.RuleConfig
	(*VoteResult)(nil),                // 11: types.VoteResult
	(*PublicVote)(nil),                // 12: types.PublicVote
}
var file_project_proto_depIdxs = []int32{
	1,  // 0: types.AutonomyProposalProject.propProject:type_name -> types.ProposalProject
	10, // 1: types.AutonomyProposalProject.curRule:type_name -> types.RuleConfig
	11, // 2: types.AutonomyProposalProject.boardVoteRes:type_name -> types.VoteResult
	12, // 3: types.AutonomyProposalProject.pubVote:type_name -> types.PublicVote
	0,  // 4: types.ReceiptProposalProject.prev:type_name -> types.AutonomyProposalProject
	0,  // 5: types.ReceiptProposalProject.current:type_name -> types.AutonomyProposalProject
	0,  // 6: types.LocalProposalProject.propPrj:type_name -> types.AutonomyProposalProject
	0,  // 7: types.ReplyQueryProposalProject.propProjects:type_name -> types.AutonomyProposalProject
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	file_lcommon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutonomyProposalProject); i {
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
		file_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalProject); i {
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
		file_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeProposalProject); i {
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
		file_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteProposalProject); i {
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
		file_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubVoteProposalProject); i {
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
		file_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateProposalProject); i {
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
		file_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptProposalProject); i {
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
		file_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalProposalProject); i {
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
		file_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQueryProposalProject); i {
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
		file_project_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyQueryProposalProject); i {
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
			RawDescriptor: file_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_rawDesc = nil
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
