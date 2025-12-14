---
title: Chapter 5 - API Gateway 实现
difficulty: intermediate  
duration: "4-5小时"
tags: ["Gin", "HTTP", "gRPC", "中间件", "REST API"]
---

# Chapter 5: API Gateway 实现

本章实现 API Gateway，作为前端和后端微服务之间的桥梁。

## 📋 学习目标

- [x] 掌握 Gin 框架基础
- [x] 实现 HTTP → gRPC 转换
- [x] 编写 JWT 认证中间件
- [x] 配置 CORS 跨域
- [x] 设计 RESTful API

## 🏗️ API Gateway 架构

### 核心职责

```
                   API Gateway
┌──────────────────────────────────────────────┐
│  HTTP 请求 → 认证 → gRPC 调用 → HTTP 响应   │
└──────────────────────────────────────────────┘
```

**功能列表**：
1. **协议转换** - HTTP/JSON → gRPC/Protobuf
2. **统一认证** - JWT Token 验证
3. **路由管理** - 请求分发到不同服务
4. **错误处理** - 统一错误响应格式
5. **CORS 处理** - 跨域资源共享

## 💻 核心实现

### 1. Gateway 结构体

```go
type Gateway struct {
	bookClient bookpb.BookServiceClient
	userClient userpb.UserServiceClient
}
```

### 2. 路由配置

```go
func main() {
	// 连接服务
	bookConn, _ := grpc.NewClient(bookServiceAddr, 
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	userConn, _ := grpc.NewClient(userServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	gateway := &Gateway{
		bookClient: bookpb.NewBookServiceClient(bookConn),
		userClient: userpb.NewUserServiceClient(userConn),
	}

	r := gin.Default()
	r.Use(cors.Default())  // CORS 中间件

	// 公开路由
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", gateway.Register)
		userGroup.POST("/login", gateway.Login)
	}

	// 需认证路由
	authGroup := r.Group("/api")
	authGroup.Use(gateway.AuthMiddleware())  // JWT 验证
	{
		authGroup.GET("/users/:id", gateway.GetUser)
		authGroup.PUT("/users/:id", gateway.UpdateUser)
		authGroup.GET("/books", gateway.ListBooks)
		authGroup.GET("/books/:id", gateway.GetBook)
		// ... 更多路由
	}

	r.Run(":8080")
}
```

### 3. JWT 认证中间件

```go
func (g *Gateway) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// 2. 解析 token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}
		token := parts[1]

		// 3. 调用 User Service 验证
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := g.userClient.ValidateToken(ctx, &userpb.ValidateTokenRequest{
			Token: token,
		})

		if err != nil || !resp.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 4. 存储用户信息到上下文
		c.Set("user_id", resp.UserId)
		c.Set("user_email", resp.Email)
		c.Next()
	}
}
```

### 4. HTTP → gRPC 转换示例

#### 用户注册

```go
func (g *Gateway) Register(c *gin.Context) {
	// 1. 解析 JSON 请求
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 2. 调用 gRPC 服务
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := g.userClient.Register(ctx, &userpb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 3. 返回 JSON 响应
	c.JSON(201, gin.H{
		"token": resp.Token,
		"user":  resp.User,
	})
}
```

#### 获取图书列表

```go
func (g *Gateway) ListBooks(c *gin.Context) {
	// 1. 解析查询参数
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("page_size", "10"), 10, 32)

	// 2. 调用 Book Service
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.bookClient.ListBooks(ctx, &bookpb.ListBooksRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 3. 返回结果
	c.JSON(200, gin.H{
		"books": resp.Books,
		"total": resp.Total,
		"page":  resp.Page,
	})
}
```

## ✅ 完整 API 列表

### 用户 API（公开）

| Method | Path | 说明 |
|--------|------|------|
| POST | `/api/users/register` | 用户注册 |
| POST | `/api/users/login` | 用户登录 |

### 用户 API（需认证）

| Method | Path | 说明 |
|--------|------|------|
| GET | `/api/users/:id` | 获取用户 |
| PUT | `/api/users/:id` | 更新用户 |

### 图书 API（需认证）

| Method | Path | 说明 |
|--------|------|------|
| GET | `/api/books` | 图书列表 |
| GET | `/api/books/:id` | 图书详情 |
| POST | `/api/books` | 创建图书 |
| PUT | `/api/books/:id` | 更新图书 |
| DELETE | `/api/books/:id` | 删除图书 |
| GET | `/api/books/search?keyword=Go` | 搜索图书 |

## 🧪 测试

```bash
# 1. 注册
curl -X POST http://localhost:8080/api/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","email":"alice@example.com","password":"password123"}'

# 2. 登录
TOKEN=$(curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password123"}' \
  | jq -r '.token')

# 3. 获取图书（需要 Token）
curl -X GET "http://localhost:8080/api/books?page=1&page_size=10" \
  -H "Authorization: Bearer $TOKEN"
```

## 🎯 下一步

**[👉 Chapter 6: 集成测试](/microservices/bookstore-project/06-integration)**

---

**🎉 API Gateway 完成！三个微服务都实现了！**
