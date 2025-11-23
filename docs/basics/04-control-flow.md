---
title: 控制流程
difficulty: beginner
duration: "4-5小时"
prerequisites: ["变量与常量", "数据类型", "运算符"]
tags: ["控制流程", "条件语句", "循环", "跳转语句"]
---

# 控制流程

控制流程决定了程序的执行顺序。Go 提供了多种控制结构，包括条件语句、循环语句和跳转语句。

## 📋 学习目标

- [ ] 掌握 if-else 条件语句
- [ ] 学会使用 switch 语句
- [ ] 理解 for 循环的各种形式
- [ ] 掌握 break 和 continue 的使用
- [ ] 学会使用 goto 和 defer
- [ ] 理解标签的作用

## 🔀 条件语句

### if-else 语句

```go
package main

import "fmt"

func main() {
	// 基本 if 语句
	age := 18
	
	if age >= 18 {
		fmt.Println("你是成年人")
	}
	
	// if-else 语句
	score := 85
	
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	
	// if 的初始化语句
	if num := 42; num > 0 {
		fmt.Printf("%d 是正数\n", num)
		// num 的作用域仅限于 if 语句块内
	}
	
	// 复杂条件判断
	height := 175
	weight := 70
	hasLicense := true
	
	// 使用括号明确优先级
	if (height >= 170 && height <= 180) && 
	   (weight >= 60 && weight <= 80) && 
	   hasLicense {
		fmt.Println("符合标准，可以考驾照")
	}
}
```

### switch 语句

```go
package main

import "fmt"

func main() {
	// 基本 switch 语句
	day := 3
	
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("无效的星期")
	}
	
	// switch 表达式
	grade := 'B'
	
	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B', 'C':
		fmt.Println("良好")
	case 'D':
		fmt.Println("及格")
	case 'F':
		fmt.Println("不及格")
	}
	
	// 无表达式的 switch
	score := 85
	
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
	
	// fallthrough 关键字
	fmt.Println("测试 fallthrough:")
	switch score / 10 {
	case 9:
		fmt.Println("90多分")
		fallthrough // 继续执行下一个 case
	case 8:
		fmt.Println("80多分")
	case 7:
		fmt.Println("70多分")
	default:
		fmt.Println("其他分数段")
	}
}
```

## 🔁 循环语句

### 基本 for 循环

```go
package main

import "fmt"

func main() {
	// 传统 for 循环
	fmt.Println("基本 for 循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	
	// 省略初始条件
	fmt.Println("\n省略初始条件:")
	j := 0
	for ; j < 3; j++ {
		fmt.Printf("j = %d\n", j)
	}
	
	// 省略条件
	fmt.Println("\n省略条件（死循环）:")
	k := 0
	for ; ; k++ {
		if k >= 2 {
			break // 跳出循环
		}
		fmt.Printf("k = %d\n", k)
	}
	
	// 省略所有条件（无限循环）
	fmt.Println("\n省略所有条件:")
	count := 0
	for {
		count++
		if count > 2 {
			break
		}
		fmt.Printf("count = %d\n", count)
	}
}
```

### for-range 循环

```go
package main

import "fmt"

func main() {
	// 遍历切片
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("遍历切片:")
	for index, value := range numbers {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}
	
	// 只要值，忽略索引
	fmt.Println("\n只要值:")
	for _, value := range numbers {
		fmt.Printf("值: %d\n", value)
	}
	
	// 遍历数组
	letters := [5]rune{'H', 'e', 'l', 'l', 'o'}
	fmt.Println("\n遍历数组:")
	for i, char := range letters {
		fmt.Printf("位置: %d, 字符: %c\n", i, char)
	}
	
	// 遍历字符串
	text := "Hello, 世界"
	fmt.Println("\n遍历字符串:")
	for i, char := range text {
		fmt.Printf("位置: %d, 字符: %c\n", i, char)
	}
	
	// 遍历映射
	ages := map[string]int{
		"张三": 25,
		"李四": 30,
		"王五": 35,
	}
	fmt.Println("\n遍历映射:")
	for name, age := range ages {
		fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
	}
	
	// 遍历通道
	ch := make(chan string, 3)
	go func() {
		ch <- "第一条消息"
		ch <- "第二条消息"
		ch <- "第三条消息"
		close(ch)
	}()
	
	fmt.Println("\n遍历通道:")
	for msg := range ch {
		fmt.Printf("收到消息: %s\n", msg)
	}
}
```

### 循环控制示例

```go
package main

import "fmt"

func main() {
	// 嵌套循环
	fmt.Println("九九乘法表:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%-2d ", j, i, i*j)
		}
		fmt.Println()
	}
	
	// 循环中的条件判断
	fmt.Println("\n查找第一个大于100的2的幂:")
	for i := 1; ; i *= 2 {
		if i > 100 {
			fmt.Printf("找到: %d\n", i)
			break
		}
		fmt.Printf("检查: %d\n", i)
	}
	
	// 使用标签控制循环
	fmt.Println("\n使用标签:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("跳过内层循环 i=%d, j=%d\n", i, j)
				continue outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
```

## 🏃‍♂️ 跳转语句

### break 和 continue

```go
package main

import "fmt"

func main() {
	// break 示例
	fmt.Println("break 示例 - 查找第一个偶数:")
	numbers := []int{3, 7, 9, 8, 11, 13, 16}
	
	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("找到第一个偶数: %d (索引: %d)\n", num, i)
			break
		}
		fmt.Printf("检查数字: %d\n", num)
	}
	
	// continue 示例
	fmt.Println("\ncontinue 示例 - 打印奇数:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// 嵌套循环中的 break/continue
	fmt.Println("\n嵌套循环中的控制:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("外层循环 i=%d\n", i)
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("  跳过内层 i=%d, j=%d\n", i, j)
				continue
			}
			if i == 3 && j == 2 {
				fmt.Printf("  跳出外层循环 i=%d, j=%d\n", i, j)
				break
			}
			fmt.Printf("  内层循环 i=%d, j=%d\n", i, j)
		}
		if i == 3 {
			fmt.Println("外层循环也被跳出")
			break
		}
	}
}
```

### 标签和 goto

```go
package main

import "fmt"

func main() {
	// 标签定义
	fmt.Println("goto 示例:")
	i := 0
	
start:
	if i < 3 {
		fmt.Printf("i = %d\n", i)
		i++
		goto start
	}
	fmt.Println("循环结束")
	
	// 标签用于错误处理
	fmt.Println("\n标签用于错误处理:")
	file := "example.txt"
	
	// 模拟文件处理
retry:
	fmt.Printf("尝试处理文件: %s\n", file)
	
	// 模拟随机成功或失败
	if i < 2 { // 假设前两次失败
		fmt.Println("处理失败，重试...")
		i++
		goto retry
	}
	
	fmt.Println("处理成功！")
}
```

## 🔄 defer 语句

### defer 基本用法

```go
package main

import "fmt"

func main() {
	// 基本 defer 示例
	fmt.Println("函数开始")
	defer fmt.Println("函数结束")
	
	fmt.Println("函数执行中...")
	
	// 多个 defer 的执行顺序（后进先出）
	fmt.Println("\n多个 defer 示例:")
	defer fmt.Println("第三个 defer")
	defer fmt.Println("第二个 defer")
	defer fmt.Println("第一个 defer")
	
	fmt.Println("主函数执行")
	
	testDefer()
}

func testDefer() {
	fmt.Println("\n在 testDefer 函数中:")
	defer fmt.Println("testDefer 结束")
	
	fmt.Println("testDefer 执行中...")
}
```

### defer 与返回值

```go
package main

import "fmt"

func main() {
	// defer 修改返回值
	result := calculate()
	fmt.Printf("最终结果: %d\n", result)
	
	// defer 与闭包
	fmt.Println("\ndefer 与闭包:")
	numbers := []int{1, 2, 3, 4, 5}
	
	for _, num := range numbers {
		defer func(n int) {
			fmt.Printf("处理数字: %d\n", n)
		}(num)
	}
	
	fmt.Println("所有 defer 函数设置完成")
}

func calculate() int {
	defer func() {
		fmt.Println("defer 1: 在返回前执行")
	}()
	
	result := 42
	
	defer func() {
		fmt.Println("defer 2: 可以访问返回值", result)
	}()
	
	defer func(r int) {
		fmt.Println("defer 3: 修改返回值为", r*2)
		return
	}(result)
	
	fmt.Println("函数即将返回")
	return result
}
```

## 🧪 实践练习

### 练习 1：成绩等级判断

```go
// 要求：
// 1. 使用 if-else 实现成绩等级判断
// 2. 使用 switch 重新实现
// 3. 添加输入验证
// 4. 测试边界条件
```

参考实现：
```go
package main

import "fmt"

func getGradeWithIf(score int) string {
	if score >= 90 {
		return "A（优秀）"
	} else if score >= 80 {
		return "B（良好）"
	} else if score >= 70 {
		return "C（中等）"
	} else if score >= 60 {
		return "D（及格）"
	} else {
		return "F（不及格）"
	}
}

func getGradeWithSwitch(score int) string {
	switch {
	case score >= 90:
		return "A（优秀）"
	case score >= 80:
		return "B（良好）"
	case score >= 70:
		return "C（中等）"
	case score >= 60:
		return "D（及格）"
	default:
		return "F（不及格）"
	}
}

func main() {
	scores := []int{95, 82, 75, 68, 55}
	
	fmt.Println("使用 if-else 判断:")
	for _, score := range scores {
		fmt.Printf("分数 %d: %s\n", score, getGradeWithIf(score))
	}
	
	fmt.Println("\n使用 switch 判断:")
	for _, score := range scores {
		fmt.Printf("分数 %d: %s\n", score, getGradeWithSwitch(score))
	}
}
```

### 练习 2：数字模式打印

```go
// 要求：
// 1. 使用循环打印数字模式
// 2. 使用 continue 控制打印
// 3. 使用 break 控制循环
// 4. 使用标签控制嵌套循环
```

### 练习 3：资源清理模拟

```go
// 要求：
// 1. 模拟资源分配和释放
// 2. 使用 defer 确保资源清理
// 3. 测试错误情况下的清理
// 4. 比较有 defer 和无 defer 的差异
```

## 🤔 思考题

1. **for-range 循环中，什么时候应该使用索引，什么时候应该用值？**
   - 提示：考虑是否需要修改集合元素

2. **goto 语句在什么情况下是合适的？**
   - 提示：考虑错误处理和资源清理

3. **defer 语句的实际应用场景有哪些？**
   - 提示：考虑资源管理、日志记录、性能测量

4. **switch 语句和 if-else 链相比有什么优势？**
   - 提示：考虑代码可读性和性能

## 📚 扩展阅读

- [Go 语言规范 - 语句](https://golang.org/ref/spec#Statements)
- [Effective Go - 控制流程](https://golang.org/doc/effective_go.html#control_flow)
- [Go by Example - 流程控制](https://gobyexample.com/flow-control)

## ⏭️ 下一章节

[函数](./05-functions.md) → 学习 Go 语言的函数定义和使用

---

**💡 小贴士**: Go 的控制流程设计简洁而强大。记住：**优先使用 for-range 处理集合，使用 defer 进行资源清理，这会让你的代码更安全、更易维护**！
