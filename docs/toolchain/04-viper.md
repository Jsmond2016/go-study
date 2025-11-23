---
title: Viper 配置管理
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["基础语法", "文件操作"]
tags: ["Viper", "配置管理", "环境变量"]
---

# Viper 配置管理

Viper 是 Go 语言最流行的配置管理库，支持多种配置格式和环境变量。

## 📋 学习目标

- [ ] 理解配置管理的重要性
- [ ] 掌握 Viper 的基本使用
- [ ] 学会读取多种配置格式
- [ ] 理解环境变量和命令行参数
- [ ] 掌握配置热加载
- [ ] 了解配置最佳实践

## 🎯 Viper 简介

### 为什么选择 Viper

- **多格式支持**: JSON、YAML、TOML、环境变量等
- **优先级管理**: 支持配置优先级
- **热加载**: 支持配置文件变更监听
- **类型安全**: 支持类型转换和验证
- **易于使用**: API 简洁直观

### 安装 Viper

```bash
go get github.com/spf13/viper
```

## 🚀 快速开始

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 设置配置文件名（不含扩展名）
	viper.SetConfigName("config")
	
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	
	// 添加配置文件搜索路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.myapp")
	
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %w", err))
	}
	
	// 获取配置值
	port := viper.GetInt("server.port")
	fmt.Printf("服务器端口: %d\n", port)
}
```

## 📝 配置文件格式

### YAML 配置

```yaml
# config.yaml
server:
  port: 8080
  host: localhost
  timeout: 30s

database:
  host: localhost
  port: 3306
  username: root
  password: password
  database: mydb
  max_connections: 100

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

app:
  name: "My Application"
  version: "1.0.0"
  debug: true
```

### JSON 配置

```json
{
  "server": {
    "port": 8080,
    "host": "localhost",
    "timeout": "30s"
  },
  "database": {
    "host": "localhost",
    "port": 3306,
    "username": "root",
    "password": "password",
    "database": "mydb"
  }
}
```

### TOML 配置

```toml
[server]
port = 8080
host = "localhost"
timeout = "30s"

[database]
host = "localhost"
port = 3306
username = "root"
password = "password"
database = "mydb"
```

## 🔧 读取配置

### 基本读取

```go
// 读取字符串
name := viper.GetString("app.name")

// 读取整数
port := viper.GetInt("server.port")

// 读取布尔值
debug := viper.GetBool("app.debug")

// 读取时间
timeout := viper.GetDuration("server.timeout")

// 读取字符串切片
tags := viper.GetStringSlice("app.tags")

// 读取映射
database := viper.GetStringMap("database")
```

### 嵌套配置

```go
// 读取嵌套配置
port := viper.GetInt("server.port")
host := viper.GetString("server.host")

// 或使用子配置
serverConfig := viper.Sub("server")
if serverConfig != nil {
	port := serverConfig.GetInt("port")
	host := serverConfig.GetString("host")
}
```

### 默认值

```go
// 设置默认值
viper.SetDefault("server.port", 8080)
viper.SetDefault("server.host", "localhost")
viper.SetDefault("app.debug", false)

// 读取配置（如果不存在则使用默认值）
port := viper.GetInt("server.port")
```

## 🌍 环境变量

### 自动绑定环境变量

```go
// 自动读取环境变量
viper.AutomaticEnv()

// 读取环境变量（优先级高于配置文件）
port := viper.GetInt("SERVER_PORT")
```

### 手动绑定环境变量

```go
// 绑定环境变量
viper.BindEnv("server.port", "SERVER_PORT")
viper.BindEnv("server.host", "SERVER_HOST")
viper.BindEnv("database.password", "DB_PASSWORD")

// 读取
port := viper.GetInt("server.port")
```

### 环境变量前缀

```go
// 设置环境变量前缀
viper.SetEnvPrefix("MYAPP")
viper.AutomaticEnv()

// 读取 MYAPP_SERVER_PORT
port := viper.GetInt("server.port")
```

## 🎛️ 命令行参数

### 绑定命令行标志

```go
import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// 绑定命令行参数
	viper.BindPFlag("server.port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.host", rootCmd.Flags().Lookup("host"))
}

var rootCmd = &cobra.Command{
	Use: "myapp",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("server.port")
		fmt.Printf("端口: %d\n", port)
	},
}

func init() {
	rootCmd.Flags().IntP("port", "p", 8080, "服务器端口")
	rootCmd.Flags().StringP("host", "H", "localhost", "服务器地址")
}
```

## 🔄 配置热加载

### 监听配置文件

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	
	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已更新:", e.Name)
		// 重新读取配置
		port := viper.GetInt("server.port")
		fmt.Printf("新端口: %d\n", port)
	})
	
	// 保持程序运行
	select {}
}
```

## 🏃‍♂️ 实践应用

### 完整的配置管理示例

```go
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	App      AppConfig      `mapstructure:"app"`
}

type ServerConfig struct {
	Port    int    `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	Timeout string `mapstructure:"timeout"`
}

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxConns     int    `mapstructure:"max_connections"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Debug   bool   `mapstructure:"debug"`
}

var AppConfigInstance *Config

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("$HOME/.myapp")
	}
	
	// 设置默认值
	setDefaults()
	
	// 读取环境变量
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()
	
	// 绑定环境变量
	bindEnvVars()
	
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("读取配置文件失败: %w", err)
		}
	}
	
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	
	AppConfigInstance = &config
	return &config, nil
}

func setDefaults() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.timeout", "30s")
	
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.max_connections", 100)
	
	viper.SetDefault("app.debug", false)
}

func bindEnvVars() {
	viper.BindEnv("server.port", "MYAPP_SERVER_PORT")
	viper.BindEnv("server.host", "MYAPP_SERVER_HOST")
	viper.BindEnv("database.host", "MYAPP_DB_HOST")
	viper.BindEnv("database.password", "MYAPP_DB_PASSWORD")
	viper.BindEnv("app.debug", "MYAPP_DEBUG")
}

// 使用示例
func main() {
	config, err := LoadConfig("")
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("服务器: %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("数据库: %s:%d\n", config.Database.Host, config.Database.Port)
}
```

## ⚠️ 注意事项

### 1. 配置优先级

```go
// 优先级从高到低：
// 1. 显式调用 Set
// 2. 命令行标志
// 3. 环境变量
// 4. 配置文件
// 5. 默认值
```

### 2. 类型转换

```go
// Viper 会自动进行类型转换
port := viper.GetInt("server.port")      // 字符串 "8080" 转换为 int
timeout := viper.GetDuration("timeout")   // 字符串 "30s" 转换为 time.Duration
```

### 3. 配置验证

```go
// 验证必需的配置项
requiredKeys := []string{"server.port", "database.host"}
for _, key := range requiredKeys {
	if !viper.IsSet(key) {
		return fmt.Errorf("缺少必需的配置项: %s", key)
	}
}
```

## 📚 扩展阅读

- [Viper 官方文档](https://github.com/spf13/viper)
- [Viper 示例](https://github.com/spf13/viper/tree/master/examples)

## ⏭️ 下一章节

[Zap 日志库](./05-zap.md) → 学习结构化日志

---

**💡 提示**: Viper 是 Go 项目中最常用的配置管理库，掌握它可以让配置管理变得简单高效！

