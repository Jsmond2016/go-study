package main

import "fmt"

func formatTitle(titles string) {
	fmt.Println()
	fmt.Printf("============= %s =================", titles)
	fmt.Println()
}

func main() {
	formatTitle("Print, Println, Printf")

	fmt.Print("在终端打印该信息。")
	name := "枯藤"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
	// ==============================================

	formatTitle("Sprint, Sprintln, Sprintf")

	hello := "hello"
	world := "world"
	s1 := fmt.Sprint("测试 Sprint", " ", hello, " ", world)

	ss1 := "测试 Sprint" + " " + hello + " " + world

	name1 := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name1, age)

	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, ss1, s2, s3)

	// ==============================================
	formatTitle("Errorf")
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)

	// ==============================================
	formatTitle("format 格式占位符")
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	fmt.Printf("%T\n", false)

	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)

	fmt.Println()

	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	// ==============================================
	formatTitle("整数数字格式化 65")
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	// 数字十进制
	fmt.Printf("%d\n", n)

	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	// ==============================================
	formatTitle("浮点数数字格式化 1200.34567")
	f := 1200.34567
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	//
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	// ==============================================
	// formatTitle("fmt.Scan 的使用")

	// var (
	// 	name11  string
	// 	age1    int
	// 	married bool
	// )
	// fmt.Scan(&name11, &age1, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name11, age1, married)

	// ==============================================
	// formatTitle("fmt.Scanf 的使用")

	// var (
	// 	n1        string
	// 	a1        int
	// 	isMarried bool
	// )
	// // 按照指定的格式输入才能取到值 1:Alan 2:18 3:false
	// fmt.Scanf("1:%s 2:%d 3:%t", &n1, &a1, &isMarried)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", n1, a1, isMarried)

	formatTitle("fmt.Scanln 的使用")
	var (
		name111    string
		age111     int
		married111 bool
	)
	fmt.Scanln(&name111, &age111, &married111)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name111, age111, married111)
}
