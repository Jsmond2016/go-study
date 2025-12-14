---
title: Chapter 6 - 集成测试
difficulty: intermediate
duration: "2-3小时"
tags: ["测试", "集成测试", "API测试"]
---

# Chapter 6: 服务集成与测试

本章将三个微服务集成在一起，并进行完整的端到端测试。

## 📋 学习目标

- [x] 启动所有服务
- [x] 进行集成测试
- [x] API 完整测试流程
- [x] 排查常见问题

## 🚀 启动所有服务

### 方法1：使用启动脚本

```bash
cd 10-bookstore-project
bash scripts/run-all.sh
```

`scripts/run-all.sh` 内容：

```bash
#!/bin/bash

echo "Starting all services..."

# 1. 生成 Proto 代码
bash scripts/gen-proto.sh

# 2. 启动 Book Service（后台）
echo "Starting Book Service..."
cd book-service && go run main.go &
BOOK_PID=$!

# 3. 启动 User Service（后台）
echo "Starting User Service..."
cd ../user-service && go run main.go &
USER_PID=$!

# 等待服务启动
sleep 3

# 4. 启动 API Gateway（前台）
echo "Starting API Gateway..."
cd ../api-gateway && go run main.go

# Ctrl+C 时清理后台进程
trap "kill $BOOK_PID $USER_PID" EXIT
```

### 方法2：手动启动

```bash
# 终端 1
cd book-service && go run main.go

# 终端 2
cd user-service && go run main.go

# 终端 3
cd api-gateway && go run main.go
```

## 🧪 完整测试流程

### 步骤 1: 用户注册

```bash
curl -X POST http://localhost:8080/api/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "email": "alice@example.com",
    "password": "password123",
    "phone": "13800138000"
  }'
```

预期响应：
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "username": "alice",
    "email": "alice@example.com",
    ...
  }
}
```

### 步骤 2: 用户登录

```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "password123"
  }'
```

### 步骤 3: 获取图书列表

```bash
# 保存 token
TOKEN="your_token_here"

curl -X GET "http://localhost:8080/api/books?page=1&page_size=10" \
  -H "Authorization: Bearer $TOKEN"
```

### 步骤 4: 搜索图书

```bash
curl -X GET "http://localhost:8080/api/books/search?keyword=Go&page=1&page_size=5" \
  -H "Authorization: Bearer $TOKEN"
```

### 步骤 5: 创建图书

```bash
curl -X POST http://localhost:8080/api/books \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试图书",
    "author": "测试作者",
    "isbn": "978-1234567890",
    "publisher": "测试出版社",
    "price": 99.99,
    "stock": 50
  }'
```

## 🐛 常见问题排查

### Q1: 端口被占用

```bash
# 检查端口
lsof -i :8080
lsof -i :50051
lsof -i :50052

# 终止进程
kill -9 <PID>
```

### Q2: gRPC 连接失败

检查服务是否启动：
```bash
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50052 list
```

### Q3: Token 验证失败

确保：
- Token 格式正确（Bearer + 空格 + token）
- Token 未过期（24小时有效期）
- User Service 正常运行

## ✅ 测试检查清单

- [ ] 三个服务都成功启动
- [ ] 用户可以注册
- [ ] 用户可以登录并获取 Token
- [ ] 可以获取图书列表
- [ ] 可以搜索图书
- [ ] 可以创建图书
- [ ] 未认证请求被正确拒绝

## 🎯 下一步

**[👉 Chapter 7: Docker 部署](/microservices/bookstore-project/07-deployment)**
