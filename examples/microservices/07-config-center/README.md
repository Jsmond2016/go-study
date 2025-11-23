# 配置中心示例

本目录包含使用 Apollo 和 Nacos 实现配置管理的示例代码。

## 📋 目录结构

```
07-config-center/
├── apollo/              # Apollo 配置中心示例
│   ├── basic/          # 基础使用示例
│   └── database/       # 数据库配置管理示例
├── nacos/              # Nacos 配置中心示例
│   ├── basic/          # 基础使用示例
│   └── publish/        # 配置发布示例
├── docker-compose.yml  # Nacos 服务配置
├── go.mod              # Go 模块定义
└── README.md           # 本文件
```

## 🚀 快速开始

### 1. 启动 Nacos

使用 Docker Compose 启动 Nacos：

```bash
docker-compose up -d
```

访问 Nacos Console：http://localhost:8848/nacos

默认账号：`nacos` / `nacos`

### 2. 启动 Apollo

参考 [Apollo 快速启动](https://www.apolloconfig.io/#/zh/deployment/quick-start)

访问 Apollo Portal：http://localhost:8070

默认账号：`apollo` / `admin`

### 3. 运行 Nacos 示例

**发布配置**：

```bash
cd nacos/publish
go run main.go
```

**监听配置**：

```bash
cd nacos/basic
go run main.go
```

### 4. 运行 Apollo 示例

```bash
cd apollo/basic
go run main.go
```

## 📚 示例说明

### Apollo 示例

#### 基础示例 (apollo/basic/)

演示 Apollo 配置中心的基础使用：

- 初始化 Apollo 客户端
- 获取配置值
- 监听配置变更
- 处理配置更新

#### 数据库配置示例 (apollo/database/)

演示如何使用 Apollo 管理数据库配置：

- 加载数据库配置
- 监听配置变更
- 配置变更时重新连接数据库

### Nacos 示例

#### 基础示例 (nacos/basic/)

演示 Nacos 配置中心的基础使用：

- 初始化 Nacos 客户端
- 获取配置
- 监听配置变更

#### 配置发布示例 (nacos/publish/)

演示如何发布和管理配置：

- 发布配置
- 获取配置
- 删除配置

## 🔧 配置说明

### Nacos 配置

默认配置：
- **Nacos Console**: http://localhost:8848/nacos
- **Namespace**: public
- **Group**: DEFAULT_GROUP

### Apollo 配置

默认配置：
- **Apollo Portal**: http://localhost:8070
- **Config Service**: http://localhost:8080
- **AppID**: your-app-id
- **Cluster**: default
- **Namespace**: application

## 💡 使用场景

### 1. 数据库配置管理

```go
dbConfig := &DatabaseConfig{
	Host:     client.GetStringValue("db.host", "localhost"),
	Port:     client.GetIntValue("db.port", 3306),
	Username: client.GetStringValue("db.username", "root"),
	Password: client.GetStringValue("db.password", ""),
	Database: client.GetStringValue("db.database", "test"),
}
```

### 2. Redis 配置管理

```go
redisConfig := &RedisConfig{
	Host:     client.GetStringValue("redis.host", "localhost"),
	Port:     client.GetIntValue("redis.port", 6379),
	Password: client.GetStringValue("redis.password", ""),
	DB:       client.GetIntValue("redis.db", 0),
}
```

### 3. 第三方服务配置

```go
thirdPartyConfig := &ThirdPartyConfig{
	APIKey:    client.GetStringValue("third-party.api-key", ""),
	APISecret: client.GetStringValue("third-party.api-secret", ""),
	BaseURL:   client.GetStringValue("third-party.base-url", ""),
}
```

## 🎯 学习要点

1. **配置集中管理**：统一管理所有配置
2. **动态更新**：配置变更实时生效
3. **环境隔离**：不同环境配置隔离
4. **版本控制**：配置变更历史追踪
5. **权限控制**：配置访问权限管理

## 📖 相关文档

- [配置中心文档](../../../docs/microservices/07-config-center.md)
- [Apollo 官方文档](https://www.apolloconfig.io/)
- [Nacos 官方文档](https://nacos.io/docs/what-is-nacos.html)

## 🐛 常见问题

### 1. 无法连接到配置中心

确保配置中心服务已启动：

```bash
# Nacos
docker-compose ps

# Apollo
# 检查 Apollo 服务状态
```

### 2. 配置获取失败

- 检查 AppID/DataId 是否正确
- 检查命名空间和分组是否正确
- 检查配置是否已发布

### 3. 配置变更不生效

- 确保已添加配置变更监听器
- 检查配置中心推送是否正常
- 检查客户端连接是否正常

## 📝 下一步

- 学习消息队列：`08-message-queue.md`
- 学习服务网格：`09-service-mesh.md`

---

**🎉 开始使用配置中心，实现配置的集中管理和动态更新！**

