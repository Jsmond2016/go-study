# Go 语言学习笔记

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)
![VitePress](https://img.shields.io/badge/VitePress-1.6.4-646CFF?style=flat-square&logo=vitepress)

**系统化的 Go 语言学习路径，从零基础到实战应用**

[📖 在线文档](https://jsmond2016.github.io/go-study/) | [🚀 快速开始](#-快速开始) | [📚 学习路径](#-学习路径)

</div>

## ✨ 项目简介

这是一个系统化的 Go 语言学习项目，旨在帮助开发者从零基础到能够独立开发 Go 应用。项目包含：

- 📚 **完整的文档体系** - 使用 VitePress 构建的现代化文档站点
- 💻 **丰富的代码示例** - 每个知识点都配有可运行的代码示例
- 🎯 **实战项目** - 提供完整的项目实战教程（TODO API、博客系统、电商系统、聊天应用）
- 🛠️ **工具链指南** - 详细介绍开发中常用的工具和框架
- 📖 **学习资源** - 整理推荐优质的学习资源

## 🎯 主要特性

- ✅ **系统化学习路径** - 从基础语法到实战项目，循序渐进
- ✅ **理论与实践结合** - 每个知识点都配有代码示例和实践练习
- ✅ **完整的项目实战** - 提供多个完整的实战项目教程
- ✅ **现代化文档** - 使用 VitePress 构建，支持搜索、暗色模式等
- ✅ **持续更新** - 内容持续更新和完善

## 📚 学习内容

### 基础语法
- 变量与常量
- 数据类型
- 运算符
- 控制流程
- 函数
- 数组、切片、映射
- 结构体
- 指针
- 接口
- 错误处理
- 并发编程
- 反射
- 测试

### 标准库
- fmt - 格式化输出
- time - 时间处理
- flag - 命令行参数
- log - 日志记录
- 文件 I/O
- strconv - 类型转换
- net/http - HTTP 客户端和服务器
- context - 上下文管理
- encoding - 编码解码
- crypto - 加密解密
- strings - 字符串操作
- math - 数学运算

### Web 开发
- HTTP 服务器
- Gin 框架基础
- 路由与中间件
- 模板渲染
- 数据验证
- 认证与授权
- RESTful API 设计

### 开发工具链
- MySQL 数据库
- GORM ORM
- Go Modules
- Viper 配置管理
- Zap 日志库
- JWT 认证
- Validator 数据验证
- CORS 跨域
- 限流
- Swagger API 文档
- Redis 缓存
- Docker 容器化
- Testify 测试框架
- Cron 定时任务
- WebSocket 实时通信

### 实战项目
- **TODO API** - RESTful API 项目，学习基础的 Web 开发
- **博客系统** - 完整的博客系统，包含文章、评论、上传等功能
- **电商系统** - 电商系统，包含商品、购物车、订单、支付等功能
- **聊天应用** - 实时聊天应用，使用 WebSocket 实现

## 🚀 快速开始

### 环境要求

- Go 1.21 或更高版本
- Node.js 18+ (用于文档开发)
- pnpm 10.23.0+ (推荐使用)

### 安装步骤

1. **克隆项目**
```bash
git clone https://github.com/Jsmond2016/go-study.git
cd go-study
```

2. **安装依赖**
```bash
# 安装文档依赖
pnpm install

# 安装 Go 依赖（如果需要）
cd code && go mod download
```

3. **启动文档站点**
```bash
# 开发模式
pnpm dev

# 构建文档
pnpm build

# 预览构建结果
pnpm preview
```

4. **运行代码示例**
```bash
# 进入示例目录
cd examples/basics/01-variables-constants

# 运行示例
go run main.go
```

## 📖 在线文档

访问 [在线文档站点](https://jsmond2016.github.io/go-study/) 查看完整的文档内容。

文档包含：
- 📘 基础语法教程
- 📗 标准库使用指南
- 📙 Web 开发教程
- 📕 工具链使用指南
- 📓 实战项目教程
- 📔 学习资源推荐

## 📁 项目结构

```
go-study/
├── docs/                    # 文档目录
│   ├── basics/              # 基础语法
│   ├── standard-library/    # 标准库
│   ├── web-development/     # Web 开发
│   ├── toolchain/           # 开发工具链
│   ├── projects/            # 实战项目
│   ├── guide/               # 学习指南
│   └── resources/           # 学习资源
├── examples/                 # 代码示例
│   ├── basics/              # 基础语法示例
│   ├── standard-library/    # 标准库示例
│   ├── web-development/     # Web 开发示例
│   └── projects/            # 项目示例
├── code/                     # 练习代码
├── tools/                    # 工具脚本
└── docs/.vitepress/          # VitePress 配置
```

## 🎓 学习路径

### 初级阶段（0-3个月）
- 环境搭建与基础语法
- 数据类型与变量
- 控制流程与函数
- 数据结构（数组、切片、映射、结构体）
- 指针与内存管理
- 接口与错误处理
- 并发编程基础

### 中级阶段（3-6个月）
- 标准库深入学习
- HTTP 服务器开发
- Gin 框架使用
- 数据库操作（SQL、GORM）
- Redis 缓存使用
- REST API 设计
- 中间件开发
- 认证与授权

### 高级阶段（6-12个月）
- 性能优化与内存管理
- 泛型与反射
- gRPC 与微服务
- 服务发现与负载均衡
- 容器化部署（Docker、K8s）
- CI/CD 流水线
- 监控与日志
- 分布式系统设计

详细的学习路径请参考 [学习路径文档](https://jsmond2016.github.io/go-study/guide/learning-path)。

## 💡 使用建议

1. **循序渐进** - 按照文档顺序学习，不要跳跃
2. **动手实践** - 每学一个知识点都要编写代码并运行
3. **完成练习** - 每章都有实践练习，务必完成
4. **做项目** - 通过实际项目巩固知识
5. **阅读源码** - 学习优秀项目的代码

## 🤝 贡献

欢迎贡献代码、文档或提出建议！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📝 常见问题

### Windows 下执行 `go mod tidy` 报错 `Permission Denied`

**解决方法**：
```bash
go mod init test  # 随机创建一个包名
go mod tidy
```

### 下载安装太慢

如果 `go get` 下载失败，可以：
1. 配置 Go 代理：`go env -w GOPROXY=https://goproxy.cn,direct`
2. 使用国内镜像源

### 文档构建失败

确保已安装正确的依赖：
```bash
pnpm install
```

## 📚 推荐资源

### 学习网站
- [Go 官方文档](https://golang.org/doc/)
- [Go 语言之旅](https://tour.golang.org/)
- [TopGoer](https://www.topgoer.com/) - 强烈推荐
- [李文周博客](https://www.liwenzhou.com/)

### 开源项目
- [7days-golang](https://github.com/geektutu/7days-golang) - 7天用Go从零实现系列
- [Go 从入门到实战](https://github.com/xinliangnote/Go) - 从零开始学 Go、Gin 框架
- [go-admin](https://github.com/go-admin-team/go-admin) - 基于Gin + Vue的前后端分离权限管理系统
- [Go语言高级编程](https://github.com/chai2010/advanced-go-programming-book) - 开源电子书
- [LeetCode-Go](https://github.com/halfrost/LeetCode-Go) - Go语言刷LeetCode

### 实战项目
- [go-gin-api](https://github.com/xinliangnote/go-gin-api) - 基于 Gin 进行模块化设计的 API 框架

更多资源请查看 [学习资源文档](https://jsmond2016.github.io/go-study/resources/)。

## 📄 许可证

本项目采用 [MIT](LICENSE) 许可证。

## 🙏 致谢

感谢所有为 Go 语言社区做出贡献的开发者！

## 📮 联系方式

- GitHub: [@Jsmond2016](https://github.com/Jsmond2016)
- Email: Jsmond2016@gmail.com

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给个 Star ⭐**

Made with ❤️ by [Jsmond2016](https://github.com/Jsmond2016)

</div>
