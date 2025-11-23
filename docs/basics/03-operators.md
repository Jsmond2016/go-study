---
title: 运算符
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["运算符", "表达式", "优先级"]
---

# 运算符

运算符是编程语言的基本构建块，用于对数据进行操作和计算。Go 语言提供了丰富的运算符集合。

## 📋 学习目标

- [ ] 掌握 Go 的算术运算符
- [ ] 理解关系和逻辑运算符
- [ ] 学会使用位运算符
- [ ] 掌握赋值和复合运算符
- [ ] 理解运算符优先级

## ➕ 算术运算符

### 基本算术运算

```go
package main

import "fmt"

func main() {
	// 基本算术运算
	a := 10
	b := 3
	
	fmt.Printf("%d + %d = %d\n", a, b, a+b)   // 加法
	fmt.Printf("%d - %d = %d\n", a, b, a-b)   // 减法
	fmt.Printf("%d * %d = %d\n", a, b, a*b)   // 乘法
	fmt.Printf("%d / %d = %d\n", a, b, a/b)   // 除法
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // 取余
	
	// 浮点数运算
	x := 10.5
	y := 3.2
	fmt.Printf("%.1f + %.1f = %.2f\n", x, y, x+y)
	fmt.Printf("%.1f / %.1f = %.2f\n", x, y, x/y)
	
	// 除零处理
	fmt.Printf("10 / 0 = %v\n", 10/0)    // +Inf (正无穷)
	fmt.Printf("0 / 0 = %v\n", 0/0)      // NaN (非数字)
}
```

### 自增自减运算符

```go
package main

import "fmt"

func main() {
	// 自增运算
	count := 5
	fmt.Printf("count: %d\n", count)
	count++ // count = count + 1
	fmt.Printf("count++: %d\n", count)
	
	// 自减运算
	counter := 10
	fmt.Printf("counter: %d\n", counter)
	counter-- // counter = counter - 1
	fmt.Printf("counter--: %d\n", counter)
	
	// 注意：Go 没有 ++count 或 --count 的前缀形式
	// ++count // 编译错误！
	// --counter // 编译错误！
	
	// 在表达式中使用
	score := 80
	fmt.Printf("原分数: %d\n", score)
	score++ // 分数增加
	fmt.Printf("新分数: %d\n", score)
}
```

## ⚖️ 关系运算符

### 比较运算

```go
package main

import "fmt"

func main() {
	// 数值比较
	a := 10
	b := 20
	
	fmt.Printf("%d == %d: %t\n", a, b, a == b) // 等于
	fmt.Printf("%d != %d: %t\n", a, b, a != b) // 不等于
	fmt.Printf("%d < %d: %t\n", a, b, a < b)   // 小于
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b) // 小于等于
	fmt.Printf("%d > %d: %t\n", a, b, a > b)   // 大于
	fmt.Printf("%d >= %d: %t\n", a, b, a >= b) // 大于等于
	
	// 字符串比较
	str1 := "hello"
	str2 := "world"
	str3 := "hello"
	
	fmt.Printf("\"%s\" == \"%s\": %t\n", str1, str2, str1 == str2)
	fmt.Printf("\"%s\" == \"%s\": %t\n", str1, str3, str1 == str3)
	
	// 浮点数比较（注意精度问题）
	f1 := 0.1 + 0.2
	f2 := 0.3
	fmt.Printf("%.1f == %.1f: %t\n", f1, f2, f1 == f2) // 可能是 false
	
	// 正确的浮点数比较
	epsilon := 0.000001
	isEqual := (f1 - f2) < epsilon && (f2 - f1) < epsilon
	fmt.Printf("浮点数比较: %t\n", isEqual)
}
```

## 🔗 逻辑运算符

### 布尔逻辑

```go
package main

import "fmt"

func main() {
	// 基本逻辑运算
	a := true
	b := false
	c := true
	
	fmt.Printf("%t && %t: %t\n", a, b, a && b) // 逻辑与
	fmt.Printf("%t && %t: %t\n", a, c, a && c) // 逻辑与
	fmt.Printf("%t || %t: %t\n", a, b, a || b) // 逻辑或
	fmt.Printf("%t || %t: %t\n", a, c, a || c) // 逻辑或
	fmt.Printf("!%t: %t\n", a, !a)           // 逻辑非
	
	// 复合逻辑表达式
	age := 25
	hasLicense := true
	hasInsurance := true
	
	// 判断是否可以开车
	canDrive := age >= 18 && hasLicense
	fmt.Printf("可以开车: %t\n", canDrive)
	
	// 判断是否可以租车
	canRentCar := age >= 21 && hasLicense && hasInsurance
	fmt.Printf("可以租车: %t\n", canRentCar)
	
	// 短路求值
	result1 := false && someFunction()
	fmt.Printf("false && func(): %t\n", result1)
	
	result2 := true || someFunction()
	fmt.Printf("true || func(): %t\n", result2)
}

func someFunction() bool {
	fmt.Println("someFunction 被调用")
	return true
}
```

## 🔢 位运算符

### 二进制操作

```go
package main

import "fmt"

func main() {
	// 基本位运算
	a := 0b1010 // 10
	b := 0b0110 // 6
	
	fmt.Printf("a = %04b (%d)\n", a, a)
	fmt.Printf("b = %04b (%d)\n", b, b)
	
	fmt.Printf("%04b & %04b = %04b (%d)\n", a, b, a&b, a&b)  // 按位与
	fmt.Printf("%04b | %04b = %04b (%d)\n", a, b, a|b, a|b)  // 按位或
	fmt.Printf("%04b ^ %04b = %04b (%d)\n", a, b, a^b, a^b)  // 按位异或
	fmt.Printf("^%04b = %04b (%d)\n", a, ^a, ^a)          // 按位取反
	
	// 移位运算
	x := 8 // 0b1000
	fmt.Printf("%04b << 2 = %04b (%d)\n", x, x<<2, x<<2) // 左移
	fmt.Printf("%04b >> 1 = %04b (%d)\n", x, x>>1, x>>1) // 右移
	
	// 实际应用：权限管理
	const (
		READ   = 1 << 0 // 1  (0b0001)
		WRITE  = 1 << 1 // 2  (0b0010)
		EXECUTE = 1 << 2 // 4  (0b0100)
		DELETE = 1 << 3 // 8  (0b1000)
	)
	
	// 用户权限
	var permissions uint8 = READ | WRITE // 3 (0b0011)
	fmt.Printf("初始权限: %04b (%d)\n", permissions, permissions)
	
	// 添加执行权限
	permissions |= EXECUTE
	fmt.Printf("添加执行权限: %04b (%d)\n", permissions, permissions)
	
	// 检查权限
	hasRead := (permissions & READ) != 0
	hasWrite := (permissions & WRITE) != 0
	hasExecute := (permissions & EXECUTE) != 0
	
	fmt.Printf("可读: %t, 可写: %t, 可执行: %t\n", hasRead, hasWrite, hasExecute)
	
	// 移除写权限
	permissions &= ^WRITE
	fmt.Printf("移除写权限: %04b (%d)\n", permissions, permissions)
}
```

## 🔄 赋值运算符

### 基本赋值

```go
package main

import "fmt"

func main() {
	// 基本赋值
	var x int
	x = 10
	fmt.Printf("x = %d\n", x)
	
	// 多重赋值
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	
	// 交换值
	x, y := 10, 20
	fmt.Printf("交换前: x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("交换后: x = %d, y = %d\n", x, y)
}
```

### 复合赋值运算符

```go
package main

import "fmt"

func main() {
	// 复合赋值运算符
	x := 10
	
	x += 5  // 等价于 x = x + 5
	fmt.Printf("x += 5: %d\n", x)
	
	x -= 3  // 等价于 x = x - 3
	fmt.Printf("x -= 3: %d\n", x)
	
	x *= 2  // 等价于 x = x * 2
	fmt.Printf("x *= 2: %d\n", x)
	
	x /= 4  // 等价于 x = x / 4
	fmt.Printf("x /= 4: %d\n", x)
	
	x %= 3  // 等价于 x = x % 3
	fmt.Printf("x %% 3: %d\n", x)
	
	// 位运算复合赋值
	var flags uint8 = 0b0001
	
	flags |= 0b0010  // 等价于 flags = flags | 0b0010
	fmt.Printf("flags |= 0b0010: %04b\n", flags)
	
	flags &= ^0b0010  // 等价于 flags = flags & ^0b0010
	fmt.Printf("flags &= ^0b0010: %04b\n", flags)
	
	flags ^= 0b0100   // 等价于 flags = flags ^ 0b0100
	fmt.Printf("flags ^= 0b0100: %04b\n", flags)
	
	flags <<= 2       // 等价于 flags = flags << 2
	fmt.Printf("flags <<= 2: %04b\n", flags)
	
	flags >>= 1       // 等价于 flags = flags >> 1
	fmt.Printf("flags >>= 1: %04b\n", flags)
}
```

## 🎯 运算符优先级

### 优先级表

```go
package main

import "fmt"

func main() {
	// 优先级示例
	result1 := 2 + 3 * 4     // 先乘后加：2 + (3 * 4) = 14
	result2 := (2 + 3) * 4   // 先加后乘：(2 + 3) * 4 = 20
	
	fmt.Printf("2 + 3 * 4 = %d\n", result1)
	fmt.Printf("(2 + 3) * 4 = %d\n", result2)
	
	// 逻辑运算优先级
	a := true
	b := false
	c := true
	
	result3 := a || b && c  // 先 && 后 ||：a || (b && c) = true
	result4 := (a || b) && c  // 先 || 后 &&：(a || b) && c = true
	
	fmt.Printf("true || false && true: %t\n", result3)
	fmt.Printf("(true || false) && true: %t\n", result4)
	
	// 复杂表达式
	age := 20
	score := 85
	isActive := true
	
	// 判断是否优秀学生
	isExcellent := isActive && age >= 18 && score >= 80
	fmt.Printf("优秀学生: %t\n", isExcellent)
	
	// 使用括号明确优先级
	isExcellent2 := isActive && (age >= 18) && (score >= 80)
	fmt.Printf("优秀学生2: %t\n", isExcellent2)
	
	// 位运算优先级
	flags := 0b0001
	mask := 0b0110
	
	result5 := flags | mask & 0b0100  // 先 & 后 |
	fmt.Printf("flags | mask & 0b0100: %04b\n", result5)
	
	result6 := (flags | mask) & 0b0100  // 先 | 后 &
	fmt.Printf("(flags | mask) & 0b0100: %04b\n", result6)
}
```

### 优先级规则

从高到低的运算符优先级：

1. `() [] .` - 括号、索引、成员访问
2. `++ --` - 自增自减（后缀）
3. `* / % << >> & & ^` - 乘除取余移位位运算
4. `+ - | ^` - 加减异或
5. `== != < <= > >=` - 关系运算
6. `&&` - 逻辑与
7. `||` - 逻辑或
8. `=` `+= -= *= /= %= &= ^= <<= >>= |=` - 赋值运算

## 🧪 特殊运算符

### 地址和指针运算符

```go
package main

import "fmt"

func main() {
	// 取地址运算符 &
	x := 42
	ptr := &x // 获取 x 的内存地址
	fmt.Printf("x 的值: %d\n", x)
	fmt.Printf("x 的地址: %p\n", ptr)
	
	// 指针解引用运算符 *
	value := *ptr // 获取指针指向的值
	fmt.Printf("指针指向的值: %d\n", value)
	
	// 修改指针指向的值
	*ptr = 100
	fmt.Printf("修改后 x 的值: %d\n", x)
}
```

### 通道运算符

```go
package main

import "fmt"

func main() {
	// 通道操作
	ch := make(chan int, 2)
	
	// 发送数据到通道
	ch <- 42
	
	// 从通道接收数据
	value := <-ch
	fmt.Printf("从通道接收: %d\n", value)
	
	// 非阻塞接收（多值返回）
	ch <- 100
	select {
	case v := <-ch:
		fmt.Printf("select 接收: %d\n", v)
	default:
		fmt.Println("通道无数据")
	}
}
```

## 🏃‍♂️ 实践练习

### 练习 1：计算器实现

```go
// 要求：
// 1. 实现基本四则运算函数
// 2. 处理除零情况
// 3. 支持复合赋值
// 4. 添加运算优先级测试
```

参考实现：
```go
package main

import "fmt"

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("除零错误")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("不支持的操作符: %s", op)
	}
}

func main() {
	// 测试运算
	result, err := calculate(10, 5, "+")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 + 5 = %.2f\n", result)
	}
	
	// 测试除零
	_, err = calculate(10, 0, "/")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
}
```

### 练习 2：位运算权限系统

```go
// 要求：
// 1. 定义权限常量
// 2. 实现权限检查函数
// 3. 实现权限添加/移除函数
// 4. 使用位运算操作权限
```

### 练习 3：表达式计算器

```go
// 要求：
// 1. 解析数学表达式
// 2. 考虑运算符优先级
// 3. 支持括号
// 4. 处理错误表达式
```

## 🤔 思考题

1. **为什么 `true || false && true` 的结果是 `true`？**
   - 提示：考虑运算符优先级和短路求值

2. **什么时候应该使用括号来明确优先级？**
   - 提示：考虑代码可读性

3. **位运算在实际编程中有哪些应用场景？**
   - 提示：考虑权限、状态标记、性能优化

4. **浮点数比较为什么需要特殊处理？**
   - 提示：考虑计算机表示浮点数的方式

## 📚 扩展阅读

- [Go 语言规范 - 运算符](https://golang.org/ref/spec#Operators)
- [Effective Go - 运算符](https://golang.org/doc/effective_go.html#operators)
- [Go by Example - 运算符](https://gobyexample.com/operators)

## ⏭️ 下一章节

[控制流程](./04-control-flow.md) → 学习条件语句、循环语句等控制流程

---

**💡 小贴士**: 理解运算符优先级和结合性对于编写正确的表达式至关重要。在复杂表达式中，**善用括号可以提高代码可读性**！
