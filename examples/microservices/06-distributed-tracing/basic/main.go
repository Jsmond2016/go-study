package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

// initTracer 初始化 Tracer Provider
func initTracer(url string) (*tracesdk.TracerProvider, error) {
	// 创建 Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	// 创建 Tracer Provider
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("basic-tracing-example"),
			semconv.ServiceVersionKey.String("1.0.0"),
		)),
	)

	// 设置为全局 Tracer Provider
	otel.SetTracerProvider(tp)

	return tp, nil
}

func main() {
	// 初始化 Tracer Provider
	tp, err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatalf("初始化 Tracer Provider 失败: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("关闭 Tracer Provider 失败: %v", err)
		}
	}()

	ctx := context.Background()
	tracer := otel.Tracer("basic-example")

	// 创建根 Span
	ctx, rootSpan := tracer.Start(ctx, "main")
	defer rootSpan.End()

	fmt.Println("开始处理订单...")

	// 模拟处理订单
	if err := processOrder(ctx, tracer, "order-123"); err != nil {
		rootSpan.RecordError(err)
		log.Printf("处理订单失败: %v", err)
		return
	}

	fmt.Println("订单处理完成！")
}

// processOrder 处理订单
func processOrder(ctx context.Context, tracer trace.Tracer, orderID string) error {
	ctx, span := tracer.Start(ctx, "processOrder")
	defer span.End()

	span.SetAttributes(
		attribute.String("order.id", orderID),
		attribute.String("order.status", "processing"),
	)

	// 验证订单
	if err := validateOrder(ctx, tracer, orderID); err != nil {
		span.RecordError(err)
		span.SetAttributes(attribute.String("order.status", "failed"))
		return err
	}

	// 检查库存
	if err := checkInventory(ctx, tracer, orderID); err != nil {
		span.RecordError(err)
		span.SetAttributes(attribute.String("order.status", "failed"))
		return err
	}

	// 创建订单
	if err := createOrder(ctx, tracer, orderID); err != nil {
		span.RecordError(err)
		span.SetAttributes(attribute.String("order.status", "failed"))
		return err
	}

	span.SetAttributes(attribute.String("order.status", "completed"))
	return nil
}

// validateOrder 验证订单
func validateOrder(ctx context.Context, tracer trace.Tracer, orderID string) error {
	ctx, span := tracer.Start(ctx, "validateOrder")
	defer span.End()

	span.SetAttributes(attribute.String("order.id", orderID))

	// 模拟验证逻辑
	time.Sleep(10 * time.Millisecond)

	span.SetAttributes(
		attribute.Bool("order.valid", true),
		attribute.String("validation.result", "success"),
	)

	return nil
}

// checkInventory 检查库存
func checkInventory(ctx context.Context, tracer trace.Tracer, orderID string) error {
	ctx, span := tracer.Start(ctx, "checkInventory")
	defer span.End()

	span.SetAttributes(attribute.String("order.id", orderID))

	// 模拟库存检查
	time.Sleep(20 * time.Millisecond)

	span.SetAttributes(
		attribute.Bool("inventory.available", true),
		attribute.Int("inventory.quantity", 100),
	)

	return nil
}

// createOrder 创建订单
func createOrder(ctx context.Context, tracer trace.Tracer, orderID string) error {
	ctx, span := tracer.Start(ctx, "createOrder")
	defer span.End()

	span.SetAttributes(attribute.String("order.id", orderID))

	// 模拟创建订单
	time.Sleep(15 * time.Millisecond)

	span.SetAttributes(
		attribute.String("order.status", "created"),
		attribute.String("order.created_at", time.Now().Format(time.RFC3339)),
	)

	return nil
}

