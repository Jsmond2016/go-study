package main

import (
	"fmt"
	"os"
)

// 需要导入其他算法文件中的类型
// 这些类型在其他文件中定义

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run main.go <algorithm>")
		fmt.Println("算法: round-robin, weighted-round-robin, least-connections, consistent-hash")
		os.Exit(1)
	}

	algorithm := os.Args[1]

	switch algorithm {
	case "round-robin":
		runRoundRobin()
	case "weighted-round-robin":
		runWeightedRoundRobin()
	case "least-connections":
		runLeastConnections()
	case "consistent-hash":
		runConsistentHash()
	default:
		fmt.Printf("未知算法: %s\n", algorithm)
		os.Exit(1)
	}
}

func runRoundRobin() {
	servers := []string{"server1:8080", "server2:8080", "server3:8080"}
	lb := NewRoundRobin(servers)

	fmt.Println("=== 轮询负载均衡 ===")
	for i := 0; i < 10; i++ {
		fmt.Printf("请求 %d -> %s\n", i+1, lb.Next())
	}
}

func runWeightedRoundRobin() {
	servers := map[string]int{
		"server1:8080": 5,
		"server2:8080": 3,
		"server3:8080": 2,
	}
	lb := NewWeightedRoundRobin(servers)

	fmt.Println("=== 加权轮询负载均衡 ===")
	for i := 0; i < 20; i++ {
		fmt.Printf("请求 %d -> %s\n", i+1, lb.Next())
	}
}

func runLeastConnections() {
	servers := []string{"server1:8080", "server2:8080", "server3:8080"}
	lb := NewLeastConnections(servers)

	fmt.Println("=== 最少连接负载均衡 ===")
	for i := 0; i < 10; i++ {
		server := lb.Next()
		fmt.Printf("请求 %d -> %s (连接数: %d)\n", i+1, server, lb.getConnections(server))
		if i%3 == 0 {
			lb.Release(server)
		}
	}
}

func runConsistentHash() {
	ring := NewHashRing(150)

	servers := []string{"server1:8080", "server2:8080", "server3:8080"}
	for _, server := range servers {
		ring.AddServer(server)
	}

	fmt.Println("=== 一致性哈希负载均衡 ===")
	keys := []string{"user1", "user2", "user3", "user4", "user5"}
	for _, key := range keys {
		server := ring.GetServer(key)
		fmt.Printf("键 %s -> %s\n", key, server)
	}
}

