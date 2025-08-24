package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	stream, err := helloServiceClient.Channel(context.Background())
	if err != nil {
		panic(err)
	}

	// 首先是向服务端发送数据
	go func() {
		count := 1
		for {
			if err := stream.Send(&hello_service.Request{Value: fmt.Sprintf("[%d] hi", count)}); err != nil {
				log.Fatal(err)
			}
			count++
			time.Sleep(time.Second)
		}
	}()

	// 然后在循环中接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
