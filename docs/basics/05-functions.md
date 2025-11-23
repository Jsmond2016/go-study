---
title: 函数
difficulty: beginner
duration: "5-6小时"
prerequisites: ["变量与常量", "数据类型", "运算符", "控制流程"]
tags: ["函数", "参数", "返回值", "闭包"]
---

# 函数

函数是 Go 语言的基本组成单元，用于封装可重用的代码逻辑。Go 的函数设计简洁而强大，支持多返回值、闭包等特性。

## 📋 学习目标

- [ ] 掌握函数的定义和调用
- [ ] 理解参数和返回值机制
- [ ] 学会使用多返回值
- [ ] 掌握可变参数函数
- [ ] 理解闭包的概念和应用
- [ ] 学会使用递归函数
- [ ] 了解函数类型和函数值

## 🔧 函数基础

### 基本函数定义

```go
package main

import "fmt"

// 基本函数定义
func greet(name string) {
	fmt.Printf("你好，%s！\n", name)
}

// 带返回值的函数
func add(a, b int) int {
	return a + b
}

// 带多个返回值的函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除零错误")
	}
	return a / b, nil
}

func main() {
	// 函数调用
	greet("张三")
	
	// 调用带返回值的函数
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)
	
	// 调用多返回值函数
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", quotient)
	}
	
	// 忽略某个返回值
	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("除零错误: %v\n", err2)
	}
}
```

### 函数命名和可见性

```go
package main

import "fmt"

// 公开函数（首字母大写）
func PublicFunction() {
	fmt.Println("这是一个公开函数")
}

// 私有函数（首字母小写）
func privateFunction() {
	fmt.Println("这是一个私有函数")
}

// 命名返回值
func calculate(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b
	return // 隐式返回
}

func main() {
	PublicFunction()
	
	s, diff := calculate(20, 5)
	fmt.Printf("和: %d, 差: %d\n", s, diff)
}
```

## 📥 参数处理

### 值传递和引用传递

```go
package main

import "fmt"

// 值传递（基本类型）
func modifyValue(x int) {
	x = 100 // 修改的是副本
}

// 引用传递（指针）
func modifyPointer(x *int) {
	*x = 100 // 修改的是原始值
}

// 切片传递（引用类型）
func modifySlice(numbers []int) {
	numbers[0] = 100 // 修改的是原始切片
}

func main() {
	// 基本类型的值传递
	value := 50
	fmt.Printf("修改前: %d\n", value)
	modifyValue(value)
	fmt.Printf("修改后: %d\n", value) // 还是 50
	
	// 指针的引用传递
	value2 := 50
	fmt.Printf("\n指针修改前: %d\n", value2)
	modifyPointer(&value2)
	fmt.Printf("指针修改后: %d\n", value2) // 变成 100
	
	// 切片的引用传递
	numbers := []int{1, 2, 3}
	fmt.Printf("\n切片修改前: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("切片修改后: %v\n", numbers) // 第一个元素变成 100
}
```

### 可变参数函数

```go
package main

import "fmt"

// 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 混合参数（固定参数 + 可变参数）
func greeting(prefix string, names ...string) {
	fmt.Printf("%s: ", prefix)
	for i, name := range names {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(name)
	}
	fmt.Println()
}

func main() {
	// 调用可变参数函数
	result1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("sum(1,2,3,4,5) = %d\n", result1)
	
	result2 := sum()
	fmt.Printf("sum() = %d\n", result2)
	
	// 调用混合参数函数
	greeting("大家好", "张三", "李四", "王五")
	greeting("欢迎", "赵六")
	
	// 传递切片给可变参数函数
	numbers := []int{10, 20, 30}
	result3 := sum(numbers...)
	fmt.Printf("sum([]int{10,20,30}...) = %d\n", result3)
}
```

## 🔄 高级函数特性

### 函数类型和函数值

```go
package main

import "fmt"

// 函数类型定义
type Operation func(int, int) int

// 使用函数类型作为参数
func calculate(a, b int, op Operation) int {
	return op(a, b)
}

// 返回函数的函数
func getOperation(operator string) Operation {
	switch operator {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int { return a / b }
	default:
		return nil
	}
}

func main() {
	// 直接使用函数类型
	var add Operation = func(a, b int) int { return a + b }
	
	result := calculate(10, 5, add)
	fmt.Printf("10 + 5 = %d\n", result)
	
	// 使用匿名函数
	result = calculate(20, 8, func(a, b int) int {
		return a * b + a - b
	})
	fmt.Printf("20 * 8 + 20 - 8 = %d\n", result)
	
	// 使用返回的函数
	multiply := getOperation("*")
	if multiply != nil {
		result = calculate(6, 7, multiply)
		fmt.Printf("6 * 7 = %d\n", result)
	}
}
```

### 闭包（Closure）

```go
package main

import "fmt"

// 返回闭包的函数
func makeAdder(increment int) func(int) int {
	return func(x int) int {
		return x + increment
	}
}

// 闭包用于状态保持
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// 基本闭包
	add5 := makeAdder(5)
	add10 := makeAdder(10)
	
	fmt.Printf("add5(20) = %d\n", add5(20))  // 25
	fmt.Printf("add10(20) = %d\n", add10(20)) // 30
	
	// 闭包保持状态
	c1 := counter()
	c2 := counter()
	
	fmt.Printf("\n计数器1: %d\n", c1()) // 1
	fmt.Printf("计数器1: %d\n", c1()) // 2
	fmt.Printf("计数器2: %d\n", c2()) // 1
	fmt.Printf("计数器1: %d\n", c1()) // 3
	fmt.Printf("计数器2: %d\n", c2()) // 2
	
	// 闭包在实际应用中
	message := "Hello"
	
	// 延迟执行闭包
	delayedFunc := func() {
		fmt.Println("延迟执行:", message)
	}
	
	message = "World"
	delayedFunc() // 输出: 延迟执行: World
}
```

### 递归函数

```go
package main

import "fmt"

// 阶乘递归
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 递归遍历目录结构（模拟）
type File struct {
	name     string
	isDir    bool
	children []File
}

func printFiles(files []File, indent int) {
	for _, file := range files {
		prefix := ""
		for i := 0; i < indent; i++ {
			prefix += "  "
		}
		
		fmt.Printf("%s%s/\n", prefix, file.name)
		if file.isDir && len(file.children) > 0 {
			printFiles(file.children, indent+1)
		}
	}
}

func main() {
	// 数学递归
	fmt.Printf("5! = %d\n", factorial(5))
	
	fmt.Printf("\n斐波那契数列前10项:\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}
	
	// 结构递归
	root := File{
		name:  "root",
		isDir: true,
		children: []File{
			{name: "src", isDir: true, children: []File{
				{name: "main.go", isDir: false},
				{name: "utils.go", isDir: false},
			}},
			{name: "docs", isDir: true, children: []File{
				{name: "README.md", isDir: false},
			}},
			{name: "go.mod", isDir: false},
		},
	}
	
	fmt.Println("\n目录结构:")
	printFiles([]File{root}, 0)
}
```

## 🎯 延迟调用和错误处理

### defer 在函数中的应用

```go
package main

import "fmt"

// 资源管理示例
func processFile(filename string) error {
	fmt.Printf("打开文件: %s\n", filename)
	
	// 模拟资源清理
	defer fmt.Printf("关闭文件: %s\n", filename)
	
	// 模拟处理过程
	if filename == "error.txt" {
		return fmt.Errorf("处理文件时出错")
	}
	
	fmt.Printf("处理文件: %s\n", filename)
	return nil
}

// 多个 defer 的执行顺序
func demonstrateDeferOrder() {
	fmt.Println("函数开始")
	
	defer fmt.Println("第一个 defer")
	defer fmt.Println("第二个 defer")
	defer fmt.Println("第三个 defer")
	
	fmt.Println("函数执行中...")
	
	defer func() {
		fmt.Println("匿名 defer 函数")
	}()
	
	fmt.Println("函数即将结束")
}

func main() {
	// 正常情况
	err := processFile("example.txt")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
	
	fmt.Println("\n---")
	
	// 错误情况
	err = processFile("error.txt")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
	
	fmt.Println("\n--- defer 执行顺序 ---")
	demonstrateDeferOrder()
}
```

### panic 和 recover

```go
package main

import "fmt"

// 可能引发 panic 的函数
func riskyOperation(shouldPanic bool) {
	if shouldPanic {
		panic("这是一个严重错误！")
	}
	fmt.Println("操作成功完成")
}

// 使用 recover 处理 panic
func safeOperation(shouldPanic bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 panic: %v\n", r)
			fmt.Println("程序继续执行")
		}
	}()
	
	fmt.Println("开始安全操作")
	riskyOperation(shouldPanic)
	fmt.Println("安全操作结束")
}

func main() {
	fmt.Println("=== 正常情况 ===")
	safeOperation(false)
	
	fmt.Println("\n=== 异常情况 ===")
	safeOperation(true)
	
	fmt.Println("程序继续正常运行")
}
```

## 🧪 实践练习

### 练习 1：计算器库

```go
// 要求：
// 1. 实现基本四则运算函数
// 2. 支持链式调用
// 3. 添加错误处理
// 4. 使用函数类型提高扩展性
```

参考实现：
```go
package main

import "fmt"

type Calculator struct {
	value float64
}

func NewCalculator(initial float64) *Calculator {
	return &Calculator{value: initial}
}

func (c *Calculator) Add(x float64) *Calculator {
	c.value += x
	return c
}

func (c *Calculator) Subtract(x float64) *Calculator {
	c.value -= x
	return c
}

func (c *Calculator) Multiply(x float64) *Calculator {
	c.value *= x
	return c
}

func (c *Calculator) Divide(x float64) (*Calculator, error) {
	if x == 0 {
		return nil, fmt.Errorf("除零错误")
	}
	c.value /= x
	return c, nil
}

func (c *Calculator) Result() float64 {
	return c.value
}

func main() {
	calc := NewCalculator(10)
	result, err := calc.Add(5).Multiply(3).Subtract(8).Divide(2)
	if err != nil {
		fmt.Printf("计算错误: %v\n", err)
		return
	}
	
	fmt.Printf("计算结果: ((10 + 5) * 3 - 8) / 2 = %.2f\n", calc.Result())
}
```

### 练习 2：函数工厂

```go
// 要求：
// 1. 创建返回不同类型函数的工厂
// 2. 使用闭包保持状态
// 3. 实现函数组合
// 4. 支持函数缓存
```

### 练习 3：递归算法实现

```go
// 要求：
// 1. 实现二叉树遍历
// 2. 实现快速排序
// 3. 实现深度优先搜索
// 4. 分析递归性能
```

## 🤔 思考题

1. **什么时候应该使用值传递，什么时候使用指针传递？**
   - 提示：考虑性能和修改需求

2. **闭包的主要应用场景有哪些？**
   - 提示：考虑状态保持、延迟执行、函数工厂

3. **递归函数的优缺点是什么？**
   - 提示：考虑代码简洁性、性能、栈溢出

4. **defer 语句在实际开发中的最佳实践？**
   - 提示：考虑资源管理、日志记录、性能测量

## 📚 扩展阅读

- [Go 语言规范 - 函数](https://golang.org/ref/spec#Function_declarations)
- [Effective Go - 函数](https://golang.org/doc/effective_go.html#functions)
- [Go by Example - 函数](https://gobyexample.com/functions)

## ⏭️ 下一章节

[数组](./06-arrays.md) → 学习 Go 语言的数组类型

---

**💡 小贴士**: Go 的函数设计强调简洁性和明确性。记住：**多返回值是 Go 的特色，闭包是函数式编程的利器，合理使用能让代码更优雅**！
