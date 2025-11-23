package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go-study/examples/microservices/05-api-gateway/middleware"
)

var (
	port = flag.Int("port", 8080, "网关端口")
)

func main() {
	flag.Parse()

	// 创建网关
	gateway := NewGateway()

	// 注册路由
	gateway.RegisterRoute("/api/users", "http://localhost:5001")
	gateway.RegisterRoute("/api/orders", "http://localhost:5002")
	gateway.RegisterRoute("/api/products", "http://localhost:5003")

	// 创建路由器
	router := mux.NewRouter()

	// 添加中间件
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.CORSMiddleware)

	// 健康检查端点
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	}).Methods("GET")

	// 监控端点
	router.HandleFunc("/metrics", gateway.MetricsHandler).Methods("GET")

	// 网关路由
	router.PathPrefix("/api/").Handler(gateway)

	// 启动服务器
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("API 网关启动在端口 %d", *port)
	log.Fatal(server.ListenAndServe())
}

// Gateway API 网关
type Gateway struct {
	routes map[string]*url.URL
	stats  *Stats
}

// NewGateway 创建网关
func NewGateway() *Gateway {
	return &Gateway{
		routes: make(map[string]*url.URL),
		stats:  NewStats(),
	}
}

// RegisterRoute 注册路由
func (g *Gateway) RegisterRoute(path string, targetURL string) error {
	u, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("解析目标 URL 失败: %w", err)
	}
	g.routes[path] = u
	log.Printf("注册路由: %s -> %s", path, targetURL)
	return nil
}

// ServeHTTP 实现 http.Handler 接口
func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 查找匹配的路由
	var target *url.URL
	var matchedPath string

	for path, u := range g.routes {
		if len(r.URL.Path) >= len(path) && r.URL.Path[:len(path)] == path {
			target = u
			matchedPath = path
			break
		}
	}

	if target == nil {
		http.NotFound(w, r)
		return
	}

	// 记录请求开始时间
	start := time.Now()

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 修改请求
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = r.URL.Path[len(matchedPath):]
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
	}

	// 错误处理
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("代理错误: %v", err)
		g.stats.IncrementError(matchedPath)
		http.Error(w, "服务不可用", http.StatusBadGateway)
	}

	// 执行代理
	proxy.ServeHTTP(w, r)

	// 记录统计信息
	duration := time.Since(start)
	g.stats.RecordRequest(matchedPath, duration)
}

// Stats 统计信息
type Stats struct {
	requests map[string]int64
	errors   map[string]int64
	duration map[string]time.Duration
	mu       sync.RWMutex
}

// NewStats 创建统计
func NewStats() *Stats {
	return &Stats{
		requests: make(map[string]int64),
		errors:   make(map[string]int64),
		duration: make(map[string]time.Duration),
	}
}

// RecordRequest 记录请求
func (s *Stats) RecordRequest(path string, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requests[path]++
	s.duration[path] += duration
}

// IncrementError 增加错误计数
func (s *Stats) IncrementError(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors[path]++
}

// MetricsHandler 指标处理器
func (g *Gateway) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	g.stats.mu.RLock()
	defer g.stats.mu.RUnlock()

	metrics := make(map[string]interface{})
	for path := range g.stats.requests {
		metrics[path] = map[string]interface{}{
			"requests": g.stats.requests[path],
			"errors":   g.stats.errors[path],
			"avg_duration_ms": g.stats.duration[path].Milliseconds() / max(g.stats.requests[path], 1),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

