// Package main 展示 Go 语言的运算符
package main

import (
	"fmt"
)

// 程序入口，演示各种运算符
func main() {
	fmt.Println("=== Go 运算符示例 ===")

	// 算术运算符
	arithmeticOperators()

	// 关系运算符
	relationalOperators()

	// 逻辑运算符
	logicalOperators()

	// 位运算符
	bitwiseOperators()

	// 赋值运算符
	assignmentOperators()

	// 特殊运算符
	specialOperators()

	// 运算符优先级
	operatorPrecedence()
}

// arithmeticOperators 算术运算符
func arithmeticOperators() {
	fmt.Println("\n--- 算术运算符 ---")

	var a, b int = 10, 3

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a + b = %d\n", a+b)    // 加法
	fmt.Printf("a - b = %d\n", a-b)    // 减法
	fmt.Printf("a * b = %d\n", a*b)    // 乘法
	fmt.Printf("a / b = %d\n", a/b)    // 除法（整数除法）
	fmt.Printf("a %% b = %d\n", a%b)   // 取模（求余数）

	// 浮点数运算
	var x, y float64 = 10.5, 3.2
	fmt.Printf("\n浮点数运算:\n")
	fmt.Printf("x = %.1f, y = %.1f\n", x, y)
	fmt.Printf("x + y = %.2f\n", x+y)
	fmt.Printf("x - y = %.2f\n", x-y)
	fmt.Printf("x * y = %.2f\n", x*y)
	fmt.Printf("x / y = %.2f\n", x/y)

	// 自增自减运算符
	fmt.Println("\n自增自减运算符:")
	count := 0
	fmt.Printf("初始值: %d\n", count)
	count++
	fmt.Printf("count++ 后: %d\n", count)
	count--
	fmt.Printf("count-- 后: %d\n", count)
}

// relationalOperators 关系运算符
func relationalOperators() {
	fmt.Println("\n--- 关系运算符 ---")

	var a, b int = 10, 20

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a == b: %t\n", a == b)   // 等于
	fmt.Printf("a != b: %t\n", a != b)   // 不等于
	fmt.Printf("a > b:  %t\n", a > b)    // 大于
	fmt.Printf("a < b:  %t\n", a < b)    // 小于
	fmt.Printf("a >= b: %t\n", a >= b)   // 大于等于
	fmt.Printf("a <= b: %t\n", a <= b)   // 小于等于

	// 字符串比较
	fmt.Println("\n字符串比较:")
	str1, str2 := "apple", "banana"
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' < '%s': %t\n", str1, str2, str1 < str2)
}

// logicalOperators 逻辑运算符
func logicalOperators() {
	fmt.Println("\n--- 逻辑运算符 ---")

	var x, y, z bool = true, false, true

	fmt.Printf("x = %t, y = %t, z = %t\n", x, y, z)
	fmt.Printf("x && y: %t\n", x && y)  // 逻辑与（AND）
	fmt.Printf("x || y: %t\n", x || y)  // 逻辑或（OR）
	fmt.Printf("!x:    %t\n", !x)       // 逻辑非（NOT）

	fmt.Println("\n复杂逻辑表达式:")
	fmt.Printf("x && y || z: %t\n", x && y || z)
	fmt.Printf("(x || y) && z: %t\n", (x || y) && z)
	fmt.Printf("!x && !y: %t\n", !x && !y)

	// 短路求值
	fmt.Println("\n短路求值示例:")
	a, b := 5, 10
	// 如果第一个条件为false，不会执行第二个条件
	result := (a > b) && (b/a > 1) // 不会执行除法，因为 a > b 为 false
	fmt.Printf("(a > b) && (b/a > 1): %t\n", result)

	// 如果第一个条件为true，不会执行第二个条件
	result = (a < b) || (a/b > 1) // 不会执行除法，因为 a < b 为 true
	fmt.Printf("(a < b) || (a/b > 1): %t\n", result)
}

// bitwiseOperators 位运算符
func bitwiseOperators() {
	fmt.Println("\n--- 位运算符 ---")

	var a, b uint8 = 12, 10 // 1100, 1010

	fmt.Printf("a = %d (%08b)\n", a, a)
	fmt.Printf("b = %d (%08b)\n", b, b)
	fmt.Printf("a & b = %d (%08b)  // 按位与\n", a&b, a&b)
	fmt.Printf("a | b = %d (%08b)  // 按位或\n", a|b, a|b)
	fmt.Printf("a ^ b = %d (%08b)  // 按位异或\n", a^b, a^b)
	fmt.Printf("<< 1  = %d (%08b)  // 左移\n", a<<1, a<<1)
	fmt.Printf(">> 1  = %d (%08b)  // 右移\n", a>>1, a>>1)
	fmt.Printf("&^ b = %d (%08b)  // 位清除\n", a&^b, a&^b)

	// 位运算的实际应用
	fmt.Println("\n位运算实际应用:")
	fmt.Printf("设置第3位: %d -> %d\n", a, a|8)
	fmt.Printf("清除第2位: %d -> %d\n", a, a&^4)
	fmt.Printf("切换第1位: %d -> %d\n", a, a^2)
	fmt.Printf("检查第4位: %t\n", (a&8) != 0)
}

// assignmentOperators 赋值运算符
func assignmentOperators() {
	fmt.Println("\n--- 赋值运算符 ---")

	var a int = 10
	fmt.Printf("初始值: a = %d\n", a)

	a += 5  // 等价于 a = a + 5
	fmt.Printf("a += 5  -> a = %d\n", a)

	a -= 3  // 等价于 a = a - 3
	fmt.Printf("a -= 3  -> a = %d\n", a)

	a *= 2  // 等价于 a = a * 2
	fmt.Printf("a *= 2  -> a = %d\n", a)

	a /= 4  // 等价于 a = a / 4
	fmt.Printf("a /= 4  -> a = %d\n", a)

	a %= 3  // 等价于 a = a % 3
	fmt.Printf("a %%= 3 -> a = %d\n", a)

	// 位运算赋值
	var b uint8 = 0x0F // 00001111
	fmt.Printf("\n初始值: b = %08b\n", b)
	b |= 0xF0 // 按位或赋值
	fmt.Printf("b |= 0xF0 -> b = %08b\n", b)
	b &= 0xCC // 按位与赋值
	fmt.Printf("b &= 0xCC -> b = %08b\n", b)
	b ^= 0x55 // 按位异或赋值
	fmt.Printf("b ^= 0x55 -> b = %08b\n", b)
	b <<= 2   // 左移赋值
	fmt.Printf("b <<= 2  -> b = %08b\n", b)
	b >>= 1   // 右移赋值
	fmt.Printf("b >>= 1  -> b = %08b\n", b)
}

// specialOperators 特殊运算符
func specialOperators() {
	fmt.Println("\n--- 特殊运算符 ---")

	// 地址运算符 &
	fmt.Println("地址运算符 &:")
	x := 42
	fmt.Printf("x = %d\n", x)
	fmt.Printf("&x = %p\n", &x)

	// 指针解引用运算符 *
	fmt.Println("\n指针解引用运算符 *:")
	ptr := &x
	fmt.Printf("ptr = %p\n", ptr)
	fmt.Printf("*ptr = %d\n", *ptr)

	// 通道操作 <- （简单示例）
	fmt.Println("\n通道操作:")
	ch := make(chan int, 1)
	ch <- 42        // 发送到通道
	value := <-ch   // 从通道接收
	fmt.Printf("通道值: %d\n", value)

	// 类型断言 .(type)
	fmt.Println("\n类型断言:")
	var i interface{} = "Hello, Go"
	if str, ok := i.(string); ok {
		fmt.Printf("类型断言成功: %s\n", str)
	}

	// 类型选择 switch x.(type)
	var val interface{} = 123
	switch v := val.(type) {
	case int:
		fmt.Printf("整型值: %d\n", v)
	case string:
		fmt.Printf("字符串值: %s\n", v)
	case float64:
		fmt.Printf("浮点值: %.2f\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// operatorPrecedence 运算符优先级
func operatorPrecedence() {
	fmt.Println("\n--- 运算符优先级 ---")

	// 优先级从高到低：
	// 1. () [] .
	// 2. ! ~ * & <-
	// 3. * / % << >> & &^
	// 4. + - | ^
	// 5. == != < <= > >=
	// 6. &&
	// 7. ||

	fmt.Println("运算符优先级示例:")

	a, b, c := 2, 3, 4

	// 没有括号，按优先级计算
	result1 := a + b*c
	fmt.Printf("a + b*c = %d + %d*%d = %d\n", a, b, c, result1)

	// 使用括号改变优先级
	result2 := (a + b) * c
	fmt.Printf("(a + b)*c = (%d + %d)*%d = %d\n", a, b, c, result2)

	// 逻辑运算符优先级
	x, y, z := true, false, true
	result3 := x || y && z      // && 优先级高于 ||
	fmt.Printf("x || y && z = %t || %t && %t = %t\n", x, y, z, result3)

	result4 := (x || y) && z    // 使用括号改变优先级
	fmt.Printf("(x || y) && z = (%t || %t) && %t = %t\n", x, y, z, result4)

	// 位运算符优先级
	m, n, o := uint8(5), uint8(3), uint8(2) // 0101, 0011, 0010
	result5 := m | n & o       // & 优先级高于 |
	fmt.Printf("%d | %d & %d = %08b | %08b & %08b = %08b (%d)\n",
		m, n, o, m, n, o, result5, result5)

	result6 := (m | n) & o     // 使用括号改变优先级
	fmt.Printf("(%d | %d) & %d = (%08b | %08b) & %08b = %08b (%d)\n",
		m, n, o, m, n, o, result6, result6)
}

// demonstratePracticalUsage 实际应用示例
func demonstratePracticalUsage() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 1. 奇偶数判断
	num := 7
	fmt.Printf("%d 是 %s\n", num, func() string {
		if num%2 == 0 {
			return "偶数"
		}
		return "奇数"
	}())

	// 2. 闰年判断
	year := 2024
	fmt.Printf("%d年 %s闰年\n", year, func() string {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			return "是"
		}
		return "不是"
	}())

	// 3. 交换变量（不使用临时变量）
	a, b := 10, 20
	fmt.Printf("交换前: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("交换后: a = %d, b = %d\n", a, b)

	// 4. 乘法优化（使用位运算）
	fmt.Printf("乘法优化: 15 * 8 = %d (使用位运算: %d)\n", 15*8, 15<<3)
	fmt.Printf("除法优化: 100 / 8 = %d (使用位运算: %d)\n", 100/8, 100>>3)
}