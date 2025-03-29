# golang 中的循环有哪些常用操作

我来介绍一下 Go 语言中常用的循环操作方法。

## 一、基本 for 循环

```go
// 1. 标准 for 循环
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 2. 类似 while 的 for 循环
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// 3. 无限循环
for {
    fmt.Println("无限循环")
    // 使用 break 跳出循环
    break
}
```

## 二、range 循环

```go
// 1. 遍历切片
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("索引:%d 值:%d\n", index, value)
}

// 2. 遍历字符串
str := "Hello, 世界"
for i, char := range str {
    fmt.Printf("位置:%d 字符:%c\n", i, char)
}

// 3. 遍历 map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("键:%s 值:%d\n", key, value)
}

// 4. 遍历通道
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    close(ch)
}()
for num := range ch {
    fmt.Println(num)
}
```

## 三、循环控制

```go
// 1. break 语句
for i := 0; i < 5; i++ {
    if i == 3 {
        break // 跳出循环
    }
    fmt.Println(i)
}

// 2. continue 语句
for i := 0; i < 5; i++ {
    if i == 2 {
        continue // 跳过本次循环
    }
    fmt.Println(i)
}

// 3. 标签跳转
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i*j == 4 {
                break outer // 跳出外层循环
            }
        }
    }
```

## 四、常用循环模式

```go
// 1. 倒计时循环
for i := 10; i > 0; i-- {
    fmt.Println(i)
}

// 2. 步进循环
for i := 0; i < 10; i += 2 {
    fmt.Println(i) // 输出偶数
}

// 3. 多变量循环
for i, j := 0, 10; i < j; i, j = i+1, j-1 {
    fmt.Printf("i=%d, j=%d\n", i, j)
}
```

## 五、并发循环

```go
// 1. 并发遍历切片
items := []int{1, 2, 3, 4, 5}
var wg sync.WaitGroup
for _, item := range items {
    wg.Add(1)
    go func(x int) {
        defer wg.Done()
        // 处理 x
        fmt.Println(x)
    }(item)
}
wg.Wait()

// 2. 带超时的循环
timeout := time.After(1 * time.Second)
for {
    select {
    case <-timeout:
        fmt.Println("超时退出")
        return
    default:
        // 执行循环体
    }
}
```

## 六、实用技巧

1. **性能优化**：
```go
// 预先计算长度，避免每次循环都计算
length := len(slice)
for i := 0; i < length; i++ {
    // 使用 slice[i]
}
```

2. **避免在循环中分配内存**：
```go
// 不好的做法
for i := 0; i < 1000; i++ {
    slice := make([]int, 100)
    // 使用 slice
}

// 好的做法
slice := make([]int, 100)
for i := 0; i < 1000; i++ {
    // 重用 slice
}
```

## 注意事项

1. 在循环中修改切片时要特别小心，可能影响循环行为
2. range 循环会复制元素，对于大对象使用索引访问更高效
3. 并发循环要注意变量捕获问题，建议通过参数传递
4. 避免在循环中进行频繁的内存分配
5. 合理使用 break 和 continue 来优化循环流程

这个文档涵盖了 Go 语言中最常用的循环操作方法。建议根据实际场景选择合适的循环方式，特别是在性能要求较高的场景下。