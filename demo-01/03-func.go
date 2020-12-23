package main

import "fmt"

func main() {
	res := getMaxNum(1, 2)
	fmt.Println("res", res)

	fmt.Println("----------------------")
	a, b := swap("Google", "Runoob")
	fmt.Println("a, b", a, b)
}

func getMaxNum(num1 int, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
