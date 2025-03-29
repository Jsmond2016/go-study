# Go 实用工具库推荐

golang 有什么好用的库吗，我看很多 切片，字符串、遍历，等操作代码量很麻烦

我来推荐一些 Go 语言中常用且好用的库，可以简化日常开发。


## 一、切片和集合操作：go-funk

[go-funk](https://github.com/thoas/go-funk) 提供了类似 JavaScript 中 Lodash 的功能。

```go
import "github.com/thoas/go-funk"

// 示例代码
func FunkExample() {
    // 切片操作
    numbers := []int{1, 2, 3, 4, 5}
    
    // 过滤
    evens := funk.Filter(numbers, func(x int) bool {
        return x%2 == 0
    })
    
    // 映射
    doubled := funk.Map(numbers, func(x int) int {
        return x * 2
    })
    
    // 查找
    found := funk.Find(numbers, func(x int) bool {
        return x > 3
    })
    
    // 包含检查
    contains := funk.Contains(numbers, 3)
}
```

## 二、字符串处理：xstrings

[xstrings](https://github.com/huandu/xstrings) 提供了丰富的字符串处理函数。

```go
import "github.com/huandu/xstrings"

// 示例代码
func XstringsExample() {
    // 驼峰命名转换
    snakeCase := xstrings.ToSnakeCase("HelloWorld")    // hello_world
    camelCase := xstrings.ToCamelCase("hello_world")   // HelloWorld
    
    // 字符串翻转
    reversed := xstrings.Reverse("Hello")              // olleH
    
    // 字符串宽度调整
    padded := xstrings.LeftPad("Hello", 10, ".")      // .....Hello
}
```

## 三、日期时间处理：carbon

[carbon](https://github.com/golang-module/carbon) 提供了类似 PHP Carbon 的时间处理功能。

```go
import "github.com/golang-module/carbon"

// 示例代码
func CarbonExample() {
    // 时间创建和格式化
    now := carbon.Now()
    
    // 时间计算
    tomorrow := now.AddDay()
    yesterday := now.SubDay()
    
    // 时间比较
    isWeekend := now.IsWeekend()
    isLeapYear := now.IsLeapYear()
    
    // 友好格式化
    diff := now.DiffForHumans() // 例如：1 小时前
}
```

## 四、数据验证：validator

[validator](https://github.com/go-playground/validator) 提供了强大的数据验证功能。

```go
import "github.com/go-playground/validator/v10"

// 示例代码
type User struct {
    Name  string `validate:"required,min=3,max=32"`
    Email string `validate:"required,email"`
    Age   int    `validate:"gte=0,lte=130"`
}

func ValidatorExample() {
    validate := validator.New()
    
    user := User{
        Name:  "John",
        Email: "john@example.com",
        Age:   25,
    }
    
    err := validate.Struct(user)
    if err != nil {
        // 处理验证错误
    }
}
```

## 五、工具集合：lo

[lo](https://github.com/samber/lo) 提供了现代的 Go 工具集。

```go
import "github.com/samber/lo"

// 示例代码
func LoExample() {
    // 切片操作
    numbers := []int{1, 2, 3, 4, 5}
    
    // 过滤
    evens := lo.Filter(numbers, func(x int, _ int) bool {
        return x%2 == 0
    })
    
    // 去重
    unique := lo.Uniq([]int{1, 2, 2, 3, 3, 4})
    
    // 分组
    grouped := lo.GroupBy([]int{1, 2, 3, 4, 5}, func(x int) string {
        if x%2 == 0 {
            return "even"
        }
        return "odd"
    })
}
```

## 六、实用建议

1. **选择标准**：
   - 优先使用标准库
   - 选择维护活跃的第三方库
   - 考虑性能影响
   - 注意版本兼容性

2. **常见使用场景**：
   - 数据处理和转换：go-funk, lo
   - 字符串操作：xstrings
   - 时间处理：carbon
   - 数据验证：validator

3. **性能考虑**：
   - 对性能敏感的场景，优先使用标准库
   - 合理使用反射相关的功能
   - 注意内存使用

4. **项目集成**：
```go
// go.mod 示例
module myproject

go 1.16

require (
    github.com/thoas/go-funk v0.9.3
    github.com/huandu/xstrings v1.4.0
    github.com/golang-module/carbon v1.7.3
    github.com/go-playground/validator/v10 v10.11.2
    github.com/samber/lo v1.37.0
)
```

这些库可以大大提高开发效率，减少重复代码。建议根据项目需求选择合适的库，并注意维护和性能平衡。


其他资源推荐：

- [gookit/goutil: Go 常用的一些工具函数](https://github.com/gookit/goutil)
- [一些值得推荐的golang库 - 黑月教主 - 博客园](https://www.cnblogs.com/achillis/p/18005973)
- [lancet: 一个全面、高效、可复用的go语言工具函数库lancet](https://juejin.cn/post/7036180968050262030)
- [Lancet API 文档](https://www.golancet.cn/api/packages/algorithm.html)