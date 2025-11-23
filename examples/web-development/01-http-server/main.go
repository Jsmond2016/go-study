// Package main 展示 Go 语言 HTTP 服务器的各种用法
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// User 用户结构体
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	CreateAt string `json:"created_at"`
}

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 模拟用户数据存储
var users = []User{
	{1, "张三", "zhangsan@example.com", 25, "2024-01-01T00:00:00Z"},
	{2, "李四", "lisi@example.com", 30, "2024-01-02T00:00:00Z"},
	{3, "王五", "wangwu@example.com", 28, "2024-01-03T00:00:00Z"},
}

func main() {
	fmt.Println("=== Go HTTP 服务器示例 ===")

	// 设置路由
	setupRoutes()

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	fmt.Println("可用的API端点:")
	fmt.Println("  GET    /           - 欢迎页面")
	fmt.Println("  GET    /users      - 获取用户列表")
	fmt.Println("  GET    /users/{id} - 获取指定用户")
	fmt.Println("  POST   /users      - 创建新用户")
	fmt.Println("  PUT    /users/{id} - 更新用户信息")
	fmt.Println("  DELETE /users/{id} - 删除用户")
	fmt.Println("  GET    /health     - 健康检查")
	fmt.Println("  GET    /headers    - 显示请求头")
	fmt.Println("  GET    /redirect   - 重定向示例")
	fmt.Println("  GET    /download   - 文件下载示例")
	fmt.Println("  POST   /upload     - 文件上传示例")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// setupRoutes 设置路由
func setupRoutes() {
	// 基础路由
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUserByID)

	// 实用路由
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/headers", handleHeaders)
	http.HandleFunc("/redirect", handleRedirect)
	http.HandleFunc("/download", handleDownload)
	http.HandleFunc("/upload", handleUpload)

	// 中间件路由
	http.HandleFunc("/middleware", loggingMiddleware(authMiddleware(handleProtected)))
}

// handleHome 处理首页请求
func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Go HTTP 服务器示例</title>
    <meta charset="utf-8">
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .endpoint { background: #f5f5f5; padding: 10px; margin: 10px 0; border-radius: 5px; }
        .method { font-weight: bold; color: #007bff; }
    </style>
</head>
<body>
    <h1>🚀 Go HTTP 服务器示例</h1>
    <p>这是一个演示Go HTTP服务器功能的示例应用。</p>

    <h2>📚 API端点</h2>
    <div class="endpoint"><span class="method">GET</span> /users - 获取用户列表</div>
    <div class="endpoint"><span class="method">GET</span> /users/1 - 获取指定用户</div>
    <div class="endpoint"><span class="method">POST</span> /users - 创建新用户</div>
    <div class="endpoint"><span class="method">PUT</span> /users/1 - 更新用户</div>
    <div class="endpoint"><span class="method">DELETE</span> /users/1 - 删除用户</div>

    <h2>🔧 实用功能</h2>
    <div class="endpoint"><span class="method">GET</span> /health - 健康检查</div>
    <div class="endpoint"><span class="method">GET</span> /headers - 显示请求头</div>
    <div class="endpoint"><span class="method">GET</span> /redirect - 重定向示例</div>
    <div class="endpoint"><span class="method">GET</span> /download - 文件下载</div>
    <div class="endpoint"><span class="method">POST</span> /upload - 文件上传</div>

    <h2>📝 测试命令示例</h2>
    <pre>
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
    </pre>
</body>
</html>
`)
}

// handleUsers 处理用户相关请求
func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r)
	case http.MethodPost:
		handleCreateUser(w, r)
	default:
		writeJSONError(w, http.StatusMethodNotAllowed, "方法不允许")
	}
}

// handleGetUsers 获取用户列表
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Code:    200,
		Message: "获取用户列表成功",
		Data:    users,
	}

	writeJSONResponse(w, response)
}

// handleCreateUser 创建新用户
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		writeJSONError(w, http.StatusBadRequest, "请求格式错误")
		return
	}

	// 验证数据
	if newUser.Name == "" || newUser.Email == "" {
		writeJSONError(w, http.StatusBadRequest, "姓名和邮箱不能为空")
		return
	}

	// 生成新用户ID
	newUser.ID = len(users) + 1
	newUser.CreateAt = time.Now().Format(time.RFC3339)

	// 添加到用户列表
	users = append(users, newUser)

	response := Response{
		Code:    201,
		Message: "用户创建成功",
		Data:    newUser,
	}

	writeJSONResponse(w, response)
}

// handleUserByID 处理单个用户的请求
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中提取用户ID
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id := strings.TrimSuffix(path, "/")

	if id == "" {
		writeJSONError(w, http.StatusBadRequest, "用户ID不能为空")
		return
	}

	// 查找用户
	var targetUser *User
	for i := range users {
		if fmt.Sprintf("%d", users[i].ID) == id {
			targetUser = &users[i]
			break
		}
	}

	if targetUser == nil {
		writeJSONError(w, http.StatusNotFound, "用户不存在")
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, targetUser)
	case http.MethodPut:
		handleUpdateUser(w, r, targetUser)
	case http.MethodDelete:
		handleDeleteUser(w, targetUser)
	default:
		writeJSONError(w, http.StatusMethodNotAllowed, "方法不允许")
	}
}

// handleGetUser 获取指定用户
func handleGetUser(w http.ResponseWriter, user *User) {
	response := Response{
		Code:    200,
		Message: "获取用户成功",
		Data:    user,
	}

	writeJSONResponse(w, response)
}

// handleUpdateUser 更新用户信息
func handleUpdateUser(w http.ResponseWriter, r *http.Request, user *User) {
	var updateData User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		writeJSONError(w, http.StatusBadRequest, "请求格式错误")
		return
	}

	// 更新用户信息（保留ID和创建时间）
	if updateData.Name != "" {
		user.Name = updateData.Name
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}
	if updateData.Age > 0 {
		user.Age = updateData.Age
	}

	response := Response{
		Code:    200,
		Message: "用户更新成功",
		Data:    user,
	}

	writeJSONResponse(w, response)
}

// handleDeleteUser 删除用户
func handleDeleteUser(w http.ResponseWriter, targetUser *User) {
	// 从切片中删除用户
	for i := range users {
		if users[i].ID == targetUser.ID {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	response := Response{
		Code:    200,
		Message: "用户删除成功",
	}

	writeJSONResponse(w, response)
}

// handleHealth 健康检查
func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"uptime":    time.Since(time.Now()).String(), // 这里应该是服务器启动时间
		"version":   "1.0.0",
	}

	writeJSONResponse(w, health)
}

// handleHeaders 显示请求头信息
func handleHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	headers := make(map[string][]string)
	for name, values := range r.Header {
		headers[name] = values
	}

	response := map[string]interface{}{
		"method":     r.Method,
		"url":        r.URL.String(),
		"headers":    headers,
		"user_agent": r.UserAgent(),
		"remote":     r.RemoteAddr,
	}

	writeJSONResponse(w, response)
}

// handleRedirect 重定向示例
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 检查查询参数
	redirectTo := r.URL.Query().Get("to")
	if redirectTo == "" {
		redirectTo = "/users"
	}

	// 设置重定向
	http.Redirect(w, r, redirectTo, http.StatusFound)
}

// handleDownload 文件下载示例
func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 创建示例文件内容
	content := `这是一个示例文件
Go HTTP 服务器下载示例
时间: ` + time.Now().Format(time.RFC3339)

	// 设置响应头
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="example.txt"`)

	// 写入文件内容
	fmt.Fprint(w, content)
}

// handleUpload 文件上传示例
func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 解析表单
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "解析表单失败: "+err.Error())
		return
	}

	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "获取文件失败: "+err.Error())
		return
	}
	defer file.Close()

	// 读取文件内容（实际应用中应该保存到磁盘或存储服务）
	content, err := io.ReadAll(file)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "读取文件失败: "+err.Error())
		return
	}

	response := map[string]interface{}{
		"filename": handler.Filename,
		"size":     handler.Size,
		"header":   handler.Header,
		"content":  string(content)[:min(len(content), 100)], // 只显示前100个字符
	}

	writeJSONResponse(w, response)
}

// writeJSONResponse 写入JSON响应
func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

// writeJSONError 写入JSON错误响应
func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	response := Response{
		Code:    statusCode,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

// 中间件示例
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 记录请求日志
		log.Printf("请求 %s %s", r.Method, r.URL.Path)

		// 创建响应记录器来捕获状态码
		recorder := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		// 调用下一个处理器
		next(recorder, r)

		// 记录响应日志
		log.Printf("响应 %s %s %d %v", r.Method, r.URL.Path, recorder.statusCode, time.Since(start))
	}
}

// authMiddleware 认证中间件
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 检查认证头
		auth := r.Header.Get("Authorization")
		if auth == "" {
			writeJSONError(w, http.StatusUnauthorized, "缺少认证头")
			return
		}

		// 简单的token验证（实际应用中应该使用更安全的方法）
		if auth != "Bearer valid-token" {
			writeJSONError(w, http.StatusUnauthorized, "无效的认证令牌")
			return
		}

		// 认证通过，调用下一个处理器
		next(w, r)
	}
}

// handleProtected 受保护的端点
func handleProtected(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Code:    200,
		Message: "认证成功，访问受保护资源",
		Data: map[string]interface{}{
			"user": "authenticated_user",
			"time": time.Now(),
		},
	}

	writeJSONResponse(w, response)
}

// responseRecorder 用于记录响应状态码
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}