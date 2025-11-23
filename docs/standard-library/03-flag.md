---
title: 命令行参数 (flag)
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["flag", "命令行", "参数解析", "CLI"]
---

# 命令行参数 (flag)

`flag` 包提供了命令行参数解析功能，是构建命令行工具的基础。

## 📋 学习目标

- [ ] 掌握 flag 包的基本用法
- [ ] 理解不同类型的标志定义
- [ ] 学会使用 FlagSet
- [ ] 了解自定义标志类型
- [ ] 掌握实际应用场景

## 🎯 基本用法

### 定义标志

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义标志
	name := flag.String("name", "Guest", "用户姓名")
	age := flag.Int("age", 0, "用户年龄")
	married := flag.Bool("married", false, "是否已婚")

	// 解析命令行参数
	flag.Parse()

	// 使用标志值（注意：返回的是指针）
	fmt.Printf("姓名: %s\n", *name)
	fmt.Printf("年龄: %d\n", *age)
	fmt.Printf("已婚: %t\n", *married)
}
```

**运行示例**：
```bash
go run main.go -name "张三" -age 30 -married
```

### 标志类型

```go
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 字符串标志
	name := flag.String("name", "default", "字符串标志")

	// 整数标志
	age := flag.Int("age", 0, "整数标志")

	// 布尔标志
	verbose := flag.Bool("verbose", false, "详细输出")

	// 浮点数标志
	height := flag.Float64("height", 0.0, "身高（米）")

	// 持续时间标志
	timeout := flag.Duration("timeout", 30*time.Second, "超时时间")

	flag.Parse()

	fmt.Printf("姓名: %s\n", *name)
	fmt.Printf("年龄: %d\n", *age)
	fmt.Printf("详细输出: %t\n", *verbose)
	fmt.Printf("身高: %.2f米\n", *height)
	fmt.Printf("超时: %v\n", *timeout)
}
```

## 🔧 高级用法

### 使用变量存储标志

```go
package main

import (
	"flag"
	"fmt"
)

var (
	name    = flag.String("name", "Guest", "用户姓名")
	age     = flag.Int("age", 0, "用户年龄")
	married = flag.Bool("married", false, "是否已婚")
)

func main() {
	flag.Parse()

	fmt.Printf("姓名: %s\n", *name)
	fmt.Printf("年龄: %d\n", *age)
	fmt.Printf("已婚: %t\n", *married)
}
```

### 获取非选项参数

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	// 获取非选项参数
	args := flag.Args()
	fmt.Printf("非选项参数: %v\n", args)
	fmt.Printf("参数数量: %d\n", flag.NArg())

	// 访问特定参数
	if flag.NArg() > 0 {
		fmt.Printf("第一个参数: %s\n", flag.Arg(0))
	}
}
```

**运行示例**：
```bash
go run main.go -name "张三" file1.txt file2.txt
```

### 自定义帮助信息

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 自定义帮助信息
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: %s [选项]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "这是一个演示 flag 包的示例程序\n\n")
		fmt.Fprintf(os.Stderr, "选项:\n")
		flag.PrintDefaults()
	}

	name := flag.String("name", "Guest", "用户姓名")
	flag.Parse()

	fmt.Printf("姓名: %s\n", *name)
}
```

## 📦 FlagSet

FlagSet 允许创建独立的标志集合，适用于子命令场景。

### 基本用法

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 创建新的标志集
	fs := flag.NewFlagSet("demo", flag.ExitOnError)

	// 在标志集中定义标志
	name := fs.String("name", "default", "名称")
	age := fs.Int("age", 0, "年龄")

	// 解析参数
	args := []string{"-name", "张三", "-age", "30"}
	err := fs.Parse(args)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}

	fmt.Printf("姓名: %s\n", *name)
	fmt.Printf("年龄: %d\n", *age)
}
```

### 错误处理模式

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// ExitOnError: 遇到错误时调用 os.Exit(2)
	fs1 := flag.NewFlagSet("exit", flag.ExitOnError)

	// ContinueOnError: 遇到错误时返回错误，不退出
	fs2 := flag.NewFlagSet("continue", flag.ContinueOnError)

	// PanicOnError: 遇到错误时 panic
	fs3 := flag.NewFlagSet("panic", flag.PanicOnError)

	_ = fs1
	_ = fs2
	_ = fs3
}
```

### 子命令示例

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: program <command> [选项]")
		fmt.Println("命令: user, admin")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "user":
		handleUserCommand(os.Args[2:])
	case "admin":
		handleAdminCommand(os.Args[2:])
	default:
		fmt.Printf("未知命令: %s\n", command)
		os.Exit(1)
	}
}

func handleUserCommand(args []string) {
	fs := flag.NewFlagSet("user", flag.ExitOnError)
	username := fs.String("username", "", "用户名")
	password := fs.String("password", "", "密码")

	fs.Parse(args)

	fmt.Printf("用户登录: %s\n", *username)
}

func handleAdminCommand(args []string) {
	fs := flag.NewFlagSet("admin", flag.ExitOnError)
	admin := fs.String("admin", "", "管理员用户名")
	level := fs.Int("level", 1, "权限级别")

	fs.Parse(args)

	fmt.Printf("管理员: %s, 级别: %d\n", *admin, *level)
}
```

## 🎨 自定义标志类型

实现 `flag.Value` 接口可以创建自定义标志类型。

```go
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 自定义类型：字符串列表
type stringList []string

func (s *stringList) String() string {
	return strings.Join(*s, ",")
}

func (s *stringList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var tags stringList
	flag.Var(&tags, "tag", "标签（可多次使用）")

	flag.Parse()

	fmt.Printf("标签: %v\n", tags)
}
```

**运行示例**：
```bash
go run main.go -tag "go" -tag "programming" -tag "tutorial"
```

## 🏃‍♂️ 实践应用

### 服务器配置

```go
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		host     = flag.String("host", "0.0.0.0", "服务器地址")
		port     = flag.Int("port", 8080, "服务器端口")
		https    = flag.Bool("https", false, "启用HTTPS")
		timeout  = flag.Duration("timeout", 30*time.Second, "请求超时")
	)

	flag.Parse()

	scheme := "http"
	if *https {
		scheme = "https"
	}

	fmt.Printf("服务器地址: %s://%s:%d\n", scheme, *host, *port)
	fmt.Printf("超时时间: %v\n", *timeout)
}
```

### 数据库配置

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		driver   = flag.String("driver", "mysql", "数据库驱动")
		host     = flag.String("db-host", "localhost", "数据库主机")
		port     = flag.Int("db-port", 3306, "数据库端口")
		database = flag.String("db-name", "myapp", "数据库名称")
		username = flag.String("db-user", "root", "数据库用户名")
		password = flag.String("db-pass", "", "数据库密码")
	)

	flag.Parse()

	fmt.Printf("数据库: %s@%s:%d/%s\n", *username, *host, *port, *database)
}
```

### 配置文件加载器

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		configFile = flag.String("config", "config.json", "配置文件路径")
		env        = flag.String("env", "development", "运行环境")
		debug      = flag.Bool("debug", false, "调试模式")
		logLevel   = flag.String("log-level", "info", "日志级别")
	)

	flag.Parse()

	fmt.Printf("配置文件: %s\n", *configFile)
	fmt.Printf("运行环境: %s\n", *env)
	fmt.Printf("调试模式: %t\n", *debug)
	fmt.Printf("日志级别: %s\n", *logLevel)
}
```

## ⚠️ 注意事项

### 1. 标志值是指针

```go
// ❌ 错误
name := flag.String("name", "default", "名称")
fmt.Println(name)  // 输出指针地址

// ✅ 正确
name := flag.String("name", "default", "名称")
fmt.Println(*name)  // 输出值
```

### 2. 必须调用 Parse

```go
// ❌ 错误：未调用 Parse
name := flag.String("name", "default", "名称")
fmt.Println(*name)  // 总是输出默认值

// ✅ 正确
name := flag.String("name", "default", "名称")
flag.Parse()  // 必须调用
fmt.Println(*name)
```

### 3. 标志定义顺序

标志定义必须在 `Parse()` 之前。

```go
// ❌ 错误
flag.Parse()
name := flag.String("name", "default", "名称")

// ✅ 正确
name := flag.String("name", "default", "名称")
flag.Parse()
```

## 🏃‍♂️ 实践练习

### 练习 1: 创建文件操作工具

使用 flag 包创建一个文件操作工具，支持复制、移动、删除等操作。

### 练习 2: 实现配置管理器

创建一个配置管理器，支持从命令行参数和配置文件加载配置。

### 练习 3: 构建 CLI 工具

创建一个完整的 CLI 工具，包含多个子命令。

## 🤔 思考题

1. flag 包和 os.Args 有什么区别？
2. 什么时候应该使用 FlagSet？
3. 如何实现标志的验证？
4. 如何处理标志之间的依赖关系？

## 📚 扩展阅读

- [flag 包官方文档](https://pkg.go.dev/flag)
- [命令行工具最佳实践](https://github.com/golang/go/wiki/CodeReviewComments#flag-names)

## ⏭️ 下一章节

[日志 (log)](./04-log.md) → 学习 Go 语言的日志记录

---

**💡 提示**: flag 包是构建命令行工具的基础，掌握它可以让你的程序更加灵活和易用！

