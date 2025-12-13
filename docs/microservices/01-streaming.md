---
title: gRPC 流式通信详解
difficulty: intermediate
duration: "4-6小时"
prerequisites: ["gRPC 基础", "Protocol Buffers"]
tags: ["gRPC", "流式RPC", "服务间通信"]
---

# gRPC 流式通信详解

本章深入讲解 gRPC 的四种 RPC 调用模式，重点介绍流式通信的实现和应用场景。

## 📋 学习目标

完成本教程后，你将能够：

- [x] 理解 gRPC 的四种 RPC 调用模式
- [x] 实现服务端流式 RPC
- [x] 实现客户端流式 RPC
- [x] 实现双向流式 RPC
- [x] 处理流式通信中的错误和超时
- [x] 编写流式 RPC 的单元测试

## 🎯 四种 RPC 模式对比

### 模式概览

| 模式 | 客户端 | 服务端 | 使用场景 |
|------|--------|--------|----------|
| **简单 RPC** | 一个请求 | 一个响应 | 查询用户信息、执行计算 |
| **服务端流式** | 一个请求 | 多个响应 | 下载文件、推送通知、数据流 |
| **客户端流式** | 多个请求 | 一个响应 | 上传文件、批量导入、日志收集 |
| **双向流式** | 多个请求 | 多个响应 | 实时聊天、协作编辑、游戏状态同步 |

### 通信模式图解

#### 1. 简单 RPC

```
Client                    Server
  |                         |
  |-------- Request ------->|
  |                         |
  |<------- Response -------|
  |                         |
```

#### 2. 服务端流式 RPC

```
Client                    Server
  |                         |
  |-------- Request ------->|
  |                         |
  |<------ Response 1 ------|
  |<------ Response 2 ------|
  |<------ Response N ------|
  |                         |
```

#### 3. 客户端流式 RPC

```
Client                    Server
  |                         |
  |------- Request 1 ------>|
  |------- Request 2 ------>|
  |------- Request N ------>|
  |                         |
  |<------- Response -------|
  |                         |
```

#### 4. 双向流式 RPC

```
Client                    Server
  |                         |
  |------- Request 1 ------>|
  |<------ Response 1 ------|
  |------- Request 2 ------>|
  |<------ Response 2 ------|
  |<------ Response 3 ------|
  |------- Request 3 ------>|
  |                         |
```

## 🔧 Protocol Buffers 定义

首先定义支持所有四种模式的 proto 文件：

```protobuf
syntax = "proto3";

package hello;

option go_package = "./proto;hello";

service Greeter {
  // 简单 RPC：一个请求 -> 一个响应
  rpc SayHello (HelloRequest) returns (HelloReply);

  // 服务端流式：一个请求 -> 多个响应（流）
  rpc SayHelloStream (HelloRequest) returns (stream HelloReply);

  // 客户端流式：多个请求（流）-> 一个响应
  rpc CollectHello (stream HelloRequest) returns (HelloReply);

  // 双向流式：多个请求（流）<-> 多个响应（流）
  rpc ChatHello (stream HelloRequest) returns (stream HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 2;
}
```

## 📝 服务端实现

### 1. 简单 RPC

```go
func (s *StreamingServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Printf("[Simple RPC] Received: %v", req.GetName())
    return &pb.HelloReply{
        Message: fmt.Sprintf("Hello, %s!", req.GetName()),
    }, nil
}
```

### 2. 服务端流式 RPC

服务端可以发送多个响应：

```go
func (s *StreamingServer) SayHelloStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServer) error {
    log.Printf("[Server Streaming] Streaming for: %v", req.GetName())

    // 发送 5 条消息
    for i := 0; i < 5; i++ {
        reply := &pb.HelloReply{
            Message: fmt.Sprintf("Hello %s, message %d (timestamp: %d)",
                req.GetName(), i+1, time.Now().Unix()),
        }

        if err := stream.Send(reply); err != nil {
            return err
        }

        time.Sleep(time.Second) // 模拟处理延迟
    }

    return nil
}
```

**关键点**：
- 使用 `stream.Send()` 发送多个响应
- 返回 `nil` 表示流结束
- 可以在循环中发送任意数量的消息

### 3. 客户端流式 RPC

服务端接收多个请求，返回一个汇总响应：

```go
func (s *StreamingServer) CollectHello(stream pb.Greeter_CollectHelloServer) error {
    var names []string

    // 接收客户端发送的所有消息
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            // 客户端发送完毕，返回汇总响应
            reply := &pb.HelloReply{
                Message: fmt.Sprintf("Collected %d names: %v", len(names), names),
            }
            return stream.SendAndClose(reply)
        }
        if err != nil {
            return err
        }

        names = append(names, req.GetName())
    }
}
```

**关键点**：
- 使用 `stream.Recv()` 接收消息
- `io.EOF` 表示客户端发送完毕
- 使用 `stream.SendAndClose()` 返回最终响应

### 4. 双向流式 RPC

客户端和服务端可以同时收发消息：

```go
func (s *StreamingServer) ChatHello(stream pb.Greeter_ChatHelloServer) error {
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil // 客户端关闭连接
        }
        if err != nil {
            return err
        }

        // 收到消息后立即回复
        reply := &pb.HelloReply{
            Message: fmt.Sprintf("Echo: Hello %s! (timestamp: %d)",
                req.GetName(), time.Now().Unix()),
        }

        if err := stream.Send(reply); err != nil {
            return err
        }
    }
}
```

**关键点**：
- 同时使用 `stream.Recv()` 和 `stream.Send()`
- 读写操作独立，可以自由组合
- 通常需要 goroutine 并发处理

## 💻 客户端实现

### 1. 简单 RPC 调用

```go
func simpleRPC(client pb.GreeterClient) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
    if err != nil {
        log.Fatalf("RPC failed: %v", err)
    }

    fmt.Printf("Response: %s\n", resp.GetMessage())
}
```

### 2. 服务端流式调用

客户端接收多个响应：

```go
func serverStreaming(client pb.GreeterClient) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Bob"})
    if err != nil {
        log.Fatalf("Failed to call: %v", err)
    }

    // 接收多个响应
    count := 0
    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break // 流结束
        }
        if err != nil {
            log.Fatalf("Error: %v", err)
        }

        count++
        fmt.Printf("[%d] %s\n", count, resp.GetMessage())
    }
}
```

### 3. 客户端流式调用

客户端发送多个请求：

```go
func clientStreaming(client pb.GreeterClient) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    stream, err := client.CollectHello(ctx)
    if err != nil {
        log.Fatalf("Failed to call: %v", err)
    }

    // 发送多个请求
    names := []string{"Charlie", "David", "Eve"}
    for _, name := range names {
        if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
            log.Fatalf("Send error: %v", err)
        }
        fmt.Printf("Sent: %s\n", name)
    }

    // 关闭发送并接收响应
    resp, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("CloseAndRecv error: %v", err)
    }

    fmt.Printf("Server response: %s\n", resp.GetMessage())
}
```

### 4. 双向流式调用

需要使用 goroutine 并发收发：

```go
func bidirectionalStreaming(client pb.GreeterClient) {
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    stream, err := client.ChatHello(ctx)
    if err != nil {
        log.Fatalf("Failed to call: %v", err)
    }

    waitc := make(chan struct{})

    // 启动接收 goroutine
    go func() {
        for {
            resp, err := stream.Recv()
            if err == io.EOF {
                close(waitc)
                return
            }
            if err != nil {
                log.Fatalf("Receive error: %v", err)
            }
            fmt.Printf("← Received: %s\n", resp.GetMessage())
        }
    }()

    // 发送消息
    names := []string{"Henry", "Iris", "Jack"}
    for _, name := range names {
        if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
            log.Fatalf("Send error: %v", err)
        }
        fmt.Printf("→ Sent: %s\n", name)
        time.Sleep(2 * time.Second)
    }

    stream.CloseSend()
    <-waitc // 等待接收完成
}
```

## 🧪 单元测试

### In-Memory gRPC Server

使用 `bufconn` 创建内存中的 gRPC 服务器进行测试：

```go
const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(bufSize)
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &StreamingServer{})
    go func() {
        s.Serve(lis)
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}
```

### 测试服务端流式

```go
func TestServerStreaming(t *testing.T) {
    ctx := context.Background()
    conn, _ := grpc.DialContext(ctx, "bufnet",
        grpc.WithContextDialer(bufDialer),
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    defer conn.Close()

    client := pb.NewGreeterClient(conn)
    stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Test"})
    assert.NoError(t, err)

    count := 0
    for {
        _, err := stream.Recv()
        if err == io.EOF {
            break
        }
        assert.NoError(t, err)
        count++
    }

    assert.Equal(t, 5, count)
}
```

### 测试客户端流式

```go
func TestClientStreaming(t *testing.T) {
    // ... setup ...

    stream, err := client.CollectHello(ctx)
    assert.NoError(t, err)

    names := []string{"Alice", "Bob", "Charlie"}
    for _, name := range names {
        err := stream.Send(&pb.HelloRequest{Name: name})
        assert.NoError(t, err)
    }

    resp, err := stream.CloseAndRecv()
    assert.NoError(t, err)
    assert.Contains(t, resp.GetMessage(), "3")
}
```

## 🎨 应用场景

### 服务端流式 RPC

✅ **适用场景**：
- 下载大文件（分块传输）
- 实时数据推送（股票行情、日志流）
- 长时间运行的任务（进度更新）
- 数据库查询结果集（分页）

❌ **不适用**：
- 一次性小数据传输
- 需要事务保证的操作

**示例**：文件下载服务

```go
func (s *FileServer) DownloadFile(req *pb.DownloadRequest, stream pb.FileService_DownloadFileServer) error {
    file, err := os.Open(req.Filename)
    if err != nil {
        return err
    }
    defer file.Close()

    buffer := make([]byte, 64*1024) // 64KB chunks
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        if err := stream.Send(&pb.FileChunk{Data: buffer[:n]}); err != nil {
            return err
        }
    }

    return nil
}
```

### 客户端流式 RPC

✅ **适用场景**：
- 上传大文件
- 批量数据导入
- 日志聚合
- 遥测数据收集

**示例**：日志收集服务

```go
func (s *LogServer) CollectLogs(stream pb.LogService_CollectLogsServer) error {
    var logs []*pb.LogEntry

    for {
        entry, err := stream.Recv()
        if err == io.EOF {
            // 所有日志收集完毕，写入数据库
            count := s.db.SaveLogs(logs)
            return stream.SendAndClose(&pb.LogResponse{
                Count: count,
                Status: "success",
            })
        }
        if err != nil {
            return err
        }

        logs = append(logs, entry)
    }
}
```

### 双向流式 RPC

✅ **适用场景**：
- 实时聊天应用
- 协作编辑（如 Google Docs）
- 多人游戏状态同步
- 视频/音频会议

**示例**：聊天服务

```go
func (s *ChatServer) Chat(stream pb.ChatService_ChatServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }

        // 广播消息给所有用户
        s.broadcast(msg)

        // 确认收到
        if err := stream.Send(&pb.ChatResponse{
            Status: "delivered",
            Timestamp: time.Now().Unix(),
        }); err != nil {
            return err
        }
    }
}
```

## 💡 最佳实践

### 1. 错误处理

```go
func (s *Server) StreamData(req *pb.Request, stream pb.Service_StreamDataServer) error {
    for i := 0; i < 10; i++ {
        data, err := s.fetchData(i)
        if err != nil {
            // 返回 gRPC 错误
            return status.Errorf(codes.Internal, "fetch failed: %v", err)
        }

        if err := stream.Send(data); err != nil {
            log.Printf("Send error: %v", err)
            return err
        }
    }
    return nil
}
```

### 2. 超时控制

```go
// 服务端设置超时
ctx, cancel := context.WithTimeout(stream.Context(), 30*time.Second)
defer cancel()

// 客户端设置超时
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

stream, err := client.StreamData(ctx, req)
```

### 3. 背压处理

```go
// 控制发送速率
rateLimiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10)

for data := range dataChannel {
    if err := rateLimiter.Wait(ctx); err != nil {
        return err
    }
    stream.Send(data)
}
```

### 4. 资源管理

```go
// 使用 defer 确保资源释放
defer stream.CloseSend()

// 设置合理的缓冲区大小
opts := []grpc.ServerOption{
    grpc.MaxConcurrentStreams(100),
    grpc.MaxRecvMsgSize(10 * 1024 * 1024), // 10MB
}
```

## 🚀 运行示例

完整的代码示例在 [examples/microservices/01-grpc/streaming/](https://github.com/Jsmond2016/go-study/tree/main/examples/microservices/01-grpc/streaming)

```bash
# 启动服务器
cd examples/microservices/01-grpc/streaming
go run server.go

# 运行客户端（另一个终端）
go run client.go

# 运行测试
go test -v

# 运行基准测试
go test -bench=. -benchmem
```

## 📚 扩展阅读

- [gRPC 官方文档 - Streaming](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)
- [Protocol Buffers 指南](https://protobuf.dev/programming-guides/proto3/)
- [gRPC-Go 示例](https://github.com/grpc/grpc-go/tree/master/examples)

## ⏭️ 下一步

- [02. Protocol Buffers](./02-protobuf) - 深入学习 Protocol Buffers
- [03. 服务发现](./03-service-discovery) - 服务注册与发现
- [微服务测试指南](/microservices/testing) - 学习如何测试微服务

---

**🎉 恭喜！** 你已经掌握了 gRPC 流式通信的核心概念和实践技巧！
