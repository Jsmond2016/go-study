---
title: Gin 路由
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["Gin 基础"]
tags: ["Gin", "路由", "参数", "路由组"]
---

# Gin 路由

Gin 提供了强大的路由功能，支持路径参数、查询参数、路由组等。

## 📋 学习目标

- [ ] 掌握路由的基本配置
- [ ] 理解路径参数和查询参数
- [ ] 学会使用路由组
- [ ] 掌握路由的优先级
- [ ] 了解路由的最佳实践

## 🎯 基本路由

### HTTP 方法

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.String(200, "GET 方法")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "POST 方法")
	})

	r.PUT("/put", func(c *gin.Context) {
		c.String(200, "PUT 方法")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "DELETE 方法")
	})

	r.PATCH("/patch", func(c *gin.Context) {
		c.String(200, "PATCH 方法")
	})

	r.HEAD("/head", func(c *gin.Context) {
		c.String(200, "HEAD 方法")
	})

	r.OPTIONS("/options", func(c *gin.Context) {
		c.String(200, "OPTIONS 方法")
	})

	r.Run(":8080")
}
```

### 任意方法

```go
r.Any("/any", func(c *gin.Context) {
	c.String(200, "任意方法")
})

// 处理所有未匹配的路由
r.NoRoute(func(c *gin.Context) {
	c.JSON(404, gin.H{"message": "路由不存在"})
})
```

## 🔗 路径参数

### 必需参数

```go
r := gin.Default()

// 单个参数
r.GET("/user/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "用户ID: %s", id)
})

// 多个参数
r.GET("/user/:id/posts/:postId", func(c *gin.Context) {
	id := c.Param("id")
	postId := c.Param("postId")
	c.String(200, "用户ID: %s, 文章ID: %s", id, postId)
})
```

### 可选参数

```go
r := gin.Default()

// 通配符参数
r.GET("/user/:id/*action", func(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.String(200, "ID: %s, Action: %s", id, action)
})
```

## 🔍 查询参数

### 获取查询参数

```go
r := gin.Default()

r.GET("/search", func(c *gin.Context) {
	// 必需参数
	keyword := c.Query("q")

	// 可选参数（带默认值）
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	// 检查参数是否存在
	if keyword == "" {
		c.JSON(400, gin.H{"error": "缺少查询参数 q"})
		return
	}

	c.JSON(200, gin.H{
		"keyword":  keyword,
		"page":     page,
		"page_size": pageSize,
	})
})
```

### 获取多个值

```go
r.GET("/tags", func(c *gin.Context) {
	tags := c.QueryArray("tag")
	c.JSON(200, gin.H{"tags": tags})
})

// 请求: /tags?tag=go&tag=web&tag=api
```

### 表单参数

```go
r := gin.Default()

// 获取表单参数
r.POST("/form", func(c *gin.Context) {
	// 必需参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 可选参数（带默认值）
	email := c.DefaultPostForm("email", "default@example.com")

	// 获取数组参数
	hobbys := c.PostFormArray("hobby")

	c.JSON(200, gin.H{
		"username": username,
		"password": password,
		"email":    email,
		"hobbys":   hobbys,
	})
})
```

### 文件上传

```go
r := gin.Default()

// 设置最大上传文件大小
r.MaxMultipartMemory = 8 << 20 // 8 MiB

// 单文件上传
r.POST("/upload", func(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 保存文件
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
})

// 多文件上传
r.POST("/uploads", func(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Uploaded successfully %d files", len(files)),
	})
})
```

## 📦 路由组

### 基本路由组

```go
r := gin.Default()

// 创建路由组
v1 := r.Group("/api/v1")
{
	v1.GET("/users", getUsers)
	v1.POST("/users", createUser)
	v1.GET("/users/:id", getUser)
}

v2 := r.Group("/api/v2")
{
	v2.GET("/users", getUsersV2)
}
```

### 嵌套路由组

```go
r := gin.Default()

api := r.Group("/api")
{
	v1 := api.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", getUsers)
			users.POST("", createUser)
			users.GET("/:id", getUser)
		}
	}
}
```

### 路由组中间件

```go
r := gin.Default()

// 路由组使用中间件
v1 := r.Group("/api/v1")
v1.Use(authMiddleware())
{
	v1.GET("/users", getUsers)
	v1.POST("/users", createUser)
}

// 公开路由组（不使用中间件）
public := r.Group("/public")
{
	public.GET("/info", getInfo)
}
```

## 🎯 路由优先级

### 路由匹配规则

```go
r := gin.Default()

// 精确匹配优先
r.GET("/user/list", listUsers)

// 参数匹配
r.GET("/user/:id", getUser)

// 通配符匹配（最低优先级）
r.GET("/user/*action", userAction)
```

### 路由冲突

```go
// ❌ 错误：路由冲突
r.GET("/user/:id", handler1)
r.GET("/user/new", handler2) // 会被 :id 匹配

// ✅ 正确：将具体路由放在前面
r.GET("/user/new", handler2)
r.GET("/user/:id", handler1)
```

## 🔧 高级路由

### 参数绑定

Gin 支持将请求参数自动绑定到结构体，支持多种 Content-Type：

```go
type Login struct {
	Name     string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// JSON 绑定
r.POST("/login-json", func(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "登录成功"})
})

// Form 表单绑定
r.POST("/login-form", func(c *gin.Context) {
	var form Login
	// 根据 Content-Type 自动推断绑定方式
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "登录成功"})
})

// URI 绑定
r.GET("/login/:name/:password", func(c *gin.Context) {
	var login Login
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, login)
})
```

### 路由绑定

```go
type User struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name"`
}

r.GET("/user/:id/:name", func(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
})
```

### 路由重定向

```go
r.GET("/old", func(c *gin.Context) {
	c.Redirect(301, "/new")
})

r.GET("/new", func(c *gin.Context) {
	c.String(200, "新页面")
})
```

## 🏃‍♂️ 实践应用

### RESTful API 路由设计

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// 用户资源
		users := api.Group("/users")
		{
			users.GET("", listUsers)           // GET /api/v1/users
			users.POST("", createUser)         // POST /api/v1/users
			users.GET("/:id", getUser)         // GET /api/v1/users/:id
			users.PUT("/:id", updateUser)      // PUT /api/v1/users/:id
			users.DELETE("/:id", deleteUser)   // DELETE /api/v1/users/:id
		}

		// 文章资源
		posts := api.Group("/posts")
		{
			posts.GET("", listPosts)
			posts.POST("", createPost)
			posts.GET("/:id", getPost)
			posts.PUT("/:id", updatePost)
			posts.DELETE("/:id", deletePost)
		}
	}

	r.Run(":8080")
}
```

### 版本控制

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

## ⚠️ 注意事项

### 1. 路由顺序

```go
// ✅ 具体路由在前
r.GET("/user/new", handler1)
r.GET("/user/:id", handler2)
```

### 2. 参数验证

```go
// ✅ 验证路径参数
id := c.Param("id")
if id == "" {
	c.JSON(400, gin.H{"error": "ID不能为空"})
	return
}
```

### 3. 路由组织

```go
// ✅ 使用路由组组织代码
api := r.Group("/api")
{
	api.GET("/users", getUsers)
	api.GET("/posts", getPosts)
}
```

## 📚 扩展阅读

- [Gin 路由文档](https://gin-gonic.com/docs/examples/routing/)
- [HTTP 路由最佳实践](https://restfulapi.net/resource-naming/)

## ⏭️ 下一章节

[Gin 中间件](./04-gin-middleware.md) → 学习 Gin 中间件的开发和使用

---

**💡 提示**: 良好的路由设计是 RESTful API 的基础，遵循 REST 规范可以让 API 更加清晰和易用！

