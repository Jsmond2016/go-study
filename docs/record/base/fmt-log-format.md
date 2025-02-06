
# fmt常用工具和配置、常用的格式化占位符

学习资料：

- [fmt](https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/fmt.html)


## Print 和 Println 和 Printf

- Print系列函数会将内容输出到系统的标准输出
- Println 函数会在输出内容的结尾添加一个换行符
- Printf 函数支持格式化输出字符串

```go
func main() {
    fmt.Print("在终端打印该信息。")
    name := "枯藤"
    fmt.Printf("我是：%s\n", name)
    fmt.Println("在终端打印单独一行显示")
}
```

## Sprint、Sprintln、Sprintf

Sprint系列函数会把传入的数据生成并返回一个字符串。

最长用：`fmt.Sprintf` 格式化拼接

```go
func main() {
  // 使用 Sprint 拼接
  hello := "hello"
	world := "world"
	s1 := fmt.Sprint("测试 Sprint", " ", hello, " ", world)

  // 使用 + 拼接
	ss1 := "测试 Sprint" + " " + hello + " " + world

  // 格式化输出 
	name1 := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name1, age)

	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, ss1, s2, s3)
}

```


## 格式占位符


### 通用占位符

最常用：`%v, %+v` 表示结构体，`%T` 表示类型

|占位符	|说明|
|  :--:  | :--:  |
|%v	    |值的默认格式表示|
|%+v	  |类似%v，但输出结构体时会添加字段名|
|%#v	   |值的Go语法表示|
|%T	    |打印值的类型|
|%%	    |百分号|

代码示例：

```go
func main() {
  // 值
  fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	
  // 结构体
  o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	
  fmt.Println()
	
  fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}
```

### 布尔型

|占位符	|说明|
|  :--:  | :--:  |
|%t	| true或false|


### 整型

常用：`%d` 十进制输出数字

|占位符	|说明|
|  :--:  | :--:  |
|%b	|表示为二进制|
|%c	|该值对应的unicode码值|
|%d	|表示为十进制|
|%o	|表示为八进制|
|%x	|表示为十六进制，使用a-f|
|%X	|表示为十六进制，使用A-F|
|%U	|表示为Unicode格式：U+1234，等价于”U+%04X”|
|%q	|该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示|


代码：

```go
func main() {
  n := 65
  fmt.Printf("%b\n", n)
  fmt.Printf("%c\n", n)
  // 十进制输出
  fmt.Printf("%d\n", n)
  
  fmt.Printf("%o\n", n)
  fmt.Printf("%x\n", n)
  fmt.Printf("%X\n", n)
}

```




### 浮点型

主要记住常用的2个：`%f, %g` 


|占位符	|说明|
|  :--:  | :--:  |
|%b |	无小数部分、二进制指数的科学计数法，如-123456p-78|
|%e |	科学计数法，如-1234.456e+78|
|%E |	科学计数法，如-1234.456E+78|
|%f |	有小数部分但无指数部分，如123.456|
|%F |	等价于%f|
|%g |	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）|
|%G |	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）|


代码：


```go
func main() {
  f := 1200.34567
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	//
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}

// 输出结果
// 5279176086062293p-42
// 1.200346e+03
// 1.200346E+03
// 1200.345670
// 1200.345670
// 1200.34567
// 1200.34567

```

## Scan，Scanf, Scanln

- Scan 从标准输入扫描文本，读取由 **空白符** 分隔的值保存到传递给本函数的参数中，**换行符** 视为空白符。

```go
func main() {
  var (
		name11  string
		age1    int
		married bool
	)
	fmt.Scan(&name11, &age1, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name11, age1, married)
}
```

- Scanf 从标准输入扫描文本，根据 **format 参数指定的格式** 去读取由空白符分隔的值保存到传递给本函数的参数中。

```go
func main() {
    var (
        name    string
        age     int
        married bool
    )
    // 按照指定的格式输入才能取到值 1:Alan 2:18 3:false
    fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

```


- Scanln 类似 Scan，它以**空格**进行分隔，在 **遇到换行** 时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。


```go
func main() {
  var (
		name111    string
		age111     int
		married111 bool
	)
	fmt.Scanln(&name111, &age111, &married111)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name111, age111, married111)
}
```