---
title: 数学运算 (math)
difficulty: beginner
duration: "2-3小时"
prerequisites: ["变量与常量", "数据类型"]
tags: ["math", "数学", "运算", "三角函数", "随机数"]
---

# 数学运算 (math)

`math` 包提供了基本的数学常数和数学函数，`math/rand` 包提供了随机数生成功能。

## 📋 学习目标

- [ ] 掌握基本数学运算函数
- [ ] 理解三角函数的使用
- [ ] 学会使用随机数生成
- [ ] 了解数学常数的使用
- [ ] 掌握数值转换和取整操作

## 🔢 基本数学运算

### 最大值和最小值

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 取最大值
	max := math.Max(3.14, 2.71)
	fmt.Printf("最大值: %f\n", max) // 3.140000

	// 取最小值
	min := math.Min(3.14, 2.71)
	fmt.Printf("最小值: %f\n", min) // 2.710000

	// 多个值比较
	max3 := math.Max(math.Max(1.0, 2.0), 3.0)
	fmt.Printf("三个数最大值: %f\n", max3) // 3.000000
}
```

### 绝对值和符号

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 绝对值
	abs := math.Abs(-3.14)
	fmt.Printf("绝对值: %f\n", abs) // 3.140000

	// 符号函数（返回 -1, 0, 或 1）
	sign := math.Copysign(1, -5)
	fmt.Printf("符号: %f\n", sign) // -1.000000

	// 检查符号位
	isNegative := math.Signbit(-3.14)
	fmt.Printf("是否为负: %t\n", isNegative) // true
}
```

### 幂运算和开方

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 幂运算
	pow := math.Pow(2, 3)
	fmt.Printf("2^3 = %f\n", pow) // 8.000000

	// 平方根
	sqrt := math.Sqrt(16)
	fmt.Printf("√16 = %f\n", sqrt) // 4.000000

	// 立方根
	cbrt := math.Cbrt(27)
	fmt.Printf("³√27 = %f\n", cbrt) // 3.000000

	// 10的幂
	pow10 := math.Pow10(3)
	fmt.Printf("10^3 = %f\n", pow10) // 1000.000000
}
```

## 📐 取整操作

### 基本取整

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	num := 3.7

	// 向上取整
	ceil := math.Ceil(num)
	fmt.Printf("向上取整: %f\n", ceil) // 4.000000

	// 向下取整
	floor := math.Floor(num)
	fmt.Printf("向下取整: %f\n", floor) // 3.000000

	// 四舍五入（Go 1.10+）
	round := math.Round(num)
	fmt.Printf("四舍五入: %f\n", round) // 4.000000

	// 截断小数部分
	trunc := math.Trunc(num)
	fmt.Printf("截断: %f\n", trunc) // 3.000000
}
```

### 舍入到指定精度

```go
package main

import (
	"fmt"
	"math"
)

func roundToDecimal(num float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(num*multiplier) / multiplier
}

func main() {
	num := 3.14159

	// 保留2位小数
	rounded := roundToDecimal(num, 2)
	fmt.Printf("保留2位小数: %.2f\n", rounded) // 3.14

	// 保留4位小数
	rounded2 := roundToDecimal(num, 4)
	fmt.Printf("保留4位小数: %.4f\n", rounded2) // 3.1416
}
```

## 📊 三角函数

### 基本三角函数

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	angle := math.Pi / 4 // 45度

	// 正弦
	sin := math.Sin(angle)
	fmt.Printf("sin(45°) = %.4f\n", sin) // 0.7071

	// 余弦
	cos := math.Cos(angle)
	fmt.Printf("cos(45°) = %.4f\n", cos) // 0.7071

	// 正切
	tan := math.Tan(angle)
	fmt.Printf("tan(45°) = %.4f\n", tan) // 1.0000

	// 反三角函数
	asin := math.Asin(sin)
	acos := math.Acos(cos)
	atan := math.Atan(tan)

	fmt.Printf("arcsin: %.4f\n", asin)
	fmt.Printf("arccos: %.4f\n", acos)
	fmt.Printf("arctan: %.4f\n", atan)
}
```

### 双曲函数

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1.0

	// 双曲正弦
	sinh := math.Sinh(x)
	fmt.Printf("sinh(1) = %.4f\n", sinh)

	// 双曲余弦
	cosh := math.Cosh(x)
	fmt.Printf("cosh(1) = %.4f\n", cosh)

	// 双曲正切
	tanh := math.Tanh(x)
	fmt.Printf("tanh(1) = %.4f\n", tanh)
}
```

## 🎲 随机数生成

### 基本随机数

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子（重要！）
	rand.Seed(time.Now().UnixNano())

	// 生成随机整数
	randomInt := rand.Int()
	fmt.Printf("随机整数: %d\n", randomInt)

	// 生成 [0, n) 范围内的随机整数
	randomIntN := rand.Intn(100)
	fmt.Printf("0-99 随机数: %d\n", randomIntN)

	// 生成随机浮点数 [0.0, 1.0)
	randomFloat := rand.Float64()
	fmt.Printf("随机浮点数: %.4f\n", randomFloat)

	// 生成随机浮点数 [0.0, n)
	randomFloatN := rand.Float64() * 100
	fmt.Printf("0-100 随机浮点数: %.2f\n", randomFloatN)
}
```

### 范围随机数

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成 [min, max] 范围内的随机整数
	min, max := 10, 20
	randomInRange := rand.Intn(max-min+1) + min
	fmt.Printf("[%d, %d] 随机数: %d\n", min, max, randomInRange)

	// 生成 [min, max) 范围内的随机浮点数
	minFloat, maxFloat := 1.0, 10.0
	randomFloatRange := rand.Float64()*(maxFloat-minFloat) + minFloat
	fmt.Printf("[%.1f, %.1f) 随机浮点数: %.2f\n", minFloat, maxFloat, randomFloatRange)
}
```

### 随机排列

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 打乱切片
	numbers := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Printf("打乱后: %v\n", numbers)

	// 随机选择
	choices := []string{"apple", "banana", "orange", "grape"}
	randomChoice := choices[rand.Intn(len(choices))]
	fmt.Printf("随机选择: %s\n", randomChoice)
}
```

## 🔢 数学常数

### 常用常数

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 圆周率
	fmt.Printf("π = %.10f\n", math.Pi) // 3.1415926536

	// 自然常数
	fmt.Printf("e = %.10f\n", math.E) // 2.7182818285

	// 最大和最小浮点数
	fmt.Printf("最大 float64: %e\n", math.MaxFloat64)
	fmt.Printf("最小 float64: %e\n", math.SmallestNonzeroFloat64)

	// 最大和最小整数
	fmt.Printf("最大 int64: %d\n", math.MaxInt64)
	fmt.Printf("最小 int64: %d\n", math.MinInt64)
}
```

## 🔄 对数运算

### 基本对数

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 100.0

	// 自然对数（以 e 为底）
	ln := math.Log(x)
	fmt.Printf("ln(100) = %.4f\n", ln)

	// 常用对数（以 10 为底）
	log10 := math.Log10(x)
	fmt.Printf("log10(100) = %.4f\n", log10) // 2.0000

	// 以 2 为底的对数
	log2 := math.Log2(x)
	fmt.Printf("log2(100) = %.4f\n", log2)

	// 对数加 1（避免 log(0)）
	log1p := math.Log1p(x - 1)
	fmt.Printf("log1p(99) = %.4f\n", log1p)
}
```

## 🎯 实用工具函数

### 数值判断

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64

	// 检查是否为 NaN
	f = math.NaN()
	fmt.Printf("是否为 NaN: %t\n", math.IsNaN(f)) // true

	// 检查是否为无穷大
	f = math.Inf(1)
	fmt.Printf("是否为无穷大: %t\n", math.IsInf(f, 1)) // true

	// 检查是否为有限数
	f = 3.14
	fmt.Printf("是否为有限数: %t\n", !math.IsInf(f, 0) && !math.IsNaN(f)) // true
}
```

### 类型转换

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 浮点数转整数（注意精度丢失）
	f := 3.7
	i := int(f)
	fmt.Printf("浮点数转整数: %d\n", i) // 3

	// 整数转浮点数
	i2 := 42
	f2 := float64(i2)
	fmt.Printf("整数转浮点数: %.1f\n", f2) // 42.0

	// 使用 math 包的转换函数
	f3 := 3.7
	i3 := int(math.Round(f3))
	fmt.Printf("四舍五入转整数: %d\n", i3) // 4
}
```

## 🏃‍♂️ 实践应用

### 计算器函数

```go
package main

import (
	"fmt"
	"math"
)

// 计算两点间距离
func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

// 计算圆的面积
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// 计算阶乘
func factorial(n int) int64 {
	if n <= 1 {
		return 1
	}
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

func main() {
	// 距离计算
	dist := distance(0, 0, 3, 4)
	fmt.Printf("两点距离: %.2f\n", dist) // 5.00

	// 圆面积
	area := circleArea(5)
	fmt.Printf("圆面积: %.2f\n", area) // 78.54

	// 阶乘
	fact := factorial(5)
	fmt.Printf("5! = %d\n", fact) // 120
}
```

### 随机密码生成

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}

func main() {
	password := generatePassword(12)
	fmt.Printf("随机密码: %s\n", password)
}
```

## ⚠️ 注意事项

### 1. 随机数种子

```go
// ❌ 错误：不设置种子，每次运行结果相同
randomInt := rand.Intn(100)

// ✅ 正确：设置随机种子
rand.Seed(time.Now().UnixNano())
randomInt := rand.Intn(100)
```

### 2. 浮点数精度

```go
// 注意浮点数精度问题
f1 := 0.1
f2 := 0.2
sum := f1 + f2
fmt.Printf("0.1 + 0.2 = %.20f\n", sum) // 可能不是精确的 0.3

// 比较浮点数时使用误差范围
epsilon := 1e-9
if math.Abs(sum-0.3) < epsilon {
	fmt.Println("相等")
}
```

### 3. 溢出检查

```go
// 检查整数溢出
func safeAdd(a, b int64) (int64, bool) {
	if a > math.MaxInt64-b {
		return 0, false // 溢出
	}
	return a + b, true
}
```

## 📚 扩展阅读

- [math 包官方文档](https://pkg.go.dev/math)
- [math/rand 包官方文档](https://pkg.go.dev/math/rand)

## ⏭️ 下一章节

完成数学运算学习后，可以继续学习其他标准库模块。

---

**💡 提示**: math 包提供了丰富的数学函数，掌握它对于科学计算、游戏开发、数据分析等场景非常重要！

