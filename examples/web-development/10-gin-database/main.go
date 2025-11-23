// Package main 展示 Gin 框架与数据库的集成
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移
	db.AutoMigrate(&User{})
}

func main() {
	r := gin.Default()

	// 创建用户
	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"data": user})
	})

	// 获取用户列表
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, gin.H{"data": users})
	})

	// 获取单个用户
	r.GET("/users/:id", func(c *gin.Context) {
		var user User
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})

	// 更新用户
	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Save(&user)
		c.JSON(200, gin.H{"data": user})
	})

	// 删除用户
	r.DELETE("/users/:id", func(c *gin.Context) {
		if err := db.Delete(&User{}, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(200, gin.H{"message": "用户已删除"})
	})

	r.Run(":8080")
}

