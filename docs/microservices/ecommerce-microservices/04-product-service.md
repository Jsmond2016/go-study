---
title: 商品服务实现
difficulty: advanced
duration: "3-4小时"
prerequisites: ["用户服务实现"]
tags: ["商品服务", "gRPC", "库存管理"]
---

# 商品服务实现

本章节将详细介绍如何实现商品服务，包括商品信息管理、库存管理和商品查询功能。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 定义商品服务的 Protocol Buffers 接口
- [ ] 实现商品服务的 gRPC 服务端
- [ ] 实现商品数据存储和访问
- [ ] 实现库存管理功能
- [ ] 实现商品查询和列表功能

## 🎯 服务功能

商品服务提供以下功能：

- ✅ 创建商品
- ✅ 获取商品信息
- ✅ 更新商品信息
- ✅ 删除商品
- ✅ 获取商品列表
- ✅ 更新库存
- ✅ 检查库存

## 📝 Protocol Buffers 定义

### 创建 product.proto

在 `proto/product/` 目录下创建 `product.proto`：

```protobuf
syntax = "proto3";

package product;

option go_package = "go-study/examples/microservices/06-ecommerce-microservices/proto/product";

// 商品服务
service ProductService {
  // 创建商品
  rpc CreateProduct(CreateProductRequest) returns (ProductResponse);
  // 获取商品
  rpc GetProduct(GetProductRequest) returns (ProductResponse);
  // 更新商品
  rpc UpdateProduct(UpdateProductRequest) returns (ProductResponse);
  // 删除商品
  rpc DeleteProduct(DeleteProductRequest) returns (DeleteProductResponse);
  // 获取商品列表
  rpc ListProducts(ListProductsRequest) returns (ListProductsResponse);
  // 更新库存
  rpc UpdateStock(UpdateStockRequest) returns (UpdateStockResponse);
  // 检查库存
  rpc CheckStock(CheckStockRequest) returns (CheckStockResponse);
}

// 创建商品请求
message CreateProductRequest {
  string name = 1;
  string description = 2;
  double price = 3;
  int32 stock = 4;
  string category = 5;
}

// 获取商品请求
message GetProductRequest {
  int64 product_id = 1;
}

// 更新商品请求
message UpdateProductRequest {
  int64 product_id = 1;
  string name = 2;
  string description = 3;
  double price = 4;
  int32 stock = 5;
  string category = 6;
}

// 删除商品请求
message DeleteProductRequest {
  int64 product_id = 1;
}

// 获取商品列表请求
message ListProductsRequest {
  int32 page = 1;
  int32 page_size = 2;
  string category = 3;
}

// 更新库存请求
message UpdateStockRequest {
  int64 product_id = 1;
  int32 quantity = 2;  // 正数增加，负数减少
}

// 检查库存请求
message CheckStockRequest {
  int64 product_id = 1;
  int32 quantity = 2;
}

// 商品响应
message ProductResponse {
  int64 product_id = 1;
  string name = 2;
  string description = 3;
  double price = 4;
  int32 stock = 5;
  string category = 6;
  string created_at = 7;
  string updated_at = 8;
}

// 获取商品列表响应
message ListProductsResponse {
  repeated ProductResponse products = 1;
  int32 total = 2;
  int32 page = 3;
  int32 page_size = 4;
}

// 更新库存响应
message UpdateStockResponse {
  bool success = 1;
  int32 new_stock = 2;
  string message = 3;
}

// 检查库存响应
message CheckStockResponse {
  bool available = 1;
  int32 current_stock = 2;
}
```

## 💻 服务实现

### 1. 初始化数据库

```go
// initDB 初始化数据库
func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "product_service.db")
	if err != nil {
		return nil, err
	}

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price REAL NOT NULL,
			stock INTEGER NOT NULL DEFAULT 0,
			category TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
```

### 2. 服务实现

```go
// ProductService 商品服务实现
type ProductService struct {
	pb.UnimplementedProductServiceServer
	db *sql.DB
}

// NewProductService 创建商品服务
func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

// CreateProduct 创建商品
func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	result, err := s.db.Exec(
		"INSERT INTO products (name, description, price, stock, category) VALUES (?, ?, ?, ?, ?)",
		req.Name, req.Description, req.Price, req.Stock, req.Category,
	)
	if err != nil {
		return nil, fmt.Errorf("创建商品失败: %w", err)
	}

	productID, _ := result.LastInsertId()
	return s.GetProduct(ctx, &pb.GetProductRequest{ProductId: productID})
}

// GetProduct 获取商品
func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	var product pb.ProductResponse
	var createdAt, updatedAt time.Time

	err := s.db.QueryRow(
		"SELECT id, name, description, price, stock, category, created_at, updated_at FROM products WHERE id = ?",
		req.ProductId,
	).Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("商品不存在")
	}
	if err != nil {
		return nil, fmt.Errorf("查询商品失败: %w", err)
	}

	product.CreatedAt = createdAt.Format(time.RFC3339)
	product.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &product, nil
}

// UpdateProduct 更新商品
func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	_, err := s.db.Exec(
		"UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		req.Name, req.Description, req.Price, req.Stock, req.Category, req.ProductId,
	)
	if err != nil {
		return nil, fmt.Errorf("更新商品失败: %w", err)
	}

	return s.GetProduct(ctx, &pb.GetProductRequest{ProductId: req.ProductId})
}

// DeleteProduct 删除商品
func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	result, err := s.db.Exec("DELETE FROM products WHERE id = ?", req.ProductId)
	if err != nil {
		return nil, fmt.Errorf("删除商品失败: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return &pb.DeleteProductResponse{Success: false, Message: "商品不存在"}, nil
	}

	return &pb.DeleteProductResponse{Success: true, Message: "商品已删除"}, nil
}

// ListProducts 获取商品列表
func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var query string
	var args []interface{}

	if req.Category != "" {
		query = "SELECT id, name, description, price, stock, category, created_at, updated_at FROM products WHERE category = ? LIMIT ? OFFSET ?"
		args = []interface{}{req.Category, pageSize, offset}
	} else {
		query = "SELECT id, name, description, price, stock, category, created_at, updated_at FROM products LIMIT ? OFFSET ?"
		args = []interface{}{pageSize, offset}
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("查询商品失败: %w", err)
	}
	defer rows.Close()

	var products []*pb.ProductResponse
	for rows.Next() {
		var product pb.ProductResponse
		var createdAt, updatedAt time.Time

		err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category, &createdAt, &updatedAt)
		if err != nil {
			continue
		}

		product.CreatedAt = createdAt.Format(time.RFC3339)
		product.UpdatedAt = updatedAt.Format(time.RFC3339)

		products = append(products, &product)
	}

	// 获取总数
	var total int32
	if req.Category != "" {
		s.db.QueryRow("SELECT COUNT(*) FROM products WHERE category = ?", req.Category).Scan(&total)
	} else {
		s.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
	}

	return &pb.ListProductsResponse{
		Products: products,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// UpdateStock 更新库存
func (s *ProductService) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {
	var currentStock int32
	err := s.db.QueryRow("SELECT stock FROM products WHERE id = ?", req.ProductId).Scan(&currentStock)
	if err == sql.ErrNoRows {
		return &pb.UpdateStockResponse{Success: false, Message: "商品不存在"}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("查询库存失败: %w", err)
	}

	newStock := currentStock + req.Quantity
	if newStock < 0 {
		return &pb.UpdateStockResponse{Success: false, Message: "库存不足"}, nil
	}

	_, err = s.db.Exec("UPDATE products SET stock = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", newStock, req.ProductId)
	if err != nil {
		return nil, fmt.Errorf("更新库存失败: %w", err)
	}

	return &pb.UpdateStockResponse{Success: true, NewStock: newStock, Message: "库存已更新"}, nil
}

// CheckStock 检查库存
func (s *ProductService) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	var currentStock int32
	err := s.db.QueryRow("SELECT stock FROM products WHERE id = ?", req.ProductId).Scan(&currentStock)
	if err == sql.ErrNoRows {
		return &pb.CheckStockResponse{Available: false, CurrentStock: 0}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("查询库存失败: %w", err)
	}

	return &pb.CheckStockResponse{
		Available:    currentStock >= req.Quantity,
		CurrentStock: currentStock,
	}, nil
}
```

## 🔑 关键功能

### 库存管理

库存管理是商品服务的核心功能：

1. **更新库存**: 支持增加和减少库存
2. **检查库存**: 检查是否有足够库存
3. **库存保护**: 防止库存为负数

### 分页查询

实现商品列表的分页查询：

- 支持按分类筛选
- 支持分页参数
- 返回总数和分页信息

## 🚀 运行服务

```bash
cd product-service
go mod tidy
go run main.go
```

## 📚 下一步

商品服务实现完成后，可以继续：

1. [订单服务](./05-order-service.md) - 实现订单服务（可能需要调用商品服务检查库存）

---

**✅ 商品服务实现完成！** 现在可以继续实现订单服务了。

