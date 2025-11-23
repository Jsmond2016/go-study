---
title: 字符串操作 (strings)
difficulty: beginner
duration: "3-4小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["strings", "字符串", "文本处理", "查找", "替换"]
---

# 字符串操作 (strings)

`strings` 包提供了丰富的字符串操作函数，是处理文本数据的基础工具。

## 📋 学习目标

- [ ] 掌握字符串的查找和匹配
- [ ] 学会字符串的替换和修改
- [ ] 理解字符串的分割和拼接
- [ ] 掌握字符串的修剪和格式化
- [ ] 了解字符串构建器的高效使用

## 🔍 字符串查找

### 基本查找

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World"

	// 检查是否包含子串
	contains := strings.Contains(str, "World")
	fmt.Printf("包含 'World': %t\n", contains) // true

	// 查找子串位置
	index := strings.Index(str, "World")
	fmt.Printf("'World' 的位置: %d\n", index) // 7

	// 查找最后出现的位置
	lastIndex := strings.LastIndex(str, "o")
	fmt.Printf("'o' 最后出现的位置: %d\n", lastIndex) // 8

	// 查找任意字符
	anyIndex := strings.IndexAny(str, "aeiou")
	fmt.Printf("第一个元音字母位置: %d\n", anyIndex) // 1 (e)

	// 统计出现次数
	count := strings.Count(str, "l")
	fmt.Printf("'l' 出现次数: %d\n", count) // 3
}
```

### 前缀和后缀

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World"

	// 检查前缀
	hasPrefix := strings.HasPrefix(str, "Hello")
	fmt.Printf("以 'Hello' 开头: %t\n", hasPrefix) // true

	// 检查后缀
	hasSuffix := strings.HasSuffix(str, "World")
	fmt.Printf("以 'World' 结尾: %t\n", hasSuffix) // true
}
```

## 🔄 字符串替换

### 基本替换

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World, World"

	// 替换第一个匹配
	newStr1 := strings.Replace(str, "World", "Go", 1)
	fmt.Printf("替换第一个: %s\n", newStr1) // Hello, Go, World

	// 替换所有匹配
	newStr2 := strings.ReplaceAll(str, "World", "Go")
	fmt.Printf("替换所有: %s\n", newStr2) // Hello, Go, Go

	// 替换指定次数
	newStr3 := strings.Replace(str, "World", "Go", 2)
	fmt.Printf("替换2次: %s\n", newStr3) // Hello, Go, Go
}
```

### 大小写转换

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World"

	// 转大写
	upper := strings.ToUpper(str)
	fmt.Printf("大写: %s\n", upper) // HELLO, WORLD

	// 转小写
	lower := strings.ToLower(str)
	fmt.Printf("小写: %s\n", lower) // hello, world

	// 首字母大写（已废弃，建议使用其他方法）
	title := strings.Title(str)
	fmt.Printf("标题格式: %s\n", title) // Hello, World

	// 每个单词首字母大写
	// 注意：Go 1.18+ 推荐使用 cases 包
}
```

## ✂️ 字符串修剪

### 去除空白字符

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Hello, World  "

	// 去除两端空白
	trimmed := strings.TrimSpace(str)
	fmt.Printf("去除空白: '%s'\n", trimmed) // 'Hello, World'

	// 去除左侧空白
	trimLeft := strings.TrimLeft(str, " ")
	fmt.Printf("去除左侧空白: '%s'\n", trimLeft)

	// 去除右侧空白
	trimRight := strings.TrimRight(str, " ")
	fmt.Printf("去除右侧空白: '%s'\n", trimRight)

	// 去除指定字符
	trimmed2 := strings.Trim("!!!Hello!!!", "!")
	fmt.Printf("去除 '!': %s\n", trimmed2) // Hello
}
```

### 去除前缀和后缀

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World"

	// 去除前缀
	withoutPrefix := strings.TrimPrefix(str, "Hello, ")
	fmt.Printf("去除前缀: %s\n", withoutPrefix) // World

	// 去除后缀
	withoutSuffix := strings.TrimSuffix(str, ", World")
	fmt.Printf("去除后缀: %s\n", withoutSuffix) // Hello
}
```

## 🔪 字符串分割

### 基本分割

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "apple,banana,orange,grape"

	// 按分隔符分割
	parts := strings.Split(str, ",")
	fmt.Printf("分割结果: %v\n", parts) // [apple banana orange grape]

	// 限制分割次数
	parts2 := strings.SplitN(str, ",", 2)
	fmt.Printf("分割2次: %v\n", parts2) // [apple banana,orange,grape]

	// 按空白字符分割
	words := strings.Fields("  hello   world  go  ")
	fmt.Printf("按空白分割: %v\n", words) // [hello world go]
}
```

### 高级分割

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a,b,c,d"

	// 分割后保留分隔符
	parts := strings.SplitAfter(str, ",")
	fmt.Printf("保留分隔符: %v\n", parts) // [a, b, c, d]

	// 限制分割次数并保留分隔符
	parts2 := strings.SplitAfterN(str, ",", 2)
	fmt.Printf("保留分隔符分割2次: %v\n", parts2) // [a, b,c,d]
}
```

## 🔗 字符串拼接

### 基本拼接

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 使用 + 运算符（少量拼接）
	str1 := "Hello" + " " + "World"
	fmt.Printf("拼接结果: %s\n", str1)

	// 使用 Join（固定分隔符）
	parts := []string{"apple", "banana", "orange"}
	str2 := strings.Join(parts, ", ")
	fmt.Printf("Join 结果: %s\n", str2) // apple, banana, orange
}
```

### 使用 Builder（高效拼接）

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	// 追加字符串
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")

	// 追加字节
	builder.WriteByte('!')

	// 追加 rune
	builder.WriteRune('🎉')

	// 获取结果
	result := builder.String()
	fmt.Printf("Builder 结果: %s\n", result) // Hello World!🎉

	// 预分配容量（性能优化）
	builder2 := strings.Builder{}
	builder2.Grow(100) // 预分配100字节
	builder2.WriteString("Pre-allocated")
	fmt.Printf("预分配结果: %s\n", builder2.String())
}
```

## 🔤 字符串比较

### 基本比较

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	str2 := "hello"

	// 区分大小写比较
	equal := str1 == str2
	fmt.Printf("相等: %t\n", equal) // false

	// 忽略大小写比较
	equalFold := strings.EqualFold(str1, str2)
	fmt.Printf("忽略大小写相等: %t\n", equalFold) // true

	// 比较大小（字典序）
	compare := strings.Compare(str1, str2)
	fmt.Printf("比较结果: %d\n", compare) // -1 (str1 < str2)
}
```

## 📊 字符串映射和转换

### Map 函数

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "Hello World 123"

	// 将字符串中的每个字符映射
	mapped := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, str)
	fmt.Printf("映射结果: %s\n", mapped) // HELLO WORLD 123
}
```

### Repeat 函数

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 重复字符串
	repeated := strings.Repeat("Go ", 3)
	fmt.Printf("重复结果: %s\n", repeated) // Go Go Go
}
```

## 🎯 实用工具函数

### 字符串转换

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串转切片
	str := "Hello"
	bytes := []byte(str)
	runes := []rune(str)

	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("Rune 切片: %v\n", runes)

	// 切片转字符串
	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("字节转字符串: %s\n", str2)
	fmt.Printf("Rune 转字符串: %s\n", str3)
}
```

### 字符串验证

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "Hello123"

	// 检查是否只包含字母
	hasOnlyLetters := true
	for _, r := range str {
		if !unicode.IsLetter(r) {
			hasOnlyLetters = false
			break
		}
	}
	fmt.Printf("只包含字母: %t\n", hasOnlyLetters) // false

	// 检查是否为空
	isEmpty := strings.TrimSpace(str) == ""
	fmt.Printf("是否为空: %t\n", isEmpty) // false
}
```

## 🏃‍♂️ 实践应用

### 文本处理工具

```go
package main

import (
	"fmt"
	"strings"
)

// 单词计数
func wordCount(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	count := make(map[string]int)
	for _, word := range words {
		word = strings.Trim(word, ".,!?;:")
		count[word]++
	}
	return count
}

// 清理文本
func cleanText(text string) string {
	// 去除多余空白
	text = strings.TrimSpace(text)
	// 替换多个空白为单个
	words := strings.Fields(text)
	return strings.Join(words, " ")
}

func main() {
	text := "  Hello   World  Hello  Go  "

	// 单词计数
	counts := wordCount(text)
	fmt.Printf("单词计数: %v\n", counts)

	// 清理文本
	cleaned := cleanText(text)
	fmt.Printf("清理后: '%s'\n", cleaned)
}
```

### URL 路径处理

```go
package main

import (
	"fmt"
	"strings"
)

func normalizePath(path string) string {
	// 去除前后斜杠
	path = strings.Trim(path, "/")
	// 确保以 / 开头
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	// 去除重复斜杠
	parts := strings.Split(path, "/")
	var cleanParts []string
	for _, part := range parts {
		if part != "" {
			cleanParts = append(cleanParts, part)
		}
	}
	return "/" + strings.Join(cleanParts, "/")
}

func main() {
	paths := []string{
		"//api//users//",
		"/api/users",
		"api/users/",
	}

	for _, path := range paths {
		normalized := normalizePath(path)
		fmt.Printf("原始: %s -> 规范化: %s\n", path, normalized)
	}
}
```

## ⚠️ 注意事项

### 1. 性能考虑

```go
// ❌ 避免：大量字符串拼接使用 +
var result string
for i := 0; i < 1000; i++ {
	result += "a" // 每次都会创建新字符串
}

// ✅ 正确：使用 Builder
var builder strings.Builder
for i := 0; i < 1000; i++ {
	builder.WriteString("a")
}
result := builder.String()
```

### 2. Unicode 处理

```go
// 注意：strings 包按字节操作，处理 Unicode 时要注意
str := "你好"
fmt.Println(len(str)) // 6 (字节数，不是字符数)

// 处理 Unicode 字符时使用 []rune
runes := []rune(str)
fmt.Println(len(runes)) // 2 (字符数)
```

### 3. 字符串不可变性

```go
// Go 中字符串是不可变的
str := "Hello"
// str[0] = 'h' // 编译错误

// 需要修改时创建新字符串
newStr := "h" + str[1:]
```

## 📚 扩展阅读

- [strings 包官方文档](https://pkg.go.dev/strings)
- [Effective Go - 字符串](https://golang.org/doc/effective_go.html#strings)

## ⏭️ 下一章节

[数学运算 (math)](./12-math.md) → 学习使用数学函数进行数值计算

---

**💡 提示**: strings 包是处理文本数据的基础工具，掌握它对于处理用户输入、配置文件、日志等文本数据非常重要！

