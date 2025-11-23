// Package main 展示 Go 语言的文件操作
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Go 文件操作示例 ===\n")

	// 1. 写入文件
	fmt.Println("--- 写入文件 ---")
	content := "Hello, Go!\n这是第二行\n"
	err := os.WriteFile("test.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
	} else {
		fmt.Println("文件写入成功")
	}

	// 2. 读取整个文件
	fmt.Println("\n--- 读取整个文件 ---")
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
	} else {
		fmt.Printf("文件内容:\n%s", string(data))
	}

	// 3. 逐行读取
	fmt.Println("\n--- 逐行读取 ---")
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("第 %d 行: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取错误: %v\n", err)
	}

	// 4. 文件信息
	fmt.Println("\n--- 文件信息 ---")
	info, err := os.Stat("test.txt")
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
	} else {
		fmt.Printf("文件名: %s\n", info.Name())
		fmt.Printf("大小: %d 字节\n", info.Size())
		fmt.Printf("修改时间: %v\n", info.ModTime())
	}

	// 5. 检查文件是否存在
	fmt.Println("\n--- 检查文件是否存在 ---")
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}

	// 清理
	os.Remove("test.txt")
}
