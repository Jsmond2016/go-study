export const baseNavConfig = { text: "Golang 基础", link: "/record/base" }
const sidebarConfig = [
  {
    text: "开发准备",
    items: [
      {
        text: "goenv 安装多版本 go",
        link: "/record/base/goenv-config",
      },
    ]
  },
  {
    text: "基础语法篇",
    items: [
      {
        text: "fmt常用工具和配置、常用的格式化占位符",
        link: "/record/base/fmt-log-format",
      },
      {
        text: '常用的 fmt 方法',
        link: '/record/base/go-fmt'
      },
      {
        text: '数字操作常用方法',
        link: '/record/base/go-number'
      },
      {
        text: '字符串常用方法',
        link: '/record/base/go-string'
      },
      {
        text: 'strconv 常用的方法',
        link: '/record/base/go-strconv'
      },
      {
        text: '切片的常用方法',
        link: '/record/base/go-slice'
      },
      {
        text: 'map常用方法',
        link: '/record/base/golang-map'
      },
      {
        text: '循环常用操作',
        link: '/record/base/go-loop'
      },
      {
        text: '基础综合练习',
        link: '/record/base/golang-basic-practice'
      },
      {
        text: "函数",
        link: "/record/base/func",
      },
      // {
      //   text: "todo-数据类型、字符串和常用操作、类型转换",
      //   link: "/record/base/data-type",
      // },
      // {
      //   text: "todo-切片、集合 及其常用操作",
      //   link: "/record/base/slice-map",
      // },
      // {
      //   text: "todo-指针和结构体",
      //   link: "/record/base/pointer-struct",
      // },
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
  {
    text: "基础拓展",
    items: [
        {
          text: '实用工具库推荐',
          link: '/record/base/golang-util-lib'
        }
    ]
  }
]

export const baseNavSideBar = {
  "/record/base/": sidebarConfig,
}