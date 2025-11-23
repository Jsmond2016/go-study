# 部署文档

## 前置要求

- Docker 和 Docker Compose
- Go 1.21+
- Protocol Buffers 编译器

## 使用 Docker Compose 部署

### 1. 生成 Protocol Buffers 代码

```bash
cd proto
make proto
```

### 2. 启动所有服务

```bash
docker-compose up -d
```

这将启动：
- Consul (端口 8500)
- 用户服务 (端口 5001)
- 订单服务 (端口 5002)
- 商品服务 (端口 5003)
- API 网关 (端口 8080)

### 3. 验证服务

```bash
# 检查 Consul
curl http://localhost:8500/v1/agent/services

# 检查网关健康状态
curl http://localhost:8080/health
```

## 手动部署

### 1. 启动 Consul

```bash
consul agent -dev
```

### 2. 生成 Protocol Buffers 代码

```bash
cd proto
make proto
```

### 3. 启动各个服务

```bash
# 终端 1: 用户服务
cd user-service
go mod tidy
go run main.go

# 终端 2: 订单服务
cd order-service
go mod tidy
go run main.go

# 终端 3: 商品服务
cd product-service
go mod tidy
go run main.go

# 终端 4: API 网关
cd gateway
go mod tidy
go run main.go
```

## 环境变量

可以通过环境变量配置服务：

- `CONSUL_ADDR`: Consul 地址（默认: localhost:8500）
- `PORT`: 服务端口（各服务默认不同）

## 生产环境建议

1. **数据库**: 使用 PostgreSQL 或 MySQL 替代 SQLite
2. **服务发现**: 使用 Consul 集群模式
3. **监控**: 集成 Prometheus 和 Grafana
4. **日志**: 使用集中式日志系统（如 ELK）
5. **安全**: 启用 TLS/SSL
6. **限流**: 在网关层添加限流
7. **熔断**: 实现熔断器保护

## 故障排查

### 服务无法注册到 Consul

- 检查 Consul 是否运行
- 检查网络连接
- 检查服务端口是否被占用

### 服务间调用失败

- 检查服务是否已注册到 Consul
- 检查服务是否健康
- 查看服务日志

### 网关无法转发请求

- 检查网关日志
- 检查服务发现配置
- 验证 gRPC 连接

