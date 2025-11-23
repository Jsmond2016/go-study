// VitePress 主题配置
import type { Theme } from 'vitepress'
import DefaultTheme from 'vitepress/theme'
import './style.css'

export default {
  extends: DefaultTheme,
  enhanceApp({ app, router, siteData }) {
    // 这里可以添加自定义主题配置
  }
} satisfies Theme
