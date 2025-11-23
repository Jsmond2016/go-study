// Package main 展示 Gin 框架的基本用法
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// User 用户结构体
type User struct {
	ID       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"gte=0,lte=150"`
	IsActive bool   `json:"is_active"`
	CreateAt string `json:"created_at"`
}

// APIResponse 通用API响应
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 分页响应
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
}

// 模拟数据库
var (
	users []User
	nextID int = 1
)

func main() {
	fmt.Println("=== Gin 框架基础示例 ===")

	// 初始化用户数据
	initUsers()

	// 创建 Gin 引擎
	r := setupRouter()

	// 启动服务器
	fmt.Println("Gin 服务器启动在 http://localhost:8080")
	fmt.Println("API 文档: http://localhost:8080/swagger")
	fmt.Println("可用端点:")
	fmt.Println("  GET    /api/v1/users       - 获取用户列表（支持分页、搜索、排序）")
	fmt.Println("  GET    /api/v1/users/:id    - 获取指定用户")
	fmt.Println("  POST   /api/v1/users       - 创建新用户")
	fmt.Println("  PUT    /api/v1/users/:id    - 更新用户信息")
	fmt.Println("  DELETE /api/v1/users/:id    - 删除用户")
	fmt.Println("  GET    /api/v1/health      - 健康检查")
	fmt.Println("  GET    /api/v1/middleware  - 中间件演示")
	fmt.Println("  GET    /api/v1/config      - 配置信息")
	fmt.Println("  POST   /api/v1/upload      - 文件上传")
	fmt.Println("  GET    /api/v1/download    - 文件下载")

	// 启动服务器
	r.Run(":8080")
}

// initUsers 初始化示例用户数据
func initUsers() {
	users = []User{
		{ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 25, IsActive: true, CreateAt: "2024-01-01T00:00:00Z"},
		{ID: 2, Name: "李四", Email: "lisi@example.com", Age: 30, IsActive: true, CreateAt: "2024-01-02T00:00:00Z"},
		{ID: 3, Name: "王五", Email: "wangwu@example.com", Age: 28, IsActive: false, CreateAt: "2024-01-03T00:00:00Z"},
		{ID: 4, Name: "赵六", Email: "zhaoliu@example.com", Age: 22, IsActive: true, CreateAt: "2024-01-04T00:00:00Z"},
		{ID: 5, Name: "钱七", Email: "qianqi@example.com", Age: 35, IsActive: true, CreateAt: "2024-01-05T00:00:00Z"},
	}
	nextID = 6
}

// setupRouter 设置路由
func setupRouter() *gin.Engine {
	// 设置 Gin 模式
	// gin.SetMode(gin.ReleaseMode) // 生产环境
	// gin.SetMode(gin.TestMode)     // 测试环境

	r := gin.Default() // 包含 Logger 和 Recovery 中间件

	// 添加自定义中间件
	r.Use(corsMiddleware())
	r.use(loggingMiddleware())

	// 创建 API 路由组
	v1 := r.Group("/api/v1")
	{
		// 基础路由
		v1.GET("/", handleHome)
		v1.GET("/health", handleHealth)
		v1.GET("/config", handleConfig)

		// 用户管理路由
		users := v1.Group("/users")
		{
			users.GET("", handleGetUsers)           // 获取用户列表
			users.GET("/:id", handleGetUserByID)     // 获取指定用户
			users.POST("", handleCreateUser)         // 创建用户
			users.PUT("/:id", handleUpdateUser)      // 更新用户
			users.DELETE("/:id", handleDeleteUser)   // 删除用户

			// 中间件演示
			users.GET("/middleware", demoMiddleware)
		}

		// 文件操作路由
		v1.POST("/upload", handleFileUpload)
		v1.GET("/download", handleFileDownload)
	}

	// 添加 Swagger 文档路由（实际项目中使用 swaggo/gin-swagger）
	r.GET("/swagger", handleSwagger)

	return r
}

// handleHome 首页
func handleHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin 框架示例",
		"version": "1.0.0",
	})
}

// handleHealth 健康检查
func handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    200,
		Message: "服务器运行正常",
		Data: gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"version":   "1.0.0",
		},
	})
}

// handleConfig 配置信息
func handleConfig(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    200,
		Message: "配置信息获取成功",
		Data: gin.H{
			"server": gin.H{
				"port":     8080,
				"mode":     gin.Mode(),
				"timezone": "Asia/Shanghai",
			},
			"database": gin.H{
				"type":     "sqlite",
				"host":     "localhost",
				"port":     5432,
				"database": "myapp",
			},
			"features": []string{"auth", "logging", "metrics"},
		},
	})
}

// handleGetUsers 获取用户列表
func handleGetUsers(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")

	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 过滤用户
	var filteredUsers []User
	for _, user := range users {
		if search == "" ||
		   user.Name == search ||
		   user.Email == search ||
		   fmt.Sprintf("%d", user.ID) == search {
			filteredUsers = append(filteredUsers, user)
		}
	}

	// 排序
	sortedUsers := sortUsers(filteredUsers, sortBy, sortOrder)

	// 分页
	total := len(sortedUsers)
	totalPages := (total + pageSize - 1) / pageSize
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	pagedUsers := sortedUsers[start:end]

	// 返回分页响应
	c.JSON(http.StatusOK, APIResponse{
		Code:    200,
		Message: "用户列表获取成功",
		Data: PaginatedResponse{
			Items:      pagedUsers,
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// handleGetUserByID 获取指定用户
func handleGetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Code:    200,
				Message: "用户获取成功",
				Data:    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Code:    404,
		Message: "用户不存在",
	})
}

// handleCreateUser 创建用户
func handleCreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 设置用户ID和创建时间
	user.ID = nextID
	user.IsActive = true
	user.CreateAt = time.Now().Format(time.RFC3339)

	// 添加到用户列表
	users = append(users, user)
	nextID++

	c.JSON(http.StatusCreated, APIResponse{
		Code:    201,
		Message: "用户创建成功",
		Data:    user,
	})
}

// handleUpdateUser 更新用户
func handleUpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	var updateData User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 查找并更新用户
	for i := range users {
		if users[i].ID == id {
			if updateData.Name != "" {
				users[i].Name = updateData.Name
			}
			if updateData.Email != "" {
				users[i].Email = updateData.Email
			}
			if updateData.Age > 0 {
				users[i].Age = updateData.Age
			}

			c.JSON(http.StatusOK, APIResponse{
				Code:    200,
				Message: "用户更新成功",
				Data:    users[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Code:    404,
		Message: "用户不存在",
	})
}

// handleDeleteUser 删除用户
func handleDeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	// 查找并删除用户
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, APIResponse{
				Code:    200,
				Message: "用户删除成功",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Code:    404,
		Message: "用户不存在",
	})
}

// handleFileUpload 文件上传
func handleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    400,
			Message: "文件上传失败: " + err.Error(),
		})
		return
	}

	// 保存文件（实际应用中应该保存到指定目录）
	c.JSON(http.StatusOK, APIResponse{
		Code:    200,
		Message: "文件上传成功",
		Data: gin.H{
			"filename": file.Filename,
			"size":     file.Size,
			"header":   file.Header,
		},
	})
}

// handleFileDownload 文件下载
func handleFileDownload(c *gin.Context) {
	filename := c.DefaultQuery("filename", "example.txt")

	// 设置下载头
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Header("Content-Type", "application/octet-stream")

	// 写入文件内容
	c.String(http.StatusOK, "这是一个示例文件，下载时间: %s", time.Now().Format(time.RFC3339))
}

// handleSwagger Swagger 文档
func handleSwagger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"swagger": "2.0",
		"info": gin.H{
			"title":       "Gin 示例 API",
			"description": "演示 Gin 框架的各种功能",
			"version":     "1.0.0",
		},
		"host": "localhost:8080",
		"basePath": "/api/v1",
		"schemes": []string{"http", "https"},
		"paths": gin.H{
			"/users": gin.H{
				"get": gin.H{
					"summary": "获取用户列表",
					"responses": gin.H{
						"200": gin.H{
							"description": "成功",
						},
					},
				},
				"post": gin.H{
					"summary": "创建用户",
					"responses": gin.H{
						"201": gin.H{
							"description": "创建成功",
						},
					},
				},
			},
		},
	})
}

// demoMiddleware 中间件演示
func demoMiddleware(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    200,
		Message: "中间件演示成功",
		Data: gin.H{
			"request_id": c.GetString("request_id"),
			"user_agent": c.GetHeader("User-Agent"),
			"client_ip":  c.ClientIP(),
		},
	})
}

// 中间件函数

// corsMiddleware CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// loggingMiddleware 日志中间件
func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

// authMiddleware 认证中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, APIResponse{
				Code:    401,
				Message: "缺少认证令牌",
			})
			c.Abort()
			return
		}

		// 这里应该验证 token 的有效性
		if token != "Bearer valid-token" {
			c.JSON(http.StatusUnauthorized, APIResponse{
				Code:    401,
				Message: "无效的认证令牌",
			})
			c.Abort()
			return
		}

		// 设置用户信息到上下文
		c.Set("user_id", 1)
		c.Set("username", "admin")

		c.Next()
	}
}

// requestIdMiddleware 请求ID中间件
func requestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// 辅助函数

// sortUsers 用户排序
func sortUsers(users []User, sortBy, sortOrder string) []User {
	// 这里简化处理，实际应用中可以使用更复杂的排序逻辑
	if sortBy == "name" && sortOrder == "asc" {
		// 按姓名升序
		for i := 0; i < len(users); i++ {
			for j := i + 1; j < len(users); j++ {
				if users[i].Name > users[j].Name {
					users[i], users[j] = users[j], users[i]
				}
			}
		}
	}
	return users
}

// use 添加中间件的辅助方法（兼容性处理）
func (r *gin.Engine) use(middlewares ...gin.HandlerFunc) {
	r.Use(middlewares...)
}