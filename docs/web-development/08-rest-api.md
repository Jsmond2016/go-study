---
title: REST API 设计
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["Gin 基础", "数据验证", "认证授权"]
tags: ["REST", "API", "设计", "最佳实践"]
---

# REST API 设计

RESTful API 是一种设计风格，遵循 REST 原则可以让 API 更加清晰、易用和可维护。

## 📋 学习目标

- [ ] 理解 REST 设计原则
- [ ] 掌握 RESTful API 设计规范
- [ ] 学会设计资源路由
- [ ] 理解 HTTP 状态码的使用
- [ ] 掌握 API 版本控制
- [ ] 了解 API 文档编写

## 🎯 REST 原则

### REST 核心概念

- **资源（Resource）**: API 操作的对象
- **URI**: 资源的唯一标识
- **HTTP 方法**: 操作资源的方式
- **状态码**: 操作结果的状态

### RESTful 设计原则

1. **使用名词表示资源**
2. **使用 HTTP 方法表示操作**
3. **使用状态码表示结果**
4. **使用 JSON 格式**
5. **无状态通信**

## 🛣️ 路由设计

### 资源路由

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	// 用户资源
	users := r.Group("/api/v1/users")
	{
		users.GET("", listUsers)           // GET /api/v1/users
		users.POST("", createUser)         // POST /api/v1/users
		users.GET("/:id", getUser)         // GET /api/v1/users/:id
		users.PUT("/:id", updateUser)      // PUT /api/v1/users/:id
		users.DELETE("/:id", deleteUser)   // DELETE /api/v1/users/:id
	}
	
	// 文章资源
	posts := r.Group("/api/v1/posts")
	{
		posts.GET("", listPosts)
		posts.POST("", createPost)
		posts.GET("/:id", getPost)
		posts.PUT("/:id", updatePost)
		posts.DELETE("/:id", deletePost)
		
		// 嵌套资源
		posts.GET("/:id/comments", getPostComments)
		posts.POST("/:id/comments", createComment)
	}
	
	r.Run(":8080")
}
```

### 路由命名规范

```go
// ✅ 好的设计
GET    /api/v1/users          // 获取用户列表
GET    /api/v1/users/:id      // 获取指定用户
POST   /api/v1/users          // 创建用户
PUT    /api/v1/users/:id      // 更新用户
DELETE /api/v1/users/:id      // 删除用户

// ❌ 不好的设计
GET    /api/v1/getUsers
POST   /api/v1/createUser
GET    /api/v1/user/:id/get
```

## 📊 HTTP 状态码

### 常用状态码

```go
// 成功响应
c.JSON(200, data)  // OK - 成功
c.JSON(201, data)  // Created - 创建成功
c.JSON(204, nil)   // No Content - 删除成功

// 客户端错误
c.JSON(400, gin.H{"error": "请求参数错误"})  // Bad Request
c.JSON(401, gin.H{"error": "未授权"})        // Unauthorized
c.JSON(403, gin.H{"error": "禁止访问"})      // Forbidden
c.JSON(404, gin.H{"error": "资源不存在"})    // Not Found
c.JSON(409, gin.H{"error": "资源冲突"})      // Conflict

// 服务器错误
c.JSON(500, gin.H{"error": "服务器内部错误"}) // Internal Server Error
```

### 状态码使用示例

```go
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// 创建用户
	if err := userService.Create(&user); err != nil {
		if err == ErrUserExists {
			c.JSON(409, gin.H{"error": "用户已存在"})
		} else {
			c.JSON(500, gin.H{"error": "创建失败"})
		}
		return
	}
	
	c.JSON(201, user)
}
```

## 📝 响应格式

### 统一响应格式

```go
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

func successResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	})
}

func errorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := APIResponse{
		Success:   false,
		Message:   message,
		Timestamp: time.Now(),
	}
	if err != nil {
		response.Error = err.Error()
	}
	c.JSON(statusCode, response)
}
```

### 分页响应

```go
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
}

func paginatedResponse(c *gin.Context, items interface{}, page, pageSize, total int) {
	totalPages := (total + pageSize - 1) / pageSize
	c.JSON(200, PaginatedResponse{
		Items:      items,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	})
}
```

## 🔄 API 版本控制

### URL 版本控制

```go
r := gin.Default()

// v1 API
v1 := r.Group("/api/v1")
{
	v1.GET("/users", getUsersV1)
}

// v2 API
v2 := r.Group("/api/v2")
{
	v2.GET("/users", getUsersV2)
}
```

### Header 版本控制

```go
func versionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := c.GetHeader("API-Version")
		if version == "" {
			version = "v1"
		}
		c.Set("api_version", version)
		c.Next()
	}
}
```

## 🏃‍♂️ 实践应用

### 完整的 REST API 示例

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User
var nextID = 1

func main() {
	r := gin.Default()
	
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", listUsers)
			users.POST("", createUser)
			users.GET("/:id", getUser)
			users.PUT("/:id", updateUser)
			users.DELETE("/:id", deleteUser)
		}
	}
	
	r.Run(":8080")
}

func listUsers(c *gin.Context) {
	c.JSON(200, users)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = nextID
	nextID++
	users = append(users, user)
	c.JSON(201, user)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	// 查找用户逻辑
	c.JSON(200, gin.H{"id": id})
}

func updateUser(c *gin.Context) {
	// 更新逻辑
}

func deleteUser(c *gin.Context) {
	// 删除逻辑
}
```

## ⚠️ 最佳实践

### 1. 资源命名

```go
// ✅ 使用复数名词
/api/v1/users
/api/v1/posts

// ❌ 避免使用动词
/api/v1/getUsers
/api/v1/createUser
```

### 2. HTTP 方法使用

```go
// ✅ 正确使用 HTTP 方法
GET    - 获取资源
POST   - 创建资源
PUT    - 更新资源（完整更新）
PATCH  - 更新资源（部分更新）
DELETE - 删除资源
```

### 3. 错误处理

```go
// ✅ 提供清晰的错误信息
c.JSON(400, gin.H{
	"error": "验证失败",
	"details": []string{
		"姓名不能为空",
		"邮箱格式不正确",
	},
})
```

## 📚 扩展阅读

- [RESTful API 设计指南](https://restfulapi.net/)
- [HTTP 状态码](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status)

## ⏭️ 下一阶段

完成 Web 开发学习后，可以进入：

- [实战项目](../projects/) - 构建完整项目
- [微服务](../microservices/) - 分布式系统

---

**💡 提示**: RESTful API 设计是 Web 开发的基础，遵循 REST 原则可以让 API 更加清晰和易用！

