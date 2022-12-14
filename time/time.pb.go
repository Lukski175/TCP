// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: time/time.proto

package time

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

type Syn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq int64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Syn) Reset() {
	*x = Syn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Syn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Syn) ProtoMessage() {}

func (x *Syn) ProtoReflect() protoreflect.Message {
	mi := &file_time_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Syn.ProtoReflect.Descriptor instead.
func (*Syn) Descriptor() ([]byte, []int) {
	return file_time_time_proto_rawDescGZIP(), []int{0}
}

func (x *Syn) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq        int64  `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Ack        int64  `protobuf:"varint,2,opt,name=ack,proto3" json:"ack,omitempty"`
	ClientData string `protobuf:"bytes,3,opt,name=clientData,proto3" json:"clientData,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_time_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_time_time_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Ack) GetAck() int64 {
	if x != nil {
		return x.Ack
	}
	return 0
}

func (x *Ack) GetClientData() string {
	if x != nil {
		return x.ClientData
	}
	return ""
}

type Synack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SynSeq int64 `protobuf:"varint,1,opt,name=synSeq,proto3" json:"synSeq,omitempty"`
	AckSeq int64 `protobuf:"varint,2,opt,name=ackSeq,proto3" json:"ackSeq,omitempty"`
}

func (x *Synack) Reset() {
	*x = Synack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_time_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Synack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Synack) ProtoMessage() {}

func (x *Synack) ProtoReflect() protoreflect.Message {
	mi := &file_time_time_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Synack.ProtoReflect.Descriptor instead.
func (*Synack) Descriptor() ([]byte, []int) {
	return file_time_time_proto_rawDescGZIP(), []int{2}
}

func (x *Synack) GetSynSeq() int64 {
	if x != nil {
		return x.SynSeq
	}
	return 0
}

func (x *Synack) GetAckSeq() int64 {
	if x != nil {
		return x.AckSeq
	}
	return 0
}

var File_time_time_proto protoreflect.FileDescriptor

var file_time_time_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x03, 0x73, 0x79, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x22, 0x49, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x06, 0x73,
	0x79, 0x6e, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x53, 0x65, 0x71, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x53, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x63, 0x6b, 0x53, 0x65, 0x71, 0x32, 0x50, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x12, 0x26, 0x0a, 0x09,
	0x67, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x61, 0x63, 0x6b, 0x12, 0x09, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x73, 0x79, 0x6e, 0x1a, 0x0c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x73, 0x79, 0x6e, 0x61,
	0x63, 0x6b, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x6b, 0x12,
	0x09, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61, 0x63, 0x6b, 0x1a, 0x09, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x61, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75,
	0x6b, 0x73, 0x6b, 0x69, 0x31, 0x37, 0x35, 0x2f, 0x54, 0x43, 0x50, 0x3b, 0x74, 0x69, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_time_time_proto_rawDescOnce sync.Once
	file_time_time_proto_rawDescData = file_time_time_proto_rawDesc
)

func file_time_time_proto_rawDescGZIP() []byte {
	file_time_time_proto_rawDescOnce.Do(func() {
		file_time_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_time_proto_rawDescData)
	})
	return file_time_time_proto_rawDescData
}

var file_time_time_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_time_time_proto_goTypes = []interface{}{
	(*Syn)(nil),    // 0: time.syn
	(*Ack)(nil),    // 1: time.ack
	(*Synack)(nil), // 2: time.synack
}
var file_time_time_proto_depIdxs = []int32{
	0, // 0: time.tcp.getSynack:input_type -> time.syn
	1, // 1: time.tcp.sendAck:input_type -> time.ack
	2, // 2: time.tcp.getSynack:output_type -> time.synack
	1, // 3: time.tcp.sendAck:output_type -> time.ack
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_time_time_proto_init() }
func file_time_time_proto_init() {
	if File_time_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_time_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Syn); i {
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
		file_time_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_time_time_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Synack); i {
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
			RawDescriptor: file_time_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_time_time_proto_goTypes,
		DependencyIndexes: file_time_time_proto_depIdxs,
		MessageInfos:      file_time_time_proto_msgTypes,
	}.Build()
	File_time_time_proto = out.File
	file_time_time_proto_rawDesc = nil
	file_time_time_proto_goTypes = nil
	file_time_time_proto_depIdxs = nil
}
