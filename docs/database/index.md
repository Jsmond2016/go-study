---
title: 数据库操作
difficulty: intermediate
duration: "6-8周"
prerequisites: ["基础语法", "错误处理", "接口"]
tags: ["数据库", "SQL", "GORM", "Redis", "MongoDB"]
---

# 数据库操作

本模块介绍 Go 语言中的数据库操作，包括 SQL 数据库、ORM 框架、NoSQL 数据库等。

## 📋 学习目标

完成本模块学习后，你将能够：

- [ ] 掌握 SQL 数据库的基本操作
- [ ] 使用 GORM 进行 ORM 开发
- [ ] 操作 Redis 缓存
- [ ] 使用 MongoDB 进行文档存储
- [ ] 理解连接池和性能优化
- [ ] 掌握数据库最佳实践

## 🎯 学习路径

### 📚 第一部分：SQL 数据库（第1-2周）

| 章节                              | 内容               | 预计时间 | 难度 |
| --------------------------------- | ------------------ | -------- | ---- |
| [SQL 基础](./01-sql.md)           | database/sql 使用  | 4-5小时  | ⭐⭐   |
| [GORM 框架](./02-gorm.md)         | ORM 操作           | 5-6小时  | ⭐⭐⭐ |

### 🗄️ 第二部分：NoSQL 数据库（第3-4周）

| 章节                              | 内容               | 预计时间 | 难度 |
| --------------------------------- | ------------------ | -------- | ---- |
| [Redis 操作](./03-redis.md)       | 缓存和数据结构     | 4-5小时  | ⭐⭐⭐ |
| [MongoDB 操作](./04-mongodb.md)   | 文档数据库         | 4-5小时  | ⭐⭐⭐ |

### ⚙️ 第三部分：高级主题（第5-6周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [连接池](./05-connection-pool.md)       | 连接池配置和优化   | 3-4小时  | ⭐⭐⭐ |

## 🚀 快速开始

### 第一个数据库程序

```go
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("数据库连接成功")
}
```

## 💡 学习建议

### 📖 学习方法

1. **实践为主**：每个概念都要动手编写代码
2. **理解原理**：理解数据库操作的底层原理
3. **性能优化**：关注查询性能和连接管理
4. **安全考虑**：注意 SQL 注入和权限控制

### 🔍 推荐资源

- [database/sql 官方文档](https://pkg.go.dev/database/sql)
- [GORM 官方文档](https://gorm.io/)
- [Redis Go 客户端](https://github.com/redis/go-redis)

## ⏭️ 下一阶段

完成数据库学习后，可以进入：

- [Web 开发](../web-development/) - 构建 Web 应用
- [实战项目](../projects/) - 完整项目开发

---

**🎉 开始你的数据库学习之旅吧！** 选择第一个章节，开始学习数据库操作。

