---
title: 环境搭建
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["Web 开发", "数据库", "工具链"]
tags: ["环境", "搭建", "配置", "初始化"]
---

# 环境搭建

本章节将指导你搭建博客系统的开发环境，包括项目初始化、依赖安装和配置管理。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 创建项目目录结构
- [ ] 初始化 Go 模块
- [ ] 安装项目依赖
- [ ] 配置数据库连接
- [ ] 配置应用参数
- [ ] 运行项目并验证

## 🚀 快速开始

### 1. 创建项目目录

```bash
# 创建项目根目录
mkdir -p blog-system
cd blog-system

# 创建标准 Go 项目目录结构
mkdir -p cmd/server
mkdir -p internal/{handler,service,repository,model}
mkdir -p pkg/{auth,upload,utils}
mkdir -p config
mkdir -p uploads
mkdir -p scripts
```

### 2. 初始化 Go 模块

```bash
# 初始化 Go 模块
go mod init blog-system

# 创建主程序文件
touch cmd/server/main.go
```

### 3. 安装依赖

```bash
# 安装核心依赖
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/golang-jwt/jwt/v5
go get github.com/spf13/viper
go get github.com/go-playground/validator/v10
go get golang.org/x/crypto
```

## 📁 项目结构详解

### 目录结构

```
blog-system/
├── cmd/
│   └── server/
│       └── main.go          # 应用入口
├── internal/                # 内部代码（不对外暴露）
│   ├── handler/             # HTTP 处理器
│   │   ├── article.go
│   │   ├── comment.go
│   │   ├── user.go
│   │   └── upload.go
│   ├── service/             # 业务逻辑层
│   │   ├── article.go
│   │   ├── comment.go
│   │   └── user.go
│   ├── repository/          # 数据访问层
│   │   ├── article.go
│   │   ├── comment.go
│   │   └── user.go
│   └── model/               # 数据模型
│       ├── article.go
│       ├── comment.go
│       ├── user.go
│       ├── category.go
│       └── tag.go
├── pkg/                     # 可复用的包
│   ├── auth/                # 认证相关
│   │   ├── jwt.go
│   │   └── middleware.go
│   ├── upload/              # 文件上传
│   │   └── upload.go
│   └── utils/               # 工具函数
│       ├── response.go
│       └── validator.go
├── config/                   # 配置文件
│   ├── config.yaml
│   └── config.example.yaml
├── uploads/                  # 上传文件目录
├── scripts/                  # 脚本文件
│   └── init_db.sql
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

### 目录说明

- **cmd/**: 应用程序的入口点
- **internal/**: 私有应用程序代码，不对外暴露
- **pkg/**: 可被外部应用程序使用的库代码
- **config/**: 配置文件目录
- **uploads/**: 用户上传的文件存储目录

## ⚙️ 配置文件

### 创建配置文件

创建 `config/config.yaml`:

```yaml
server:
  port: 8080
  mode: debug  # debug, release, test
  host: localhost

database:
  driver: mysql  # mysql, postgres
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: blog_db
  charset: utf8mb4
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 3600s

jwt:
  secret: your-secret-key-change-in-production
  expire: 24h
  issuer: blog-system

upload:
  path: ./uploads
  max_size: 10485760  # 10MB
  allowed_types:
    - image/jpeg
    - image/png
    - image/gif
    - image/webp

log:
  level: info  # debug, info, warn, error
  format: json  # json, text
  output: stdout  # stdout, file
  file_path: ./logs/app.log
```

### 创建配置示例文件

创建 `config/config.example.yaml`（用于版本控制，不包含敏感信息）:

```yaml
server:
  port: 8080
  mode: debug
  host: localhost

database:
  driver: mysql
  host: localhost
  port: 3306
  user: root
  password: CHANGE_ME
  dbname: blog_db
  charset: utf8mb4

jwt:
  secret: CHANGE_ME
  expire: 24h

upload:
  path: ./uploads
  max_size: 10485760
```

## 🔧 配置管理

### 使用 Viper 加载配置

创建 `pkg/config/config.go`:

```go
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Upload   UploadConfig
	Log      LogConfig
}

type ServerConfig struct {
	Port string
	Mode string
	Host string
}

type DatabaseConfig struct {
	Driver            string
	Host             string
	Port             int
	User             string
	Password         string
	DBName           string
	Charset          string
	MaxIdleConns     int
	MaxOpenConns     int
	ConnMaxLifetime  int
}

type JWTConfig struct {
	Secret string
	Expire string
	Issuer string
}

type UploadConfig struct {
	Path        string
	MaxSize     int64
	AllowedTypes []string
}

type LogConfig struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

var AppConfig *Config

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置默认值
	setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 从环境变量读取（可选）
	viper.AutomaticEnv()

	// 解析配置
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	AppConfig = config
	return config, nil
}

func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("database.max_idle_conns", 10)
	viper.SetDefault("database.max_open_conns", 100)
}
```

## 🗄️ 数据库初始化

### 创建数据库

```sql
-- 创建数据库
CREATE DATABASE IF NOT EXISTS blog_db 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE blog_db;
```

### 数据库连接

创建 `internal/repository/database.go`:

```go
package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	// 获取底层 sql.DB 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库实例失败: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

func BuildDSN(config DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Charset,
	)
}
```

## 🚀 主程序初始化

### 创建主程序

创建 `cmd/server/main.go`:

```go
package main

import (
	"fmt"
	"log"
	"blog-system/internal/repository"
	"blog-system/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("加载配置失败:", err)
	}

	// 初始化数据库
	dsn := repository.BuildDSN(cfg.Database)
	db, err := repository.InitDB(dsn)
	if err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 创建 Gin 引擎
	r := gin.Default()

	// 设置路由（后续章节实现）
	setupRoutes(r, db)

	// 启动服务器
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("服务器启动在 http://%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "服务运行正常",
		})
	})
}
```

## ✅ 验证安装

### 运行项目

```bash
# 运行项目
go run cmd/server/main.go

# 或者构建后运行
go build -o blog-system cmd/server/main.go
./blog-system
```

### 测试健康检查

```bash
curl http://localhost:8080/health
```

预期响应：
```json
{
  "status": "ok",
  "message": "服务运行正常"
}
```

## 🔍 常见问题

### 1. 数据库连接失败

**问题**: 无法连接到数据库

**解决方案**:
- 检查数据库服务是否启动
- 验证配置文件中的数据库连接信息
- 确认数据库用户权限

### 2. 配置文件找不到

**问题**: `config file not found`

**解决方案**:
- 确认配置文件路径正确
- 检查配置文件名称是否为 `config.yaml`
- 确认配置文件在 `config/` 目录下

### 3. 依赖安装失败

**问题**: `go get` 失败

**解决方案**:
- 检查网络连接
- 设置 Go 代理: `go env -w GOPROXY=https://goproxy.cn,direct`
- 清理模块缓存: `go clean -modcache`

## 📝 下一步

环境搭建完成后，下一步是：

- [数据模型设计](./02-models.md) - 设计数据库表和模型

---

**🎉 环境搭建完成！** 现在你可以开始设计数据模型了。

