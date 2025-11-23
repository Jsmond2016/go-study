# 结构体示例

本示例展示了 Go 语言中结构体的使用方法和特性。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 结构体基本特性

- **复合类型**：组合多个不同类型的字段
- **值类型**：结构体是值类型，赋值会复制整个结构
- **字段访问**：使用点号 `.` 访问字段
- **可比较性**：所有字段可比较的结构体可以比较

### 2. 结构体定义

```go
type Person struct {
    Name    string  // 公开字段（首字母大写）
    age     int     // 私有字段（首字母小写）
    Email   string
    IsAdmin bool
}
```

### 3. 结构体初始化

```go
// 零值初始化
var p1 Person

// 空值初始化
p2 := Person{}

// 字段名初始化（推荐）
p3 := Person{
    Name:  "张三",
    Age:   25,
    Email: "zhangsan@example.com",
}

// 按顺序初始化
p4 := Person{"张三", 25, "zhangsan@example.com", false}

// 部分初始化（其余为零值）
p5 := Person{
    Name: "张三",
    Age:  25,
    // Email 和 IsAdmin 为零值
}
```

### 4. 结构体指针

```go
// 创建指针
p := &Person{Name: "张三", Age: 25}

// 访问字段（自动解引用）
fmt.Println(p.Name)  // 等价于 (*p).Name

// 修改字段
p.Age = 26  // 等价于 (*p).Age = 26
```

### 5. 结构体方法

```go
// 值接收者（不修改原结构体）
func (p Person) GetName() string {
    return p.Name
}

// 指针接收者（可以修改原结构体）
func (p *Person) SetName(newName string) {
    p.Name = newName
}

// 方法调用
person := Person{Name: "张三"}
name := person.GetName()  // 值接收者
person.SetName("李四")    // 指针接收者
```

### 6. 结构体组合

```go
type Address struct {
    City    string
    Street  string
}

type Person struct {
    Name    string
    Address Address  // 嵌套结构体
}

// 访问嵌套字段
person.Address.City

// 如果外层结构体没有同名字段，可以提升访问
// （需要使用匿名结构体）
type Person struct {
    Name string
    Address  // 匿名结构体
}
person.City  // 直接访问
```

### 7. 结构体标签

```go
type User struct {
    ID    int    `json:"id" db:"user_id" bson:"_id"`
    Name  string `json:"name" db:"name" bson:"name"`
    Email string `json:"email,omitempty" db:"email"` // omitempty: 空值时忽略
    Secret string `json:"-" db:"secret"`           // -: 忽略序列化
}
```

### 8. 匿名结构体

```go
// 定义匿名结构体变量
user := struct {
    ID   int
    Name string
}{
    ID:   1,
    Name: "张三",
}

// 作为字段类型
type Config struct {
    Server struct {
        Host string
        Port int
    }
}
```

## 重要概念

### 1. 值类型 vs 指针类型

```go
// 值传递（会复制整个结构体）
func modifyPerson(p Person) {
    p.Name = "新名字"  // 不会影响原结构体
}

// 指针传递（会修改原结构体）
func modifyPersonPtr(p *Person) {
    p.Name = "新名字"  // 会影响原结构体
}
```

### 2. 构造函数模式

```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
```

### 3. 方法接收者选择

- **值接收者**：不修改结构体，结构体较小
- **指针接收者**：需要修改结构体，结构体较大
- 一致性：同一类型的方法最好使用一致的接收者类型

### 4. 接口实现

结构体通过方法实现接口：

```go
type Stringer interface {
    String() string
}

func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}
```

## 实际应用场景

### 1. 数据模型

```go
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}
```

### 2. 配置结构

```go
type Config struct {
    Database DatabaseConfig `json:"database"`
    Server   ServerConfig   `json:"server"`
    Debug    bool           `json:"debug"`
}
```

### 3. API 响应

```go
type APIResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}
```

## 最佳实践

1. **字段命名**：使用驼峰命名法，公开字段首字母大写
2. **初始化**：优先使用字段名初始化，提高可读性
3. **指针使用**：大结构体使用指针，避免复制开销
4. **方法接收者**：需要修改状态时使用指针接收者
5. **标签使用**：合理使用结构体标签，特别是JSON/数据库映射
6. **构造函数**：为复杂结构体提供构造函数

## 常见陷阱

1. **指针误用**：修改值接收者方法内的字段不会影响原结构体
2. **零值比较**：包含不可比较字段的结构体不能直接比较
3. **匿名结构体限制**：匿名结构体不能声明方法
4. **递归类型**：结构体不能包含自身（但可以包含指向自身的指针）
5. **内存布局**：结构体内存对齐可能影响大小

## 性能考虑

1. **结构体大小**：合理安排字段顺序减少内存占用
2. **指针 vs 值**：大结构体使用指针避免复制
3. **方法调用**：值接收者方法会复制整个结构体
4. **内存对齐**：相同类型字段放在一起可能减少填充

## 练习

1. 实现一个完整的图书管理系统
2. 创建一个配置文件解析器
3. 实现一个简单的游戏角色系统
4. 创建一个响应式Web API的数据结构
5. 实现一个树形数据结构
6. 创建一个链表或栈数据结构