---
title: API 网关
difficulty: advanced
duration: "5-6小时"
prerequisites: ["gRPC 基础", "服务发现", "负载均衡"]
tags: ["API 网关", "网关", "路由", "认证", "限流", "熔断"]
---

# API 网关

API 网关是微服务架构中的入口点，它提供了统一的 API 接口，处理路由、认证、限流、监控等横切关注点。

## 📋 学习目标

完成本教程后，你将能够：

- [ ] 理解 API 网关的概念和架构
- [ ] 实现路由配置和管理
- [ ] 实现请求转发和代理
- [ ] 集成认证和授权
- [ ] 实现限流和熔断机制
- [ ] 实现请求日志和监控
- [ ] 管理 API 版本
- [ ] 优化网关性能

## 🎯 API 网关简介

### 什么是 API 网关

API 网关是微服务架构中的单一入口点，它充当客户端和后端服务之间的中间层，提供：

- **统一入口**：所有客户端请求通过网关
- **路由转发**：将请求路由到相应的后端服务
- **横切关注点**：认证、授权、限流、日志等
- **协议转换**：HTTP 到 gRPC 等
- **聚合服务**：组合多个服务响应

### 为什么需要 API 网关

- **简化客户端**：客户端只需知道网关地址
- **解耦**：后端服务变更不影响客户端
- **安全**：集中处理认证和授权
- **监控**：集中收集日志和指标
- **限流**：保护后端服务

### API 网关架构

```
客户端 → API 网关 → 路由 → 认证 → 限流 → 转发 → 后端服务
                ↓
            监控/日志
```

## 🚀 基础实现

### 简单的 HTTP 网关

```go
package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Gateway struct {
	routes map[string]*url.URL
	proxy  *httputil.ReverseProxy
}

func NewGateway() *Gateway {
	return &Gateway{
		routes: make(map[string]*url.URL),
	}
}

func (g *Gateway) RegisterRoute(path string, targetURL string) error {
	u, err := url.Parse(targetURL)
	if err != nil {
		return err
	}
	g.routes[path] = u
	return nil
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 查找路由
	target, ok := g.routes[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 修改请求
	r.URL.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = target.Host

	// 转发请求
	proxy.ServeHTTP(w, r)
}

func main() {
	gateway := NewGateway()

	// 注册路由
	gateway.RegisterRoute("/api/users", "http://localhost:8081")
	gateway.RegisterRoute("/api/orders", "http://localhost:8082")
	gateway.RegisterRoute("/api/products", "http://localhost:8083")

	log.Println("Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", gateway))
}
```

### 使用 Gin 框架

```go
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Gateway struct {
	routes map[string]*url.URL
}

func NewGateway() *Gateway {
	return &Gateway{
		routes: make(map[string]*url.URL),
	}
}

func (g *Gateway) RegisterRoute(path, target string) error {
	u, err := url.Parse(target)
	if err != nil {
		return err
	}
	g.routes[path] = u
	return nil
}

func (g *Gateway) Proxy(c *gin.Context) {
	target, ok := g.routes[c.Request.URL.Path]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()

	gateway := NewGateway()
	gateway.RegisterRoute("/api/users", "http://localhost:8081")
	gateway.RegisterRoute("/api/orders", "http://localhost:8082")

	r.Any("/api/*path", gateway.Proxy)

	r.Run(":8080")
}
```

## 🛣️ 路由管理

### 动态路由配置

```go
package main

import (
	"encoding/json"
	"sync"
)

type Route struct {
	Path   string `json:"path"`
	Target string `json:"target"`
	Method string `json:"method"`
}

type RouteManager struct {
	routes map[string]*Route
	mu     sync.RWMutex
}

func NewRouteManager() *RouteManager {
	return &RouteManager{
		routes: make(map[string]*Route),
	}
}

func (rm *RouteManager) AddRoute(route *Route) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.routes[route.Path] = route
}

func (rm *RouteManager) RemoveRoute(path string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.routes, path)
}

func (rm *RouteManager) GetRoute(path string) (*Route, bool) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	route, ok := rm.routes[path]
	return route, ok
}

func (rm *RouteManager) ListRoutes() []*Route {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	routes := make([]*Route, 0, len(rm.routes))
	for _, route := range rm.routes {
		routes = append(routes, route)
	}
	return routes
}

// 从配置文件加载路由
func (rm *RouteManager) LoadFromConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var routes []*Route
	if err := json.Unmarshal(data, &routes); err != nil {
		return err
	}

	for _, route := range routes {
		rm.AddRoute(route)
	}

	return nil
}
```

### 路由匹配

```go
import (
	"regexp"
	"strings"
)

type RouteMatcher struct {
	pattern *regexp.Regexp
	target  string
}

func NewRouteMatcher(pattern, target string) (*RouteMatcher, error) {
	// 将路径模式转换为正则表达式
	regexPattern := strings.ReplaceAll(pattern, "*", ".*")
	regexPattern = "^" + regexPattern + "$"

	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, err
	}

	return &RouteMatcher{
		pattern: re,
		target:  target,
	}, nil
}

func (rm *RouteMatcher) Match(path string) bool {
	return rm.pattern.MatchString(path)
}
```

## 🔐 认证和授权

### JWT 认证中间件

```go
package main

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// 提取 token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 解析 token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 将用户信息存储到 context
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(403, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		roleStr := role.(string)
		allowed := false
		for _, r := range roles {
			if r == roleStr {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(403, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 公开路由
	r.POST("/api/login", loginHandler)

	// 需要认证的路由
	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/users", getUserHandler)

		// 需要特定角色的路由
		admin := api.Group("/admin")
		admin.Use(RoleMiddleware("admin"))
		{
			admin.DELETE("/users/:id", deleteUserHandler)
		}
	}

	r.Run(":8080")
}
```

## 🚦 限流

### 令牌桶算法

```go
package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity    int64
	tokens      int64
	refillRate  int64 // tokens per second
	lastRefill  time.Time
	mu          sync.Mutex
}

func NewTokenBucket(capacity, refillRate int64) *TokenBucket {
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

	// 补充令牌
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tokensToAdd := int64(elapsed * float64(tb.refillRate))

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

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
```

### 限流中间件

```go
func RateLimitMiddleware(tb *TokenBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !tb.Allow() {
			c.JSON(429, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 为不同路径设置不同的限流
	userLimiter := NewTokenBucket(100, 10) // 100 tokens, 10 per second
	apiLimiter := NewTokenBucket(1000, 100) // 1000 tokens, 100 per second

	r.Use(RateLimitMiddleware(apiLimiter))

	users := r.Group("/api/users")
	users.Use(RateLimitMiddleware(userLimiter))
	{
		users.GET("/", getUserHandler)
	}

	r.Run(":8080")
}
```

### 漏桶算法

```go
type LeakyBucket struct {
	capacity    int64
	water       int64
	leakRate    int64 // water per second
	lastLeak    time.Time
	mu          sync.Mutex
}

func NewLeakyBucket(capacity, leakRate int64) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// 漏水
	now := time.Now()
	elapsed := now.Sub(lb.lastLeak).Seconds()
	waterToLeak := int64(elapsed * float64(lb.leakRate))

	if waterToLeak > 0 {
		lb.water = max(0, lb.water-waterToLeak)
		lb.lastLeak = now
	}

	// 检查是否有空间
	if lb.water < lb.capacity {
		lb.water++
		return true
	}

	return false
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
```

## ⚡ 熔断器

### 熔断器实现

```go
package main

import (
	"sync"
	"time"
)

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	maxFailures    int
	resetTimeout   time.Duration
	failures       int
	state          CircuitState
	lastFailureTime time.Time
	mu             sync.Mutex
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()

	// 检查状态
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
		} else {
			cb.mu.Unlock()
			return fmt.Errorf("circuit breaker is open")
		}
	}

	cb.mu.Unlock()

	// 执行函数
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failures++
		cb.lastFailureTime = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
		}
		return err
	}

	// 成功
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failures = 0
	} else {
		cb.failures = 0
	}

	return nil
}
```

### 熔断器中间件

```go
func CircuitBreakerMiddleware(cb *CircuitBreaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := cb.Call(func() error {
			c.Next()

			// 检查响应状态码
			if c.Writer.Status() >= 500 {
				return fmt.Errorf("server error: %d", c.Writer.Status())
			}
			return nil
		})

		if err != nil {
			c.JSON(503, gin.H{"error": "service unavailable"})
			c.Abort()
		}
	}
}
```

## 📊 监控和日志

### 请求日志中间件

```go
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[%s] %s %s %d %v",
			method,
			path,
			c.ClientIP(),
			status,
			latency,
		)
	}
}
```

### 指标收集

```go
type Metrics struct {
	requests    int64
	errors      int64
	latency     time.Duration
	mu          sync.RWMutex
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) RecordRequest(latency time.Duration, isError bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.requests++
	if isError {
		m.errors++
	}
	m.latency = latency
}

func (m *Metrics) GetStats() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return map[string]interface{}{
		"requests": m.requests,
		"errors":   m.errors,
		"error_rate": float64(m.errors) / float64(m.requests),
		"latency":  m.latency.String(),
	}
}

func MetricsMiddleware(metrics *Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		isError := c.Writer.Status() >= 400
		metrics.RecordRequest(latency, isError)
	}
}
```

### 健康检查端点

```go
func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
		"timestamp": time.Now().Unix(),
	})
}

func MetricsHandler(metrics *Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, metrics.GetStats())
	}
}
```

## 🔄 gRPC 网关

### HTTP 到 gRPC 转换

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	pb "your-project/proto"
)

type GRPCGateway struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

func NewGRPCGateway(target string) (*GRPCGateway, error) {
	conn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &GRPCGateway{
		conn:   conn,
		client: pb.NewUserServiceClient(conn),
	}, nil
}

func (gw *GRPCGateway) GetUser(c *gin.Context) {
	var req pb.GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := gw.client.GetUser(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 转换为 JSON
	data, err := protojson.Marshal(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Data(200, "application/json", data)
}

func main() {
	gw, err := NewGRPCGateway("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer gw.conn.Close()

	r := gin.Default()
	r.POST("/api/users", gw.GetUser)

	r.Run(":8080")
}
```

## 💡 最佳实践

### 1. 路由配置

- 使用配置文件管理路由
- 支持动态更新路由
- 实现路由优先级

### 2. 认证授权

- 集中处理认证
- 支持多种认证方式
- 实现细粒度授权

### 3. 限流策略

- 不同路径不同限流
- 基于用户/IP 限流
- 实现限流降级

### 4. 监控告警

- 收集关键指标
- 设置告警阈值
- 实现可视化

### 5. 性能优化

- 使用连接池
- 实现缓存
- 异步处理

## 📝 实践练习

1. **基础练习**：实现一个简单的 HTTP 网关
2. **路由练习**：实现动态路由配置
3. **认证练习**：集成 JWT 认证
4. **限流练习**：实现令牌桶和漏桶算法
5. **综合练习**：构建一个完整的 API 网关

## 🔗 相关资源

- [Kong API 网关](https://konghq.com/kong/)
- [Envoy 代理](https://www.envoyproxy.io/)
- [代码示例](../../examples/microservices/05-api-gateway/)

## ⏭️ 下一步

完成 API 网关学习后，你已经掌握了微服务架构的核心组件。可以：

- 构建完整的微服务项目
- 学习服务网格
- 深入学习分布式系统

---

**🎉 恭喜！** 你已经掌握了 API 网关的核心知识。现在可以构建完整的微服务架构了！

