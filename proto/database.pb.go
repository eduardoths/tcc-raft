// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/database.proto

package proto

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

type SetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetArgs) Reset() {
	*x = SetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetArgs) ProtoMessage() {}

func (x *SetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetArgs.ProtoReflect.Descriptor instead.
func (*SetArgs) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{0}
}

func (x *SetArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetArgs) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Noted bool  `protobuf:"varint,2,opt,name=noted,proto3" json:"noted,omitempty"`
}

func (x *SetReply) Reset() {
	*x = SetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReply) ProtoMessage() {}

func (x *SetReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReply.ProtoReflect.Descriptor instead.
func (*SetReply) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{1}
}

func (x *SetReply) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SetReply) GetNoted() bool {
	if x != nil {
		return x.Noted
	}
	return false
}

type DeleteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteArgs) Reset() {
	*x = DeleteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArgs) ProtoMessage() {}

func (x *DeleteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArgs.ProtoReflect.Descriptor instead.
func (*DeleteArgs) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Noted bool  `protobuf:"varint,2,opt,name=noted,proto3" json:"noted,omitempty"`
}

func (x *DeleteReply) Reset() {
	*x = DeleteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReply) ProtoMessage() {}

func (x *DeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReply.ProtoReflect.Descriptor instead.
func (*DeleteReply) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteReply) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DeleteReply) GetNoted() bool {
	if x != nil {
		return x.Noted
	}
	return false
}

type GetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetArgs) Reset() {
	*x = GetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArgs) ProtoMessage() {}

func (x *GetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArgs.ProtoReflect.Descriptor instead.
func (*GetArgs) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{4}
}

func (x *GetArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{5}
}

func (x *GetReply) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type EmptyDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyDB) Reset() {
	*x = EmptyDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyDB) ProtoMessage() {}

func (x *EmptyDB) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyDB.ProtoReflect.Descriptor instead.
func (*EmptyDB) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{6}
}

type GetLeaderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Term int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *GetLeaderReply) Reset() {
	*x = GetLeaderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaderReply) ProtoMessage() {}

func (x *GetLeaderReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaderReply.ProtoReflect.Descriptor instead.
func (*GetLeaderReply) Descriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{7}
}

func (x *GetLeaderReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLeaderReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

var File_proto_database_proto protoreflect.FileDescriptor

var file_proto_database_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x36, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x22, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x39, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x64, 0x22, 0x1b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x20, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x42, 0x22, 0x34, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x32, 0xbf, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x32, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x42, 0x1a, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x75, 0x61, 0x72, 0x64, 0x6f, 0x74, 0x68, 0x73, 0x2f, 0x74,
	0x63, 0x63, 0x2d, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_database_proto_rawDescOnce sync.Once
	file_proto_database_proto_rawDescData = file_proto_database_proto_rawDesc
)

func file_proto_database_proto_rawDescGZIP() []byte {
	file_proto_database_proto_rawDescOnce.Do(func() {
		file_proto_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_database_proto_rawDescData)
	})
	return file_proto_database_proto_rawDescData
}

var file_proto_database_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_database_proto_goTypes = []any{
	(*SetArgs)(nil),        // 0: proto.SetArgs
	(*SetReply)(nil),       // 1: proto.SetReply
	(*DeleteArgs)(nil),     // 2: proto.DeleteArgs
	(*DeleteReply)(nil),    // 3: proto.DeleteReply
	(*GetArgs)(nil),        // 4: proto.GetArgs
	(*GetReply)(nil),       // 5: proto.GetReply
	(*EmptyDB)(nil),        // 6: proto.EmptyDB
	(*GetLeaderReply)(nil), // 7: proto.GetLeaderReply
}
var file_proto_database_proto_depIdxs = []int32{
	0, // 0: proto.Database.Set:input_type -> proto.SetArgs
	2, // 1: proto.Database.Delete:input_type -> proto.DeleteArgs
	4, // 2: proto.Database.Get:input_type -> proto.GetArgs
	6, // 3: proto.Database.GetLeader:input_type -> proto.EmptyDB
	1, // 4: proto.Database.Set:output_type -> proto.SetReply
	3, // 5: proto.Database.Delete:output_type -> proto.DeleteReply
	5, // 6: proto.Database.Get:output_type -> proto.GetReply
	7, // 7: proto.Database.GetLeader:output_type -> proto.GetLeaderReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_database_proto_init() }
func file_proto_database_proto_init() {
	if File_proto_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_database_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetArgs); i {
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
		file_proto_database_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetReply); i {
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
		file_proto_database_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteArgs); i {
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
		file_proto_database_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteReply); i {
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
		file_proto_database_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetArgs); i {
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
		file_proto_database_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetReply); i {
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
		file_proto_database_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyDB); i {
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
		file_proto_database_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetLeaderReply); i {
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
			RawDescriptor: file_proto_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_database_proto_goTypes,
		DependencyIndexes: file_proto_database_proto_depIdxs,
		MessageInfos:      file_proto_database_proto_msgTypes,
	}.Build()
	File_proto_database_proto = out.File
	file_proto_database_proto_rawDesc = nil
	file_proto_database_proto_goTypes = nil
	file_proto_database_proto_depIdxs = nil
}
