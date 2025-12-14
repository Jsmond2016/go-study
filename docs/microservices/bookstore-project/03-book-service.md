---
title: Chapter 3 - Book Service 实现
difficulty: intermediate
duration: "3-4小时"
tags: ["gRPC", "SQLite", "CRUD", "微服务实现"]
---

# Chapter 3: Book Service 实现

本章将从零开始实现 Book Service（图书服务），这是我们的第一个完整微服务。

## 📋 学习目标

完成本章后，你将：
- [x] 掌握 gRPC 服务端实现方法
- [x] 学会 SQLite 数据库操作
- [x] 实现完整的 CRUD 功能
- [x] 实现搜索和分页功能
- [x] 掌握错误处理最佳实践
- [x] 编写结构化日志

## 🏗️ Book Service 架构

### 服务职责

Book Service 负责：
1. **图书 CRUD** - 创建、读取、更新、删除图书
2. **图书搜索** - 按标题、作者、ISBN 搜索
3. **分页查询** - 支持大数据量分页
4. **库存管理** - 跟踪图书库存数量

### 代码结构

```
book-service/
└── main.go          # 所有代码（500+ 行）
    ├── 数据库初始化
    ├── gRPC 服务实现
    ├── CRUD 方法
    ├── 搜索方法
    └── 主函数
```

##完整代码实现

我们将完整实现 `book-service/main.go`，代码分为以下几个部分：

### 第一部分：包导入和常量定义

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"

	_ "github.com/mattn/go-sqlite3"  // SQLite 驱动
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "bookstore/proto/book"  // 导入生成的 proto 代码
)

const (
	port = ":50051"  // gRPC 服务端口
)
```

**要点说明**：
- `_ "github.com/mattn/go-sqlite3"` - 匿名导入，注册 SQLite 驱动
- `pb "bookstore/proto/book"` - 使用别名导入 proto 包
- 使用 `const` 定义常量，便于维护

### 第二部分：服务结构体定义

```go
// BookServer 实现 BookService gRPC 服务
type BookServer struct {
	pb.UnimplementedBookServiceServer  // 嵌入未实现的服务
	db *sql.DB                          // 数据库连接
}
```

**要点说明**：
- 嵌入 `UnimplementedBookServiceServer` 确保向后兼容
- 持有数据库连接 `*sql.DB`，所有方法共享

### 第三部分：主函数

```go
func main() {
	// 1. 初始化数据库
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// 2. 创建 gRPC 服务器
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &BookServer{db: db})

	log.Printf("Book Service is running on port %s", port)

	// 3. 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

**流程说明**：
1. 初始化数据库连接
2. 创建 TCP 监听器
3. 创建 gRPC 服务器并注册服务
4. 启动服务监听

### 第四部分：数据库初始化

```go
// initDatabase 初始化 SQLite 数据库
func initDatabase() (*sql.DB, error) {
	// 1. 打开数据库连接
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		return nil, err
	}

	// 2. 创建图书表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT NOT NULL UNIQUE,
		publisher TEXT,
		price REAL NOT NULL,
		stock INTEGER NOT NULL DEFAULT 0,
		description TEXT,
		cover_image TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	// 3. 创建索引（提高查询性能）
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);",
		"CREATE INDEX IF NOT EXISTS idx_books_author ON books(author);",
		"CREATE INDEX IF NOT EXISTS idx_books_isbn ON books(isbn);",
	}

	for _, idx := range indexes {
		_, err = db.Exec(idx)
		if err != nil {
			return nil, err
		}
	}

	// 4. 插入示例数据
	err = insertSampleData(db)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}
```

**设计亮点**：
- ✅ 使用 `IF NOT EXISTS` 避免重复创建
- ✅ 为常用查询字段创建索引
- ✅ 自动插入示例数据方便测试

### 第五部分：插入示例数据

```go
// insertSampleData 插入示例图书数据
func insertSampleData(db *sql.DB) error {
	// 检查是否已有数据
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Sample data already exists, skipping...")
		return nil
	}

	// 示例图书数据
	books := []struct {
		title       string
		author      string
		isbn        string
		publisher   string
		price       float64
		stock       int
		description string
		coverImage  string
	}{
		{
			"Go语言圣经",
			"Alan A. A. Donovan",
			"978-0134190440",
			"Addison-Wesley",
			89.00,
			100,
			"Go语言权威指南，适合所有Go开发者",
			"https://example.com/covers/go-bible.jpg",
		},
		{
			"深入理解计算机系统",
			"Randal E. Bryant",
			"978-0136108040",
			"机械工业出版社",
			139.00,
			50,
			"经典计算机系统教材",
			"https://example.com/covers/csapp.jpg",
		},
		// ... 更多示例数据
	}

	// 批量插入
	for _, book := range books {
		_, err := db.Exec(`
			INSERT INTO books (title, author, isbn, publisher, price, stock, description, cover_image)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			book.title, book.author, book.isbn, book.publisher,
			book.price, book.stock, book.description, book.coverImage,
		)
		if err != nil {
			return err
		}
	}

	log.Printf("Inserted %d sample books", len(books))
	return nil
}
```

### 第六部分：GetBook - 获取单个图书

```go
// GetBook 根据 ID 获取图书信息
func (s *BookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.BookResponse, error) {
	log.Printf("GetBook called: id=%d", req.Id)

	// 1. 参数验证
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid book ID")
	}

	// 2. 查询数据库
	var book pb.Book
	var createdAt, updatedAt string

	err := s.db.QueryRow(`
		SELECT id, title, author, isbn, publisher, price, stock,
		       description, cover_image, created_at, updated_at
		FROM books WHERE id = ?
	`, req.Id).Scan(
		&book.Id, &book.Title, &book.Author, &book.Isbn,
		&book.Publisher, &book.Price, &book.Stock,
		&book.Description, &book.CoverImage,
		&createdAt, &updatedAt,
	)

	// 3. 错误处理
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "book not found")
	}
	if err != nil {
		log.Printf("Error querying book: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	// 4. 设置时间字段
	book.CreatedAt = createdAt
	book.UpdatedAt = updatedAt

	log.Printf("Book found: %s by %s", book.Title, book.Author)

	return &pb.BookResponse{Book: &book}, nil
}
```

**要点说明**：
- ✅ 参数验证在最前面
- ✅ 使用 gRPC status codes（NotFound, Internal 等）
- ✅ 记录日志方便调试
- ✅ sql.ErrNoRows 返回 NotFound 而不是 Internal Error

### 第七部分：ListBooks - 分页查询

```go
// ListBooks 获取图书列表（分页）
func (s *BookServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	log.Printf("ListBooks called: page=%d, page_size=%d", req.Page, req.PageSize)

	// 1. 设置默认值和验证
	page := req.Page
	pageSize := req.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // 限制最大页面大小
	}

	// 2. 计算 offset
	offset := (page - 1) * pageSize

	// 3. 查询总数
	var total int32
	err := s.db.QueryRow("SELECT COUNT(*) FROM books").Scan(&total)
	if err != nil {
		log.Printf("Error counting books: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	// 4. 查询分页数据
	rows, err := s.db.Query(`
		SELECT id, title, author, isbn, publisher, price, stock,
		       description, cover_image, created_at, updated_at
		FROM books
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`, pageSize, offset)

	if err != nil {
		log.Printf("Error querying books: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	defer rows.Close()

	// 5. 遍历结果集
	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		var createdAt, updatedAt string

		err := rows.Scan(
			&book.Id, &book.Title, &book.Author, &book.Isbn,
			&book.Publisher, &book.Price, &book.Stock,
			&book.Description, &book.CoverImage,
			&createdAt, &updatedAt,
		)
		if err != nil {
			log.Printf("Error scanning book: %v", err)
			continue
		}

		book.CreatedAt = createdAt
		book.UpdatedAt = updatedAt
		books = append(books, &book)
	}

	log.Printf("Returned %d books (total: %d)", len(books), total)

	return &pb.ListBooksResponse{
		Books: books,
		Total: total,
		Page:  page,
	}, nil
}
```

**分页实现要点**：
- ✅ 设置合理的默认值（page=1, pageSize=10）
- ✅ 限制最大页面大小防止内存溢出
- ✅ 先查询总数，再查询分页数据
- ✅ 使用 `LIMIT` 和 `OFFSET` 实现分页
- ✅ 按 ID 降序排列（最新的在前面）

### 第八部分：CreateBook - 创建图书

```go
// CreateBook 创建新图书
func (s *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	log.Printf("CreateBook called: title=%s, author=%s", req.Title, req.Author)

	// 1. 参数验证
	if req.Title == "" || req.Author == "" || req.Isbn == "" {
		return nil, status.Error(codes.InvalidArgument, "title, author and ISBN are required")
	}

	if req.Price <= 0 {
		return nil, status.Error(codes.InvalidArgument, "price must be greater than 0")
	}

	if req.Stock < 0 {
		return nil, status.Error(codes.InvalidArgument, "stock cannot be negative")
	}

	// 2. 检查 ISBN 是否已存在
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE isbn = ?)", req.Isbn).Scan(&exists)
	if err != nil {
		log.Printf("Error checking ISBN: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "book with this ISBN already exists")
	}

	// 3. 插入数据
	result, err := s.db.Exec(`
		INSERT INTO books (title, author, isbn, publisher, price, stock, description, cover_image)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		req.Title, req.Author, req.Isbn, req.Publisher,
		req.Price, req.Stock, req.Description, req.CoverImage,
	)

	if err != nil {
		log.Printf("Error creating book: %v", err)
		return nil, status.Error(codes.Internal, "failed to create book")
	}

	// 4. 获取插入的 ID
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return nil, status.Error(codes.Internal, "failed to get book ID")
	}

	// 5. 查询并返回创建的图书
	return s.GetBook(ctx, &pb.GetBookRequest{Id: id})
}
```

**创建操作要点**：
- ✅ 严格的参数验证
- ✅ 检查唯一约束（ISBN）避免重复
- ✅ 使用 `LastInsertId()` 获取新记录 ID
- ✅ 返回完整的创建结果（复用 GetBook）

### 第九部分：UpdateBook - 更新图书

```go
// UpdateBook 更新图书信息
func (s *BookServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.BookResponse, error) {
	log.Printf("UpdateBook called: id=%d", req.Id)

	// 1. 参数验证
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid book ID")
	}

	// 2. 检查图书是否存在
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?)", req.Id).Scan(&exists)
	if err != nil {
		log.Printf("Error checking book: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if !exists {
		return nil, status.Error(codes.NotFound, "book not found")
	}

	// 3. 构建动态 UPDATE 语句
	query := "UPDATE books SET updated_at = CURRENT_TIMESTAMP"
	args := []interface{}{}

	if req.Title != "" {
		query += ", title = ?"
		args = append(args, req.Title)
	}
	if req.Author != "" {
		query += ", author = ?"
		args = append(args, req.Author)
	}
	if req.Publisher != "" {
		query += ", publisher = ?"
		args = append(args, req.Publisher)
	}
	if req.Price > 0 {
		query += ", price = ?"
		args = append(args, req.Price)
	}
	if req.Stock >= 0 {
		query += ", stock = ?"
		args = append(args, req.Stock)
	}
	if req.Description != "" {
		query += ", description = ?"
		args = append(args, req.Description)
	}
	if req.CoverImage != "" {
		query += ", cover_image = ?"
		args = append(args, req.CoverImage)
	}

	query += " WHERE id = ?"
	args = append(args, req.Id)

	// 4. 执行更新
	_, err = s.db.Exec(query, args...)
	if err != nil {
		log.Printf("Error updating book: %v", err)
		return nil, status.Error(codes.Internal, "failed to update book")
	}

	log.Printf("Book %d updated successfully", req.Id)

	// 5. 返回更新后的图书
	return s.GetBook(ctx, &pb.GetBookRequest{Id: req.Id})
}
```

**更新操作要点**：
- ✅ 动态构建 SQL（只更新提供的字段）
- ✅ 自动更新 `updated_at` 时间戳
- ✅ 先检查存在性再更新
- ✅ 返回更新后的完整数据

### 第十部分：DeleteBook - 删除图书

```go
// DeleteBook 删除图书
func (s *BookServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	log.Printf("DeleteBook called: id=%d", req.Id)

	// 1. 参数验证
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid book ID")
	}

	// 2. 执行删除
	result, err := s.db.Exec("DELETE FROM books WHERE id = ?", req.Id)
	if err != nil {
		log.Printf("Error deleting book: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete book")
	}

	// 3. 检查是否实际删除了记录
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "book not found")
	}

	log.Printf("Book %d deleted successfully", req.Id)

	return &pb.DeleteBookResponse{
		Message: fmt.Sprintf("Book %d deleted successfully", req.Id),
	}, nil
}
```

**删除操作要点**：
- ✅ 使用 `RowsAffected()` 判断是否真正删除
- ✅ 未找到记录返回 NotFound 而不是成功
- ✅ 返回友好的成功消息

### 第十一部分：SearchBooks - 搜索图书

```go
// SearchBooks 搜索图书（按标题、作者、ISBN）
func (s *BookServer) SearchBooks(ctx context.Context, req *pb.SearchBooksRequest) (*pb.ListBooksResponse, error) {
	log.Printf("SearchBooks called: keyword=%s, page=%d", req.Keyword, req.Page)

	// 1. 参数验证和默认值
	if req.Keyword == "" {
		return nil, status.Error(codes.InvalidArgument, "keyword is required")
	}

	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// 2. 构建搜索模式
	searchPattern := "%" + strings.ToLower(req.Keyword) + "%"

	// 3. 查询匹配的总数
	var total int32
	err := s.db.QueryRow(`
		SELECT COUNT(*)
		FROM books
		WHERE LOWER(title) LIKE ? OR LOWER(author) LIKE ? OR LOWER(isbn) LIKE ?
	`, searchPattern, searchPattern, searchPattern).Scan(&total)

	if err != nil {
		log.Printf("Error counting search results: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	// 4. 查询分页结果
	rows, err := s.db.Query(`
		SELECT id, title, author, isbn, publisher, price, stock,
		       description, cover_image, created_at, updated_at
		FROM books
		WHERE LOWER(title) LIKE ? OR LOWER(author) LIKE ? OR LOWER(isbn) LIKE ?
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`, searchPattern, searchPattern, searchPattern, pageSize, offset)

	if err != nil {
		log.Printf("Error searching books: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	defer rows.Close()

	// 5. 处理结果集
	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		var createdAt, updatedAt string

		err := rows.Scan(
			&book.Id, &book.Title, &book.Author, &book.Isbn,
			&book.Publisher, &book.Price, &book.Stock,
			&book.Description, &book.CoverImage,
			&createdAt, &updatedAt,
		)
		if err != nil {
			log.Printf("Error scanning book: %v", err)
			continue
		}

		book.CreatedAt = createdAt
		book.UpdatedAt = updatedAt
		books = append(books, &book)
	}

	log.Printf("Found %d books matching '%s'", len(books), req.Keyword)

	return &pb.ListBooksResponse{
		Books: books,
		Total: total,
		Page:  page,
	}, nil
}
```

**搜索功能要点**：
- ✅ 使用 `LIKE` 进行模糊搜索
- ✅ 使用 `LOWER()` 实现大小写不敏感
- ✅ 在标题、作者、ISBN 三个字段中搜索
- ✅ 支持分页
- ✅ 使用 `%keyword%` 匹配任意位置

## 🧪 测试 Book Service

### 启动服务

```bash
cd book-service
go run main.go
```

预期输出：
```
2024/01/15 10:00:00 Database initialized successfully
2024/01/15 10:00:00 Inserted 10 sample books
2024/01/15 10:00:00 Book Service is running on port :50051
```

### 使用 grpcurl 测试

#### 1. 获取图书详情

```bash
grpcurl -plaintext -d '{"id": 1}' localhost:50051 book.BookService/GetBook
```

#### 2. 获取图书列表

```bash
grpcurl -plaintext -d '{"page": 1, "page_size": 5}' \
  localhost:50051 book.BookService/ListBooks
```

#### 3. 创建图书

```bash
grpcurl -plaintext -d '{
  "title": "测试图书",
  "author": "测试作者",
  "isbn": "978-1234567890",
  "publisher": "测试出版社",
  "price": 99.99,
  "stock": 50
}' localhost:50051 book.BookService/CreateBook
```

#### 4. 搜索图书

```bash
grpcurl -plaintext -d '{"keyword": "Go", "page": 1, "page_size": 10}' \
  localhost:50051 book.BookService/SearchBooks
```

## ✅ 本章检查清单

- [ ] 创建了 `book-service/main.go`
- [ ] 理解了 gRPC 服务实现方法
- [ ] 掌握了 SQLite CRUD 操作
- [ ] 实现了分页查询
- [ ] 实现了搜索功能
- [ ] 服务能够正常启动
- [ ] 使用 grpcurl 测试成功

## 🎯 下一步

Book Service 完成！接下来实现用户服务：

**[👉 Chapter 4: User Service](/microservices/bookstore-project/04-user-service)**

下一章我们将学习：
- 用户注册和登录
- 密码加密（bcrypt）
- JWT Token 生成和验证
- 安全最佳实践

---

**🎉 恭喜完成第一个微服务！**
