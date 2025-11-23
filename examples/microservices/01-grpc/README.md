# gRPC 基础示例

本示例展示了 gRPC 的基础使用，包括服务端、客户端、流式传输和拦截器。

## 📋 目录结构

```
01-grpc/
├── proto/              # Protocol Buffers 定义文件
│   ├── hello.proto     # Hello 服务定义
│   └── user.proto      # User 服务定义
├── server/             # 服务端代码
│   ├── main.go         # 基础服务端
│   └── with_interceptors.go  # 带拦截器的服务端
├── client/             # 客户端代码
│   ├── main.go         # 基础客户端
│   └── with_auth.go    # 带认证的客户端
├── interceptors/       # 拦截器实现
│   ├── logging.go      # 日志拦截器
│   ├── auth.go         # 认证拦截器
│   └── ratelimit.go    # 限流拦截器
├── go.mod              # Go 模块定义
├── Makefile            # 构建脚本
└── README.md           # 本文件
```

## 🚀 快速开始

### 1. 安装依赖

```bash
# 安装 Protocol Buffers 编译器
# macOS
brew install protobuf

# Linux
sudo apt-get install protobuf-compiler

# 安装 Go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 2. 生成代码

```bash
make proto
```

或者手动执行：

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       --proto_path=proto \
       proto/*.proto
```

### 3. 安装 Go 依赖

```bash
go mod download
go mod tidy
```

### 4. 运行服务端

```bash
# 基础服务端
make server
# 或
go run server/main.go

# 带拦截器的服务端
go run server/with_interceptors.go
```

### 5. 运行客户端

在另一个终端：

```bash
# 基础客户端
make client
# 或
go run client/main.go

# 带认证的客户端（需要先启动带拦截器的服务端）
go run client/with_auth.go
```

## 📚 示例说明

### 1. 基础 RPC 调用

**服务端** (`server/main.go`):
- 实现简单的 `SayHello` RPC 方法
- 监听 `:50051` 端口

**客户端** (`client/main.go`):
- 连接到服务端
- 调用 `SayHello` 方法

### 2. 流式传输

示例包含四种流式传输模式：

- **服务端流** (`SayHelloStream`): 客户端发送一个请求，服务端返回一个流
- **客户端流** (`CollectHello`): 客户端发送一个流，服务端返回一个响应
- **双向流** (`ChatHello`): 客户端和服务端都发送流

### 3. 拦截器

#### 日志拦截器 (`interceptors/logging.go`)
- 记录所有 RPC 调用的请求、响应和耗时
- 支持 Unary 和 Stream 两种模式

#### 认证拦截器 (`interceptors/auth.go`)
- 从 metadata 中获取 token
- 验证 token 有效性
- 将用户信息存储到 context

#### 限流拦截器 (`interceptors/ratelimit.go`)
- 实现令牌桶算法
- 限制每秒请求数

## 🔧 使用拦截器

### 服务端

```go
import (
    "go-study/examples/microservices/01-grpc/interceptors"
)

s := grpc.NewServer(
    grpc.ChainUnaryInterceptor(
        interceptors.LoggingInterceptor,
        interceptors.AuthInterceptor,
        interceptors.RateLimitInterceptor(rateLimiter),
    ),
)
```

### 客户端（带认证）

```go
ctx := metadata.AppendToOutgoingContext(ctx, "authorization", "valid-token-123")
reply, err := client.SayHello(ctx, req)
```

## 📝 测试

### 测试基础功能

1. 启动服务端：`go run server/main.go`
2. 运行客户端：`go run client/main.go`

### 测试拦截器

1. 启动带拦截器的服务端：`go run server/with_interceptors.go`
2. 运行带认证的客户端：`go run client/with_auth.go`

### 测试限流

连续快速发送请求，观察限流效果：

```bash
for i in {1..20}; do
    go run client/with_auth.go
    sleep 0.1
done
```

## 🐛 常见问题

### 1. protoc 命令未找到

确保已安装 Protocol Buffers 编译器：

```bash
protoc --version
```

### 2. 代码生成失败

检查 proto 文件路径和 go_package 选项是否正确。

### 3. 连接被拒绝

确保服务端已启动，并且端口号正确（默认 50051）。

### 4. 认证失败

使用带拦截器的服务端时，需要在客户端 metadata 中添加正确的 token：

```go
ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "valid-token-123")
```

## 📖 相关文档

- [gRPC 官方文档](https://grpc.io/)
- [Protocol Buffers 文档](https://protobuf.dev/)
- [项目文档](../../../docs/microservices/01-grpc.md)

## ⚠️ 注意事项

1. 本示例中的认证 token 是硬编码的，仅用于演示。生产环境应使用 JWT 等安全方案。
2. 限流器的参数可以根据实际需求调整。
3. 流式传输示例中使用了简单的 sleep，实际应用中应根据业务逻辑调整。

## 🎯 下一步

完成本示例后，可以继续学习：

- [Protocol Buffers 示例](../02-protobuf/) - 深入学习 Protocol Buffers
- [服务发现示例](../03-service-discovery/) - 实现服务注册和发现
- [负载均衡示例](../04-load-balancing/) - 实现负载均衡

