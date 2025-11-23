---
title: 上下文 (context)
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["并发编程", "错误处理"]
tags: ["context", "上下文", "超时", "取消", "并发"]
---

# 上下文 (context)

`context` 包提供了上下文管理功能，用于控制 goroutine 的生命周期、传递请求范围的值和取消信号。

## 📋 学习目标

- [ ] 理解 context 的作用和用途
- [ ] 掌握 context 的创建和使用
- [ ] 学会实现超时和取消
- [ ] 理解 context 在并发编程中的应用
- [ ] 掌握 context 的最佳实践

## 🎯 基本用法

### 创建 Context

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// 创建根 context
	ctx := context.Background()
	
	// 创建带值的 context
	ctx = context.WithValue(ctx, "userID", "123")
	
	// 获取值
	userID := ctx.Value("userID")
	fmt.Println(userID)
}
```

### 超时控制

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建带超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// 在 goroutine 中使用
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("操作超时或被取消")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("执行中...")
			}
		}
	}()
	
	time.Sleep(6 * time.Second)
}
```

### 取消操作

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(3 * time.Second)
		cancel() // 取消操作
	}()
	
	select {
	case <-ctx.Done():
		fmt.Println("操作被取消")
	case <-time.After(5 * time.Second):
		fmt.Println("操作完成")
	}
}
```

## 🔧 高级用法

### Context 传递

```go
package main

import (
	"context"
	"fmt"
)

func processRequest(ctx context.Context, userID string) {
	// 传递 context 到子函数
	ctx = context.WithValue(ctx, "userID", userID)
	
	processData(ctx)
}

func processData(ctx context.Context) {
	userID := ctx.Value("userID")
	fmt.Printf("处理用户 %v 的数据\n", userID)
	
	// 检查是否被取消
	select {
	case <-ctx.Done():
		fmt.Println("操作被取消")
		return
	default:
		// 继续处理
	}
}
```

### HTTP 请求中的 Context

```go
package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func makeRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	return io.ReadAll(resp.Body)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	data, err := makeRequest(ctx, "https://api.example.com/data")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(data))
}
```

## 🏃‍♂️ 实践应用

### 数据库查询超时

```go
func queryWithTimeout(ctx context.Context, db *sql.DB, query string) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	return db.QueryContext(ctx, query)
}
```

### 并发任务控制

```go
func processTasks(ctx context.Context, tasks []Task) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	errCh := make(chan error, len(tasks))
	
	for _, task := range tasks {
		go func(t Task) {
			if err := processTask(ctx, t); err != nil {
				errCh <- err
				cancel() // 取消其他任务
			}
		}(task)
	}
	
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
```

## ⚠️ 注意事项

### 1. Context 值

```go
// ✅ 使用类型安全的 key
type key string

const userIDKey key = "userID"

ctx := context.WithValue(ctx, userIDKey, "123")
userID := ctx.Value(userIDKey).(string)
```

### 2. Context 传递

```go
// ✅ 总是将 context 作为第一个参数
func process(ctx context.Context, data string) error {
	// ...
}
```

### 3. 不要存储 Context

```go
// ❌ 不要将 context 存储在结构体中
type Service struct {
	ctx context.Context // 错误
}

// ✅ 在函数参数中传递
func (s *Service) Process(ctx context.Context) error {
	// ...
}
```

## 📚 扩展阅读

- [context 包官方文档](https://pkg.go.dev/context)
- [Go Context 最佳实践](https://blog.golang.org/context)

## ⏭️ 下一章节

[编码解码 (encoding)](./09-encoding.md) → 学习 Go 语言的编码解码

---

**💡 提示**: Context 是 Go 并发编程的重要工具，掌握它对于编写健壮的并发程序至关重要！

