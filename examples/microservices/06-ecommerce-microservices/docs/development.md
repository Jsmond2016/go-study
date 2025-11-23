# 开发环境搭建

## 前置要求

- Go 1.21 或更高版本
- Protocol Buffers 编译器 (protoc)
- Consul（用于服务发现）

## 安装依赖

### 1. 安装 Protocol Buffers

```bash
# macOS
brew install protobuf

# Linux
sudo apt-get install protobuf-compiler

# 验证
protoc --version
```

### 2. 安装 Go 插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. 安装 Consul

```bash
# macOS
brew install consul

# Linux
wget https://releases.hashicorp.com/consul/1.16.0/consul_1.16.0_linux_amd64.zip
unzip consul_1.16.0_linux_amd64.zip
sudo mv consul /usr/local/bin/
```

## 项目设置

### 1. 克隆项目

```bash
cd examples/microservices/06-ecommerce-microservices
```

### 2. 生成 Protocol Buffers 代码

```bash
cd proto
make proto
```

### 3. 安装 Go 依赖

```bash
# 用户服务
cd ../user-service
go mod tidy

# 订单服务
cd ../order-service
go mod tidy

# 商品服务
cd ../product-service
go mod tidy

# 网关
cd ../gateway
go mod tidy
```

## 开发流程

### 1. 启动 Consul

```bash
consul agent -dev
```

访问 Consul UI: http://localhost:8500

### 2. 启动服务

按以下顺序启动服务：

1. 用户服务
2. 商品服务
3. 订单服务（依赖用户服务）
4. API 网关

### 3. 测试 API

使用 curl 或 Postman 测试 API：

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

## 代码结构

```
06-ecommerce-microservices/
├── proto/              # Protocol Buffers 定义
│   ├── user.proto
│   ├── order.proto
│   └── product.proto
├── user-service/       # 用户服务
│   ├── main.go
│   └── go.mod
├── order-service/      # 订单服务
│   ├── main.go
│   └── go.mod
├── product-service/    # 商品服务
│   ├── main.go
│   └── go.mod
├── gateway/            # API 网关
│   ├── main.go
│   └── go.mod
└── docs/               # 文档
```

## 调试技巧

1. **查看 Consul 服务注册**:
   ```bash
   consul members
   curl http://localhost:8500/v1/agent/services
   ```

2. **查看服务日志**: 每个服务都会输出详细日志

3. **测试 gRPC 服务**: 可以使用 grpcurl 工具

4. **数据库查看**: SQLite 数据库文件在各服务目录下

## 常见问题

### Protocol Buffers 生成失败

- 检查 protoc 版本
- 检查 go_package 选项
- 确保 Go 插件已安装

### 服务无法连接

- 检查 Consul 是否运行
- 检查端口是否被占用
- 检查防火墙设置

### 依赖问题

- 运行 `go mod tidy` 更新依赖
- 检查 Go 版本是否符合要求

