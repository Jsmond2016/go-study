package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int = 10

	s1 := "hello, world"

	fmt.Println("a", a)
	fmt.Println("s1", s1)

	fmt.Println("-------------string--strings-----------------")

	fmt.Println("s1.length", len(s1))
	fmt.Println("s1.Split", strings.Split(s1, ""))
	fmt.Println("s1.Contains", strings.Contains(s1, "llo"))
	fmt.Println("s1.HasPrefix", strings.HasPrefix(s1, "he"))
	fmt.Println("s1.HasSuffix", strings.HasSuffix(s1, "ld"))
	fmt.Println("s1.Index", strings.Index(s1, "wo"))

	fmt.Println("---------------number-int-----------------")

	var num int = 1234
	var num2 int = 4567
	// 没有赋初值，int 默认初值为0
	var num3 int
	var total = num + num2
	fmt.Println("total", total)
	fmt.Println("num3", num3)

	fmt.Println("---------------bool-false-----------------")

	var trueFlag bool = true
	var falseFlag bool = false
	// 没有赋初值，默认为 false
	var nonFlag bool

	fmt.Println("trueFlag", trueFlag)
	fmt.Println("falseFlag", falseFlag)
	fmt.Println("nonFlag", nonFlag)

}
