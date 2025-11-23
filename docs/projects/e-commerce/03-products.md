---
title: 商品管理
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["数据模型设计"]
tags: ["商品", "CRUD", "分类", "SKU", "库存"]
---

# 商品管理

本章节将实现商品的完整管理功能，包括商品CRUD、分类管理、SKU管理、图片管理和库存管理。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现商品的CRUD操作
- [ ] 实现商品分类管理
- [ ] 实现商品SKU管理
- [ ] 实现商品图片管理
- [ ] 实现商品库存管理
- [ ] 实现商品列表查询和筛选

## 📦 商品服务

创建 `internal/service/product.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"errors"
	"fmt"
)

type ProductService interface {
	Create(product *model.Product) error
	GetByID(id uint) (*model.Product, error)
	GetBySlug(slug string) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
	List(filter ProductFilter) ([]model.Product, int64, error)
	UpdateStock(id uint, quantity int) error
	IncreaseSales(id uint, quantity int) error
}

type ProductServiceImpl struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepo: productRepo}
}

type ProductFilter struct {
	CategoryID uint
	Keyword    string
	MinPrice   float64
	MaxPrice   float64
	Status     string
	Page       int
	PageSize   int
}

func (s *ProductServiceImpl) Create(product *model.Product) error {
	// 生成slug
	if product.Slug == "" {
		product.Slug = generateSlug(product.Name)
	}

	// 设置默认状态
	if product.Status == "" {
		product.Status = "active"
	}

	return s.productRepo.Create(product)
}

func (s *ProductServiceImpl) UpdateStock(id uint, quantity int) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New("商品不存在")
	}

	if product.Stock+quantity < 0 {
		return errors.New("库存不足")
	}

	product.Stock += quantity
	
	// 如果库存为0，更新状态
	if product.Stock == 0 {
		product.Status = "sold_out"
	} else if product.Status == "sold_out" {
		product.Status = "active"
	}

	return s.productRepo.Update(product)
}

func (s *ProductServiceImpl) IncreaseSales(id uint, quantity int) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New("商品不存在")
	}

	product.SalesCount += quantity
	return s.productRepo.Update(product)
}
```

## 🖼️ 商品图片管理

### 图片服务

```go
type ProductImageService interface {
	AddImage(productID uint, image ProductImageRequest) error
	UpdateImage(id uint, image ProductImageRequest) error
	DeleteImage(id uint) error
	SetPrimary(productID, imageID uint) error
}

func (s *ProductImageServiceImpl) AddImage(productID uint, image ProductImageRequest) error {
	productImage := &model.ProductImage{
		ProductID: productID,
		URL:       image.URL,
		Sort:      image.Sort,
		IsPrimary: image.IsPrimary,
	}

	// 如果设置为主图，取消其他主图
	if image.IsPrimary {
		s.productImageRepo.UnsetPrimary(productID)
	}

	return s.productImageRepo.Create(productImage)
}
```

## 🏷️ SKU 管理

### SKU 服务

```go
type ProductSKUService interface {
	CreateSKU(productID uint, sku ProductSKURequest) error
	UpdateSKU(id uint, sku ProductSKURequest) error
	DeleteSKU(id uint) error
	GetByProductID(productID uint) ([]model.ProductSKU, error)
	UpdateStock(skuID uint, quantity int) error
}

func (s *ProductSKUServiceImpl) CreateSKU(productID uint, sku ProductSKURequest) error {
	// 生成SKU编号
	if sku.SKU == "" {
		sku.SKU = generateSKU(productID)
	}

	productSKU := &model.ProductSKU{
		ProductID:  productID,
		SKU:        sku.SKU,
		Price:      sku.Price,
		Stock:      sku.Stock,
		Attributes: sku.Attributes,
	}

	return s.productSKURepo.Create(productSKU)
}

func generateSKU(productID uint) string {
	return fmt.Sprintf("SKU%d%d", productID, time.Now().Unix())
}
```

## 📝 商品处理器

创建 `internal/handler/product.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/model"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Status:      "active",
	}

	if err := h.productService.Create(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "创建商品失败",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "商品创建成功",
		"data":    product,
	})
}

func (h *ProductHandler) List(c *gin.Context) {
	filter := service.ProductFilter{
		Page:     1,
		PageSize: 20,
	}

	// 解析查询参数
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		filter.Page = page
	}
	if pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20")); err == nil {
		filter.PageSize = pageSize
	}
	if categoryID, err := strconv.ParseUint(c.Query("category_id"), 10, 32); err == nil {
		filter.CategoryID = uint(categoryID)
	}
	filter.Keyword = c.Query("keyword")
	if minPrice, err := strconv.ParseFloat(c.Query("min_price"), 64); err == nil {
		filter.MinPrice = minPrice
	}
	if maxPrice, err := strconv.ParseFloat(c.Query("max_price"), 64); err == nil {
		filter.MaxPrice = maxPrice
	}
	filter.Status = c.Query("status")

	products, total, err := h.productService.List(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取商品列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":     products,
			"total":     total,
			"page":      filter.Page,
			"page_size": filter.PageSize,
		},
	})
}
```

## 🏷️ 分类管理

### 分类服务

```go
type CategoryService interface {
	Create(category *model.Category) error
	GetByID(id uint) (*model.Category, error)
	GetTree() ([]model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error
}

func (s *CategoryServiceImpl) GetTree() ([]model.Category, error) {
	// 获取所有分类
	categories, err := s.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// 构建树形结构
	return s.buildCategoryTree(categories), nil
}

func (s *CategoryServiceImpl) buildCategoryTree(categories []model.Category) []model.Category {
	categoryMap := make(map[uint]*model.Category)
	var roots []model.Category

	// 创建映射
	for i := range categories {
		categoryMap[categories[i].ID] = &categories[i]
		categories[i].Children = []model.Category{}
	}

	// 构建树
	for i := range categories {
		if categories[i].ParentID == nil {
			roots = append(roots, categories[i])
		} else {
			if parent, ok := categoryMap[*categories[i].ParentID]; ok {
				parent.Children = append(parent.Children, categories[i])
			}
		}
	}

	return roots
}
```

## 🔧 路由设置

```go
func setupProductRoutes(r *gin.RouterGroup, productHandler *handler.ProductHandler) {
	products := r.Group("/products")
	{
		products.GET("", productHandler.List)
		products.GET("/:id", productHandler.GetByID)
		products.GET("/slug/:slug", productHandler.GetBySlug)
	}

	// 需要认证的路由
	admin := r.Group("/admin/products")
	admin.Use(auth.AuthMiddleware(), auth.AdminOnly())
	{
		admin.POST("", productHandler.Create)
		admin.PUT("/:id", productHandler.Update)
		admin.DELETE("/:id", productHandler.Delete)
		admin.POST("/:id/images", productHandler.AddImage)
		admin.POST("/:id/skus", productHandler.AddSKU)
	}
}
```

## 📝 API 使用示例

### 创建商品

```bash
curl -X POST http://localhost:8080/api/admin/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "name": "iPhone 15 Pro",
    "description": "最新款iPhone",
    "price": 8999.00,
    "stock": 100,
    "category_id": 1
  }'
```

### 获取商品列表

```bash
curl "http://localhost:8080/api/products?category_id=1&min_price=1000&max_price=10000&page=1&page_size=20"
```

## 💡 最佳实践

### 1. 商品状态管理

- **active**: 正常销售
- **inactive**: 下架
- **sold_out**: 售罄

### 2. 库存管理

- 实时更新库存
- 库存预警机制
- 防止超卖

### 3. 性能优化

- 缓存热门商品
- 使用索引优化查询
- 分页加载商品列表

## ⏭️ 下一步

商品管理完成后，下一步是：

- [购物车](./04-cart.md) - 实现购物车功能

---

**🎉 商品管理完成！** 现在你可以开始实现购物车功能了。

