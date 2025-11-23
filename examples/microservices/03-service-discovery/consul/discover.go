package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var (
	serviceName = flag.String("service", "user-service", "要发现的服务名称")
	consulAddr  = flag.String("consul", "localhost:8500", "Consul 地址")
)

func main() {
	flag.Parse()

	// 创建 Consul 客户端
	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}

	// 服务发现
	discoverer := NewServiceDiscoverer(client)

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
				log.Printf("  [%d] %s@%s:%d (健康: %v)", i+1, svc.Service, svc.Address, svc.Port, svc.Healthy)
			}

		case <-ctx.Done():
			return
		}
	}
}

// ServiceInstance 服务实例信息
type ServiceInstance struct {
	ID      string
	Service string
	Address string
	Port    int
	Tags    []string
	Healthy bool
}

// ServiceDiscoverer 服务发现器
type ServiceDiscoverer struct {
	client *consulapi.Client
}

// NewServiceDiscoverer 创建服务发现器
func NewServiceDiscoverer(client *consulapi.Client) *ServiceDiscoverer {
	return &ServiceDiscoverer{client: client}
}

// Discover 发现服务实例
func (sd *ServiceDiscoverer) Discover(serviceName string) ([]ServiceInstance, error) {
	// 查询健康服务
	services, _, err := sd.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, fmt.Errorf("查询服务失败: %w", err)
	}

	instances := make([]ServiceInstance, 0, len(services))
	for _, svc := range services {
		instances = append(instances, ServiceInstance{
			ID:      svc.Service.ID,
			Service: svc.Service.Service,
			Address: svc.Service.Address,
			Port:    svc.Service.Port,
			Tags:    svc.Service.Tags,
			Healthy: true, // 通过 Health().Service 查询的默认都是健康的
		})
	}

	return instances, nil
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

	// 返回第一个健康实例的地址
	instance := instances[0]
	return fmt.Sprintf("%s:%d", instance.Address, instance.Port), nil
}

