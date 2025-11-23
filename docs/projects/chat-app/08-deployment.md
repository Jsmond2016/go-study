---
title: 部署优化
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["消息推送"]
tags: ["部署", "Docker", "性能优化", "监控"]
---

# 部署优化

本章节将介绍聊天应用的部署方案、性能优化和监控配置。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 使用Docker容器化应用
- [ ] 配置WebSocket负载均衡
- [ ] 实现连接池管理
- [ ] 配置监控和日志
- [ ] 进行性能调优

## 🐳 Docker 部署

### Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chat-app cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/chat-app .
EXPOSE 8080
CMD ["./chat-app"]
```

## ⚡ 性能优化

### WebSocket连接池

```go
// 限制每个用户的连接数
func (h *Hub) RegisterClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// 检查用户是否已有连接
	existingConnections := 0
	for c := range h.clients {
		if c.UserID == client.UserID {
			existingConnections++
		}
	}

	if existingConnections >= 3 {
		client.Conn.Close()
		return
	}

	h.clients[client] = true
}
```

## 📊 监控配置

### 连接数监控

```go
var (
	activeConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "websocket_connections_active",
			Help: "当前活跃的WebSocket连接数",
		},
	)
)
```

## 💡 最佳实践

### 1. 连接管理

- 限制每个用户的连接数
- 实现连接超时机制
- 使用连接池管理

### 2. 消息处理

- 使用消息队列处理大量消息
- 实现消息批量处理
- 优化数据库查询

### 3. 扩展性

- 使用Redis实现分布式
- 实现水平扩展
- 使用消息队列解耦

---

**🎉 部署优化完成！** 恭喜你完成了整个聊天应用的开发！

