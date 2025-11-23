# Context 示例

本示例展示了 Go 语言 context 包的使用，包括上下文创建、超时控制、取消操作和截止时间等。

## 运行示例

```bash
go run main.go
```

## 知识点

### Context 创建

- `context.Background()`：创建根 context
- `context.TODO()`：创建待实现的 context
- `context.WithValue()`：创建带值的 context

### 超时和取消

- `context.WithTimeout()`：创建带超时的 context
- `context.WithCancel()`：创建可取消的 context
- `context.WithDeadline()`：创建带截止时间的 context

### Context 使用

- `ctx.Done()`：返回一个 channel，当 context 被取消时关闭
- `ctx.Err()`：返回 context 的错误信息
- `ctx.Value()`：获取 context 中的值

## 最佳实践

1. 将 context 作为函数的第一个参数
2. 不要存储 context，而是传递它
3. 使用 context 控制 goroutine 的生命周期
4. 在超时或取消时及时清理资源

## 练习

1. 实现一个带超时的 HTTP 请求函数
2. 创建一个可以取消的长时间运行任务
3. 使用 context 实现请求链路的超时控制

