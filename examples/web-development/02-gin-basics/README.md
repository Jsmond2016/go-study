# Gin 框架基础示例

本示例展示了 Gin HTTP 框架的核心功能，包括路由处理、中间件、参数验证、JSON 处理、文件操作等。

## 运行示例

```bash
# 初始化 go.mod 并安装依赖
go mod init gin-example
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors

# 启动服务器
go run main.go

# 服务器将在 http://localhost:8080 启动
```

## 功能演示

### 1. 用户管理 API

```bash
# 获取用户列表（支持分页、搜索、排序）
curl "http://localhost:8080/api/v1/users?page=1&page_size=5&search=张三&sort_by=name&sort_order=asc"

# 获取指定用户
curl http://localhost:8080/api/v1/users/1

# 创建新用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新用户","email":"new@example.com","age":25}'

# 更新用户信息
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"更新的用户","age":26}'

# 删除用户
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### 2. 实用功能

```bash
# 健康检查
curl http://localhost:8080/api/v1/health

# 配置信息
curl http://localhost:8080/api/v1/config

# Swagger 文档
curl http://localhost:8080/api/v1/swagger

# 中间件演示
curl -H "X-Request-ID: test-123" http://localhost:8080/api/v1/users/middleware

# 文件下载
curl -O "http://localhost:8080/api/v1/download?filename=test.txt"

# 文件上传
curl -X POST http://localhost:8080/api/v1/upload \
  -F "file=@README.md"
```

## 知识点

### 1. Gin 基础

```go
import "github.com/gin-gonic/gin"

// 创建 Gin 引擎
r := gin.Default() // 包含 Logger 和 Recovery 中间件
r := gin.New()     // 不包含默认中间件

// 设置运行模式
gin.SetMode(gin.DebugMode)   // 开发模式
gin.SetMode(gin.ReleaseMode) // 生产模式
gin.SetMode(gin.TestMode)     // 测试模式

// 启动服务器
r.Run(":8080")
```

### 2. 路由定义

```go
// 基础路由
r.GET("/users", handleGetUsers)
r.POST("/users", handleCreateUser)
r.PUT("/users/:id", handleUpdateUser)
r.DELETE("/users/:id", handleDeleteUser)

// 路由组
v1 := r.Group("/api/v1")
{
    v1.GET("/users", handleGetUsers)
    v1.POST("/users", handleCreateUser)
}

// 支持所有HTTP方法
r.Any("/test", handleTest)
r.NoRoute(handleNotFound) // 处理未匹配的路由
```

### 3. 参数处理

```go
// 路径参数
r.GET("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
})

// 查询参数
r.GET("/users", func(c *gin.Context) {
    page := c.DefaultQuery("page", "1")
    search := c.Query("search")
})

// 表单参数
r.POST("/users", func(c *gin.Context) {
    name := c.PostForm("name")
    email := c.PostForm("email")
})

// 获取所有参数
r.GET("/test", func(c *gin.Context) {
    params := c.Params     // 路径参数
    query := c.QueryMap("") // 查询参数
    form := c.PostFormMap("") // 表单参数
})
```

### 4. JSON 处理

```go
// 结构体绑定和验证
type User struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=150"`
}

// 解析 JSON 请求
func handleCreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // 使用 user...
}

// 返回 JSON 响应
func handleGetUsers(c *gin.Context) {
    users := []User{{ID: 1, Name: "张三"}}
    c.JSON(http.StatusOK, gin.H{"users": users})
}
```

### 5. 响应处理

```go
// JSON 响应
c.JSON(http.StatusOK, gin.H{"message": "success"})

// HTML 响应
c.HTML(http.StatusOK, "index.html", gin.H{"title": "首页"})

// 字符串响应
c.String(http.StatusOK, "Hello, World!")

// 文件响应
c.File("./uploads/file.txt")

// 文件下载
c.FileAttachment("./uploads/file.txt", "download.txt")

// 重定向
c.Redirect(http.StatusMovedPermanently, "/new-location")

// 设置状态码
c.Status(http.StatusCreated)
c.JSON(http.StatusCreated, gin.H{"id": 1})

// 设置响应头
c.Header("X-Custom-Header", "value")
```

### 6. 中间件

```go
// 全局中间件
r.Use(gin.Logger())
r.Use(gin.Recovery())

// 路由组中间件
v1 := r.Group("/api/v1")
v1.Use(authMiddleware())
{
    v1.GET("/users", handleGetUsers)
}

// 单个路由中间件
r.GET("/protected", authMiddleware(), handleProtected)

// 自定义中间件
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if !isValidToken(token) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort() // 阻止继续处理
            return
        }
        c.Next() // 继续处理
    }
}

// 获取上下文值
func handleProtected(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if exists {
        fmt.Println(userID)
    }
}
```

### 7. 数据验证

```go
// 内置验证标签
type User struct {
    Name     string `json:"name" binding:"required"`              // 必填
    Email    string `json:"email" binding:"required,email"`       // 必填且是有效邮箱
    Age      int    `json:"age" binding:"gte=0,lte=150"`         // 0-150之间
    Password string `json:"password" binding:"min=6,max=20"`      // 6-20字符
    URL      string `json:"url" binding:"url"`                   // 有效URL
    Image    string `json:"image" binding:"url,startswith=http"` // 以http开头的URL
}

// 自定义验证器
func customValidator(fl validator.FieldLevel) bool {
    return true // 验证逻辑
}

// 注册自定义验证器
if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("custom", customValidator)
}
```

### 8. 文件处理

```go
// 单文件上传
r.POST("/upload", func(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 保存文件
    c.SaveUploadedFile(file, "./uploads/"+file.Filename)

    c.JSON(http.StatusOK, gin.H{
        "filename": file.Filename,
        "size":     file.Size,
    })
})

// 多文件上传
r.POST("/uploads", func(c *gin.Context) {
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    files := form.File["files"]
    for _, file := range files {
        c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    }

    c.JSON(http.StatusOK, gin.H{"count": len(files)})
})

// 文件下载
r.GET("/download/:filename", func(c *gin.Context) {
    filename := c.Param("filename")
    c.FileAttachment("./uploads/"+filename, filename)
})
```

## 高级功能

### 1. CORS 支持

```go
import "github.com/gin-contrib/cors"

r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}))
```

### 2. 请求限流

```go
func rateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Every(time.Second), 10) // 每秒10个请求

    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

### 3. 静态文件服务

```go
// 静态文件服务器
r.Static("/static", "./static")

// 静态文件（单个）
r.StaticFile("/favicon.ico", "./resources/favicon.ico")

// 虚拟目录
r.StaticFS("/public", http.Dir("./public"))
```

### 4. 模板渲染

```go
// 加载模板
r.LoadHTMLGlob("templates/*")

// 使用模板
r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "首页",
        "users": []User{{Name: "张三"}},
    })
})
```

### 5. WebSocket 支持

```go
import "github.com/gin-gonic/gin"
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

r.GET("/ws", func(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    defer conn.Close()

    // WebSocket 处理逻辑
})
```

## 测试

### 1. 单元测试

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
    gin.SetMode(gin.TestMode)

    router := setupRouter()
    req, _ := http.NewRequest("GET", "/api/v1/users", nil)
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "users")
}
```

### 2. 集成测试

```go
func TestAPIIntegration(t *testing.T) {
    router := setupRouter()

    // 创建用户
    req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(`{
        "name": "测试用户",
        "email": "test@example.com",
        "age": 25
    }`))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, 201, w.Code)

    // 获取用户
    req, _ = http.NewRequest("GET", "/api/v1/users", nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "测试用户")
}
```

## 性能优化

### 1. 连接池配置

```go
func startServer() {
    router := setupRouter()

    server := &http.Server{
        Addr:         ":8080",
        Handler:      router,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    server.ListenAndServe()
}
```

### 2. JSON 优化

```go
// 使用 jsoniter 替代标准库
import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 或在 Gin 中配置
router.Use(func(c *gin.Context) {
    c.Writer = &responseWriter{c.Writer}
    c.Next()
})
```

### 3. 中间件优化

```go
// 避免不必要的中间件
r.Use(gin.Logger()) // 开发环境
r.Use(gin.Recovery()) // 生产环境

// 针对特定路由使用中间件
r.POST("/api/upload", authMiddleware(), handleUpload)
```

## 安全考虑

### 1. 输入验证

```go
// 严格的数据验证
type CreateUserRequest struct {
    Name     string `json:"name" binding:"required,min=2,max=50"`
    Email    string `json:"email" binding:"required,email,max=100"`
    Password string `json:"password" binding:"required,min=8,max=128,containsany=!@#$%^&*"`
}
```

### 2. 安全头设置

```go
func securityHeadersMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("X-Frame-Options", "DENY")
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
        c.Next()
    }
}
```

### 3. HTTPS 支持

```go
func startTLSServer() {
    router := setupRouter()

    router.RunTLS(":443", "cert.pem", "key.pem")
}
```

## 最佳实践

1. **路由设计**：使用 RESTful 风格的路由设计
2. **错误处理**：统一的错误响应格式
3. **日志记录**：使用结构化日志
4. **配置管理**：使用环境变量或配置文件
5. **数据库操作**：使用事务确保数据一致性
6. **API 版本控制**：使用 URL 路径或请求头进行版本控制
7. **文档化**：使用 Swagger 自动生成 API 文档
8. **测试覆盖**：编写完整的单元测试和集成测试

## 常见问题

### 1. 优雅关闭

```go
func main() {
    router := setupRouter()

    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    go server.ListenAndServe()

    // 监听关闭信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    // 优雅关闭
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    server.Shutdown(ctx)
}
```

### 2. 并发控制

```go
var userMutex sync.RWMutex

func handleGetUsers(c *gin.Context) {
    userMutex.RLock()
    defer userMutex.RUnlock()
    // 读取用户数据...
}

func handleCreateUser(c *gin.Context) {
    userMutex.Lock()
    defer userMutex.Unlock()
    // 修改用户数据...
}
```

## 练习

1. 实现完整的用户认证系统（注册、登录、JWT）
2. 添加 API 限流和缓存功能
3. 实现文件上传的进度显示
4. 添加 WebSocket 支持的实时通知
5. 实现 API 版本控制
6. 添加监控和指标收集
7. 实现数据库迁移工具
8. 添加单元测试和集成测试