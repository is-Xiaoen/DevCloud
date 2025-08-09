<template>
  <div class="login-container">
    <!-- 背景装饰元素 -->
    <div class="decoration-circle circle-1"></div>
    <div class="decoration-circle circle-2"></div>
    <div class="decoration-wave"></div>

    <!-- 登录卡片 -->
    <div class="login-card">
      <!-- 平台Logo和标题 -->
      <div class="login-header">
        <h2 class="platform-name">研发交付平台</h2>
        <p class="welcome-text">持续集成 · 持续交付 · 高效协作</p>
      </div>

      <!-- 登录表单 -->
      <a-form class="login-form" :model="form" size="large" @submit="handleSubmit">
        <a-form-item hide-label field="parameter.username" :rules="{ required: true, message: '请输入用户名' }">
          <a-input v-model="form.parameter.username" placeholder="请输入用户名" allow-clear>
            <template #prefix>
              <icon-user style="color: var(--color-text-3)" />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item hide-label field="parameter.password" :rules="{ required: true, message: '请输入密码' }">
          <a-input-password :invisible-button="false" v-model="form.parameter.password" placeholder="请输入密码" allow-clear>
            <template #prefix>
              <icon-lock style="color: var(--color-text-3)" />
            </template>
          </a-input-password>
        </a-form-item>

        <div class="form-actions">
          <a-form-item hide-label field="remember">
            <a-checkbox v-model="form.remember">记住我</a-checkbox>
          </a-form-item>
          <a-link style="width: 80px;">忘记密码?</a-link>
        </div>

        <a-form-item hide-label>
          <a-button type="primary" long html-type="submit" class="login-btn">登 录</a-button>
        </a-form-item>

        <div class="other-login-methods">
          <a-divider>其他登录方式</a-divider>
          <div class="oauth-icons">
            <a-tooltip content="GitLab登录">
              <a-button shape="circle" class="oauth-btn">
                <icon-gitlab />
              </a-button>
            </a-tooltip>
            <a-tooltip content="GitHub登录">
              <a-button shape="circle" class="oauth-btn">
                <icon-github />
              </a-button>
            </a-tooltip>
            <a-tooltip content="LDAP登录">
              <a-button shape="circle" class="oauth-btn">
                <icon-user-group />
              </a-button>
            </a-tooltip>
          </div>
        </div>
      </a-form>
    </div>

    <!-- 页脚 -->
    <div class="login-footer">
      <p>© 2025 研发交付平台 · 让软件交付更高效</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import mcenter from '@/api/mcenter'
import token from '@/storage/token'
import { useRouter } from 'vue-router';

const form = ref({
  issuer: 'password',
  parameter: {
    username: '',
    password: ''
  },
  remember: false,
});

const router = useRouter();

const handleSubmit = async (data) => {
  if (data.errors == null) {
    const resp = await mcenter.Login(data.values);
    token.value = resp;
    router.push({ name: 'HomePage' });
  }
};
</script>

<style lang="less" scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  width: 100%;
  background-color: #f5f7fa;
  position: relative;
  overflow: hidden;

  .decoration-circle {
    position: absolute;
    border-radius: 50%;
    background: linear-gradient(135deg, rgba(24, 144, 255, 0.1) 0%, rgba(24, 144, 255, 0.05) 100%);
    z-index: 0;

    &.circle-1 {
      width: 600px;
      height: 600px;
      top: -200px;
      left: -200px;
    }

    &.circle-2 {
      width: 400px;
      height: 400px;
      bottom: -100px;
      right: -100px;
    }
  }

  .decoration-wave {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 100px;
    background: url('@/assets/svg/wave.svg') repeat-x;
    background-size: 1000px 100px;
    opacity: 0.1;
  }

  .login-card {
    width: 420px;
    padding: 40px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
    position: relative;
    z-index: 1;

    .login-header {
      text-align: center;
      margin-bottom: 30px;

      .logo {
        width: 60px;
        height: 60px;
        margin-bottom: 16px;
      }

      .platform-name {
        color: var(--color-text-1);
        margin-bottom: 8px;
        font-weight: 600;
      }

      .welcome-text {
        color: var(--color-text-3);
        font-size: 14px;
        margin: 0;
      }
    }

    .login-form {
      .form-actions {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16px;
      }

      .login-btn {
        height: 44px;
        font-size: 16px;
        border-radius: 4px;
        margin-top: 8px;
      }

      .other-login-methods {
        margin-top: 24px;

        .oauth-icons {
          display: flex;
          justify-content: center;
          gap: 16px;
          margin-top: 16px;

          .oauth-btn {
            color: var(--color-text-2);
            background-color: var(--color-fill-2);
            border: none;

            &:hover {
              background-color: var(--color-fill-3);
            }
          }
        }
      }
    }
  }

  .login-footer {
    position: absolute;
    bottom: 20px;
    width: 100%;
    text-align: center;
    color: var(--color-text-3);
    font-size: 12px;
  }
}
</style>
