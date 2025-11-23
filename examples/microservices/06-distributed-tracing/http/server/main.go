package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	port     = flag.Int("port", 8080, "HTTP 服务器端口")
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
			semconv.ServiceNameKey.String("http-server"),
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

	// 创建 HTTP 处理器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("http-handler")
		ctx, span := tracer.Start(ctx, "handleRequest")
		defer span.End()

		span.SetAttributes(
			attribute.String("http.method", r.Method),
			attribute.String("http.path", r.URL.Path),
			attribute.String("http.scheme", r.URL.Scheme),
			attribute.String("http.host", r.Host),
		)

		// 模拟业务处理
		processRequest(ctx, tracer)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, World! Request processed at %s", time.Now().Format(time.RFC3339))
	})

	// 使用 otelhttp 包装处理器
	wrappedHandler := otelhttp.NewHandler(handler, "http-server")

	http.Handle("/", wrappedHandler)
	http.Handle("/api", wrappedHandler)

	log.Printf("HTTP 服务器（带追踪）监听在端口 %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func processRequest(ctx context.Context, tracer trace.Tracer) {
	ctx, span := tracer.Start(ctx, "processRequest")
	defer span.End()

	// 模拟处理时间
	time.Sleep(10 * time.Millisecond)

	span.SetAttributes(attribute.String("request.status", "processed"))
}

