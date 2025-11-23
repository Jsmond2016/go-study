package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

func main() {
	// 配置 Apollo 客户端
	appConfig := &config.AppConfig{
		AppID:          "your-app-id",
		Cluster:        "default",
		NamespaceName:  "application",
		IP:             "http://localhost:8080",
		IsBackupConfig: true,
	}

	// 创建客户端
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return appConfig, nil
	})
	if err != nil {
		log.Fatalf("启动 Apollo 客户端失败: %v", err)
	}

	// 等待配置加载
	time.Sleep(1 * time.Second)

	// 获取配置
	value := client.GetStringValue("key", "default-value")
	fmt.Printf("配置值: %s\n", value)

	// 获取不同类型的配置
	intValue := client.GetIntValue("int-key", 0)
	boolValue := client.GetBoolValue("bool-key", false)
	fmt.Printf("整数配置: %d, 布尔配置: %v\n", intValue, boolValue)

	// 监听配置变更
	client.AddChangeListener(&CustomChangeListener{})

	fmt.Println("配置监听已启动，等待配置变更...")

	// 保持运行
	select {}
}

// CustomChangeListener 自定义变更监听器
type CustomChangeListener struct{}

func (l *CustomChangeListener) OnChange(event *agollo.ChangeEvent) {
	for key, change := range event.Changes {
		fmt.Printf("配置变更: %s, 旧值: %s, 新值: %s\n",
			key, change.OldValue, change.NewValue)
	}
}

func (l *CustomChangeListener) OnNewestChange(event *agollo.FullChangeEvent) {
	fmt.Printf("配置全量更新: %+v\n", event)
}

