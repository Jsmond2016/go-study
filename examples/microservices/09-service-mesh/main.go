package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 服务版本和名称（从环境变量读取，Istio 场景中常用）
var (
	serviceName    = getEnv("SERVICE_NAME", "demo-service")
	serviceVersion = getEnv("SERVICE_VERSION", "v1")
	servicePort    = getEnv("PORT", "8080")
)

// Prometheus 指标
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status", "version"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latencies in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "version"},
	)
)

func init() {
	// 注册 Prometheus 指标
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func main() {
	// 创建 Gin 路由
	r := gin.Default()

	// 注册中间件
	r.Use(prometheusMiddleware())
	r.Use(istioHeadersMiddleware())

	// 健康检查端点（Kubernetes liveness probe）
	r.GET("/health", healthHandler)

	// 就绪检查端点（Kubernetes readiness probe）
	r.GET("/ready", readyHandler)

	// Prometheus 指标端点
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 业务 API 端点
	r.GET("/api/hello", helloHandler)
	r.GET("/api/version", versionHandler)
	r.GET("/api/call-service", callServiceHandler)

	// 模拟不同版本的行为
	r.GET("/api/feature", featureHandler)

	log.Printf("Starting %s %s on port %s", serviceName, serviceVersion, servicePort)
	log.Printf("Health check: http://localhost:%s/health", servicePort)
	log.Printf("Metrics: http://localhost:%s/metrics", servicePort)

	if err := r.Run(":" + servicePort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// prometheusMiddleware 记录 HTTP 请求指标
func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录指标
		duration := time.Since(start).Seconds()
		status := fmt.Sprintf("%d", c.Writer.Status())

		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			status,
			serviceVersion,
		).Inc()

		httpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			serviceVersion,
		).Observe(duration)
	}
}

// istioHeadersMiddleware 传播 Istio 追踪头
// Istio 需要应用程序传播这些 header 来实现分布式追踪
func istioHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Istio 追踪需要的 header
		tracingHeaders := []string{
			"x-request-id",
			"x-b3-traceid",
			"x-b3-spanid",
			"x-b3-parentspanid",
			"x-b3-sampled",
			"x-b3-flags",
			"x-ot-span-context",
		}

		// 保存到 context，供下游调用使用
		headers := make(map[string]string)
		for _, h := range tracingHeaders {
			if value := c.GetHeader(h); value != "" {
				headers[h] = value
			}
		}

		c.Set("tracing-headers", headers)
		c.Next()
	}
}

// healthHandler 健康检查（liveness probe）
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": serviceName,
		"version": serviceVersion,
	})
}

// readyHandler 就绪检查（readiness probe）
func readyHandler(c *gin.Context) {
	// 这里可以检查依赖服务的健康状态
	// 示例：检查数据库连接、缓存连接等
	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"service": serviceName,
		"version": serviceVersion,
	})
}

// helloHandler 简单的问候接口
func helloHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello, %s!", name),
		"service": serviceName,
		"version": serviceVersion,
		"time":    time.Now().Format(time.RFC3339),
	})
}

// versionHandler 返回服务版本信息
func versionHandler(c *gin.Context) {
	// 读取 Istio 注入的环境变量
	podName := getEnv("HOSTNAME", "unknown")
	podIP := getEnv("POD_IP", "unknown")
	nodeName := getEnv("NODE_NAME", "unknown")

	c.JSON(http.StatusOK, gin.H{
		"service": serviceName,
		"version": serviceVersion,
		"pod": gin.H{
			"name": podName,
			"ip":   podIP,
			"node": nodeName,
		},
		"istio": gin.H{
			"sidecar": c.GetHeader("x-envoy-peer-metadata") != "",
		},
	})
}

// callServiceHandler 调用下游服务（演示服务间调用）
func callServiceHandler(c *gin.Context) {
	// 获取要调用的服务地址
	targetService := c.DefaultQuery("target", "http://demo-service:8080")

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", targetService+"/api/hello", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 传播 Istio 追踪 header（重要！）
	if headers, exists := c.Get("tracing-headers"); exists {
		for k, v := range headers.(map[string]string) {
			req.Header.Set(k, v)
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to call service: %v", err),
		})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"message":       "Service call successful",
		"target":        targetService,
		"status":        resp.Status,
		"from_service":  serviceName,
		"from_version":  serviceVersion,
		"tracing_id":    c.GetHeader("x-request-id"),
	})
}

// featureHandler 模拟不同版本的功能差异
func featureHandler(c *gin.Context) {
	var response gin.H

	// 不同版本返回不同的响应（用于演示 Istio 流量管理）
	switch serviceVersion {
	case "v1":
		response = gin.H{
			"feature": "basic",
			"message": "This is version 1 with basic features",
			"color":   "blue",
		}
	case "v2":
		response = gin.H{
			"feature": "enhanced",
			"message": "This is version 2 with enhanced features",
			"color":   "green",
			"extra":   "Advanced feature enabled",
		}
	case "v3":
		response = gin.H{
			"feature": "premium",
			"message": "This is version 3 with premium features",
			"color":   "gold",
			"extra":   "Premium features enabled",
			"bonus":   "Beta feature available",
		}
	default:
		response = gin.H{
			"feature": "unknown",
			"message": "Unknown version",
		}
	}

	response["version"] = serviceVersion
	response["service"] = serviceName

	c.JSON(http.StatusOK, response)
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
