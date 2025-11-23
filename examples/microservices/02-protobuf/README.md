# Protocol Buffers 示例

本示例展示了 Protocol Buffers 的各种用法，包括基础消息、复杂结构、服务定义和版本兼容性。

## 📋 目录结构

```
02-protobuf/
├── basic/              # 基础消息示例
│   ├── user.proto     # 基础用户消息定义
│   └── main.go        # 序列化/反序列化示例
├── complex/            # 复杂消息示例
│   ├── product.proto  # 复杂产品消息定义
│   └── main.go        # 复杂消息使用示例
├── service/            # 服务定义示例
│   └── user_service.proto  # 用户服务定义
├── versioning/         # 版本兼容性示例
│   ├── user_v1.proto  # 版本 1
│   └── user_v2.proto  # 版本 2
├── go.mod             # Go 模块定义
├── Makefile           # 构建脚本
└── README.md          # 本文件
```

## 🚀 快速开始

### 1. 安装依赖

```bash
# 安装 Protocol Buffers 编译器
# macOS
brew install protobuf

# Linux
sudo apt-get install protobuf-compiler

# 安装 Go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### 2. 生成代码

```bash
make proto
```

### 3. 安装 Go 依赖

```bash
go mod download
go mod tidy
```

### 4. 运行示例

```bash
# 基础示例
make basic

# 复杂消息示例
make complex
```

## 📚 示例说明

### 1. 基础消息 (`basic/`)

演示最基本的 Protocol Buffers 消息定义和使用：

- 定义简单的用户消息
- 序列化和反序列化
- 消息比较

### 2. 复杂消息 (`complex/`)

演示复杂的数据结构：

- **枚举类型**: ProductStatus
- **嵌套消息**: Address, Supplier
- **数组**: repeated 字段
- **Map**: map<string, string>
- **时间戳**: google.protobuf.Timestamp
- **Oneof**: 互斥字段

### 3. 服务定义 (`service/`)

演示如何定义 gRPC 服务：

- 服务接口定义
- 请求和响应消息
- 使用 google.protobuf.Empty

### 4. 版本兼容性 (`versioning/`)

演示如何保持向后兼容：

- 添加新字段
- 使用 reserved 保留已删除字段
- 字段编号管理

## 🔧 使用示例

### 基础消息

```go
import (
    "google.golang.org/protobuf/proto"
    pb "go-study/examples/microservices/02-protobuf/basic/proto"
)

// 创建消息
user := &pb.User{
    Id:     1,
    Name:   "Alice",
    Email:  "alice@example.com",
    Age:    25,
    Active: true,
}

// 序列化
data, err := proto.Marshal(user)

// 反序列化
var newUser pb.User
proto.Unmarshal(data, &newUser)
```

### 复杂消息

```go
product := &pb.Product{
    Id:    1,
    Name:  "Go Book",
    Price: 49.99,
    Status: pb.ProductStatus_PRODUCT_STATUS_ACTIVE,
    Tags: []string{"programming", "go"},
    Attributes: map[string]string{
        "isbn": "978-0123456789",
    },
    Discount: &pb.Product_Percentage{
        Percentage: 10.0,
    },
}
```

## 📝 最佳实践

### 1. 字段编号

- 常用字段使用 1-15（编码更高效）
- 不要重用已删除的字段编号
- 为未来扩展预留编号范围

### 2. 命名规范

- 消息名使用 PascalCase
- 字段名使用 snake_case
- 枚举值使用 UPPER_SNAKE_CASE

### 3. 版本兼容

- 可以安全地添加新字段
- 使用 reserved 标记删除的字段
- 不要更改字段类型

## 🐛 常见问题

### 1. 导入 google/protobuf 类型失败

确保安装了完整的 Protocol Buffers：

```bash
# 检查 protobuf 包含文件
ls /usr/local/include/google/protobuf/
```

### 2. 代码生成失败

检查 proto 文件路径和 go_package 选项：

```protobuf
option go_package = "./proto;package_name";
```

### 3. 时间戳类型

使用 `google.golang.org/protobuf/types/known/timestamppb`：

```go
import "google.golang.org/protobuf/types/known/timestamppb"

createdAt := timestamppb.Now()
```

## 📖 相关文档

- [Protocol Buffers 官方文档](https://protobuf.dev/)
- [Protocol Buffers 语言指南](https://protobuf.dev/programming-guides/proto3/)
- [项目文档](../../../docs/microservices/02-protobuf.md)

## ⚠️ 注意事项

1. 字段编号一旦使用，不应更改
2. 删除字段时使用 reserved 保留编号
3. 新字段应使用新的编号，不要重用旧编号

## 🎯 下一步

完成本示例后，可以继续学习：

- [gRPC 示例](../01-grpc/) - 使用 Protocol Buffers 构建 gRPC 服务
- [服务发现示例](../03-service-discovery/) - 实现服务注册和发现

