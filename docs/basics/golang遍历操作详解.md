# Golang 遍历操作详解

## 概述

Go 语言提供了多种遍历数据结构的方式，主要通过 `for` 循环和 `for...range` 语法来实现。本文将详细介绍 Go 语言中各种数据结构的遍历方法，特别关注字符和数组的遍历。

## 1. 基础遍历语法

### 1.1 传统 for 循环
```go
// 传统索引遍历
for i := 0; i < len(slice); i++ {
    fmt.Printf("索引: %d, 值: %v\n", i, slice[i])
}
```

### 1.2 for...range 循环
```go
// range 遍历（推荐）
for index, value := range collection {
    fmt.Printf("索引: %d, 值: %v\n", index, value)
}

// 只需要值
for _, value := range collection {
    fmt.Printf("值: %v\n", value)
}

// 只需要索引
for index := range collection {
    fmt.Printf("索引: %d\n", index)
}
```

## 2. 字符串遍历

### 2.1 按字节遍历（ASCII 字符）
```go
str := "Hello"

// 传统方式（按字节）
for i := 0; i < len(str); i++ {
    fmt.Printf("索引: %d, 字符: %c\n", i, str[i])
}

// 输出结果：
// 索引: 0, 字符: H
// 索引: 1, 字符: e
// 索引: 2, 字符: l
// 索引: 3, 字符: l
// 索引: 4, 字符: o

// range 方式（按字节）
for i, b := range []byte(str) {
    fmt.Printf("索引: %d, 字节: %c\n", i, b)
}

// 输出结果：
// 索引: 0, 字节: H
// 索引: 1, 字节: e
// 索引: 2, 字节: l
// 索引: 3, 字节: l
// 索引: 4, 字节: o
```

### 2.2 按字符遍历（Unicode）
```go
str := "你好，世界"

// range 方式（推荐，支持 Unicode）
for i, char := range str {
    fmt.Printf("索引: %d, 字符: %c, Unicode: %U\n", i, char, char)
}

// 输出结果：
// 索引: 0, 字符: 你, Unicode: U+4F60
// 索引: 3, 字符: 好, Unicode: U+597D
// 索引: 6, 字符: ，, Unicode: U+FF0C
// 索引: 9, 字符: 世, Unicode: U+4E16
// 索引: 12, 字符: 界, Unicode: U+754C

// 只获取字符
for _, char := range str {
    fmt.Printf("字符: %c\n", char)
}

// 输出结果：
// 字符: 你
// 字符: 好
// 字符: ，
// 字符: 世
// 字符: 界
```

### 2.3 字符串遍历示例
```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // ASCII 字符串
    asciiStr := "Hello Go"
    fmt.Println("ASCII 字符串遍历:")
    for i, char := range asciiStr {
        fmt.Printf("位置: %d, 字符: %c\n", i, char)
    }
    
    // 输出结果：
    // ASCII 字符串遍历:
    // 位置: 0, 字符: H
    // 位置: 1, 字符: e
    // 位置: 2, 字符: l
    // 位置: 3, 字符: l
    // 位置: 4, 字符: o
    // 位置: 5, 字符:  
    // 位置: 6, 字符: G
    // 位置: 7, 字符: o
    
    // Unicode 字符串
    unicodeStr := "你好，Go语言"
    fmt.Println("\nUnicode 字符串遍历:")
    for i, char := range unicodeStr {
        fmt.Printf("位置: %d, 字符: %c, 字节数: %d\n", 
            i, char, utf8.RuneLen(char))
    }
    
    // 输出结果：
    // Unicode 字符串遍历:
    // 位置: 0, 字符: 你, 字节数: 3
    // 位置: 3, 字符: 好, 字节数: 3
    // 位置: 6, 字符: ，, 字节数: 3
    // 位置: 9, 字符: G, 字节数: 1
    // 位置: 10, 字符: o, 字节数: 1
    // 位置: 11, 字符: 语, 字节数: 3
    // 位置: 14, 字符: 言, 字节数: 3
    
    // 获取字符串长度
    fmt.Printf("\n字节长度: %d\n", len(unicodeStr))
    fmt.Printf("字符数量: %d\n", utf8.RuneCountInString(unicodeStr))
    
    // 输出结果：
    // 字节长度: 17
    // 字符数量: 7
}
```

## 3. 数组遍历

### 3.1 固定长度数组
```go
// 声明数组
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// 传统 for 循环遍历
fmt.Println("传统 for 循环:")
for i := 0; i < len(arr); i++ {
    fmt.Printf("arr[%d] = %d\n", i, arr[i])
}

// 输出结果：
// 传统 for 循环:
// arr[0] = 1
// arr[1] = 2
// arr[2] = 3
// arr[3] = 4
// arr[4] = 5

// range 遍历（推荐）
fmt.Println("\nrange 遍历:")
for index, value := range arr {
    fmt.Printf("arr[%d] = %d\n", index, value)
}

// 输出结果：
// range 遍历:
// arr[0] = 1
// arr[1] = 2
// arr[2] = 3
// arr[3] = 4
// arr[4] = 5
```

### 3.2 多维数组遍历
```go
// 二维数组
var matrix [3][3]int = [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// 遍历二维数组
for i, row := range matrix {
    for j, value := range row {
        fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
    }
}

// 输出结果：
// matrix[0][0] = 1
// matrix[0][1] = 2
// matrix[0][2] = 3
// matrix[1][0] = 4
// matrix[1][1] = 5
// matrix[1][2] = 6
// matrix[2][0] = 7
// matrix[2][1] = 8
// matrix[2][2] = 9
```

## 4. 切片遍历

### 4.1 基本切片遍历
```go
slice := []int{10, 20, 30, 40, 50}

// range 遍历（最常用）
for index, value := range slice {
    fmt.Printf("slice[%d] = %d\n", index, value)
}

// 输出结果：
// slice[0] = 10
// slice[1] = 20
// slice[2] = 30
// slice[3] = 40
// slice[4] = 50

// 只遍历值
for _, value := range slice {
    fmt.Printf("值: %d\n", value)
}

// 输出结果：
// 值: 10
// 值: 20
// 值: 30
// 值: 40
// 值: 50
```

### 4.2 切片遍历注意事项
```go
slice := []string{"Apple", "Banana", "Cherry"}

// 注意：range 中的 value 是副本
for i, value := range slice {
    // 修改 value 不会影响原切片
    value = value + " Modified"
    fmt.Printf("索引: %d, 修改后的值: %s\n", i, value)
}

// 输出结果：
// 索引: 0, 修改后的值: Apple Modified
// 索引: 1, 修改后的值: Banana Modified
// 索引: 2, 修改后的值: Cherry Modified

fmt.Println("原切片:", slice) // 原切片未改变

// 输出结果：
// 原切片: [Apple Banana Cherry]

// 正确的修改方式
for i := range slice {
    slice[i] = slice[i] + " Modified"
}
fmt.Println("修改后的切片:", slice)

// 输出结果：
// 修改后的切片: [Apple Modified Banana Modified Cherry Modified]
```

## 5. 映射（Map）遍历

### 5.1 基本 Map 遍历
```go
m := map[string]int{
    "Apple":  1,
    "Banana": 2,
    "Cherry": 3,
}

// 遍历键值对
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// 可能的输出结果（顺序随机）：
// Apple: 1
// Banana: 2
// Cherry: 3

// 只遍历键
for key := range m {
    fmt.Printf("键: %s\n", key)
}

// 可能的输出结果（顺序随机）：
// 键: Apple
// 键: Banana
// 键: Cherry

// 只遍历值
for _, value := range m {
    fmt.Printf("值: %d\n", value)
}

// 可能的输出结果（顺序随机）：
// 值: 1
// 值: 2
// 值: 3
```

### 5.2 Map 遍历顺序
```go
// 注意：Map 的遍历顺序是随机的
m := map[int]string{
    1: "one",
    2: "two", 
    3: "three",
}

fmt.Println("第一次遍历:")
for k, v := range m {
    fmt.Printf("%d: %s\n", k, v)
}

// 可能的输出结果：
// 第一次遍历:
// 3: three
// 1: one
// 2: two

fmt.Println("\n第二次遍历:")
for k, v := range m {
    fmt.Printf("%d: %s\n", k, v)
}

// 可能的输出结果（顺序可能不同）：
// 第二次遍历:
// 2: two
// 3: three
// 1: one
// 两次遍历的顺序可能不同
```

## 6. 通道（Channel）遍历

```go
ch := make(chan int)

// 启动 goroutine 发送数据
go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}()

// 遍历通道（直到通道关闭）
for value := range ch {
    fmt.Printf("接收到: %d\n", value)
}

// 输出结果：
// 接收到: 0
// 接收到: 1
// 接收到: 2
// 接收到: 3
// 接收到: 4
```

## 7. 遍历最佳实践

### 7.1 性能考虑
```go
// 对于大型切片，避免不必要的拷贝
largeSlice := make([]int, 1000000)

// 好的做法：只获取索引
for i := range largeSlice {
    // 使用 largeSlice[i]
    process(largeSlice[i])
}

// 避免的做法：拷贝整个值
// for _, value := range largeSlice {
//     process(value) // 每次都会拷贝
// }
```

### 7.2 并发安全遍历
```go
// 在遍历过程中修改集合是危险的
slice := []int{1, 2, 3, 4, 5}

// 危险的做法
for i, value := range slice {
    if value == 3 {
        // 不能在遍历过程中修改切片
        // slice = append(slice, 6) // 会导致未定义行为
        fmt.Printf("找到: %d\n", value)
    }
}

// 安全的做法：创建副本
safeSlice := make([]int, len(slice))
copy(safeSlice, slice)

for i, value := range safeSlice {
    if value == 3 {
        slice = append(slice, 6) // 安全
        fmt.Printf("找到: %d\n", value)
    }
}
```

## 8. 常用遍历模式

### 8.1 过滤模式
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var evenNumbers []int

for _, num := range numbers {
    if num%2 == 0 {
        evenNumbers = append(evenNumbers, num)
    }
}
fmt.Println("偶数:", evenNumbers)

// 输出结果：
// 偶数: [2 4 6 8 10]
```

### 8.2 映射转换模式
```go
words := []string{"hello", "world", "golang"}
var upperWords []string

for _, word := range words {
    upperWords = append(upperWords, strings.ToUpper(word))
}
fmt.Println("大写:", upperWords)

// 输出结果：
// 大写: [HELLO WORLD GOLANG]
```

### 8.3 累积模式
```go
numbers := []int{1, 2, 3, 4, 5}
sum := 0

for _, num := range numbers {
    sum += num
}
fmt.Println("总和:", sum)

// 输出结果：
// 总和: 15
```

## 9. 总结

### 9.1 字符遍历选择
- **ASCII 字符串**：直接使用 `for i := 0; i < len(str); i++`
- **Unicode 字符串**：使用 `for i, char := range str`（推荐）

### 9.2 数组/切片遍历选择
- **需要索引**：`for i, v := range slice`
- **只需要值**：`for _, v := range slice`
- **需要修改原值**：`for i := range slice { slice[i] = newValue }`
- **性能敏感**：使用传统 for 循环避免拷贝

### 9.3 推荐做法
1. 优先使用 `for...range` 语法，代码更简洁
2. 注意 Unicode 字符的处理
3. 避免在遍历过程中修改集合
4. 对于大型数据结构，考虑性能影响
5. 使用有意义的变量名提高代码可读性

这些遍历操作是 Go 语言编程的基础，掌握它们对于高效编写 Go 代码至关重要。
