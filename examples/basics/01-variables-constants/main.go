// Package main 提供变量与常量的基础示例
package main

import (
	"fmt"
)

// 程序入口，演示各种变量和常量的使用
func main() {
	fmt.Println("=== Go 变量与常量示例 ===")

	// 演示常量定义
	demonstrateConstants()

	// 演示各种变量声明方式
	demonstrateVariables()

	// 演示类型转换
	demonstrateTypeConversion()

	// 演示变量作用域
	demonstrateVariableScope()
}

// demonstrateConstants 展示常量的定义和使用
func demonstrateConstants() {
	fmt.Println("\n--- 常量定义 ---")

	// 使用 const 关键字定义常量
	const PI = 3.14159
	const Language = "Go"
	const IsCompleted = true

	fmt.Printf("圆周率: %.5f\n", PI)
	fmt.Printf("编程语言: %s\n", Language)
	fmt.Printf("是否完成: %t\n", IsCompleted)

	// 常量组定义
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
	)

	fmt.Printf("工作日编号: 周一=%d, 周五=%d\n", Monday, Friday)

	// iota 枚举
	const (
		Spring = iota
		Summer
		Autumn
		Winter
		SeasonCount
	)

	fmt.Printf("季节枚举: 春=%d, 夏=%d, 秋=%d, 冬=%d, 季节数量=%d\n",
		Spring, Summer, Autumn, Winter, SeasonCount)
}

// demonstrateVariables 展示变量的声明和使用
func demonstrateVariables() {
	fmt.Println("\n--- 变量声明 ---")

	// 1. 使用 var 关键字声明变量
	var name string = "张三"
	var age int = 25
	var height float64 = 175.5
	var isStudent bool = false

	fmt.Printf("姓名: %s, 年龄: %d, 身高: %.1fcm, 是否学生: %t\n",
		name, age, height, isStudent)

	// 2. 类型推断
	var city = "北京"
	var score = 95.5

	fmt.Printf("城市: %s, 分数: %.1f\n", city, score)

	// 3. 短变量声明（只能在函数内部使用）
	email := "zhangsan@example.com"
	phone := "13800138000"

	fmt.Printf("邮箱: %s, 电话: %s\n", email, phone)

	// 4. 多变量声明
	var (
		company  string = "科技公司"
		position string = "工程师"
		salary   float64 = 15000.0
	)

	fmt.Printf("公司: %s, 职位: %s, 薪资: %.1f元\n", company, position, salary)

	// 5. 多变量并行声明
	x, y, z := 10, 20, 30
	fmt.Printf("坐标: (%d, %d, %d)\n", x, y, z)

	// 6. 变量交换
	a, b := 100, 200
	fmt.Printf("交换前: a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("交换后: a=%d, b=%d\n", a, b)
}

// demonstrateTypeConversion 展示类型转换
func demonstrateTypeConversion() {
	fmt.Println("\n--- 类型转换 ---")

	// 数值类型转换
	var integer int = 42
	var floating float64 = 3.14

	// int 转 float64
	floatFromInt := float64(integer)
	fmt.Printf("%d 转 float64: %.2f\n", integer, floatFromInt)

	// float64 转 int（注意：会丢失小数部分）
	intFromFloat := int(floating)
	fmt.Printf("%.2f 转 int: %d\n", floating, intFromFloat)

	// 字符串转换
	var text string = "123"
	var number int

	// 注意：字符串不能直接转换为数字
	fmt.Printf("字符串 '%s' 需要使用 strconv 包来转换为数字\n", text)

	// 布尔值与字符串
	var isReady bool = true
	fmt.Printf("布尔值: %t\n", isReady)
}

// demonstrateVariableScope 展示变量作用域
func demonstrateVariableScope() {
	fmt.Println("\n--- 变量作用域 ---")

	// 全局变量作用域演示
	globalVar := "全局变量"

	// 内部作用域
	{
		localVar := "局部变量"
		fmt.Printf("%s\n", globalVar)  // 可以访问全局变量
		fmt.Printf("%s\n", localVar)   // 可以访问局部变量

		// 声明同名的局部变量（遮蔽全局变量）
		globalVar := "遮蔽的全局变量"
		fmt.Printf("在内部作用域中: %s\n", globalVar)
	}

	// 外部作用域
	fmt.Printf("在外部作用域中: %s\n", globalVar) // 仍然是原来的全局变量
	// fmt.Printf(localVar) // 编译错误：无法访问内部作用域的变量
}

// 全局变量
var globalVersion = "1.0.0"
var globalAuthor = "Go学习项目"

// init 函数在 main 函数之前执行
func init() {
	fmt.Printf("\n项目版本: %s, 作者: %s\n", globalVersion, globalAuthor)
}