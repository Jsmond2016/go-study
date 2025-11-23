---
title: GORM 框架
difficulty: intermediate
duration: "5-6小时"
prerequisites: ["MySQL 基础", "结构体"]
tags: ["GORM", "ORM", "数据库", "关联"]
---

# GORM 框架

GORM 是 Go 语言最流行的 ORM 框架，提供了简洁的 API 和强大的功能。

## 📋 学习目标

- [ ] 理解 ORM 的概念和优势
- [ ] 掌握 GORM 的基本使用
- [ ] 学会模型定义和迁移
- [ ] 掌握 CRUD 操作
- [ ] 理解关联关系处理
- [ ] 学会事务和钩子函数
- [ ] 了解性能优化

## 🎯 GORM 简介

### 为什么选择 GORM

- **功能完整**: 支持 CRUD、关联、事务等
- **API 简洁**: 链式调用，易于使用
- **性能优秀**: 查询优化，支持预加载
- **文档完善**: 官方文档详细
- **社区活跃**: 使用广泛

### 安装 GORM

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 🚀 快速开始

### 基本连接

```go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	// 获取底层 sql.DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
```

## 📊 模型定义

### 基本模型

```go
package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Age       int            `gorm:"default:0" json:"age"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
```

### 模型标签

```go
type User struct {
	ID        uint   `gorm:"primaryKey"`              // 主键
	Name      string `gorm:"size:100;not null"`       // 长度限制，非空
	Email     string `gorm:"uniqueIndex"`             // 唯一索引
	Age       int    `gorm:"default:0"`               // 默认值
	Status    string `gorm:"type:enum('active','inactive')"` // 枚举类型
	CreatedAt time.Time                               // 自动时间戳
	UpdatedAt time.Time                               // 自动更新时间
}
```

### 自动迁移

```go
func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	// 自动迁移
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
```

## 🔧 CRUD 操作

### 创建

```go
// 创建单条记录
user := User{Name: "张三", Email: "zhangsan@example.com", Age: 25}
result := db.Create(&user)
if result.Error != nil {
	panic(result.Error)
}

// 批量创建
users := []User{
	{Name: "李四", Email: "lisi@example.com", Age: 30},
	{Name: "王五", Email: "wangwu@example.com", Age: 28},
}
db.Create(&users)

// 指定字段创建
db.Select("Name", "Email").Create(&user)

// 忽略字段创建
db.Omit("Age").Create(&user)
```

### 查询

```go
// 查询单条记录
var user User
db.First(&user, 1)                    // 主键查询
db.First(&user, "name = ?", "张三")   // 条件查询

// 查询多条记录
var users []User
db.Find(&users)                       // 查询所有
db.Where("age > ?", 18).Find(&users)  // 条件查询

// 条件查询
db.Where("name = ?", "张三").First(&user)
db.Where("age BETWEEN ? AND ?", 20, 30).Find(&users)
db.Where("name IN ?", []string{"张三", "李四"}).Find(&users)
db.Where("name LIKE ?", "%张%").Find(&users)

// 排序和分页
db.Order("created_at DESC").Limit(10).Offset(0).Find(&users)

// 计数
var count int64
db.Model(&User{}).Where("age > ?", 18).Count(&count)
```

### 更新

```go
// 更新单条记录
db.Model(&user).Update("name", "李四")
db.Model(&user).Updates(User{Name: "李四", Age: 30})

// 更新多条记录
db.Model(&User{}).Where("age < ?", 18).Update("status", "inactive")

// 更新指定字段
db.Model(&user).Select("name", "age").Updates(User{Name: "李四", Age: 30})

// 忽略字段更新
db.Model(&user).Omit("email").Updates(User{Name: "李四", Email: "new@example.com"})
```

### 删除

```go
// 软删除
db.Delete(&user)

// 硬删除
db.Unscoped().Delete(&user)

// 批量删除
db.Where("age < ?", 18).Delete(&User{})

// 永久删除
db.Unscoped().Where("age < ?", 18).Delete(&User{})
```

## 🔗 关联关系

### 一对一

```go
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Profile Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint
	Bio    string
	Avatar string
}

// 查询关联
var user User
db.Preload("Profile").First(&user, 1)
```

### 一对多

```go
type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string
	Posts  []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint
	Title  string
	Content string
}

// 查询关联
var user User
db.Preload("Posts").First(&user, 1)
```

### 多对多

```go
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Users []User `gorm:"many2many:user_roles;"`
}

// 关联操作
var user User
var role Role
db.First(&user, 1)
db.First(&role, 1)
db.Model(&user).Association("Roles").Append(&role)
```

## 🔄 事务

```go
// 自动事务
db.Transaction(func(tx *gorm.DB) error {
	if err := tx.Create(&user1).Error; err != nil {
		return err
	}
	
	if err := tx.Create(&user2).Error; err != nil {
		return err
	}
	
	return nil
})

// 手动事务
tx := db.Begin()
defer func() {
	if r := recover(); r != nil {
		tx.Rollback()
	}
}()

if err := tx.Create(&user1).Error; err != nil {
	tx.Rollback()
	return err
}

if err := tx.Create(&user2).Error; err != nil {
	tx.Rollback()
	return err
}

tx.Commit()
```

## 🎣 钩子函数

```go
type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 创建前钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 加密密码
	u.Password = hashPassword(u.Password)
	return nil
}

// 创建后钩子
func (u *User) AfterCreate(tx *gorm.DB) error {
	// 发送欢迎邮件
	sendWelcomeEmail(u.Email)
	return nil
}

// 更新前钩子
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// 更新时间戳
	u.UpdatedAt = time.Now()
	return nil
}
```

## 🏃‍♂️ 实践应用

### 完整的 CRUD 服务

```go
package services

import (
	"gorm.io/gorm"
	"myproject/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Create(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) List(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	
	offset := (page - 1) * pageSize
	
	s.db.Model(&models.User{}).Count(&total)
	err := s.db.Offset(offset).Limit(pageSize).Find(&users).Error
	
	return users, total, err
}

func (s *UserService) Update(id uint, user *models.User) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (s *UserService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *UserService) Search(keyword string) ([]models.User, error) {
	var users []models.User
	err := s.db.Where("name LIKE ? OR email LIKE ?", 
		"%"+keyword+"%", "%"+keyword+"%").Find(&users).Error
	return users, err
}
```

## ⚠️ 注意事项

### 1. 性能优化

```go
// ✅ 使用预加载避免 N+1 查询
db.Preload("Posts").Find(&users)

// ✅ 使用 Select 指定字段
db.Select("id", "name").Find(&users)

// ✅ 使用索引
db.Model(&User{}).Where("email = ?", email).First(&user)
```

### 2. 错误处理

```go
// ✅ 检查错误
if err := db.Create(&user).Error; err != nil {
	return err
}

// ✅ 检查记录是否存在
if err := db.First(&user, id).Error; err != nil {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("用户不存在")
	}
	return err
}
```

### 3. 连接池配置

```go
sqlDB, _ := db.DB()
sqlDB.SetMaxOpenConns(25)
sqlDB.SetMaxIdleConns(5)
sqlDB.SetConnMaxLifetime(5 * time.Minute)
```

## 📚 扩展阅读

- [GORM 官方文档](https://gorm.io/)
- [GORM 中文文档](https://gorm.io/zh_CN/)

## ⏭️ 下一章节

[Go Modules](./03-go-modules.md) → 学习依赖管理

---

**💡 提示**: GORM 是 Go 开发中最常用的 ORM 框架，掌握它对于 Web 开发非常重要！

