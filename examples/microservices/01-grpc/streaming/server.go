package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "go-study/examples/microservices/01-grpc/proto"
)

// StreamingServer 实现所有流式 RPC
type StreamingServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现简单的 RPC 调用（非流式）
func (s *StreamingServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("[Simple RPC] Received: %v", req.GetName())
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

// SayHelloStream 实现服务端流式 RPC
// 服务端流式：客户端发送一个请求，服务端返回多个响应（流）
func (s *StreamingServer) SayHelloStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServer) error {
	log.Printf("[Server Streaming] Streaming for: %v", req.GetName())

	// 服务端发送多个消息
	for i := 0; i < 5; i++ {
		reply := &pb.HelloReply{
			Message: fmt.Sprintf("Hello %s, message %d (timestamp: %d)",
				req.GetName(), i+1, time.Now().Unix()),
		}

		log.Printf("[Server Streaming] Sending message %d to %s", i+1, req.GetName())

		if err := stream.Send(reply); err != nil {
			log.Printf("[Server Streaming] Error sending: %v", err)
			return err
		}

		// 模拟处理延迟
		time.Sleep(time.Second)
	}

	log.Printf("[Server Streaming] Finished streaming for %s", req.GetName())
	return nil
}

// CollectHello 实现客户端流式 RPC
// 客户端流式：客户端发送多个请求（流），服务端返回一个响应
func (s *StreamingServer) CollectHello(stream pb.Greeter_CollectHelloServer) error {
	log.Printf("[Client Streaming] Started receiving names")

	var names []string
	count := 0

	// 接收客户端发送的多个消息
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 客户端发送完毕
			log.Printf("[Client Streaming] Client finished sending, received %d names", count)

			// 返回一个汇总响应
			reply := &pb.HelloReply{
				Message: fmt.Sprintf("Collected %d names: %v", count, names),
			}
			return stream.SendAndClose(reply)
		}
		if err != nil {
			log.Printf("[Client Streaming] Error receiving: %v", err)
			return err
		}

		count++
		names = append(names, req.GetName())
		log.Printf("[Client Streaming] Received name #%d: %s", count, req.GetName())
	}
}

// ChatHello 实现双向流式 RPC
// 双向流式：客户端和服务端都可以发送多个消息（流）
func (s *StreamingServer) ChatHello(stream pb.Greeter_ChatHelloServer) error {
	log.Printf("[Bidirectional Streaming] Chat session started")

	messageCount := 0

	// 同时接收和发送消息
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("[Bidirectional Streaming] Client closed connection after %d messages", messageCount)
			return nil
		}
		if err != nil {
			log.Printf("[Bidirectional Streaming] Error receiving: %v", err)
			return err
		}

		messageCount++
		log.Printf("[Bidirectional Streaming] Received message #%d from %s", messageCount, req.GetName())

		// 立即回复
		reply := &pb.HelloReply{
			Message: fmt.Sprintf("Echo #%d: Hello %s! (timestamp: %d)",
				messageCount, req.GetName(), time.Now().Unix()),
		}

		if err := stream.Send(reply); err != nil {
			log.Printf("[Bidirectional Streaming] Error sending: %v", err)
			return err
		}

		log.Printf("[Bidirectional Streaming] Sent reply #%d", messageCount)
	}
}

func main() {
	// 创建 TCP 监听器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册 Greeter 服务
	pb.RegisterGreeterServer(s, &StreamingServer{})

	log.Printf("Server listening at %v", lis.Addr())
	log.Println("支持以下 RPC 方法:")
	log.Println("  1. SayHello - 简单 RPC")
	log.Println("  2. SayHelloStream - 服务端流式 RPC")
	log.Println("  3. CollectHello - 客户端流式 RPC")
	log.Println("  4. ChatHello - 双向流式 RPC")

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
