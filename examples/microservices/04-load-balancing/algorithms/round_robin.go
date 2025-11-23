package main

import (
	"fmt"
	"sync"
)

// RoundRobin 轮询负载均衡器
type RoundRobin struct {
	servers []string
	current int
	mu      sync.Mutex
}

// NewRoundRobin 创建轮询负载均衡器
func NewRoundRobin(servers []string) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		current: 0,
	}
}

// Next 获取下一个服务器
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

// AddServer 添加服务器
func (rr *RoundRobin) AddServer(server string) {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	rr.servers = append(rr.servers, server)
}

// RemoveServer 移除服务器
func (rr *RoundRobin) RemoveServer(server string) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	for i, s := range rr.servers {
		if s == server {
			rr.servers = append(rr.servers[:i], rr.servers[i+1:]...)
			if rr.current >= len(rr.servers) {
				rr.current = 0
			}
			break
		}
	}
}

func main() {
	servers := []string{"server1:8080", "server2:8080", "server3:8080"}
	lb := NewRoundRobin(servers)

	fmt.Println("轮询负载均衡示例:")
	for i := 0; i < 10; i++ {
		fmt.Printf("请求 %d -> %s\n", i+1, lb.Next())
	}
}

