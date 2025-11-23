// Package main 展示 Go 语言的指针
package main

import (
	"fmt"
	"unsafe"
)

// 程序入口，演示指针操作
func main() {
	fmt.Println("=== Go 指针示例 ===")

	// 基本指针操作
	basicPointers()

	// 指针与函数
	pointersAndFunctions()

	// 指针与结构体
	pointersAndStructs()

	// 指针与切片映射
	pointersWithCollections()

	// 指针高级用法
	advancedPointers()

	// 指针安全
	pointerSafety()

	// 实际应用示例
	practicalExample()
}

// 基本指针操作
func basicPointers() {
	fmt.Println("\n--- 基本指针操作 ---")

	// 1. 声明和使用指针
	var x int = 42
	var p *int = &x // p 指向 x 的地址

	fmt.Printf("变量 x 的值: %d\n", x)
	fmt.Printf("变量 x 的地址: %p\n", &x)
	fmt.Printf("指针 p 的值: %p\n", p)
	fmt.Printf("通过指针访问 x 的值: %d\n", *p)

	// 2. 通过指针修改值
	fmt.Printf("\n通过指针修改值:")
	fmt.Printf("修改前 x = %d\n", x)
	*p = 100
	fmt.Printf("修改后 x = %d\n", x)

	// 3. 空指针
	var nilPtr *int
	fmt.Printf("\n空指针:")
	fmt.Printf("nilPtr = %v\n", nilPtr)
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil)

	// 解引用空指针会panic，这里只是演示
	// fmt.Printf(*nilPtr) // 这会导致panic

	// 4. 指针的大小
	fmt.Printf("\n指针的大小:")
	fmt.Printf("int 指针大小: %d 字节\n", unsafe.Sizeof(p))
	fmt.Printf("string 指针大小: %d 字节\n", unsafe.Sizeof((*string)(nil)))
	fmt.Printf("float64 指针大小: %d 字节\n", unsafe.Sizeof((*float64)(nil)))
}

// 指针与函数
func pointersAndFunctions() {
	fmt.Println("\n--- 指针与函数 ---")

	// 1. 值传递（不会修改原变量）
	fmt.Println("值传递:")
	num1 := 10
	fmt.Printf("调用前: num1 = %d\n", num1)
	modifyByValue(num1)
	fmt.Printf("调用后: num1 = %d\n", num1) // 仍然是10

	// 2. 指针传递（会修改原变量）
	fmt.Println("\n指针传递:")
	num2 := 10
	fmt.Printf("调用前: num2 = %d\n", num2)
	modifyByPointer(&num2)
	fmt.Printf("调用后: num2 = %d\n", num2) // 变成100

	// 3. 返回指针
	fmt.Println("\n返回指针:")
	p1 := createIntPointer(42)
	fmt.Printf("返回的指针值: %d\n", *p1)

	// 4. 指针作为返回值的注意事项
	fmt.Println("\n返回局部变量的指针:")
	// 错误示例（函数返回后局部变量会被销毁）
	// p2 := dangerousPointer() // 危险！

	// 正确示例：返回指向堆上分配的内存
	p3 := safePointer()
	fmt.Printf("安全的指针: %d\n", *p3)

	// 5. 多返回值中的指针
	fmt.Println("\n多返回值中的指针:")
	result, err := divideWithPointer(10, 2)
	if err == nil {
		fmt.Printf("结果: %d\n", *result)
	} else {
		fmt.Printf("错误: %v\n", err)
	}
}

// modifyByValue 值传递函数
func modifyByValue(x int) {
	x = 100
	fmt.Printf("函数内部: x = %d\n", x)
}

// modifyByPointer 指针传递函数
func modifyByPointer(p *int) {
	*p = 100
	fmt.Printf("函数内部: *p = %d\n", *p)
}

// createIntPointer 创建指向整数的指针
func createIntPointer(value int) *int {
	return &value
}

// dangerousPointer 危险的指针返回（不要这样做！）
/*
func dangerousPointer() *int {
	x := 42
	return &x // 危险！x是局部变量
}
*/

// safePointer 安全的指针返回
func safePointer() *int {
	// 使用 new 分配内存（在堆上）
	x := new(int)
	*x = 42
	return x
}

// divideWithPointer 返回结果指针和错误
func divideWithPointer(a, b int) (*int, error) {
	if b == 0 {
		return nil, fmt.Errorf("除数不能为零")
	}
	result := a / b
	return &result, nil
}

// 指针与结构体
func pointersAndStructs() {
	fmt.Println("\n--- 指针与结构体 ---")

	// 1. 结构体指针的创建
	fmt.Println("结构体指针创建:")
	p1 := &Person{Name: "张三", Age: 25} // 直接创建指针
	var p2 *Person                       // 先声明指针
	p2 = &Person{Name: "李四", Age: 30}  // 再赋值

	fmt.Printf("p1: %+v\n", *p1)
	fmt.Printf("p2: %+v\n", *p2)

	// 2. 通过指针访问字段
	fmt.Println("\n通过指针访问字段:")
	fmt.Printf("p1.Name = %s\n", p1.Name)     // 自动解引用
	fmt.Printf("(*p1).Name = %s\n", (*p1).Name) // 显式解引用

	// 3. 通过指针修改字段
	fmt.Println("\n通过指针修改字段:")
	fmt.Printf("修改前: %+v\n", *p1)
	p1.Age = 26
	fmt.Printf("修改后: %+v\n", *p1)

	// 4. 结构体指针作为函数参数
	fmt.Println("\n结构体指针作为参数:")
	modifyPerson(p1)
	fmt.Printf("修改后: %+v\n", *p1)

	// 5. 结构体方法中的指针接收者
	fmt.Println("\n指针接收者方法:")
	fmt.Printf("生日前年龄: %d\n", p2.Age)
	p2.Birthday()
	fmt.Printf("生日后年龄: %d\n", p2.Age)
}

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// modifyPerson 修改人员信息
func modifyPerson(p *Person) {
	p.Name = "王五"
	fmt.Printf("函数内部修改: %+v\n", *p)
}

// Birthday 生日方法（指针接收者）
func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("生日快乐！现在 %d 岁了\n", p.Age)
}

// 指针与切片映射
func pointersWithCollections() {
	fmt.Println("\n--- 指针与切片映射 ---")

	// 1. 指针切片
	fmt.Println("指针切片:")
	people := []*Person{
		{Name: "张三", Age: 25},
		{Name: "李四", Age: 30},
		{Name: "王五", Age: 28},
	}

	fmt.Printf("原始人员列表:\n")
	for i, person := range people {
		fmt.Printf("%d. %s (%d岁)\n", i+1, person.Name, person.Age)
	}

	// 修改切片中的指针元素
	fmt.Println("\n修改切片中的元素:")
	people[0].Age = 26
	fmt.Printf("修改后张三的年龄: %d\n", people[0].Age)

	// 2. 指针映射
	fmt.Println("\n指针映射:")
	userMap := map[string]*Person{
		"user1": {Name: "用户1", Age: 20},
		"user2": {Name: "用户2", Age: 25},
	}

	fmt.Printf("用户映射:\n")
	for id, user := range userMap {
		fmt.Printf("%s: %s (%d岁)\n", id, user.Name, user.Age)
	}

	// 修改映射中的指针元素
	fmt.Println("\n修改映射中的元素:")
	userMap["user1"].Age = 21
	fmt.Printf("修改后用户1的年龄: %d\n", userMap["user1"].Age)

	// 3. 指针的指针
	fmt.Println("\n指针的指针:")
	var ptr1 *int = new(int)
	*ptr1 = 42
	var ptr2 **int = &ptr1

	fmt.Printf("ptr1 的值: %p, 指向的值: %d\n", ptr1, *ptr1)
	fmt.Printf("ptr2 的值: %p, 指向的值: %p, 二级解引用: %d\n", ptr2, *ptr2, **ptr2)
}

// 指针高级用法
func advancedPointers() {
	fmt.Println("\n--- 指针高级用法 ---")

	// 1. 函数指针
	fmt.Println("函数指针:")
	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }

	operations := map[string]func(int, int) int{
		"add":      add,
		"multiply": multiply,
	}

	fmt.Printf("10 + 5 = %d\n", operations["add"](10, 5))
	fmt.Printf("10 * 5 = %d\n", operations["multiply"](10, 5))

	// 2. 接口与指针
	fmt.Println("\n接口与指针:")
	var shaper Shaper = &Circle{Radius: 5}
	fmt.Printf("圆形面积: %.2f\n", shaper.Area())

	// 注意：值类型也可以实现接口，但指针类型更灵活
	// var shaper2 Shaper = Circle{Radius: 3} // 也可以

	// 3. 空接口与指针
	fmt.Println("\n空接口与指针:")
	var data interface{} = &Person{Name: "指针人员", Age: 25}

	if person, ok := data.(*Person); ok {
		fmt.Printf("类型断言成功: %+v\n", *person)
	}

	// 4. unsafe.Pointer
	fmt.Println("\nunsafe.Pointer:")
	num := int64(0x123456789ABCDEF0)
	ptr := unsafe.Pointer(&num)

	fmt.Printf("原始值: 0x%x\n", num)
	fmt.Printf("指针地址: %p\n", ptr)

	// 将 unsafe.Pointer 转换为其他类型指针
	intPtr := (*int64)(ptr)
	fmt.Printf("转换后值: 0x%x\n", *intPtr)
}

// Shaper 接口
type Shaper interface {
	Area() float64
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形面积（指针接收者）
func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// 指针安全
func pointerSafety() {
	fmt.Println("\n--- 指针安全 ---")

	// 1. 空指针检查
	fmt.Println("空指针检查:")
	var p *int
	if p != nil {
		fmt.Printf("指针值: %d\n", *p)
	} else {
		fmt.Println("指针是 nil，不能解引用")
	}

	// 2. 安全的指针操作
	fmt.Println("\n安全的指针操作:")
	safeAccess(nil)

	safePtr := &Person{Name: "张三", Age: 25}
	safeAccess(safePtr)

	// 3. 避免悬垂指针
	fmt.Println("\n避免悬垂指针:")
	// 好的实践：不要返回局部变量的地址
	badPractice()
	goodPractice()
}

// safeAccess 安全访问指针
func safeAccess(p *Person) {
	if p != nil {
		fmt.Printf("安全访问: %s (%d岁)\n", p.Name, p.Age)
	} else {
		fmt.Println("指针为 nil，无法访问")
	}
}

// badPractice 不好的实践（示例）
func badPractice() {
	// 不要这样做！
	/*
	localVar := 42
	return &localVar // localVar 函数结束后被销毁
	*/
	fmt.Println("不要返回局部变量的地址")
}

// goodPractice 好的实践
func goodPractice() {
	// 好的做法：使用 new 或分配到全局变量
	result := new(int)
	*result = 42
	fmt.Printf("安全的指针操作: %d\n", *result)
}

// 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1：链表实现
	fmt.Println("1. 链表实现:")
	// 创建链表: 1 -> 2 -> 3
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}

	fmt.Printf("链表内容: ")
	for current := head; current != nil; current = current.Next {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()

	// 在链表头部插入新节点
	newNode := &ListNode{Value: 0, Next: head}
	head = newNode

	fmt.Printf("插入后: ")
	for current := head; current != nil; current = current.Next {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()

	// 示例2：树结构
	fmt.Println("\n2. 树结构:")
	root := &TreeNode{
		Value: "Root",
		Left: &TreeNode{
			Value: "Left",
			Left:  &TreeNode{Value: "Left-Left"},
			Right: &TreeNode{Value: "Left-Right"},
		},
		Right: &TreeNode{
			Value: "Right",
			Left:  &TreeNode{Value: "Right-Left"},
		},
	}

	fmt.Printf("树的先序遍历: ")
	preorder(root)
	fmt.Println()

	// 示例3：缓存系统
	fmt.Println("\n3. 简单缓存系统:")
	cache := NewCache()

	// 存储数据
	cache.Set("user:1", &User{ID: 1, Name: "张三"})
	cache.Set("user:2", &User{ID: 2, Name: "李四"})

	// 获取数据
	if user, found := cache.Get("user:1"); found {
		fmt.Printf("缓存命中: %+v\n", *user)
	}

	// 修改缓存中的数据
	if user, found := cache.Get("user:2"); found {
		user.Name = "李四改"
		fmt.Printf("修改后: %+v\n", *user)
	}

	// 示例4：对象池
	fmt.Println("\n4. 对象池模式:")
	pool := NewBufferPool(3)

	// 获取缓冲区
	buf1 := pool.Get()
	buf2 := pool.Get()
	buf3 := pool.Get()

	fmt.Printf("获取缓冲区: %p, %p, %p\n", buf1, buf2, buf3)

	// 归还缓冲区
	pool.Put(buf1)
	pool.Put(buf2)

	buf4 := pool.Get() // 应该重用刚归还的缓冲区
	fmt.Printf("重用缓冲区: %p\n", buf4)
}

// ListNode 链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

// TreeNode 树节点
type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

// User 用户结构
type User struct {
	ID   int
	Name string
}

// preorder 先序遍历
func preorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%s ", node.Value)
	preorder(node.Left)
	preorder(node.Right)
}

// Cache 简单缓存
type Cache struct {
	data map[string]interface{}
}

// NewCache 创建新缓存
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set 设置缓存
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get 获取缓存
func (c *Cache) Get(key string) (interface{}, bool) {
	value, exists := c.data[key]
	return value, exists
}

// BufferPool 缓冲区池
type BufferPool struct {
	pool [][]byte
	used int
}

// NewBufferPool 创建缓冲区池
func NewBufferPool(size int) *BufferPool {
	return &BufferPool{
		pool: make([][]byte, size),
		used: 0,
	}
}

// Get 获取缓冲区
func (bp *BufferPool) Get() []byte {
	if bp.used < len(bp.pool) {
		buf := bp.pool[bp.used]
		bp.used++
		if buf == nil {
			buf = make([]byte, 1024) // 1KB 缓冲区
		}
		return buf
	}
	return make([]byte, 1024) // 池满时创建新的
}

// Put 归还缓冲区
func (bp *BufferPool) Put(buf []byte) {
	if bp.used > 0 {
		bp.used--
		bp.pool[bp.used] = buf[:0] // 重置长度但保留容量
	}
}