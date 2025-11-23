---
title: HTTP 客户端 (net/http)
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["基础语法", "错误处理"]
tags: ["http", "网络", "客户端", "请求", "响应"]
---

# HTTP 客户端 (net/http)

`net/http` 包提供了 HTTP 客户端和服务器功能。本文档主要介绍 HTTP 客户端的使用。

## 📋 学习目标

- [ ] 掌握 HTTP 请求的发送
- [ ] 理解请求和响应的处理
- [ ] 学会设置请求头和处理 Cookie
- [ ] 掌握超时和重试机制
- [ ] 了解 HTTP 客户端的最佳实践

## 🎯 基本用法

### GET 请求

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(body))
}
```

### POST 请求

```go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	data := map[string]string{
		"name": "张三",
		"age":  "30",
	}
	jsonData, _ := json.Marshal(data)
	
	resp, err := http.Post("https://api.example.com/users", 
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

## 🔧 高级用法

### 自定义请求

```go
package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	// 创建请求
	req, err := http.NewRequest("POST", "https://api.example.com/data", 
		bytes.NewBuffer([]byte("data")))
	if err != nil {
		log.Fatal(err)
	}
	
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

### 设置超时

```go
package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Get("https://api.example.com/data")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

### 处理 Cookie

```go
package main

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	
	// 第一次请求，服务器设置 Cookie
	resp, _ := client.Get("https://api.example.com/login")
	resp.Body.Close()
	
	// 后续请求自动携带 Cookie
	resp2, _ := client.Get("https://api.example.com/data")
	defer resp2.Body.Close()
	
	body, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(body))
}
```

## 🏃‍♂️ 实践应用

### HTTP 客户端封装

```go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		headers: make(map[string]string),
	}
}

func (c *HTTPClient) SetHeader(key, value string) {
	c.headers[key] = value
}

func (c *HTTPClient) Get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	return io.ReadAll(resp.Body)
}

func (c *HTTPClient) Post(path string, data interface{}) ([]byte, error) {
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", c.baseURL+path, 
		bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	return io.ReadAll(resp.Body)
}
```

## ⚠️ 注意事项

### 1. 关闭响应体

```go
// ✅ 总是关闭响应体
resp, err := http.Get(url)
if err != nil {
	return err
}
defer resp.Body.Close()
```

### 2. 错误处理

```go
// ✅ 检查 HTTP 状态码
if resp.StatusCode != http.StatusOK {
	return fmt.Errorf("unexpected status: %d", resp.StatusCode)
}
```

### 3. 超时设置

```go
// ✅ 设置合理的超时时间
client := &http.Client{
	Timeout: 10 * time.Second,
}
```

## 📚 扩展阅读

- [net/http 包官方文档](https://pkg.go.dev/net/http)

## ⏭️ 下一章节

[上下文 (context)](./08-context.md) → 学习 Go 语言的上下文管理

---

**💡 提示**: HTTP 客户端是 Web 开发的基础，掌握它对于构建 API 客户端和集成第三方服务非常重要！

