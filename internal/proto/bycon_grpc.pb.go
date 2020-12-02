// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BYCONClient is the client API for BYCON service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BYCONClient interface {
	PrePrepare(ctx context.Context, in *PrePrepareArgs, opts ...grpc.CallOption) (*empty.Empty, error)
	Prepare(ctx context.Context, in *PrepareArgs, opts ...grpc.CallOption) (*empty.Empty, error)
	Commit(ctx context.Context, in *CommitArgs, opts ...grpc.CallOption) (*empty.Empty, error)
	// view change...
	PreRequestVote(ctx context.Context, in *PreRequestVoteArgs, opts ...grpc.CallOption) (*PreRequestVoteReply, error)
	RequestVote(ctx context.Context, in *RequestVoteArgs, opts ...grpc.CallOption) (*empty.Empty, error)
	VoteConfirm(ctx context.Context, in *VoteConfirmArgs, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bYCONClient struct {
	cc grpc.ClientConnInterface
}

func NewBYCONClient(cc grpc.ClientConnInterface) BYCONClient {
	return &bYCONClient{cc}
}

func (c *bYCONClient) PrePrepare(ctx context.Context, in *PrePrepareArgs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.BYCON/PrePrepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bYCONClient) Prepare(ctx context.Context, in *PrepareArgs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.BYCON/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bYCONClient) Commit(ctx context.Context, in *CommitArgs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.BYCON/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bYCONClient) PreRequestVote(ctx context.Context, in *PreRequestVoteArgs, opts ...grpc.CallOption) (*PreRequestVoteReply, error) {
	out := new(PreRequestVoteReply)
	err := c.cc.Invoke(ctx, "/proto.BYCON/PreRequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bYCONClient) RequestVote(ctx context.Context, in *RequestVoteArgs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.BYCON/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bYCONClient) VoteConfirm(ctx context.Context, in *VoteConfirmArgs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.BYCON/VoteConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BYCONServer is the server API for BYCON service.
// All implementations must embed UnimplementedBYCONServer
// for forward compatibility
type BYCONServer interface {
	PrePrepare(context.Context, *PrePrepareArgs) (*empty.Empty, error)
	Prepare(context.Context, *PrepareArgs) (*empty.Empty, error)
	Commit(context.Context, *CommitArgs) (*empty.Empty, error)
	// view change...
	PreRequestVote(context.Context, *PreRequestVoteArgs) (*PreRequestVoteReply, error)
	RequestVote(context.Context, *RequestVoteArgs) (*empty.Empty, error)
	VoteConfirm(context.Context, *VoteConfirmArgs) (*empty.Empty, error)
	mustEmbedUnimplementedBYCONServer()
}

// UnimplementedBYCONServer must be embedded to have forward compatible implementations.
type UnimplementedBYCONServer struct {
}

func (UnimplementedBYCONServer) PrePrepare(context.Context, *PrePrepareArgs) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrePrepare not implemented")
}
func (UnimplementedBYCONServer) Prepare(context.Context, *PrepareArgs) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedBYCONServer) Commit(context.Context, *CommitArgs) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedBYCONServer) PreRequestVote(context.Context, *PreRequestVoteArgs) (*PreRequestVoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreRequestVote not implemented")
}
func (UnimplementedBYCONServer) RequestVote(context.Context, *RequestVoteArgs) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedBYCONServer) VoteConfirm(context.Context, *VoteConfirmArgs) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteConfirm not implemented")
}
func (UnimplementedBYCONServer) mustEmbedUnimplementedBYCONServer() {}

// UnsafeBYCONServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BYCONServer will
// result in compilation errors.
type UnsafeBYCONServer interface {
	mustEmbedUnimplementedBYCONServer()
}

func RegisterBYCONServer(s grpc.ServiceRegistrar, srv BYCONServer) {
	s.RegisterService(&_BYCON_serviceDesc, srv)
}

func _BYCON_PrePrepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrePrepareArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BYCONServer).PrePrepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BYCON/PrePrepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BYCONServer).PrePrepare(ctx, req.(*PrePrepareArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _BYCON_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BYCONServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BYCON/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BYCONServer).Prepare(ctx, req.(*PrepareArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _BYCON_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BYCONServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BYCON/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BYCONServer).Commit(ctx, req.(*CommitArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _BYCON_PreRequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreRequestVoteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BYCONServer).PreRequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BYCON/PreRequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BYCONServer).PreRequestVote(ctx, req.(*PreRequestVoteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _BYCON_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BYCONServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BYCON/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BYCONServer).RequestVote(ctx, req.(*RequestVoteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _BYCON_VoteConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteConfirmArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BYCONServer).VoteConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BYCON/VoteConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BYCONServer).VoteConfirm(ctx, req.(*VoteConfirmArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _BYCON_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BYCON",
	HandlerType: (*BYCONServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrePrepare",
			Handler:    _BYCON_PrePrepare_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _BYCON_Prepare_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _BYCON_Commit_Handler,
		},
		{
			MethodName: "PreRequestVote",
			Handler:    _BYCON_PreRequestVote_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _BYCON_RequestVote_Handler,
		},
		{
			MethodName: "VoteConfirm",
			Handler:    _BYCON_VoteConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bycon.proto",
}
