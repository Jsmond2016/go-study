---
title: 映射
difficulty: beginner
duration: "3-4小时"
prerequisites: ["变量与常量", "数据类型", "切片"]
tags: ["映射", "字典", "哈希表", "键值对"]
---

# 映射

映射（Map）是 Go 语言中最重要的数据结构之一，用于存储键值对的集合。映射提供了快速的查找、插入和删除操作。

## 📋 学习目标

- [ ] 掌握映射的声明和初始化
- [ ] 学会映射的基本操作
- [ ] 理解映射的内部实现
- [ ] 掌握映射的遍历方法
- [ ] 学会映射的高级用法
- [ ] 了解映射的性能特性

## 🏗️ 映射基础

### 映射声明和初始化

```go
package main

import "fmt"

func main() {
	// 声明映射（nil 映射）
	var m1 map[string]int
	fmt.Printf("nil 映射: %v, 是否为 nil: %t\n", m1, m1 == nil)
	
	// 使用 make 创建映射
	m2 := make(map[string]int)
	fmt.Printf("空映射: %v, 长度: %d\n", m2, len(m2))
	
	// 创建时指定容量
	m3 := make(map[string]int, 10)
	fmt.Printf("指定容量的映射: %v, 容量: %d\n", m3, cap(m3))
	
	// 直接初始化
	m4 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	fmt.Printf("初始化映射: %v, 长度: %d\n", m4, len(m4))
	
	// 不同类型的映射
	stringMap := map[string]string{"hello": "world"}
	intMap := map[int]string{1: "one", 2: "two"}
	boolMap := map[bool]string{true: "yes", false: "no"}
	
	fmt.Printf("字符串映射: %v\n", stringMap)
	fmt.Printf("整数映射: %v\n", intMap)
	fmt.Printf("布尔映射: %v\n", boolMap)
}
```

### 映射的基本操作

```go
package main

import "fmt"

func main() {
	// 创建映射
	scores := make(map[string]int)
	
	// 添加元素
	scores["张三"] = 95
	scores["李四"] = 88
	scores["王五"] = 76
	fmt.Printf("添加元素后: %v\n", scores)
	
	// 获取元素
	score := scores["张三"]
	fmt.Printf("张三的成绩: %d\n", score)
	
	// 获取不存在的键
	unknownScore := scores["赵六"]
	fmt.Printf("赵六的成绩: %d (零值)\n", unknownScore)
	
	// 检查键是否存在
	value, exists := scores["李四"]
	fmt.Printf("李四是否存在: %t, 成绩: %d\n", exists, value)
	
	_, exists2 := scores["赵六"]
	fmt.Printf("赵六是否存在: %t\n", exists2)
	
	// 修改元素
	scores["张三"] = 98
	fmt.Printf("修改后: %v\n", scores)
	
	// 删除元素
	delete(scores, "王五")
	fmt.Printf("删除王五后: %v\n", scores)
	
	// 删除不存在的键（不会报错）
	delete(scores, "赵六")
	fmt.Printf("删除不存在的键后: %v\n", scores)
}
```

## 🔍 映射遍历

### 遍历键值对

```go
package main

import "fmt"

func main() {
	// 创建示例映射
	studentGrades := map[string]int{
		"张三": 95,
		"李四": 88,
		"王五": 76,
		"赵六": 92,
		"钱七": 84,
	}
	
	// 遍历键值对
	fmt.Println("遍历键值对:")
	for name, grade := range studentGrades {
		fmt.Printf("%s: %d\n", name, grade)
	}
	
	// 只遍历键
	fmt.Println("\n只遍历键:")
	for name := range studentGrades {
		fmt.Printf("学生: %s\n", name)
	}
	
	// 只遍历值
	fmt.Println("\n只遍历值:")
	for _, grade := range studentGrades {
		fmt.Printf("成绩: %d\n", grade)
	}
	
	// 收集所有键
	fmt.Println("\n收集所有键:")
	var names []string
	for name := range studentGrades {
		names = append(names, name)
	}
	fmt.Printf("所有学生: %v\n", names)
}
```

### 遍历顺序注意事项

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建映射
	ages := map[string]int{
		"张三": 25,
		"李四": 30,
		"王五": 28,
		"赵六": 22,
	}
	
	fmt.Println("随机顺序遍历:")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}
	
	// 如果需要有序遍历，先排序键
	fmt.Println("\n按键名排序后遍历:")
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, ages[name])
	}
	
	// 按值排序遍历
	fmt.Println("\n按年龄排序后遍历:")
	type NameAge struct {
		Name string
		Age  int
	}
	
	var nameAges []NameAge
	for name, age := range ages {
		nameAges = append(nameAges, NameAge{name, age})
	}
	
	// 按年龄排序
	for i := 0; i < len(nameAges)-1; i++ {
		for j := i + 1; j < len(nameAges); j++ {
			if nameAges[i].Age > nameAges[j].Age {
				nameAges[i], nameAges[j] = nameAges[j], nameAges[i]
			}
		}
	}
	
	for _, na := range nameAges {
		fmt.Printf("%s: %d\n", na.Name, na.Age)
	}
}
```

## 🎯 映射高级操作

### 映射作为集合使用

```go
package main

import "fmt"

func main() {
	// 使用映射模拟集合
	set := make(map[string]bool)
	
	// 添加元素到集合
	items := []string{"apple", "banana", "orange", "apple", "grape"}
	for _, item := range items {
		set[item] = true
	}
	
	fmt.Printf("集合元素: ")
	for item := range set {
		fmt.Printf("%s ", item)
	}
	fmt.Printf("(总数: %d)\n", len(set))
	
	// 检查元素是否存在
	fmt.Println("\n集合操作:")
	fmt.Printf("apple 在集合中: %t\n", set["apple"])
	fmt.Printf("pear 在集合中: %t\n", set["pear"])
	
	// 从集合中删除
	delete(set, "orange")
	fmt.Printf("删除 orange 后集合大小: %d\n", len(set))
	
	// 集合运算：交集
	set2 := map[string]bool{
		"apple":  true,
		"grape":  true,
		"pear":   true,
		"melon":  true,
	}
	
	fmt.Println("\n两个集合的交集:")
	for item := range set {
		if set2[item] {
			fmt.Printf("%s ", item)
		}
	}
	fmt.Println()
}
```

### 嵌套映射

```go
package main

import "fmt"

func main() {
	// 嵌套映射：学生各科成绩
	students := map[string]map[string]int{
		"张三": {
			"数学": 95,
			"英语": 88,
			"物理": 92,
		},
		"李四": {
			"数学": 87,
			"英语": 92,
			"物理": 85,
		},
		"王五": {
			"数学": 76,
			"英语": 84,
			"物理": 79,
		},
	}
	
	// 访问嵌套数据
	fmt.Println("学生成绩:")
	for name, subjects := range students {
		fmt.Printf("%s: ", name)
		for subject, score := range subjects {
			fmt.Printf("%s=%d ", subject, score)
		}
		fmt.Println()
	}
	
	// 获取特定学生的特定科目
	mathScore := students["张三"]["数学"]
	fmt.Printf("\n张三的数学成绩: %d\n", mathScore)
	
	// 添加新科目
	students["张三"]["化学"] = 90
	fmt.Printf("张三的化学成绩: %d\n", students["张三"]["化学"])
	
	// 安全访问嵌套映射
	if student, exists := students["赵六"]; exists {
		if mathScore, hasMath := student["数学"]; hasMath {
			fmt.Printf("赵六的数学成绩: %d\n", mathScore)
		} else {
			fmt.Println("赵六没有数学成绩")
		}
	} else {
		fmt.Println("赵六不存在")
	}
}
```

### 映射的函数式操作

```go
package main

import (
	"fmt"
	"strings"
)

// 映射过滤
func filterMap(m map[string]int, predicate func(string, int) bool) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// 映射转换
func transformMap(m map[string]int, transformer func(string, int) (string, int)) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		newKey, newValue := transformer(k, v)
		result[newKey] = newValue
	}
	return result
}

// 映射聚合
func reduceMap(m map[string]int, initial int, reducer func(int, string, int) int) int {
	result := initial
	for k, v := range m {
		result = reducer(result, k, v)
	}
	return result
}

func main() {
	// 示例映射
	wordCounts := map[string]int{
		"hello":   3,
		"world":   2,
		"go":      5,
		"programming": 1,
	}
	
	fmt.Printf("原始映射: %v\n", wordCounts)
	
	// 过滤：只保留出现次数 >= 3 的单词
	filtered := filterMap(wordCounts, func(word string, count int) bool {
		return count >= 3
	})
	fmt.Printf("过滤后: %v\n", filtered)
	
	// 转换：将单词转为大写，计数加倍
	transformed := transformMap(wordCounts, func(word string, count int) (string, int) {
		return strings.ToUpper(word), count * 2
	})
	fmt.Printf("转换后: %v\n", transformed)
	
	// 聚合：计算总词数
	totalCount := reduceMap(wordCounts, 0, func(acc int, word string, count int) int {
		return acc + count
	})
	fmt.Printf("总词数: %d\n", totalCount)
}
```

## 💡 实用技巧

### 使用空结构体作为集合

```go
package main

import "fmt"

func main() {
	// 使用空结构体作为集合元素，节省内存
	seen := make(map[string]struct{})
	
	items := []string{"apple", "banana", "apple", "orange", "banana"}
	for _, item := range items {
		seen[item] = struct{}{} // 空结构体不占用内存
	}
	
	fmt.Printf("唯一元素: ")
	for key := range seen {
		fmt.Printf("%s ", key)
	}
	fmt.Println()
}
```

### 键值对调

```go
package main

import "fmt"

func main() {
	// 原始映射
	original := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	
	// 键值对调
	reverse := make(map[int]string)
	for k, v := range original {
		reverse[v] = k
	}
	
	fmt.Printf("原始映射: %v\n", original)
	fmt.Printf("反转映射: %v\n", reverse)
}
```

### 复合键

```go
package main

import "fmt"

// 使用结构体作为复合键
type Key struct {
	Path string
	Type string
}

func main() {
	// 使用复合键
	m := make(map[Key]string)
	m[Key{Path: "/api/users", Type: "GET"}] = "用户列表"
	m[Key{Path: "/api/users", Type: "POST"}] = "创建用户"
	m[Key{Path: "/api/users/1", Type: "GET"}] = "用户详情"
	
	// 访问
	value := m[Key{Path: "/api/users", Type: "GET"}]
	fmt.Printf("路由处理: %s\n", value)
}
```

### Map 作为函数缓存

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var cache = make(map[string]interface{})
var cacheMu sync.RWMutex

// 模拟耗时计算
func computeExpensiveOperation(key string) interface{} {
	time.Sleep(100 * time.Millisecond) // 模拟耗时操作
	return fmt.Sprintf("result for %s", key)
}

func getWithCache(key string) interface{} {
	// 先读锁检查缓存
	cacheMu.RLock()
	if val, ok := cache[key]; ok {
		cacheMu.RUnlock()
		return val
	}
	cacheMu.RUnlock()
	
	// 计算值
	val := computeExpensiveOperation(key)
	
	// 写锁更新缓存
	cacheMu.Lock()
	cache[key] = val
	cacheMu.Unlock()
	
	return val
}

func main() {
	// 第一次调用，需要计算
	start := time.Now()
	result1 := getWithCache("key1")
	fmt.Printf("第一次: %v (耗时: %v)\n", result1, time.Since(start))
	
	// 第二次调用，从缓存获取
	start = time.Now()
	result2 := getWithCache("key1")
	fmt.Printf("第二次: %v (耗时: %v)\n", result2, time.Since(start))
}
```

### 使用 sync.Map 实现并发安全

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	
	// 存储
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")
	
	// 读取
	if value, ok := sm.Load("key1"); ok {
		fmt.Printf("key1: %v\n", value)
	}
	
	// 删除
	sm.Delete("key2")
	
	// 遍历
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true // 返回 false 停止遍历
	})
	
	// 加载或存储（如果不存在则存储）
	actual, loaded := sm.LoadOrStore("key3", "value3")
	fmt.Printf("key3 loaded: %t, value: %v\n", loaded, actual)
}
```

### 预分配容量

```go
package main

import "fmt"

func main() {
	// 当知道大约容量时，预分配可提高性能
	expectedSize := 100
	m := make(map[string]int, expectedSize)
	
	// 添加元素
	for i := 0; i < expectedSize; i++ {
		key := fmt.Sprintf("key%d", i)
		m[key] = i
	}
	
	fmt.Printf("映射大小: %d\n", len(m))
}
```

## 📊 实际应用

### 配置管理

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

// 配置管理器
type ConfigManager struct {
	configs map[string]string
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		configs: make(map[string]string),
	}
}

func (cm *ConfigManager) Set(key, value string) {
	cm.configs[key] = value
	fmt.Printf("设置配置: %s = %s\n", key, value)
}

func (cm *ConfigManager) GetString(key, defaultValue string) string {
	if value, exists := cm.configs[key]; exists {
		return value
	}
	return defaultValue
}

func (cm *ConfigManager) GetInt(key string, defaultValue int) int {
	if value, exists := cm.configs[key]; exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func (cm *ConfigManager) GetBool(key string, defaultValue bool) bool {
	if value, exists := cm.configs[key]; exists {
		return value == "true" || value == "1"
	}
	return defaultValue
}

func (cm *ConfigManager) PrintAll() {
	fmt.Println("当前配置:")
	for key, value := range cm.configs {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

func main() {
	config := NewConfigManager()
	
	// 设置配置
	config.Set("server.port", "8080")
	config.Set("server.host", "localhost")
	config.Set("debug.mode", "true")
	config.Set("max.connections", "100")
	
	config.PrintAll()
	
	// 获取配置
	port := config.GetInt("server.port", 3000)
	host := config.GetString("server.host", "127.0.0.1")
	debug := config.GetBool("debug.mode", false)
	maxConn := config.GetInt("max.connections", 50)
	
	fmt.Printf("\n使用配置:\n")
	fmt.Printf("服务器地址: %s:%d\n", host, port)
	fmt.Printf("调试模式: %t\n", debug)
	fmt.Printf("最大连接数: %d\n", maxConn)
}
```

### 缓存实现

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 简单的内存缓存
type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]interface{})
}

func main() {
	cache := NewCache()
	
	// 缓存操作
	cache.Set("user:1", "张三")
	cache.Set("user:2", "李四")
	cache.Set("session:abc123", map[string]interface{}{
		"user_id": 1,
		"login_time": time.Now(),
	})
	
	// 读取缓存
	if user1, exists := cache.Get("user:1"); exists {
		fmt.Printf("用户1: %v\n", user1)
	}
	
	if session, exists := cache.Get("session:abc123"); exists {
		fmt.Printf("会话信息: %v\n", session)
	}
	
	// 缓存统计
	fmt.Printf("缓存大小: %d\n", cache.Size())
	
	// 删除和清空
	cache.Delete("user:2")
	fmt.Printf("删除后缓存大小: %d\n", cache.Size())
	
	cache.Clear()
	fmt.Printf("清空后缓存大小: %d\n", cache.Size())
}
```

## 🧪 实践练习

### 练习 1：词频统计

```go
// 要求：
// 1. 统计文本中每个单词的出现次数
// 2. 按出现次数排序输出
// 3. 支持大小写不敏感
// 4. 过滤停用词
```

参考实现：
```go
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	freq := make(map[string]int)
	
	// 停用词
	stopWords := map[string]bool{
		"the": true, "a": true, "an": true, "and": true, "or": true,
		"but": true, "in": true, "on": true, "at": true, "to": true,
	}
	
	for _, word := range words {
		// 转为小写并清理
		word = strings.ToLower(word)
		word = strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		
		// 跳过停用词和空字符串
		if word == "" || stopWords[word] {
			continue
		}
		
		freq[word]++
	}
	
	return freq
}

func printWordFrequency(freq map[string]int) {
	type WordCount struct {
		Word  string
		Count int
	}
	
	var wordCounts []WordCount
	for word, count := range freq {
		wordCounts = append(wordCounts, WordCount{word, count})
	}
	
	// 按频率排序
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})
	
	fmt.Println("词频统计结果:")
	for i, wc := range wordCounts {
		fmt.Printf("%d. %s: %d\n", i+1, wc.Word, wc.Count)
		if i >= 9 { // 只显示前10个
			break
		}
	}
}

func main() {
	text := `The Go programming language is an open source project to make programmers more productive. 
	Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs 
	that get the most out of multicore and networked machines.`
	
	freq := wordFrequency(text)
	printWordFrequency(freq)
}
```

### 练习 2：数据库索引模拟

```go
// 要求：
// 1. 实现简单的倒排索引
// 2. 支持多字段索引
// 3. 实现基本的查询功能
// 4. 添加索引更新机制
```

### 练习 3：图结构实现

```go
// 要求：
// 1. 使用映射实现邻接表
// 2. 实现深度优先搜索
// 3. 实现广度优先搜索
// 4. 计算最短路径
```

## 🤔 思考题

1. **映射和切片的区别是什么？**
   - 提示：考虑键值对 vs 有序集合、复杂度

2. **为什么映射的遍历顺序是随机的？**
   - 提示：考虑哈希表的实现

3. **什么时候应该使用映射而不是切片？**
   - 提示：考虑查找频率、数据关联性

4. **映射的线程安全性如何保证？**
   - 提示：考虑并发访问、锁机制

## 📚 扩展阅读

- [Go 语言规范 - 映射类型](https://golang.org/ref/spec#Map_types)
- [Effective Go - 映射](https://golang.org/doc/effective_go.html#maps)
- [Go by Example - 映射](https://gobyexample.com/maps)

## ⏭️ 下一章节

[结构体](./09-structs.md) → 学习 Go 语言的结构体类型

---

**💡 小贴士**: 映射是 Go 语言中最重要的数据结构之一。记住：**映射提供了 O(1) 的平均查找时间，是构建高效算法的基础工具**！
