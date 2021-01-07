package main

// https://www.liwenzhou.com/posts/Go/14_concurrence/

import (
	"fmt"
	"time"
)

func main() {
	go testgo() //使用关键字go调用函数或者方法 开启一个goroutine
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	time.Sleep(3000 * time.Millisecond) //加上休眠让主程序休眠3秒钟。
	fmt.Println("main 函数结束")

}

//自定义函数
func testgo() {
	for i := 0; i < 10; i++ {
		fmt.Println("测试goroutine", i)
	}
}
