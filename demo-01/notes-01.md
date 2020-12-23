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

egg:

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

egg:

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