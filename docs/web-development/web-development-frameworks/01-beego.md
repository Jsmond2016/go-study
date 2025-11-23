---
title: Beego 框架
difficulty: intermediate
duration: "5-6小时"
prerequisites: ["Gin 基础", "数据库"]
tags: ["Beego", "框架", "MVC", "ORM"]
---

# Beego 框架

Beego 是一个基于 Go 语言的高性能 Web 框架，采用 MVC 架构模式，提供了完整的 Web 开发解决方案。

## 📋 学习目标

- [ ] 理解 Beego 框架的特点和架构
- [ ] 掌握 Beego 项目的创建和配置
- [ ] 学会使用路由和控制器
- [ ] 掌握模型和数据库操作
- [ ] 理解视图和模板使用
- [ ] 学会使用中间件和过滤器
- [ ] 掌握会话管理和认证

## 🎯 Beego 简介

### 为什么选择 Beego

- **MVC 架构**: 清晰的代码组织结构
- **功能完整**: 内置 ORM、Session、日志等
- **开发效率**: Bee 工具提供快速开发
- **文档完善**: 官方文档详细
- **社区活跃**: 国内使用广泛

### Beego vs Gin

| 特性 | Beego | Gin |
|------|-------|-----|
| 架构 | MVC | 轻量级 |
| ORM | 内置 | 需集成 GORM |
| 模板 | 内置 | 需配置 |
| 学习曲线 | 中等 | 简单 |
| 适用场景 | 全栈应用 | API 服务 |

## 🚀 快速开始

### 安装 Beego

```bash
# 安装 Beego
go get github.com/beego/beego/v2/server/web

# 安装 Bee 工具
go get github.com/beego/bee/v2
```

### 创建项目

```bash
# 使用 Bee 工具创建项目
bee new myproject

# 进入项目目录
cd myproject

# 运行项目
bee run
```

### 第一个 Beego 应用

```go
package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Router("/", &MainController{})
	web.Run()
}

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("Hello, Beego!")
}
```

## 🏗️ 项目结构

### 标准项目结构

```
myproject/
├── conf/              # 配置文件
│   └── app.conf       # 应用配置
├── controllers/       # 控制器
│   └── default.go
├── models/            # 模型
│   └── models.go
├── routers/           # 路由
│   └── router.go
├── static/            # 静态文件
│   ├── css/
│   ├── js/
│   └── img/
├── tests/             # 测试文件
├── views/             # 视图模板
│   └── index.tpl
├── main.go            # 入口文件
└── go.mod
```

### 配置文件

```ini
# conf/app.conf
appname = myproject
httpport = 8080
runmode = dev

# 数据库配置
dbdriver = mysql
dbuser = root
dbpass = password
dbhost = 127.0.0.1
dbport = 3306
dbname = mydb
```

## 🛣️ 路由和控制器

### 基本路由

```go
package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// 基本路由
	web.Router("/", &MainController{})
	web.Router("/user/:id", &UserController{})

	// RESTful 路由
	web.Router("/api/user", &UserController{}, "get:GetUser;post:CreateUser")
	web.Router("/api/user/:id", &UserController{}, "get:GetUser;put:UpdateUser;delete:DeleteUser")

	web.Run()
}
```

### 控制器

```go
package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

// GET /user/:id
func (c *UserController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["json"] = map[string]interface{}{
		"id":   id,
		"name": "张三",
	}
	c.ServeJSON()
}

// POST /user
func (c *UserController) Post() {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := web.BindJSON(&c.Controller, &user); err != nil {
		c.Ctx.WriteString("参数错误")
		return
	}

	c.Data["json"] = map[string]interface{}{
		"message": "创建成功",
		"user":    user,
	}
	c.ServeJSON()
}
```

### 路由参数

```go
// 路径参数
web.Router("/user/:id", &UserController{})
// 访问: /user/123
// 获取: c.Ctx.Input.Param(":id")

// 查询参数
// 访问: /search?q=keyword
// 获取: c.GetString("q")

// 表单参数
// 获取: c.GetString("name")
```

## 📊 模型和数据库

### ORM 模型定义

```go
package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int       `orm:"auto;pk"`
	Name     string    `orm:"size(100)"`
	Email    string    `orm:"size(100);unique"`
	Age      int       `orm:"default(0)"`
	Created  time.Time `orm:"auto_now_add"`
	Updated  time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(User))
}
```

### 数据库连接

```go
package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql",
		"root:password@tcp(127.0.0.1:3306)/mydb?charset=utf8")
}

func main() {
	web.Run()
}
```

### ORM 操作

```go
package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"myproject/models"
)

type UserController struct {
	web.Controller
}

// 查询
func (c *UserController) Get() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)

	if err == nil {
		c.Data["json"] = user
		c.ServeJSON()
	}
}

// 创建
func (c *UserController) Post() {
	o := orm.NewOrm()
	user := models.User{
		Name:  "张三",
		Email: "zhangsan@example.com",
		Age:   25,
	}

	id, err := o.Insert(&user)
	if err == nil {
		c.Data["json"] = map[string]interface{}{
			"id": id,
			"user": user,
		}
		c.ServeJSON()
	}
}

// 更新
func (c *UserController) Put() {
	o := orm.NewOrm()
	user := models.User{Id: 1}

	if o.Read(&user) == nil {
		user.Name = "李四"
		o.Update(&user)
		c.ServeJSON()
	}
}

// 删除
func (c *UserController) Delete() {
	o := orm.NewOrm()
	user := models.User{Id: 1}

	if o.Read(&user) == nil {
		o.Delete(&user)
		c.ServeJSON()
	}
}

// 查询列表
func (c *UserController) List() {
	o := orm.NewOrm()
	var users []models.User

	qs := o.QueryTable("user")
	qs.Filter("age__gte", 18).All(&users)

	c.Data["json"] = users
	c.ServeJSON()
}
```

## 🎨 视图和模板

### 模板语法

```html
<!-- views/user/index.tpl -->
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>用户列表</h1>
    {{range .Users}}
    <div>
        <p>姓名: {{.Name}}</p>
        <p>邮箱: {{.Email}}</p>
    </div>
    {{end}}
</body>
</html>
```

### 渲染模板

```go
func (c *UserController) Get() {
	c.Data["Title"] = "用户列表"
	c.Data["Users"] = []models.User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
	}
	c.TplName = "user/index.tpl"
}
```

### 模板函数

```go
package main

import (
	"strings"
	"github.com/beego/beego/v2/server/web"
	"html/template"
)

func init() {
	web.AddFuncMap("upper", func(s string) string {
		return strings.ToUpper(s)
	})
}

// 在模板中使用
// {{.Name | upper}}
```

## 🔧 中间件和过滤器

### 过滤器

```go
package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 前置过滤器
func beforeRouter(ctx *context.Context) {
	// 在路由之前执行
	ctx.Output.Header("X-Custom-Header", "value")
}

// 后置过滤器
func afterExec(ctx *context.Context) {
	// 在控制器执行后执行
}

// 完成过滤器
func finishRouter(ctx *context.Context) {
	// 请求完成后执行
}

func init() {
	web.InsertFilter("/*", web.BeforeRouter, beforeRouter)
	web.InsertFilter("/*", web.AfterExec, afterExec)
	web.InsertFilter("/*", web.FinishRouter, finishRouter)
}
```

### 认证过滤器

```go
func authFilter(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	if token == "" {
		ctx.Output.Status = 401
		ctx.Output.Body([]byte("未授权"))
		return
	}
	// 验证 token
}

func init() {
	web.InsertFilter("/api/*", web.BeforeRouter, authFilter)
}
```

## 🔐 会话管理

### Session 配置

```go
// conf/app.conf
sessionon = true
sessionprovider = memory
sessionname = beegosessionID
sessiongcmaxlifetime = 3600
```

### 使用 Session

```go
func (c *UserController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")

	// 验证用户
	if username == "admin" && password == "admin123" {
		c.SetSession("user_id", 1)
		c.SetSession("username", username)
		c.Data["json"] = map[string]interface{}{
			"message": "登录成功",
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"error": "用户名或密码错误",
		}
	}
	c.ServeJSON()
}

func (c *UserController) Profile() {
	userID := c.GetSession("user_id")
	if userID == nil {
		c.Ctx.WriteString("请先登录")
		return
	}

	c.Data["json"] = map[string]interface{}{
		"user_id": userID,
	}
	c.ServeJSON()
}

func (c *UserController) Logout() {
	c.DelSession("user_id")
	c.DelSession("username")
	c.Data["json"] = map[string]interface{}{
		"message": "退出成功",
	}
	c.ServeJSON()
}
```

## 🏃‍♂️ 实践应用

### 完整的 CRUD 示例

```go
package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"myproject/models"
)

type UserController struct {
	web.Controller
}

// 获取用户列表
func (c *UserController) List() {
	o := orm.NewOrm()
	var users []models.User

	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("page_size", 10)

	qs := o.QueryTable("user")
	qs.Limit(pageSize, (page-1)*pageSize).All(&users)

	c.Data["json"] = map[string]interface{}{
		"users": users,
		"page":  page,
	}
	c.ServeJSON()
}

// 获取单个用户
func (c *UserController) Get() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	user := models.User{Id: id}

	if err := o.Read(&user); err == nil {
		c.Data["json"] = user
	} else {
		c.Ctx.Output.Status = 404
		c.Data["json"] = map[string]string{"error": "用户不存在"}
	}
	c.ServeJSON()
}

// 创建用户
func (c *UserController) Post() {
	var user models.User
	if err := web.BindJSON(&c.Controller, &user); err != nil {
		c.Ctx.Output.Status = 400
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err == nil {
		user.Id = int(id)
		c.Data["json"] = user
	} else {
		c.Ctx.Output.Status = 500
		c.Data["json"] = map[string]string{"error": "创建失败"}
	}
	c.ServeJSON()
}

// 更新用户
func (c *UserController) Put() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	user := models.User{Id: id}

	if o.Read(&user) != nil {
		c.Ctx.Output.Status = 404
		c.Data["json"] = map[string]string{"error": "用户不存在"}
		c.ServeJSON()
		return
	}

	if err := web.BindJSON(&c.Controller, &user); err != nil {
		c.Ctx.Output.Status = 400
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	o.Update(&user)
	c.Data["json"] = user
	c.ServeJSON()
}

// 删除用户
func (c *UserController) Delete() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	user := models.User{Id: id}

	if o.Read(&user) == nil {
		o.Delete(&user)
		c.Data["json"] = map[string]string{"message": "删除成功"}
	} else {
		c.Ctx.Output.Status = 404
		c.Data["json"] = map[string]string{"error": "用户不存在"}
	}
	c.ServeJSON()
}
```

## ⚠️ 注意事项

### 1. 配置管理

```go
// ✅ 使用配置文件管理配置
web.AppConfig.String("dbuser")
web.AppConfig.Int("httpport")
```

### 2. 错误处理

```go
// ✅ 统一错误处理
func (c *UserController) errorResponse(message string, code int) {
	c.Ctx.Output.Status = code
	c.Data["json"] = map[string]string{"error": message}
	c.ServeJSON()
}
```

### 3. 性能优化

```go
// ✅ 使用连接池
orm.SetMaxIdleConns("default", 30)
orm.SetMaxOpenConns("default", 30)

// ✅ 启用模板缓存
web.BConfig.WebConfig.TemplateLeft = "{{"
web.BConfig.WebConfig.TemplateRight = "}}"
```

## 📚 扩展阅读

- [Beego 官方文档](https://beego.me/)
- [Beego GitHub](https://github.com/beego/beego)
- [Bee 工具文档](https://beego.me/docs/install/bee.md)

## ⏭️ 下一阶段

完成 Beego 学习后，可以：

- [Iris 框架](./02-iris.md) - 学习另一个 Go Web 框架
- [实战项目](../projects/) - 使用 Beego 构建完整项目
- [数据库](../database/) - 深入学习数据库操作

---

**💡 提示**: Beego 是一个功能完整的 MVC 框架，适合构建全栈 Web 应用。掌握 Beego 可以快速开发企业级应用！

