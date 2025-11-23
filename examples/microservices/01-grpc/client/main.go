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

func main() {
	// 连接到服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewGreeterClient(conn)

	// 示例 1: 简单的 RPC 调用
	fmt.Println("=== 示例 1: 简单 RPC 调用 ===")
	simpleRPC(client)

	// 示例 2: 服务端流式 RPC
	fmt.Println("\n=== 示例 2: 服务端流式 RPC ===")
	serverStreaming(client)

	// 示例 3: 客户端流式 RPC
	fmt.Println("\n=== 示例 3: 客户端流式 RPC ===")
	clientStreaming(client)

	// 示例 4: 双向流式 RPC
	fmt.Println("\n=== 示例 4: 双向流式 RPC ===")
	bidirectionalStreaming(client)
}

// simpleRPC 演示简单的 RPC 调用
func simpleRPC(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: "World"}
	reply, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Printf("Response: %s\n", reply.GetMessage())
}

// serverStreaming 演示服务端流式 RPC
func serverStreaming(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: "Stream"}
	stream, err := client.SayHelloStream(ctx, req)
	if err != nil {
		log.Fatalf("could not stream: %v", err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving: %v", err)
		}
		fmt.Printf("Received: %s\n", reply.GetMessage())
	}
}

// clientStreaming 演示客户端流式 RPC
func clientStreaming(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.CollectHello(ctx)
	if err != nil {
		log.Fatalf("could not collect: %v", err)
	}

	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending: %v", err)
		}
		fmt.Printf("Sent: %s\n", name)
		time.Sleep(500 * time.Millisecond)
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}
	fmt.Printf("Response: %s\n", reply.GetMessage())
}

// bidirectionalStreaming 演示双向流式 RPC
func bidirectionalStreaming(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ChatHello(ctx)
	if err != nil {
		log.Fatalf("could not chat: %v", err)
	}

	// 发送消息的 goroutine
	go func() {
		messages := []string{"Hello", "How are you?", "Goodbye"}
		for _, msg := range messages {
			req := &pb.HelloRequest{Name: msg}
			if err := stream.Send(req); err != nil {
				log.Printf("error sending: %v", err)
				return
			}
			fmt.Printf("Sent: %s\n", msg)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// 接收消息
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving: %v", err)
		}
		fmt.Printf("Received: %s\n", reply.GetMessage())
	}
}

