package main

import (
	"crypto/md5"
	"fmt"
	"sort"
	"sync"
)

// HashRing 一致性哈希环
type HashRing struct {
	ring     map[uint32]string
	sortedKeys []uint32
	mu       sync.RWMutex
	replicas int // 虚拟节点数量
}

// NewHashRing 创建一致性哈希环
func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		ring:     make(map[uint32]string),
		sortedKeys: make([]uint32, 0),
		replicas: replicas,
	}
}

// AddServer 添加服务器
func (hr *HashRing) AddServer(server string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	for i := 0; i < hr.replicas; i++ {
		key := hr.hashKey(fmt.Sprintf("%s:%d", server, i))
		hr.ring[key] = server
		hr.sortedKeys = append(hr.sortedKeys, key)
	}

	sort.Slice(hr.sortedKeys, func(i, j int) bool {
		return hr.sortedKeys[i] < hr.sortedKeys[j]
	})
}

// RemoveServer 移除服务器
func (hr *HashRing) RemoveServer(server string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	for i := 0; i < hr.replicas; i++ {
		key := hr.hashKey(fmt.Sprintf("%s:%d", server, i))
		delete(hr.ring, key)

		// 从排序的键中移除
		for j, k := range hr.sortedKeys {
			if k == key {
				hr.sortedKeys = append(hr.sortedKeys[:j], hr.sortedKeys[j+1:]...)
				break
			}
		}
	}
}

// GetServer 根据键获取服务器
func (hr *HashRing) GetServer(key string) string {
	hr.mu.RLock()
	defer hr.mu.RUnlock()

	if len(hr.sortedKeys) == 0 {
		return ""
	}

	hash := hr.hashKey(key)

	// 找到第一个大于等于 hash 的键
	idx := sort.Search(len(hr.sortedKeys), func(i int) bool {
		return hr.sortedKeys[i] >= hash
	})

	// 如果没找到，使用第一个（环形）
	if idx == len(hr.sortedKeys) {
		idx = 0
	}

	return hr.ring[hr.sortedKeys[idx]]
}

// hashKey 计算键的哈希值
func (hr *HashRing) hashKey(key string) uint32 {
	h := md5.Sum([]byte(key))
	return uint32(h[0])<<24 | uint32(h[1])<<16 | uint32(h[2])<<8 | uint32(h[3])
}

func main() {
	ring := NewHashRing(150) // 每个服务器 150 个虚拟节点

	servers := []string{"server1:8080", "server2:8080", "server3:8080"}
	for _, server := range servers {
		ring.AddServer(server)
	}

	fmt.Println("一致性哈希负载均衡示例:")
	keys := []string{"user1", "user2", "user3", "user4", "user5"}
	for _, key := range keys {
		server := ring.GetServer(key)
		fmt.Printf("键 %s -> %s\n", key, server)
	}

	// 移除一个服务器
	fmt.Println("\n移除 server1 后:")
	ring.RemoveServer("server1:8080")
	for _, key := range keys {
		server := ring.GetServer(key)
		fmt.Printf("键 %s -> %s\n", key, server)
	}
}

