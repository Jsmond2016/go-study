# 函数示例

本示例全面展示了 Go 语言中的函数概念，包括基本函数、多返回值、可变参数、闭包、递归、高阶函数和方法等。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 基本函数定义

```go
// 基本语法
func functionName(parameters) (returnTypes) {
    // 函数体
    return values
}

// 无参数无返回值
func sayHello() {
    fmt.Println("Hello, World!")
}

// 有参数有返回值
func add(a, b int) int {
    return a + b
}

// 命名返回值
func rectangle(width, height int) (area, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return // 自动返回 area 和 perimeter
}
```

### 2. 多返回值

Go 函数可以返回多个值：

```go
func divide(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return
}

func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

// 使用
q, r := divide(17, 5)
result, err := safeDivide(10, 2)
if err != nil {
    // 处理错误
}
```

### 3. 可变参数函数

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// 调用方式
total := sum(1, 2, 3, 4, 5)
numbers := []int{10, 20, 30}
total := sum(numbers...) // 使用切片
```

### 4. 闭包和匿名函数

```go
// 匿名函数
add := func(a, b int) int {
    return a + b
}

// 闭包：捕获外部变量
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// 计数器闭包
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### 5. 递归函数

```go
// 基本递归
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// 尾递归优化
func factorialTail(n, acc int) int {
    if n <= 1 {
        return acc
    }
    return factorialTail(n-1, n*acc)
}
```

### 6. 高阶函数

函数可以作为参数和返回值：

```go
// 函数作为参数
func filter(numbers []int, predicate func(int) bool) []int {
    var result []int
    for _, num := range numbers {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}

// 函数作为返回值
func adder(n int) func(int) int {
    return func(x int) int {
        return x + n
    }
}

// 函数组合
func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}
```

### 7. 方法

为类型定义方法：

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// 值接收者方法
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 指针接收者方法（可修改原值）
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// 实现 Stringer 接口
func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{%.1fx%.1f}", r.Width, r.Height)
}
```

### 8. defer 延迟执行

```go
func fileOperation() {
    fmt.Println("打开文件")
    defer fmt.Println("关闭文件") // 函数退出前执行

    fmt.Println("读取文件内容")
    // 即使函数中途返回，defer 也会执行
}

// 多个 defer 按栈式执行（后进先出）
func deferExample() {
    defer fmt.Println("第3个")
    defer fmt.Println("第2个")
    defer fmt.Println("第1个") // 最先执行
}
```

## 重要概念

### 1. 值传递 vs 引用传递

Go 默认是值传递，但可以通过指针实现引用传递：
- 基本类型：值传递
- 指针：引用传递
- 切片、映射、通道：引用传递（内部是指针）

### 2. 命名返回值

- 命名的返回变量在函数开始时被初始化为零值
- `return` 语句会返回当前命名变量的值
- 适合复杂的返回逻辑

### 3. 错误处理

Go 推荐使用多返回值处理错误：
```go
func safeOperation() (result string, err error) {
    // 操作
    if errorCondition {
        return "", errors.New("操作失败")
    }
    return "成功", nil
}
```

### 4. 接口方法

方法可以实现接口：
```go
type Stringer interface {
    String() string
}
```

## 最佳实践

1. **函数命名**：使用动词或动词短语，遵循驼峰命名法
2. **参数设计**：避免过多参数（建议不超过4个）
3. **错误处理**：总是检查和处理错误
4. **资源清理**：使用 `defer` 确保资源释放
5. **文档注释**：为公开函数添加文档注释

## 常见陷阱

1. **nil 指针调用方法**：先检查指针是否为 nil
2. **闭包循环变量**：循环中使用闭包要注意捕获时机
3. **递归深度**：避免过深的递归导致栈溢出
4. **defer 性能**：defer 有一定的性能开销
5. **返回值忽略**：使用 `_` 忽略不需要的返回值

## 练习

1. 实现计算器的各种运算函数
2. 编写递归函数解决汉诺塔问题
3. 创建一个简单的状态机
4. 实现函数式编程工具（map、reduce、filter）
5. 编写一个带有缓存功能的斐波那契函数