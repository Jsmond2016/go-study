import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Go 语言学习笔记',
  description: 'Golang 学习的点滴记录',
  lang: 'zh-CN',
  
  // 主题配置
  themeConfig: {
    // 导航菜单
    nav: [
      { text: '首页', link: '/' },
      { text: '学习指南', link: '/guide/getting-started' },
      { text: '基础语法', link: '/basics/' },
      { text: '标准库', link: '/standard-library/' },
      { text: 'Web开发', link: '/web-development/' },
      { text: '数据库', link: '/database/' },
      { text: '实战项目', link: '/projects/' },
      { text: '学习资源', link: '/resources/' }
    ],

    // 侧边栏配置
    sidebar: {
      '/guide/': [
        {
          text: '学习指南',
          items: [
            { text: '快速开始', link: '/guide/getting-started' }
          ]
        }
      ],

      '/basics/': [
        {
          text: 'Go 基础语法',
          items: [
            { text: '总览', link: '/basics/' },
            { text: '变量与常量', link: '/basics/01-variables-constants' },
            { text: '数据类型', link: '/basics/02-data-types' },
            { text: '运算符', link: '/basics/03-operators' },
            { text: '控制流程', link: '/basics/04-control-flow' },
            { text: '函数', link: '/basics/05-functions' },
            { text: '数组', link: '/basics/06-arrays' },
            { text: '切片', link: '/basics/07-slices' },
            { text: '映射', link: '/basics/08-maps' },
            { text: '结构体', link: '/basics/09-structs' },
            { text: '指针', link: '/basics/10-pointers' }
          ]
        }
      ]
    },

    // 社交链接
    socialLinks: [
      { icon: 'github', link: 'https://github.com/Jsmond2016/go-study' }
    ],

    // 页脚
    footer: {
      message: '基于 MIT 许可发布',
      copyright: `版权所有 © 2025-${new Date().getFullYear()} Jsmond2016`
    },

    // 搜索
    search: {
      provider: 'local'
    },

    // 编辑链接
    editLink: {
      pattern: 'https://github.com/Jsmond2016/go-study/edit/main/docs/:path',
      text: '在 GitHub 上编辑此页面'
    },

    // 上次更新
    lastUpdated: {
      text: '最后更新于',
      formatOptions: {
        dateStyle: 'short',
        timeStyle: 'medium'
      }
    },

    // 文档页脚
    docFooter: {
      prev: '上一页',
      next: '下一页'
    },

    // 大纲标题
    outline: {
      label: '页面导航'
    },

    // 返回顶部
    returnToTopLabel: '回到顶部'
  },

  // Markdown 配置
  markdown: {
    lineNumbers: true,
    config: (md) => {
      // 可以在这里添加 markdown-it 插件
    }
  },

  // 构建配置
  base: '/',
  outDir: '.vitepress/dist',
  cacheDir: '.vitepress/cache',

  // Head 配置
  head: [
    ['link', { rel: 'icon', href: '/favicon.ico' }],
    ['meta', { name: 'theme-color', content: '#3c82f6' }],
    ['meta', { name: 'keywords', content: 'Go, Golang, 学习, 教程, 编程' }]
  ]
})
