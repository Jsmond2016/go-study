import { DefaultTheme } from 'vitepress'

export const sidebar: DefaultTheme.Sidebar = {
  // 学习指南
  '/guide/': [
    {
      text: '学习指南',
      collapsed: false,
      items: [
        { text: '快速开始', link: '/guide/getting-started' },
        { text: '学习路径', link: '/guide/learning-path' },
        { text: '开发环境', link: '/guide/development-environment' },
        {
          text: '学习资源',
          collapsed: false,
          items: [
            { text: '总览', link: '/guide/resources/' },
            { text: '推荐书籍', link: '/guide/resources/books' },
            { text: '视频教程', link: '/guide/resources/videos' },
            { text: '优质博客', link: '/guide/resources/blogs' },
            { text: '开发工具', link: '/guide/resources/tools' }
          ]
        }
      ]
    }
  ],

  // 学习资源（子目录页面访问时使用）
  '/guide/resources/': [
    {
      text: '学习指南',
      collapsed: false,
      items: [
        { text: '快速开始', link: '/guide/getting-started' },
        { text: '学习路径', link: '/guide/learning-path' },
        { text: '开发环境', link: '/guide/development-environment' },
        {
          text: '学习资源',
          collapsed: false,
          items: [
            { text: '总览', link: '/guide/resources/' },
            { text: '推荐书籍', link: '/guide/resources/books' },
            { text: '视频教程', link: '/guide/resources/videos' },
            { text: '优质博客', link: '/guide/resources/blogs' },
            { text: '开发工具', link: '/guide/resources/tools' }
          ]
        }
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
        { text: '16. 测试', link: '/basics/16-testing' },
        { text: '17. 基础练习题', link: '/basics/17-practice-exercises' },
        {
          text: '重点总结',
          collapsed: false,
          items: [
            { text: '01.遍历操作详解', link: '/basics/key-summaries/18-traversal' }
          ]
        }
      ]
    }
  ],

  // 重点总结（子目录页面访问时使用）
  '/basics/key-summaries/': [
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
        { text: '16. 测试', link: '/basics/16-testing' },
        { text: '17. 基础练习题', link: '/basics/17-practice-exercises' },
        {
          text: '重点总结',
          collapsed: false,
          items: [
            { text: '遍历操作详解', link: '/basics/key-summaries/18-traversal' }
          ]
        }
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
        { text: '10. 加密 (crypto)', link: '/standard-library/10-crypto' },
        { text: '11. 字符串操作 (strings)', link: '/standard-library/11-strings' },
        { text: '12. 数学运算 (math)', link: '/standard-library/12-math' }
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
        { text: '08. REST API 设计', link: '/web-development/08-rest-api' },
        { text: '09. Gin 数据库操作', link: '/web-development/09-gin-database' },
        {
          text: '10. Web 开发拓展框架',
          collapsed: false,
          items: [
            { text: '总览', link: '/web-development/web-development-frameworks/' },
            { text: '01. Beego 框架', link: '/web-development/web-development-frameworks/01-beego' },
            { text: '02. Iris 框架', link: '/web-development/web-development-frameworks/02-iris' }
          ]
        }
      ]
    }
  ],

  // Web 开发拓展框架（子目录页面访问时使用）
  '/web-development/web-development-frameworks/': [
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
        { text: '08. REST API 设计', link: '/web-development/08-rest-api' },
        { text: '09. Gin 数据库操作', link: '/web-development/09-gin-database' },
        {
          text: '10. Web 开发拓展框架',
          collapsed: false,
          items: [
            { text: '总览', link: '/web-development/web-development-frameworks/' },
            { text: '01. Beego 框架', link: '/web-development/web-development-frameworks/01-beego' },
            { text: '02. Iris 框架', link: '/web-development/web-development-frameworks/02-iris' }
          ]
        }
      ]
    }
  ],

  // 开发工具链
  '/toolchain/': [
    {
      text: '开发工具链',
      collapsed: false,
      items: [
        { text: '总览', link: '/toolchain/' },
        { text: '01. MySQL 基础', link: '/toolchain/01-mysql' },
        { text: '02. GORM 框架', link: '/toolchain/02-gorm' },
        { text: '03. Go Modules', link: '/toolchain/03-go-modules' },
        { text: '04. Viper 配置管理', link: '/toolchain/04-viper' },
        { text: '05. Zap 日志库', link: '/toolchain/05-zap' },
        { text: '06. JWT 鉴权', link: '/toolchain/06-jwt' },
        { text: '07. Validator 验证', link: '/toolchain/07-validator' },
        { text: '08. CORS 跨域', link: '/toolchain/08-cors' },
        { text: '09. 限流与熔断', link: '/toolchain/09-rate-limit' },
        { text: '10. Swagger API 文档', link: '/toolchain/10-swagger' },
        { text: '11. Redis 缓存', link: '/toolchain/11-redis' },
        { text: '12. Docker 基础', link: '/toolchain/12-docker' },
        { text: '13. 测试框架 Testify', link: '/toolchain/13-testify' },
        { text: '14. 任务调度 Cron', link: '/toolchain/14-cron' },
        { text: '15. WebSocket', link: '/toolchain/15-websocket' }
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
        // 第一部分：gRPC 基础
        { text: '01. gRPC', link: '/microservices/01-grpc' },
        { text: '02. Protocol Buffers', link: '/microservices/02-protobuf' },
        // 第二部分：服务治理
        { text: '03. 服务发现', link: '/microservices/03-service-discovery' },
        { text: '04. 负载均衡', link: '/microservices/04-load-balancing' },
        { text: '05. API 网关', link: '/microservices/05-api-gateway' },
        // 第三部分：进阶主题
        { text: '06. 分布式追踪', link: '/microservices/06-distributed-tracing' },
        { text: '07. 配置中心', link: '/microservices/07-config-center' },
        { text: '08. 消息队列', link: '/microservices/08-message-queue' },
        { text: '09. 服务网格', link: '/microservices/09-service-mesh' },
        // 第四部分：实战项目
        {
          text: '10. 电商微服务实战',
          collapsed: false,
          items: [
            { text: '项目概述', link: '/microservices/ecommerce-microservices/' },
            { text: '01. 环境搭建', link: '/microservices/ecommerce-microservices/01-setup' },
            { text: '02. 架构设计', link: '/microservices/ecommerce-microservices/02-architecture' },
            { text: '03. 用户服务', link: '/microservices/ecommerce-microservices/03-user-service' },
            { text: '04. 商品服务', link: '/microservices/ecommerce-microservices/04-product-service' },
            { text: '05. 订单服务', link: '/microservices/ecommerce-microservices/05-order-service' },
            { text: '06. API 网关', link: '/microservices/ecommerce-microservices/06-gateway' },
            { text: '07. 部署运维', link: '/microservices/ecommerce-microservices/07-deployment' }
          ]
        }
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
        {
          text: '02. 博客系统项目',
          collapsed: false,
          items: [
            { text: '项目概述', link: '/projects/blog-system/' },
            { text: '01. 环境搭建', link: '/projects/blog-system/01-setup' },
            { text: '02. 数据模型设计', link: '/projects/blog-system/02-models' },
            { text: '03. 用户认证', link: '/projects/blog-system/03-auth' },
            { text: '04. 文章管理', link: '/projects/blog-system/04-articles' },
            { text: '05. 评论系统', link: '/projects/blog-system/05-comments' },
            { text: '06. 文件上传', link: '/projects/blog-system/06-upload' },
            { text: '07. 搜索功能', link: '/projects/blog-system/07-search' },
            { text: '08. 部署优化', link: '/projects/blog-system/08-deployment' }
          ]
        },
        {
          text: '03. 电商系统项目',
          collapsed: false,
          items: [
            { text: '项目概述', link: '/projects/e-commerce/' },
            { text: '01. 环境搭建', link: '/projects/e-commerce/01-setup' },
            { text: '02. 数据模型设计', link: '/projects/e-commerce/02-models' },
            { text: '03. 商品管理', link: '/projects/e-commerce/03-products' },
            { text: '04. 购物车', link: '/projects/e-commerce/04-cart' },
            { text: '05. 订单系统', link: '/projects/e-commerce/05-orders' },
            { text: '06. 支付集成', link: '/projects/e-commerce/06-payment' },
            { text: '07. 库存管理', link: '/projects/e-commerce/07-inventory' },
            { text: '08. 部署优化', link: '/projects/e-commerce/08-deployment' }
          ]
        },
        {
          text: '04. 聊天应用项目',
          collapsed: false,
          items: [
            { text: '项目概述', link: '/projects/chat-app/' },
            { text: '01. 环境搭建', link: '/projects/chat-app/01-setup' },
            { text: '02. 数据模型设计', link: '/projects/chat-app/02-models' },
            { text: '03. WebSocket基础', link: '/projects/chat-app/03-websocket' },
            { text: '04. 消息系统', link: '/projects/chat-app/04-messages' },
            { text: '05. 用户状态', link: '/projects/chat-app/05-presence' },
            { text: '06. 群组聊天', link: '/projects/chat-app/06-rooms' },
            { text: '07. 消息推送', link: '/projects/chat-app/07-notifications' },
            { text: '08. 部署优化', link: '/projects/chat-app/08-deployment' }
          ]
        }
      ]
    }
  ],

}

