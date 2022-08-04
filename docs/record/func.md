# 函数

- https://www.liwenzhou.com/posts/Go/09_function/
- https://www.topgoer.com/%E5%87%BD%E6%95%B0/


**主要知识要点：**

- 参数和可变参数；
- 返回值：
  - 多返回值
  - 返回值命名
  - 返回值补充：切片返回值时，可以返回 nil，没必要返回一个长度为 0 的切片
- 变量作用域：全局变量和局部变量
- 函数类型和类型变量
- 高阶函数
- 匿名函数和闭包 
- defer 语句延缓执行和倒序执行
- 内置错误处理函数：panic / recover， 注意 recover 只能在 defer 中结合使用；