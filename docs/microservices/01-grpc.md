---
title: gRPC 基础
difficulty: advanced
duration: "5-6小时"
prerequisites: ["Web 开发", "并发编程"]
tags: ["gRPC", "微服务", "RPC", "Protocol Buffers"]
---

# gRPC 基础

gRPC 是一个高性能、开源的通用 RPC 框架，由 Google 开发。它使用 Protocol Buffers 作为接口定义语言和消息交换格式。

## 📋 学习目标

完成本教程后，你将能够：

- [ ] 理解 gRPC 的概念和架构
- [ ] 了解 gRPC 与 REST 的区别
- [ ] 掌握 gRPC 服务端的实现
- [ ] 掌握 gRPC 客户端的实现
- [ ] 理解和使用 gRPC 拦截器
- [ ] 掌握 gRPC 错误处理
- [ ] 理解和使用 gRPC 流式传输

## 🎯 gRPC 简介

### 什么是 gRPC

gRPC（gRPC Remote Procedure Calls）是一个现代的开源高性能 RPC 框架，可以在任何环境中运行。它可以高效地连接数据中心内和跨数据中心的服务，支持负载均衡、追踪、健康检查和身份验证。

### 为什么选择 gRPC

- **高性能**: 基于 HTTP/2，支持多路复用和流式传输
- **跨语言**: 支持多种编程语言（Go、Java、Python、C++ 等）
- **类型安全**: 使用 Protocol Buffers 定义接口，编译时类型检查
- **流式支持**: 支持单向流、双向流等多种通信模式
- **内置功能**: 支持认证、负载均衡、超时、重试等

### gRPC vs REST

| 特性       | gRPC                       | REST           |
| ---------- | -------------------------- | -------------- |
| 协议       | HTTP/2                     | HTTP/1.1       |
| 数据格式   | Protocol Buffers（二进制） | JSON（文本）   |
| 性能       | 更高（二进制、多路复用）   | 较低           |
| 浏览器支持 | 有限（需要 gRPC-Web）      | 原生支持       |
| 流式传输   | 原生支持                   | 需要 WebSocket |
| 代码生成   | 自动生成客户端/服务端代码  | 手动编写       |

## 🚀 快速开始

### 安装依赖

```bash
# 安装 gRPC Go 包
go get google.golang.org/grpc

# 安装 Protocol Buffers 编译器插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 第一个 gRPC 服务

#### 1. 定义 .proto 文件

```protobuf
// hello.proto
syntax = "proto3";

package hello;

option go_package = "./proto;hello";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 2;
}
```

#### 2. 生成 Go 代码

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       hello.proto
```

#### 3. 实现服务端

```go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "your-project/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

#### 4. 实现客户端

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "your-project/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
```

## 📚 核心概念

### 服务定义

gRPC 使用 Protocol Buffers 定义服务接口。服务由多个 RPC 方法组成，每个方法有请求和响应类型。

```protobuf
service UserService {
  rpc GetUser (GetUserRequest) returns (User);
  rpc ListUsers (ListUsersRequest) returns (stream User);
  rpc CreateUser (CreateUserRequest) returns (User);
}
```

### 四种 RPC 类型

1. **Unary RPC（一元 RPC）**: 客户端发送一个请求，服务端返回一个响应
2. **Server Streaming（服务端流）**: 客户端发送一个请求，服务端返回一个流
3. **Client Streaming（客户端流）**: 客户端发送一个流，服务端返回一个响应
4. **Bidirectional Streaming（双向流）**: 客户端和服务端都发送流

## 🔧 服务端实现

### 基本服务端

```go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "your-project/proto"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// 实现获取用户逻辑
	return &pb.User{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userServer{})

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

### 服务端选项配置

```go
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	// 配置服务端选项
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     15 * time.Minute,
			MaxConnectionAge:      30 * time.Minute,
			MaxConnectionAgeGrace: 5 * time.Minute,
			Time:                  5 * time.Second,
			Timeout:               1 * time.Second,
		}),
		grpc.MaxRecvMsgSize(4 * 1024 * 1024), // 4MB
		grpc.MaxSendMsgSize(4 * 1024 * 1024), // 4MB
	}

	s := grpc.NewServer(opts...)
	// ...
}
```

## 💻 客户端实现

### 基本客户端

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "your-project/proto"
)

func main() {
	// 建立连接
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewUserServiceClient(conn)

	// 调用服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "123"})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User: %+v", user)
}
```

### 客户端选项配置

```go
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

conn, err := grpc.Dial("localhost:50051",
	grpc.WithTransportCredentials(insecure.NewCredentials()),
	grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	}),
	grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(4*1024*1024),
		grpc.MaxCallSendMsgSize(4*1024*1024),
	),
)
```

## 🛡️ 拦截器（Interceptor）

拦截器允许你在 RPC 调用前后执行自定义逻辑，类似于 HTTP 中间件。

### 服务端拦截器

```go
import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// 日志拦截器
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	log.Printf("Method: %s, Request: %+v", info.FullMethod, req)

	resp, err := handler(ctx, req)

	duration := time.Since(start)
	if err != nil {
		log.Printf("Method: %s, Error: %v, Duration: %v", info.FullMethod, err, duration)
	} else {
		log.Printf("Method: %s, Response: %+v, Duration: %v", info.FullMethod, resp, duration)
	}

	return resp, err
}

// 认证拦截器
func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 从 context 中获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// 验证 token
	tokens := md.Get("authorization")
	if len(tokens) == 0 || tokens[0] != "valid-token" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return handler(ctx, req)
}

func main() {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc.ChainUnaryInterceptor(
			loggingInterceptor,
			authInterceptor,
		)),
	}

	s := grpc.NewServer(opts...)
	// ...
}
```

### 客户端拦截器

```go
import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 客户端日志拦截器
func clientLoggingInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("Calling method: %s, Request: %+v", method, req)

	err := invoker(ctx, method, req, reply, cc, opts...)

	if err != nil {
		log.Printf("Method: %s, Error: %v", method, err)
	} else {
		log.Printf("Method: %s, Response: %+v", method, reply)
	}

	return err
}

// 客户端认证拦截器
func clientAuthInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 添加认证信息到 metadata
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "valid-token")
	return invoker(ctx, method, req, reply, cc, opts...)
}

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc.ChainUnaryInterceptor(
			clientLoggingInterceptor,
			clientAuthInterceptor,
		)),
	)
	// ...
}
```

## ⚠️ 错误处理

gRPC 使用标准的错误码和错误消息。

### 错误码

```go
import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user id is required")
	}

	user, err := s.findUser(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return user, nil
}
```

### 客户端错误处理

```go
user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "123"})
if err != nil {
	st, ok := status.FromError(err)
	if ok {
		switch st.Code() {
		case codes.NotFound:
			log.Printf("User not found: %s", st.Message())
		case codes.InvalidArgument:
			log.Printf("Invalid argument: %s", st.Message())
		default:
			log.Printf("Error: %s", st.Message())
		}
	} else {
		log.Printf("Unknown error: %v", err)
	}
	return
}
```

## 🌊 流式传输

### 服务端流（Server Streaming）

```protobuf
service UserService {
  rpc ListUsers (ListUsersRequest) returns (stream User);
}
```

服务端实现：

```go
func (s *userServer) ListUsers(req *pb.ListUsersRequest, stream pb.UserService_ListUsersServer) error {
	users := []*pb.User{
		{Id: "1", Name: "Alice"},
		{Id: "2", Name: "Bob"},
		{Id: "3", Name: "Charlie"},
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
```

客户端实现：

```go
stream, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
if err != nil {
	log.Fatalf("could not list users: %v", err)
}

for {
	user, err := stream.Recv()
	if err == io.EOF {
		break
	}
	if err != nil {
		log.Fatalf("error receiving user: %v", err)
	}
	log.Printf("User: %+v", user)
}
```

### 客户端流（Client Streaming）

```protobuf
service UserService {
  rpc CreateUsers (stream CreateUserRequest) returns (CreateUsersResponse);
}
```

服务端实现：

```go
func (s *userServer) CreateUsers(stream pb.UserService_CreateUsersServer) error {
	var users []*pb.User

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CreateUsersResponse{
				Count: int32(len(users)),
				Users: users,
			})
		}
		if err != nil {
			return err
		}

		user := &pb.User{
			Id:   generateID(),
			Name: req.Name,
		}
		users = append(users, user)
	}
}
```

客户端实现：

```go
stream, err := client.CreateUsers(ctx)
if err != nil {
	log.Fatalf("could not create users: %v", err)
}

names := []string{"Alice", "Bob", "Charlie"}
for _, name := range names {
	if err := stream.Send(&pb.CreateUserRequest{Name: name}); err != nil {
		log.Fatalf("error sending request: %v", err)
	}
}

resp, err := stream.CloseAndRecv()
if err != nil {
	log.Fatalf("error receiving response: %v", err)
}
log.Printf("Created %d users", resp.Count)
```

### 双向流（Bidirectional Streaming）

```protobuf
service ChatService {
  rpc Chat (stream ChatMessage) returns (stream ChatMessage);
}
```

服务端实现：

```go
func (s *chatServer) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// 处理消息并回复
		reply := &pb.ChatMessage{
			User:    "Server",
			Message: "Echo: " + msg.Message,
		}

		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}
```

客户端实现：

```go
stream, err := client.Chat(ctx)
if err != nil {
	log.Fatalf("could not chat: %v", err)
}

// 发送消息的 goroutine
go func() {
	messages := []string{"Hello", "How are you?", "Goodbye"}
	for _, msg := range messages {
		if err := stream.Send(&pb.ChatMessage{
			User:    "Client",
			Message: msg,
		}); err != nil {
			log.Fatalf("error sending: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
}()

// 接收消息的 goroutine
for {
	msg, err := stream.Recv()
	if err == io.EOF {
		break
	}
	if err != nil {
		log.Fatalf("error receiving: %v", err)
	}
	log.Printf("Received: %s", msg.Message)
}
```

## 🔐 TLS/SSL 安全连接

### 生成证书

```bash
# 生成私钥
openssl genrsa -out server.key 2048

# 生成证书签名请求
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365
```

### 服务端 TLS

```go
import (
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	// ...
}
```

### 客户端 TLS

```go
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

creds, err := credentials.NewClientTLSFromFile("server.crt", "")
if err != nil {
	log.Fatalf("failed to load credentials: %v", err)
}

conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
```

## 💡 最佳实践

### 1. 使用 Context 管理超时和取消

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resp, err := client.SomeMethod(ctx, req)
```

### 2. 合理设置消息大小限制

```go
s := grpc.NewServer(
	grpc.MaxRecvMsgSize(4*1024*1024), // 4MB
	grpc.MaxSendMsgSize(4*1024*1024),
)
```

### 3. 使用拦截器实现横切关注点

- 日志记录
- 认证授权
- 限流
- 监控和追踪

### 4. 错误处理要明确

```go
if err != nil {
	st, _ := status.FromError(err)
	log.Printf("Error code: %s, Message: %s", st.Code(), st.Message())
}
```

### 5. 流式传输要注意资源管理

```go
defer stream.CloseSend()
// 或
defer stream.CloseRecv()
```

## 📝 实践练习

1. **基础练习**：创建一个计算器服务，支持加减乘除运算
2. **流式练习**：实现一个文件上传服务，使用客户端流传输文件块
3. **拦截器练习**：实现一个限流拦截器，限制每秒请求数
4. **错误处理练习**：实现完善的错误处理和重试机制
5. **综合练习**：创建一个聊天服务，使用双向流实现实时通信

## 🔗 相关资源

- [gRPC 官方文档](https://grpc.io/)
- [gRPC Go 文档](https://pkg.go.dev/google.golang.org/grpc)
- [Protocol Buffers 文档](https://protobuf.dev/)
- [代码示例](../../examples/microservices/01-grpc/)

## ⏭️ 下一步

完成 gRPC 基础学习后，可以继续学习：

- [Protocol Buffers](./02-protobuf.md) - 深入学习 Protocol Buffers
- [服务发现](./03-service-discovery.md) - 实现服务注册和发现
- [负载均衡](./04-load-balancing.md) - 实现负载均衡

---

**🎉 恭喜！** 你已经掌握了 gRPC 的基础知识。继续实践，构建更复杂的微服务应用！

