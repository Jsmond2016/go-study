# RESTful API 示例

本示例展示了如何使用 Gin 框架实现 RESTful API，包括 CRUD 操作。

## 运行示例

```bash
go run main.go
```

## 测试

```bash
# 获取用户列表
curl http://localhost:8080/api/v1/users

# 获取单个用户
curl http://localhost:8080/api/v1/users/1

# 创建用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"David","email":"david@example.com"}'

# 更新用户
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.updated@example.com"}'

# 删除用户
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 知识点

### RESTful 设计原则

- 使用 HTTP 方法表示操作：GET、POST、PUT、DELETE
- 使用名词表示资源：/users、/products
- 使用状态码表示结果：200、201、404、500

### API 设计

- 版本控制：/api/v1/
- 统一响应格式
- 错误处理
- 分页支持

## 练习

1. 添加分页功能
2. 实现搜索和过滤
3. 添加 API 文档（Swagger）

