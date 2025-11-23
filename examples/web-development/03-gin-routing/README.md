# Gin 路由示例

本示例展示了 Gin 框架的路由功能，包括基本路由、路径参数、查询参数和路由组等。

## 运行示例

```bash
go run main.go
```

## 测试

```bash
# GET 请求
curl http://localhost:8080/get

# POST 请求
curl -X POST http://localhost:8080/post

# 路径参数
curl http://localhost:8080/user/123

# 查询参数
curl "http://localhost:8080/search?q=go&page=2"

# 路由组
curl http://localhost:8080/api/v1/users
```

## 知识点

### 基本路由

- `r.GET()`：GET 请求
- `r.POST()`：POST 请求
- `r.PUT()`：PUT 请求
- `r.DELETE()`：DELETE 请求

### 路径参数

- 使用 `:param` 定义路径参数
- 使用 `c.Param()` 获取路径参数

### 查询参数

- 使用 `c.Query()` 获取查询参数
- 使用 `c.DefaultQuery()` 获取带默认值的查询参数

### 路由组

- 使用 `r.Group()` 创建路由组
- 可以为路由组添加中间件

## 练习

1. 创建一个 RESTful API，包含 CRUD 操作
2. 实现一个带分页的列表接口
3. 创建一个多版本 API 路由结构

