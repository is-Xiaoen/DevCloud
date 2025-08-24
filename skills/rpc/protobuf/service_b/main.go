package main

import (
	"context"
	"fmt"
	"log"

	"122.51.31.227/go-course/go18/skills/rpc/protobuf/hello_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// grpc.Dial负责和gRPC服务建立链接
	conn, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 使用SDK调用远程函数
	helloServiceClient := hello_service.NewHelloServiceClient(conn)
	resp, err := helloServiceClient.Hello(context.Background(), &hello_service.Request{
		Value: "bob",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Value)
}
