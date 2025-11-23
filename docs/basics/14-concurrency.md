---
title: 并发编程
difficulty: advanced
duration: "8-10小时"
prerequisites: ["函数", "接口", "错误处理"]
tags: ["并发", "goroutine", "channel", "同步", "并发模式"]
---

# 并发编程

Go 语言最强大的特性之一就是内置的并发支持。通过 goroutine 和 channel，Go 让并发编程变得简单而优雅。

## 📋 学习目标

- [ ] 理解并发和并行的概念
- [ ] 掌握 goroutine 的使用
- [ ] 理解 channel 的通信机制
- [ ] 掌握同步原语的使用
- [ ] 了解常见的并发模式
- [ ] 学会处理并发安全问题

## 🎯 并发基础

### 什么是并发

并发是同时处理多个任务的能力，Go 通过 goroutine 实现轻量级并发。

```go
package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello, %s!\n", name)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 顺序执行
	sayHello("Alice")
	sayHello("Bob")
	
	// 并发执行
	go sayHello("Charlie")
	go sayHello("David")
	
	// 等待 goroutine 完成
	time.Sleep(1 * time.Second)
}
```

## 🚀 Goroutine

### 创建 Goroutine

使用 `go` 关键字启动一个 goroutine：

```go
go function()
```

### 示例

```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("数字: %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("字母: %c\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	go printLetters()
	
	// 等待 goroutine 完成
	time.Sleep(2 * time.Second)
	fmt.Println("完成")
}
```

### Goroutine 的特点

- **轻量级**: 每个 goroutine 只占用几 KB 内存
- **快速启动**: 创建和销毁成本低
- **M:N 调度**: Go 运行时在多个 OS 线程上调度 goroutine

## 📡 Channel

Channel 是 goroutine 之间通信的管道。

### 创建 Channel

```go
ch := make(chan int)        // 无缓冲 channel
ch := make(chan int, 10)    // 有缓冲 channel
```

### 发送和接收

```go
ch <- value    // 发送
value := <-ch  // 接收
```

### 示例

```go
package main

import "fmt"

func sendData(ch chan<- string) {
	ch <- "Hello"
	ch <- "World"
	close(ch)
}

func receiveData(ch <-chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func main() {
	ch := make(chan string)
	
	go sendData(ch)
	receiveData(ch)
}
```

### 无缓冲 vs 有缓冲 Channel

```go
// 无缓冲 channel（同步）
ch1 := make(chan int)
// 发送会阻塞直到有接收者

// 有缓冲 channel（异步）
ch2 := make(chan int, 10)
// 可以发送 10 个值而不阻塞
```

## 🔒 同步原语

### sync.WaitGroup

等待一组 goroutine 完成：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("Worker %d 开始工作\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d 完成工作\n", id)
}

func main() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	
	wg.Wait()
	fmt.Println("所有工作完成")
}
```

### sync.Mutex

互斥锁，保护共享资源：

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	
	wg.Wait()
	fmt.Printf("最终值: %d\n", counter.Value())
}
```

### sync.RWMutex

读写锁，允许多个读或一个写：

```go
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func (m *SafeMap) Get(key string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data[key]
}

func (m *SafeMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}
```

## 🎭 并发模式

### 1. 生产者-消费者模式

```go
package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Printf("消费: %d\n", val)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 3)
	
	go producer(ch)
	consumer(ch)
}
```

### 2. 工作池模式

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d 处理任务 %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// 启动 3 个 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// 发送任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 收集结果
	for r := 1; r <= 5; r++ {
		fmt.Printf("结果: %d\n", <-results)
	}
}
```

### 3. 扇出扇入模式

```go
func fanOut(in <-chan int, out1, out2 chan<- int) {
	for val := range in {
		select {
		case out1 <- val:
		case out2 <- val:
		}
	}
	close(out1)
	close(out2)
}

func fanIn(input1, input2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case val, ok := <-input1:
				if !ok {
					input1 = nil
					continue
				}
				out <- val
			case val, ok := <-input2:
				if !ok {
					input2 = nil
					continue
				}
				out <- val
			}
			if input1 == nil && input2 == nil {
				break
			}
		}
	}()
	return out
}
```

## ⚠️ 常见陷阱

### 1. 竞态条件

```go
// ❌ 错误：竞态条件
var counter int

func increment() {
	counter++
}

// ✅ 正确：使用互斥锁
var counter int
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}
```

### 2. 死锁

```go
// ❌ 可能导致死锁
ch1 := make(chan int)
ch2 := make(chan int)

go func() {
	ch1 <- 1
	<-ch2
}()

<-ch1
ch2 <- 2
```

### 3. Goroutine 泄漏

```go
// ❌ 可能导致 goroutine 泄漏
func leak() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	// 忘记接收，goroutine 永远阻塞
}
```

## 🏃‍♂️ 实践练习

### 练习 1: 并发下载

实现一个并发下载多个文件的程序。

### 练习 2: 并发计算

使用 goroutine 并发计算斐波那契数列。

### 练习 3: 并发爬虫

实现一个简单的并发网页爬虫。

## 🤔 思考题

1. goroutine 和线程有什么区别？
2. 什么时候使用无缓冲 channel，什么时候使用有缓冲 channel？
3. 如何避免死锁？
4. 什么是竞态条件？如何检测？
5. context 包在并发编程中的作用是什么？

## 📚 扩展阅读

- [Go 并发模式](https://go.dev/blog/pipelines)
- [并发编程最佳实践](https://golang.org/doc/effective_go.html#concurrency)
- [Go 内存模型](https://golang.org/ref/mem)

## ⏭️ 下一章节

[反射](./15-reflection.md) → 学习 Go 语言的反射机制

---

**💡 提示**: 并发是 Go 语言的核心优势，掌握 goroutine 和 channel 是编写高效 Go 程序的关键！

