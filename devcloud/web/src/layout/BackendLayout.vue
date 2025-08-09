<template>
  <div class="layout-container">
    <!-- 背景装饰元素 -->
    <div class="decoration-circle circle-1"></div>
    <div class="decoration-circle circle-2"></div>
    <div class="decoration-wave"></div>

    <!-- 固定顶部导航栏 -->
    <div class="fixed-header">
      <div class="header-content">
        <div class="logo-section">
          <h1 class="platform-name">研发交付平台</h1>
        </div>
        <div class="main-nav-section">
          <a-menu mode="horizontal" :default-selected-keys="['1']">
            <a-menu-item key="1">工作台</a-menu-item>
            <a-menu-item key="2">项目管理</a-menu-item>
            <a-menu-item key="3">研发交付</a-menu-item>
            <a-menu-item key="4">制品库</a-menu-item>
            <a-menu-item key="5">测试中心</a-menu-item>
            <a-menu-item key="6">运维中心</a-menu-item>
          </a-menu>
        </div>
        <div class="user-section">
          <a-dropdown position="bottom">
            <a-avatar :size="32" class="user-avatar">
              <icon-user />
            </a-avatar>
            <template #content>
              <a-doption>个人中心</a-doption>
              <a-doption>系统设置</a-doption>
              <a-doption>退出登录</a-doption>
            </template>
          </a-dropdown>
        </div>
      </div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content-wrapper">
      <!-- 可收缩侧边栏 -->
      <div class="fixed-sidebar" :class="{ 'collapsed': isSidebarCollapsed }">
        <!-- 使用原生a-menu确保折叠效果 -->
        <a-menu :theme="light" :default-selected-keys="['1']" :collapsed="isSidebarCollapsed" :collapsed-width="64"
          :style="{ width: '100%', height: '100%' }">
          <a-menu-item key="1">
            <template #icon><icon-dashboard /></template>
            <span class="menu-title">流水线列表</span>
          </a-menu-item>
          <a-menu-item key="2">
            <template #icon><icon-branch /></template>
            <span class="menu-title">分支管理</span>
          </a-menu-item>
          <a-menu-item key="3">
            <template #icon><icon-history /></template>
            <span class="menu-title">执行历史</span>
          </a-menu-item>
          <a-menu-item key="4">
            <template #icon><icon-settings /></template>
            <span class="menu-title">流水线模板</span>
          </a-menu-item>
          <a-menu-item key="5">
            <template #icon><icon-monitor /></template>
            <span class="menu-title">监控中心</span>
          </a-menu-item>
        </a-menu>

        <!-- 自定义收缩按钮 -->
        <div class="collapse-btn-wrapper">
          <a-button class="collapse-btn" @click="toggleSidebar">
            <template #icon>
              <icon-left v-if="!isSidebarCollapsed" />
              <icon-right v-else />
            </template>
          </a-button>
        </div>
      </div>

      <!-- 内容滚动区域（关键修改） -->
      <div class="scrollable-content" :class="{ 'expanded': isSidebarCollapsed }">
        <div class="content-scroll-wrapper">
          <!-- 面包屑（会随内容滚动） -->
          <div class="breadcrumb">
            <a-breadcrumb>
              <a-breadcrumb-item>首页</a-breadcrumb-item>
              <a-breadcrumb-item>流水线系统</a-breadcrumb-item>
              <a-breadcrumb-item>流水线列表</a-breadcrumb-item>
            </a-breadcrumb>
          </div>

          <!-- 主内容区 -->
          <main class="router-view-wrapper">
            <router-view />
          </main>

          <!-- 页脚（会随内容滚动） -->
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

const isSidebarCollapsed = ref(false);

const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value;
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

/* 顶部导航 */
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
  display: flex;
  padding-top: 60px;
  height: calc(100vh - 60px);
}

/* 侧边栏 */
.fixed-sidebar {
  position: fixed;
  left: 0;
  top: 60px;
  padding: 12px 8px;
  bottom: 0;
  width: 220px;
  background-color: var(--color-bg-2);
  z-index: 90;
  transition: width 0.3s ease;

  &.collapsed {
    width: 64px;
  }

  .sidebar-header {
    padding: 16px;
    background-color: var(--color-primary-light-1);
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;

    h3 {
      color: var(--color-text-1);
      font-size: 16px;
      font-weight: 500;
      margin: 0;
      white-space: nowrap;
    }
  }

  /* 强制修正折叠菜单样式 */
  :deep(.arco-menu) {
    height: 100%;
    border-right: none;

    &.arco-menu-collapsed {
      .arco-menu-item {
        justify-content: center !important;
        padding: 0 12px !important;

        .arco-menu-icon {
          margin: 0 auto !important;
        }

        .arco-menu-title {
          display: inline-block;
          width: 0;
          overflow: hidden;
        }
      }
    }
  }
}

/* 自定义收缩按钮 */
.collapse-btn-wrapper {
  position: absolute;
  right: -12px;
  top: 50%;
  transform: translateY(-50%);
  z-index: 110;
  width: 24px;
  height: 32px;

  .collapse-btn {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #fff;
    border: 1px solid var(--color-border-2);
    border-radius: 4px;
    padding: 0;
    transition: all 0.2s;
    cursor: pointer;

    // &:hover {
    //   background: var(--color-fill-2);
    //   border-color: var(--color-border-3);
    //   box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
    // }

    &> :deep(svg) {
      font-size: 14px;
      color: var(--color-text-2);
    }
  }
}

/* 内容滚动区域（关键修改） */
.scrollable-content {
  flex: 1;
  margin-left: 220px;
  height: 100%;
  overflow-y: scroll;
  scrollbar-width: none;
  /* Firefox */
  -ms-overflow-style: none;
  /* IE/Edge */
  transition: margin-left 0.3s ease;

  &::-webkit-scrollbar {
    display: none;
    /* Chrome/Safari */
  }

  &.expanded {
    margin-left: 64px;
  }

  .content-scroll-wrapper {
    min-height: 100%;
    display: flex;
    flex-direction: column;
  }
}

/* 面包屑 */
.breadcrumb {
  padding: 16px 24px 8px;
  flex-shrink: 0;
}

/* 路由视图容器 */
.router-view-wrapper {
  flex: 1;
  padding: 0 24px;
  min-height: calc(100vh - 180px);
  /* 动态计算最小高度 */
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
