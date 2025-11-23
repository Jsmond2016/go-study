---
title: Redis 缓存
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["基础语法", "数据库基础"]
tags: ["Redis", "缓存", "消息队列", "go-redis"]
---

# Redis 缓存

Redis 是一个高性能的内存数据库，常用于缓存、消息队列、会话存储等场景。

## 📋 学习目标

- [ ] 理解 Redis 的概念和用途
- [ ] 掌握 go-redis 的使用
- [ ] 学会基本的缓存操作
- [ ] 理解 Redis 数据结构
- [ ] 掌握发布订阅功能
- [ ] 了解连接池和性能优化

## 🎯 Redis 简介

### 为什么使用 Redis

- **高性能**: 内存存储，读写速度快
- **丰富的数据结构**: String、Hash、List、Set、ZSet
- **持久化**: 支持 RDB 和 AOF
- **分布式**: 支持集群和哨兵模式
- **多功能**: 缓存、消息队列、计数器等

### 安装 go-redis

```bash
go get github.com/redis/go-redis/v9
```

## 🚀 快速开始

### 基本连接

```go
package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 无密码
		DB:       0,  // 默认数据库
	})
	
	ctx := context.Background()
	
	// 测试连接
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	
	defer rdb.Close()
}
```

## 📊 基本操作

### String 操作

```go
// 设置值
err := rdb.Set(ctx, "key", "value", 0).Err()

// 设置值（带过期时间）
err := rdb.Set(ctx, "key", "value", time.Hour).Err()

// 获取值
val, err := rdb.Get(ctx, "key").Result()

// 获取值（如果不存在返回错误）
val, err := rdb.Get(ctx, "key").Result()
if err == redis.Nil {
	fmt.Println("key 不存在")
}

// 删除键
err := rdb.Del(ctx, "key").Err()

// 检查键是否存在
exists, err := rdb.Exists(ctx, "key").Result()

// 设置过期时间
err := rdb.Expire(ctx, "key", time.Hour).Err()

// 自增
val, err := rdb.Incr(ctx, "counter").Result()

// 自增指定值
val, err := rdb.IncrBy(ctx, "counter", 10).Result()
```

### Hash 操作

```go
// 设置字段
err := rdb.HSet(ctx, "user:1", "name", "张三", "age", 25).Err()

// 获取字段
name, err := rdb.HGet(ctx, "user:1", "name").Result()

// 获取所有字段
user, err := rdb.HGetAll(ctx, "user:1").Result()

// 删除字段
err := rdb.HDel(ctx, "user:1", "age").Err()

// 检查字段是否存在
exists, err := rdb.HExists(ctx, "user:1", "name").Result()

// 获取所有字段名
keys, err := rdb.HKeys(ctx, "user:1").Result()

// 获取所有字段值
vals, err := rdb.HVals(ctx, "user:1").Result()
```

### List 操作

```go
// 从左侧推入
err := rdb.LPush(ctx, "list", "value1", "value2").Err()

// 从右侧推入
err := rdb.RPush(ctx, "list", "value3", "value4").Err()

// 从左侧弹出
val, err := rdb.LPop(ctx, "list").Result()

// 从右侧弹出
val, err := rdb.RPop(ctx, "list").Result()

// 获取列表长度
length, err := rdb.LLen(ctx, "list").Result()

// 获取指定范围元素
vals, err := rdb.LRange(ctx, "list", 0, -1).Result()

// 根据索引获取元素
val, err := rdb.LIndex(ctx, "list", 0).Result()
```

### Set 操作

```go
// 添加成员
err := rdb.SAdd(ctx, "set", "member1", "member2").Err()

// 获取所有成员
members, err := rdb.SMembers(ctx, "set").Result()

// 检查成员是否存在
exists, err := rdb.SIsMember(ctx, "set", "member1").Result()

// 获取成员数量
count, err := rdb.SCard(ctx, "set").Result()

// 删除成员
err := rdb.SRem(ctx, "set", "member1").Err()

// 集合运算
// 交集
intersect, err := rdb.SInter(ctx, "set1", "set2").Result()

// 并集
union, err := rdb.SUnion(ctx, "set1", "set2").Result()

// 差集
diff, err := rdb.SDiff(ctx, "set1", "set2").Result()
```

### Sorted Set 操作

```go
// 添加成员（带分数）
err := rdb.ZAdd(ctx, "zset", redis.Z{
	Score:  100,
	Member: "member1",
}).Err()

// 获取排名
rank, err := rdb.ZRank(ctx, "zset", "member1").Result()

// 获取分数
score, err := rdb.ZScore(ctx, "zset", "member1").Result()

// 获取指定范围的成员（按分数排序）
members, err := rdb.ZRange(ctx, "zset", 0, -1).Result()

// 获取指定分数范围的成员
members, err := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{
	Min: "0",
	Max: "100",
}).Result()
```

## 🔄 发布订阅

### 发布消息

```go
// 发布消息到频道
err := rdb.Publish(ctx, "channel", "message").Err()
```

### 订阅消息

```go
pubsub := rdb.Subscribe(ctx, "channel")
defer pubsub.Close()

// 接收消息
ch := pubsub.Channel()
for msg := range ch {
	fmt.Printf("收到消息: %s\n", msg.Payload)
}
```

### 模式订阅

```go
// 订阅匹配模式的频道
pubsub := rdb.PSubscribe(ctx, "channel:*")
defer pubsub.Close()

ch := pubsub.Channel()
for msg := range ch {
	fmt.Printf("频道: %s, 消息: %s\n", msg.Channel, msg.Payload)
}
```

## 🏃‍♂️ 实践应用

### 缓存服务封装

```go
package cache

import (
	"context"
	"encoding/json"
	"time"
	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService(addr, password string, db int) *CacheService {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	
	return &CacheService{
		client: client,
		ctx:    context.Background(),
	}
}

// Set 设置缓存
func (c *CacheService) Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, data, expiration).Err()
}

// Get 获取缓存
func (c *CacheService) Get(key string, dest interface{}) error {
	data, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

// Delete 删除缓存
func (c *CacheService) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

// Exists 检查键是否存在
func (c *CacheService) Exists(key string) (bool, error) {
	count, err := c.client.Exists(c.ctx, key).Result()
	return count > 0, err
}

// 使用示例
func Example() {
	cache := NewCacheService("localhost:6379", "", 0)
	
	// 设置缓存
	cache.Set("user:1", map[string]interface{}{
		"id":   1,
		"name": "张三",
	}, time.Hour)
	
	// 获取缓存
	var user map[string]interface{}
	cache.Get("user:1", &user)
}
```

### 分布式锁

```go
package cache

import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
)

type DistributedLock struct {
	client *redis.Client
	ctx    context.Context
}

func NewDistributedLock(client *redis.Client) *DistributedLock {
	return &DistributedLock{
		client: client,
		ctx:    context.Background(),
	}
}

// Lock 获取锁
func (dl *DistributedLock) Lock(key string, expiration time.Duration) (bool, error) {
	result, err := dl.client.SetNX(dl.ctx, key, "locked", expiration).Result()
	return result, err
}

// Unlock 释放锁
func (dl *DistributedLock) Unlock(key string) error {
	return dl.client.Del(dl.ctx, key).Err()
}

// 使用示例
func ExampleLock() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	lock := NewDistributedLock(client)
	
	// 获取锁
	acquired, err := lock.Lock("resource:1", time.Minute)
	if err != nil {
		return err
	}
	
	if !acquired {
		return fmt.Errorf("获取锁失败")
	}
	
	defer lock.Unlock("resource:1")
	
	// 执行需要加锁的操作
	// ...
}
```

## ⚡ 性能优化

### 连接池配置

```go
rdb := redis.NewClient(&redis.Options{
	Addr:         "localhost:6379",
	Password:     "",
	DB:           0,
	PoolSize:     10,              // 连接池大小
	MinIdleConns: 5,               // 最小空闲连接
	MaxRetries:   3,                // 最大重试次数
	DialTimeout:  5 * time.Second, // 连接超时
	ReadTimeout:  3 * time.Second,  // 读取超时
	WriteTimeout: 3 * time.Second, // 写入超时
})
```

### Pipeline 批量操作

```go
pipe := rdb.Pipeline()

pipe.Set(ctx, "key1", "value1", 0)
pipe.Set(ctx, "key2", "value2", 0)
pipe.Get(ctx, "key1")
pipe.Get(ctx, "key2")

cmds, err := pipe.Exec(ctx)
if err != nil {
	panic(err)
}

// 处理结果
for _, cmd := range cmds {
	fmt.Println(cmd.(*redis.StringCmd).Val())
}
```

## ⚠️ 注意事项

### 1. 错误处理

```go
// ✅ 检查 Redis.Nil
val, err := rdb.Get(ctx, "key").Result()
if err == redis.Nil {
	// 键不存在
} else if err != nil {
	// 其他错误
}
```

### 2. 连接管理

```go
// ✅ 使用连接池
// ✅ 设置合理的超时时间
// ✅ 处理连接错误
```

### 3. 数据序列化

```go
// ✅ 使用 JSON 序列化复杂对象
// ✅ 考虑序列化性能
// ✅ 注意数据大小限制
```

## 📚 扩展阅读

- [go-redis 官方文档](https://github.com/redis/go-redis)
- [Redis 官方文档](https://redis.io/docs/)

## ⏭️ 下一章节

[Docker 基础](./12-docker.md) → 学习容器化部署

---

**💡 提示**: Redis 是企业级项目中必不可少的缓存工具，掌握它对于提高系统性能非常重要！

