---
title: 日志 (log)
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "函数"]
tags: ["log", "日志", "日志记录", "调试"]
---

# 日志 (log)

`log` 包提供了简单的日志记录功能，适合基础日志需求。

## 📋 学习目标

- [ ] 掌握 log 包的基本用法
- [ ] 理解日志级别和格式
- [ ] 学会自定义日志输出
- [ ] 了解日志文件操作
- [ ] 掌握日志的最佳实践

## 🎯 基本用法

### 简单日志输出

```go
package main

import (
	"log"
)

func main() {
	// 基本日志输出
	log.Print("这是一条普通日志")
	log.Println("这是一条带换行的日志")
	log.Printf("这是一条格式化日志: %s", "值")
}
```

### 日志级别

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Fatal 系列：输出日志后调用 os.Exit(1)
	log.Fatal("致命错误，程序退出")
	log.Fatalf("致命错误: %s", "详情")
	log.Fatalln("致命错误，程序退出")
	
	// Panic 系列：输出日志后调用 panic
	log.Panic("程序崩溃")
	log.Panicf("程序崩溃: %s", "详情")
	log.Panicln("程序崩溃")
}
```

## 🔧 自定义日志

### 设置日志前缀

```go
package main

import (
	"log"
)

func main() {
	// 设置日志前缀
	log.SetPrefix("[APP] ")
	
	log.Println("这条日志有前缀")
	// 输出: [APP] 2024/01/15 14:30:00 这条日志有前缀
}
```

### 设置日志标志

```go
package main

import (
	"log"
)

func main() {
	// 设置日志标志
	// Ldate: 日期 (2009/01/23)
	// Ltime: 时间 (01:23:23)
	// Lmicroseconds: 微秒
	// Llongfile: 完整文件路径和行号
	// Lshortfile: 文件名和行号
	// LstdFlags: Ldate | Ltime (默认)
	
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("自定义格式的日志")
}
```

### 设置日志输出

```go
package main

import (
	"log"
	"os"
)

func main() {
	// 输出到文件
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开日志文件:", err)
	}
	defer file.Close()
	
	// 设置日志输出到文件
	log.SetOutput(file)
	log.Println("这条日志写入文件")
	
	// 同时输出到文件和控制台
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.Println("这条日志同时输出到控制台和文件")
}
```

## 📝 创建自定义 Logger

```go
package main

import (
	"log"
	"os"
)

func main() {
	// 创建自定义 Logger
	logger := log.New(os.Stdout, "[CUSTOM] ", log.Ldate|log.Ltime|log.Lshortfile)
	
	logger.Println("自定义 Logger 的日志")
	logger.Printf("格式化日志: %s", "值")
}
```

### 多 Logger 示例

```go
package main

import (
	"log"
	"os"
)

func main() {
	// 创建不同级别的 Logger
	infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	warnLogger := log.New(os.Stdout, "[WARN] ", log.LstdFlags)
	errorLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	
	infoLogger.Println("信息日志")
	warnLogger.Println("警告日志")
	errorLogger.Println("错误日志")
}
```

## 🏃‍♂️ 实践应用

### 日志工具函数

```go
package main

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	WarnLogger = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("应用启动")
	WarnLogger.Println("这是一个警告")
	ErrorLogger.Println("这是一个错误")
}
```

### 日志文件轮转

```go
package main

import (
	"log"
	"os"
	"time"
)

func setupLogging() {
	// 创建日志目录
	os.MkdirAll("logs", 0755)
	
	// 使用日期作为日志文件名
	logFile := "logs/app-" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法创建日志文件:", err)
	}
	
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	setupLogging()
	log.Println("应用启动")
}
```

## ⚠️ 注意事项

### 1. Fatal 和 Panic 的区别

```go
// Fatal: 调用 os.Exit(1)，defer 不会执行
log.Fatal("致命错误")

// Panic: 调用 panic，defer 会执行
log.Panic("程序崩溃")
```

### 2. 日志性能

```go
// ❌ 避免：即使日志级别关闭也会执行字符串格式化
log.Printf("值: %v", expensiveOperation())

// ✅ 正确：先检查日志级别
if debugMode {
	log.Printf("值: %v", expensiveOperation())
}
```

## 📚 扩展阅读

- [log 包官方文档](https://pkg.go.dev/log)
- [第三方日志库推荐](https://github.com/sirupsen/logrus)

## ⏭️ 下一章节

[文件操作](./05-file-io.md) → 学习 Go 语言的文件 I/O 操作

---

**💡 提示**: 对于生产环境，建议使用更强大的日志库如 logrus 或 zap，它们提供更多功能如日志级别、结构化日志等。

