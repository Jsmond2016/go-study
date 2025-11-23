package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "go-study/examples/microservices/01-grpc/proto"
)

func main() {
	// 连接到服务器
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewGreeterClient(conn)

	// 创建带认证信息的 context
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "valid-token-123")

	// 调用需要认证的服务
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: "Authenticated User"}
	reply, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s", reply.GetMessage())
}

