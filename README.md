# go-study

> 学习资料：[李文周博客](https://www.liwenzhou.com/)

## Go 基础

- 数组

- 切片

  - [go语言笔记——切片函数常见操作，增删改查和搜索、排序](https://www.cnblogs.com/bonelee/p/6862627.html)

- map

- 结构体

- 指针


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

- [Go 从入门到实战](https://github.com/xinliangnote/Go) 从零开始学 Go、Gin 框架，基本语法包括 26 个Demo，Gin 框架包括：Gin 自定义路由配置、Gin 使用 Logrus 进行日志记录、Gin 数据绑定和验证、Gin 自定义错误处理、Go gRPC Hello World... 持续更新中
- [基于Gin + Vue + Element UI的前后端分离权限管理系统脚手架](https://github.com/go-admin-team/go-admin) 开源项目，综合学习 Go 开发
- [Go语言高级编程](https://github.com/chai2010/advanced-go-programming-book) 开源电子书
- [Go语言刷LeetCode](https://github.com/halfrost/LeetCode-Go)