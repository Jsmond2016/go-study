---
title: 实战项目
difficulty: intermediate
duration: "12-16周"
prerequisites: ["Web 开发", "数据库", "标准库"]
tags: ["项目", "实战", "API", "系统设计"]
---

# 实战项目

本模块提供完整的实战项目，帮助你将所学知识应用到实际开发中。

## 📋 学习目标

完成本模块学习后，你将能够：

- [ ] 设计和实现完整的 Web 应用
- [ ] 使用数据库存储数据
- [ ] 实现用户认证和授权
- [ ] 构建 RESTful API
- [ ] 部署和运维应用
- [ ] 掌握项目开发流程

## 🎯 项目列表

### 📚 项目 1：TODO API（第1-3周）

| 章节                              | 内容            | 预计时间 | 难度 |
| --------------------------------- | --------------- | -------- | ---- |
| [TODO API 项目](./01-todo-api.md) | 完整的 TODO API | 8-10小时 | ⭐⭐   |

**项目特点**：
- RESTful API 设计
- 数据库操作
- 用户认证
- 单元测试

### 🏗️ 项目 2：博客系统（第4-6周）

| 章节                           | 内容           | 预计时间  | 难度 |
| ------------------------------ | -------------- | --------- | ---- |
| [博客系统项目](./blog-system/) | 完整的博客系统 | 12-15小时 | ⭐⭐⭐  |

**项目特点**：
- 文章管理
- 评论系统
- 用户管理
- 文件上传

**学习路径**：
- [环境搭建](./blog-system/01-setup.md) - 项目初始化和配置
- [数据模型设计](./blog-system/02-models.md) - 数据库设计和模型定义
- [用户认证](./blog-system/03-auth.md) - 用户注册、登录、JWT认证
- [文章管理](./blog-system/04-articles.md) - 文章CRUD、分类、标签
- [评论系统](./blog-system/05-comments.md) - 评论、回复、审核
- [文件上传](./blog-system/06-upload.md) - 图片上传、文件管理
- [搜索功能](./blog-system/07-search.md) - 全文搜索、标签搜索
- [部署优化](./blog-system/08-deployment.md) - 部署、性能优化、监控

### 🛒 项目 3：电商系统（第7-10周）

| 章节                          | 内容         | 预计时间  | 难度 |
| ----------------------------- | ------------ | --------- | ---- |
| [电商系统项目](./e-commerce/) | 电商平台开发 | 20-25小时 | ⭐⭐⭐⭐ |

**项目特点**：
- 商品管理
- 购物车
- 订单系统
- 支付集成

**学习路径**：
- [环境搭建](./e-commerce/01-setup.md) - 项目初始化和配置
- [数据模型设计](./e-commerce/02-models.md) - 数据库设计和模型定义
- [商品管理](./e-commerce/03-products.md) - 商品CRUD、分类、库存
- [购物车](./e-commerce/04-cart.md) - 购物车管理
- [订单系统](./e-commerce/05-orders.md) - 订单创建、状态管理
- [支付集成](./e-commerce/06-payment.md) - 支付接口集成
- [库存管理](./e-commerce/07-inventory.md) - 库存管理、预警
- [部署优化](./e-commerce/08-deployment.md) - 部署、性能优化、监控

### 💬 项目 4：聊天应用（第11-14周）

| 章节                        | 内容         | 预计时间  | 难度 |
| --------------------------- | ------------ | --------- | ---- |
| [聊天应用项目](./chat-app/) | 实时聊天应用 | 20-25小时 | ⭐⭐⭐⭐ |

**项目特点**：
- WebSocket 通信
- 实时消息
- 用户在线状态
- 消息历史

**学习路径**：
- [环境搭建](./chat-app/01-setup.md) - 项目初始化和WebSocket配置
- [数据模型设计](./chat-app/02-models.md) - 数据库设计和模型定义
- [WebSocket基础](./chat-app/03-websocket.md) - WebSocket连接和消息处理
- [消息系统](./chat-app/04-messages.md) - 消息发送、接收、存储
- [用户状态](./chat-app/05-presence.md) - 在线状态、心跳检测
- [群组聊天](./chat-app/06-rooms.md) - 聊天室、群组管理
- [消息推送](./chat-app/07-notifications.md) - 离线推送、通知
- [部署优化](./chat-app/08-deployment.md) - 部署、性能优化、监控

## 🚀 快速开始

### 项目结构

```
project/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/
│   └── utils/
├── api/
│   └── routes.go
├── go.mod
└── README.md
```

## 💡 学习建议

### 📖 学习方法

1. **循序渐进**：从简单项目开始
2. **完整实现**：完成整个项目流程
3. **代码质量**：注重代码质量和测试
4. **文档编写**：编写项目文档

### 🔍 推荐资源

- [Go 项目布局](https://github.com/golang-standards/project-layout)
- [RESTful API 设计](https://restfulapi.net/)
- [系统设计指南](https://github.com/donnemartin/system-design-primer)

## 📚 前置知识

在开始实战项目前，建议先学习：

- [开发工具链](../toolchain/) - MySQL、GORM、配置管理、日志等工程化工具
- [Web 开发](../web-development/) - Web 框架和 API 开发

## ⏭️ 下一阶段

完成项目学习后，可以：

- 开发自己的项目
- 参与开源项目
- [微服务](../microservices/) - 分布式系统开发
- [运维部署](../devops/) - 部署和运维
- 深入学习特定领域

---

**🎉 开始你的项目实战之旅吧！** 选择第一个项目，开始构建完整的应用。

