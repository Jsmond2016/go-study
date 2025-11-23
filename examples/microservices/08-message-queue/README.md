# 消息队列示例

本目录包含使用 RabbitMQ 和 Kafka 实现消息队列的示例代码。

## 📋 目录结构

```
08-message-queue/
├── rabbitmq/          # RabbitMQ 示例
│   ├── producer/     # 生产者
│   └── consumer/     # 消费者
├── kafka/            # Kafka 示例
│   ├── producer/     # 生产者
│   └── consumer/     # 消费者
├── go.mod            # Go 模块定义
└── README.md         # 本文件
```

## 🚀 快速开始

### 1. 启动 RabbitMQ

```bash
docker run -d --name rabbitmq \
  -p 5672:5672 \
  -p 15672:15672 \
  -e RABBITMQ_DEFAULT_USER=guest \
  -e RABBITMQ_DEFAULT_PASS=guest \
  rabbitmq:3-management
```

访问管理界面：http://localhost:15672

### 2. 启动 Kafka

参考 Kafka 官方文档启动 Kafka 服务。

### 3. 运行 RabbitMQ 示例

**启动消费者**：

```bash
cd rabbitmq/consumer
go run main.go
```

**运行生产者**：

```bash
cd rabbitmq/producer
go run main.go
```

### 4. 运行 Kafka 示例

**启动消费者**：

```bash
cd kafka/consumer
go run main.go
```

**运行生产者**：

```bash
cd kafka/producer
go run main.go
```

## 📝 测试

### 1. 测试 RabbitMQ 消息发送和接收

```bash
# 终端 1：启动消费者
cd rabbitmq/consumer
go run main.go

# 终端 2：运行生产者
cd rabbitmq/producer
go run main.go

# 观察消费者终端，应该能看到接收到的消息
```

### 2. 测试 RabbitMQ 消息确认

```bash
# 1. 启动消费者（自动确认模式）
cd rabbitmq/consumer
go run main.go

# 2. 发送消息
cd rabbitmq/producer
go run main.go

# 3. 在 RabbitMQ 管理界面查看消息状态
# http://localhost:15672
```

### 3. 测试 Kafka 消息发送和接收

```bash
# 终端 1：启动消费者
cd kafka/consumer
go run main.go

# 终端 2：运行生产者
cd kafka/producer
go run main.go

# 观察消费者终端，应该能看到接收到的消息
```

### 4. 测试消息持久化

```bash
# 1. 发送持久化消息
cd rabbitmq/producer
go run main.go

# 2. 停止消费者
# 3. 重启消费者
# 4. 应该能收到之前未确认的消息
```

### 5. 测试消息重试

```bash
# 1. 修改消费者代码，模拟处理失败
# 2. 发送消息
# 3. 观察消息是否被重新入队
```

## 🐛 常见问题

### 1. 无法连接到 RabbitMQ

**检查项**：
- RabbitMQ 服务是否运行：`docker ps | grep rabbitmq`
- 端口是否正确：默认 5672
- 用户名密码是否正确：默认 guest/guest
- 防火墙是否阻止连接

**解决方法**：
```bash
# 检查 RabbitMQ 状态
docker ps | grep rabbitmq

# 查看 RabbitMQ 日志
docker logs rabbitmq

# 重启 RabbitMQ
docker restart rabbitmq
```

### 2. 无法连接到 Kafka

**检查项**：
- Kafka 服务是否运行
- Broker 地址是否正确
- 网络连接是否正常

**解决方法**：
```bash
# 检查 Kafka 是否运行
# 根据你的 Kafka 安装方式检查

# 测试连接
telnet <kafka-host> <kafka-port>
```

### 3. 消息丢失

**可能原因**：
- 消息未持久化
- 消费者未确认消息
- 队列配置错误

**解决方法**：
- 使用持久化消息：`amqp.Persistent`
- 确保消费者正确确认消息：`d.Ack(false)`
- 配置队列持久化：`durable: true`

### 4. 消息重复消费

**可能原因**：
- 消费者处理时间过长，消息超时
- 网络问题导致连接断开
- 消费者异常退出

**解决方法**：
- 实现幂等性处理
- 使用消息去重机制
- 合理设置超时时间

### 5. 消费者无法接收消息

**可能原因**：
- 队列名称不匹配
- 路由键不匹配
- 消费者未正确绑定队列

**解决方法**：
- 检查队列名称是否一致
- 检查路由键配置
- 检查 Exchange 和 Queue 绑定关系

## 💡 最佳实践

### 1. 消息确认

- **自动确认**：适合处理快速、可靠的消息
- **手动确认**：适合处理耗时、可能失败的消息

### 2. 消息持久化

- 重要消息必须持久化
- 队列和 Exchange 也要持久化

### 3. 错误处理

- 实现重试机制
- 使用死信队列处理失败消息
- 记录错误日志

### 4. 性能优化

- 使用批量发送
- 合理设置预取数量
- 使用连接池

## 📚 相关文档

- [消息队列文档](../../../docs/microservices/08-message-queue.md)
- [RabbitMQ 官方文档](https://www.rabbitmq.com/documentation.html)
- [Kafka 官方文档](https://kafka.apache.org/documentation/)

---

**🎉 开始使用消息队列，实现异步通信和解耦！**

