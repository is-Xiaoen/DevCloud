package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 把我们的对象注册成一个rpc的 receiver
	// 其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	// 所有注册的方法会放在“HelloService”服务空间之下
	// err := HelloService.FnName(req, resp)
	rpc.RegisterName("HelloService", new(HelloService))

	// RPC服务器还没启动起来
	// 然后我们建立一个唯一的TCP链接，
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	// 没Accept一个请求，就创建一个goroutie进行处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 前面都是tcp的知识, 到这个RPC就接管了
		// 因此 你可以认为 rpc 帮我们封装消息到函数调用的这个逻辑,
		// 提升了工作效率, 逻辑比较简洁，可以看看他代码
		go rpc.ServeConn(conn)
	}
}

// 要通过 net/rpc 把一个对象的方法 暴露为rpc，需要符合 net/rpc的方法签名: Hello(request string, reply *string) error
// FnName(req any, resp *any) error

type HelloService struct{}

// HTTP Framwork Handler Hello(request string, reply *string) error
func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}
