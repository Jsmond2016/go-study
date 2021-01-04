# go 语言学习

> [VS Code配置Go语言开发环境](https://www.liwenzhou.com/posts/Go/00_go_in_vscode/)


## go 安装

- 安装包地址：[go](https://golang.google.cn/dl/) 下载安装即可
- 配置 vs code 环境，参考上面的教程
- 安装 go 语言格式化以及相关工具
- 配置简单的 snippets


## hello，world

```go

package main
import "fmt"

func main() {
  fmt.Println("hello, world")
}

```

## string 类型


```go
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

	fmt.Println("--------------------------------")

	fmt.Println("s1.length", len(s1))
	fmt.Println("s1.Split", strings.Split(s1, ""))
	fmt.Println("s1.Contains", strings.Contains(s1, "llo"))
	fmt.Println("s1.HasPrefix", strings.HasPrefix(s1, "he"))
	fmt.Println("s1.HasSuffix", strings.HasSuffix(s1, "ld"))
	fmt.Println("s1.Index", strings.Index(s1, "wo"))
}

```

## int 类型

> 还包括 float, int32, int64 等

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------number-int-----------------")

	var num int = 1234
	var num2 int = 4567
	// 没有赋初值，int 默认初值为0
	var num3 int
	var total = num + num2
	fmt.Println("total", total)
	fmt.Println("num3", num3)

}

```

## bool 类型


```go
package main

import (
	"fmt"
)

func main() {

	fmt.Println("---------------bool-false-----------------")

	var trueFlag bool = true
	var falseFlag bool = false
	// 没有赋初值，默认为 false
	var nonFlag bool

	fmt.Println("trueFlag", trueFlag)
	fmt.Println("falseFlag", falseFlag)
	fmt.Println("nonFlag", nonFlag)

}

```

## 条件控制语句

- if ... else
- switch
- select


参考资料：

- [【golang】select关键字用法](https://www.jianshu.com/p/2a1146dc42c3)


## 循环语句

- for
- for 循环嵌套


```go
for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                continue
            }
    }
```
- for ... continue

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* for 循环 */
   for a < 20 {
      if a == 15 {
         /* 跳过此次循环 */
         a = a + 1;
         continue;
      }
      fmt.Printf("a 的值为 : %d\n", a);
      a++;    
   }  
}
```

- [资料-菜鸟教程](https://www.runoob.com/go/go-continue-statement.html)


## 函数

形式

```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

eg:

```go
package main

import "fmt"

func main() {
	res := getMaxNum(1, 2)
	fmt.Println("res", res)
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

```

## 方法

> Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集

形式：

```go
func (variable_name variable_data_type) function_name() [return_type]{
   /* 函数体*/
}
```

eg:

```go
package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

```

## 数组

- 数组的定义：

```go
var 数组变量名 [元素数量]T


var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型
```

- 数组初始化-方式1：使用初始化列表来设置数组元素的值

```go
package main

import "fmt"

func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化

	fmt.Println(testArray) //[0 0 0]
	fmt.Println(numArray)  //[1 2 0]
	fmt.Println(cityArray) //[北京 上海 深圳]

	fmt.Println("---------------------")
	var arr = [2]int{10, 20}
	fmt.Println(arr)
	var arrStr = [3]string{"hello", "world", "!"}
	fmt.Println(arrStr)
}

```

- 数组初始化-方式2：让编译器根据初始值的个数自行推断数组的长度

```go
func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

- 数组初始化-方式3：使用指定索引值的方式来初始化数组

```go
func main() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
```

- 数组的遍历

```go
func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}

	fmt.Println("===================")
	var test = [...]string{"hello", "world=======", "略略略"}
	for _, val := range test {
		fmt.Println(val)
	}
}
```

- 多维数组，定义

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
```

- 二维数组的遍历

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
```

**注意：** 多维数组**只有第一层**可以使用`...`来让编译器推导数组长度。例如：

```go
//支持的写法
a := [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
```

- 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
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
}
```



## 切片

切片： slice

数组的缺点：固定长度，不灵活，因此有很大的局限性

切片的特点：

- 是一个拥有相同类型元素的**可变长度**的序列
- 是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。

切片的定义：

```go
var name []T
```

eg:

```go
func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)         //切片是引用类型，不支持直接比较，只能和nil比较
}
```

**注意：**切片是引用类型，不支持直接比较，只能和nil比较

- 切片的长度和容量：
  - 长度：`len(s)`
  - 容量：`cap(s)`