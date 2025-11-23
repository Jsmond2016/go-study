---
title: Zap 日志库
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["基础语法", "接口"]
tags: ["Zap", "日志", "结构化日志", "性能"]
---

# Zap 日志库

Zap 是 Uber 开源的高性能结构化日志库，专为 Go 语言设计，性能优异。

## 📋 学习目标

- [ ] 理解结构化日志的概念
- [ ] 掌握 Zap 的基本使用
- [ ] 学会配置日志级别和输出
- [ ] 理解结构化字段的使用
- [ ] 掌握日志采样和性能优化
- [ ] 了解日志最佳实践

## 🎯 Zap 简介

### 为什么选择 Zap

- **高性能**: 零分配设计，性能优异
- **结构化日志**: 支持结构化字段
- **日志级别**: 支持多种日志级别
- **灵活配置**: 支持多种输出格式
- **生产就绪**: 适合生产环境使用

### 安装 Zap

```bash
go get go.uber.org/zap
```

## 🚀 快速开始

### 基本使用

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	// 使用默认配置
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	logger.Info("这是一条信息日志")
	logger.Warn("这是一条警告日志")
	logger.Error("这是一条错误日志")
}
```

### 开发模式

```go
// 开发模式（更详细的输出）
logger, _ := zap.NewDevelopment()
defer logger.Sync()

logger.Info("开发模式日志")
```

### 生产模式

```go
// 生产模式（JSON 格式，性能优化）
logger, _ := zap.NewProduction()
defer logger.Sync()

logger.Info("生产模式日志")
```

## 📝 日志级别

### 基本级别

```go
logger.Debug("调试信息")
logger.Info("一般信息")
logger.Warn("警告信息")
logger.Error("错误信息")
logger.DPanic("开发环境 panic")
logger.Panic("panic")
logger.Fatal("致命错误，程序退出")
```

### 自定义级别

```go
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 创建自定义配置
config := zap.NewProductionConfig()
config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel) // 设置最低日志级别

logger, _ := config.Build()
defer logger.Sync()
```

## 🔧 结构化字段

### 基本字段

```go
logger.Info("用户登录",
	zap.String("username", "zhangsan"),
	zap.Int("user_id", 123),
	zap.Bool("success", true),
)

logger.Error("数据库连接失败",
	zap.String("host", "localhost"),
	zap.Int("port", 3306),
	zap.Error(err),
)
```

### 常用字段类型

```go
// 字符串
zap.String("key", "value")

// 整数
zap.Int("age", 25)
zap.Int64("timestamp", time.Now().Unix())

// 布尔值
zap.Bool("enabled", true)

// 浮点数
zap.Float64("score", 95.5)

// 时间
zap.Time("created_at", time.Now())
zap.Duration("latency", time.Second)

// 错误
zap.Error(err)

// 对象
zap.Any("user", user)

// 字节数组
zap.Binary("data", []byte("binary"))
```

### 字段组合

```go
// 使用 Fields
logger.Info("请求处理",
	zap.String("method", "GET"),
	zap.String("path", "/api/users"),
	zap.Int("status", 200),
	zap.Duration("latency", time.Millisecond*150),
)

// 使用结构化对象
type Request struct {
	Method  string
	Path    string
	Status  int
	Latency time.Duration
}

req := Request{
	Method:  "GET",
	Path:    "/api/users",
	Status:  200,
	Latency: time.Millisecond * 150,
}

logger.Info("请求处理", zap.Any("request", req))
```

## 🎨 自定义配置

### 完整配置示例

```go
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() (*zap.Logger, error) {
	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	
	// 配置日志级别
	atom := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	
	// 配置输出
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		atom,
	)
	
	// 创建 logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	
	return logger, nil
}

func main() {
	logger, _ := NewLogger()
	defer logger.Sync()
	
	logger.Info("自定义配置的日志")
}
```

### 多输出配置

```go
func NewMultiOutputLogger() (*zap.Logger, error) {
	// 控制台输出（开发环境）
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	
	// 文件输出（生产环境）
	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.InfoLevel)
	
	// 组合多个 core
	core := zapcore.NewTee(consoleCore, fileCore)
	
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}
```

## 🏃‍♂️ 实践应用

### 在 Web 应用中使用

```go
package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		
		c.Next()
		
		latency := time.Since(start)
		
		logger.Info("HTTP 请求",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		)
	}
}

func main() {
	r := gin.New()
	r.Use(ZapLogger())
	
	r.GET("/", func(c *gin.Context) {
		logger.Info("处理请求", zap.String("handler", "index"))
		c.JSON(200, gin.H{"message": "Hello"})
	})
	
	r.Run(":8080")
}
```

### 日志服务封装

```go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(level zapcore.Level) (*Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	
	return &Logger{Logger: logger}, nil
}

func (l *Logger) LogRequest(method, path string, status int, latency time.Duration) {
	l.Info("HTTP 请求",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status", status),
		zap.Duration("latency", latency),
	)
}

func (l *Logger) LogError(err error, context ...zap.Field) {
	l.Error("错误发生",
		append([]zap.Field{zap.Error(err)}, context...)...,
	)
}

func (l *Logger) LogInfo(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}
```

## ⚡ 性能优化

### 日志采样

```go
import "go.uber.org/zap/zapcore"

// 配置采样
config := zap.NewProductionConfig()
config.Sampling = &zap.SamplingConfig{
	Initial:    100, // 初始采样率
	Thereafter: 100, // 后续采样率
}

logger, _ := config.Build()
```

### 避免不必要的日志

```go
// ✅ 使用条件检查
if logger.Core().Enabled(zapcore.DebugLevel) {
	logger.Debug("调试信息", expensiveFields()...)
}

// ❌ 避免直接调用（即使被禁用也会执行函数）
logger.Debug("调试信息", expensiveFields()...)
```

## ⚠️ 注意事项

### 1. 同步日志

```go
// ✅ 程序退出前同步日志
defer logger.Sync()

// 或手动同步
logger.Sync()
```

### 2. 错误处理

```go
// ✅ 检查日志初始化错误
logger, err := zap.NewProduction()
if err != nil {
	panic(err)
}
```

### 3. 性能考虑

```go
// ✅ 使用结构化字段而不是字符串拼接
logger.Info("用户登录", zap.String("username", username))

// ❌ 避免字符串拼接
logger.Info(fmt.Sprintf("用户登录: %s", username))
```

## 📚 扩展阅读

- [Zap 官方文档](https://github.com/uber-go/zap)
- [Zap 性能基准](https://github.com/uber-go/zap#performance)

## ⏭️ 下一章节

[JWT 鉴权](./06-jwt.md) → 学习 JWT 认证实现

---

**💡 提示**: Zap 是生产环境中最常用的日志库，掌握结构化日志可以让日志分析更加高效！

