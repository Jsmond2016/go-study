import mdTaskListPlugin from 'markdown-it-task-lists'
import { defineConfig } from 'vitepress'
import { sidebar } from './sidebar'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Go 语言学习笔记',
  description: '从零基础到实战应用的完整 Go 语言学习路径',

  // 基础路径配置
  base: '/',

  // 语言配置
  lang: 'zh-CN',

  // 主题配置
  themeConfig: {
    // 网站标题
    siteTitle: 'Go 语言学习笔记',

    // Logo
    logo: '/hand-coding.png',

    // 导航栏配置
    nav: [
      { text: '首页', link: '/' },
      { text: '学习指南', link: '/guide/getting-started' },
      { text: '基础语法', link: '/basics/' },
      { text: '标准库', link: '/standard-library/' },
      { text: 'Web 开发', link: '/web-development/' },
      { text: '开发工具链', link: '/toolchain/' },
      { text: '实战项目', link: '/projects/' },
    ],

    // 侧边栏配置
    sidebar,

    // 社交链接
    socialLinks: [
      { icon: 'github', link: 'https://github.com/Jsmond2016/go-study' }
    ],

    // 搜索配置
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

    // 编辑链接
    editLink: {
      pattern: 'https://github.com/Jsmond2016/go-study/edit/main/docs/:path',
      text: '在 GitHub 上编辑此页'
    },

    // 最后更新时间
    lastUpdated: {
      text: '最后更新于',
      formatOptions: {
        dateStyle: 'short',
        timeStyle: 'medium'
      }
    },

    // 页脚配置
    footer: {
      message: '基于 VitePress 构建',
      copyright: 'Copyright © 2025 Go 语言学习笔记'
    },

    // 大纲配置
    outline: {
      level: [2, 3],
      label: '页面导航'
    },

    // 返回顶部按钮
    returnToTopLabel: '返回顶部',

    // 深色模式切换
    darkModeSwitchLabel: '外观',

    // 侧边栏菜单标签
    sidebarMenuLabel: '菜单',

    // 上一页/下一页文本
    docFooter: {
      prev: '上一页',
      next: '下一页'
    }
  },

  // Markdown 配置
  markdown: {
    lineNumbers: true,
    theme: {
      light: 'github-light',
      dark: 'github-dark'
    },
    config: (md) => {
      // 支持任务列表
      md.use(mdTaskListPlugin, { enabled: true })
    }
  },

  // 开发服务器配置
  vite: {
    server: {
      port: 3000,
      host: '127.0.0.1'
    }
  }
})
