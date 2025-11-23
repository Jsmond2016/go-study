package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// LoggingInterceptor 日志拦截器
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	log.Printf("Method: %s, Request: %+v", info.FullMethod, req)

	resp, err := handler(ctx, req)

	duration := time.Since(start)
	if err != nil {
		log.Printf("Method: %s, Error: %v, Duration: %v", info.FullMethod, err, duration)
	} else {
		log.Printf("Method: %s, Response: %+v, Duration: %v", info.FullMethod, resp, duration)
	}

	return resp, err
}

// LoggingStreamInterceptor 流式日志拦截器
func LoggingStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("Stream Method: %s", info.FullMethod)
	return handler(srv, ss)
}

