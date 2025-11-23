---
title: 切片
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["变量与常量", "数据类型", "数组"]
tags: ["切片", "动态数组", "扩容", "内存管理"]
---

# 切片

切片（Slice）是 Go 语言中最重要的数据结构之一，提供了动态数组的功能。切片比数组更灵活，是 Go 程序中处理序列数据的首选方式。

## 📋 学习目标

- [ ] 理解切片的概念和用途
- [ ] 掌握切片的声明和初始化
- [ ] 理解切片的底层实现原理
- [ ] 学会切片的常用操作
- [ ] 掌握切片的扩容机制
- [ ] 理解切片的内存管理
- [ ] 学会使用多维切片

## 🎯 切片基础

### 什么是切片

切片是对数组的抽象，提供了动态大小的序列。切片包含三个部分：
- **指针**：指向底层数组
- **长度（length）**：切片中元素的数量
- **容量（capacity）**：底层数组从切片起始位置到数组末尾的元素数量

```go
package main

import "fmt"

func main() {
	// 声明切片（nil 切片）
	var slice1 []int
	fmt.Printf("nil 切片: %v (len=%d, cap=%d, isNil=%t)\n",
		slice1, len(slice1), cap(slice1), slice1 == nil)

	// 直接初始化
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("直接初始化: %v (len=%d, cap=%d)\n",
		slice2, len(slice2), cap(slice2))

	// 使用 make 创建
	slice3 := make([]int, 5)        // 长度为5，容量为5
	slice4 := make([]int, 3, 10)    // 长度为3，容量为10
	fmt.Printf("make 创建: %v (len=%d, cap=%d)\n",
		slice3, len(slice3), cap(slice3))
	fmt.Printf("make 指定容量: %v (len=%d, cap=%d)\n",
		slice4, len(slice4), cap(slice4))
}
```

### 从数组创建切片

```go
package main

import "fmt"

func main() {
	// 创建数组
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 从数组创建切片
	slice1 := arr[2:7]  // 包含索引2到6的元素
	fmt.Printf("arr[2:7] = %v\n", slice1)

	// 切片的切片
	slice2 := slice1[1:4]  // 从现有切片创建新切片
	fmt.Printf("slice1[1:4] = %v\n", slice2)

	// 完整切片表达式
	slice3 := arr[2:7:8]  // [low:high:max]，容量为 max-low
	fmt.Printf("arr[2:7:8] = %v (len=%d, cap=%d)\n",
		slice3, len(slice3), cap(slice3))
}
```

## 🔍 切片的底层原理

### 切片的结构

切片本身不存储数据，而是引用底层数组的一部分。

```go
package main

import "fmt"

func main() {
	// 创建一个切片
	s := make([]int, 3, 5)
	s[0], s[1], s[2] = 10, 20, 30

	fmt.Printf("原始切片: %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	// 创建新切片，共享底层数组
	s2 := s[1:3]
	fmt.Printf("新切片: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	// 修改新切片会影响原切片（共享底层数组）
	s2[0] = 99
	fmt.Printf("修改新切片后:\n")
	fmt.Printf("原切片: %v\n", s)
	fmt.Printf("新切片: %v\n", s2)
}
```

### 切片的零值

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("零值切片: %v\n", s)
	fmt.Printf("是否为 nil: %t\n", s == nil)
	fmt.Printf("长度: %d, 容量: %d\n", len(s), cap(s))

	// nil 切片可以安全使用
	s = append(s, 1, 2, 3)
	fmt.Printf("append 后: %v\n", s)
}
```

## 🛠️ 切片操作

### 添加元素（append）

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("原始切片: %v (len=%d, cap=%d)\n",
		slice, len(slice), cap(slice))

	// 添加单个元素
	slice = append(slice, 4)
	fmt.Printf("添加单个元素: %v (len=%d, cap=%d)\n",
		slice, len(slice), cap(slice))

	// 添加多个元素
	slice = append(slice, 5, 6, 7)
	fmt.Printf("添加多个元素: %v\n", slice)

	// 添加另一个切片
	another := []int{8, 9, 10}
	slice = append(slice, another...)
	fmt.Printf("添加另一个切片: %v\n", slice)
}
```

### 删除元素

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片: %v\n", slice)

	// 删除索引为 2 的元素
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("删除索引 %d: %v\n", index, slice)

	// 删除第一个元素
	slice = slice[1:]
	fmt.Printf("删除第一个元素: %v\n", slice)

	// 删除最后一个元素
	slice = slice[:len(slice)-1]
	fmt.Printf("删除最后一个元素: %v\n", slice)

	// 删除指定范围的元素
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = append(slice[:2], slice[5:]...)  // 删除索引2到4
	fmt.Printf("删除范围 [2:5]: %v\n", slice)
}
```

### 复制切片

```go
package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}

	// 深拷贝
	copy1 := make([]int, len(original))
	copy(copy1, original)
	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("深拷贝: %v\n", copy1)

	// 修改拷贝不影响原切片
	copy1[0] = 99
	fmt.Printf("修改拷贝后:\n")
	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("深拷贝: %v\n", copy1)

	// 浅拷贝（共享底层数组）
	copy2 := original
	copy2[0] = 88
	fmt.Printf("浅拷贝后:\n")
	fmt.Printf("原切片: %v\n", original)  // 也被修改了
	fmt.Printf("浅拷贝: %v\n", copy2)
}
```

### 切片截取

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// 截取切片
	slice1 := slice[1:4]    // [2,3,4]
	slice2 := slice[2:]     // [3,4,5]
	slice3 := slice[:3]     // [1,2,3]

	fmt.Printf("原切片: %v\n", slice)
	fmt.Printf("slice[1:4]: %v\n", slice1)
	fmt.Printf("slice[2:]: %v\n", slice2)
	fmt.Printf("slice[:3]: %v\n", slice3)

	// 注意：截取的切片共享底层数组
	slice1[0] = 99
	fmt.Printf("修改 slice1[0] 后:\n")
	fmt.Printf("原切片: %v\n", slice)  // 也被修改了
}
```

## 📈 切片扩容

### 扩容机制

当切片容量不足时，Go 会自动扩容。扩容策略：
- 如果容量 < 1024，容量翻倍
- 如果容量 >= 1024，每次增加 25%

```go
package main

import "fmt"

func main() {
	var s []int

	fmt.Println("演示切片扩容:")
	for i := 0; i < 20; i++ {
		oldCap := cap(s)
		s = append(s, i)
		newCap := cap(s)

		if oldCap != newCap {
			fmt.Printf("添加 %d: len=%d, cap=%d -> %d (扩容)\n",
				i, len(s), oldCap, newCap)
		}
	}
}
```

### 预分配容量

```go
package main

import "fmt"

func main() {
	// 不预分配容量
	var s1 []int
	for i := 0; i < 1000; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("不预分配: 最终 cap=%d\n", cap(s1))

	// 预分配容量（性能更好）
	s2 := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		s2 = append(s2, i)
	}
	fmt.Printf("预分配: 最终 cap=%d\n", cap(s2))
}
```

## 💾 内存管理

### 避免内存泄漏

```go
package main

import "fmt"

func main() {
	// 大切片
	large := make([]int, 1000)
	for i := range large {
		large[i] = i
	}

	// 只使用前10个元素
	small := large[:10]
	fmt.Printf("small: len=%d, cap=%d\n", len(small), cap(small))

	// 问题：small 仍然引用整个 large 的底层数组
	// 解决：创建新的切片，只复制需要的元素
	small2 := make([]int, 10)
	copy(small2, large[:10])
	fmt.Printf("small2: len=%d, cap=%d\n", len(small2), cap(small2))

	// 现在 large 可以被垃圾回收
	large = nil
}
```

### 切片作为函数参数

```go
package main

import "fmt"

// 修改切片（会影响原切片）
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

// 追加元素（不会影响原切片，除非重新赋值）
func appendToSlice(s []int) {
	s = append(s, 100)
	fmt.Printf("函数内: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("原始: %v\n", slice)

	// 修改元素会影响原切片
	modifySlice(slice)
	fmt.Printf("修改后: %v\n", slice)

	// 追加元素不会影响原切片（除非重新赋值）
	appendToSlice(slice)
	fmt.Printf("追加后: %v\n", slice)  // 未改变

	// 需要重新赋值
	slice = append(slice, 200)
	fmt.Printf("重新赋值后: %v\n", slice)
}
```

## 🔄 多维切片

### 二维切片

```go
package main

import "fmt"

func main() {
	// 创建二维切片
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
		for j := range matrix[i] {
			matrix[i][j] = i*4 + j + 1
		}
	}

	fmt.Println("二维切片:")
	for i, row := range matrix {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 不规则的二维切片
	jagged := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}

	fmt.Println("\n不规则二维切片:")
	for i, row := range jagged {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 动态添加行和列
	matrix = append(matrix, []int{13, 14, 15, 16})  // 添加新行
	matrix[0] = append(matrix[0], 17)               // 添加列

	fmt.Println("\n添加后的矩阵:")
	for i, row := range matrix {
		fmt.Printf("  行 %d: %v\n", i, row)
	}
}
```

## 🎯 常用操作模式

### 过滤

```go
func filter(slice []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// 使用
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := filter(numbers, func(n int) bool {
	return n%2 == 0
})
```

### 映射转换

```go
func mapSlice(slice []int, transform func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// 使用
numbers := []int{1, 2, 3, 4, 5}
doubled := mapSlice(numbers, func(n int) int {
	return n * 2
})
```

### 查找

```go
func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func indexOf(slice []int, item int) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}
```

### 删除操作

```go
// 删除指定位置的元素
slice := []int{1, 2, 3, 4, 5}
i := 2 // 要删除的索引
slice = append(slice[:i], slice[i+1:]...)  // 删除索引i的元素

// 删除指定范围的元素
slice = append(slice[:i], slice[j:]...)    // 删除索引i到j的元素

// 删除并保持顺序（使用copy）
copy(slice[i:], slice[i+1:])
slice = slice[:len(slice)-1]
```

### 排序操作

```go
import "sort"

// 基本类型排序
numbers := []int{4, 2, 1, 3, 5}
sort.Ints(numbers)               // 升序排序
sort.Sort(sort.Reverse(sort.IntSlice(numbers))) // 降序排序

// 自定义类型排序
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 20},
}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

### 搜索操作

```go
import "sort"

// 线性查找
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    if v == 3 {
        fmt.Printf("找到元素 %d 在位置 %d\n", v, i)
        break
    }
}

// 二分查找（要求切片已排序）
numbers := []int{1, 2, 3, 4, 5}
index := sort.SearchInts(numbers, 3)  // 返回值为2
```

## ⚠️ 常见陷阱

### 1. 在循环中使用 append

```go
// ❌ 错误：可能导致意外的行为
var result []int
for i := 0; i < 3; i++ {
	result = append(result, i)
	// 如果 result 在循环外被其他 goroutine 修改，会有问题
}

// ✅ 正确：预分配容量
result := make([]int, 0, 3)
for i := 0; i < 3; i++ {
	result = append(result, i)
}
```

### 2. 共享底层数组

```go
// 注意：多个切片可能共享底层数组
s1 := []int{1, 2, 3, 4, 5}
s2 := s1[1:4]  // 共享底层数组

s2[0] = 99
fmt.Println(s1)  // [1, 99, 3, 4, 5] - s1 也被修改了

// 如果需要独立，使用 copy
s3 := make([]int, len(s1))
copy(s3, s1)
```

### 3. 切片作为函数参数

```go
// 注意：切片是引用类型，但 append 需要重新赋值
func addElement(s []int, val int) {
	s = append(s, val)  // 不会影响原切片
	// 需要返回新切片
}

// 正确做法
func addElement(s []int, val int) []int {
	return append(s, val)
}
```

## 🏃‍♂️ 实践练习

### 练习 1: 实现栈

使用切片实现一个栈数据结构。

### 练习 2: 切片去重

实现一个函数，去除切片中的重复元素。

### 练习 3: 切片合并

实现一个函数，合并多个切片。

## 🤔 思考题

1. 切片和数组有什么区别？
2. 切片的扩容机制是什么？
3. 什么时候应该预分配切片容量？
4. 如何避免切片的内存泄漏？
5. 切片作为函数参数时需要注意什么？

## 📚 扩展阅读

- [Go 切片官方文档](https://golang.org/ref/spec#Slice_types)
- [切片内部实现](https://github.com/golang/go/blob/master/src/runtime/slice.go)
- [切片最佳实践](https://github.com/golang/go/wiki/SliceTricks)

## ⏭️ 下一章节

[映射](./08-maps.md) → 学习 Go 语言的映射（Map）数据结构

---

**💡 提示**: 切片是 Go 语言中最常用的数据结构，理解其底层原理和内存管理对于编写高效的 Go 程序至关重要！
