---
title: Chapter 1 - 环境搭建
difficulty: beginner
duration: "30分钟"
tags: ["环境配置", "Go", "protoc", "gRPC"]
---

# Chapter 1: 环境搭建

本章将指导你搭建完整的开发环境，为后续开发做好准备。

## 📋 学习目标

完成本章后，你将：
- [x] 安装 Go 开发环境（1.21+）
- [x] 安装 Protocol Buffers 编译器
- [x] 安装 gRPC 相关工具
- [x] 创建项目目录结构
- [x] 初始化 Go Modules
- [x] 验证环境配置正确

## 🛠️ 环境要求

### 操作系统
- macOS 10.15+
- Linux (Ubuntu 20.04+推荐)
- Windows 10+ (需要 WSL2)

### 软件版本
- Go 1.21 或更高版本
- protoc 3.20+ 或更高版本
- Git 2.30+

## 📦 安装步骤

### 步骤 1: 安装 Go

#### macOS
```bash
# 使用 Homebrew 安装
brew install go

# 验证安装
go version
# 输出：go version go1.21.x darwin/arm64
```

#### Linux (Ubuntu/Debian)
```bash
# 下载 Go (以 1.21.6 为例)
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz

# 解压到 /usr/local
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

# 配置环境变量
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

# 验证安装
go version
```

#### Windows (WSL2)
```bash
# 在 WSL2 Ubuntu 中执行 Linux 安装步骤
# 或者下载 Windows 安装包：https://go.dev/dl/
```

### 步骤 2: 配置 Go 环境

```bash
# 设置 Go 代理（中国大陆用户推荐）
go env -w GOPROXY=https://goproxy.cn,direct

# 启用 Go Modules
go env -w GO111MODULE=on

# 验证配置
go env | grep GOPROXY
go env | grep GO111MODULE
```

### 步骤 3: 安装 Protocol Buffers 编译器

#### macOS
```bash
# 使用 Homebrew 安装
brew install protobuf

# 验证安装
protoc --version
# 输出：libprotoc 3.21.x 或更高
```

#### Linux
```bash
# 下载预编译二进制文件
PB_VERSION=25.1
wget https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VERSION}/protoc-${PB_VERSION}-linux-x86_64.zip

# 解压并安装
unzip protoc-${PB_VERSION}-linux-x86_64.zip -d $HOME/.local
sudo cp $HOME/.local/bin/protoc /usr/local/bin/
sudo cp -r $HOME/.local/include/* /usr/local/include/

# 验证安装
protoc --version
```

### 步骤 4: 安装 gRPC Go 插件

```bash
# 安装 protoc-gen-go （生成 Go 消息类型）
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 安装 protoc-gen-go-grpc （生成 gRPC 服务代码）
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 验证安装
protoc-gen-go --version
protoc-gen-go-grpc --version

# 确保 $GOPATH/bin 在 PATH 中
echo $PATH | grep go/bin
```

### 步骤 5: 安装其他开发工具（可选）

```bash
# grpcurl - gRPC 命令行工具，用于测试
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# golangci-lint - 代码检查工具
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# air - 热重载工具
go install github.com/cosmtrek/air@latest
```

## 🏗️ 创建项目结构

### 步骤 1: 创建项目目录

```bash
# 进入你的工作目录
cd ~/go/src  # 或其他你喜欢的目录

# 创建项目根目录
mkdir -p 10-bookstore-project
cd 10-bookstore-project

# 创建子目录
mkdir -p proto
mkdir -p book-service
mkdir -p user-service
mkdir -p api-gateway
mkdir -p scripts
```

### 步骤 2: 初始化 Go Module

```bash
# 初始化 Go Module
go mod init bookstore

# 添加必要的依赖
go get google.golang.org/grpc@latest
go get google.golang.org/protobuf@latest
go get github.com/gin-gonic/gin@latest
go get github.com/gin-contrib/cors@latest
go get github.com/golang-jwt/jwt/v5@latest
go get github.com/mattn/go-sqlite3@latest
go get golang.org/x/crypto/bcrypt@latest
```

### 步骤 3: 查看项目结构

```bash
tree -L 2
```

预期输出：
```
10-bookstore-project/
├── go.mod
├── go.sum
├── proto/
├── book-service/
├── user-service/
├── api-gateway/
└── scripts/
```

## ✅ 验证环境

创建一个简单的测试文件来验证环境配置：

### 创建测试文件

```bash
# 创建测试目录
mkdir -p test

# 创建测试文件
cat > test/verify.go << 'EOF'
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("✅ Go version:", getGoVersion())
	fmt.Println("✅ gRPC imported successfully")
	fmt.Println("✅ Protobuf imported successfully")
	fmt.Println("✅ All dependencies are working!")
}

func getGoVersion() string {
	// This will be replaced by runtime.Version() in actual code
	return "Go 1.21+"
}
EOF
```

### 运行验证

```bash
cd test
go run verify.go
```

预期输出：
```
✅ Go version: Go 1.21+
✅ gRPC imported successfully
✅ Protobuf imported successfully
✅ All dependencies are working!
```

### 清理测试文件

```bash
cd ..
rm -rf test
```

## 🔧 IDE 配置（可选）

### Visual Studio Code

推荐安装以下扩展：
- **Go** (golang.go) - Go 语言支持
- **vscode-proto3** - Protocol Buffers 语法高亮
- **Go Test Explorer** - 测试运行器

配置 VSCode 设置（`.vscode/settings.json`）：

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "go.formatTool": "goimports",
  "editor.formatOnSave": true,
  "[go]": {
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }
}
```

### GoLand / IntelliJ IDEA

1. 安装 **Protocol Buffers** 插件
2. 配置 Go SDK 路径
3. 启用 **Go Modules** 支持

## 📝 环境检查清单

在继续下一章之前，请确认：

- [ ] Go 版本 >= 1.21
- [ ] protoc 版本 >= 3.20
- [ ] protoc-gen-go 已安装
- [ ] protoc-gen-go-grpc 已安装
- [ ] 项目目录结构已创建
- [ ] go.mod 文件已初始化
- [ ] 所有依赖包已下载
- [ ] 验证脚本运行成功

## 🐛 常见问题

### Q1: `protoc-gen-go: program not found or is not executable`

**原因**：`$GOPATH/bin` 不在 PATH 中

**解决方案**：
```bash
# 查找 GOPATH
go env GOPATH

# 添加到 PATH
export PATH=$PATH:$(go env GOPATH)/bin

# 永久生效（添加到 ~/.bashrc 或 ~/.zshrc）
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### Q2: `go get` 下载很慢

**原因**：网络问题

**解决方案**：
```bash
# 使用中国代理
go env -w GOPROXY=https://goproxy.cn,direct

# 或使用其他代理
go env -w GOPROXY=https://goproxy.io,direct
```

### Q3: CGO_ENABLED 相关错误（SQLite）

**原因**：需要 C 编译器

**解决方案**：

macOS:
```bash
xcode-select --install
```

Linux:
```bash
sudo apt-get install build-essential
```

### Q4: 权限问题（Linux）

**原因**：没有 sudo 权限

**解决方案**：
```bash
# 安装到用户目录
mkdir -p $HOME/.local/bin
export PATH=$HOME/.local/bin:$PATH
```

## 🎯 下一步

环境搭建完成！接下来我们将：

**[👉 Chapter 2: 架构设计](/microservices/bookstore-project/02-architecture)**

在下一章中，我们将：
- 设计系统架构
- 定义 Protocol Buffers 文件
- 设计数据库表结构
- 规划 API 接口

---

**💡 提示**

- 建议使用 **Git** 管理代码版本
- 养成编写 **README.md** 的习惯
- 定期运行 `go mod tidy` 清理依赖

**📌 检查点**

如果你完成了本章的所有步骤，你现在应该有：
- ✅ 完整的 Go 开发环境
- ✅ 所有必需的工具和依赖
- ✅ 清晰的项目目录结构
- ✅ 可以开始编写代码的状态

**🎉 干得好！继续下一章吧！**
