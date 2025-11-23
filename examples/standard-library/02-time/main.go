// Package main 展示 time 包的各种用法
package main

import (
	"fmt"
	"time"
)

// 程序入口，演示 time 包的各种功能
func main() {
	fmt.Println("=== time 包示例 ===")

	// 基本时间操作
	basicTimeOperations()

	// 时间格式化和解析
	timeFormattingAndParsing()

	// 时间计算
	timeCalculations()

	// 定时器和计时器
	timersAndTickers()

	// 时区操作
	timezoneOperations()

	// 实际应用示例
	practicalExample()
}

// basicTimeOperations 基本时间操作
func basicTimeOperations() {
	fmt.Println("\n--- 基本时间操作 ---")

	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间: %v\n", now)
	fmt.Printf("时间戳(秒): %d\n", now.Unix())
	fmt.Printf("时间戳(纳秒): %d\n", now.UnixNano())

	// 2. 创建指定时间
	// 年, 月, 日, 时, 分, 秒, 纳秒, 位置
	specificTime := time.Date(2024, 1, 15, 14, 30, 0, 0, time.Local)
	fmt.Printf("指定时间: %v\n", specificTime)

	// 3. 使用时间戳创建时间
	timestamp := time.Unix(1642248600, 0)
	fmt.Printf("时间戳转时间: %v\n", timestamp)

	// 4. 时间的组成部分
	fmt.Printf("\n时间的组成部分:")
	fmt.Printf("年份: %d\n", now.Year())
	fmt.Printf("月份: %d\n", now.Month())
	fmt.Printf("日期: %d\n", now.Day())
	fmt.Printf("小时: %d\n", now.Hour())
	fmt.Printf("分钟: %d\n", now.Minute())
	fmt.Printf("秒: %d\n", now.Second())
	fmt.Printf("纳秒: %d\n", now.Nanosecond())
	fmt.Printf("星期几: %s\n", now.Weekday())
	fmt.Printf("一年中的第几天: %d\n", now.YearDay())

	// 5. 时间比较
	fmt.Printf("\n时间比较:")
	time1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	time2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.Local)

	fmt.Printf("time1: %v\n", time1)
	fmt.Printf("time2: %v\n", time2)
	fmt.Printf("time1 == time2: %t\n", time1.Equal(time2))
	fmt.Printf("time1 before time2: %t\n", time1.Before(time2))
	fmt.Printf("time1 after time2: %t\n", time1.After(time2))
}

// timeFormattingAndParsing 时间格式化和解析
func timeFormattingAndParsing() {
	fmt.Println("\n--- 时间格式化和解析 ---")

	now := time.Now()

	// 1. 预定义格式
	fmt.Printf("预定义格式:\n")
	fmt.Printf("RFC1123: %s\n", now.Format(time.RFC1123))
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("ANSIC: %s\n", now.Format(time.ANSIC))
	fmt.Printf("UnixDate: %s\n", now.Format(time.UnixDate))
	fmt.Printf("RubyDate: %s\n", now.Format(time.RubyDate))

	// 2. 自定义格式（重要：格式字符串必须是参考时间 2006-01-02 15:04:05 的格式）
	fmt.Printf("\n自定义格式:\n")
	fmt.Printf("2006-01-02: %s\n", now.Format("2006-01-02"))
	fmt.Printf("2006/01/02 15:04:05: %s\n", now.Format("2006/01/02 15:04:05"))
	fmt.Printf("2006年01月02日 15时04分05秒: %s\n", now.Format("2006年01月02日 15时04分05秒"))
	fmt.Printf("2006-01-02T15:04:05Z07:00: %s\n", now.Format("2006-01-02T15:04:05Z07:00"))

	// 3. 解析时间字符串
	fmt.Printf("\n解析时间字符串:\n")

	// 解析自定义格式
	timeStr1 := "2024-01-15 14:30:00"
	parsedTime1, err := time.Parse("2006-01-02 15:04:05", timeStr1)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果1: %v\n", parsedTime1)
	}

	// 解析RFC3339格式
	timeStr2 := "2024-01-15T14:30:00Z08:00"
	parsedTime2, err := time.Parse(time.RFC3339, timeStr2)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果2: %v\n", parsedTime2)
	}

	// 解析中文格式
	timeStr3 := "2024年01月15日 14时30分00秒"
	parsedTime3, err := time.Parse("2006年01月02日 15时04分05秒", timeStr3)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果3: %v\n", parsedTime3)
	}
}

// timeCalculations 时间计算
func timeCalculations() {
	fmt.Println("\n--- 时间计算 ---")

	now := time.Now()

	// 1. 时间加减
	fmt.Printf("时间加减:\n")

	// 加时间
	fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))

	later1 := now.Add(24 * time.Hour)      // 加24小时
	later2 := now.Add(30 * time.Minute)    // 加30分钟
	later3 := now.Add(5 * time.Second)     // 加5秒
	later4 := now.Add(500 * time.Millisecond) // 加500毫秒

	fmt.Printf("加24小时: %s\n", later1.Format("2006-01-02 15:04:05"))
	fmt.Printf("加30分钟: %s\n", later2.Format("2006-01-02 15:04:05"))
	fmt.Printf("加5秒: %s\n", later3.Format("2006-01-02 15:04:05"))
	fmt.Printf("加500毫秒: %s\n", later4.Format("2006-01-02 15:04:05.000"))

	// 减时间
	earlier := now.Add(-24 * time.Hour)
	fmt.Printf("减24小时: %s\n", earlier.Format("2006-01-02 15:04:05"))

	// 2. 添加年月日
	fmt.Printf("\n添加年月日:\n")
	nextYear := now.AddDate(1, 0, 0)        // 加1年
	nextMonth := now.AddDate(0, 1, 0)       // 加1月
	nextDay := now.AddDate(0, 0, 1)         // 加1天
	specificDate := now.AddDate(1, 2, 15)   // 加1年2月15天

	fmt.Printf("加1年: %s\n", nextYear.Format("2006-01-02"))
	fmt.Printf("加1月: %s\n", nextMonth.Format("2006-01-02"))
	fmt.Printf("加1天: %s\n", nextDay.Format("2006-01-02"))
	fmt.Printf("加1年2月15天: %s\n", specificDate.Format("2006-01-02"))

	// 3. 时间差计算
	fmt.Printf("\n时间差计算:\n")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.Local)

	duration := end.Sub(start)
	fmt.Printf("从 %s 到 %s 的时间差: %v\n",
		start.Format("2006-01-02"), end.Format("2006-01-02"), duration)

	fmt.Printf("小时: %.1f\n", duration.Hours())
	fmt.Printf("分钟: %.1f\n", duration.Minutes())
	fmt.Printf("秒数: %.1f\n", duration.Seconds())

	// 4. Truncate 和 Round
	fmt.Printf("\nTruncate 和 Round:\n")
	testTime := time.Date(2024, 1, 15, 14, 30, 45, 123456789, time.Local)

	truncated := testTime.Truncate(time.Minute)  // 截断到分钟
	rounded := testTime.Round(time.Second)       // 四舍五入到秒

	fmt.Printf("原始时间: %s\n", testTime.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("截断到分钟: %s\n", truncated.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("四舍五入到秒: %s\n", rounded.Format("2006-01-02 15:04:05.000000000"))
}

// timersAndTickers 定时器和计时器
func timersAndTickers() {
	fmt.Println("\n--- 定时器和计时器 ---")

	// 1. 简单计时器
	fmt.Printf("计时器示例:\n")
	start := time.Now()

	// 模拟一些工作
	time.Sleep(100 * time.Millisecond)

	elapsed := time.Since(start)
	fmt.Printf("执行时间: %v\n", elapsed)

	// 2. Timer - 单次触发
	fmt.Printf("\nTimer 示例:\n")
	fmt.Printf("创建Timer，2秒后触发...\n")

	timer := time.NewTimer(2 * time.Second)

	// 使用 select 处理定时器
	select {
	case <-timer.C:
		fmt.Printf("Timer 触发了！\n")
	case <-time.After(1 * time.Second):
		fmt.Printf("1秒超时了\n")
		timer.Stop() // 停止timer
	}

	// 3. Ticker - 周期性触发
	fmt.Printf("\nTicker 示例 (3次):\n")
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("Tick: %s\n", t.Format("15:04:05.000"))
			count++
			if count >= 3 {
				fmt.Printf("Ticker 示例结束\n")
				return
			}
		case <-time.After(2 * time.Second):
			fmt.Printf("Ticker 示例超时结束\n")
			return
		}
	}
}

// timezoneOperations 时区操作
func timezoneOperations() {
	fmt.Println("\n--- 时区操作 ---")

	// 1. 获取时区信息
	fmt.Printf("时区信息:\n")
	local := time.Local
	fmt.Printf("本地时区: %v\n", local)
	fmt.Printf("UTC时区: %v\n", time.UTC)

	// 2. 加载特定时区
	fmt.Printf("\n加载特定时区:")

	// 纽约时区
	nyTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("加载纽约时区失败: %v\n", err)
	} else {
		fmt.Printf("纽约时区: %v\n", nyTime)
	}

	// 伦敦时区
	londonTime, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Printf("加载伦敦时区失败: %v\n", err)
	} else {
		fmt.Printf("伦敦时区: %v\n", londonTime)
	}

	// 东京时区
	tokyoTime, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Printf("加载东京时区失败: %v\n", err)
	} else {
		fmt.Printf("东京时区: %v\n", tokyoTime)
	}

	// 3. 同一时间在不同时区的显示
	fmt.Printf("\n同一时间在不同时区:\n")
	utcTime := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

	fmt.Printf("UTC: %s\n", utcTime.Format("2006-01-02 15:04:05 MST"))
	if nyTime != nil {
		nyLocalTime := utcTime.In(nyTime)
		fmt.Printf("纽约: %s\n", nyLocalTime.Format("2006-01-02 15:04:05 MST"))
	}
	if londonTime != nil {
		londonLocalTime := utcTime.In(londonTime)
		fmt.Printf("伦敦: %s\n", londonLocalTime.Format("2006-01-02 15:04:05 MST"))
	}
	if tokyoTime != nil {
		tokyoLocalTime := utcTime.In(tokyoTime)
		fmt.Printf("东京: %s\n", tokyoLocalTime.Format("2006-01-02 15:04:05 MST"))
	}
}

// practicalExample 实际应用示例
func practicalExample() {
	fmt.Println("\n--- 实际应用示例 ---")

	// 示例1: 计算程序运行时间
	fmt.Println("1. 程序运行时间计算:")
	calculateExecutionTime()

	// 示例2: 日期工具函数
	fmt.Println("\n2. 日期工具函数:")
	demonstrateDateUtils()

	// 示例3: 缓存过期时间
	fmt.Println("\n3. 缓存过期时间:")
	demonstrateCacheExpiration()

	// 示例4: 任务调度
	fmt.Println("\n4. 简单任务调度:")
	demonstrateTaskScheduling()
}

// calculateExecutionTime 计算执行时间
func calculateExecutionTime() {
	start := time.Now()

	// 模拟计算密集任务
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}

	elapsed := time.Since(start)
	fmt.Printf("计算结果: %d\n", sum)
	fmt.Printf("执行时间: %v\n", elapsed)
	fmt.Printf("执行时间(毫秒): %d\n", elapsed.Milliseconds())
}

// demonstrateDateUtils 日期工具函数
func demonstrateDateUtils() {
	now := time.Now()

	// 判断是否是今天
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	fmt.Printf("今天: %s\n", today.Format("2006-01-02"))

	// 获取本周第一天
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7 // 周日转换为7
	}
	thisWeek := now.AddDate(0, 0, -(weekday - 1))
	fmt.Printf("本周第一天: %s\n", thisWeek.Format("2006-01-02 Monday"))

	// 获取本月第一天
	thisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	fmt.Printf("本月第一天: %s\n", thisMonth.Format("2006-01-02"))

	// 计算年龄
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, now.Location())
	age := calculateAge(birthDate, now)
	fmt.Printf("出生日期: %s, 当前年龄: %d岁\n", birthDate.Format("2006-01-02"), age)

	// 判断是否是工作日
	fmt.Printf("今天是否是工作日: %t\n", isWeekday(now))
}

// calculateAge 计算年龄
func calculateAge(birthDate, now time.Time) int {
	age := now.Year() - birthDate.Year()

	// 如果还没过生日，年龄减1
	if now.Month() < birthDate.Month() ||
	   (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age
}

// isWeekday 判断是否是工作日
func isWeekday(t time.Time) bool {
	weekday := t.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}

// demonstrateCacheExpiration 缓存过期时间
func demonstrateCacheExpiration() {
	// 模拟缓存项
	type CacheItem struct {
		Value     interface{}
		ExpiresAt time.Time
	}

	cache := make(map[string]*CacheItem)

	// 添加缓存项，5分钟后过期
	cacheKey := "user:123"
	cache[cacheKey] = &CacheItem{
		Value:     "张三",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	fmt.Printf("缓存添加时间: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("缓存过期时间: %s\n", cache[cacheKey].ExpiresAt.Format("15:04:05"))

	// 模拟检查缓存
	if item, exists := cache[cacheKey]; exists {
		if time.Now().Before(item.ExpiresAt) {
			fmt.Printf("缓存有效: %v\n", item.Value)
		} else {
			fmt.Printf("缓存已过期\n")
		}
	}
}

// demonstrateTaskScheduling 任务调度
func demonstrateTaskScheduling() {
	fmt.Printf("启动任务调度器...\n")

	// 模拟定时任务
	taskCount := 0
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeout := time.After(10 * time.Second) // 10秒后停止

	for {
		select {
		case <-ticker.C:
			taskCount++
			fmt.Printf("执行任务 #%d - %s\n", taskCount, time.Now().Format("15:04:05"))

			// 模拟任务执行
			time.Sleep(500 * time.Millisecond)

		case <-timeout:
			fmt.Printf("任务调度器停止，共执行 %d 个任务\n", taskCount)
			return
		}
	}
}