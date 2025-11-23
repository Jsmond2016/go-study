---
title: 评论系统
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["文章管理"]
tags: ["评论", "回复", "审核", "树形结构"]
---

# 评论系统

本章节将实现文章的评论系统，包括评论发表、回复功能、评论审核和评论树形结构展示。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现评论的CRUD操作
- [ ] 实现评论回复功能
- [ ] 构建评论树形结构
- [ ] 实现评论审核机制
- [ ] 实现评论列表查询和分页
- [ ] 处理评论的关联关系

## 💬 评论服务

创建 `internal/service/comment.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"errors"
)

type CommentService interface {
	Create(comment *model.Comment) error
	GetByArticleID(articleID uint, page, pageSize int) ([]model.Comment, int64, error)
	GetByID(id uint) (*model.Comment, error)
	Approve(id uint) error
	Reject(id uint) error
	Delete(id uint, userID uint) error
}

type CommentServiceImpl struct {
	commentRepo repository.CommentRepository
	articleRepo repository.ArticleRepository
}

func NewCommentService(commentRepo repository.CommentRepository, articleRepo repository.ArticleRepository) CommentService {
	return &CommentServiceImpl{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
	}
}

func (s *CommentServiceImpl) Create(comment *model.Comment) error {
	// 检查文章是否存在
	article, err := s.articleRepo.GetByID(comment.ArticleID)
	if err != nil {
		return errors.New("文章不存在")
	}

	// 如果文章不允许评论
	if article.Status != "published" {
		return errors.New("文章未发布，无法评论")
	}

	// 创建评论
	if err := s.commentRepo.Create(comment); err != nil {
		return err
	}

	// 更新文章评论数
	go s.updateArticleCommentCount(comment.ArticleID)

	return nil
}

func (s *CommentServiceImpl) GetByArticleID(articleID uint, page, pageSize int) ([]model.Comment, int64, error) {
	// 只获取已审核的评论
	comments, err := s.commentRepo.GetByArticleID(articleID, "approved", page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 构建评论树
	tree := s.buildCommentTree(comments)
	
	total, _ := s.commentRepo.CountByArticleID(articleID, "approved")
	
	return tree, total, nil
}

func (s *CommentServiceImpl) buildCommentTree(comments []model.Comment) []model.Comment {
	// 构建评论树结构
	commentMap := make(map[uint]*model.Comment)
	var roots []model.Comment

	// 第一遍：创建映射
	for i := range comments {
		commentMap[comments[i].ID] = &comments[i]
		comments[i].Replies = []model.Comment{}
	}

	// 第二遍：构建树
	for i := range comments {
		if comments[i].ParentID == nil {
			roots = append(roots, comments[i])
		} else {
			if parent, ok := commentMap[*comments[i].ParentID]; ok {
				parent.Replies = append(parent.Replies, comments[i])
			}
		}
	}

	return roots
}

func (s *CommentServiceImpl) updateArticleCommentCount(articleID uint) {
	count, _ := s.commentRepo.CountByArticleID(articleID, "approved")
	s.articleRepo.UpdateCommentCount(articleID, count)
}
```

## 📝 评论处理器

创建 `internal/handler/comment.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/model"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) Create(c *gin.Context) {
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的文章ID",
		})
		return
	}

	var req CommentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	userID := c.GetUint("user_id")
	
	comment := &model.Comment{
		ArticleID: uint(articleID),
		UserID:    userID,
		Content:   req.Content,
		ParentID:  req.ParentID,
		Status:    "pending", // 默认待审核
	}

	// 获取客户端IP和User-Agent
	comment.IP = c.ClientIP()
	comment.UserAgent = c.GetHeader("User-Agent")

	if err := h.commentService.Create(comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "评论发表成功，等待审核",
		"data":    comment,
	})
}

func (h *CommentHandler) List(c *gin.Context) {
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的文章ID",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	comments, total, err := h.commentService.GetByArticleID(uint(articleID), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取评论列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":     comments,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}
```

## 🔍 评论审核

### 审核服务方法

```go
func (s *CommentServiceImpl) Approve(id uint) error {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return errors.New("评论不存在")
	}

	comment.Status = "approved"
	if err := s.commentRepo.Update(comment); err != nil {
		return err
	}

	// 更新文章评论数
	go s.updateArticleCommentCount(comment.ArticleID)

	return nil
}

func (s *CommentServiceImpl) Reject(id uint) error {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return errors.New("评论不存在")
	}

	comment.Status = "rejected"
	return s.commentRepo.Update(comment)
}
```

## 📊 评论统计

### 获取评论统计

```go
func (h *CommentHandler) Statistics(c *gin.Context) {
	articleID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	stats, err := h.commentService.GetStatistics(uint(articleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取统计失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}
```

## 🔧 路由设置

```go
func setupCommentRoutes(r *gin.RouterGroup, commentHandler *handler.CommentHandler) {
	comments := r.Group("/articles/:id/comments")
	{
		comments.GET("", commentHandler.List)
		comments.POST("", auth.AuthMiddleware(), commentHandler.Create)
		comments.DELETE("/:comment_id", auth.AuthMiddleware(), commentHandler.Delete)
	}

	// 管理员路由
	admin := r.Group("/admin/comments")
	admin.Use(auth.AuthMiddleware(), auth.AdminOnly())
	{
		admin.GET("", commentHandler.AdminList)
		admin.PATCH("/:id/approve", commentHandler.Approve)
		admin.PATCH("/:id/reject", commentHandler.Reject)
	}
}
```

## 📝 API 使用示例

### 发表评论

```bash
curl -X POST http://localhost:8080/api/articles/1/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "content": "这是一条评论",
    "parent_id": null
  }'
```

### 回复评论

```bash
curl -X POST http://localhost:8080/api/articles/1/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "content": "这是回复",
    "parent_id": 1
  }'
```

### 获取评论列表

```bash
curl http://localhost:8080/api/articles/1/comments?page=1&page_size=10
```

## 💡 最佳实践

### 1. 评论审核策略

- **自动审核**: 注册用户评论自动通过
- **人工审核**: 新用户或敏感内容需要审核
- **关键词过滤**: 过滤敏感词汇

### 2. 性能优化

- **缓存热门评论**: 使用Redis缓存
- **异步更新统计**: 使用消息队列
- **分页加载**: 避免一次性加载大量评论

### 3. 安全考虑

- **防刷机制**: 限制评论频率
- **内容过滤**: 过滤恶意内容
- **IP记录**: 记录评论者IP

## ⏭️ 下一步

评论系统完成后，下一步是：

- [文件上传](./06-upload.md) - 实现图片上传功能

---

**🎉 评论系统完成！** 现在你可以开始实现文件上传功能了。

