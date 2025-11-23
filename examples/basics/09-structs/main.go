// Package main 展示 Go 语言的结构体
package main

import (
	"fmt"
)

// 程序入口，演示结构体操作
func main() {
	fmt.Println("=== Go 结构体示例 ===")

	// 基本结构体
	basicStructs()

	// 结构体方法
	structMethods()

	// 结构体组合
	structComposition()

	// 结构体标签
	structTags()

	// 匿名结构体
	anonymousStructs()

	// 实际应用示例
	practicalExample()
}

// 定义基本结构体
type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdmin bool
}

type Address struct {
	Province string
	City     string
	Street   string
	Number   int
}

// 嵌套结构体
type User struct {
	Person  Person
	Address Address
	ID      int
	Phone   string
}

// 带指针的结构体
type Company struct {
	Name     string
	Employees []*Employee
	CEO      *Person
}

type Employee struct {
	Name    string
	Salary  float64
	Manager *Employee
}

// 基本结构体操作
func basicStructs() {
	fmt.Println("\n--- 基本结构体 ---")

	// 1. 结构体初始化
	var p1 Person                    // 零值初始化
	p2 := Person{}                   // 空值初始化
	p3 := Person{                   // 字段名初始化
		Name:    "张三",
		Age:     25,
		Email:   "zhangsan@example.com",
		IsAdmin: false,
	}
	p4 := Person{"李四", 30, "lisi@example.com", true} // 按顺序初始化

	fmt.Printf("零值初始化: %+v\n", p1)
	fmt.Printf("空值初始化: %+v\n", p2)
	fmt.Printf("字段名初始化: %+v\n", p3)
	fmt.Printf("按顺序初始化: %+v\n", p4)

	// 2. 访问和修改字段
	fmt.Printf("\n字段访问:")
	fmt.Printf("姓名: %s, 年龄: %d\n", p3.Name, p3.Age)

	p3.Age = 26
	fmt.Printf("修改后年龄: %d\n", p3.Age)

	// 3. 结构体指针
	fmt.Printf("\n结构体指针:")
	p5 := &Person{
		Name:  "王五",
		Age:   35,
		Email: "wangwu@example.com",
	}

	fmt.Printf("通过指针访问: %s\n", p5.Name)
	fmt.Printf("解引用访问: %s\n", (*p5).Name)

	// 自动解引用
	fmt.Printf("自动解引用: %s\n", p5.GetName())

	// 4. 比较结构体
	fmt.Printf("\n结构体比较:")
	p6 := Person{"张三", 25, "zhangsan@example.com", false}
	p7 := Person{"张三", 25, "zhangsan@example.com", false}
	p8 := Person{"张三", 26, "zhangsan@example.com", false}

	fmt.Printf("p6 == p7: %t\n", p6 == p7) // true
	fmt.Printf("p6 == p8: %t\n", p6 == p8) // false (年龄不同)

	// 包含不可比较字段的不能用 == 比较
}

// 结构体方法
func structMethods() {
	fmt.Println("\n--- 结构体方法 ---")

	user := User{
		Person: Person{
			Name:    "赵六",
			Age:     28,
			Email:   "zhaoliu@example.com",
			IsAdmin: true,
		},
		Address: Address{
			Province: "北京",
			City:     "北京",
			Street:   "长安街",
			Number:   1,
		},
		ID:    1001,
		Phone: "13800138000",
	}

	// 调用值接收者方法
	fmt.Printf("用户信息: %s\n", user.GetUserInfo())
	fmt.Printf("完整地址: %s\n", user.GetFullAddress())
	fmt.Printf("是否成年: %t\n", user.IsAdult())

	// 调用指针接收者方法
	user.SetEmail("newemail@example.com")
	fmt.Printf("更新后邮箱: %s\n", user.Email)

	user.Birthday()
	fmt.Printf("生日后年龄: %d\n", user.Age)
}

// GetName 值接收者方法
func (p Person) GetName() string {
	return p.Name
}

// GetUserInfo 值接收者方法
func (u User) GetUserInfo() string {
	return fmt.Sprintf("%s (ID: %d, 年龄: %d)", u.Name, u.ID, u.Age)
}

// GetFullAddress 值接收者方法
func (u User) GetFullAddress() string {
	return fmt.Sprintf("%s省%s市%s%s号", u.Address.Province, u.Address.City, u.Address.Street, u.Address.Number)
}

// IsAdult 值接收者方法
func (u User) IsAdult() bool {
	return u.Age >= 18
}

// SetEmail 指针接收者方法（修改原值）
func (u *User) SetEmail(newEmail string) {
	u.Email = newEmail
	fmt.Printf("邮箱已更新为: %s\n", newEmail)
}

// Birthday 指针接收者方法（修改原值）
func (u *User) Birthday() {
	u.Age++
	fmt.Printf("生日快乐！现在 %d 岁了\n", u.Age)
}

// 结构体组合
func structComposition() {
	fmt.Println("\n--- 结构体组合 ---")

	// 创建公司
	ceo := &Person{
		Name:  "马云",
		Age:   55,
		Email: "jack@alibaba.com",
	}

	company := Company{
		Name:     "阿里巴巴",
		CEO:      ceo,
		Employees: make([]*Employee, 0),
	}

	// 添加员工
	emp1 := &Employee{
		Name:    "张三",
		Salary:  15000.0,
		Manager: nil,
	}

	emp2 := &Employee{
		Name:    "李四",
		Salary:  12000.0,
		Manager: emp1,
	}

	company.Employees = append(company.Employees, emp1, emp2)

	fmt.Printf("公司: %s\n", company.Name)
	fmt.Printf("CEO: %s (%d岁)\n", company.CEO.Name, company.CEO.Age)

	fmt.Printf("\n员工列表:\n")
	for i, emp := range company.Employees {
		manager := "无"
		if emp.Manager != nil {
			manager = emp.Manager.Name
		}
		fmt.Printf("%d. %s - 薪水: %.1f, 上级: %s\n", i+1, emp.Name, emp.Salary, manager)
	}
}

// 结构体标签
func structTags() {
	fmt.Println("\n--- 结构体标签 ---")

	type Product struct {
		ID          int     `json:"id" db:"product_id" bson:"_id"`
		Name        string  `json:"name" db:"name" bson:"name"`
		Price       float64 `json:"price" db:"price" bson:"price"`
		Description string  `json:"description,omitempty" db:"desc" bson:"desc,omitempty"`
		InStock     bool    `json:"in_stock" db:"in_stock" bson:"in_stock"`
		InternalID string  `json:"-" db:"internal_id"` // 忽略JSON序列化
	}

	product := Product{
		ID:          1,
		Name:        "苹果手机",
		Price:       6999.99,
		Description: "最新款iPhone",
		InStock:     true,
		InternalID:  "PROD-001-INTERNAL",
	}

	fmt.Printf("产品: %+v\n", product)

	// 在实际应用中，可以使用反射来读取标签
	// 这里只是演示结构体的定义
}

// 匿名结构体
func anonymousStructs() {
	fmt.Println("\n--- 匿名结构体 ---")

	// 1. 基本匿名结构体
	user := struct {
		ID   int
		Name string
		Age  int
	}{
		ID:   1,
		Name: "匿名用户",
		Age:  30,
	}

	fmt.Printf("匿名结构体用户: %+v\n", user)

	// 2. 作为字段类型
	type Config struct {
		Server struct {
			Host string
			Port int
		}
		Database struct {
			Host string
			Port int
			Name string
		}
		Debug bool
	}

	config := Config{
		Server: struct {
			Host string
			Port int
		}{
			Host: "localhost",
			Port: 8080,
		},
		Database: struct {
			Host string
			Port int
			Name string
		}{
			Host: "localhost",
			Port: 5432,
			Name: "myapp",
		},
		Debug: true,
	}

	fmt.Printf("配置: %+v\n", config)

	// 3. 切片中的匿名结构体
	users := []struct {
		Name string
		Age  int
	}{
		{"用户1", 25},
		{"用户2", 30},
		{"用户3", 28},
	}

	fmt.Printf("\n用户列表:\n")
	for i, user := range users {
		fmt.Printf("%d. %s (%d岁)\n", i+1, user.Name, user.Age)
	}
}

// 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1：图书管理系统
	type Book struct {
		ISBN     string  `json:"isbn"`
		Title    string  `json:"title"`
		Author   string  `json:"author"`
		Price    float64 `json:"price"`
		Stock    int     `json:"stock"`
		Category string  `json:"category"`
	}

	type Library struct {
		Name  string
		Books map[string]*Book  // ISBN -> Book
	}

	library := Library{
		Name:  "市图书馆",
		Books: make(map[string]*Book),
	}

	// 添加图书
	books := []*Book{
		{ISBN: "978-7-111-21382-6", Title: "Go程序设计语言", Author: "Alan A.A. Donovan", Price: 89.00, Stock: 10, Category: "编程"},
		{ISBN: "978-7-115-42787-0", Title: "Go并发编程实战", Author: "聂品", Price: 59.00, Stock: 5, Category: "编程"},
		{ISBN: "978-7-121-32749-6", Title: "Go语言实战", Author: "William Kennedy", Price: 79.00, Stock: 8, Category: "编程"},
	}

	for _, book := range books {
		library.Books[book.ISBN] = book
	}

	fmt.Printf("%s - 共有 %d 本图书\n", library.Name, len(library.Books))

	// 查找图书
	if book, found := library.Books["978-7-111-21382-6"]; found {
		fmt.Printf("找到图书: %s - %s - ¥%.2f (库存: %d)\n",
			book.Title, book.Author, book.Price, book.Stock)
	}

	// 借书
	if book, found := library.Books["978-7-111-21382-6"]; found && book.Stock > 0 {
		book.Stock--
		fmt.Printf("借出《%s》后，剩余库存: %d\n", book.Title, book.Stock)
	}

	// 示例2：学生成绩管理
	type Student struct {
		ID      int
		Name    string
		Class   string
		Grades  map[string]float64  // 科目 -> 分数
		Contact struct {
			Phone string
			Email string
		}
	}

	type Class struct {
		Name     string
		Teacher  string
		Students []*Student
	}

	mathClass := Class{
		Name:    "高三(1)班",
		Teacher: "张老师",
		Students: []*Student{
			{
				ID:    1,
				Name:  "张小明",
				Class: "高三(1)班",
				Grades: map[string]float64{
					"数学": 95.5,
					"物理": 88.0,
					"化学": 92.5,
				},
				Contact: struct {
					Phone string
					Email string
				}{
					Phone: "13800138001",
					Email: "zhang@email.com",
				},
			},
			{
				ID:    2,
				Name:  "李小红",
				Class: "高三(1)班",
				Grades: map[string]float64{
					"数学": 89.0,
					"物理": 95.5,
					"化学": 87.0,
				},
				Contact: struct {
					Phone string
					Email string
				}{
					Phone: "13800138002",
					Email: "li@email.com",
				},
			},
		},
	}

	fmt.Printf("\n%s - 老师: %s\n", mathClass.Name, mathClass.Teacher)
	fmt.Printf("学生成绩:\n")

	for _, student := range mathClass.Students {
		var total float64
		var count int

		for subject, score := range student.Grades {
			total += score
			count++
			fmt.Printf("  %s - %s: %.1f\n", student.Name, subject, score)
		}

		avg := total / float64(count)
		fmt.Printf("    平均分: %.1f, 联系方式: %s, %s\n\n", avg, student.Contact.Phone, student.Contact.Email)
	}

	// 示例3：响应结构
	type APIResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	type PageInfo struct {
		Page      int `json:"page"`
		PageSize  int `json:"page_size"`
		Total     int `json:"total"`
		TotalPage int `json:"total_page"`
	}

	successResponse := APIResponse{
		Code:    200,
		Message: "success",
		Data: map[string]interface{}{
			"users": books,
			"page": PageInfo{
				Page:      1,
				PageSize:  10,
				Total:     3,
				TotalPage: 1,
			},
		},
	}

	fmt.Printf("API响应: %+v\n", successResponse)
}