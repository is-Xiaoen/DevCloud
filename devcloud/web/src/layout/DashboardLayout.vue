<template>
  <div class="layout-container">
    <!-- 背景装饰元素 -->
    <div class="decoration-circle circle-1"></div>
    <div class="decoration-circle circle-2"></div>
    <div class="decoration-wave"></div>

    <!-- 使用顶部导航组件 -->
    <HeaderNav :active-key="activeMenuKey" @menu-change="handleMenuChange" @user-option-click="handleUserOption" />

    <!-- 主内容区 -->
    <div class="main-content-wrapper">
      <!-- 内容滚动区域 -->
      <div class="scrollable-content">
        <div class="content-scroll-wrapper">
          <!-- 主内容区 -->
          <a-watermark :content="token.user_name" :font="{ color: 'rgba(0, 0, 0, 0.06)' }">
            <main class="router-view-wrapper">
              <router-view />
            </main>
          </a-watermark>

          <!-- 页脚 -->
          <footer class="layout-footer">
            <p>© 2025 研发交付平台 · 让软件交付更高效</p>
          </footer>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import HeaderNav from './components/HeaderNav.vue';
import { useRouter } from 'vue-router';
import token from '@/storage/token'

const router = useRouter()

const activeMenuKey = ref('DashBoard');

const handleMenuChange = (key) => {
  activeMenuKey.value = key;
  router.push({ name: key })
  // 这里可以添加路由跳转逻辑
  console.log('菜单切换:', key);
};

const handleUserOption = (option) => {
  console.log('用户操作:', option);
  // 根据不同的option执行不同的操作
  switch (option) {
    case 'profile':
      // 跳转到个人中心
      break;
    case 'settings':
      // 跳转到系统设置
      break;
    case 'logout':
      // 执行退出登录
      break;
  }
};
</script>

<style lang="less" scoped>
/* 基础布局 */
.layout-container {
  min-height: 100vh;
  background-color: #f5f7fa;
  position: relative;
  overflow: hidden;

  .decoration-circle {
    position: fixed;
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
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 70px;
    background: url('@/assets/svg/wave.svg') repeat-x;
    background-size: 1000px 100px;
    opacity: 0.1;
  }
}

/* 主内容区 */
.main-content-wrapper {
  padding-top: 60px;
  height: calc(100vh - 60px);
}

/* 内容滚动区域 */
.scrollable-content {
  height: 100%;
  overflow-y: scroll;
  scrollbar-width: none;
  /* Firefox */
  -ms-overflow-style: none;
  /* IE/Edge */

  &::-webkit-scrollbar {
    display: none;
    /* Chrome/Safari */
  }

  .content-scroll-wrapper {
    min-height: 100%;
    display: flex;
    flex-direction: column;
  }
}

/* 路由视图容器 */
.router-view-wrapper {
  flex: 1;
  padding: 20px;
  min-height: calc(100vh - 180px);
  /* 调整最小高度 */
}

/* 页脚 */
.layout-footer {
  padding: 16px 24px;
  margin-top: auto;
  /* 关键：使页脚始终在底部 */
  text-align: center;
  color: var(--color-text-3);
  font-size: 12px;
  border-top: 1px solid var(--color-border-2);
  background: var(--color-bg-2);
}
</style>
