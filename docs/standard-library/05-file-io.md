---
title: 文件操作
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["变量与常量", "错误处理"]
tags: ["文件", "I/O", "读写", "目录"]
---

# 文件操作

Go 语言提供了丰富的文件操作功能，主要通过 `os`、`io`、`ioutil` 和 `bufio` 包实现。

## 📋 学习目标

- [ ] 掌握文件的基本读写操作
- [ ] 理解文件权限和模式
- [ ] 学会目录操作
- [ ] 掌握文件路径处理
- [ ] 了解文件监控
- [ ] 学会文件操作的最佳实践

## 🎯 文件读取

### 读取整个文件

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// 读取整个文件
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("文件内容:\n%s", data)
}
```

### 逐行读取

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
```

### 分块读取

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	buf := make([]byte, 1024) // 1KB 缓冲区
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(buf[:n]))
	}
}
```

## ✍️ 文件写入

### 写入整个文件

```go
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data := []byte("Hello, Go!")
	
	err := ioutil.WriteFile("output.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```

### 追加写入

```go
package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	_, err = file.WriteString("追加的内容\n")
	if err != nil {
		log.Fatal(err)
	}
}
```

### 使用 bufio 写入

```go
package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	
	writer.WriteString("第一行\n")
	writer.WriteString("第二行\n")
}
```

## 📁 目录操作

### 创建目录

```go
package main

import (
	"os"
)

func main() {
	// 创建单个目录
	err := os.Mkdir("mydir", 0755)
	if err != nil {
		log.Fatal(err)
	}
	
	// 创建多级目录
	err = os.MkdirAll("path/to/dir", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
```

### 读取目录

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	
	for _, entry := range entries {
		fmt.Println(entry.Name(), entry.IsDir())
	}
}
```

### 删除文件和目录

```go
package main

import (
	"os"
)

func main() {
	// 删除文件
	err := os.Remove("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	// 删除目录（必须为空）
	err = os.Remove("mydir")
	if err != nil {
		log.Fatal(err)
	}
	
	// 删除目录及其内容
	err = os.RemoveAll("path/to/dir")
	if err != nil {
		log.Fatal(err)
	}
}
```

## 🔍 文件信息

### 获取文件信息

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("文件名: %s\n", info.Name())
	fmt.Printf("大小: %d 字节\n", info.Size())
	fmt.Printf("模式: %v\n", info.Mode())
	fmt.Printf("修改时间: %v\n", info.ModTime())
	fmt.Printf("是否目录: %t\n", info.IsDir())
}
```

### 检查文件是否存在

```go
package main

import (
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func main() {
	if fileExists("file.txt") {
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
	}
}
```

## 🛤️ 文件路径

### 路径操作

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 路径拼接
	path := filepath.Join("dir", "subdir", "file.txt")
	fmt.Println(path) // dir/subdir/file.txt (Unix) 或 dir\subdir\file.txt (Windows)
	
	// 获取目录
	dir := filepath.Dir(path)
	fmt.Println(dir)
	
	// 获取文件名
	filename := filepath.Base(path)
	fmt.Println(filename)
	
	// 获取扩展名
	ext := filepath.Ext(path)
	fmt.Println(ext)
	
	// 获取绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(absPath)
}
```

### 路径匹配

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 路径匹配
	matched, err := filepath.Match("*.go", "main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(matched) // true
	
	// 文件路径匹配
	matched, err = filepath.Match("test/*.go", "test/main.go")
	fmt.Println(matched) // true
}
```

## 🔄 文件复制和移动

### 复制文件

```go
package main

import (
	"io"
	"os"
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()
	
	_, err = io.Copy(destFile, sourceFile)
	return err
}

func main() {
	err := copyFile("source.txt", "dest.txt")
	if err != nil {
		log.Fatal(err)
	}
}
```

### 移动文件

```go
package main

import (
	"os"
)

func main() {
	err := os.Rename("old.txt", "new.txt")
	if err != nil {
		log.Fatal(err)
	}
}
```

## 🏃‍♂️ 实践应用

### 文件工具函数

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 读取文件所有行
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// 写入文件所有行
func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	// 处理行
	for i, line := range lines {
		lines[i] = strings.ToUpper(line)
	}
	
	err = writeLines(lines, "output.txt")
	if err != nil {
		log.Fatal(err)
	}
}
```

## ⚠️ 注意事项

### 1. 文件权限

```go
// 文件权限模式
// 0644: 所有者读写，其他只读
// 0755: 所有者读写执行，其他读执行
// 0600: 仅所有者读写
```

### 2. 资源清理

```go
// ✅ 总是使用 defer 关闭文件
file, err := os.Open("file.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()
```

### 3. 错误处理

```go
// ✅ 总是检查错误
data, err := ioutil.ReadFile("file.txt")
if err != nil {
	// 处理错误
	return err
}
```

## 📚 扩展阅读

- [os 包官方文档](https://pkg.go.dev/os)
- [io 包官方文档](https://pkg.go.dev/io)
- [path/filepath 包官方文档](https://pkg.go.dev/path/filepath)

## ⏭️ 下一章节

[字符串转换 (strconv)](./06-strconv.md) → 学习 Go 语言的字符串转换

---

**💡 提示**: 文件操作是程序的基础功能，掌握文件 I/O 对于构建实用程序至关重要！

