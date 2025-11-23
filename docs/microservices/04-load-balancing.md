---
title: 负载均衡
difficulty: advanced
duration: "4-5小时"
prerequisites: ["gRPC 基础", "服务发现"]
tags: ["负载均衡", "算法", "微服务", "高可用"]
---

# 负载均衡

负载均衡是分布式系统中的重要组件，它能够将请求分发到多个服务实例，提高系统的可用性和性能。

## 📋 学习目标

完成本教程后，你将能够：

- [ ] 理解负载均衡的概念和原理
- [ ] 掌握常见的负载均衡算法
- [ ] 实现客户端负载均衡
- [ ] 实现服务端负载均衡
- [ ] 集成 gRPC 负载均衡
- [ ] 配置 Nginx 负载均衡
- [ ] 进行性能测试和优化

## 🎯 负载均衡简介

### 什么是负载均衡

负载均衡（Load Balancing）是将网络流量或工作负载分发到多个服务器或服务实例的过程，目的是：

- **提高性能**：分散请求，避免单点过载
- **提高可用性**：单个实例故障不影响整体服务
- **扩展性**：可以动态添加或移除实例
- **资源利用**：充分利用所有服务器资源

### 负载均衡类型

#### 1. 客户端负载均衡（Client-Side Load Balancing）

客户端负责选择服务实例：

```
客户端 → 服务发现 → 选择实例 → 直接调用
```

**优点**：
- 减少中间层，降低延迟
- 客户端可以自定义选择策略

**缺点**：
- 客户端逻辑复杂
- 需要客户端实现

#### 2. 服务端负载均衡（Server-Side Load Balancing）

通过负载均衡器分发请求：

```
客户端 → 负载均衡器 → 选择实例 → 转发请求
```

**优点**：
- 客户端简单
- 集中管理
- 支持 SSL 终止

**缺点**：
- 增加一跳，可能成为瓶颈
- 需要额外的硬件或软件

## 🔄 负载均衡算法

### 1. 轮询（Round Robin）

按顺序轮流分配请求到每个服务器。

```go
package main

import (
	"sync"
)

type RoundRobin struct {
	servers []string
	current int
	mu      sync.Mutex
}

func NewRoundRobin(servers []string) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		current: 0,
	}
}

func (rr *RoundRobin) Next() string {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if len(rr.servers) == 0 {
		return ""
	}

	server := rr.servers[rr.current]
	rr.current = (rr.current + 1) % len(rr.servers)
	return server
}
```

### 2. 加权轮询（Weighted Round Robin）

根据服务器权重分配请求。

```go
type WeightedServer struct {
	Address string
	Weight  int
	Current int
}

type WeightedRoundRobin struct {
	servers []*WeightedServer
	mu      sync.Mutex
}

func NewWeightedRoundRobin(servers []*WeightedServer) *WeightedRoundRobin {
	return &WeightedRoundRobin{servers: servers}
}

func (wrr *WeightedRoundRobin) Next() string {
	wrr.mu.Lock()
	defer wrr.mu.Unlock()

	if len(wrr.servers) == 0 {
		return ""
	}

	// 找到当前权重最大的服务器
	maxWeight := -1
	selectedIndex := 0

	for i, server := range wrr.servers {
		server.Current += server.Weight
		if server.Current > maxWeight {
			maxWeight = server.Current
			selectedIndex = i
		}
	}

	// 减少选中服务器的当前权重
	wrr.servers[selectedIndex].Current -= wrr.getTotalWeight()

	return wrr.servers[selectedIndex].Address
}

func (wrr *WeightedRoundRobin) getTotalWeight() int {
	total := 0
	for _, server := range wrr.servers {
		total += server.Weight
	}
	return total
}
```

### 3. 最少连接（Least Connections）

选择当前连接数最少的服务器。

```go
type ServerStats struct {
	Address    string
	Connections int
	mu         sync.RWMutex
}

type LeastConnections struct {
	servers []*ServerStats
	mu      sync.Mutex
}

func NewLeastConnections(servers []string) *LeastConnections {
	stats := make([]*ServerStats, len(servers))
	for i, addr := range servers {
		stats[i] = &ServerStats{Address: addr}
	}
	return &LeastConnections{servers: stats}
}

func (lc *LeastConnections) Next() string {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if len(lc.servers) == 0 {
		return ""
	}

	minConnections := lc.servers[0].Connections
	selectedIndex := 0

	for i, server := range lc.servers {
		server.mu.RLock()
		conns := server.Connections
		server.mu.RUnlock()

		if conns < minConnections {
			minConnections = conns
			selectedIndex = i
		}
	}

	lc.servers[selectedIndex].mu.Lock()
	lc.servers[selectedIndex].Connections++
	lc.servers[selectedIndex].mu.Unlock()

	return lc.servers[selectedIndex].Address
}

func (lc *LeastConnections) Release(address string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	for _, server := range lc.servers {
		if server.Address == address {
			server.mu.Lock()
			if server.Connections > 0 {
				server.Connections--
			}
			server.mu.Unlock()
			break
		}
	}
}
```

### 4. 一致性哈希（Consistent Hashing）

根据请求的 key 哈希值选择服务器，相同 key 总是路由到同一服务器。

```go
import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashRing struct {
	servers map[uint32]string
	sortedKeys []uint32
	mu      sync.RWMutex
}

func NewHashRing(servers []string) *HashRing {
	hr := &HashRing{
		servers: make(map[uint32]string),
	}

	for _, server := range servers {
		hr.AddServer(server)
	}

	return hr
}

func (hr *HashRing) AddServer(server string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	// 为每个服务器创建多个虚拟节点
	for i := 0; i < 150; i++ {
		hash := hr.hashKey(server + ":" + strconv.Itoa(i))
		hr.servers[hash] = server
		hr.sortedKeys = append(hr.sortedKeys, hash)
	}

	sort.Slice(hr.sortedKeys, func(i, j int) bool {
		return hr.sortedKeys[i] < hr.sortedKeys[j]
	})
}

func (hr *HashRing) RemoveServer(server string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	for i := 0; i < 150; i++ {
		hash := hr.hashKey(server + ":" + strconv.Itoa(i))
		delete(hr.servers, hash)

		// 从排序的 keys 中删除
		index := sort.Search(len(hr.sortedKeys), func(i int) bool {
			return hr.sortedKeys[i] >= hash
		})
		if index < len(hr.sortedKeys) && hr.sortedKeys[index] == hash {
			hr.sortedKeys = append(hr.sortedKeys[:index], hr.sortedKeys[index+1:]...)
		}
	}
}

func (hr *HashRing) GetServer(key string) string {
	hr.mu.RLock()
	defer hr.mu.RUnlock()

	if len(hr.sortedKeys) == 0 {
		return ""
	}

	hash := hr.hashKey(key)

	// 找到第一个大于等于 hash 的服务器
	index := sort.Search(len(hr.sortedKeys), func(i int) bool {
		return hr.sortedKeys[i] >= hash
	})

	// 如果没找到，使用第一个（环形）
	if index >= len(hr.sortedKeys) {
		index = 0
	}

	return hr.servers[hr.sortedKeys[index]]
}

func (hr *HashRing) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
```

## 🔧 客户端负载均衡实现

### gRPC 客户端负载均衡

```go
package main

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	pb "your-project/proto"
)

type LoadBalancer struct {
	servers []string
	lb      *RoundRobin
	mu      sync.RWMutex
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		lb:      NewRoundRobin(servers),
	}
}

func (lb *LoadBalancer) UpdateServers(servers []string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.servers = servers
	lb.lb = NewRoundRobin(servers)
}

func (lb *LoadBalancer) GetConnection() (*grpc.ClientConn, error) {
	lb.mu.RLock()
	server := lb.lb.Next()
	lb.mu.RUnlock()

	if server == "" {
		return nil, fmt.Errorf("no servers available")
	}

	conn, err := grpc.Dial(server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (lb *LoadBalancer) CallRPC(ctx context.Context, method func(*grpc.ClientConn) error) error {
	conn, err := lb.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	return method(conn)
}

func main() {
	lb := NewLoadBalancer([]string{
		"localhost:50051",
		"localhost:50052",
		"localhost:50053",
	})

	// 调用 RPC
	err := lb.CallRPC(context.Background(), func(conn *grpc.ClientConn) error {
		client := pb.NewUserServiceClient(conn)
		user, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: "123"})
		if err != nil {
			return err
		}
		log.Printf("User: %+v", user)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
```

### 使用 gRPC 内置负载均衡

```go
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

func main() {
	// 注册轮询负载均衡器
	resolver.Register(&customResolverBuilder{})

	// 使用负载均衡连接
	conn, err := grpc.Dial(
		"custom:///user-service",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	// 使用客户端...
}
```

## 🌐 服务端负载均衡（Nginx）

### Nginx 配置

```nginx
upstream grpc_servers {
    # 轮询
    server localhost:50051;
    server localhost:50052;
    server localhost:50053;

    # 或者使用最少连接
    # least_conn;

    # 或者使用加权轮询
    # server localhost:50051 weight=3;
    # server localhost:50052 weight=2;
    # server localhost:50053 weight=1;
}

server {
    listen 50050 http2;

    location / {
        grpc_pass grpc://grpc_servers;

        # 健康检查
        grpc_next_upstream error timeout invalid_header http_500;
        grpc_next_upstream_tries 3;
        grpc_next_upstream_timeout 10s;
    }
}
```

### HTTP 负载均衡

```nginx
upstream http_servers {
    server localhost:8080;
    server localhost:8081;
    server localhost:8082;

    # 健康检查
    keepalive 32;
}

server {
    listen 80;

    location / {
        proxy_pass http://http_servers;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

        # 健康检查
        proxy_next_upstream error timeout http_500 http_502 http_503;
    }
}
```

## 📊 健康检查和故障转移

### 健康检查实现

```go
type HealthChecker struct {
	servers map[string]bool
	mu      sync.RWMutex
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		servers: make(map[string]bool),
	}
}

func (hc *HealthChecker) CheckHealth(server string) bool {
	conn, err := grpc.Dial(server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		return false
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	healthClient := grpc_health_v1.NewHealthClient(conn)
	resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return false
	}

	return resp.Status == grpc_health_v1.HealthCheckResponse_SERVING
}

func (hc *HealthChecker) StartHealthCheck(servers []string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, server := range servers {
			healthy := hc.CheckHealth(server)
			hc.mu.Lock()
			hc.servers[server] = healthy
			hc.mu.Unlock()
		}
	}
}

func (hc *HealthChecker) GetHealthyServers(allServers []string) []string {
	hc.mu.RLock()
	defer hc.mu.RUnlock()

	var healthy []string
	for _, server := range allServers {
		if hc.servers[server] {
			healthy = append(healthy, server)
		}
	}
	return healthy
}
```

### 故障转移

```go
type FailoverLoadBalancer struct {
	lb           *RoundRobin
	healthChecker *HealthChecker
	servers      []string
	mu           sync.RWMutex
}

func (flb *FailoverLoadBalancer) Next() string {
	flb.mu.Lock()
	defer flb.mu.Unlock()

	healthy := flb.healthChecker.GetHealthyServers(flb.servers)
	if len(healthy) == 0 {
		// 如果没有健康服务器，返回所有服务器（降级）
		healthy = flb.servers
	}

	flb.lb = NewRoundRobin(healthy)
	return flb.lb.Next()
}
```

## 🧪 性能测试

### 基准测试

```go
func BenchmarkRoundRobin(b *testing.B) {
	lb := NewRoundRobin([]string{
		"server1", "server2", "server3",
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.Next()
	}
}

func BenchmarkConsistentHash(b *testing.B) {
	hr := NewHashRing([]string{
		"server1", "server2", "server3",
	})

	keys := []string{"key1", "key2", "key3", "key4", "key5"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := keys[i%len(keys)]
		hr.GetServer(key)
	}
}
```

### 负载测试

```go
func TestLoadDistribution(t *testing.T) {
	servers := []string{"server1", "server2", "server3"}
	lb := NewRoundRobin(servers)

	distribution := make(map[string]int)
	requests := 1000

	for i := 0; i < requests; i++ {
		server := lb.Next()
		distribution[server]++
	}

	// 验证负载分布
	expected := requests / len(servers)
	tolerance := expected / 10 // 10% 容差

	for server, count := range distribution {
		if count < expected-tolerance || count > expected+tolerance {
			t.Errorf("Server %s: expected ~%d, got %d", server, expected, count)
		}
	}
}
```

## 💡 最佳实践

### 1. 选择合适的算法

- **轮询**：服务器性能相近时使用
- **加权轮询**：服务器性能不同时使用
- **最少连接**：长连接场景使用
- **一致性哈希**：需要会话保持时使用

### 2. 实现健康检查

定期检查服务器健康状态，自动移除不健康的服务器。

### 3. 实现故障转移

当服务器故障时，自动切换到其他服务器。

### 4. 监控和日志

记录负载均衡决策，监控服务器状态和性能。

### 5. 动态更新

支持动态添加和移除服务器，无需重启。

## 📝 实践练习

1. **算法实现**：实现所有常见的负载均衡算法
2. **健康检查**：实现健康检查机制
3. **故障转移**：实现自动故障转移
4. **性能测试**：对比不同算法的性能
5. **综合练习**：构建一个完整的负载均衡系统

## 🔗 相关资源

- [gRPC 负载均衡](https://grpc.io/blog/grpc-load-balancing/)
- [Nginx 负载均衡](https://nginx.org/en/docs/http/load_balancing.html)
- [代码示例](../../examples/microservices/04-load-balancing/)

## ⏭️ 下一步

完成负载均衡学习后，可以继续学习：

- [API 网关](./05-api-gateway.md) - 构建 API 网关

---

**🎉 恭喜！** 你已经掌握了负载均衡的核心知识。现在可以构建高可用的微服务系统了！

