// Package main 展示 Go 语言的日志记录
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("=== Go 日志示例 ===\n")

	// 1. 基本日志输出
	fmt.Println("--- 基本日志输出 ---")
	log.Print("这是一条普通日志")
	log.Println("这是一条带换行的日志")
	log.Printf("这是一条格式化日志: %s", "值")

	// 2. 日志级别
	fmt.Println("\n--- 日志级别 ---")
	log.Println("普通日志")
	// log.Fatal("致命错误，程序退出") // 取消注释会导致程序退出
	// log.Panic("程序崩溃") // 取消注释会导致程序崩溃

	// 3. 设置日志前缀
	fmt.Println("\n--- 设置日志前缀 ---")
	log.SetPrefix("[APP] ")
	log.Println("这条日志有前缀")

	// 4. 设置日志标志
	fmt.Println("\n--- 设置日志标志 ---")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("带日期、时间和文件名的日志")

	// 5. 输出到文件
	fmt.Println("\n--- 输出到文件 ---")
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开日志文件:", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("这条日志写入文件")
}

