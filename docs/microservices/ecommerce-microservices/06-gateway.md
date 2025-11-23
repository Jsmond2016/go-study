---
title: API 网关实现
difficulty: advanced
duration: "4-5小时"
prerequisites: ["用户服务实现", "商品服务实现", "订单服务实现"]
tags: ["API 网关", "网关", "路由", "协议转换"]
---

# API 网关实现

本章节将详细介绍如何实现 API 网关，包括路由转发、协议转换（HTTP 到 gRPC）、服务发现集成和统一入口管理。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 理解 API 网关的作用和架构
- [ ] 实现 HTTP 路由和转发
- [ ] 实现 HTTP 到 gRPC 的协议转换
- [ ] 集成服务发现（Consul）
- [ ] 实现请求日志和监控
- [ ] 实现认证和授权（可选）

## 🎯 网关功能

API 网关提供以下功能：

- ✅ 统一 API 入口
- ✅ 路由转发到各个微服务
- ✅ HTTP 到 gRPC 协议转换
- ✅ 服务发现集成
- ✅ 请求日志记录
- ✅ CORS 支持
- ✅ 健康检查端点

## 🏗️ 网关架构

```
客户端 (HTTP)
    ↓
API 网关 (Gin)
    ↓
服务发现 (Consul)
    ↓
gRPC 客户端
    ↓
微服务 (gRPC)
```

## 💻 网关实现

### 1. 创建网关结构

```go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderpb "go-study/examples/microservices/06-ecommerce-microservices/proto/order"
	productpb "go-study/examples/microservices/06-ecommerce-microservices/proto/product"
	userpb "go-study/examples/microservices/06-ecommerce-microservices/proto/user"
)

var (
	port       = flag.Int("port", 8080, "网关端口")
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
)

func main() {
	flag.Parse()

	// 创建网关
	gateway := NewGateway(*consulAddr)

	// 创建 Gin 路由
	router := gin.Default()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API 路由
	api := router.Group("/api")
	{
		// 用户服务路由
		users := api.Group("/users")
		{
			users.POST("", gateway.CreateUser)
			users.GET("/:id", gateway.GetUser)
			users.PUT("/:id", gateway.UpdateUser)
			users.DELETE("/:id", gateway.DeleteUser)
			users.POST("/login", gateway.Login)
		}

		// 订单服务路由
		orders := api.Group("/orders")
		{
			orders.POST("", gateway.CreateOrder)
			orders.GET("/:id", gateway.GetOrder)
			orders.GET("/user/:user_id", gateway.GetUserOrders)
			orders.PUT("/:id/status", gateway.UpdateOrderStatus)
			orders.DELETE("/:id", gateway.CancelOrder)
		}

		// 商品服务路由
		products := api.Group("/products")
		{
			products.POST("", gateway.CreateProduct)
			products.GET("/:id", gateway.GetProduct)
			products.PUT("/:id", gateway.UpdateProduct)
			products.DELETE("/:id", gateway.DeleteProduct)
			products.GET("", gateway.ListProducts)
		}
	}

	// 启动服务器
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("API 网关启动在端口 %d", *port)
	log.Fatal(server.ListenAndServe())
}
```

### 2. 网关核心结构

```go
// Gateway API 网关
type Gateway struct {
	consulAddr string
	clients    map[string]*grpc.ClientConn
}

// NewGateway 创建网关
func NewGateway(consulAddr string) *Gateway {
	return &Gateway{
		consulAddr: consulAddr,
		clients:    make(map[string]*grpc.ClientConn),
	}
}

// getServiceClient 获取服务客户端
func (g *Gateway) getServiceClient(serviceName string) (interface{}, error) {
	// 从 Consul 发现服务
	config := consulapi.DefaultConfig()
	config.Address = g.consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil || len(services) == 0 {
		return nil, fmt.Errorf("未找到服务: %s", serviceName)
	}

	service := services[0].Service
	address := fmt.Sprintf("%s:%d", service.Address, service.Port)

	// 检查是否已有连接
	if conn, ok := g.clients[address]; ok {
		return g.createClient(serviceName, conn)
	}

	// 创建新连接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	g.clients[address] = conn

	return g.createClient(serviceName, conn)
}

// createClient 创建服务客户端
func (g *Gateway) createClient(serviceName string, conn *grpc.ClientConn) (interface{}, error) {
	switch serviceName {
	case "user-service":
		return userpb.NewUserServiceClient(conn), nil
	case "order-service":
		return orderpb.NewOrderServiceClient(conn), nil
	case "product-service":
		return productpb.NewProductServiceClient(conn), nil
	default:
		return nil, fmt.Errorf("未知服务: %s", serviceName)
	}
}
```

### 3. 用户服务路由处理

```go
// CreateUser 创建用户
func (g *Gateway) CreateUser(c *gin.Context) {
	var req userpb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUser 获取用户
func (g *Gateway) GetUser(c *gin.Context) {
	userID := c.Param("id")
	req := &userpb.GetUserRequest{UserId: parseID(userID)}

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.GetUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
```

### 4. CORS 中间件

```go
// CORSMiddleware CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
```

### 5. 辅助函数

```go
// parseID 解析 ID
func parseID(s string) int64 {
	var id int64
	fmt.Sscanf(s, "%d", &id)
	return id
}

// parseInt 解析整数
func parseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	var val int
	fmt.Sscanf(s, "%d", &val)
	if val <= 0 {
		return defaultValue
	}
	return val
}
```

## 🔑 关键功能

### 1. 服务发现

网关通过 Consul 动态发现服务：

- 自动发现服务实例
- 支持服务健康检查
- 支持负载均衡

### 2. 协议转换

将 HTTP 请求转换为 gRPC 调用：

- HTTP JSON → gRPC Protobuf
- HTTP 响应 ← gRPC 响应

### 3. 连接管理

- 复用 gRPC 连接
- 连接池管理
- 自动重连

## 🚀 运行网关

```bash
# 确保所有服务已启动
# 1. 启动 Consul
consul agent -dev

# 2. 启动用户服务
cd user-service && go run main.go

# 3. 启动商品服务
cd product-service && go run main.go

# 4. 启动订单服务
cd order-service && go run main.go

# 5. 启动网关
cd gateway
go mod tidy
go run main.go
```

## 🧪 测试 API

```bash
# 创建用户
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User"
  }'

# 获取用户
curl http://localhost:8080/api/users/1

# 创建商品
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "description": "Test Description",
    "price": 99.99,
    "stock": 100,
    "category": "electronics"
  }'

# 创建订单
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "items": [
      {
        "product_id": 1,
        "quantity": 2,
        "price": 99.99
      }
    ],
    "shipping_address": "123 Main St"
  }'
```

## 📚 下一步

API 网关实现完成后，可以继续：

1. [部署运维](./07-deployment.md) - 部署和运维实践

---

**✅ API 网关实现完成！** 现在可以部署整个系统了。

