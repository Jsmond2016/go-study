# Beego 框架示例

本示例展示了 Beego 框架的基本用法，包括路由和控制器。

## 运行示例

```bash
go mod tidy
go run main.go
```

## 测试

```bash
# 访问首页
curl http://localhost:8080/

# 获取用户列表
curl http://localhost:8080/api/users

# 创建用户
curl -X POST http://localhost:8080/api/users \
  -d "name=Alice&email=alice@example.com"
```

## 知识点

### Beego 框架

- Beego 是一个完整的 MVC 框架
- 使用控制器处理请求
- 支持 RESTful 路由

### 控制器

- 继承 `web.Controller`
- 实现 HTTP 方法对应的函数（Get、Post、Put、Delete）
- 使用 `c.Data` 传递数据
- 使用 `c.ServeJSON()` 返回 JSON

### 路由

- `web.Router()`：注册路由
- 支持路径参数和查询参数

## 注意事项

Beego 是一个完整的框架，包含更多功能：
- ORM
- 模板引擎
- Session 管理
- 日志系统

## 练习

1. 实现完整的 CRUD 操作
2. 添加数据库集成
3. 实现用户认证

