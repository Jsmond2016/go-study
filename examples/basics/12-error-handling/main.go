// Package main 展示 Go 语言的错误处理
package main

import (
	"errors"
	"fmt"
	"os"
)

// FileNotFoundError 自定义错误类型
type FileNotFoundError struct {
	FileName string
	Path     string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("文件 %s 在路径 %s 中未找到", e.FileName, e.Path)
}

// 基本错误处理
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 使用 fmt.Errorf 创建错误
func readFile(filename string) (string, error) {
	if filename == "" {
		return "", fmt.Errorf("文件名不能为空")
	}
	// 模拟文件不存在
	if filename == "notfound.txt" {
		return "", &FileNotFoundError{
			FileName: filename,
			Path:     "/tmp",
		}
	}
	return "文件内容", nil
}

// 错误包装
func processFile(filename string) error {
	content, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("处理文件失败: %w", err)
	}
	fmt.Printf("文件内容: %s\n", content)
	return nil
}

// 错误检查
func checkError(err error) {
	if err != nil {
		fmt.Printf("发生错误: %v\n", err)
		// 检查是否是自定义错误类型
		var fileErr *FileNotFoundError
		if errors.As(err, &fileErr) {
			fmt.Printf("文件错误详情: %s\n", fileErr.Error())
		}
	}
}

// Panic 和 Recover 示例
func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 panic: %v\n", r)
		}
	}()

	panic("这是一个 panic")
	fmt.Println("这行不会执行")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("安全捕获: %v\n", r)
		}
	}()

	mayPanic()
	fmt.Println("程序继续执行")
}

func main() {
	fmt.Println("=== Go 错误处理示例 ===\n")

	// 1. 基本错误处理
	fmt.Println("--- 基本错误处理 ---")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	// 2. 自定义错误
	fmt.Println("\n--- 自定义错误 ---")
	_, err = readFile("notfound.txt")
	checkError(err)

	// 3. 错误包装
	fmt.Println("\n--- 错误包装 ---")
	err = processFile("notfound.txt")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		// 检查原始错误
		var fileErr *FileNotFoundError
		if errors.As(err, &fileErr) {
			fmt.Printf("原始错误: %v\n", fileErr)
		}
	}

	// 4. Panic 和 Recover
	fmt.Println("\n--- Panic 和 Recover ---")
	safeCall()
	fmt.Println("程序正常结束")

	// 5. 错误比较
	fmt.Println("\n--- 错误比较 ---")
	err1 := errors.New("错误1")
	err2 := errors.New("错误1")
	fmt.Printf("err1 == err2: %v\n", err1 == err2) // false，因为是指针比较

	// 使用 errors.Is 比较错误值
	targetErr := errors.New("目标错误")
	wrappedErr := fmt.Errorf("包装: %w", targetErr)
	fmt.Printf("errors.Is(wrappedErr, targetErr): %v\n", errors.Is(wrappedErr, targetErr)) // true
}

