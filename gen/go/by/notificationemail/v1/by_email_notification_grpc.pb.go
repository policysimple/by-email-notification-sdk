// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package notificationemailv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotificationEmailAPIServiceClient is the client API for NotificationEmailAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationEmailAPIServiceClient interface {
	// SendEmail ...
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
	SendEmailWithCustomDomain(ctx context.Context, in *SendEmailWithCustomDomainRequest, opts ...grpc.CallOption) (*SendEmailWithCustomDomainResponse, error)
}

type notificationEmailAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationEmailAPIServiceClient(cc grpc.ClientConnInterface) NotificationEmailAPIServiceClient {
	return &notificationEmailAPIServiceClient{cc}
}

func (c *notificationEmailAPIServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/by.notificationemail.v1.NotificationEmailAPIService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationEmailAPIServiceClient) SendEmailWithCustomDomain(ctx context.Context, in *SendEmailWithCustomDomainRequest, opts ...grpc.CallOption) (*SendEmailWithCustomDomainResponse, error) {
	out := new(SendEmailWithCustomDomainResponse)
	err := c.cc.Invoke(ctx, "/by.notificationemail.v1.NotificationEmailAPIService/SendEmailWithCustomDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationEmailAPIServiceServer is the server API for NotificationEmailAPIService service.
// All implementations should embed UnimplementedNotificationEmailAPIServiceServer
// for forward compatibility
type NotificationEmailAPIServiceServer interface {
	// SendEmail ...
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
	SendEmailWithCustomDomain(context.Context, *SendEmailWithCustomDomainRequest) (*SendEmailWithCustomDomainResponse, error)
}

// UnimplementedNotificationEmailAPIServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotificationEmailAPIServiceServer struct {
}

func (UnimplementedNotificationEmailAPIServiceServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedNotificationEmailAPIServiceServer) SendEmailWithCustomDomain(context.Context, *SendEmailWithCustomDomainRequest) (*SendEmailWithCustomDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailWithCustomDomain not implemented")
}

// UnsafeNotificationEmailAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationEmailAPIServiceServer will
// result in compilation errors.
type UnsafeNotificationEmailAPIServiceServer interface {
	mustEmbedUnimplementedNotificationEmailAPIServiceServer()
}

func RegisterNotificationEmailAPIServiceServer(s grpc.ServiceRegistrar, srv NotificationEmailAPIServiceServer) {
	s.RegisterService(&NotificationEmailAPIService_ServiceDesc, srv)
}

func _NotificationEmailAPIService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationEmailAPIServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/by.notificationemail.v1.NotificationEmailAPIService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationEmailAPIServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationEmailAPIService_SendEmailWithCustomDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailWithCustomDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationEmailAPIServiceServer).SendEmailWithCustomDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/by.notificationemail.v1.NotificationEmailAPIService/SendEmailWithCustomDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationEmailAPIServiceServer).SendEmailWithCustomDomain(ctx, req.(*SendEmailWithCustomDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationEmailAPIService_ServiceDesc is the grpc.ServiceDesc for NotificationEmailAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationEmailAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "by.notificationemail.v1.NotificationEmailAPIService",
	HandlerType: (*NotificationEmailAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _NotificationEmailAPIService_SendEmail_Handler,
		},
		{
			MethodName: "SendEmailWithCustomDomain",
			Handler:    _NotificationEmailAPIService_SendEmailWithCustomDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "by/notificationemail/v1/by_email_notification.proto",
}
