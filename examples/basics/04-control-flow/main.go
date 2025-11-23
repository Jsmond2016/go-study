// Package main 展示 Go 语言的控制流程
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 程序入口，演示各种控制流程
func main() {
	fmt.Println("=== Go 控制流程示例 ===")

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 条件语句
	conditionalStatements()

	// 循环语句
	loopStatements()

	// 分支语句
	branchStatements()

	// 跳转语句
	jumpStatements()

	// 综合示例
	comprehensiveExample()
}

// conditionalStatements 条件语句
func conditionalStatements() {
	fmt.Println("\n--- 条件语句 ---")

	// 1. if 语句
	fmt.Println("if 语句:")
	age := 20
	if age >= 18 {
		fmt.Printf("%d 岁，已成年\n", age)
	}

	// 2. if-else 语句
	fmt.Println("\nif-else 语句:")
	score := 75
	if score >= 60 {
		fmt.Printf("分数 %d，及格\n", score)
	} else {
		fmt.Printf("分数 %d，不及格\n", score)
	}

	// 3. if-else if-else 语句
	fmt.Println("\nif-else if-else 语句:")
	testScore := 85
	if testScore >= 90 {
		fmt.Printf("分数 %d，优秀\n", testScore)
	} else if testScore >= 80 {
		fmt.Printf("分数 %d，良好\n", testScore)
	} else if testScore >= 70 {
		fmt.Printf("分数 %d，中等\n", testScore)
	} else if testScore >= 60 {
		fmt.Printf("分数 %d，及格\n", testScore)
	} else {
		fmt.Printf("分数 %d，不及格\n", testScore)
	}

	// 4. if 初始化语句
	fmt.Println("\nif 初始化语句:")
	if num := generateRandomNumber(1, 100); num > 50 {
		fmt.Printf("随机数 %d 大于 50\n", num)
	} else {
		fmt.Printf("随机数 %d 小于等于 50\n", num)
	}

	// 5. 复杂条件
	fmt.Println("\n复杂条件:")
	username := "admin"
	password := "123456"
	isLoggedIn := true

	if username == "admin" && password == "123456" && isLoggedIn {
		fmt.Println("管理员登录成功")
	} else if username == "guest" && isLoggedIn {
		fmt.Println("访客登录成功")
	} else {
		fmt.Println("登录失败")
	}
}

// loopStatements 循环语句
func loopStatements() {
	fmt.Println("\n--- 循环语句 ---")

	// 1. 基本 for 循环
	fmt.Println("基本 for 循环:")
	fmt.Print("数字 1-5: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 2. 只有条件的 for 循环（类似 while）
	fmt.Println("\n只有条件的 for 循环:")
	count := 1
	fmt.Print("计数到 3: ")
	for count <= 3 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// 3. 无限循环
	fmt.Println("\n无限循环（使用 break 退出）:")
	loopCount := 0
	for {
		fmt.Printf("循环 %d\n", loopCount+1)
		loopCount++
		if loopCount >= 3 {
			break
		}
	}

	// 4. for range 遍历
	fmt.Println("\nfor range 遍历:")

	// 遍历数组
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Print("数组元素: ")
	for index, value := range numbers {
		fmt.Printf("[%d]=%d ", index, value)
	}
	fmt.Println()

	// 遍历切片
	fruits := []string{"苹果", "香蕉", "橙子"}
	fmt.Print("切片元素: ")
	for i, fruit := range fruits {
		fmt.Printf("%d:%s ", i, fruit)
	}
	fmt.Println()

	// 遍历字符串
	text := "Hello"
	fmt.Print("字符串字符: ")
	for i, char := range text {
		fmt.Printf("[%d]=%c ", i, char)
	}
	fmt.Println()

	// 遍历映射
	grades := map[string]int{"数学": 90, "英语": 85, "编程": 95}
	fmt.Print("映射键值对: ")
	for subject, score := range grades {
		fmt.Printf("%s:%d ", subject, score)
	}
	fmt.Println()

	// 只要键
	fmt.Print("映射键: ")
	for subject := range grades {
		fmt.Printf("%s ", subject)
	}
	fmt.Println()

	// 只要值
	fmt.Print("映射值: ")
	for _, score := range grades {
		fmt.Printf("%d ", score)
	}
	fmt.Println()
}

// branchStatements 分支语句
func branchStatements() {
	fmt.Println("\n--- 分支语句 ---")

	// 1. 基本 switch 语句
	fmt.Println("基本 switch 语句:")
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("无效的星期")
	}

	// 2. switch 初始化语句
	fmt.Println("\nswitch 初始化语句:")
	switch num := generateRandomNumber(1, 10); num % 3 {
	case 0:
		fmt.Printf("随机数 %d 能被3整除\n", num)
	case 1:
		fmt.Printf("随机数 %d 除以3余1\n", num)
	case 2:
		fmt.Printf("随机数 %d 除以3余2\n", num)
	}

	// 3. 无表达式的 switch（用于条件判断）
	fmt.Println("\n无表达式的 switch:")
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Printf("温度 %d°C：极寒\n", temperature)
	case temperature < 10:
		fmt.Printf("温度 %d°C：寒冷\n", temperature)
	case temperature < 20:
		fmt.Printf("温度 %d°C：凉爽\n", temperature)
	case temperature < 30:
		fmt.Printf("温度 %d°C：温暖\n", temperature)
	default:
		fmt.Printf("温度 %d°C：炎热\n", temperature)
	}

	// 4. fallthrough 关键字
	fmt.Println("\nfallthrough 关键字:")
	score := 75
	switch {
	case score >= 90:
		fmt.Println("优秀")
		fallthrough
	case score >= 80:
		fmt.Println("良好")
		fallthrough
	case score >= 70:
		fmt.Println("中等")
		fallthrough
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 5. 类型 switch
	fmt.Println("\n类型 switch:")
	testType(42)
	testType("Hello")
	testType(true)
	testType(3.14)
}

// jumpStatements 跳转语句
func jumpStatements() {
	fmt.Println("\n--- 跳转语句 ---")

	// 1. break 语句
	fmt.Println("break 语句:")
	fmt.Print("寻找第一个偶数: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("找到 %d，停止搜索\n", i)
			break
		}
		fmt.Printf("%d ", i)
	}

	// 2. continue 语句
	fmt.Println("\n\ncontinue 语句:")
	fmt.Print("打印1-10中的奇数: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 3. goto 语句
	fmt.Println("\ngoto 语句:")
	fmt.Println("模拟循环:")
	i := 1
loop:
	if i <= 3 {
		fmt.Printf("第 %d 次循环\n", i)
		i++
		goto loop
	}
	fmt.Println("循环结束")

	// 4. 带标签的 break 和 continue
	fmt.Println("\n带标签的 break 和 continue:")
outer:
	for i := 1; i <= 3; i++ {
		fmt.Printf("外层循环 %d:\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("  内层循环 %d\n", j)
			if i == 2 && j == 2 {
				fmt.Println("  跳出外层循环")
				break outer
			}
			if j == 2 {
				fmt.Println("  继续下一次外层循环")
				continue outer
			}
		}
	}
}

// comprehensiveExample 综合示例
func comprehensiveExample() {
	fmt.Println("\n--- 综合示例：学生成绩管理系统 ---")

	// 模拟学生成绩数据
	students := []struct {
		Name    string
		Chinese int
		Math    int
		English int
	}{
		{"张三", 85, 90, 88},
		{"李四", 76, 85, 92},
		{"王五", 92, 78, 85},
		{"赵六", 68, 72, 75},
	}

	fmt.Printf("共有 %d 名学生\n\n", len(students))

	// 计算并显示每个学生的成绩
	for _, student := range students {
		total := student.Chinese + student.Math + student.English
		average := float64(total) / 3.0

		fmt.Printf("学生：%s\n", student.Name)
		fmt.Printf("  语文：%d，数学：%d，英语：%d\n", student.Chinese, student.Math, student.English)
		fmt.Printf("  总分：%d，平均分：%.1f\n", total, average)

		// 成绩等级评定
		switch {
		case average >= 90:
			fmt.Printf("  等级：优秀 ⭐⭐⭐\n")
		case average >= 80:
			fmt.Printf("  等级：良好 ⭐⭐\n")
		case average >= 70:
			fmt.Printf("  等级：中等 ⭐\n")
		case average >= 60:
			fmt.Printf("  等级：及格\n")
		default:
			fmt.Printf("  等级：不及格 ❌\n")
		}

		// 奖励检查
		rewards := 0
		if student.Chinese >= 90 {
			fmt.Printf("  语文优秀奖！\n")
			rewards++
		}
		if student.Math >= 90 {
			fmt.Printf("  数学优秀奖！\n")
			rewards++
		}
		if student.English >= 90 {
			fmt.Printf("  英语优秀奖！\n")
			rewards++
		}

		if rewards >= 2 {
			fmt.Printf("  多科优秀，获得特别奖！🏆\n")
		}

		fmt.Println()
	}

	// 统计信息
	var totalStudents, excellentCount int
	var totalScore float64

	for _, student := range students {
		studentTotal := student.Chinese + student.Math + student.English
		studentAvg := float64(studentTotal) / 3.0
		totalScore += studentAvg
		totalStudents++

		if studentAvg >= 85 {
			excellentCount++
		}
	}

	classAverage := totalScore / float64(totalStudents)
	excellentRate := float64(excellentCount) / float64(totalStudents) * 100

	fmt.Printf("班级统计：\n")
	fmt.Printf("  总人数：%d\n", totalStudents)
	fmt.Printf("  班级平均分：%.1f\n", classAverage)
	fmt.Printf("  优秀率：%.1f%%\n", excellentRate)
}

// 辅助函数
func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func testType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整型值：%d\n", v)
	case string:
		fmt.Printf("字符串值：%s\n", v)
	case bool:
		fmt.Printf("布尔值：%t\n", v)
	case float64:
		fmt.Printf("浮点值：%.2f\n", v)
	default:
		fmt.Printf("未知类型：%T\n", v)
	}
}