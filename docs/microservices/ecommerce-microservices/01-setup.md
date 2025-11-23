---
title: 环境搭建
difficulty: advanced
duration: "2-3小时"
prerequisites: ["gRPC 基础", "服务发现"]
tags: ["环境", "搭建", "配置", "初始化"]
---

# 环境搭建

本章节将指导你搭建电商微服务系统的开发环境，包括项目初始化、依赖安装、Protocol Buffers 配置和服务发现环境。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 创建微服务项目目录结构
- [ ] 初始化各个服务的 Go 模块
- [ ] 安装项目依赖
- [ ] 配置 Protocol Buffers
- [ ] 启动 Consul 服务发现
- [ ] 运行项目并验证

## 🚀 快速开始

### 1. 创建项目目录

```bash
# 创建项目根目录
cd examples/microservices
mkdir -p 06-ecommerce-microservices
cd 06-ecommerce-microservices

# 创建服务目录
mkdir -p user-service
mkdir -p order-service
mkdir -p product-service
mkdir -p gateway
mkdir -p proto/{user,order,product}
mkdir -p docs
```

### 2. 初始化 Go 模块

```bash
# 根模块
go mod init go-study/examples/microservices/06-ecommerce-microservices

# 用户服务
cd user-service
go mod init go-study/examples/microservices/06-ecommerce-microservices/user-service

# 订单服务
cd ../order-service
go mod init go-study/examples/microservices/06-ecommerce-microservices/order-service

# 商品服务
cd ../product-service
go mod init go-study/examples/microservices/06-ecommerce-microservices/product-service

# 网关
cd ../gateway
go mod init go-study/examples/microservices/06-ecommerce-microservices/gateway
```

### 3. 安装核心依赖

#### 用户服务依赖

```bash
cd user-service
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/hashicorp/consul/api
go get github.com/mattn/go-sqlite3
```

#### 订单服务依赖

```bash
cd ../order-service
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/hashicorp/consul/api
go get github.com/mattn/go-sqlite3
```

#### 商品服务依赖

```bash
cd ../product-service
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/hashicorp/consul/api
go get github.com/mattn/go-sqlite3
```

#### 网关依赖

```bash
cd ../gateway
go get github.com/gin-gonic/gin
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/hashicorp/consul/api
```

## 📁 项目结构详解

### 完整目录结构

```
06-ecommerce-microservices/
├── proto/                    # Protocol Buffers 定义
│   ├── user/                 # 用户服务 proto
│   │   └── user.proto
│   ├── order/                # 订单服务 proto
│   │   └── order.proto
│   ├── product/              # 商品服务 proto
│   │   └── product.proto
│   └── Makefile              # 代码生成脚本
├── user-service/             # 用户服务
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── order-service/            # 订单服务
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── product-service/          # 商品服务
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── gateway/                  # API 网关
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── docs/                     # 文档
│   ├── architecture.md
│   ├── deployment.md
│   ├── development.md
│   └── api.md
├── docker-compose.yml        # Docker Compose 配置
├── go.mod                    # 根模块
└── README.md                 # 项目说明
```

### 目录说明

- **proto/**: 存放所有服务的 Protocol Buffers 定义文件
- **user-service/**: 用户服务，处理用户相关业务
- **order-service/**: 订单服务，处理订单相关业务
- **product-service/**: 商品服务，处理商品相关业务
- **gateway/**: API 网关，统一入口和路由转发
- **docs/**: 项目文档

## 🔧 Protocol Buffers 配置

### 1. 安装 Protocol Buffers 编译器

```bash
# macOS
brew install protobuf

# Linux
sudo apt-get install protobuf-compiler

# 验证安装
protoc --version
```

### 2. 安装 Go 插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. 创建 Makefile

在 `proto/` 目录下创建 `Makefile`：

```makefile
.PHONY: proto clean

proto:
	@echo "生成 Go 代码..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		user/user.proto order/order.proto product/product.proto

clean:
	@echo "清理生成的文件..."
	rm -rf user/*.pb.go
	rm -rf order/*.pb.go
	rm -rf product/*.pb.go
```

### 4. 生成代码

```bash
cd proto
make proto
```

## 🏗️ Consul 服务发现配置

### 1. 安装 Consul

```bash
# macOS
brew install consul

# Linux
wget https://releases.hashicorp.com/consul/1.16.0/consul_1.16.0_linux_amd64.zip
unzip consul_1.16.0_linux_amd64.zip
sudo mv consul /usr/local/bin/
```

### 2. 启动 Consul

```bash
# 开发模式启动
consul agent -dev

# 或者使用 Docker
docker run -d --name consul -p 8500:8500 consul:latest
```

### 3. 验证 Consul

访问 Consul UI: http://localhost:8500

或使用命令行：

```bash
consul members
curl http://localhost:8500/v1/agent/services
```

## 📝 配置文件

### 环境变量

各服务可以通过环境变量配置：

```bash
# Consul 地址
export CONSUL_ADDR=localhost:8500

# 服务端口
export USER_SERVICE_PORT=5001
export ORDER_SERVICE_PORT=5002
export PRODUCT_SERVICE_PORT=5003
export GATEWAY_PORT=8080
```

### Docker Compose 配置

创建 `docker-compose.yml`：

```yaml
version: '3.8'

services:
  consul:
    image: consul:latest
    container_name: consul
    ports:
      - "8500:8500"
    command: consul agent -dev -client=0.0.0.0 -ui
    networks:
      - microservices

networks:
  microservices:
    driver: bridge
```

启动 Consul：

```bash
docker-compose up -d consul
```

## ✅ 验证环境

### 1. 检查 Go 环境

```bash
go version
```

### 2. 检查 Protocol Buffers

```bash
protoc --version
which protoc-gen-go
which protoc-gen-go-grpc
```

### 3. 检查 Consul

```bash
consul version
curl http://localhost:8500/v1/status/leader
```

## 🐛 常见问题

### Protocol Buffers 生成失败

**问题**: `protoc: command not found`

**解决**: 确保已安装 protobuf 编译器，并添加到 PATH

### Go 插件未找到

**问题**: `protoc-gen-go: program not found`

**解决**:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Consul 连接失败

**问题**: 服务无法连接到 Consul

**解决**:
- 检查 Consul 是否运行: `consul members`
- 检查端口是否被占用: `lsof -i :8500`
- 检查防火墙设置

### 端口冲突

**问题**: 端口已被占用

**解决**:
- 修改服务端口配置
- 或停止占用端口的进程

## 📚 下一步

环境搭建完成后，可以继续：

1. [架构设计](./02-architecture.md) - 了解系统架构设计
2. [用户服务](./03-user-service.md) - 实现第一个微服务

---

**✅ 环境搭建完成！** 现在可以开始实现各个微服务了。

