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

## 📚 相关文档

- [消息队列文档](../../../docs/microservices/08-message-queue.md)
- [RabbitMQ 官方文档](https://www.rabbitmq.com/documentation.html)
- [Kafka 官方文档](https://kafka.apache.org/documentation/)

---

**🎉 开始使用消息队列，实现异步通信和解耦！**

