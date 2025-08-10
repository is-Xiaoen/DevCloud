<template>
  <div class="layout-container">
    <!-- 背景装饰元素 -->
    <div class="decoration-circle circle-1"></div>
    <div class="decoration-circle circle-2"></div>
    <div class="decoration-wave"></div>

    <!-- 使用顶部导航组件 -->
    <div class="fixed-header">
      <div class="header-content">
        <div class="logo-section">
          <h1 class="platform-name">研发交付平台</h1>
        </div>
        <div class="main-nav-section">
          <a-menu mode="horizontal" :default-selected-keys="[activeKey]">
            <a-menu-item v-for="item in menuItems" :key="item.key" @click="handleMenuClick(item)">
              {{ item.label }}
            </a-menu-item>
          </a-menu>
        </div>
        <div class="user-section">
          <a-button type="text" @click="() => $router.push({ name: 'LoginPage' })">登录</a-button>
        </div>
      </div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content-wrapper">
      <!-- 内容滚动区域 -->
      <div class="scrollable-content">
        <div class="content-scroll-wrapper">
          <!-- 主内容区 -->
          <main class="router-view-wrapper">
            <router-view />
          </main>

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

.fixed-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  z-index: 100;

  .header-content {
    max-width: 100%;
    height: 100%;
    padding: 0 24px;
    margin: 0 auto;
    display: flex;
    align-items: center;

    .logo-section {
      .platform-name {
        color: var(--color-text-1);
        font-size: 18px;
        font-weight: 600;
        margin: 0;
        white-space: nowrap;
      }
    }

    .main-nav-section {
      flex: 1;
      margin: 0 40px;
      overflow-x: auto;
      overflow-y: hidden;

      :deep(.arco-menu-horizontal) {
        height: 60px;
        background: transparent;
        border-bottom: none;
      }
    }

    .user-avatar {
      background-color: var(--color-primary-light-3);
      color: var(--color-primary);
      cursor: pointer;
      transition: all 0.2s;

      &:hover {
        background-color: var(--color-primary-light-2);
      }
    }
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
