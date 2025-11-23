// Package main 展示 Gin 框架的中间件功能
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志中间件
func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		// 调用下一个处理器
		c.Next()

		// 请求后处理
		latency := time.Since(start)
		fmt.Printf("[%s] %s %s %v\n",
			c.Request.Method,
			path,
			c.Writer.Status(),
			latency,
		)
	}
}

// 认证中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "未授权"})
			c.Abort()
			return
		}
		// 这里可以验证 token
		c.Set("user_id", "123")
		c.Next()
	}
}

// 错误处理中间件
func errorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(500, gin.H{
			"error": "服务器内部错误",
		})
		c.Abort()
	})
}

func main() {
	r := gin.Default()

	// 全局中间件
	r.Use(loggingMiddleware())
	r.Use(errorHandlerMiddleware())

	// 公开路由
	r.GET("/public", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "公开接口"})
	})

	// 需要认证的路由组
	api := r.Group("/api")
	api.Use(authMiddleware())
	{
		api.GET("/user", func(c *gin.Context) {
			userID := c.MustGet("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	r.Run(":8080")
}

