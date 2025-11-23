package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	serviceName = flag.String("service", "user-service", "服务名称")
	servicePort = flag.Int("port", 8080, "服务端口")
	etcdAddr    = flag.String("etcd", "localhost:2379", "etcd 地址")
	ttl         = flag.Int("ttl", 30, "服务 TTL（秒）")
)

func main() {
	flag.Parse()

	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*etcdAddr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("创建 etcd 客户端失败: %v", err)
	}
	defer cli.Close()

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

	// 服务注册键
	serviceKey := fmt.Sprintf("/services/%s/%s-%d", *serviceName, hostname, *servicePort)
	serviceValue := fmt.Sprintf("%s:%d", serviceIP, *servicePort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建租约
	lease, err := cli.Grant(ctx, int64(*ttl))
	if err != nil {
		log.Fatalf("创建租约失败: %v", err)
	}

	// 注册服务（带租约）
	_, err = cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Fatalf("注册服务失败: %v", err)
	}
	log.Printf("服务注册成功: %s@%s:%d (租约ID: %x)", *serviceName, serviceIP, *servicePort, lease.ID)

	// 启动心跳续约
	keepAliveChan, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		log.Fatalf("启动心跳失败: %v", err)
	}

	// 启动 gRPC 服务（示例）
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *servicePort))
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	log.Printf("gRPC 服务启动在端口 %d", *servicePort)

	// 优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 监听心跳响应
		for ka := range keepAliveChan {
			log.Printf("心跳续约成功: TTL=%d", ka.TTL)
		}
	}()

	// 等待关闭信号
	<-sigChan
	log.Println("收到关闭信号，开始优雅关闭...")

	// 注销服务
	_, err = cli.Delete(ctx, serviceKey)
	if err != nil {
		log.Printf("注销服务失败: %v", err)
	} else {
		log.Println("服务已注销")
	}

	// 撤销租约
	_, err = cli.Revoke(ctx, lease.ID)
	if err != nil {
		log.Printf("撤销租约失败: %v", err)
	}

	log.Println("服务已关闭")
}

