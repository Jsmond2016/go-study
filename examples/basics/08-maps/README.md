# 映射示例

本示例展示了 Go 语言中映射的使用方法和特性。

## 运行示例

```bash
go run main.go
```

## 知识点

### 1. 映射基本特性

- **键值对集合**：存储唯一的键到值的映射
- **无序存储**：映射的元素没有固定顺序
- **动态大小**：可以动态添加和删除键值对
- **引用类型**：映射是引用类型，函数间共享

### 2. 映射创建方式

```go
// nil 映射
var map1 map[string]int

// 空映射
map2 := make(map[string]int)

// 直接初始化
map3 := map[string]int{
    "key1": value1,
    "key2": value2,
}

// 指定初始容量
map4 := make(map[string]int, 100)  // 预分配100个键值对的容量
```

### 3. 基本操作

```go
// 添加/修改元素
m[key] = value

// 访问元素
value, exists := m[key]
if exists {
    fmt.Println("键存在:", value)
} else {
    fmt.Println("键不存在")
}

// 删除元素
delete(m, key)

// 获取长度
length := len(m)
```

### 4. 映射遍历

```go
// 遍历键值对
for key, value := range m {
    fmt.Printf("%s: %v\n", key, value)
}

// 只遍历键
for key := range m {
    fmt.Println(key)
}

// 只遍历值
for _, value := range m {
    fmt.Println(value)
}
```

### 5. 映射比较和复制

```go
// 复制映射（深拷贝）
copyMap := make(map[string]int)
for k, v := range original {
    copyMap[k] = v
}

// Go 1.21+ 使用 maps 包
import "maps"

equal := maps.Equal(map1, map2)  // 比较两个映射是否相等
```

### 6. 映射作为函数参数

```go
func modifyMap(m map[string]int) {
    m["new_key"] = 100  // 会修改原映射
}
```

### 7. 嵌套映射

```go
// 二维映射
nested := map[string]map[string]int{
    "outer1": {
        "inner1": 1,
        "inner2": 2,
    },
    "outer2": {
        "inner3": 3,
    },
}
```

## 重要概念

### 1. 映射内部实现

Go 的映射使用哈希表实现：
- 哈希函数将键转换为数组索引
- 冲突时使用链表或开放寻址解决
- 动态扩容，装载因子过高时重新哈希

### 2. 键的类型要求

- 可比较的类型才能作为键：`int`、`string`、`bool`、指针、数组等
- 不能作为键的类型：`slice`、`map`、`function`
- 结构体作为键要求所有字段都是可比较的

### 3. 零值和 nil 映射

```go
var m map[string]int  // nil 映射，len = 0，不能写入
m = make(map[string]int)  // 空映射，len = 0，可以写入
```

### 4. 并发安全

- 内置映射不是并发安全的
- 多 goroutine 同时读写需要同步机制
- 使用 `sync.Map` 或互斥锁保护

## 性能优化

### 1. 预分配容量

```go
// 知道大概大小时预分配
m := make(map[string]int, 1000)
```

### 2. 合理选择键类型

- 优先使用基本类型：`int`、`string`
- 避免复杂的结构体作为键
- 考虑键的哈希性能

### 3. 内存管理

```go
// 清空映射
for k := range m {
    delete(m, k)
}

// 或者重新创建
m = make(map[string]int)
```

## 实际应用场景

### 1. 配置管理

```go
config := map[string]interface{}{
    "database": map[string]interface{}{
        "host": "localhost",
        "port": 5432,
    },
    "debug": true,
}
```

### 2. 缓存系统

```go
cache := make(map[string]interface{})
cache["user:1"] = userData
```

### 3. 数据统计

```go
wordCount := make(map[string]int)
wordCount["hello"]++
wordCount["world"]++
```

### 4. 关联数据

```go
userToProfile := make(map[int]*Profile)
userToProfile[123] = &Profile{...}
```

## 最佳实践

1. **预分配容量**：知道大小时预分配提高性能
2. **检查键存在**：使用 `value, ok := m[key]` 模式
3. **避免并发竞争**：多 goroutine 使用时加同步
4. **选择合适键类型**：优先简单类型
5. **及时清理**：不需要的数据及时删除

## 常见陷阱

1. **并发读写**：没有同步的并发访问会导致 panic
2. **nil 映射写入**：向 nil 映射写入会 panic
3. **键类型选择**：不可比较的类型不能作为键
4. **内存泄漏**：大映射长期持有数据
5. **顺序依赖**：映射遍历顺序不确定

## Go 1.21+ maps 包

```go
import "maps"

// 比较
equal := maps.Equal(map1, map2)

// 复制
maps.Copy(dest, src)

// 查找键
key, ok := maps.FindKey(m, "target")

// 遍历所有键
maps.Keys(m)
maps.Values(m)
```

## 练习

1. 实现一个单词频率统计器
2. 创建一个简单的内存缓存系统
3. 实现图的邻接表表示
4. 创建一个配置管理系统
5. 实现一个简单的会话管理器
6. 创建一个频率限制器