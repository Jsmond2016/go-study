# 切片示例

本示例展示了 Go 语言中切片的使用方法和特性。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 切片基本特性

- **动态大小**：切片长度可以动态改变
- **引用类型**：切片是引用类型，共享底层数组
- **三要素**：指针、长度、容量
- **灵活操作**：支持切片、追加、复制等操作

### 2. 切片创建方式

```go
// nil 切片
var slice1 []int

// 直接初始化
slice2 := []int{1, 2, 3, 4, 5}

// 使用 make 创建
slice3 := make([]int, 5)      // 长度5，容量5
slice4 := make([]int, 3, 10)  // 长度3，容量10

// 从数组/切片创建
arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
slice5 := arr[2:7]  // 包含索引2到6的元素
```

### 3. 切片底层结构

```go
type slice struct {
    ptr unsafe.Pointer  // 指向底层数组的指针
    len int             // 切片长度
    cap int             // 切片容量
}
```

### 4. 切片操作

```go
// 添加元素
slice = append(slice, element)
slice = append(slice, elements...)
slice = append(slice, anotherSlice...)

// 删除元素（索引为 i）
slice = append(slice[:i], slice[i+1:]...)

// 复制切片
copy1 := make([]int, len(original))
copy(copy1, original)

// 切片的切片
newSlice := slice[low:high]  // 包含 low，不包含 high
```

### 5. 切片扩容机制

- 当切片容量不足时，`append` 会自动扩容
- 扩容策略：容量小于1024时翻倍，大于1024时增加1/4
- 扩容可能会分配新的底层数组

### 6. 常用切片操作

```go
// 检查元素是否存在
func contains(slice []int, item int) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// 过滤切片
func filter(slice []int, predicate func(int) bool) []int {
    result := make([]int, 0)
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// 反转切片
func reverse(slice []int) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}
```

## 重要概念

### 1. 切片 vs 数组

| 特性 | 数组 | 切片 |
|------|------|------|
| 大小 | 固定 | 动态 |
| 传递方式 | 值传递 | 引用传递 |
| 性能 | 访问快 | 灵活性高 |
| 内存布局 | 连续 | 灵活 |
| 使用场景 | 固定大小集合 | 动态数据集合 |

### 2. 切片共享底层数组

多个切片可以共享同一个底层数组：
```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1[1:4]  // 共享底层数组
slice2[0] = 99        // 会影响 slice1
```

### 3. 容量管理

```go
// 预分配容量提高性能
slice := make([]int, 0, 1000)  // 预分配1000容量
// 在循环中 append 时避免频繁扩容
```

### 4. 内存泄漏注意

切片持有底层数组的引用，要注意避免内存泄漏：
```go
// 大数组的切片
largeArray := [1000000]int{...}
smallSlice := largeArray[100:200]  // 会持有整个数组的引用

// 解决方案：复制到新切片
safeSlice := make([]int, 100)
copy(safeSlice, largeArray[100:200])
```

## 标准库 slices 包（Go 1.21+）

```go
import "slices"

// 排序
slices.Sort(slice)                    // 升序
slices.SortFunc(slice, func(a, b int) int {
    return b - a                      // 降序
})

// 查找
index := slices.Index(slice, element)           // 查找元素
index := slices.IndexFunc(slice, predicate)     // 条件查找

// 比较
equal := slices.Equal(slice1, slice2)           // 比较相等
equalFunc := slices.EqualFunc(slice1, slice2, cmpFunc)

// 其他操作
max := slices.Max(slice)                        // 最大值
min := slices.Min(slice)                        // 最小值
contains := slices.Contains(slice, element)     // 是否包含
```

## 性能优化建议

1. **预分配容量**：在知道大概大小时预分配
2. **避免频繁扩容**：批量操作时预留足够容量
3. **及时释放大切片**：设置为 nil 释放内存
4. **使用复制替代共享**：避免意外的大数组引用

## 最佳实践

1. **优先使用切片**：除非确实需要固定大小
2. **理解共享机制**：注意切片间的数据共享
3. **合理预分配**：根据使用场景预分配容量
4. **使用标准库**：Go 1.21+ 使用 slices 包

## 常见陷阱

1. **意外的数据修改**：切片间共享数据
2. **内存泄漏**：大数组的切片引用
3. **扩容性能**：频繁 append 导致的性能问题
4. **nil 切片 vs 空切片**：注意区别

## 练习

1. 实现一个高效的切片去重函数
2. 创建一个可动态增长的栈数据结构
3. 实现切片的分页功能
4. 创建一个缓冲池管理系统
5. 使用切片实现一个简单的内存数据库