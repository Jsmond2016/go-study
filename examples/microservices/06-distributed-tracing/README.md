# 分布式追踪示例

本目录包含使用 OpenTelemetry 和 Jaeger 实现分布式追踪的示例代码。

## 📋 目录结构

```
06-distributed-tracing/
├── basic/              # 基础追踪示例
├── grpc/              # gRPC 服务追踪示例
│   ├── server/        # gRPC 服务器
│   └── client/        # gRPC 客户端
├── http/              # HTTP 服务追踪示例
│   ├── server/        # HTTP 服务器
│   └── client/        # HTTP 客户端
├── docker-compose.yml # Jaeger 服务配置
├── go.mod             # Go 模块定义
└── README.md          # 本文件
```

## 🚀 快速开始

### 1. 启动 Jaeger

使用 Docker Compose 启动 Jaeger：

```bash
docker-compose up -d
```

访问 Jaeger UI：http://localhost:16686

### 2. 运行基础示例

```bash
cd basic
go run main.go
```

在 Jaeger UI 中查看追踪数据。

### 3. 运行 HTTP 服务示例

**启动服务器**：

```bash
cd http/server
go run main.go
```

**运行客户端**：

```bash
cd http/client
go run main.go -url http://localhost:8080
```

### 4. 运行 gRPC 服务示例

**启动服务器**：

```bash
cd grpc/server
go run main.go
```

**运行客户端**：

```bash
cd grpc/client
go run main.go -addr localhost:50051
```

## 📚 示例说明

### 基础示例 (basic/)

演示如何使用 OpenTelemetry 创建 Span 和追踪请求链路。

**功能**：
- 初始化 Tracer Provider
- 创建根 Span 和子 Span
- 添加标签和属性
- 记录错误

### gRPC 服务示例 (grpc/)

演示如何在 gRPC 服务中集成分布式追踪。

**功能**：
- 服务端自动追踪
- 客户端自动追踪
- Context 传递
- 跨服务追踪

### HTTP 服务示例 (http/)

演示如何在 HTTP 服务中集成分布式追踪。

**功能**：
- HTTP 服务器追踪
- HTTP 客户端追踪
- 请求/响应追踪
- 中间件集成

## 🔧 配置说明

### Jaeger 配置

默认配置：
- **Jaeger UI**: http://localhost:16686
- **Collector HTTP**: http://localhost:14268/api/traces
- **Collector gRPC**: localhost:14250

### 环境变量

可以通过环境变量配置 Jaeger 地址：

```bash
export JAEGER_ENDPOINT=http://localhost:14268/api/traces
```

## 📊 查看追踪数据

1. 打开 Jaeger UI：http://localhost:16686
2. 选择服务名称（如 `basic-tracing-example`）
3. 选择时间范围
4. 点击 "Find Traces"
5. 点击单个 Trace 查看详细信息

## 🎯 学习要点

1. **Trace 和 Span**：理解追踪的基本概念
2. **Context 传递**：如何在服务间传递追踪上下文
3. **自动插桩**：使用 OpenTelemetry 自动追踪
4. **手动追踪**：如何手动创建 Span
5. **标签和属性**：添加有意义的元数据

## 📖 相关文档

- [分布式追踪文档](../../../docs/microservices/06-distributed-tracing.md)
- [OpenTelemetry 官方文档](https://opentelemetry.io/docs/)
- [Jaeger 官方文档](https://www.jaegertracing.io/docs/)

## 🐛 常见问题

### 1. 无法连接到 Jaeger

确保 Jaeger 服务已启动：

```bash
docker-compose ps
```

### 2. 看不到追踪数据

- 检查 Jaeger 地址配置是否正确
- 确保服务名称匹配
- 检查时间范围设置

### 3. 追踪数据延迟

Jaeger 使用批处理，数据可能有几秒延迟。

## 📝 下一步

- 学习配置中心：`07-config-center.md`
- 学习消息队列：`08-message-queue.md`

---

**🎉 开始使用分布式追踪，提升微服务的可观测性！**

