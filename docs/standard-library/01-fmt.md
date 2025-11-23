---
title: 格式化输出 (fmt)
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["fmt", "格式化", "输出", "输入"]
---

# 格式化输出 (fmt)

`fmt` 包提供了格式化 I/O 功能，是 Go 语言中最常用的包之一。

## 📋 学习目标

- [ ] 掌握 fmt 包的基本用法
- [ ] 理解格式化占位符
- [ ] 学会格式化输出和输入
- [ ] 了解 fmt 包的高级功能

## 🎯 基本输出

### Print 系列函数

```go
package main

import "fmt"

func main() {
	// Print: 直接输出，不换行
	fmt.Print("Hello")
	fmt.Print("World")
	// 输出: HelloWorld
	
	// Println: 输出后换行
	fmt.Println("Hello")
	fmt.Println("World")
	// 输出:
	// Hello
	// World
	
	// Printf: 格式化输出
	fmt.Printf("姓名: %s, 年龄: %d\n", "张三", 30)
	// 输出: 姓名: 张三, 年龄: 30
}
```

## 🔤 格式化占位符

### 通用占位符

```go
%v    // 默认格式
%+v   // 打印结构体时包含字段名
%#v   // Go 语法表示
%T    // 类型
%%    // 字面量 %
```

### 布尔值

```go
%t    // true 或 false
```

### 整数

```go
%b    // 二进制
%o    // 八进制
%d    // 十进制
%x    // 十六进制（小写）
%X    // 十六进制（大写）
%c    // Unicode 字符
%U    // Unicode 格式
```

### 浮点数

```go
%e    // 科学计数法（小写）
%E    // 科学计数法（大写）
%f    // 十进制表示
%g    // 自动选择 %e 或 %f
%G    // 自动选择 %E 或 %f
```

### 字符串和字节切片

```go
%s    // 字符串
%q    // 带引号的字符串
%x    // 十六进制（小写）
%X    // 十六进制（大写）
```

### 指针

```go
%p    // 指针地址（十六进制）
```

## 💻 代码示例

### 基本格式化

```go
package main

import "fmt"

func main() {
	name := "张三"
	age := 30
	height := 175.5
	
	// 基本格式化
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年龄: %d\n", age)
	fmt.Printf("身高: %.1f cm\n", height)
	
	// 使用 %v（默认格式）
	fmt.Printf("信息: %v\n", name)
	fmt.Printf("信息: %+v\n", name)
	fmt.Printf("信息: %#v\n", name)
	fmt.Printf("类型: %T\n", name)
}
```

### 宽度和精度

```go
package main

import "fmt"

func main() {
	num := 123
	
	// 宽度
	fmt.Printf("|%5d|\n", num)    // |  123|
	fmt.Printf("|%-5d|\n", num)   // |123  |
	fmt.Printf("|%05d|\n", num)   // |00123|
	
	// 浮点数精度
	pi := 3.14159
	fmt.Printf("%.2f\n", pi)     // 3.14
	fmt.Printf("%.4f\n", pi)     // 3.1416
}
```

### 结构体格式化

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "张三", Age: 30}
	
	fmt.Printf("%v\n", p)   // {张三 30}
	fmt.Printf("%+v\n", p)  // {Name:张三 Age:30}
	fmt.Printf("%#v\n", p)  // main.Person{Name:"张三", Age:30}
}
```

## 📥 输入函数

### Scan 系列函数

```go
package main

import "fmt"

func main() {
	var name string
	var age int
	
	// Scan: 从标准输入读取
	fmt.Print("请输入姓名和年龄: ")
	fmt.Scan(&name, &age)
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
	
	// Scanf: 格式化输入
	fmt.Print("请输入姓名: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("姓名: %s\n", name)
	
	// Scanln: 读取一行
	fmt.Print("请输入一行文本: ")
	var line string
	fmt.Scanln(&line)
	fmt.Printf("输入: %s\n", line)
}
```

## 🔧 高级功能

### Sprintf - 格式化字符串

```go
package main

import "fmt"

func main() {
	name := "张三"
	age := 30
	
	// 格式化但不输出，返回字符串
	message := fmt.Sprintf("姓名: %s, 年龄: %d", name, age)
	fmt.Println(message)
}
```

### Fprintf - 格式化写入

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 写入到文件
	file, _ := os.Create("output.txt")
	defer file.Close()
	
	fmt.Fprintf(file, "姓名: %s, 年龄: %d\n", "张三", 30)
}
```

### Errorf - 格式化错误

```go
package main

import (
	"fmt"
	"errors"
)

func divide(a, b float64) error {
	if b == 0 {
		return fmt.Errorf("除数不能为零: %.2f / %.2f", a, b)
	}
	return nil
}

func main() {
	err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}
```

## 🏃‍♂️ 实践练习

### 练习 1: 格式化输出表格

使用 fmt 包输出一个格式化的表格。

### 练习 2: 输入验证

编写一个程序，从用户输入读取数据并进行验证。

### 练习 3: 日志格式化

使用 fmt 包格式化日志输出。

## 🤔 思考题

1. Print、Println、Printf 有什么区别？
2. 什么时候使用 Sprintf？
3. 如何控制浮点数的精度？
4. 格式化占位符的宽度和精度如何设置？

## 📚 扩展阅读

- [fmt 包官方文档](https://pkg.go.dev/fmt)
- [格式化占位符详解](https://golang.org/pkg/fmt/)

## ⏭️ 下一章节

[时间处理 (time)](./02-time.md) → 学习 Go 语言的时间操作

---

**💡 提示**: fmt 包是 Go 语言中最常用的包之一，熟练掌握格式化占位符可以让你更高效地处理输出！

