# gRPC Streaming 示例

本示例演示 gRPC 的四种 RPC 调用模式：
1. 简单 RPC (Simple RPC)
2. 服务端流式 RPC (Server Streaming RPC)
3. 客户端流式 RPC (Client Streaming RPC)
4. 双向流式 RPC (Bidirectional Streaming RPC)

## 目录结构

```
streaming/
├── server.go    # 服务端实现（包含所有四种模式）
├── client.go    # 客户端示例（演示所有四种模式）
└── README.md    # 本文件
```

## Proto 定义

流式 RPC 在 `../proto/hello.proto` 中定义：

```protobuf
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
```

## 四种 RPC 模式详解

### 1. 简单 RPC (Simple RPC)

**模式**: 客户端 → 一个请求 → 服务端 → 一个响应 → 客户端

```
Client                    Server
  |                         |
  |-------- Request ------->|
  |                         |
  |<------- Response -------|
  |                         |
```

**使用场景**:
- 查询单个用户信息
- 执行简单的计算
- 获取配置信息

### 2. 服务端流式 RPC (Server Streaming RPC)

**模式**: 客户端 → 一个请求 → 服务端 → 多个响应（流）→ 客户端

```
Client                    Server
  |                         |
  |-------- Request ------->|
  |                         |
  |<------ Response 1 ------|
  |<------ Response 2 ------|
  |<------ Response 3 ------|
  |<------- ... ------------|
  |<------ Response N ------|
  |                         |
```

**使用场景**:
- 下载大文件（分块传输）
- 实时推送通知
- 订阅事件流
- 数据库结果集分页

**特点**:
- 客户端发送一次请求
- 服务端返回多个响应
- 使用 `stream.Recv()` 接收，`io.EOF` 表示流结束

### 3. 客户端流式 RPC (Client Streaming RPC)

**模式**: 客户端 → 多个请求（流）→ 服务端 → 一个响应 → 客户端

```
Client                    Server
  |                         |
  |------- Request 1 ------>|
  |------- Request 2 ------>|
  |------- Request 3 ------>|
  |-------- ... ----------->|
  |------- Request N ------>|
  |                         |
  |<------- Response -------|
  |                         |
```

**使用场景**:
- 上传大文件（分块传输）
- 批量数据导入
- 日志聚合
- 遥测数据收集

**特点**:
- 客户端发送多个请求
- 服务端返回一个汇总响应
- 使用 `stream.Send()` 发送，`stream.CloseAndRecv()` 接收最终响应

### 4. 双向流式 RPC (Bidirectional Streaming RPC)

**模式**: 客户端 ↔ 多个请求/响应（流）↔ 服务端

```
Client                    Server
  |                         |
  |------- Request 1 ------>|
  |<------ Response 1 ------|
  |------- Request 2 ------>|
  |<------ Response 2 ------|
  |------- Request 3 ------>|
  |<------ Response 3 ------|
  |-------- ... ----------->|
  |<------- ... ------------|
  |                         |
```

**使用场景**:
- 实时聊天
- 协作编辑
- 游戏状态同步
- 视频会议

**特点**:
- 双方都可以发送多个消息
- 读写操作独立
- 需要使用 goroutine 并发处理

## 运行示例

### 前置条件

确保已生成 proto 代码：

```bash
cd ../proto
make gen
```

### 启动服务器

```bash
cd streaming
go run server.go
```

输出：
```
Server listening at [::]:50051
支持以下 RPC 方法:
  1. SayHello - 简单 RPC
  2. SayHelloStream - 服务端流式 RPC
  3. CollectHello - 客户端流式 RPC
  4. ChatHello - 双向流式 RPC
```

### 运行客户端

在另一个终端：

```bash
cd streaming
go run client.go
```

## 示例输出

### 服务端流式 RPC 输出

```
--- 2. 服务端流式 RPC (Server Streaming RPC) ---
客户端发送一个请求，服务端返回多个响应流...

  [1] Hello Bob, message 1 (timestamp: 1234567890)
  [2] Hello Bob, message 2 (timestamp: 1234567891)
  [3] Hello Bob, message 3 (timestamp: 1234567892)
  [4] Hello Bob, message 4 (timestamp: 1234567893)
  [5] Hello Bob, message 5 (timestamp: 1234567894)

服务端流式 RPC 完成，共接收 5 条消息
```

### 客户端流式 RPC 输出

```
--- 3. 客户端流式 RPC (Client Streaming RPC) ---
客户端发送多个请求流，服务端返回一个汇总响应...

  [1] 发送: Charlie
  [2] 发送: David
  [3] 发送: Eve
  [4] 发送: Frank
  [5] 发送: Grace

服务端汇总响应: Collected 5 names: [Charlie David Eve Frank Grace]
```

### 双向流式 RPC 输出

```
--- 4. 双向流式 RPC (Bidirectional Streaming RPC) ---
客户端和服务端同时发送和接收消息流...

  → 发送 [1]: Henry
  ← 接收 [1]: Echo #1: Hello Henry! (timestamp: 1234567890)
  → 发送 [2]: Iris
  ← 接收 [2]: Echo #2: Hello Iris! (timestamp: 1234567892)
  → 发送 [3]: Jack
  ← 接收 [3]: Echo #3: Hello Jack! (timestamp: 1234567894)

双向流式 RPC 完成
```

## 关键代码解析

### 服务端流式实现

```go
func (s *StreamingServer) SayHelloStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServer) error {
    for i := 0; i < 5; i++ {
        reply := &pb.HelloReply{
            Message: fmt.Sprintf("Hello %s, message %d", req.GetName(), i+1),
        }
        // 发送响应到流
        if err := stream.Send(reply); err != nil {
            return err
        }
    }
    return nil
}
```

### 客户端流式实现

```go
func (s *StreamingServer) CollectHello(stream pb.Greeter_CollectHelloServer) error {
    var names []string
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

### 双向流式实现

```go
func (s *StreamingServer) ChatHello(stream pb.Greeter_ChatHelloServer) error {
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }

        // 立即回复
        reply := &pb.HelloReply{
            Message: fmt.Sprintf("Echo: Hello %s!", req.GetName()),
        }
        if err := stream.Send(reply); err != nil {
            return err
        }
    }
}
```

## 最佳实践

### 1. 错误处理

- 使用 `io.EOF` 判断流结束
- 区分网络错误和业务错误
- 设置合理的超时时间

### 2. 资源管理

- 及时关闭流 (`stream.CloseSend()`)
- 使用 `defer` 确保资源释放
- 设置 context 超时

### 3. 并发处理

- 双向流使用 goroutine 并发收发
- 使用 channel 同步
- 注意 goroutine 泄漏

### 4. 性能优化

- 控制消息大小
- 使用缓冲区
- 批量处理数据

## 常见问题

### Q: 什么时候使用流式 RPC？

A:
- **服务端流式**: 需要返回大量数据或实时推送
- **客户端流式**: 需要上传大量数据
- **双向流式**: 需要实时双向通信

### Q: 流式 RPC 和 WebSocket 有什么区别？

A:
- gRPC 基于 HTTP/2，支持多路复用
- gRPC 有类型安全的 protobuf
- gRPC 内置负载均衡和服务发现支持

### Q: 如何处理大文件传输？

A: 使用流式 RPC 分块传输：
- 定义合适的块大小（如 64KB）
- 使用客户端流上传
- 使用服务端流下载
- 添加进度回调

## 相关资源

- [gRPC 官方文档 - Streaming](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)
- [gRPC-Go 示例](https://github.com/grpc/grpc-go/tree/master/examples)
- [Protocol Buffers 文档](https://protobuf.dev/)
