<template>
  <div class="login">
    <!-- form登录表单 -->
    <div :style="{ width: '480px', height: '400px' }">
      <div style="display: flex; height: 40px;justify-content: center;align-items:center;margin-bottom: 12px;">
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
          <a-input-password :invisible-button="false" v-model="form.password" placeholder="Please enter something"
            allow-clear>
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


// 表单对象
const form = ref({
  name: '',
  password: '',
  remember: false,
})

// 表单的提交
const handleSubmit = (data) => {
  console.log(data);
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
