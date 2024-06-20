// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/raft.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Raft_RequestVote_FullMethodName = "/raft.Raft/RequestVote"
	Raft_Heartbeat_FullMethodName   = "/raft.Raft/Heartbeat"
	Raft_Set_FullMethodName         = "/raft.Raft/Set"
	Raft_SearchLog_FullMethodName   = "/raft.Raft/SearchLog"
)

// RaftClient is the client API for Raft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftClient interface {
	RequestVote(ctx context.Context, in *RequestVoteArgs, opts ...grpc.CallOption) (*RequestVoteReply, error)
	Heartbeat(ctx context.Context, in *HeartbeatArgs, opts ...grpc.CallOption) (*HeartbeatReply, error)
	Set(ctx context.Context, in *SetArgs, opts ...grpc.CallOption) (*SetReply, error)
	SearchLog(ctx context.Context, in *SearchLogArgs, opts ...grpc.CallOption) (*SearchLogReply, error)
}

type raftClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftClient(cc grpc.ClientConnInterface) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) RequestVote(ctx context.Context, in *RequestVoteArgs, opts ...grpc.CallOption) (*RequestVoteReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestVoteReply)
	err := c.cc.Invoke(ctx, Raft_RequestVote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Heartbeat(ctx context.Context, in *HeartbeatArgs, opts ...grpc.CallOption) (*HeartbeatReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatReply)
	err := c.cc.Invoke(ctx, Raft_Heartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) Set(ctx context.Context, in *SetArgs, opts ...grpc.CallOption) (*SetReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetReply)
	err := c.cc.Invoke(ctx, Raft_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) SearchLog(ctx context.Context, in *SearchLogArgs, opts ...grpc.CallOption) (*SearchLogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchLogReply)
	err := c.cc.Invoke(ctx, Raft_SearchLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftServer is the server API for Raft service.
// All implementations must embed UnimplementedRaftServer
// for forward compatibility
type RaftServer interface {
	RequestVote(context.Context, *RequestVoteArgs) (*RequestVoteReply, error)
	Heartbeat(context.Context, *HeartbeatArgs) (*HeartbeatReply, error)
	Set(context.Context, *SetArgs) (*SetReply, error)
	SearchLog(context.Context, *SearchLogArgs) (*SearchLogReply, error)
	mustEmbedUnimplementedRaftServer()
}

// UnimplementedRaftServer must be embedded to have forward compatible implementations.
type UnimplementedRaftServer struct {
}

func (UnimplementedRaftServer) RequestVote(context.Context, *RequestVoteArgs) (*RequestVoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedRaftServer) Heartbeat(context.Context, *HeartbeatArgs) (*HeartbeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedRaftServer) Set(context.Context, *SetArgs) (*SetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedRaftServer) SearchLog(context.Context, *SearchLogArgs) (*SearchLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLog not implemented")
}
func (UnimplementedRaftServer) mustEmbedUnimplementedRaftServer() {}

// UnsafeRaftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftServer will
// result in compilation errors.
type UnsafeRaftServer interface {
	mustEmbedUnimplementedRaftServer()
}

func RegisterRaftServer(s grpc.ServiceRegistrar, srv RaftServer) {
	s.RegisterService(&Raft_ServiceDesc, srv)
}

func _Raft_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Raft_RequestVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).RequestVote(ctx, req.(*RequestVoteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Raft_Heartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Heartbeat(ctx, req.(*HeartbeatArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Raft_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).Set(ctx, req.(*SetArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_SearchLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLogArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).SearchLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Raft_SearchLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).SearchLog(ctx, req.(*SearchLogArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// Raft_ServiceDesc is the grpc.ServiceDesc for Raft service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Raft_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raft.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _Raft_RequestVote_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Raft_Heartbeat_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Raft_Set_Handler,
		},
		{
			MethodName: "SearchLog",
			Handler:    _Raft_SearchLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/raft.proto",
}
