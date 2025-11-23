package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// 配置生产者
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("创建生产者失败: %v", err)
	}
	defer producer.Close()

	// 发送消息
	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello Kafka!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("发送消息失败: %v", err)
	}
	log.Printf("消息已发送: partition=%d, offset=%d", partition, offset)
}

