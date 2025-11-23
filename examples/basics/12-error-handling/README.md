# 错误处理示例

本示例展示了 Go 语言的错误处理机制，包括 error 接口、错误创建、错误包装、panic 和 recover 等。

## 运行示例

```bash
go run main.go
```

## 知识点

### Error 接口

Go 语言中的错误是一个接口：
```go
type error interface {
    Error() string
}
```

### 创建错误

1. **errors.New()**：创建简单错误
2. **fmt.Errorf()**：创建格式化错误
3. **自定义错误类型**：实现 error 接口

### 错误处理模式

- 检查错误：`if err != nil { ... }`
- 返回错误：`return nil, err`
- 包装错误：`fmt.Errorf("context: %w", err)`

### Panic 和 Recover

- `panic()`：触发程序崩溃
- `recover()`：捕获 panic，只能在 defer 中使用
- 使用 defer + recover 实现错误恢复

### 错误比较

- `errors.Is()`：检查错误链中是否包含目标错误
- `errors.As()`：检查错误链中是否包含目标错误类型

## 练习

1. 创建一个自定义错误类型 `ValidationError`
2. 实现一个函数，验证用户输入，返回相应的错误
3. 使用错误包装，在错误链中保留原始错误信息
4. 实现一个安全的函数，使用 recover 捕获可能的 panic

