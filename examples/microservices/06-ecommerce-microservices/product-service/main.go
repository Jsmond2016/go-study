package main

import (
	"context"
	"database/sql"
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

	pb "go-study/examples/microservices/06-ecommerce-microservices/proto/product"
)

var (
	port       = flag.Int("port", 5003, "服务端口")
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
	service := NewProductService(db)

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, service)

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

	log.Printf("商品服务启动在端口 %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务运行失败: %v", err)
	}

	<-ctx.Done()
	log.Println("服务已关闭")
}

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
	serviceID := fmt.Sprintf("product-service-%s-%d", hostname, port)

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "product-service",
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

