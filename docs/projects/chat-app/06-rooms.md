---
title: 群组聊天
difficulty: advanced
duration: "4-5小时"
prerequisites: ["用户状态"]
tags: ["聊天室", "群组", "成员管理"]
---

# 群组聊天

本章节将实现聊天室和群组聊天功能，包括创建聊天室、成员管理和群组消息。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现聊天室的创建和管理
- [ ] 实现成员加入和退出
- [ ] 实现群组消息发送
- [ ] 实现成员权限管理
- [ ] 实现聊天室列表查询

## 🏠 聊天室服务

创建 `internal/service/room.go`:

```go
package service

import (
	"chat-app/internal/model"
	"chat-app/internal/repository"
	"errors"
)

type RoomService interface {
	CreateRoom(room *model.Room, creatorID uint) error
	JoinRoom(roomID, userID uint) error
	LeaveRoom(roomID, userID uint) error
	GetRoom(roomID uint) (*model.Room, error)
	GetUserRooms(userID uint) ([]model.Room, error)
	SendRoomMessage(roomID uint, senderID uint, content string) error
}

func (s *RoomServiceImpl) CreateRoom(room *model.Room, creatorID uint) error {
	room.CreatedBy = creatorID
	room.Type = "public"

	if err := s.roomRepo.Create(room); err != nil {
		return err
	}

	// 创建者自动加入
	member := &model.RoomMember{
		RoomID:   room.ID,
		UserID:   creatorID,
		Role:     "owner",
		JoinedAt: time.Now(),
	}

	return s.roomMemberRepo.Create(member)
}

func (s *RoomServiceImpl) SendRoomMessage(roomID uint, senderID uint, content string) error {
	// 检查用户是否在房间中
	if !s.roomMemberRepo.IsMember(roomID, senderID) {
		return errors.New("不是房间成员")
	}

	message := &model.Message{
		SenderID: senderID,
		RoomID:   &roomID,
		Content:  content,
		Type:     "text",
		Status:   "sent",
	}

	return s.messageService.SendMessage(message)
}
```

## ⏭️ 下一步

群组聊天完成后，下一步是：

- [消息推送](./07-notifications.md) - 实现离线消息推送

---

**🎉 群组聊天完成！** 现在你可以开始实现消息推送了。

