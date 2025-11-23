// Package main 实现一个完整的博客系统
package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

// Article 文章模型
type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:"type:text"`
	AuthorID  uint      `json:"author_id"`
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`
	ViewCount int       `json:"view_count" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comment 评论模型
type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	ArticleID uint      `json:"article_id"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleID"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移
	db.AutoMigrate(&User{}, &Article{}, &Comment{})
}

func main() {
	r := gin.Default()

	// 用户注册
	r.POST("/api/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 简单密码处理（实际应使用 bcrypt）
		user.Password = user.Password // 实际应加密
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "注册成功", "user": user})
	})

	// 创建文章
	r.POST("/api/articles", func(c *gin.Context) {
		var article Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		article.AuthorID = 1 // 简化处理，实际应从 JWT 获取
		if err := db.Create(&article).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Author").First(&article, article.ID)
		c.JSON(http.StatusCreated, gin.H{"data": article})
	})

	// 获取文章列表
	r.GET("/api/articles", func(c *gin.Context) {
		var articles []Article
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
		offset := (page - 1) * pageSize

		db.Preload("Author").Offset(offset).Limit(pageSize).Find(&articles)
		var total int64
		db.Model(&Article{}).Count(&total)

		c.JSON(http.StatusOK, gin.H{
			"data": articles,
			"page": page,
			"page_size": pageSize,
			"total": total,
		})
	})

	// 获取文章详情
	r.GET("/api/articles/:id", func(c *gin.Context) {
		var article Article
		id := c.Param("id")
		if err := db.Preload("Author").First(&article, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
			return
		}
		// 增加阅读量
		db.Model(&article).Update("view_count", article.ViewCount+1)
		c.JSON(http.StatusOK, gin.H{"data": article})
	})

	// 添加评论
	r.POST("/api/articles/:id/comments", func(c *gin.Context) {
		var comment Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		articleID, _ := strconv.Atoi(c.Param("id"))
		comment.ArticleID = uint(articleID)
		comment.UserID = 1 // 简化处理
		if err := db.Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("User").First(&comment, comment.ID)
		c.JSON(http.StatusCreated, gin.H{"data": comment})
	})

	// 获取文章评论
	r.GET("/api/articles/:id/comments", func(c *gin.Context) {
		var comments []Comment
		articleID := c.Param("id")
		db.Preload("User").Where("article_id = ?", articleID).Find(&comments)
		c.JSON(http.StatusOK, gin.H{"data": comments})
	})

	log.Println("博客系统启动在 :8080")
	r.Run(":8080")
}

