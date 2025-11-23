---
title: Docker 基础
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["基础语法", "Linux 基础"]
tags: ["Docker", "容器化", "部署", "镜像"]
---

# Docker 基础

Docker 是一个容器化平台，可以将应用及其依赖打包成容器，实现快速部署和扩展。

## 📋 学习目标

- [ ] 理解 Docker 的概念和优势
- [ ] 掌握 Dockerfile 的编写
- [ ] 学会构建和运行容器
- [ ] 理解 Docker Compose 的使用
- [ ] 掌握多阶段构建
- [ ] 了解最佳实践

## 🎯 Docker 简介

### 为什么使用 Docker

- **环境一致性**: 开发、测试、生产环境一致
- **快速部署**: 容器启动速度快
- **资源隔离**: 容器之间相互隔离
- **易于扩展**: 可以快速扩展容器数量
- **版本管理**: 镜像版本化管理

### 安装 Docker

```bash
# macOS
brew install docker

# Linux
curl -fsSL https://get.docker.com | sh

# 验证安装
docker --version
```

## 🚀 快速开始

### 基本命令

```bash
# 运行容器
docker run hello-world

# 运行交互式容器
docker run -it ubuntu /bin/bash

# 后台运行容器
docker run -d nginx

# 查看运行中的容器
docker ps

# 查看所有容器
docker ps -a

# 停止容器
docker stop <container_id>

# 删除容器
docker rm <container_id>
```

## 📝 Dockerfile

### 基本 Dockerfile

```dockerfile
# 使用官方 Go 镜像作为基础镜像
FROM golang:1.21-alpine

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]
```

### 多阶段构建

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
```

### 优化 Dockerfile

```dockerfile
# 使用多阶段构建减小镜像大小
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 先复制依赖文件，利用 Docker 缓存
COPY go.mod go.sum ./
RUN go mod download

# 再复制源代码
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 运行阶段使用最小镜像
FROM scratch

# 复制必要的文件
COPY --from=builder /app/main /main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/main"]
```

## 🏗️ 构建和运行

### 构建镜像

```bash
# 构建镜像
docker build -t myapp:latest .

# 指定 Dockerfile
docker build -f Dockerfile.prod -t myapp:prod .

# 构建时传递参数
docker build --build-arg VERSION=1.0.0 -t myapp:1.0.0 .
```

### 运行容器

```bash
# 运行容器
docker run -p 8080:8080 myapp:latest

# 后台运行
docker run -d -p 8080:8080 --name myapp myapp:latest

# 挂载卷
docker run -d -p 8080:8080 -v $(pwd)/config:/app/config myapp:latest

# 设置环境变量
docker run -d -p 8080:8080 -e DB_HOST=localhost myapp:latest

# 使用 .env 文件
docker run -d -p 8080:8080 --env-file .env myapp:latest
```

## 🐳 Docker Compose

### docker-compose.yml

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=db
      - DB_PORT=3306
      - REDIS_HOST=redis
    depends_on:
      - db
      - redis
    volumes:
      - ./config:/app/config

  db:
    image: mysql:8.0
    environment:
      - MYSQL_ROOT_PASSWORD=password
      - MYSQL_DATABASE=myapp
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

volumes:
  db_data:
  redis_data:
```

### Compose 命令

```bash
# 启动服务
docker-compose up

# 后台启动
docker-compose up -d

# 停止服务
docker-compose down

# 查看日志
docker-compose logs -f

# 重新构建
docker-compose build

# 执行命令
docker-compose exec app sh
```

## 🏃‍♂️ 实践应用

### Go 应用 Dockerfile

```dockerfile
# 多阶段构建示例
FROM golang:1.21-alpine AS builder

# 安装构建工具
RUN apk add --no-cache git

WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 运行阶段
FROM alpine:latest

# 安装 CA 证书（用于 HTTPS）
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 创建非 root 用户
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# 更改文件所有者
RUN chown -R appuser:appuser /root

# 切换到非 root 用户
USER appuser

EXPOSE 8080

CMD ["./main"]
```

### .dockerignore

```
# .dockerignore
.git
.gitignore
README.md
.env
.idea
.vscode
*.md
dist/
tmp/
```

## ⚠️ 最佳实践

### 1. 镜像大小优化

```dockerfile
# ✅ 使用多阶段构建
# ✅ 使用 alpine 基础镜像
# ✅ 清理不必要的文件
RUN rm -rf /tmp/* /var/cache/apk/*
```

### 2. 安全性

```dockerfile
# ✅ 使用非 root 用户运行
USER appuser

# ✅ 只复制必要的文件
COPY --from=builder /app/main .

# ✅ 使用最小基础镜像
FROM alpine:latest
```

### 3. 缓存优化

```dockerfile
# ✅ 先复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# ✅ 再复制源代码
COPY . .
```

## 📚 扩展阅读

- [Docker 官方文档](https://docs.docker.com/)
- [Docker Compose 文档](https://docs.docker.com/compose/)
- [Go Docker 最佳实践](https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/)

## ⏭️ 下一章节

[测试框架](./13-testify.md) → 学习单元测试

---

**💡 提示**: Docker 是现代应用部署的标准方式，掌握它对于开发和运维都非常重要！

