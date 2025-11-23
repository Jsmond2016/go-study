# flag 包示例

flag 包提供了命令行参数解析功能，是 Go 应用程序中处理用户输入的标准工具。本示例展示了 flag 包的各种用法和高级特性。

## 运行示例

```bash
# 基本用法
go run main.go

# 带参数运行
go run main.go -name="张三" -age=25 -married=true -height=1.75

# 帮助信息
go run main.go -h
```

## 知识点

### 1. 基本用法

```go
import "flag"

// 定义标志变量
var name = flag.String("name", "Guest", "用户姓名")
var age = flag.Int("age", 0, "用户年龄")
var married = flag.Bool("married", false, "是否已婚")

// 解析命令行参数
func main() {
    flag.Parse()
    fmt.Printf("姓名: %s, 年龄: %d\n", *name, *age)
}
```

### 2. 标志类型

flag 包支持多种数据类型：

```go
// 基本类型
flag.String("str", "default", "字符串")
flag.Int("int", 0, "整数")
flag.Bool("bool", false, "布尔值")
flag.Float64("float", 0.0, "浮点数")

// 时间相关
flag.Duration("timeout", time.Second, "持续时间")

// 自定义类型
flag.Var(&customValue, "custom", "自定义标志")
```

### 3. 标志集合 (FlagSet)

```go
// 创建新的标志集
userFS := flag.NewFlagSet("user", flag.ExitOnError)
adminFS := flag.NewFlagSet("admin", flag.ExitOnError)

// 在不同标志集中定义不同的标志
var username = userFS.String("username", "", "用户名")
var adminUser = adminFS.String("admin", "", "管理员")

// 分别解析
userFS.Parse(userArgs)
adminFS.Parse(adminArgs)
```

### 4. 非选项参数

```go
flag.Parse() // 解析后

// 获取非选项参数
args := flag.Args()     // []string
argCount := flag.NArg() // int
```

### 5. 自定义帮助信息

```go
flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "用法: %s [选项]\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "程序描述\n\n")
    fmt.Fprintf(os.Stderr, "选项:\n")
    flag.PrintDefaults()
}
```

### 6. 自定义标志类型

```go
type customType string

func (c *customType) String() string {
    return string(*c)
}

func (c *customType) Set(value string) error {
    // 验证和设置值
    *c = customType(value)
    return nil
}

// 使用自定义标志
var status customType
flag.Var(&status, "status", "自定义状态")
```

## 实际应用场景

### 1. 配置管理

```go
type Config struct {
    ConfigFile string
    Env        string
    Debug      bool
    LogLevel   string
    MaxConn    int
}

func parseConfig() *Config {
    fs := flag.NewFlagSet("config", flag.ExitOnError)

    c := &Config{}
    fs.StringVar(&c.ConfigFile, "config", "config.json", "配置文件")
    fs.StringVar(&c.Env, "env", "development", "运行环境")
    fs.BoolVar(&c.Debug, "debug", false, "调试模式")
    fs.StringVar(&c.LogLevel, "log-level", "info", "日志级别")
    fs.IntVar(&c.MaxConn, "max-conn", 100, "最大连接数")

    fs.Parse(os.Args[1:])
    return c
}
```

### 2. 服务器配置

```go
func parseServerConfig() {
    var (
        host    = flag.String("host", "0.0.0.0", "服务器地址")
        port    = flag.Int("port", 8080, "服务器端口")
        https   = flag.Bool("https", false, "启用HTTPS")
        certFile = flag.String("cert", "", "SSL证书")
        keyFile = flag.String("key", "", "SSL私钥")
    )

    flag.Parse()

    // 启动服务器
    if *https {
        startHTTPS(*host, *port, *certFile, *keyFile)
    } else {
        startHTTP(*host, *port)
    }
}
```

### 3. 子命令支持

```go
func main() {
    if len(os.Args) < 2 {
        fmt.Println("需要指定子命令")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "start":
        startCommand(os.Args[2:])
    case "stop":
        stopCommand(os.Args[2:])
    case "status":
        statusCommand(os.Args[2:])
    default:
        fmt.Printf("未知命令: %s\n", os.Args[1])
    }
}

func startCommand(args []string) {
    fs := flag.NewFlagSet("start", flag.ExitOnError)
    var daemon = fs.Bool("d", false, "后台运行")
    fs.Parse(args)
    // 启动逻辑
}
```

## 最佳实践

### 1. 标志命名规范

- 使用短横线分隔：`--max-conn` 而不是 `--max_conn`
- 保持一致性：同一项目中使用相同的命名风格
- 避免缩写：使用 `--timeout` 而不是 `--to`

### 2. 默认值设置

- 为所有标志提供合理的默认值
- 考虑不同环境的默认值
- 敏感信息不要设置默认值

### 3. 帮助信息

- 提供清晰简洁的帮助文本
- 包含使用示例
- 说明必需参数和可选参数

### 4. 错误处理

- 使用合适的 FlagSet 错误处理策略
- 验证参数值的有效性
- 提供友好的错误信息

## 高级技巧

### 1. 环境变量支持

```go
func stringEnvFlag(name, envName, defaultValue, usage string) *string {
    value := os.Getenv(envName)
    if value == "" {
        value = defaultValue
    }
    return flag.String(name, value, fmt.Sprintf("%s (环境变量: %s)", usage, envName))
}
```

### 2. 配置文件集成

```go
func loadConfigFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // 解析配置文件
    return json.NewDecoder(file).Decode(&config)
}
```

### 3. 参数验证

```go
func validateFlags() error {
    if *port < 1 || *port > 65535 {
        return fmt.Errorf("端口必须在 1-65535 范围内")
    }
    if *timeout <= 0 {
        return fmt.Errorf("超时时间必须大于0")
    }
    return nil
}
```

## 常见陷阱

### 1. 忘记调用 flag.Parse()

```go
// 错误 - 忘记解析
var name = flag.String("name", "", "姓名")
func main() {
    fmt.Printf("姓名: %s\n", *name) // 总是显示默认值
}

// 正确
func main() {
    flag.Parse()
    fmt.Printf("姓名: %s\n", *name)
}
```

### 2. 标志名冲突

```go
// 避免在不同的 FlagSet 中使用相同的标志名
```

### 3. 指针解引用

```go
var name = flag.String("name", "", "姓名")
fmt.Printf("姓名: %s\n", name)   // 错误 - 打印指针地址
fmt.Printf("姓名: %s\n", *name)  // 正确 - 解引用指针
```

## 替代方案

对于复杂的命令行应用，可以考虑：

1. **cobra** - 功能强大的 CLI 框架
2. **urfave/cli** - 简单易用的 CLI 库
3. **viper** - 配置管理库，与 flag 集成良好

## 练习

1. 创建一个支持多个子命令的 CLI 工具
2. 实现一个配置文件和命令行参数的优先级系统
3. 创建一个带有参数验证的服务启动器
4. 实现一个带有环境变量支持的配置解析器
5. 创建一个自定义的标志类型