// Package main 展示 Go 语言的加密解密
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("=== Go 加密示例 ===\n")

	data := []byte("Hello, Go!")

	// 1. MD5 哈希
	fmt.Println("--- MD5 哈希 ---")
	md5Hash := md5.Sum(data)
	fmt.Printf("MD5: %x\n", md5Hash)

	// 2. SHA1 哈希
	fmt.Println("\n--- SHA1 哈希 ---")
	sha1Hash := sha1.Sum(data)
	fmt.Printf("SHA1: %x\n", sha1Hash)

	// 3. SHA256 哈希
	fmt.Println("\n--- SHA256 哈希 ---")
	sha256Hash := sha256.Sum256(data)
	fmt.Printf("SHA256: %x\n", sha256Hash)

	// 4. 使用 Hash 接口
	fmt.Println("\n--- 使用 Hash 接口 ---")
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	fmt.Printf("SHA256 (接口): %x\n", hash)
}

