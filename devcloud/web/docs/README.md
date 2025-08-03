# 页面开发流程

vue3 + ui组件库

## 安装UI组件

- [element plus (vue3版本)](https://element-plus.sxtxhy.com/zh-CN/)
- [TDesign](https://tdesign.tencent.com/vue-next/overview)
- [ant design for vue](https://2x.antdv.com/components/page-header-cn)
- [arco design](https://arco.design/vue/docs/start)

从UI的精美和官方的维护度来考虑, 这里选择使用 arco design

```sh
npm install --save-dev @arco-design/web-vue
```

```js
// UI 组件库
import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'
app.use(ArcoVue)
```

## ajax

如何将表单数据 提交给服务器, 封装的比较友好的js的http客户端, [axios](https://axios-http.com/docs/intro)

```sh
import axios from 'axios'

// 封装一个axios的实例，http cient实例
// https://axios-http.com/zh/docs/instance
const client = axios.create({
  baseURL: 'http://127.0.0.1:8080',
  timeout: 3000,
})

export default client
```

基于客户端封装的mcenter client

```sh
import client from './client'

var API = {
  Login: (data) => {
    return client.post('/api/devcloud/v1/token', data)
  },
}

export default API
```

## 配置vite代理开发

```js
// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), vueDevTools()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    proxy: {
      '/api': 'http://127.0.0.1:8080',
    },
  },
})
```

去掉baseURL，应用使用代理

```js
// 封装一个axios的实例，http cient实例
// https://axios-http.com/zh/docs/instance
const client = axios.create({
  timeout: 3000,
})
```

## 统一处理 API的 异常

设置axios的中间件, 统一识别并处理报错

```js
// 拦截API的返回结果, 如果是异常 提取异常信息，并展示
// 如果正常 则直接返回 API的响应结果
// https://axios-http.com/zh/docs/interceptors
client.interceptors.response.use(
  (data) => {
    return data.data
  },
  (error) => {
    // 提取错误
    var msg = error.message

    // 如果有更加详细的信息
    if (error.response.data.message) {
      msg = error.response.data.message
    }

    Message.error(msg)
    return error
  },
)
```

## 登陆页面

![alt text](login.png)

1. UI

```vue
<template>
  <div class="login">
    <!-- form登录表单 -->
    <div :style="{ width: '480px', height: '400px' }">
      <div
        style="display: flex; height: 40px;justify-content: center;align-items:center;margin-bottom: 12px;"
      >
        <h2>欢迎登录</h2>
      </div>
      <a-form :model="form" size="large" @submit="handleSubmit">
        <a-form-item hide-label field="name" tooltip="Please enter username" label="Username">
          <a-input v-model="form.name" placeholder="please enter your username...">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item hide-label field="password" label="Password">
          <a-input-password
            :invisible-button="false"
            v-model="form.password"
            placeholder="Please enter something"
            allow-clear
          >
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item hide-label field="remember">
          <a-checkbox v-model="form.remember"> 记住 </a-checkbox>
        </a-form-item>
        <a-form-item hide-label>
          <a-button type="primary" :style="{ width: '100%' }" html-type="submit">登录</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>
```

2. 对接

```js
// 表单的提交
const handleSubmit = async (data) => {
  // 判断表单是否验证成功
  if (data.errors == null) {
    // 通过HttpClient把数据提交给后端
    const resp = await mcenter.Login(data.values)
    console.log(resp)
  }
}
```

3. 状态处理(useStorage)

```sh
npm i @vueuse/core
```

```sh
import { useStorage } from '@vueuse/core'

export default useStorage('token', {}, localStorage, { mergeDefaults: true })
```

```js
// 表单的提交
const handleSubmit = async (data) => {
  // 判断表单是否验证成功
  if (data.errors == null) {
    // 通过HttpClient把数据提交给后端
    const resp = await mcenter.Login(data.values)
    // 用户登录后，这些用户信息我们怎么保存喃 1. cookie, session storage, localstorage
    // localStorage.setItem('token', JSON.stringify(resp))
    // localStorage.getItem('token')
    // 把 localStorage 做成响应式的
    // https://cn.vuejs.org/guide/reusability/composables.html
    token.value = resp
  }
}
```

4. 登录成功后进行跳转

```js
// 登录完成后，需要跳转到后台页面
// location.assign('/users/eduardo')
// vue router对象, 提供路由跳转的功能
// 字符串路径
//  router.push('/users/eduardo')
// 指定到组件名称
router.push({ name: 'HomePage' })
```

## Layout(多套子系统)

## 具体子页面
