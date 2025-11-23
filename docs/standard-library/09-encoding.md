---
title: 编码解码 (encoding)
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["基础语法", "数据结构"]
tags: ["encoding", "JSON", "XML", "Base64", "序列化"]
---

# 编码解码 (encoding)

Go 语言提供了丰富的编码解码功能，包括 JSON、XML、Base64 等格式的处理。

## 📋 学习目标

- [ ] 掌握 JSON 编码和解码
- [ ] 理解 XML 处理
- [ ] 学会 Base64 编码
- [ ] 了解其他编码格式
- [ ] 掌握编码的最佳实践

## 🎯 JSON 编码解码

### 基本用法

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 编码（序列化）
	p := Person{Name: "张三", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
	// 输出: {"name":"张三","age":30}
	
	// 解码（反序列化）
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p2)
}
```

### 流式处理

```go
package main

import (
	"encoding/json"
	"os"
)

func main() {
	// 编码到流
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(Person{Name: "张三", Age: 30})
	
	// 从流解码
	decoder := json.NewDecoder(os.Stdin)
	var p Person
	decoder.Decode(&p)
}
```

### 处理嵌套结构

```go
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func main() {
	user := User{
		Name: "张三",
		Address: Address{
			City:    "北京",
			Country: "中国",
		},
	}
	
	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))
}
```

## 📄 XML 编码解码

### 基本用法

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	// 编码
	p := Person{Name: "张三", Age: 30}
	xmlData, err := xml.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlData))
	
	// 解码
	var p2 Person
	xml.Unmarshal(xmlData, &p2)
	fmt.Printf("%+v\n", p2)
}
```

## 🔤 Base64 编码

### 基本用法

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 编码
	data := []byte("Hello, Go!")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)
	
	// 解码
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decoded))
}
```

### URL 安全的 Base64

```go
func main() {
	data := []byte("Hello, Go!")
	
	// URL 安全编码
	encoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println(encoded)
	
	// URL 安全解码
	decoded, _ := base64.URLEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
}
```

## 🏃‍♂️ 实践应用

### JSON 配置文件

```go
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	var config Config
	err = json.Unmarshal(data, &config)
	return &config, err
}

func saveConfig(config *Config, filename string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
```

### API 响应处理

```go
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func sendJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	response := APIResponse{
		Code:    code,
		Message: "success",
		Data:    data,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```

## ⚠️ 注意事项

### 1. JSON 标签

```go
// ✅ 使用 json 标签控制序列化
type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`              // 不序列化
	Email    string `json:"email,omitempty"` // 空值时不序列化
}
```

### 2. 错误处理

```go
// ✅ 总是检查错误
data, err := json.Marshal(obj)
if err != nil {
	return err
}
```

### 3. 性能考虑

```go
// ✅ 对于大量数据，使用流式处理
encoder := json.NewEncoder(writer)
for _, item := range items {
	encoder.Encode(item)
}
```

## 📚 扩展阅读

- [encoding/json 包官方文档](https://pkg.go.dev/encoding/json)
- [encoding/xml 包官方文档](https://pkg.go.dev/encoding/xml)
- [encoding/base64 包官方文档](https://pkg.go.dev/encoding/base64)

## ⏭️ 下一章节

[加密 (crypto)](./10-crypto.md) → 学习 Go 语言的加密功能

---

**💡 提示**: 编码解码是数据交换的基础，掌握 JSON、XML 等格式对于构建 API 和配置文件处理非常重要！

