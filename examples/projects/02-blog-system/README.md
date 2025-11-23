# 博客系统示例

这是一个完整的博客系统示例，包含用户管理、文章管理、评论系统等核心功能。

## 功能特性

- ✅ 用户注册
- ✅ 文章创建、查询、列表
- ✅ 文章评论
- ✅ 阅读量统计
- ✅ 分页查询

## 运行示例

```bash
# 安装依赖
go mod tidy

# 运行程序
go run main.go
```

## 测试 API

### 1. 用户注册

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","email":"alice@example.com","password":"123456"}'
```

### 2. 创建文章

```bash
curl -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -d '{
    "title":"Go 语言学习笔记",
    "content":"这是一篇关于 Go 语言的文章..."
  }'
```

### 3. 获取文章列表

```bash
curl http://localhost:8080/api/articles?page=1&page_size=10
```

### 4. 获取文章详情

```bash
curl http://localhost:8080/api/articles/1
```

### 5. 添加评论

```bash
curl -X POST http://localhost:8080/api/articles/1/comments \
  -H "Content-Type: application/json" \
  -d '{"content":"这是一条评论"}'
```

### 6. 获取文章评论

```bash
curl http://localhost:8080/api/articles/1/comments
```

## 项目结构

```
02-blog-system/
├── main.go          # 主程序
├── go.mod           # 依赖管理
└── README.md        # 说明文档
```

## 扩展功能

可以在此基础上添加：
- JWT 认证
- 文件上传
- 文章搜索
- 标签系统
- 分类管理

