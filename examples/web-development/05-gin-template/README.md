# Gin 模板示例

本示例展示了 Gin 框架的模板功能，包括 HTML 模板渲染和静态文件服务。

## 运行示例

```bash
# 创建模板目录和文件
mkdir -p templates static

# 创建模板文件 templates/index.html
# 创建模板文件 templates/users.html

go run main.go
```

## 知识点

### 模板加载

- `r.LoadHTMLGlob()`：加载模板文件
- `r.LoadHTMLFiles()`：加载指定模板文件

### 模板渲染

- `c.HTML()`：渲染 HTML 模板
- 传递数据到模板

### 静态文件

- `r.Static()`：提供静态文件服务
- `r.StaticFS()`：提供静态文件系统服务

## 模板文件示例

### templates/index.html
```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
</head>
<body>
    <h1>欢迎使用 {{.name}}</h1>
</body>
</html>
```

## 练习

1. 创建一个完整的博客系统模板
2. 实现模板继承和包含
3. 创建一个响应式布局模板

