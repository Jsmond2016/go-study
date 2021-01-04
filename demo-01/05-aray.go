package main

import "fmt"

// func main() {
// 	var testArray [3]int                        //数组会初始化为int类型的零值
// 	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
// 	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化

// 	fmt.Println(testArray) //[0 0 0]
// 	fmt.Println(numArray)  //[1 2 0]
// 	fmt.Println(cityArray) //[北京 上海 深圳]

// 	fmt.Println("---------------------")
// 	var arr = [2]int{10, 20}
// 	fmt.Println(arr)
// 	var arrStr = [3]string{"hello", "world", "!"}
// 	fmt.Println(arrStr)
// }

// func main() {
// 	var testArray [3]int
// 	var numArray = [...]int{1, 2}
// 	var cityArray = [...]string{"北京", "上海", "深圳"}
// 	fmt.Println(testArray)                          //[0 0 0]
// 	fmt.Println(numArray)                           //[1 2]
// 	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
// 	fmt.Println(cityArray)                          //[北京 上海 深圳]
// 	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
// }

// func main() {
// 	var a = [...]string{"北京", "上海", "深圳"}
// 	// 方法1：for循环遍历
// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(a[i])
// 	}

// 	// 方法2：for range遍历
// 	for index, value := range a {
// 		fmt.Println(index, value)
// 	}

// 	fmt.Println("===================")
// 	var test = [...]string{"hello", "world=======", "略略略"}
// 	for _, val := range test {
// 		fmt.Println(val)
// 	}
// }

// func main() {
// 	a := [3][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}
// 	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
// 	fmt.Println(a[2][1]) //支持索引取值:重庆
// }

// func main() {
// 	a := [3][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}
// 	for _, v1 := range a {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s\t", v2)
// 		}
// 		fmt.Println()
// 	}
// }

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func getTotal(x [6]int) int {
	total := 0
	for _, val := range x {
		total += val
	}
	return total
}

func main() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
	fmt.Println("==============")

	myTest := [6]int{1, 2, 3, 4, 5, 6}
	_total := getTotal(myTest)
	fmt.Println(_total)
}
