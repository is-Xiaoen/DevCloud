package main

import (
	"context"
	"fmt"
	"log"

	"122.51.31.227/go-course/go18/skills/rpc/hello_world/service"
)

func main() {
	client, err := service.NewHelloServiceClient("localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// SDK 提供的方法调用
	reply, err := client.Hello(context.Background(), "bob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
