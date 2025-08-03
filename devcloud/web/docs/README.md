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

## 登陆页面

![alt text](login.png)

## Layout(多套子系统)

## 具体子页面
