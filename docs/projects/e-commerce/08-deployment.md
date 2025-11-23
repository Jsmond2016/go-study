---
title: 部署优化
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["库存管理"]
tags: ["部署", "Docker", "性能优化", "监控"]
---

# 部署优化

本章节将介绍电商系统的部署方案、性能优化和监控配置。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 使用Docker容器化应用
- [ ] 配置Nginx反向代理
- [ ] 实现数据库优化
- [ ] 配置缓存策略
- [ ] 实现监控和日志
- [ ] 进行性能调优

## 🐳 Docker 部署

### Dockerfile

创建 `Dockerfile`:

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o e-commerce cmd/server/main.go

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/e-commerce .
COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./e-commerce"]
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
      MYSQL_DATABASE: ecommerce_db
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
    command: --innodb-buffer-pool-size=1G

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: redis-server --appendonly yes

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - app

volumes:
  mysql_data:
  redis_data:
```

## ⚡ 性能优化

### 数据库优化

```go
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 预编译语句
		NowFunc:     time.Now,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 连接池优化
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}
```

### Redis 缓存

```go
// 商品缓存
func (s *ProductServiceImpl) GetByID(id uint) (*model.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)
	
	var product model.Product
	if err := redis.Get(ctx, cacheKey).Scan(&product); err == nil {
		return &product, nil
	}

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	redis.Set(ctx, cacheKey, product, time.Hour)
	return &product, nil
}

// 购物车缓存
func (s *CartServiceImpl) GetCart(userID uint) ([]model.CartItem, float64, error) {
	cacheKey := fmt.Sprintf("cart:%d", userID)
	
	var items []model.CartItem
	if err := redis.Get(ctx, cacheKey).Scan(&items); err == nil {
		total := calculateTotal(items)
		return items, total, nil
	}

	// 从数据库获取并缓存
	items, total, err := s.getCartFromDB(userID)
	if err == nil {
		redis.Set(ctx, cacheKey, items, 30*time.Minute)
	}
	return items, total, err
}
```

### 索引优化

```sql
-- 商品表索引
CREATE INDEX idx_product_category ON products(category_id);
CREATE INDEX idx_product_status ON products(status);
CREATE INDEX idx_product_price ON products(price);

-- 订单表索引
CREATE INDEX idx_order_user ON orders(user_id);
CREATE INDEX idx_order_status ON orders(status);
CREATE INDEX idx_order_created ON orders(created_at);

-- 订单项索引
CREATE INDEX idx_order_item_order ON order_items(order_id);
CREATE INDEX idx_order_item_product ON order_items(product_id);
```

## 📊 监控配置

### Prometheus 指标

```go
import "github.com/prometheus/client_golang/prometheus"

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request duration",
		},
		[]string{"method", "endpoint"},
	)

	orderTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "orders_total",
			Help: "Total orders",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(orderTotal)
}
```

### 日志配置

```go
import "go.uber.org/zap"

func InitLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "./logs/app.log"}
	config.ErrorOutputPaths = []string{"stderr", "./logs/error.log"}
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	return config.Build()
}
```

## 🔧 Nginx 配置

创建 `nginx.conf`:

```nginx
upstream ecommerce_backend {
    server app:8080;
    keepalive 32;
}

server {
    listen 80;
    server_name your-domain.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name your-domain.com;

    ssl_certificate /etc/nginx/ssl/cert.pem;
    ssl_certificate_key /etc/nginx/ssl/key.pem;

    # 静态文件
    location /uploads/ {
        alias /app/uploads/;
        expires 30d;
        add_header Cache-Control "public, immutable";
    }

    # API代理
    location /api/ {
        proxy_pass http://ecommerce_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # 压缩
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css application/json application/javascript;
}
```

## 🚀 部署步骤

### 1. 构建镜像

```bash
docker build -t e-commerce:latest .
```

### 2. 启动服务

```bash
docker-compose up -d
```

### 3. 数据库迁移

```bash
docker-compose exec app ./e-commerce migrate
```

### 4. 查看日志

```bash
docker-compose logs -f app
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
- 使用连接池

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

**🎉 部署优化完成！** 恭喜你完成了整个电商系统的开发！

