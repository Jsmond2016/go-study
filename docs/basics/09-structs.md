---
title: 结构体
difficulty: beginner
duration: "4-5小时"
prerequisites: ["变量与常量", "数据类型", "函数", "映射"]
tags: ["结构体", "方法", "组合", "面向对象"]
---

# 结构体

结构体（Struct）是 Go 语言中最重要的复合数据类型，用于将不同类型的数据组合成一个逻辑单元。结构体是 Go 实现面向对象编程的基础。

## 📋 学习目标

- [ ] 掌握结构体的定义和创建
- [ ] 学会结构体的初始化方法
- [ ] 理解结构体的内存布局
- [ ] 掌握结构体的嵌套和组合
- [ ] 学会为结构体定义方法
- [ ] 了解结构体标签的用途

## 🏗️ 结构体基础

### 结构体定义

```go
package main

import "fmt"

// 基本结构体定义
type Person struct {
	Name string
	Age  int
	City string
}

// 带标签的结构体
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

// 匿名结构体
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

func main() {
	// 结构体零值
	var p1 Person
	fmt.Printf("零值结构体: %+v\n", p1)
	
	// 结构体字面量初始化
	p2 := Person{
		Name: "张三",
		Age:  25,
		City: "北京",
	}
	fmt.Printf("完整初始化: %+v\n", p2)
	
	// 部分初始化
	p3 := Person{
		Name: "李四",
		Age:  30,
	}
	fmt.Printf("部分初始化: %+v\n", p3)
	
	// 按顺序初始化（不推荐，易出错）
	p4 := Person{"王五", 28, "上海"}
	fmt.Printf("顺序初始化: %+v\n", p4)
	
	// 使用 new 创建
	p5 := new(Person) // 返回指针
	p5.Name = "赵六"
	p5.Age = 22
	p5.City = "广州"
	fmt.Printf("new创建: %+v\n", p5)
	
	// 使用 & 创建
	p6 := &Person{
		Name: "钱七",
		Age:  35,
		City: "深圳",
	}
	fmt.Printf("&创建: %+v\n", *p6)
	fmt.Printf("地址: %p\n", p6)
}
```

### 结构体字段访问和修改

```go
package main

import "fmt"

type Employee struct {
	ID        int
	Name      string
	Department string
	Salary    float64
	IsActive  bool
}

func main() {
	// 创建员工
	emp := Employee{
		ID:        1001,
		Name:      "张三",
		Department: "技术部",
		Salary:    15000.00,
		IsActive:  true,
	}
	
	// 访问字段
	fmt.Printf("员工信息:\n")
	fmt.Printf("  ID: %d\n", emp.ID)
	fmt.Printf("  姓名: %s\n", emp.Name)
	fmt.Printf("  部门: %s\n", emp.Department)
	fmt.Printf("  薪资: %.2f\n", emp.Salary)
	fmt.Printf("  在职: %t\n", emp.IsActive)
	
	// 修改字段
	emp.Salary = 18000.00
	emp.Department = "研发部"
	fmt.Printf("\n调整后:\n")
	fmt.Printf("  部门: %s\n", emp.Department)
	fmt.Printf("  薪资: %.2f\n", emp.Salary)
	
	// 使用指针修改
	empPtr := &emp
	empPtr.Name = "张三丰"
	fmt.Printf("\n指针修改后姓名: %s\n", emp.Name)
}
```

## 🎯 结构体方法

### 方法定义

```go
package main

import "fmt"

// Rectangle 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// 值接收者方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 指针接收者方法
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (r *Rectangle) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}

func main() {
	// 创建矩形
	rect := Rectangle{Width: 10.0, Height: 5.0}
	
	// 调用值方法
	area := rect.Area()
	perimeter := rect.Perimeter()
	
	fmt.Printf("矩形: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)
	fmt.Printf("面积: %.1f\n", area)
	fmt.Printf("周长: %.1f\n", perimeter)
	
	// 调用指针方法（自动取地址）
	rect.Scale(2.0)
	fmt.Printf("\n缩放2倍后: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)
	fmt.Printf("新面积: %.1f\n", rect.Area())
	
	// 使用指针
	rectPtr := &rect
	rectPtr.SetDimensions(15.0, 8.0)
	fmt.Printf("\n重新设置后: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)
}
```

### 方法选择：值接收者 vs 指针接收者

```go
package main

import "fmt"

type Counter struct {
	count int
}

// 值接收者方法（不修改状态）
func (c Counter) Value() int {
	return c.count
}

// 指针接收者方法（修改状态）
func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Reset() {
	c.count = 0
}

// 值接收者（返回新值）
func (c Counter) Add(n int) Counter {
	return Counter{count: c.count + n}
}

func main() {
	// 使用值
	counter := Counter{count: 5}
	fmt.Printf("初始值: %d\n", counter.Value())
	
	// 修改状态需要指针
	counter.Increment()
	fmt.Printf("递增后: %d\n", counter.Value())
	
	// 值方法不会修改原对象
	newCounter := counter.Add(10)
	fmt.Printf("原计数器: %d\n", counter.Value())
	fmt.Printf("新计数器: %d\n", newCounter.Value())
	
	// 重置
	counter.Reset()
	fmt.Printf("重置后: %d\n", counter.Value())
}
```

## 🔄 结构体组合

### 匿名字段组合

```go
package main

import "fmt"

// 基础结构体
type Animal struct {
	Name string
	Age  int
}

func (a Animal) Speak() string {
	return fmt.Sprintf("%s 发出了声音", a.Name)
}

// 组合结构体
type Dog struct {
	Animal  // 匿名字段（嵌入）
	Breed string
}

func (d Dog) Bark() string {
	return fmt.Sprintf("%s 汪汪叫", d.Name)
}

// 重写方法
func (d Dog) Speak() string {
	return fmt.Sprintf("%s 摇着尾巴汪汪叫", d.Name)
}

func main() {
	// 创建狗
	dog := Dog{
		Animal: Animal{
			Name: "旺财",
			Age:  3,
		},
		Breed: "金毛",
	}
	
	// 访问嵌入字段
	fmt.Printf("狗名: %s\n", dog.Name) // 直接访问
	fmt.Printf("年龄: %d\n", dog.Age)
	fmt.Printf("品种: %s\n", dog.Breed)
	
	// 调用方法
	fmt.Printf("Speak: %s\n", dog.Speak())     // 调用重写的方法
	fmt.Printf("Bark: %s\n", dog.Bark())       // 调用自己的方法
	fmt.Printf("Animal.Speak: %s\n", dog.Animal.Speak()) // 调用嵌入的方法
}
```

### 有名字段组合

```go
package main

import "fmt"

type Engine struct {
	Power  int
	Fuel   string
	Volume float64
}

func (e Engine) Start() {
	fmt.Printf("发动机启动，功率: %d马力\n", e.Power)
}

func (e Engine) Stop() {
	fmt.Println("发动机关闭")
}

type Car struct {
	engine  Engine // 有名字段
	Brand   string
	Model   string
	Year    int
}

func (c Car) Start() {
	c.engine.Start()
	fmt.Printf("%s %s 启动成功\n", c.Brand, c.Model)
}

func (c Car) Stop() {
	fmt.Printf("%s %s 停车\n", c.Brand, c.Model)
	c.engine.Stop()
}

func (c Car) GetEngineInfo() {
	fmt.Printf("发动机功率: %d马力, 燃料: %s\n", 
		c.engine.Power, c.engine.Fuel)
}

func main() {
	// 创建汽车
	car := Car{
		engine: Engine{
			Power:  200,
			Fuel:   "汽油",
			Volume: 2.0,
		},
		Brand: "丰田",
		Model: "凯美瑞",
		Year:  2020,
	}
	
	// 访问字段
	fmt.Printf("汽车: %s %s %d年\n", car.Brand, car.Model, car.Year)
	car.GetEngineInfo()
	
	// 调用方法
	car.Start()
	car.Stop()
}
```

## 🏷️ 结构体标签

### JSON 标签

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"product_name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
	InStock     bool    `json:"in_stock"`
	Tags        []string `json:"tags"`
}

func main() {
	// 创建产品
	product := Product{
		ID:      1001,
		Name:    "智能手机",
		Price:   2999.99,
		InStock: true,
		Tags:    []string{"电子", "手机", "5G"},
	}
	
	// 转换为 JSON
	jsonData, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		fmt.Printf("JSON 编码错误: %v\n", err)
		return
	}
	
	fmt.Printf("JSON 输出:\n%s\n", string(jsonData))
	
	// 从 JSON 解析
	jsonStr := `{
  "id": 1002,
  "product_name": "笔记本电脑",
  "price": 5999.99,
  "in_stock": false
}`
	
	var newProduct Product
	err = json.Unmarshal([]byte(jsonStr), &newProduct)
	if err != nil {
		fmt.Printf("JSON 解码错误: %v\n", err)
		return
	}
	
	fmt.Printf("\n解析结果:\n")
	fmt.Printf("ID: %d\n", newProduct.ID)
	fmt.Printf("名称: %s\n", newProduct.Name)
	fmt.Printf("价格: %.2f\n", newProduct.Price)
	fmt.Printf("库存: %t\n", newProduct.InStock)
}
```

### 自定义标签解析

```go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 自定义标签解析器
type Validator struct{}

func (v *Validator) Validate(s interface{}) error {
	val := reflect.ValueOf(s).Elem()
	typ := val.Type()
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")
		
		if tag != "" {
			// 解析验证规则
			rules := strings.Split(tag, ",")
			for _, rule := range rules {
				switch rule {
				case "required":
					if field.IsZero() {
						return fmt.Errorf("字段 %s 是必需的", fieldType.Name)
					}
				case "min=3":
					if str := field.String(); len(str) < 3 {
						return fmt.Errorf("字段 %s 长度不能少于3个字符", fieldType.Name)
					}
				case "email":
					email := field.String()
					if !strings.Contains(email, "@") {
						return fmt.Errorf("字段 %s 必须是有效的邮箱地址", fieldType.Name)
					}
				}
			}
		}
	}
	
	return nil
}

type UserProfile struct {
	Username string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"required"`
	Country  string `validate:"min=2"`
}

func main() {
	// 测试验证
	validator := Validator{}
	
	// 有效数据
	validProfile := UserProfile{
		Username: "john_doe",
		Email:    "john@example.com",
		Age:      25,
		Country:  "CN",
	}
	
	err := validator.Validate(&validProfile)
	if err == nil {
		fmt.Println("✅ 有效配置验证通过")
	} else {
		fmt.Printf("❌ 验证失败: %v\n", err)
	}
	
	// 无效数据
	invalidProfile := UserProfile{
		Username: "jo", // 太短
		Email:    "invalid-email", // 无效邮箱
	}
	
	err = validator.Validate(&invalidProfile)
	if err != nil {
		fmt.Printf("\n❌ 无效配置: %v\n", err)
	} else {
		fmt.Println("✅ 验证通过（不应该出现）")
	}
}
```

## 🎯 实际应用

### 配置结构体

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 数据库配置
type DatabaseConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

// 服务器配置
type ServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// 完整应用配置
type AppConfig struct {
	Server   ServerConfig   `json:"server" yaml:"server"`
	Database DatabaseConfig `json:"database" yaml:"database"`
	Debug    bool           `json:"debug" yaml:"debug"`
	Version  string         `json:"version" yaml:"version"`
}

func (c *AppConfig) Print() {
	fmt.Printf("=== 应用配置 ===\n")
	fmt.Printf("版本: %s\n", c.Version)
	fmt.Printf("调试模式: %t\n", c.Debug)
	
	fmt.Printf("\n--- 服务器配置 ---\n")
	fmt.Printf("地址: %s:%d\n", c.Server.Host, c.Server.Port)
	fmt.Printf("模式: %s\n", c.Server.Mode)
	
	fmt.Printf("\n--- 数据库配置 ---\n")
	fmt.Printf("数据库: %s@%s:%d/%s\n", 
		c.Database.Username, c.Database.Host, c.Database.Port, c.Database.Database)
}

func (c *AppConfig) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}
	
	err = json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}
	
	return nil
}

func (c *AppConfig) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}
	
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}
	
	return nil
}

func main() {
	// 创建默认配置
	config := AppConfig{
		Server: ServerConfig{
			Host: "localhost",
			Port: 8080,
			Mode: "development",
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "admin",
			Password: "password",
			Database: "myapp",
		},
		Debug:   true,
		Version: "1.0.0",
	}
	
	fmt.Println("默认配置:")
	config.Print()
	
	// 保存到文件
	err := config.SaveToFile("config.json")
	if err != nil {
		fmt.Printf("保存配置失败: %v\n", err)
		return
	}
	
	// 从文件加载
	var newConfig AppConfig
	err = newConfig.LoadFromFile("config.json")
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}
	
	fmt.Println("\n从文件加载的配置:")
	newConfig.Print()
}
```

### 链表实现

```go
package main

import "fmt"

// 链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

// 链表
type LinkedList struct {
	Head *ListNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Add(value int) {
	newNode := &ListNode{Value: value}
	
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

func (ll *LinkedList) Contains(value int) bool {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList) Print() {
	current := ll.Head
	fmt.Printf("链表内容 (长度: %d): ", ll.Size)
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
	// 创建链表
	ll := NewLinkedList()
	
	// 添加元素
	values := []int{10, 20, 30, 40, 50}
	for _, value := range values {
		ll.Add(value)
	}
	ll.Print()
	
	// 查找元素
	fmt.Printf("包含 30: %t\n", ll.Contains(30))
	fmt.Printf("包含 99: %t\n", ll.Contains(99))
	
	// 删除元素
	ll.Remove(30)
	fmt.Println("删除 30 后:")
	ll.Print()
	
	ll.Remove(10)
	fmt.Println("删除 10 后:")
	ll.Print()
}
```

## 🧪 实践练习

### 练习 1：图书管理系统

```go
// 要求：
// 1. 定义 Book 结构体
// 2. 实现图书的增删改查
// 3. 支持按作者、分类查找
// 4. 实现借阅状态管理
```

参考实现：
```go
package main

import (
	"fmt"
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publish_date"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	IsAvailable bool      `json:"is_available"`
}

type Library struct {
	books map[int]*Book
	nextID int
}

func NewLibrary() *Library {
	return &Library{
		books:  make(map[int]*Book),
		nextID: 1,
	}
}

func (l *Library) AddBook(title, author, isbn, publisher, category string, 
	publishDate time.Time, price float64) *Book {
	book := &Book{
		ID:          l.nextID,
		Title:       title,
		Author:      author,
		ISBN:        isbn,
		Publisher:   publisher,
		PublishDate: publishDate,
		Category:    category,
		Price:       price,
		IsAvailable: true,
	}
	
	l.books[book.ID] = book
	l.nextID++
	
	return book
}

func (l *Library) FindByID(id int) (*Book, bool) {
	book, exists := l.books[id]
	return book, exists
}

func (l *Library) FindByAuthor(author string) []*Book {
	var books []*Book
	for _, book := range l.books {
		if book.Author == author {
			books = append(books, book)
		}
	}
	return books
}

func (l *Library) BorrowBook(id int) error {
	book, exists := l.FindByID(id)
	if !exists {
		return fmt.Errorf("图书不存在")
	}
	
	if !book.IsAvailable {
		return fmt.Errorf("图书已被借出")
	}
	
	book.IsAvailable = false
	return nil
}

func (l *Library) ReturnBook(id int) error {
	book, exists := l.FindByID(id)
	if !exists {
		return fmt.Errorf("图书不存在")
	}
	
	book.IsAvailable = true
	return nil
}

func main() {
	library := NewLibrary()
	
	// 添加图书
	book1 := library.AddBook("Go语言圣经", "Alan A. A. Donovan", "978-0134190440", 
		"Addison-Wesley", "编程", time.Date(2015, 10, 26, 0, 0, 0, 0), 89.99)
	
	book2 := library.AddBook("深入理解计算机系统", "Randal E. Bryant", "978-7111544937", 
		"机械工业出版社", "计算机", time.Date(2016, 11, 1, 0, 0, 0, 0), 139.00)
	
	fmt.Printf("添加图书:\n")
	fmt.Printf("  %s - %s\n", book1.Title, book1.Author)
	fmt.Printf("  %s - %s\n", book2.Title, book2.Author)
	
	// 借阅图书
	err := library.BorrowBook(book1.ID)
	if err != nil {
		fmt.Printf("借阅失败: %v\n", err)
	} else {
		fmt.Printf("\n成功借阅: %s\n", book1.Title)
	}
	
	// 再次尝试借阅
	err = library.BorrowBook(book1.ID)
	if err != nil {
		fmt.Printf("再次借阅失败: %v\n", err)
	}
	
	// 归还图书
	err = library.ReturnBook(book1.ID)
	if err != nil {
		fmt.Printf("归还失败: %v\n", err)
	} else {
		fmt.Printf("\n成功归还: %s\n", book1.Title)
	}
}
```

### 练习 2：学生信息管理

```go
// 要求：
// 1. 定义 Student 和 Course 结构体
// 2. 实现成绩计算功能
// 3. 支持课程添加和删除
// 4. 实现学生排名功能
```

### 练习 3：简单的ORM实现

```go
// 要求：
// 1. 定义 Model 接口
// 2. 实现基础的 CRUD 操作
// 3. 支持查询过滤
// 4. 实现批量操作
```

## 🤔 思考题

1. **结构体和类的区别是什么？**
   - 提示：考虑继承、方法重写、访问控制

2. **什么时候使用值接收者，什么时候使用指针接收者？**
   - 提示：考虑修改状态、性能、nil 检查

3. **结构体组合 vs 继承有什么优势？**
   - 提示：考虑灵活性、耦合性、组合原则

4. **结构体标签的实际应用场景有哪些？**
   - 提示：考虑序列化、验证、ORM、API

## 📚 扩展阅读

- [Go 语言规范 - 结构体类型](https://golang.org/ref/spec#Struct_types)
- [Effective Go - 结构体](https://golang.org/doc/effective_go.html#structs)
- [Go by Example - 结构体](https://gobyexample.com/structs)

## ⏭️ 下一章节

[指针](./10-pointers.md) → 学习 Go 语言的指针

---

**💡 小贴士**: 结构体是 Go 面向对象编程的基础。记住：**Go 推崇组合优于继承，善用结构体组合能让代码更灵活、更易维护**！
