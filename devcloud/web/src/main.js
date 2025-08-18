import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

// UI 组件库
import ArcoVue from '@arco-design/web-vue'
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
import '@arco-design/web-vue/dist/arco.css'
app.use(ArcoVue,)
app.use(ArcoVueIcon)

// 样式加载
import './assets/main.css'

app.use(router)
app.mount('#app')
