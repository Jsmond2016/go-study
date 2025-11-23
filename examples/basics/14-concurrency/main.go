// Package main 展示 Go 语言的并发编程
package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine 示例
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello, %s! (%d)\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Channel 示例
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

// 带缓冲的 Channel
func bufferedChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("缓冲通道已发送 3 个值")
	close(ch)
}

// Select 语句
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自 ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自 ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Printf("收到: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("收到: %s\n", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("超时")
	}
}

// WaitGroup 示例
func waitGroupExample() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d 执行\n", id)
			time.Sleep(500 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有 Goroutine 完成")
}

// Mutex 示例
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

func mutexExample() {
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("计数器值: %d\n", counter.Value())
}

func main() {
	fmt.Println("=== Go 并发编程示例 ===\n")

	// 1. Goroutine
	fmt.Println("--- Goroutine ---")
	go sayHello("Alice")
	go sayHello("Bob")
	time.Sleep(1 * time.Second)

	// 2. Channel
	fmt.Println("\n--- Channel ---")
	ch := make(chan int)
	go sendNumbers(ch)
	for num := range ch {
		fmt.Printf("收到数字: %d\n", num)
	}

	// 3. 缓冲 Channel
	fmt.Println("\n--- 缓冲 Channel ---")
	bufferedChannel()

	// 4. Select
	fmt.Println("\n--- Select ---")
	selectExample()

	// 5. WaitGroup
	fmt.Println("\n--- WaitGroup ---")
	waitGroupExample()

	// 6. Mutex
	fmt.Println("\n--- Mutex ---")
	mutexExample()
}

