# 在线书店微服务项目 - 快速上手指南

## ✅ 项目已全部完成

### 项目位置
```
examples/microservices/10-bookstore-project/
```

### 项目包含

1. **✅ Proto 定义**
   - `proto/book.proto` - 图书服务接口（6个RPC方法）
   - `proto/user.proto` - 用户服务接口（5个RPC方法）

2. **✅ Book Service（图书服务）**
   - `book-service/main.go` - 完整实现（500+ 行）
   - 功能：图书CRUD、搜索、分页、库存管理
   - 数据库：SQLite
   - 端口：50051

3. **✅ User Service（用户服务）**
   - `user-service/main.go` - 完整实现（400+ 行）
   - 功能：用户注册、登录、JWT认证、用户信息管理
   - 数据库：SQLite
   - 端口：50052

4. **✅ API Gateway（API 网关）**
   - `api-gateway/main.go` - 完整实现（400+ 行）
   - 功能：HTTP→gRPC转换、JWT验证、路由、CORS
   - 框架：Gin
   - 端口：8080

5. **✅ 脚本和配置**
   - `scripts/gen-proto.sh` - Proto代码生成脚本
   - `scripts/run-all.sh` - 启动所有服务
   - `go.mod` - Go模块配置
   - `README.md` - 完整文档（400+ 行）

6. **✅ 在线文档**
   - VitePress 文档已集成
   - 访问路径：`/microservices/bookstore-project/`

## 🚀 快速开始

### 第一步：生成 Proto 代码

```bash
cd examples/microservices/10-bookstore-project
bash scripts/gen-proto.sh
```

### 第二步：安装依赖和启动服务

所有服务都已完整实现，可以直接启动：

```bash
# 安装依赖
go mod download

# 使用脚本启动所有服务（推荐）
bash scripts/run-all.sh

# 或者手动启动每个服务
# 终端 1: Book Service
cd book-service && go run main.go

# 终端 2: User Service
cd user-service && go run main.go

# 终端 3: API Gateway
cd api-gateway && go run main.go
```

### 第三步：测试 API

完整的 API 测试示例请查看 [README.md](./README.md)

**快速测试**：

```bash
# 1. 用户注册
curl -X POST http://localhost:8080/api/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","email":"alice@example.com","password":"password123"}'

# 2. 用户登录（获取 token）
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password123"}'

# 3. 获取图书列表（使用上一步返回的 token）
curl -X GET "http://localhost:8080/api/books?page=1&page_size=10" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## 📚 项目特点

### 1. 学习友好
- ✅ 完整的架构图和说明
- ✅ 详细的代码注释
- ✅ API 测试示例
- ✅ 常见问题解答

### 2. 实用功能
- ✅ gRPC 服务实现
- ✅ Protocol Buffers 定义
- ✅ 数据库操作（SQLite）
- ✅ 错误处理
- ✅ 日志记录
- ✅ 示例数据

### 3. 完整实现
- ✅ User Service - 用户注册、登录、JWT认证
- ✅ Book Service - 图书CRUD、搜索、库存管理
- ✅ API Gateway - HTTP→gRPC、JWT验证、路由

### 4. 扩展性强
- 可添加 Redis 缓存
- 可添加服务发现
- 可添加分布式追踪
- 可部署到 Kubernetes

## 📖 完整文档

详细说明请查看：
```
examples/microservices/10-bookstore-project/README.md
```

包含：
- 📋 项目概述和架构图
- 🎯 学习目标
- 📁 项目结构说明
- 🚀 启动步骤
- 📚 完整API文档
- 💡 核心概念解析
- 🏗️ 代码结构说明
- 🧪 测试方法
- 📝 常见问题
- 🎓 学习路径（初级→中级→高级）

## 🎯 下一步

1. **理解架构**
   - 阅读 README.md 中的架构图
   - 理解三个服务的职责

2. **运行项目**
   ```bash
   bash scripts/run-all.sh
   ```

3. **测试 API**
   - 使用上面的 curl 命令测试
   - 或使用 Postman 导入测试

4. **阅读代码**
   - 从 proto 定义开始
   - 理解每个服务的实现
   - 学习 JWT 认证流程

5. **扩展功能**
   - 添加订单服务
   - 集成 Redis 缓存
   - 实现服务发现

## 💡 技术亮点

相比现有的 `06-ecommerce-microservices`：

| 特性 | 电商项目 | 书店项目 |
|------|---------| ---------|
| 复杂度 | 高 | 中等 |
| 服务数量 | 4个 | 3个 |
| 学习曲线 | 陡峭 | 平缓 |
| 文档完整度 | 部分 | 完整 |
| 代码注释 | 部分 | 详细 |
| 适合人群 | 有经验 | 初学者 |
| 完整度 | 100% | 100% |

**书店项目**专为学习设计，包含：
- ✅ 更详细的文档
- ✅ 更清晰的代码注释
- ✅ 完整的API测试示例
- ✅ 渐进式学习路径
- ✅ 三个服务全部完整实现

---

**🎉 项目已完整实现！**

完整代码包括：
1. ✅ Book Service - 图书服务完整实现
2. ✅ User Service - 用户服务完整实现
3. ✅ API Gateway - API网关完整实现
4. ✅ Proto 定义 - gRPC 接口定义
5. ✅ 脚本工具 - 自动化启动脚本
6. ✅ 详细文档 - 完整的使用文档

这样设计可以让学习者：
- 通过完整的三个服务学习微服务开发
- 理解 gRPC 服务间通信
- 掌握 JWT 认证流程
- 在实践中深入理解微服务架构

**开始学习微服务开发吧！** 🚀
