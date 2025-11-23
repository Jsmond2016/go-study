package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundRobin(t *testing.T) {
	servers := []string{"server1", "server2", "server3"}
	lb := NewRoundRobin(servers)

	// 测试轮询
	assert.Equal(t, "server1", lb.Next())
	assert.Equal(t, "server2", lb.Next())
	assert.Equal(t, "server3", lb.Next())
	assert.Equal(t, "server1", lb.Next()) // 回到第一个
}

func TestWeightedRoundRobin(t *testing.T) {
	servers := map[string]int{
		"server1": 5,
		"server2": 3,
		"server3": 2,
	}
	lb := NewWeightedRoundRobin(servers)

	// 测试加权轮询
	results := make(map[string]int)
	for i := 0; i < 100; i++ {
		server := lb.Next()
		results[server]++
	}

	// server1 应该被选中的次数最多
	assert.Greater(t, results["server1"], results["server2"])
	assert.Greater(t, results["server2"], results["server3"])
}

func TestLeastConnections(t *testing.T) {
	servers := []string{"server1", "server2", "server3"}
	lb := NewLeastConnections(servers)

	// 初始时所有服务器连接数为 0，应该选择第一个
	assert.Equal(t, "server1", lb.Next())

	// 增加 server1 的连接数
	lb.Next() // server1
	lb.Next() // server1

	// 现在应该选择连接数最少的
	lb.Release("server1")
	server := lb.Next()
	assert.Contains(t, []string{"server2", "server3"}, server)
}

func TestConsistentHash(t *testing.T) {
	ring := NewHashRing(150)
	servers := []string{"server1", "server2", "server3"}

	for _, server := range servers {
		ring.AddServer(server)
	}

	// 相同键应该映射到相同服务器
	key := "test-key"
	server1 := ring.GetServer(key)
	server2 := ring.GetServer(key)
	assert.Equal(t, server1, server2)

	// 移除服务器后，部分键应该重新映射
	ring.RemoveServer("server1")
	server3 := ring.GetServer(key)
	// server3 可能和 server1 相同（如果 key 原本就映射到其他服务器）
	assert.NotEmpty(t, server3)
}

