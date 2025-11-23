// Package main 实现一个完整的电商系统
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
}

// CartItem 购物车项
type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
}

// Order 订单模型
type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status" gorm:"default:pending"` // pending, paid, shipped, completed, cancelled
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time   `json:"created_at"`
}

// OrderItem 订单项
type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	db.AutoMigrate(&Product{}, &CartItem{}, &Order{}, &OrderItem{})
}

func main() {
	r := gin.Default()

	// 商品管理
	products := r.Group("/api/products")
	{
		// 创建商品
		products.POST("", func(c *gin.Context) {
			var product Product
			if err := c.ShouldBindJSON(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Create(&product)
			c.JSON(http.StatusCreated, gin.H{"data": product})
		})

		// 获取商品列表
		products.GET("", func(c *gin.Context) {
			var products []Product
			db.Find(&products)
			c.JSON(http.StatusOK, gin.H{"data": products})
		})

		// 获取商品详情
		products.GET("/:id", func(c *gin.Context) {
			var product Product
			if err := db.First(&product, c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "商品不存在"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": product})
		})
	}

	// 购物车管理
	cart := r.Group("/api/cart")
	{
		// 添加到购物车
		cart.POST("/items", func(c *gin.Context) {
			var item CartItem
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			item.UserID = 1 // 简化处理
			// 检查是否已存在
			var existing CartItem
			if err := db.Where("user_id = ? AND product_id = ?", item.UserID, item.ProductID).First(&existing).Error; err == nil {
				existing.Quantity += item.Quantity
				db.Save(&existing)
				db.Preload("Product").First(&existing, existing.ID)
				c.JSON(http.StatusOK, gin.H{"data": existing})
				return
			}
			db.Create(&item)
			db.Preload("Product").First(&item, item.ID)
			c.JSON(http.StatusCreated, gin.H{"data": item})
		})

		// 获取购物车
		cart.GET("", func(c *gin.Context) {
			var items []CartItem
			userID := 1 // 简化处理
			db.Preload("Product").Where("user_id = ?", userID).Find(&items)
			var total float64
			for _, item := range items {
				total += item.Product.Price * float64(item.Quantity)
			}
			c.JSON(http.StatusOK, gin.H{
				"items": items,
				"total": total,
			})
		})

		// 删除购物车项
		cart.DELETE("/items/:id", func(c *gin.Context) {
			db.Delete(&CartItem{}, c.Param("id"))
			c.JSON(http.StatusOK, gin.H{"message": "已删除"})
		})
	}

	// 订单管理
	orders := r.Group("/api/orders")
	{
		// 创建订单
		orders.POST("", func(c *gin.Context) {
			userID := uint(1) // 简化处理
			var cartItems []CartItem
			db.Preload("Product").Where("user_id = ?", userID).Find(&cartItems)

			if len(cartItems) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "购物车为空"})
				return
			}

			var totalPrice float64
			var orderItems []OrderItem
			for _, item := range cartItems {
				// 检查库存
				if item.Product.Stock < item.Quantity {
					c.JSON(http.StatusBadRequest, gin.H{"error": "库存不足"})
					return
				}
				totalPrice += item.Product.Price * float64(item.Quantity)
				orderItems = append(orderItems, OrderItem{
					ProductID: item.ProductID,
					Quantity:  item.Quantity,
					Price:     item.Product.Price,
				})
			}

			order := Order{
				UserID:     userID,
				TotalPrice: totalPrice,
				Status:     "pending",
				Items:      orderItems,
			}
			db.Create(&order)

			// 扣减库存
			for _, item := range cartItems {
				db.Model(&Product{}).Where("id = ?", item.ProductID).
					Update("stock", gorm.Expr("stock - ?", item.Quantity))
			}

			// 清空购物车
			db.Where("user_id = ?", userID).Delete(&CartItem{})

			db.Preload("Items.Product").First(&order, order.ID)
			c.JSON(http.StatusCreated, gin.H{"data": order})
		})

		// 获取订单列表
		orders.GET("", func(c *gin.Context) {
			var orders []Order
			userID := 1 // 简化处理
			db.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders)
			c.JSON(http.StatusOK, gin.H{"data": orders})
		})

		// 获取订单详情
		orders.GET("/:id", func(c *gin.Context) {
			var order Order
			if err := db.Preload("Items.Product").First(&order, c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": order})
		})

		// 更新订单状态
		orders.PATCH("/:id/status", func(c *gin.Context) {
			var req struct {
				Status string `json:"status"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Model(&Order{}).Where("id = ?", c.Param("id")).Update("status", req.Status)
			c.JSON(http.StatusOK, gin.H{"message": "状态已更新"})
		})
	}

	log.Println("电商系统启动在 :8080")
	r.Run(":8080")
}
