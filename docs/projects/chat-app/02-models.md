---
title: 数据模型设计
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["环境搭建"]
tags: ["数据模型", "数据库", "GORM", "消息", "聊天室"]
---

# 数据模型设计

本章节将详细介绍聊天应用的数据模型设计，包括用户、消息、聊天室等模型的定义和关联关系。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 设计聊天系统的数据库表结构
- [ ] 使用 GORM 定义数据模型
- [ ] 理解消息和聊天室的关系
- [ ] 实现数据库自动迁移
- [ ] 掌握实时系统的数据模型设计

## 🗄️ 数据库设计

### 实体关系图

```
User (用户)
  ├── Messages (消息) 1:N (发送者)
  ├── Messages (消息) 1:N (接收者)
  ├── RoomMembers (聊天室成员) 1:N
  └── UserSessions (用户会话) 1:N

Message (消息)
  ├── User (发送者) N:1
  ├── User (接收者) N:1 (私聊)
  └── Room (聊天室) N:1 (群聊)

Room (聊天室)
  ├── Messages (消息) 1:N
  └── RoomMembers (成员) 1:N

RoomMember (聊天室成员)
  ├── Room (聊天室) N:1
  └── User (用户) N:1
```

## 📦 模型定义

### 1. 用户模型

创建 `internal/model/user.go`:

```go
package model

import (
	"time"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password  string    `gorm:"not null;size:255" json:"-"`
	Nickname  string    `gorm:"size:50" json:"nickname"`
	Avatar    string    `gorm:"size:255" json:"avatar"`
	Status    string    `gorm:"default:offline;size:20" json:"status"` // online, offline, away
	LastSeen  *time.Time `json:"last_seen,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	SentMessages    []Message     `gorm:"foreignKey:SenderID" json:"-"`
	ReceivedMessages []Message    `gorm:"foreignKey:ReceiverID" json:"-"`
	RoomMembers     []RoomMember  `gorm:"foreignKey:UserID" json:"-"`
	Sessions        []UserSession `gorm:"foreignKey:UserID" json:"-"`
}

func (User) TableName() string {
	return "users"
}
```

### 2. 消息模型

创建 `internal/model/message.go`:

```go
package model

import (
	"time"
)

// Message 消息模型
type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Type      string    `gorm:"default:text;size:20" json:"type"` // text, image, file
	Status    string    `gorm:"default:sent;size:20" json:"status"` // sent, delivered, read
	CreatedAt time.Time `json:"created_at"`

	// 外键
	SenderID   uint  `gorm:"not null;index" json:"sender_id"`
	ReceiverID *uint `gorm:"index" json:"receiver_id,omitempty"` // 私聊接收者
	RoomID     *uint `gorm:"index" json:"room_id,omitempty"`     // 群聊房间

	// 关联
	Sender   User  `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	Receiver *User `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
	Room     *Room `gorm:"foreignKey:RoomID" json:"room,omitempty"`
}

func (Message) TableName() string {
	return "messages"
}
```

### 3. 聊天室模型

创建 `internal/model/room.go`:

```go
package model

import (
	"time"
)

// Room 聊天室模型
type Room struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;size:100" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Type        string    `gorm:"default:public;size:20" json:"type"` // public, private
	Avatar      string    `gorm:"size:255" json:"avatar"`
	CreatedBy   uint      `gorm:"not null;index" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// 关联
	Messages     []Message     `gorm:"foreignKey:RoomID" json:"messages,omitempty"`
	Members      []RoomMember  `gorm:"foreignKey:RoomID" json:"members,omitempty"`
	Creator      User          `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// RoomMember 聊天室成员
type RoomMember struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RoomID    uint      `gorm:"not null;index" json:"room_id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Role      string    `gorm:"default:member;size:20" json:"role"` // owner, admin, member
	JoinedAt  time.Time `json:"joined_at"`

	Room Room `gorm:"foreignKey:RoomID" json:"room,omitempty"`
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Room) TableName() string {
	return "rooms"
}

func (RoomMember) TableName() string {
	return "room_members"
}
```

### 4. 用户会话模型

创建 `internal/model/session.go`:

```go
package model

import (
	"time"
)

// UserSession 用户会话模型
type UserSession struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null;size:255" json:"token"`
	IP        string    `gorm:"size:45" json:"ip"`
	UserAgent string    `gorm:"size:255" json:"user_agent"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
```

## 🗄️ 数据库迁移

创建 `internal/repository/migrate.go`:

```go
package repository

import (
	"chat-app/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Message{},
		&model.Room{},
		&model.RoomMember{},
		&model.UserSession{},
	)
}
```

## 💡 索引优化

### 创建索引

```go
// 消息表索引
db.Exec("CREATE INDEX idx_message_sender ON messages(sender_id)")
db.Exec("CREATE INDEX idx_message_receiver ON messages(receiver_id)")
db.Exec("CREATE INDEX idx_message_room ON messages(room_id)")
db.Exec("CREATE INDEX idx_message_created ON messages(created_at)")

// 聊天室成员索引
db.Exec("CREATE INDEX idx_room_member_room ON room_members(room_id)")
db.Exec("CREATE INDEX idx_room_member_user ON room_members(user_id)")
```

## ⏭️ 下一步

数据模型设计完成后，下一步是：

- [WebSocket基础](./03-websocket.md) - 实现WebSocket连接和消息处理

---

**🎉 数据模型设计完成！** 现在你可以开始实现WebSocket功能了。

