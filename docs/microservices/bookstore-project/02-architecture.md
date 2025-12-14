---
title: Chapter 2 - 架构设计
difficulty: intermediate
duration: "1小时"
tags: ["架构设计", "Protocol Buffers", "数据库设计", "API设计"]
---

# Chapter 2: 架构设计

本章将详细讲解系统架构设计、Protocol Buffers 定义、数据库设计和 API 规划。

## 📋 学习目标

完成本章后，你将：
- [x] 理解微服务拆分原则
- [x] 掌握 Protocol Buffers 定义方法
- [x] 设计数据库表结构
- [x] 规划 RESTful API 接口
- [x] 生成 gRPC 代码

## 🏗️ 系统架构设计

### 服务拆分原则

我们按照**领域驱动设计（DDD）**的思想，将系统拆分为三个核心服务：

#### 1. Book Service（图书服务）
**职责**：
- 图书的增删改查（CRUD）
- 图书搜索功能
- 库存管理
- 图书分页查询

**为什么独立**：
- 图书是核心业务实体
- 可能需要独立扩展（高并发查询）
- 业务逻辑相对独立

#### 2. User Service（用户服务）
**职责**：
- 用户注册和登录
- JWT Token 生成和验证
- 用户信息管理

**为什么独立**：
- 用户认证是通用能力
- 可被多个服务复用
- 安全性要求高，需要独立管理

#### 3. API Gateway（API 网关）
**职责**：
- HTTP → gRPC 协议转换
- 统一入口和路由
- 认证授权
- 跨域处理

**为什么需要**：
- 前端通常使用 HTTP/REST
- 后端服务使用 gRPC
- 需要统一的安全控制

### 服务交互流程

```
用户请求流程：

1. 用户登录
   Client → API Gateway → User Service
   ↓
   返回 JWT Token

2. 查询图书（需要认证）
   Client → API Gateway（验证Token）→ User Service（验证）
                    ↓
                Book Service → 返回图书列表

3. 创建图书（需要认证）
   Client → API Gateway（验证Token）→ Book Service
                    ↓
                插入数据库 → 返回结果
```

## 📝 Protocol Buffers 定义

### Book Service Proto

创建 `proto/book.proto`：

```protobuf
syntax = "proto3";

package book;

option go_package = "./proto/book;book";

// ============ 图书服务定义 ============
service BookService {
  // 获取单个图书
  rpc GetBook (GetBookRequest) returns (BookResponse);

  // 获取图书列表（分页）
  rpc ListBooks (ListBooksRequest) returns (ListBooksResponse);

  // 创建图书
  rpc CreateBook (CreateBookRequest) returns (BookResponse);

  // 更新图书
  rpc UpdateBook (UpdateBookRequest) returns (BookResponse);

  // 删除图书
  rpc DeleteBook (DeleteBookRequest) returns (DeleteBookResponse);

  // 搜索图书
  rpc SearchBooks (SearchBooksRequest) returns (ListBooksResponse);
}

// ============ 消息定义 ============

// 图书实体
message Book {
  int64 id = 1;              // 图书ID
  string title = 2;          // 书名
  string author = 3;         // 作者
  string isbn = 4;           // ISBN号
  string publisher = 5;      // 出版社
  double price = 6;          // 价格
  int32 stock = 7;           // 库存
  string description = 8;    // 描述
  string cover_image = 9;    // 封面图片URL
  string created_at = 10;    // 创建时间
  string updated_at = 11;    // 更新时间
}

// 获取图书请求
message GetBookRequest {
  int64 id = 1;
}

// 图书响应
message BookResponse {
  Book book = 1;
}

// 获取图书列表请求
message ListBooksRequest {
  int32 page = 1;       // 页码（从1开始）
  int32 page_size = 2;  // 每页数量
}

// 图书列表响应
message ListBooksResponse {
  repeated Book books = 1;  // 图书列表
  int32 total = 2;          // 总数
  int32 page = 3;           // 当前页
}

// 创建图书请求
message CreateBookRequest {
  string title = 1;
  string author = 2;
  string isbn = 3;
  string publisher = 4;
  double price = 5;
  int32 stock = 6;
  string description = 7;
  string cover_image = 8;
}

// 更新图书请求
message UpdateBookRequest {
  int64 id = 1;
  string title = 2;
  string author = 3;
  string publisher = 4;
  double price = 5;
  int32 stock = 6;
  string description = 7;
  string cover_image = 8;
}

// 删除图书请求
message DeleteBookRequest {
  int64 id = 1;
}

// 删除图书响应
message DeleteBookResponse {
  string message = 1;
}

// 搜索图书请求
message SearchBooksRequest {
  string keyword = 1;   // 搜索关键词
  int32 page = 2;
  int32 page_size = 3;
}
```

**设计要点**：
- ✅ 使用 `int64` 作为 ID 类型（支持大数据量）
- ✅ 价格使用 `double` 类型
- ✅ 时间使用 `string` 类型（便于格式化）
- ✅ 列表请求都支持分页
- ✅ 响应包装在专门的 Response 消息中

### User Service Proto

创建 `proto/user.proto`：

```protobuf
syntax = "proto3";

package user;

option go_package = "./proto/user;user";

// ============ 用户服务定义 ============
service UserService {
  // 用户注册
  rpc Register (RegisterRequest) returns (AuthResponse);

  // 用户登录
  rpc Login (LoginRequest) returns (AuthResponse);

  // 获取用户信息
  rpc GetUser (GetUserRequest) returns (UserResponse);

  // 更新用户信息
  rpc UpdateUser (UpdateUserRequest) returns (UserResponse);

  // 验证Token
  rpc ValidateToken (ValidateTokenRequest) returns (ValidateTokenResponse);
}

// ============ 消息定义 ============

// 用户实体
message User {
  int64 id = 1;           // 用户ID
  string username = 2;    // 用户名
  string email = 3;       // 邮箱
  string phone = 4;       // 手机号
  string avatar = 5;      // 头像URL
  string created_at = 6;  // 创建时间
  string updated_at = 7;  // 更新时间
}

// 注册请求
message RegisterRequest {
  string username = 1;
  string email = 2;
  string password = 3;
  string phone = 4;
}

// 登录请求
message LoginRequest {
  string email = 1;
  string password = 2;
}

// 认证响应（注册和登录共用）
message AuthResponse {
  string token = 1;  // JWT Token
  User user = 2;     // 用户信息
}

// 获取用户请求
message GetUserRequest {
  int64 id = 1;
}

// 用户响应
message UserResponse {
  User user = 1;
}

// 更新用户请求
message UpdateUserRequest {
  int64 id = 1;
  string username = 2;
  string phone = 3;
  string avatar = 4;
}

// 验证Token请求
message ValidateTokenRequest {
  string token = 1;
}

// 验证Token响应
message ValidateTokenResponse {
  bool valid = 1;     // Token是否有效
  int64 user_id = 2;  // 用户ID
  string email = 3;   // 用户邮箱
  string error = 4;   // 错误信息（如果有）
}
```

**设计要点**：
- ✅ 密码不包含在 User 实体中（安全）
- ✅ AuthResponse 同时返回 Token 和用户信息
- ✅ ValidateToken 用于 API Gateway 验证
- ✅ 更新用户时只允许修改特定字段

## 🗄️ 数据库设计

### Books 表（图书表）

```sql
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,              -- 书名
    author TEXT NOT NULL,              -- 作者
    isbn TEXT NOT NULL UNIQUE,         -- ISBN号（唯一）
    publisher TEXT,                    -- 出版社
    price REAL NOT NULL,               -- 价格
    stock INTEGER NOT NULL DEFAULT 0,  -- 库存
    description TEXT,                  -- 描述
    cover_image TEXT,                  -- 封面图片URL
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 索引：提高查询性能
CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
CREATE INDEX IF NOT EXISTS idx_books_author ON books(author);
CREATE INDEX IF NOT EXISTS idx_books_isbn ON books(isbn);
```

**设计说明**：
- `id`: 自增主键
- `isbn`: 唯一约束（每本书的 ISBN 是唯一的）
- `stock`: 默认值为 0
- 为常用查询字段创建索引

### Users 表（用户表）

```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,     -- 用户名（唯一）
    email TEXT NOT NULL UNIQUE,        -- 邮箱（唯一）
    password_hash TEXT NOT NULL,       -- 密码哈希
    phone TEXT,                        -- 手机号
    avatar TEXT,                       -- 头像URL
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 索引
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
```

**设计说明**：
- `username` 和 `email` 都是唯一的
- 存储密码哈希而不是明文密码
- 邮箱和用户名都建立索引（登录查询）

### ER 关系图

```
┌─────────────────┐
│     Users       │
├─────────────────┤
│ id (PK)         │
│ username        │
│ email (UK)      │
│ password_hash   │
│ phone           │
│ avatar          │
│ created_at      │
│ updated_at      │
└─────────────────┘

┌─────────────────┐
│     Books       │
├─────────────────┤
│ id (PK)         │
│ title           │
│ author          │
│ isbn (UK)       │
│ publisher       │
│ price           │
│ stock           │
│ description     │
│ cover_image     │
│ created_at      │
│ updated_at      │
└─────────────────┘

注：本项目中两个表相互独立，
    后续可扩展添加订单表建立关联
```

## 🌐 API 接口设计

### RESTful API 规划

#### 用户相关（公开接口）

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/api/users/register` | 用户注册 | 否 |
| POST | `/api/users/login` | 用户登录 | 否 |

#### 用户相关（需认证）

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| GET | `/api/users/:id` | 获取用户信息 | 是 |
| PUT | `/api/users/:id` | 更新用户信息 | 是 |

#### 图书相关（需认证）

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| GET | `/api/books/:id` | 获取图书详情 | 是 |
| GET | `/api/books` | 获取图书列表 | 是 |
| POST | `/api/books` | 创建图书 | 是 |
| PUT | `/api/books/:id` | 更新图书 | 是 |
| DELETE | `/api/books/:id` | 删除图书 | 是 |
| GET | `/api/books/search` | 搜索图书 | 是 |

### API 请求/响应示例

#### 1. 用户注册

**请求**：
```bash
POST /api/users/register
Content-Type: application/json

{
  "username": "alice",
  "email": "alice@example.com",
  "password": "password123",
  "phone": "13800138000"
}
```

**响应**：
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "username": "alice",
    "email": "alice@example.com",
    "phone": "13800138000",
    "avatar": "",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

#### 2. 获取图书列表

**请求**：
```bash
GET /api/books?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

**响应**：
```json
{
  "books": [
    {
      "id": 1,
      "title": "Go语言圣经",
      "author": "Alan A. A. Donovan",
      "isbn": "978-0134190440",
      "publisher": "Addison-Wesley",
      "price": 89.00,
      "stock": 100,
      "description": "Go语言权威指南",
      "cover_image": "https://example.com/cover1.jpg",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "total": 10,
  "page": 1
}
```

## 🔧 生成 gRPC 代码

### 创建生成脚本

创建 `scripts/gen-proto.sh`：

```bash
#!/bin/bash

echo "生成 Protocol Buffers 代码..."

# 创建输出目录
mkdir -p proto/book proto/user

# 生成 Book Service proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/book.proto

# 生成 User Service proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto

echo "Proto 代码生成完成！"
```

### 执行生成

```bash
# 添加执行权限
chmod +x scripts/gen-proto.sh

# 执行生成
bash scripts/gen-proto.sh
```

### 验证生成结果

```bash
# 查看生成的文件
tree proto/

# 预期输出：
# proto/
# ├── book/
# │   ├── book.pb.go          # 消息定义
# │   └── book_grpc.pb.go     # gRPC 服务定义
# ├── user/
# │   ├── user.pb.go
# │   └── user_grpc.pb.go
# ├── book.proto
# └── user.proto
```

## 📊 架构决策记录（ADR）

### 为什么选择 SQLite？

**决策**：使用 SQLite 作为数据库

**理由**：
- ✅ 零配置，嵌入式数据库
- ✅ 适合学习和演示
- ✅ 单文件存储，易于管理
- ✅ 支持 SQL 标准

**权衡**：
- ❌ 不适合高并发生产环境
- 💡 生产环境建议使用 PostgreSQL/MySQL

### 为什么使用 JWT？

**决策**：使用 JWT 进行认证

**理由**：
- ✅ 无状态认证
- ✅ 适合微服务架构
- ✅ 易于在服务间传递

**权衡**：
- ❌ Token 无法撤销（除非加黑名单）
- 💡 可以设置较短的过期时间

### 为什么需要 API Gateway？

**决策**：使用独立的 API Gateway

**理由**：
- ✅ 前端统一入口
- ✅ 协议转换（HTTP → gRPC）
- ✅ 统一认证和日志

**权衡**：
- ❌ 增加一个服务
- 💡 简化客户端调用

## ✅ 本章检查清单

完成本章后，你应该有：

- [ ] 理解了服务拆分原则
- [ ] 创建了 `proto/book.proto`
- [ ] 创建了 `proto/user.proto`
- [ ] 创建了生成脚本 `scripts/gen-proto.sh`
- [ ] 成功生成了 gRPC 代码
- [ ] 理解了数据库表设计
- [ ] 熟悉了 API 接口规划

## 🐛 常见问题

### Q1: Proto 生成失败

**错误**：`protoc-gen-go: program not found`

**解决**：
```bash
# 确认安装了插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 确认 PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### Q2: 为什么不用 MySQL？

SQLite 是为了简化学习环境。生产环境建议：
- 小型应用：PostgreSQL
- 大型应用：MySQL + Redis

### Q3: API 为什么都需要认证？

为了保护数据安全。实际项目中：
- 图书列表可以公开
- 创建/更新/删除需要认证
- 可以根据需求调整

## 🎯 下一步

架构设计完成！接下来我们开始实现第一个微服务：

**[👉 Chapter 3: Book Service](/microservices/bookstore-project/03-book-service)**

在下一章中，我们将：
- 实现 Book Service 的所有 RPC 方法
- 操作 SQLite 数据库
- 处理错误和日志
- 编写单元测试

---

**💡 提示**

- 保存好 proto 文件，后续修改需要重新生成
- 理解 Proto 定义是微服务开发的基础
- 好的架构设计能让后续开发事半功倍

**🎉 准备好实现第一个服务了吗？继续前进！**
