// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wasm.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type WasmAction struct {
	// Types that are valid to be assigned to Value:
	//	*WasmAction_Create
	//	*WasmAction_Call
	Value                isWasmAction_Value `protobuf_oneof:"value"`
	Ty                   int32              `protobuf:"varint,3,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WasmAction) Reset()         { *m = WasmAction{} }
func (m *WasmAction) String() string { return proto.CompactTextString(m) }
func (*WasmAction) ProtoMessage()    {}
func (*WasmAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{0}
}

func (m *WasmAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmAction.Unmarshal(m, b)
}
func (m *WasmAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmAction.Marshal(b, m, deterministic)
}
func (m *WasmAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmAction.Merge(m, src)
}
func (m *WasmAction) XXX_Size() int {
	return xxx_messageInfo_WasmAction.Size(m)
}
func (m *WasmAction) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmAction.DiscardUnknown(m)
}

var xxx_messageInfo_WasmAction proto.InternalMessageInfo

type isWasmAction_Value interface {
	isWasmAction_Value()
}

type WasmAction_Create struct {
	Create *WasmCreate `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type WasmAction_Call struct {
	Call *WasmCall `protobuf:"bytes,2,opt,name=call,proto3,oneof"`
}

func (*WasmAction_Create) isWasmAction_Value() {}

func (*WasmAction_Call) isWasmAction_Value() {}

func (m *WasmAction) GetValue() isWasmAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *WasmAction) GetCreate() *WasmCreate {
	if x, ok := m.GetValue().(*WasmAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *WasmAction) GetCall() *WasmCall {
	if x, ok := m.GetValue().(*WasmAction_Call); ok {
		return x.Call
	}
	return nil
}

func (m *WasmAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WasmAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WasmAction_Create)(nil),
		(*WasmAction_Call)(nil),
	}
}

type WasmCreate struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 []byte   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WasmCreate) Reset()         { *m = WasmCreate{} }
func (m *WasmCreate) String() string { return proto.CompactTextString(m) }
func (*WasmCreate) ProtoMessage()    {}
func (*WasmCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{1}
}

func (m *WasmCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmCreate.Unmarshal(m, b)
}
func (m *WasmCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmCreate.Marshal(b, m, deterministic)
}
func (m *WasmCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmCreate.Merge(m, src)
}
func (m *WasmCreate) XXX_Size() int {
	return xxx_messageInfo_WasmCreate.Size(m)
}
func (m *WasmCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmCreate.DiscardUnknown(m)
}

var xxx_messageInfo_WasmCreate proto.InternalMessageInfo

func (m *WasmCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WasmCreate) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type WasmCall struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Parameters           []int64  `protobuf:"varint,3,rep,packed,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WasmCall) Reset()         { *m = WasmCall{} }
func (m *WasmCall) String() string { return proto.CompactTextString(m) }
func (*WasmCall) ProtoMessage()    {}
func (*WasmCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{2}
}

func (m *WasmCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmCall.Unmarshal(m, b)
}
func (m *WasmCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmCall.Marshal(b, m, deterministic)
}
func (m *WasmCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmCall.Merge(m, src)
}
func (m *WasmCall) XXX_Size() int {
	return xxx_messageInfo_WasmCall.Size(m)
}
func (m *WasmCall) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmCall.DiscardUnknown(m)
}

var xxx_messageInfo_WasmCall proto.InternalMessageInfo

func (m *WasmCall) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *WasmCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *WasmCall) GetParameters() []int64 {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type QueryCheckContract struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryCheckContract) Reset()         { *m = QueryCheckContract{} }
func (m *QueryCheckContract) String() string { return proto.CompactTextString(m) }
func (*QueryCheckContract) ProtoMessage()    {}
func (*QueryCheckContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{3}
}

func (m *QueryCheckContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCheckContract.Unmarshal(m, b)
}
func (m *QueryCheckContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCheckContract.Marshal(b, m, deterministic)
}
func (m *QueryCheckContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCheckContract.Merge(m, src)
}
func (m *QueryCheckContract) XXX_Size() int {
	return xxx_messageInfo_QueryCheckContract.Size(m)
}
func (m *QueryCheckContract) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCheckContract.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCheckContract proto.InternalMessageInfo

func (m *QueryCheckContract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CustomLog struct {
	Info                 []string `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomLog) Reset()         { *m = CustomLog{} }
func (m *CustomLog) String() string { return proto.CompactTextString(m) }
func (*CustomLog) ProtoMessage()    {}
func (*CustomLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{4}
}

func (m *CustomLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomLog.Unmarshal(m, b)
}
func (m *CustomLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomLog.Marshal(b, m, deterministic)
}
func (m *CustomLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomLog.Merge(m, src)
}
func (m *CustomLog) XXX_Size() int {
	return xxx_messageInfo_CustomLog.Size(m)
}
func (m *CustomLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomLog.DiscardUnknown(m)
}

var xxx_messageInfo_CustomLog proto.InternalMessageInfo

func (m *CustomLog) GetInfo() []string {
	if m != nil {
		return m.Info
	}
	return nil
}

type CreateContractLog struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateContractLog) Reset()         { *m = CreateContractLog{} }
func (m *CreateContractLog) String() string { return proto.CompactTextString(m) }
func (*CreateContractLog) ProtoMessage()    {}
func (*CreateContractLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{5}
}

func (m *CreateContractLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateContractLog.Unmarshal(m, b)
}
func (m *CreateContractLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateContractLog.Marshal(b, m, deterministic)
}
func (m *CreateContractLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContractLog.Merge(m, src)
}
func (m *CreateContractLog) XXX_Size() int {
	return xxx_messageInfo_CreateContractLog.Size(m)
}
func (m *CreateContractLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContractLog.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContractLog proto.InternalMessageInfo

func (m *CreateContractLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateContractLog) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CallContractLog struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Result               int32    `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallContractLog) Reset()         { *m = CallContractLog{} }
func (m *CallContractLog) String() string { return proto.CompactTextString(m) }
func (*CallContractLog) ProtoMessage()    {}
func (*CallContractLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{6}
}

func (m *CallContractLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallContractLog.Unmarshal(m, b)
}
func (m *CallContractLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallContractLog.Marshal(b, m, deterministic)
}
func (m *CallContractLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractLog.Merge(m, src)
}
func (m *CallContractLog) XXX_Size() int {
	return xxx_messageInfo_CallContractLog.Size(m)
}
func (m *CallContractLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractLog.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractLog proto.InternalMessageInfo

func (m *CallContractLog) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *CallContractLog) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CallContractLog) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type LocalDataLog struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDataLog) Reset()         { *m = LocalDataLog{} }
func (m *LocalDataLog) String() string { return proto.CompactTextString(m) }
func (*LocalDataLog) ProtoMessage()    {}
func (*LocalDataLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{7}
}

func (m *LocalDataLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDataLog.Unmarshal(m, b)
}
func (m *LocalDataLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDataLog.Marshal(b, m, deterministic)
}
func (m *LocalDataLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDataLog.Merge(m, src)
}
func (m *LocalDataLog) XXX_Size() int {
	return xxx_messageInfo_LocalDataLog.Size(m)
}
func (m *LocalDataLog) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDataLog.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDataLog proto.InternalMessageInfo

func (m *LocalDataLog) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalDataLog) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*WasmAction)(nil), "types.wasmAction")
	proto.RegisterType((*WasmCreate)(nil), "types.wasmCreate")
	proto.RegisterType((*WasmCall)(nil), "types.wasmCall")
	proto.RegisterType((*QueryCheckContract)(nil), "types.queryCheckContract")
	proto.RegisterType((*CustomLog)(nil), "types.customLog")
	proto.RegisterType((*CreateContractLog)(nil), "types.createContractLog")
	proto.RegisterType((*CallContractLog)(nil), "types.callContractLog")
	proto.RegisterType((*LocalDataLog)(nil), "types.localDataLog")
}

func init() {
	proto.RegisterFile("wasm.proto", fileDescriptor_7d78909ad64e3bbb)
}

var fileDescriptor_7d78909ad64e3bbb = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xa6, 0x8d, 0xcd, 0x58, 0xac, 0x5d, 0xa4, 0x04, 0x0f, 0x1a, 0x16, 0x84, 0x80,
	0xd0, 0x83, 0x8a, 0x17, 0x4f, 0x5a, 0x0f, 0x3d, 0x78, 0xda, 0xbb, 0xc2, 0xba, 0x5d, 0x6d, 0xe9,
	0x26, 0x5b, 0x37, 0x13, 0x25, 0xdf, 0x5e, 0xf6, 0x4f, 0x25, 0x07, 0xf1, 0xe0, 0x6d, 0x66, 0xf6,
	0x37, 0xef, 0x65, 0x66, 0x02, 0xf0, 0xc5, 0xeb, 0x72, 0xbe, 0x33, 0x1a, 0x35, 0x19, 0x62, 0xbb,
	0x93, 0x35, 0x6d, 0x7d, 0xf1, 0x5e, 0xe0, 0x46, 0x57, 0xe4, 0x12, 0x12, 0x61, 0x24, 0x47, 0x99,
	0x45, 0x79, 0x54, 0x1c, 0x5e, 0x4d, 0xe7, 0x8e, 0x9a, 0x5b, 0x64, 0xe1, 0x1e, 0x96, 0x3d, 0x16,
	0x10, 0x72, 0x01, 0x03, 0xc1, 0x95, 0xca, 0xfa, 0x0e, 0x9d, 0x74, 0x51, 0xae, 0xd4, 0xb2, 0xc7,
	0xdc, 0x33, 0x39, 0x82, 0x3e, 0xb6, 0x59, 0x9c, 0x47, 0xc5, 0x90, 0xf5, 0xb1, 0x7d, 0x38, 0x80,
	0xe1, 0x27, 0x57, 0x8d, 0xa4, 0x37, 0xde, 0xda, 0xeb, 0x12, 0x02, 0x83, 0x8a, 0x97, 0xde, 0x38,
	0x65, 0x2e, 0xb6, 0x35, 0xa1, 0x57, 0xd2, 0x39, 0x8c, 0x99, 0x8b, 0xe9, 0x0b, 0x8c, 0xf6, 0x16,
	0xe4, 0x14, 0x46, 0x42, 0x57, 0x68, 0xb8, 0xc0, 0xd0, 0xf7, 0x93, 0x93, 0x19, 0x24, 0xa5, 0xc4,
	0xb5, 0x5e, 0xb9, 0xee, 0x94, 0x85, 0x8c, 0x9c, 0x01, 0xec, 0xb8, 0xe1, 0xa5, 0x44, 0x69, 0xea,
	0x2c, 0xce, 0xe3, 0x22, 0x66, 0x9d, 0x0a, 0x2d, 0x80, 0x7c, 0x34, 0xd2, 0xb4, 0x8b, 0xb5, 0x14,
	0xdb, 0xc5, 0x5e, 0xed, 0x97, 0xaf, 0xa3, 0xe7, 0x90, 0x8a, 0xa6, 0x46, 0x5d, 0x3e, 0xe9, 0x77,
	0x0b, 0x6c, 0xaa, 0x37, 0x9d, 0x45, 0x79, 0x6c, 0x01, 0x1b, 0xd3, 0x3b, 0x98, 0xfa, 0x55, 0xed,
	0x65, 0x02, 0xf8, 0xe7, 0x9c, 0x69, 0x98, 0xf3, 0x19, 0x26, 0x76, 0x7d, 0xdd, 0xd6, 0xff, 0x8c,
	0x3b, 0x83, 0xc4, 0xc8, 0xba, 0x51, 0x18, 0x2e, 0x10, 0x32, 0x7a, 0x0b, 0x63, 0xa5, 0x05, 0x57,
	0x8f, 0x1c, 0xb9, 0xd5, 0x3e, 0x86, 0x78, 0x2b, 0x5b, 0x27, 0x3b, 0x66, 0x36, 0x24, 0x27, 0xe1,
	0x4e, 0x61, 0xfb, 0x3e, 0x79, 0x4d, 0xdc, 0xdf, 0x73, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0xc2, 0x9e, 0xbb, 0x4b, 0x02, 0x00, 0x00,
}
