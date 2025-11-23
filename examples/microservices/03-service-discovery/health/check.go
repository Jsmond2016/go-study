package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var (
	serviceID  = flag.String("id", "", "服务 ID")
	checkURL   = flag.String("url", "http://localhost:8080/health", "健康检查 URL")
	consulAddr = flag.String("consul", "localhost:8500", "Consul 地址")
	interval   = flag.Duration("interval", 10*time.Second, "检查间隔")
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

	// 创建 HTTP 客户端
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 启动健康检查
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("开始健康检查，服务 ID: %s, URL: %s, 间隔: %v", *serviceID, *checkURL, *interval)

	for {
		select {
		case <-ticker.C:
			// 执行健康检查
			healthy := performHealthCheck(ctx, httpClient, *checkURL)

			// 更新 Consul 中的服务状态
			status := consulapi.HealthPassing
			output := "健康检查通过"
			if !healthy {
				status = consulapi.HealthCritical
				output = "健康检查失败"
			}

			err := client.Agent().UpdateTTL(*serviceID, output, status)
			if err != nil {
				log.Printf("更新服务状态失败: %v", err)
			} else {
				log.Printf("健康检查结果: %s - %s", status, output)
			}

		case <-ctx.Done():
			return
		}
	}
}

// performHealthCheck 执行健康检查
func performHealthCheck(ctx context.Context, client *http.Client, url string) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("健康检查请求失败: %v", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

