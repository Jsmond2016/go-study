# golang 中  strconv 常用的方法有哪些


我来介绍一下 Go 语言中 strconv 包的常用方法。我将创建一个新的文档来详细说明。

```markdown:/Users/huangjin/Desktop/github/golang-pro/go-study/docs/record/base/go-strconv.md
# Go 语言 strconv 包常用方法

## 一、字符串与整数转换

```go
import "strconv"

// 1. string 转 int
num1, err := strconv.Atoi("42")              // 字符串转 int
num2, err := strconv.ParseInt("42", 10, 64)  // 字符串转 int64，第二个参数是进制，第三个参数是位数

// 2. int 转 string
str1 := strconv.Itoa(42)                     // int 转字符串
str2 := strconv.FormatInt(42, 10)            // int64 转字符串，第二个参数是进制
```

## 二、字符串与浮点数转换

```go
// 1. string 转 float
f1, err := strconv.ParseFloat("3.14", 64)    // 转换为 float64
f2, err := strconv.ParseFloat("3.14", 32)    // 转换为 float32

// 2. float 转 string
str1 := strconv.FormatFloat(3.14, 'f', 2, 64)  // 'f' 格式，保留2位小数
str2 := strconv.FormatFloat(3.14, 'e', 2, 64)  // 科学计数法
str3 := strconv.FormatFloat(3.14, 'g', 2, 64)  // 自动选择格式
```

## 三、字符串与布尔值转换

```go
// 1. string 转 bool
b1, err := strconv.ParseBool("true")   // 转换 "true", "1", "t", "T", "TRUE"
b2, err := strconv.ParseBool("false")  // 转换 "false", "0", "f", "F", "FALSE"

// 2. bool 转 string
str1 := strconv.FormatBool(true)       // "true"
str2 := strconv.FormatBool(false)      // "false"
```

## 四、进制转换

```go
// 1. 其他进制字符串转换为十进制整数
bin, err := strconv.ParseInt("1010", 2, 64)    // 二进制转十进制
oct, err := strconv.ParseInt("777", 8, 64)     // 八进制转十进制
hex, err := strconv.ParseInt("FF", 16, 64)     // 十六进制转十进制

// 2. 十进制转其他进制字符串
binStr := strconv.FormatInt(10, 2)             // 转二进制字符串
octStr := strconv.FormatInt(64, 8)             // 转八进制字符串
hexStr := strconv.FormatInt(255, 16)           // 转十六进制字符串
```

## 五、特殊格式化选项

```go
// 1. 浮点数格式化选项
f := 3.1415926
s1 := strconv.FormatFloat(f, 'f', -1, 64)   // 无小数位限制
s2 := strconv.FormatFloat(f, 'f', 2, 64)    // 保留2位小数
s3 := strconv.FormatFloat(f, 'e', 2, 64)    // 科学计数法
s4 := strconv.FormatFloat(f, 'g', 2, 64)    // 根据值自动选择格式

// 2. 整数格式化为指定宽度的字符串
uint := uint64(42)
s5 := strconv.FormatUint(uint, 10)          // 无符号整数转字符串
```

## 六、引号处理

```go
// 1. 添加引号
str := strconv.Quote("Hello, 世界")           // 添加双引号并转义
str2 := strconv.QuoteToASCII("Hello, 世界")  // 转换为ASCII并添加引号

// 2. 去除引号
unquoted, err := strconv.Unquote("\"Hello\"")  // 去除双引号
```

## 七、实用技巧

1. **错误处理最佳实践**：
```go
// 始终检查错误
if num, err := strconv.Atoi("42"); err == nil {
    // 使用 num
} else {
    // 处理错误
}
```

2. **性能优化**：
```go
// 对于已知范围的整数，使用 ParseInt 可以避免溢出
smallNum, err := strconv.ParseInt("256", 10, 8)  // 确保数值在 int8 范围内
```

## 注意事项

1. 进行数值转换时必须注意处理错误返回值
2. 浮点数转换要注意精度问题
3. ParseInt/ParseUint 的位数参数(bitSize)要根据实际需求选择
4. 在处理用户输入时，要考虑输入验证和安全性
5. 对于大量的字符串转换操作，注意性能优化


此文档涵盖了 strconv 包中最常用的方法。建议根据实际场景选择合适的方法，并始终注意错误处理。