# HTTP 服务器示例

本示例展示了如何使用 Go 标准库 `net/http` 创建功能完整的 HTTP 服务器，包括路由处理、中间件、文件操作等。

## 运行示例

```bash
# 启动服务器
go run main.go

# 服务器将在 http://localhost:8080 启动
```

## 功能演示

### 1. 基础 HTTP 操作

访问 `http://localhost:8080` 查看欢迎页面和API文档。

### 2. REST API 端点

```bash
# 获取用户列表
curl http://localhost:8080/users

# 创建新用户
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新用户","email":"new@example.com","age":25}'

# 获取指定用户
curl http://localhost:8080/users/1

# 更新用户
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"更新的用户","age":26}'

# 删除用户
curl -X DELETE http://localhost:8080/users/1
```

### 3. 实用功能

```bash
# 健康检查
curl http://localhost:8080/health

# 查看请求头
curl -H "Custom-Header: Test" http://localhost:8080/headers

# 重定向示例
curl http://localhost:8080/redirect?to=/health

# 文件下载
curl -O http://localhost:8080/download

# 文件上传
curl -X POST http://localhost:8080/upload \
  -F "file=@README.md"
```

### 4. 中间件演示

```bash
# 使用认证中间件
curl -H "Authorization: Bearer valid-token" \
  http://localhost:8080/middleware
```

## 知识点

### 1. HTTP 服务器基础

```go
import "net/http"

// 创建简单的服务器
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. 路由处理

```go
// 方法路由
func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleGet(w, r)
    case http.MethodPost:
        handlePost(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// 路径参数处理
func handleUser(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/users/")
    id := strings.TrimSuffix(path, "/")
    // 处理用户ID
}
```

### 3. JSON 处理

```go
// 发送 JSON 响应
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

// 接收 JSON 请求
var user User
if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
}
```

### 4. 请求和响应处理

```go
// 设置响应头
w.Header().Set("Content-Type", "application/json")
w.Header().Set("Cache-Control", "no-cache")

// 设置状态码
w.WriteHeader(http.StatusCreated)

// 获取请求信息
method := r.Method
url := r.URL.String()
userAgent := r.UserAgent()
remoteAddr := r.RemoteAddr

// 获取请求头
contentType := r.Header.Get("Content-Type")
auth := r.Header.Get("Authorization")

// 获取查询参数
query := r.URL.Query()
name := query.Get("name")

// 获取表单数据
r.ParseForm()
username := r.FormValue("username")
```

### 5. 文件操作

```go
// 文件下载
func handleDownload(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Header().Set("Content-Disposition", `attachment; filename="example.txt"`)
    fmt.Fprint(w, "文件内容")
}

// 文件上传
func handleUpload(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10MB
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // 处理文件内容
    content, _ := io.ReadAll(file)
    fmt.Printf("上传文件: %s, 大小: %d\n", handler.Filename, len(content))
}
```

### 6. 中间件

```go
// 中间件函数类型
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 日志中间件
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next(w, r)
        log.Printf("Completed in %v", time.Since(start))
    }
}

// 认证中间件
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if !isValidToken(token) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}

// 使用中间件
http.HandleFunc("/protected", loggingMiddleware(authMiddleware(handleProtected)))
```

### 7. 重定向

```go
// 临时重定向 (302)
http.Redirect(w, r, "/new-location", http.StatusFound)

// 永久重定向 (301)
http.Redirect(w, r, "/new-location", http.StatusMovedPermanently)
```

## 实际应用场景

### 1. REST API 服务器

```go
type APIServer struct {
    users map[int]*User
}

func (s *APIServer) Start() {
    http.HandleFunc("/api/users", s.handleUsers)
    http.HandleFunc("/api/users/", s.handleUser)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. 静态文件服务器

```go
func staticFileServer() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 3. 代理服务器

```go
func proxyServer() {
    target, _ := url.Parse("http://backend-server:8080")
    proxy := httputil.NewSingleHostReverseProxy(target)
    http.Handle("/", proxy)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 性能优化

### 1. 使用缓冲区

```go
func bufferedResponse(w http.ResponseWriter, data []byte) {
    buf := bufio.NewWriter(w)
    buf.Write(data)
    buf.Flush()
}
```

### 2. 连接池

```go
var transport = &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 100,
    IdleConnTimeout:     90 * time.Second,
}

var client = &http.Client{Transport: transport}
```

### 3. 启用 HTTP/2

```go
func startHTTP2Server() {
    server := &http.Server{
        Addr:    ":8443",
        Handler: myHandler,
    }

    log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

## 安全考虑

### 1. 输入验证

```go
func validateUserInput(input string) error {
    if len(input) == 0 || len(input) > 100 {
        return errors.New("invalid input length")
    }
    // 更多验证逻辑
    return nil
}
```

### 2. CORS 支持

```go
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        next(w, r)
    }
}
```

### 3. HTTPS 支持

```go
func startTLSServer() {
    log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
}
```

## 常见问题

### 1. 并发安全性

```go
var usersMutex sync.RWMutex

func handleUsers(w http.ResponseWriter, r *http.Request) {
    usersMutex.RLock()
    defer usersMutex.RUnlock()
    // 安全读取 users
}
```

### 2. 内存泄漏

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // 使用 defer 确保资源清理
    defer r.Body.Close()

    // 处理请求...
}
```

### 3. 超时处理

```go
func startServerWithTimeout() {
    server := &http.Server{
        Addr:         ":8080",
        Handler:      myHandler,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
    log.Fatal(server.ListenAndServe())
}
```

## 测试

### 1. 单元测试

```go
func TestHandleUsers(t *testing.T) {
    req, _ := http.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()
    handleUsers(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}
```

### 2. 集成测试

```go
func TestAPIServer(t *testing.T) {
    server := startTestServer()
    defer server.Close()

    resp, _ := http.Get(server.URL + "/users")
    // 测试响应...
}
```

## 最佳实践

1. **使用结构化日志**：使用 `log` 包或第三方日志库
2. **错误处理**：统一错误响应格式
3. **资源管理**：使用 `defer` 确保资源清理
4. **并发安全**：正确使用互斥锁保护共享数据
5. **性能监控**：添加性能指标和健康检查
6. **配置管理**：使用环境变量或配置文件
7. **优雅关闭**：实现服务器的优雅关闭

## 练习

1. 扩展用户管理系统，添加分页和搜索功能
2. 实现基于 JWT 的认证系统
3. 添加请求限流功能
4. 实现文件上传的进度显示
5. 创建 WebSocket 支持的实时聊天功能
6. 添加 API 版本控制
7. 实现完整的日志和监控系统