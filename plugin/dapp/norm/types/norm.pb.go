// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.1
// source: norm.proto

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

type Norm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NormId     []byte `protobuf:"bytes,1,opt,name=normId,proto3" json:"normId,omitempty"`
	CreateTime int64  `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Key        []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Norm) Reset() {
	*x = Norm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_norm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Norm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Norm) ProtoMessage() {}

func (x *Norm) ProtoReflect() protoreflect.Message {
	mi := &file_norm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Norm.ProtoReflect.Descriptor instead.
func (*Norm) Descriptor() ([]byte, []int) {
	return file_norm_proto_rawDescGZIP(), []int{0}
}

func (x *Norm) GetNormId() []byte {
	if x != nil {
		return x.NormId
	}
	return nil
}

func (x *Norm) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Norm) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Norm) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type NormAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*NormAction_Nput
	Value isNormAction_Value `protobuf_oneof:"value"`
	Ty    int32              `protobuf:"varint,5,opt,name=ty,proto3" json:"ty,omitempty"`
}

func (x *NormAction) Reset() {
	*x = NormAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_norm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormAction) ProtoMessage() {}

func (x *NormAction) ProtoReflect() protoreflect.Message {
	mi := &file_norm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormAction.ProtoReflect.Descriptor instead.
func (*NormAction) Descriptor() ([]byte, []int) {
	return file_norm_proto_rawDescGZIP(), []int{1}
}

func (m *NormAction) GetValue() isNormAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *NormAction) GetNput() *NormPut {
	if x, ok := x.GetValue().(*NormAction_Nput); ok {
		return x.Nput
	}
	return nil
}

func (x *NormAction) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type isNormAction_Value interface {
	isNormAction_Value()
}

type NormAction_Nput struct {
	Nput *NormPut `protobuf:"bytes,1,opt,name=nput,proto3,oneof"`
}

func (*NormAction_Nput) isNormAction_Value() {}

type NormPut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NormPut) Reset() {
	*x = NormPut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_norm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormPut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormPut) ProtoMessage() {}

func (x *NormPut) ProtoReflect() protoreflect.Message {
	mi := &file_norm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormPut.ProtoReflect.Descriptor instead.
func (*NormPut) Descriptor() ([]byte, []int) {
	return file_norm_proto_rawDescGZIP(), []int{2}
}

func (x *NormPut) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *NormPut) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type NormGetKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *NormGetKey) Reset() {
	*x = NormGetKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_norm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormGetKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormGetKey) ProtoMessage() {}

func (x *NormGetKey) ProtoReflect() protoreflect.Message {
	mi := &file_norm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormGetKey.ProtoReflect.Descriptor instead.
func (*NormGetKey) Descriptor() ([]byte, []int) {
	return file_norm_proto_rawDescGZIP(), []int{3}
}

func (x *NormGetKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_norm_proto protoreflect.FileDescriptor

var file_norm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x04, 0x4e, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x72,
	0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x0a, 0x4e,
	0x6f, 0x72, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4e, 0x6f, 0x72, 0x6d, 0x50, 0x75, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x79, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x07, 0x4e, 0x6f, 0x72, 0x6d,
	0x50, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x0a, 0x4e,
	0x6f, 0x72, 0x6d, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_norm_proto_rawDescOnce sync.Once
	file_norm_proto_rawDescData = file_norm_proto_rawDesc
)

func file_norm_proto_rawDescGZIP() []byte {
	file_norm_proto_rawDescOnce.Do(func() {
		file_norm_proto_rawDescData = protoimpl.X.CompressGZIP(file_norm_proto_rawDescData)
	})
	return file_norm_proto_rawDescData
}

var file_norm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_norm_proto_goTypes = []interface{}{
	(*Norm)(nil),       // 0: types.Norm
	(*NormAction)(nil), // 1: types.NormAction
	(*NormPut)(nil),    // 2: types.NormPut
	(*NormGetKey)(nil), // 3: types.NormGetKey
}
var file_norm_proto_depIdxs = []int32{
	2, // 0: types.NormAction.nput:type_name -> types.NormPut
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_norm_proto_init() }
func file_norm_proto_init() {
	if File_norm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_norm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Norm); i {
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
		file_norm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormAction); i {
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
		file_norm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormPut); i {
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
		file_norm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormGetKey); i {
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
	file_norm_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*NormAction_Nput)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_norm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_norm_proto_goTypes,
		DependencyIndexes: file_norm_proto_depIdxs,
		MessageInfos:      file_norm_proto_msgTypes,
	}.Build()
	File_norm_proto = out.File
	file_norm_proto_rawDesc = nil
	file_norm_proto_goTypes = nil
	file_norm_proto_depIdxs = nil
}
