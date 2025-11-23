---
title: Gin 数据库操作
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["Gin 基础", "数据库基础"]
tags: ["Gin", "数据库", "MySQL", "CRUD"]
---

# Gin 数据库操作

本指南介绍如何在 Gin 应用中操作数据库，包括 MySQL 的连接、查询、插入、更新和删除操作。

## 📋 学习目标

- [ ] 掌握数据库连接和初始化
- [ ] 学会执行 CRUD 操作
- [ ] 理解数据库连接池管理
- [ ] 掌握错误处理和事务
- [ ] 了解数据库最佳实践

## 🔌 数据库连接

### 初始化数据库

```go
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func initDB() (*sql.DB, error) {
	// 数据库连接字符串
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	db.SetMaxOpenConns(100)                 // 最大打开连接数
	db.SetMaxIdleConns(10)                  // 最大空闲连接数
	db.SetConnMaxLifetime(time.Hour)        // 连接最大生存时间

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
```

## 📊 CRUD 操作

### 定义模型

```go
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
```

### 查询操作

#### 查询单个记录

```go
func getUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	query := "SELECT id, username, password FROM user_info WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
```

#### 查询多个记录

```go
func getAllUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, username, password FROM user_info"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
```

#### 模糊查询

```go
func searchUsers(db *sql.DB, keyword string) ([]User, error) {
	query := "SELECT id, username, password FROM user_info WHERE username LIKE ?"

	rows, err := db.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
```

### 插入操作

#### 单个插入

```go
func addUser(db *sql.DB, user User) (int, error) {
	query := "INSERT INTO user_info(username, password) VALUES (?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
```

#### 批量插入

```go
func addUsers(db *sql.DB, users []User) error {
	query := "INSERT INTO user_info(username, password) VALUES (?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, user := range users {
		_, err := stmt.Exec(user.Username, user.Password)
		if err != nil {
			return err
		}
	}

	return nil
}
```

### 更新操作

#### 单个更新

```go
func updateUser(db *sql.DB, user User) (int64, error) {
	query := "UPDATE user_info SET username=?, password=? WHERE id=?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password, user.ID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
```

### 删除操作

#### 单个删除

```go
func deleteUser(db *sql.DB, id int) (int64, error) {
	query := "DELETE FROM user_info WHERE id=?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
```

## 🏃‍♂️ 完整示例

### 路由处理

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()

	// 查询所有用户
	router.GET("/users", func(c *gin.Context) {
		users, err := getAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
	})

	// 添加用户
	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := addUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s 插入成功, id 为 %d", user.Username, id),
		})
	})

	// 更新用户
	router.PUT("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, err := updateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("修改 id: %d 成功，影响行数: %d", user.ID, rowsAffected),
		})
	})

	// 删除用户
	router.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
			return
		}

		rowsAffected, err := deleteUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("成功删除用户，影响行数: %d", rowsAffected),
		})
	})

	router.Run(":8080")
}

func getAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, username, password FROM user_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func addUser(user User) (int, error) {
	stmt, err := db.Prepare("INSERT INTO user_info(username, password) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func updateUser(user User) (int64, error) {
	stmt, err := db.Prepare("UPDATE user_info SET username=?, password=? WHERE id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password, user.ID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func deleteUser(id int) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM user_info WHERE id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
```

## ⚠️ 注意事项

### 1. 连接管理

```go
// ✅ 正确：使用连接池
db.SetMaxOpenConns(100)
db.SetMaxIdleConns(10)

// ✅ 正确：及时关闭连接
defer db.Close()
```

### 2. 资源释放

```go
// ✅ 正确：关闭 rows
rows, err := db.Query("SELECT ...")
defer rows.Close()

// ✅ 正确：关闭 stmt
stmt, err := db.Prepare("INSERT ...")
defer stmt.Close()
```

### 3. 错误处理

```go
// ✅ 正确：检查所有错误
if err != nil {
	return err
}

// ✅ 正确：检查 rows.Err()
if err := rows.Err(); err != nil {
	return err
}
```

### 4. SQL 注入防护

```go
// ✅ 正确：使用参数化查询
db.Query("SELECT * FROM users WHERE id = ?", id)

// ❌ 错误：字符串拼接（容易 SQL 注入）
db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %s", id))
```

## 📚 扩展阅读

- [database/sql 包文档](https://pkg.go.dev/database/sql)
- [MySQL 驱动文档](https://github.com/go-sql-driver/mysql)
- [GORM 文档](https://gorm.io/) - 推荐使用 ORM 框架

## ⏭️ 下一章节

完成数据库操作学习后，可以继续学习：

- [数据验证](./06-gin-validation.md) - 学习请求验证
- [认证授权](./07-gin-auth.md) - 学习用户认证

---

**💡 提示**: 在实际项目中，建议使用 GORM 等 ORM 框架来简化数据库操作，提高开发效率！

