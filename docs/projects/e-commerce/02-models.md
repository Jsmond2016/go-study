---
title: 数据模型设计
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["环境搭建"]
tags: ["数据模型", "数据库", "GORM", "电商"]
---

# 数据模型设计

本章节将详细介绍电商系统的数据模型设计，包括商品、订单、购物车、支付等模型的定义和关联关系。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 设计电商系统的数据库表结构
- [ ] 使用 GORM 定义数据模型
- [ ] 理解电商业务模型之间的关系
- [ ] 实现数据库自动迁移
- [ ] 掌握复杂业务模型的设计

## 🗄️ 数据库设计

### 实体关系图

```
User (用户)
  ├── Orders (订单) 1:N
  ├── CartItems (购物车) 1:N
  └── Addresses (地址) 1:N

Product (商品)
  ├── Category (分类) N:1
  ├── ProductImages (商品图片) 1:N
  ├── ProductSKUs (SKU) 1:N
  └── CartItems (购物车项) 1:N

Order (订单)
  ├── User (用户) N:1
  ├── OrderItems (订单项) 1:N
  └── Payment (支付) 1:1

OrderItem (订单项)
  ├── Order (订单) N:1
  └── ProductSKU (SKU) N:1
```

## 📦 模型定义

### 1. 商品模型

创建 `internal/model/product.go`:

```go
package model

import (
	"time"
	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;size:200;index" json:"name"`
	Slug        string    `gorm:"uniqueIndex;not null;size:255" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"not null;index" json:"price"`
	Stock       int       `gorm:"default:0" json:"stock"`
	Status      string    `gorm:"default:active;size:20;index" json:"status"` // active, inactive, sold_out
	SalesCount  int       `gorm:"default:0" json:"sales_count"`
	ViewCount   int       `gorm:"default:0" json:"view_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 外键
	CategoryID uint `gorm:"not null;index" json:"category_id"`

	// 关联
	Category    Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Images      []ProductImage `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	SKUs        []ProductSKU  `gorm:"foreignKey:ProductID" json:"skus,omitempty"`
	CartItems   []CartItem    `gorm:"foreignKey:ProductID" json:"-"`
}

// ProductImage 商品图片
type ProductImage struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProductID uint   `gorm:"not null;index" json:"product_id"`
	URL       string `gorm:"not null;size:500" json:"url"`
	Sort      int    `gorm:"default:0" json:"sort"`
	IsPrimary bool   `gorm:"default:false" json:"is_primary"`
}

// ProductSKU 商品SKU
type ProductSKU struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `gorm:"not null;index" json:"product_id"`
	SKU       string  `gorm:"uniqueIndex;not null;size:100" json:"sku"`
	Price     float64 `gorm:"not null" json:"price"`
	Stock     int     `gorm:"default:0" json:"stock"`
	Attributes string `gorm:"type:json" json:"attributes"` // 规格属性JSON
}
```

### 2. 分类模型

```go
// Category 分类模型
type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Slug        string    `gorm:"uniqueIndex;not null;size:100" json:"slug"`
	Description string    `gorm:"size:255" json:"description"`
	ParentID    *uint     `gorm:"index" json:"parent_id,omitempty"`
	Sort        int       `gorm:"default:0" json:"sort"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Products  []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
	Children  []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Parent    *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
}
```

### 3. 购物车模型

```go
// CartItem 购物车项
type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ProductID uint     `gorm:"not null;index" json:"product_id"`
	SKUID     *uint     `gorm:"index" json:"sku_id,omitempty"`
	Quantity  int       `gorm:"not null;default:1" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Product Product    `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	SKU     *ProductSKU `gorm:"foreignKey:SKUID" json:"sku,omitempty"`
}
```

### 4. 订单模型

```go
// Order 订单模型
type Order struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderNo       string    `gorm:"uniqueIndex;not null;size:50" json:"order_no"`
	UserID        uint      `gorm:"not null;index" json:"user_id"`
	TotalAmount   float64   `gorm:"not null" json:"total_amount"`
	Status        string    `gorm:"default:pending;size:20;index" json:"status"` // pending, paid, shipped, completed, cancelled
	PaymentStatus string    `gorm:"default:unpaid;size:20" json:"payment_status"` // unpaid, paid, refunded
	ShippingStatus string   `gorm:"default:unshipped;size:20" json:"shipping_status"` // unshipped, shipped, received
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`
	ShippedAt     *time.Time `json:"shipped_at,omitempty"`

	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	Payment    *Payment    `gorm:"foreignKey:OrderID" json:"payment,omitempty"`
	Address    *Address    `gorm:"foreignKey:OrderID" json:"address,omitempty"`
}

// OrderItem 订单项
type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null;index" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	SKUID     *uint   `json:"sku_id,omitempty"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
	Total     float64 `gorm:"not null" json:"total"`

	Order   Order      `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Product Product    `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	SKU     *ProductSKU `gorm:"foreignKey:SKUID" json:"sku,omitempty"`
}
```

### 5. 支付模型

```go
// Payment 支付模型
type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `gorm:"uniqueIndex;not null" json:"order_id"`
	PaymentNo     string    `gorm:"uniqueIndex;not null;size:50" json:"payment_no"`
	Amount        float64   `gorm:"not null" json:"amount"`
	Method        string    `gorm:"not null;size:20" json:"method"` // alipay, wechat
	Status        string    `gorm:"default:pending;size:20" json:"status"` // pending, success, failed
	TransactionID string    `gorm:"size:100" json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`

	Order Order `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}
```

### 6. 地址模型

```go
// Address 地址模型
type Address struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	OrderID   *uint     `gorm:"index" json:"order_id,omitempty"`
	Name      string    `gorm:"not null;size:50" json:"name"`
	Phone     string    `gorm:"not null;size:20" json:"phone"`
	Province  string    `gorm:"not null;size:50" json:"province"`
	City      string    `gorm:"not null;size:50" json:"city"`
	District  string    `gorm:"not null;size:50" json:"district"`
	Detail    string    `gorm:"not null;size:200" json:"detail"`
	IsDefault bool      `gorm:"default:false" json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Order *Order `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}
```

## 🗄️ 数据库迁移

```go
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Category{},
		&Product{},
		&ProductImage{},
		&ProductSKU{},
		&CartItem{},
		&Order{},
		&OrderItem{},
		&Payment{},
		&Address{},
	)
}
```

## 💡 最佳实践

### 1. 订单号生成

```go
func GenerateOrderNo() string {
	return fmt.Sprintf("ORD%d%s", time.Now().Unix(), generateRandomString(6))
}
```

### 2. 库存管理

```go
func (p *Product) DecreaseStock(quantity int) error {
	if p.Stock < quantity {
		return errors.New("库存不足")
	}
	p.Stock -= quantity
	return nil
}
```

## ⏭️ 下一步

数据模型设计完成后，下一步是：

- [商品管理](./03-products.md) - 实现商品的CRUD操作

---

**🎉 数据模型设计完成！** 现在你可以开始实现商品管理功能了。

