// Package main 展示 Gin 框架的模板功能
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("templates/*")

	// 静态文件
	r.Static("/static", "./static")

	// 渲染模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Go 学习笔记",
			"name":   "Gin",
		})
	})

	// 用户列表
	r.GET("/users", func(c *gin.Context) {
		users := []map[string]interface{}{
			{"id": 1, "name": "Alice", "age": 30},
			{"id": 2, "name": "Bob", "age": 25},
			{"id": 3, "name": "Charlie", "age": 35},
		}
		c.HTML(http.StatusOK, "users.html", gin.H{
			"users": users,
		})
	})

	r.Run(":8080")
}

