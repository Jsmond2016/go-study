// Package main 展示 Go 语言的反射
package main

import (
	"fmt"
	"reflect"
)

// 反射基础示例
func reflectionBasics() {
	var x float64 = 3.14

	// 获取类型
	t := reflect.TypeOf(x)
	fmt.Printf("类型: %v\n", t)

	// 获取值
	v := reflect.ValueOf(x)
	fmt.Printf("值: %v\n", v)
	fmt.Printf("类型: %v\n", v.Type())
	fmt.Printf("种类: %v\n", v.Kind())
}

// 修改值
func modifyValue() {
	x := 3.14

	// 必须使用指针才能修改
	v := reflect.ValueOf(&x)
	fmt.Printf("修改前: %v\n", x)

	// 获取指针指向的值
	v = v.Elem()
	v.SetFloat(2.71)
	fmt.Printf("修改后: %v\n", x)
}

// 结构体反射
type Person struct {
	Name string
	Age  int
}

func structReflection() {
	p := Person{Name: "Alice", Age: 30}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Printf("结构体类型: %v\n", t)
	fmt.Printf("字段数量: %d\n", t.NumField())

	// 遍历字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("字段 %d: %s = %v\n", i, field.Name, value.Interface())
	}
}

// 方法反射
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, %d years old", p.Name, p.Age)
}

func methodReflection() {
	p := Person{Name: "Bob", Age: 25}
	v := reflect.ValueOf(p)

	// 获取方法
	method := v.MethodByName("Greet")
	if method.IsValid() {
		result := method.Call(nil)
		fmt.Printf("方法调用结果: %v\n", result[0].String())
	}
}

// 类型断言与反射
func typeAssertion() {
	var x interface{} = 42

	// 使用类型断言
	if v, ok := x.(int); ok {
		fmt.Printf("类型断言: %d\n", v)
	}

	// 使用反射
	v := reflect.ValueOf(x)
	fmt.Printf("反射类型: %v, 值: %v\n", v.Type(), v.Interface())
}

func main() {
	fmt.Println("=== Go 反射示例 ===\n")

	// 1. 反射基础
	fmt.Println("--- 反射基础 ---")
	reflectionBasics()

	// 2. 修改值
	fmt.Println("\n--- 修改值 ---")
	modifyValue()

	// 3. 结构体反射
	fmt.Println("\n--- 结构体反射 ---")
	structReflection()

	// 4. 方法反射
	fmt.Println("\n--- 方法反射 ---")
	methodReflection()

	// 5. 类型断言
	fmt.Println("\n--- 类型断言与反射 ---")
	typeAssertion()
}

