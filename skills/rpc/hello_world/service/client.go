package service

import (
	"context"
	"net/rpc"
)

// 首先是通过rpc.Dial拨号RPC服务, 建立连接
// client, err := rpc.Dial("tcp", "localhost:1234")
// if err != nil {
// 	log.Fatal("dialing:", err)
// }

// // 然后通过client.Call调用具体的RPC方法
// // 在调用client.Call时:
// // 		第一个参数是用点号链接的RPC服务名字和方法名字，
// // 		第二个参数是 请求参数
// //      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值
// var reply string
// // HelloServiceClient.Hello(ctx, req) (resp, error)
// err = client.Call("HelloService.Hello", "bob", &reply)
// if err != nil {
// 	log.Fatal(err)
// }

func NewHelloServiceClient(address string) (HelloService, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{client: client}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (h *HelloServiceClient) Hello(ctx context.Context, request string) (string, error) {
	var reply string
	err := h.client.Call("HelloService.Hello", request, &reply)
	if err != nil {
		return "", err
	}
	return reply, nil
}
