---
title: 架构设计
difficulty: advanced
duration: "2-3小时"
prerequisites: ["环境搭建"]
tags: ["架构", "设计", "微服务", "系统设计"]
---

# 架构设计

本章节将详细介绍电商微服务系统的架构设计，包括服务划分、通信方式、数据存储和部署方案。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 理解微服务架构的设计原则
- [ ] 掌握服务边界划分方法
- [ ] 设计服务间通信协议
- [ ] 规划数据存储方案
- [ ] 设计部署和运维方案

## 🏗️ 系统架构

### 整体架构图

```
┌─────────────────────────────────────────────────────────┐
│                     客户端层                              │
│              (Web/Mobile/API Client)                     │
└──────────────────────┬──────────────────────────────────┘
                       │ HTTP
                       │
┌──────────────────────▼──────────────────────────────────┐
│                    API 网关层                            │
│  ┌──────────────────────────────────────────────────┐  │
│  │  API Gateway (Port 8080)                        │  │
│  │  - 路由转发                                       │  │
│  │  - 协议转换 (HTTP → gRPC)                        │  │
│  │  - 认证授权                                       │  │
│  │  - 限流熔断                                       │  │
│  └──────────────────────────────────────────────────┘  │
└──────────────────────┬──────────────────────────────────┘
                       │ gRPC
        ┌──────────────┼──────────────┐
        │              │              │
┌───────▼──────┐ ┌─────▼──────┐ ┌─────▼──────┐
│  用户服务     │ │  订单服务   │ │  商品服务   │
│ (Port 5001)  │ │ (Port 5002) │ │ (Port 5003) │
│              │ │             │ │             │
│ - 用户注册    │ │ - 订单创建   │ │ - 商品管理   │
│ - 用户登录    │ │ - 订单查询   │ │ - 库存管理   │
│ - 用户信息    │ │ - 状态更新   │ │ - 商品查询   │
│ - 用户验证    │ │             │ │             │
└───────┬──────┘ └─────┬──────┘ └─────────────┘
        │              │
        │              │ gRPC (调用用户服务)
        │              │
        └──────────────┘
              │
              │ 服务注册/发现
              │
       ┌──────▼──────┐
       │   Consul    │
       │ (Port 8500) │
       │             │
       │ - 服务注册   │
       │ - 服务发现   │
       │ - 健康检查   │
       └─────────────┘
```

## 🎯 服务划分

### 服务边界原则

1. **业务边界**: 按业务领域划分服务
2. **数据边界**: 每个服务拥有独立的数据存储
3. **团队边界**: 服务可以由不同团队独立开发
4. **技术边界**: 服务可以使用不同的技术栈

### 服务职责

#### 1. 用户服务 (User Service)

**职责**:
- 用户注册和登录
- 用户信息管理
- 用户认证和验证
- 用户权限管理

**数据**:
- 用户基本信息
- 用户认证信息
- 用户配置信息

**接口**:
- `CreateUser`: 创建用户
- `GetUser`: 获取用户信息
- `UpdateUser`: 更新用户信息
- `DeleteUser`: 删除用户
- `ValidateUser`: 验证用户
- `Login`: 用户登录

#### 2. 商品服务 (Product Service)

**职责**:
- 商品信息管理
- 商品库存管理
- 商品分类管理
- 商品搜索

**数据**:
- 商品基本信息
- 商品库存信息
- 商品分类信息

**接口**:
- `CreateProduct`: 创建商品
- `GetProduct`: 获取商品信息
- `UpdateProduct`: 更新商品信息
- `DeleteProduct`: 删除商品
- `ListProducts`: 获取商品列表
- `UpdateStock`: 更新库存
- `CheckStock`: 检查库存

#### 3. 订单服务 (Order Service)

**职责**:
- 订单创建和管理
- 订单状态跟踪
- 订单查询
- 订单统计

**数据**:
- 订单基本信息
- 订单项信息
- 订单状态信息

**依赖**:
- 调用用户服务验证用户
- 调用商品服务检查库存（可选）

**接口**:
- `CreateOrder`: 创建订单
- `GetOrder`: 获取订单信息
- `GetUserOrders`: 获取用户订单列表
- `UpdateOrderStatus`: 更新订单状态
- `CancelOrder`: 取消订单

#### 4. API 网关 (API Gateway)

**职责**:
- 统一 API 入口
- 路由转发
- 协议转换 (HTTP → gRPC)
- 认证和授权
- 限流和熔断
- 请求日志和监控

**特点**:
- 无状态服务
- 可水平扩展
- 不存储业务数据

## 🔌 服务间通信

### 通信方式

#### 1. gRPC 同步调用

**使用场景**:
- 需要立即响应的操作
- 服务间数据查询
- 服务间数据验证

**示例**:
```go
// 订单服务调用用户服务验证用户
userClient.ValidateUser(ctx, &userpb.ValidateUserRequest{
    UserId: order.UserId,
})
```

#### 2. 服务发现

**使用 Consul**:
- 服务自动注册
- 服务自动发现
- 健康检查
- 负载均衡

**流程**:
1. 服务启动时注册到 Consul
2. 服务调用时从 Consul 发现目标服务
3. 使用 gRPC 客户端连接服务
4. 定期健康检查

### 通信协议

#### Protocol Buffers 定义

```protobuf
// user.proto
service UserService {
  rpc GetUser(GetUserRequest) returns (UserResponse);
  rpc ValidateUser(ValidateUserRequest) returns (ValidateUserResponse);
}

// order.proto
service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (OrderResponse);
}
```

## 💾 数据存储

### 存储方案

#### 1. 数据隔离

每个服务拥有独立的数据存储：

- **用户服务**: `user_service.db`
- **订单服务**: `order_service.db`
- **商品服务**: `product_service.db`

#### 2. 数据库选择

**开发环境**: SQLite（简单、快速）

**生产环境**:
- PostgreSQL（推荐）
- MySQL
- MongoDB（文档存储）

#### 3. 数据一致性

**最终一致性**:
- 服务间通过事件或消息队列同步
- 允许短暂的数据不一致

**强一致性**:
- 使用分布式事务（2PC、Saga）
- 性能开销较大

### 数据模型

#### 用户服务数据模型

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    full_name TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

#### 订单服务数据模型

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    items TEXT NOT NULL,  -- JSON 格式
    total_amount REAL NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    shipping_address TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

#### 商品服务数据模型

```sql
CREATE TABLE products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    price REAL NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## 🚀 部署方案

### 部署架构

```
┌─────────────────────────────────────────┐
│         Docker Compose / K8s            │
│                                         │
│  ┌──────────┐  ┌──────────┐          │
│  │ Consul    │  │ Gateway  │          │
│  │ (1实例)   │  │ (多实例)  │          │
│  └──────────┘  └──────────┘          │
│                                         │
│  ┌──────────┐  ┌──────────┐          │
│  │ User      │  │ Order    │          │
│  │ (多实例)  │  │ (多实例)  │          │
│  └──────────┘  └──────────┘          │
│                                         │
│  ┌──────────┐                         │
│  │ Product  │                         │
│  │ (多实例)  │                         │
│  └──────────┘                         │
└─────────────────────────────────────────┘
```

### 扩展策略

1. **水平扩展**: 每个服务可以运行多个实例
2. **负载均衡**: 通过 Consul 实现客户端负载均衡
3. **服务隔离**: 每个服务独立部署和扩展

## 🔒 安全设计

### 认证和授权

1. **JWT Token**: 用户登录后获取 Token
2. **Token 验证**: 网关验证 Token
3. **服务间认证**: 使用服务密钥或 mTLS

### 数据安全

1. **密码加密**: 使用 bcrypt 加密密码
2. **敏感数据**: 不传输敏感信息
3. **HTTPS**: 生产环境使用 HTTPS

## 📊 监控和日志

### 监控指标

1. **服务指标**: 请求数、响应时间、错误率
2. **系统指标**: CPU、内存、网络
3. **业务指标**: 订单数、用户数、商品数

### 日志收集

1. **集中式日志**: 使用 ELK 或 Loki
2. **日志格式**: 结构化日志（JSON）
3. **日志级别**: DEBUG、INFO、WARN、ERROR

## 🎯 最佳实践

### 1. 服务设计

- ✅ 单一职责原则
- ✅ 高内聚低耦合
- ✅ 无状态服务
- ✅ 版本管理

### 2. 通信设计

- ✅ 异步优先
- ✅ 超时设置
- ✅ 重试机制
- ✅ 熔断保护

### 3. 数据设计

- ✅ 数据隔离
- ✅ 最终一致性
- ✅ 缓存策略
- ✅ 数据备份

## 📚 下一步

架构设计完成后，可以开始实现各个服务：

1. [用户服务](./03-user-service.md) - 实现用户服务
2. [商品服务](./04-product-service.md) - 实现商品服务
3. [订单服务](./05-order-service.md) - 实现订单服务

---

**✅ 架构设计完成！** 现在可以开始实现各个微服务了。

