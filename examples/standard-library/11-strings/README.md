# 字符串操作示例

本示例展示了 Go 语言 strings 包的使用，包括字符串查找、替换、分割、拼接等操作。

## 运行示例

```bash
go run main.go
```

## 知识点

### 字符串查找

- `strings.Contains()`：检查是否包含子串
- `strings.Index()`：查找子串位置
- `strings.Count()`：统计出现次数

### 字符串操作

- `strings.Replace()`：替换字符串
- `strings.ToUpper()` / `ToLower()`：大小写转换
- `strings.TrimSpace()`：去除空白字符
- `strings.Split()`：分割字符串
- `strings.Join()`：拼接字符串

### 字符串构建器

- `strings.Builder`：高效的字符串构建
- `WriteString()`：追加字符串
- `String()`：获取结果

## 练习

1. 实现一个文本处理工具，支持查找和替换
2. 创建一个 URL 路径规范化函数
3. 实现一个简单的模板引擎，使用字符串操作

