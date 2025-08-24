# protobuf

```go
// RPC 协议，约束服务端的实现，约束 客户端的调用
type HelloService interface {
	Hello(ctx context.Context, request string) (string, error)
}
```

+ RPC方法的定义
+ PRC数据结构的定义

## 插件安装

```sh
➜  protobuf go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go: downloading google.golang.org/protobuf v1.36.8
➜  protobuf go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go: downloading google.golang.org/grpc v1.75.0
```


## 接口声明

```proto
syntax = "proto3";

package hello;
option go_package="122.51.31.227/go-course/go18/skills/rpc/protobuf/hello_service";

// The HelloService service definition.
service HelloService {
    // rpc 声明接口
    rpc Hello (Request) returns (Response);
}

message Request {
    string value = 1;
}

message Response {
    string value = 1;
}
```

## 使用protobuf 定义数据结构

使用protoc-gen-go 插件来生成struct
```proto
message Request {
    string value = 1;
}

message Response {
    string value = 1;
}
```

```sh
# -- go_out 是参数, go out: go使用go语言插件(protoc-gen-go), go语言插件的参数 out
# out: 生成的代码放哪个目录, 默认项目根目录
# opt: module="xxx", 用于指定文件所属go module
# 应用使用了module="xxx"最好在工程根目录执行
# -I: 指定protobuf include的其他包的搜索位置, 通常情况 项目根目录比较合适
# cd go18
protoc -I=. --go_out=. --go_opt=module="122.51.31.227/go-course/go18" skills/rpc/protobuf/hello_service/interface.proto 
```


interface.pb.go
```go
type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
```

## 使用protobuf 定义RPC (grpc)

使用 protoc-gen-go-grpc来生成 接口描述(service定义相关部分)
```proto
// The HelloService service definition.
service HelloService {
    // rpc 声明接口
    rpc Hello (Request) returns (Response);
}
```

```sh
# -- go_out 是参数, go out: go使用go语言插件(protoc-gen-go), go语言插件的参数 out
# out: 生成的代码放哪个目录, 默认项目根目录
# opt: module="xxx", 用于指定文件所属go module
# 应用使用了module="xxx"最好在工程根目录执行
# -I: 指定protobuf include的其他包的搜索位置, 通常情况 项目根目录比较合适
# cd go18
protoc -I=. --go-grpc_out=. --go-grpc_opt=module="122.51.31.227/go-course/go18" skills/rpc/protobuf/hello_service/interface.proto 
```


interface_grpc.pb.go
```go
// 服务端接口
// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility.
//
// The HelloService service definition.
type HelloServiceServer interface {
	// rpc 声明接口
	Hello(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// 客户端接口
// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The HelloService service definition.
type HelloServiceClient interface {
	// rpc 声明接口
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

// 客户端的具体实现
type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, HelloService_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
```

## 服务端实现接口，注册给grpc服务



## 其他服务 使用生成的client，调用grpc服务的方法