---
title: 环境搭建
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["Web 开发", "数据库", "工具链"]
tags: ["环境", "搭建", "配置", "WebSocket"]
---

# 环境搭建

本章节将指导你搭建聊天应用的开发环境，包括项目初始化、依赖安装和WebSocket配置。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 创建项目目录结构
- [ ] 初始化 Go 模块
- [ ] 安装项目依赖（包括WebSocket）
- [ ] 配置数据库连接
- [ ] 配置WebSocket服务器
- [ ] 运行项目并验证

## 🚀 快速开始

### 1. 创建项目目录

```bash
# 创建项目根目录
mkdir -p chat-app
cd chat-app

# 创建标准 Go 项目目录结构
mkdir -p cmd/server
mkdir -p internal/{handler,service,repository,model}
mkdir -p pkg/{websocket,utils}
mkdir -p config
```

### 2. 初始化 Go 模块

```bash
# 初始化 Go 模块
go mod init chat-app

# 创建主程序文件
touch cmd/server/main.go
```

### 3. 安装依赖

```bash
# 安装核心依赖
go get github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/golang-jwt/jwt/v5
go get github.com/spf13/viper
go get github.com/go-redis/redis/v8
```

## ⚙️ 配置文件

创建 `config/config.yaml`:

```yaml
server:
  port: 8080
  mode: debug
  host: localhost

websocket:
  read_buffer_size: 1024
  write_buffer_size: 1024
  handshake_timeout: 10s
  ping_period: 54s
  pong_wait: 60s
  max_message_size: 512

database:
  driver: mysql
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: chat_db
  charset: utf8mb4

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  secret: your-secret-key
  expire: 24h
```

## 🔌 WebSocket 配置

创建 `pkg/websocket/upgrader.go`:

```go
package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 生产环境应该验证Origin
		return true
	},
	HandshakeTimeout: 10 * time.Second,
}

// 配置WebSocket连接参数
func ConfigureUpgrader(readBuffer, writeBuffer int, checkOrigin func(*http.Request) bool) {
	Upgrader.ReadBufferSize = readBuffer
	Upgrader.WriteBufferSize = writeBuffer
	if checkOrigin != nil {
		Upgrader.CheckOrigin = checkOrigin
	}
}
```

## 🚀 主程序初始化

创建 `cmd/server/main.go`:

```go
package main

import (
	"fmt"
	"log"
	"chat-app/internal/repository"
	"chat-app/pkg/config"
	"chat-app/pkg/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("加载配置失败:", err)
	}

	// 初始化数据库
	dsn := repository.BuildDSN(cfg.Database)
	db, err := repository.InitDB(dsn)
	if err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	// 配置WebSocket
	websocket.ConfigureUpgrader(
		cfg.WebSocket.ReadBufferSize,
		cfg.WebSocket.WriteBufferSize,
		nil,
	)

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 创建 Gin 引擎
	r := gin.Default()

	// 设置路由
	setupRoutes(r, db)

	// 启动服务器
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("聊天服务器启动在 ws://%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "服务运行正常",
		})
	})

	// WebSocket端点
	r.GET("/ws", handleWebSocket)
}
```

## ✅ 验证安装

### 运行项目

```bash
# 运行项目
go run cmd/server/main.go

# 或者构建后运行
go build -o chat-app cmd/server/main.go
./chat-app
```

### 测试WebSocket连接

使用浏览器控制台测试：

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');
ws.onopen = () => console.log('连接成功');
ws.onmessage = (event) => console.log('收到消息:', event.data);
ws.onerror = (error) => console.error('错误:', error);
ws.onclose = () => console.log('连接关闭');
```

## 🔍 常见问题

### 1. WebSocket连接失败

**问题**: 无法建立WebSocket连接

**解决方案**:
- 检查服务器是否启动
- 验证WebSocket端点路径
- 检查防火墙设置

### 2. 跨域问题

**问题**: WebSocket连接被CORS阻止

**解决方案**:
- 配置CheckOrigin函数
- 设置正确的Origin验证

## 📝 下一步

环境搭建完成后，下一步是：

- [数据模型设计](./02-models.md) - 设计数据库表和模型

---

**🎉 环境搭建完成！** 现在你可以开始设计数据模型了。

