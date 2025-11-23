package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	serviceName = flag.String("service", "user-service", "要发现的服务名称")
	etcdAddr    = flag.String("etcd", "localhost:2379", "etcd 地址")
)

func main() {
	flag.Parse()

	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*etcdAddr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("创建 etcd 客户端失败: %v", err)
	}
	defer cli.Close()

	// 服务发现
	discoverer := NewServiceDiscoverer(cli)

	// 定期发现服务
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case <-ticker.C:
			services, err := discoverer.Discover(*serviceName)
			if err != nil {
				log.Printf("发现服务失败: %v", err)
				continue
			}

			if len(services) == 0 {
				log.Printf("未找到服务: %s", *serviceName)
				continue
			}

			log.Printf("发现 %d 个服务实例:", len(services))
			for i, svc := range services {
				log.Printf("  [%d] %s@%s (键: %s)", i+1, svc.Service, svc.Address, svc.Key)
			}

		case <-ctx.Done():
			return
		}
	}
}

// ServiceInstance 服务实例信息
type ServiceInstance struct {
	Key     string
	Service string
	Address string
}

// ServiceDiscoverer 服务发现器
type ServiceDiscoverer struct {
	client *clientv3.Client
}

// NewServiceDiscoverer 创建服务发现器
func NewServiceDiscoverer(client *clientv3.Client) *ServiceDiscoverer {
	return &ServiceDiscoverer{client: client}
}

// Discover 发现服务实例
func (sd *ServiceDiscoverer) Discover(serviceName string) ([]ServiceInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 查询服务键前缀
	prefix := fmt.Sprintf("/services/%s/", serviceName)
	resp, err := sd.client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("查询服务失败: %w", err)
	}

	instances := make([]ServiceInstance, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		// 解析键：/services/{serviceName}/{instanceId}
		key := string(kv.Key)
		parts := strings.Split(key, "/")
		if len(parts) < 4 {
			continue
		}

		instances = append(instances, ServiceInstance{
			Key:     key,
			Service: serviceName,
			Address: string(kv.Value),
		})
	}

	return instances, nil
}

// Watch 监听服务变化
func (sd *ServiceDiscoverer) Watch(serviceName string, callback func([]ServiceInstance)) error {
	ctx := context.Background()
	prefix := fmt.Sprintf("/services/%s/", serviceName)

	// 初始获取
	instances, err := sd.Discover(serviceName)
	if err != nil {
		return err
	}
	callback(instances)

	// 监听变化
	watchChan := sd.client.Watch(ctx, prefix, clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			log.Printf("服务变化: Type=%s, Key=%s", event.Type, event.Kv.Key)

			// 重新获取所有服务
			instances, err := sd.Discover(serviceName)
			if err != nil {
				log.Printf("重新发现服务失败: %v", err)
				continue
			}
			callback(instances)
		}
	}

	return nil
}

// GetServiceAddress 获取服务地址（格式：host:port）
func (sd *ServiceDiscoverer) GetServiceAddress(serviceName string) (string, error) {
	instances, err := sd.Discover(serviceName)
	if err != nil {
		return "", err
	}

	if len(instances) == 0 {
		return "", fmt.Errorf("未找到服务实例: %s", serviceName)
	}

	// 返回第一个实例的地址
	return instances[0].Address, nil
}

