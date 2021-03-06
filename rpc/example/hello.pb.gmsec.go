// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package example

import (
	context "context"
	micro "github.com/gmsec/micro"
	client "github.com/gmsec/micro/client"
	server "github.com/gmsec/micro/server"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface
var _ server.Server
var _ client.Client
var _ micro.Service

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	// 定义SayHello方法
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloClient struct {
	cc client.Client
}

// GetHelloName get client name(package.class)
func GetHelloName() string {
	return "proto.Hello"
}

// GetHelloClient get client by clientname
func GetHelloClient() HelloClient {
	cc := micro.GetClient(GetHelloName())
	return &helloClient{cc}
}

// GetHelloClientByName get client by custom name
func GetHelloClientByName(name string) HelloClient {
	cc := micro.GetClient(name)
	return &helloClient{cc}
}

func NewHelloClient(cc client.Client) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	conn, err := c.cc.Next()
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	out := new(HelloReply)
	err = conn.Invoke(ctx, "/proto.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	// 定义SayHello方法
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloServer(s server.Server, srv HelloServer) {
	s.GetServer().RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/hello.proto",
}
