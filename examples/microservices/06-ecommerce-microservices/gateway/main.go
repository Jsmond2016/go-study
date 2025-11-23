package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderpb "go-study/examples/microservices/06-ecommerce-microservices/proto/order"
	productpb "go-study/examples/microservices/06-ecommerce-microservices/proto/product"
	userpb "go-study/examples/microservices/06-ecommerce-microservices/proto/user"
)

var (
	port       = flag.Int("port", 8080, "网关端口")
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
)

func main() {
	flag.Parse()

	// 创建网关
	gateway := NewGateway(*consulAddr)

	// 创建 Gin 路由
	router := gin.Default()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API 路由
	api := router.Group("/api")
	{
		// 用户服务路由
		users := api.Group("/users")
		{
			users.POST("", gateway.CreateUser)
			users.GET("/:id", gateway.GetUser)
			users.PUT("/:id", gateway.UpdateUser)
			users.DELETE("/:id", gateway.DeleteUser)
			users.POST("/login", gateway.Login)
		}

		// 订单服务路由
		orders := api.Group("/orders")
		{
			orders.POST("", gateway.CreateOrder)
			orders.GET("/:id", gateway.GetOrder)
			orders.GET("/user/:user_id", gateway.GetUserOrders)
			orders.PUT("/:id/status", gateway.UpdateOrderStatus)
			orders.DELETE("/:id", gateway.CancelOrder)
		}

		// 商品服务路由
		products := api.Group("/products")
		{
			products.POST("", gateway.CreateProduct)
			products.GET("/:id", gateway.GetProduct)
			products.PUT("/:id", gateway.UpdateProduct)
			products.DELETE("/:id", gateway.DeleteProduct)
			products.GET("", gateway.ListProducts)
		}
	}

	// 启动服务器
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("API 网关启动在端口 %d", *port)
	log.Fatal(server.ListenAndServe())
}

// Gateway API 网关
type Gateway struct {
	consulAddr string
	clients    map[string]*grpc.ClientConn
}

// NewGateway 创建网关
func NewGateway(consulAddr string) *Gateway {
	return &Gateway{
		consulAddr: consulAddr,
		clients:    make(map[string]*grpc.ClientConn),
	}
}

// getServiceClient 获取服务客户端
func (g *Gateway) getServiceClient(serviceName string) (interface{}, error) {
	// 从 Consul 发现服务
	config := consulapi.DefaultConfig()
	config.Address = g.consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}

	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil || len(services) == 0 {
		return nil, fmt.Errorf("未找到服务: %s", serviceName)
	}

	service := services[0].Service
	address := fmt.Sprintf("%s:%d", service.Address, service.Port)

	// 检查是否已有连接
	if conn, ok := g.clients[address]; ok {
		return conn, nil
	}

	// 创建新连接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	g.clients[address] = conn

	switch serviceName {
	case "user-service":
		return userpb.NewUserServiceClient(conn), nil
	case "order-service":
		return orderpb.NewOrderServiceClient(conn), nil
	case "product-service":
		return productpb.NewProductServiceClient(conn), nil
	default:
		return nil, fmt.Errorf("未知服务: %s", serviceName)
	}
}

// CORSMiddleware CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// 用户服务处理函数
func (g *Gateway) CreateUser(c *gin.Context) {
	var req userpb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) GetUser(c *gin.Context) {
	userID := c.Param("id")
	req := &userpb.GetUserRequest{UserId: parseID(userID)}

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.GetUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var req userpb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = parseID(userID)

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	req := &userpb.DeleteUserRequest{UserId: parseID(userID)}

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.DeleteUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) Login(c *gin.Context) {
	var req userpb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := g.getServiceClient("user-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userClient := client.(userpb.UserServiceClient)
	resp, err := userClient.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// 订单服务处理函数
func (g *Gateway) CreateOrder(c *gin.Context) {
	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := g.getServiceClient("order-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderClient := client.(orderpb.OrderServiceClient)
	resp, err := orderClient.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) GetOrder(c *gin.Context) {
	orderID := c.Param("id")
	req := &orderpb.GetOrderRequest{OrderId: parseID(orderID)}

	client, err := g.getServiceClient("order-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderClient := client.(orderpb.OrderServiceClient)
	resp, err := orderClient.GetOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) GetUserOrders(c *gin.Context) {
	userID := c.Param("user_id")
	page := parseInt(c.Query("page"), 1)
	pageSize := parseInt(c.Query("page_size"), 10)

	req := &orderpb.GetUserOrdersRequest{
		UserId:   parseID(userID),
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

	client, err := g.getServiceClient("order-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderClient := client.(orderpb.OrderServiceClient)
	resp, err := orderClient.GetUserOrders(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var req orderpb.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.OrderId = parseID(orderID)

	client, err := g.getServiceClient("order-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderClient := client.(orderpb.OrderServiceClient)
	resp, err := orderClient.UpdateOrderStatus(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) CancelOrder(c *gin.Context) {
	orderID := c.Param("id")
	req := &orderpb.CancelOrderRequest{OrderId: parseID(orderID)}

	client, err := g.getServiceClient("order-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderClient := client.(orderpb.OrderServiceClient)
	resp, err := orderClient.CancelOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// 商品服务处理函数
func (g *Gateway) CreateProduct(c *gin.Context) {
	var req productpb.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := g.getServiceClient("product-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productClient := client.(productpb.ProductServiceClient)
	resp, err := productClient.CreateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) GetProduct(c *gin.Context) {
	productID := c.Param("id")
	req := &productpb.GetProductRequest{ProductId: parseID(productID)}

	client, err := g.getServiceClient("product-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productClient := client.(productpb.ProductServiceClient)
	resp, err := productClient.GetProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var req productpb.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ProductId = parseID(productID)

	client, err := g.getServiceClient("product-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productClient := client.(productpb.ProductServiceClient)
	resp, err := productClient.UpdateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (g *Gateway) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	req := &productpb.DeleteProductRequest{ProductId: parseID(productID)}

	client, err := g.getServiceClient("product-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productClient := client.(productpb.ProductServiceClient)
	_, err = productClient.DeleteProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "商品已删除"})
}

func (g *Gateway) ListProducts(c *gin.Context) {
	page := parseInt(c.Query("page"), 1)
	pageSize := parseInt(c.Query("page_size"), 10)
	category := c.Query("category")

	req := &productpb.ListProductsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Category: category,
	}

	client, err := g.getServiceClient("product-service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productClient := client.(productpb.ProductServiceClient)
	resp, err := productClient.ListProducts(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// 辅助函数
func parseID(s string) int64 {
	var id int64
	fmt.Sscanf(s, "%d", &id)
	return id
}

func parseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	var val int
	fmt.Sscanf(s, "%d", &val)
	if val <= 0 {
		return defaultValue
	}
	return val
}

