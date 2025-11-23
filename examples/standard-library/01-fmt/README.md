# fmt 包示例

fmt 包是 Go 语言中最重要的基础包之一，提供了格式化 I/O 功能。本示例展示了 fmt 包的各种用法和高级特性。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 基本输出函数

```go
// Print - 不换行输出
fmt.Print("Hello, ")
fmt.Print("Go!\n")

// Println - 自动换行，参数间添加空格
fmt.Println("Hello,", "Go!")

// Printf - 格式化输出
fmt.Printf("姓名: %s, 年龄: %d\n", "张三", 25)
```

### 2. 输入函数

```go
// Scan - 从标准输入读取
var name string
var age int
fmt.Scan(&name, &age)

// Scanln - 读取到换行符
fmt.Scanln(&name, &age)

// Scanf - 格式化输入
fmt.Scanf("%s,%d", &name, &age)
```

### 3. 格式化动词详解

#### 通用动词
- `%v` - 默认格式
- `%+v` - 结构体显示字段名
- `%#v` - Go 语法表示
- `%T` - 类型信息

#### 布尔值
- `%t` - true 或 false

#### 整数
- `%b` - 二进制
- `%c` - 对应的 Unicode 字符
- `%d` - 十进制
- `%o` - 八进制
- `%x`, `%X` - 十六进制
- `%U` - Unicode 格式

#### 浮点数
- `%b` - 科学记数法（二进制指数）
- `%e`, `%E` - 科学记数法
- `%f`, `%F` - 小数形式
- `%g`, `%G` - 最紧凑格式

#### 字符串和字节
- `%s` - 字符串
- `%q` - 带引号的字符串
- `%x`, `%X` - 十六进制字节值

#### 指针
- `%p` - 指针地址

### 4. 宽度和精度控制

```go
fmt.Printf("|%6d|", 42)        // 宽度为6
fmt.Printf("|%-6d|", 42)       // 左对齐
fmt.Printf("|%06d|", 42)       // 用0填充
fmt.Printf("|%6.2f|", 3.14)    // 宽度6，精度2
fmt.Printf("|%.3s|", "Hello")  // 字符串截断
```

### 5. 文件操作函数

```go
// Fprint 系列 - 输出到指定文件
fmt.Fprintln(os.Stdout, "Hello")

// Sprint 系列 - 返回字符串
str := fmt.Sprintf("结果: %d", 42)

// Errorf - 创建错误
err := fmt.Errorf("错误代码: %d", 404)
```

## 实际应用场景

### 1. 格式化表格

```go
fmt.Printf("|%-10s|%-20s|\n", "姓名", "邮箱")
fmt.Printf("+----------+--------------------+\n")
for _, user := range users {
    fmt.Printf("|%-10s|%-20s|\n", user.Name, user.Email)
}
```

### 2. 进度条显示

```go
bar := strings.Repeat("=", completed) + strings.Repeat(" ", total-completed)
fmt.Printf("\r[%s] %d%%", bar, percentage)
```

### 3. 日志格式化

```go
fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
```

### 4. 配置显示

```go
fmt.Printf("服务器: %s:%d\n", config.Host, config.Port)
fmt.Printf("数据库: %s@%s:%d/%s\n", config.User, config.Host, config.Port, config.DB)
```

## 性能考虑

1. **避免频繁格式化**：在循环中预分配字符串
2. **使用缓冲区**：大量输出时使用 `bufio.Writer`
3. **简化格式**：简单场景使用 `strconv` 包

## 最佳实践

1. **错误信息**：使用 `fmt.Errorf` 创建描述性错误
2. **日志格式**：保持一致的日志格式
3. **用户输出**：使用友好的格式和语言
4. **调试信息**：使用 `%+v` 显示详细结构

## 常见陷阱

1. **格式字符串不匹配**：参数数量和类型要匹配格式动词
2. **换行符处理**：注意 `Print` 和 `Println` 的区别
3. **Unicode 处理**：正确处理中文字符的宽度
4. **缓冲问题**：及时刷新输出缓冲区

## 与其他包的配合

```go
// 与 strings 包配合
result := fmt.Sprintf("%s %s", strings.Title(name), strings.ToUpper(status))

// 与 strconv 包配合
number := fmt.Sprintf("价格: ¥%.2f", price)

// 与 log 包配合
log.Printf("处理完成: %v", result)
```

## 实用技巧

### 1. 调试输出
```go
fmt.Printf("调试: %+v\n", complexStruct)
```

### 2. JSON 风格格式化
```go
fmt.Printf("{\n  \"name\": %q,\n  \"age\": %d\n}\n", name, age)
```

### 3. 条件格式化
```go
status := "OK"
if !success {
    status = "FAILED"
}
fmt.Printf("操作状态: %s\n", status)
```

## 练习

1. 创建一个格式化的报表生成器
2. 实现一个彩色控制台输出工具
3. 创建一个配置文件格式化器
4. 实现一个简单的模板系统
5. 创建一个进度条组件