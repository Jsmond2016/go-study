# 并发编程示例

本示例展示了 Go 语言的并发编程，包括 goroutine、channel、select、WaitGroup 和 Mutex 等。

## 运行示例

```bash
go run main.go
```

## 知识点

### Goroutine

- 使用 `go` 关键字启动一个 goroutine
- Goroutine 是轻量级线程，由 Go 运行时管理
- 主程序退出时，所有 goroutine 也会退出

### Channel

- Channel 是 goroutine 之间通信的管道
- 无缓冲 channel：发送和接收必须同时准备好
- 有缓冲 channel：可以缓存一定数量的值
- 使用 `close()` 关闭 channel

### Select 语句

- 用于在多个 channel 操作中选择
- 类似于 switch，但用于 channel 操作
- 可以设置超时

### 同步原语

- **WaitGroup**：等待一组 goroutine 完成
- **Mutex**：互斥锁，保护共享资源
- **RWMutex**：读写锁，允许多个读操作

## 练习

1. 创建一个生产者-消费者模式，使用 channel 传递数据
2. 实现一个并发下载器，使用 WaitGroup 等待所有下载完成
3. 使用 Mutex 保护一个共享的计数器
4. 实现一个超时机制，使用 select 和 time.After

