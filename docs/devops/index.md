---
title: 运维部署
difficulty: intermediate
duration: "6-8周"
prerequisites: ["Web 开发", "数据库"]
tags: ["Docker", "Kubernetes", "CI/CD", "监控", "日志"]
---

# 运维部署

本模块介绍 Go 应用的运维和部署，包括容器化、编排、CI/CD 等。

## 📋 学习目标

完成本模块学习后，你将能够：

- [ ] 使用 Docker 容器化应用
- [ ] 使用 Kubernetes 编排服务
- [ ] 配置 CI/CD 流水线
- [ ] 实现监控和告警
- [ ] 管理日志和追踪
- [ ] 掌握运维最佳实践

## 🎯 学习路径

### 📚 第一部分：容器化（第1-2周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [Docker 容器化](./01-docker.md)         | Docker 基础和使用  | 4-5小时  | ⭐⭐   |

### 🚀 第二部分：编排和部署（第3-4周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [Kubernetes 编排](./02-kubernetes.md)   | K8s 基础和使用     | 5-6小时  | ⭐⭐⭐ |
| [CI/CD 流水线](./03-ci-cd.md)          | 持续集成和部署     | 4-5小时  | ⭐⭐⭐ |

### 📊 第三部分：监控和运维（第5-6周）

| 章节                                    | 内容               | 预计时间 | 难度 |
| --------------------------------------- | ------------------ | -------- | ---- |
| [监控告警](./04-monitoring.md)          | 监控系统配置       | 4-5小时  | ⭐⭐⭐ |
| [日志管理](./05-logging.md)             | 日志收集和分析     | 3-4小时  | ⭐⭐   |

## 🚀 快速开始

### Dockerfile 示例

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/app .
CMD ["./app"]
```

## 💡 学习建议

### 📖 学习方法

1. **实践为主**：每个工具都要动手实践
2. **理解原理**：理解容器化和编排原理
3. **工具使用**：掌握常用运维工具
4. **最佳实践**：学习运维最佳实践

### 🔍 推荐资源

- [Docker 官方文档](https://docs.docker.com/)
- [Kubernetes 官方文档](https://kubernetes.io/docs/)
- [CI/CD 最佳实践](https://www.thoughtworks.com/insights/blog/cicd)

## ⏭️ 下一阶段

完成运维部署学习后，可以进入：

- [实战项目](../projects/) - 完整项目开发

---

**🎉 开始你的运维部署学习之旅吧！** 选择第一个章节，开始学习容器化和部署。

