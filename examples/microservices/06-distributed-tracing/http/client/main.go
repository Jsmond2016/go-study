package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	serverURL = flag.String("url", "http://localhost:8080", "服务器 URL")
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
			semconv.ServiceNameKey.String("http-client"),
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

	// 创建 HTTP 客户端，使用 otelhttp 包装
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", *serverURL, nil)
	if err != nil {
		log.Fatalf("创建请求失败: %v", err)
	}

	log.Printf("发送请求到: %s", *serverURL)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}

	log.Printf("响应状态: %s", resp.Status)
	log.Printf("响应内容: %s", string(body))
}

