---
title: 测试
difficulty: intermediate
duration: "4-5小时"
prerequisites: ["函数", "包管理"]
tags: ["测试", "单元测试", "基准测试", "覆盖率"]
---

# 测试

测试是保证代码质量的重要手段。Go 语言内置了强大的测试工具，让编写和运行测试变得简单。

## 📋 学习目标

- [ ] 理解测试的重要性
- [ ] 掌握单元测试的编写
- [ ] 学会编写基准测试
- [ ] 理解测试覆盖率
- [ ] 掌握表格驱动测试
- [ ] 了解测试的最佳实践

## 🎯 测试基础

### 测试文件命名

测试文件必须以 `_test.go` 结尾：

```
math.go          # 源代码
math_test.go     # 测试文件
```

### 测试函数命名

测试函数必须以 `Test` 开头，接受 `*testing.T` 参数：

```go
func TestFunctionName(t *testing.T)
```

## 🧪 单元测试

### 基本测试

```go
package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; 期望 %d", result, expected)
	}
}
```

### 运行测试

```bash
# 运行当前包的测试
go test

# 运行所有测试
go test ./...

# 显示详细信息
go test -v

# 运行特定测试
go test -run TestAdd
```

### 使用 t.Run 子测试

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"正数", 2, 3, 5},
		{"负数", -1, -2, -3},
		{"零", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; 期望 %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
```

## 📊 表格驱动测试

表格驱动测试是 Go 中常用的测试模式：

```go
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "正数相乘",
			a:        2,
			b:        3,
			expected: 6,
		},
		{
			name:     "负数相乘",
			a:        -2,
			b:        3,
			expected: -6,
		},
		{
			name:     "零相乘",
			a:        0,
			b:        5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; 期望 %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
```

## ⚡ 基准测试

基准测试用于测量代码性能。

### 基本语法

```go
func BenchmarkFunctionName(b *testing.B)
```

### 示例

```go
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}
```

### 运行基准测试

```bash
# 运行基准测试
go test -bench=.

# 运行特定基准测试
go test -bench=BenchmarkAdd

# 显示内存分配
go test -bench=. -benchmem

# 比较基准测试结果
go test -bench=. -benchmem -count=3
```

### 基准测试示例

```go
func BenchmarkStringConcat(b *testing.B) {
	b.Run("使用 +", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = "Hello" + " " + "World"
		}
	})

	b.Run("使用 fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s %s", "Hello", "World")
		}
	})

	b.Run("使用 strings.Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var builder strings.Builder
			builder.WriteString("Hello")
			builder.WriteString(" ")
			builder.WriteString("World")
			_ = builder.String()
		}
	})
}
```

## 📈 测试覆盖率

### 生成覆盖率报告

```bash
# 生成覆盖率文件
go test -coverprofile=coverage.out

# 查看覆盖率
go tool cover -func=coverage.out

# 生成 HTML 报告
go tool cover -html=coverage.out
```

### 覆盖率目标

```bash
# 设置覆盖率阈值
go test -cover -coverprofile=coverage.out
go tool cover -func=coverage.out | grep total | awk '{print $3}'
```

## 🎯 示例测试

示例测试既是测试也是文档：

```go
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func ExampleMultiply() {
	result := Multiply(2, 3)
	fmt.Println(result)
	// Output: 6
}
```

## 🔧 测试辅助函数

### 创建测试辅助函数

```go
func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("期望 %d, 得到 %d", want, got)
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5)
}
```

### 使用 t.Helper()

`t.Helper()` 标记函数为辅助函数，错误信息会指向调用者：

```go
func assertEqual(t *testing.T, got, want int) {
	t.Helper()  // 重要！
	if got != want {
		t.Errorf("期望 %d, 得到 %d", want, got)
	}
}
```

## 🧹 测试清理

### 使用 t.Cleanup()

```go
func TestWithCleanup(t *testing.T) {
	// 设置测试环境
	tempFile := createTempFile(t)

	// 注册清理函数
	t.Cleanup(func() {
		os.Remove(tempFile.Name())
	})

	// 测试代码
	// ...
}
```

## 🎭 测试技巧

### 1. 使用测试表

```go
tests := []struct {
	name     string
	input    int
	expected int
}{
	{"case1", 1, 2},
	{"case2", 2, 4},
}
```

### 2. 并行测试

```go
func TestParallel(t *testing.T) {
	t.Parallel()
	// 测试代码
}
```

### 3. 跳过测试

```go
func TestSkip(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过长时间测试")
	}
	// 测试代码
}
```

### 4. 条件测试

```go
func TestConditional(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows 不支持此功能")
	}
	// 测试代码
}
```

## 🏃‍♂️ 实践练习

### 练习 1: 编写单元测试

为你的函数编写完整的单元测试。

### 练习 2: 编写基准测试

比较不同实现的性能。

### 练习 3: 提高测试覆盖率

确保测试覆盖率 > 80%。

## 🤔 思考题

1. 测试应该测试什么？
2. 如何组织测试代码？
3. 什么时候使用基准测试？
4. 如何提高测试覆盖率？
5. 测试和文档的关系是什么？

## 📚 扩展阅读

- [Go 测试官方文档](https://golang.org/pkg/testing/)
- [测试最佳实践](https://golang.org/doc/effective_go.html#testing)
- [高级测试技巧](https://github.com/golang/go/wiki/TestComments)

## ⏭️ 下一阶段

完成基础语法学习后，可以进入：

- [进阶主题](../advanced/) - 泛型、性能优化等
- [标准库](../standard-library/) - 常用标准库详解
- [Web 开发](../web-development/) - Gin 框架和 Web 应用

---

**💡 提示**: 良好的测试是代码质量的保证。编写测试不仅验证功能，也是最好的文档！

