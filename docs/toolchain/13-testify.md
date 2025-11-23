---
title: 测试框架 Testify
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["基础语法", "测试基础"]
tags: ["Testify", "测试", "单元测试", "Mock"]
---

# 测试框架 Testify

Testify 是 Go 语言最流行的测试工具包，提供了丰富的断言和 Mock 功能。

## 📋 学习目标

- [ ] 理解测试的重要性
- [ ] 掌握 Testify 的基本使用
- [ ] 学会使用断言
- [ ] 理解 Mock 的使用
- [ ] 掌握测试套件
- [ ] 了解测试最佳实践

## 🎯 Testify 简介

### 为什么使用 Testify

- **丰富的断言**: 提供多种断言方法
- **Mock 支持**: 支持接口 Mock
- **测试套件**: 支持测试组织和共享
- **易于使用**: API 简洁直观
- **社区活跃**: 使用广泛

### 安装 Testify

```bash
go get github.com/stretchr/testify
```

## 🚀 快速开始

### 基本测试

```go
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result)
}

func Add(a, b int) int {
	return a + b
}
```

## ✅ 断言

### 基本断言

```go
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAssertions(t *testing.T) {
	// 相等断言
	assert.Equal(t, 5, 5)
	assert.NotEqual(t, 5, 6)
	
	// 布尔断言
	assert.True(t, true)
	assert.False(t, false)
	
	// Nil 断言
	var ptr *int
	assert.Nil(t, ptr)
	assert.NotNil(t, &ptr)
	
	// 错误断言
	err := someFunction()
	assert.NoError(t, err)
	assert.Error(t, err)
	
	// 包含断言
	assert.Contains(t, "Hello World", "World")
	assert.NotContains(t, "Hello", "World")
	
	// 长度断言
	assert.Len(t, []int{1, 2, 3}, 3)
	
	// 空断言
	assert.Empty(t, "")
	assert.NotEmpty(t, "hello")
}
```

### 使用 require

```go
import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRequire(t *testing.T) {
	// require 在失败时会立即停止测试
	require.Equal(t, 5, 5)
	require.NotNil(t, someObject)
}
```

## 🎭 Mock

### 创建 Mock

```go
package main

import (
	"testing"
	"github.com/stretchr/testify/mock"
)

// 定义接口
type UserRepository interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) error
}

// Mock 实现
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

// 使用 Mock
func TestGetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	
	// 设置期望
	expectedUser := &User{ID: 1, Name: "张三"}
	mockRepo.On("GetUser", 1).Return(expectedUser, nil)
	
	// 执行测试
	user, err := mockRepo.GetUser(1)
	
	// 验证
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}
```

## 📦 测试套件

### 使用 Suite

```go
package main

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	service *UserService
}

// SetupTest 在每个测试前执行
func (suite *UserServiceTestSuite) SetupTest() {
	suite.service = NewUserService()
}

// TearDownTest 在每个测试后执行
func (suite *UserServiceTestSuite) TearDownTest() {
	// 清理工作
}

// SetupSuite 在套件开始前执行一次
func (suite *UserServiceTestSuite) SetupSuite() {
	// 初始化工作
}

// TearDownSuite 在套件结束后执行一次
func (suite *UserServiceTestSuite) TearDownSuite() {
	// 清理工作
}

// 测试方法
func (suite *UserServiceTestSuite) TestCreateUser() {
	user := &User{Name: "张三", Email: "zhangsan@example.com"}
	err := suite.service.CreateUser(user)
	suite.NoError(err)
	suite.NotNil(user.ID)
}

func (suite *UserServiceTestSuite) TestGetUser() {
	user, err := suite.service.GetUser(1)
	suite.NoError(err)
	suite.NotNil(user)
}

// 运行测试套件
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
```

## 🏃‍♂️ 实践应用

### HTTP 测试

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	r := gin.New()
	r.GET("/users/:id", GetUser)
	
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
}
```

### 数据库测试

```go
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepository(t *testing.T) {
	// 使用内存数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	
	// 自动迁移
	db.AutoMigrate(&User{})
	
	repo := NewUserRepository(db)
	
	// 创建用户
	user := &User{Name: "张三", Email: "zhangsan@example.com"}
	err = repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	
	// 查询用户
	found, err := repo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, found.Name)
}
```

## ⚠️ 最佳实践

### 1. 测试命名

```go
// ✅ 清晰的测试名称
func TestUserService_CreateUser_Success(t *testing.T) {}
func TestUserService_CreateUser_InvalidEmail(t *testing.T) {}
```

### 2. 测试组织

```go
// ✅ 使用表驱动测试
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
		{"zero", 0, 0, 0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}
```

### 3. Mock 使用

```go
// ✅ 验证 Mock 调用
mockRepo.AssertExpectations(t)

// ✅ 验证调用次数
mockRepo.AssertNumberOfCalls(t, "GetUser", 1)
```

## 📚 扩展阅读

- [Testify 官方文档](https://github.com/stretchr/testify)
- [Go 测试文档](https://golang.org/pkg/testing/)

## ⏭️ 下一章节

[任务调度](./14-cron.md) → 学习定时任务

---

**💡 提示**: 良好的测试是代码质量的保证，Testify 让编写测试变得更加简单高效！

