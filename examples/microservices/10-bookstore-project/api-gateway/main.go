package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bookpb "bookstore/proto/book"
	userpb "bookstore/proto/user"
)

const (
	port            = ":8080"
	bookServiceAddr = "localhost:50051"
	userServiceAddr = "localhost:50052"
)

type Gateway struct {
	bookClient bookpb.BookServiceClient
	userClient userpb.UserServiceClient
}

func main() {
	// 连接到 Book Service
	bookConn, err := grpc.NewClient(
		bookServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Book Service: %v", err)
	}
	defer bookConn.Close()

	// 连接到 User Service
	userConn, err := grpc.NewClient(
		userServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer userConn.Close()

	// 创建 Gateway
	gateway := &Gateway{
		bookClient: bookpb.NewBookServiceClient(bookConn),
		userClient: userpb.NewUserServiceClient(userConn),
	}

	// 设置 Gin
	r := gin.Default()

	// CORS 配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 用户相关路由（公开）
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", gateway.Register)
		userGroup.POST("/login", gateway.Login)
	}

	// 需要认证的路由
	authGroup := r.Group("/api")
	authGroup.Use(gateway.AuthMiddleware())
	{
		// 用户
		authGroup.GET("/users/:id", gateway.GetUser)
		authGroup.PUT("/users/:id", gateway.UpdateUser)

		// 图书
		authGroup.GET("/books/:id", gateway.GetBook)
		authGroup.GET("/books", gateway.ListBooks)
		authGroup.POST("/books", gateway.CreateBook)
		authGroup.PUT("/books/:id", gateway.UpdateBook)
		authGroup.DELETE("/books/:id", gateway.DeleteBook)
		authGroup.GET("/books/search", gateway.SearchBooks)
	}

	log.Printf("API Gateway is running on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// ============ 中间件 ============

// AuthMiddleware - JWT 认证中间件
func (g *Gateway) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// 解析 token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}
		token := parts[1]

		// 验证 token
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := g.userClient.ValidateToken(ctx, &userpb.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to validate token"})
			c.Abort()
			return
		}

		if !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": resp.Error})
			c.Abort()
			return
		}

		// 将用户 ID 和 Email 存入上下文
		c.Set("user_id", resp.UserId)
		c.Set("user_email", resp.Email)
		c.Next()
	}
}

// ============ 用户相关处理器 ============

// Register - 用户注册
func (g *Gateway) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := g.userClient.Register(ctx, &userpb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": resp.Token,
		"user":  resp.User,
	})
}

// Login - 用户登录
func (g *Gateway) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := g.userClient.Login(ctx, &userpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": resp.Token,
		"user":  resp.User,
	})
}

// GetUser - 获取用户信息
func (g *Gateway) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.User)
}

// UpdateUser - 更新用户信息
func (g *Gateway) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// 验证权限：只能修改自己的信息
	userID, exists := c.Get("user_id")
	if !exists || userID.(int64) != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	var req struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.userClient.UpdateUser(ctx, &userpb.UpdateUserRequest{
		Id:       id,
		Username: req.Username,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.User)
}

// ============ 图书相关处理器 ============

// GetBook - 获取图书详情
func (g *Gateway) GetBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.bookClient.GetBook(ctx, &bookpb.GetBookRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Book)
}

// ListBooks - 获取图书列表
func (g *Gateway) ListBooks(c *gin.Context) {
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("page_size", "10"), 10, 32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.bookClient.ListBooks(ctx, &bookpb.ListBooksRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": resp.Books,
		"total": resp.Total,
		"page":  resp.Page,
	})
}

// CreateBook - 创建图书
func (g *Gateway) CreateBook(c *gin.Context) {
	var req struct {
		Title       string  `json:"title" binding:"required"`
		Author      string  `json:"author" binding:"required"`
		ISBN        string  `json:"isbn" binding:"required"`
		Publisher   string  `json:"publisher"`
		Price       float64 `json:"price" binding:"required,gt=0"`
		Stock       int32   `json:"stock" binding:"required,gte=0"`
		Description string  `json:"description"`
		CoverImage  string  `json:"cover_image"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.bookClient.CreateBook(ctx, &bookpb.CreateBookRequest{
		Title:       req.Title,
		Author:      req.Author,
		Isbn:        req.ISBN,
		Publisher:   req.Publisher,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
		CoverImage:  req.CoverImage,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp.Book)
}

// UpdateBook - 更新图书
func (g *Gateway) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	var req struct {
		Title       string  `json:"title"`
		Author      string  `json:"author"`
		Publisher   string  `json:"publisher"`
		Price       float64 `json:"price"`
		Stock       int32   `json:"stock"`
		Description string  `json:"description"`
		CoverImage  string  `json:"cover_image"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.bookClient.UpdateBook(ctx, &bookpb.UpdateBookRequest{
		Id:          id,
		Title:       req.Title,
		Author:      req.Author,
		Publisher:   req.Publisher,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
		CoverImage:  req.CoverImage,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Book)
}

// DeleteBook - 删除图书
func (g *Gateway) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = g.bookClient.DeleteBook(ctx, &bookpb.DeleteBookRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Book %d deleted successfully", id)})
}

// SearchBooks - 搜索图书
func (g *Gateway) SearchBooks(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "keyword is required"})
		return
	}

	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("page_size", "10"), 10, 32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.bookClient.SearchBooks(ctx, &bookpb.SearchBooksRequest{
		Keyword:  keyword,
		Page:     int32(page),
		PageSize: int32(pageSize),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": resp.Books,
		"total": resp.Total,
		"page":  resp.Page,
	})
}
