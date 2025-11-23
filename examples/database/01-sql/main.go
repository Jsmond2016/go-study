// Package main 展示 Go 语言中 SQL 数据库操作的基础用法
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite 驱动
)

// User 用户结构体
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Article 文章结构体
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	ViewCount int       `json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	fmt.Println("=== Go SQL 数据库操作示例 ===")

	// 连接数据库
	db := connectDatabase()
	defer db.Close()

	// 初始化数据库表
	initDatabase(db)

	// 演示各种SQL操作
	demonstrateSQLBasics(db)
	demonstrateTransactions(db)
	demonstratePreparedStatements(db)
	demonstrateJoins(db)
	demonstrateAggregations(db)

	fmt.Println("所有数据库操作完成！")
}

// connectDatabase 连接数据库
func connectDatabase() *sql.DB {
	// 使用SQLite作为示例数据库
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("数据库连接测试失败:", err)
	}

	fmt.Println("数据库连接成功")
	return db
}

// initDatabase 初始化数据库表
func initDatabase(db *sql.DB) {
	fmt.Println("\n--- 初始化数据库表 ---")

	// 创建用户表
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		age INTEGER CHECK(age >= 0),
		is_active BOOLEAN DEFAULT 1,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatal("创建用户表失败:", err)
	}
	fmt.Println("用户表创建成功")

	// 创建文章表
	createArticlesTable := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(200) NOT NULL,
		content TEXT,
		user_id INTEGER NOT NULL,
		view_count INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)`

	_, err = db.Exec(createArticlesTable)
	if err != nil {
		log.Fatal("创建文章表失败:", err)
	}
	fmt.Println("文章表创建成功")

	// 创建索引
	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)")
	if err != nil {
		log.Fatal("创建索引失败:", err)
	}
	fmt.Println("索引创建成功")
}

// demonstrateSQLBasics 演示SQL基础操作
func demonstrateSQLBasics(db *sql.DB) {
	fmt.Println("\n--- SQL基础操作演示 ---")

	// 清理现有数据
	db.Exec("DELETE FROM articles")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM sqlite_sequence WHERE name IN ('users', 'articles')")

	// 1. INSERT - 插入数据
	fmt.Println("1. 插入数据:")
	users := []User{
		{Username: "张三", Email: "zhangsan@example.com", Age: 25, IsActive: true},
		{Username: "李四", Email: "lisi@example.com", Age: 30, IsActive: true},
		{Username: "王五", Email: "wangwu@example.com", Age: 28, IsActive: false},
	}

	for _, user := range users {
		result, err := db.Exec(
			"INSERT INTO users (username, email, age, is_active) VALUES (?, ?, ?, ?)",
			user.Username, user.Email, user.Age, user.IsActive,
		)
		if err != nil {
			log.Printf("插入用户失败: %v", err)
			continue
		}
		id, _ := result.LastInsertId()
		fmt.Printf("  插入用户: %s (ID: %d)\n", user.Username, id)
	}

	// 2. SELECT - 查询数据
	fmt.Println("\n2. 查询数据:")
	queryUsers(db, "SELECT id, username, email, age, is_active, created_at, updated_at FROM users")

	// 3. UPDATE - 更新数据
	fmt.Println("\n3. 更新数据:")
	result, err := db.Exec("UPDATE users SET age = ?, updated_at = CURRENT_TIMESTAMP WHERE username = ?", 26, "张三")
	if err != nil {
		log.Printf("更新失败: %v", err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("  更新了 %d 行数据\n", rowsAffected)
	}

	// 4. DELETE - 删除数据
	fmt.Println("\n4. 删除数据:")
	result, err = db.Exec("DELETE FROM users WHERE is_active = ?", false)
	if err != nil {
		log.Printf("删除失败: %v", err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("  删除了 %d 行数据\n", rowsAffected)
	}

	// 5. 查询更新后的数据
	fmt.Println("\n5. 更新后的用户列表:")
	queryUsers(db, "SELECT id, username, email, age, is_active, created_at, updated_at FROM users")

	// 添加文章数据
	fmt.Println("\n6. 添加文章数据:")
	articles := []Article{
		{Title: "Go语言入门", Content: "Go是一门优秀的编程语言...", UserID: 1, ViewCount: 100},
		{Title: "数据库基础", Content: "数据库是现代应用的核心...", UserID: 1, ViewCount: 50},
		{Title: "Web开发实践", Content: "Web开发需要掌握多种技术...", UserID: 2, ViewCount: 200},
	}

	for _, article := range articles {
		result, err := db.Exec(
			"INSERT INTO articles (title, content, user_id, view_count) VALUES (?, ?, ?, ?)",
			article.Title, article.Content, article.UserID, article.ViewCount,
		)
		if err != nil {
			log.Printf("插入文章失败: %v", err)
			continue
		}
		id, _ := result.LastInsertId()
		fmt.Printf("  插入文章: %s (ID: %d)\n", article.Title, id)
	}
}

// queryUsers 查询用户并打印
func queryUsers(db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("  用户列表:")
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		fmt.Printf("    ID: %d, 用户名: %s, 邮箱: %s, 年龄: %d, 状态: %t\n",
			user.ID, user.Username, user.Email, user.Age, user.IsActive)
	}
}

// demonstrateTransactions 演示事务处理
func demonstrateTransactions(db *sql.DB) {
	fmt.Println("\n--- 事务处理演示 ---")

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		log.Printf("开始事务失败: %v", err)
		return
	}

	// 设置事务保存点
	savepoint, err := tx.Exec("SAVEPOINT user_insert")
	if err != nil {
		log.Printf("创建保存点失败: %v", err)
		tx.Rollback()
		return
	}

	fmt.Println("1. 在事务中插入新用户:")
	result, err := tx.Exec("INSERT INTO users (username, email, age) VALUES (?, ?, ?)", "新用户", "new@example.com", 22)
	if err != nil {
		log.Printf("插入用户失败: %v", err)
		tx.Rollback()
		return
	}
	userID, _ := result.LastInsertId()
	fmt.Printf("  插入用户ID: %d\n", userID)

	fmt.Println("2. 为新用户插入文章:")
	_, err = tx.Exec("INSERT INTO articles (title, content, user_id) VALUES (?, ?, ?)", "我的第一篇文章", "这是文章内容...", userID)
	if err != nil {
		log.Printf("插入文章失败: %v", err)
		tx.Rollback()
		return
	}
	fmt.Println("  插入文章成功")

	fmt.Println("3. 模拟错误，回滚到保存点:")
	_, err = tx.Exec("INSERT INTO users (username, email, age) VALUES (?, ?, ?)", "重复用户", "new@example.com", 25)
	if err != nil {
		fmt.Printf("  预期错误: %v\n", err)
		// 回滚到保存点
		_, err = tx.Exec("ROLLBACK TO SAVEPOINT user_insert")
		if err != nil {
			log.Printf("回滚到保存点失败: %v", err)
			tx.Rollback()
			return
		}
		fmt.Println("  已回滚到保存点")
	}

	fmt.Println("4. 提交事务:")
	err = tx.Commit()
	if err != nil {
		log.Printf("提交事务失败: %v", err)
		return
	}
	fmt.Println("  事务提交成功")

	// 查询结果
	fmt.Println("\n5. 事务结果查询:")
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "新用户").Scan(&count)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("  新用户存在: %t\n", count > 0)

	err = db.QueryRow("SELECT COUNT(*) FROM articles WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("  用户文章存在: %t\n", count > 0)
}

// demonstratePreparedStatements 演示预处理语句
func demonstratePreparedStatements(db *sql.DB) {
	fmt.Println("\n--- 预处理语句演示 ---")

	// 1. 准备插入语句
	insertStmt, err := db.Prepare("INSERT INTO users (username, email, age) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("准备插入语句失败: %v", err)
		return
	}
	defer insertStmt.Close()

	fmt.Println("1. 使用预处理语句插入多个用户:")
	newUsers := []struct {
		Username string
		Email    string
		Age      int
	}{
		{"赵六", "zhaoliu@example.com", 32},
		{"钱七", "qianqi@example.com", 27},
		{"孙八", "sunba@example.com", 29},
	}

	for _, user := range newUsers {
		result, err := insertStmt.Exec(user.Username, user.Email, user.Age)
		if err != nil {
			log.Printf("插入用户 %s 失败: %v", user.Username, err)
			continue
		}
		id, _ := result.LastInsertId()
		fmt.Printf("  插入用户: %s (ID: %d)\n", user.Username, id)
	}

	// 2. 准备查询语句
	fmt.Println("\n2. 使用预处理语句查询用户:")
	selectStmt, err := db.Prepare("SELECT id, username, email, age FROM users WHERE age > ? ORDER BY age")
	if err != nil {
		log.Printf("准备查询语句失败: %v", err)
		return
	}
	defer selectStmt.Close()

	rows, err := selectStmt.Query(25)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("  年龄大于25的用户:")
	for rows.Next() {
		var id, age int
		var username, email string
		err := rows.Scan(&id, &username, &email, &age)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		fmt.Printf("    %s (%d岁)\n", username, age)
	}

	// 3. 准备更新语句
	fmt.Println("\n3. 使用预处理语句批量更新:")
	updateStmt, err := db.Prepare("UPDATE users SET age = age + ? WHERE id = ?")
	if err != nil {
		log.Printf("准备更新语句失败: %v", err)
		return
	}
	defer updateStmt.Close()

	updates := []struct {
		ID     int
		AgeAdd int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	for _, update := range updates {
		result, err := updateStmt.Exec(update.AgeAdd, update.ID)
		if err != nil {
			log.Printf("更新用户 %d 失败: %v", update.ID, err)
			continue
		}
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("  更新用户 %d，影响行数: %d\n", update.ID, rowsAffected)
	}
}

// demonstrateJoins 演示表连接
func demonstrateJoins(db *sql.DB) {
	fmt.Println("\n--- 表连接演示 ---")

	// 1. INNER JOIN
	fmt.Println("1. INNER JOIN - 用户及其文章:")
	rows, err := db.Query(`
		SELECT u.id, u.username, a.id as article_id, a.title, a.view_count
		FROM users u
		INNER JOIN articles a ON u.id = a.user_id
		ORDER BY u.id, a.id
	`)
	if err != nil {
		log.Printf("INNER JOIN 查询失败: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userID int
		var username, title string
		var articleID, viewCount int
		err := rows.Scan(&userID, &username, &articleID, &title, &viewCount)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		fmt.Printf("  %s: %s (浏览: %d)\n", username, title, viewCount)
	}

	// 2. LEFT JOIN
	fmt.Println("\n2. LEFT JOIN - 所有用户及其文章数量:")
	rows, err = db.Query(`
		SELECT u.id, u.username, COUNT(a.id) as article_count
		FROM users u
		LEFT JOIN articles a ON u.id = a.user_id
		GROUP BY u.id, u.username
		ORDER BY u.id
	`)
	if err != nil {
		log.Printf("LEFT JOIN 查询失败: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, articleCount int
		var username string
		err := rows.Scan(&id, &username, &articleCount)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		fmt.Printf("  %s: %d 篇文章\n", username, articleCount)
	}

	// 3. 子查询
	fmt.Println("\n3. 子查询 - 发表文章最多的用户:")
	err = db.QueryRow(`
		SELECT u.username, COUNT(a.id) as article_count
		FROM users u
		WHERE u.id = (
			SELECT user_id
			FROM articles
			GROUP BY user_id
			ORDER BY COUNT(*) DESC
			LIMIT 1
		)
		LEFT JOIN articles a ON u.id = a.user_id
		GROUP BY u.username
	`).Scan(&username, &articleCount)

	if err != nil {
		log.Printf("子查询失败: %v", err)
	} else {
		fmt.Printf("  %s: %d 篇文章\n", username, articleCount)
	}
}

// demonstrateAggregations 演示聚合函数
func demonstrateAggregations(db *sql.DB) {
	fmt.Println("\n--- 聚合函数演示 ---")

	// 1. 基本聚合
	fmt.Println("1. 基本聚合函数:")
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	fmt.Printf("  用户总数: %d\n", count)

	var avgAge float64
	db.QueryRow("SELECT AVG(age) FROM users").Scan(&avgAge)
	fmt.Printf("  平均年龄: %.1f\n", avgAge)

	var maxAge, minAge int
	db.QueryRow("SELECT MAX(age), MIN(age) FROM users").Scan(&maxAge, &minAge)
	fmt.Printf("  最大年龄: %d, 最小年龄: %d\n", maxAge, minAge)

	var totalViews int
	db.QueryRow("SELECT SUM(view_count) FROM articles").Scan(&totalViews)
	fmt.Printf("  文章总浏览量: %d\n", totalViews)

	// 2. GROUP BY 聚合
	fmt.Println("\n2. GROUP BY 聚合:")
	rows, err := db.Query(`
		SELECT
			CASE
				WHEN age < 25 THEN '年轻'
				WHEN age BETWEEN 25 AND 30 THEN '中年'
				ELSE '成熟'
			END as age_group,
			COUNT(*) as user_count,
			AVG(age) as avg_age
		FROM users
		GROUP BY age_group
		ORDER BY avg_age
	`)
	if err != nil {
		log.Printf("GROUP BY 查询失败: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ageGroup string
		var userCount int
		var avgAge float64
		err := rows.Scan(&ageGroup, &userCount, &avgAge)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		fmt.Printf("  %s: %d 人, 平均年龄: %.1f\n", ageGroup, userCount, avgAge)
	}

	// 3. HAVING 过滤
	fmt.Println("\n3. HAVING 过滤:")
	rows, err = db.Query(`
		SELECT u.id, u.username, COUNT(a.id) as article_count
		FROM users u
		LEFT JOIN articles a ON u.id = a.user_id
		GROUP BY u.id, u.username
		HAVING COUNT(a.id) > 0
		ORDER BY article_count DESC
	`)
	if err != nil {
		log.Printf("HAVING 查询失败: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("  发表文章的用户:")
	for rows.Next() {
		var id, articleCount int
		var username string
		err := rows.Scan(&id, &username, &articleCount)
		if err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		fmt.Printf("    %s: %d 篇\n", username, articleCount)
	}
}

// 演示数据库元数据
func demonstrateMetadata(db *sql.DB) {
	fmt.Println("\n--- 数据库元数据演示 ---")

	// 获取表信息
	fmt.Println("1. 数据库表信息:")
	tables, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Printf("查询表信息失败: %v", err)
		return
	}
	defer tables.Close()

	for tables.Next() {
		var tableName string
		tables.Scan(&tableName)
		fmt.Printf("  表: %s\n", tableName)

		// 获取表结构
		columns, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s)", tableName))
		if err != nil {
			log.Printf("查询表结构失败: %v", err)
			continue
		}

		for columns.Next() {
			var cid int
			var name, dataType string
			var notNull, pk int
			var defaultValue interface{}
			columns.Scan(&cid, &name, &dataType, &notNull, &defaultValue, &pk)
			fmt.Printf("    %s: %s (NOT NULL: %t, PK: %t)\n", name, dataType, notNull == 1, pk == 1)
		}
		columns.Close()
	}
}

// 清理函数
func cleanupDatabase() {
	// 删除数据库文件
	if _, err := os.Stat("./data.db"); err == nil {
		os.Remove("./data.db")
		fmt.Println("数据库文件已清理")
	}
}