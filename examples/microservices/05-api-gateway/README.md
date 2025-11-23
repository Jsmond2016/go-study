# API 网关示例

本示例展示了如何实现一个功能完整的 API 网关，包括路由、认证、限流、熔断等功能。

## 📋 目录结构

```
05-api-gateway/
├── gateway/             # 网关核心实现
│   └── main.go         # 网关主程序
├── middleware/         # 中间件实现
│   ├── auth.go         # JWT 认证中间件
│   ├── ratelimit.go    # 限流中间件
│   ├── logging.go      # 日志中间件
│   └── cors.go         # CORS 中间件
├── circuit/            # 熔断器实现
│   ├── breaker.go      # 熔断器核心逻辑
│   └── middleware.go   # 熔断器中间件
├── go.mod              # Go 模块定义
└── README.md           # 本文件
```

## 🚀 快速开始

### 1. 安装依赖

```bash
go mod download
go mod tidy
```

### 2. 启动网关

```bash
go run gateway/main.go -port=8080
```

### 3. 测试网关

```bash
# 健康检查
curl http://localhost:8080/health

# 查看指标
curl http://localhost:8080/metrics

# 访问后端服务（需要先启动后端服务）
curl http://localhost:8080/api/users
```

## 📚 功能说明

### 1. 路由转发

网关支持将请求转发到不同的后端服务：

```go
gateway.RegisterRoute("/api/users", "http://localhost:5001")
gateway.RegisterRoute("/api/orders", "http://localhost:5002")
gateway.RegisterRoute("/api/products", "http://localhost:5003")
```

### 2. JWT 认证

使用 JWT token 进行身份认证：

```bash
# 生成 token（示例）
# 实际应用中应该由认证服务生成

# 使用 token 访问受保护的路由
curl -H "Authorization: Bearer <token>" http://localhost:8080/api/users
```

### 3. 限流

支持令牌桶和漏桶两种限流算法：

```go
// 令牌桶：允许突发流量
router.Use(middleware.TokenBucketRateLimit(10, 20))

// 漏桶：固定速率处理
router.Use(middleware.LeakyBucketRateLimit(10))
```

### 4. 熔断器

自动检测服务故障并熔断：

```go
router.Use(circuit.CircuitBreakerMiddleware(5, 30*time.Second))
```

### 5. 日志记录

记录所有请求的详细信息：

- 请求方法
- 请求路径
- 客户端 IP
- 响应状态码
- 处理时间

### 6. CORS 支持

支持跨域请求：

```go
router.Use(middleware.CORSMiddleware)
```

### 7. 监控指标

提供 `/metrics` 端点查看统计信息：

- 请求数量
- 错误数量
- 平均响应时间

## 🔧 使用示例

### 完整网关配置

```go
package main

import (
    "go-study/examples/microservices/05-api-gateway/middleware"
    "go-study/examples/microservices/05-api-gateway/circuit"
    "time"
)

func main() {
    gateway := NewGateway()

    // 注册路由
    gateway.RegisterRoute("/api/users", "http://localhost:5001")
    gateway.RegisterRoute("/api/orders", "http://localhost:5002")

    router := mux.NewRouter()

    // 添加中间件
    router.Use(middleware.LoggingMiddleware)
    router.Use(middleware.CORSMiddleware)
    router.Use(middleware.AuthMiddleware) // 需要认证
    router.Use(middleware.TokenBucketRateLimit(10, 20)) // 限流
    router.Use(circuit.CircuitBreakerMiddleware(5, 30*time.Second)) // 熔断

    // 健康检查（不需要认证）
    router.HandleFunc("/health", healthHandler).Methods("GET")

    // 网关路由
    router.PathPrefix("/api/").Handler(gateway)

    http.ListenAndServe(":8080", router)
}
```

### 生成 JWT Token（示例）

```go
package main

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
)

func generateToken(userID, username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  userID,
        "username": username,
        "exp":       time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString([]byte("your-secret-key"))
}
```

## 📝 测试

### 1. 测试认证中间件

```bash
# 不带 token 的请求（应该返回 401）
curl http://localhost:8080/api/users

# 带有效 token 的请求
curl -H "Authorization: Bearer <your-token>" http://localhost:8080/api/users
```

### 2. 测试限流

```bash
# 快速发送多个请求，观察限流效果
for i in {1..20}; do
  curl http://localhost:8080/api/users
  sleep 0.1
done
```

### 3. 测试熔断器

```bash
# 停止后端服务，然后发送请求
# 观察熔断器状态变化
curl http://localhost:8080/api/users
```

### 4. 测试路由转发

```bash
# 测试不同的路由
curl -H "Authorization: Bearer <token>" http://localhost:8080/api/users
curl -H "Authorization: Bearer <token>" http://localhost:8080/api/products
curl -H "Authorization: Bearer <token>" http://localhost:8080/api/orders
```

### 5. 测试健康检查

```bash
# 健康检查不需要认证
curl http://localhost:8080/health
```

## 🐛 常见问题

### 1. 认证失败

- 检查 token 格式是否正确：`Bearer <token>`
- 检查 token 是否过期
- 检查密钥是否匹配

### 2. 限流触发

- 降低请求频率
- 增加限流配置（rps 和 burst）

### 3. 熔断器打开

- 检查后端服务是否正常
- 等待重置超时时间
- 手动重置熔断器

### 4. 路由不匹配

- 检查路由路径是否正确
- 检查后端服务是否启动
- 查看日志了解详细错误

## 📖 相关文档

- [Gorilla Mux 文档](https://github.com/gorilla/mux)
- [JWT 文档](https://github.com/golang-jwt/jwt)
- [项目文档](../../../docs/microservices/05-api-gateway.md)

## ⚠️ 注意事项

1. **生产环境**：
   - 使用 HTTPS
   - 配置正确的 CORS 策略
   - 使用环境变量管理密钥
   - 配置适当的超时时间

2. **安全性**：
   - JWT 密钥应保密
   - 使用强密钥
   - 定期轮换密钥
   - 验证 token 签名

3. **性能**：
   - 使用连接池
   - 配置适当的超时
   - 监控网关性能
   - 考虑使用缓存

4. **可扩展性**：
   - 支持动态路由配置
   - 集成服务发现
   - 支持多实例部署
   - 使用负载均衡

## 🎯 下一步

完成本示例后，可以继续学习：

- [服务发现示例](../03-service-discovery/) - 集成服务发现
- [负载均衡示例](../04-load-balancing/) - 集成负载均衡
- 微服务实战项目 - 构建完整的微服务系统

