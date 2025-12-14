package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "bookstore/proto/user"
)

const (
	port      = ":50052"
	jwtSecret = "your-secret-key-change-in-production" // 生产环境请使用环境变量
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

// JWT Claims
type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func main() {
	// 初始化数据库
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{db: db})

	log.Printf("User Service is running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// 初始化数据库
func initDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}

	// 创建用户表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		phone TEXT,
		avatar TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}

// Register - 用户注册
func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	log.Printf("Register request: username=%s, email=%s", req.Username, req.Email)

	// 验证输入
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "username, email and password are required")
	}

	if len(req.Password) < 6 {
		return nil, status.Error(codes.InvalidArgument, "password must be at least 6 characters")
	}

	// 检查用户名是否已存在
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", req.Username).Scan(&exists)
	if err != nil {
		log.Printf("Error checking username: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	}

	// 检查邮箱是否已存在
	err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&exists)
	if err != nil {
		log.Printf("Error checking email: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "email already exists")
	}

	// 哈希密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, status.Error(codes.Internal, "failed to process password")
	}

	// 插入用户
	result, err := s.db.Exec(
		"INSERT INTO users (username, email, password_hash, phone) VALUES (?, ?, ?, ?)",
		req.Username, req.Email, string(passwordHash), req.Phone,
	)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting user ID: %v", err)
		return nil, status.Error(codes.Internal, "failed to get user ID")
	}

	// 生成 JWT Token
	token, err := generateToken(userID, req.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// 获取用户信息
	user, err := s.getUserByID(userID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	log.Printf("User registered successfully: id=%d, username=%s", userID, req.Username)

	return &pb.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// Login - 用户登录
func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	log.Printf("Login request: email=%s", req.Email)

	// 验证输入
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	// 查询用户
	var userID int64
	var username, passwordHash string
	err := s.db.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE email = ?",
		req.Email,
	).Scan(&userID, &username, &passwordHash)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if err != nil {
		log.Printf("Error querying user: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// 生成 JWT Token
	token, err := generateToken(userID, req.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	// 获取用户信息
	user, err := s.getUserByID(userID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	log.Printf("User logged in successfully: id=%d, username=%s", userID, username)

	return &pb.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// GetUser - 获取用户信息
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	log.Printf("GetUser request: id=%d", req.Id)

	user, err := s.getUserByID(req.Id)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}

	return &pb.UserResponse{User: user}, nil
}

// UpdateUser - 更新用户信息
func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	log.Printf("UpdateUser request: id=%d", req.Id)

	// 检查用户是否存在
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", req.Id).Scan(&exists)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return nil, status.Error(codes.Internal, "database error")
	}
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// 构建更新 SQL
	query := "UPDATE users SET updated_at = CURRENT_TIMESTAMP"
	args := []interface{}{}

	if req.Username != "" {
		// 检查用户名是否被其他用户占用
		var otherUserExists bool
		err := s.db.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM users WHERE username = ? AND id != ?)",
			req.Username, req.Id,
		).Scan(&otherUserExists)
		if err != nil {
			log.Printf("Error checking username: %v", err)
			return nil, status.Error(codes.Internal, "database error")
		}
		if otherUserExists {
			return nil, status.Error(codes.AlreadyExists, "username already exists")
		}
		query += ", username = ?"
		args = append(args, req.Username)
	}

	if req.Phone != "" {
		query += ", phone = ?"
		args = append(args, req.Phone)
	}

	if req.Avatar != "" {
		query += ", avatar = ?"
		args = append(args, req.Avatar)
	}

	query += " WHERE id = ?"
	args = append(args, req.Id)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	// 获取更新后的用户信息
	user, err := s.getUserByID(req.Id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	log.Printf("User updated successfully: id=%d", req.Id)

	return &pb.UserResponse{User: user}, nil
}

// ValidateToken - 验证 JWT Token
func (s *UserServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	log.Printf("ValidateToken request")

	claims, err := parseToken(req.Token)
	if err != nil {
		return &pb.ValidateTokenResponse{
			Valid: false,
			Error: err.Error(),
		}, nil
	}

	// 检查用户是否存在
	var exists bool
	err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", claims.UserID).Scan(&exists)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return &pb.ValidateTokenResponse{
			Valid: false,
			Error: "database error",
		}, nil
	}
	if !exists {
		return &pb.ValidateTokenResponse{
			Valid: false,
			Error: "user not found",
		}, nil
	}

	log.Printf("Token validated successfully: user_id=%d", claims.UserID)

	return &pb.ValidateTokenResponse{
		Valid:  true,
		UserId: claims.UserID,
		Email:  claims.Email,
	}, nil
}

// 辅助函数：根据 ID 获取用户
func (s *UserServer) getUserByID(id int64) (*pb.User, error) {
	var user pb.User
	var createdAt, updatedAt string
	var phone, avatar sql.NullString

	err := s.db.QueryRow(`
		SELECT id, username, email, phone, avatar, created_at, updated_at
		FROM users WHERE id = ?
	`, id).Scan(
		&user.Id, &user.Username, &user.Email,
		&phone, &avatar, &createdAt, &updatedAt,
	)

	if err != nil {
		return nil, err
	}

	if phone.Valid {
		user.Phone = phone.String
	}
	if avatar.Valid {
		user.Avatar = avatar.String
	}
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt

	return &user, nil
}

// 生成 JWT Token
func generateToken(userID int64, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// 解析 JWT Token
func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
