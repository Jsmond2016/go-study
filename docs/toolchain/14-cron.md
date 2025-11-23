---
title: 任务调度 Cron
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["基础语法", "并发编程"]
tags: ["Cron", "定时任务", "调度", "robfig/cron"]
---

# 任务调度 Cron

Cron 用于定时执行任务，是后台系统中常用的功能。

## 📋 学习目标

- [ ] 理解定时任务的概念
- [ ] 掌握 cron 表达式的使用
- [ ] 学会使用 robfig/cron
- [ ] 理解任务管理
- [ ] 掌握最佳实践

## 🎯 Cron 简介

### 为什么需要定时任务

- **数据同步**: 定期同步数据
- **清理任务**: 定期清理过期数据
- **报表生成**: 定期生成报表
- **监控检查**: 定期检查系统状态
- **通知发送**: 定期发送通知

### 安装 robfig/cron

```bash
go get github.com/robfig/cron/v3
```

## 🚀 快速开始

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	
	// 添加任务：每秒执行
	c.AddFunc("@every 1s", func() {
		fmt.Println("每秒执行")
	})
	
	// 添加任务：每分钟执行
	c.AddFunc("@every 1m", func() {
		fmt.Println("每分钟执行")
	})
	
	// 启动
	c.Start()
	defer c.Stop()
	
	// 保持程序运行
	select {}
}
```

## 📅 Cron 表达式

### 标准格式

```
秒 分 时 日 月 星期
*  *  *  *  *  *
```

### 常用表达式

```go
// 每分钟执行
"0 * * * * *"

// 每小时执行
"0 0 * * * *"

// 每天执行（凌晨）
"0 0 0 * * *"

// 每周执行（周一凌晨）
"0 0 0 * * 1"

// 每月执行（1号凌晨）
"0 0 0 1 * *"

// 每5分钟执行
"0 */5 * * * *"

// 工作日上午9点执行
"0 0 9 * * 1-5"
```

### 预定义表达式

```go
// 每秒
"@every 1s"

// 每分钟
"@every 1m"

// 每小时
"@every 1h"

// 每天
"@daily"

// 每周
"@weekly"

// 每月
"@monthly"

// 每年
"@yearly"
```

## 🔧 高级用法

### 带参数的任务

```go
type Task struct {
	Name string
	Func func()
}

func main() {
	c := cron.New()
	
	task := Task{
		Name: "清理任务",
		Func: func() {
			fmt.Println("执行清理任务")
		},
	}
	
	c.AddFunc("@every 1m", task.Func)
	c.Start()
	
	select {}
}
```

### 任务管理

```go
func main() {
	c := cron.New()
	
	// 添加任务并获取 ID
	id, err := c.AddFunc("@every 1m", func() {
		fmt.Println("任务执行")
	})
	if err != nil {
		panic(err)
	}
	
	// 启动
	c.Start()
	
	// 移除任务
	c.Remove(id)
	
	// 停止所有任务
	c.Stop()
}
```

### 时区设置

```go
import (
	"time"
	"github.com/robfig/cron/v3"
)

func main() {
	// 使用中国时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(loc))
	
	c.AddFunc("0 0 9 * * *", func() {
		fmt.Println("每天9点执行（北京时间）")
	})
	
	c.Start()
	select {}
}
```

## 🏃‍♂️ 实践应用

### 完整的任务调度服务

```go
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(cron.WithSeconds()),
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("任务调度器已启动")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Println("任务调度器已停止")
}

func (s *Scheduler) AddTask(spec string, cmd func()) (cron.EntryID, error) {
	return s.cron.AddFunc(spec, cmd)
}

func (s *Scheduler) RemoveTask(id cron.EntryID) {
	s.cron.Remove(id)
}

func main() {
	scheduler := NewScheduler()
	
	// 添加清理任务：每天凌晨2点执行
	scheduler.AddTask("0 0 2 * * *", func() {
		fmt.Println("执行清理任务:", time.Now())
		// 清理逻辑
	})
	
	// 添加数据同步任务：每5分钟执行
	scheduler.AddTask("0 */5 * * * *", func() {
		fmt.Println("执行数据同步:", time.Now())
		// 同步逻辑
	})
	
	// 添加健康检查任务：每分钟执行
	scheduler.AddTask("@every 1m", func() {
		fmt.Println("执行健康检查:", time.Now())
		// 检查逻辑
	})
	
	scheduler.Start()
	defer scheduler.Stop()
	
	// 保持程序运行
	select {}
}
```

### 在 Web 应用中使用

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var scheduler *cron.Cron

func init() {
	scheduler = cron.New()
	
	// 添加定时任务
	scheduler.AddFunc("@every 1h", cleanupTask)
	scheduler.AddFunc("@daily", reportTask)
	
	scheduler.Start()
}

func cleanupTask() {
	// 清理逻辑
}

func reportTask() {
	// 报表生成逻辑
}

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	
	r.Run(":8080")
}
```

## ⚠️ 注意事项

### 1. 任务执行时间

```go
// ✅ 避免长时间运行的任务
// ✅ 使用 goroutine 执行耗时任务
c.AddFunc("@every 1m", func() {
	go longRunningTask()
})
```

### 2. 错误处理

```go
// ✅ 处理任务中的错误
c.AddFunc("@every 1m", func() {
	if err := doTask(); err != nil {
		log.Printf("任务执行失败: %v", err)
	}
})
```

### 3. 资源清理

```go
// ✅ 程序退出时停止调度器
defer scheduler.Stop()
```

## 📚 扩展阅读

- [robfig/cron 文档](https://github.com/robfig/cron)
- [Cron 表达式](https://en.wikipedia.org/wiki/Cron)

## ⏭️ 下一章节

[WebSocket](./15-websocket.md) → 学习实时通信

---

**💡 提示**: 定时任务是后台系统的重要组成部分，合理使用可以提高系统自动化程度！

