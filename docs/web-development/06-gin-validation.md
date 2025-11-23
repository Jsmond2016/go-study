---
title: 数据验证
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["Gin 基础", "结构体"]
tags: ["Gin", "验证", "绑定", "validator"]
---

# 数据验证

Gin 使用 `validator` 包进行数据验证，支持丰富的验证规则。

## 📋 学习目标

- [ ] 理解数据验证的重要性
- [ ] 掌握结构体标签验证
- [ ] 学会自定义验证规则
- [ ] 理解不同绑定方式
- [ ] 掌握验证错误处理

## 🎯 基本验证

### 结构体标签

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=0,lte=150"`
	Password string `json:"password" binding:"required,min=6"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.Run(":8080")
}
```

### 常用验证标签

```go
type User struct {
	// 必需字段
	Name string `binding:"required"`

	// 字符串长度
	Title string `binding:"min=1,max=100"`

	// 数值范围
	Age int `binding:"gte=0,lte=150"`

	// 邮箱格式
	Email string `binding:"email"`

	// URL 格式
	Website string `binding:"url"`

	// 枚举值
	Status string `binding:"oneof=pending completed cancelled"`

	// 正则表达式
	Phone string `binding:"regexp=^1[3-9]\\d{9}$"`

	// 忽略空值
	Description string `binding:"omitempty,min=0,max=500"`
}
```

## 🔧 绑定方式

### JSON 绑定

```go
r.POST("/user", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
})
```

### 表单绑定

```go
r.POST("/form", func(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
})
```

### URI 绑定

```go
type User struct {
	ID int `uri:"id" binding:"required"`
}

r.GET("/user/:id", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
})
```

### 查询参数绑定

```go
type Query struct {
	Page     int    `form:"page" binding:"gte=1"`
	PageSize int    `form:"page_size" binding:"gte=1,lte=100"`
	Keyword  string `form:"keyword"`
}

r.GET("/search", func(c *gin.Context) {
	var query Query
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, query)
})
```

## 🎨 自定义验证

### 自定义验证函数

```go
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("custom_rule", customValidation)
}

func customValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	// 自定义验证逻辑
	return len(value) > 5
}

type User struct {
	Name string `binding:"required,custom_rule"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.Run(":8080")
}
```

## 🏃‍♂️ 实践应用

### 完整的验证示例

```go
package main

import (
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Age      int    `json:"age" binding:"gte=18,lte=100"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty" binding:"omitempty,min=2,max=50"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
	Age   *int    `json:"age,omitempty" binding:"omitempty,gte=18,lte=100"`
}

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": formatValidationError(err)})
			return
		}
		c.JSON(201, gin.H{"message": "用户创建成功"})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var req UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": formatValidationError(err)})
			return
		}
		c.JSON(200, gin.H{"message": "用户更新成功"})
	})

	r.Run(":8080")
}

func formatValidationError(err error) string {
	// 格式化验证错误
	return err.Error()
}
```

## ⚠️ 注意事项

### 1. 验证顺序

```go
// 验证按标签顺序执行
type User struct {
	Email string `binding:"required,email"` // 先检查必需，再检查格式
}
```

### 2. 错误处理

```go
// ✅ 提供友好的错误信息
if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(400, gin.H{
		"error": "验证失败",
		"details": err.Error(),
	})
	return
}
```

### 3. 性能考虑

```go
// ✅ 验证是同步操作，避免在验证中执行耗时操作
```

## 📚 扩展阅读

- [validator 包文档](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [Gin 绑定文档](https://gin-gonic.com/docs/examples/binding-and-validation/)

## ⏭️ 下一章节

[认证授权](./07-gin-auth.md) → 学习 Gin 的认证和授权机制

---

**💡 提示**: 数据验证是 API 安全的基础，严格的验证可以防止无效数据和潜在的安全问题！

