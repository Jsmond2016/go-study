package main

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	pb "go-study/examples/microservices/01-grpc/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

// 初始化测试用的 in-memory gRPC 服务器
func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &StreamingServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

// 创建测试用的 gRPC 连接
func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

// TestSayHello 测试简单 RPC
func TestSayHello(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 测试用例
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"简单请求", "Alice", "Hello, Alice!"},
		{"空字符串", "", "Hello, !"},
		{"中文名字", "张三", "Hello, 张三!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: tt.input})
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, resp.GetMessage())
		})
	}
}

// TestSayHelloStream 测试服务端流式 RPC
func TestSayHelloStream(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 发起流式调用
	stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Bob"})
	assert.NoError(t, err)

	// 接收流式响应
	count := 0
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		assert.NoError(t, err)
		count++

		// 验证消息格式
		assert.Contains(t, resp.GetMessage(), "Bob")
		assert.Contains(t, resp.GetMessage(), "message")
	}

	// 验证接收到 5 条消息
	assert.Equal(t, 5, count, "应该接收到 5 条消息")
}

// TestCollectHello 测试客户端流式 RPC
func TestCollectHello(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 发起流式调用
	stream, err := client.CollectHello(ctx)
	assert.NoError(t, err)

	// 发送多个请求
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		err := stream.Send(&pb.HelloRequest{Name: name})
		assert.NoError(t, err)
	}

	// 关闭发送并接收响应
	resp, err := stream.CloseAndRecv()
	assert.NoError(t, err)

	// 验证响应
	assert.Contains(t, resp.GetMessage(), "3")
	assert.Contains(t, resp.GetMessage(), "Alice")
	assert.Contains(t, resp.GetMessage(), "Bob")
	assert.Contains(t, resp.GetMessage(), "Charlie")
}

// TestChatHello 测试双向流式 RPC
func TestChatHello(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 发起双向流式调用
	stream, err := client.ChatHello(ctx)
	assert.NoError(t, err)

	// 使用 channel 同步
	waitc := make(chan struct{})
	receivedCount := 0

	// 启动接收 goroutine
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			assert.NoError(t, err)
			receivedCount++

			// 验证响应格式
			assert.Contains(t, resp.GetMessage(), "Echo")
		}
	}()

	// 发送多个请求
	names := []string{"David", "Eve", "Frank"}
	for _, name := range names {
		err := stream.Send(&pb.HelloRequest{Name: name})
		assert.NoError(t, err)
		time.Sleep(100 * time.Millisecond) // 给服务端处理时间
	}

	// 关闭发送端
	err = stream.CloseSend()
	assert.NoError(t, err)

	// 等待接收完成
	<-waitc

	// 验证接收到的消息数量
	assert.Equal(t, len(names), receivedCount, "应该接收到与发送相同数量的消息")
}

// TestStreamingServer_Timeout 测试超时场景
func TestStreamingServer_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 这个调用应该超时（因为服务端流会发送 5 条消息，每条间隔 1 秒）
	stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Timeout"})
	assert.NoError(t, err)

	// 尝试接收，应该在超时前只能收到部分消息
	count := 0
	for {
		_, err := stream.Recv()
		if err != nil {
			// 预期会超时
			assert.Error(t, err)
			break
		}
		count++
	}

	// 由于超时时间是 100ms，应该收不到消息或只收到很少的消息
	assert.Less(t, count, 5, "超时情况下不应该收到所有消息")
}

// TestStreamingServer_ConcurrentCalls 测试并发调用
func TestStreamingServer_ConcurrentCalls(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 并发调用 10 次
	const concurrency = 10
	done := make(chan bool, concurrency)

	for i := 0; i < concurrency; i++ {
		go func(id int) {
			resp, err := client.SayHello(ctx, &pb.HelloRequest{
				Name: fmt.Sprintf("User-%d", id),
			})
			assert.NoError(t, err)
			assert.Contains(t, resp.GetMessage(), fmt.Sprintf("User-%d", id))
			done <- true
		}(i)
	}

	// 等待所有调用完成
	for i := 0; i < concurrency; i++ {
		<-done
	}
}

// BenchmarkSayHello 基准测试
func BenchmarkSayHello(b *testing.B) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Benchmark"})
		if err != nil {
			b.Fatalf("RPC failed: %v", err)
		}
	}
}

// BenchmarkSayHelloStream 流式 RPC 基准测试
func BenchmarkSayHelloStream(b *testing.B) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Benchmark"})
		if err != nil {
			b.Fatalf("Failed to call stream: %v", err)
		}

		for {
			_, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				b.Fatalf("Stream error: %v", err)
			}
		}
	}
}
