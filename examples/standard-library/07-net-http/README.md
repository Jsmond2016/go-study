# HTTP 客户端示例

本示例展示了 Go 语言 net/http 包的 HTTP 客户端使用，包括 GET、POST 请求和自定义请求等。

## 运行示例

```bash
go run main.go
```

## 知识点

### 基本请求

- `http.Get()`：发送 GET 请求
- `http.Post()`：发送 POST 请求
- `http.NewRequest()`：创建自定义请求

### 请求处理

- 设置请求头：`req.Header.Set()`
- 设置超时：`http.Client{Timeout: ...}`
- 读取响应：`io.ReadAll(resp.Body)`

### HTTP 客户端

- `http.Client`：HTTP 客户端，可以设置超时、重定向等
- `http.DefaultClient`：默认客户端

## 练习

1. 实现一个 HTTP 客户端封装，支持重试机制
2. 创建一个下载文件的函数
3. 实现一个简单的 API 客户端库

