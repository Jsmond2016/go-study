# TODO API 实战项目

这是一个功能完整的 TODO 任务管理 REST API 应用，展示了 Go 语言在企业级项目中的最佳实践。

## 🚀 快速开始

### 1. 环境准备

```bash
# 确保 Go 版本 >= 1.21
go version

# 克隆项目或复制代码到本地
# 初始化模块
go mod tidy
```

### 2. 运行项目

```bash
# 启动服务器
go run main.go

# 服务器将启动在 http://localhost:8080
```

### 3. 测试 API

```bash
# 获取健康状态
curl http://localhost:8080/api/health

# 获取任务列表
curl http://localhost:8080/api/todos

# 创建新任务
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"学习Go","description":"完成教程学习","priority":"high"}'

# 获取指定任务
curl http://localhost:8080/api/todos/1

# 更新任务
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed","description":"教程学习已完成"}'

# 切换任务状态
curl -X PATCH http://localhost:8080/api/todos/1/toggle

# 删除任务
curl -X DELETE http://localhost:8080/api/todos/1

# 获取统计信息
curl http://localhost:8080/api/todos/statistics
```

## 📋 项目特性

### 核心功能

- ✅ **完整的 CRUD 操作** - 创建、读取、更新、删除任务
- ✅ **状态管理** - pending、completed、cancelled 三种状态
- ✅ **优先级管理** - low、medium、high 三个优先级
- ✅ **截止日期** - 支持任务截止日期设置
- ✅ **搜索功能** - 按标题和描述搜索任务
- ✅ **分页查询** - 支持大数据量的分页显示
- ✅ **过滤功能** - 按状态、优先级过滤
- ✅ **统计功能** - 任务统计数据

### 技术特性

- ✅ **RESTful API 设计** - 遵循 REST 设计原则
- ✅ **数据验证** - 严格的输入验证和错误处理
- ✅ **数据库操作** - 使用 SQLite 数据库
- ✅ **中间件支持** - CORS、日志、错误处理中间件
- ✅ **API 文档** - 内置 API 文档端点
- ✅ **JSON 响应** - 统一的 JSON 响应格式
- ✅ **错误处理** - 完善的错误处理机制

## 🏗️ 项目架构

```
01-todo-api/
├── main.go           # 主程序文件
├── go.mod           # Go 模块文件
├── README.md        # 项目文档
└── todos.db         # SQLite 数据库文件（运行时生成）
```

### 代码架构

```
main.go
├── 数据模型 (Models)
│   ├── Todo              # 任务结构体
│   ├── TodoCreateRequest # 创建请求结构体
│   ├── TodoUpdateRequest # 更新请求结构体
│   ├── APIResponse      # API响应结构体
│   └── PaginatedResponse # 分页响应结构体
├── 服务层 (Services)
│   ├── TodoService       # 任务服务接口
│   └── TodoServiceImpl  # 任务服务实现
├── 路由层 (Handlers)
│   ├── handleHome        # 首页处理
│   ├── handleHealth      # 健康检查
│   ├── handleListTodos   # 获取任务列表
│   ├── handleCreateTodo  # 创建任务
│   ├── handleGetTodo     # 获取任务详情
│   ├── handleUpdateTodo  # 更新任务
│   ├── handleDeleteTodo  # 删除任务
│   ├── handleToggleTodo  # 切换任务状态
│   └── handleTodoStatistics # 统计信息
└── 中间件 (Middleware)
    ├── corsMiddleware    # CORS 跨域中间件
    ├── loggingMiddleware # 日志中间件
    └── errorHandlerMiddleware # 错误处理中间件
```

## 🗄️ 数据库设计

### 任务表结构

```sql
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    status VARCHAR(20) DEFAULT 'pending'
        CHECK(status IN ('pending', 'completed', 'cancelled')),
    priority VARCHAR(10) DEFAULT 'medium'
        CHECK(priority IN ('low', 'medium', 'high')),
    due_date DATE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME
);
```

### 索引设计

```sql
CREATE INDEX idx_todos_status ON todos(status);
CREATE INDEX idx_todos_priority ON todos(priority);
CREATE INDEX idx_todos_created_at ON todos(created_at);
```

## 🔧 API 端点

### 任务管理

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/api/todos` | 获取任务列表（支持分页、搜索、过滤） |
| POST | `/api/todos` | 创建新任务 |
| GET | `/api/todos/{id}` | 获取指定任务详情 |
| PUT | `/api/todos/{id}` | 更新任务信息 |
| DELETE | `/api/todos/{id}` | 删除任务 |
| PATCH | `/api/todos/{id}/toggle` | 切换任务状态 |
| GET | `/api/todos/statistics` | 获取任务统计信息 |

### 系统管理

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/api/health` | 健康检查 |
| GET | `/api/docs` | API 文档 |
| GET | `/` | 主页 |

## 📝 请求/响应格式

### 创建任务请求

```json
{
  "title": "学习Go语言",
  "description": "完成Go语言基础教程的学习",
  "priority": "high",
  "due_date": "2024-01-15"
}
```

### 任务响应

```json
{
  "success": true,
  "message": "操作成功",
  "data": {
    "id": 1,
    "title": "学习Go语言",
    "description": "完成Go语言基础教程的学习",
    "status": "pending",
    "priority": "high",
    "due_date": "2024-01-15",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z",
    "completed_at": null
  },
  "timestamp": "2024-01-01T10:00:00Z"
}
```

### 分页响应

```json
{
  "success": true,
  "message": "获取任务列表成功",
  "data": {
    "items": [...],
    "page": 1,
    "page_size": 10,
    "total": 25,
    "total_pages": 3
  },
  "timestamp": "2024-01-01T10:00:00Z"
}
```

## 🔍 查询参数

### 获取任务列表

- `page` - 页码（默认：1）
- `page_size` - 每页大小（默认：10，最大：100）
- `status` - 状态过滤：pending、completed、cancelled
- `priority` - 优先级过滤：low、medium、high
- `search` - 搜索关键词（搜索标题和描述）

### 示例查询

```bash
# 获取第一页，每页5个，状态为pending的任务
curl "http://localhost:8080/api/todos?page=1&page_size=5&status=pending"

# 搜索包含"Go"关键词的任务
curl "http://localhost:8080/api/todos?search=Go"

# 获取高优先级的已完成任务
curl "http://localhost:8080/api/todos?status=completed&priority=high"
```

## 🛡️ 数据验证

### 输入验证规则

- **title**: 必填，1-200 字符
- **description**: 可选，最大 1000 字符
- **status**: 枚举值：pending、completed、cancelled
- **priority**: 枚举值：low、medium、high
- **due_date**: 日期格式：YYYY-MM-DD

### 错误响应

```json
{
  "success": false,
  "message": "请求参数无效",
  "error": "Title is required and must be between 1 and 200 characters",
  "timestamp": "2024-01-01T10:00:00Z"
}
```

## 🔧 开发和部署

### 开发环境

```bash
# 开发模式启动
gin.SetMode(gin.DebugMode)
go run main.go

# 生产模式启动
gin.SetMode(gin.ReleaseMode)
go run main.go
```

### 构建和部署

```bash
# 构建
go build -o todo-api

# 运行
./todo-api

# 使用 Docker
docker build -t todo-api .
docker run -p 8080:8080 todo-api
```

### 环境变量配置

```bash
# 数据库路径
export DB_PATH="./todos.db"

# 服务器端口
export SERVER_PORT="8080"

# 运行模式
export GIN_MODE="release"
```

## 🧪 测试

### 手动测试

```bash
# 健康检查
curl -f http://localhost:8080/api/health

# 创建任务测试
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"测试任务","priority":"medium"}' \
  -w "\nHTTP Status: %{http_code}\n"

# 分页测试
curl "http://localhost:8080/api/todos?page=1&page_size=5"
```

### 压力测试

```bash
# 使用 ab 进行压力测试
ab -n 1000 -c 10 http://localhost:8080/api/todos

# 使用 hey 进行测试
hey -n 1000 -c 10 http://localhost:8080/api/todos
```

## 📈 性能优化

### 数据库优化

1. **索引优化** - 为常用查询字段创建索引
2. **分页查询** - 避免一次性加载大量数据
3. **连接池** - 配置合适的连接池大小

### API 优化

1. **缓存** - 对热点数据使用缓存
2. **压缩** - 启用 gzip 压缩
3. **限流** - 实现请求限流机制

## 🔒 安全考虑

### 输入验证

- 严格的输入验证和数据类型检查
- SQL 注入防护（使用参数化查询）
- XSS 防护（输出转义）

### 认证授权

```go
// 可以添加 JWT 认证中间件
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        // 验证 token
        c.Next()
    }
}
```

### CORS 配置

```go
// 生产环境应该限制具体域名
cors.New(cors.Config{
    AllowOrigins:     []string{"https://yourdomain.com"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowCredentials: true,
})
```

## 📝 日志和监控

### 日志格式

```json
{
  "timestamp": "2024-01-01T10:00:00Z",
  "method": "GET",
  "path": "/api/todos",
  "status": 200,
  "latency": "2.5ms",
  "client_ip": "127.0.0.1",
  "user_agent": "curl/7.68.0"
}
```

### 监控指标

```go
// 可以添加 Prometheus 监控
func metricsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 记录请求指标
        c.Next()
    }
}
```

## 🚀 扩展功能

### 功能扩展建议

1. **用户认证** - 添加用户注册、登录功能
2. **任务标签** - 支持为任务添加标签
3. **文件附件** - 支持为任务添加附件
4. **任务分享** - 支持任务分享功能
5. **邮件通知** - 截止日期提醒
6. **移动端 API** - 专门的移动端 API

### 技术扩展

1. **数据库** - 支持 MySQL、PostgreSQL
2. **缓存** - 集成 Redis 缓存
3. **消息队列** - 使用 RabbitMQ、Kafka
4. **容器化** - Docker、Kubernetes 部署
5. **微服务** - 拆分为微服务架构

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 联系方式

- 项目链接: [https://github.com/your-username/todo-api](https://github.com/your-username/todo-api)
- 问题反馈: [Issues](https://github.com/your-username/todo-api/issues)

## 🙏 致谢

感谢以下开源项目：

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [SQLite3 Driver](https://github.com/mattn/go-sqlite3)
- [CORS Middleware](https://github.com/gin-contrib/cors)