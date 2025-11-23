---
title: Go Modules
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["基础语法", "包管理"]
tags: ["Go Modules", "依赖管理", "版本控制"]
---

# Go Modules

Go Modules 是 Go 语言的官方依赖管理系统，从 Go 1.11 开始引入，Go 1.13 成为默认选项。

## 📋 学习目标

- [ ] 理解 Go Modules 的概念
- [ ] 掌握模块的创建和管理
- [ ] 学会依赖的添加和更新
- [ ] 理解版本控制和语义化版本
- [ ] 掌握私有模块的使用
- [ ] 了解常见问题和解决方案

## 🎯 Go Modules 简介

### 为什么使用 Go Modules

- **官方支持**: Go 官方依赖管理方案
- **版本控制**: 支持语义化版本
- **依赖隔离**: 每个项目独立的依赖
- **可重现构建**: go.sum 确保依赖一致性
- **代理支持**: 支持模块代理加速

### 基本概念

- **模块 (Module)**: 一个包含 go.mod 文件的目录树
- **模块路径 (Module Path)**: 模块的唯一标识符
- **版本 (Version)**: 遵循语义化版本规范
- **go.mod**: 模块定义文件
- **go.sum**: 依赖校验文件

## 🚀 快速开始

### 初始化模块

```bash
# 创建新模块
go mod init github.com/username/project

# 在现有项目中初始化
cd existing-project
go mod init
```

### go.mod 文件

```go
module github.com/username/project

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	gorm.io/gorm v1.25.5
)

require (
	github.com/bytedance/sonic v1.9.1 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
)
```

## 📦 依赖管理

### 添加依赖

```bash
# 添加依赖（自动下载）
go get github.com/gin-gonic/gin

# 添加指定版本
go get github.com/gin-gonic/gin@v1.9.1

# 添加最新版本
go get github.com/gin-gonic/gin@latest

# 添加特定提交
go get github.com/gin-gonic/gin@abc1234
```

### 更新依赖

```bash
# 更新所有依赖
go get -u ./...

# 更新指定依赖
go get -u github.com/gin-gonic/gin

# 更新到最新版本
go get -u github.com/gin-gonic/gin@latest

# 更新到主版本
go get github.com/gin-gonic/gin/v2
```

### 删除依赖

```bash
# 删除未使用的依赖
go mod tidy

# 手动删除（然后运行 go mod tidy）
# 编辑 go.mod 文件，删除不需要的 require 行
```

## 🔢 版本控制

### 语义化版本

版本格式：`v主版本号.次版本号.修订号`

- **主版本号**: 不兼容的 API 修改
- **次版本号**: 向下兼容的功能性新增
- **修订号**: 向下兼容的问题修正

示例：
- `v1.0.0` - 初始版本
- `v1.1.0` - 新增功能
- `v1.1.1` - 修复 bug
- `v2.0.0` - 重大更新（不兼容）

### 版本选择

```go
// go.mod
require (
	github.com/gin-gonic/gin v1.9.1  // 精确版本
	github.com/some/package v1.2.3   // 精确版本
)

// Go 会自动选择满足要求的最新版本
require github.com/gin-gonic/gin v1.9.0  // 会选择 >= v1.9.0 的最新版本
```

### 版本替换

```go
// go.mod
replace github.com/old/package => github.com/new/package v1.0.0

// 使用本地路径
replace github.com/local/package => ../local-package

// 使用相对路径
replace github.com/local/package => ./local-package
```

## 🔧 常用命令

### go mod init

```bash
# 初始化模块
go mod init module-path

# 示例
go mod init github.com/username/project
```

### go mod tidy

```bash
# 整理依赖
go mod tidy

# 功能：
# - 添加缺失的依赖
# - 删除未使用的依赖
# - 更新 go.mod 和 go.sum
```

### go mod download

```bash
# 下载依赖到本地缓存
go mod download

# 下载指定模块
go mod download github.com/gin-gonic/gin
```

### go mod verify

```bash
# 验证依赖的完整性
go mod verify
```

### go mod graph

```bash
# 查看依赖关系图
go mod graph
```

### go mod why

```bash
# 查看为什么需要某个依赖
go mod why github.com/gin-gonic/gin
```

## 🌐 私有模块

### 配置 GOPRIVATE

```bash
# 设置私有模块（不走代理）
go env -w GOPRIVATE=github.com/company/*,gitlab.com/company/*

# 查看当前配置
go env GOPRIVATE
```

### 配置 Git 认证

```bash
# 配置 Git 凭据
git config --global url."https://username:token@github.com/company/".insteadOf "https://github.com/company/"
```

### 使用 .netrc

```bash
# ~/.netrc
machine github.com
login username
password token
```

## 🏗️ Go Workspace

### 创建 Workspace

```bash
# 创建 workspace
go work init ./module1 ./module2

# 添加模块到 workspace
go work use ./module3
```

### go.work 文件

```go
go 1.21

use (
	./module1
	./module2
	../external-module
)

replace github.com/old/package => github.com/new/package v1.0.0
```

## 🏃‍♂️ 实践应用

### 完整的项目结构

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   ├── models/
│   └── handlers/
└── pkg/
    └── utils/
```

### go.mod 示例

```go
module github.com/username/myproject

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	gorm.io/gorm v1.25.5
	gorm.io/driver/mysql v1.5.2
	github.com/spf13/viper v1.16.0
	go.uber.org/zap v1.25.0
)

require (
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	// ... 更多 indirect 依赖
)
```

### 依赖管理最佳实践

```bash
# 1. 初始化模块
go mod init github.com/username/project

# 2. 添加依赖
go get github.com/gin-gonic/gin

# 3. 整理依赖
go mod tidy

# 4. 验证依赖
go mod verify

# 5. 查看依赖图
go mod graph | grep gin
```

## ⚠️ 常见问题

### 1. 依赖下载失败

```bash
# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct

# 或使用多个代理
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

### 2. 版本冲突

```bash
# 查看依赖树
go mod graph

# 更新冲突的依赖
go get -u package@version

# 使用 replace 替换
go mod edit -replace old/package=new/package@version
```

### 3. 私有模块访问

```bash
# 配置私有模块
go env -w GOPRIVATE=github.com/company/*

# 配置 Git 认证
git config --global url."https://token@github.com/company/".insteadOf "https://github.com/company/"
```

### 4. 清理缓存

```bash
# 清理模块缓存
go clean -modcache

# 清理构建缓存
go clean -cache
```

## 📚 扩展阅读

- [Go Modules 官方文档](https://go.dev/ref/mod)
- [语义化版本规范](https://semver.org/)
- [Go Modules Wiki](https://github.com/golang/go/wiki/Modules)

## ⏭️ 下一章节

[Viper 配置管理](./04-viper.md) → 学习配置管理

---

**💡 提示**: Go Modules 是 Go 开发的基础，掌握它对于管理项目依赖非常重要！

