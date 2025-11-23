package main

import (
	"fmt"
	"sync"
)

// ServerStats 服务器统计信息
type ServerStats struct {
	Address    string
	Connections int
}

// LeastConnections 最少连接负载均衡器
type LeastConnections struct {
	servers []*ServerStats
	mu      sync.Mutex
}

// NewLeastConnections 创建最少连接负载均衡器
func NewLeastConnections(servers []string) *LeastConnections {
	lc := &LeastConnections{
		servers: make([]*ServerStats, 0, len(servers)),
	}

	for _, addr := range servers {
		lc.servers = append(lc.servers, &ServerStats{
			Address:     addr,
			Connections: 0,
		})
	}

	return lc
}

// Next 获取下一个服务器（最少连接的）
func (lc *LeastConnections) Next() string {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if len(lc.servers) == 0 {
		return ""
	}

	// 找到连接数最少的服务器
	minConnections := lc.servers[0].Connections
	selected := lc.servers[0]

	for _, server := range lc.servers {
		if server.Connections < minConnections {
			minConnections = server.Connections
			selected = server
		}
	}

	// 增加连接数
	selected.Connections++
	return selected.Address
}

// Release 释放连接
func (lc *LeastConnections) Release(server string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	for _, s := range lc.servers {
		if s.Address == server && s.Connections > 0 {
			s.Connections--
			break
		}
	}
}

func main() {
	servers := []string{"server1:8080", "server2:8080", "server3:8080"}
	lb := NewLeastConnections(servers)

	fmt.Println("最少连接负载均衡示例:")
	for i := 0; i < 10; i++ {
		server := lb.Next()
		fmt.Printf("请求 %d -> %s (连接数: %d)\n", i+1, server, lb.getConnections(server))
		// 模拟释放连接
		if i%3 == 0 {
			lb.Release(server)
		}
	}
}

func (lc *LeastConnections) getConnections(server string) int {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	for _, s := range lc.servers {
		if s.Address == server {
			return s.Connections
		}
	}
	return 0
}

