---
title: HTTP 服务器
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["基础语法", "标准库"]
tags: ["HTTP", "服务器", "net/http", "路由", "中间件"]
---

# HTTP 服务器

使用 Go 标准库 `net/http` 构建 HTTP 服务器是学习 Web 开发的基础。

## 📋 学习目标

- [ ] 理解 HTTP 服务器的基本概念
- [ ] 掌握 net/http 包的使用
- [ ] 学会处理 HTTP 请求和响应
- [ ] 理解路由和处理器
- [ ] 掌握中间件的实现
- [ ] 学会文件上传和下载

## 🎯 基本 HTTP 服务器

### 最简单的服务器

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})

	http.ListenAndServe(":8080", nil)
}
```

### 处理不同方法

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET 请求")
	case http.MethodPost:
		fmt.Fprintf(w, "POST 请求")
	case http.MethodPut:
		fmt.Fprintf(w, "PUT 请求")
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE 请求")
	default:
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

## 📥 处理请求

### 读取请求参数

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 查询参数
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(w, "姓名: %s, 年龄: %s", name, age)
}

func main() {
	http.HandleFunc("/user", handler)
	http.ListenAndServe(":8080", nil)
}
```

### 读取请求体

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "读取请求体失败", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "解析JSON失败", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "用户: %+v", user)
}

func main() {
	http.HandleFunc("/user", handler)
	http.ListenAndServe(":8080", nil)
}
```

## 📤 发送响应

### JSON 响应

```go
package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Code:    200,
		Message: "成功",
		Data:    map[string]string{"name": "张三"},
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

### HTML 响应

```go
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>Go Server</title></head>
		<body><h1>Hello, Go!</h1></body>
		</html>
	`)
}
```

## 🛣️ 路由处理

### 基本路由

```go
package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUserByID)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "首页")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "用户列表")
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
	// 从路径中提取ID
	id := r.URL.Path[len("/users/"):]
	fmt.Fprintf(w, "用户ID: %s", id)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
```

### 使用 ServeMux

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/users", handleUsers)

	http.ListenAndServe(":8080", mux)
}
```

## 🔧 中间件

### 日志中间件

```go
package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("请求: %s %s", r.Method, r.URL.Path)

		next(w, r)

		log.Printf("响应: %s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", loggingMiddleware(handler))
	http.ListenAndServe(":8080", nil)
}
```

### 认证中间件

```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "未授权", http.StatusUnauthorized)
			return
		}

		// 验证 token
		if token != "Bearer valid-token" {
			http.Error(w, "无效token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
```

## 📁 文件操作

### 文件上传

```go
func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "解析表单失败", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "获取文件失败", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 保存文件
	// ...

	fmt.Fprintf(w, "文件上传成功: %s", handler.Filename)
}
```

### 文件下载

```go
func handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, "./file.txt")
}
```

## 🏃‍♂️ 实践应用

### 完整的 REST API

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{1, "张三"},
	{2, "李四"},
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		user.ID = len(users) + 1
		users = append(users, user)
		json.NewEncoder(w).Encode(user)
	}
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))

	switch r.Method {
	case http.MethodGet:
		for _, user := range users {
			if user.ID == id {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
		http.Error(w, "用户不存在", http.StatusNotFound)
	case http.MethodDelete:
		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				fmt.Fprintf(w, "删除成功")
				return
			}
		}
		http.Error(w, "用户不存在", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUserByID)
	http.ListenAndServe(":8080", nil)
}
```

## ⚠️ 注意事项

### 1. 错误处理

```go
// ✅ 总是检查错误
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
```

### 2. 资源清理

```go
// ✅ 关闭请求体
defer r.Body.Close()

// ✅ 关闭文件
defer file.Close()
```

### 3. 安全考虑

```go
// ✅ 验证输入
if len(name) > 100 {
	http.Error(w, "名称过长", http.StatusBadRequest)
	return
}

// ✅ 防止路径遍历
if strings.Contains(filename, "..") {
	http.Error(w, "无效文件名", http.StatusBadRequest)
	return
}
```

## 📚 扩展阅读

- [net/http 包官方文档](https://pkg.go.dev/net/http)
- [HTTP 协议详解](https://developer.mozilla.org/zh-CN/docs/Web/HTTP)

## ⏭️ 下一章节

[Gin 基础](./02-gin-basics.md) → 学习 Gin 框架的使用

---

**💡 提示**: 理解标准库 HTTP 服务器是学习 Web 框架的基础，掌握它有助于理解框架的工作原理！

