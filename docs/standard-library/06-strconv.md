---
title: 字符串转换 (strconv)
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["strconv", "字符串", "转换", "格式化"]
---

# 字符串转换 (strconv)

`strconv` 包提供了字符串与基本数据类型之间的转换功能。

## 📋 学习目标

- [ ] 掌握字符串与数字的转换
- [ ] 理解不同进制的转换
- [ ] 学会布尔值转换
- [ ] 了解引号处理
- [ ] 掌握转换的最佳实践

## 🔢 数字转换

### 字符串转整数

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转 int
	i, err := strconv.Atoi("123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("整数: %d\n", i)
	
	// 字符串转 int64（指定进制）
	i64, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("int64: %d\n", i64)
	
	// 不同进制
	hex, _ := strconv.ParseInt("FF", 16, 64)
	oct, _ := strconv.ParseInt("777", 8, 64)
	bin, _ := strconv.ParseInt("1010", 2, 64)
	
	fmt.Printf("十六进制 FF: %d\n", hex)
	fmt.Printf("八进制 777: %d\n", oct)
	fmt.Printf("二进制 1010: %d\n", bin)
}
```

### 整数转字符串

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int 转字符串
	s := strconv.Itoa(123)
	fmt.Printf("字符串: %s\n", s)
	
	// int64 转字符串（指定进制）
	s2 := strconv.FormatInt(255, 16) // 十六进制
	fmt.Printf("255 的十六进制: %s\n", s2)
	
	s3 := strconv.FormatInt(255, 2) // 二进制
	fmt.Printf("255 的二进制: %s\n", s3)
}
```

### 字符串转浮点数

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转 float64
	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("浮点数: %f\n", f)
	
	// 字符串转 float32
	f32, err := strconv.ParseFloat("3.14", 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("float32: %f\n", float32(f32))
}
```

### 浮点数转字符串

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// float64 转字符串
	s := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("保留2位小数: %s\n", s) // 3.14
	
	// 格式说明符
	// 'f': 普通格式
	// 'e': 科学计数法
	// 'g': 自动选择最短格式
	
	s2 := strconv.FormatFloat(1234.567, 'e', 2, 64)
	fmt.Printf("科学计数法: %s\n", s2) // 1.23e+03
}
```

## ✅ 布尔值转换

### 字符串转布尔值

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转布尔值
	b1, err := strconv.ParseBool("true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true: %t\n", b1)
	
	b2, _ := strconv.ParseBool("false")
	fmt.Printf("false: %t\n", b2)
	
	b3, _ := strconv.ParseBool("1") // "1", "t", "T", "true", "TRUE" 都是 true
	fmt.Printf("1: %t\n", b3)
	
	b4, _ := strconv.ParseBool("0") // "0", "f", "F", "false", "FALSE" 都是 false
	fmt.Printf("0: %t\n", b4)
}
```

### 布尔值转字符串

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := strconv.FormatBool(true)
	fmt.Printf("true: %s\n", s1) // "true"
	
	s2 := strconv.FormatBool(false)
	fmt.Printf("false: %s\n", s2) // "false"
}
```

## 🔤 引号处理

### 添加和移除引号

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 添加引号
	quoted := strconv.Quote("Hello, World!")
	fmt.Printf("带引号: %s\n", quoted) // "Hello, World!"
	
	// 移除引号
	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("移除引号: %s\n", unquoted) // Hello, World!
	
	// ASCII 引号
	asciiQuoted := strconv.QuoteToASCII("你好")
	fmt.Printf("ASCII 引号: %s\n", asciiQuoted) // "\u4f60\u597d"
}
```

## 🎯 高级用法

### 无符号整数转换

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转 uint64
	u, err := strconv.ParseUint("123", 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("uint64: %d\n", u)
	
	// uint64 转字符串
	s := strconv.FormatUint(123, 10)
	fmt.Printf("字符串: %s\n", s)
}
```

### 带单位的数字解析

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseSize(sizeStr string) (int64, error) {
	sizeStr = strings.TrimSpace(sizeStr)
	
	// 提取数字和单位
	var numStr string
	var unit string
	
	for i, r := range sizeStr {
		if r >= '0' && r <= '9' || r == '.' {
			numStr += string(r)
		} else {
			unit = strings.ToUpper(sizeStr[i:])
			break
		}
	}
	
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}
	
	multiplier := int64(1)
	switch unit {
	case "K", "KB":
		multiplier = 1024
	case "M", "MB":
		multiplier = 1024 * 1024
	case "G", "GB":
		multiplier = 1024 * 1024 * 1024
	}
	
	return int64(num) * multiplier, nil
}

func main() {
	size, _ := parseSize("1.5MB")
	fmt.Printf("1.5MB = %d 字节\n", size)
}
```

## 🏃‍♂️ 实践应用

### 配置解析

```go
package main

import (
	"fmt"
	"strconv"
)

type Config struct {
	Port     int
	Debug    bool
	Timeout  float64
	MaxConns int
}

func parseConfig(args map[string]string) (*Config, error) {
	config := &Config{}
	
	// 解析端口
	if portStr, ok := args["port"]; ok {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		config.Port = port
	}
	
	// 解析调试模式
	if debugStr, ok := args["debug"]; ok {
		debug, err := strconv.ParseBool(debugStr)
		if err != nil {
			return nil, err
		}
		config.Debug = debug
	}
	
	return config, nil
}
```

### 数字验证

```go
package main

import (
	"fmt"
	"strconv"
)

func isValidNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isValidInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	fmt.Println(isValidNumber("123.45"))  // true
	fmt.Println(isValidNumber("abc"))     // false
	fmt.Println(isValidInt("123"))        // true
	fmt.Println(isValidInt("123.45"))     // false
}
```

## ⚠️ 注意事项

### 1. 错误处理

```go
// ✅ 总是检查错误
i, err := strconv.Atoi("123")
if err != nil {
	// 处理错误
	return err
}

// ❌ 不要忽略错误
i, _ := strconv.Atoi("abc") // i = 0，但这是错误的
```

### 2. 进制转换

```go
// ParseInt 的第二个参数是进制（2-36）
i, err := strconv.ParseInt("FF", 16, 64) // 十六进制

// FormatInt 的第二个参数是进制（2-36）
s := strconv.FormatInt(255, 16) // 转换为十六进制字符串
```

### 3. 精度问题

```go
// 浮点数转换时注意精度
f, _ := strconv.ParseFloat("3.14159", 64)
s := strconv.FormatFloat(f, 'f', 2, 64) // 保留2位小数
```

## 📚 扩展阅读

- [strconv 包官方文档](https://pkg.go.dev/strconv)

## ⏭️ 下一章节

[HTTP 客户端 (net/http)](./07-net-http.md) → 学习 Go 语言的 HTTP 客户端

---

**💡 提示**: strconv 包是处理字符串和基本类型转换的基础工具，掌握它对于处理用户输入和配置文件非常重要！

