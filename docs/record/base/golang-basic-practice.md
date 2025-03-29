# golang 基础练习题

帮我出一份 golang 基础练习题，涉及 不同的类型转换，切片操作，遍历，map操作，每个小分类举例 5 个以上不同场景的练习题，并给出答案；

我来为您准备一份 Go 语言基础练习题及答案。


## 一、类型转换练习

### 练习 1.1：整数和浮点数转换
问题：将整数 42 转换为 float64，然后再转换回 int，观察是否有精度损失。

```go
func Practice1_1() {
    num := 42
    f := float64(num)
    backToInt := int(f)
    fmt.Printf("原始值：%d, 转换后：%f, 再转换回：%d\n", num, f, backToInt)
}
```

### 练习 1.2：字符串和数字转换
问题：将字符串 "3.14159" 转换为 float64，然后将结果四舍五入到两位小数。

```go
func Practice1_2() {
    str := "3.14159"
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("转换错误:", err)
        return
    }
    rounded := math.Round(f*100) / 100
    fmt.Printf("原始字符串：%s, 转换后：%.2f\n", str, rounded)
}
```

### 练习 1.3：字节切片和字符串转换
问题：将字符串 "Hello, 世界" 转换为字节切片，再转换回字符串。

```go
func Practice1_3() {
    str := "Hello, 世界"
    bytes := []byte(str)
    backToStr := string(bytes)
    fmt.Printf("原始字符串：%s\n字节切片：%v\n转换回的字符串：%s\n", str, bytes, backToStr)
}
```

### 练习 1.4：rune 和字符串转换
问题：将字符串中的每个字符转换为其 Unicode 码点值，然后再转换回字符。

```go
func Practice1_4() {
    str := "Hello, 世界"
    for _, r := range str {
        fmt.Printf("字符：%c, Unicode：%d\n", r, r)
    }
}
```

### 练习 1.5：布尔值和字符串转换
问题：将字符串 "true" 和 "false" 转换为布尔值，然后再转换回字符串。

```go
func Practice1_5() {
    strs := []string{"true", "false", "True", "False"}
    for _, s := range strs {
        b, err := strconv.ParseBool(s)
        if err != nil {
            fmt.Printf("转换错误: %s\n", err)
            continue
        }
        fmt.Printf("字符串：%s -> 布尔值：%v -> 字符串：%s\n", s, b, strconv.FormatBool(b))
    }
}
```

## 二、切片操作练习

### 练习 2.1：切片追加和删除
问题：创建一个整数切片，实现：追加元素、删除指定位置的元素、在指定位置插入元素。

```go
func Practice2_1() {
    slice := []int{1, 2, 3, 4, 5}
    
    // 追加元素
    slice = append(slice, 6)
    
    // 删除索引为 2 的元素
    slice = append(slice[:2], slice[3:]...)
    
    // 在索引 1 处插入元素 10
    slice = append(slice[:1], append([]int{10}, slice[1:]...)...)
    
    fmt.Printf("结果：%v\n", slice)
}
```

### 练习 2.2：切片排序和搜索
问题：创建一个字符串切片，实现排序并查找特定元素的位置。

```go
func Practice2_2() {
    names := []string{"Alice", "Bob", "Charlie", "David"}
    
    // 排序
    sort.Strings(names)
    
    // 二分查找
    target := "Charlie"
    index := sort.SearchStrings(names, target)
    
    fmt.Printf("排序后：%v\n目标 %s 的位置：%d\n", names, target, index)
}
```

### 练习 2.3：切片复制和克隆
问题：实现切片的深度复制，确保修改新切片不影响原切片。

```go
func Practice2_3() {
    original := []int{1, 2, 3, 4, 5}
    
    // 方法1：使用 copy
    clone1 := make([]int, len(original))
    copy(clone1, original)
    
    // 方法2：使用 append
    clone2 := append([]int(nil), original...)
    
    // 修改克隆
    clone1[0] = 10
    clone2[0] = 20
    
    fmt.Printf("原始：%v\n克隆1：%v\n克隆2：%v\n", original, clone1, clone2)
}
```

### 练习 2.4：切片过滤
问题：实现一个函数，过滤出切片中的偶数。

```go
func Practice2_4() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    evenNumbers := make([]int, 0)
    for _, n := range numbers {
        if n%2 == 0 {
            evenNumbers = append(evenNumbers, n)
        }
    }
    
    fmt.Printf("原始数组：%v\n偶数：%v\n", numbers, evenNumbers)
}
```

### 练习 2.5：切片去重
问题：实现一个函数，去除切片中的重复元素。

```go
func Practice2_5() {
    numbers := []int{1, 2, 2, 3, 3, 4, 5, 5}
    
    seen := make(map[int]bool)
    unique := make([]int, 0)
    
    for _, n := range numbers {
        if !seen[n] {
            seen[n] = true
            unique = append(unique, n)
        }
    }
    
    fmt.Printf("原始：%v\n去重后：%v\n", numbers, unique)
}
```

## 三、Map 操作练习

### 练习 3.1：单词计数
问题：统计一个字符串中每个单词出现的次数。

```go
func Practice3_1() {
    text := "the quick brown fox jumps over the lazy dog"
    words := strings.Fields(text)
    
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    
    fmt.Printf("单词计数：%v\n", wordCount)
}
```

### 练习 3.2：Map 合并
问题：实现两个 map 的合并操作。

```go
func Practice3_2() {
    map1 := map[string]int{"a": 1, "b": 2}
    map2 := map[string]int{"b": 3, "c": 4}
    
    result := make(map[string]int)
    
    // 复制 map1
    for k, v := range map1 {
        result[k] = v
    }
    
    // 合并 map2
    for k, v := range map2 {
        result[k] = v
    }
    
    fmt.Printf("合并结果：%v\n", result)
}
```

### 练习 3.3：Map 过滤
问题：过滤出 map 中值大于 5 的键值对。

```go
func Practice3_3() {
    scores := map[string]int{
        "Alice": 8,
        "Bob":   4,
        "Charlie": 6,
        "David": 3,
    }
    
    filtered := make(map[string]int)
    for k, v := range scores {
        if v > 5 {
            filtered[k] = v
        }
    }
    
    fmt.Printf("原始：%v\n过滤后：%v\n", scores, filtered)
}
```

### 练习 3.4：Map 键值对调
问题：实现 map 的键值对调（假设值是唯一的）。

```go
func Practice3_4() {
    original := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    inverted := make(map[int]string)
    for k, v := range original {
        inverted[v] = k
    }
    
    fmt.Printf("原始：%v\n对调后：%v\n", original, inverted)
}
```

### 练习 3.5：嵌套 Map
问题：创建一个学生成绩管理系统，使用嵌套 map 存储学生各科成绩。

```go
func Practice3_5() {
    grades := map[string]map[string]int{
        "Alice": {
            "Math":    90,
            "English": 85,
            "Science": 88,
        },
        "Bob": {
            "Math":    78,
            "English": 92,
            "Science": 85,
        },
    }
    
    // 添加新学生
    grades["Charlie"] = map[string]int{
        "Math":    95,
        "English": 89,
        "Science": 92,
    }
    
    // 修改成绩
    grades["Alice"]["Math"] = 92
    
    // 打印所有学生的数学成绩
    for student, scores := range grades {
        fmt.Printf("%s 的数学成绩：%d\n", student, scores["Math"])
    }
}
```

## 四、遍历练习

### 练习 4.1：多维切片遍历
问题：遍历一个二维整数切片，计算每行的和。

```go
func Practice4_1() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    for i, row := range matrix {
        sum := 0
        for _, val := range row {
            sum += val
        }
        fmt.Printf("第 %d 行的和：%d\n", i+1, sum)
    }
}
```

### 练习 4.2：字符串遍历
问题：分别使用 byte 和 rune 遍历字符串，观察区别。

```go
func Practice4_2() {
    str := "Hello, 世界"
    
    // 按字节遍历
    for i := 0; i < len(str); i++ {
        fmt.Printf("%x ", str[i])
    }
    fmt.Println()
    
    // 按字符遍历
    for _, r := range str {
        fmt.Printf("%c ", r)
    }
    fmt.Println()
}
```

### 练习 4.3：Map 按键排序遍历
问题：实现按键排序遍历 map。

```go
func Practice4_3() {
    m := map[string]int{
        "banana": 3,
        "apple":  1,
        "cherry": 2,
    }
    
    // 获取所有键并排序
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    // 按排序后的键遍历
    for _, k := range keys {
        fmt.Printf("%s: %d\n", k, m[k])
    }
}
```

### 练习 4.4：通道遍历
问题：创建一个通道，发送一些数据并遍历接收。

```go
func Practice4_4() {
    ch := make(chan int, 5)
    
    // 发送数据
    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
        }
        close(ch)
    }()
    
    // 遍历接收
    for num := range ch {
        fmt.Printf("接收到：%d\n", num)
    }
}
```

### 练习 4.5：复合结构遍历
问题：定义一个结构体切片，实现多种遍历方式。

```go
type Person struct {
    Name string
    Age  int
}

func Practice4_5() {
    people := []Person{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 35},
    }
    
    // 普通遍历
    for i, p := range people {
        fmt.Printf("%d: %s is %d years old\n", i, p.Name, p.Age)
    }
    
    // 按条件遍历
    for _, p := range people {
        if p.Age > 30 {
            fmt.Printf("%s is over 30\n", p.Name)
        }
    }
}
```

这些练习题涵盖了 Go 语言的基础操作，建议逐个实践，加深对各种操作的理解。每个练习都可以扩展或修改来创建更复杂的场景。