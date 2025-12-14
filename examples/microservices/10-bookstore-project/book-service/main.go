package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "bookstore/proto/book"
)

var (
	port = flag.Int("port", 50051, "服务端口")
)

// BookServer 实现 BookService
type BookServer struct {
	pb.UnimplementedBookServiceServer
	db *sql.DB
}

// NewBookServer 创建图书服务
func NewBookServer(db *sql.DB) *BookServer {
	return &BookServer{db: db}
}

// GetBook 获取图书详情
func (s *BookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.BookResponse, error) {
	log.Printf("GetBook: id=%d", req.Id)

	query := `SELECT id, title, author, isbn, publisher, price, stock, category, description, cover_url, created_at, updated_at
	          FROM books WHERE id = ?`

	var book pb.Book
	var createdAt, updatedAt int64

	err := s.db.QueryRowContext(ctx, query, req.Id).Scan(
		&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Publisher,
		&book.Price, &book.Stock, &book.Category, &book.Description,
		&book.CoverUrl, &createdAt, &updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "图书不存在")
	}
	if err != nil {
		log.Printf("查询失败: %v", err)
		return nil, status.Error(codes.Internal, "查询失败")
	}

	book.CreatedAt = createdAt
	book.UpdatedAt = updatedAt

	return &pb.BookResponse{Book: &book}, nil
}

// ListBooks 获取图书列表
func (s *BookServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	log.Printf("ListBooks: page=%d, pageSize=%d, category=%s", req.Page, req.PageSize, req.Category)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize

	var query string
	var args []interface{}

	if req.Category != "" {
		query = `SELECT id, title, author, isbn, publisher, price, stock, category, description, cover_url, created_at, updated_at
		         FROM books WHERE category = ? ORDER BY id DESC LIMIT ? OFFSET ?`
		args = []interface{}{req.Category, req.PageSize, offset}
	} else {
		query = `SELECT id, title, author, isbn, publisher, price, stock, category, description, cover_url, created_at, updated_at
		         FROM books ORDER BY id DESC LIMIT ? OFFSET ?`
		args = []interface{}{req.PageSize, offset}
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return nil, status.Error(codes.Internal, "查询失败")
	}
	defer rows.Close()

	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		var createdAt, updatedAt int64

		err := rows.Scan(
			&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Publisher,
			&book.Price, &book.Stock, &book.Category, &book.Description,
			&book.CoverUrl, &createdAt, &updatedAt,
		)
		if err != nil {
			log.Printf("扫描失败: %v", err)
			continue
		}

		book.CreatedAt = createdAt
		book.UpdatedAt = updatedAt
		books = append(books, &book)
	}

	// 查询总数
	var total int32
	countQuery := "SELECT COUNT(*) FROM books"
	if req.Category != "" {
		countQuery += " WHERE category = ?"
		s.db.QueryRowContext(ctx, countQuery, req.Category).Scan(&total)
	} else {
		s.db.QueryRowContext(ctx, countQuery).Scan(&total)
	}

	return &pb.ListBooksResponse{
		Books:    books,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// CreateBook 创建图书
func (s *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	log.Printf("CreateBook: title=%s, author=%s", req.Title, req.Author)

	// 验证必填字段
	if req.Title == "" || req.Author == "" {
		return nil, status.Error(codes.InvalidArgument, "书名和作者不能为空")
	}

	now := time.Now().Unix()
	query := `INSERT INTO books (title, author, isbn, publisher, price, stock, category, description, cover_url, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := s.db.ExecContext(ctx, query,
		req.Title, req.Author, req.Isbn, req.Publisher, req.Price, req.Stock,
		req.Category, req.Description, req.CoverUrl, now, now,
	)
	if err != nil {
		log.Printf("插入失败: %v", err)
		return nil, status.Error(codes.Internal, "创建失败")
	}

	id, _ := result.LastInsertId()

	book := &pb.Book{
		Id:          id,
		Title:       req.Title,
		Author:      req.Author,
		Isbn:        req.Isbn,
		Publisher:   req.Publisher,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Description: req.Description,
		CoverUrl:    req.CoverUrl,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return &pb.BookResponse{Book: book}, nil
}

// UpdateBook 更新图书
func (s *BookServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.BookResponse, error) {
	log.Printf("UpdateBook: id=%d", req.Id)

	now := time.Now().Unix()
	query := `UPDATE books SET title=?, author=?, publisher=?, price=?, stock=?, category=?, description=?, cover_url=?, updated_at=?
	          WHERE id=?`

	result, err := s.db.ExecContext(ctx, query,
		req.Title, req.Author, req.Publisher, req.Price, req.Stock,
		req.Category, req.Description, req.CoverUrl, now, req.Id,
	)
	if err != nil {
		log.Printf("更新失败: %v", err)
		return nil, status.Error(codes.Internal, "更新失败")
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return nil, status.Error(codes.NotFound, "图书不存在")
	}

	// 返回更新后的图书
	return s.GetBook(ctx, &pb.GetBookRequest{Id: req.Id})
}

// DeleteBook 删除图书
func (s *BookServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	log.Printf("DeleteBook: id=%d", req.Id)

	query := "DELETE FROM books WHERE id=?"
	result, err := s.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		log.Printf("删除失败: %v", err)
		return nil, status.Error(codes.Internal, "删除失败")
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return &pb.DeleteBookResponse{
			Success: false,
			Message: "图书不存在",
		}, nil
	}

	return &pb.DeleteBookResponse{
		Success: true,
		Message: "删除成功",
	}, nil
}

// SearchBooks 搜索图书
func (s *BookServer) SearchBooks(ctx context.Context, req *pb.SearchBooksRequest) (*pb.ListBooksResponse, error) {
	log.Printf("SearchBooks: keyword=%s", req.Keyword)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize
	keyword := "%" + req.Keyword + "%"

	query := `SELECT id, title, author, isbn, publisher, price, stock, category, description, cover_url, created_at, updated_at
	          FROM books WHERE title LIKE ? OR author LIKE ? OR description LIKE ?
	          ORDER BY id DESC LIMIT ? OFFSET ?`

	rows, err := s.db.QueryContext(ctx, query, keyword, keyword, keyword, req.PageSize, offset)
	if err != nil {
		log.Printf("搜索失败: %v", err)
		return nil, status.Error(codes.Internal, "搜索失败")
	}
	defer rows.Close()

	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		var createdAt, updatedAt int64

		err := rows.Scan(
			&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Publisher,
			&book.Price, &book.Stock, &book.Category, &book.Description,
			&book.CoverUrl, &createdAt, &updatedAt,
		)
		if err != nil {
			continue
		}

		book.CreatedAt = createdAt
		book.UpdatedAt = updatedAt
		books = append(books, &book)
	}

	// 查询总数
	var total int32
	countQuery := "SELECT COUNT(*) FROM books WHERE title LIKE ? OR author LIKE ? OR description LIKE ?"
	s.db.QueryRowContext(ctx, countQuery, keyword, keyword, keyword).Scan(&total)

	return &pb.ListBooksResponse{
		Books:    books,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// initDB 初始化数据库
func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./bookstore.db")
	if err != nil {
		return nil, err
	}

	// 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT,
		publisher TEXT,
		price REAL DEFAULT 0,
		stock INTEGER DEFAULT 0,
		category TEXT,
		description TEXT,
		cover_url TEXT,
		created_at INTEGER,
		updated_at INTEGER
	);

	CREATE INDEX IF NOT EXISTS idx_category ON books(category);
	CREATE INDEX IF NOT EXISTS idx_title ON books(title);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	// 插入示例数据
	insertSampleData(db)

	return db, nil
}

// insertSampleData 插入示例数据
func insertSampleData(db *sql.DB) {
	// 检查是否已有数据
	var count int
	db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	if count > 0 {
		return
	}

	log.Println("插入示例数据...")

	now := time.Now().Unix()
	books := []struct {
		title, author, isbn, publisher, category, description string
		price                                                  float64
		stock                                                  int
	}{
		{"Go 语言编程", "许式伟", "9787115290366", "人民邮电出版社", "编程", "Go 语言入门经典", 59.00, 100},
		{"深入理解计算机系统", "Randal E. Bryant", "9787111544937", "机械工业出版社", "计算机", "CSAPP 经典教材", 139.00, 50},
		{"算法导论", "Thomas H. Cormen", "9787111407010", "机械工业出版社", "算法", "算法圣经", 128.00, 30},
		{"设计模式", "Erich Gamma", "9787111211358", "机械工业出版社", "设计", "GoF 设计模式", 49.00, 80},
		{"重构", "Martin Fowler", "9787115508645", "人民邮电出版社", "编程", "代码重构指南", 89.00, 60},
	}

	for _, book := range books {
		db.Exec(`INSERT INTO books (title, author, isbn, publisher, category, description, price, stock, created_at, updated_at)
		         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			book.title, book.author, book.isbn, book.publisher, book.category, book.description,
			book.price, book.stock, now, now)
	}

	log.Println("示例数据插入完成")
}

func main() {
	flag.Parse()

	// 初始化数据库
	db, err := initDB()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	s := grpc.NewServer()
	bookServer := NewBookServer(db)
	pb.RegisterBookServiceServer(s, bookServer)

	// 优雅关闭
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		log.Println("收到关闭信号，正在关闭服务...")
		s.GracefulStop()
	}()

	log.Printf("Book Service 启动在端口 %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务失败: %v", err)
	}
}
