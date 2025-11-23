---
title: 用户状态
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["消息系统"]
tags: ["在线状态", "心跳", "状态管理"]
---

# 用户状态

本章节将实现用户在线状态管理、心跳检测和状态同步功能。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现用户在线状态管理
- [ ] 实现心跳检测机制
- [ ] 实现状态同步和广播
- [ ] 处理用户上线和下线
- [ ] 实现最后在线时间记录

## 💚 状态管理服务

创建 `internal/service/presence.go`:

```go
package service

import (
	"chat-app/internal/model"
	"chat-app/internal/repository"
	"time"
)

type PresenceService interface {
	SetOnline(userID uint) error
	SetOffline(userID uint) error
	SetAway(userID uint) error
	GetStatus(userID uint) (string, error)
	GetOnlineUsers() ([]uint, error)
	UpdateLastSeen(userID uint) error
}

type PresenceServiceImpl struct {
	userRepo repository.UserRepository
	redis    *redis.Client
}

func NewPresenceService(userRepo repository.UserRepository, redis *redis.Client) PresenceService {
	return &PresenceServiceImpl{
		userRepo: userRepo,
		redis:    redis,
	}
}

func (s *PresenceServiceImpl) SetOnline(userID uint) error {
	// 更新数据库状态
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	user.Status = "online"
	now := time.Now()
	user.LastSeen = &now

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	// 更新Redis缓存
	key := fmt.Sprintf("user:status:%d", userID)
	s.redis.Set(ctx, key, "online", 5*time.Minute)

	// 广播状态变更
	s.broadcastStatusChange(userID, "online")

	return nil
}

func (s *PresenceServiceImpl) SetOffline(userID uint) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	user.Status = "offline"
	now := time.Now()
	user.LastSeen = &now

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	// 从Redis删除
	key := fmt.Sprintf("user:status:%d", userID)
	s.redis.Del(ctx, key)

	// 广播状态变更
	s.broadcastStatusChange(userID, "offline")

	return nil
}
```

## ⏭️ 下一步

用户状态完成后，下一步是：

- [群组聊天](./06-rooms.md) - 实现聊天室和群组功能

---

**🎉 用户状态完成！** 现在你可以开始实现群组聊天了。

