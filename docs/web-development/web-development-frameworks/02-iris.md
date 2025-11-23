---
title: Iris 框架
difficulty: intermediate
duration: "5-6小时"
prerequisites: ["Gin 基础", "数据库"]
tags: ["Iris", "框架", "Web", "高性能"]
---

# Iris 框架

Iris 是一个快速、简单但功能强大的 Go Web 框架，提供了丰富的功能和优秀的性能。

## 📋 学习目标

- [ ] 理解 Iris 框架的特点和优势
- [ ] 掌握 Iris 项目的创建和配置
- [ ] 学会使用路由和中间件
- [ ] 掌握视图渲染和模板
- [ ] 理解依赖注入和 MVC 模式
- [ ] 学会使用 Iris 的扩展功能

## 🎯 Iris 简介

### 为什么选择 Iris

- **高性能**: 基于 fasthttp，性能优秀
- **功能丰富**: 内置模板引擎、会话管理、WebSocket 等
- **易于使用**: API 设计简洁直观
- **文档完善**: 官方文档详细，示例丰富
- **活跃维护**: 持续更新，社区活跃

### Iris vs Gin vs Beego

| 特性 | Iris | Gin | Beego |
|------|------|-----|-------|
| 性能 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ |
| 功能完整性 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| 学习曲线 | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ |
| 适用场景 | 全栈应用 | API 服务 | MVC 应用 |

## 🚀 快速开始

### 安装 Iris

```bash
# 安装 Iris
go get github.com/kataras/iris/v12@latest
```

### 第一个 Iris 应用

```go
package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "Hello, Iris!",
		})
	})

	app.Listen(":8080")
}
```

## 🛣️ 路由

### 基本路由

```go
package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// GET 请求
	app.Get("/", handler)
	
	// POST 请求
	app.Post("/users", createUser)
	
	// PUT 请求
	app.Put("/users/{id:int}", updateUser)
	
	// DELETE 请求
	app.Delete("/users/{id:int}", deleteUser)
	
	// 任意方法
	app.Any("/any", anyHandler)
	
	// 多个方法
	app.HandleMany("GET POST", "/many", manyHandler)

	app.Listen(":8080")
}

func handler(ctx iris.Context) {
	ctx.WriteString("Hello, Iris!")
}
```

### 路由参数

```go
app := iris.New()

// 路径参数
app.Get("/users/{id:int}", func(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	ctx.JSON(iris.Map{"id": id})
})

// 字符串参数
app.Get("/users/{name}", func(ctx iris.Context) {
	name := ctx.Params().Get("name")
	ctx.JSON(iris.Map{"name": name})
})

// 多个参数
app.Get("/users/{id:int}/posts/{postId:int}", func(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	postId := ctx.Params().GetIntDefault("postId", 0)
	ctx.JSON(iris.Map{
		"userId": id,
		"postId": postId,
	})
})

// 可选参数
app.Get("/users/{id:int?}", func(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	ctx.JSON(iris.Map{"id": id})
})
```

### 路由组

```go
app := iris.New()

// 创建路由组
users := app.Party("/users")
{
	users.Get("/", listUsers)
	users.Get("/{id:int}", getUser)
	users.Post("/", createUser)
	users.Put("/{id:int}", updateUser)
	users.Delete("/{id:int}", deleteUser)
}

// 带中间件的路由组
api := app.Party("/api", authMiddleware)
{
	api.Get("/users", getUsers)
	api.Post("/users", createUser)
}
```

## 🔧 中间件

### 全局中间件

```go
app := iris.New()

// 全局中间件
app.Use(func(ctx iris.Context) {
	ctx.Header("X-Custom-Header", "value")
	ctx.Next()
})

app.Get("/", handler)
```

### 路由组中间件

```go
app := iris.New()

// 认证中间件
authMiddleware := func(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "未授权"})
		return
	}
	// 验证 token
	ctx.Next()
}

// 日志中间件
loggerMiddleware := func(ctx iris.Context) {
	start := time.Now()
	ctx.Next()
	duration := time.Since(start)
	log.Printf("%s %s - %v", ctx.Method(), ctx.Path(), duration)
}

// 应用到路由组
api := app.Party("/api", authMiddleware, loggerMiddleware)
{
	api.Get("/users", getUsers)
}
```

### 内置中间件

```go
app := iris.New()

// 恢复中间件（处理 panic）
app.Use(recover.New())

// 日志中间件
app.Use(logger.New())

// CORS 中间件
app.Use(cors.New(cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
}))

// 压缩中间件
app.Use(compress.New())
```

## 📝 请求处理

### 获取请求参数

```go
// 查询参数
app.Get("/search", func(ctx iris.Context) {
	query := ctx.URLParam("q")
	page := ctx.URLParamIntDefault("page", 1)
	ctx.JSON(iris.Map{
		"query": query,
		"page":  page,
	})
})

// 表单参数
app.Post("/login", func(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	// 处理登录
})

// JSON 请求体
app.Post("/users", func(ctx iris.Context) {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	
	// 处理用户创建
	ctx.JSON(iris.Map{"message": "创建成功", "user": user})
})

// 文件上传
app.Post("/upload", func(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	defer file.Close()
	
	// 保存文件
	dst, _ := os.Create("./uploads/" + info.Filename)
	defer dst.Close()
	io.Copy(dst, file)
	
	ctx.JSON(iris.Map{"message": "上传成功"})
})
```

### 响应处理

```go
// JSON 响应
app.Get("/json", func(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "Hello",
		"status":  "success",
	})
})

// 文本响应
app.Get("/text", func(ctx iris.Context) {
	ctx.Text("Hello, Iris!")
})

// HTML 响应
app.Get("/html", func(ctx iris.Context) {
	ctx.HTML("<h1>Hello, Iris!</h1>")
})

// XML 响应
app.Get("/xml", func(ctx iris.Context) {
	ctx.XML(iris.Map{"message": "Hello"})
})

// 文件下载
app.Get("/download", func(ctx iris.Context) {
	ctx.SendFile("./files/document.pdf", "document.pdf")
})
```

## 🎨 视图和模板

### 模板配置

```go
app := iris.New()

// 配置模板
app.RegisterView(iris.HTML("./views", ".html"))

app.Get("/", func(ctx iris.Context) {
	ctx.ViewData("Title", "首页")
	ctx.ViewData("Users", []string{"张三", "李四"})
	ctx.View("index.html")
})
```

### 模板语法

```html
<!-- views/index.html -->
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <ul>
        {{range .Users}}
        <li>{{.}}</li>
        {{end}}
    </ul>
</body>
</html>
```

### 布局模板

```html
<!-- views/layouts/main.html -->
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <header>Header</header>
    {{render}}
    <footer>Footer</footer>
</body>
</html>
```

```go
app.RegisterView(iris.HTML("./views", ".html").Layout("layouts/main.html"))
```

## 📊 数据库集成

### 使用 GORM

```go
package main

import (
	"github.com/kataras/iris/v12"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Email string
}

var db *gorm.DB

func init() {
	var err error
	dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	
	db.AutoMigrate(&User{})
}

func main() {
	app := iris.New()

	app.Get("/users", func(ctx iris.Context) {
		var users []User
		db.Find(&users)
		ctx.JSON(users)
	})

	app.Post("/users", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		
		db.Create(&user)
		ctx.JSON(user)
	})

	app.Listen(":8080")
}
```

## 🔐 会话管理

### Session 配置

```go
app := iris.New()

// 配置 Session
sess := sessions.New(sessions.Config{
	Cookie:  "sessionid",
	Expires: 24 * time.Hour,
})

app.Use(sess.Handler())

app.Post("/login", func(ctx iris.Context) {
	session := sessions.Get(ctx)
	session.Set("user_id", 1)
	session.Set("username", "admin")
	
	ctx.JSON(iris.Map{"message": "登录成功"})
})

app.Get("/profile", func(ctx iris.Context) {
	session := sessions.Get(ctx)
	userID := session.GetIntDefault("user_id", 0)
	
	if userID == 0 {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}
	
	ctx.JSON(iris.Map{"user_id": userID})
})
```

## 🏗️ MVC 模式

### 控制器

```go
package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	// 可以注入依赖
}

// Get 处理 GET /users
func (c *UserController) Get() iris.Map {
	return iris.Map{
		"message": "获取用户列表",
	}
}

// GetBy 处理 GET /users/{id:int}
func (c *UserController) GetBy(id int) iris.Map {
	return iris.Map{
		"id":      id,
		"message": "获取用户",
	}
}

// Post 处理 POST /users
func (c *UserController) Post() iris.Map {
	return iris.Map{
		"message": "创建用户",
	}
}
```

### 注册控制器

```go
app := iris.New()

mvc.New(app.Party("/users")).Handle(new(controllers.UserController))
```

## 🚀 实践应用

### 完整的 CRUD API

```go
package main

import (
	"github.com/kataras/iris/v12"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func init() {
	var err error
	dsn := "root:password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.AutoMigrate(&User{})
}

func main() {
	app := iris.New()

	// 获取用户列表
	app.Get("/users", func(ctx iris.Context) {
		var users []User
		db.Find(&users)
		ctx.JSON(users)
	})

	// 获取单个用户
	app.Get("/users/{id:int}", func(ctx iris.Context) {
		id := ctx.Params().GetIntDefault("id", 0)
		var user User
		if err := db.First(&user, id).Error; err != nil {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": "用户不存在"})
			return
		}
		ctx.JSON(user)
	})

	// 创建用户
	app.Post("/users", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		db.Create(&user)
		ctx.StatusCode(iris.StatusCreated)
		ctx.JSON(user)
	})

	// 更新用户
	app.Put("/users/{id:int}", func(ctx iris.Context) {
		id := ctx.Params().GetIntDefault("id", 0)
		var user User
		if err := db.First(&user, id).Error; err != nil {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{"error": "用户不存在"})
			return
		}
		if err := ctx.ReadJSON(&user); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		db.Save(&user)
		ctx.JSON(user)
	})

	// 删除用户
	app.Delete("/users/{id:int}", func(ctx iris.Context) {
		id := ctx.Params().GetIntDefault("id", 0)
		db.Delete(&User{}, id)
		ctx.JSON(iris.Map{"message": "删除成功"})
	})

	app.Listen(":8080")
}
```

## ⚠️ 注意事项

### 1. 错误处理

```go
app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
	ctx.JSON(iris.Map{"error": "页面未找到"})
})

app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
	ctx.JSON(iris.Map{"error": "服务器内部错误"})
})
```

### 2. 性能优化

```go
app := iris.New()

// 启用压缩
app.Use(compress.New())

// 启用缓存
app.Use(cache.Handler(2 * time.Minute))

// 静态文件服务
app.HandleDir("/static", "./static")
```

### 3. 配置管理

```go
app := iris.New()

app.Configure(iris.WithOptimizations) // 生产环境优化

app.Listen(":8080", iris.WithConfiguration(iris.Configuration{
	DisableStartupLog: false,
	EnableOptimizations: true,
}))
```

## 📚 扩展阅读

- [Iris 官方文档](https://www.iris-go.com/)
- [Iris GitHub](https://github.com/kataras/iris)
- [Iris 示例](https://github.com/kataras/iris/tree/master/_examples)

## ⏭️ 下一阶段

完成 Iris 学习后，可以：

- [Beego 框架](./01-beego.md) - 学习另一个 Go Web 框架
- [实战项目](../projects/) - 使用 Iris 构建完整项目
- [微服务](../microservices/) - 学习微服务架构

---

**💡 提示**: Iris 是一个功能强大且性能优秀的 Web 框架，适合构建各种类型的 Web 应用。掌握 Iris 可以快速开发高性能的 Web 服务！

