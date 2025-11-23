// Package main 演示如何使用 Consul 进行服务注册
//
// 本示例展示了：
// 1. 如何连接到 Consul
// 2. 如何获取本机 IP 地址
// 3. 如何注册服务到 Consul
// 4. 如何实现优雅关闭和自动注销
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

var (
	// serviceName 服务名称，用于在 Consul 中标识服务
	serviceName = flag.String("service", "user-service", "服务名称")
	// servicePort 服务端口，gRPC 服务监听的端口
	servicePort = flag.Int("port", 8080, "服务端口")
	// consulAddr Consul 服务器地址
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
)

// main 主函数，演示服务注册流程
func main() {
	flag.Parse()

	// 创建 Consul 客户端
	// 使用默认配置，但指定 Consul 服务器地址
	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}

	// 获取本机主机名，用于生成唯一的服务 ID
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("获取主机名失败: %v", err)
	}

	// 获取本机 IP 地址
	// 通过连接外部地址来获取本机的网络接口 IP
	// 注意：这种方法在某些网络环境下可能不准确，生产环境建议使用更可靠的方法
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalf("获取本机 IP 失败: %v", err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	serviceIP := localAddr.IP.String()

	// 构建服务注册信息
	// ID: 唯一标识符，使用服务名-主机名-端口组合
	// Name: 服务名称，用于服务发现
	// Tags: 服务标签，可用于过滤和路由
	// Address/Port: 服务地址和端口
	// Check: 健康检查配置
	registration := &consulapi.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", *serviceName, hostname, *servicePort),
		Name:    *serviceName,
		Tags:    []string{"v1", "grpc"},
		Address: serviceIP,
		Port:    *servicePort,
		Check: &consulapi.AgentServiceCheck{
			// TCP 健康检查：检查服务端口是否可访问
			TCP: fmt.Sprintf("%s:%d", serviceIP, *servicePort),
			// Interval: 健康检查间隔
			Interval: "10s",
			// Timeout: 健康检查超时时间
			Timeout: "5s",
			// DeregisterCriticalServiceAfter: 服务不健康后自动注销的时间
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	// 注册服务到 Consul
	// 注册成功后，其他服务可以通过服务名称发现该服务
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("注册服务失败: %v", err)
	}
	log.Printf("服务注册成功: %s@%s:%d", *serviceName, serviceIP, *servicePort)

	// 启动 gRPC 服务（示例）
	// 在实际应用中，这里应该注册具体的 gRPC 服务实现
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *servicePort))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	log.Printf("gRPC 服务启动在端口 %d", *servicePort)

	// 启动 HTTP 健康检查服务（可选）
	// 如果启用，Consul 可以通过 HTTP 检查服务健康状态
	// go startHealthCheck(*servicePort)

	// 实现优雅关闭
	// 当收到关闭信号时，先注销服务，再停止 gRPC 服务
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听系统信号，实现优雅关闭
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("收到关闭信号，开始优雅关闭...")

		// 从 Consul 注销服务
		// 这确保其他服务不会继续尝试连接已关闭的服务
		err := client.Agent().ServiceDeregister(registration.ID)
		if err != nil {
			log.Printf("注销服务失败: %v", err)
		} else {
			log.Println("服务已注销")
		}

		// 优雅停止 gRPC 服务
		// GracefulStop 会等待所有正在处理的请求完成
		s.GracefulStop()
		cancel()
	}()

	// 运行 gRPC 服务
	// Serve 会阻塞直到服务停止
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务运行失败: %v", err)
	}

	// 等待优雅关闭完成
	<-ctx.Done()
	log.Println("服务已关闭")
}

// startHealthCheck 启动 HTTP 健康检查服务
//
// 参数:
//   - port: 基础端口号，健康检查服务会在该端口基础上加 10000
//
// 健康检查端点: /health
// 返回 200 OK 表示服务健康
func startHealthCheck(port int) {
	httpPort := port + 10000 // 使用不同的端口避免与 gRPC 服务冲突
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: mux,
	}

	log.Printf("健康检查服务启动在端口 %d", httpPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("健康检查服务启动失败: %v", err)
	}
}

