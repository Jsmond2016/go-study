// Package main 展示 Go 语言的包管理
package main

import (
	"fmt"
	"math"
)

// 包级别变量
var packageVar = "包级别变量"

// 包级别函数
func packageFunction() {
	fmt.Println("包级别函数")
}

// init 函数在 main 之前执行
func init() {
	fmt.Println("init 函数执行")
}

func main() {
	fmt.Println("=== Go 包管理示例 ===\n")

	// 1. 使用标准库包
	fmt.Println("--- 使用标准库包 ---")
	fmt.Printf("π = %.2f\n", math.Pi)
	fmt.Printf("√2 = %.2f\n", math.Sqrt(2))

	// 2. 包级别变量和函数
	fmt.Println("\n--- 包级别变量和函数 ---")
	fmt.Println(packageVar)
	packageFunction()

	// 3. 导入别名
	fmt.Println("\n--- 导入别名 ---")
	// import m "math"
	// fmt.Println(m.Pi)
}

