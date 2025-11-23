---
title: TODO API 项目
difficulty: intermediate
duration: "8-10小时"
prerequisites: ["Web 开发", "数据库", "工具链"]
tags: ["项目", "API", "RESTful", "Gin", "SQLite", "CRUD"]
---

# TODO API 项目

这是一个功能完整的 TODO 任务管理 REST API 应用，展示了 Go 语言在企业级项目中的最佳实践。

## 📋 学习目标

完成本项目后，你将能够：

- [ ] 设计和实现完整的 RESTful API
- [ ] 使用 Gin 框架构建 Web 服务
- [ ] 实现数据库操作和 CRUD 功能
- [ ] 使用中间件处理 CORS、日志和错误
- [ ] 实现数据验证和错误处理
- [ ] 掌握项目架构设计和代码组织
- [ ] 编写项目文档和测试

## 🎯 项目概述

### 项目功能

本项目实现一个完整的任务管理 API，包括：

- ✅ **任务管理** - 创建、读取、更新、删除任务
- ✅ **状态管理** - pending、completed、cancelled 三种状态
- ✅ **优先级管理** - low、medium、high 三个优先级
- ✅ **截止日期** - 支持任务截止日期设置
- ✅ **搜索功能** - 按标题和描述搜索任务
- ✅ **分页查询** - 支持大数据量的分页显示
- ✅ **过滤功能** - 按状态、优先级过滤
- ✅ **统计功能** - 任务统计数据

### 技术栈

- **Web 框架**: Gin
- **数据库**: SQLite
- **ORM**: database/sql（标准库）
- **中间件**: CORS、日志、错误处理
- **数据验证**: Gin 绑定验证

## 🏗️ 项目结构

```
01-todo-api/
├── main.go           # 主程序文件
├── go.mod           # Go 模块文件
├── README.md        # 项目文档
├── test-api.sh      # API 测试脚本
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

## 🚀 快速开始

### 1. 环境准备

```bash
# 确保 Go 版本 >= 1.21
go version

# 进入项目目录
cd examples/projects/01-todo-api

# 初始化模块（如果还没有）
go mod init todo-api

# 安装依赖
go mod tidy
```

### 2. 安装依赖

```bash
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/mattn/go-sqlite3
```

### 3. 运行项目

```bash
# 启动服务器
go run main.go

# 服务器将启动在 http://localhost:8080
```

### 4. 测试 API

```bash
# 获取健康状态
curl http://localhost:8080/api/health

# 获取任务列表
curl http://localhost:8080/api/todos

# 创建新任务
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"学习Go","description":"完成教程学习","priority":"high"}'
```

## 📚 详细实现

### 第一步：数据模型设计

#### 1.1 任务结构体

```go
// Todo 任务结构体
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required,min=1,max=200"`
	Description string    `json:"description" binding:"max=1000"`
	Status      string    `json:"status" binding:"oneof=pending completed cancelled"`
	Priority    string    `json:"priority" binding:"oneof=low medium high"`
	DueDate     *string   `json:"due_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
```

**要点说明**：
- 使用 `json` 标签定义 JSON 序列化
- 使用 `binding` 标签进行数据验证
- 使用指针类型 `*string` 和 `*time.Time` 表示可选字段

#### 1.2 请求结构体

```go
// TodoCreateRequest 创建任务请求
type TodoCreateRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description" binding:"max=1000"`
	Priority    string `json:"priority" binding:"oneof=low medium high"`
	DueDate     string `json:"due_date,omitempty"`
}

// TodoUpdateRequest 更新任务请求
type TodoUpdateRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty" binding:"omitempty,oneof=pending completed cancelled"`
	Priority    *string `json:"priority,omitempty" binding:"omitempty,oneof=low medium high"`
	DueDate     *string `json:"due_date,omitempty"`
}
```

**要点说明**：
- 创建请求使用必填字段
- 更新请求使用指针类型，支持部分更新
- 使用 `omitempty` 标签，空值不序列化

#### 1.3 响应结构体

```go
// APIResponse 通用API响应
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// PaginatedResponse 分页响应
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
}
```

### 第二步：数据库设计

#### 2.1 创建数据库表

```go
func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(200) NOT NULL,
		description TEXT,
		status VARCHAR(20) DEFAULT 'pending' CHECK(status IN ('pending', 'completed', 'cancelled')),
		priority VARCHAR(10) DEFAULT 'medium' CHECK(priority IN ('low', 'medium', 'high')),
		due_date DATE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed_at DATETIME
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("创建任务表失败: %w", err)
	}

	// 创建索引
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_todos_status ON todos(status)",
		"CREATE INDEX IF NOT EXISTS idx_todos_priority ON todos(priority)",
		"CREATE INDEX IF NOT EXISTS idx_todos_created_at ON todos(created_at)",
	}

	for _, index := range indexes {
		_, err := db.Exec(index)
		if err != nil {
			return fmt.Errorf("创建索引失败: %w", err)
		}
	}

	return nil
}
```

**要点说明**：
- 使用 `CHECK` 约束限制状态和优先级的值
- 为常用查询字段创建索引
- 使用 `IF NOT EXISTS` 避免重复创建

#### 2.2 初始化数据库

```go
func initDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("数据库连接测试失败:", err)
	}

	// 创建表
	if err := createTables(db); err != nil {
		log.Fatal("创建数据库表失败:", err)
	}

	// 初始化示例数据
	if err := seedData(db); err != nil {
		log.Fatal("初始化示例数据失败:", err)
	}

	fmt.Println("数据库初始化完成")
	return db
}
```

### 第三步：服务层实现

#### 3.1 定义服务接口

```go
// TodoService 任务服务接口
type TodoService interface {
	Create(todo *Todo) error
	GetByID(id int) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int) error
	List(filter TodoFilter) ([]Todo, int, error)
	ToggleStatus(id int, status string) error
}

// TodoFilter 任务查询过滤器
type TodoFilter struct {
	Status   string
	Priority string
	Search   string
	Page     int
	PageSize int
}
```

**要点说明**：
- 使用接口定义服务，便于测试和扩展
- 过滤器结构体封装查询条件

#### 3.2 实现服务方法

```go
// TodoServiceImpl 任务服务实现
type TodoServiceImpl struct {
	db *sql.DB
}

func NewTodoService(db *sql.DB) TodoService {
	return &TodoServiceImpl{db: db}
}

// Create 创建任务
func (s *TodoServiceImpl) Create(todo *Todo) error {
	query := `
		INSERT INTO todos (title, description, status, priority, due_date, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	todo.CreatedAt = now
	todo.UpdatedAt = now
	if todo.Status == "" {
		todo.Status = "pending"
	}

	result, err := s.db.Exec(query, todo.Title, todo.Description, todo.Status,
		todo.Priority, todo.DueDate, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		return fmt.Errorf("创建任务失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取任务ID失败: %w", err)
	}

	todo.ID = int(id)
	return nil
}
```

**要点说明**：
- 使用参数化查询防止 SQL 注入
- 自动设置创建时间和更新时间
- 返回插入的 ID

#### 3.3 实现列表查询（支持分页和过滤）

```go
// List 获取任务列表
func (s *TodoServiceImpl) List(filter TodoFilter) ([]Todo, int, error) {
	// 构建WHERE条件
	whereClause := "WHERE 1=1"
	args := []interface{}{}

	if filter.Status != "" {
		whereClause += " AND status = ?"
		args = append(args, filter.Status)
	}

	if filter.Priority != "" {
		whereClause += " AND priority = ?"
		args = append(args, filter.Priority)
	}

	if filter.Search != "" {
		whereClause += " AND (title LIKE ? OR description LIKE ?)"
		searchTerm := "%" + filter.Search + "%"
		args = append(args, searchTerm, searchTerm)
	}

	// 获取总数
	countQuery := "SELECT COUNT(*) FROM todos " + whereClause
	var total int
	err := s.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("获取任务总数失败: %w", err)
	}

	// 分页查询
	query := `
		SELECT id, title, description, status, priority, due_date,
		       created_at, updated_at, completed_at
		FROM todos ` + whereClause + `
		ORDER BY created_at DESC, priority DESC
		LIMIT ? OFFSET ?
	`

	offset := (filter.Page - 1) * filter.PageSize
	args = append(args, filter.PageSize, offset)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("查询任务列表失败: %w", err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		todo := &Todo{}
		var completedAt sql.NullTime

		err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Priority,
			&todo.DueDate, &todo.CreatedAt, &todo.UpdatedAt, &completedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("扫描任务数据失败: %w", err)
		}

		if completedAt.Valid {
			todo.CompletedAt = &completedAt.Time
		}

		todos = append(todos, *todo)
	}

	return todos, total, nil
}
```

**要点说明**：
- 动态构建 WHERE 条件
- 先查询总数，再分页查询数据
- 使用 `sql.NullTime` 处理可空的日期字段

### 第四步：路由处理

#### 4.1 设置路由

```go
func setupRoutes(r *gin.Engine) {
	// API路由组
	api := r.Group("/api")
	{
		// 健康检查
		api.GET("/health", handleHealth)

		// API文档
		api.GET("/docs", handleDocs)

		// 任务路由
		todos := api.Group("/todos")
		{
			todos.GET("", handleListTodos)
			todos.POST("", handleCreateTodo)
			todos.GET("/:id", handleGetTodo)
			todos.PUT("/:id", handleUpdateTodo)
			todos.DELETE("/:id", handleDeleteTodo)
			todos.PATCH("/:id/toggle", handleToggleTodo)
			todos.GET("/statistics", handleTodoStatistics)
		}
	}
}
```

#### 4.2 实现处理函数

```go
// handleCreateTodo 创建任务
func handleCreateTodo(c *gin.Context) {
	var req TodoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success:   false,
			Message:   "请求参数无效",
			Error:     err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	// 创建任务
	todo := &Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      "pending",
		Priority:    req.Priority,
	}

	if req.DueDate != "" {
		todo.DueDate = &req.DueDate
	}

	if err := todoService.Create(todo); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success:   false,
			Message:   "创建任务失败",
			Error:     err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "任务创建成功",
		Data:    todo,
		Timestamp: time.Now(),
	})
}
```

**要点说明**：
- 使用 `ShouldBindJSON` 进行数据绑定和验证
- 统一的错误响应格式
- 正确的 HTTP 状态码

### 第五步：中间件实现

#### 5.1 CORS 中间件

```go
func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
```

#### 5.2 日志中间件

```go
func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
```

#### 5.3 错误处理中间件

```go
func errorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success:   false,
			Message:   "服务器内部错误",
			Error:     fmt.Sprintf("%v", recovered),
			Timestamp: time.Now(),
		})
		c.Abort()
	})
}
```

## 🔧 API 端点

### 任务管理

| 方法   | 端点                     | 描述                                 |
| ------ | ------------------------ | ------------------------------------ |
| GET    | `/api/todos`             | 获取任务列表（支持分页、搜索、过滤） |
| POST   | `/api/todos`             | 创建新任务                           |
| GET    | `/api/todos/{id}`        | 获取指定任务详情                     |
| PUT    | `/api/todos/{id}`        | 更新任务信息                         |
| DELETE | `/api/todos/{id}`        | 删除任务                             |
| PATCH  | `/api/todos/{id}/toggle` | 切换任务状态                         |
| GET    | `/api/todos/statistics`  | 获取任务统计信息                     |

### 系统管理

| 方法 | 端点          | 描述     |
| ---- | ------------- | -------- |
| GET  | `/api/health` | 健康检查 |
| GET  | `/api/docs`   | API 文档 |

## 📝 请求/响应示例

### 创建任务

**请求**:
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "学习Go语言",
    "description": "完成Go语言基础教程的学习",
    "priority": "high",
    "due_date": "2024-01-15"
  }'
```

**响应**:
```json
{
  "success": true,
  "message": "任务创建成功",
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

### 获取任务列表（分页）

**请求**:
```bash
curl "http://localhost:8080/api/todos?page=1&page_size=10&status=pending&priority=high"
```

**响应**:
```json
{
  "success": true,
  "message": "获取任务列表成功",
  "data": {
    "items": [
      {
        "id": 1,
        "title": "学习Go语言",
        "status": "pending",
        "priority": "high"
      }
    ],
    "page": 1,
    "page_size": 10,
    "total": 25,
    "total_pages": 3
  },
  "timestamp": "2024-01-01T10:00:00Z"
}
```

## 🛡️ 数据验证

### 验证规则

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
  "error": "Key: 'TodoCreateRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag",
  "timestamp": "2024-01-01T10:00:00Z"
}
```

## 🧪 测试

### 手动测试

```bash
# 健康检查
curl http://localhost:8080/api/health

# 创建任务
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"测试任务","priority":"medium"}'

# 获取任务列表
curl "http://localhost:8080/api/todos?page=1&page_size=5"

# 更新任务
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed"}'

# 删除任务
curl -X DELETE http://localhost:8080/api/todos/1
```

## 🚀 扩展功能

### 功能扩展建议

1. **用户认证** - 添加 JWT 认证，支持用户注册、登录
2. **任务标签** - 支持为任务添加多个标签
3. **文件附件** - 支持为任务添加附件
4. **任务分享** - 支持任务分享功能
5. **邮件通知** - 截止日期提醒
6. **任务评论** - 支持任务评论功能

### 技术扩展

1. **数据库** - 支持 MySQL、PostgreSQL
2. **缓存** - 集成 Redis 缓存热点数据
3. **消息队列** - 使用 RabbitMQ、Kafka 处理异步任务
4. **容器化** - Docker、Kubernetes 部署
5. **微服务** - 拆分为微服务架构
6. **API 文档** - 集成 Swagger 自动生成文档

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
        if !isValidToken(token) {
            c.JSON(http.StatusUnauthorized, APIResponse{
                Success: false,
                Message: "未授权",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
```

## 💡 最佳实践

### 代码组织

1. **分层架构** - Model、Service、Handler 分离
2. **接口设计** - 使用接口便于测试和扩展
3. **错误处理** - 统一的错误处理机制
4. **代码复用** - 提取公共逻辑

### API 设计

1. **RESTful 规范** - 遵循 REST 设计原则
2. **统一响应** - 统一的响应格式
3. **状态码** - 正确使用 HTTP 状态码
4. **版本控制** - API 版本管理

## 📚 相关资源

- [Gin 官方文档](https://gin-gonic.com/)
- [SQLite 文档](https://www.sqlite.org/docs.html)
- [RESTful API 设计指南](https://restfulapi.net/)
- [Go 项目布局](https://github.com/golang-standards/project-layout)

## ⏭️ 下一步

完成本项目后，可以：

- 尝试实现扩展功能
- 学习下一个项目：[博客系统](./02-blog-system.md)
- 深入学习 [微服务](../microservices/) 架构
- 学习 [运维部署](../devops/) 相关知识

---

**🎉 恭喜完成 TODO API 项目！** 你已经掌握了 Go Web 开发的核心技能，可以开始构建更复杂的应用了。

