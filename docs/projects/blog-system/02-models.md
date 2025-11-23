---
title: 数据模型设计
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["环境搭建"]
tags: ["数据模型", "数据库", "GORM", "关联"]
---

# 数据模型设计

本章节将详细介绍博客系统的数据模型设计，包括用户、文章、分类、标签和评论等模型的定义和关联关系。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 设计博客系统的数据库表结构
- [ ] 使用 GORM 定义数据模型
- [ ] 理解模型之间的关联关系
- [ ] 实现数据库自动迁移
- [ ] 掌握 GORM 的高级特性

## 🗄️ 数据库设计

### 实体关系图

```
User (用户)
  ├── Articles (文章) 1:N
  └── Comments (评论) 1:N

Article (文章)
  ├── User (作者) N:1
  ├── Category (分类) N:1
  ├── Tags (标签) N:M
  └── Comments (评论) 1:N

Category (分类)
  └── Articles (文章) 1:N

Tag (标签)
  └── Articles (文章) N:M

Comment (评论)
  ├── Article (文章) N:1
  ├── User (用户) N:1
  └── Parent (父评论) N:1 (自关联)
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
	Role      string    `gorm:"default:user;size:20" json:"role"` // admin, user
	Status    string    `gorm:"default:active;size:20" json:"status"` // active, inactive
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
	
	// 关联
	Articles []Article `gorm:"foreignKey:UserID" json:"articles,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
```

**字段说明**：
- `Username`: 用户名，唯一索引
- `Email`: 邮箱，唯一索引
- `Password`: 密码（不序列化到JSON）
- `Role`: 角色（admin/user）
- `Status`: 状态（active/inactive）
- `DeletedAt`: 软删除字段

### 2. 分类模型

创建 `internal/model/category.go`:

```go
package model

import (
	"time"
)

// Category 分类模型
type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Slug        string    `gorm:"uniqueIndex;not null;size:100" json:"slug"`
	Description string    `gorm:"size:255" json:"description"`
	Sort        int       `gorm:"default:0" json:"sort"` // 排序字段
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	Articles    []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
```

### 3. 标签模型

创建 `internal/model/tag.go`:

```go
package model

import (
	"time"
)

// Tag 标签模型
type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Slug      string    `gorm:"uniqueIndex;not null;size:100" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	
	Articles  []Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
}

func (Tag) TableName() string {
	return "tags"
}
```

### 4. 文章模型

创建 `internal/model/article.go`:

```go
package model

import (
	"time"
)

// Article 文章模型
type Article struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Title        string     `gorm:"not null;size:200;index" json:"title"`
	Slug         string     `gorm:"uniqueIndex;not null;size:255" json:"slug"`
	Content      string     `gorm:"type:text" json:"content"`
	Summary      string     `gorm:"size:500" json:"summary"`
	CoverImage   string     `gorm:"size:255" json:"cover_image"`
	Status       string     `gorm:"default:draft;size:20;index" json:"status"` // draft, published, archived
	ViewCount    int        `gorm:"default:0" json:"view_count"`
	LikeCount    int        `gorm:"default:0" json:"like_count"`
	CommentCount int        `gorm:"default:0" json:"comment_count"`
	PublishedAt  *time.Time `json:"published_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	
	// 外键
	UserID     uint `gorm:"not null;index" json:"user_id"`
	CategoryID uint `gorm:"index" json:"category_id"`
	
	// 关联
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Category  Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags      []Tag     `gorm:"many2many:article_tags;" json:"tags,omitempty"`
	Comments  []Comment `gorm:"foreignKey:ArticleID" json:"comments,omitempty"`
}

func (Article) TableName() string {
	return "articles"
}

// BeforeCreate 创建前钩子
func (a *Article) BeforeCreate(tx *gorm.DB) error {
	// 生成slug（如果未提供）
	if a.Slug == "" {
		a.Slug = generateSlug(a.Title)
	}
	
	// 如果状态是published，设置发布时间
	if a.Status == "published" && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}
	
	return nil
}

// BeforeUpdate 更新前钩子
func (a *Article) BeforeUpdate(tx *gorm.DB) error {
	// 如果状态从draft变为published，设置发布时间
	if a.Status == "published" && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}
	
	return nil
}
```

### 5. 评论模型

创建 `internal/model/comment.go`:

```go
package model

import (
	"time"
)

// Comment 评论模型
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Status    string    `gorm:"default:pending;size:20;index" json:"status"` // pending, approved, rejected
	IP        string    `gorm:"size:45" json:"-"` // IP地址
	UserAgent string    `gorm:"size:255" json:"-"` // 用户代理
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// 外键
	ArticleID uint  `gorm:"not null;index" json:"article_id"`
	UserID    uint  `gorm:"index" json:"user_id"`
	ParentID  *uint `gorm:"index" json:"parent_id,omitempty"` // 回复的评论ID
	
	// 关联
	Article Article  `gorm:"foreignKey:ArticleID" json:"article,omitempty"`
	User    User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Parent  *Comment `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
}

func (Comment) TableName() string {
	return "comments"
}
```

## 🔗 关联关系

### 一对一关系

```go
// User 和 UserProfile 一对一
type UserProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"uniqueIndex"`
	User   User `gorm:"foreignKey:UserID"`
}
```

### 一对多关系

```go
// User 和 Article 一对多
type User struct {
	Articles []Article `gorm:"foreignKey:UserID"`
}

type Article struct {
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
```

### 多对多关系

```go
// Article 和 Tag 多对多
type Article struct {
	Tags []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	Articles []Article `gorm:"many2many:article_tags;"`
}
```

### 自关联关系

```go
// Comment 自关联（评论回复）
type Comment struct {
	ParentID *uint
	Parent   *Comment `gorm:"foreignKey:ParentID"`
	Replies  []Comment `gorm:"foreignKey:ParentID"`
}
```

## 🗄️ 数据库迁移

### 自动迁移

创建 `internal/repository/migrate.go`:

```go
package repository

import (
	"blog-system/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Tag{},
		&model.Article{},
		&model.Comment{},
	)
}
```

### 手动迁移（可选）

如果需要更精细的控制，可以创建迁移脚本：

```go
func Migrate(db *gorm.DB) error {
	// 创建用户表
	if err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			nickname VARCHAR(50),
			avatar VARCHAR(255),
			role VARCHAR(20) DEFAULT 'user',
			status VARCHAR(20) DEFAULT 'active',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at DATETIME NULL,
			INDEX idx_username (username),
			INDEX idx_email (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`).Error; err != nil {
		return err
	}

	// 创建其他表...
	
	return nil
}
```

## 🔧 模型方法

### 业务方法

```go
// User 模型方法
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) IsActive() bool {
	return u.Status == "active"
}

// Article 模型方法
func (a *Article) IsPublished() bool {
	return a.Status == "published"
}

func (a *Article) IncrementViewCount() {
	a.ViewCount++
}

// Comment 模型方法
func (c *Comment) IsApproved() bool {
	return c.Status == "approved"
}

func (c *Comment) IsReply() bool {
	return c.ParentID != nil
}
```

### 查询作用域

```go
// 定义查询作用域
func PublishedArticles(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", "published")
}

func ActiveUsers(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", "active")
}

// 使用作用域
var articles []Article
db.Scopes(PublishedArticles).Find(&articles)
```

## 📝 索引优化

### 创建索引

```go
// 在模型定义中使用索引标签
type Article struct {
	Title string `gorm:"index"`                    // 单列索引
	Status string `gorm:"index:idx_status"`        // 命名索引
	UserID uint `gorm:"index:idx_user_status"`     // 复合索引
	CategoryID uint `gorm:"index:idx_user_status"` // 复合索引
}

// 手动创建索引
db.Exec("CREATE INDEX idx_article_status ON articles(status)")
db.Exec("CREATE INDEX idx_article_user_status ON articles(user_id, status)")
```

## ✅ 验证和测试

### 测试模型

```go
func TestUserModel(t *testing.T) {
	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	
	// 测试创建
	if err := db.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	
	// 测试查询
	var foundUser model.User
	db.First(&foundUser, user.ID)
	
	if foundUser.Username != user.Username {
		t.Errorf("期望用户名 %s, 得到 %s", user.Username, foundUser.Username)
	}
}
```

## 📚 最佳实践

### 1. 模型设计原则

- **单一职责**: 每个模型只负责一个实体
- **合理关联**: 避免过度关联
- **索引优化**: 为常用查询字段创建索引
- **软删除**: 重要数据使用软删除

### 2. 命名规范

- 表名使用复数形式（users, articles）
- 字段名使用下划线命名（user_id, created_at）
- 关联字段使用明确的命名

### 3. 性能优化

- 使用预加载避免 N+1 查询
- 合理使用索引
- 避免过度使用关联查询

## ⏭️ 下一步

数据模型设计完成后，下一步是：

- [用户认证](./03-auth.md) - 实现用户注册、登录和JWT认证

---

**🎉 数据模型设计完成！** 现在你可以开始实现用户认证功能了。

