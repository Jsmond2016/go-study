package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var (
	serviceID  = flag.String("id", "", "服务 ID")
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
	interval   = flag.Duration("interval", 10*time.Second, "心跳间隔")
)

func main() {
	flag.Parse()

	if *serviceID == "" {
		log.Fatal("服务 ID 不能为空")
	}

	// 创建 Consul 客户端
	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}

	// 启动心跳
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("开始发送心跳，服务 ID: %s, 间隔: %v", *serviceID, *interval)

	for {
		select {
		case <-ticker.C:
			err := client.Agent().UpdateTTL(*serviceID, "heartbeat", consulapi.HealthPassing)
			if err != nil {
				log.Printf("发送心跳失败: %v", err)
			} else {
				log.Printf("心跳发送成功: %s", time.Now().Format("2006-01-02 15:04:05"))
			}

		case <-ctx.Done():
			return
		}
	}
}

