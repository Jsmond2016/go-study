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

	// 发送消息
	body := "Hello World!"
	err = ch.Publish(
		"",     // 交换机
		q.Name, // 路由键
		false,  // 强制
		false,  // 立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("发送消息失败: %v", err)
	}
	log.Printf("发送消息: %s", body)
}

