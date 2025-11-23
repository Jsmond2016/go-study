---
title: 消息系统
difficulty: advanced
duration: "4-5小时"
prerequisites: ["WebSocket基础"]
tags: ["消息", "发送", "接收", "存储"]
---

# 消息系统

本章节将实现完整的消息系统，包括消息发送、接收、存储和消息历史查询。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现消息的发送和接收
- [ ] 实现消息存储到数据库
- [ ] 实现私聊消息功能
- [ ] 实现群聊消息功能
- [ ] 实现消息状态管理
- [ ] 实现消息历史查询

## 💬 消息服务

创建 `internal/service/message.go`:

```go
package service

import (
	"chat-app/internal/model"
	"chat-app/internal/repository"
	"errors"
	"time"
)

type MessageService interface {
	SendMessage(msg *model.Message) error
	GetMessages(userID uint, targetID *uint, roomID *uint, page, pageSize int) ([]model.Message, int64, error)
	MarkAsRead(messageID uint, userID uint) error
	GetUnreadCount(userID uint) (int64, error)
}

type MessageServiceImpl struct {
	messageRepo repository.MessageRepository
	hub         *websocket.Hub
}

func NewMessageService(messageRepo repository.MessageRepository, hub *websocket.Hub) MessageService {
	return &MessageServiceImpl{
		messageRepo: messageRepo,
		hub:         hub,
	}
}

func (s *MessageServiceImpl) SendMessage(msg *model.Message) error {
	// 保存消息到数据库
	if err := s.messageRepo.Create(msg); err != nil {
		return err
	}

	// 发送消息到WebSocket
	s.broadcastMessage(msg)

	return nil
}

func (s *MessageServiceImpl) broadcastMessage(msg *model.Message) {
	messageData := map[string]interface{}{
		"type":      "message",
		"id":        msg.ID,
		"sender_id": msg.SenderID,
		"content":   msg.Content,
		"created_at": msg.CreatedAt,
	}

	// 私聊消息
	if msg.ReceiverID != nil {
		messageData["receiver_id"] = *msg.ReceiverID
		s.hub.SendToUser(*msg.ReceiverID, messageData)
	}

	// 群聊消息
	if msg.RoomID != nil {
		messageData["room_id"] = *msg.RoomID
		s.hub.BroadcastToRoom(*msg.RoomID, messageData)
	}
}
```

## 📝 消息处理器

创建 `internal/handler/message.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"chat-app/internal/model"
	"chat-app/internal/service"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数无效",
		})
		return
	}

	message := &model.Message{
		SenderID:   userID,
		ReceiverID: req.ReceiverID,
		RoomID:     req.RoomID,
		Content:    req.Content,
		Type:       req.Type,
		Status:     "sent",
	}

	if err := h.messageService.SendMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "发送消息失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "消息发送成功",
		"data":    message,
	})
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	var targetID *uint
	var roomID *uint

	if targetIDStr := c.Query("target_id"); targetIDStr != "" {
		if id, err := strconv.ParseUint(targetIDStr, 10, 32); err == nil {
			targetID = new(uint)
			*targetID = uint(id)
		}
	}

	if roomIDStr := c.Query("room_id"); roomIDStr != "" {
		if id, err := strconv.ParseUint(roomIDStr, 10, 32); err == nil {
			roomID = new(uint)
			*roomID = uint(id)
		}
	}

	messages, total, err := h.messageService.GetMessages(userID, targetID, roomID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取消息失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items":     messages,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}
```

## ⏭️ 下一步

消息系统完成后，下一步是：

- [用户状态](./05-presence.md) - 实现在线状态管理

---

**🎉 消息系统完成！** 现在你可以开始实现用户状态管理了。

