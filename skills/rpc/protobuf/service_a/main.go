package main

import (
	"context"
	"io"
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

// 接收来自客户端的流式请求，然后不断返回
func (h *HelloService) Channel(stream grpc.BidiStreamingServer[hello_service.Request, hello_service.Response]) error {
	for {
		// 读取客户端发送过来的数据
		msg, err := stream.Recv()
		if err != nil {
			// 如果遇到io.EOF表示客户端流被关闭
			if err == io.EOF {
				return nil
			}
			return err
		}
		// 处理消息并返回响应
		if err := stream.Send(&hello_service.Response{
			Value: "Hello " + msg.Value,
		}); err != nil {
			return err
		}
	}
}
