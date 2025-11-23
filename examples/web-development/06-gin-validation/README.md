# Gin 数据验证示例

本示例展示了 Gin 框架的数据验证功能，包括结构体标签验证和错误处理。

## 运行示例

```bash
go run main.go
```

## 测试

```bash
# 创建用户（成功）
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com","age":30,"password":"123456"}'

# 创建用户（验证失败）
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name":"A","email":"invalid","age":200}'
```

## 知识点

### 验证标签

- `required`：必填字段
- `min` / `max`：长度限制
- `gte` / `lte`：数值范围
- `email`：邮箱格式
- `url`：URL 格式
- `oneof`：枚举值

### 绑定方式

- `ShouldBindJSON()`：JSON 绑定
- `ShouldBindQuery()`：查询参数绑定
- `ShouldBindUri()`：路径参数绑定

## 练习

1. 实现一个完整的用户注册验证
2. 创建自定义验证规则
3. 实现验证错误的友好提示

