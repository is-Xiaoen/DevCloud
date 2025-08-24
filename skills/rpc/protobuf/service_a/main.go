package main

import (
	"context"
	"log"
	"net"

	"122.51.31.227/go-course/go18/skills/rpc/protobuf/hello_service"
	"google.golang.org/grpc"
)

func main() {
	// 首先是通过grpc.NewServer()构造一个gRPC服务对象
	grpcServer := grpc.NewServer()

	// SDK 提供 服务实现对象的注册
	hello_service.RegisterHelloServiceServer(grpcServer, &HelloService{})

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 然后通过grpcServer.Serve(lis)在一个监听端口上提供gRPC服务
	grpcServer.Serve(lis)
}

var _ hello_service.HelloServiceServer = (*HelloService)(nil)

// 实现一个GRPC的对象, 并行实现了 HelloServiceServer 接口
// 该对象就可以注册给GRPC框架
type HelloService struct {
	hello_service.UnimplementedHelloServiceServer
}

func (h *HelloService) Hello(ctx context.Context, req *hello_service.Request) (*hello_service.Response, error) {
	return &hello_service.Response{
		Value: "Hello " + req.Value,
	}, nil
}
