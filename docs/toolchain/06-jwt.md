---
title: JWT 鉴权
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["基础语法", "Web 开发"]
tags: ["JWT", "认证", "授权", "Token"]
---

# JWT 鉴权

JWT (JSON Web Token) 是一种开放标准，用于在各方之间安全地传输信息。本教程介绍 JWT 的原理和在 Go 中的应用。

## 📋 学习目标

- [ ] 理解 JWT 的原理和结构
- [ ] 掌握 JWT 的生成和解析
- [ ] 学会在 Web 应用中实现 JWT 认证
- [ ] 理解 Token 刷新机制
- [ ] 掌握安全最佳实践
- [ ] 了解常见攻击和防护

## 🎯 JWT 简介

### 什么是 JWT

JWT 是一种紧凑且自包含的方式，用于在各方之间安全地传输信息。它由三部分组成：
- Header（头部）
- Payload（载荷）
- Signature（签名）

### JWT 的优势

- **无状态**: 服务器不需要存储会话信息
- **跨域友好**: 可以在不同域名间使用
- **自包含**: Token 包含所有必要信息
- **可扩展**: 易于添加自定义字段

### 安装 JWT 库

```bash
go get github.com/golang-jwt/jwt/v5
```

## 🚀 快速开始

### 基本结构

```go
package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 定义 Claims
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key")

func main() {
	// 生成 Token
	token, _ := generateToken(1, "zhangsan")
	fmt.Printf("Token: %s\n", token)
	
	// 解析 Token
	claims, _ := parseToken(token)
	fmt.Printf("Claims: %+v\n", claims)
}
```

## 🔑 生成 Token

### 基本生成

```go
func generateToken(userID int, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "myapp",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
```

### 自定义 Claims

```go
type CustomClaims struct {
	UserID   int      `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func generateTokenWithRoles(userID int, username string, roles []string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
```

## 🔍 解析 Token

### 基本解析

```go
func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, fmt.Errorf("无效的 token")
}
```

### 错误处理

```go
func parseTokenWithError(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token 格式错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token 已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("token 尚未生效")
			} else {
				return nil, fmt.Errorf("token 无效")
			}
		}
		return nil, err
	}
	
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, fmt.Errorf("无法解析 claims")
}
```

## 🌐 在 Web 应用中使用

### Gin 中间件

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少认证头"})
			c.Abort()
			return
		}
		
		// 提取 Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "认证头格式错误"})
			c.Abort()
			return
		}
		
		tokenString := parts[1]
		
		// 解析 Token
		claims, err := parseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}
		
		// 设置用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		
		c.Next()
	}
}

func main() {
	r := gin.Default()
	
	// 登录接口
	r.POST("/login", loginHandler)
	
	// 受保护的路由
	protected := r.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.GET("/profile", getProfileHandler)
	}
	
	r.Run(":8080")
}

func loginHandler(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 验证用户名密码（实际应用中查询数据库）
	if loginReq.Username == "admin" && loginReq.Password == "admin123" {
		token, err := generateToken(1, loginReq.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"type":  "Bearer",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
	}
}

func getProfileHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	username := c.GetString("username")
	
	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"username": username,
	})
}
```

## 🔄 Token 刷新

### 刷新机制

```go
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func generateTokenPair(userID int, username string) (*TokenPair, error) {
	// 生成 Access Token（短期有效）
	accessToken, err := generateToken(userID, username)
	if err != nil {
		return nil, err
	}
	
	// 生成 Refresh Token（长期有效）
	refreshClaims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7天
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	
	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
	}, nil
}

func refreshTokenHandler(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 解析 Refresh Token
	claims, err := parseToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 refresh token"})
		return
	}
	
	// 生成新的 Access Token
	newAccessToken, err := generateToken(claims.UserID, claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
```

## 🏃‍♂️ 实践应用

### 完整的 JWT 服务

```go
package jwt

import (
	"errors"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret     []byte
	expiration time.Duration
}

func NewJWTService(secret string, expiration time.Duration) *JWTService {
	return &JWTService{
		secret:     []byte(secret),
		expiration: expiration,
	}
}

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Roles    []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

func (s *JWTService) GenerateToken(userID int, username string, roles []string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myapp",
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JWTService) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, errors.New("无效的 token")
}

func (s *JWTService) ValidateToken(tokenString string) error {
	_, err := s.ParseToken(tokenString)
	return err
}
```

## ⚠️ 安全注意事项

### 1. 密钥管理

```go
// ✅ 使用环境变量或配置文件
secret := os.Getenv("JWT_SECRET")
if secret == "" {
	panic("JWT_SECRET 环境变量未设置")
}

// ❌ 不要硬编码密钥
secret := []byte("hardcoded-secret")
```

### 2. Token 过期时间

```go
// ✅ 设置合理的过期时间
// Access Token: 15分钟 - 1小时
// Refresh Token: 7天 - 30天

accessExpiration := 1 * time.Hour
refreshExpiration := 7 * 24 * time.Hour
```

### 3. HTTPS 传输

```go
// ✅ 生产环境必须使用 HTTPS
// ❌ 不要在 HTTP 上传输敏感 Token
```

### 4. Token 存储

```go
// ✅ 客户端存储（localStorage 或 httpOnly cookie）
// ❌ 不要存储在普通 cookie（容易被 XSS 攻击）
```

### 5. 防止重放攻击

```go
// 添加 jti (JWT ID) 和 nonce
claims := Claims{
	RegisteredClaims: jwt.RegisteredClaims{
		ID:        generateUUID(), // jti
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	},
}
```

## 📚 扩展阅读

- [JWT 官方文档](https://jwt.io/)
- [RFC 7519](https://tools.ietf.org/html/rfc7519)
- [JWT 最佳实践](https://tools.ietf.org/html/draft-ietf-oauth-jwt-bcp-07)

## ⏭️ 下一章节

[Validator 验证](./07-validator.md) → 学习数据验证

---

**💡 提示**: JWT 是现代 Web 应用中最常用的认证方式，掌握它对于构建安全的 API 非常重要！

