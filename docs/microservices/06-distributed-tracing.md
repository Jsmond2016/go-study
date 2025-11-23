---
title: 分布式追踪
difficulty: advanced
duration: "4-5小时"
prerequisites: ["gRPC 基础", "微服务架构"]
tags: ["分布式追踪", "OpenTelemetry", "Jaeger", "可观测性"]
---

# 分布式追踪

分布式追踪是微服务架构中实现可观测性的重要工具，它帮助我们理解请求在多个服务之间的流转路径，定位性能瓶颈和故障。

## 📋 学习目标

完成本教程后，你将能够：

- [ ] 理解分布式追踪的概念和原理
- [ ] 使用 OpenTelemetry 进行追踪
- [ ] 集成 Jaeger 进行可视化
- [ ] 在 gRPC 服务中实现追踪
- [ ] 在 HTTP 服务中实现追踪
- [ ] 分析追踪数据并优化性能

## 🎯 分布式追踪简介

### 什么是分布式追踪

分布式追踪是一种监控技术，用于跟踪请求在分布式系统中的执行路径。它记录请求经过的每个服务、每个操作，以及它们之间的调用关系。

### 为什么需要分布式追踪

在微服务架构中，一个请求可能经过多个服务：

```
客户端 → API 网关 → 用户服务 → 订单服务 → 商品服务 → 支付服务
```

**问题**：
- 请求在哪个服务变慢了？
- 哪个服务调用失败了？
- 服务之间的调用关系是什么？
- 如何定位性能瓶颈？

**解决方案**：分布式追踪可以回答这些问题。

### 核心概念

#### 1. Trace（追踪）

一个完整的请求链路，包含多个 Span。

```
Trace: 用户下单请求
├── Span: API 网关接收请求
│   ├── Span: 认证验证
│   └── Span: 路由转发
├── Span: 订单服务处理
│   ├── Span: 调用用户服务验证
│   ├── Span: 调用商品服务检查库存
│   └── Span: 创建订单
└── Span: 返回响应
```

#### 2. Span（跨度）

一个操作单元，包含：
- **操作名称**：如 "GetUser"
- **开始时间**：操作开始时间戳
- **结束时间**：操作结束时间戳
- **标签（Tags）**：键值对，如 `http.method=GET`
- **日志（Logs）**：事件和错误信息
- **上下文（Context）**：用于关联父子 Span

#### 3. Trace Context（追踪上下文）

用于在服务间传递追踪信息，包含：
- **Trace ID**：唯一标识一个 Trace
- **Span ID**：当前 Span 的 ID
- **Parent Span ID**：父 Span 的 ID
- **Flags**：采样标志等

### 分布式追踪架构

```
服务 A → 服务 B → 服务 C
  ↓        ↓        ↓
生成 Span → 传递 Context → 生成 Span
  ↓        ↓        ↓
    收集器（Collector）
          ↓
    追踪后端（Jaeger/Zipkin）
          ↓
    可视化界面
```

## 🔧 OpenTelemetry 基础

### 什么是 OpenTelemetry

OpenTelemetry 是一个开源的可观测性框架，提供了统一的 API 和 SDK，用于收集、处理和导出遥测数据（追踪、指标、日志）。

### 核心组件

1. **API**：定义接口和抽象
2. **SDK**：实现 API，提供具体功能
3. **Instrumentation**：自动或手动插桩代码
4. **Exporters**：导出数据到后端系统

### 安装 OpenTelemetry

```bash
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/trace
go get go.opentelemetry.io/otel/exporters/jaeger
go get go.opentelemetry.io/otel/sdk/trace
go get go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
```

## 📝 基础使用

### 1. 初始化 Tracer Provider

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func initTracer(url string) (*tracesdk.TracerProvider, error) {
	// 创建 Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	// 创建 Tracer Provider
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-service"),
			semconv.ServiceVersionKey.String("1.0.0"),
		)),
	)

	// 设置为全局 Tracer Provider
	otel.SetTracerProvider(tp)

	return tp, nil
}

func main() {
	tp, err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	// 使用追踪
	tracer := otel.Tracer("example-tracer")
	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	fmt.Println("Tracing initialized")
}
```

### 2. 创建 Span

```go
func processOrder(ctx context.Context, orderID string) error {
	tracer := otel.Tracer("order-service")
	ctx, span := tracer.Start(ctx, "processOrder")
	defer span.End()

	// 添加标签
	span.SetAttributes(
		attribute.String("order.id", orderID),
		attribute.String("order.status", "processing"),
	)

	// 执行业务逻辑
	if err := validateOrder(ctx, orderID); err != nil {
		span.RecordError(err)
		return err
	}

	span.SetAttributes(attribute.String("order.status", "completed"))
	return nil
}
```

### 3. 传递 Context

```go
// 服务 A
func serviceA(ctx context.Context) {
	tracer := otel.Tracer("service-a")
	ctx, span := tracer.Start(ctx, "serviceA")
	defer span.End()

	// 调用服务 B，传递 context
	serviceB(ctx)
}

// 服务 B
func serviceB(ctx context.Context) {
	tracer := otel.Tracer("service-b")
	ctx, span := tracer.Start(ctx, "serviceB")
	defer span.End()

	// 这个 span 会自动成为 serviceA span 的子 span
}
```

## 🔌 gRPC 服务集成

### 1. 服务端集成

```go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func main() {
	// 初始化 Tracer Provider
	tp, err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	// 创建 gRPC 服务器，添加追踪拦截器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	// 注册服务
	// pb.RegisterYourServiceServer(s, &server{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server with tracing listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

### 2. 客户端集成

```go
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func main() {
	// 初始化 Tracer Provider
	tp, err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	// 创建 gRPC 连接，添加追踪拦截器
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 使用连接调用服务
	// client := pb.NewYourServiceClient(conn)
	// ctx := context.Background()
	// resp, err := client.YourMethod(ctx, &pb.Request{})
}
```

## 🌐 HTTP 服务集成

### 1. HTTP 服务器集成

```go
package main

import (
	"context"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	// 初始化 Tracer Provider
	tp, err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	// 创建 HTTP 处理器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("http-server")
		ctx, span := tracer.Start(ctx, "handleRequest")
		defer span.End()

		span.SetAttributes(
			attribute.String("http.method", r.Method),
			attribute.String("http.path", r.URL.Path),
		)

		w.Write([]byte("Hello, World!"))
	})

	// 使用 otelhttp 包装处理器
	wrappedHandler := otelhttp.NewHandler(handler, "my-service")

	http.Handle("/", wrappedHandler)
	log.Println("HTTP server with tracing listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. HTTP 客户端集成

```go
package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	// 初始化 Tracer Provider
	tp, err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	// 创建 HTTP 客户端，使用 otelhttp 包装
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/api", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)
}
```

## 📊 Jaeger 集成

### 什么是 Jaeger

Jaeger 是 Uber 开源的分布式追踪系统，用于监控和诊断微服务架构。

### 安装 Jaeger

使用 Docker 运行 Jaeger：

```bash
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
```

访问 Jaeger UI：http://localhost:16686

### 配置 Jaeger Exporter

```go
func initTracer(serviceName string) (*tracesdk.TracerProvider, error) {
	// 创建 Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://localhost:14268/api/traces"),
	))
	if err != nil {
		return nil, err
	}

	// 创建 Tracer Provider
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
```

## 🎯 实践示例

### 示例：多服务追踪

```go
package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	// 初始化 Tracer Provider
	tp, err := initTracer("example-service")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	ctx := context.Background()
	tracer := otel.Tracer("example-tracer")

	// 创建根 Span
	ctx, rootSpan := tracer.Start(ctx, "processOrder")
	defer rootSpan.End()

	// 模拟处理订单
	processOrder(ctx, tracer, "order-123")
}

func processOrder(ctx context.Context, tracer trace.Tracer, orderID string) {
	ctx, span := tracer.Start(ctx, "processOrder")
	defer span.End()

	span.SetAttributes(
		attribute.String("order.id", orderID),
	)

	// 验证订单
	validateOrder(ctx, tracer, orderID)

	// 检查库存
	checkInventory(ctx, tracer, orderID)

	// 创建订单
	createOrder(ctx, tracer, orderID)
}

func validateOrder(ctx context.Context, tracer trace.Tracer, orderID string) {
	ctx, span := tracer.Start(ctx, "validateOrder")
	defer span.End()

	time.Sleep(10 * time.Millisecond)
	span.SetAttributes(attribute.Bool("order.valid", true))
}

func checkInventory(ctx context.Context, tracer trace.Tracer, orderID string) {
	ctx, span := tracer.Start(ctx, "checkInventory")
	defer span.End()

	time.Sleep(20 * time.Millisecond)
	span.SetAttributes(attribute.Bool("inventory.available", true))
}

func createOrder(ctx context.Context, tracer trace.Tracer, orderID string) {
	ctx, span := tracer.Start(ctx, "createOrder")
	defer span.End()

	time.Sleep(15 * time.Millisecond)
	span.SetAttributes(attribute.String("order.status", "created"))
}
```

## 🔍 追踪数据分析

### 1. 查看 Trace 列表

在 Jaeger UI 中：
- 选择服务名称
- 选择时间范围
- 点击 "Find Traces"

### 2. 分析单个 Trace

- **时间线视图**：查看每个 Span 的执行时间
- **服务依赖图**：查看服务之间的调用关系
- **Span 详情**：查看标签、日志、错误信息

### 3. 性能分析

- **慢请求**：找出执行时间最长的 Span
- **错误率**：查看哪些服务错误最多
- **调用链**：理解服务间的依赖关系

## 💡 最佳实践

### 1. 采样策略

```go
// 总是采样
tp := tracesdk.NewTracerProvider(
	tracesdk.WithSampler(tracesdk.AlwaysSample()),
	// ...
)

// 概率采样（采样 10% 的请求）
tp := tracesdk.NewTracerProvider(
	tracesdk.WithSampler(tracesdk.TraceIDRatioBased(0.1)),
	// ...
)
```

### 2. 添加有意义的标签

```go
span.SetAttributes(
	attribute.String("user.id", userID),
	attribute.String("order.id", orderID),
	attribute.Int("order.amount", amount),
)
```

### 3. 记录错误

```go
if err != nil {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
	return err
}
```

### 4. 避免过度追踪

- 不要追踪每个函数调用
- 只追踪重要的业务操作
- 使用采样减少开销

## 🚀 总结

分布式追踪是微服务架构中实现可观测性的关键工具。通过 OpenTelemetry 和 Jaeger，我们可以：

1. **理解系统行为**：查看请求在服务间的流转
2. **定位性能问题**：找出慢请求和瓶颈
3. **诊断故障**：快速定位错误来源
4. **优化系统**：基于数据做出优化决策

## 📚 扩展阅读

- [OpenTelemetry 官方文档](https://opentelemetry.io/docs/)
- [Jaeger 官方文档](https://www.jaegertracing.io/docs/)
- [分布式追踪最佳实践](https://opentelemetry.io/docs/best-practices/)

## 💻 代码示例

完整的代码示例请参考：
- [代码示例](../../examples/microservices/06-distributed-tracing/)

示例包括：
- 基础追踪示例
- gRPC 服务追踪
- HTTP 服务追踪
- Jaeger 集成

## ⏭️ 下一步

- [配置中心](./07-config-center.md) - 学习配置管理
- [消息队列](./08-message-queue.md) - 学习异步通信

---

**🎉 恭喜！** 你已经掌握了分布式追踪的基础知识。继续学习下一个主题，构建更强大的微服务系统！

