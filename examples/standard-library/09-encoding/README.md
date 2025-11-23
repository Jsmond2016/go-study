# 编码解码示例

本示例展示了 Go 语言的编码解码功能，包括 JSON 和 Base64 编码解码。

## 运行示例

```bash
go run main.go
```

## 知识点

### JSON 编码解码

- `json.Marshal()`：将 Go 对象编码为 JSON
- `json.Unmarshal()`：将 JSON 解码为 Go 对象
- 使用结构体标签 `json:"field"` 控制字段名

### Base64 编码解码

- `base64.StdEncoding`：标准 Base64 编码
- `base64.URLEncoding`：URL 安全的 Base64 编码
- `EncodeToString()`：编码为字符串
- `DecodeString()`：解码字符串

## 练习

1. 实现一个配置文件解析器，支持 JSON 格式
2. 创建一个函数，将图片文件编码为 Base64
3. 实现一个 API 客户端，处理 JSON 请求和响应

