---
title: 订单服务实现
difficulty: advanced
duration: "4-5小时"
prerequisites: ["用户服务实现", "商品服务实现"]
tags: ["订单服务", "gRPC", "服务间调用"]
---

# 订单服务实现

本章节将详细介绍如何实现订单服务，包括订单创建、状态管理和服务间调用（调用用户服务验证用户）。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 定义订单服务的 Protocol Buffers 接口
- [ ] 实现订单服务的 gRPC 服务端
- [ ] 实现订单数据存储和访问
- [ ] 实现服务间调用（调用用户服务）
- [ ] 处理服务间调用的错误和超时

## 🎯 服务功能

订单服务提供以下功能：

- ✅ 创建订单
- ✅ 获取订单信息
- ✅ 获取用户订单列表
- ✅ 更新订单状态
- ✅ 取消订单

**依赖关系**:
- 调用用户服务验证用户是否存在

## 📝 Protocol Buffers 定义

### 创建 order.proto

在 `proto/order/` 目录下创建 `order.proto`：

```protobuf
syntax = "proto3";

package order;

option go_package = "go-study/examples/microservices/06-ecommerce-microservices/proto/order";

// 订单服务
service OrderService {
  // 创建订单
  rpc CreateOrder(CreateOrderRequest) returns (OrderResponse);
  // 获取订单
  rpc GetOrder(GetOrderRequest) returns (OrderResponse);
  // 获取用户订单列表
  rpc GetUserOrders(GetUserOrdersRequest) returns (GetUserOrdersResponse);
  // 更新订单状态
  rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (OrderResponse);
  // 取消订单
  rpc CancelOrder(CancelOrderRequest) returns (CancelOrderResponse);
}

// 创建订单请求
message CreateOrderRequest {
  int64 user_id = 1;
  repeated OrderItem items = 2;
  string shipping_address = 3;
}

// 订单项
message OrderItem {
  int64 product_id = 1;
  int32 quantity = 2;
  double price = 3;
}

// 获取订单请求
message GetOrderRequest {
  int64 order_id = 1;
}

// 获取用户订单请求
message GetUserOrdersRequest {
  int64 user_id = 1;
  int32 page = 2;
  int32 page_size = 3;
}

// 更新订单状态请求
message UpdateOrderStatusRequest {
  int64 order_id = 1;
  string status = 2;
}

// 取消订单请求
message CancelOrderRequest {
  int64 order_id = 1;
}

// 订单响应
message OrderResponse {
  int64 order_id = 1;
  int64 user_id = 2;
  repeated OrderItem items = 3;
  double total_amount = 4;
  string status = 5;
  string shipping_address = 6;
  string created_at = 7;
  string updated_at = 8;
}

// 获取用户订单响应
message GetUserOrdersResponse {
  repeated OrderResponse orders = 1;
  int32 total = 2;
  int32 page = 3;
  int32 page_size = 4;
}

// 取消订单响应
message CancelOrderResponse {
  bool success = 1;
  string message = 2;
}
```

## 💻 服务实现

### 1. 初始化数据库

```go
// initDB 初始化数据库
func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "order_service.db")
	if err != nil {
		return nil, err
	}

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			items TEXT NOT NULL,
			total_amount REAL NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending',
			shipping_address TEXT,
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

### 2. 服务间调用

```go
// OrderService 订单服务实现
type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	db         *sql.DB
	consulAddr string
}

// NewOrderService 创建订单服务
func NewOrderService(db *sql.DB, consulAddr string) *OrderService {
	return &OrderService{db: db, consulAddr: consulAddr}
}

// getUserServiceClient 获取用户服务客户端
func (s *OrderService) getUserServiceClient(ctx context.Context) (userpb.UserServiceClient, error) {
	// 从 Consul 发现用户服务
	config := consulapi.DefaultConfig()
	config.Address = s.consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	services, _, err := client.Health().Service("user-service", "", true, nil)
	if err != nil || len(services) == 0 {
		return nil, fmt.Errorf("未找到用户服务")
	}

	service := services[0].Service
	address := fmt.Sprintf("%s:%d", service.Address, service.Port)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return userpb.NewUserServiceClient(conn), nil
}
```

### 3. 创建订单

```go
// CreateOrder 创建订单
func (s *OrderService) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	// 验证用户是否存在
	userClient, err := s.getUserServiceClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("连接用户服务失败: %w", err)
	}

	_, err = userClient.ValidateUser(ctx, &userpb.ValidateUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("用户验证失败: %w", err)
	}

	// 计算总金额
	var totalAmount float64
	for _, item := range req.Items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	// 序列化订单项
	itemsJSON, _ := json.Marshal(req.Items)

	// 插入订单
	result, err := s.db.Exec(
		"INSERT INTO orders (user_id, items, total_amount, status, shipping_address) VALUES (?, ?, ?, ?, ?)",
		req.UserId, string(itemsJSON), totalAmount, "pending", req.ShippingAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("创建订单失败: %w", err)
	}

	orderID, _ := result.LastInsertId()
	return s.GetOrder(ctx, &orderpb.GetOrderRequest{OrderId: orderID})
}
```

### 4. 其他方法实现

```go
// GetOrder 获取订单
func (s *OrderService) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.OrderResponse, error) {
	var order orderpb.OrderResponse
	var itemsJSON string
	var createdAt, updatedAt time.Time

	err := s.db.QueryRow(
		"SELECT id, user_id, items, total_amount, status, shipping_address, created_at, updated_at FROM orders WHERE id = ?",
		req.OrderId,
	).Scan(&order.OrderId, &order.UserId, &itemsJSON, &order.TotalAmount, &order.Status, &order.ShippingAddress, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("订单不存在")
	}
	if err != nil {
		return nil, fmt.Errorf("查询订单失败: %w", err)
	}

	// 反序列化订单项
	json.Unmarshal([]byte(itemsJSON), &order.Items)

	order.CreatedAt = createdAt.Format(time.RFC3339)
	order.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &order, nil
}

// GetUserOrders 获取用户订单列表
func (s *OrderService) GetUserOrders(ctx context.Context, req *orderpb.GetUserOrdersRequest) (*orderpb.GetUserOrdersResponse, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	rows, err := s.db.Query(
		"SELECT id, user_id, items, total_amount, status, shipping_address, created_at, updated_at FROM orders WHERE user_id = ? LIMIT ? OFFSET ?",
		req.UserId, pageSize, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("查询订单失败: %w", err)
	}
	defer rows.Close()

	var orders []*orderpb.OrderResponse
	for rows.Next() {
		var order orderpb.OrderResponse
		var itemsJSON string
		var createdAt, updatedAt time.Time

		err := rows.Scan(&order.OrderId, &order.UserId, &itemsJSON, &order.TotalAmount, &order.Status, &order.ShippingAddress, &createdAt, &updatedAt)
		if err != nil {
			continue
		}

		json.Unmarshal([]byte(itemsJSON), &order.Items)
		order.CreatedAt = createdAt.Format(time.RFC3339)
		order.UpdatedAt = updatedAt.Format(time.RFC3339)

		orders = append(orders, &order)
	}

	// 获取总数
	var total int32
	s.db.QueryRow("SELECT COUNT(*) FROM orders WHERE user_id = ?", req.UserId).Scan(&total)

	return &orderpb.GetUserOrdersResponse{
		Orders:   orders,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// UpdateOrderStatus 更新订单状态
func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.OrderResponse, error) {
	_, err := s.db.Exec(
		"UPDATE orders SET status = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		req.Status, req.OrderId,
	)
	if err != nil {
		return nil, fmt.Errorf("更新订单状态失败: %w", err)
	}

	return s.GetOrder(ctx, &orderpb.GetOrderRequest{OrderId: req.OrderId})
}

// CancelOrder 取消订单
func (s *OrderService) CancelOrder(ctx context.Context, req *orderpb.CancelOrderRequest) (*orderpb.CancelOrderResponse, error) {
	result, err := s.db.Exec(
		"UPDATE orders SET status = 'cancelled', updated_at = CURRENT_TIMESTAMP WHERE id = ? AND status != 'cancelled'",
		req.OrderId,
	)
	if err != nil {
		return nil, fmt.Errorf("取消订单失败: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return &orderpb.CancelOrderResponse{Success: false, Message: "订单不存在或已取消"}, nil
	}

	return &orderpb.CancelOrderResponse{Success: true, Message: "订单已取消"}, nil
}
```

## 🔑 关键功能

### 服务间调用

订单服务需要调用用户服务验证用户：

1. **服务发现**: 从 Consul 发现用户服务
2. **连接管理**: 创建 gRPC 连接
3. **错误处理**: 处理服务不可用的情况
4. **超时控制**: 设置调用超时

### 错误处理

```go
// 设置超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 调用用户服务
resp, err := userClient.ValidateUser(ctx, &userpb.ValidateUserRequest{UserId: req.UserId})
if err != nil {
	// 处理错误
	return nil, fmt.Errorf("用户验证失败: %w", err)
}
```

## 🚀 运行服务

```bash
# 确保用户服务已启动
cd user-service
go run main.go

# 启动订单服务
cd ../order-service
go mod tidy
go run main.go
```

## 📚 下一步

订单服务实现完成后，可以继续：

1. [API 网关](./06-gateway.md) - 实现 API 网关统一入口

---

**✅ 订单服务实现完成！** 现在可以继续实现 API 网关了。

