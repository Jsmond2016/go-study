package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "服务器地址")
	jaegerURL  = flag.String("jaeger", "http://localhost:14268/api/traces", "Jaeger 地址")
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
			semconv.ServiceNameKey.String("grpc-client"),
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

	// 创建 gRPC 连接，添加追踪拦截器
	conn, err := grpc.Dial(
		*serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	log.Printf("已连接到 gRPC 服务器: %s", *serverAddr)

	// 使用连接调用服务
	// client := pb.NewYourServiceClient(conn)
	// ctx := context.Background()
	// resp, err := client.YourMethod(ctx, &pb.Request{})
	// if err != nil {
	// 	log.Fatalf("调用失败: %v", err)
	// }
	// log.Printf("响应: %v", resp)
}

