---
title: 消息推送
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["群组聊天"]
tags: ["推送", "通知", "离线消息"]
---

# 消息推送

本章节将实现离线消息推送、通知管理和消息提醒功能。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现离线消息推送
- [ ] 实现消息通知管理
- [ ] 实现推送队列处理
- [ ] 集成第三方推送服务
- [ ] 实现消息提醒功能

## 📲 推送服务

创建 `internal/service/notification.go`:

```go
package service

import (
	"chat-app/internal/model"
	"chat-app/internal/repository"
)

type NotificationService interface {
	SendNotification(userID uint, notification *Notification) error
	GetNotifications(userID uint, page, pageSize int) ([]Notification, int64, error)
	MarkAsRead(notificationID uint, userID uint) error
	PushOfflineMessages(userID uint) error
}

func (s *NotificationServiceImpl) PushOfflineMessages(userID uint) error {
	// 获取离线消息
	messages, err := s.messageRepo.GetUnreadMessages(userID)
	if err != nil {
		return err
	}

	// 发送推送通知
	for _, msg := range messages {
		notification := &Notification{
			Type:    "message",
			UserID:  userID,
			Title:   "新消息",
			Content: msg.Content,
			Data:    msg,
		}
		s.SendNotification(userID, notification)
	}

	return nil
}
```

## ⏭️ 下一步

消息推送完成后，下一步是：

- [部署优化](./08-deployment.md) - 部署和性能优化

---

**🎉 消息推送完成！** 现在你可以开始学习部署和优化了。

