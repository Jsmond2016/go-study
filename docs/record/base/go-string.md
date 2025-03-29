# golang 中，字符串常用方法有哪些?

Go 语言中常用的字符串操作方法

## 一、字符串基本操作

```go
str := "Hello, 世界"

// 1. 获取字符串长度
length := len(str)                // 字节长度
runeCount := utf8.RuneCountInString(str)  // Unicode 字符数

// 2. 字符串拼接
s1 := "Hello" + " World"         // 使用 + 运算符
s2 := strings.Join([]string{"Hello", "World"}, " ")  // 使用 Join
s3 := fmt.Sprintf("%s %s", "Hello", "World")        // 使用 Sprintf
```

## 二、字符串查找和替换

```go
str := "Hello, World"

// 1. 查找子串
contains := strings.Contains(str, "World")     // 是否包含
index := strings.Index(str, "World")          // 查找位置
lastIndex := strings.LastIndex(str, "o")      // 最后出现位置

// 2. 替换
newStr := strings.Replace(str, "World", "Go", 1)  // 替换一次
newStr2 := strings.ReplaceAll(str, "l", "L")     // 替换所有
```

## 三、字符串转换

```go
// 1. 大小写转换
upper := strings.ToUpper("hello")    // "HELLO"
lower := strings.ToLower("HELLO")    // "hello"
title := strings.Title("hello world") // "Hello World"

// 2. 类型转换
str := strconv.Itoa(123)            // int 转字符串
num, _ := strconv.Atoi("123")       // 字符串转 int
f, _ := strconv.ParseFloat("3.14", 64)  // 字符串转 float64
```

## 四、字符串修剪和分割

```go
// 1. 去除空白字符
trimmed := strings.TrimSpace(" hello ")  // "hello"
trimLeft := strings.TrimLeft("!!!hello", "!")   // "hello"
trimRight := strings.TrimRight("hello!!!", "!") // "hello"

// 2. 分割字符串
parts := strings.Split("a,b,c", ",")    // ["a", "b", "c"]
lines := strings.Split("a\nb\nc", "\n") // ["a", "b", "c"]
```

## 五、字符串比较和前缀后缀

```go
// 1. 字符串比较
equal := strings.EqualFold("Hello", "hello")  // 忽略大小写比较

// 2. 前缀后缀判断
hasPrefix := strings.HasPrefix("Hello", "He")  // true
hasSuffix := strings.HasSuffix("Hello", "lo")  // true
```

## 六、字符串构建器（适用于频繁拼接）

```go
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
result := builder.String()  // "Hello World"
```

## 七、实用技巧

1. **高效拼接**：
   - 少量拼接用 `+` 运算符
   - 大量拼接用 `strings.Builder`
   - 固定分隔符用 `strings.Join`

2. **性能考虑**：
   - `strings.Builder` 比 `+` 和 `fmt.Sprintf` 性能更好
   - 预知容量时可用 `builder.Grow(n)` 预分配内存

3. **编码处理**：
   - 处理 Unicode 字符时使用 `[]rune` 转换
   - 注意区分字节长度和字符长度

4. **正则表达式**：
```go
reg := regexp.MustCompile(`\d+`)
matches := reg.FindAllString("abc123def456", -1)  // ["123", "456"]
```

## 注意事项
1. 字符串在 Go 中是不可变的，每次修改都会创建新的字符串
2. 使用 `strings.Builder` 进行字符串拼接比 `+` 操作符更高效
3. 处理 Unicode 字符时需要特别注意编码问题
4. 大小写转换要考虑语言环境（locale）的影响
```

此文涵盖了 Go 语言中最常用的字符串操作方法。建议根据实际场景选择合适的方法，特别是在性能要求较高的场景下。