---
title: Gin 基础
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["HTTP 服务器", "基础语法"]
tags: ["Gin", "框架", "Web", "路由"]
---

# Gin 基础

Gin 是一个高性能的 Go Web 框架，提供了简洁的 API 和丰富的功能。

## 📋 学习目标

- [ ] 理解 Gin 框架的特点
- [ ] 掌握 Gin 的基本使用
- [ ] 学会创建路由和处理请求
- [ ] 理解 Gin 的响应类型
- [ ] 掌握 Gin 的中间件使用
- [ ] 了解 Gin 的最佳实践

## 🎯 Gin 简介

### 为什么选择 Gin

- **高性能**: 基于 httprouter，性能优秀
- **简洁**: API 设计简洁易用
- **功能丰富**: 支持路由、中间件、参数绑定等
- **活跃**: 社区活跃，文档完善

### 安装 Gin

```bash
go get -u github.com/gin-gonic/gin
```

## 🚀 快速开始

### 第一个 Gin 应用

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	r.Run(":8080")
}
```

### Gin 模式

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	// 开发模式（默认）
	gin.SetMode(gin.DebugMode)

	// 发布模式
	gin.SetMode(gin.ReleaseMode)

	// 测试模式
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	// ...
}
```

## 📥 处理请求

### 获取参数

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路径参数
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "用户ID: %s", id)
	})

	// 查询参数
	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("q")
		page := c.DefaultQuery("page", "1")
		c.String(200, "搜索: %s, 页码: %s", keyword, page)
	})

	// 表单参数
	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.DefaultPostForm("email", "default@example.com")
		c.String(200, "姓名: %s, 邮箱: %s", name, email)
	})

	r.Run(":8080")
}
```

### 参数绑定

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Email string `form:"email" json:"email" binding:"required,email"`
	Age   int    `form:"age" json:"age" binding:"gte=0,lte=150"`
}

func main() {
	r := gin.Default()

	// JSON 绑定
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// 表单绑定
	r.POST("/form", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}
```

## 📤 响应类型

### JSON 响应

```go
r.GET("/json", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"name": "张三",
			"age":  30,
		},
	})
})
```

### 字符串响应

```go
r.GET("/string", func(c *gin.Context) {
	c.String(200, "Hello, Gin!")
})
```

### HTML 响应

```go
r.GET("/html", func(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Gin 示例",
	})
})
```

### XML 响应

```go
r.GET("/xml", func(c *gin.Context) {
	c.XML(200, gin.H{
		"message": "success",
	})
})
```

### YAML 响应

```go
r.GET("/yaml", func(c *gin.Context) {
	c.YAML(200, gin.H{
		"message": "success",
		"status":  200,
	})
})
```

### Protobuf 响应

```go
import "github.com/gin-gonic/gin/testdata/protoexample"

r.GET("/protobuf", func(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	// 注意：响应为二进制数据
	c.ProtoBuf(200, data)
})
```

### 文件响应

```go
r.GET("/file", func(c *gin.Context) {
	c.File("./file.txt")
})

// 文件下载
r.GET("/download", func(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=file.txt")
	c.File("./file.txt")
})
```

### 自定义结构体响应

```go
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

r.GET("/custom", func(c *gin.Context) {
	response := Response{
		Code:    0,
		Data:    gin.H{"name": "Tony", "age": 30},
		Message: "success",
	}
	c.JSON(200, response)
})
```

## ⚡ 异步处理

### 同步处理

```go
r.GET("/sync", func(c *gin.Context) {
	// 同步处理，会阻塞直到完成
	time.Sleep(5 * time.Second)
	log.Println("Done! in path " + c.Request.URL.Path)
	c.JSON(200, gin.H{"message": "同步处理完成，耗时 5s"})
})
```

### 异步处理

```go
r.GET("/async", func(c *gin.Context) {
	// 创建上下文的副本（只读）
	cCp := c.Copy()

	// 在 goroutine 中处理耗时任务
	go func() {
		time.Sleep(5 * time.Second)
		// 注意：必须使用只读上下文
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()

	// 立即返回响应
	c.JSON(200, gin.H{"message": "异步请求已提交，请稍后查看结果"})
})
```

**注意事项**：
- 异步处理中必须使用 `c.Copy()` 创建只读上下文
- 原始上下文 `c` 在 handler 返回后可能被回收
- 适合处理耗时任务，如文件上传、邮件发送等

## 🛣️ 路由

### 基本路由

```go
r := gin.Default()

r.GET("/get", handler)
r.POST("/post", handler)
r.PUT("/put", handler)
r.DELETE("/delete", handler)
r.PATCH("/patch", handler)
r.HEAD("/head", handler)
r.OPTIONS("/options", handler)
```

### 路由组

```go
r := gin.Default()

v1 := r.Group("/api/v1")
{
	v1.GET("/users", getUsers)
	v1.POST("/users", createUser)
	v1.GET("/users/:id", getUser)
	v1.PUT("/users/:id", updateUser)
	v1.DELETE("/users/:id", deleteUser)
}
```

### 路由参数

```go
// 必需参数
r.GET("/user/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "用户ID: %s", id)
})

// 可选参数
r.GET("/user/:id/*action", func(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.String(200, "ID: %s, Action: %s", id, action)
})
```

## 🔧 中间件

### 使用中间件

```go
r := gin.Default()

// 全局中间件
r.Use(loggingMiddleware())

// 路由组中间件
v1 := r.Group("/api/v1")
v1.Use(authMiddleware())
{
	v1.GET("/users", getUsers)
}

// 单个路由中间件
r.GET("/protected", authMiddleware(), protectedHandler)
```

### 创建中间件

```go
func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
```

## 🏃‍♂️ 实践应用

### 完整的用户管理 API

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

var users []User
var nextID = 1

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUser)
		api.POST("/users", createUser)
		api.PUT("/users/:id", updateUser)
		api.DELETE("/users/:id", deleteUser)
	}

	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	// 查找用户逻辑
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = nextID
	nextID++
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	// 更新逻辑
}

func deleteUser(c *gin.Context) {
	// 删除逻辑
}
```

## ⚠️ 注意事项

### 1. 错误处理

```go
// ✅ 总是检查绑定错误
if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
```

### 2. 中间件顺序

```go
// 中间件按添加顺序执行
r.Use(middleware1())
r.Use(middleware2())
// 执行顺序: middleware1 -> middleware2 -> handler
```

### 3. 性能优化

```go
// ✅ 生产环境使用 ReleaseMode
gin.SetMode(gin.ReleaseMode)

// ✅ 使用路由组组织代码
api := r.Group("/api")
```

## 📚 扩展阅读

- [Gin 官方文档](https://gin-gonic.com/)
- [Gin GitHub](https://github.com/gin-gonic/gin)

## ⏭️ 下一章节

[Gin 路由](./03-gin-routing.md) → 深入学习 Gin 路由配置

---

**💡 提示**: Gin 是 Go 语言最流行的 Web 框架之一，掌握它对于构建 Web 应用非常重要！

