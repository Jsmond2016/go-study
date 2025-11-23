// Package main 展示 Go 语言的映射
package main

import (
	"fmt"
	"maps"
	"sort"
)

// 程序入口，演示映射操作
func main() {
	fmt.Println("=== Go 映射示例 ===")

	// 基本映射操作
	basicMaps()

	// 映射高级操作
	advancedMapOperations()

	// 映射遍历
	mapIteration()

	// 映射作为函数参数
	mapAsParameter()

	// 嵌套映射
	nestedMaps()

	// 实际应用示例
	practicalExample()
}

// basicMaps 基本映射操作
func basicMaps() {
	fmt.Println("\n--- 基本映射操作 ---")

	// 1. 映射声明和初始化
	var map1 map[string]int                    // nil 映射
	map2 := make(map[string]int)              // 空映射
	map3 := map[string]int{                   // 直接初始化
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	fmt.Printf("nil 映射: %v (isNil=%t)\n", map1, map1 == nil)
	fmt.Printf("空映射: %v (len=%d)\n", map2, len(map2))
	fmt.Printf("初始化映射: %v (len=%d)\n", map3, len(map3))

	// 2. 添加和修改元素
	map2["grape"] = 4
	map2["apple"] = 5
	fmt.Printf("\n添加元素后: %v\n", map2)

	// 修改元素
	map2["apple"] = 10
	fmt.Printf("修改元素后: %v\n", map2)

	// 3. 访问元素
	fmt.Printf("\n访问元素:")
	value, exists := map2["apple"]
	if exists {
		fmt.Printf("apple = %d\n", value)
	} else {
		fmt.Println("apple 不存在")
	}

	value, exists = map2["nonexistent"]
	if exists {
		fmt.Printf("nonexistent = %d\n", value)
	} else {
		fmt.Println("nonexistent 不存在")
	}

	// 4. 删除元素
	fmt.Printf("\n删除操作:")
	fmt.Printf("删除前: %v\n", map2)
	delete(map2, "apple")
	fmt.Printf("删除 apple 后: %v\n", map2)
	delete(map2, "nonexistent") // 删除不存在的键不会有任何效果
}

// advancedMapOperations 映射高级操作
func advancedMapOperations() {
	fmt.Println("\n--- 映射高级操作 ---")

	// 1. 映射的复制
	original := map[string]int{
		"A": 10,
		"B": 20,
		"C": 30,
	}

	// 深拷贝
	copyMap := make(map[string]int)
	for key, value := range original {
		copyMap[key] = value
	}

	fmt.Printf("原映射: %v\n", original)
	fmt.Printf("复制映射: %v\n", copyMap)

	// 修改复制品不影响原映射
	copyMap["A"] = 100
	fmt.Printf("修改复制品后:\n")
	fmt.Printf("原映射: %v\n", original)
	fmt.Printf("复制映射: %v\n", copyMap)

	// 2. 映射的比较（使用 Go 1.21+ 的 maps 包）
	otherMap := map[string]int{
		"A": 10,
		"B": 20,
		"C": 30,
	}

	fmt.Printf("\n映射比较:")
	fmt.Printf("original == otherMap: %t\n", maps.Equal(original, otherMap))

	otherMap["A"] = 100
	fmt.Printf("修改后 original == otherMap: %t\n", maps.Equal(original, otherMap))

	// 3. 映射的合并
	map1 := map[string]int{"A": 1, "B": 2}
	map2 := map[string]int{"B": 20, "C": 3, "D": 4}

	// 合并映射（map2 会覆盖 map1 的相同键）
	merged := make(map[string]int)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}

	fmt.Printf("\n映射合并:\n")
	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)
	fmt.Printf("merged: %v\n", merged)
}

// mapIteration 映射遍历
func mapIteration() {
	fmt.Println("\n--- 映射遍历 ---")

	grades := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
		"赵六": 96,
		"钱七": 88,
	}

	// 1. 基本遍历
	fmt.Println("基本遍历:")
	for name, score := range grades {
		fmt.Printf("%s: %d\n", name, score)
	}

	// 2. 只遍历键
	fmt.Println("\n只遍历键:")
	for name := range grades {
		fmt.Printf("%s\n", name)
	}

	// 3. 只遍历值
	fmt.Println("\n只遍历值:")
	for _, score := range grades {
		fmt.Printf("%d\n", score)
	}

	// 4. 有序遍历（映射本身是无序的）
	fmt.Println("\n按姓名排序遍历:")
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names) // 对姓名排序

	for _, name := range names {
		fmt.Printf("%s: %d\n", name, grades[name])
	}

	// 5. 按值排序遍历
	fmt.Println("\n按分数排序遍历:")
	type nameScore struct {
		name  string
		score int
	}
	var pairs []nameScore
	for name, score := range grades {
		pairs = append(pairs, nameScore{name, score})
	}

	// 按分数降序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair.name, pair.score)
	}
}

// mapAsParameter 映射作为函数参数
func mapAsParameter() {
	fmt.Println("\n--- 映射作为函数参数 ---")

	// 映射是引用类型，函数内修改会影响原映射
	grades := map[string]int{
		"数学": 90,
		"英语": 85,
		"编程": 95,
	}

	fmt.Printf("原映射: %v\n", grades)

	increaseScore(grades, "数学", 5)
	fmt.Printf("修改后: %v\n", grades)

	// 获取最高分科目
	subject, score := findHighestScore(grades)
	fmt.Printf("最高分科目: %s (%d分)\n", subject, score)
}

// increaseScore 增加指定科目的分数
func increaseScore(grades map[string]int, subject string, points int) {
	if _, exists := grades[subject]; exists {
		grades[subject] += points
		fmt.Printf("增加了 %s 的分数 %d 分\n", subject, points)
	} else {
		fmt.Printf("科目 %s 不存在\n", subject)
	}
}

// findHighestScore 查找最高分科目
func findHighestScore(grades map[string]int) (string, int) {
	var highestSubject string
	var highestScore int

	for subject, score := range grades {
		if score > highestScore {
			highestSubject = subject
			highestScore = score
		}
	}

	return highestSubject, highestScore
}

// nestedMaps 嵌套映射
func nestedMaps() {
	fmt.Println("\n--- 嵌套映射 ---")

	// 1. 二维映射
	departmentEmployees := map[string]map[string]int{
		"技术部": {
			"张三": 25,
			"李四": 30,
			"王五": 28,
		},
		"销售部": {
			"赵六": 32,
			"钱七": 26,
		},
		"人事部": {
			"孙八": 29,
		},
	}

	fmt.Printf("公司员工信息:\n")
	for department, employees := range departmentEmployees {
		fmt.Printf("%s:\n", department)
		for name, age := range employees {
			fmt.Printf("  %s: %d岁\n", name, age)
		}
	}

	// 2. 动态添加部门和员工
	fmt.Printf("\n动态添加:\n")
	if _, exists := departmentEmployees["财务部"]; !exists {
		departmentEmployees["财务部"] = make(map[string]int)
	}
	departmentEmployees["财务部"]["周九"] = 35

	// 为技术部添加新员工
	departmentEmployees["技术部"]["吴十"] = 27

	fmt.Printf("添加后的员工信息:\n")
	for department, employees := range departmentEmployees {
		fmt.Printf("%s (%d人):\n", department, len(employees))
		for name, age := range employees {
			fmt.Printf("  %s: %d岁\n", name, age)
		}
	}
}

// practicalExample 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1：单词计数器
	fmt.Println("1. 单词计数器:")
	text := "Go is an open source programming language that makes it easy to build simple reliable and efficient software"
	wordCount := countWords(text)
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("单词统计:\n")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d\n", word, count)
	}

	// 示例2：配置管理
	fmt.Println("\n2. 配置管理:")
	config := map[string]interface{}{
		"database": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"name":     "myapp",
			"user":     "admin",
			"password": "secret123",
		},
		"server": map[string]interface{}{
			"host": "0.0.0.0",
			"port": 8080,
			"debug": false,
		},
		"features": []string{"auth", "logging", "metrics"},
	}

	fmt.Printf("系统配置:\n")
	printConfig(config)

	// 示例3：缓存系统
	fmt.Println("\n3. 简单缓存系统:")
	cache := NewCache()

	// 存储数据
	cache.Set("user:1", "张三")
	cache.Set("user:2", "李四")
	cache.Set("token:abc123", "valid_token")

	// 获取数据
	if user, found := cache.Get("user:1"); found {
		fmt.Printf("缓存命中: user:1 = %s\n", user)
	}

	if user, found := cache.Get("user:999"); found {
		fmt.Printf("缓存命中: user:999 = %s\n", user)
	} else {
		fmt.Println("缓存未命中: user:999")
	}

	// 显示所有缓存
	fmt.Printf("当前缓存内容: %v\n", cache.GetAll())

	// 示例4：数据统计
	fmt.Println("\n4. 数据统计:")
	salesData := []struct {
		Product  string
		Quantity int
		Price    float64
	}{
		{"苹果", 100, 5.5},
		{"香蕉", 80, 3.2},
		{"橙子", 120, 4.8},
		{"苹果", 50, 5.5},
		{"葡萄", 60, 8.9},
		{"香蕉", 90, 3.2},
		{"橙子", 70, 4.8},
	}

	statistics := analyzeSales(salesData)
	fmt.Printf("销售统计:\n")
	for product, stats := range statistics {
		fmt.Printf("  %s: 销量=%d, 总收入=¥%.1f, 平均价格=¥%.2f\n",
			product, stats.Quantity, stats.TotalRevenue, stats.AvgPrice)
	}
}

// countWords 单词计数
func countWords(text string) map[string]int {
	words := make(map[string]int)
	for _, word := range text {
		if word != ' ' && word != ',' && word != '.' {
			words[string(word)]++
		}
	}
	return words
}

// printConfig 打印配置
func printConfig(config map[string]interface{}) {
	for key, value := range config {
		switch v := value.(type) {
		case map[string]interface{}:
			fmt.Printf("  %s:\n", key)
			for subKey, subValue := range v {
				fmt.Printf("    %s: %v\n", subKey, subValue)
			}
		case []string:
			fmt.Printf("  %s: %v\n", key, v)
		default:
			fmt.Printf("  %s: %v\n", key, v)
		}
	}
}

// Cache 简单缓存
type Cache struct {
	data map[string]interface{}
}

// NewCache 创建新缓存
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set 存储键值对
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get 获取值
func (c *Cache) Get(key string) (interface{}, bool) {
	value, exists := c.data[key]
	return value, exists
}

// GetAll 获取所有数据
func (c *Cache) GetAll() map[string]interface{} {
	// 返回副本以避免外部修改
	copy := make(map[string]interface{})
	for k, v := range c.data {
		copy[k] = v
	}
	return copy
}

// Delete 删除键
func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

// Stats 销售统计
type Stats struct {
	Quantity    int
	TotalRevenue float64
	AvgPrice    float64
}

// analyzeSales 分析销售数据
func analyzeSales(sales []struct {
	Product  string
	Quantity int
	Price    float64
}) map[string]*Stats {
	statistics := make(map[string]*Stats)

	for _, sale := range sales {
		if _, exists := statistics[sale.Product]; !exists {
			statistics[sale.Product] = &Stats{}
		}

		stats := statistics[sale.Product]
		stats.Quantity += sale.Quantity
		stats.TotalRevenue += float64(sale.Quantity) * sale.Price
		stats.AvgPrice = stats.TotalRevenue / float64(stats.Quantity)
	}

	return statistics
}