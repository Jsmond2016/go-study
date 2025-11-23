// Package main 展示 Go 语言的字符串操作
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Go 字符串操作示例 ===\n")

	str := "Hello, World"

	// 1. 字符串查找
	fmt.Println("--- 字符串查找 ---")
	contains := strings.Contains(str, "World")
	fmt.Printf("包含 'World': %t\n", contains)

	index := strings.Index(str, "World")
	fmt.Printf("'World' 的位置: %d\n", index)

	count := strings.Count(str, "l")
	fmt.Printf("'l' 出现次数: %d\n", count)

	// 2. 前缀和后缀
	fmt.Println("\n--- 前缀和后缀 ---")
	hasPrefix := strings.HasPrefix(str, "Hello")
	fmt.Printf("以 'Hello' 开头: %t\n", hasPrefix)

	hasSuffix := strings.HasSuffix(str, "World")
	fmt.Printf("以 'World' 结尾: %t\n", hasSuffix)

	// 3. 字符串替换
	fmt.Println("\n--- 字符串替换 ---")
	newStr := strings.Replace(str, "World", "Go", 1)
	fmt.Printf("替换后: %s\n", newStr)

	newStr2 := strings.ReplaceAll(str, "l", "L")
	fmt.Printf("全部替换: %s\n", newStr2)

	// 4. 大小写转换
	fmt.Println("\n--- 大小写转换 ---")
	upper := strings.ToUpper(str)
	fmt.Printf("大写: %s\n", upper)

	lower := strings.ToLower(str)
	fmt.Printf("小写: %s\n", lower)

	// 5. 字符串修剪
	fmt.Println("\n--- 字符串修剪 ---")
	trimmed := strings.TrimSpace("  Hello, World  ")
	fmt.Printf("去除空白: '%s'\n", trimmed)

	trimmed2 := strings.Trim("!!!Hello!!!", "!")
	fmt.Printf("去除 '!': %s\n", trimmed2)

	// 6. 字符串分割
	fmt.Println("\n--- 字符串分割 ---")
	parts := strings.Split("apple,banana,orange", ",")
	fmt.Printf("分割结果: %v\n", parts)

	words := strings.Fields("  hello   world  go  ")
	fmt.Printf("按空白分割: %v\n", words)

	// 7. 字符串拼接
	fmt.Println("\n--- 字符串拼接 ---")
	joined := strings.Join(parts, " | ")
	fmt.Printf("拼接结果: %s\n", joined)

	// 8. 字符串构建器
	fmt.Println("\n--- 字符串构建器 ---")
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("Go")
	result := builder.String()
	fmt.Printf("Builder 结果: %s\n", result)
}

