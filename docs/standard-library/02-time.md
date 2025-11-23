---
title: 时间处理 (time)
difficulty: intermediate
duration: "3-4小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["time", "时间", "日期", "定时器", "时区"]
---

# 时间处理 (time)

`time` 包提供了时间相关的功能，包括时间获取、格式化、计算、定时器和时区操作。

## 📋 学习目标

- [ ] 掌握时间的基本操作
- [ ] 理解时间的格式化和解析
- [ ] 学会时间计算和比较
- [ ] 掌握定时器和计时器的使用
- [ ] 理解时区操作
- [ ] 了解时间在实际应用中的用法

## 🎯 基本时间操作

### 获取当前时间

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间: %v\n", now)
	
	// 获取时间戳
	fmt.Printf("时间戳(秒): %d\n", now.Unix())
	fmt.Printf("时间戳(纳秒): %d\n", now.UnixNano())
}
```

### 创建指定时间

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用 Date 创建指定时间
	// 参数：年, 月, 日, 时, 分, 秒, 纳秒, 时区
	specificTime := time.Date(2024, 1, 15, 14, 30, 0, 0, time.Local)
	fmt.Printf("指定时间: %v\n", specificTime)
	
	// 使用时间戳创建时间
	timestamp := time.Unix(1642248600, 0)
	fmt.Printf("时间戳转时间: %v\n", timestamp)
}
```

### 时间的组成部分

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	fmt.Printf("年份: %d\n", now.Year())
	fmt.Printf("月份: %d (%s)\n", now.Month(), now.Month())
	fmt.Printf("日期: %d\n", now.Day())
	fmt.Printf("小时: %d\n", now.Hour())
	fmt.Printf("分钟: %d\n", now.Minute())
	fmt.Printf("秒: %d\n", now.Second())
	fmt.Printf("纳秒: %d\n", now.Nanosecond())
	fmt.Printf("星期几: %s\n", now.Weekday())
	fmt.Printf("一年中的第几天: %d\n", now.YearDay())
}
```

### 时间比较

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	time2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.Local)
	
	fmt.Printf("time1 == time2: %t\n", time1.Equal(time2))
	fmt.Printf("time1 before time2: %t\n", time1.Before(time2))
	fmt.Printf("time1 after time2: %t\n", time1.After(time2))
}
```

## 📅 时间格式化和解析

### 预定义格式

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	fmt.Printf("RFC1123: %s\n", now.Format(time.RFC1123))
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("ANSIC: %s\n", now.Format(time.ANSIC))
	fmt.Printf("UnixDate: %s\n", now.Format(time.UnixDate))
	fmt.Printf("RubyDate: %s\n", now.Format(time.RubyDate))
}
```

### 自定义格式

**重要**：Go 使用参考时间 `2006-01-02 15:04:05` 来定义格式。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	// 自定义格式
	fmt.Printf("2006-01-02: %s\n", now.Format("2006-01-02"))
	fmt.Printf("2006/01/02 15:04:05: %s\n", now.Format("2006/01/02 15:04:05"))
	fmt.Printf("2006年01月02日 15时04分05秒: %s\n", 
		now.Format("2006年01月02日 15时04分05秒"))
	fmt.Printf("2006-01-02T15:04:05Z07:00: %s\n", 
		now.Format("2006-01-02T15:04:05Z07:00"))
}
```

### 解析时间字符串

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 解析自定义格式
	timeStr1 := "2024-01-15 14:30:00"
	parsedTime1, err := time.Parse("2006-01-02 15:04:05", timeStr1)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果: %v\n", parsedTime1)
	}
	
	// 解析 RFC3339 格式
	timeStr2 := "2024-01-15T14:30:00+08:00"
	parsedTime2, err := time.Parse(time.RFC3339, timeStr2)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果: %v\n", parsedTime2)
	}
```

## ⏰ 时间计算

### 时间加减

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	// 加时间
	later1 := now.Add(24 * time.Hour)      // 加24小时
	later2 := now.Add(30 * time.Minute)    // 加30分钟
	later3 := now.Add(5 * time.Second)     // 加5秒
	later4 := now.Add(500 * time.Millisecond) // 加500毫秒
	
	fmt.Printf("加24小时: %s\n", later1.Format("2006-01-02 15:04:05"))
	fmt.Printf("加30分钟: %s\n", later2.Format("2006-01-02 15:04:05"))
	
	// 减时间
	earlier := now.Add(-24 * time.Hour)
	fmt.Printf("减24小时: %s\n", earlier.Format("2006-01-02 15:04:05"))
}
```

### 添加年月日

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	nextYear := now.AddDate(1, 0, 0)        // 加1年
	nextMonth := now.AddDate(0, 1, 0)       // 加1月
	nextDay := now.AddDate(0, 0, 1)         // 加1天
	specificDate := now.AddDate(1, 2, 15)   // 加1年2月15天
	
	fmt.Printf("加1年: %s\n", nextYear.Format("2006-01-02"))
	fmt.Printf("加1月: %s\n", nextMonth.Format("2006-01-02"))
	fmt.Printf("加1天: %s\n", nextDay.Format("2006-01-02"))
	fmt.Printf("加1年2月15天: %s\n", specificDate.Format("2006-01-02"))
}
```

### 时间差计算

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local)
	
	duration := end.Sub(start)
	fmt.Printf("时间差: %v\n", duration)
	fmt.Printf("小时: %.1f\n", duration.Hours())
	fmt.Printf("分钟: %.1f\n", duration.Minutes())
	fmt.Printf("秒数: %.1f\n", duration.Seconds())
}
```

### Truncate 和 Round

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	testTime := time.Date(2024, 1, 15, 14, 30, 45, 123456789, time.Local)
	
	truncated := testTime.Truncate(time.Minute)  // 截断到分钟
	rounded := testTime.Round(time.Second)       // 四舍五入到秒
	
	fmt.Printf("原始时间: %s\n", testTime.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("截断到分钟: %s\n", truncated.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("四舍五入到秒: %s\n", rounded.Format("2006-01-02 15:04:05.000000000"))
}
```

## ⏲️ 定时器和计时器

### 简单计时

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	// 模拟一些工作
	time.Sleep(100 * time.Millisecond)
	
	elapsed := time.Since(start)
	fmt.Printf("执行时间: %v\n", elapsed)
	fmt.Printf("执行时间(毫秒): %d\n", elapsed.Milliseconds())
}
```

### Timer - 单次触发

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("创建Timer，2秒后触发...")
	
	timer := time.NewTimer(2 * time.Second)
	
	// 使用 select 处理定时器
	select {
	case <-timer.C:
		fmt.Println("Timer 触发了！")
	case <-time.After(1 * time.Second):
		fmt.Println("1秒超时了")
		timer.Stop() // 停止timer
	}
}
```

### Ticker - 周期性触发

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	count := 0
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("Tick: %s\n", t.Format("15:04:05.000"))
			count++
			if count >= 3 {
				fmt.Println("Ticker 示例结束")
				return
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Ticker 示例超时结束")
			return
		}
	}
}
```

## 🌍 时区操作

### 获取时区信息

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("本地时区: %v\n", time.Local)
	fmt.Printf("UTC时区: %v\n", time.UTC)
}
```

### 加载特定时区

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 加载时区
	nyTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("加载失败: %v\n", err)
	} else {
		fmt.Printf("纽约时区: %v\n", nyTime)
	}
	
	londonTime, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Printf("加载失败: %v\n", err)
	} else {
		fmt.Printf("伦敦时区: %v\n", londonTime)
	}
}
```

### 时区转换

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	utcTime := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	
	fmt.Printf("UTC: %s\n", utcTime.Format("2006-01-02 15:04:05 MST"))
	
	// 转换为其他时区
	if nyTime, err := time.LoadLocation("America/New_York"); err == nil {
		nyLocalTime := utcTime.In(nyTime)
		fmt.Printf("纽约: %s\n", nyLocalTime.Format("2006-01-02 15:04:05 MST"))
	}
}
```

## 🏃‍♂️ 实践应用

### 计算程序运行时间

```go
func calculateExecutionTime() {
	start := time.Now()
	
	// 执行一些操作
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	
	elapsed := time.Since(start)
	fmt.Printf("执行时间: %v\n", elapsed)
}
```

### 日期工具函数

```go
func getToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func getThisWeek() time.Time {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return now.AddDate(0, 0, -(weekday - 1))
}

func isWeekday(t time.Time) bool {
	weekday := t.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}
```

### 缓存过期时间

```go
type CacheItem struct {
	Value     interface{}
	ExpiresAt time.Time
}

func isExpired(item *CacheItem) bool {
	return time.Now().After(item.ExpiresAt)
}

// 使用
item := &CacheItem{
	Value:     "data",
	ExpiresAt: time.Now().Add(5 * time.Minute),
}
```

## 🤔 思考题

1. Go 的时间格式化为什么使用 `2006-01-02 15:04:05` 作为参考时间？
2. Timer 和 Ticker 有什么区别？
3. 如何计算两个日期之间的天数？
4. 时区转换时需要注意什么？

## 📚 扩展阅读

- [time 包官方文档](https://pkg.go.dev/time)
- [时间格式化详解](https://golang.org/pkg/time/#Time.Format)

## ⏭️ 下一章节

[命令行参数 (flag)](./03-flag.md) → 学习 Go 语言的命令行参数解析

---

**💡 提示**: Go 的时间格式化使用固定的参考时间 `2006-01-02 15:04:05`，这是 Go 语言的诞生时间！

