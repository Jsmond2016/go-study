package main

import "fmt"

// 切片学习：https://www.liwenzhou.com/posts/Go/06_slice/

func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)         //切片是引用类型，不支持直接比较，只能和nil比较

	fmt.Println("============>>>>>>>>>>>>>>>>>>>>>>>==============")

	// 使用 make
	var slice1 []int
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(len(slice1))

	slice2 := make([]string, 2)
	slice2 = append(slice2, "Hello", "world")
	fmt.Println(slice2)

	// 遍历
	var list []int
	list = append(list, 1, 2, 3, 4, 5)
	for i, v := range list {
		fmt.Println("key: ", i, "==> value :", v)
	}

	strList := []string{"Hello", "World"}
	for i, v := range strList {
		fmt.Println("key: ", i, "==> value :", v)
	}

	for i := 0; i < len(strList); i++ {
		fmt.Println("key: ", i, "==> value :", strList[i])
	}

	fmt.Println("============>>>>>>>>>> 使用 copy >>>>>>>>>>>>>==============")

	// 使用 copy
	strList2 := make([]string, len(strList))
	copy(strList2, strList)
	fmt.Println(strList2)

	fmt.Println("============>>>>>>>>>>> 删除元素 append >>>>>>>>>>>>==============")

	// 删除元素 -- 删除索引为2 的元素
	// 总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	sourceList := []int{30, 31, 32, 33, 34, 35, 36, 37}
	fmt.Println(sourceList)
	sourceList = append(sourceList[:1], sourceList[2:]...)
	fmt.Println(sourceList)
}
