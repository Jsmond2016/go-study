---
title: 快速开始
difficulty: beginner
duration: "30分钟"
prerequisites: []
tags: ["入门", "环境配置", "快速开始"]
---

# 快速开始

欢迎来到 Go 语言学习世界！本指南将帮助你快速搭建 Go 开发环境并开始你的第一个 Go 程序。

## 📋 学习目标

- [ ] 安装和配置 Go 开发环境
- [ ] 理解 Go 的工作目录结构
- [ ] 编写并运行第一个 Go 程序
- [ ] 了解 Go 的基本命令
- [ ] 配置代码编辑器

## 🛠️ 环境要求

### 系统要求
- **操作系统**: Windows 10+, macOS 10.15+, Linux (大多数发行版)
- **内存**: 至少 4GB RAM
- **存储**: 至少 1GB 可用空间
- **网络**: 用于下载 Go 和依赖包

### Go 版本
推荐使用 **Go 1.21+** 版本以获得最佳体验和新特性支持。

## 📥 安装 Go

### 方法一：官方安装包（推荐）

1. **访问官网下载**
   - 打开 [Go 官方下载页面](https://golang.org/dl/)
   - 选择适合你操作系统的版本

2. **安装步骤**
   
   **Windows:**
   ```bash
   # 下载 .msi 安装包
   # 双击运行，按照提示完成安装
   ```
   
   **macOS:**
   ```bash
   # 使用 Homebrew（推荐）
   brew install go
   
   # 或下载 .pkg 安装包
   ```
   
   **Linux:**
   ```bash
   # 下载 tar.gz 包
   wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
   
   # 解压到 /usr/local
   sudo rm -rf /usr/local/go
   sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
   ```

3. **配置环境变量**
   
   添加到 `~/.bashrc`、`~/.zshrc` 或系统环境变量：
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go
   export GO111MODULE=on
   ```

### 方法二：使用包管理器

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang-go
```

**CentOS/RHEL:**
```bash
sudo yum install golang
# 或使用 dnf（较新版本）
sudo dnf install golang
```

## ✅ 验证安装

打开终端，运行以下命令：

```bash
go version
```

你应该看到类似输出：
```
go version go1.21.0 darwin/arm64
```

检查 Go 环境信息：
```bash
go env
```

## 💻 第一个 Go 程序

### 1. 创建工作目录

```bash
mkdir hello-go
cd hello-go
```

### 2. 初始化模块

```bash
go mod init hello-go
```

### 3. 创建主程序文件

创建 `main.go` 文件：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go 世界!")
	fmt.Println("欢迎学习 Go 语言！")
}
```

### 4. 运行程序

```bash
# 直接运行
go run main.go

# 编译后运行
go build
./hello-go  # Windows: hello-go.exe
```

预期输出：
```
Hello, Go 世界!
欢迎学习 Go 语言！
```

## 🧪 Go 基本命令

### 常用命令速查

| 命令          | 功能         | 示例                |
| ------------- | ------------ | ------------------- |
| `go run`      | 直接运行程序 | `go run main.go`    |
| `go build`    | 编译程序     | `go build`          |
| `go mod init` | 初始化模块   | `go mod init myapp` |
| `go mod tidy` | 整理依赖     | `go mod tidy`       |
| `go test`     | 运行测试     | `go test ./...`     |
| `go fmt`      | 格式化代码   | `go fmt ./...`      |
| `go vet`      | 静态检查     | `go vet ./...`      |

### Go 模块基础

Go 使用模块来管理依赖包：

```bash
# 创建新模块
go mod init github.com/username/project

# 添加依赖
go get github.com/gin-gonic/gin

# 更新依赖
go get -u github.com/gin-gonic/gin

# 移除未使用的依赖
go mod tidy
```

## 🛠️ 开发工具配置

### VS Code 配置（推荐）

1. **安装扩展**
   - Go (官方)
   - GoLand (如果你使用 JetBrains)

2. **安装 Go 工具**
   ```bash
   # VS Code 会提示安装工具，或手动安装
   go install golang.org/x/tools/gopls@latest
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

3. **VS Code 设置**
   
   创建 `.vscode/settings.json`：
   ```json
   {
     "go.useLanguageServer": true,
     "go.formatTool": "goimports",
     "go.lintTool": "golangci-lint",
     "go.testFlags": ["-v", "-race"],
     "go.coverOnSave": true
   }
   ```

### 其他编辑器

- **GoLand**: JetBrains 出品的专用 Go IDE
- **Vim/Neovim**: 使用 vim-go 插件
- **Emacs**: 使用 go-mode

## 📁 项目结构建议

### 基本项目结构

```
my-project/
├── go.mod          # 模块定义文件
├── go.sum          # 依赖校验文件
├── main.go         # 主程序入口
├── cmd/            # 命令行程序
│   └── app/
│       └── main.go
├── internal/       # 内部包
│   ├── config/
│   └── handler/
├── pkg/           # 外部可用包
├── docs/          # 文档
├── examples/      # 示例代码
└── tests/         # 测试文件
```

### 本项目结构

本学习项目采用以下结构：

```
go-study/
├── docs/           # VitePress 文档
├── examples/       # 代码示例
│   ├── basics/     # 基础语法示例
│   ├── advanced/   # 进阶主题示例
│   └── projects/   # 实战项目
├── tools/          # 开发工具
└── go.work         # Go 工作空间配置
```

## 🎯 下一步

完成环境配置后，建议按以下顺序学习：

1. **基础语法** → [变量与常量](../basics/01-variables-constants.md)
2. **数据类型** → [数据类型详解](../basics/02-data-types.md)
3. **控制流程** → [条件与循环](../basics/04-control-flow.md)
4. **函数** → [函数基础](../basics/05-functions.md)

## 🆘 常见问题

### Q: Go 安装后 `go: command not found`
A: 检查 PATH 环境变量是否包含 Go 的 bin 目录

### Q: 代理下载失败
A: 配置 Go 代理：
```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn
```

### Q: VS Code Go 扩展报错
A: 运行 `Go: Install/Update Tools` 安装所有必要工具

### Q: 模块下载缓慢
A: 使用国内镜像：
```bash
go env -w GOPROXY=https://goproxy.io,direct
```

## 📚 扩展资源

- [Go 官方教程](https://tour.golang.org/)
- [Go 语言规范](https://golang.org/ref/spec)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)

## ⏭️ 下一章节

[学习路径规划](./learning-path.md) → 了解完整的学习路径和进度安排

---

**💡 小贴士**: 记得经常运行 `go mod tidy` 来保持依赖的整洁，这是良好的 Go 开发习惯！
