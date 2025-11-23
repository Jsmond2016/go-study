package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 连接 RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("连接 RabbitMQ 失败: %v", err)
	}
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("创建通道失败: %v", err)
	}
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 持久化
		false,   // 自动删除
		false,   // 排他
		false,   // 无等待
		nil,     // 参数
	)
	if err != nil {
		log.Fatalf("声明队列失败: %v", err)
	}

	// 消费消息
	msgs, err := ch.Consume(
		q.Name, // 队列
		"",     // 消费者标签
		true,   // 自动确认
		false,  // 排他
		false,  // 无本地
		false,  // 无等待
		nil,    // 参数
	)
	if err != nil {
		log.Fatalf("注册消费者失败: %v", err)
	}

	log.Println("等待消息...")
	for d := range msgs {
		log.Printf("收到消息: %s", d.Body)
	}
}

