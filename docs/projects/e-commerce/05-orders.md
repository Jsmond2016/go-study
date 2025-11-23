---
title: 订单系统
difficulty: advanced
duration: "4-5小时"
prerequisites: ["购物车"]
tags: ["订单", "状态管理", "订单流转"]
---

# 订单系统

本章节将实现订单系统，包括订单创建、状态管理、订单查询和订单统计等功能。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现订单创建功能
- [ ] 实现订单状态流转
- [ ] 实现订单查询和分页
- [ ] 处理订单取消和退款
- [ ] 实现订单统计功能
- [ ] 处理订单超时自动取消

## 📦 订单服务

创建 `internal/service/order.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"errors"
	"fmt"
	"time"
)

type OrderService interface {
	CreateFromCart(userID uint, addressID uint) (*model.Order, error)
	GetByID(id uint, userID uint) (*model.Order, error)
	GetByUserID(userID uint, page, pageSize int) ([]model.Order, int64, error)
	UpdateStatus(orderID uint, status string) error
	Cancel(orderID uint, userID uint) error
	Complete(orderID uint, userID uint) error
}

type OrderServiceImpl struct {
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, cartRepo repository.CartRepository, productRepo repository.ProductRepository) OrderService {
	return &OrderServiceImpl{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *OrderServiceImpl) CreateFromCart(userID uint, addressID uint) (*model.Order, error) {
	// 获取购物车
	cartItems, err := s.cartRepo.GetByUserID(userID)
	if err != nil || len(cartItems) == 0 {
		return nil, errors.New("购物车为空")
	}

	// 验证购物车
	if err := s.validateCart(cartItems); err != nil {
		return nil, err
	}

	// 计算总价
	var totalAmount float64
	var orderItems []model.OrderItem

	for _, item := range cartItems {
		price := item.Product.Price
		if item.SKU != nil {
			price = item.SKU.Price
		}

		itemTotal := price * float64(item.Quantity)
		totalAmount += itemTotal

		orderItems = append(orderItems, model.OrderItem{
			ProductID: item.ProductID,
			SKUID:     item.SKUID,
			Quantity:  item.Quantity,
			Price:     price,
			Total:     itemTotal,
		})
	}

	// 创建订单
	order := &model.Order{
		OrderNo:      generateOrderNo(),
		UserID:       userID,
		TotalAmount:  totalAmount,
		Status:       "pending",
		PaymentStatus: "unpaid",
		ShippingStatus: "unshipped",
		Items:        orderItems,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// 扣减库存
	for _, item := range cartItems {
		if item.SKU != nil {
			s.productRepo.UpdateSKUStock(item.SKUID, -item.Quantity)
		} else {
			s.productRepo.UpdateStock(item.ProductID, -item.Quantity)
		}
	}

	// 清空购物车
	s.cartRepo.ClearByUserID(userID)

	return order, nil
}

func (s *OrderServiceImpl) validateCart(cartItems []model.CartItem) error {
	for _, item := range cartItems {
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

func (s *OrderServiceImpl) UpdateStatus(orderID uint, status string) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return errors.New("订单不存在")
	}

	// 状态流转验证
	if !s.canChangeStatus(order.Status, status) {
		return errors.New("订单状态不能变更")
	}

	order.Status = status

	// 更新相关时间
	now := time.Now()
	switch status {
	case "paid":
		order.PaymentStatus = "paid"
		order.PaidAt = &now
	case "shipped":
		order.ShippingStatus = "shipped"
		order.ShippedAt = &now
	case "completed":
		order.ShippingStatus = "received"
	case "cancelled":
		// 恢复库存
		s.restoreStock(order)
	}

	return s.orderRepo.Update(order)
}

func (s *OrderServiceImpl) canChangeStatus(current, new string) bool {
	// 定义状态流转规则
	allowed := map[string][]string{
		"pending":   {"paid", "cancelled"},
		"paid":      {"shipped", "cancelled"},
		"shipped":   {"completed", "cancelled"},
		"completed": {},
		"cancelled": {},
	}

	for _, status := range allowed[current] {
		if status == new {
			return true
		}
	}

	return false
}

func (s *OrderServiceImpl) restoreStock(order *model.Order) {
	for _, item := range order.Items {
		if item.SKUID != nil {
			s.productRepo.UpdateSKUStock(item.SKUID, item.Quantity)
		} else {
			s.productRepo.UpdateStock(item.ProductID, item.Quantity)
		}
	}
}

func generateOrderNo() string {
	return fmt.Sprintf("ORD%d%s", time.Now().Unix(), generateRandomString(6))
}
```

## 📝 订单处理器

创建 `internal/handler/order.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req struct {
		AddressID uint `json:"address_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	order, err := h.orderService.CreateFromCart(userID, req.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "订单创建成功",
		"data":    order,
	})
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	userID := c.GetUint("user_id")
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	order, err := h.orderService.GetByID(uint(orderID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "订单不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    order,
	})
}

func (h *OrderHandler) List(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, total, err := h.orderService.GetByUserID(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取订单列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":     orders,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	userID := c.GetUint("user_id")
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.orderService.Cancel(uint(orderID), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "订单已取消",
	})
}
```

## ⏰ 订单超时处理

### 定时任务

```go
func (s *OrderServiceImpl) CancelExpiredOrders() error {
	// 查找30分钟未支付的订单
	expiredTime := time.Now().Add(-30 * time.Minute)
	
	orders, err := s.orderRepo.GetExpiredPendingOrders(expiredTime)
	if err != nil {
		return err
	}

	for _, order := range orders {
		s.Cancel(order.ID, order.UserID)
	}

	return nil
}
```

## 🔧 路由设置

```go
func setupOrderRoutes(r *gin.RouterGroup, orderHandler *handler.OrderHandler) {
	orders := r.Group("/orders")
	orders.Use(auth.AuthMiddleware())
	{
		orders.POST("", orderHandler.Create)
		orders.GET("", orderHandler.List)
		orders.GET("/:id", orderHandler.GetByID)
		orders.PATCH("/:id/cancel", orderHandler.Cancel)
		orders.PATCH("/:id/complete", orderHandler.Complete)
	}
}
```

## 📝 API 使用示例

### 创建订单

```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "address_id": 1
  }'
```

### 获取订单列表

```bash
curl http://localhost:8080/api/orders?page=1&page_size=10 \
  -H "Authorization: Bearer <token>"
```

## 💡 最佳实践

### 1. 订单状态管理

- **pending**: 待支付
- **paid**: 已支付
- **shipped**: 已发货
- **completed**: 已完成
- **cancelled**: 已取消

### 2. 库存管理

- 创建订单时扣减库存
- 取消订单时恢复库存
- 防止超卖

### 3. 订单超时

- 定时检查未支付订单
- 自动取消超时订单
- 恢复库存

## ⏭️ 下一步

订单系统完成后，下一步是：

- [支付集成](./06-payment.md) - 集成支付接口

---

**🎉 订单系统完成！** 现在你可以开始集成支付功能了。

