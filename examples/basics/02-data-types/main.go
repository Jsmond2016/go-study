// Package main 展示 Go 语言的各种数据类型
package main

import (
	"fmt"
	"unsafe"
)

// 程序入口，演示各种数据类型
func main() {
	fmt.Println("=== Go 数据类型示例 ===")

	// 基本数据类型
	basicTypes()

	// 复合数据类型
	compoundTypes()

	// 自定义类型
	customTypes()

	// 类型信息
	typeInfo()
}

// basicTypes 展示基本数据类型
func basicTypes() {
	fmt.Println("\n--- 基本数据类型 ---")

	// 1. 整型
	fmt.Println("整型类型:")
	var (
		int8Var   int8   = 127    // -128 到 127
		int16Var  int16  = 32767  // -32768 到 32767
		int32Var  int32  = 2147483647
		int64Var  int64  = 9223372036854775807
		uint8Var  uint8  = 255    // 0 到 255
		uint16Var uint16 = 65535  // 0 到 65535
		uint32Var uint32 = 4294967295
		uint64Var uint64 = 18446744073709551615
		intVar    int    = 42     // 32位系统上是int32，64位系统上是int64
		uintVar   uint   = 42     // 32位系统上是uint32，64位系统上是uint64
	)

	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n",
		int8Var, int16Var, int32Var, int64Var)
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d, uint64: %d\n",
		uint8Var, uint16Var, uint32Var, uint64Var)
	fmt.Printf("int: %d, uint: %d\n", intVar, uintVar)

	// 特殊整型
	var ptrVar uintptr = 0x12345678
	fmt.Printf("uintptr: 0x%x\n", ptrVar)

	// 2. 浮点型
	fmt.Println("\n浮点型类型:")
	var (
		float32Var float32 = 3.14159
		float64Var float64 = 2.718281828459045
	)

	fmt.Printf("float32: %.2f, float64: %.15f\n", float32Var, float64Var)

	// 3. 复数类型
	fmt.Println("\n复数类型:")
	var (
		complex64Var  complex64  = complex(3, 4)      // 3 + 4i
		complex128Var complex128 = complex(1.5, 2.5)  // 1.5 + 2.5i
	)

	fmt.Printf("complex64: %.1f + %.1fi\n", real(complex64Var), imag(complex64Var))
	fmt.Printf("complex128: %.1f + %.1fi\n", real(complex128Var), imag(complex128Var))

	// 4. 布尔型
	fmt.Println("\n布尔型:")
	var (
		boolTrue  bool = true
		boolFalse bool = false
	)

	fmt.Printf("true: %t, false: %t\n", boolTrue, boolFalse)

	// 5. 字符串
	fmt.Println("\n字符串:")
	var (
		stringVar   string = "Hello, Go!"
		runeVar     rune   = '中'   // Unicode 字符
		byteVar     byte   = 65     // ASCII 字符 'A'
	)

	fmt.Printf("字符串: %s\n", stringVar)
	fmt.Printf("字符: %c (Unicode: %U)\n", runeVar, runeVar)
	fmt.Printf("字节: %c (ASCII: %d)\n", byteVar, byteVar)

	// 字符串操作
	fmt.Printf("字符串长度: %d\n", len(stringVar))
	fmt.Printf("字符串字节数: %d\n", len([]byte(stringVar)))

	// 6. 指针类型
	fmt.Println("\n指针类型:")
	var (
		intValue int = 100
		intPtr   *int = &intValue
	)

	fmt.Printf("整数值: %d\n", intValue)
	fmt.Printf("指针地址: %p\n", intPtr)
	fmt.Printf("通过指针访问值: %d\n", *intPtr)
	fmt.Printf("指针的大小: %d 字节\n", unsafe.Sizeof(intPtr))
}

// compoundTypes 展示复合数据类型
func compoundTypes() {
	fmt.Println("\n--- 复合数据类型 ---")

	// 1. 数组
	fmt.Println("数组:")
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := [...]string{"Go", "Python", "Java"} // 编译器推断长度

	fmt.Printf("整数数组: %v\n", arr1)
	fmt.Printf("字符串数组: %v\n", arr2)
	fmt.Printf("数组长度: %d\n", len(arr1))

	// 2. 切片
	fmt.Println("\n切片:")
	var (
		slice1 []int = []int{1, 2, 3, 4, 5}
		slice2 = make([]string, 3) // 使用 make 创建
	)

	slice2[0] = "苹果"
	slice2[1] = "香蕉"
	slice2[2] = "橙子"

	fmt.Printf("整数切片: %v\n", slice1)
	fmt.Printf("字符串切片: %v\n", slice2)

	// 切片操作
	fmt.Printf("切片长度: %d, 容量: %d\n", len(slice1), cap(slice1))
	fmt.Printf("切片子区间: %v\n", slice1[1:4])

	// 动态添加元素
	slice1 = append(slice1, 6, 7, 8)
	fmt.Printf("添加元素后: %v\n", slice1)

	// 3. 映射
	fmt.Println("\n映射:")
	var (
		map1 map[string]int = make(map[string]int)
		map2 = map[string]string{
			"name":    "张三",
			"city":    "北京",
			"country": "中国",
		}
	)

	map1["apple"] = 1
	map1["banana"] = 2
	map1["orange"] = 3

	fmt.Printf("数字映射: %v\n", map1)
	fmt.Printf("字符串映射: %v\n", map2)

	// 映射操作
	fmt.Printf("apple 的值: %d\n", map1["apple"])
	fmt.Printf("映射长度: %d\n", len(map1))

	// 检查键是否存在
	if value, exists := map1["apple"]; exists {
		fmt.Printf("键 'apple' 存在，值为: %d\n", value)
	}

	// 4. 结构体
	fmt.Println("\n结构体:")
	type Person struct {
		Name string
		Age  int
		City string
	}

	person1 := Person{
		Name: "李四",
		Age:  30,
		City: "上海",
	}

	person2 := Person{王五, 25, "广州"}

	fmt.Printf("人员1: %+v\n", person1)
	fmt.Printf("人员2: %+v\n", person2)

	// 结构体指针
	personPtr := &person1
	fmt.Printf("通过指针访问: %s\n", personPtr.Name)
}

// customTypes 展示自定义类型
func customTypes() {
	fmt.Println("\n--- 自定义类型 ---")

	// 1. 类型别名
	type (
		ID       int
		Username string
		Age      uint8
	)

	var userID ID = 1001
	var username Username = "admin"
	var userAge Age = 25

	fmt.Printf("用户信息: ID=%d, 用户名=%s, 年龄=%d\n", userID, username, userAge)

	// 2. 自定义结构体
	type Address struct {
		Province string
		City     string
		Street   string
		Number   int
	}

	type User struct {
		ID     ID
		Name   Username
		Age    Age
		Addr   Address
		Email  string
		Phone  string
	}

	user := User{
		ID:   2001,
		Name: "赵六",
		Age:  35,
		Addr: Address{
			Province: "浙江省",
			City:     "杭州市",
			Street:   "西湖大道",
			Number:   123,
		},
		Email: "zhaoliu@example.com",
		Phone: "13800138000",
	}

	fmt.Printf("完整用户信息: %+v\n", user)

	// 3. 自定义方法
	func (u User) GetDisplayName() string {
		return fmt.Sprintf("%s (%d岁)", u.Name, u.Age)
	}

	fmt.Printf("用户显示名称: %s\n", user.GetDisplayName())

	// 4. 接口实现
	type Stringer interface {
		String() string
	}

	func (u User) String() string {
		return fmt.Sprintf("User{ID:%d, Name:%s, City:%s}",
			u.ID, u.Name, u.Addr.City)
	}

	var stringer Stringer = user
	fmt.Printf("接口输出: %s\n", stringer.String())
}

// typeInfo 展示类型信息
func typeInfo() {
	fmt.Println("\n--- 类型信息 ---")

	// 使用 reflect 包获取类型信息
	fmt.Println("类型大小信息:")
	var (
		intVal    int     = 42
		floatVal  float64 = 3.14
		stringVal string  = "Hello"
		boolVal   bool    = true
	)

	fmt.Printf("int 大小: %d 字节\n", unsafe.Sizeof(intVal))
	fmt.Printf("float64 大小: %d 字节\n", unsafe.Sizeof(floatVal))
	fmt.Printf("string 大小: %d 字节\n", unsafe.Sizeof(stringVal))
	fmt.Printf("bool 大小: %d 字节\n", unsafe.Sizeof(boolVal))

	// 零值
	fmt.Println("\n零值:")
	var (
		zeroInt    int
		zeroFloat  float64
		zeroString string
		zeroBool   bool
		zeroSlice  []int
		zeroMap    map[string]int
		zeroPtr    *int
	)

	fmt.Printf("int 零值: %d\n", zeroInt)
	fmt.Printf("float64 零值: %f\n", zeroFloat)
	fmt.Printf("string 零值: '%s'\n", zeroString)
	fmt.Printf("bool 零值: %t\n", zeroBool)
	fmt.Printf("slice 零值: %v\n", zeroSlice)
	fmt.Printf("map 零值: %v\n", zeroMap)
	fmt.Printf("pointer 零值: %v\n", zeroPtr)
}