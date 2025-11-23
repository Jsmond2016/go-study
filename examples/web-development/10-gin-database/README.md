# Gin 数据库集成示例

本示例展示了 Gin 框架与 GORM 数据库的集成，包括 CRUD 操作。

## 运行示例

```bash
go mod tidy
go run main.go
```

## 测试

```bash
# 创建用户
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'

# 获取用户列表
curl http://localhost:8080/users

# 获取单个用户
curl http://localhost:8080/users/1

# 更新用户
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.updated@example.com"}'

# 删除用户
curl -X DELETE http://localhost:8080/users/1
```

## 知识点

### GORM 基础

- `gorm.Open()`：连接数据库
- `db.AutoMigrate()`：自动迁移表结构
- `db.Create()`：创建记录
- `db.Find()`：查询记录
- `db.First()`：查询单条记录
- `db.Save()`：更新记录
- `db.Delete()`：删除记录

### 数据库集成

- 使用 GORM 简化数据库操作
- 模型定义和自动迁移
- 错误处理

## 练习

1. 添加数据库连接池配置
2. 实现事务处理
3. 添加查询条件和分页

