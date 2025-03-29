import { defineConfig } from "vitepress";
import mdTaskListPlugin from "markdown-it-task-lists";
import {
  baseNavConfig,
  baseNavSideBar,
} from "./navSideConfig/base.navConfig.mts";
// import { advanceNavConfig, advanceNavSideBar } from './navSideConfig/advance.navConfig.mts'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  lang: "zh-CN",
  title: "Go语言学习笔记",
  description: "golang study notes",
  base: "/go-study/",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "资料", link: "/record/info" },
      // { text: "Golang 基础", link: "/record/base" },
      baseNavConfig,
      // advanceNavConfig,
      { text: "Golang 框架-gin", link: "/record/gin" },
      // { text: "Golang 项目", link: "/projects" },
    ],

    sidebar: {
      // 首页
      "/record/info/": getInfoSidebar(),
      // golang 基础
      ...baseNavSideBar,
      //  ...advanceNavSideBar,
      // golang 进阶
      "/record/advance/": getAdvanceSidebar(),
      // gin
      "/record/gin": getGinSidebar(),
    },

    socialLinks: [
      {
        icon: "github",
        link: "https://vitepress.dev/zh/guide/what-is-vitepress",
      },
    ],
  },
  markdown: {
    config: (md) => {
      md.use(mdTaskListPlugin);
    },
  },
});

function getInfoSidebar() {
  return [
    {
      text: "基础篇",
      items: [
        { text: "Go 语言学习目标及成果", link: "/record/info/goal-result" },
        { text: "学习大纲", link: "/record/info/structure" },
      ],
    },
  ];
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
  ];
}

function getGinSidebar() {
  return [
    {
      text: "Gin基础",
      items: [
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
  ];
}
