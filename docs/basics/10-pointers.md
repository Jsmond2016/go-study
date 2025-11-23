---
title: 指针
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["变量与常量", "数据类型", "函数", "结构体"]
tags: ["指针", "内存地址", "引用", "解引用"]
---

# 指针

指针是 Go 语言中一个重要但相对简单的概念，用于存储变量的内存地址。理解指针对于编写高效的 Go 程序至关重要。

## 📋 学习目标

- [ ] 理解指针的概念和用途
- [ ] 掌握指针的声明和使用
- [ ] 学会指针的解引用操作
- [ ] 理解指针与函数参数的关系
- [ ] 掌握指针的高级用法
- [ ] 了解指针的安全性和最佳实践

## 🎯 指针基础

### 什么是指针

```go
package main

import "fmt"

func main() {
	// 基本变量
	var x int = 42
	
	// 获取变量的地址
	var p *int = &x
	
	fmt.Printf("变量 x 的值: %d\n", x)
	fmt.Printf("变量 x 的地址: %p\n", &x)
	fmt.Printf("指针 p 的值: %p\n", p)
	fmt.Printf("指针 p 指向的值: %d\n", *p)
	fmt.Printf("指针 p 自身的地址: %p\n", &p)
	
	// 通过指针修改原变量
	*p = 100
	fmt.Printf("通过指针修改后 x 的值: %d\n", x)
	
	// 指针的零值
	var nilPtr *int
	fmt.Printf("nil 指针: %v\n", nilPtr)
	fmt.Printf("nil 指针是否为 nil: %t\n", nilPtr == nil)
	
	// 解引用 nil 指针会导致 panic
	// fmt.Printf(*nilPtr) // panic: runtime error: invalid memory address
}
```

### 指针声明和初始化

```go
package main

import "fmt"

func main() {
	// 声明指针变量
	var p1 *int
	var p2 *string
	var p3 *bool
	
	fmt.Printf("未初始化的指针: %p, %p, %p\n", p1, p2, p3)
	fmt.Printf("是否为 nil: %t, %t, %t\n", p1 == nil, p2 == nil, p3 == nil)
	
	// 使用 new 创建指针
	p1 = new(int)
	p2 = new(string)
	p3 = new(bool)
	
	fmt.Printf("\n使用 new 创建后:\n")
	fmt.Printf("p1 指向的值: %d\n", *p1) // 零值
	fmt.Printf("p2 指向的值: %q\n", *p2) // 空字符串
	fmt.Printf("p3 指向的值: %t\n", *p3) // false
	
	// 使用取地址操作符
	x := 42
	name := "Go"
	flag := true
	
	p1 = &x
	p2 = &name
	p3 = &flag
	
	fmt.Printf("\n取地址后:\n")
	fmt.Printf("p1 -> %d\n", *p1)
	fmt.Printf("p2 -> %s\n", *p2)
	fmt.Printf("p3 -> %t\n", *p3)
}
```

## 🔗 指针与函数

### 值传递 vs 指针传递

```go
package main

import "fmt"

// 值传递（不会修改原变量）
func modifyValue(x int) {
	x = 100
	fmt.Printf("函数内部 x 的值: %d\n", x)
}

// 指针传递（会修改原变量）
func modifyPointer(x *int) {
	*x = 100
	fmt.Printf("函数内部 *x 的值: %d\n", *x)
}

func main() {
	// 值传递测试
	value1 := 42
	fmt.Printf("修改前 value1: %d\n", value1)
	modifyValue(value1)
	fmt.Printf("修改后 value1: %d\n", value1) // 值未改变
	
	fmt.Println("---")
	
	// 指针传递测试
	value2 := 42
	fmt.Printf("修改前 value2: %d\n", value2)
	modifyPointer(&value2)
	fmt.Printf("修改后 value2: %d\n", value2) // 值被改变
}
```

### 结构体指针

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 值接收者方法
func (p Person) SetName(name string) {
	p.Name = name // 修改的是副本
}

// 指针接收者方法
func (p *Person) SetNamePtr(name string) {
	p.Name = name // 修改的是原对象
}

func (p *Person) Birthday() {
	p.Age++
}

func (p Person) Print() {
	fmt.Printf("姓名: %s, 年龄: %d\n", p.Name, p.Age)
}

func main() {
	// 值调用
	person1 := Person{Name: "张三", Age: 25}
	person1.SetName("李四")
	person1.Print() // 姓名未改变
	
	// 指针调用
	person2 := &Person{Name: "张三", Age: 25}
	person2.SetNamePtr("李四")
	person2.Print() // 姓名已改变
	
	// Go 会自动解引用
	person2.Birthday() // 等价于 (&person2).Birthday()
	person2.Print()
	
	// 结构体字面量指针
	person3 := &Person{Name: "王五", Age: 30}
	person3.Print()
}
```

## 🎯 指针操作

### 指针解引用和地址获取

```go
package main

import "fmt"

func main() {
	// 基本操作
	x := 42
	p := &x
	
	fmt.Printf("变量值: %d\n", x)
	fmt.Printf("变量地址: %p\n", &x)
	fmt.Printf("指针值: %p\n", p)
	fmt.Printf("指针指向的值: %d\n", *p)
	fmt.Printf("指针地址: %p\n", &p)
	
	// 多重解引用
	pp := &p
	fmt.Printf("\n双重指针:\n")
	fmt.Printf("pp 的值: %p\n", pp)
	fmt.Printf("*pp 的值: %p\n", *pp)
	fmt.Printf("**pp 的值: %d\n", **pp)
	
	// 通过双重指针修改原变量
	**pp = 200
	fmt.Printf("修改后 x 的值: %d\n", x)
}
```

### 指针算术（Go 不支持）

```go
package main

import "fmt"

func main() {
	// Go 不支持指针算术运算
	// 这是 Go 与 C 的一个重要区别
	
	arr := [5]int{10, 20, 30, 40, 50}
	p := &arr[0]
	
	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("第一个元素地址: %p\n", p)
	fmt.Printf("第一个元素值: %d\n", *p)
	
	// 以下操作在 Go 中不被允许：
	// p++     // 编译错误
	// p = p + 1 // 编译错误
	// p += 2  // 编译错误
	
	// 正确的遍历方式
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d (地址: %p)\n", i, arr[i], &arr[i])
	}
}
```

## 🔄 指针的高级用法

### 函数返回指针

```go
package main

import "fmt"

// 返回局部变量指针（不推荐，但在 Go 中是安全的）
func createInt() *int {
	x := 42
	return &x // Go 会进行逃逸分析，将 x 分配到堆上
}

// 返回堆分配的指针
func newInt() *int {
	return new(int)
}

// 返回指向结构体的指针
func newPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// 从函数获取指针
	p1 := createInt()
	fmt.Printf("createInt() 返回: %d\n", *p1)
	
	p2 := newInt()
	*p2 = 100
	fmt.Printf("newInt() 赋值后: %d\n", *p2)
	
	// 从函数获取结构体指针
	p3 := newPerson("张三", 25)
	fmt.Printf("newPerson() 返回: %+v\n", *p3)
	
	// 修改结构体
	p3.Age = 26
	fmt.Printf("修改年龄后: %+v\n", *p3)
}
```

### 指针数组和数组指针

```go
package main

import "fmt"

func main() {
	// 指针数组
	x, y, z := 10, 20, 30
	ptrArray := [3]*int{&x, &y, &z}
	
	fmt.Printf("指针数组:\n")
	for i, p := range ptrArray {
		fmt.Printf("ptrArray[%d] = %p -> %d\n", i, p, *p)
	}
	
	// 修改通过指针数组访问的值
	*ptrArray[0] = 100
	fmt.Printf("修改后 x = %d\n", x)
	
	fmt.Println("---")
	
	// 数组指针
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	
	fmt.Printf("数组指针:\n")
	fmt.Printf("数组地址: %p\n", &arr)
	fmt.Printf("指针值: %p\n", arrPtr)
	fmt.Printf("通过指针访问数组: %v\n", *arrPtr)
	fmt.Printf("访问第一个元素: %d\n", (*arrPtr)[0])
	fmt.Printf("简写访问第一个元素: %d\n", arrPtr[0]) // Go 自动解引用
	
	// 修改数组元素
	arrPtr[1] = 200
	fmt.Printf("修改后数组: %v\n", arr)
}
```

### 切片和指针

```go
package main

import "fmt"

func main() {
	// 切片头信息
	slice := []int{10, 20, 30, 40, 50}
	
	fmt.Printf("切片: %v\n", slice)
	fmt.Printf("切片地址: %p\n", &slice)
	fmt.Printf("底层数组地址: %p\n", slice)
	
	// 获取切片头指针
	slicePtr := &slice
	fmt.Printf("切片头指针: %p\n", slicePtr)
	fmt.Printf("通过指针访问切片: %v\n", *slicePtr)
	
	// 获取第一个元素的指针
	firstElementPtr := &slice[0]
	fmt.Printf("第一个元素指针: %p\n", firstElementPtr)
	fmt.Printf("第一个元素值: %d\n", *firstElementPtr)
	
	// 通过指针修改元素
	*firstElementPtr = 100
	fmt.Printf("修改后切片: %v\n", slice)
	
	// 创建指向切片的指针并修改
	newSlice := []string{"Go", "Python", "Java"}
	newSlicePtr := &newSlice
	
	(*newSlicePtr)[1] = "Rust"
	fmt.Printf("修改后: %v\n", newSlice)
}
```

## 🛡️ 指针安全性

### nil 指针检查

```go
package main

import "fmt"

func safeDereference(ptr *int) {
	if ptr == nil {
		fmt.Println("指针为 nil，无法解引用")
		return
	}
	
	fmt.Printf("指针指向的值: %d\n", *ptr)
}

func modifySafely(ptr *int, value int) error {
	if ptr == nil {
		return fmt.Errorf("不能修改 nil 指针")
	}
	
	*ptr = value
	return nil
}

func main() {
	var nilPtr *int
	
	// 安全的解引用
	safeDereference(nilPtr)
	
	// 安全的修改
	err := modifySafely(nilPtr, 42)
	if err != nil {
		fmt.Printf("修改失败: %v\n", err)
	}
	
	// 正常情况
	value := 10
	safeDereference(&value)
	
	err = modifySafely(&value, 20)
	if err == nil {
		fmt.Printf("修改成功，新值: %d\n", value)
	}
}
```

### 避免指针陷阱

```go
package main

import "fmt"

func main() {
	// 陷阱 1：返回局部变量的地址（在 Go 中是安全的）
	fmt.Println("=== 陷阱 1：返回局部变量地址 ===")
	p1 := getLocalVar()
	fmt.Printf("*p1 = %d\n", *p1) // Go 会进行逃逸分析
	
	// 陷阱 2：使用临时变量的指针
	fmt.Println("\n=== 陷阱 2：临时变量指针 ===")
	var p2 *int
	{
		temp := 42
		p2 = &temp // temp 会逃逸到堆上
	}
	fmt.Printf("*p2 = %d\n", *p2) // 在 Go 中是安全的
	
	// 陷阱 3：循环中的指针问题
	fmt.println("\n=== 陷阱 3：循环中的指针 ===")
	values := []int{10, 20, 30}
	var ptrs []*int
	
	for _, v := range values {
		v := v // 创建新变量（重要！）
		ptrs = append(ptrs, &v)
	}
	
	for i, p := range ptrs {
		fmt.Printf("ptrs[%d] = %d\n", i, *p)
	}
	
	// 陷阱 4：悬空指针（在 Go 中较少见）
	fmt.Println("\n=== 陷阱 4：避免悬空指针 ===")
	safePtr := createSafePtr()
	fmt.Printf("*safePtr = %d\n", *safePtr)
}

func getLocalVar() *int {
	x := 100
	return &x // Go 会将 x 分配到堆上
}

func createSafePtr() *int {
	return new(int) // 显式创建
}
```

## 🎯 实际应用

### 链表节点

```go
package main

import "fmt"

// 链表节点
type Node struct {
	Value int
	Next  *Node
}

// 链表
type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Add(value int) {
	newNode := &Node{Value: value}
	
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Remove(value int) bool {
	if ll.Head == nil {
		return false
	}
	
	// 删除头节点
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}
	
	// 删除其他节点
	prev := ll.Head
	current := ll.Head.Next
	
	for current != nil {
		if current.Value == value {
			prev.Next = current.Next
			ll.Size--
			return true
		}
		prev = current
		current = current.Next
	}
	
	return false
}

func (ll *LinkedList) Print() {
	current := ll.Head
	fmt.Printf("链表 (长度: %d): ", ll.Size)
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Printf(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	ll := NewLinkedList()
	
	// 添加元素
	for i := 1; i <= 5; i++ {
		ll.Add(i * 10)
	}
	ll.Print()
	
	// 删除元素
	ll.Remove(30)
	ll.Print()
	
	ll.Remove(10)
	ll.Print()
}
```

### 内存池

```go
package main

import (
	"fmt"
	"sync"
)

// 对象池
type Object struct {
	ID   int
	Data string
}

type ObjectPool struct {
	pool    sync.Pool
	counter int
	mu      sync.Mutex
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &Object{}
			},
		},
	}
}

func (op *ObjectPool) Get() *Object {
	obj := op.pool.Get().(*Object)
	
	op.mu.Lock()
	op.counter++
	obj.ID = op.counter
	op.mu.Unlock()
	
	return obj
}

func (op *ObjectPool) Put(obj *Object) {
	obj.Data = "" // 重置数据
	op.pool.Put(obj)
}

func main() {
	pool := NewObjectPool()
	
	// 获取对象
	obj1 := pool.Get()
	obj1.Data = "Hello World"
	fmt.Printf("obj1: %+v\n", obj1)
	
	obj2 := pool.Get()
	obj2.Data = "Go Programming"
	fmt.Printf("obj2: %+v\n", obj2)
	
	// 归还对象
	pool.Put(obj1)
	pool.Put(obj2)
	
	// 再次获取（可能会重用之前归还的对象）
	obj3 := pool.Get()
	obj3.Data = "New Object"
	fmt.Printf("obj3: %+v\n", obj3)
}
```

## 🧪 实践练习

### 练习 1：二叉树实现

```go
// 要求：
// 1. 定义 TreeNode 结构体
// 2. 实现插入、查找、删除功能
// 3. 实现前序、中序、后序遍历
// 4. 计算树的深度和节点数
```

参考实现：
```go
package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(value int) {
	newNode := &TreeNode{Value: value}
	
	if bt.Root == nil {
		bt.Root = newNode
		return
	}
	
	current := bt.Root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				break
			}
			current = current.Right
		}
	}
}

func (bt *BinaryTree) InOrderTraversal(node *TreeNode) {
	if node != nil {
		bt.InOrderTraversal(node.Left)
		fmt.Printf("%d ", node.Value)
		bt.InOrderTraversal(node.Right)
	}
}

func main() {
	bt := NewBinaryTree()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	
	for _, value := range values {
		bt.Insert(value)
	}
	
	fmt.Print("中序遍历: ")
	bt.InOrderTraversal(bt.Root)
	fmt.Println()
}
```

### 练习 2：智能指针

```go
// 要求：
// 1. 实现引用计数
// 2. 自动内存管理
// 3. 支持共享所有权
// 4. 线程安全实现
```

### 练习 3：缓存系统

```go
// 要求：
// 1. 实现 LRU 缓存
// 2. 使用指针优化性能
// 3. 支持并发访问
// 4. 实现缓存淘汰策略
```

## 🤔 思考题

1. **什么时候应该使用指针，什么时候使用值？**
   - 提示：考虑性能、内存使用、修改需求

2. **Go 的指针和 C 的指针有什么区别？**
   - 提示：考虑指针算术、类型安全、垃圾回收

3. **为什么 Go 不支持指针算术？**
   - 提示：考虑安全性、内存管理

4. **指针的逃逸分析是什么？**
   - 提示：考虑栈分配 vs 堆分配、性能影响

## 📚 扩展阅读

- [Go 语言规范 - 指针类型](https://golang.org/ref/spec#Pointer_types)
- [Effective Go - 指针](https://golang.org/doc/effective_go.html#pointers_vs_values)
- [Go by Example - 指针](https://gobyexample.com/pointers)

## ⏭️ 下一章节

[接口](./11-interfaces.md) → 学习 Go 语言的接口

---

**💡 小贴士**: 指针是 Go 语言中提高性能和实现修改的重要工具。记住：**Go 的指针比 C 更安全，不支持算术运算，但足够高效**！
