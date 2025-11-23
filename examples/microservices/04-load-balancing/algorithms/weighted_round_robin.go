package main

import (
	"fmt"
	"sync"
)

// Server 服务器信息
type Server struct {
	Address    string
	Weight     int
	CurrentWeight int
}

// WeightedRoundRobin 加权轮询负载均衡器
type WeightedRoundRobin struct {
	servers []*Server
	mu      sync.Mutex
}

// NewWeightedRoundRobin 创建加权轮询负载均衡器
func NewWeightedRoundRobin(servers map[string]int) *WeightedRoundRobin {
	wr := &WeightedRoundRobin{
		servers: make([]*Server, 0, len(servers)),
	}

	for addr, weight := range servers {
		wr.servers = append(wr.servers, &Server{
			Address:       addr,
			Weight:        weight,
			CurrentWeight: 0,
		})
	}

	return wr
}

// Next 获取下一个服务器（平滑加权轮询算法）
func (wr *WeightedRoundRobin) Next() string {
	wr.mu.Lock()
	defer wr.mu.Unlock()

	if len(wr.servers) == 0 {
		return ""
	}

	// 计算总权重
	totalWeight := 0
	for _, server := range wr.servers {
		totalWeight += server.Weight
		server.CurrentWeight += server.Weight
	}

	// 选择当前权重最大的服务器
	maxWeight := -1
	var selected *Server
	for _, server := range wr.servers {
		if server.CurrentWeight > maxWeight {
			maxWeight = server.CurrentWeight
			selected = server
		}
	}

	// 减少选中服务器的当前权重
	selected.CurrentWeight -= totalWeight

	return selected.Address
}

func main() {
	servers := map[string]int{
		"server1:8080": 5,
		"server2:8080": 3,
		"server3:8080": 2,
	}
	lb := NewWeightedRoundRobin(servers)

	fmt.Println("加权轮询负载均衡示例:")
	for i := 0; i < 20; i++ {
		fmt.Printf("请求 %d -> %s\n", i+1, lb.Next())
	}
}

