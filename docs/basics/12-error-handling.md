---
title: 错误处理
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["函数", "接口"]
tags: ["错误处理", "error", "panic", "recover"]
---

# 错误处理

Go 语言采用显式的错误处理机制，通过返回 `error` 值来表示操作是否成功。这是 Go 语言设计哲学的重要体现。

## 📋 学习目标

- [ ] 理解 Go 语言的错误处理哲学
- [ ] 掌握 error 接口的使用
- [ ] 学会创建和包装错误
- [ ] 理解 panic 和 recover 机制
- [ ] 掌握错误处理的最佳实践
- [ ] 了解错误处理的常见模式

## 🎯 错误处理基础

### error 接口

Go 语言中的错误是一个接口：

```go
type error interface {
	Error() string
}
```

### 基本用法

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("结果: %.2f\n", result)
	
	// 尝试除以零
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
}
```

## 🔍 创建错误

### 使用 errors.New()

```go
import "errors"

err := errors.New("这是一个错误")
```

### 使用 fmt.Errorf()

```go
import "fmt"

name := "test"
err := fmt.Errorf("文件 %s 不存在", name)
```

### 自定义错误类型

```go
package main

import "fmt"

type FileNotFoundError struct {
	FileName string
	Path     string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("文件 %s 在路径 %s 中未找到", e.FileName, e.Path)
}

func openFile(name, path string) error {
	// 模拟文件不存在
	return &FileNotFoundError{
		FileName: name,
		Path:     path,
	}
}

func main() {
	err := openFile("test.txt", "/tmp")
	if err != nil {
		fmt.Println(err)
		
		// 类型断言获取详细信息
		if fileErr, ok := err.(*FileNotFoundError); ok {
			fmt.Printf("文件名: %s\n", fileErr.FileName)
			fmt.Printf("路径: %s\n", fileErr.Path)
		}
	}
}
```

## 🔗 错误包装

Go 1.13+ 引入了错误包装机制，使用 `fmt.Errorf()` 和 `%w` 动词：

```go
package main

import (
	"errors"
	"fmt"
)

func readFile(filename string) error {
	// 模拟底层错误
	baseErr := errors.New("文件系统错误")
	
	// 包装错误
	return fmt.Errorf("读取文件 %s 失败: %w", filename, baseErr)
}

func main() {
	err := readFile("test.txt")
	if err != nil {
		fmt.Println(err)
		
		// 解包错误
		baseErr := errors.Unwrap(err)
		if baseErr != nil {
			fmt.Printf("底层错误: %v\n", baseErr)
		}
		
		// 检查错误链
		if errors.Is(err, baseErr) {
			fmt.Println("错误链中包含底层错误")
		}
	}
}
```

### errors.Is() 和 errors.As()

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("未找到")

func findItem(id int) error {
	if id < 0 {
		return fmt.Errorf("查找失败: %w", ErrNotFound)
	}
	return nil
}

func main() {
	err := findItem(-1)
	
	// 检查错误链中是否包含特定错误
	if errors.Is(err, ErrNotFound) {
		fmt.Println("确实是未找到错误")
	}
	
	// 类型断言（支持错误链）
	var notFoundErr *FileNotFoundError
	if errors.As(err, &notFoundErr) {
		fmt.Printf("文件错误: %v\n", notFoundErr)
	}
}
```

## ⚠️ Panic 和 Recover

### Panic

`panic` 用于处理程序无法恢复的错误：

```go
package main

import "fmt"

func mustDivide(a, b int) int {
	if b == 0 {
		panic("除数不能为零")
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 panic: %v\n", r)
		}
	}()
	
	result := mustDivide(10, 0)
	fmt.Printf("结果: %d\n", result)
}
```

### Recover

`recover` 用于捕获 `panic`，只能在 `defer` 函数中使用：

```go
package main

import "fmt"

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("发生 panic: %v", r)
		}
	}()
	
	if b == 0 {
		panic("除数不能为零")
	}
	
	result = a / b
	return
}

func main() {
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}
}
```

## 📝 错误处理模式

### 模式 1: 立即返回错误

```go
func process() error {
	if err := step1(); err != nil {
		return err
	}
	if err := step2(); err != nil {
		return err
	}
	return nil
}
```

### 模式 2: 错误包装

```go
func process() error {
	if err := step1(); err != nil {
		return fmt.Errorf("步骤1失败: %w", err)
	}
	if err := step2(); err != nil {
		return fmt.Errorf("步骤2失败: %w", err)
	}
	return nil
}
```

### 模式 3: 错误日志记录

```go
import "log"

func process() error {
	if err := step1(); err != nil {
		log.Printf("步骤1失败: %v", err)
		return err
	}
	return nil
}
```

### 模式 4: 错误重试

```go
func processWithRetry(maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		if err := process(); err != nil {
			if i == maxRetries-1 {
				return fmt.Errorf("重试 %d 次后仍然失败: %w", maxRetries, err)
			}
			continue
		}
		return nil
	}
	return nil
}
```

## 🎯 最佳实践

### 1. 总是检查错误

```go
// ❌ 错误做法
result, _ := divide(10, 2)

// ✅ 正确做法
result, err := divide(10, 2)
if err != nil {
	return err
}
```

### 2. 提供有意义的错误信息

```go
// ❌ 错误做法
return errors.New("错误")

// ✅ 正确做法
return fmt.Errorf("读取文件 %s 失败: %v", filename, err)
```

### 3. 使用错误变量

```go
var (
	ErrNotFound    = errors.New("未找到")
	ErrInvalidData = errors.New("无效数据")
)

func find(id int) error {
	if id < 0 {
		return ErrInvalidData
	}
	// ...
	return ErrNotFound
}
```

### 4. 不要忽略错误

```go
// ❌ 错误做法
_ = process()

// ✅ 正确做法
if err := process(); err != nil {
	log.Printf("处理失败: %v", err)
}
```

### 5. 合理使用 panic

`panic` 应该只用于：
- 程序无法恢复的错误
- 编程错误（如空指针解引用）
- 不应该用于正常的错误处理

## 🏃‍♂️ 实践练习

### 练习 1: 文件操作错误处理

实现一个文件读取函数，包含完整的错误处理。

### 练习 2: 网络请求错误处理

实现一个 HTTP 请求函数，处理各种网络错误。

### 练习 3: 数据库操作错误处理

实现一个数据库查询函数，处理连接错误和查询错误。

## 🤔 思考题

1. error 和 panic 的区别是什么？
2. 什么时候应该使用 panic？
3. 错误包装有什么好处？
4. 如何设计一个好的错误类型？
5. 错误处理对程序性能有什么影响？

## 📚 扩展阅读

- [Go 错误处理最佳实践](https://golang.org/doc/effective_go.html#errors)
- [错误处理设计](https://blog.golang.org/error-handling-and-go)
- [错误包装详解](https://go.dev/blog/go1.13-errors)

## ⏭️ 下一章节

[包管理](./13-packages.md) → 学习 Go 语言的包管理和模块系统

---

**💡 提示**: 错误处理是 Go 语言的核心特性，良好的错误处理可以让程序更加健壮和可维护！

