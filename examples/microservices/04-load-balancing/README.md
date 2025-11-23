# 负载均衡示例

本示例展示了各种负载均衡算法的实现，包括轮询、加权轮询、最少连接和一致性哈希。

## 📋 目录结构

```
04-load-balancing/
├── algorithms/          # 负载均衡算法实现
│   ├── round_robin.go              # 轮询算法
│   ├── weighted_round_robin.go     # 加权轮询算法
│   ├── least_connections.go        # 最少连接算法
│   ├── consistent_hash.go          # 一致性哈希算法
│   ├── main.go                     # 算法测试主程序
│   └── load_balancer_test.go       # 单元测试
├── grpc/                # gRPC 负载均衡示例
│   └── client.go        # gRPC 客户端负载均衡
├── nginx/               # Nginx 负载均衡配置
│   ├── nginx.conf       # Nginx 配置文件
│   ├── docker-compose.yml
│   └── server*.html     # 后端服务器示例页面
├── go.mod               # Go 模块定义
└── README.md            # 本文件
```

## 🚀 快速开始

### 1. 安装依赖

```bash
go mod download
go mod tidy
```

### 2. 运行算法示例

```bash
# 轮询算法
go run algorithms/main.go round-robin

# 加权轮询算法
go run algorithms/main.go weighted-round-robin

# 最少连接算法
go run algorithms/main.go least-connections

# 一致性哈希算法
go run algorithms/main.go consistent-hash
```

### 3. 运行单元测试

```bash
go test ./algorithms/...
```

### 4. 运行 gRPC 负载均衡示例

```bash
# 需要先启动多个 gRPC 服务实例
go run grpc/client.go \
  -servers=localhost:50051,localhost:50052,localhost:50053 \
  -algorithm=round-robin
```

### 5. 运行 Nginx 负载均衡

```bash
cd nginx
docker-compose up -d

# 测试负载均衡
curl http://localhost
```

## 📚 算法说明

### 1. 轮询（Round Robin）

**原理**：按顺序轮流分配请求到每个服务器。

**优点**：
- 实现简单
- 请求均匀分布

**缺点**：
- 不考虑服务器性能差异
- 不考虑服务器当前负载

**适用场景**：服务器性能相近，请求处理时间相似。

### 2. 加权轮询（Weighted Round Robin）

**原理**：根据服务器权重分配请求，权重高的服务器获得更多请求。

**优点**：
- 考虑服务器性能差异
- 平滑分配请求

**缺点**：
- 需要手动配置权重
- 不考虑实时负载

**适用场景**：服务器性能差异明显。

### 3. 最少连接（Least Connections）

**原理**：将请求分配给当前连接数最少的服务器。

**优点**：
- 考虑服务器当前负载
- 适合长连接场景

**缺点**：
- 需要维护连接计数
- 实现相对复杂

**适用场景**：长连接、服务器处理能力不同。

### 4. 一致性哈希（Consistent Hashing）

**原理**：使用哈希函数将请求和服务器映射到哈希环上，请求分配给顺时针最近的服务器。

**优点**：
- 服务器增减时影响最小
- 适合分布式缓存
- 支持会话保持

**缺点**：
- 实现复杂
- 可能出现负载不均（需要虚拟节点）

**适用场景**：分布式缓存、需要会话保持的场景。

## 🔧 gRPC 负载均衡

### 客户端负载均衡

gRPC 支持客户端负载均衡，客户端从服务发现获取服务列表，然后使用负载均衡算法选择服务实例。

```go
// 创建负载均衡器
lb := NewRoundRobinLB(servers)

// 创建连接池
pool := NewConnectionPool(lb)

// 获取连接
conn, err := pool.GetConnection(ctx)
```

### 服务端负载均衡

使用 Nginx 或 Envoy 等代理进行服务端负载均衡。

## 🌐 Nginx 负载均衡配置

### 轮询（默认）

```nginx
upstream backend {
    server server1:8080;
    server server2:8080;
    server server3:8080;
}
```

### 加权轮询

```nginx
upstream backend {
    server server1:8080 weight=5;
    server server2:8080 weight=3;
    server server3:8080 weight=2;
}
```

### 最少连接

```nginx
upstream backend {
    least_conn;
    server server1:8080;
    server server2:8080;
    server server3:8080;
}
```

### IP 哈希（会话保持）

```nginx
upstream backend {
    ip_hash;
    server server1:8080;
    server server2:8080;
    server server3:8080;
}
```

### 健康检查

```nginx
upstream backend {
    server server1:8080 max_fails=3 fail_timeout=30s;
    server server2:8080 max_fails=3 fail_timeout=30s;
    server server3:8080 backup;
}
```

## 📊 性能测试

### 测试负载均衡算法

```bash
# 运行性能测试
go test -bench=. ./algorithms/
```

### 测试 Nginx 负载均衡

```bash
# 使用 Apache Bench
ab -n 1000 -c 10 http://localhost/

# 使用 wrk
wrk -t4 -c100 -d30s http://localhost/
```

## 🐛 常见问题

### 1. 负载不均

- **原因**：算法选择不当或配置错误
- **解决**：使用加权算法或最少连接算法

### 2. 会话丢失

- **原因**：请求被分配到不同服务器
- **解决**：使用 IP 哈希或一致性哈希

### 3. 服务器故障

- **原因**：未配置健康检查
- **解决**：配置健康检查，自动移除故障服务器

## 📖 相关文档

- [Nginx 负载均衡文档](https://nginx.org/en/docs/http/load_balancing.html)
- [gRPC 负载均衡](https://grpc.io/blog/grpc-load-balancing/)
- [项目文档](../../../docs/microservices/04-load-balancing.md)

## ⚠️ 注意事项

1. **算法选择**：
   - 根据实际场景选择合适的算法
   - 考虑服务器性能、请求特性、会话需求

2. **健康检查**：
   - 必须配置健康检查
   - 及时移除故障服务器

3. **监控**：
   - 监控服务器负载
   - 监控请求分布
   - 监控错误率

4. **扩展性**：
   - 支持动态添加/移除服务器
   - 考虑使用服务发现

## 🎯 下一步

完成本示例后，可以继续学习：

- [API 网关示例](../05-api-gateway/) - 实现 API 网关
- [服务发现示例](../03-service-discovery/) - 结合服务发现使用负载均衡

