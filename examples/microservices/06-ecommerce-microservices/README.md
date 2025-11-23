# 电商微服务系统

这是一个完整的微服务实战项目，展示了如何使用 gRPC、服务发现、负载均衡和 API 网关构建一个电商系统。

## 🏗️ 系统架构

```
                    ┌─────────────┐
                    │  API Gateway │
                    └──────┬───────┘
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
   ┌────▼────┐      ┌──────▼──────┐    ┌─────▼─────┐
   │  User   │      │   Order     │    │  Product  │
   │ Service │      │   Service   │    │  Service  │
   └─────────┘      └─────────────┘    └───────────┘
        │                  │                  │
        └──────────────────┼──────────────────┘
                           │
                    ┌──────▼──────┐
                    │   Consul     │
                    │ (Service     │
                    │  Discovery)  │
                    └─────────────┘
```

## 📋 服务说明

### 1. 用户服务 (User Service)
- 用户注册和登录
- 用户信息管理
- 用户认证和授权

### 2. 订单服务 (Order Service)
- 订单创建和管理
- 订单状态跟踪
- 调用用户服务验证用户信息

### 3. 商品服务 (Product Service)
- 商品信息管理
- 商品库存管理
- 商品搜索

### 4. API 网关 (API Gateway)
- 统一入口
- 路由转发
- 认证和授权
- 限流和熔断
- 服务发现集成

## 🚀 快速开始

### 1. 启动 Consul

```bash
docker-compose up -d consul
```

或手动启动：

```bash
consul agent -dev
```

### 2. 启动各个服务

```bash
# 终端 1: 启动用户服务
cd user-service
go run main.go

# 终端 2: 启动订单服务
cd order-service
go run main.go

# 终端 3: 启动商品服务
cd product-service
go run main.go

# 终端 4: 启动 API 网关
cd gateway
go run main.go
```

### 3. 测试 API

```bash
# 通过网关访问用户服务
curl http://localhost:8080/api/users/1

# 通过网关访问订单服务
curl http://localhost:8080/api/orders/1

# 通过网关访问商品服务
curl http://localhost:8080/api/products/1
```

## 📚 详细文档

- [项目架构文档](./docs/architecture.md)
- [部署文档](./docs/deployment.md)
- [开发环境搭建](./docs/development.md)
- [API 文档](./docs/api.md)

## 🔧 技术栈

- **gRPC**: 服务间通信
- **Consul**: 服务发现
- **SQLite**: 数据存储（示例）
- **JWT**: 认证和授权
- **Gin**: HTTP 框架（网关）

## 📝 项目结构

```
06-ecommerce-microservices/
├── proto/              # Protocol Buffers 定义
│   ├── user.proto
│   ├── order.proto
│   └── product.proto
├── user-service/       # 用户服务
├── order-service/      # 订单服务
├── product-service/    # 商品服务
├── gateway/            # API 网关
├── docker-compose.yml  # Docker Compose 配置
└── README.md           # 本文件
```

