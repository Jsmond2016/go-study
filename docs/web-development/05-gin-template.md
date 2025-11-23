---
title: Gin 模板
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["Gin 基础"]
tags: ["Gin", "模板", "渲染", "HTML"]
---

# Gin 模板

Gin 支持多种模板引擎，可以用于渲染 HTML、XML 等格式的响应。

## 📋 学习目标

- [ ] 理解模板渲染的概念
- [ ] 掌握 HTML 模板的使用
- [ ] 学会模板语法和函数
- [ ] 理解模板继承和包含
- [ ] 掌握模板的最佳实践

## 🎯 基本用法

### 加载模板

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("templates/*")
	// 或
	r.LoadHTMLFiles("templates/index.html", "templates/about.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "首页",
		})
	})

	r.Run(":8080")
}
```

### 渲染模板

```go
r := gin.Default()
r.LoadHTMLGlob("templates/*")

r.GET("/", func(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "首页",
		"name":  "张三",
	})
})
```

## 📝 模板语法

### 变量输出

```html
<!-- templates/index.html -->
<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
</head>
<body>
    <h1>欢迎, {{.name}}</h1>
    <p>年龄: {{.age}}</p>
</body>
</html>
```

### 条件判断

```html
{{if .isAdmin}}
    <p>管理员面板</p>
{{else}}
    <p>普通用户</p>
{{end}}

{{if gt .age 18}}
    <p>已成年</p>
{{end}}
```

### 循环

```html
{{range .users}}
    <div>
        <p>姓名: {{.Name}}</p>
        <p>邮箱: {{.Email}}</p>
    </div>
{{end}}

{{range $index, $user := .users}}
    <p>{{$index}}: {{$user.Name}}</p>
{{end}}
```

### 模板函数

```html
{{.title | upper}}
{{.date | formatDate}}
{{len .items}}
```

## 🔧 自定义模板函数

```go
package main

import (
	"html/template"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"date": time.Now(),
		})
	})

	r.Run(":8080")
}
```

## 📦 模板继承

### 基础模板

```html
<!-- templates/layout.html -->
<!DOCTYPE html>
<html>
<head>
    <title>{{block "title" .}}默认标题{{end}}</title>
</head>
<body>
    {{block "content" .}}{{end}}
</body>
</html>
```

### 继承模板

```html
<!-- templates/index.html -->
{{template "layout.html" .}}

{{define "title"}}首页{{end}}

{{define "content"}}
<h1>欢迎来到首页</h1>
<p>这是首页内容</p>
{{end}}
```

## 🏃‍♂️ 实践应用

### 完整的模板示例

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "首页",
			"users": []User{
				{Name: "张三", Age: 25},
				{Name: "李四", Age: 30},
			},
		})
	})

	r.Run(":8080")
}
```

## 📁 静态文件服务

### 基本用法

```go
r := gin.Default()

// 静态文件服务
r.Static("/assets", "./assets")
// 访问: http://localhost:8080/assets/img.png

// 静态文件系统
r.StaticFS("/files", http.Dir("./public"))
// 显示目录下的所有文件

// 单个静态文件
r.StaticFile("/favicon.ico", "./favicon.ico")
```

### 完整示例

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态资源
	r.Static("/assets", "./assets")
	r.StaticFS("/show-dir", http.Dir("."))
	r.StaticFile("/image", "./images/img.png")

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.example.com/")
	})

	r.Run(":8080")
}
```

## ⚠️ 注意事项

### 1. 模板路径

```go
// ✅ 使用相对路径或绝对路径
r.LoadHTMLGlob("templates/*")
r.LoadHTMLFiles("./templates/index.html")

// ✅ 支持多层级模板
r.LoadHTMLGlob("templates/**/*")
```

### 2. 模板缓存

```go
// 生产环境启用模板缓存
gin.SetMode(gin.ReleaseMode)
```

### 3. XSS 防护

```go
// ✅ 使用 html/template 自动转义
// ❌ 不要使用 text/template 渲染用户输入
```

### 4. 静态文件安全

```go
// ✅ 限制静态文件目录
r.Static("/assets", "./public/assets")

// ❌ 避免暴露敏感目录
// r.StaticFS("/", http.Dir(".")) // 危险！
```

## 📚 扩展阅读

- [Go 模板文档](https://golang.org/pkg/html/template/)
- [Gin 模板示例](https://gin-gonic.com/docs/examples/html-rendering/)

## ⏭️ 下一章节

[数据验证](./06-gin-validation.md) → 学习 Gin 的数据验证和绑定

---

**💡 提示**: 模板渲染是 Web 开发的重要功能，掌握模板语法可以让页面渲染更加灵活！

