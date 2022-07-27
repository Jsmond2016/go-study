# go-study

- [推荐学习地址-强烈推荐](https://www.topgoer.com/)
- [李文周博客](https://www.liwenzhou.com/)


其他资料：

- [Go 从入门到实战](https://github.com/xinliangnote/Go) 学习笔记，从零开始学 Go、Gin 框架


## Go 基础


- [数组](https://www.liwenzhou.com/posts/Go/05_array/)
- [切片](https://www.liwenzhou.com/posts/Go/06_slice/)
  - [go语言笔记——切片函数常见操作，增删改查和搜索、排序](https://www.cnblogs.com/bonelee/p/6862627.html)
- [map](https://www.liwenzhou.com/posts/Go/08_map/)
- [函数](https://www.liwenzhou.com/posts/Go/09_function/)
- [指针](https://www.liwenzhou.com/posts/Go/07_pointer/)
- [结构体](https://www.liwenzhou.com/posts/Go/10-struct/)
- [包](https://www.liwenzhou.com/posts/Go/11-package/)
- 接口
- error接口
- 反射
- 并发

## Go 语言基础库

- fmt与格式化占位符
- time
- flag
- log
- 文件操作
- strconv
- net/http
- context


## Gin 框架

> [Gin框架安装与使用](https://www.liwenzhou.com/posts/Go/Gin_framework/)
> [Gin 文档](https://gin-gonic.com/zh-cn/docs/)
> [Gin 视频教程-网易云课堂](https://study.163.com/course/courseLearn.htm?courseId=1210182958#/learn/video?lessonId=1281052216&courseId=1210182958)


遇到问题：

windows 下执行 `go mod tidy`，报错为  `Permission Denied`



看到解决办法：

- go mod init test 随机创建一个包名
- go mod tidy

- [go: go.mod file not found in current directory or any parent directory; see 'go help modules'](https://stackoverflow.com/questions/66894200/go-go-mod-file-not-found-in-current-directory-or-any-parent-directory-see-go)


下载安装太慢?：

- [go get github.com/gin-gonic/gin 下载失败](https://www.cnblogs.com/kevin-yang123/p/14799091.html)


**相关资料** ：

- [7days-golang](https://github.com/geektutu/7days-golang)
- [Go 从入门到实战](https://github.com/xinliangnote/Go) 从零开始学 Go、Gin 框架，基本语法包括 26 个Demo，Gin 框架包括：Gin 自定义路由配置、Gin 使用 Logrus 进行日志记录、Gin 数据绑定和验证、Gin 自定义错误处理、Go gRPC Hello World... 持续更新中
- [基于Gin + Vue + Element UI的前后端分离权限管理系统脚手架](https://github.com/go-admin-team/go-admin) 开源项目，综合学习 Go 开发
- [Go语言高级编程](https://github.com/chai2010/advanced-go-programming-book) 开源电子书
- [Go 从入门到实战](https://github.com/xinliangnote/Go) 从零开始学 Go、Gin 框架，基本语法包括 26 个Demo，Gin 框架包括：Gin 自定义路由配置、Gin 使用 Logrus 进行日志记录、Gin 数据绑定和验证、Gin 自定义错误处理、Go gRPC Hello World... 持续更新中
- [Go语言刷LeetCode](https://github.com/halfrost/LeetCode-Go)

## 项目实战资料

- [基于 Gin 进行模块化设计的 API 框架](https://github.com/xinliangnote/go-gin-api)
