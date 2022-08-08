package main

import (
	"fmt"
	"sync"
)

// 学习文档 https://www.liwenzhou.com/posts/Go/concurrence/

// 声明全局等待组变量
var wg sync.WaitGroup

// func hello() {
// 	fmt.Println("hello")
// 	wg.Done() // 告知当前goroutine完成
// }

// func main() {
// 	wg.Add(1) // 登记1个goroutine
// 	go hello()
// 	fmt.Println("你好")
// 	wg.Wait() // 阻塞等待登记的goroutine完成
// }

// 启动多个 goroutine

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
