// Package main 展示 flag 包的各种用法
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// 定义全局变量存储命令行参数
var (
	name    = flag.String("name", "Guest", "用户姓名")
	age     = flag.Int("age", 0, "用户年龄")
	married = flag.Bool("married", false, "是否已婚")
	height  = flag.Float64("height", 0.0, "身高（米）")
	timeout = flag.Duration("timeout", 30*time.Second, "超时时间")
)

func main() {
	fmt.Println("=== flag 包示例 ===")

	// 自定义帮助信息
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: %s [选项]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "这是一个演示 flag 包的示例程序\n\n")
		fmt.Fprintf(os.Stderr, "选项:\n")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	// 显示解析结果
	showParsedFlags()

	// 演示不同类型的标志类型
	demonstrateFlagTypes()

	// 演示标志集合
	demonstrateFlagSet()

	// 实际应用示例
	practicalExample()
}

// showParsedFlags 显示解析结果
func showParsedFlags() {
	fmt.Println("\n--- 解析结果 ---")

	fmt.Printf("姓名: %s\n", *name)
	fmt.Printf("年龄: %d\n", *age)
	fmt.Printf("已婚: %t\n", *married)
	fmt.Printf("身高: %.2f米\n", *height)
	fmt.Printf("超时: %v\n", *timeout)

	// 显示非选项参数
	fmt.Printf("\n非选项参数: %v\n", flag.Args())
	fmt.Printf("参数数量: %d\n", flag.NArg())
}

// demonstrateFlagTypes 演示不同类型的标志类型
func demonstrateFlagTypes() {
	fmt.Println("\n--- 标志类型演示 ---")

	// 创建新的标志集用于演示
	fs := flag.NewFlagSet("demo", flag.ContinueOnError)

	// 基本类型
	var (
		stringFlag    = fs.String("str", "default", "字符串标志")
		intFlag       = fs.Int("int", 42, "整数标志")
		boolFlag      = fs.Bool("bool", true, "布尔标志")
		floatFlag     = fs.Float64("float", 3.14, "浮点数标志")
		durationFlag  = fs.Duration("duration", time.Second, "持续时间标志")
	)

	// 解析示例参数
	demoArgs := []string{"-str", "Hello", "-int", "100", "-bool", "false", "-float", "2.71", "-duration", "5s"}
	err := fs.Parse(demoArgs)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}

	fmt.Printf("字符串: %s\n", *stringFlag)
	fmt.Printf("整数: %d\n", *intFlag)
	fmt.Printf("布尔: %t\n", *boolFlag)
	fmt.Printf("浮点数: %.2f\n", *floatFlag)
	fmt.Printf("持续时间: %v\n", *durationFlag)
}

// demonstrateFlagSet 演示标志集合
func demonstrateFlagSet() {
	fmt.Println("\n--- 标志集合演示 ---")

	// 创建两个不同的标志集
	userFS := flag.NewFlagSet("user", flag.ExitOnError)
	adminFS := flag.NewFlagSet("admin", flag.ExitOnError)

	// 用户标志集
	var (
		username = userFS.String("username", "", "用户名")
		password = userFS.String("password", "", "密码")
		remember = userFS.Bool("remember", false, "记住密码")
	)

	// 管理员标志集
	var (
		adminUser = adminFS.String("admin", "", "管理员用户名")
		adminPass = adminFS.String("pass", "", "管理员密码")
		privLevel = adminFS.Int("level", 1, "权限级别")
	)

	// 模拟不同的命令场景
	fmt.Println("用户登录场景:")
	if err := userFS.Parse([]string{"-username", "zhangsan", "-password", "123456", "-remember"}); err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}
	fmt.Printf("用户名: %s, 记住密码: %t\n", *username, *remember)

	fmt.Println("\n管理员登录场景:")
	if err := adminFS.Parse([]string{"-admin", "admin", "-pass", "admin123", "-level", "5"}); err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}
	fmt.Printf("管理员: %s, 权限级别: %d\n", *adminUser, *privLevel)
}

// practicalExample 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1: 配置文件加载器
	demonstrateConfigLoader()

	// 示例2: 服务器启动参数
	demonstrateServerConfig()

	// 示例3: 数据库连接配置
	demonstrateDatabaseConfig()
}

// demonstrateConfigLoader 配置文件加载器
func demonstrateConfigLoader() {
	fmt.Println("1. 配置文件加载器:")

	fs := flag.NewFlagSet("config", flag.ExitOnError)

	var (
		configFile = fs.String("config", "config.json", "配置文件路径")
		env        = fs.String("env", "development", "运行环境")
		debug      = fs.Bool("debug", false, "调试模式")
		logLevel   = fs.String("log-level", "info", "日志级别")
		maxConn    = fs.Int("max-conn", 100, "最大连接数")
	)

	// 模拟配置文件参数
	configArgs := []string{
		"-config", "prod.json",
		"-env", "production",
		"-debug",
		"-log-level", "warn",
		"-max-conn", "1000",
	}

	if err := fs.Parse(configArgs); err != nil {
		fmt.Printf("配置解析错误: %v\n", err)
		return
	}

	fmt.Printf("配置文件: %s\n", *configFile)
	fmt.Printf("运行环境: %s\n", *env)
	fmt.Printf("调试模式: %t\n", *debug)
	fmt.Printf("日志级别: %s\n", *logLevel)
	fmt.Printf("最大连接数: %d\n", *maxConn)
}

// demonstrateServerConfig 服务器启动参数
func demonstrateServerConfig() {
	fmt.Println("\n2. 服务器启动参数:")

	fs := flag.NewFlagSet("server", flag.ExitOnError)

	var (
		host     = fs.String("host", "0.0.0.0", "服务器地址")
		port     = fs.Int("port", 8080, "服务器端口")
		https    = fs.Bool("https", false, "启用HTTPS")
		certFile = fs.String("cert", "", "SSL证书文件")
		keyFile  = fs.String("key", "", "SSL私钥文件")
		timeout  = fs.Duration("timeout", 30*time.Second, "请求超时")
	)

	// 模拟服务器启动参数
	serverArgs := []string{
		"-host", "localhost",
		"-port", "443",
		"-https",
		"-cert", "/path/to/cert.pem",
		"-key", "/path/to/key.pem",
		"-timeout", "1m",
	}

	if err := fs.Parse(serverArgs); err != nil {
		fmt.Printf("服务器配置解析错误: %v\n", err)
		return
	}

	scheme := "http"
	if *https {
		scheme = "https"
	}

	fmt.Printf("服务器地址: %s://%s:%d\n", scheme, *host, *port)
	if *https {
		fmt.Printf("SSL证书: %s\n", *certFile)
		fmt.Printf("SSL私钥: %s\n", *keyFile)
	}
	fmt.Printf("超时时间: %v\n", *timeout)
}

// demonstrateDatabaseConfig 数据库连接配置
func demonstrateDatabaseConfig() {
	fmt.Println("\n3. 数据库连接配置:")

	fs := flag.NewFlagSet("database", flag.ExitOnError)

	var (
		driver   = fs.String("driver", "mysql", "数据库驱动")
		host     = fs.String("db-host", "localhost", "数据库主机")
		port     = fs.Int("db-port", 3306, "数据库端口")
		database = fs.String("db-name", "myapp", "数据库名称")
		username = fs.String("db-user", "root", "数据库用户名")
		password = fs.String("db-pass", "", "数据库密码")
		sslMode  = fs.String("ssl-mode", "disable", "SSL模式")
	)

	// 模拟数据库连接参数
	dbArgs := []string{
		"-driver", "postgres",
		"-db-host", "db.example.com",
		"-db-port", "5432",
		"-db-name", "production_db",
		"-db-user", "app_user",
		"-db-pass", "secure_password",
		"-ssl-mode", "require",
	}

	if err := fs.Parse(dbArgs); err != nil {
		fmt.Printf("数据库配置解析错误: %v\n", err)
		return
	}

	fmt.Printf("数据库驱动: %s\n", *driver)
	fmt.Printf("连接地址: %s@%s:%d/%s\n", *username, *host, *port, *database)
	fmt.Printf("SSL模式: %s\n", *sslMode)

	// 构建连接字符串
	var connStr string
	switch *driver {
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
			*username, *password, *host, *port, *database)
	case "postgres":
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			*host, *port, *username, *password, *database, *sslMode)
	}

	fmt.Printf("连接字符串: %s\n", connStr)
}

// 自定义标志类型示例
type customType string

func (c *customType) String() string {
	return string(*c)
}

func (c *customType) Set(value string) error {
	if value != "valid" && value != "invalid" {
		return fmt.Errorf("必须是 'valid' 或 'invalid'")
	}
	*c = customType(value)
	return nil
}

// 演示自定义标志类型
func demonstrateCustomFlag() {
	fmt.Println("\n--- 自定义标志类型 ---")

	var status customType
	flag.Var(&status, "status", "状态 (valid/invalid)")

	// 这里可以设置自定义标志
	// status.Set("valid")

	fmt.Printf("状态: %s\n", status)
}