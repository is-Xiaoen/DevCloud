package service

import "context"

// RPC 协议，约束服务端的实现，约束 客户端的调用
type HelloService interface {
	Hello(ctx context.Context, request string) (string, error)
}
