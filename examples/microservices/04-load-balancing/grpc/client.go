package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	servers = flag.String("servers", "localhost:50051,localhost:50052,localhost:50053", "服务器地址列表（逗号分隔）")
	algorithm = flag.String("algorithm", "round-robin", "负载均衡算法: round-robin, weighted-round-robin, least-connections")
)

func main() {
	flag.Parse()

	// 解析服务器列表
	serverList := parseServers(*servers)

	// 创建负载均衡器
	var lb LoadBalancer
	switch *algorithm {
	case "round-robin":
		lb = NewRoundRobinLB(serverList)
	case "weighted-round-robin":
		weights := make(map[string]int)
		for i, server := range serverList {
			weights[server] = 5 - i // 第一个服务器权重最高
		}
		lb = NewWeightedRoundRobinLB(weights)
	case "least-connections":
		lb = NewLeastConnectionsLB(serverList)
	default:
		log.Fatalf("未知算法: %s", *algorithm)
	}

	// 创建 gRPC 连接池
	pool := NewConnectionPool(lb)

	// 发送请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		conn, err := pool.GetConnection(ctx)
		if err != nil {
			log.Printf("获取连接失败: %v", err)
			continue
		}

		// 这里应该调用实际的 gRPC 服务
		// 示例：假设有一个 HelloService
		fmt.Printf("请求 %d -> 使用连接: %s\n", i+1, conn.Target())

		// 模拟使用连接
		time.Sleep(100 * time.Millisecond)

		// 释放连接
		pool.ReleaseConnection(conn)
	}
}

func parseServers(servers string) []string {
	var result []string
	for _, s := range split(servers, ",") {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

func split(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
		}
	}
	result = append(result, s[start:])
	return result
}

// LoadBalancer 负载均衡器接口
type LoadBalancer interface {
	Next() string
}

// RoundRobinLB 轮询负载均衡器
type RoundRobinLB struct {
	servers []string
	current int
}

func NewRoundRobinLB(servers []string) *RoundRobinLB {
	return &RoundRobinLB{servers: servers, current: 0}
}

func (rr *RoundRobinLB) Next() string {
	if len(rr.servers) == 0 {
		return ""
	}
	server := rr.servers[rr.current]
	rr.current = (rr.current + 1) % len(rr.servers)
	return server
}

// WeightedRoundRobinLB 加权轮询负载均衡器
type WeightedRoundRobinLB struct {
	servers []*WeightedServer
	totalWeight int
}

type WeightedServer struct {
	Address string
	Weight  int
	CurrentWeight int
}

func NewWeightedRoundRobinLB(weights map[string]int) *WeightedRoundRobinLB {
	wr := &WeightedRoundRobinLB{
		servers: make([]*WeightedServer, 0, len(weights)),
	}

	for addr, weight := range weights {
		wr.servers = append(wr.servers, &WeightedServer{
			Address: addr,
			Weight: weight,
		})
		wr.totalWeight += weight
	}

	return wr
}

func (wr *WeightedRoundRobinLB) Next() string {
	if len(wr.servers) == 0 {
		return ""
	}

	// 平滑加权轮询
	for _, server := range wr.servers {
		server.CurrentWeight += server.Weight
	}

	maxWeight := -1
	var selected *WeightedServer
	for _, server := range wr.servers {
		if server.CurrentWeight > maxWeight {
			maxWeight = server.CurrentWeight
			selected = server
		}
	}

	selected.CurrentWeight -= wr.totalWeight
	return selected.Address
}

// LeastConnectionsLB 最少连接负载均衡器
type LeastConnectionsLB struct {
	servers []*ConnectionStats
}

type ConnectionStats struct {
	Address     string
	Connections int
}

func NewLeastConnectionsLB(servers []string) *LeastConnectionsLB {
	lc := &LeastConnectionsLB{
		servers: make([]*ConnectionStats, 0, len(servers)),
	}

	for _, addr := range servers {
		lc.servers = append(lc.servers, &ConnectionStats{
			Address:     addr,
			Connections: 0,
		})
	}

	return lc
}

func (lc *LeastConnectionsLB) Next() string {
	if len(lc.servers) == 0 {
		return ""
	}

	minConnections := lc.servers[0].Connections
	selected := lc.servers[0]

	for _, server := range lc.servers {
		if server.Connections < minConnections {
			minConnections = server.Connections
			selected = server
		}
	}

	selected.Connections++
	return selected.Address
}

func (lc *LeastConnectionsLB) Release(server string) {
	for _, s := range lc.servers {
		if s.Address == server && s.Connections > 0 {
			s.Connections--
			break
		}
	}
}

// ConnectionPool 连接池
type ConnectionPool struct {
	lb     LoadBalancer
	conns  map[string]*grpc.ClientConn
	lc     *LeastConnectionsLB
}

func NewConnectionPool(lb LoadBalancer) *ConnectionPool {
	pool := &ConnectionPool{
		lb:    lb,
		conns: make(map[string]*grpc.ClientConn),
	}

	// 如果是最少连接算法，保存引用
	if lc, ok := lb.(*LeastConnectionsLB); ok {
		pool.lc = lc
	}

	return pool
}

func (cp *ConnectionPool) GetConnection(ctx context.Context) (*grpc.ClientConn, error) {
	server := cp.lb.Next()

	// 检查是否已有连接
	if conn, ok := cp.conns[server]; ok {
		return conn, nil
	}

	// 创建新连接
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("连接失败: %w", err)
	}

	cp.conns[server] = conn
	return conn, nil
}

func (cp *ConnectionPool) ReleaseConnection(conn *grpc.ClientConn) {
	// 如果是最少连接算法，释放连接计数
	if cp.lc != nil {
		cp.lc.Release(conn.Target())
	}
	// 注意：这里不关闭连接，保持连接池
}

func (cp *ConnectionPool) Close() {
	for _, conn := range cp.conns {
		conn.Close()
	}
}

