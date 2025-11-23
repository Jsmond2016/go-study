---
title: 文章管理
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["用户认证"]
tags: ["文章", "CRUD", "分类", "标签"]
---

# 文章管理

本章节将实现文章的完整管理功能，包括创建、编辑、删除、发布文章，以及分类和标签管理。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现文章的CRUD操作
- [ ] 实现文章分类管理
- [ ] 实现文章标签管理
- [ ] 实现文章状态管理
- [ ] 实现文章列表查询和分页
- [ ] 实现文章统计功能

## 📝 文章服务

创建 `internal/service/article.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"errors"
	"strings"
	"time"
)

type ArticleService interface {
	Create(article *model.Article) error
	GetByID(id uint) (*model.Article, error)
	GetBySlug(slug string) (*model.Article, error)
	Update(article *model.Article) error
	Delete(id uint, userID uint) error
	List(filter ArticleFilter) ([]model.Article, int64, error)
	IncreaseViewCount(id uint) error
}

type ArticleServiceImpl struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{articleRepo: articleRepo}
}

type ArticleFilter struct {
	Status     string
	CategoryID uint
	TagID      uint
	UserID     uint
	Keyword    string
	Page       int
	PageSize   int
}

func (s *ArticleServiceImpl) Create(article *model.Article) error {
	// 生成slug
	if article.Slug == "" {
		article.Slug = generateSlug(article.Title)
	}
	
	// 如果状态是published，设置发布时间
	if article.Status == "published" && article.PublishedAt == nil {
		now := time.Now()
		article.PublishedAt = &now
	}
	
	return s.articleRepo.Create(article)
}

func (s *ArticleServiceImpl) GetBySlug(slug string) (*model.Article, error) {
	article, err := s.articleRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	
	// 增加阅读量
	go s.IncreaseViewCount(article.ID)
	
	return article, nil
}

func generateSlug(title string) string {
	// 简单的slug生成逻辑
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	// 可以添加更多处理逻辑
	return slug
}
```

## 📝 文章处理器

创建 `internal/handler/article.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/model"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req ArticleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	userID := c.GetUint("user_id")
	
	article := &model.Article{
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		Status:      req.Status,
		CategoryID:  req.CategoryID,
		UserID:     userID,
	}

	if err := h.articleService.Create(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "创建文章失败",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "文章创建成功",
		"data":    article,
	})
}

func (h *ArticleHandler) List(c *gin.Context) {
	filter := service.ArticleFilter{
		Page:     1,
		PageSize: 10,
	}

	// 解析查询参数
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		filter.Page = page
	}
	if pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10")); err == nil {
		filter.PageSize = pageSize
	}
	filter.Status = c.Query("status")
	if categoryID, err := strconv.ParseUint(c.Query("category_id"), 10, 32); err == nil {
		filter.CategoryID = uint(categoryID)
	}
	filter.Keyword = c.Query("keyword")

	articles, total, err := h.articleService.List(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取文章列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":      articles,
			"total":      total,
			"page":       filter.Page,
			"page_size":  filter.PageSize,
		},
	})
}
```

## ⏭️ 下一步

文章管理完成后，下一步是：

- [评论系统](./05-comments.md) - 实现评论和回复功能

---

**🎉 文章管理完成！** 现在你可以开始实现评论系统了。

