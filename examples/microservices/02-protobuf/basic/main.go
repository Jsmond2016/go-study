package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	pb "go-study/examples/microservices/02-protobuf/basic/proto"
)

func main() {
	fmt.Println("=== Protocol Buffers 基础示例 ===")

	// 创建消息
	user := &pb.User{
		Id:     1,
		Name:   "Alice",
		Email:  "alice@example.com",
		Age:    25,
		Active: true,
	}

	fmt.Printf("原始消息: %+v\n", user)

	// 序列化
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("序列化大小: %d bytes\n", len(data))

	// 反序列化
	var newUser pb.User
	if err := proto.Unmarshal(data, &newUser); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("反序列化消息: %+v\n", newUser)
	fmt.Printf("消息相等: %v\n", proto.Equal(user, &newUser))
}

