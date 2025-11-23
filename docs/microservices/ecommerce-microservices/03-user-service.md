---
title: 用户服务实现
difficulty: advanced
duration: "3-4小时"
prerequisites: ["架构设计"]
tags: ["用户服务", "gRPC", "服务实现"]
---

# 用户服务实现

本章节将详细介绍如何实现用户服务，包括 Protocol Buffers 定义、gRPC 服务实现、数据存储和服务注册。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 定义用户服务的 Protocol Buffers 接口
- [ ] 实现用户服务的 gRPC 服务端
- [ ] 实现用户数据存储和访问
- [ ] 实现服务注册到 Consul
- [ ] 实现用户认证和验证功能

## 🎯 服务功能

用户服务提供以下功能：

- ✅ 用户注册
- ✅ 用户登录
- ✅ 获取用户信息
- ✅ 更新用户信息
- ✅ 删除用户
- ✅ 验证用户（供其他服务调用）

## 📝 Protocol Buffers 定义

### 创建 user.proto

在 `proto/user/` 目录下创建 `user.proto`：

```protobuf
syntax = "proto3";

package user;

option go_package = "go-study/examples/microservices/06-ecommerce-microservices/proto/user";

// 用户服务
service UserService {
  // 创建用户
  rpc CreateUser(CreateUserRequest) returns (UserResponse);
  // 获取用户
  rpc GetUser(GetUserRequest) returns (UserResponse);
  // 更新用户
  rpc UpdateUser(UpdateUserRequest) returns (UserResponse);
  // 删除用户
  rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);
  // 验证用户
  rpc ValidateUser(ValidateUserRequest) returns (ValidateUserResponse);
  // 用户登录
  rpc Login(LoginRequest) returns (LoginResponse);
}

// 创建用户请求
message CreateUserRequest {
  string username = 1;
  string email = 2;
  string password = 3;
  string full_name = 4;
}

// 获取用户请求
message GetUserRequest {
  int64 user_id = 1;
}

// 更新用户请求
message UpdateUserRequest {
  int64 user_id = 1;
  string email = 2;
  string full_name = 3;
}

// 删除用户请求
message DeleteUserRequest {
  int64 user_id = 1;
}

// 验证用户请求
message ValidateUserRequest {
  int64 user_id = 1;
}

// 用户登录请求
message LoginRequest {
  string username = 1;
  string password = 2;
}

// 用户响应
message UserResponse {
  int64 user_id = 1;
  string username = 2;
  string email = 3;
  string full_name = 4;
  string created_at = 5;
  string updated_at = 6;
}

// 删除用户响应
message DeleteUserResponse {
  bool success = 1;
  string message = 2;
}

// 验证用户响应
message ValidateUserResponse {
  bool valid = 1;
  UserResponse user = 2;
}

// 登录响应
message LoginResponse {
  bool success = 1;
  string token = 2;
  UserResponse user = 3;
}
```

### 生成代码

```bash
cd proto
make proto
```

## 💻 服务实现

### 1. 初始化数据库

创建 `user-service/main.go`：

```go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"

	pb "go-study/examples/microservices/06-ecommerce-microservices/proto/user"
)

var (
	port       = flag.Int("port", 5001, "服务端口")
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
)

func main() {
	flag.Parse()

	// 初始化数据库
	db, err := initDB()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 创建服务实现
	service := NewUserService(db)

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service)

	// 注册到 Consul
	registerToConsul(*port, *consulAddr)

	// 优雅关闭
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("收到关闭信号，开始优雅关闭...")
		s.GracefulStop()
		cancel()
	}()

	log.Printf("用户服务启动在端口 %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务运行失败: %v", err)
	}

	<-ctx.Done()
	log.Println("服务已关闭")
}

// initDB 初始化数据库
func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "user_service.db")
	if err != nil {
		return nil, err
	}

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			full_name TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
```

### 2. 服务注册

```go
// registerToConsul 注册服务到 Consul
func registerToConsul(port int, consulAddr string) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Printf("创建 Consul 客户端失败: %v", err)
		return
	}

	hostname, _ := os.Hostname()
	serviceID := fmt.Sprintf("user-service-%s-%d", hostname, port)

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "user-service",
		Tags:    []string{"grpc", "v1"},
		Address: "localhost",
		Port:    port,
		Check: &consulapi.AgentServiceCheck{
			TCP:      fmt.Sprintf("localhost:%d", port),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Printf("注册服务失败: %v", err)
	} else {
		log.Printf("服务已注册到 Consul: %s", serviceID)
	}
}
```

### 3. 服务实现

```go
// UserService 用户服务实现
type UserService struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

// NewUserService 创建用户服务
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// 检查用户名是否已存在
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", req.Username).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	if count > 0 {
		return nil, fmt.Errorf("用户名已存在")
	}

	// 插入用户（实际应用中应该加密密码）
	result, err := s.db.Exec(
		"INSERT INTO users (username, email, password, full_name) VALUES (?, ?, ?, ?)",
		req.Username, req.Email, req.Password, req.FullName,
	)
	if err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	userID, _ := result.LastInsertId()
	return s.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
}

// GetUser 获取用户
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	var user pb.UserResponse
	var createdAt, updatedAt time.Time

	err := s.db.QueryRow(
		"SELECT id, username, email, full_name, created_at, updated_at FROM users WHERE id = ?",
		req.UserId,
	).Scan(&user.UserId, &user.Username, &user.Email, &user.FullName, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("用户不存在")
	}
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	user.CreatedAt = createdAt.Format(time.RFC3339)
	user.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &user, nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	_, err := s.db.Exec(
		"UPDATE users SET email = ?, full_name = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		req.Email, req.FullName, req.UserId,
	)
	if err != nil {
		return nil, fmt.Errorf("更新用户失败: %w", err)
	}

	return s.GetUser(ctx, &pb.GetUserRequest{UserId: req.UserId})
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	result, err := s.db.Exec("DELETE FROM users WHERE id = ?", req.UserId)
	if err != nil {
		return nil, fmt.Errorf("删除用户失败: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return &pb.DeleteUserResponse{Success: false, Message: "用户不存在"}, nil
	}

	return &pb.DeleteUserResponse{Success: true, Message: "用户已删除"}, nil
}

// ValidateUser 验证用户
func (s *UserService) ValidateUser(ctx context.Context, req *pb.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	user, err := s.GetUser(ctx, &pb.GetUserRequest{UserId: req.UserId})
	if err != nil {
		return &pb.ValidateUserResponse{Valid: false}, nil
	}

	return &pb.ValidateUserResponse{Valid: true, User: user}, nil
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var userID int64
	var password string

	err := s.db.QueryRow(
		"SELECT id, password FROM users WHERE username = ?",
		req.Username,
	).Scan(&userID, &password)

	if err == sql.ErrNoRows {
		return &pb.LoginResponse{Success: false}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 验证密码（实际应用中应该使用加密比较）
	if password != req.Password {
		return &pb.LoginResponse{Success: false}, nil
	}

	// 获取用户信息
	user, err := s.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	// 生成 token（简化版，实际应使用 JWT）
	token := fmt.Sprintf("token-%d-%d", userID, time.Now().Unix())

	return &pb.LoginResponse{
		Success: true,
		Token:   token,
		User:    user,
	}, nil
}
```

## 🚀 运行服务

### 1. 启动 Consul

```bash
consul agent -dev
```

### 2. 运行服务

```bash
cd user-service
go mod tidy
go run main.go
```

### 3. 验证服务

```bash
# 检查 Consul 中的服务
curl http://localhost:8500/v1/agent/services

# 使用 grpcurl 测试（需要安装 grpcurl）
grpcurl -plaintext localhost:5001 list
```

## 🧪 测试

### 创建用户测试

```go
// 测试代码示例
conn, _ := grpc.Dial("localhost:5001", grpc.WithInsecure())
client := pb.NewUserServiceClient(conn)

resp, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
	Username: "testuser",
	Email:    "test@example.com",
	Password: "password123",
	FullName: "Test User",
})
```

## 📚 下一步

用户服务实现完成后，可以继续：

1. [商品服务](./04-product-service.md) - 实现商品服务
2. [订单服务](./05-order-service.md) - 实现订单服务（会调用用户服务）

---

**✅ 用户服务实现完成！** 现在可以继续实现其他服务了。

