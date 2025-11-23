# 指针示例

本示例展示了 Go 语言中指针的使用方法和特性。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 指针基本概念

- **地址存储**：指针存储变量的内存地址
- **解引用**：使用 `*` 操作符访问指针指向的值
- **取地址**：使用 `&` 操作符获取变量的地址
- **零值**：未初始化的指针为零值 `nil`

### 2. 指针声明和初始化

```go
// 基本语法
var p *int           // 声明指针变量
x := 42
p = &x              // p 指向 x 的地址

// 直接创建指针
p2 := new(int)       // 分配内存并返回指针
*p2 = 100           // 通过指针设置值

// 取地址
p3 := &someVariable  // 获取变量地址
```

### 3. 指针操作

```go
// 解引用
value := *p          // 获取指针指向的值
*p = 100            // 修改指针指向的值

// 比较指针
if p1 == p2 { ... }  // 比较地址
if p != nil { ... }  // 检查是否为空指针
```

### 4. 指针与函数

```go
// 值传递（不会修改原变量）
func modifyByValue(x int) {
    x = 100  // 只修改副本
}

// 指针传递（会修改原变量）
func modifyByPointer(p *int) {
    *p = 100  // 修改原变量
}

// 返回指针
func createPointer() *int {
    x := new(int)  // 在堆上分配
    *x = 42
    return x
}
```

### 5. 结构体指针

```go
type Person struct {
    Name string
    Age  int
}

// 创建结构体指针
p1 := &Person{Name: "张三", Age: 25}
p2 := new(Person)
p2.Name = "李四"
p2.Age = 30

// 访问字段（自动解引用）
fmt.Println(p1.Name)     // 等价于 (*p1).Name
p1.Age = 26              // 等价于 (*p1).Age = 26
```

### 6. 指针接收者

```go
// 值接收者（不会修改原结构体）
func (p Person) GetName() string {
    return p.Name
}

// 指针接收者（会修改原结构体）
func (p *Person) SetName(name string) {
    p.Name = name
}
```

### 7. 指针与集合

```go
// 指针切片
people := []*Person{
    {Name: "张三", Age: 25},
    {Name: "李四", Age: 30},
}

// 指针映射
userMap := map[string]*Person{
    "user1": {Name: "用户1", Age: 20},
}

// 指针的指针
var ptr **int
```

### 8. 函数指针

```go
// 函数类型
type Operation func(int, int) int

// 函数指针映射
operations := map[string]Operation{
    "add": func(a, b int) int { return a + b },
    "mul": func(a, b int) int { return a * b },
}
```

## 重要概念

### 1. 值类型 vs 引用类型

Go 中的类型分类：
- **值类型**：int、float、bool、string、struct、array
- **引用类型**：slice、map、channel、interface、pointer、function

### 2. 内存分配

```go
// 栈上分配（局部变量）
x := 42

// 堆上分配
p := new(int)        // 显式堆分配
p2 := &struct{}{}    // 编译器可能优化为堆分配
```

### 3. 指针大小

在64位系统上，所有指针类型都占用8字节：
```go
fmt.Println(unsafe.Sizeof(&int{}))     // 8
fmt.Println(unsafe.Sizeof(&string{}))  // 8
fmt.Println(unsafe.Sizeof(&struct{}{})) // 8
```

### 4. 空指针安全

```go
var p *int
if p != nil {
    fmt.Println(*p)  // 安全访问
} else {
    fmt.Println("指针为 nil")
}
```

## 指针使用场景

### 1. 修改函数参数

```go
func modify(x *int) {
    *x = 100  // 修改调用者的变量
}
```

### 2. 避免大对象复制

```go
func processLargeData(data *LargeStruct) {
    // 避免复制大结构体
}
```

### 3. 实现可选参数

```go
func CreateUser(name string, age *int) {
    if age != nil {
        // 使用提供的年龄
    } else {
        // 使用默认年龄
    }
}
```

### 4. 链表等数据结构

```go
type ListNode struct {
    Value int
    Next  *ListNode  // 自引用指针
}
```

### 5. 对象池模式

```go
type BufferPool struct {
    pool [][]byte
}
```

## 指针安全

### 1. 空指针检查

```go
func safeAccess(p *Person) {
    if p != nil {
        fmt.Println(p.Name)
    }
}
```

### 2. 避免悬垂指针

```go
// 错误：返回局部变量地址
func bad() *int {
    x := 42
    return &x  // x 在函数结束后被销毁
}

// 正确：使用 new
func good() *int {
    x := new(int)
    *x = 42
    return x
}
```

### 3. 指针生命周期

确保指针指向的数据在使用时仍然有效：
- 避免返回局部变量的地址
- 注意并发访问时的安全性
- 及时释放不再需要的内存

## 性能考虑

### 1. 指针 vs 值

- **小对象**：值传递可能更快
- **大对象**：指针传递避免复制开销
- **需要修改**：必须使用指针

### 2. 内存分配

- 指针可能导致额外的堆分配
- 减少不必要的指针使用
- 使用对象池重用内存

## 最佳实践

1. **明确使用意图**：只在需要时使用指针
2. **检查空指针**：解引用前检查 nil
3. **避免过度使用**：不要为所有类型都用指针
4. **考虑线程安全**：多 goroutine 访问时注意同步
5. **文档说明**：在函数文档中说明指针参数的用途

## 常见陷阱

1. **空指针解引用**：导致 panic
2. **悬垂指针**：访问已释放的内存
3. **指针比较**：比较的是地址而不是值
4. **过度使用**：不必要的指针增加复杂性
5. **并发竞争**：多 goroutine 同时访问指针数据

## Go 指针 vs C/C++ 指针

| 特性 | Go | C/C++ |
|------|-----|-------|
| 指针运算 | 不支持 | 支持 |
| 空指针 | nil | NULL/nullptr |
| 自动内存管理 | 有 | 无 |
| 指针类型安全 | 强类型 | 可以强制转换 |
| 函数指针 | 安全函数类型 | 不安全函数指针 |

## 练习

1. 实现双向链表
2. 创建二叉搜索树
3. 实现内存缓存系统
4. 创建对象池管理器
5. 实现图数据结构
6. 创建智能指针包装器