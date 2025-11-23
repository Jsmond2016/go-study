# time 包示例

time 包提供了时间和日期的相关功能，包括时间测量、格式化、解析、计算等。本示例全面展示了 time 包的各种用法。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 基本时间操作

```go
// 获取当前时间
now := time.Now()

// 创建指定时间
t := time.Date(2024, 1, 15, 14, 30, 0, 0, time.Local)

// 从时间戳创建
t := time.Unix(1642248600, 0)

// 获取时间组成部分
year := t.Year()
month := t.Month()
day := t.Day()
hour := t.Hour()
minute := t.Minute()
second := t.Second()
weekday := t.Weekday()
```

### 2. 时间比较

```go
t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
t2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.Local)

t1.Equal(t2)   // 相等
t1.Before(t2)  // t1 在 t2 之前
t1.After(t2)   // t1 在 t2 之后
```

### 3. 时间格式化

```go
now := time.Now()

// 预定义格式
now.Format(time.RFC1123)    // "Mon, 02 Jan 2006 15:04:05 MST"
now.Format(time.RFC3339)    // "2006-01-02T15:04:05Z07:00"
now.Format(time.ANSIC)      // "Mon Jan 2 15:04:05 2006"

// 自定义格式（重要：使用参考时间 2006-01-02 15:04:05）
now.Format("2006-01-02")                    // "2024-01-15"
now.Format("2006/01/02 15:04:05")           // "2024/01/15 14:30:00"
now.Format("2006年01月02日 15时04分05秒")    // "2024年01月15日 14时30分00秒"
```

### 4. 时间解析

```go
// 解析自定义格式
t, err := time.Parse("2006-01-02 15:04:05", "2024-01-15 14:30:00")

// 解析标准格式
t, err := time.Parse(time.RFC3339, "2024-01-15T14:30:00Z08:00")

// 解析中文格式
t, err := time.Parse("2006年01月02日", "2024年01月15日")
```

### 5. 时间计算

```go
now := time.Now()

// 时间加减
future := now.Add(24 * time.Hour)      // 加24小时
past := now.Add(-30 * time.Minute)     // 减30分钟

// 年月日加减
future := now.AddDate(1, 2, 15)        // 加1年2月15天

// 时间差
duration := future.Sub(now)           // 时间差
hours := duration.Hours()             // 小时数
minutes := duration.Minutes()         // 分钟数
seconds := duration.Seconds()         // 秒数
```

### 6. Timer 和 Ticker

```go
// Timer - 单次触发
timer := time.NewTimer(2 * time.Second)
select {
case <-timer.C:
    fmt.Println("Timer 触发")
}

// Ticker - 周期性触发
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for {
    select {
    case t := <-ticker.C:
        fmt.Println("Tick:", t)
    }
}
```

### 7. 时区操作

```go
// 加载时区
loc, err := time.LoadLocation("America/New_York")

// 转换时区
utcTime := time.Now().UTC()
nyTime := utcTime.In(loc)

// 固定时区
local := time.Now()
utc := local.UTC()
```

## 重要概念

### 1. 参考时间

Go 的时间格式化使用特定的参考时间：`Mon Jan 2 15:04:05 MST 2006`
- `2006` - 年
- `01` - 月
- `02` - 日
- `15` - 小时
- `04` - 分钟
- `05` - 秒

### 2. 时区

- **UTC** - 协调世界时
- **Local** - 本地时区
- **IANA时区** - 如 "Asia/Shanghai"、"America/New_York"

### 3. 时间戳

- `Unix()` - 秒级时间戳
- `UnixNano()` - 纳秒级时间戳
- `UnixMilli()` - 毫秒级时间戳（Go 1.17+）

### 4. Duration 类型

```go
type Duration int64

const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

## 实际应用场景

### 1. 性能测量

```go
start := time.Now()
// 执行代码
elapsed := time.Since(start)
fmt.Printf("执行时间: %v\n", elapsed)
```

### 2. 缓存过期

```go
type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

if time.Now().Before(item.ExpiresAt) {
    // 缓存有效
} else {
    // 缓存过期
}
```

### 3. 任务调度

```go
ticker := time.NewTicker(time.Hour)
for {
    select {
    case <-ticker.C:
        // 执行每小时任务
    }
}
```

### 4. 日期计算

```go
// 获取第一天和最后一天
firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

// 计算年龄
age := now.Year() - birth.Year()
if now.Before(birth.AddDate(age, 0, 0)) {
    age--
}
```

## 最佳实践

### 1. 时间格式化

- 使用常量时间格式，避免每次创建
- 统一时间格式标准
- 考虑国际化需求

### 2. 时区处理

- 数据库存储使用 UTC
- 显示时转换为用户时区
- 明确时区信息

### 3. 性能考虑

- 避免频繁的时间解析
- 使用缓存预格式化时间
- 选择合适的时间精度

### 4. 错误处理

- 总是检查时间解析错误
- 处理时区加载失败
- 处理无效时间值

## 常见陷阱

### 1. 格式化字符串错误

```go
// 错误 - 使用了传统格式
fmt.Println(now.Format("YYYY-MM-DD"))

// 正确 - 使用Go参考时间
fmt.Println(now.Format("2006-01-02"))
```

### 2. 时区混淆

```go
// 注意区分时间和时区
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // 明确UTC
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local) // 本地时间
```

### 3. 时间精度问题

```go
// 注意纳秒精度可能导致的问题
if t1 == t2 { ... } // 可能因为纳秒不同而不相等

// 使用 Equal 方法比较
if t1.Equal(t2) { ... } // 正确的比较方式
```

### 4. 夏令时处理

```go
// 夏令时可能导致时间不存在或重复
// 使用时间.AddDate 而不是手动计算
```

## 实用工具函数

### 1. 获取时间范围

```go
func GetDayRange(t time.Time) (time.Time, time.Time) {
    start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
    end := start.Add(24 * time.Hour)
    return start, end
}
```

### 2. 时间格式化工具

```go
var (
    DateFormat     = "2006-01-02"
    DateTimeFormat = "2006-01-02 15:04:05"
    TimeFormat     = "15:04:05"
)

func FormatDate(t time.Time) string {
    return t.Format(DateFormat)
}
```

### 3. 时区转换

```go
func ToUTC(t time.Time) time.Time {
    return t.UTC()
}

func ToLocation(t time.Time, loc *time.Location) time.Time {
    return t.In(loc)
}
```

## 练习

1. 创建一个时间日期格式化工具
2. 实现一个简单的日程管理系统
3. 创建一个倒计时器应用
4. 实现一个工作日计算器
5. 创建一个时区转换工具