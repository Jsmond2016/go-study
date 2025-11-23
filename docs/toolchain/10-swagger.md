---
title: Swagger API 文档
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["Web 开发", "REST API"]
tags: ["Swagger", "API 文档", "OpenAPI", "swaggo"]
---

# Swagger API 文档

Swagger 是 API 文档生成和测试工具，可以帮助开发者设计、构建、记录和使用 RESTful API。

## 📋 学习目标

- [ ] 理解 API 文档的重要性
- [ ] 掌握 Swagger/OpenAPI 规范
- [ ] 学会使用 swaggo 生成文档
- [ ] 掌握注释编写规范
- [ ] 理解文档部署和访问
- [ ] 了解最佳实践

## 🎯 Swagger 简介

### 为什么需要 API 文档

- **团队协作**: 前后端开发分离，需要清晰的接口文档
- **接口测试**: 可以直接在文档中测试接口
- **版本管理**: 文档随代码更新，保持同步
- **降低沟通成本**: 减少口头沟通，提高效率

### Swagger vs OpenAPI

- **Swagger**: 工具集和规范
- **OpenAPI**: 规范标准（原 Swagger 规范）
- **swaggo**: Go 语言的 Swagger 实现

### 安装工具

```bash
# 安装 swag
go install github.com/swaggo/swag/cmd/swag@latest

# 安装 Gin Swagger
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

## 🚀 快速开始

### 基本配置

```go
package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API 文档
// @version         1.0
// @description     这是一个示例 API 文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	r := gin.Default()
	
	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run(":8080")
}
```

## 📝 注释规范

### API 注释

```go
// @Summary      创建用户
// @Description  创建一个新用户
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserRequest  true  "用户信息"
// @Success      200   {object}  Response{data=User}
// @Failure      400   {object}  Response
// @Failure      500   {object}  Response
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	// 处理逻辑
}
```

### 结构体注释

```go
// User 用户信息
type User struct {
	ID       int    `json:"id" example:"1"`                    // 用户ID
	Name     string `json:"name" example:"张三"`                // 用户名
	Email    string `json:"email" example:"zhangsan@example.com"` // 邮箱
	Age      int    `json:"age" example:"25"`                  // 年龄
	CreateAt string `json:"created_at" example:"2024-01-01T00:00:00Z"` // 创建时间
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"张三"`      // 用户名
	Email    string `json:"email" binding:"required,email" example:"zhangsan@example.com"` // 邮箱
	Age      int    `json:"age" binding:"required,gte=0,lte=150" example:"25"` // 年龄
	Password string `json:"password" binding:"required,min=8" example:"password123"` // 密码
}

// Response 通用响应
type Response struct {
	Code    int         `json:"code" example:"200"`    // 状态码
	Message string      `json:"message" example:"成功"` // 消息
	Data    interface{} `json:"data,omitempty"`         // 数据
}
```

## 🔧 完整示例

### 完整的 API 文档

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           用户管理 API
// @version         1.0
// @description     用户管理系统的 API 文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 输入 "Bearer {token}"

func main() {
	r := gin.Default()
	
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", ListUsers)
			users.POST("", CreateUser)
			users.GET("/:id", GetUser)
			users.PUT("/:id", UpdateUser)
			users.DELETE("/:id", DeleteUser)
		}
	}
	
	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run(":8080")
}

// ListUsers 获取用户列表
// @Summary      获取用户列表
// @Description  获取所有用户列表，支持分页
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "页码"     default(1)
// @Param        page_size query     int     false  "每页数量"  default(10)
// @Success      200       {object}  Response{data=[]User}
// @Failure      500       {object}  Response
// @Router       /users [get]
func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功",
		"data":    []User{},
	})
}

// CreateUser 创建用户
// @Summary      创建用户
// @Description  创建一个新用户
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserRequest  true  "用户信息"
// @Success      201   {object}  Response{data=User}
// @Failure      400   {object}  Response
// @Failure      500   {object}  Response
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "创建成功",
		"data":    User{},
	})
}

// GetUser 获取用户
// @Summary      获取用户
// @Description  根据ID获取用户信息
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "用户ID"
// @Success      200  {object}  Response{data=User}
// @Failure      404  {object}  Response
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功",
		"data":    User{ID: 1},
	})
}

// UpdateUser 更新用户
// @Summary      更新用户
// @Description  更新用户信息
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        id    path      int                true  "用户ID"
// @Param        user  body      UpdateUserRequest   true  "用户信息"
// @Success      200   {object}  Response{data=User}
// @Failure      400   {object}  Response
// @Failure      404   {object}  Response
// @Router       /users/{id} [put]
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

// DeleteUser 删除用户
// @Summary      删除用户
// @Description  根据ID删除用户
// @Tags         用户管理
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "用户ID"
// @Success      200  {object}  Response
// @Failure      404  {object}  Response
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
```

## 🔐 认证配置

### JWT 认证

```go
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 输入 "Bearer {token}"

// 在需要认证的接口上添加
// @Security Bearer
func ProtectedHandler(c *gin.Context) {
	// 处理逻辑
}
```

## 📊 生成文档

### 生成命令

```bash
# 生成文档
swag init

# 指定输出目录
swag init -o ./docs

# 解析依赖
swag init --parseDependency

# 解析内部依赖
swag init --parseInternal
```

### 目录结构

```
project/
├── main.go
├── docs/
│   ├── swagger.json
│   ├── swagger.yaml
│   └── docs.go
└── go.mod
```

## 🏃‍♂️ 实践应用

### 在项目中使用

```go
package main

import (
	_ "myproject/docs" // 导入生成的文档
	
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	
	// API 路由
	setupRoutes(r)
	
	// Swagger 文档（仅开发环境）
	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	
	r.Run(":8080")
}
```

## ⚠️ 注意事项

### 1. 注释格式

```go
// ✅ 正确的注释格式
// @Summary 创建用户
// @Tags 用户管理

// ❌ 错误的注释格式
// @Summary创建用户  // 缺少空格
```

### 2. 类型定义

```go
// ✅ 使用 example 标签
type User struct {
	ID int `json:"id" example:"1"`
}

// ✅ 使用 swaggertype 标签
type Status string // @Description 状态 @Enum(active,inactive) @Default(active)
```

### 3. 文档更新

```bash
# ✅ 每次修改注释后重新生成
swag init

# ✅ 提交文档到版本控制
git add docs/
```

## 📚 扩展阅读

- [Swagger 官方文档](https://swagger.io/)
- [swaggo 文档](https://github.com/swaggo/swag)
- [OpenAPI 规范](https://swagger.io/specification/)

## ⏭️ 下一章节

[Redis 缓存](./11-redis.md) → 学习 Redis 缓存使用

---

**💡 提示**: Swagger 是 API 开发中必不可少的工具，良好的文档可以大大提高开发效率！

