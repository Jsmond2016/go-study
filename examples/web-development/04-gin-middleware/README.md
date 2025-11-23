# Gin 中间件示例

本示例展示了 Gin 框架的中间件功能，包括日志、认证和错误处理中间件。

## 运行示例

```bash
go run main.go
```

## 测试

```bash
# 公开接口
curl http://localhost:8080/public

# 需要认证的接口（无 token）
curl http://localhost:8080/api/user

# 需要认证的接口（有 token）
curl -H "Authorization: token123" http://localhost:8080/api/user
```

## 知识点

### 中间件创建

- 中间件是返回 `gin.HandlerFunc` 的函数
- 使用 `c.Next()` 调用下一个处理器
- 使用 `c.Abort()` 终止处理链

### 中间件使用

- `r.Use()`：全局中间件
- `group.Use()`：路由组中间件
- 中间件按添加顺序执行

### 常用中间件

- 日志记录
- 认证授权
- 错误处理
- CORS 处理
- 限流

## 练习

1. 实现一个请求限流中间件
2. 创建一个性能监控中间件
3. 实现一个请求日志中间件，记录到文件

