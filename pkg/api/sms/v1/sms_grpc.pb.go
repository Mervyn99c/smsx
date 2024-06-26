// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Rosas99/smsx.
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/sms/v1/sms.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SmsServer_DeleteOrder_FullMethodName = "/sms.v1.SmsServer/DeleteOrder"
)

// SmsServerClient is the client API for SmsServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsServerClient interface {
	DeleteOrder(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type smsServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServerClient(cc grpc.ClientConnInterface) SmsServerClient {
	return &smsServerClient{cc}
}

func (c *smsServerClient) DeleteOrder(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SmsServer_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServerServer is the server API for SmsServer service.
// All implementations must embed UnimplementedSmsServerServer
// for forward compatibility
type SmsServerServer interface {
	DeleteOrder(context.Context, *CreateTemplateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSmsServerServer()
}

// UnimplementedSmsServerServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServerServer struct {
}

func (UnimplementedSmsServerServer) DeleteOrder(context.Context, *CreateTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedSmsServerServer) mustEmbedUnimplementedSmsServerServer() {}

// UnsafeSmsServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServerServer will
// result in compilation errors.
type UnsafeSmsServerServer interface {
	mustEmbedUnimplementedSmsServerServer()
}

func RegisterSmsServerServer(s grpc.ServiceRegistrar, srv SmsServerServer) {
	s.RegisterService(&SmsServer_ServiceDesc, srv)
}

func _SmsServer_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServerServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsServer_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServerServer).DeleteOrder(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsServer_ServiceDesc is the grpc.ServiceDesc for SmsServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.v1.SmsServer",
	HandlerType: (*SmsServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteOrder",
			Handler:    _SmsServer_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sms/v1/sms.proto",
}
