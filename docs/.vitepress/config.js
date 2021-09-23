const taskLists = require('markdown-it-task-lists');


module.exports = {
  lang: 'zh-CN',
  title: 'Go语言学习笔记',
  description: 'Go语言学习笔记',
  base: "/go-study/",
  themeConfig: {
    repo: 'Jsmond2016/go-study',
    docsDir: 'docs',
    editLinks: true,
    editLinkText: '编辑本页',
    lastUpdated: '最后更新时间',

    nav: [
      { text: '首页', link: '/' },
      // {
      //   text: '归档',
      //   link: '/config/basics',
      //   activeMatch: '^/config/'
      // },
      // {
      //   text: '分类',
      //   link: '/category'
      // }
    ],

    sidebar: {
      '/': getGuideSidebar()
    }
  },
  markdown: {

    config: (md) => {
      // use more markdown-it plugins!
      md.use(taskLists)
    }
  }
}

function getGuideSidebar() {
  return [
    {
      text: 'Go语言学习',
      children: [
        { text: '学习目标', link: '/record/goal' },
        { text: '模板', link: '/record/template' },
        { text: 'TODOS', link: '/record/todos' },
        ...require('./sidebar')
      ]
    },

  ]
}
