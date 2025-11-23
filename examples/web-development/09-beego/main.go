// Package main 展示 Beego 框架的基本用法
package main

import (
	"github.com/beego/beego/v2/server/web"
)

// MainController 主控制器
type MainController struct {
	web.Controller
}

// Get 处理 GET 请求
func (c *MainController) Get() {
	c.Ctx.WriteString("Hello, Beego!")
}

// UserController 用户控制器
type UserController struct {
	web.Controller
}

// Get 获取用户列表
func (c *UserController) Get() {
	users := []map[string]interface{}{
		{"id": 1, "name": "Alice"},
		{"id": 2, "name": "Bob"},
	}
	c.Data["json"] = users
	c.ServeJSON()
}

// Post 创建用户
func (c *UserController) Post() {
	name := c.GetString("name")
	email := c.GetString("email")
	
	c.Data["json"] = map[string]interface{}{
		"message": "用户创建成功",
		"name":    name,
		"email":   email,
	}
	c.ServeJSON()
}

func main() {
	// 路由配置
	web.Router("/", &MainController{})
	web.Router("/api/users", &UserController{})

	// 启动服务器
	web.Run()
}

