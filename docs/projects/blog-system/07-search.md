---
title: 搜索功能
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["文章管理"]
tags: ["搜索", "全文搜索", "Elasticsearch", "标签搜索"]
---

# 搜索功能

本章节将实现文章搜索功能，包括全文搜索、标签搜索、分类搜索和高级搜索选项。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现数据库全文搜索
- [ ] 集成Elasticsearch搜索
- [ ] 实现标签和分类搜索
- [ ] 实现搜索高亮
- [ ] 实现搜索建议和自动完成
- [ ] 优化搜索性能

## 🔍 数据库全文搜索

### 基础搜索实现

创建 `internal/service/search.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"fmt"
	"strings"
)

type SearchService interface {
	Search(keyword string, page, pageSize int) ([]model.Article, int64, error)
	SearchByTag(tagID uint, page, pageSize int) ([]model.Article, int64, error)
	SearchByCategory(categoryID uint, page, pageSize int) ([]model.Article, int64, error)
	GetSuggestions(keyword string, limit int) ([]string, error)
}

type SearchServiceImpl struct {
	articleRepo repository.ArticleRepository
}

func NewSearchService(articleRepo repository.ArticleRepository) SearchService {
	return &SearchServiceImpl{articleRepo: articleRepo}
}

func (s *SearchServiceImpl) Search(keyword string, page, pageSize int) ([]model.Article, int64, error) {
	if keyword == "" {
		return s.articleRepo.ListPublished(page, pageSize)
	}

	// 使用LIKE进行搜索
	keyword = "%" + strings.TrimSpace(keyword) + "%"
	return s.articleRepo.Search(keyword, page, pageSize)
}
```

### 搜索仓库实现

创建 `internal/repository/search.go`:

```go
package repository

import (
	"blog-system/internal/model"
	"gorm.io/gorm"
)

func (r *ArticleRepository) Search(keyword string, page, pageSize int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Where("status = ?", "published").
		Where("(title LIKE ? OR content LIKE ? OR summary LIKE ?)", keyword, keyword, keyword).
		Order("created_at DESC")

	// 获取总数
	query.Model(&model.Article{}).Count(&total)

	// 获取数据
	err := query.Offset(offset).Limit(pageSize).
		Preload("User").Preload("Category").Preload("Tags").
		Find(&articles).Error

	return articles, total, err
}
```

## 🔎 Elasticsearch 集成

### Elasticsearch 客户端

创建 `pkg/search/elasticsearch.go`:

```go
package search

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticsearchService struct {
	client *elasticsearch.Client
	index  string
}

func NewElasticsearchService(addresses []string, index string) (*ElasticsearchService, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchService{
		client: client,
		index:  index,
	}, nil
}

func (s *ElasticsearchService) Search(keyword string, page, pageSize int) ([]map[string]interface{}, int64, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"title^3", "content", "summary^2"},
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"highlight": map[string]interface{}{
			"fields": map[string]interface{}{
				"title":   map[string]interface{}{},
				"content": map[string]interface{}{},
			},
		},
	}

	queryJSON, _ := json.Marshal(query)

	req := esapi.SearchRequest{
		Index: []string{s.index},
		Body:  strings.NewReader(string(queryJSON)),
	}

	res, err := req.Do(context.Background(), s.client)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, err
	}

	// 解析结果
	hits := result["hits"].(map[string]interface{})
	total := int64(hits["total"].(map[string]interface{})["value"].(float64))

	var articles []map[string]interface{}
	for _, hit := range hits["hits"].([]interface{}) {
		articles = append(articles, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}

	return articles, total, nil
}
```

## 🏷️ 标签和分类搜索

### 标签搜索

```go
func (s *SearchServiceImpl) SearchByTag(tagID uint, page, pageSize int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Where("status = ?", "published").
		Joins("JOIN article_tags ON articles.id = article_tags.article_id").
		Where("article_tags.tag_id = ?", tagID).
		Order("created_at DESC")

	query.Model(&model.Article{}).Count(&total)

	err := query.Offset(offset).Limit(pageSize).
		Preload("User").Preload("Category").Preload("Tags").
		Find(&articles).Error

	return articles, total, err
}
```

### 分类搜索

```go
func (s *SearchServiceImpl) SearchByCategory(categoryID uint, page, pageSize int) ([]model.Article, int64, error) {
	return s.articleRepo.GetByCategoryID(categoryID, "published", page, pageSize)
}
```

## 💡 搜索建议

### 实现搜索建议

```go
func (s *SearchServiceImpl) GetSuggestions(keyword string, limit int) ([]string, error) {
	if keyword == "" {
		return []string{}, nil
	}

	// 从标题中提取建议
	var suggestions []string
	keyword = "%" + strings.TrimSpace(keyword) + "%"

	err := r.db.Model(&model.Article{}).
		Where("status = ?", "published").
		Where("title LIKE ?", keyword).
		Select("DISTINCT title").
		Limit(limit).
		Pluck("title", &suggestions).Error

	return suggestions, err
}
```

## 📝 搜索处理器

创建 `internal/handler/search.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchService service.SearchService
}

func NewSearchHandler(searchService service.SearchService) *SearchHandler {
	return &SearchHandler{searchService: searchService}
}

func (h *SearchHandler) Search(c *gin.Context) {
	keyword := c.Query("q")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articles, total, err := h.searchService.Search(keyword, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "搜索失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":     articles,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
			"keyword":   keyword,
		},
	})
}

func (h *SearchHandler) Suggestions(c *gin.Context) {
	keyword := c.Query("q")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	suggestions, err := h.searchService.GetSuggestions(keyword, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取建议失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    suggestions,
	})
}
```

## 🔧 路由设置

```go
func setupSearchRoutes(r *gin.RouterGroup, searchHandler *handler.SearchHandler) {
	search := r.Group("/search")
	{
		search.GET("", searchHandler.Search)
		search.GET("/suggestions", searchHandler.Suggestions)
		search.GET("/tag/:id", searchHandler.SearchByTag)
		search.GET("/category/:id", searchHandler.SearchByCategory)
	}
}
```

## 📝 API 使用示例

### 搜索文章

```bash
curl "http://localhost:8080/api/search?q=Go语言&page=1&page_size=10"
```

### 获取搜索建议

```bash
curl "http://localhost:8080/api/search/suggestions?q=Go&limit=5"
```

## 💡 性能优化

### 1. 索引优化

- 为搜索字段创建索引
- 使用全文索引（FULLTEXT）
- 定期优化索引

### 2. 缓存策略

- 缓存热门搜索
- 缓存搜索结果
- 使用Redis缓存

### 3. 搜索优化

- 限制搜索关键词长度
- 使用分词器
- 实现搜索去重

## ⏭️ 下一步

搜索功能完成后，下一步是：

- [部署优化](./08-deployment.md) - 部署和性能优化

---

**🎉 搜索功能完成！** 现在你可以开始学习部署和优化了。

