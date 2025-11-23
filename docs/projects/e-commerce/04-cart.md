---
title: 购物车
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["商品管理"]
tags: ["购物车", "购物", "结算"]
---

# 购物车

本章节将实现购物车功能，包括添加商品、更新数量、删除商品和购物车结算。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现购物车的CRUD操作
- [ ] 实现购物车商品数量更新
- [ ] 实现购物车结算功能
- [ ] 处理购物车库存验证
- [ ] 实现购物车合并功能

## 🛒 购物车服务

创建 `internal/service/cart.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"errors"
)

type CartService interface {
	AddItem(userID uint, item CartItemRequest) error
	UpdateQuantity(userID uint, itemID uint, quantity int) error
	RemoveItem(userID uint, itemID uint) error
	GetCart(userID uint) ([]model.CartItem, float64, error)
	ClearCart(userID uint) error
	ValidateCart(userID uint) error
}

type CartServiceImpl struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) CartService {
	return &CartServiceImpl{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

type CartItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	SKUID     *uint `json:"sku_id,omitempty"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

func (s *CartServiceImpl) AddItem(userID uint, item CartItemRequest) error {
	// 检查商品是否存在
	product, err := s.productRepo.GetByID(item.ProductID)
	if err != nil {
		return errors.New("商品不存在")
	}

	// 检查商品状态
	if product.Status != "active" {
		return errors.New("商品已下架")
	}

	// 检查库存
	if product.Stock < item.Quantity {
		return errors.New("库存不足")
	}

	// 检查购物车中是否已存在
	existingItem, err := s.cartRepo.GetByUserAndProduct(userID, item.ProductID, item.SKUID)
	if err == nil {
		// 更新数量
		newQuantity := existingItem.Quantity + item.Quantity
		if newQuantity > product.Stock {
			return errors.New("库存不足")
		}
		existingItem.Quantity = newQuantity
		return s.cartRepo.Update(existingItem)
	}

	// 创建新购物车项
	cartItem := &model.CartItem{
		UserID:    userID,
		ProductID: item.ProductID,
		SKUID:     item.SKUID,
		Quantity:  item.Quantity,
	}

	return s.cartRepo.Create(cartItem)
}

func (s *CartServiceImpl) UpdateQuantity(userID uint, itemID uint, quantity int) error {
	if quantity <= 0 {
		return s.RemoveItem(userID, itemID)
	}

	cartItem, err := s.cartRepo.GetByID(itemID)
	if err != nil {
		return errors.New("购物车项不存在")
	}

	if cartItem.UserID != userID {
		return errors.New("无权操作")
	}

	// 检查库存
	product, err := s.productRepo.GetByID(cartItem.ProductID)
	if err != nil {
		return errors.New("商品不存在")
	}

	if product.Stock < quantity {
		return errors.New("库存不足")
	}

	cartItem.Quantity = quantity
	return s.cartRepo.Update(cartItem)
}

func (s *CartServiceImpl) GetCart(userID uint) ([]model.CartItem, float64, error) {
	items, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, 0, err
	}

	// 计算总价
	var total float64
	for _, item := range items {
		price := item.Product.Price
		if item.SKU != nil {
			price = item.SKU.Price
		}
		total += price * float64(item.Quantity)
	}

	return items, total, nil
}

func (s *CartServiceImpl) ValidateCart(userID uint) error {
	items, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	for _, item := range items {
		// 检查商品状态
		if item.Product.Status != "active" {
			return fmt.Errorf("商品 %s 已下架", item.Product.Name)
		}

		// 检查库存
		stock := item.Product.Stock
		if item.SKU != nil {
			stock = item.SKU.Stock
		}

		if stock < item.Quantity {
			return fmt.Errorf("商品 %s 库存不足", item.Product.Name)
		}
	}

	return nil
}
```

## 📝 购物车处理器

创建 `internal/handler/cart.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req service.CartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	if err := h.cartService.AddItem(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "添加成功",
	})
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	items, total, err := h.cartService.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取购物车失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items": items,
			"total": total,
		},
	})
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {
	userID := c.GetUint("user_id")
	itemID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var req struct {
		Quantity int `json:"quantity" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	if err := h.cartService.UpdateQuantity(userID, uint(itemID), req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "更新成功",
	})
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	itemID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.cartService.RemoveItem(userID, uint(itemID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "删除成功",
	})
}
```

## 🔧 路由设置

```go
func setupCartRoutes(r *gin.RouterGroup, cartHandler *handler.CartHandler) {
	cart := r.Group("/cart")
	cart.Use(auth.AuthMiddleware())
	{
		cart.GET("", cartHandler.GetCart)
		cart.POST("/items", cartHandler.AddItem)
		cart.PUT("/items/:id", cartHandler.UpdateQuantity)
		cart.DELETE("/items/:id", cartHandler.RemoveItem)
		cart.DELETE("", cartHandler.ClearCart)
	}
}
```

## 📝 API 使用示例

### 添加商品到购物车

```bash
curl -X POST http://localhost:8080/api/cart/items \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "product_id": 1,
    "quantity": 2
  }'
```

### 获取购物车

```bash
curl http://localhost:8080/api/cart \
  -H "Authorization: Bearer <token>"
```

### 更新数量

```bash
curl -X PUT http://localhost:8080/api/cart/items/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "quantity": 3
  }'
```

## 💡 最佳实践

### 1. 库存验证

- 添加时验证库存
- 更新时验证库存
- 结算前再次验证

### 2. 性能优化

- 使用Redis缓存购物车
- 批量查询商品信息
- 异步更新统计

### 3. 用户体验

- 实时显示库存
- 价格变化提醒
- 购物车过期清理

## ⏭️ 下一步

购物车完成后，下一步是：

- [订单系统](./05-orders.md) - 实现订单创建和管理

---

**🎉 购物车完成！** 现在你可以开始实现订单系统了。

