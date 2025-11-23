import { DefaultTheme } from 'vitepress'

export const sidebar: DefaultTheme.Sidebar = {
  // 学习指南
  '/guide/': [
    {
      text: '学习指南',
      items: [
        { text: '快速开始', link: '/guide/getting-started' },
        { text: '学习路径', link: '/guide/learning-path' },
        { text: '开发环境', link: '/guide/development-environment' }
      ]
    }
  ],
  
  // 基础语法
  '/basics/': [
    {
      text: '基础语法',
      collapsed: false,
      items: [
        { text: '总览', link: '/basics/' },
        { text: '01. 变量与常量', link: '/basics/01-variables-constants' },
        { text: '02. 数据类型', link: '/basics/02-data-types' },
        { text: '03. 运算符', link: '/basics/03-operators' },
        { text: '04. 控制流程', link: '/basics/04-control-flow' },
        { text: '05. 函数', link: '/basics/05-functions' },
        { text: '06. 数组', link: '/basics/06-arrays' },
        { text: '07. 切片', link: '/basics/07-slices' },
        { text: '08. 映射', link: '/basics/08-maps' },
        { text: '09. 结构体', link: '/basics/09-structs' },
        { text: '10. 指针', link: '/basics/10-pointers' },
        { text: '11. 接口', link: '/basics/11-interfaces' },
        { text: '12. 错误处理', link: '/basics/12-error-handling' },
        { text: '13. 包管理', link: '/basics/13-packages' },
        { text: '14. 并发编程', link: '/basics/14-concurrency' },
        { text: '15. 反射', link: '/basics/15-reflection' },
        { text: '16. 测试', link: '/basics/16-testing' }
      ]
    }
  ],
  
  // 进阶主题
  '/advanced/': [
    {
      text: '进阶主题',
      collapsed: false,
      items: [
        { text: '总览', link: '/advanced/' },
        { text: '01. 泛型', link: '/advanced/01-generics' },
        { text: '02. 性能优化', link: '/advanced/02-performance' },
        { text: '03. 内存管理', link: '/advanced/03-memory-management' },
        { text: '04. 垃圾回收调优', link: '/advanced/04-gc-tuning' },
        { text: '05. 性能分析', link: '/advanced/05-pprof' }
      ]
    }
  ],
  
  // 标准库
  '/standard-library/': [
    {
      text: '标准库',
      collapsed: false,
      items: [
        { text: '总览', link: '/standard-library/' },
        { text: '01. 格式化输出 (fmt)', link: '/standard-library/01-fmt' },
        { text: '02. 时间处理 (time)', link: '/standard-library/02-time' },
        { text: '03. 命令行参数 (flag)', link: '/standard-library/03-flag' },
        { text: '04. 日志 (log)', link: '/standard-library/04-log' },
        { text: '05. 文件操作 (file-io)', link: '/standard-library/05-file-io' },
        { text: '06. 字符串转换 (strconv)', link: '/standard-library/06-strconv' },
        { text: '07. HTTP 客户端 (net/http)', link: '/standard-library/07-net-http' },
        { text: '08. 上下文 (context)', link: '/standard-library/08-context' },
        { text: '09. 编码解码 (encoding)', link: '/standard-library/09-encoding' },
        { text: '10. 加密 (crypto)', link: '/standard-library/10-crypto' }
      ]
    }
  ],
  
  // 数据库操作
  '/database/': [
    {
      text: '数据库操作',
      collapsed: false,
      items: [
        { text: '总览', link: '/database/' },
        { text: '01. SQL 基础', link: '/database/01-sql' },
        { text: '02. GORM 框架', link: '/database/02-gorm' },
        { text: '03. Redis 操作', link: '/database/03-redis' },
        { text: '04. MongoDB 操作', link: '/database/04-mongodb' },
        { text: '05. 连接池', link: '/database/05-connection-pool' }
      ]
    }
  ],
  
  // Web 开发
  '/web-development/': [
    {
      text: 'Web 开发',
      collapsed: false,
      items: [
        { text: '总览', link: '/web-development/' },
        { text: '01. HTTP 服务器', link: '/web-development/01-http-server' },
        { text: '02. Gin 基础', link: '/web-development/02-gin-basics' },
        { text: '03. Gin 路由', link: '/web-development/03-gin-routing' },
        { text: '04. Gin 中间件', link: '/web-development/04-gin-middleware' },
        { text: '05. Gin 模板', link: '/web-development/05-gin-template' },
        { text: '06. 数据验证', link: '/web-development/06-gin-validation' },
        { text: '07. 认证授权', link: '/web-development/07-gin-auth' },
        { text: '08. REST API 设计', link: '/web-development/08-rest-api' }
      ]
    }
  ],
  
  // 微服务
  '/microservices/': [
    {
      text: '微服务',
      collapsed: false,
      items: [
        { text: '总览', link: '/microservices/' },
        { text: '01. gRPC', link: '/microservices/01-grpc' },
        { text: '02. Protocol Buffers', link: '/microservices/02-protobuf' },
        { text: '03. 服务发现', link: '/microservices/03-service-discovery' },
        { text: '04. 负载均衡', link: '/microservices/04-load-balancing' },
        { text: '05. API 网关', link: '/microservices/05-api-gateway' }
      ]
    }
  ],
  
  // 运维部署
  '/devops/': [
    {
      text: '运维部署',
      collapsed: false,
      items: [
        { text: '总览', link: '/devops/' },
        { text: '01. Docker 容器化', link: '/devops/01-docker' },
        { text: '02. Kubernetes 编排', link: '/devops/02-kubernetes' },
        { text: '03. CI/CD 流水线', link: '/devops/03-ci-cd' },
        { text: '04. 监控告警', link: '/devops/04-monitoring' },
        { text: '05. 日志管理', link: '/devops/05-logging' }
      ]
    }
  ],
  
  // 实战项目
  '/projects/': [
    {
      text: '实战项目',
      collapsed: false,
      items: [
        { text: '总览', link: '/projects/' },
        { text: '01. TODO API 项目', link: '/projects/01-todo-api' },
        { text: '02. 博客系统项目', link: '/projects/02-blog-system' },
        { text: '03. 电商系统项目', link: '/projects/03-e-commerce' },
        { text: '04. 聊天应用项目', link: '/projects/04-chat-app' }
      ]
    }
  ],
  
  // 学习资源
  '/resources/': [
    {
      text: '学习资源',
      collapsed: false,
      items: [
        { text: '总览', link: '/resources/' },
        { text: '推荐书籍', link: '/resources/books' },
        { text: '视频教程', link: '/resources/videos' },
        { text: '优质博客', link: '/resources/blogs' },
        { text: '开发工具', link: '/resources/tools' }
      ]
    }
  ]
}

