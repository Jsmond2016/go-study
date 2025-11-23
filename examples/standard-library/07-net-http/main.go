// Package main 展示 Go 语言的 HTTP 客户端
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== Go HTTP 客户端示例 ===\n")

	// 1. GET 请求
	fmt.Println("--- GET 请求 ---")
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}
	fmt.Printf("响应状态: %s\n", resp.Status)
	fmt.Printf("响应长度: %d 字节\n", len(body))

	// 2. POST 请求
	fmt.Println("\n--- POST 请求 ---")
	data := map[string]string{
		"name": "Go",
		"type": "Language",
	}
	jsonData, _ := json.Marshal(data)

	resp, err = http.Post("https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("POST 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("POST 响应状态: %s\n", resp.Status)

	// 3. 自定义请求
	fmt.Println("\n--- 自定义请求 ---")
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	req.Header.Set("User-Agent", "Go-Client/1.0")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("自定义请求状态: %s\n", resp.Status)
}

