// Package main 展示 Go 语言的切片
package main

import (
	"fmt"
	"slices"
)

// 程序入口，演示切片操作
func main() {
	fmt.Println("=== Go 切片示例 ===")

	// 基本切片操作
	basicSlices()

	// 切片底层原理
	sliceInternals()

	// 切片操作
	sliceOperations()

	// 切片函数
	sliceFunctions()

	// 多维切片
	multiDimensionalSlices()

	// 切片实际应用
	practicalExample()
}

// basicSlices 基本切片操作
func basicSlices() {
	fmt.Println("\n--- 基本切片操作 ---")

	// 1. 切片声明和初始化
	var slice1 []int                    // nil 切片
	var slice2 = []int{1, 2, 3, 4, 5}   // 直接初始化
	slice3 := make([]int, 5)            // 使用 make 创建，长度为5，容量为5
	slice4 := make([]int, 3, 10)        // 使用 make 创建，长度为3，容量为10

	fmt.Printf("nil 切片: %v (len=%d, cap=%d, isNil=%t)\n", slice1, len(slice1), cap(slice1), slice1 == nil)
	fmt.Printf("直接初始化: %v (len=%d, cap=%d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("make 创建: %v (len=%d, cap=%d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("make 指定容量: %v (len=%d, cap=%d)\n", slice4, len(slice4), cap(slice4))

	// 2. 从数组创建切片
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice5 := arr[2:7] // 创建切片，包含索引2到6的元素
	fmt.Printf("\n从数组创建切片: arr[2:7] = %v\n", slice5)

	// 3. 切片的切片
	slice6 := slice5[1:4] // 从现有切片创建新切片
	fmt.Printf("切片的切片: slice5[1:4] = %v\n", slice6)
}

// sliceInternals 切片底层原理
func sliceInternals() {
	fmt.Println("\n--- 切片底层原理 ---")

	// 创建一个切片
	s := make([]int, 3, 5)
	s[0], s[1], s[2] = 10, 20, 30

	fmt.Printf("原始切片: %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	// 演示共享底层数组
	s2 := s[1:3] // 创建新切片，共享底层数组
	fmt.Printf("新切片: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	// 修改新切片会影响原切片
	s2[0] = 99
	fmt.Printf("修改新切片后:\n")
	fmt.Printf("原切片: %v\n", s)
	fmt.Printf("新切片: %v\n", s2)

	// 使用 append 可能会触发扩容
	fmt.Printf("\n使用 append:")
	s2 = append(s2, 40)
	fmt.Printf("扩容后新切片: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))
	fmt.Printf("原切片不受影响: %v\n", s)
}

// sliceOperations 切片操作
func sliceOperations() {
	fmt.Println("\n--- 切片操作 ---")

	// 1. 添加元素
	slice := []int{1, 2, 3}
	fmt.Printf("原始切片: %v\n", slice)

	// 添加单个元素
	slice = append(slice, 4)
	fmt.Printf("添加单个元素: %v\n", slice)

	// 添加多个元素
	slice = append(slice, 5, 6, 7)
	fmt.Printf("添加多个元素: %v\n", slice)

	// 添加另一个切片
	another := []int{8, 9, 10}
	slice = append(slice, another...)
	fmt.Printf("添加另一个切片: %v\n", slice)

	// 2. 删除元素
	fmt.Printf("\n删除元素:")

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

	// 3. 复制切片
	fmt.Printf("\n复制切片:")
	original := []int{1, 2, 3, 4, 5}
	copy1 := make([]int, len(original))
	copy(copy1, original)
	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("复制品: %v\n", copy1)

	// 浅拷贝（共享底层数组）
	copy2 := original
	copy2[0] = 99
	fmt.Printf("浅拷贝后 - 原: %v, 复制: %v\n", original, copy2)
}

// sliceFunctions 切片函数
func sliceFunctions() {
	fmt.Println("\n--- 切片函数 ---")

	// 1. 检查元素是否存在
	numbers := []int{1, 3, 5, 7, 9}
	fmt.Printf("切片: %v\n", numbers)
	fmt.Printf("3 是否存在: %t\n", contains(numbers, 3))
	fmt.Printf("4 是否存在: %t\n", contains(numbers, 4))

	// 2. 过滤切片
	filtered := filter(numbers, func(n int) bool {
		return n > 5
	})
	fmt.Printf("大于5的元素: %v\n", filtered)

	// 3. 映射切片
	mapped := mapSlice(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("每个元素乘以2: %v\n", mapped)

	// 4. 反转切片
	fmt.Printf("反转前: %v\n", numbers)
	reverse(numbers)
	fmt.Printf("反转后: %v\n", numbers)

	// 5. 查找最大值和最小值
	max, min := findMaxMin(numbers)
	fmt.Printf("最大值: %d, 最小值: %d\n", max, min)
}

// contains 检查元素是否存在
func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// filter 过滤切片
func filter(slice []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// mapSlice 映射切片
func mapSlice(slice []int, transform func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// reverse 反转切片
func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// findMaxMin 查找最大值和最小值
func findMaxMin(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}

	max, min := slice[0], slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

// multiDimensionalSlices 多维切片
func multiDimensionalSlices() {
	fmt.Println("\n--- 多维切片 ---")

	// 1. 二维切片
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
		for j := range matrix[i] {
			matrix[i][j] = i*4 + j + 1
		}
	}

	fmt.Printf("二维切片:\n")
	for i, row := range matrix {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 2. 不规则的二维切片
	jagged := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}

	fmt.Printf("\n不规则二维切片:\n")
	for i, row := range jagged {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 3. 添加行和列
	fmt.Printf("\n动态添加行和列:")
	matrix = append(matrix, []int{13, 14, 15, 16}) // 添加新行
	matrix[0] = append(matrix[0], 17)               // 添加列
	matrix[1] = append(matrix[1], 18, 19)           // 添加多列

	fmt.Printf("添加后的矩阵:\n")
	for i, row := range matrix {
		fmt.Printf("  行 %d: %v\n", i, row)
	}
}

// practicalExample 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1：购物车管理
	type Product struct {
		ID    int
		Name  string
		Price float64
	}

	type CartItem struct {
		Product  Product
		Quantity int
	}

	// 创建商品列表
	products := []Product{
		{1, "苹果", 5.5},
		{2, "香蕉", 3.2},
		{3, "橙子", 4.8},
		{4, "葡萄", 8.9},
	}

	fmt.Println("商品列表:")
	for i, product := range products {
		fmt.Printf("%d. %s - ¥%.1f\n", i+1, product.Name, product.Price)
	}

	// 购物车
	var cart []CartItem

	// 添加商品到购物车
	cart = append(cart, CartItem{products[0], 2})  // 2个苹果
	cart = append(cart, CartItem{products[2], 1})  // 1个橙子
	cart = append(cart, CartItem{products[3], 3})  // 3个葡萄

	fmt.Printf("\n购物车商品 (%d 种):\n", len(cart))
	var total float64
	for i, item := range cart {
		subtotal := float64(item.Quantity) * item.Product.Price
		total += subtotal
		fmt.Printf("%d. %s x%d = ¥%.1f\n",
			i+1, item.Product.Name, item.Quantity, subtotal)
	}
	fmt.Printf("总计: ¥%.1f\n", total)

	// 使用新的 slices 包操作（Go 1.21+）
	fmt.Printf("\n使用标准库 slices 包:")

	// 排序购物车（按商品ID）
	slices.SortFunc(cart, func(a, b CartItem) int {
		return a.Product.ID - b.Product.ID
	})
	fmt.Printf("排序后的购物车:\n")
	for i, item := range cart {
		fmt.Printf("%d. %s\n", i+1, item.Product.Name)
	}

	// 查找特定商品
	appleIndex := slices.IndexFunc(cart, func(item CartItem) bool {
		return item.Product.Name == "苹果"
	})
	if appleIndex != -1 {
		fmt.Printf("找到苹果在购物车中的位置: %d\n", appleIndex)
	}

	// 示例2：数据分析
	scores := []int{85, 92, 78, 96, 88, 73, 89, 95, 82, 90}
	fmt.Printf("\n学生成绩分析:\n")
	fmt.Printf("成绩列表: %v\n", scores)
	fmt.Printf("最高分: %d\n", slices.Max(scores))
	fmt.Printf("最低分: %d\n", slices.Min(scores))

	// 统计各等级
	excellent := filter(scores, func(score int) bool { return score >= 90 })
	good := filter(scores, func(score int) bool { return score >= 80 && score < 90 })
	pass := filter(scores, func(score int) bool { return score >= 60 && score < 80 })
	fmt.Printf("优秀 (>=90): %d 人: %v\n", len(excellent), excellent)
	fmt.Printf("良好 (80-89): %d 人: %v\n", len(good), good)
	fmt.Printf("及格 (60-79): %d 人: %v\n", len(pass), pass)

	// 计算平均分
	sum := 0
	for _, score := range scores {
		sum += score
	}
	average := float64(sum) / float64(len(scores))
	fmt.Printf("平均分: %.1f\n", average)
}