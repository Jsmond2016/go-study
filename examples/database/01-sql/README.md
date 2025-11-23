# SQL 数据库操作示例

本示例展示了 Go 语言中使用标准库 `database/sql` 进行数据库操作的完整指南，包括连接管理、CRUD操作、事务处理、预处理语句等。

## 运行示例

```bash
# 安装SQLite驱动
go get github.com/mattn/go-sqlite3

# 运行示例
go run main.go

# 运行后会生成 data.db 文件
```

## 功能演示

示例程序会演示以下功能：

1. **数据库连接** - 使用SQLite数据库
2. **表结构创建** - 用户表和文章表
3. **CRUD操作** - 增删改查
4. **事务处理** - 事务提交和回滚
5. **预处理语句** - 参数化查询
6. **表连接** - INNER JOIN和LEFT JOIN
7. **聚合函数** - COUNT、AVG、SUM等
8. **元数据查询** - 表结构信息

## 知识点

### 1. 数据库连接

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3" // 导入驱动
)

// 连接数据库
db, err := sql.Open("sqlite3", "./data.db")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// 测试连接
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```

### 2. 基本CRUD操作

```go
// INSERT - 插入数据
result, err := db.Exec(
    "INSERT INTO users (username, email, age) VALUES (?, ?, ?)",
    username, email, age,
)
if err != nil {
    return err
}
id, _ := result.LastInsertId()

// SELECT - 查询单行
var user User
err := db.QueryRow(
    "SELECT id, username, email FROM users WHERE id = ?",
    id,
).Scan(&user.ID, &user.Username, &user.Email)

// SELECT - 查询多行
rows, err := db.Query("SELECT id, username, email FROM users")
defer rows.Close()
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Username, &u.Email)
}

// UPDATE - 更新数据
result, err := db.Exec(
    "UPDATE users SET age = ? WHERE id = ?",
    newAge, id,
)
rowsAffected, _ := result.RowsAffected()

// DELETE - 删除数据
result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
```

### 3. 事务处理

```go
// 开始事务
tx, err := db.Begin()
if err != nil {
    return err
}

// 执行SQL语句
_, err = tx.Exec("INSERT INTO users (username) VALUES (?)", "张三")
if err != nil {
    tx.Rollback() // 出错回滚
    return err
}

// 设置保存点
tx.Exec("SAVEPOINT savepoint1")

// 提交事务
err = tx.Commit()
if err != nil {
    tx.Rollback()
    return err
}
```

### 4. 预处理语句

```go
// 准备语句
stmt, err := db.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
if err != nil {
    return err
}
defer stmt.Close()

// 执行预处理语句
result, err := stmt.Exec("用户名", "email@example.com")

// 预处理查询
queryStmt, err := db.Prepare("SELECT id, username FROM users WHERE age > ?")
rows, err := queryStmt.Query(25)
```

### 5. 表连接

```go
// INNER JOIN
rows, err := db.Query(`
    SELECT u.username, a.title
    FROM users u
    INNER JOIN articles a ON u.id = a.user_id
`)

// LEFT JOIN
rows, err := db.Query(`
    SELECT u.username, COUNT(a.id) as article_count
    FROM users u
    LEFT JOIN articles a ON u.id = a.user_id
    GROUP BY u.id
`)
```

### 6. 聚合函数

```go
// 基本聚合
var count int
db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

var avgAge float64
db.QueryRow("SELECT AVG(age) FROM users").Scan(&avgAge)

// GROUP BY 聚合
rows, err := db.Query(`
    SELECT age_group, COUNT(*) as count
    FROM users
    GROUP BY
        CASE
            WHEN age < 25 THEN 'young'
            ELSE 'adult'
        END
`)
```

### 7. 错误处理

```go
result, err := db.Exec("INSERT INTO users (username) VALUES (?)", username)
if err != nil {
    // 检查是否是约束错误
    if strings.Contains(err.Error(), "UNIQUE constraint") {
        return fmt.Errorf("用户名已存在")
    }
    return err
}
```

## 数据库驱动

### 常用数据库驱动

```bash
# SQLite
go get github.com/mattn/go-sqlite3

# MySQL
go get github.com/go-sql-driver/mysql

# PostgreSQL
go get github.com/lib/pq

# SQL Server
go get github.com/denisenkom/go-mssqldb

# Oracle
go get github.com/sijms/go-ora/v2
```

### 连接字符串示例

```go
// SQLite
db, err := sql.Open("sqlite3", "file:./data.db?cache=shared")

// MySQL
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8")

// PostgreSQL
db, err := sql.Open("postgres", "user=username password=password dbname=dbname host=localhost port=5432 sslmode=disable")

// SQL Server
db, err := sql.Open("sqlserver", "server=localhost;user id=user;password=password;database=dbname")
```

## 最佳实践

### 1. 连接池配置

```go
db.SetMaxOpenConns(25)        // 最大连接数
db.SetMaxIdleConns(5)         // 最大空闲连接数
db.SetConnMaxLifetime(5 * time.Minute) // 连接最大生存时间
```

### 2. 防止SQL注入

```go
// 正确 - 使用参数化查询
db.Query("SELECT * FROM users WHERE id = ?", userID)

// 错误 - 直接拼接SQL
db.Query("SELECT * FROM users WHERE id = " + userID)
```

### 3. 事务管理

```go
func updateUser(db *sql.DB, userID int, updateData map[string]interface{}) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        }
    }()

    // 执行多个操作...
    _, err = tx.Exec("UPDATE users SET name = ? WHERE id = ?", updateData["name"], userID)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}
```

### 4. 批量操作

```go
// 批量插入
stmt, err := db.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
if err != nil {
    return err
}
defer stmt.Close()

for _, user := range users {
    _, err := stmt.Exec(user.Username, user.Email)
    if err != nil {
        return err
    }
}
```

## 性能优化

### 1. 索引优化

```sql
-- 创建索引
CREATE INDEX idx_users_email ON users(email)
CREATE INDEX idx_articles_user_id ON articles(user_id)

-- 复合索引
CREATE INDEX idx_users_age_active ON users(age, is_active)
```

### 2. 查询优化

```go
// 使用LIMIT限制结果集
rows, err := db.Query("SELECT id, username FROM users LIMIT ? OFFSET ?", pageSize, offset)

// 只查询需要的列
db.QueryRow("SELECT id, username FROM users WHERE id = ?", id)

// 避免SELECT *
```

### 3. 连接池调优

```go
db.SetMaxOpenConns(100)        // 根据应用需求调整
db.SetMaxIdleConns(10)         // 保持一定的空闲连接
db.SetConnMaxLifetime(time.Hour) // 避免长时间连接
```

## 常见问题

### 1. 资源泄漏

```go
// 正确 - 确保关闭Rows
rows, err := db.Query("SELECT * FROM users")
if err != nil {
    return err
}
defer rows.Close()

for rows.Next() {
    // 处理数据
}
```

### 2. nil值处理

```go
// 使用sql.NullType处理可能的null值
var name sql.NullString
err := db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)

if name.Valid {
    fmt.Println(name.String) // 有值
} else {
    fmt.Println("name is NULL") // null值
}
```

### 3. 时间处理

```go
import "time"

// 插入时间
now := time.Now()
db.Exec("INSERT INTO logs (message, created_at) VALUES (?, ?)", message, now)

// 查询时间
var createdAt time.Time
err := db.QueryRow("SELECT created_at FROM logs WHERE id = ?", id).Scan(&createdAt)
```

## ORM框架

虽然本示例使用标准库，但实际项目中可以考虑使用ORM：

### 1. GORM

```go
import "gorm.io/gorm"

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Email    string
    Age      int
}

// 自动迁移
db.AutoMigrate(&User{})

// 创建记录
db.Create(&User{Username: "张三", Email: "zhangsan@example.com"})

// 查询记录
var user User
db.First(&user, 1)
```

### 2. SQLBoiler

```go
// 生成代码
sqlboiler sqlite3 --output models ./data.db

// 使用生成的代码
user, err := models.Users(models.UserWhere.Username.EQ("张三")).One(db)
```

## 测试

### 单元测试

```go
func TestUserRepository(t *testing.T) {
    // 使用内存数据库进行测试
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // 初始化测试数据
    initTestDB(db)

    // 执行测试...
}
```

## 练习

1. 实现一个完整的用户管理系统
2. 添加分页查询功能
3. 实现软删除机制
4. 添加数据库迁移功能
5. 实现连接池监控
6. 添加查询性能分析
7. 实现读写分离
8. 添加数据库备份功能