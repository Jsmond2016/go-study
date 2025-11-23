---
title: 部署优化
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["搜索功能"]
tags: ["部署", "Docker", "性能优化", "监控"]
---

# 部署优化

本章节将介绍博客系统的部署方案、性能优化和监控配置。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 使用Docker容器化应用
- [ ] 配置Nginx反向代理
- [ ] 实现数据库连接池优化
- [ ] 配置日志和监控
- [ ] 实现缓存策略
- [ ] 进行性能调优

## 🐳 Docker 部署

### Dockerfile

创建 `Dockerfile`:

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制go mod文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o blog-system cmd/server/main.go

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/blog-system .

# 复制配置文件
COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./blog-system"]
```

### Docker Compose

创建 `docker-compose.yml`:

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    depends_on:
      - mysql
      - redis
    volumes:
      - ./uploads:/app/uploads
      - ./logs:/app/logs

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: blog_db
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - app

volumes:
  mysql_data:
  redis_data:
```

## 🔧 Nginx 配置

创建 `nginx.conf`:

```nginx
upstream blog_backend {
    server app:8080;
}

server {
    listen 80;
    server_name your-domain.com;

    # 静态文件
    location /uploads/ {
        alias /app/uploads/;
        expires 30d;
    }

    # API代理
    location /api/ {
        proxy_pass http://blog_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # 压缩
    gzip on;
    gzip_types text/plain text/css application/json application/javascript;
}
```

## ⚡ 性能优化

### 数据库连接池

```go
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
```

### Redis 缓存

```go
import "github.com/go-redis/redis/v8"

func InitRedis(addr, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
```

### 缓存文章列表

```go
func (s *ArticleServiceImpl) List(filter ArticleFilter) ([]Article, int64, error) {
	// 生成缓存key
	cacheKey := fmt.Sprintf("articles:list:%v", filter)
	
	// 尝试从缓存获取
	var cached []Article
	if err := redis.Get(ctx, cacheKey).Scan(&cached); err == nil {
		return cached, 0, nil
	}

	// 从数据库获取
	articles, total, err := s.articleRepo.List(filter)
	if err != nil {
		return nil, 0, err
	}

	// 存入缓存
	redis.Set(ctx, cacheKey, articles, time.Hour)

	return articles, total, nil
}
```

## 📊 日志配置

### 结构化日志

```go
import "go.uber.org/zap"

func InitLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "./logs/app.log"}
	config.ErrorOutputPaths = []string{"stderr", "./logs/error.log"}

	return config.Build()
}
```

### 日志中间件

```go
func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.Info("HTTP Request",
			zap.String("method", param.Method),
			zap.String("path", param.Path),
			zap.Int("status", param.StatusCode),
			zap.Duration("latency", param.Latency),
			zap.String("ip", param.ClientIP),
		)
		return ""
	})
}
```

## 📈 监控配置

### Prometheus 指标

```go
import "github.com/prometheus/client_golang/prometheus"

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}
```

### 健康检查

```go
func HealthCheck(c *gin.Context) {
	// 检查数据库
	if err := db.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "database connection failed",
		})
		return
	}

	// 检查Redis
	if err := redis.Ping(ctx).Err(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "redis connection failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
```

## 🚀 部署步骤

### 1. 构建镜像

```bash
docker build -t blog-system:latest .
```

### 2. 启动服务

```bash
docker-compose up -d
```

### 3. 查看日志

```bash
docker-compose logs -f app
```

### 4. 数据库迁移

```bash
docker-compose exec app ./blog-system migrate
```

## 💡 最佳实践

### 1. 安全配置

- 使用HTTPS
- 配置防火墙
- 限制数据库访问
- 使用环境变量存储敏感信息

### 2. 性能优化

- 启用Gzip压缩
- 使用CDN加速静态资源
- 配置缓存策略
- 优化数据库查询

### 3. 监控告警

- 配置应用监控
- 设置告警规则
- 定期检查日志
- 监控系统资源

## 📚 相关资源

- [Docker 官方文档](https://docs.docker.com/)
- [Nginx 配置指南](https://nginx.org/en/docs/)
- [Prometheus 文档](https://prometheus.io/docs/)

---

**🎉 部署优化完成！** 恭喜你完成了整个博客系统的开发！

