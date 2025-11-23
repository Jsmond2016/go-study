# 字符串转换示例

本示例展示了 Go 语言 strconv 包的使用，包括字符串与数字、布尔值的转换，以及不同进制的转换。

## 运行示例

```bash
go run main.go
```

## 知识点

### 数字转换

- `strconv.Atoi()`：字符串转 int
- `strconv.Itoa()`：int 转字符串
- `strconv.ParseInt()`：字符串转 int64（指定进制）
- `strconv.FormatInt()`：int64 转字符串（指定进制）

### 浮点数转换

- `strconv.ParseFloat()`：字符串转浮点数
- `strconv.FormatFloat()`：浮点数转字符串

### 布尔值转换

- `strconv.ParseBool()`：字符串转布尔值
- `strconv.FormatBool()`：布尔值转字符串

### 进制转换

- 支持 2-36 进制
- 常用：2（二进制）、8（八进制）、10（十进制）、16（十六进制）

## 练习

1. 实现一个进制转换工具，支持不同进制之间的转换
2. 创建一个函数，安全地将字符串转换为数字，处理错误情况
3. 实现一个配置解析器，将字符串配置转换为相应的类型

