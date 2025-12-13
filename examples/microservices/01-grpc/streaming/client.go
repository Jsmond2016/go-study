package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "go-study/examples/microservices/01-grpc/proto"
)

const (
	serverAddr = "localhost:50051"
)

func main() {
	// 连接到服务器
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 演示所有四种 RPC 类型
	fmt.Println("=== gRPC Streaming 示例 ===\n")

	// 1. 简单 RPC
	simpleRPC(client)

	time.Sleep(2 * time.Second)

	// 2. 服务端流式 RPC
	serverStreaming(client)

	time.Sleep(2 * time.Second)

	// 3. 客户端流式 RPC
	clientStreaming(client)

	time.Sleep(2 * time.Second)

	// 4. 双向流式 RPC
	bidirectionalStreaming(client)
}

// simpleRPC 演示简单的 RPC 调用（非流式）
func simpleRPC(client pb.GreeterClient) {
	fmt.Println("--- 1. 简单 RPC (Simple RPC) ---")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("Simple RPC failed: %v", err)
	}

	fmt.Printf("服务端响应: %s\n\n", resp.GetMessage())
}

// serverStreaming 演示服务端流式 RPC
// 客户端发送一个请求，服务端返回多个响应（流）
func serverStreaming(client pb.GreeterClient) {
	fmt.Println("--- 2. 服务端流式 RPC (Server Streaming RPC) ---")
	fmt.Println("客户端发送一个请求，服务端返回多个响应流...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 发送一个请求
	stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Bob"})
	if err != nil {
		log.Fatalf("Server streaming failed: %v", err)
	}

	// 接收多个响应
	count := 0
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			// 流结束
			fmt.Printf("\n服务端流式 RPC 完成，共接收 %d 条消息\n\n", count)
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		count++
		fmt.Printf("  [%d] %s\n", count, resp.GetMessage())
	}
}

// clientStreaming 演示客户端流式 RPC
// 客户端发送多个请求（流），服务端返回一个响应
func clientStreaming(client pb.GreeterClient) {
	fmt.Println("--- 3. 客户端流式 RPC (Client Streaming RPC) ---")
	fmt.Println("客户端发送多个请求流，服务端返回一个汇总响应...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.CollectHello(ctx)
	if err != nil {
		log.Fatalf("Client streaming failed: %v", err)
	}

	// 发送多个请求
	names := []string{"Charlie", "David", "Eve", "Frank", "Grace"}
	for i, name := range names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		fmt.Printf("  [%d] 发送: %s\n", i+1, name)
		time.Sleep(500 * time.Millisecond)
	}

	// 关闭发送并接收响应
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}

	fmt.Printf("\n服务端汇总响应: %s\n\n", resp.GetMessage())
}

// bidirectionalStreaming 演示双向流式 RPC
// 客户端和服务端都可以发送多个消息（流）
func bidirectionalStreaming(client pb.GreeterClient) {
	fmt.Println("--- 4. 双向流式 RPC (Bidirectional Streaming RPC) ---")
	fmt.Println("客户端和服务端同时发送和接收消息流...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.ChatHello(ctx)
	if err != nil {
		log.Fatalf("Bidirectional streaming failed: %v", err)
	}

	// 使用 channel 来同步发送和接收
	waitc := make(chan struct{})

	// 启动接收 goroutine
	go func() {
		count := 0
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				// 服务端关闭了流
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error receiving: %v", err)
			}

			count++
			fmt.Printf("  ← 接收 [%d]: %s\n", count, resp.GetMessage())
		}
	}()

	// 发送多个请求
	names := []string{"Henry", "Iris", "Jack"}
	for i, name := range names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending: %v", err)
		}
		fmt.Printf("  → 发送 [%d]: %s\n", i+1, name)
		time.Sleep(2 * time.Second)
	}

	// 关闭发送端
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing send: %v", err)
	}

	// 等待接收完成
	<-waitc

	fmt.Println("\n双向流式 RPC 完成\n")
}
