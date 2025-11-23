# 服务发现示例

本示例展示了如何使用 Consul 和 etcd 实现服务注册和发现。

## 📋 目录结构

```
03-service-discovery/
├── consul/              # Consul 服务发现示例
│   ├── register.go      # 服务注册示例
│   └── discover.go      # 服务发现示例
├── etcd/                # etcd 服务发现示例
│   ├── register.go      # 服务注册示例
│   └── discover.go      # 服务发现示例
├── health/              # 健康检查示例
│   ├── heartbeat.go    # 心跳机制示例
│   └── check.go         # 健康检查示例
├── docker-compose.yml   # Docker Compose 配置
├── go.mod               # Go 模块定义
└── README.md            # 本文件
```

## 🚀 快速开始

### 1. 启动服务注册中心

#### 使用 Docker Compose（推荐）

```bash
docker-compose up -d
```

这将启动：
- Consul: http://localhost:8500
- etcd: http://localhost:2379

#### 手动启动 Consul

```bash
# macOS
brew install consul
consul agent -dev

# Linux
wget https://releases.hashicorp.com/consul/1.16.0/consul_1.16.0_linux_amd64.zip
unzip consul_1.16.0_linux_amd64.zip
sudo mv consul /usr/local/bin/
consul agent -dev
```

访问 Consul UI: http://localhost:8500

#### 手动启动 etcd

```bash
# 使用 Docker
docker run -d -p 2379:2379 -p 2380:2380 \
  --name etcd quay.io/coreos/etcd:v3.5.11 \
  etcd --listen-client-urls=http://0.0.0.0:2379 \
  --advertise-client-urls=http://localhost:2379
```

### 2. 安装依赖

```bash
go mod download
go mod tidy
```

### 3. 运行示例

#### Consul 服务注册

```bash
# 终端 1: 注册服务
go run consul/register.go -service=user-service -port=8080

# 终端 2: 发现服务
go run consul/discover.go -service=user-service
```

#### etcd 服务注册

```bash
# 终端 1: 注册服务
go run etcd/register.go -service=user-service -port=8080

# 终端 2: 发现服务
go run etcd/discover.go -service=user-service
```

#### 健康检查

```bash
# 启动健康检查（需要先注册服务并获取服务 ID）
go run health/check.go -id=<service-id> -url=http://localhost:8080/health

# 启动心跳
go run health/heartbeat.go -id=<service-id>
```

## 📚 示例说明

### 1. Consul 服务注册 (`consul/register.go`)

**功能**：
- 自动获取本机 IP 地址
- 注册服务到 Consul
- 配置健康检查
- 优雅关闭时自动注销服务

**使用**：
```bash
go run consul/register.go \
  -service=user-service \
  -port=8080 \
  -consul=localhost:8500
```

### 2. Consul 服务发现 (`consul/discover.go`)

**功能**：
- 定期查询健康服务实例
- 显示服务实例信息
- 支持获取服务地址

**使用**：
```bash
go run consul/discover.go \
  -service=user-service \
  -consul=localhost:8500
```

### 3. etcd 服务注册 (`etcd/register.go`)

**功能**：
- 使用租约机制注册服务
- 自动心跳续约
- 优雅关闭时自动注销

**使用**：
```bash
go run etcd/register.go \
  -service=user-service \
  -port=8080 \
  -etcd=localhost:2379 \
  -ttl=30
```

### 4. etcd 服务发现 (`etcd/discover.go`)

**功能**：
- 通过键前缀查询服务
- 支持监听服务变化
- 获取服务地址

**使用**：
```bash
go run etcd/discover.go \
  -service=user-service \
  -etcd=localhost:2379
```

### 5. 健康检查 (`health/check.go`)

**功能**：
- 定期检查服务健康状态
- 更新 Consul 中的服务状态
- 支持自定义检查 URL 和间隔

**使用**：
```bash
go run health/check.go \
  -id=user-service-hostname-8080 \
  -url=http://localhost:8080/health \
  -interval=10s
```

### 6. 心跳机制 (`health/heartbeat.go`)

**功能**：
- 定期发送心跳到 Consul
- 保持服务在线状态
- 可配置心跳间隔

**使用**：
```bash
go run health/heartbeat.go \
  -id=user-service-hostname-8080 \
  -interval=10s
```

## 🔧 集成到 gRPC 服务

### 在 gRPC 服务中使用服务发现

```go
package main

import (
    "go-study/examples/microservices/03-service-discovery/consul"
)

func main() {
    // 创建服务发现器
    discoverer := consul.NewServiceDiscoverer(consulClient)
    
    // 发现服务
    address, err := discoverer.GetServiceAddress("user-service")
    if err != nil {
        log.Fatal(err)
    }
    
    // 连接到服务
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    // ...
}
```

## 🐛 常见问题

### 1. Consul 连接失败

确保 Consul 已启动：
```bash
consul members
```

检查 Consul 地址是否正确（默认：localhost:8500）

### 2. etcd 连接失败

确保 etcd 已启动：
```bash
docker ps | grep etcd
```

检查 etcd 地址是否正确（默认：localhost:2379）

### 3. 服务注册失败

- 检查服务注册中心是否运行
- 检查网络连接
- 检查服务名称和端口是否冲突

### 4. 服务发现为空

- 确保服务已成功注册
- 检查服务名称是否匹配
- 对于 Consul，确保服务健康检查通过

## 📖 相关文档

- [Consul 官方文档](https://www.consul.io/docs)
- [etcd 官方文档](https://etcd.io/docs/)
- [项目文档](../../../docs/microservices/03-service-discovery.md)

## ⚠️ 注意事项

1. **生产环境**：
   - 使用集群模式而非开发模式
   - 配置适当的健康检查间隔
   - 实现服务降级和故障转移

2. **服务 ID**：
   - 确保服务 ID 唯一
   - 建议使用格式：`{service-name}-{hostname}-{port}`

3. **租约和 TTL**：
   - etcd 使用租约机制，需要定期续约
   - Consul 使用健康检查，需要定期更新状态

4. **网络**：
   - 确保服务注册中心可访问
   - 考虑使用服务网格（如 Istio）进行更高级的服务发现

## 🎯 下一步

完成本示例后，可以继续学习：

- [负载均衡示例](../04-load-balancing/) - 实现负载均衡
- [API 网关示例](../05-api-gateway/) - 实现 API 网关

