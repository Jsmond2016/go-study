# 控制流程示例

本示例展示了 Go 语言中的各种控制流程语句，包括条件语句、循环语句、分支语句和跳转语句。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 条件语句 (if-else)

**基本语法**：
```go
if condition {
    // 条件为真时执行
}

if condition {
    // 条件为真时执行
} else {
    // 条件为假时执行
}

if condition1 {
    // 条件1为真时执行
} else if condition2 {
    // 条件1为假且条件2为真时执行
} else {
    // 所有条件都为假时执行
}
```

**if 初始化语句**：
```go
if x := 10; x > 5 {
    // 可以使用 x
}
// x 在这里不可访问
```

### 2. 循环语句 (for)

Go 只有一种循环结构：`for` 循环，但它有多种形式：

**基本形式**：
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**while 形式**：
```go
for condition {
    // 循环体
}
```

**无限循环**：
```go
for {
    // 无限循环，需要 break 退出
}
```

**for-range 遍历**：
```go
// 遍历数组/切片
for index, value := range array {
    fmt.Printf("索引：%d, 值：%v\n", index, value)
}

// 只要值
for _, value := range array {
    fmt.Printf("值：%v\n", value)
}
```

### 3. 分支语句 (switch)

Go 的 `switch` 非常灵活：

**基本形式**：
```go
switch variable {
case value1:
    // 执行语句
case value2, value3:
    // 多个值
default:
    // 默认情况
}
```

**switch 初始化语句**：
```go
switch x := calculateValue(); x {
case 1:
    fmt.Println("结果是1")
}
```

**无表达式的 switch**：
```go
switch {
case x > 100:
    fmt.Println("大于100")
case x > 50:
    fmt.Println("大于50")
default:
    fmt.Println("小于等于50")
}
```

**fallthrough 关键字**：
```go
switch x {
case 1:
    fmt.Println("情况1")
    fallthrough // 继续执行下一个 case
case 2:
    fmt.Println("情况2") // 总会执行
}
```

### 4. 跳转语句

**break**：
- 跳出当前循环
- 可以指定标签跳出外层循环

**continue**：
- 跳过当前迭代的剩余部分
- 继续下一次迭代
- 可以指定标签

**goto**：
- 无条件跳转到指定标签
- 慎用，通常可以用其他结构替代

**标签**：
```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // 跳出外层循环
        }
    }
}
```

### 5. 类型 switch

用于类型断言和类型判断：

```go
switch v := i.(type) {
case int:
    fmt.Printf("整数：%d\n", v)
case string:
    fmt.Printf("字符串：%s\n", v)
case bool:
    fmt.Printf("布尔值：%t\n", v)
default:
    fmt.Printf("未知类型：%T\n", v)
}
```

## 最佳实践

### 1. 条件表达式

- 避免复杂的嵌套条件
- 使用清晰的变量名
- 考虑将复杂条件提取为函数

### 2. 循环优化

- 尽量减少循环内的计算
- 合理使用 `break` 和 `continue`
- 优先使用 `for-range` 遍历集合

### 3. switch 语句

- 使用 `fallthrough` 要谨慎
- 优先使用无表达式的 `switch` 替代复杂的 `if-else` 链
- 考虑使用 `default` 处理意外情况

### 4. 错误处理

Go 中常见的错误处理模式：
```go
if err != nil {
    // 处理错误
    return
}
```

## 常见陷阱

1. **未使用的变量**：Go 不允许未使用的变量
2. **作用域错误**：注意变量的作用域
3. **死循环**：确保循环有退出条件
4. **类型匹配**：`switch` 中的类型要匹配
5. **标签使用**：标签必须定义在使用之前

## 练习

1. 实现一个简单的计算器，使用 `switch` 处理不同的运算
2. 编写程序打印九九乘法表
3. 实现猜数字游戏
4. 使用 `for-range` 遍历不同类型的数据结构
5. 实现一个简单的菜单系统
6. 编写程序统计字符串中各字符出现的次数