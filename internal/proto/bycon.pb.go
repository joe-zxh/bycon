// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: bycon.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PrePrepareArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View     uint32     `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Seq      uint32     `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Commands []*Command `protobuf:"bytes,3,rep,name=Commands,proto3" json:"Commands,omitempty"`
}

func (x *PrePrepareArgs) Reset() {
	*x = PrePrepareArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bycon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrePrepareArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrePrepareArgs) ProtoMessage() {}

func (x *PrePrepareArgs) ProtoReflect() protoreflect.Message {
	mi := &file_bycon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrePrepareArgs.ProtoReflect.Descriptor instead.
func (*PrePrepareArgs) Descriptor() ([]byte, []int) {
	return file_bycon_proto_rawDescGZIP(), []int{0}
}

func (x *PrePrepareArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *PrePrepareArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PrePrepareArgs) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

// 可以通过设置context的方式来获取Sender，但简单起见，直接设置Sender。
type PrepareArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View   uint32 `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Seq    uint32 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Digest []byte `protobuf:"bytes,3,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Sender uint32 `protobuf:"varint,4,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *PrepareArgs) Reset() {
	*x = PrepareArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bycon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareArgs) ProtoMessage() {}

func (x *PrepareArgs) ProtoReflect() protoreflect.Message {
	mi := &file_bycon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareArgs.ProtoReflect.Descriptor instead.
func (*PrepareArgs) Descriptor() ([]byte, []int) {
	return file_bycon_proto_rawDescGZIP(), []int{1}
}

func (x *PrepareArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *PrepareArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PrepareArgs) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *PrepareArgs) GetSender() uint32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

type CommitArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View   uint32 `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Seq    uint32 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Digest []byte `protobuf:"bytes,3,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Sender uint32 `protobuf:"varint,4,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *CommitArgs) Reset() {
	*x = CommitArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bycon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitArgs) ProtoMessage() {}

func (x *CommitArgs) ProtoReflect() protoreflect.Message {
	mi := &file_bycon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitArgs.ProtoReflect.Descriptor instead.
func (*CommitArgs) Descriptor() ([]byte, []int) {
	return file_bycon_proto_rawDescGZIP(), []int{2}
}

func (x *CommitArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *CommitArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *CommitArgs) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *CommitArgs) GetSender() uint32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bycon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_bycon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_bycon_proto_rawDescGZIP(), []int{3}
}

func (x *Command) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_bycon_proto protoreflect.FileDescriptor

var file_bycon_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x79, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x62, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x08, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x63, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x1d,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0xb6, 0x01,
	0x0a, 0x05, 0x42, 0x59, 0x43, 0x4f, 0x4e, 0x12, 0x3d, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x65, 0x2d, 0x7a, 0x78, 0x68, 0x2f, 0x62, 0x79, 0x63,
	0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bycon_proto_rawDescOnce sync.Once
	file_bycon_proto_rawDescData = file_bycon_proto_rawDesc
)

func file_bycon_proto_rawDescGZIP() []byte {
	file_bycon_proto_rawDescOnce.Do(func() {
		file_bycon_proto_rawDescData = protoimpl.X.CompressGZIP(file_bycon_proto_rawDescData)
	})
	return file_bycon_proto_rawDescData
}

var file_bycon_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bycon_proto_goTypes = []interface{}{
	(*PrePrepareArgs)(nil), // 0: proto.PrePrepareArgs
	(*PrepareArgs)(nil),    // 1: proto.PrepareArgs
	(*CommitArgs)(nil),     // 2: proto.CommitArgs
	(*Command)(nil),        // 3: proto.Command
	(*empty.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_bycon_proto_depIdxs = []int32{
	3, // 0: proto.PrePrepareArgs.Commands:type_name -> proto.Command
	0, // 1: proto.BYCON.PrePrepare:input_type -> proto.PrePrepareArgs
	1, // 2: proto.BYCON.Prepare:input_type -> proto.PrepareArgs
	2, // 3: proto.BYCON.Commit:input_type -> proto.CommitArgs
	4, // 4: proto.BYCON.PrePrepare:output_type -> google.protobuf.Empty
	4, // 5: proto.BYCON.Prepare:output_type -> google.protobuf.Empty
	4, // 6: proto.BYCON.Commit:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bycon_proto_init() }
func file_bycon_proto_init() {
	if File_bycon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bycon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrePrepareArgs); i {
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
		file_bycon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareArgs); i {
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
		file_bycon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitArgs); i {
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
		file_bycon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_bycon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bycon_proto_goTypes,
		DependencyIndexes: file_bycon_proto_depIdxs,
		MessageInfos:      file_bycon_proto_msgTypes,
	}.Build()
	File_bycon_proto = out.File
	file_bycon_proto_rawDesc = nil
	file_bycon_proto_goTypes = nil
	file_bycon_proto_depIdxs = nil
}
