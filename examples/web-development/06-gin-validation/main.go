// Package main 展示 Gin 框架的数据验证功能
package main

import (
	"github.com/gin-gonic/gin"
)

// User 用户结构体
type User struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=0,lte=150"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 创建用户
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": "验证失败",
				"details": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "用户创建成功",
			"user":    user,
		})
	})

	// 登录
	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": "验证失败",
				"details": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "登录成功",
		})
	})

	r.Run(":8080")
}

