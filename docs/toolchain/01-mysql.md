---
title: MySQL 基础
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["基础语法", "数据库基础"]
tags: ["MySQL", "数据库", "SQL", "驱动"]
---

# MySQL 基础

MySQL 是最流行的关系型数据库之一，本教程介绍如何在 Go 中使用 MySQL 进行数据库操作。

## 📋 学习目标

- [ ] 理解 MySQL 数据库基础
- [ ] 掌握 Go 连接 MySQL 的方法
- [ ] 学会执行 SQL 查询和操作
- [ ] 理解连接池的使用
- [ ] 掌握事务处理
- [ ] 了解数据库最佳实践

## 🎯 MySQL 简介

### 为什么选择 MySQL

- **流行度高**: 使用广泛，社区活跃
- **性能优秀**: 适合 Web 应用
- **易于使用**: 学习曲线平缓
- **功能完整**: 支持事务、索引等

### Go MySQL 驱动

Go 标准库提供了 `database/sql` 接口，常用的 MySQL 驱动：

- `github.com/go-sql-driver/mysql` - 官方推荐
- `github.com/go-gorm/mysql` - GORM 驱动

## 🚀 快速开始

### 安装驱动

```bash
go get github.com/go-sql-driver/mysql
```

### 基本连接

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据源名称格式: username:password@tcp(host:port)/dbname?charset=utf8mb4
	dsn := "root:password@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	// 测试连接
	if err := db.Ping(); err != nil {
		panic(err)
	}
	
	fmt.Println("数据库连接成功")
}
```

## 📊 数据库操作

### 创建表

```go
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		age INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	
	_, err := db.Exec(query)
	return err
}
```

### 插入数据

```go
func insertUser(db *sql.DB, name, email string, age int) (int64, error) {
	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	
	result, err := db.Exec(query, name, email, age)
	if err != nil {
		return 0, err
	}
	
	id, err := result.LastInsertId()
	return id, err
}

// 批量插入
func insertUsers(db *sql.DB, users []User) error {
	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	for _, user := range users {
		_, err := stmt.Exec(user.Name, user.Email, user.Age)
		if err != nil {
			return err
		}
	}
	
	return nil
}
```

### 查询数据

```go
type User struct {
	ID        int
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 查询单条记录
func getUserByID(db *sql.DB, id int) (*User, error) {
	query := "SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?"
	
	var user User
	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// 查询多条记录
func getUsers(db *sql.DB, limit, offset int) ([]User, error) {
	query := "SELECT id, name, email, age, created_at, updated_at FROM users LIMIT ? OFFSET ?"
	
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, rows.Err()
}

// 条件查询
func searchUsers(db *sql.DB, keyword string) ([]User, error) {
	query := `
		SELECT id, name, email, age, created_at, updated_at 
		FROM users 
		WHERE name LIKE ? OR email LIKE ?
	`
	
	keyword = "%" + keyword + "%"
	rows, err := db.Query(query, keyword, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, rows.Err()
}
```

### 更新数据

```go
func updateUser(db *sql.DB, id int, name, email string, age int) error {
	query := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
	
	result, err := db.Exec(query, name, email, age, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}
	
	return nil
}
```

### 删除数据

```go
func deleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}
	
	return nil
}
```

## 🔄 事务处理

```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	
	// 扣除转出账户金额
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		return err
	}
	
	// 增加转入账户金额
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		return err
	}
	
	return nil
}
```

## 🔧 连接池配置

```go
func setupDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	// 设置最大打开连接数
	db.SetMaxOpenConns(25)
	
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	
	// 设置连接最大生存时间
	db.SetConnMaxLifetime(5 * time.Minute)
	
	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}
```

## 🏃‍♂️ 实践应用

### 完整的 CRUD 示例

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService struct {
	db *sql.DB
}

func NewUserService(dsn string) (*UserService, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	// 配置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	return &UserService{db: db}, nil
}

func (s *UserService) Create(name, email string, age int) (*User, error) {
	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	
	result, err := s.db.Exec(query, name, email, age)
	if err != nil {
		return nil, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	return s.GetByID(int(id))
}

func (s *UserService) GetByID(id int) (*User, error) {
	query := "SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?"
	
	var user User
	err := s.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (s *UserService) List(limit, offset int) ([]User, error) {
	query := "SELECT id, name, email, age, created_at, updated_at FROM users LIMIT ? OFFSET ?"
	
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, rows.Err()
}

func (s *UserService) Update(id int, name, email string, age int) error {
	query := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
	
	result, err := s.db.Exec(query, name, email, age, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}
	
	return nil
}

func (s *UserService) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}
	
	return nil
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	
	service, err := NewUserService(dsn)
	if err != nil {
		log.Fatal(err)
	}
	
	// 创建用户
	user, err := service.Create("张三", "zhangsan@example.com", 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("创建用户: %+v\n", user)
	
	// 查询用户
	user, err = service.GetByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查询用户: %+v\n", user)
	
	// 更新用户
	err = service.Update(user.ID, "李四", "lisi@example.com", 30)
	if err != nil {
		log.Fatal(err)
	}
	
	// 查询列表
	users, err := service.List(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("用户列表: %+v\n", users)
}
```

## ⚠️ 注意事项

### 1. SQL 注入防护

```go
// ✅ 使用参数化查询
db.Query("SELECT * FROM users WHERE id = ?", id)

// ❌ 不要拼接 SQL
db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %d", id))
```

### 2. 连接管理

```go
// ✅ 总是关闭连接和结果集
rows, err := db.Query(...)
defer rows.Close()

// ✅ 使用连接池
db.SetMaxOpenConns(25)
```

### 3. 错误处理

```go
// ✅ 检查所有错误
if err != nil {
	return err
}

// ✅ 检查查询结果
if rowsAffected == 0 {
	return fmt.Errorf("记录不存在")
}
```

## 📚 扩展阅读

- [database/sql 官方文档](https://pkg.go.dev/database/sql)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [MySQL 官方文档](https://dev.mysql.com/doc/)

## ⏭️ 下一章节

[GORM 框架](./02-gorm.md) → 学习更高级的 ORM 操作

---

**💡 提示**: 掌握原生 SQL 操作是理解 ORM 框架的基础，建议先熟悉 database/sql 的使用！

