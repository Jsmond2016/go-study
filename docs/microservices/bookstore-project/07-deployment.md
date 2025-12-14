---
title: Chapter 7 - Docker 部署
difficulty: advanced
duration: "2-3小时"
tags: ["Docker", "Docker Compose", "部署", "容器化"]
---

# Chapter 7: Docker 部署

本章使用 Docker 容器化部署整个微服务系统。

## 📋 学习目标

- [x] 编写 Dockerfile
- [x] 使用 Docker Compose 编排服务
- [x] 一键部署整个系统
- [x] 生产环境优化建议

## 🐳 Dockerfile 编写

### Book Service Dockerfile

创建 `book-service/Dockerfile`：

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o book-service main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/book-service .

EXPOSE 50051
CMD ["./book-service"]
```

### User Service Dockerfile

创建 `user-service/Dockerfile`（类似结构）

### API Gateway Dockerfile

创建 `api-gateway/Dockerfile`（类似结构）

## 🎼 Docker Compose 配置

创建 `docker-compose.yml`：

```yaml
version: '3.8'

services:
  book-service:
    build: ./book-service
    container_name: book-service
    ports:
      - "50051:50051"
    volumes:
      - ./data:/root/data
    healthcheck:
      test: ["CMD", "grpc-health-probe", "-addr=:50051"]
      interval: 10s
      timeout: 5s
      retries: 3

  user-service:
    build: ./user-service
    container_name: user-service
    ports:
      - "50052:50052"
    volumes:
      - ./data:/root/data
    healthcheck:
      test: ["CMD", "grpc-health-probe", "-addr=:50052"]
      interval: 10s
      timeout: 5s
      retries: 3

  api-gateway:
    build: ./api-gateway
    container_name: api-gateway
    ports:
      - "8080:8080"
    depends_on:
      - book-service
      - user-service
    environment:
      - BOOK_SERVICE_ADDR=book-service:50051
      - USER_SERVICE_ADDR=user-service:50052
    healthcheck:
      test: ["CMD", "wget", "--spider", "http://localhost:8080/health"]
      interval: 10s
      timeout: 5s
      retries: 3

networks:
  default:
    name: bookstore-network
```

## 🚀 部署命令

### 构建镜像

```bash
docker-compose build
```

### 启动所有服务

```bash
docker-compose up -d
```

### 查看日志

```bash
# 所有服务
docker-compose logs -f

# 单个服务
docker-compose logs -f api-gateway
```

### 停止服务

```bash
docker-compose down
```

## 📊 监控和调试

### 查看容器状态

```bash
docker-compose ps
```

### 进入容器

```bash
docker exec -it api-gateway sh
```

### 查看资源使用

```bash
docker stats
```

## 🔧 生产环境优化

### 1. 使用环境变量

```yaml
services:
  api-gateway:
    environment:
      - JWT_SECRET=${JWT_SECRET}
      - LOG_LEVEL=${LOG_LEVEL:-info}
```

### 2. 添加重启策略

```yaml
services:
  api-gateway:
    restart: unless-stopped
```

### 3. 限制资源

```yaml
services:
  book-service:
    deploy:
      resources:
        limits:
          cpus: '0.50'
          memory: 512M
```

### 4. 使用外部数据库

生产环境建议使用 PostgreSQL/MySQL：

```yaml
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: bookstore
      POSTGRES_USER: bookstore
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

## ✅ 部署检查清单

- [ ] 所有 Dockerfile 已创建
- [ ] docker-compose.yml 配置正确
- [ ] 所有服务成功构建
- [ ] 所有服务成功启动
- [ ] 健康检查通过
- [ ] API 测试通过

## 🎉 恭喜完成！

你已经完成了整个微服务项目的学习！

### 下一步建议

1. **Kubernetes 部署** - 学习 K8s 编排
2. **添加监控** - Prometheus + Grafana
3. **服务网格** - Istio 集成
4. **CI/CD** - GitHub Actions 自动化

### 扩展阅读

- [Kubernetes 编排](/devops/02-kubernetes)
- [监控告警](/devops/04-monitoring)
- [服务网格](/microservices/09-service-mesh)

---

**🎉 项目完成！你已经掌握了完整的微服务开发技能！**
