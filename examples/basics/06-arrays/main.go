// Package main 展示 Go 语言的数组
package main

import (
	"fmt"
)

// 程序入口，演示数组操作
func main() {
	fmt.Println("=== Go 数组示例 ===")

	// 基本数组操作
	basicArrays()

	// 数组遍历
	arrayTraversal()

	// 数组比较
	arrayComparison()

	// 数组作为函数参数
	arrayAsParameter()

	// 多维数组
	multiDimensionalArrays()

	// 数组实际应用
	practicalExample()
}

// basicArrays 基本数组操作
func basicArrays() {
	fmt.Println("\n--- 基本数组操作 ---")

	// 1. 数组声明和初始化
	var arr1 [5]int                    // 零值数组
	var arr2 = [5]int{1, 2, 3, 4, 5}  // 指定值
	arr3 := [5]int{10, 20, 30}         // 部分初始化，其余为零值
	arr4 := [...]int{1, 2, 3, 4, 5}    // 编译器推断长度
	arr5 := [5]int{0: 100, 4: 500}     // 指定索引初始化

	fmt.Printf("零值数组: %v\n", arr1)
	fmt.Printf("指定值数组: %v\n", arr2)
	fmt.Printf("部分初始化: %v\n", arr3)
	fmt.Printf("推断长度: %v (长度: %d)\n", arr4, len(arr4))
	fmt.Printf("索引初始化: %v\n", arr5)

	// 2. 访问和修改元素
	fmt.Println("\n访问和修改元素:")
	fmt.Printf("arr2[0] = %d\n", arr2[0])
	arr2[0] = 100
	fmt.Printf("修改后 arr2[0] = %d\n", arr2[0])
	fmt.Printf("修改后的数组: %v\n", arr2)

	// 3. 数组长度和容量
	fmt.Printf("数组长度: %d\n", len(arr2))
	// 数组容量等于长度
	fmt.Printf("数组容量: %d\n", cap(arr2))
}

// arrayTraversal 数组遍历
func arrayTraversal() {
	fmt.Println("\n--- 数组遍历 ---")

	numbers := [5]int{10, 20, 30, 40, 50}

	// 1. 传统 for 循环
	fmt.Println("传统 for 循环:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("索引 %d: %d\n", i, numbers[i])
	}

	// 2. for-range 循环
	fmt.Println("\nfor-range 循环:")
	for index, value := range numbers {
		fmt.Printf("索引 %d: %d\n", index, value)
	}

	// 3. 只要值
	fmt.Println("\n只要值:")
	for _, value := range numbers {
		fmt.Printf("值: %d\n", value)
	}

	// 4. 只要索引
	fmt.Println("\n只要索引:")
	for index := range numbers {
		fmt.Printf("索引: %d\n", index)
	}
}

// arrayComparison 数组比较
func arrayComparison() {
	fmt.Println("\n--- 数组比较 ---")

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	// 相同类型和长度的数组可以比较
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)
	fmt.Printf("arr2 == arr3: %t\n", arr2 == arr3)

	// 不同长度的数组不能比较
	// var arr4 [4]int = [4]int{1, 2, 3, 4}
	// fmt.Printf("arr1 == arr4: %t\n", arr1 == arr4) // 编译错误
}

// arrayAsParameter 数组作为函数参数
func arrayAsParameter() {
	fmt.Println("\n--- 数组作为函数参数 ---")

	// Go 中数组是值传递
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("原数组: %v\n", arr)

	modifyArray(arr)
	fmt.Printf("函数调用后: %v\n", arr) // 数组没有被修改

	// 如果要修改原数组，需要使用指针
	modifyArrayPtr(&arr)
	fmt.Printf("使用指针修改后: %v\n", arr) // 数组被修改
}

// modifyArray 值传递（不会修改原数组）
func modifyArray(arr [5]int) {
	arr[0] = 100
	fmt.Printf("函数内部修改: %v\n", arr)
}

// modifyArrayPtr 指针传递（会修改原数组）
func modifyArrayPtr(arrPtr *[5]int) {
	(*arrPtr)[0] = 100
	fmt.Printf("函数内部修改: %v\n", *arrPtr)
}

// multiDimensionalArrays 多维数组
func multiDimensionalArrays() {
	fmt.Println("\n--- 多维数组 ---")

	// 1. 二维数组声明
	var matrix [3][4]int
	fmt.Printf("二维数组零值: %v\n", matrix)

	// 2. 二维数组初始化
	matrix2 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Printf("初始化的二维数组:\n")
	for i, row := range matrix2 {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 3. 访问二维数组元素
	fmt.Printf("matrix2[1][2] = %d\n", matrix2[1][2])
	matrix2[1][2] = 100
	fmt.Printf("修改后 matrix2[1][2] = %d\n", matrix2[1][2])

	// 4. 三维数组示例
	cube := [2][2][3]int{
		{
			{1, 2, 3},
			{4, 5, 6},
		},
		{
			{7, 8, 9},
			{10, 11, 12},
		},
	}
	fmt.Printf("三维数组: %v\n", cube)
}

// practicalExample 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1：学生成绩管理
	type Student struct {
		Name   string
		Scores [5]int // 5门课程成绩
	}

	students := [3]Student{
		{"张三", [5]int{85, 90, 78, 92, 88}},
		{"李四", [5]int{76, 85, 92, 78, 95}},
		{"王五", [5]int{92, 88, 76, 85, 90}},
	}

	fmt.Println("学生成绩管理:")
	for i, student := range students {
		total := 0
		for _, score := range student.Scores {
			total += score
		}
		average := float64(total) / float64(len(student.Scores))
		fmt.Printf("学生 %d: %s, 总分: %d, 平均分: %.1f\n",
			i+1, student.Name, total, average)
	}

	// 示例2：游戏棋盘
	type Position struct {
		X, Y int
	}

	const BOARD_SIZE = 3
	var board [BOARD_SIZE][BOARD_SIZE]string

	// 初始化棋盘
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			board[i][j] = "-"
		}
	}

	// 模拟下棋
	moves := []Position{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {1, 0}}
	players := []string{"X", "O"}

	fmt.Println("\n棋盘游戏:")
	printBoard(board)

	for i, move := range moves {
		player := players[i%2]
		board[move.X][move.Y] = player
		fmt.Printf("\n玩家 %s 下在 (%d,%d):\n", player, move.X, move.Y)
		printBoard(board)
	}
}

// printBoard 打印棋盘
func printBoard(board [3][3]string) {
	for _, row := range board {
		fmt.Printf(" ")
		for j, cell := range row {
			fmt.Printf(cell)
			if j < len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}