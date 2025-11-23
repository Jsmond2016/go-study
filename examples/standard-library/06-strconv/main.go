// Package main 展示 Go 语言的字符串转换
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Go 字符串转换示例 ===\n")

	// 1. 字符串转整数
	fmt.Println("--- 字符串转整数 ---")
	i, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("整数: %d\n", i)
	}

	// 不同进制
	hex, _ := strconv.ParseInt("FF", 16, 64)
	oct, _ := strconv.ParseInt("777", 8, 64)
	bin, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Printf("十六进制 FF: %d\n", hex)
	fmt.Printf("八进制 777: %d\n", oct)
	fmt.Printf("二进制 1010: %d\n", bin)

	// 2. 整数转字符串
	fmt.Println("\n--- 整数转字符串 ---")
	s := strconv.Itoa(123)
	fmt.Printf("字符串: %s\n", s)

	s2 := strconv.FormatInt(255, 16) // 十六进制
	fmt.Printf("255 的十六进制: %s\n", s2)

	s3 := strconv.FormatInt(255, 2) // 二进制
	fmt.Printf("255 的二进制: %s\n", s3)

	// 3. 字符串转浮点数
	fmt.Println("\n--- 字符串转浮点数 ---")
	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("浮点数: %.2f\n", f)
	}

	// 4. 浮点数转字符串
	fmt.Println("\n--- 浮点数转字符串 ---")
	fs := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("浮点数字符串: %s\n", fs)

	// 5. 布尔值转换
	fmt.Println("\n--- 布尔值转换 ---")
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Printf("转换失败: %v\n", err)
	} else {
		fmt.Printf("布尔值: %t\n", b)
	}

	bs := strconv.FormatBool(true)
	fmt.Printf("布尔值字符串: %s\n", bs)

	// 6. 引号处理
	fmt.Println("\n--- 引号处理 ---")
	quoted := strconv.Quote("Hello, World")
	fmt.Printf("加引号: %s\n", quoted)

	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		fmt.Printf("去引号失败: %v\n", err)
	} else {
		fmt.Printf("去引号: %s\n", unquoted)
	}
}

