# golang 中，数字操作常用方法

我来介绍一下 Go 语言中常用的数字操作方法。我将创建一个新的文档来详细说明。


## 一、基本数字类型
```go
// 整数类型
var i8 int8    // -128 到 127
var i16 int16  // -32768 到 32767
var i32 int32  // -2^31 到 2^31-1
var i64 int64  // -2^63 到 2^63-1
var ui uint    // 无符号整数

// 浮点类型
var f32 float32  // 32位浮点数
var f64 float64  // 64位浮点数
```

## 二、数字转换

```go
// 1. 类型转换
intNum := 42
float64Num := float64(intNum)    // int 转 float64
uint32Num := uint32(intNum)      // int 转 uint32

// 2. 字符串与数字转换
import "strconv"

// 字符串转数字
intVal, _ := strconv.Atoi("42")           // 字符串转 int
int64Val, _ := strconv.ParseInt("42", 10, 64)  // 字符串转 int64
floatVal, _ := strconv.ParseFloat("3.14", 64)  // 字符串转 float64

// 数字转字符串
strVal := strconv.Itoa(42)                // int 转字符串
str64 := strconv.FormatInt(42, 10)        // int64 转字符串
strFloat := strconv.FormatFloat(3.14, 'f', 2, 64)  // float64 转字符串
```

## 三、数学运算
```go
import "math"

// 1. 基本运算
max := math.Max(1.0, 2.0)    // 取最大值
min := math.Min(1.0, 2.0)    // 取最小值
abs := math.Abs(-3.14)       // 绝对值
pow := math.Pow(2, 3)        // 幂运算
sqrt := math.Sqrt(16)        // 平方根

// 2. 取整操作
ceil := math.Ceil(3.14)      // 向上取整
floor := math.Floor(3.14)    // 向下取整
round := math.Round(3.14)    // 四舍五入

// 3. 三角函数
sin := math.Sin(math.Pi / 2)  // 正弦
cos := math.Cos(0)            // 余弦
tan := math.Tan(math.Pi / 4)  // 正切
```

## 四、随机数生成
```go
import (
    "math/rand"
    "time"
)

// 1. 基本随机数
rand.Seed(time.Now().UnixNano())  // 设置随机种子
n := rand.Int()                   // 生成随机整数
f := rand.Float64()               // 生成随机浮点数

// 2. 范围随机数
randInt := rand.Intn(100)         // 生成 [0,100) 的随机整数
randFloat := rand.Float64() * 100  // 生成 [0,100) 的随机浮点数
```

## 五、位运算
```go
a, b := 60, 13

// 常用位运算
and := a & b     // 按位与
or := a | b      // 按位或
xor := a ^ b     // 按位异或
not := ^a        // 按位取反
left := a << 2   // 左移
right := a >> 2  // 右移
```

## 六、数字判断
```go
import "math"

// 1. 特殊值判断
isNaN := math.IsNaN(f)       // 是否是 NaN
isInf := math.IsInf(f, 0)    // 是否是无穷大

// 2. 符号判断
isPositive := math.Signbit(f)  // 判断符号位
```

## 七、实用技巧

1. **精度处理**：
```go
// 保留两位小数
rounded := math.Round(3.1415926*100) / 100

// 金额计算（使用 decimal 包）
import "github.com/shopspring/decimal"

d1 := decimal.NewFromFloat(2.1)
d2 := decimal.NewFromFloat(1.2)
sum := d1.Add(d2)  // 精确计算
```

2. **溢出检查**：
```go
// 检查加法溢出
if a > math.MaxInt64 - b {
    // 处理溢出情况
}
```

## 注意事项

1. 浮点数计算可能存在精度问题，需要特别注意
2. 金融计算建议使用 decimal 包
3. 进行类型转换时要注意可能的数据溢出
4. 随机数生成前必须设置随机种子
5. 位运算要注意符号位的影响
```

这个文档涵盖了 Go 语言中最常用的数字操作方法。建议根据实际场景选择合适的方法，特别是在需要高精度计算的场景下。