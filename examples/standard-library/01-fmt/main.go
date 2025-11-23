// Package main 展示 fmt 包的各种用法
package main

import (
	"fmt"
	"os"
)

// 程序入口，演示 fmt 包的各种功能
func main() {
	fmt.Println("=== fmt 包示例 ===")

	// 基本输出函数
	basicPrinting()

	// 格式化打印
	formattedPrinting()

	// 输入函数
	inputFunctions()

	// 格式化动词
	formattingVerbs()

	// 高级格式化
	advancedFormatting()

	// 错误和文件操作
	errorAndFileOperations()

	// 实际应用示例
	practicalExample()
}

// basicPrinting 基本输出函数
func basicPrinting() {
	fmt.Println("\n--- 基本输出函数 ---")

	// 1. Print
	fmt.Print("Hello, ")
	fmt.Print("Go!\n") // 需要手动添加换行

	// 2. Println - 自动添加空格和换行
	fmt.Println("Hello,", "Go!")
	fmt.Println("数字:", 42, "字符串:", "test")

	// 3. Printf - 格式化输出
	name := "张三"
	age := 25
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)

	// 4. Fprint - 输出到指定文件
	fmt.Fprintln(os.Stdout, "输出到标准输出")

	// 5. Sprint - 返回字符串而不是打印
	result := fmt.Sprint("拼接", "多个", "字符串")
	fmt.Printf("Sprint 结果: %s\n", result)

	// 6. Sprintf - 格式化返回字符串
	formatted := fmt.Sprintf("姓名: %s, 年龄: %d", name, age)
	fmt.Printf("Sprintf 结果: %s\n", formatted)
}

// formattedPrinting 格式化打印
func formattedPrinting() {
	fmt.Println("\n--- 格式化打印 ---")

	// 1. 整数格式化
	fmt.Printf("十进制: %d\n", 42)
	fmt.Printf("二进制: %b\n", 42)
	fmt.Printf("八进制: %o\n", 42)
	fmt.Printf("十六进制: %x\n", 42)
	fmt.Printf("十六进制(大写): %X\n", 42)

	// 2. 浮点数格式化
	fmt.Printf("默认浮点数: %f\n", 3.14159)
	fmt.Printf("保留2位小数: %.2f\n", 3.14159)
	fmt.Printf("科学计数法: %e\n", 3.14159)
	fmt.Printf("科学计数法(大写): %E\n", 3.14159)
	fmt.Printf("最简格式: %g\n", 3.14159)

	// 3. 字符串格式化
	fmt.Printf("字符串: %s\n", "Hello, Go!")
	fmt.Printf("引用字符串: %q\n", "Hello, Go!")
	fmt.Printf("字节切片: %x\n", []byte("Hello"))

	// 4. 布尔值格式化
	fmt.Printf("布尔值: %t\n", true)
	fmt.Printf("布尔值: %t\n", false)

	// 5. 指针格式化
	value := 42
	ptr := &value
	fmt.Printf("指针地址: %p\n", ptr)
	fmt.Printf("指针值: %v\n", ptr)
}

// inputFunctions 输入函数
func inputFunctions() {
	fmt.Println("\n--- 输入函数 ---")

	// 演示 Scan（需要用户交互，这里用注释说明）
	/*
	// 1. Scan - 从标准输入读取
	var name string
	var age int
	fmt.Print("请输入姓名: ")
	fmt.Scan(&name)
	fmt.Print("请输入年龄: ")
	fmt.Scan(&age)
	fmt.Printf("输入结果: 姓名=%s, 年龄=%d\n", name, age)

	// 2. Scanln - 读取到换行符
	fmt.Print("请输入姓名和年龄(用空格分隔): ")
	var scanName string
	var scanAge int
	fmt.Scanln(&scanName, &scanAge)
	fmt.Printf("Scanln 结果: 姓名=%s, 年龄=%d\n", scanName, scanAge)

	// 3. Scanf - 格式化输入
	var fName string
	var fAge int
	fmt.Print("请输入格式化数据(姓名,年龄): ")
	fmt.Scanf("%s,%d", &fName, &fAge)
	fmt.Printf("Scanf 结果: 姓名=%s, 年龄=%d\n", fName, fAge)
	*/

	fmt.Println("输入函数需要交互式环境，这里展示基本用法")
	fmt.Println("// fmt.Scan(&name) // 读取输入")
	fmt.Println("// fmt.Scanf(\"%s\", &name) // 格式化读取")
}

// formattingVerbs 格式化动词详解
func formattingVerbs() {
	fmt.Println("\n--- 格式化动词详解 ---")

	// 通用动词
	fmt.Printf("%v (默认格式): %v\n", "Hello")
	fmt.Printf("%+v (结构体显示字段名): %+v\n", struct{ Name string }{"张三"})
	fmt.Printf("%#v (Go语法表示): %#v\n", struct{ Name string }{"张三"})
	fmt.Printf("%T (类型): %T\n", "Hello")

	// 布尔值
	fmt.Printf("%t (布尔): %t\n", true)

	// 整数
	fmt.Printf("%b (二进制): %b\n", 13)
	fmt.Printf("%c (对应Unicode字符): %c\n", 65)
	fmt.Printf("%d (十进制): %d\n", 13)
	fmt.Printf("%o (八进制): %o\n", 13)
	fmt.Printf("%q (单引号字符): %q\n", 65)
	fmt.Printf("%x (十六进制): %x\n", 13)
	fmt.Printf("%X (十六进制大写): %X\n", 13)
	fmt.Printf("%U (Unicode格式): %U\n", 65)

	// 浮点数
	fmt.Printf("%b (科学记数法二进制): %b\n", 3.14159)
	fmt.Printf("%e (科学记数法): %e\n", 3.14159)
	fmt.Printf("%E (科学记数法大写): %E\n", 3.14159)
	fmt.Printf("%f (小数): %f\n", 3.14159)
	fmt.Printf("%F (小数大写): %F\n", 3.14159)
	fmt.Printf("%g (最紧凑): %g\n", 3.14159)
	fmt.Printf("%G (最紧凑大写): %G\n", 3.14159)

	// 字符串和字节切片
	fmt.Printf("%s (字符串): %s\n", "Hello")
	fmt.Printf("%q (带引号字符串): %q\n", "Hello")
	fmt.Printf("%x (十六进制字节): %x\n", "Hello")
	fmt.Printf("%X (十六进制字节大写): %X\n", "Hello")

	// 指针
	value := 42
	ptr := &value
	fmt.Printf("%p (指针地址): %p\n", ptr)
}

// advancedFormatting 高级格式化
func advancedFormatting() {
	fmt.Println("\n--- 高级格式化 ---")

	// 1. 宽度和精度
	fmt.Printf("|%6d| (宽度6): |%6d|\n", 42)
	fmt.Printf("|%-6d| (左对齐): |%-6d|\n", 42)
	fmt.Printf("|%06d| (补零): |%06d|\n", 42)
	fmt.Printf("|%6.2f| (宽度6精度2): |%6.2f|\n", 3.14159)

	// 2. 字符串对齐
	fmt.Printf("|%10s| (右对齐): |%10s|\n", "Hello")
	fmt.Printf("|%-10s| (左对齐): |%-10s|\n", "Hello")
	fmt.Printf("|%10.3s| (截断): |%10.3s|\n", "Hello, World!")

	// 3. 多行格式化
	fmt.Printf("多行格式化:\n")
	fmt.Printf("第一行: %d\n", 1)
	fmt.Printf("第二行: %d\n", 2)

	// 4. 转义字符
	fmt.Printf("转义字符示例:\n")
	fmt.Printf("制表符: Hello\tWorld\n")
	fmt.Printf("换行符: Hello\nWorld\n")
	fmt.Printf("反斜杠: C:\\\\Windows\n")
	fmt.Printf("引号: \"Hello, Go!\"\n")
}

// errorAndFileOperations 错误和文件操作
func errorAndFileOperations() {
	fmt.Println("\n--- 错误和文件操作 ---")

	// 1. Errorf - 创建错误
	err := fmt.Errorf("用户 %d 不存在", 404)
	fmt.Printf("创建的错误: %v\n", err)

	// 2. Fprint - 写入文件（示例）
	fmt.Println("文件输出示例:")
	fmt.Fprintf(os.Stdout, "写入标准输出: %s\n", "Hello, File!")

	// 3. 自定义错误类型
	type MyError struct {
		Code    int
		Message string
	}

	func (e MyError) Error() string {
		return fmt.Sprintf("错误 %d: %s", e.Code, e.Message)
	}

	myErr := MyError{Code: 500, Message: "服务器内部错误"}
	fmt.Printf("自定义错误: %v\n", myErr)
}

// practicalExample 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1: 格式化表格
	fmt.Println("1. 格式化表格:")
	formatTable()

	// 示例2: 进度条
	fmt.Println("\n2. 简单进度条:")
	showProgress()

	// 示例3: 日志格式化
	fmt.Println("\n3. 日志格式化:")
	formatLog()

	// 示例4: 配置显示
	fmt.Println("\n4. 配置显示:")
	showConfig()
}

// formatTable 格式化表格
func formatTable() {
	type User struct {
		ID     int
		Name   string
		Email  string
		Score  float64
		Active bool
	}

	users := []User{
		{1, "张三", "zhangsan@example.com", 85.5, true},
		{2, "李四", "lisi@example.com", 92.0, true},
		{3, "王五", "wangwu@example.com", 78.5, false},
		{4, "赵六", "zhaoliu@example.com", 88.0, true},
	}

	// 表头
	fmt.Printf("|%-4s|%-10s|%-25s|%-8s|%-6s|\n", "ID", "姓名", "邮箱", "分数", "活跃")
	fmt.Printf("+----+----------+--------------------------+--------+------+\n")

	// 数据行
	for _, user := range users {
		fmt.Printf("|%-4d|%-10s|%-25s|%-8.1f|%-6t|\n",
			user.ID, user.Name, user.Email, user.Score, user.Active)
	}
}

// showProgress 显示进度条
func showProgress() {
	total := 100
	for i := 0; i <= total; i += 10 {
		percent := i * 100 / total
		barLength := 30
		filled := i * barLength / total

		bar := ""
		for j := 0; j < barLength; j++ {
			if j < filled {
				bar += "="
			} else {
				bar += " "
			}
		}

		fmt.Printf("\r进度: [%s] %d%% (%d/%d)", bar, percent, i, total)
	}
	fmt.Println() // 换行
}

// formatLog 格式化日志
func formatLog() {
	logs := []struct {
		Level   string
		Message string
		Time    string
		File    string
		Line    int
	}{
		{"INFO", "应用启动成功", "2024-01-15 10:30:00", "main.go", 25},
		{"WARN", "内存使用率较高", "2024-01-15 10:31:15", "monitor.go", 42},
		{"ERROR", "数据库连接失败", "2024-01-15 10:32:30", "database.go", 18},
	}

	for _, log := range logs {
		fmt.Printf("[%s] %s - %s (%s:%d)\n",
			log.Level, log.Time, log.Message, log.File, log.Line)
	}
}

// showConfig 显示配置
func showConfig() {
	type Config struct {
		Server struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		} `json:"server"`
		Database struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"database"`
		Features []string `json:"features"`
	}

	config := Config{
		Server: struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		}{
			Host: "localhost",
			Port: 8080,
		},
		Database: struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Name     string `json:"name"`
			Username string `json:"username"`
		}{
			Host:     "localhost",
			Port:     5432,
			Name:     "myapp",
			Username: "admin",
		},
		Features: []string{"auth", "logging", "metrics"},
	}

	fmt.Printf("服务器配置:\n")
	fmt.Printf("  地址: %s:%d\n", config.Server.Host, config.Server.Port)

	fmt.Printf("\n数据库配置:\n")
	fmt.Printf("  连接: %s:%d\n", config.Database.Host, config.Database.Port)
	fmt.Printf("  名称: %s\n", config.Database.Name)
	fmt.Printf("  用户: %s\n", config.Database.Username)

	fmt.Printf("\n启用功能: %v\n", config.Features)
}