// Package main 展示 Go 语言的接口使用
package main

import (
	"fmt"
)

// Animal 定义动物接口
type Animal interface {
	Speak() string
	Move() string
}

// Dog 实现 Animal 接口
type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s 说: 汪汪汪", d.name)
}

func (d Dog) Move() string {
	return fmt.Sprintf("%s 在跑", d.name)
}

// Cat 实现 Animal 接口
type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s 说: 喵喵喵", c.name)
}

func (c Cat) Move() string {
	return fmt.Sprintf("%s 在走", c.name)
}

// Writer 定义写入接口
type Writer interface {
	Write([]byte) (int, error)
}

// File 实现 Writer 接口
type File struct {
	name string
}

func (f *File) Write(data []byte) (int, error) {
	fmt.Printf("写入文件 %s: %s\n", f.name, string(data))
	return len(data), nil
}

// 空接口示例
func printAnything(v interface{}) {
	fmt.Printf("类型: %T, 值: %v\n", v, v)
}

// 类型断言示例
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

// 接口组合示例
type Reader interface {
	Read() []byte
}

type ReadWriter interface {
	Reader
	Writer
}

type Buffer struct {
	data []byte
}

func (b *Buffer) Read() []byte {
	return b.data
}

func (b *Buffer) Write(data []byte) (int, error) {
	b.data = append(b.data, data...)
	return len(data), nil
}

func main() {
	fmt.Println("=== Go 接口示例 ===")
	fmt.Println()

	// 1. 基本接口使用
	fmt.Println("--- 基本接口使用 ---")
	var animals []Animal
	animals = append(animals, Dog{name: "旺财"})
	animals = append(animals, Cat{name: "咪咪"})

	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
	}

	// 2. Writer 接口
	fmt.Println("\n--- Writer 接口 ---")
	var w Writer
	file := &File{name: "test.txt"}
	w = file
	data := []byte("Hello, Go!")
	n, err := w.Write(data)
	if err != nil {
		fmt.Printf("写入错误: %v\n", err)
	} else {
		fmt.Printf("成功写入 %d 字节\n", n)
	}

	// 3. 空接口
	fmt.Println("\n--- 空接口 ---")
	printAnything(42)
	printAnything("Hello")
	printAnything(3.14)
	printAnything([]int{1, 2, 3})

	// 4. 类型断言
	fmt.Println("\n--- 类型断言 ---")
	processValue("Hello")
	processValue(42)
	processValue(3.14)
	processValue(true)

	// 5. 接口组合
	fmt.Println("\n--- 接口组合 ---")
	buf := &Buffer{}
	var rw ReadWriter = buf
	rw.Write([]byte("Hello"))
	fmt.Printf("读取: %s\n", string(rw.Read()))
}
