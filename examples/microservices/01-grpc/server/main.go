package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "go-study/examples/microservices/01-grpc/proto"
)

// GreeterServer 实现 Greeter 服务
type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现简单的 RPC 调用
func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

// SayHelloStream 实现服务端流式 RPC
func (s *GreeterServer) SayHelloStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServer) error {
	log.Printf("Streaming for: %v", req.GetName())
	
	for i := 0; i < 5; i++ {
		reply := &pb.HelloReply{
			Message: fmt.Sprintf("Hello %s, message %d", req.GetName(), i+1),
		}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
	
	return nil
}

// CollectHello 实现客户端流式 RPC
func (s *GreeterServer) CollectHello(stream pb.Greeter_CollectHelloServer) error {
	var names []string
	
	for {
		req, err := stream.Recv()
		if err != nil {
			// 流结束
			break
		}
		names = append(names, req.GetName())
		log.Printf("Received name: %s", req.GetName())
	}
	
	reply := &pb.HelloReply{
		Message: fmt.Sprintf("Collected %d names: %v", len(names), names),
	}
	return stream.SendAndClose(reply)
}

// ChatHello 实现双向流式 RPC
func (s *GreeterServer) ChatHello(stream pb.Greeter_ChatHelloServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		
		log.Printf("Received: %v", req.GetName())
		
		reply := &pb.HelloReply{
			Message: fmt.Sprintf("Echo: Hello, %s!", req.GetName()),
		}
		
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	
	// 注册服务
	pb.RegisterGreeterServer(s, &GreeterServer{})

	log.Println("gRPC server listening on :50051")
	
	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

