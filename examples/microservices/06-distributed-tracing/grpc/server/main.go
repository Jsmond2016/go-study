package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	port     = flag.Int("port", 50051, "服务端口")
	jaegerURL = flag.String("jaeger", "http://localhost:14268/api/traces", "Jaeger 地址")
)

// initTracer 初始化 Tracer Provider
func initTracer(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("grpc-server"),
			semconv.ServiceVersionKey.String("1.0.0"),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	flag.Parse()

	// 初始化 Tracer Provider
	tp, err := initTracer(*jaegerURL)
	if err != nil {
		log.Fatalf("初始化 Tracer Provider 失败: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("关闭 Tracer Provider 失败: %v", err)
		}
	}()

	// 创建 gRPC 服务器，添加追踪拦截器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	// 注册服务（这里需要实际的 proto 定义）
	// pb.RegisterYourServiceServer(s, &server{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	log.Printf("gRPC 服务器（带追踪）监听在端口 %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

