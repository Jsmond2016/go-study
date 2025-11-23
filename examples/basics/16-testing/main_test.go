package main

import (
	"testing"
)

// 单元测试
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; 期望 %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"正数", 2, 3, 6},
		{"零", 0, 5, 0},
		{"负数", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; 期望 %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("正常除法", func(t *testing.T) {
		result, err := Divide(10, 2)
		if err != nil {
			t.Errorf("Divide(10, 2) 返回错误: %v", err)
		}
		if result != 5.0 {
			t.Errorf("Divide(10, 2) = %f; 期望 5.0", result)
		}
	})

	t.Run("除以零", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) 应该返回错误")
		}
	})
}

// 基准测试
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(2, 3)
	}
}

// 示例测试
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

