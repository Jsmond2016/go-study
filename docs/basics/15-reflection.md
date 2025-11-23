---
title: 反射
difficulty: advanced
duration: "4-5小时"
prerequisites: ["接口", "类型系统"]
tags: ["反射", "reflect", "类型检查", "动态调用"]
---

# 反射

反射是 Go 语言提供的一种在运行时检查类型和值的能力。虽然反射功能强大，但应该谨慎使用。

## 📋 学习目标

- [ ] 理解反射的概念和用途
- [ ] 掌握 reflect 包的使用
- [ ] 学会类型检查和值操作
- [ ] 理解反射的性能影响
- [ ] 了解反射的使用场景
- [ ] 掌握反射的最佳实践

## 🎯 反射基础

### 什么是反射

反射允许程序在运行时检查类型信息和操作值，这在某些场景下非常有用，如 JSON 序列化、ORM 框架等。

### reflect 包

Go 的反射功能通过 `reflect` 包提供：

```go
import "reflect"
```

## 🔍 类型检查

### Type 和 Value

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42

	// 获取类型
	t := reflect.TypeOf(x)
	fmt.Printf("类型: %v\n", t)
	fmt.Printf("类型名称: %v\n", t.Name())
	fmt.Printf("类型种类: %v\n", t.Kind())

	// 获取值
	v := reflect.ValueOf(x)
	fmt.Printf("值: %v\n", v)
	fmt.Printf("值的类型: %v\n", v.Type())
}
```

### Kind

`Kind` 表示类型的基础种类：

```go
package main

import (
	"fmt"
	"reflect"
)

func printKind(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("类型: %v, 种类: %v\n", t, t.Kind())
}

func main() {
	printKind(42)              // int
	printKind(3.14)            // float64
	printKind("hello")         // string
	printKind([]int{1, 2, 3})  // slice
	printKind(map[string]int{}) // map
}
```

## 💡 值操作

### 获取和设置值

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42

	// 获取值
	v := reflect.ValueOf(x)
	fmt.Printf("值: %v\n", v.Int())

	// 设置值（需要指针）
	p := reflect.ValueOf(&x)
	e := p.Elem()
	e.SetInt(100)
	fmt.Printf("新值: %v\n", x)
}
```

### 检查值是否可设置

```go
func setValue(v interface{}, newValue int) {
	rv := reflect.ValueOf(v)

	// 必须是指针
	if rv.Kind() != reflect.Ptr {
		fmt.Println("不是指针，无法设置")
		return
	}

	// 获取指向的值
	elem := rv.Elem()

	// 检查是否可设置
	if !elem.CanSet() {
		fmt.Println("值不可设置")
		return
	}

	// 设置值
	if elem.Kind() == reflect.Int {
		elem.SetInt(int64(newValue))
	}
}
```

## 🏗️ 结构体反射

### 遍历结构体字段

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	City string
}

func printStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fmt.Printf("结构体: %v\n", t.Name())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("  字段: %s, 类型: %v, 值: %v\n",
			field.Name, field.Type, value.Interface())
	}
}

func main() {
	p := Person{
		Name: "张三",
		Age:  30,
		City: "北京",
	}
	printStruct(p)
}
```

### 获取标签

```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID    int    `json:"id" db:"user_id"`
	Name  string `json:"name" db:"user_name"`
	Email string `json:"email" db:"email"`
}

func printTags(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")

		fmt.Printf("字段: %s\n", field.Name)
		fmt.Printf("  JSON 标签: %s\n", jsonTag)
		fmt.Printf("  DB 标签: %s\n", dbTag)
	}
}

func main() {
	u := User{}
	printTags(u)
}
```

## 🔧 方法反射

### 调用方法

```go
package main

import (
	"fmt"
	"reflect"
)

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func main() {
	calc := Calculator{}
	v := reflect.ValueOf(calc)

	// 调用 Add 方法
	method := v.MethodByName("Add")
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}
	result := method.Call(args)
	fmt.Printf("结果: %v\n", result[0].Int())
}
```

### 遍历方法

```go
func printMethods(v interface{}) {
	t := reflect.TypeOf(v)

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("方法: %s\n", method.Name)
		fmt.Printf("  类型: %v\n", method.Type)
	}
}
```

## 🎯 常见应用

### JSON 序列化（简化版）

```go
func toJSON(v interface{}) string {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}

	result := "{"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf(`"%s":`, jsonTag)

		switch value.Kind() {
		case reflect.String:
			result += fmt.Sprintf(`"%v"`, value.String())
		case reflect.Int:
			result += fmt.Sprintf(`%v`, value.Int())
		default:
			result += fmt.Sprintf(`"%v"`, value.Interface())
		}
	}
	result += "}"
	return result
}
```

### 函数调用

```go
func callFunction(fn interface{}, args ...interface{}) []reflect.Value {
	fnValue := reflect.ValueOf(fn)

	// 准备参数
	argsValue := make([]reflect.Value, len(args))
	for i, arg := range args {
		argsValue[i] = reflect.ValueOf(arg)
	}

	// 调用函数
	return fnValue.Call(argsValue)
}

func add(a, b int) int {
	return a + b
}

func main() {
	result := callFunction(add, 10, 20)
	fmt.Println(result[0].Int()) // 30
}
```

## ⚠️ 注意事项

### 1. 性能影响

反射比直接调用慢很多，应该避免在性能关键路径使用。

```go
// ❌ 慢
method := reflect.ValueOf(obj).MethodByName("Method")
method.Call(args)

// ✅ 快
obj.Method(args)
```

### 2. 类型安全

反射绕过了编译时的类型检查，可能导致运行时错误。

### 3. 代码可读性

过度使用反射会让代码难以理解和维护。

## 🏃‍♂️ 实践练习

### 练习 1: 结构体验证器

使用反射实现一个结构体字段验证器。

### 练习 2: 配置加载器

使用反射从配置文件加载结构体。

### 练习 3: ORM 查询构建器

使用反射构建 SQL 查询。

## 🤔 思考题

1. 什么时候应该使用反射？
2. 反射的性能影响有多大？
3. 如何安全地使用反射？
4. 反射和接口有什么区别？
5. 有哪些替代反射的方案？

## 📚 扩展阅读

- [Go 反射官方文档](https://golang.org/pkg/reflect/)
- [反射的使用场景](https://blog.golang.org/laws-of-reflection)
- [反射最佳实践](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)

## ⏭️ 下一章节

[测试](./16-testing.md) → 学习 Go 语言的测试方法

---

**💡 提示**: 反射是强大的工具，但应该谨慎使用。优先考虑接口和类型断言，只在必要时使用反射！

