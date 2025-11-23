---
title: 数组
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "数据类型", "循环"]
tags: ["数组", "固定长度", "内存布局", "多维数组"]
---

# 数组

数组是 Go 语言中最基本的数据结构，用于存储固定长度的同类型元素序列。理解数组对于掌握 Go 的内存管理和性能优化很重要。

## 📋 学习目标

- [ ] 掌握数组的声明和初始化
- [ ] 理解数组的内存布局
- [ ] 学会使用多维数组
- [ ] 掌握数组的遍历和操作
- [ ] 了解数组与切片的区别
- [ ] 学会使用数组作为函数参数

## 🏗️ 数组基础

### 数组声明和初始化

```go
package main

import "fmt"

func main() {
	// 声明数组（零值）
	var arr1 [5]int
	fmt.Printf("零值数组: %v\n", arr1)
	
	// 声明并初始化
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("完整初始化: %v\n", arr2)
	
	// 自动推断长度
	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("自动长度: %v (长度: %d)\n", arr3, len(arr3))
	
	// 指定索引初始化
	arr4 := [5]int{0: 100, 2: 200, 4: 300}
	fmt.Printf("指定索引: %v\n", arr4)
	
	// 部分初始化（其余为零值）
	arr5 := [5]int{1, 2}
	fmt.Printf("部分初始化: %v\n", arr5)
	
	// 不同类型的数组
	var stringArr [3]string
	var boolArr [4]bool
	var floatArr [3]float64
	
	fmt.Printf("字符串数组: %v\n", stringArr)
	fmt.Printf("布尔数组: %v\n", boolArr)
	fmt.Printf("浮点数组: %v\n", floatArr)
}
```

### 数组的长度和容量

```go
package main

import "fmt"

func main() {
	// 数组的长度是固定的
	arr := [5]int{1, 2, 3, 4, 5}
	
	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("长度: %d\n", len(arr))
	fmt.Printf("容量: %d\n", cap(arr)) // 数组容量等于长度
	
	// 尝试修改长度（编译错误）
	// arr = append(arr, 6) // 编译错误：数组是固定长度的
	
	// 不同长度的数组是不同类型
	var shortArr [3]int
	var longArr [5]int
	
	// shortArr = longArr // 编译错误：类型不匹配
	_ = shortArr // 避免未使用变量错误
	
	fmt.Println("数组长度在编译时就确定了")
}
```

## 🎯 数组操作

### 数组访问和修改

```go
package main

import "fmt"

func main() {
	// 基本访问
	numbers := [5]int{10, 20, 30, 40, 50}
	
	fmt.Printf("numbers[0] = %d\n", numbers[0])  // 10
	fmt.Printf("numbers[4] = %d\n", numbers[4])  // 50
	
	// 修改元素
	numbers[2] = 300
	fmt.Printf("修改后: %v\n", numbers)
	
	// 边界检查
	if len(numbers) > 3 {
		fmt.Printf("numbers[3] = %d\n", numbers[3])
	}
	
	// 超出边界访问（运行时 panic）
	// fmt.Printf("numbers[5] = %d\n", numbers[5]) // panic: index out of range
	
	// 使用常量作为索引
	const lastIndex = 4
	fmt.Printf("最后元素: %d\n", numbers[lastIndex])
	
	// 动态索引访问
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("索引 %d: %d\n", i, numbers[i])
	}
}
```

### 数组遍历

```go
package main

import "fmt"

func main() {
	// 传统 for 循环遍历
	arr := [5]string{"苹果", "香蕉", "橙子", "葡萄", "西瓜"}
	
	fmt.Println("传统 for 循环:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %s\n", i, arr[i])
	}
	
	// for-range 遍历
	fmt.Println("\nfor-range 遍历:")
	for index, value := range arr {
		fmt.Printf("索引 %d: %s\n", index, value)
	}
	
	// 只要值，忽略索引
	fmt.Println("\n只获取值:")
	for _, value := range arr {
		fmt.Printf("水果: %s\n", value)
	}
	
	// 只要索引，忽略值（不常用）
	fmt.Println("\n只获取索引:")
	for index := range arr {
		fmt.Printf("索引: %d\n", index)
	}
}
```

## 📐 多维数组

### 二维数组

```go
package main

import "fmt"

func main() {
	// 二维数组声明和初始化
	var matrix [3][4]int // 3行4列的零值矩阵
	
	// 完整初始化
	matrix2 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	
	fmt.Printf("二维数组: %v\n", matrix2)
	
	// 访问和修改
	matrix2[1][2] = 100
	fmt.Printf("修改后: %v\n", matrix2)
	fmt.Printf("matrix[1][2] = %d\n", matrix2[1][2])
	
	// 遍历二维数组
	fmt.Println("\n遍历二维数组:")
	for i := 0; i < len(matrix2); i++ {
		for j := 0; j < len(matrix2[i]); j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix2[i][j])
		}
	}
	
	// for-range 遍历二维数组
	fmt.Println("\nfor-range 遍历:")
	for i, row := range matrix2 {
		for j, value := range row {
			fmt.Printf("[%d][%d] = %d\n", i, j, value)
		}
	}
}
```

### 三维数组

```go
package main

import "fmt"

func main() {
	// 三维数组（3x2x4）
	cube := [3][2][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		{
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		{
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}
	
	fmt.Printf("三维数组大小: %d x %d x %d\n", 
		len(cube), len(cube[0]), len(cube[0][0]))
	
	// 访问三维数组元素
	fmt.Printf("cube[1][0][2] = %d\n", cube[1][0][2])
	
	// 修改元素
	cube[2][1][3] = 999
	fmt.Printf("修改后 cube[2][1][3] = %d\n", cube[2][1][3])
	
	// 遍历三维数组
	fmt.Println("\n遍历三维数组:")
	for x := 0; x < len(cube); x++ {
		for y := 0; y < len(cube[x]); y++ {
			for z := 0; z < len(cube[x][y]); z++ {
				fmt.Printf("cube[%d][%d][%d] = %d\n", 
					x, y, z, cube[x][y][z])
			}
		}
	}
}
```

## 🔄 数组与函数

### 数组作为函数参数

```go
package main

import "fmt"

// 数组作为参数（需要指定长度）
func processArray(arr [5]int) {
	fmt.Printf("处理数组: %v\n", arr)
	for i, value := range arr {
		arr[i] = value * 2 // 修改的是副本
	}
	fmt.Printf("处理后: %v\n", arr)
}

// 使用指针修改原数组
func modifyArray(arr *[5]int) {
	fmt.Printf("修改原数组: %v\n", *arr)
	for i := 0; i < len(arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
	fmt.Printf("修改后: %v\n", *arr)
}

// 返回数组
func createArray() [5]int {
	return [5]int{1, 1, 2, 3, 5} // 斐波那契数列前5项
}

func main() {
	// 传递数组（值传递）
	original := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("原数组: %v\n", original)
	
	processArray(original)
	fmt.Printf("原数组未被改变: %v\n", original)
	
	// 传递数组指针
	modifyArray(&original)
	fmt.Printf("原数组被改变: %v\n", original)
	
	// 接收返回的数组
	newArray := createArray()
	fmt.Printf("新数组: %v\n", newArray)
}
```

### 数组切片操作

```go
package main

import "fmt"

// 从数组创建切片
func arrayToSlice(arr [5]int) []int {
	// 创建整个数组的切片
	slice1 := arr[:]
	fmt.Printf("整个数组切片: %v (长度: %d, 容量: %d)\n", 
		slice1, len(slice1), cap(slice1))
	
	// 创建部分切片
	slice2 := arr[1:4]
	fmt.Printf("部分切片 [1:4]: %v (长度: %d, 容量: %d)\n", 
		slice2, len(slice2), cap(slice2))
	
	// 创建从开始的切片
	slice3 := arr[:3]
	fmt.Printf("开头切片 [:3]: %v\n", slice3)
	
	// 创建到末尾的切片
	slice4 := arr[2:]
	fmt.Printf("末尾切片 [2:]: %v\n", slice4)
	
	return slice1
}

func main() {
	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原数组: %v\n", numbers)
	
	slice := arrayToSlice(numbers)
	
	// 修改切片会影响原数组
	slice[0] = 100
	fmt.Printf("修改切片后原数组: %v\n", numbers)
	fmt.Printf("修改后的切片: %v\n", slice)
}
```

## 🎯 实际应用

### 数组作为缓存和缓冲区

```go
package main

import "fmt"

// 使用数组作为循环缓冲区
type CircularBuffer struct {
	buffer [5]int
	head   int
	tail   int
	count  int
}

func NewCircularBuffer() *CircularBuffer {
	return &CircularBuffer{}
}

func (cb *CircularBuffer) Push(value int) bool {
	if cb.count >= len(cb.buffer) {
		return false // 缓冲区满
	}
	
	cb.buffer[cb.tail] = value
	cb.tail = (cb.tail + 1) % len(cb.buffer)
	cb.count++
	return true
}

func (cb *CircularBuffer) Pop() (int, bool) {
	if cb.count == 0 {
		return 0, false // 缓冲区空
	}
	
	value := cb.buffer[cb.head]
	cb.head = (cb.head + 1) % len(cb.buffer)
	cb.count--
	return value, true
}

func (cb *CircularBuffer) IsEmpty() bool {
	return cb.count == 0
}

func (cb *CircularBuffer) IsFull() bool {
	return cb.count >= len(cb.buffer)
}

func main() {
	buffer := NewCircularBuffer()
	
	// 测试缓冲区
	fmt.Println("添加元素:")
	for i := 1; i <= 7; i++ {
		if ok := buffer.Push(i); ok {
			fmt.Printf("添加: %d\n", i)
		} else {
			fmt.Printf("缓冲区满，无法添加: %d\n", i)
		}
	}
	
	fmt.Println("\n取出元素:")
	for !buffer.IsEmpty() {
		if value, ok := buffer.Pop(); ok {
			fmt.Printf("取出: %d\n", value)
		}
	}
}
```

### 数组排序

```go
package main

import "fmt"

// 冒泡排序（原地排序）
func bubbleSort(arr *[5]int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// 测试排序
	numbers := [5]int{64, 34, 25, 12, 22}
	fmt.Printf("排序前: %v\n", numbers)
	
	bubbleSort(&numbers)
	fmt.Printf("排序后: %v\n", numbers)
	
	// 验证排序结果
	isSorted := true
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			isSorted = false
			break
		}
	}
	
	fmt.Printf("排序正确: %t\n", isSorted)
}
```

## 🧪 实践练习

### 练习 1：矩阵运算

```go
// 要求：
// 1. 实现矩阵加法
// 2. 实现矩阵转置
// 3. 实现矩阵乘法
// 4. 添加矩阵打印功能
```

参考实现：
```go
package main

import "fmt"

// 矩阵加法
func matrixAdd(a, b [2][3]int) ([2][3]int, error) {
	// 检查维度
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return [2][3]int{}, fmt.Errorf("矩阵维度不匹配")
	}
	
	var result [2][3]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	
	return result, nil
}

// 矩阵打印
func printMatrix(name string, matrix [2][3]int) {
	fmt.Printf("%s:\n", name)
	for i, row := range matrix {
		for _, value := range row {
			fmt.Printf("%4d", value)
		}
		fmt.Println()
	}
}

func main() {
	matrix1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	matrix2 := [2][3]int{{7, 8, 9}, {10, 11, 12}}
	
	printMatrix("矩阵1", matrix1)
	printMatrix("矩阵2", matrix2)
	
	result, err := matrixAdd(matrix1, matrix2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	
	printMatrix("结果", result)
}
```

### 练习 2：统计计算

```go
// 要求：
// 1. 实现数字统计（最大值、最小值、平均值）
// 2. 实现标准差计算
// 3. 实现中位数计算
// 4. 使用数组存储计算结果
```

### 练习 3：数组合并和查找

```go
// 要求：
// 1. 实现两个有序数组的合并
// 2. 实现二分查找
// 3. 实现数组去重
// 4. 分析算法复杂度
```

## 🤔 思考题

1. **数组和切片的主要区别是什么？**
   - 提示：考虑长度、内存分配、灵活性

2. **什么时候应该使用数组而不是切片？**
   - 提示：考虑性能需求、固定大小场景

3. **多维数组在内存中是如何存储的？**
   - 提示：考虑行主序和列主序

4. **为什么数组是值类型？**
   - 提示：考虑 Go 的设计哲学和内存安全

## 📚 扩展阅读

- [Go 语言规范 - 数组类型](https://golang.org/ref/spec#Array_types)
- [Effective Go - 数组](https://golang.org/doc/effective_go.html#arrays)
- [Go by Example - 数组](https://gobyexample.com/arrays)

## ⏭️ 下一章节

[切片](./07-slices.md) → 学习 Go 语言的切片类型

---

**💡 小贴士**: 数组在 Go 中是值类型，这意味着赋值和传参都是复制整个数组。记住：**当需要修改原数据或处理大量数据时，优先考虑使用切片**！
