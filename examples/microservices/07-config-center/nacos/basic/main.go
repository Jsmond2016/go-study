package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// 配置 Nacos 客户端
	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		log.Fatalf("创建配置客户端失败: %v", err)
	}

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "example",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		log.Fatalf("获取配置失败: %v", err)
	}
	fmt.Printf("配置内容: %s\n", content)

	// 监听配置变更
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "example",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("配置变更: namespace=%s, group=%s, dataId=%s, data=%s\n",
				namespace, group, dataId, data)
		},
	})
	if err != nil {
		log.Fatalf("监听配置失败: %v", err)
	}

	fmt.Println("配置监听已启动，等待配置变更...")

	// 保持运行
	select {}
}

