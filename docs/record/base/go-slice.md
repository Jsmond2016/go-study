# golang 中切片的常用方法有哪些？

我来介绍一下 Go 语言中切片的常用操作方法。



## 一、切片的创建和初始化

```go
// 1. 直接创建
slice1 := []int{1, 2, 3, 4, 5}

// 2. make 创建
slice2 := make([]int, 5)    // 长度为5，容量为5
slice3 := make([]int, 3, 5) // 长度为3，容量为5

// 3. 从数组创建
array := [5]int{1, 2, 3, 4, 5}
slice4 := array[1:4]  // [2,3,4]
```

## 二、切片的基本操作

```go
// 1. 追加元素
slice := []int{1, 2, 3}
slice = append(slice, 4)        // 追加单个元素
slice = append(slice, 5, 6, 7)  // 追加多个元素
slice = append(slice, []int{8, 9}...) // 追加另一个切片

// 2. 截取切片
slice := []int{1, 2, 3, 4, 5}
slice1 := slice[1:4]    // [2,3,4]
slice2 := slice[2:]     // [3,4,5]
slice3 := slice[:3]     // [1,2,3]

// 3. 复制切片
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

## 三、切片的长度和容量

```go
slice := make([]int, 3, 5)

// 获取长度和容量
length := len(slice)  // 3
capacity := cap(slice) // 5

// 扩展切片
slice = slice[:4]  // 长度变为4，但不超过容量5
```

## 四、切片的删除操作

```go
// 1. 删除指定位置的元素
slice := []int{1, 2, 3, 4, 5}
i := 2 // 要删除的索引
slice = append(slice[:i], slice[i+1:]...)  // 删除索引i的元素

// 2. 删除指定范围的元素
slice = append(slice[:i], slice[j:]...)    // 删除索引i到j的元素

// 3. 删除并保持顺序
copy(slice[i:], slice[i+1:])
slice = slice[:len(slice)-1]
```

## 五、切片的排序

```go
import "sort"

// 1. 基本类型排序
numbers := []int{4, 2, 1, 3, 5}
sort.Ints(numbers)               // 升序排序
sort.Sort(sort.Reverse(sort.IntSlice(numbers))) // 降序排序

// 2. 自定义类型排序
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 20},
}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

## 六、切片的搜索

```go
// 1. 查找元素
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    if v == 3 {
        fmt.Printf("找到元素 %d 在位置 %d\n", v, i)
        break
    }
}

// 2. 二分查找（要求切片已排序）
numbers := []int{1, 2, 3, 4, 5}
index := sort.SearchInts(numbers, 3)  // 返回值为2
```

## 七、实用技巧

1. **预分配内存**：
```go
// 当知道最终大小时，预分配可以提高性能
slice := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    slice = append(slice, i)
}
```

2. **清空切片**：
```go
// 方法1：重置长度
slice = slice[:0]

// 方法2：重新赋值
slice = nil
```

## 注意事项

1. 切片是引用类型，传递时会传递指针
2. 使用 append 时要注意容量扩展可能创建新数组
3. 在大量追加操作时，预先分配足够的容量可以提高性能
4. 删除切片元素时要注意内存泄漏问题
5. 使用 copy 函数而不是赋值来复制切片
6. 截取操作 slice[i:j] 的底层数组仍然共享，可能造成内存泄漏


这个文档涵盖了 Go 语言中最常用的切片操作方法。建议根据实际场景选择合适的操作方式，特别是在性能要求较高的场景下。