---
title: 限流与熔断
difficulty: advanced
duration: "3-4小时"
prerequisites: ["Web 开发", "中间件"]
tags: ["限流", "熔断", "保护", "性能"]
---

# 限流与熔断

限流和熔断是保护系统的重要机制，防止系统过载和故障扩散。

## 📋 学习目标

- [ ] 理解限流和熔断的概念
- [ ] 掌握限流算法的实现
- [ ] 学会使用限流中间件
- [ ] 理解熔断器的工作原理
- [ ] 掌握熔断器的实现
- [ ] 了解最佳实践

## 🎯 限流简介

### 为什么需要限流

- **防止过载**: 保护系统不被大量请求压垮
- **公平分配**: 确保资源公平分配
- **防止滥用**: 防止恶意请求
- **成本控制**: 控制 API 调用成本

### 限流算法

- **固定窗口**: 固定时间窗口内的请求数限制
- **滑动窗口**: 滑动时间窗口内的请求数限制
- **令牌桶**: 以固定速率生成令牌
- **漏桶**: 以固定速率处理请求

## 🚀 快速开始

### 使用 go-rate-limiter

```bash
go get github.com/ulule/limiter/v3
go get github.com/ulule/limiter/v3/drivers/store/memory
```

### 基本限流

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limiterMiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
	r := gin.Default()
	
	// 创建限流器：每秒 10 个请求
	rate, _ := limiter.NewRateFromFormatted("10-S")
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	
	// 使用限流中间件
	middleware := limiterMiddleware.NewMiddleware(instance)
	r.Use(middleware)
	
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	
	r.Run(":8080")
}
```

## 🔧 限流实现

### 固定窗口限流

```go
package main

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	limit    int
	window   time.Duration
	count    int
	windowStart time.Time
	mu       sync.Mutex
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:       limit,
		window:      window,
		windowStart: time.Now(),
	}
}

func (l *FixedWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	now := time.Now()
	
	// 如果超过窗口时间，重置计数
	if now.Sub(l.windowStart) >= l.window {
		l.count = 0
		l.windowStart = now
	}
	
	// 检查是否超过限制
	if l.count >= l.limit {
		return false
	}
	
	l.count++
	return true
}
```

### 令牌桶限流

```go
package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int        // 每秒补充的令牌数
	lastRefill   time.Time
	mu           sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	
	// 补充令牌
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}
	
	// 检查是否有可用令牌
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 🔌 熔断器

### 熔断器概念

熔断器是一种保护机制，当服务出现故障时，快速失败，避免故障扩散。

### 熔断器状态

- **关闭 (Closed)**: 正常状态，请求正常通过
- **打开 (Open)**: 故障状态，请求直接失败
- **半开 (Half-Open)**: 尝试恢复，允许少量请求通过

### 基本实现

```go
package main

import (
	"sync"
	"time"
)

type CircuitBreaker struct {
	maxFailures int
	timeout     time.Duration
	failures    int
	lastFailure time.Time
	state       string // "closed", "open", "half-open"
	mu          sync.Mutex
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       "closed",
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// 检查状态
	if cb.state == "open" {
		// 检查是否应该进入半开状态
		if time.Since(cb.lastFailure) >= cb.timeout {
			cb.state = "half-open"
		} else {
			return fmt.Errorf("熔断器打开")
		}
	}
	
	// 执行函数
	err := fn()
	
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		
		if cb.failures >= cb.maxFailures {
			cb.state = "open"
		}
		return err
	}
	
	// 成功，重置
	if cb.state == "half-open" {
		cb.state = "closed"
	}
	cb.failures = 0
	
	return nil
}
```

## 🏃‍♂️ 实践应用

### Gin 限流中间件

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RateLimiter struct {
	limiter *FixedWindowLimiter
}

func NewRateLimiterMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	limiter := NewFixedWindowLimiter(limit, window)
	
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	
	// 应用限流：每分钟 100 个请求
	r.Use(NewRateLimiterMiddleware(100, time.Minute))
	
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	
	r.Run(":8080")
}
```

### 使用 go-resilience

```bash
go get github.com/eapache/go-resilience
```

```go
package main

import (
	"github.com/eapache/go-resilience/breaker"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	// 创建熔断器
	cb := breaker.New(10, 1, time.Minute)
	
	r.GET("/api/data", func(c *gin.Context) {
		err := cb.Run(func() error {
			// 调用外部服务
			return callExternalService()
		})
		
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(200, gin.H{"message": "Success"})
	})
	
	r.Run(":8080")
}
```

## ⚠️ 注意事项

### 1. 限流策略

```go
// ✅ 根据业务需求设置合理的限流值
// API 接口：100 req/min
// 登录接口：5 req/min
// 文件上传：10 req/min
```

### 2. 熔断器配置

```go
// ✅ 合理设置熔断参数
// maxFailures: 连续失败次数
// timeout: 熔断持续时间
// 半开状态：允许少量请求测试
```

### 3. 监控和告警

```go
// ✅ 监控限流和熔断状态
// 记录被限流的请求
// 记录熔断器状态变化
// 设置告警阈值
```

## 📚 扩展阅读

- [限流算法详解](https://en.wikipedia.org/wiki/Rate_limiting)
- [熔断器模式](https://martinfowler.com/bliki/CircuitBreaker.html)
- [go-resilience](https://github.com/eapache/go-resilience)

## ⏭️ 下一阶段

完成开发工具链学习后，可以进入：

- [实战项目](../projects/) - 使用这些工具构建完整项目

---

**💡 提示**: 限流和熔断是生产环境中必须考虑的保护机制，合理使用可以大大提高系统的稳定性！

