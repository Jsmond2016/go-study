// Package main 展示 Gin 框架的路由功能
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. 基本路由
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "GET 方法")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "POST 方法")
	})

	r.PUT("/put", func(c *gin.Context) {
		c.String(200, "PUT 方法")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "DELETE 方法")
	})

	// 2. 路径参数
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"user_id": id})
	})

	// 3. 查询参数
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		page := c.DefaultQuery("page", "1")
		c.JSON(200, gin.H{
			"query": query,
			"page":  page,
		})
	})

	// 4. 路由组
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "获取用户列表"})
		})
		v1.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "创建用户"})
		})
	}

	// 5. 处理 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "路由不存在"})
	})

	r.Run(":8080")
}

