---
title: Chapter 4 - User Service 实现
difficulty: intermediate
duration: "3-4小时"
tags: ["gRPC", "JWT", "bcrypt", "认证", "安全"]
---

# Chapter 4: User Service 实现

本章实现用户服务，包含用户注册、登录、JWT Token 生成和验证等核心功能。

## 📋 学习目标

完成本章后，你将：
- [x] 掌握用户注册和登录流程
- [x] 学会使用 bcrypt 加密密码
- [x] 实现 JWT Token 生成和验证
- [x] 理解认证授权最佳实践
- [x] 掌握安全编码规范

## 🏗️ User Service 架构

### 核心功能

1. **用户注册** - bcrypt 密码哈希 + 数据验证
2. **用户登录** - 密码验证 + JWT Token 生成
3. **Token 验证** - API Gateway 使用此接口验证请求
4. **用户信息管理** - 获取和更新用户信息

### 安全设计

```
注册流程：
用户提交密码 → bcrypt 哈希 → 存储哈希值
(明文password123) → ($2a$10$...) → 数据库

登录流程：
用户提交密码 → bcrypt 比较 → 生成 JWT → 返回 Token
(password123) → bcrypt.Compare → JWT.Sign → eyJhbGc...

Token 验证流程：
请求携带 Token → JWT.Parse → 验证签名 → 返回用户信息
(Bearer eyJ...) → jwt.ParseWithClaims → user_id, email
```

## 💻 核心代码实现

### 第一部分：JWT Claims 定义

```go
package main

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	// ... 其他导入
)

const (
	port      = ":50052"
	jwtSecret = "your-secret-key-change-in-production"  // ⚠️ 生产环境使用环境变量
)

// JWT Claims 结构
type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type UserServer struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}
```

**安全要点**：
- ✅ JWT Secret 应存储在环境变量中
- ✅ Claims 包含最少必要信息（user_id, email）
- ✅ 使用 RegisteredClaims（包含过期时间等标准字段）

### 第二部分：用户注册

```go
// Register - 用户注册
func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	log.Printf("Register request: username=%s, email=%s", req.Username, req.Email)

	// 1. 参数验证
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument,
			"username, email and password are required")
	}

	if len(req.Password) < 6 {
		return nil, status.Error(codes.InvalidArgument,
			"password must be at least 6 characters")
	}

	// 2. 检查用户名是否存在
	var exists bool
	err := s.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)",
		req.Username,
	).Scan(&exists)

	if err != nil {
		log.Printf("Error checking username: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	}

	// 3. 检查邮箱是否存在
	err = s.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)",
		req.Email,
	).Scan(&exists)

	if err != nil {
		log.Printf("Error checking email: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "email already exists")
	}

	// 4. 哈希密码（核心安全点）
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,  // Cost = 10
	)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, status.Error(codes.Internal, "failed to process password")
	}

	// 5. 插入用户
	result, err := s.db.Exec(
		"INSERT INTO users (username, email, password_hash, phone) VALUES (?, ?, ?, ?)",
		req.Username, req.Email, string(passwordHash), req.Phone,
	)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	userID, _ := result.LastInsertId()

	// 6. 生成 JWT Token
	token, err := generateToken(userID, req.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// 7. 获取用户信息
	user, err := s.getUserByID(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	log.Printf("User registered successfully: id=%d, username=%s", userID, req.Username)

	return &pb.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
```

**关键点解析**：

1. **密码强度验证**：最少 6 个字符
2. **bcrypt 哈希**：
   - DefaultCost = 10（2^10 = 1024 轮）
   - 自动生成 salt
   - 输出格式：`$2a$10$salt+hash`
3. **唯一性检查**：防止重复注册
4. **立即返回 Token**：用户注册后无需再次登录

### 第三部分：用户登录

```go
// Login - 用户登录
func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	log.Printf("Login request: email=%s", req.Email)

	// 1. 参数验证
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument,
			"email and password are required")
	}

	// 2. 查询用户
	var userID int64
	var username, passwordHash string

	err := s.db.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE email = ?",
		req.Email,
	).Scan(&userID, &username, &passwordHash)

	if err == sql.ErrNoRows {
		// ⚠️ 安全：不要透露用户是否存在
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	if err != nil {
		log.Printf("Error querying user: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	// 3. 验证密码（核心）
	err = bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(req.Password),
	)
	if err != nil {
		// 密码错误，记录日志但返回通用错误
		log.Printf("Invalid password attempt for user: %s", req.Email)
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// 4. 生成 JWT Token
	token, err := generateToken(userID, req.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// 5. 获取用户信息
	user, err := s.getUserByID(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	log.Printf("User logged in successfully: id=%d, username=%s", userID, username)

	return &pb.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
```

**安全最佳实践**：
- ✅ 用户不存在和密码错误返回相同错误信息
- ✅ 防止用户枚举攻击
- ✅ 记录登录失败日志（安全审计）
- ✅ 使用 bcrypt.CompareHashAndPassword 验证

### 第四部分：JWT Token 生成

```go
// generateToken 生成 JWT Token
func generateToken(userID int64, email string) (string, error) {
	// 1. 创建 Claims
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 2. 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. 签名
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
```

**JWT 结构**：
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.    # Header
eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImFsaWNl... # Payload (Claims)
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_...  # Signature
```

**Claims 包含**：
- `user_id`: 用户 ID
- `email`: 用户邮箱
- `exp`: 过期时间（24小时后）
- `iat`: 签发时间
- `nbf`: 生效时间

### 第五部分：JWT Token 验证

```go
// ValidateToken - 验证 JWT Token
func (s *UserServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	log.Printf("ValidateToken request")

	// 1. 解析 Token
	claims, err := parseToken(req.Token)
	if err != nil {
		return &pb.ValidateTokenResponse{
			Valid: false,
			Error: err.Error(),
		}, nil  // 不返回 error，返回 valid=false
	}

	// 2. 检查用户是否存在（防止已删除用户使用旧 Token）
	var exists bool
	err = s.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)",
		claims.UserID,
	).Scan(&exists)

	if err != nil {
		log.Printf("Error checking user: %v", err)
		return &pb.ValidateTokenResponse{
			Valid: false,
			Error: "database error",
		}, nil
	}

	if !exists {
		return &pb.ValidateTokenResponse{
			Valid: false,
			Error: "user not found",
		}, nil
	}

	log.Printf("Token validated successfully: user_id=%d", claims.UserID)

	// 3. 返回验证结果
	return &pb.ValidateTokenResponse{
		Valid:  true,
		UserId: claims.UserID,
		Email:  claims.Email,
	}, nil
}

// parseToken 解析 JWT Token
func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
```

**验证流程**：
1. 解析 Token 并验证签名
2. 检查过期时间
3. 验证用户是否仍然存在
4. 返回用户信息

### 第六部分：更新用户信息

```go
// UpdateUser - 更新用户信息
func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	log.Printf("UpdateUser request: id=%d", req.Id)

	// 1. 检查用户是否存在
	var exists bool
	err := s.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)",
		req.Id,
	).Scan(&exists)

	if err != nil {
		return nil, status.Error(codes.Internal, "database error")
	}
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// 2. 检查用户名是否被占用（如果要更新用户名）
	if req.Username != "" {
		var otherUserExists bool
		err := s.db.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM users WHERE username = ? AND id != ?)",
			req.Username, req.Id,
		).Scan(&otherUserExists)

		if err != nil {
			return nil, status.Error(codes.Internal, "database error")
		}
		if otherUserExists {
			return nil, status.Error(codes.AlreadyExists, "username already exists")
		}
	}

	// 3. 动态构建 UPDATE 语句
	query := "UPDATE users SET updated_at = CURRENT_TIMESTAMP"
	args := []interface{}{}

	if req.Username != "" {
		query += ", username = ?"
		args = append(args, req.Username)
	}
	if req.Phone != "" {
		query += ", phone = ?"
		args = append(args, req.Phone)
	}
	if req.Avatar != "" {
		query += ", avatar = ?"
		args = append(args, req.Avatar)
	}

	query += " WHERE id = ?"
	args = append(args, req.Id)

	// 4. 执行更新
	_, err = s.db.Exec(query, args...)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	// 5. 返回更新后的用户信息
	user, err := s.getUserByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	log.Printf("User updated successfully: id=%d", req.Id)

	return &pb.UserResponse{User: user}, nil
}
```

## 🔒 安全最佳实践总结

### 1. 密码安全

```go
// ✅ 正确：使用 bcrypt
passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// ❌ 错误：使用 MD5/SHA1
hash := md5.Sum([]byte(password))  // 不安全！
```

**为什么 bcrypt？**
- 自动加盐（Salt）
- 可配置计算成本（防暴力破解）
- 行业标准

### 2. JWT 安全

```go
// ✅ 设置合理的过期时间
ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

// ✅ 验证签名算法
if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	return nil, fmt.Errorf("unexpected signing method")
}

// ❌ 不要在 JWT 中存储敏感信息
type Claims struct {
	Password string  // ❌ 危险！
	CreditCard string  // ❌ 危险！
}
```

### 3. 错误信息

```go
// ✅ 正确：不透露用户是否存在
return status.Error(codes.Unauthenticated, "invalid credentials")

// ❌ 错误：泄露信息
return status.Error(codes.NotFound, "user not found")  // 暴露用户存在性
return status.Error(codes.Unauthenticated, "wrong password")  // 暴露密码错误
```

### 4. 数据验证

```go
// ✅ 严格验证
if len(password) < 6 {
	return status.Error(codes.InvalidArgument, "password must be at least 6 characters")
}

// ✅ 验证邮箱格式（可选）
if !isValidEmail(email) {
	return status.Error(codes.InvalidArgument, "invalid email format")
}
```

## 🧪 测试 User Service

### 启动服务

```bash
cd user-service
go run main.go
```

### 使用 grpcurl 测试

#### 1. 用户注册

```bash
grpcurl -plaintext -d '{
  "username": "alice",
  "email": "alice@example.com",
  "password": "password123",
  "phone": "13800138000"
}' localhost:50052 user.UserService/Register
```

#### 2. 用户登录

```bash
grpcurl -plaintext -d '{
  "email": "alice@example.com",
  "password": "password123"
}' localhost:50052 user.UserService/Login
```

#### 3. 验证 Token

```bash
grpcurl -plaintext -d '{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}' localhost:50052 user.UserService/ValidateToken
```

## ✅ 本章检查清单

- [ ] 理解了 bcrypt 密码哈希
- [ ] 掌握了 JWT Token 生成和验证
- [ ] 实现了用户注册和登录
- [ ] 了解了安全最佳实践
- [ ] 服务能够正常启动
- [ ] 测试通过

## 🎯 下一步

User Service 完成！接下来实现 API Gateway：

**[👉 Chapter 5: API Gateway](/microservices/bookstore-project/05-gateway)**

下一章我们将学习：
- Gin 框架 HTTP 服务
- HTTP → gRPC 协议转换
- JWT 认证中间件
- RESTful API 设计

---

**🎉 认证服务完成！现在我们有了安全的用户系统！**
