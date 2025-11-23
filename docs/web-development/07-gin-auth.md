---
title: 认证授权
difficulty: advanced
duration: "4-5小时"
prerequisites: ["Gin 中间件", "数据验证"]
tags: ["Gin", "认证", "授权", "JWT", "Session"]
---

# 认证授权

认证和授权是 Web 应用安全的核心，本教程介绍在 Gin 中实现认证和授权。

## 📋 学习目标

- [ ] 理解认证和授权的概念
- [ ] 掌握 JWT 认证实现
- [ ] 学会 Session 管理
- [ ] 理解权限控制
- [ ] 掌握安全最佳实践

## 🎯 认证基础

### 基本认证

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func basicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
		"user":  "user123",
	})
}

func main() {
	r := gin.Default()

	protected := r.Group("/admin")
	protected.Use(basicAuthMiddleware())
	{
		protected.GET("/secrets", func(c *gin.Context) {
			user := c.MustGet(gin.AuthUserKey).(string)
			c.JSON(200, gin.H{"user": user})
		})
	}

	r.Run(":8080")
}
```

## 🔑 JWT 认证

### 安装依赖

```bash
go get github.com/golang-jwt/jwt/v5
```

### JWT 实现

```go
package main

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your-secret-key")

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(userID int, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "缺少认证令牌"})
			c.Abort()
			return
		}

		// 移除 "Bearer " 前缀
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "无效的认证令牌"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("user_id", claims.UserID)
			c.Set("username", claims.Username)
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 登录接口
	r.POST("/login", func(c *gin.Context) {
		var loginReq struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 验证用户名密码（实际应用中应该查询数据库）
		if loginReq.Username == "admin" && loginReq.Password == "admin123" {
			token, _ := generateToken(1, loginReq.Username)
			c.JSON(200, gin.H{"token": token})
		} else {
			c.JSON(401, gin.H{"error": "用户名或密码错误"})
		}
	})

	// 受保护的路由
	protected := r.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetInt("user_id")
			username := c.GetString("username")
			c.JSON(200, gin.H{
				"user_id":  userID,
				"username": username,
			})
		})
	}

	r.Run(":8080")
}
```

## 🔐 Session 管理

### 使用 Session

```go
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置 Session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user_id", 1)
		session.Set("username", "admin")
		session.Save()

		c.JSON(200, gin.H{"message": "登录成功"})
	})

	r.GET("/profile", func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.JSON(401, gin.H{"error": "未登录"})
			return
		}

		c.JSON(200, gin.H{
			"user_id": userID,
		})
	})

	r.Run(":8080")
}
```

## 🛡️ 权限控制

### 基于角色的访问控制

```go
func roleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")

		hasRole := false
		for _, role := range roles {
			if userRole == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(403, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	api.Use(authMiddleware())
	{
		// 管理员路由
		admin := api.Group("/admin")
		admin.Use(roleMiddleware("admin"))
		{
			admin.GET("/users", getUsers)
		}

		// 普通用户路由
		user := api.Group("/user")
		user.Use(roleMiddleware("user", "admin"))
		{
			user.GET("/profile", getProfile)
		}
	}

	r.Run(":8080")
}
```

## 🏃‍♂️ 实践应用

### 完整的认证系统

```go
package main

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	jwtSecret []byte
}

func NewAuthService(secret string) *AuthService {
	return &AuthService{
		jwtSecret: []byte(secret),
	}
}

func (s *AuthService) GenerateToken(userID int, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func main() {
	authService := NewAuthService("your-secret-key")

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var loginReq struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 验证用户（实际应用中查询数据库）
		if loginReq.Username == "admin" && loginReq.Password == "admin123" {
			token, _ := authService.GenerateToken(1, loginReq.Username)
			c.JSON(200, gin.H{"token": token})
		} else {
			c.JSON(401, gin.H{"error": "认证失败"})
		}
	})

	r.Run(":8080")
}
```

## ⚠️ 安全注意事项

### 1. 密码安全

```go
// ✅ 使用 bcrypt 哈希密码
import "golang.org/x/crypto/bcrypt"

hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

### 2. Token 安全

```go
// ✅ 使用 HTTPS
// ✅ 设置合理的过期时间
// ✅ 使用强密钥
```

### 3. 防止攻击

```go
// ✅ 实现限流
// ✅ 防止暴力破解
// ✅ 使用 CSRF 保护
```

## 📚 扩展阅读

- [JWT 官方文档](https://jwt.io/)
- [Gin Session](https://github.com/gin-contrib/sessions)

## ⏭️ 下一章节

[REST API 设计](./08-rest-api.md) → 学习 RESTful API 设计最佳实践

---

**💡 提示**: 认证和授权是 Web 应用安全的核心，实现时要特别注意安全性！

