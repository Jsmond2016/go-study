// Package main 展示 Go 语言的上下文管理
package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, name string) {
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 操作被取消\n", name)
			return
		default:
			fmt.Printf("%s: 执行中... (%d)\n", name, i+1)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("%s: 完成\n", name)
}

func main() {
	fmt.Println("=== Go Context 示例 ===\n")

	// 1. 基本 Context
	fmt.Println("--- 基本 Context ---")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "123")
	userID := ctx.Value("userID")
	fmt.Printf("用户ID: %v\n", userID)

	// 2. 超时控制
	fmt.Println("\n--- 超时控制 ---")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doWork(ctx, "任务1")
	time.Sleep(3 * time.Second)

	// 3. 取消操作
	fmt.Println("\n--- 取消操作 ---")
	ctx, cancel = context.WithCancel(context.Background())
	go doWork(ctx, "任务2")
	time.Sleep(1 * time.Second)
	cancel() // 取消操作
	time.Sleep(500 * time.Millisecond)

	// 4. 截止时间
	fmt.Println("\n--- 截止时间 ---")
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("截止时间到达: %v\n", ctx.Err())
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("执行中...")
			}
		}
	}()
	time.Sleep(3 * time.Second)
}

