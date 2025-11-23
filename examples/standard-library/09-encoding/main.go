// Package main 展示 Go 语言的编码解码
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("=== Go 编码解码示例 ===\n")

	// 1. JSON 编码
	fmt.Println("--- JSON 编码 ---")
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("JSON 编码失败: %v\n", err)
	} else {
		fmt.Printf("JSON 数据: %s\n", jsonData)
	}

	// 2. JSON 解码
	fmt.Println("\n--- JSON 解码 ---")
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Printf("JSON 解码失败: %v\n", err)
	} else {
		fmt.Printf("解码结果: %+v\n", p2)
	}

	// 3. Base64 编码
	fmt.Println("\n--- Base64 编码 ---")
	data := []byte("Hello, Go!")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("原始数据: %s\n", data)
	fmt.Printf("Base64 编码: %s\n", encoded)

	// 4. Base64 解码
	fmt.Println("\n--- Base64 解码 ---")
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Printf("Base64 解码失败: %v\n", err)
	} else {
		fmt.Printf("解码结果: %s\n", decoded)
	}

	// 5. URL 安全的 Base64
	fmt.Println("\n--- URL 安全的 Base64 ---")
	urlEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Printf("URL 编码: %s\n", urlEncoded)

	urlDecoded, _ := base64.URLEncoding.DecodeString(urlEncoded)
	fmt.Printf("URL 解码: %s\n", urlDecoded)
}

