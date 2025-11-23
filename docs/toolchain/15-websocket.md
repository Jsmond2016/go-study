---
title: WebSocket
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["Web 开发", "并发编程"]
tags: ["WebSocket", "实时通信", "gorilla/websocket"]
---

# WebSocket

WebSocket 是一种在单个 TCP 连接上进行全双工通信的协议，适用于实时通信场景。

## 📋 学习目标

- [ ] 理解 WebSocket 的概念和优势
- [ ] 掌握 gorilla/websocket 的使用
- [ ] 学会实现 WebSocket 服务器
- [ ] 理解消息广播
- [ ] 掌握连接管理
- [ ] 了解最佳实践

## 🎯 WebSocket 简介

### 为什么使用 WebSocket

- **实时通信**: 双向实时数据传输
- **低延迟**: 比 HTTP 轮询更高效
- **全双工**: 客户端和服务器可以同时发送数据
- **减少开销**: 建立连接后开销小
- **适用场景**: 聊天、实时通知、游戏等

### 安装 gorilla/websocket

```bash
go get github.com/gorilla/websocket
```

## 🚀 快速开始

### 基本 WebSocket 服务器

```go
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}
	defer conn.Close()
	
	// 读取消息
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}
		
		log.Printf("收到消息: %s", message)
		
		// 回显消息
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("发送消息失败: %v", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket 服务器启动在 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 🔧 在 Gin 中使用

### Gin WebSocket 处理

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}
	defer conn.Close()
	
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}
		
		// 处理消息
		log.Printf("收到消息: %s", message)
		
		// 发送响应
		if err := conn.WriteMessage(messageType, []byte("收到: "+string(message))); err != nil {
			log.Printf("发送消息失败: %v", err)
			break
		}
	}
}

func main() {
	r := gin.Default()
	
	r.GET("/ws", handleWebSocket)
	
	r.Run(":8080")
}
```

## 📡 消息广播

### Hub 模式

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	User    string `json:"user"`
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Println("客户端连接")
			
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Println("客户端断开")
			}
			
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		
		// 广播消息
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()
	
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}
	
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	
	client.hub.register <- client
	
	go client.writePump()
	go client.readPump()
}

func main() {
	hub := newHub()
	go hub.run()
	
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(hub, w, r)
	})
	
	log.Println("WebSocket 服务器启动在 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 🏃‍♂️ 实践应用

### 聊天室实现

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan ChatMessage)

func handleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}
	defer conn.Close()
	
	clients[conn] = true
	
	for {
		var msg ChatMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			delete(clients, conn)
			break
		}
		
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("发送消息失败: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	r := gin.Default()
	
	// 静态文件
	r.Static("/static", "./static")
	
	// WebSocket 路由
	r.GET("/ws", handleConnections)
	
	// 启动消息处理
	go handleMessages()
	
	r.Run(":8080")
}
```

## ⚠️ 注意事项

### 1. 连接管理

```go
// ✅ 正确关闭连接
defer conn.Close()

// ✅ 处理连接错误
if err != nil {
	log.Printf("连接错误: %v", err)
	return
}
```

### 2. 消息格式

```go
// ✅ 使用 JSON 格式
type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// ✅ 统一消息格式
msg, _ := json.Marshal(Message{
	Type:    "chat",
	Content: "Hello",
})
```

### 3. 安全性

```go
// ✅ 验证来源
upgrader.CheckOrigin = func(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return origin == "https://example.com"
}

// ✅ 使用 WSS（WebSocket Secure）
// ✅ 实现认证机制
```

## 📚 扩展阅读

- [gorilla/websocket 文档](https://github.com/gorilla/websocket)
- [WebSocket 规范](https://tools.ietf.org/html/rfc6455)

## ⏭️ 下一阶段

完成开发工具链学习后，可以进入：

- [实战项目](../projects/) - 使用这些工具构建完整项目

---

**💡 提示**: WebSocket 是实现实时通信的最佳选择，掌握它对于构建现代 Web 应用非常重要！

