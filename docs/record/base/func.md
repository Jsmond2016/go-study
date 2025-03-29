# 函数



我来总结一下 Go 语言中函数的常用操作和特性。建议在 <mcfile name="index.md" path="/Users/huangjin/Desktop/github/golang-pro/go-study/docs/index.md"></mcfile> 中添加相关内容。

1. **基本函数定义**
```go
// 基本函数
func functionName(param1 Type1, param2 Type2) ReturnType {
    // 函数体
    return value
}

// 多返回值
func multiReturn(x, y int) (int, int) {
    return x + y, x * y
}
```

2. **命名返回值**
```go
func namedReturn(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return // 裸返回，自动返回 sum 和 product
}
```

3. **可变参数函数**
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

4. **函数作为参数**
```go
func calculate(x int, y int, operation func(int, int) int) int {
    return operation(x, y)
}

// 使用示例
add := func(x, y int) int { return x + y }
result := calculate(10, 20, add)
```

5. **闭包**
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

6. **方法（绑定到结构体）**
```go
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}
```

7. **延迟执行 defer**
```go
func readFile() {
    file := openFile()
    defer file.Close() // 函数结束前执行
    // 处理文件...
}
```

8. **init 函数**
```go
func init() {
    // 包初始化时自动执行
    // 每个包可以有多个 init 函数
}
```

重要注意事项：

1. Go 不支持函数重载
2. 函数是一等公民，可以作为值传递
3. 支持匿名函数
4. 闭包可以捕获外部变量
5. defer 语句按 LIFO（后进先出）顺序执行
6. init 函数在 main 函数之前自动执行

建议在文档中添加这些基础函数操作的示例和说明，可以帮助更好地理解 Go 语言的函数特性。


相关资料：

- [Go语言基础之函数 | 李文周的博客](https://www.liwenzhou.com/posts/Go/function/)
- [函数 · Go语言中文文档](https://www.topgoer.com/%E5%87%BD%E6%95%B0/)