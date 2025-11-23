---
title: 服务发现
difficulty: advanced
duration: "4-5小时"
prerequisites: ["gRPC 基础", "Protocol Buffers"]
tags: ["服务发现", "服务注册", "Consul", "etcd", "微服务"]
---

# 服务发现

服务发现是微服务架构中的核心组件，它允许服务自动注册和发现其他服务，无需硬编码服务地址。

## 📋 学习目标

完成本教程后，你将能够：

- [ ] 理解服务发现的概念和原理
- [ ] 掌握服务注册机制
- [ ] 掌握服务发现机制
- [ ] 学会使用 Consul 实现服务发现
- [ ] 学会使用 etcd 实现服务发现
- [ ] 实现服务健康检查
- [ ] 处理服务注销和故障转移

## 🎯 服务发现简介

### 什么是服务发现

服务发现是微服务架构中的一个关键模式，它解决了以下问题：

- **动态服务地址**：服务实例可能动态启动和停止
- **负载均衡**：多个服务实例需要负载均衡
- **服务健康**：需要知道哪些服务实例是健康的
- **配置解耦**：避免硬编码服务地址

### 服务发现模式

#### 1. 客户端发现（Client-Side Discovery）

客户端直接查询服务注册表，然后选择服务实例进行调用。

```
客户端 → 服务注册表 → 获取服务列表 → 选择实例 → 调用服务
```

#### 2. 服务端发现（Server-Side Discovery）

客户端通过负载均衡器调用服务，负载均衡器查询服务注册表。

```
客户端 → 负载均衡器 → 服务注册表 → 选择实例 → 调用服务
```

### 服务发现组件

- **服务注册表（Service Registry）**：存储服务实例信息的数据库
- **服务注册（Service Registration）**：服务启动时注册自己
- **服务发现（Service Discovery）**：客户端查询可用服务实例
- **健康检查（Health Check）**：定期检查服务健康状态

## 🚀 使用 Consul

Consul 是 HashiCorp 开发的服务发现和配置管理工具。

### Consul 特性

- 服务发现
- 健康检查
- Key/Value 存储
- 多数据中心支持
- Web UI

### 安装 Consul

```bash
# macOS
brew install consul

# Linux
wget https://releases.hashicorp.com/consul/1.16.0/consul_1.16.0_linux_amd64.zip
unzip consul_1.16.0_linux_amd64.zip
sudo mv consul /usr/local/bin/

# 验证
consul --version
```

### 启动 Consul

```bash
# 开发模式（单节点）
consul agent -dev

# 生产模式
consul agent -server -bootstrap-expect=1 -data-dir=/tmp/consul -ui
```

访问 Web UI: http://localhost:8500

### Go 客户端集成

#### 安装依赖

```bash
go get github.com/hashicorp/consul/api
```

#### 服务注册

```go
package main

import (
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
)

func registerService(consulClient *api.Client, serviceName string, port int) error {
	registration := &api.AgentServiceRegistration{
		ID:      serviceName + "-1",
		Name:    serviceName,
		Tags:    []string{"v1", "golang"},
		Port:    port,
		Address: getLocalIP(),
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + getLocalIP() + ":8080/health",
			Interval: "10s",
			Timeout:  "3s",
		},
	}

	return consulClient.Agent().ServiceRegister(registration)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func main() {
	// 创建 Consul 客户端
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 注册服务
	if err := registerService(client, "user-service", 50051); err != nil {
		log.Fatal(err)
	}

	log.Println("Service registered successfully")

	// 保持运行
	select {}
}
```

#### 服务发现

```go
package main

import (
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

func discoverService(consulClient *api.Client, serviceName string) ([]*api.ServiceEntry, error) {
	services, _, err := consulClient.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func getServiceAddress(services []*api.ServiceEntry) string {
	if len(services) == 0 {
		return ""
	}
	// 简单的负载均衡：选择第一个健康服务
	service := services[0]
	return service.Service.Address + ":" + string(rune(service.Service.Port))
}

func main() {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 定期发现服务
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		services, err := discoverService(client, "user-service")
		if err != nil {
			log.Printf("Error discovering service: %v", err)
			continue
		}

		if len(services) == 0 {
			log.Println("No healthy services found")
			continue
		}

		address := getServiceAddress(services)
		log.Printf("Found service at: %s", address)
	}
}
```

### gRPC 集成

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "your-project/proto"
)

type ServiceDiscovery struct {
	consulClient *api.Client
	serviceName  string
}

func NewServiceDiscovery(consulAddr, serviceName string) (*ServiceDiscovery, error) {
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ServiceDiscovery{
		consulClient: client,
		serviceName:  serviceName,
	}, nil
}

func (sd *ServiceDiscovery) GetService() (string, error) {
	services, _, err := sd.consulClient.Health().Service(sd.serviceName, "", true, nil)
	if err != nil {
		return "", err
	}

	if len(services) == 0 {
		return "", fmt.Errorf("no healthy services found")
	}

	service := services[0]
	return fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port), nil
}

func main() {
	sd, err := NewServiceDiscovery("localhost:8500", "user-service")
	if err != nil {
		log.Fatal(err)
	}

	// 获取服务地址
	address, err := sd.GetService()
	if err != nil {
		log.Fatal(err)
	}

	// 连接 gRPC 服务
	conn, err := grpc.Dial(address,
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

## 🔧 使用 etcd

etcd 是一个分布式键值存储系统，常用于服务发现和配置管理。

### etcd 特性

- 分布式一致性
- 高可用性
- Watch 机制
- TTL（Time To Live）

### 安装 etcd

```bash
# macOS
brew install etcd

# Linux
wget https://github.com/etcd-io/etcd/releases/download/v3.5.9/etcd-v3.5.9-linux-amd64.tar.gz
tar -xzf etcd-v3.5.9-linux-amd64.tar.gz
sudo mv etcd-v3.5.9-linux-amd64/etcd* /usr/local/bin/
```

### 启动 etcd

```bash
# 单节点模式
etcd

# 生产模式（需要配置）
etcd --name node1 \
     --data-dir /tmp/etcd \
     --listen-client-urls http://localhost:2379 \
     --advertise-client-urls http://localhost:2379
```

### Go 客户端集成

#### 安装依赖

```bash
go get go.etcd.io/etcd/client/v3
```

#### 服务注册

```go
package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func registerService(etcdClient *clientv3.Client, serviceName, address string, ttl int64) error {
	// 创建租约
	lease, err := etcdClient.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}

	// 注册服务
	key := "/services/" + serviceName + "/" + address
	_, err = etcdClient.Put(context.Background(), key, address, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	// 保持租约活跃
	go func() {
		ch, kaerr := etcdClient.KeepAlive(context.Background(), lease.ID)
		if kaerr != nil {
			log.Printf("KeepAlive error: %v", kaerr)
			return
		}

		for ka := range ch {
			log.Printf("Lease kept alive: %v", ka)
		}
	}()

	return nil
}

func main() {
	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 注册服务
	if err := registerService(cli, "user-service", "localhost:50051", 10); err != nil {
		log.Fatal(err)
	}

	log.Println("Service registered successfully")

	// 保持运行
	select {}
}
```

#### 服务发现

```go
package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func discoverServices(etcdClient *clientv3.Client, serviceName string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := etcdClient.Get(ctx, "/services/"+serviceName+"/", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, kv := range resp.Kvs {
		addresses = append(addresses, string(kv.Value))
	}

	return addresses, nil
}

func watchServices(etcdClient *clientv3.Client, serviceName string) {
	watchChan := etcdClient.Watch(context.Background(), "/services/"+serviceName+"/", clientv3.WithPrefix())

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				log.Printf("Service added: %s = %s", event.Kv.Key, event.Kv.Value)
			case clientv3.EventTypeDelete:
				log.Printf("Service removed: %s", event.Kv.Key)
			}
		}
	}
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 发现服务
	addresses, err := discoverServices(cli, "user-service")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found services: %v", addresses)

	// 监听服务变化
	go watchServices(cli, "user-service")

	// 保持运行
	select {}
}
```

## 🏥 健康检查

### HTTP 健康检查

```go
package main

import (
	"net/http"
	"time"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// 检查数据库连接、外部服务等
	if isHealthy() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Unhealthy"))
	}
}

func isHealthy() bool {
	// 实现健康检查逻辑
	return true
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	http.ListenAndServe(":8080", nil)
}
```

### gRPC 健康检查

```go
import (
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// 创建健康检查服务
	healthServer := health.NewServer()
	healthServer.SetServingStatus("user-service", grpc_health_v1.HealthCheckResponse_SERVING)

	// 注册到 gRPC 服务器
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
}
```

## 🔄 服务注销

### Consul 服务注销

```go
func deregisterService(consulClient *api.Client, serviceID string) error {
	return consulClient.Agent().ServiceDeregister(serviceID)
}

// 优雅关闭
func gracefulShutdown(consulClient *api.Client, serviceID string) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down...")

	// 注销服务
	if err := deregisterService(consulClient, serviceID); err != nil {
		log.Printf("Error deregistering service: %v", err)
	}

	log.Println("Service deregistered")
}
```

### etcd 服务注销

```go
func deregisterService(etcdClient *clientv3.Client, serviceName, address string) error {
	key := "/services/" + serviceName + "/" + address
	_, err := etcdClient.Delete(context.Background(), key)
	return err
}
```

## 💡 最佳实践

### 1. 使用心跳机制

定期发送心跳以保持服务注册：

```go
func keepAlive(consulClient *api.Client, serviceID string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		consulClient.Agent().PassTTL("service:"+serviceID, "ok")
	}
}
```

### 2. 实现重试机制

服务发现失败时重试：

```go
func getServiceWithRetry(sd *ServiceDiscovery, maxRetries int) (string, error) {
	for i := 0; i < maxRetries; i++ {
		address, err := sd.GetService()
		if err == nil {
			return address, nil
		}
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	return "", fmt.Errorf("failed after %d retries", maxRetries)
}
```

### 3. 缓存服务列表

减少对注册表的查询：

```go
type ServiceCache struct {
	services  []string
	lastUpdate time.Time
	ttl        time.Duration
	mu         sync.RWMutex
}

func (sc *ServiceCache) GetServices() []string {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.services
}

func (sc *ServiceCache) Refresh(discover func() ([]string, error)) error {
	services, err := discover()
	if err != nil {
		return err
	}

	sc.mu.Lock()
	sc.services = services
	sc.lastUpdate = time.Now()
	sc.mu.Unlock()

	return nil
}
```

### 4. 负载均衡

实现简单的负载均衡：

```go
type LoadBalancer struct {
	services []string
	current  int
	mu       sync.Mutex
}

func (lb *LoadBalancer) Next() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.services) == 0 {
		return ""
	}

	service := lb.services[lb.current]
	lb.current = (lb.current + 1) % len(lb.services)
	return service
}
```

## 📝 实践练习

1. **基础练习**：使用 Consul 实现简单的服务注册和发现
2. **健康检查练习**：实现 HTTP 和 gRPC 健康检查
3. **Watch 练习**：使用 etcd Watch 机制监听服务变化
4. **负载均衡练习**：实现基于服务发现的负载均衡
5. **综合练习**：构建一个完整的服务发现系统

## 🔗 相关资源

- [Consul 官方文档](https://www.consul.io/docs)
- [etcd 官方文档](https://etcd.io/docs/)
- [代码示例](../../examples/microservices/03-service-discovery/)

## ⏭️ 下一步

完成服务发现学习后，可以继续学习：

- [负载均衡](./04-load-balancing.md) - 实现负载均衡策略
- [API 网关](./05-api-gateway.md) - 构建 API 网关

---

**🎉 恭喜！** 你已经掌握了服务发现的核心知识。现在可以构建动态的微服务架构了！

