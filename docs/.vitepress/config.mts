import mdTaskListPlugin from "markdown-it-task-lists";
import { defineConfig } from "vitepress";
import { sidebar } from "./sidebar";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  lang: "zh-CN",
  title: "Go语言学习笔记",
  description: "从零基础到实战应用的完整 Go 语言学习路径",
  base: "/go-study/",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    siteTitle: "Go 语言学习笔记",
    logo: "/hand-coding.png",

    nav: [
      { text: "首页", link: "/" },
      { text: "学习指南", link: "/guide/getting-started" },
      { text: "基础语法", link: "/basics/" },
      { text: "标准库", link: "/standard-library/" },
      { text: "Web 开发", link: "/web-development/" },
      { text: "开发工具链", link: "/toolchain/" },
      { text: "实战项目", link: "/projects/" },
    ],

    sidebar: sidebar,

    socialLinks: [
      {
        icon: "github",
        link: "https://github.com/Jsmond2016/go-study",
      },
    ],

    search: {
      provider: 'local',
      options: {
        translations: {
          button: {
            buttonText: '搜索文档',
            buttonAriaLabel: '搜索文档'
          },
          modal: {
            noResultsText: '无法找到相关结果',
            resetButtonTitle: '清除查询条件',
            footer: {
              selectText: '选择',
              navigateText: '切换'
            }
          }
        }
      }
    },

    editLink: {
      pattern: 'https://github.com/Jsmond2016/go-study/edit/main/docs/:path',
      text: '在 GitHub 上编辑此页'
    },

    lastUpdated: {
      text: '最后更新于'
    },

    footer: {
      message: '基于 VitePress 构建',
      copyright: 'Copyright © 2025 Go 语言学习笔记'
    },
  },
  markdown: {
    lineNumbers: true,
    config: (md) => {
      // 支持任务列表
      md.use(mdTaskListPlugin, { enabled: true });
    },
  },
});

// 注意：record 目录已删除，相关内容已整合到正式文档中
// 如需配置侧边栏，请使用 config.ts 中的 sidebar 配置
