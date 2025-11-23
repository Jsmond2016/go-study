---
title: 变量与常量
difficulty: beginner
duration: "2-3小时"
prerequisites: []
tags: ["变量", "常量", "作用域", "初始化"]
---

# 变量与常量

变量和常量是编程的基础概念。Go 语言提供了灵活的变量声明方式和严格的常量定义机制。

## 📋 学习目标

- [ ] 掌握 Go 语言的变量声明方式
- [ ] 理解变量的初始化和零值
- [ ] 学会使用常量和枚举
- [ ] 理解变量的作用域
- [ ] 掌握多重赋值和匿名变量

## 📦 变量声明

### 基本声明语法

Go 提供了多种变量声明方式：

```go
package main

import "fmt"

func main() {
	// 方式1：var 变量名 类型
	var name string
	var age int
	var isStudent bool
	
	// 方式2：var 变量名 类型 = 值
	var city string = "北京"
	var score int = 95
	
	// 方式3：var 变量名1, 变量名2 类型
	var x, y, z int
	var a, b string
	
	// 方式4：var 变量名1, 变量名2 类型 = 值1, 值2
	var width, height int = 100, 200
	var firstName, lastName string = "张", "三"
	
	// 方式5：短变量声明（函数内）
	count := 10
	rate := 0.85
	isActive := true
	
	fmt.Printf("姓名: %s, 年龄: %d, 学生: %t\n", name, age, isStudent)
	fmt.Printf("城市: %s, 分数: %d\n", city, score)
	fmt.Printf("坐标: (%d, %d, %d)\n", x, y, z)
	fmt.Printf("字符串: %s, %s\n", a, b)
	fmt.Printf("尺寸: %d x %d\n", width, height)
	fmt.Printf("全名: %s%s\n", firstName, lastName)
	fmt.Printf("数量: %d, 比例: %.2f, 状态: %t\n", count, rate, isActive)
}
```

### 变量初始化和零值

```go
package main

import "fmt"

func main() {
	// 未显式初始化的变量会有零值
	var i int          // 0
	var f float64      // 0.0
	var b bool         // false
	var s string       // "" (空字符串)
	var p *int        // nil
	var arr [5]int     // [0 0 0 0 0]
	var slice []int    // nil
	var m map[string]int // nil
	
	fmt.Printf("int 零值: %d\n", i)
	fmt.Printf("float64 零值: %f\n", f)
	fmt.Printf("bool 零值: %t\n", b)
	fmt.Printf("string 零值: '%s' (长度: %d)\n", s, len(s))
	fmt.Printf("指针零值: %v\n", p)
	fmt.Printf("数组零值: %v\n", arr)
	fmt.Printf("切片零值: %v\n", slice)
	fmt.Printf("映射零值: %v\n", m)
	
	// 检查是否为零值
	if p == nil {
		fmt.Println("指针是 nil")
	}
	if slice == nil {
		fmt.Println("切片是 nil")
	}
}
```

## 🏷️ 常量定义

### 基本常量

```go
package main

import "fmt"

const (
	PI     = 3.141592653589793
	E      = 2.718281828459045
	Golden = 1.618033988749895
)

func main() {
	fmt.Printf("π = %.15f\n", PI)
	fmt.Printf("e = %.15f\n", E)
	fmt.Printf("黄金比例 = %.15f\n", Golden)
	
	// 常量在编译时确定
	const Version = "1.0.0"
	const MaxUsers = 1000
	const DebugMode = false
	
	fmt.Printf("版本: %s, 最大用户: %d, 调试模式: %t\n", Version, MaxUsers, DebugMode)
}
```

### 枚举常量

```go
package main

import "fmt"

// 星期枚举
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday             // 1
	Tuesday            // 2
	Wednesday          // 3
	Thursday          // 4
	Friday            // 5
	Saturday          // 6
)

// 月份枚举
type Month int

const (
	January Month = 1 + iota // 1
	February              // 2
	March                 // 3
	April                 // 4
	May                   // 5
	June                  // 6
	July                  // 7
	August                // 8
	September             // 9
	October              // 10
	November              // 11
	December             // 12
)

// 文件权限枚举
type FileMode int

const (
	ReadOnly FileMode = 1 << iota // 1 << 0 = 1
	WriteOnly                  // 1 << 1 = 2
	ReadWrite                  // 1 << 2 = 4
	Execute                   // 1 << 3 = 8
)

func main() {
	// 星期枚举使用
	today := Wednesday
	fmt.Printf("今天是星期 %d\n", today)
	
	// 月份枚举使用
	currentMonth := September
	fmt.Printf("当前是 %d 月\n", currentMonth)
	
	// 权限枚举使用
	var permissions FileMode = ReadWrite | Execute
	fmt.Printf("权限值: %d\n", permissions)
	
	// 检查权限
	if permissions&ReadWrite != 0 {
		fmt.Println("有读写权限")
	}
	if permissions&Execute != 0 {
		fmt.Println("有执行权限")
	}
}
```

### 复杂常量表达式

```go
package main

import "fmt"

const (
	// 时间常量
	Second = 1
	Minute = 60 * Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
	
	// 存储单位
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	
	// 应用配置
	AppName    = "Go学习系统"
	AppVersion = "2.1.0"
	MaxRetries = 3
	Timeout    = 30 * Second
)

func main() {
	fmt.Printf("1天 = %d 秒\n", Day)
	fmt.Printf("1GB = %d 字节\n", GB)
	
	fmt.Printf("应用: %s v%s\n", AppName, AppVersion)
	fmt.Printf("最大重试次数: %d, 超时时间: %d 秒\n", MaxRetries, Timeout)
}
```

## 🔄 多重赋值和匿名变量

### 多重赋值

```go
package main

import "fmt"

func main() {
	// 交换变量
	a, b := 10, 20
	fmt.Printf("交换前: a = %d, b = %d\n", a, b)
	
	a, b = b, a
	fmt.Printf("交换后: a = %d, b = %d\n", a, b)
	
	// 函数返回多值
	name, age := getUserInfo()
	fmt.Printf("用户信息: %s, %d 岁\n", name, age)
	
	// 循环中的多重赋值
	sum, count := 0, 0
	for i := 1; i <= 10; i++ {
		sum, count = sum+i, count+1
	}
	fmt.Printf("1到10的和: %d, 个数: %d\n", sum, count)
}

func getUserInfo() (string, int) {
	return "李四", 25
}
```

### 匿名变量（空白标识符）

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 忽略不需要的返回值
	_, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Printf("文件打开失败: %v\n", err)
	}
	
	// 忽略不需要的循环变量
	data := []int{5, 2, 8, 1, 9, 3}
	sum := 0
	
	for _, value := range data {
		sum += value
	}
	fmt.Printf("数组元素和: %d\n", sum)
	
	// 忽略索引，只要值
	for _, value := range data {
		if value > 5 {
			fmt.Printf("大于5的值: %d\n", value)
		}
	}
	
	// 忽略值，只要索引
	for index := range data {
		if index%2 == 0 {
			fmt.Printf("偶数索引: %d\n", index)
		}
	}
}
```

## 🌍 作用域

### 局部作用域

```go
package main

import "fmt"

func main() {
	// main 函数的作用域
	x := 100
	fmt.Printf("main 中的 x: %d\n", x)
	
	// if 语句块中的作用域
	if x > 50 {
		y := 200
		fmt.Printf("if 块中的 x: %d, y: %d\n", x, y)
	}
	
	// y 在这里不可访问
	// fmt.Printf("y 的值: %d\n", y) // 编译错误
	
	// for 循环中的作用域
	for i := 0; i < 3; i++ {
		fmt.Printf("循环 %d: x = %d, i = %d\n", i, x, i)
	}
	
	// i 在这里不可访问
	// fmt.Printf("i 的值: %d\n", i) // 编译错误
}
```

### 包级别作用域

```go
package main

import "fmt"

// 包级别变量（全局变量）
var globalCounter int = 0

// 包级别常量
const AppName = "作用域示例"

func incrementCounter() {
	globalCounter++ // 可以访问包级别变量
}

func showLocal() {
	// 局部变量，与全局变量同名
	var globalCounter int = 100
	fmt.Printf("函数内 globalCounter: %d\n", globalCounter)
	fmt.Printf("全局 globalCounter: %d\n", main.globalCounter)
}

func main() {
	fmt.Printf("初始 globalCounter: %d\n", globalCounter)
	fmt.Printf("应用名称: %s\n", AppName)
	
	incrementCounter()
	incrementCounter()
	
	fmt.Printf("递增后 globalCounter: %d\n", globalCounter)
	
	showLocal()
}
```

## 🏪 变量类型推断

### 自动类型推断

```go
package main

import "fmt"

func main() {
	// := 会根据值自动推断类型
	i := 42                    // int
	f := 3.14                  // float64
	s := "Hello, Go!"           // string
	b := true                   // bool
	arr := [3]int{1, 2, 3}    // [3]int
	slice := []int{4, 5, 6}    // []int
	m := map[string]int{"a": 1}  // map[string]int
	
	fmt.Printf("i 的类型: %T, 值: %v\n", i, i)
	fmt.Printf("f 的类型: %T, 值: %v\n", f, f)
	fmt.Printf("s 的类型: %T, 值: %v\n", s, s)
	fmt.Printf("b 的类型: %T, 值: %v\n", b, b)
	fmt.Printf("arr 的类型: %T, 值: %v\n", arr, arr)
	fmt.Printf("slice 的类型: %T, 值: %v\n", slice, slice)
	fmt.Printf("m 的类型: %T, 值: %v\n", m, m)
	
	// 特殊情况：数字字面量的默认类型
	var pi = 3.14159         // float64
	var maxInt = 2147483647    // int
	
	fmt.Printf("pi 类型: %T, maxInt 类型: %T\n", pi, maxInt)
	
	// 明确指定类型
	var pi32 float32 = 3.14159
	var maxInt8 int8 = 127
	
	fmt.Printf("pi32 类型: %T, maxInt8 类型: %T\n", pi32, maxInt8)
}
```

## 🧪 实践练习

### 练习 1：学生信息管理

```go
// 要求：
// 1. 定义学生相关的常量（年级、状态等）
// 2. 声明学生信息变量
// 3. 使用多重赋值交换学生信息
// 4. 测试不同作用域的变量访问
```

参考实现：
```go
package main

import "fmt"

// 年级常量
const (
	Freshman = 1
	Sophomore = 2
	Junior   = 3
	Senior   = 4
)

// 学生状态常量
const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusGraduated = "graduated"
)

func main() {
	// 学生信息
	var studentName string = "王小明"
	var studentID int = 2021001
	var grade int = Sophomore
	var status string = StatusActive
	var gpa float64 = 3.75
	
	fmt.Printf("学生姓名: %s\n", studentName)
	fmt.Printf("学号: %d\n", studentID)
	fmt.Printf("年级: %d\n", grade)
	fmt.Printf("状态: %s\n", status)
	fmt.Printf("GPA: %.2f\n", gpa)
	
	// 交换学生信息
	newName, newID := "李小红", 2021002
	studentName, studentID = newName, newID
	fmt.Printf("更新后姓名: %s, 学号: %d\n", studentName, studentID)
}
```

### 练习 2：配置管理器

```go
// 要求：
// 1. 定义应用配置常量
// 2. 实现配置变量声明
// 3. 添加配置验证函数
// 4. 测试配置的作用域
```

### 练习 3：数据转换工具

```go
// 要求：
// 1. 定义转换常量
// 2. 实现温度转换（摄氏度 ↔ 华氏度）
// 3. 使用多重赋值处理输入输出
// 4. 匿名变量处理错误
```

## 🤔 思考题

1. **什么时候应该使用 `var` 声明，什么时候用 `:=`？**
   - 提示：考虑作用域和明确性

2. **常量为什么不能使用 `:=` 声明？**
   - 提示：考虑编译时和运行时的区别

3. **匿名变量的主要用途是什么？**
   - 提示：考虑函数多返回值的情况

4. **包级别变量和局部变量有什么区别？**
   - 提示：考虑生命周期和作用域

## 📚 扩展阅读

- [Go 语言规范 - 常量](https://golang.org/ref/spec#Constants)
- [Go 语言规范 - 变量](https://golang.org/ref/spec#Variable_declarations)
- [Effective Go - 初始化](https://golang.org/doc/effective_go.html#initialization)

## ⏭️ 下一章节

[数据类型](./02-data-types.md) → 深入了解 Go 语言的各种数据类型

---

**💡 小贴士**: Go 的类型系统设计哲学是"显式优于隐式"，这有助于编写更安全、更易维护的代码。养成明确类型的好习惯！
