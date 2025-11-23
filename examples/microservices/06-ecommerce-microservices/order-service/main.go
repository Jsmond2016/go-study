package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderpb "go-study/examples/microservices/06-ecommerce-microservices/proto/order"
	userpb "go-study/examples/microservices/06-ecommerce-microservices/proto/user"
)

var (
	port       = flag.Int("port", 5002, "服务端口")
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
)

func main() {
	flag.Parse()

	// 初始化数据库
	db, err := initDB()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 创建服务实现
	service := NewOrderService(db, *consulAddr)

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s, service)

	// 注册到 Consul
	registerToConsul(*port, *consulAddr)

	// 优雅关闭
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("收到关闭信号，开始优雅关闭...")
		s.GracefulStop()
		cancel()
	}()

	log.Printf("订单服务启动在端口 %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务运行失败: %v", err)
	}

	<-ctx.Done()
	log.Println("服务已关闭")
}

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

// registerToConsul 注册服务到 Consul
func registerToConsul(port int, consulAddr string) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Printf("创建 Consul 客户端失败: %v", err)
		return
	}

	hostname, _ := os.Hostname()
	serviceID := fmt.Sprintf("order-service-%s-%d", hostname, port)

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "order-service",
		Tags:    []string{"grpc", "v1"},
		Address: "localhost",
		Port:    port,
		Check: &consulapi.AgentServiceCheck{
			TCP:      fmt.Sprintf("localhost:%d", port),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Printf("注册服务失败: %v", err)
	} else {
		log.Printf("服务已注册到 Consul: %s", serviceID)
	}
}

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

