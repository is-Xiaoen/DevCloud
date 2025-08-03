<template>
  <div class="login">
    <!-- form登录表单 -->
    <div :style="{ width: '480px', height: '400px' }">
      <div style="display: flex; height: 40px;justify-content: center;align-items:center;margin-bottom: 12px;">
        <h2>欢迎登录</h2>
      </div>
      <a-form :model="form" size="large" @submit="handleSubmit">
        <a-form-item hide-label field="parameter.username" :rules="{
          required: true,
          message: '请输入用户名'
        }">
          <a-input v-model="form.parameter.username" placeholder="请输入用户名">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item hide-label field="parameter.password" required :rules="{
          required: true,
          message: '请输入密码'
        }">
          <a-input-password :invisible-button="false" v-model="form.parameter.password" placeholder="请输入密码" allow-clear>
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

<script setup>
import { ref } from 'vue';
import mcenter from '@/api/mcenter'
import token from '@/storage/token'
import { useRouter } from 'vue-router';

// 表单对象
const form = ref({
  issuer: 'password',
  parameter: {
    username: '',
    password: ''
  },
  remember: false,
})

// 获取router对象
const router = useRouter()

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

    // 登录完成后，需要跳转到后台页面
    // location.assign('/users/eduardo')
    // vue router对象, 提供路由跳转的功能
    // 字符串路径
    //  router.push('/users/eduardo')
    // 指定到组件名称
    router.push({ name: 'HomePage' })
  }
}

</script>

<style lang="css" scoped>
.login {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  height: 100%;
  width: 100%;
}
</style>
