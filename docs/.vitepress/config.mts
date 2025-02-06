import { defineConfig } from "vitepress"
import mdTaskListPlugin from 'markdown-it-task-lists'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  lang: 'zh-CN',
  title: 'Go语言学习笔记',
  description: "golang study notes",
  base: '/golang-study/',
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Home", link: "/record/home" },
      { text: "Golang 基础", link: "/record/base" },
      { text: "Golang 进阶", link: "/record/advance" },
      { text: "Golang 框架-gin", link: "/record/gin" },
      // { text: "Golang 项目", link: "/projects" },
    ],

    sidebar: {
      // golang 基础
      "/record/base/": getBaseSidebar(),
      // golang 进阶
      "/record/advance/": getAdvanceSidebar(),
      // gin
      "/record/gin": getGinSidebar(),
    },

    socialLinks: [
      { icon: "github", link: "https://github.com/vuejs/vitepress" },
    ],
  },
  markdown: {
    config: (md) => {
      md.use(mdTaskListPlugin)
    },
  },
})

function getBaseSidebar() {
  return [
    {
      text: "基础篇",
      items: [
        {
          text: "fmt常用工具和配置、常用的格式化占位符",
          link: "/record/base/fmt-log-format",
        },
        {
          text: "函数",
          link: "/record/base/func",
        },
        {
          text: "todo-数据类型、字符串和常用操作、类型转换",
          link: "/record/base/data-type",
        },
        {
          text: "todo-切片、集合 及其常用操作",
          link: "/record/base/slice-map",
        },
        {
          text: "todo-指针和结构体",
          link: "/record/base/pointer-struct",
        },
        // 001-第1阶段
        // 方法 https://www.topgoer.com/%E6%96%B9%E6%B3%95/
        // 包
        // 文件操作
        // 接口
        // err 处理
        // 反射

        // 网络编程
        // 并发

        // 常用标准库 time strings strvc
        // 可选-单元测试
        // 数据库操作 MySql、Redis、Mongodb

        // 002-第2阶段 Web 开发阶段
        //
      ],
    },
  ]
}

function getAdvanceSidebar() {
  return [
    {
      text: "go进阶",
      items: [
        {
          text: "fmt常用工具和配置、常用的格式化占位符",
          link: "/record/advance/fmt-log-format",
        },
        // 实战准备：来源-李文周 web 开发视频相关
        // Go连接 MySql
        // databasesql 的增删查改操作
        // mysql 的预处理和事务
        // sqlx 连接 MySql、以及使用方式
        // go redis 连接和使用
        // pipeline 和 watch 事务
        // zap 日志库介绍
        // gin框架集成 zap 日之苦
        // viper 配置加载器的介绍和使用
      ],
    },
  ]
}

function getGinSidebar() {
  return [
    {
      text: "Gin基础",
      items: [
        {
          text: "大纲",
          link: "/record/gin/home",
        },
        {
          text: "开始",
          link: "/record/gin/gin-begin",
        },
        {
          text: "ping 测试",
          link: "/record/gin/gin-ping",
        },
        {
          text: "路由",
          link: "/record/gin/gin-router",
        },
        {
          text: "路由参数",
          link: "/record/gin/gin-router-params",
        },
        {
          text: "路由分组",
          link: "/record/gin/gin-router-group",
        },
        {
          text: "参数-模型-绑定",
          link: "/record/gin/gin-router-params-bind",
        },
        {
          text: "响应类型",
          link: "/record/gin/gin-response-type",
        },
        {
          text: "模板渲染",
          link: "/record/gin/gin-template-render",
        },
        {
          text: "静态文件",
          link: "/record/gin/gin-static",
        },
        {
          text: "同步和异步任务处理",
          link: "/record/gin/gin-sync-async",
        },
        {
          text: "中间件",
          link: "/record/gin/gin-middleware",
        },
        {
          text: "MySql-数据库",
          link: "/record/gin/gin-mysql",
        },
      ],
    },
  ]
}
