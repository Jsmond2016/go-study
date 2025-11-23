package main

import (
	"fmt"
	"log"

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

	// 发布配置
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  "example",
		Group:   "DEFAULT_GROUP",
		Content: "key=value\nkey2=value2\ndb.host=localhost\ndb.port=3306",
	})
	if err != nil {
		log.Fatalf("发布配置失败: %v", err)
	}
	fmt.Printf("发布配置成功: %v\n", success)

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "example",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		log.Fatalf("获取配置失败: %v", err)
	}
	fmt.Printf("配置内容:\n%s\n", content)

	// 删除配置（可选）
	// success, err = configClient.DeleteConfig(vo.ConfigParam{
	// 	DataId: "example",
	// 	Group:  "DEFAULT_GROUP",
	// })
	// if err != nil {
	// 	log.Fatalf("删除配置失败: %v", err)
	// }
	// fmt.Printf("删除配置成功: %v\n", success)
}

