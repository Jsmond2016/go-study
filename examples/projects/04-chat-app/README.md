# 聊天应用示例

这是一个完整的实时聊天应用示例，包含 WebSocket 实时通信、消息管理、聊天室等功能。

## 功能特性

- ✅ 用户管理
- ✅ WebSocket 实时通信
- ✅ 私聊功能
- ✅ 群组聊天（聊天室）
- ✅ 消息历史记录
- ✅ 在线状态管理

## 运行示例

```bash
# 安装依赖
go mod tidy

# 运行程序
go run main.go
```

## 测试 API

### 1. 创建用户

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"alice"}'
```

### 2. 创建聊天室

```bash
curl -X POST http://localhost:8080/api/rooms \
  -H "Content-Type: application/json" \
  -d '{"name":"Go 学习群"}'
```

### 3. 发送消息（HTTP）

```bash
curl -X POST http://localhost:8080/api/messages \
  -H "Content-Type: application/json" \
  -d '{
    "to_id":2,
    "content":"Hello, World!"
  }'
```

### 4. 获取消息历史

```bash
# 私聊消息
curl "http://localhost:8080/api/messages?to_id=2"

# 群组消息
curl "http://localhost:8080/api/messages?room_id=1"
```

### 5. WebSocket 连接测试

使用 WebSocket 客户端工具（如 Postman 或 wscat）连接：

```
ws://localhost:8080/ws
```

发送消息格式：
```json
{
  "to_id": 2,
  "room_id": 0,
  "content": "Hello via WebSocket",
  "type": "text"
}
```

## 使用 wscat 测试 WebSocket

```bash
# 安装 wscat
npm install -g wscat

# 连接
wscat -c ws://localhost:8080/ws

# 发送消息
{"to_id":2,"content":"Hello"}
```

## 项目结构

```
04-chat-app/
├── main.go          # 主程序
├── go.mod           # 依赖管理
└── README.md        # 说明文档
```

## 扩展功能

可以在此基础上添加：
- JWT 认证
- 文件传输
- 消息推送
- 消息已读状态
- 用户头像和昵称

