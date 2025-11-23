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
            { text: '快速开始', link: '/guide/getting-started' },
            { text: '学习路径', link: '/guide/learning-path' },
            { text: '开发环境配置', link: '/guide/development-environment' }
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
            { text: '指针', link: '/basics/10-pointers' },
            { text: '接口', link: '/basics/11-interfaces' },
            { text: '错误处理', link: '/basics/12-error-handling' },
            { text: '包管理', link: '/basics/13-packages' },
            { text: '并发编程', link: '/basics/14-concurrency' },
            { text: '反射', link: '/basics/15-reflection' },
            { text: '测试', link: '/basics/16-testing' }
          ]
        }
      ],

      '/advanced/': [
        {
          text: '进阶主题',
          items: [
            { text: '总览', link: '/advanced/' },
            { text: '泛型', link: '/advanced/01-generics' },
            { text: '性能优化', link: '/advanced/02-performance' },
            { text: '内存管理', link: '/advanced/03-memory-management' },
            { text: '垃圾回收调优', link: '/advanced/04-gc-tuning' },
            { text: '性能分析', link: '/advanced/05-pprof' }
          ]
        }
      ],

      '/standard-library/': [
        {
          text: '标准库',
          items: [
            { text: '总览', link: '/standard-library/' },
            { text: '格式化输出', link: '/standard-library/01-fmt' },
            { text: '时间处理', link: '/standard-library/02-time' },
            { text: '命令行参数', link: '/standard-library/03-flag' },
            { text: '日志', link: '/standard-library/04-log' },
            { text: '文件操作', link: '/standard-library/05-file-io' },
            { text: '字符串转换', link: '/standard-library/06-strconv' },
            { text: 'HTTP 客户端', link: '/standard-library/07-net-http' },
            { text: '上下文', link: '/standard-library/08-context' },
            { text: '编码解码', link: '/standard-library/09-encoding' },
            { text: '加密', link: '/standard-library/10-crypto' }
          ]
        }
      ],

      '/database/': [
        {
          text: '数据库操作',
          items: [
            { text: '总览', link: '/database/' },
            { text: 'SQL 基础', link: '/database/01-sql' },
            { text: 'GORM 框架', link: '/database/02-gorm' },
            { text: 'Redis 操作', link: '/database/03-redis' },
            { text: 'MongoDB 操作', link: '/database/04-mongodb' },
            { text: '连接池', link: '/database/05-connection-pool' }
          ]
        }
      ],

      '/web-development/': [
        {
          text: 'Web 开发',
          items: [
            { text: '总览', link: '/web-development/' },
            { text: 'HTTP 服务器', link: '/web-development/01-http-server' },
            { text: 'Gin 基础', link: '/web-development/02-gin-basics' },
            { text: 'Gin 路由', link: '/web-development/03-gin-routing' },
            { text: 'Gin 中间件', link: '/web-development/04-gin-middleware' },
            { text: 'Gin 模板', link: '/web-development/05-gin-template' },
            { text: '数据验证', link: '/web-development/06-gin-validation' },
            { text: '认证授权', link: '/web-development/07-gin-auth' },
            { text: 'REST API 设计', link: '/web-development/08-rest-api' }
          ]
        }
      ],

      '/microservices/': [
        {
          text: '微服务',
          items: [
            { text: '总览', link: '/microservices/' },
            { text: 'gRPC', link: '/microservices/01-grpc' },
            { text: 'Protocol Buffers', link: '/microservices/02-protobuf' },
            { text: '服务发现', link: '/microservices/03-service-discovery' },
            { text: '负载均衡', link: '/microservices/04-load-balancing' },
            { text: 'API 网关', link: '/microservices/05-api-gateway' }
          ]
        }
      ],

      '/devops/': [
        {
          text: '运维部署',
          items: [
            { text: '总览', link: '/devops/' },
            { text: 'Docker 容器化', link: '/devops/01-docker' },
            { text: 'Kubernetes 编排', link: '/devops/02-kubernetes' },
            { text: 'CI/CD 流水线', link: '/devops/03-ci-cd' },
            { text: '监控告警', link: '/devops/04-monitoring' },
            { text: '日志管理', link: '/devops/05-logging' }
          ]
        }
      ],

      '/projects/': [
        {
          text: '实战项目',
          items: [
            { text: '总览', link: '/projects/' },
            { text: 'TODO API 项目', link: '/projects/01-todo-api' },
            { text: '博客系统项目', link: '/projects/02-blog-system' },
            { text: '电商系统项目', link: '/projects/03-e-commerce' },
            { text: '聊天应用项目', link: '/projects/04-chat-app' }
          ]
        }
      ],

      '/resources/': [
        {
          text: '学习资源',
          items: [
            { text: '总览', link: '/resources/' },
            { text: '推荐书籍', link: '/resources/books' },
            { text: '视频教程', link: '/resources/videos' },
            { text: '优质博客', link: '/resources/blogs' },
            { text: '开发工具', link: '/resources/tools' }
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
