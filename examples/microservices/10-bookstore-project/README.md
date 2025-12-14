# 在线书店微服务项目

一个完整的 Go 微服务学习项目，展示了微服务架构的核心概念和最佳实践。

## 📋 项目概述

本项目是一个简化的在线书店系统，包含以下微服务：

- **Book Service** (图书服务) - 图书的 CRUD 操作和搜索
- **User Service** (用户服务) - 用户注册、登录、认证
- **API Gateway** (API 网关) - 统一入口、路由、认证

### 架构图

```
┌─────────────────────────────────────────────────┐
│                 Client (浏览器/移动端)               │
└───────────────────┬─────────────────────────────┘
                    │ HTTP/REST
                    ▼
┌─────────────────────────────────────────────────┐
│            API Gateway (Gin)                    │
│  - HTTP → gRPC 转换                             │
│  - JWT 认证                                     │
│  - 路由转发                                     │
│  - 限流、日志                                   │
└──────────┬──────────────────┬───────────────────┘
           │ gRPC             │ gRPC
           ▼                  ▼
┌──────────────────┐   ┌─────────────────┐
│   User Service   │   │  Book Service   │
│  (用户服务)       │   │  (图书服务)      │
│  - 用户注册       │   │  - 图书管理      │
│  - 用户登录       │   │  - 图书搜索      │
│  - Token 验证     │   │  - 库存管理      │
│  - SQLite        │   │  - SQLite       │
└──────────────────┘   └─────────────────┘
```

## 🎯 学习目标

通过本项目，你将学会：

- ✅ 使用 gRPC 构建微服务
- ✅ 使用 Protocol Buffers 定义服务接口
- ✅ 实现 API 网关模式
- ✅ JWT Token 认证和授权
- ✅ 服务间通信
- ✅ 数据库操作（SQLite）
- ✅ 错误处理和日志记录
- ✅ 项目结构和代码组织

## 📁 项目结构

```
10-bookstore-project/
├── proto/                      # Protocol Buffers 定义
│   ├── book.proto             # 图书服务接口
│   └── user.proto             # 用户服务接口
├── book-service/              # 图书服务
│   └── main.go
├── user-service/              # 用户服务
│   └── main.go
├── api-gateway/               # API 网关
│   └── main.go
├── scripts/                   # 脚本
│   ├── gen-proto.sh          # 生成 proto 代码
│   └── run-all.sh            # 启动所有服务
├── go.mod
├── go.sum
├── docker-compose.yml        # Docker Compose 配置
└── README.md                 # 本文件
```

## 🚀 快速开始

### 前置条件

- Go 1.21+
- Protocol Buffers 编译器 (protoc)
- gRPC 插件

#### 安装 protoc 和插件

```bash
# macOS
brew install protobuf

# 安装 Go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 步骤 1: 生成 Proto 代码

```bash
# 进入项目目录
cd 10-bookstore-project

# 生成 proto 代码
bash scripts/gen-proto.sh
```

### 步骤 2: 启动服务

#### 方式一：手动启动（推荐用于学习）

```bash
# 终端 1: 启动 Book Service
cd book-service
go run main.go

# 终端 2: 启动 User Service
cd user-service
go run main.go -port 50052

# 终端 3: 启动 API Gateway
cd api-gateway
go run main.go
```

#### 方式二：使用脚本

```bash
bash scripts/run-all.sh
```

#### 方式三：使用 Docker Compose

```bash
docker-compose up
```

### 步骤 3: 测试 API

#### 1. 用户注册

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "password": "password123",
    "email": "alice@example.com"
  }'
```

响应：
```json
{
  "user": {
    "id": 1,
    "username": "alice",
    "email": "alice@example.com",
    "created_at": 1234567890
  }
}
```

#### 2. 用户登录

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "password": "password123"
  }'
```

响应：
```json
{
  "user": {...},
  "token": "eyJhbGci...",
  "expires_at": 1234567890
}
```

#### 3. 获取图书列表

```bash
curl http://localhost:8080/api/books?page=1&page_size=10
```

响应：
```json
{
  "books": [
    {
      "id": 1,
      "title": "Go 语言编程",
      "author": "许式伟",
      "price": 59.00,
      "stock": 100
    }
  ],
  "total": 5,
  "page": 1,
  "page_size": 10
}
```

#### 4. 获取图书详情

```bash
curl http://localhost:8080/api/books/1
```

#### 5. 搜索图书

```bash
curl "http://localhost:8080/api/books/search?keyword=Go&page=1"
```

#### 6. 创建图书（需要认证）

```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "title": "Kubernetes 实战",
    "author": "Brendan Burns",
    "isbn": "9787115489234",
    "publisher": "人民邮电出版社",
    "price": 89.00,
    "stock": 50,
    "category": "云原生"
  }'
```

## 📚 API 文档

### User Service APIs

| 接口 | 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|------|
| 注册 | POST | /api/register | 用户注册 | ❌ |
| 登录 | POST | /api/login | 用户登录 | ❌ |
| 获取用户信息 | GET | /api/users/:id | 获取用户详情 | ✅ |

### Book Service APIs

| 接口 | 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|------|
| 图书列表 | GET | /api/books | 分页获取图书 | ❌ |
| 图书详情 | GET | /api/books/:id | 获取图书详情 | ❌ |
| 搜索图书 | GET | /api/books/search | 关键词搜索 | ❌ |
| 创建图书 | POST | /api/books | 添加新图书 | ✅ |
| 更新图书 | PUT | /api/books/:id | 更新图书信息 | ✅ |
| 删除图书 | DELETE | /api/books/:id | 删除图书 | ✅ |

## 🔧 技术栈

### 核心技术

- **gRPC** - 服务间通信
- **Protocol Buffers** - 数据序列化
- **Gin** - HTTP 框架（API Gateway）
- **SQLite** - 数据库
- **JWT** - 身份认证

### 项目依赖

```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/golang-jwt/jwt/v5 v5.2.0
    github.com/mattn/go-sqlite3 v1.14.18
    google.golang.org/grpc v1.59.0
    google.golang.org/protobuf v1.31.0
)
```

## 💡 核心概念解析

### 1. gRPC 服务定义

```protobuf
service BookService {
  rpc GetBook(GetBookRequest) returns (BookResponse);
  rpc ListBooks(ListBooksRequest) returns (ListBooksResponse);
  rpc CreateBook(CreateBookRequest) returns (BookResponse);
}
```

- 使用 Protocol Buffers 定义服务接口
- 强类型，跨语言支持
- 自动生成客户端和服务端代码

### 2. API 网关模式

```
Client → API Gateway (HTTP) → Microservices (gRPC)
```

**优点**：
- 统一入口
- 协议转换（HTTP → gRPC）
- 集中认证和授权
- 限流、日志、监控

**API Gateway 职责**：
- 路由请求到对应的微服务
- JWT Token 验证
- HTTP 请求转换为 gRPC 调用
- 统一错误处理

### 3. JWT 认证流程

```
1. 用户登录 → User Service
2. 验证成功 → 生成 JWT Token
3. 返回 Token 给客户端
4. 客户端请求携带 Token
5. API Gateway 验证 Token
6. 验证通过 → 调用微服务
```

### 4. 服务间通信

```go
// API Gateway 调用 Book Service
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
client := pb.NewBookServiceClient(conn)

resp, err := client.GetBook(ctx, &pb.GetBookRequest{Id: 1})
```

## 🏗️ 代码结构

### Book Service (图书服务)

```go
type BookServer struct {
    db *sql.DB
}

func (s *BookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.BookResponse, error) {
    // 1. 从数据库查询图书
    // 2. 封装为 protobuf 消息
    // 3. 返回响应
}
```

**职责**：
- 图书数据管理
- 数据库操作
- 业务逻辑处理

### User Service (用户服务)

```go
type UserServer struct {
    db *sql.DB
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    // 1. 验证用户名密码
    // 2. 生成 JWT Token
    // 3. 返回用户信息和 Token
}
```

**职责**：
- 用户认证和授权
- JWT Token 管理
- 用户信息管理

### API Gateway (API 网关)

```go
func main() {
    r := gin.Default()

    // 公开接口
    r.POST("/api/register", registerHandler)
    r.POST("/api/login", loginHandler)
    r.GET("/api/books", listBooksHandler)

    // 需要认证的接口
    authorized := r.Group("/api")
    authorized.Use(authMiddleware())
    {
        authorized.POST("/books", createBookHandler)
        authorized.PUT("/books/:id", updateBookHandler)
    }
}
```

**职责**：
- HTTP 请求处理
- 路由转发
- JWT 认证中间件
- gRPC 客户端管理

## 🧪 测试

### 单元测试

```bash
# 测试 Book Service
cd book-service
go test -v

# 测试 User Service
cd user-service
go test -v
```

### 集成测试

```bash
# 启动所有服务后运行集成测试
go test -v ./tests/integration/...
```

### 使用 Postman

导入 Postman Collection:
```bash
# 导入 postman/bookstore.postman_collection.json
```

## 📝 常见问题

### Q1: Proto 代码生成失败？

**A**: 确保已安装 protoc 和相关插件：
```bash
which protoc
which protoc-gen-go
which protoc-gen-go-grpc
```

### Q2: 服务无法连接？

**A**: 检查端口是否被占用：
```bash
lsof -i :50051  # Book Service
lsof -i :50052  # User Service
lsof -i :8080   # API Gateway
```

### Q3: 数据库文件位置？

**A**: SQLite 数据库文件在各服务目录下：
- `book-service/bookstore.db`
- `user-service/users.db`

### Q4: JWT Token 过期？

**A**: 默认过期时间为 24 小时，可在 User Service 中配置：
```go
expirationTime := time.Now().Add(24 * time.Hour)
```

## 🎓 学习路径

### 初级（理解基础）

1. 阅读 proto 定义，理解服务接口
2. 运行服务，测试 API
3. 阅读 Book Service 代码，理解 gRPC 服务实现
4. 阅读 API Gateway 代码，理解 HTTP 到 gRPC 的转换

### 中级（深入理解）

1. 修改 proto 定义，添加新字段
2. 实现新的 API 接口
3. 添加数据验证和错误处理
4. 实现分类筛选和排序功能

### 高级（实战扩展）

1. 添加 Redis 缓存
2. 实现服务发现（Consul）
3. 添加分布式追踪（Jaeger）
4. 部署到 Kubernetes
5. 添加监控和告警（Prometheus）

## 🚀 扩展功能建议

### 可以添加的功能

- [ ] 图书评论系统
- [ ] 购物车功能
- [ ] 订单服务
- [ ] 支付集成
- [ ] 邮件通知
- [ ] 图片上传
- [ ] 全文搜索（Elasticsearch）
- [ ] 推荐系统

### 可以改进的地方

- [ ] 添加单元测试和集成测试
- [ ] 使用连接池优化数据库连接
- [ ] 实现分布式事务
- [ ] 添加 API 限流
- [ ] 实现熔断器
- [ ] 添加健康检查
- [ ] 使用配置中心（Consul/Nacos）

## 📚 参考资源

- [gRPC 官方文档](https://grpc.io/)
- [Protocol Buffers 文档](https://protobuf.dev/)
- [Gin 框架文档](https://gin-gonic.com/)
- [JWT 介绍](https://jwt.io/)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

---

**💡 提示**: 这是一个学习项目，不建议直接用于生产环境。生产环境需要考虑更多安全性、可靠性和性能优化。
