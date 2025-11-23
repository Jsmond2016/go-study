---
title: Gin 中间件
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["Gin 基础", "Gin 路由"]
tags: ["Gin", "中间件", "认证", "日志", "CORS"]
---

# Gin 中间件

中间件是 Gin 框架的核心功能之一，用于在请求处理前后执行通用逻辑。

## 📋 学习目标

- [ ] 理解中间件的概念和作用
- [ ] 掌握中间件的创建和使用
- [ ] 学会实现常用中间件
- [ ] 理解中间件的执行顺序
- [ ] 掌握中间件的最佳实践

## 🎯 中间件基础

### 什么是中间件

中间件是在请求处理前后执行的函数，可以用于：
- 日志记录
- 认证授权
- 错误处理
- 请求验证
- 性能监控

### 基本语法

```go
func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前处理
		fmt.Println("请求前")
		
		// 调用下一个处理器
		c.Next()
		
		// 请求后处理
		fmt.Println("请求后")
	}
}
```

## 🔧 使用中间件

### 全局中间件

```go
r := gin.Default()

// 全局中间件
r.Use(middleware1())
r.Use(middleware2())

r.GET("/", handler)
```

### 路由组中间件

```go
r := gin.Default()

// 路由组中间件
v1 := r.Group("/api/v1")
v1.Use(authMiddleware())
{
	v1.GET("/users", getUsers)
}
```

### 单个路由中间件

```go
r := gin.Default()

r.GET("/protected", authMiddleware(), protectedHandler)
```

## 📝 常用中间件

### 日志中间件

```go
func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
```

### 认证中间件

```go
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "缺少认证令牌"})
			c.Abort()
			return
		}
		
		// 验证 token
		if !isValidToken(token) {
			c.JSON(401, gin.H{"error": "无效的认证令牌"})
			c.Abort()
			return
		}
		
		// 设置用户信息
		c.Set("user_id", 1)
		c.Set("username", "admin")
		
		c.Next()
	}
}
```

### CORS 中间件

```go
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}
```

### 错误处理中间件

```go
func errorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(500, gin.H{
			"error": "服务器内部错误",
			"message": fmt.Sprintf("%v", recovered),
		})
		c.Abort()
	})
}
```

### 请求ID中间件

```go
func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
```

## 🔄 中间件执行顺序

### 执行流程

```go
r := gin.Default()

r.Use(middleware1())
r.Use(middleware2())

r.GET("/", func(c *gin.Context) {
	// 执行顺序:
	// 1. middleware1 请求前
	// 2. middleware2 请求前
	// 3. handler
	// 4. middleware2 请求后
	// 5. middleware1 请求后
})
```

### 使用 c.Next() 和 c.Abort()

```go
func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查条件
		if !checkCondition() {
			c.JSON(403, gin.H{"error": "权限不足"})
			c.Abort() // 停止执行后续中间件和处理器
			return
		}
		
		c.Next() // 继续执行
	}
}
```

## 🏃‍♂️ 实践应用

### 完整的中间件示例

```go
package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	
	// 使用自定义中间件
	r.Use(loggingMiddleware())
	r.Use(requestIDMiddleware())
	r.Use(errorHandlerMiddleware())
	
	// 公开路由
	public := r.Group("/public")
	{
		public.GET("/info", getInfo)
	}
	
	// 受保护路由
	protected := r.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.GET("/users", getUsers)
	}
	
	r.Run(":8080")
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		c.Next()
		
		latency := time.Since(start)
		fmt.Printf("[%s] %s %s %d %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			latency,
		)
	}
}

func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := fmt.Sprintf("%d", time.Now().UnixNano())
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

func errorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(500, gin.H{
			"error": "服务器内部错误",
		})
		c.Abort()
	})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "未授权"})
			c.Abort()
			return
		}
		c.Next()
	}
}
```

## ⚠️ 注意事项

### 1. 中间件顺序

```go
// ✅ 注意中间件的执行顺序
r.Use(middleware1()) // 先执行
r.Use(middleware2()) // 后执行
```

### 2. 使用 c.Abort()

```go
// ✅ 需要停止执行时使用 Abort
if !authorized {
	c.JSON(401, gin.H{"error": "未授权"})
	c.Abort()
	return
}
```

### 3. 性能考虑

```go
// ✅ 避免在中间件中执行耗时操作
func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ❌ 避免：同步数据库查询
		// db.Query(...)
		
		// ✅ 正确：异步处理或缓存
		c.Next()
	}
}
```

## 📚 扩展阅读

- [Gin 中间件文档](https://gin-gonic.com/docs/examples/using-middleware/)
- [中间件最佳实践](https://github.com/gin-gonic/gin#middleware)

## ⏭️ 下一章节

[Gin 模板](./05-gin-template.md) → 学习 Gin 的模板渲染功能

---

**💡 提示**: 中间件是 Web 开发中非常重要的概念，合理使用中间件可以让代码更加模块化和可维护！

