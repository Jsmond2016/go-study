package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// 配置消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("创建消费者失败: %v", err)
	}
	defer consumer.Close()

	// 创建分区消费者
	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("创建分区消费者失败: %v", err)
	}
	defer partitionConsumer.Close()

	log.Println("等待消息...")
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("收到消息: topic=%s, partition=%d, offset=%d, value=%s",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("消费错误: %v", err)
		}
	}
}

