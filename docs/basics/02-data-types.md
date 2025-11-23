---
title: 数据类型
difficulty: beginner
duration: "3-4小时"
prerequisites: ["变量与常量"]
tags: ["数据类型", "类型转换", "类型推断"]
---

# 数据类型

Go 语言是一门静态类型语言，变量的类型在编译时就确定。理解 Go 的类型系统是编写高质量 Go 代码的基础。

## 📋 学习目标

- [ ] 理解 Go 语言的基本数据类型
- [ ] 掌握类型转换的方法
- [ ] 学会使用类型推断
- [ ] 了解自定义类型
- [ ] 掌握类型别名和类型定义

## 🔍 基本数据类型

### 整数类型

Go 提供了多种整数类型，以适应不同的使用场景：

```go
package main

import "fmt"

func main() {
	// 有符号整数
	var i8 int8 = 127        // -128 到 127
	var i16 int16 = 32767     // -32768 到 32767
	var i32 int32 = 2147483647  // -2147483648 到 2147483647
	var i64 int64 = 9223372036854775807
	
	// 无符号整数
	var ui8 uint8 = 255       // 0 到 255
	var ui16 uint16 = 65535   // 0 到 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	
	// 平台相关的类型
	var i int = 42           // 32位系统为int32，64位系统为int64
	var u uint = 42          // 32位系统为uint32，64位系统为uint64
	var ptr uintptr = 0x123456 // 足够存放指针的整数
	
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n", i8, i16, i32, i64)
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d, uint64: %d\n", ui8, ui16, ui32, ui64)
}
```

**输出：**
```
int8: 127, int16: 32767, int32: 2147483647, int64: 9223372036854775807
uint8: 255, uint16: 65535, uint32: 4294967295, uint64: 18446744073709551615
```

### 浮点数类型

```go
package main

import "fmt"

func main() {
	var f32 float32 = 3.14159
	var f64 float64 = 2.718281828459045
	
	fmt.Printf("float32: %f, float64: %f\n", f32, f64)
	fmt.Printf("float32 精度: %.10f\n", f32)
	fmt.Printf("float64 精度: %.15f\n", f64)
	
	// 浮点数特殊值
	var posInf = 1.0 / 0.0   // 正无穷
	var negInf = -1.0 / 0.0  // 负无穷
	var nan = 0.0 / 0.0     // 非数字
	
	fmt.Printf("正无穷: %f, 负无穷: %f, NaN: %f\n", posInf, negInf, nan)
}
```

### 布尔类型

```go
package main

import "fmt"

func main() {
	var isTrue bool = true
	var isFalse bool = false
	var defaultBool bool // 默认值为 false
	
	fmt.Printf("true: %t, false: %t, default: %t\n", isTrue, isFalse, defaultBool)
	
	// 布尔运算
	fmt.Printf("true && false: %t\n", isTrue && isFalse)
	fmt.Printf("true || false: %t\n", isTrue || isFalse)
	fmt.Printf("!true: %t\n", !isTrue)
}
```

### 字符串类型

```go
package main

import "fmt"

func main() {
	// 字符串声明
	var str1 string = "Hello"
	str2 := "世界"
	str3 := `多行字符串
可以包含换行符
和特殊字符 "引号"`
	
	fmt.Println(str1, str2)
	fmt.Println(str3)
	
	// 字符串操作
	combined := str1 + " " + str2
	fmt.Println("拼接:", combined)
	
	// 字符串长度
	fmt.Printf("字节长度: %d\n", len(combined))
	
	// 获取字符数量（Unicode）
	runeCount := 0
	for range combined {
		runeCount++
	}
	fmt.Printf("字符数量: %d\n", runeCount)
	
	// 字符串切片
	fmt.Printf("前3个字符: %s\n", combined[:3])
}
```

## 🔤 字符和字节类型

### rune 类型（字符）

```go
package main

import "fmt"

func main() {
	// rune 实际上是 int32 的别名
	var r1 rune = 'A'        // ASCII 字符
	var r2 rune = '中'       // Unicode 字符
	var r3 rune = '\u4e2d'   // Unicode 转义
	var r4 rune = '\U0001f600' // emoji 😊
	
	fmt.Printf("字符: %c, Unicode: %U\n", r1, r1)
	fmt.Printf("字符: %c, Unicode: %U\n", r2, r2)
	fmt.Printf("字符: %c, Unicode: %U\n", r3, r3)
	fmt.Printf("字符: %c, Unicode: %U\n", r4, r4)
	
	// 字符串遍历
	text := "Hello世界"
	fmt.Println("按字符遍历:")
	for i, char := range text {
		fmt.Printf("位置: %d, 字符: %c, Unicode: %U\n", i, char, char)
	}
}
```

### byte 类型

```go
package main

import "fmt"

func main() {
	// byte 实际上是 uint8 的别名
	var b1 byte = 65        // 十进制
	var b2 byte = 0x41      // 十六进制
	var b3 byte = 'A'       // 字符
	
	fmt.Printf("b1: %d, 字符: %c\n", b1, b1)
	fmt.Printf("b2: %d, 字符: %c\n", b2, b2)
	fmt.Printf("b3: %d, 字符: %c\n", b3, b3)
	
	// 字符串转字节数组
	str := "Hello"
	bytes := []byte(str)
	fmt.Printf("字符串 %s 的字节: %v\n", str, bytes)
	
	// 字节数组转字符串
	newStr := string(bytes)
	fmt.Printf("字节数组 %v 转字符串: %s\n", bytes, newStr)
}
```

## 🏗️ 复合数据类型

### 数组

```go
package main

import "fmt"

func main() {
	// 数组声明和初始化
	var arr1 [5]int                    // 零值数组
	var arr2 = [5]int{1, 2, 3, 4, 5} // 完整初始化
	arr3 := [5]int{1, 2}              // 部分初始化，其余为零值
	arr4 := [...]int{1, 2, 3, 4, 5}  // 自动推断长度
	arr5 := [5]int{0: 1, 4: 5}       // 指定索引初始化
	
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println("arr4:", arr4)
	fmt.Println("arr5:", arr5)
	
	// 数组操作
	fmt.Printf("arr2 长度: %d\n", len(arr2))
	fmt.Printf("arr2[1]: %d\n", arr2[1])
	
	// 遍历数组
	fmt.Println("遍历 arr2:")
	for i, v := range arr2 {
		fmt.Printf("索引: %d, 值: %d\n", i, v)
	}
}
```

### 切片

```go
package main

import "fmt"

func main() {
	// 切片声明和初始化
	var slice1 []int                    // nil 切片
	var slice2 = []int{1, 2, 3, 4, 5}  // 直接初始化
	slice3 := make([]int, 5)             // 使用 make 创建
	slice4 := make([]int, 5, 10)         // 指定长度和容量
	
	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))
	
	// 切片操作
	slice2 = append(slice2, 6, 7)  // 追加元素
	fmt.Println("追加后:", slice2)
	
	subSlice := slice2[1:4]        // 切片操作
	fmt.Println("子切片:", subSlice)
	
	// 从数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice5 := arr[1:4]
	fmt.Println("从数组创建:", slice5)
}
```

## 🔄 类型转换

### 基本类型转换

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数值类型转换
	var i int = 42
	var f float64 = float64(i)
	var b byte = byte(i)
	
	fmt.Printf("int: %d, float64: %f, byte: %d\n", i, f, b)
	
	// 字符串与数值转换
	str := strconv.Itoa(i)           // int 转 string
	newI, _ := strconv.Atoi(str)     // string 转 int
	
	fmt.Printf("int->string: %s, string->int: %d\n", str, newI)
	
	// 字符串与浮点数转换
	f := 3.14159
	strFloat := strconv.FormatFloat(f, 'f', 2, 64)
	newF, _ := strconv.ParseFloat(strFloat, 64)
	
	fmt.Printf("float64->string: %s, string->float64: %f\n", strFloat, newF)
	
	// 注意：不同类型之间不能直接运算
	// var result = i + f  // 编译错误！
	var result = float64(i) + f // 正确做法
	fmt.Printf("42 + 3.14159 = %f\n", result)
}
```

## 🏷️ 自定义类型

### 类型定义

```go
package main

import "fmt"

// 定义新类型
type UserID int
type UserName string
type Score float64

// 为自定义类型添加方法
func (u UserID) String() string {
	return fmt.Sprintf("用户#%d", u)
}

func (n UserName) Greet() string {
	return fmt.Sprintf("你好，%s！", n)
}

func main() {
	var id UserID = 1001
	var name UserName = "张三"
	var score Score = 95.5
	
	fmt.Println(id.String())
	fmt.Println(name.Greet())
	fmt.Printf("分数: %.1f\n", score)
	
	// 类型转换（即使底层类型相同，也是不同类型）
	var normalInt int = 42
	var customID UserID = UserID(normalInt) // 需要显式转换
	fmt.Printf("普通int: %d, UserID: %s\n", normalInt, customID)
}
```

### 类型别名

```go
package main

import "fmt"

// 类型别名（Go 1.9+）
type MyInt = int      // 类型别名
type MyString = string

func main() {
	var i MyInt = 42
	var s MyString = "Hello"
	
	fmt.Printf("MyInt: %d, MyString: %s\n", i, s)
	
	// 类型别名与原类型完全相同，可以直接赋值
	var normalInt int = i  // 不需要显式转换
	var normalStr string = s
	
	fmt.Printf("转换后的: %d, %s\n", normalInt, normalStr)
}
```

## 🧪 类型检查和断言

### 类型推断

```go
package main

import "fmt"

func main() {
	// 自动类型推断
	i := 42                    // int
	f := 3.14                  // float64
	s := "Hello"                // string
	b := true                   // bool
	arr := [3]int{1, 2, 3}     // [3]int
	slice := []int{1, 2, 3}     // []int
	
	fmt.Printf("类型: %T, 值: %v\n", i, i)
	fmt.Printf("类型: %T, 值: %v\n", f, f)
	fmt.Printf("类型: %T, 值: %v\n", s, s)
	fmt.Printf("类型: %T, 值: %v\n", b, b)
	fmt.Printf("类型: %T, 值: %v\n", arr, arr)
	fmt.Printf("类型: %T, 值: %v\n", slice, slice)
}
```

## 🏃‍♂️ 实践练习

### 练习 1：温度转换器

编写一个程序，将摄氏温度转换为华氏温度：

```go
// 要求：
// 1. 定义 Celsius 和 Fahrenheit 类型
// 2. 实现转换函数
// 3. 实现字符串输出方法
```

### 练习 2：学生成绩统计

```go
// 要求：
// 1. 定义 Student 结构体（下一章会详细讲解）
// 2. 使用合适的类型存储姓名、年龄、成绩
// 3. 计算平均分
```

### 练习 3：字符串处理

```go
// 要求：
// 1. 统计字符串中的字符数、字节数
// 2. 将字符串反转
// 3. 检查是否为回文字符串
```

## 🤔 思考题

1. **为什么 Go 要区分这么多种整数类型？**
   - 提示：考虑内存使用和性能

2. **rune 和 byte 的区别是什么？**
   - 提示：考虑 Unicode 编码

3. **什么时候应该使用自定义类型？**
   - 提示：考虑代码可读性和类型安全

## 📚 扩展阅读

- [Go 语言规范 - 类型](https://golang.org/ref/spec#Types)
- [Effective Go - 数据](https://golang.org/doc/effective_go.html#data)
- [Go by Example - 类型](https://gobyexample.com/values)

## ⏭️ 下一章节

[运算符](./03-operators.md) → 学习 Go 语言的各种运算符和表达式

---

**💡 小贴士**: 选择合适的数据类型不仅能提高程序性能，还能避免潜在的错误。记住：**类型安全是 Go 语言的重要特性！**
