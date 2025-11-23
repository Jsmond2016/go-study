---
title: 环境搭建
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["Web 开发", "数据库", "工具链"]
tags: ["环境", "搭建", "配置", "初始化"]
---

# 环境搭建

本章节将指导你搭建电商系统的开发环境，包括项目初始化、依赖安装和配置管理。

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
mkdir -p e-commerce
cd e-commerce

# 创建标准 Go 项目目录结构
mkdir -p cmd/server
mkdir -p internal/{handler,service,repository,model}
mkdir -p pkg/{payment,utils}
mkdir -p config
mkdir -p scripts
```

### 2. 初始化 Go 模块

```bash
# 初始化 Go 模块
go mod init e-commerce

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
go get github.com/go-redis/redis/v8
```

## ⚙️ 配置文件

创建 `config/config.yaml`:

```yaml
server:
  port: 8080
  mode: debug

database:
  driver: mysql
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: ecommerce_db
  charset: utf8mb4

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  secret: your-secret-key
  expire: 24h

payment:
  alipay:
    app_id: your_app_id
    private_key: your_private_key
    public_key: your_public_key
  wechat:
    app_id: your_app_id
    mch_id: your_mch_id
    api_key: your_api_key
```

## 🗄️ 数据库初始化

创建数据库：

```sql
CREATE DATABASE IF NOT EXISTS ecommerce_db 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;
```

## ✅ 验证安装

运行项目并测试：

```bash
go run cmd/server/main.go
```

## ⏭️ 下一步

环境搭建完成后，下一步是：

- [数据模型设计](./02-models.md) - 设计数据库表和模型

---

**🎉 环境搭建完成！** 现在你可以开始设计数据模型了。

