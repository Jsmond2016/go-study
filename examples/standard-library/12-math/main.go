// Package main 展示 Go 语言的数学运算
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Go 数学运算示例 ===\n")

	// 1. 基本数学函数
	fmt.Println("--- 基本数学函数 ---")
	fmt.Printf("π = %.2f\n", math.Pi)
	fmt.Printf("e = %.2f\n", math.E)
	fmt.Printf("√2 = %.2f\n", math.Sqrt(2))
	fmt.Printf("2^3 = %.0f\n", math.Pow(2, 3))

	// 2. 三角函数
	fmt.Println("\n--- 三角函数 ---")
	angle := math.Pi / 4 // 45度
	fmt.Printf("sin(45°) = %.2f\n", math.Sin(angle))
	fmt.Printf("cos(45°) = %.2f\n", math.Cos(angle))
	fmt.Printf("tan(45°) = %.2f\n", math.Tan(angle))

	// 3. 对数函数
	fmt.Println("\n--- 对数函数 ---")
	fmt.Printf("ln(e) = %.2f\n", math.Log(math.E))
	fmt.Printf("log10(100) = %.2f\n", math.Log10(100))
	fmt.Printf("log2(8) = %.2f\n", math.Log2(8))

	// 4. 取整函数
	fmt.Println("\n--- 取整函数 ---")
	fmt.Printf("ceil(3.7) = %.0f\n", math.Ceil(3.7))
	fmt.Printf("floor(3.7) = %.0f\n", math.Floor(3.7))
	fmt.Printf("round(3.7) = %.0f\n", math.Round(3.7))

	// 5. 最大值和最小值
	fmt.Println("\n--- 最大值和最小值 ---")
	fmt.Printf("max(3, 5) = %.0f\n", math.Max(3, 5))
	fmt.Printf("min(3, 5) = %.0f\n", math.Min(3, 5))

	// 6. 绝对值
	fmt.Println("\n--- 绝对值 ---")
	fmt.Printf("abs(-5) = %.0f\n", math.Abs(-5))
}

