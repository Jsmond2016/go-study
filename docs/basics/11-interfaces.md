---
title: 接口
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["结构体", "方法", "指针"]
tags: ["接口", "多态", "类型断言", "空接口"]
---

# 接口

接口是 Go 语言中实现多态的核心机制。通过接口，我们可以定义行为契约，而不需要关心具体的实现类型。

## 📋 学习目标

- [ ] 理解接口的概念和用途
- [ ] 掌握接口的定义和实现
- [ ] 学会使用接口实现多态
- [ ] 理解空接口和类型断言
- [ ] 掌握接口的组合和嵌套
- [ ] 了解接口的最佳实践

## 🎯 接口基础

### 什么是接口

接口定义了一组方法的签名，任何实现了这些方法的类型都实现了该接口。

```go
package main

import "fmt"

// 定义接口
type Writer interface {
	Write([]byte) (int, error)
}

// 实现接口的类型
type File struct {
	name string
}

// 实现接口方法
func (f *File) Write(data []byte) (int, error) {
	fmt.Printf("写入文件 %s: %s\n", f.name, string(data))
	return len(data), nil
}

func main() {
	var w Writer
	w = &File{name: "test.txt"}
	
	data := []byte("Hello, Go!")
	n, err := w.Write(data)
	if err != nil {
		fmt.Printf("写入错误: %v\n", err)
	} else {
		fmt.Printf("成功写入 %d 字节\n", n)
	}
}
```

### 接口的特点

1. **隐式实现**: 不需要显式声明实现接口
2. **类型安全**: 编译时检查接口实现
3. **解耦**: 接口与实现分离

## 🔍 接口定义

### 基本语法

```go
type InterfaceName interface {
	Method1(param1 type1) returnType1
	Method2(param2 type2) returnType2
}
```

### 示例：动物接口

```go
package main

import "fmt"

// 定义动物接口
type Animal interface {
	Speak() string
	Move() string
}

// 实现 Animal 接口的类型
type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s 说: 汪汪汪", d.name)
}

func (d Dog) Move() string {
	return fmt.Sprintf("%s 在跑", d.name)
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s 说: 喵喵喵", c.name)
}

func (c Cat) Move() string {
	return fmt.Sprintf("%s 在走", c.name)
}

func main() {
	animals := []Animal{
		Dog{name: "旺财"},
		Cat{name: "小花"},
	}
	
	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
		fmt.Println()
	}
}
```

## 💡 接口的多态性

接口实现了多态，同一个接口可以表示不同的类型。

```go
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 5}
	
	printShapeInfo(rect)
	printShapeInfo(circle)
}
```

## 🔑 空接口

空接口 `interface{}` 可以表示任何类型，类似于其他语言中的 `any` 或 `Object`。

```go
package main

import "fmt"

func printAnything(v interface{}) {
	fmt.Printf("类型: %T, 值: %v\n", v, v)
}

func main() {
	printAnything(42)
	printAnything("Hello")
	printAnything(3.14)
	printAnything([]int{1, 2, 3})
	printAnything(map[string]int{"a": 1})
}
```

### Go 1.18+ 的 any 类型

Go 1.18 引入了 `any` 作为 `interface{}` 的别名：

```go
func printAnything(v any) {
	fmt.Printf("类型: %T, 值: %v\n", v, v)
}
```

## 🎭 类型断言

类型断言用于检查接口值的具体类型。

### 基本语法

```go
value, ok := interfaceValue.(Type)
```

### 示例

```go
package main

import "fmt"

func processValue(v interface{}) {
	// 类型断言
	if str, ok := v.(string); ok {
		fmt.Printf("字符串值: %s\n", str)
	} else if num, ok := v.(int); ok {
		fmt.Printf("整数值: %d\n", num)
	} else {
		fmt.Printf("未知类型: %T\n", v)
	}
}

func main() {
	processValue("Hello")
	processValue(42)
	processValue(3.14)
}
```

### 类型开关

使用 `switch` 进行类型判断：

```go
func processValue(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Printf("字符串: %s\n", val)
	case int:
		fmt.Printf("整数: %d\n", val)
	case float64:
		fmt.Printf("浮点数: %f\n", val)
	default:
		fmt.Printf("未知类型: %T\n", val)
	}
}
```

## 🔗 接口组合

接口可以组合，形成新的接口。

```go
package main

import "fmt"

// 基础接口
type Reader interface {
	Read() []byte
}

type Writer interface {
	Write([]byte) error
}

// 组合接口
type ReadWriter interface {
	Reader
	Writer
}

// 实现组合接口
type File struct {
	data []byte
}

func (f *File) Read() []byte {
	return f.data
}

func (f *File) Write(data []byte) error {
	f.data = data
	return nil
}

func main() {
	var rw ReadWriter = &File{}
	rw.Write([]byte("Hello, Go!"))
	fmt.Println(string(rw.Read()))
}
```

## 📚 标准库中的接口

### io.Reader 和 io.Writer

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func readFromReader(r io.Reader) {
	data := make([]byte, 100)
	n, err := r.Read(data)
	if err != nil && err != io.EOF {
		fmt.Printf("读取错误: %v\n", err)
		return
	}
	fmt.Printf("读取了 %d 字节: %s\n", n, string(data[:n]))
}

func main() {
	// 从字符串读取
	reader := strings.NewReader("Hello, Go!")
	readFromReader(reader)
	
	// 从字节切片读取
	readFromReader(strings.NewReader("测试数据"))
}
```

### fmt.Stringer

实现 `String()` 方法可以让类型自定义字符串表示：

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d 岁)", p.Name, p.Age)
}

func main() {
	p := Person{Name: "张三", Age: 30}
	fmt.Println(p) // 输出: 张三 (30 岁)
}
```

## 🏃‍♂️ 实践练习

### 练习 1: 实现排序接口

实现一个可以对自定义类型进行排序的接口：

```go
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

### 练习 2: 实现数据库接口

设计一个数据库接口，包含 `Connect()`, `Query()`, `Close()` 方法。

### 练习 3: 实现缓存接口

设计一个缓存接口，支持 `Get()`, `Set()`, `Delete()` 操作。

## 🤔 思考题

1. 接口和结构体有什么区别？
2. 为什么 Go 使用隐式接口实现？
3. 空接口的使用场景有哪些？
4. 类型断言和类型转换有什么区别？
5. 如何判断一个类型是否实现了某个接口？

## 📚 扩展阅读

- [Go 语言接口详解](https://golang.org/doc/effective_go.html#interfaces)
- [接口最佳实践](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)
- [接口设计原则](https://golang.org/doc/effective_go.html#interfaces_and_types)

## ⏭️ 下一章节

[错误处理](./12-error-handling.md) → 学习 Go 语言的错误处理机制

---

**💡 提示**: 接口是 Go 语言实现多态的核心，理解接口有助于编写更灵活、可扩展的代码！

