---
title: 加密 (crypto)
difficulty: advanced
duration: "3-4小时"
prerequisites: ["基础语法", "错误处理"]
tags: ["crypto", "加密", "哈希", "签名", "安全"]
---

# 加密 (crypto)

`crypto` 包提供了加密相关的功能，包括哈希、加密、签名等。

## 📋 学习目标

- [ ] 掌握哈希函数的使用
- [ ] 理解加密和解密
- [ ] 学会数字签名
- [ ] 了解密码学基础
- [ ] 掌握安全最佳实践

## 🎯 哈希函数

### MD5 和 SHA

```go
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("Hello, Go!")

	// MD5
	md5Hash := md5.Sum(data)
	fmt.Printf("MD5: %x\n", md5Hash)

	// SHA1
	sha1Hash := sha1.Sum(data)
	fmt.Printf("SHA1: %x\n", sha1Hash)

	// SHA256
	sha256Hash := sha256.Sum256(data)
	fmt.Printf("SHA256: %x\n", sha256Hash)

	// SHA512
	sha512Hash := sha512.Sum512(data)
	fmt.Printf("SHA512: %x\n", sha512Hash)
}
```

### 使用 Hash 接口

```go
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func hashData(data []byte) []byte {
	h := sha256.New()
	io.WriteString(h, string(data))
	return h.Sum(nil)
}

func main() {
	data := []byte("Hello, Go!")
	hash := hashData(data)
	fmt.Printf("Hash: %x\n", hash)
}
```

## 🔐 密码哈希

### bcrypt

```go
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "myPassword123"
	hash, _ := hashPassword(password)

	fmt.Printf("密码: %s\n", password)
	fmt.Printf("哈希: %s\n", hash)
	fmt.Printf("验证: %t\n", checkPassword(password, hash))
}
```

## 🔒 对称加密

### AES 加密

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encryptAES(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
```

## 🔑 非对称加密

### RSA 加密

```go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func generateRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func encryptRSA(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
	return rsa.EncryptOAEP(
		nil,
		rand.Reader,
		publicKey,
		plaintext,
		nil,
	)
}

func decryptRSA(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(
		nil,
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
}
```

## ✍️ 数字签名

### RSA 签名

```go
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func signData(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
}

func verifySignature(publicKey *rsa.PublicKey, data []byte, signature []byte) bool {
	hash := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	return err == nil
}
```

## 🏃‍♂️ 实践应用

### 密码存储

```go
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
```

### 数据完整性验证

```go
func calculateFileHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
```

## ⚠️ 注意事项

### 1. 密码哈希

```go
// ✅ 使用 bcrypt 或 argon2 进行密码哈希
// ❌ 不要使用 MD5 或 SHA 进行密码哈希
```

### 2. 随机数生成

```go
// ✅ 使用 crypto/rand 生成加密安全的随机数
// ❌ 不要使用 math/rand
```

### 3. 密钥管理

```go
// ✅ 安全存储密钥
// ❌ 不要将密钥硬编码在代码中
```

## 📚 扩展阅读

- [crypto 包官方文档](https://pkg.go.dev/crypto)
- [密码学基础](https://en.wikipedia.org/wiki/Cryptography)

## ⏭️ 下一阶段

完成标准库学习后，可以进入：

- [Web 开发](../web-development/) - 构建 Web 应用
- [数据库](../database/) - 数据库操作

---

**💡 提示**: 加密是安全编程的基础，掌握加密技术对于构建安全的应用程序至关重要！

