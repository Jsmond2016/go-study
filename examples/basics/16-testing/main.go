// Package main 展示 Go 语言的测试
package main

import (
	"fmt"
)

// Add 函数用于测试
func Add(a, b int) int {
	return a + b
}

// Multiply 函数用于测试
func Multiply(a, b int) int {
	return a * b
}

// Divide 函数用于测试
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== Go 测试示例 ===\n")
	fmt.Println("运行测试: go test")
	fmt.Println("运行测试并查看覆盖率: go test -cover")
	fmt.Println("运行基准测试: go test -bench=.")
}

