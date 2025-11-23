---
title: WebSocket基础
difficulty: advanced
duration: "4-5小时"
prerequisites: ["数据模型设计"]
tags: ["WebSocket", "实时通信", "连接管理"]
---

# WebSocket基础

本章节将实现WebSocket连接管理、消息处理和连接池管理。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现WebSocket连接升级
- [ ] 管理WebSocket连接池
- [ ] 实现消息的发送和接收
- [ ] 处理连接断开和重连
- [ ] 实现心跳检测机制
- [ ] 处理并发连接

## 🔌 WebSocket 连接管理

创建 `pkg/websocket/client.go`:

```go
package websocket

import (
	"encoding/json"
	"log"
	"sync"
	"time"
	"github.com/gorilla/websocket"
)

// Client WebSocket客户端
type Client struct {
	ID       string
	UserID   uint
	Conn     *websocket.Conn
	Send     chan []byte
	Hub      *Hub
	mu       sync.Mutex
}

// Hub 管理所有客户端连接
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// ReadPump 从WebSocket连接读取消息
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket错误: %v", err)
			}
			break
		}

		// 处理接收到的消息
		c.handleMessage(message)
	}
}

// WritePump 向WebSocket连接写入消息
func (c *Client) WritePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 批量发送队列中的消息
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleMessage(message []byte) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Printf("解析消息失败: %v", err)
		return
	}

	// 处理不同类型的消息
	switch msg.Type {
	case "ping":
		c.sendPong()
	case "message":
		c.Hub.broadcast <- message
	default:
		log.Printf("未知消息类型: %s", msg.Type)
	}
}
```

## 📝 消息处理

创建 `pkg/websocket/message.go`:

```go
package websocket

// Message WebSocket消息结构
type Message struct {
	Type      string      `json:"type"`      // ping, pong, message, join, leave
	From      uint        `json:"from,omitempty"`
	To        *uint       `json:"to,omitempty"`
	RoomID    *uint       `json:"room_id,omitempty"`
	Content   string      `json:"content,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// NewMessage 创建新消息
func NewMessage(msgType string, from uint, content string) *Message {
	return &Message{
		Type:      msgType,
		From:      from,
		Content:   content,
		Timestamp: time.Now().Unix(),
	}
}
```

## 🔧 WebSocket 处理器

创建 `internal/handler/websocket.go`:

```go
package handler

import (
	"chat-app/pkg/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
)

var hub = websocket.NewHub()

func init() {
	go hub.Run()
}

func HandleWebSocket(c *gin.Context) {
	// 获取用户ID（从JWT token中）
	userID := c.GetUint("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "未授权",
		})
		return
	}

	// 升级HTTP连接为WebSocket
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "WebSocket升级失败",
		})
		return
	}

	// 创建客户端
	client := &websocket.Client{
		ID:     generateClientID(),
		UserID: userID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Hub:    hub,
	}

	// 注册客户端
	hub.Register <- client

	// 启动读写协程
	go client.WritePump()
	go client.ReadPump()
}

func generateClientID() string {
	return fmt.Sprintf("client_%d", time.Now().UnixNano())
}
```

## ⏭️ 下一步

WebSocket基础完成后，下一步是：

- [消息系统](./04-messages.md) - 实现消息的发送、接收和存储

---

**🎉 WebSocket基础完成！** 现在你可以开始实现消息系统了。

