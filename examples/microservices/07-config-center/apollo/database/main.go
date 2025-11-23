package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func loadDatabaseConfig(client *agollo.Client) *DatabaseConfig {
	config := &DatabaseConfig{
		Host:     client.GetStringValue("db.host", "localhost"),
		Port:     client.GetIntValue("db.port", 3306),
		Username: client.GetStringValue("db.username", "root"),
		Password: client.GetStringValue("db.password", ""),
		Database: client.GetStringValue("db.database", "test"),
	}
	return config
}

func main() {
	// 初始化 Apollo 客户端
	appConfig := &config.AppConfig{
		AppID:          "my-app",
		Cluster:        "default",
		NamespaceName:  "application",
		IP:             "http://localhost:8080",
		IsBackupConfig: true,
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return appConfig, nil
	})
	if err != nil {
		log.Fatalf("启动 Apollo 客户端失败: %v", err)
	}

	// 等待配置加载
	time.Sleep(1 * time.Second)

	// 加载数据库配置
	dbConfig := loadDatabaseConfig(client)
	fmt.Printf("数据库配置: %+v\n", dbConfig)

	// 监听配置变更
	client.AddChangeListener(&DatabaseConfigListener{
		onChange: func(key string, oldValue, newValue string) {
			fmt.Printf("数据库配置变更: %s, 旧值: %s, 新值: %s\n",
				key, oldValue, newValue)
			// 这里可以重新连接数据库
		},
	})

	fmt.Println("数据库配置监听已启动...")

	// 保持运行
	select {}
}

type DatabaseConfigListener struct {
	onChange func(key, oldValue, newValue string)
}

func (l *DatabaseConfigListener) OnChange(event *agollo.ChangeEvent) {
	for key, change := range event.Changes {
		if key == "db.host" || key == "db.port" || key == "db.username" ||
			key == "db.password" || key == "db.database" {
			if l.onChange != nil {
				l.onChange(key, change.OldValue, change.NewValue)
			}
		}
	}
}

func (l *DatabaseConfigListener) OnNewestChange(event *agollo.FullChangeEvent) {
	fmt.Printf("数据库配置全量更新\n")
}

