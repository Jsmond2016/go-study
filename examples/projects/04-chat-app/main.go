// Package main 实现一个完整的聊天应用
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Online   bool   `json:"online" gorm:"default:false"`
}

// Message 消息模型
type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FromID    uint      `json:"from_id"`
	ToID      uint      `json:"to_id"`   // 0 表示群组消息
	RoomID    uint      `json:"room_id"` // 0 表示私聊
	Content   string    `json:"content" gorm:"type:text"`
	Type      string    `json:"type" gorm:"default:text"` // text, image, file
	CreatedAt time.Time `json:"created_at"`
}

// Room 聊天室模型
type Room struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

// Client WebSocket 客户端
type Client struct {
	ID   uint
	Conn *websocket.Conn
	Send chan []byte
	Hub  *Hub
}

// Hub 管理所有客户端
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

var (
	db       *gorm.DB
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	hub = &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	db.AutoMigrate(&User{}, &Message{}, &Room{})
}

func (h *Hub) run() {
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

func (c *Client) readPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			continue
		}

		msg.FromID = c.ID
		msg.CreatedAt = time.Now()
		db.Create(&msg)

		// 广播消息
		c.Hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket 升级失败:", err)
		return
	}

	userID := uint(1) // 简化处理，实际应从认证获取
	client := &Client{
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  hub,
	}

	client.Hub.register <- client

	// 更新用户在线状态
	db.Model(&User{}).Where("id = ?", userID).Update("online", true)

	go client.writePump()
	go client.readPump()
}

func main() {
	go hub.run()

	r := gin.Default()

	// 用户管理
	r.POST("/api/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusCreated, gin.H{"data": user})
	})

	// WebSocket 连接
	r.GET("/ws", handleWebSocket)

	// 发送消息
	r.POST("/api/messages", func(c *gin.Context) {
		var msg Message
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg.FromID = 1 // 简化处理
		msg.CreatedAt = time.Now()
		db.Create(&msg)

		// 通过 WebSocket 发送
		msgJSON, _ := json.Marshal(msg)
		hub.broadcast <- msgJSON

		c.JSON(http.StatusCreated, gin.H{"data": msg})
	})

	// 获取消息历史
	r.GET("/api/messages", func(c *gin.Context) {
		var messages []Message
		toID := c.Query("to_id")
		roomID := c.Query("room_id")

		query := db
		if toID != "" {
			query = query.Where("(from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?)", 1, toID, toID, 1)
		}
		if roomID != "" {
			query = query.Where("room_id = ?", roomID)
		}

		query.Order("created_at DESC").Limit(50).Find(&messages)
		c.JSON(http.StatusOK, gin.H{"data": messages})
	})

	// 创建聊天室
	r.POST("/api/rooms", func(c *gin.Context) {
		var room Room
		if err := c.ShouldBindJSON(&room); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&room)
		c.JSON(http.StatusCreated, gin.H{"data": room})
	})

	// 获取聊天室列表
	r.GET("/api/rooms", func(c *gin.Context) {
		var rooms []Room
		db.Find(&rooms)
		c.JSON(http.StatusOK, gin.H{"data": rooms})
	})

	log.Println("聊天应用启动在 :8080")
	log.Println("WebSocket 连接: ws://localhost:8080/ws")
	r.Run(":8080")
}
