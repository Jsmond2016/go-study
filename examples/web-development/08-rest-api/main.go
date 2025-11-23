// Package main 展示 RESTful API 的实现
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User 用户结构体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 模拟数据库
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}

// 获取用户列表
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"total": len(users),
	})
}

// 获取单个用户
func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": user})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
}

// 创建用户
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// 更新用户
func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{"data": updatedUser})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
}

// 删除用户
func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "用户已删除"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
}

func main() {
	r := gin.Default()

	// RESTful API 路由
	api := r.Group("/api/v1")
	{
		api.GET("/users", getUsers)       // GET /api/v1/users
		api.GET("/users/:id", getUser)    // GET /api/v1/users/:id
		api.POST("/users", createUser)    // POST /api/v1/users
		api.PUT("/users/:id", updateUser) // PUT /api/v1/users/:id
		api.DELETE("/users/:id", deleteUser) // DELETE /api/v1/users/:id
	}

	r.Run(":8080")
}

