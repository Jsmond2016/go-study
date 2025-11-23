// Package main 展示 Go 语言的函数
package main

import (
	"errors"
	"fmt"
	"math"
)

// 程序入口，演示各种函数
func main() {
	fmt.Println("=== Go 函数示例 ===")

	// 基本函数
	basicFunctions()

	// 多返回值函数
	multipleReturnValues()

	// 可变参数函数
	variadicFunctions()

	// 闭包和匿名函数
	closuresAndAnonymousFunctions()

	// 递归函数
	recursiveFunctions()

	// 高阶函数
	higherOrderFunctions()

	// 方法定义
	methods()

	// 延迟执行
	deferStatements()
}

// basicFunctions 基本函数
func basicFunctions() {
	fmt.Println("\n--- 基本函数 ---")

	// 调用无参数无返回值函数
	sayHello()

	// 调用有参数无返回值函数
	greet("张三")

	// 调用有参数有返回值函数
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	// 调用命名返回值函数
	area, perimeter := rectangle(5, 3)
	fmt.Printf("矩形(5x3): 面积=%d, 周长=%d\n", area, perimeter)
}

// sayHello 无参数无返回值函数
func sayHello() {
	fmt.Println("你好，Go！")
}

// greet 有参数无返回值函数
func greet(name string) {
	fmt.Printf("你好，%s！\n", name)
}

// add 有参数有返回值函数
func add(a, b int) int {
	return a + b
}

// rectangle 命名返回值函数
func rectangle(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 自动返回 area 和 perimeter
}

// multipleReturnValues 多返回值函数
func multipleReturnValues() {
	fmt.Println("\n--- 多返回值函数 ---")

	// 调用多返回值函数
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)

	// 使用空白标识符忽略某个返回值
	_, rem := divide(10, 3)
	fmt.Printf("10 ÷ 3 的余数是 %d\n", rem)

	// 调用返回错误的函数
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}
}

// divide 多返回值函数
func divide(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

// safeDivide 返回错误的函数
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// variadicFunctions 可变参数函数
func variadicFunctions() {
	fmt.Println("\n--- 可变参数函数 ---")

	// 调用可变参数函数
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", total)

	// 使用切片调用可变参数函数
	numbers := []int{10, 20, 30}
	total = sum(numbers...)
	fmt.Printf("10+20+30 = %d\n", total)

	// 混合参数
	result := combine("Numbers:", 1, 2, 3, 4, 5)
	fmt.Println(result)
}

// sum 可变参数函数
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// combine 混合参数函数
func combine(prefix string, nums ...int) string {
	result := prefix
	for i, num := range nums {
		if i > 0 {
			result += "+"
		}
		result += fmt.Sprintf("%d", num)
	}
	result += fmt.Sprintf("=%d", sum(nums...))
	return result
}

// closuresAndAnonymousFunctions 闭包和匿名函数
func closuresAndAnonymousFunctions() {
	fmt.Println("\n--- 闭包和匿名函数 ---")

	// 1. 基本匿名函数
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("匿名函数结果: %d\n", add(3, 5))

	// 2. 闭包（捕获外部变量）
	multiplier := createMultiplier(3)
	fmt.Printf("3 × 10 = %d\n", multiplier(10))
	fmt.Printf("3 × 7 = %d\n", multiplier(7))

	// 3. 返回闭包
	getCounter := createCounter()
	fmt.Printf("计数器: %d\n", getCounter())
	fmt.Printf("计数器: %d\n", getCounter())
	fmt.Printf("计数器: %d\n", getCounter())

	// 4. 立即执行匿名函数
	message := func(name string) string {
		return fmt.Sprintf("你好，%s！", name)
	}("世界")
	fmt.Println(message)
}

// createMultiplier 创建乘法器闭包
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// createCounter 创建计数器闭包
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// recursiveFunctions 递归函数
func recursiveFunctions() {
	fmt.Println("\n--- 递归函数 ---")

	// 阶乘
	n := 5
	factorialResult := factorial(n)
	fmt.Printf("%d! = %d\n", n, factorialResult)

	// 斐波那契数列
	fmt.Printf("前10个斐波那契数: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()

	// 使用尾递归优化的阶乘
	factorialTailResult := factorialTail(n, 1)
	fmt.Printf("尾递归 %d! = %d\n", n, factorialTailResult)
}

// factorial 阶乘函数
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// fibonacci 斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// factorialTail 尾递归优化的阶乘
func factorialTail(n, acc int) int {
	if n <= 1 {
		return acc
	}
	return factorialTail(n-1, n*acc)
}

// higherOrderFunctions 高阶函数
func higherOrderFunctions() {
	fmt.Println("\n--- 高阶函数 ---")

	// 1. 函数作为参数
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 使用函数处理切片
	evenNumbers := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("偶数: %v\n", evenNumbers)

	oddNumbers := filter(numbers, func(n int) bool {
		return n%2 != 0
	})
	fmt.Printf("奇数: %v\n", oddNumbers)

	// 2. 函数作为返回值
	add5 := adder(5)
	fmt.Printf("5 + 10 = %d\n", add5(10))

	multiply3 := multiplier(3)
	fmt.Printf("3 × 10 = %d\n", multiply3(10))

	// 3. 函数组合
	composed := compose(multiply3, add5)
	fmt.Printf("(5 + 10) × 3 = %d\n", composed(10))
}

// filter 过滤函数
func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// adder 加法器
func adder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// multiplier 乘法器
func multiplier(n int) func(int) int {
	return func(x int) int {
		return x * n
	}
}

// compose 函数组合
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// methods 方法定义
func methods() {
	fmt.Println("\n--- 方法定义 ---")

	// 创建 Rectangle 实例
	rect := Rectangle{Width: 10, Height: 5}

	// 调用值接收者方法
	fmt.Printf("矩形面积: %.1f\n", rect.Area())
	fmt.Printf("矩形周长: %.1f\n", rect.Perimeter())

	// 调用指针接收者方法
	rect.Scale(2.0)
	fmt.Printf("缩放后的面积: %.1f\n", rect.Area())

	// 调用 Stringer 接口方法
	fmt.Printf("矩形信息: %s\n", rect.String())

	// 使用指针调用方法
	rectPtr := &rect
	fmt.Printf("通过指针调用面积: %.1f\n", rectPtr.Area())
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算面积（值接收者）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算周长（值接收者）
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Scale 缩放矩形（指针接收者，修改原值）
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// String 实现 Stringer 接口
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width:%.1f, Height:%.1f}", r.Width, r.Height)
}

// deferStatements 延迟执行
func deferStatements() {
	fmt.Println("\n--- 延迟执行 ---")

	// 基本 defer
	fmt.Println("基本 defer 示例:")
	defer fmt.Println("最后执行：defer 语句")
	fmt.Println("中间执行：普通语句")
	fmt.Println("即将结束")

	// 多个 defer（栈式执行）
	fmt.Println("\n多个 defer 示例:")
	defer fmt.Println("第3个 defer")
	defer fmt.Println("第2个 defer")
	defer fmt.Println("第1个 defer")

	// 在函数中使用 defer
	fmt.Println("\n文件操作 defer 示例:")
	fileOperation()

	// 在错误处理中使用 defer
	fmt.Println("\n错误处理 defer 示例:")
	deferExample()
}

// fileOperation 模拟文件操作
func fileOperation() {
	fmt.Println("打开文件")
	defer fmt.Println("关闭文件")

	fmt.Println("读取文件内容")
	fmt.Println("处理文件数据")
	// 即使函数中途返回，defer 也会执行
}

// deferExample defer 使用示例
func deferExample() {
	defer fmt.Println("清理资源")

	result := someCalculation()
	fmt.Printf("计算结果: %d\n", result)
}

// someCalculation 模拟可能出错的计算
func someCalculation() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("恢复从 panic: %v\n", r)
		}
	}()

	num := 10
	fmt.Printf("开始计算，num = %d\n", num)

	// 模拟 panic
	if num > 5 {
		fmt.Println("触发 panic")
		panic("数值过大")
	}

	return num * 2
}