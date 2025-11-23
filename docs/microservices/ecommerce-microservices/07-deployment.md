---
title: 部署和运维
difficulty: advanced
duration: "2-3小时"
prerequisites: ["API 网关实现"]
tags: ["部署", "运维", "Docker", "监控"]
---

# 部署和运维

本章节将详细介绍如何部署电商微服务系统，包括 Docker 部署、监控配置和运维最佳实践。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 使用 Docker Compose 部署微服务
- [ ] 配置服务监控和日志
- [ ] 实现服务健康检查
- [ ] 掌握微服务运维最佳实践
- [ ] 处理服务故障和恢复

## 🐳 Docker 部署

### 1. 创建 Dockerfile

#### 用户服务 Dockerfile

```dockerfile
# user-service/Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建
RUN CGO_ENABLED=1 GOOS=linux go build -o user-service main.go

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite

WORKDIR /root/

COPY --from=builder /app/user-service .
COPY --from=builder /app/proto ./proto

EXPOSE 5001

CMD ["./user-service"]
```

### 2. Docker Compose 配置

```yaml
version: '3.8'

services:
  consul:
    image: consul:latest
    container_name: consul
    ports:
      - "8500:8500"
      - "8600:8600/udp"
    command: consul agent -dev -client=0.0.0.0 -ui
    networks:
      - microservices

  user-service:
    build:
      context: .
      dockerfile: user-service/Dockerfile
    container_name: user-service
    ports:
      - "5001:5001"
    environment:
      - CONSUL_ADDR=consul:8500
    depends_on:
      - consul
    networks:
      - microservices
    volumes:
      - user-data:/data

  order-service:
    build:
      context: .
      dockerfile: order-service/Dockerfile
    container_name: order-service
    ports:
      - "5002:5002"
    environment:
      - CONSUL_ADDR=consul:8500
    depends_on:
      - consul
      - user-service
    networks:
      - microservices
    volumes:
      - order-data:/data

  product-service:
    build:
      context: .
      dockerfile: product-service/Dockerfile
    container_name: product-service
    ports:
      - "5003:5003"
    environment:
      - CONSUL_ADDR=consul:8500
    depends_on:
      - consul
    networks:
      - microservices
    volumes:
      - product-data:/data

  gateway:
    build:
      context: .
      dockerfile: gateway/Dockerfile
    container_name: gateway
    ports:
      - "8080:8080"
    environment:
      - CONSUL_ADDR=consul:8500
    depends_on:
      - consul
      - user-service
      - order-service
      - product-service
    networks:
      - microservices

volumes:
  user-data:
  order-data:
  product-data:

networks:
  microservices:
    driver: bridge
```

### 3. 部署命令

```bash
# 构建并启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down

# 停止并删除数据卷
docker-compose down -v
```

## 📊 监控和日志

### 1. 健康检查

#### Consul 健康检查

各服务已配置 Consul 健康检查：

```go
Check: &consulapi.AgentServiceCheck{
    TCP:      fmt.Sprintf("localhost:%d", port),
    Interval: "10s",
    Timeout:  "5s",
}
```

#### 网关健康检查端点

```go
router.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "healthy"})
})
```

### 2. 日志配置

#### 结构化日志

```go
import "go.uber.org/zap"

logger, _ := zap.NewProduction()
defer logger.Sync()

logger.Info("服务启动",
    zap.String("service", "user-service"),
    zap.Int("port", port),
)
```

#### 日志收集

使用 ELK 或 Loki 收集日志：

```yaml
# docker-compose.yml 添加
  loki:
    image: grafana/loki:latest
    ports:
      - "3100:3100"
    networks:
      - microservices
```

### 3. 指标监控

#### Prometheus 集成

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    requestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
)
```

## 🔧 运维最佳实践

### 1. 服务配置

#### 环境变量配置

```bash
# .env 文件
CONSUL_ADDR=consul:8500
USER_SERVICE_PORT=5001
ORDER_SERVICE_PORT=5002
PRODUCT_SERVICE_PORT=5003
GATEWAY_PORT=8080
```

#### 配置文件

```yaml
# config.yaml
services:
  user:
    port: 5001
    database: user_service.db
  order:
    port: 5002
    database: order_service.db
  product:
    port: 5003
    database: product_service.db
gateway:
  port: 8080
consul:
  address: localhost:8500
```

### 2. 数据库管理

#### 数据备份

```bash
# 备份 SQLite 数据库
sqlite3 user_service.db ".backup user_service_backup.db"
```

#### 数据迁移

使用数据库迁移工具（如 migrate）：

```bash
migrate -path ./migrations -database "sqlite3://user_service.db" up
```

### 3. 服务扩展

#### 水平扩展

```yaml
# docker-compose.yml
  user-service:
    deploy:
      replicas: 3
```

#### 负载均衡

Consul 自动提供负载均衡：

```go
services, _, err := client.Health().Service("user-service", "", true, nil)
// 随机选择一个健康的服务实例
service := services[rand.Intn(len(services))].Service
```

### 4. 故障处理

#### 服务降级

```go
func (g *Gateway) GetUser(c *gin.Context) {
    client, err := g.getServiceClient("user-service")
    if err != nil {
        // 服务降级：返回缓存数据或默认值
        c.JSON(http.StatusOK, gin.H{
            "user_id": 0,
            "message": "服务暂时不可用",
        })
        return
    }
    // ...
}
```

#### 重试机制

```go
func retryCall(fn func() error, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err == nil {
            return nil
        }
        time.Sleep(time.Duration(i+1) * time.Second)
    }
    return fmt.Errorf("重试失败")
}
```

#### 熔断器

使用之前实现的熔断器：

```go
router.Use(circuit.CircuitBreakerMiddleware(5, 30*time.Second))
```

## 🚀 生产环境建议

### 1. 安全

- ✅ 使用 HTTPS/TLS
- ✅ 实现 JWT 认证
- ✅ 服务间 mTLS
- ✅ 密码加密存储
- ✅ 输入验证和清理

### 2. 性能

- ✅ 连接池管理
- ✅ 缓存策略（Redis）
- ✅ 数据库索引优化
- ✅ 异步处理

### 3. 可靠性

- ✅ 服务健康检查
- ✅ 自动故障恢复
- ✅ 数据备份和恢复
- ✅ 监控和告警

### 4. 可观测性

- ✅ 分布式追踪（Jaeger）
- ✅ 指标监控（Prometheus）
- ✅ 日志聚合（ELK/Loki）
- ✅ 性能分析

## 📚 总结

完成本教程后，你已经：

- ✅ 构建了完整的微服务系统
- ✅ 实现了服务发现和注册
- ✅ 实现了服务间通信
- ✅ 实现了 API 网关
- ✅ 掌握了部署和运维

## 🎯 下一步

可以继续学习：

1. **分布式事务** - Saga 模式
2. **服务网格** - Istio/Linkerd
3. **容器编排** - Kubernetes
4. **消息队列** - RabbitMQ/Kafka
5. **缓存系统** - Redis

---

**🎉 恭喜完成电商微服务实战项目！** 现在你已经掌握了微服务架构的核心技能。

