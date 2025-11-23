---
title: Web 开发
difficulty: intermediate
duration: "8-10周"
prerequisites: ["基础语法", "标准库", "数据库"]
tags: ["Web", "HTTP", "Gin", "API", "REST"]
---

# Web 开发

本模块介绍使用 Go 语言进行 Web 开发，包括 HTTP 服务器、Gin 框架、REST API 设计等。

## 📋 学习目标

完成本模块学习后，你将能够：

- [ ] 构建 HTTP 服务器
- [ ] 使用 Gin 框架开发 Web 应用
- [ ] 使用 Beego 框架开发 MVC 应用
- [ ] 设计 RESTful API
- [ ] 实现中间件和认证
- [ ] 处理请求和响应
- [ ] 部署 Web 应用

## 🎯 学习路径

### 📚 第一部分：HTTP 基础（第1-2周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [HTTP 服务器](./01-http-server.md)      | net/http 使用      | 3-4小时  | ⭐⭐   |
| [Gin 基础](./02-gin-basics.md)          | Gin 框架入门       | 4-5小时  | ⭐⭐   |

### 🛠️ 第二部分：路由和中间件（第3-4周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [Gin 路由](./03-gin-routing.md)         | 路由配置和参数     | 3-4小时  | ⭐⭐   |
| [Gin 中间件](./04-gin-middleware.md)    | 中间件开发         | 4-5小时  | ⭐⭐⭐ |

### 🎨 第三部分：高级功能（第5-6周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [Gin 模板](./05-gin-template.md)        | 模板渲染           | 3-4小时  | ⭐⭐   |
| [数据验证](./06-gin-validation.md)      | 请求验证和绑定     | 3-4小时  | ⭐⭐⭐ |
| [认证授权](./07-gin-auth.md)            | JWT、Session        | 4-5小时  | ⭐⭐⭐ |
| [REST API 设计](./08-rest-api.md)       | API 设计最佳实践   | 4-5小时  | ⭐⭐⭐ |

### 🐝 第四部分：Beego 框架（第7-8周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [Beego 框架](./09-beego.md)             | MVC 架构、ORM、模板 | 5-6小时  | ⭐⭐⭐ |

## 🚀 快速开始

### 第一个 Web 服务器

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
			"message": "Hello, Go!",
		})
	})

	r.Run(":8080")
}
```

## 💡 学习建议

### 📖 学习方法

1. **循序渐进**：从简单到复杂
2. **实践项目**：通过项目学习
3. **理解原理**：理解 HTTP 协议和框架原理
4. **最佳实践**：学习 Web 开发最佳实践

### 🔍 推荐资源

- [Gin 官方文档](https://gin-gonic.com/)
- [Beego 官方文档](https://beego.me/)
- [HTTP 协议](https://developer.mozilla.org/zh-CN/docs/Web/HTTP)
- [RESTful API 设计指南](https://restfulapi.net/)

## ⏭️ 下一阶段

完成 Web 开发学习后，可以进入：

- [开发工具链](../toolchain/) - 学习工程化工具和技术栈
- [实战项目](../projects/) - 构建完整项目
- [微服务](../microservices/) - 分布式系统
- [运维部署](../devops/) - 部署和运维

---

**🎉 开始你的 Web 开发之旅吧！** 选择第一个章节，开始学习 Web 开发。

