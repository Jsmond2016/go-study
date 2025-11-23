package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go-study/examples/microservices/01-grpc/interceptors"
	pb "go-study/examples/microservices/01-grpc/proto"
)

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建限流器（100 tokens, 10 per second）
	rateLimiter := interceptors.NewTokenBucket(100, 10)

	// 创建 gRPC 服务器，添加拦截器
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptors.LoggingInterceptor,
			interceptors.AuthInterceptor,
			interceptors.RateLimitInterceptor(rateLimiter),
		),
		grpc.ChainStreamInterceptor(interceptors.LoggingStreamInterceptor),
	}
	s := grpc.NewServer(opts...)

	// 注册服务
	pb.RegisterGreeterServer(s, &GreeterServer{})

	log.Println("gRPC server with interceptors listening on :50052")
	log.Println("Note: Use 'authorization: valid-token-123' in metadata to authenticate")

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

