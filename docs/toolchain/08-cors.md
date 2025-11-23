---
title: CORS 跨域
difficulty: beginner
duration: "2-3小时"
prerequisites: ["Web 开发", "HTTP"]
tags: ["CORS", "跨域", "HTTP", "安全"]
---

# CORS 跨域

CORS (Cross-Origin Resource Sharing) 是一种机制，允许 Web 应用从不同域访问资源。

## 📋 学习目标

- [ ] 理解跨域的概念和原因
- [ ] 掌握 CORS 的工作原理
- [ ] 学会在 Go 中实现 CORS
- [ ] 理解预检请求
- [ ] 掌握安全配置
- [ ] 了解常见问题

## 🎯 CORS 简介

### 什么是跨域

跨域是指浏览器从一个域名的网页去请求另一个域名的资源。同源策略限制了这种行为。

### 同源策略

同源需要满足：
- 协议相同 (http/https)
- 域名相同
- 端口相同

### 为什么需要 CORS

- **前后端分离**: 前端和后端可能部署在不同域名
- **微服务架构**: 不同服务可能在不同域名
- **第三方 API**: 需要允许跨域访问

## 🚀 快速开始

### 基本 CORS 配置

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	
	// 配置 CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	
	r.Use(cors.New(config))
	
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello CORS"})
	})
	
	r.Run(":8080")
}
```

### 允许所有来源

```go
config := cors.DefaultConfig()
config.AllowAllOrigins = true
r.Use(cors.New(config))
```

## 🔧 详细配置

### 完整配置示例

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	r := gin.Default()
	
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	
	r.Use(cors.New(config))
	
	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"users": []string{"user1", "user2"}})
	})
	
	r.Run(":8080")
}
```

### 自定义 CORS 中间件

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		
		// 允许的源列表
		allowedOrigins := []string{
			"http://localhost:3000",
			"https://example.com",
		}
		
		// 检查源是否允许
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}
		
		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Max-Age", "86400")
		}
		
		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	
	r.Run(":8080")
}
```

## 🔍 预检请求

### 什么是预检请求

对于复杂请求，浏览器会先发送 OPTIONS 请求进行预检。

### 处理预检请求

```go
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 CORS 头
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		
		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	}
}
```

## 🏃‍♂️ 实践应用

### 环境相关配置

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"os"
)

func getCORSConfig() cors.Config {
	config := cors.DefaultConfig()
	
	// 根据环境设置允许的源
	env := os.Getenv("ENV")
	if env == "production" {
		config.AllowOrigins = []string{
			"https://example.com",
			"https://www.example.com",
		}
	} else {
		config.AllowAllOrigins = true
	}
	
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	
	return config
}

func main() {
	r := gin.Default()
	r.Use(cors.New(getCORSConfig()))
	
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	
	r.Run(":8080")
}
```

## ⚠️ 安全注意事项

### 1. 不要允许所有源

```go
// ❌ 生产环境不要这样做
config.AllowAllOrigins = true

// ✅ 明确指定允许的源
config.AllowOrigins = []string{"https://example.com"}
```

### 2. 凭证安全

```go
// 使用凭证时，不能使用通配符
config.AllowCredentials = true
config.AllowOrigins = []string{"https://example.com"} // 必须明确指定
```

### 3. 方法限制

```go
// ✅ 只允许需要的方法
config.AllowMethods = []string{"GET", "POST"}

// ❌ 不要允许所有方法
config.AllowMethods = []string{"*"}
```

## 📚 扩展阅读

- [MDN CORS 文档](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS)
- [gin-contrib/cors](https://github.com/gin-contrib/cors)

## ⏭️ 下一章节

[限流与熔断](./09-rate-limit.md) → 学习限流和熔断器

---

**💡 提示**: CORS 是前后端分离开发中必须掌握的知识，正确配置可以确保应用安全运行！

