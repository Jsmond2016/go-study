---
title: 库存管理
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["商品管理"]
tags: ["库存", "预警", "盘点", "出入库"]
---

# 库存管理

本章节将实现库存管理功能，包括库存扣减、库存预警、库存盘点和出入库记录。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现库存扣减和恢复
- [ ] 实现库存预警机制
- [ ] 实现库存盘点功能
- [ ] 记录库存变动历史
- [ ] 处理库存并发问题
- [ ] 实现库存统计功能

## 📦 库存服务

创建 `internal/service/inventory.go`:

```go
package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"errors"
	"fmt"
	"sync"
)

type InventoryService interface {
	DecreaseStock(productID uint, skuID *uint, quantity int) error
	IncreaseStock(productID uint, skuID *uint, quantity int) error
	GetStock(productID uint, skuID *uint) (int, error)
	CheckStock(productID uint, skuID *uint, quantity int) (bool, error)
	GetLowStockProducts(threshold int) ([]model.Product, error)
	RecordInventoryChange(productID uint, skuID *uint, changeType string, quantity int, reason string) error
}

type InventoryServiceImpl struct {
	productRepo repository.ProductRepository
	skuRepo     repository.ProductSKURepository
	inventoryRepo repository.InventoryRepository
	mu          sync.Mutex
}

func NewInventoryService(productRepo repository.ProductRepository, skuRepo repository.ProductSKURepository, inventoryRepo repository.InventoryRepository) InventoryService {
	return &InventoryServiceImpl{
		productRepo:  productRepo,
		skuRepo:       skuRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *InventoryServiceImpl) DecreaseStock(productID uint, skuID *uint, quantity int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if skuID != nil {
		// SKU库存扣减
		sku, err := s.skuRepo.GetByID(*skuID)
		if err != nil {
			return errors.New("SKU不存在")
		}

		if sku.Stock < quantity {
			return errors.New("库存不足")
		}

		sku.Stock -= quantity
		if err := s.skuRepo.Update(sku); err != nil {
			return err
		}

		// 记录库存变动
		s.RecordInventoryChange(productID, skuID, "decrease", quantity, "订单扣减")
	} else {
		// 商品库存扣减
		product, err := s.productRepo.GetByID(productID)
		if err != nil {
			return errors.New("商品不存在")
		}

		if product.Stock < quantity {
			return errors.New("库存不足")
		}

		product.Stock -= quantity
		if product.Stock == 0 {
			product.Status = "sold_out"
		}

		if err := s.productRepo.Update(product); err != nil {
			return err
		}

		// 记录库存变动
		s.RecordInventoryChange(productID, nil, "decrease", quantity, "订单扣减")
	}

	return nil
}

func (s *InventoryServiceImpl) IncreaseStock(productID uint, skuID *uint, quantity int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if skuID != nil {
		sku, err := s.skuRepo.GetByID(*skuID)
		if err != nil {
			return errors.New("SKU不存在")
		}

		sku.Stock += quantity
		if err := s.skuRepo.Update(sku); err != nil {
			return err
		}

		s.RecordInventoryChange(productID, skuID, "increase", quantity, "订单取消恢复")
	} else {
		product, err := s.productRepo.GetByID(productID)
		if err != nil {
			return errors.New("商品不存在")
		}

		product.Stock += quantity
		if product.Status == "sold_out" && product.Stock > 0 {
			product.Status = "active"
		}

		if err := s.productRepo.Update(product); err != nil {
			return err
		}

		s.RecordInventoryChange(productID, nil, "increase", quantity, "订单取消恢复")
	}

	return nil
}

func (s *InventoryServiceImpl) GetLowStockProducts(threshold int) ([]model.Product, error) {
	return s.productRepo.GetLowStock(threshold)
}
```

## ⚠️ 库存预警

### 预警服务

```go
type InventoryAlertService interface {
	CheckLowStock() error
	SendAlert(productID uint, currentStock int, threshold int) error
}

func (s *InventoryAlertServiceImpl) CheckLowStock() error {
	// 获取低库存商品
	products, err := s.inventoryService.GetLowStockProducts(10)
	if err != nil {
		return err
	}

	for _, product := range products {
		// 发送预警通知
		s.SendAlert(product.ID, product.Stock, 10)
	}

	return nil
}

func (s *InventoryAlertServiceImpl) SendAlert(productID uint, currentStock int, threshold int) error {
	// 发送邮件或短信通知
	// 这里可以集成邮件服务或短信服务
	return nil
}
```

## 📊 库存盘点

### 盘点服务

```go
type InventoryCheckService interface {
	CreateCheck(check InventoryCheckRequest) error
	GetCheckHistory(productID uint, page, pageSize int) ([]InventoryCheck, int64, error)
}

type InventoryCheck struct {
	ID          uint      `json:"id"`
	ProductID   uint      `json:"product_id"`
	SKUID       *uint     `json:"sku_id,omitempty"`
	BookStock   int       `json:"book_stock"`   // 账面库存
	ActualStock int       `json:"actual_stock"` // 实际库存
	Difference  int       `json:"difference"`   // 差异
	Reason      string    `json:"reason"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s *InventoryCheckServiceImpl) CreateCheck(check InventoryCheckRequest) error {
	// 获取当前库存
	currentStock, err := s.inventoryService.GetStock(check.ProductID, check.SKUID)
	if err != nil {
		return err
	}

	difference := check.ActualStock - currentStock

	// 创建盘点记录
	inventoryCheck := &InventoryCheck{
		ProductID:   check.ProductID,
		SKUID:       check.SKUID,
		BookStock:   currentStock,
		ActualStock: check.ActualStock,
		Difference:  difference,
		Reason:      check.Reason,
	}

	if err := s.inventoryCheckRepo.Create(inventoryCheck); err != nil {
		return err
	}

	// 如果有差异，调整库存
	if difference != 0 {
		if difference > 0 {
			s.inventoryService.IncreaseStock(check.ProductID, check.SKUID, difference)
		} else {
			s.inventoryService.DecreaseStock(check.ProductID, check.SKUID, -difference)
		}
	}

	return nil
}
```

## 📝 库存变动记录

### 库存变动模型

```go
// InventoryChange 库存变动记录
type InventoryChange struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ProductID   uint      `gorm:"not null;index" json:"product_id"`
	SKUID       *uint     `gorm:"index" json:"sku_id,omitempty"`
	ChangeType  string    `gorm:"not null;size:20" json:"change_type"` // increase, decrease
	Quantity    int       `gorm:"not null" json:"quantity"`
	BeforeStock int       `gorm:"not null" json:"before_stock"`
	AfterStock  int       `gorm:"not null" json:"after_stock"`
	Reason      string    `gorm:"size:200" json:"reason"`
	CreatedAt   time.Time `json:"created_at"`
}
```

## 🔒 并发控制

### 使用数据库锁

```go
func (s *InventoryServiceImpl) DecreaseStockWithLock(productID uint, quantity int) error {
	// 使用数据库事务和行锁
	return s.productRepo.Transaction(func(tx *gorm.DB) error {
		var product model.Product
		// 使用SELECT FOR UPDATE锁定行
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", productID).
			First(&product).Error; err != nil {
			return err
		}

		if product.Stock < quantity {
			return errors.New("库存不足")
		}

		product.Stock -= quantity
		return tx.Save(&product).Error
	})
}
```

## 📝 库存处理器

创建 `internal/handler/inventory.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/service"
	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	inventoryService service.InventoryService
}

func NewInventoryHandler(inventoryService service.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) GetStock(c *gin.Context) {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var skuID *uint
	if skuIDStr := c.Query("sku_id"); skuIDStr != "" {
		if id, err := strconv.ParseUint(skuIDStr, 10, 32); err == nil {
			skuID = new(uint)
			*skuID = uint(id)
		}
	}

	stock, err := h.inventoryService.GetStock(uint(productID), skuID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "获取库存失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"product_id": productID,
			"sku_id":     skuID,
			"stock":      stock,
		},
	})
}

func (h *InventoryHandler) GetLowStock(c *gin.Context) {
	threshold, _ := strconv.Atoi(c.DefaultQuery("threshold", "10"))

	products, err := h.inventoryService.GetLowStockProducts(threshold)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取低库存商品失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
	})
}
```

## 🔧 路由设置

```go
func setupInventoryRoutes(r *gin.RouterGroup, inventoryHandler *handler.InventoryHandler) {
	inventory := r.Group("/inventory")
	inventory.Use(auth.AuthMiddleware(), auth.AdminOnly())
	{
		inventory.GET("/products/:id/stock", inventoryHandler.GetStock)
		inventory.GET("/low-stock", inventoryHandler.GetLowStock)
		inventory.POST("/check", inventoryHandler.CreateCheck)
	}
}
```

## 💡 最佳实践

### 1. 库存扣减策略

- 下单时预扣库存
- 支付成功后确认扣减
- 取消订单恢复库存

### 2. 并发控制

- 使用数据库锁
- 使用Redis分布式锁
- 乐观锁控制

### 3. 库存预警

- 设置预警阈值
- 定时检查库存
- 及时通知管理员

## ⏭️ 下一步

库存管理完成后，下一步是：

- [部署优化](./08-deployment.md) - 部署和性能优化

---

**🎉 库存管理完成！** 现在你可以开始学习部署和优化了。

