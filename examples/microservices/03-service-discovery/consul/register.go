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
	serviceName = flag.String("service", "user-service", "服务名称")
	servicePort = flag.Int("port", 8080, "服务端口")
	consulAddr  = flag.String("consul", "localhost:8500", "Consul 地址")
)

func main() {
	flag.Parse()

	// 创建 Consul 客户端
	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}

	// 获取本机 IP
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("获取主机名失败: %v", err)
	}

	// 获取本机 IP 地址
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalf("获取本机 IP 失败: %v", err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	serviceIP := localAddr.IP.String()

	// 服务注册信息
	registration := &consulapi.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", *serviceName, hostname, *servicePort),
		Name:    *serviceName,
		Tags:    []string{"v1", "grpc"},
		Address: serviceIP,
		Port:    *servicePort,
		Check: &consulapi.AgentServiceCheck{
			TCP:                            fmt.Sprintf("%s:%d", serviceIP, *servicePort),
			Interval:                       "10s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	// 注册服务
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("注册服务失败: %v", err)
	}
	log.Printf("服务注册成功: %s@%s:%d", *serviceName, serviceIP, *servicePort)

	// 启动 gRPC 服务（示例）
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *servicePort))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	log.Printf("gRPC 服务启动在端口 %d", *servicePort)

	// 启动 HTTP 健康检查服务（可选）
	// go startHealthCheck(*servicePort)

	// 优雅关闭
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("收到关闭信号，开始优雅关闭...")

		// 注销服务
		err := client.Agent().ServiceDeregister(registration.ID)
		if err != nil {
			log.Printf("注销服务失败: %v", err)
		} else {
			log.Println("服务已注销")
		}

		// 停止 gRPC 服务
		s.GracefulStop()
		cancel()
	}()

	// 运行服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务运行失败: %v", err)
	}

	<-ctx.Done()
	log.Println("服务已关闭")
}

// startHealthCheck 启动 HTTP 健康检查服务
func startHealthCheck(port int) {
	httpPort := port + 10000 // 使用不同的端口避免冲突
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

