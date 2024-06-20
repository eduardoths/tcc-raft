// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/raft.proto

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

type HeartbeatArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32       `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     string      `protobuf:"bytes,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PrevLogIndex int32       `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  int32       `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit int32       `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
}

func (x *HeartbeatArgs) Reset() {
	*x = HeartbeatArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatArgs) ProtoMessage() {}

func (x *HeartbeatArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatArgs.ProtoReflect.Descriptor instead.
func (*HeartbeatArgs) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *HeartbeatArgs) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *HeartbeatArgs) GetPrevLogIndex() int32 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *HeartbeatArgs) GetPrevLogTerm() int32 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *HeartbeatArgs) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *HeartbeatArgs) GetLeaderCommit() int32 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type HeartbeatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Term      int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	NextIndex int32 `protobuf:"varint,3,opt,name=nextIndex,proto3" json:"nextIndex,omitempty"`
}

func (x *HeartbeatReply) Reset() {
	*x = HeartbeatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReply) ProtoMessage() {}

func (x *HeartbeatReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatReply.ProtoReflect.Descriptor instead.
func (*HeartbeatReply) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *HeartbeatReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *HeartbeatReply) GetNextIndex() int32 {
	if x != nil {
		return x.NextIndex
	}
	return 0
}

type RequestVoteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId string `protobuf:"bytes,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
}

func (x *RequestVoteArgs) Reset() {
	*x = RequestVoteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteArgs) ProtoMessage() {}

func (x *RequestVoteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteArgs.ProtoReflect.Descriptor instead.
func (*RequestVoteArgs) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{2}
}

func (x *RequestVoteArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteArgs) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

type RequestVoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteReply) Reset() {
	*x = RequestVoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteReply) ProtoMessage() {}

func (x *RequestVoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteReply.ProtoReflect.Descriptor instead.
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{3}
}

func (x *RequestVoteReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteReply) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int32   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Index   int32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Command *LogCMD `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{4}
}

func (x *LogEntry) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LogEntry) GetCommand() *LogCMD {
	if x != nil {
		return x.Command
	}
	return nil
}

type LogCMD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LogCMD) Reset() {
	*x = LogCMD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCMD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCMD) ProtoMessage() {}

func (x *LogCMD) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCMD.ProtoReflect.Descriptor instead.
func (*LogCMD) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{5}
}

func (x *LogCMD) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *LogCMD) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LogCMD) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

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
		mi := &file_proto_raft_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetArgs) ProtoMessage() {}

func (x *SetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[6]
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
	return file_proto_raft_proto_rawDescGZIP(), []int{6}
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
		mi := &file_proto_raft_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReply) ProtoMessage() {}

func (x *SetReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[7]
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
	return file_proto_raft_proto_rawDescGZIP(), []int{7}
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

type SearchLogArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *SearchLogArgs) Reset() {
	*x = SearchLogArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLogArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLogArgs) ProtoMessage() {}

func (x *SearchLogArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLogArgs.ProtoReflect.Descriptor instead.
func (*SearchLogArgs) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{8}
}

func (x *SearchLogArgs) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type SearchLogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command  *LogCMD `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Found    bool    `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
	Commited bool    `protobuf:"varint,3,opt,name=commited,proto3" json:"commited,omitempty"`
}

func (x *SearchLogReply) Reset() {
	*x = SearchLogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raft_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLogReply) ProtoMessage() {}

func (x *SearchLogReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raft_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLogReply.ProtoReflect.Descriptor instead.
func (*SearchLogReply) Descriptor() ([]byte, []int) {
	return file_proto_raft_proto_rawDescGZIP(), []int{9}
}

func (x *SearchLogReply) GetCommand() *LogCMD {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *SearchLogReply) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *SearchLogReply) GetCommited() bool {
	if x != nil {
		return x.Commited
	}
	return false
}

var File_proto_raft_proto protoreflect.FileDescriptor

var file_proto_raft_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x66, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x28, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x5c,
	0x0a, 0x0e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x47, 0x0a, 0x0f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a,
	0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22,
	0x5c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4c, 0x6f,
	0x67, 0x43, 0x4d, 0x44, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x4e, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x43, 0x4d, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x36, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4c, 0x6f, 0x67, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x6a, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x26, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x43, 0x4d, 0x44,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x32, 0xda, 0x01, 0x0a, 0x04,
	0x52, 0x61, 0x66, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12,
	0x13, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x12, 0x0d, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73,
	0x1a, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x36, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x12, 0x13, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x41, 0x72,
	0x67, 0x73, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x75, 0x61, 0x72, 0x64, 0x6f, 0x74, 0x68,
	0x73, 0x2f, 0x74, 0x63, 0x63, 0x2d, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raft_proto_rawDescOnce sync.Once
	file_proto_raft_proto_rawDescData = file_proto_raft_proto_rawDesc
)

func file_proto_raft_proto_rawDescGZIP() []byte {
	file_proto_raft_proto_rawDescOnce.Do(func() {
		file_proto_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raft_proto_rawDescData)
	})
	return file_proto_raft_proto_rawDescData
}

var file_proto_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_raft_proto_goTypes = []any{
	(*HeartbeatArgs)(nil),    // 0: raft.HeartbeatArgs
	(*HeartbeatReply)(nil),   // 1: raft.HeartbeatReply
	(*RequestVoteArgs)(nil),  // 2: raft.RequestVoteArgs
	(*RequestVoteReply)(nil), // 3: raft.RequestVoteReply
	(*LogEntry)(nil),         // 4: raft.LogEntry
	(*LogCMD)(nil),           // 5: raft.LogCMD
	(*SetArgs)(nil),          // 6: raft.SetArgs
	(*SetReply)(nil),         // 7: raft.SetReply
	(*SearchLogArgs)(nil),    // 8: raft.SearchLogArgs
	(*SearchLogReply)(nil),   // 9: raft.SearchLogReply
}
var file_proto_raft_proto_depIdxs = []int32{
	4, // 0: raft.HeartbeatArgs.entries:type_name -> raft.LogEntry
	5, // 1: raft.LogEntry.command:type_name -> raft.LogCMD
	5, // 2: raft.SearchLogReply.command:type_name -> raft.LogCMD
	2, // 3: raft.Raft.RequestVote:input_type -> raft.RequestVoteArgs
	0, // 4: raft.Raft.Heartbeat:input_type -> raft.HeartbeatArgs
	6, // 5: raft.Raft.Set:input_type -> raft.SetArgs
	8, // 6: raft.Raft.SearchLog:input_type -> raft.SearchLogArgs
	3, // 7: raft.Raft.RequestVote:output_type -> raft.RequestVoteReply
	1, // 8: raft.Raft.Heartbeat:output_type -> raft.HeartbeatReply
	7, // 9: raft.Raft.Set:output_type -> raft.SetReply
	9, // 10: raft.Raft.SearchLog:output_type -> raft.SearchLogReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_raft_proto_init() }
func file_proto_raft_proto_init() {
	if File_proto_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_raft_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HeartbeatArgs); i {
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
		file_proto_raft_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HeartbeatReply); i {
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
		file_proto_raft_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RequestVoteArgs); i {
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
		file_proto_raft_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RequestVoteReply); i {
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
		file_proto_raft_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LogEntry); i {
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
		file_proto_raft_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LogCMD); i {
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
		file_proto_raft_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
		file_proto_raft_proto_msgTypes[7].Exporter = func(v any, i int) any {
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
		file_proto_raft_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SearchLogArgs); i {
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
		file_proto_raft_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SearchLogReply); i {
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
			RawDescriptor: file_proto_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raft_proto_goTypes,
		DependencyIndexes: file_proto_raft_proto_depIdxs,
		MessageInfos:      file_proto_raft_proto_msgTypes,
	}.Build()
	File_proto_raft_proto = out.File
	file_proto_raft_proto_rawDesc = nil
	file_proto_raft_proto_goTypes = nil
	file_proto_raft_proto_depIdxs = nil
}
