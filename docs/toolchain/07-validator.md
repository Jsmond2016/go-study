---
title: Validator 验证
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["基础语法", "结构体"]
tags: ["Validator", "验证", "数据校验"]
---

# Validator 验证

Validator 是 Go 语言最流行的数据验证库，支持结构体标签验证。

## 📋 学习目标

- [ ] 理解数据验证的重要性
- [ ] 掌握 Validator 的基本使用
- [ ] 学会使用验证标签
- [ ] 理解自定义验证规则
- [ ] 掌握错误处理
- [ ] 了解验证最佳实践

## 🎯 Validator 简介

### 为什么选择 Validator

- **功能强大**: 支持丰富的验证规则
- **易于使用**: 结构体标签验证
- **性能优秀**: 编译时优化
- **扩展性强**: 支持自定义验证器
- **文档完善**: 官方文档详细

### 安装 Validator

```bash
go get github.com/go-playground/validator/v10
```

## 🚀 快速开始

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `validate:"required,min=2,max=50"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"required,gte=0,lte=150"`
	Password string `validate:"required,min=8"`
}

func main() {
	validate := validator.New()
	
	user := User{
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Age:      25,
		Password: "password123",
	}
	
	if err := validate.Struct(user); err != nil {
		fmt.Printf("验证失败: %v\n", err)
		return
	}
	
	fmt.Println("验证成功")
}
```

## 📝 验证标签

### 字符串验证

```go
type User struct {
	Name     string `validate:"required"`              // 必需
	Email    string `validate:"email"`                 // 邮箱格式
	URL      string `validate:"url"`                   // URL 格式
	Phone    string `validate:"numeric,len=11"`        // 数字，长度11
	Username string `validate:"alphanum,min=3,max=20"` // 字母数字，3-20字符
}
```

### 数值验证

```go
type Product struct {
	Price    float64 `validate:"required,gt=0"`        // 大于0
	Discount int     `validate:"gte=0,lte=100"`        // 0-100
	Stock    int     `validate:"min=0"`                // 最小0
	Rating   float64 `validate:"min=0,max=5"`          // 0-5
}
```

### 切片和数组验证

```go
type Order struct {
	Items    []string `validate:"required,min=1"`      // 至少1个元素
	Tags     []string `validate:"dive,required"`      // 每个元素必需
	IDs      []int    `validate:"dive,min=1"`         // 每个元素最小1
}
```

### 嵌套结构体验证

```go
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Zip    string `validate:"required,len=6"`
}

type User struct {
	Name    string  `validate:"required"`
	Address Address `validate:"required"`
}
```

## 🔧 自定义验证

### 自定义验证函数

```go
package main

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	
	// 至少包含一个大写字母
	hasUpper := false
	// 至少包含一个小写字母
	hasLower := false
	// 至少包含一个数字
	hasNumber := false
	
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		}
	}
	
	return hasUpper && hasLower && hasNumber
}

func main() {
	validate := validator.New()
	
	// 注册自定义验证器
	validate.RegisterValidation("strong_password", validatePassword)
	
	type User struct {
		Password string `validate:"required,strong_password"`
	}
	
	user := User{Password: "Password123"}
	if err := validate.Struct(user); err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println("验证成功")
	}
}
```

### 自定义错误消息

```go
func customErrorMessages(err error) map[string]string {
	errors := make(map[string]string)
	
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			tag := fieldError.Tag()
			
			switch tag {
			case "required":
				errors[field] = field + " 是必需的"
			case "email":
				errors[field] = field + " 必须是有效的邮箱地址"
			case "min":
				errors[field] = field + " 长度不能小于 " + fieldError.Param()
			case "max":
				errors[field] = field + " 长度不能大于 " + fieldError.Param()
			default:
				errors[field] = field + " 验证失败"
			}
		}
	}
	
	return errors
}
```

## 🏃‍♂️ 实践应用

### 在 Gin 中使用

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=0,lte=150"`
	Password string `json:"password" binding:"required,min=8"`
}

func main() {
	r := gin.Default()
	
	r.POST("/users", func(c *gin.Context) {
		var req CreateUserRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			// 处理验证错误
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				errors := make(map[string]string)
				for _, fieldError := range validationErrors {
					errors[fieldError.Field()] = getErrorMessage(fieldError)
				}
				c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "用户创建成功"})
	})
	
	r.Run(":8080")
}

func getErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "该字段是必需的"
	case "email":
		return "必须是有效的邮箱地址"
	case "min":
		return "长度不能小于 " + fieldError.Param()
	case "max":
		return "长度不能大于 " + fieldError.Param()
	case "gte":
		return "值不能小于 " + fieldError.Param()
	case "lte":
		return "值不能大于 " + fieldError.Param()
	default:
		return "验证失败"
	}
}
```

## ⚠️ 注意事项

### 1. 验证顺序

```go
// 验证按标签顺序执行
type User struct {
	Email string `validate:"required,email"` // 先检查必需，再检查格式
}
```

### 2. 性能考虑

```go
// ✅ 复用 validator 实例
validate := validator.New()

// ❌ 不要每次都创建新实例
for _, user := range users {
	validate := validator.New() // 错误
	validate.Struct(user)
}
```

### 3. 错误处理

```go
// ✅ 提供友好的错误信息
if err := validate.Struct(user); err != nil {
	// 格式化错误信息
	errors := formatValidationErrors(err)
	return errors
}
```

## 📚 扩展阅读

- [Validator 官方文档](https://github.com/go-playground/validator)
- [验证标签列表](https://github.com/go-playground/validator#baked-in-validations)

## ⏭️ 下一章节

[CORS 跨域](./08-cors.md) → 学习跨域处理

---

**💡 提示**: Validator 是数据验证的标准库，掌握它可以让数据验证变得简单高效！

