---
title: 基础练习题
difficulty: beginner
duration: "6-8小时"
prerequisites: ["变量与常量", "数据类型", "控制流程", "函数", "切片", "映射"]
tags: ["练习", "实践", "类型转换", "切片", "映射", "遍历"]
---

# Go 语言基础练习题

通过实践练习巩固 Go 语言基础知识，涵盖类型转换、切片操作、遍历、映射等核心概念。

## 📋 学习目标

- [ ] 通过练习掌握类型转换
- [ ] 熟练使用切片的各种操作
- [ ] 掌握映射的常用操作
- [ ] 理解不同遍历方式的使用场景
- [ ] 能够解决实际问题

## 🔄 类型转换练习

### 练习 1.1：整数和浮点数转换

**问题**：将整数 42 转换为 float64，然后再转换回 int，观察是否有精度损失。

```go
package main

import "fmt"

func Practice1_1() {
	num := 42
	f := float64(num)
	backToInt := int(f)
	fmt.Printf("原始值：%d, 转换后：%f, 再转换回：%d\n", num, f, backToInt)
}

func main() {
	Practice1_1()
}
```

### 练习 1.2：字符串和数字转换

**问题**：将字符串 "3.14159" 转换为 float64，然后将结果四舍五入到两位小数。

```go
package main

import (
	"fmt"
	"math"
	"strconv"
)

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

func main() {
	Practice1_2()
}
```

### 练习 1.3：字节切片和字符串转换

**问题**：将字符串 "Hello, 世界" 转换为字节切片，再转换回字符串。

```go
package main

import "fmt"

func Practice1_3() {
	str := "Hello, 世界"
	bytes := []byte(str)
	backToStr := string(bytes)
	fmt.Printf("原始字符串：%s\n字节切片：%v\n转换回的字符串：%s\n", str, bytes, backToStr)
}

func main() {
	Practice1_3()
}
```

### 练习 1.4：rune 和字符串转换

**问题**：将字符串中的每个字符转换为其 Unicode 码点值，然后再转换回字符。

```go
package main

import "fmt"

func Practice1_4() {
	str := "Hello, 世界"
	for _, r := range str {
		fmt.Printf("字符：%c, Unicode：%d\n", r, r)
	}
}

func main() {
	Practice1_4()
}
```

### 练习 1.5：布尔值和字符串转换

**问题**：将字符串 "true" 和 "false" 转换为布尔值，然后再转换回字符串。

```go
package main

import (
	"fmt"
	"strconv"
)

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

func main() {
	Practice1_5()
}
```

## 🔪 切片操作练习

### 练习 2.1：切片追加和删除

**问题**：创建一个整数切片，实现：追加元素、删除指定位置的元素、在指定位置插入元素。

```go
package main

import "fmt"

func Practice2_1() {
	slice := []int{1, 2, 3, 4, 5}

	// 追加元素
	slice = append(slice, 6)
	fmt.Printf("追加后：%v\n", slice)

	// 删除索引为 2 的元素
	slice = append(slice[:2], slice[3:]...)
	fmt.Printf("删除索引2后：%v\n", slice)

	// 在索引 1 处插入元素 10
	slice = append(slice[:1], append([]int{10}, slice[1:]...)...)
	fmt.Printf("插入后：%v\n", slice)
}

func main() {
	Practice2_1()
}
```

### 练习 2.2：切片排序和搜索

**问题**：创建一个字符串切片，实现排序并查找特定元素的位置。

```go
package main

import (
	"fmt"
	"sort"
)

func Practice2_2() {
	names := []string{"Alice", "Bob", "Charlie", "David"}

	// 排序
	sort.Strings(names)
	fmt.Printf("排序后：%v\n", names)

	// 二分查找
	target := "Charlie"
	index := sort.SearchStrings(names, target)
	if index < len(names) && names[index] == target {
		fmt.Printf("目标 %s 的位置：%d\n", target, index)
	} else {
		fmt.Printf("未找到 %s\n", target)
	}
}

func main() {
	Practice2_2()
}
```

### 练习 2.3：切片复制和克隆

**问题**：实现切片的深度复制，确保修改新切片不影响原切片。

```go
package main

import "fmt"

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

func main() {
	Practice2_3()
}
```

### 练习 2.4：切片过滤

**问题**：实现一个函数，过滤出切片中的偶数。

```go
package main

import "fmt"

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

func main() {
	Practice2_4()
}
```

### 练习 2.5：切片去重

**问题**：实现一个函数，去除切片中的重复元素。

```go
package main

import "fmt"

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

func main() {
	Practice2_5()
}
```

## 🗺️ Map 操作练习

### 练习 3.1：单词计数

**问题**：统计一个字符串中每个单词出现的次数。

```go
package main

import (
	"fmt"
	"strings"
)

func Practice3_1() {
	text := "the quick brown fox jumps over the lazy dog"
	words := strings.Fields(text)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Printf("单词计数：%v\n", wordCount)
}

func main() {
	Practice3_1()
}
```

### 练习 3.2：Map 合并

**问题**：实现两个 map 的合并操作。

```go
package main

import "fmt"

func Practice3_2() {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	result := make(map[string]int)

	// 复制 map1
	for k, v := range map1 {
		result[k] = v
	}

	// 合并 map2（覆盖相同键）
	for k, v := range map2 {
		result[k] = v
	}

	fmt.Printf("合并结果：%v\n", result)
}

func main() {
	Practice3_2()
}
```

### 练习 3.3：Map 过滤

**问题**：过滤出 map 中值大于 5 的键值对。

```go
package main

import "fmt"

func Practice3_3() {
	scores := map[string]int{
		"Alice":   8,
		"Bob":     4,
		"Charlie": 6,
		"David":   3,
	}

	filtered := make(map[string]int)
	for k, v := range scores {
		if v > 5 {
			filtered[k] = v
		}
	}

	fmt.Printf("原始：%v\n过滤后：%v\n", scores, filtered)
}

func main() {
	Practice3_3()
}
```

### 练习 3.4：Map 键值对调

**问题**：实现 map 的键值对调（假设值是唯一的）。

```go
package main

import "fmt"

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

func main() {
	Practice3_4()
}
```

### 练习 3.5：嵌套 Map

**问题**：创建一个学生成绩管理系统，使用嵌套 map 存储学生各科成绩。

```go
package main

import "fmt"

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

func main() {
	Practice3_5()
}
```

## 🔁 遍历练习

### 练习 4.1：多维切片遍历

**问题**：遍历一个二维整数切片，计算每行的和。

```go
package main

import "fmt"

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

func main() {
	Practice4_1()
}
```

### 练习 4.2：字符串遍历

**问题**：分别使用 byte 和 rune 遍历字符串，观察区别。

```go
package main

import "fmt"

func Practice4_2() {
	str := "Hello, 世界"

	// 按字节遍历
	fmt.Println("按字节遍历:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()

	// 按字符遍历
	fmt.Println("按字符遍历:")
	for _, r := range str {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}

func main() {
	Practice4_2()
}
```

### 练习 4.3：Map 按键排序遍历

**问题**：实现按键排序遍历 map。

```go
package main

import (
	"fmt"
	"sort"
)

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

func main() {
	Practice4_3()
}
```

### 练习 4.4：通道遍历

**问题**：创建一个通道，发送一些数据并遍历接收。

```go
package main

import "fmt"

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

func main() {
	Practice4_4()
}
```

### 练习 4.5：复合结构遍历

**问题**：定义一个结构体切片，实现多种遍历方式。

```go
package main

import "fmt"

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
	fmt.Println("所有人员:")
	for i, p := range people {
		fmt.Printf("%d: %s is %d years old\n", i, p.Name, p.Age)
	}

	// 按条件遍历
	fmt.Println("\n超过30岁的人员:")
	for _, p := range people {
		if p.Age > 30 {
			fmt.Printf("%s is over 30\n", p.Name)
		}
	}
}

func main() {
	Practice4_5()
}
```

## 🎯 综合练习

### 练习 5.1：学生管理系统

**问题**：实现一个简单的学生管理系统，包括添加、查询、删除学生信息。

```go
package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

type StudentManager struct {
	students map[int]*Student
	nextID   int
}

func NewStudentManager() *StudentManager {
	return &StudentManager{
		students: make(map[int]*Student),
		nextID:   1,
	}
}

func (sm *StudentManager) AddStudent(name string, age int, grade string) int {
	id := sm.nextID
	sm.students[id] = &Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	sm.nextID++
	return id
}

func (sm *StudentManager) GetStudent(id int) (*Student, bool) {
	student, exists := sm.students[id]
	return student, exists
}

func (sm *StudentManager) DeleteStudent(id int) bool {
	if _, exists := sm.students[id]; exists {
		delete(sm.students, id)
		return true
	}
	return false
}

func (sm *StudentManager) ListStudents() {
	fmt.Println("学生列表:")
	for _, student := range sm.students {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n",
			student.ID, student.Name, student.Age, student.Grade)
	}
}

func main() {
	sm := NewStudentManager()

	// 添加学生
	id1 := sm.AddStudent("张三", 20, "大二")
	id2 := sm.AddStudent("李四", 21, "大三")

	// 查询学生
	if student, ok := sm.GetStudent(id1); ok {
		fmt.Printf("查询到学生: %s\n", student.Name)
	}

	// 列出所有学生
	sm.ListStudents()

	// 删除学生
	sm.DeleteStudent(id2)
	fmt.Println("\n删除后:")
	sm.ListStudents()
}
```

## 📚 练习建议

1. **循序渐进**：从简单的练习开始，逐步增加难度
2. **理解原理**：不仅要写出代码，还要理解为什么这样写
3. **尝试变体**：完成基础练习后，尝试修改需求或增加功能
4. **代码审查**：对比参考答案，学习更好的实现方式
5. **实际应用**：将学到的知识应用到实际项目中

## 🤔 思考题

1. 类型转换时什么时候会有精度损失？
2. 切片和数组在性能上有什么区别？
3. Map 的遍历顺序为什么是随机的？
4. 什么时候应该使用 range 遍历，什么时候使用索引遍历？
5. 如何优化大量字符串拼接的性能？

## ⏭️ 下一阶段

完成这些练习后，你可以：

- [测试](./16-testing.md) - 学习如何编写测试
- [并发编程](./14-concurrency.md) - 学习 Go 的并发特性
- [项目实战](../projects/) - 开始实际项目开发

---

**💡 提示**: 实践是学习编程的最佳方式。完成这些练习后，你将更加熟练地使用 Go 语言的各种特性！

