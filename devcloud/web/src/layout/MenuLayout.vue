<template>
  <div class="layout-container">
    <!-- 背景装饰元素 -->
    <div class="decoration-circle circle-1"></div>
    <div class="decoration-circle circle-2"></div>
    <div class="decoration-wave"></div>

    <!-- 使用新的顶部导航组件 -->
    <HeaderNav @system-change="handleSystemChange" @user-option-click="handleUserOption" />

    <!-- 主内容区 -->
    <div class="main-content-wrapper">
      <!-- 可收缩侧边栏 -->
      <div class="fixed-sidebar" :class="{ 'collapsed': isSidebarCollapsed }">
        <!-- 使用原生a-menu确保折叠效果 -->
        <a-menu theme="light" :selected-keys="[currentSelectMenuItem]" :collapsed="isSidebarCollapsed"
          :collapsed-width="64" @sub-menu-click="handleMenuClick" @menu-item-click="handleMenuItemClick" breakpoint="xl"
          :style="{ width: '100%', height: '100%' }">
          <a-menu-item v-for="menu in currentMenus" :key="menu.key">
            <template #icon>
              <component :is="menu.icon"></component>
            </template>
            <span class="menu-title">{{ menu.title }}</span>
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
              <a-breadcrumb-item v-for="m in $route.matched" :key="m.name">{{ m.meta.title }}</a-breadcrumb-item>
            </a-breadcrumb>
          </div>

          <!-- 主内容区 -->
          <a-watermark :content="token.user_name" :font="{ color: 'rgba(0, 0, 0, 0.06)' }">
            <main class="router-view-wrapper">
              <router-view />
            </main>
          </a-watermark>

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
import { computed, ref } from 'vue';
import HeaderNav from './components/HeaderNav.vue';
import { useRoute, useRouter } from 'vue-router';
import { useWindowSize } from '@vueuse/core'
import token from '@/storage/token'
import app from '@/storage/app'
import { watch } from 'vue';

const router = useRouter()
const route = useRoute()

const isSidebarCollapsed = ref(false);

const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value;
};

// 侧边栏响应式
const { width } = useWindowSize()
watch(width, (newWidth) => {
  isSidebarCollapsed.value = newWidth < 1200; // 小于1200px时自动折叠
}, { immediate: true });

const currentSelectMenuItem = computed(() => {
  return app.value.system_menu.find(item => item.key === app.value.current_system)?.current_menu_item || app.value.current_system
})

const handleSystemChange = (system) => {
  // 这里可以添加路由跳转逻辑
  console.log('菜单切换:', system);

  // 跳转到当前选中的页面
  try {
    router.push({ name: currentSelectMenuItem.value })
  } catch (error) {
    console.log("swich system error, %s", error)
    router.push({ name: app.value.current_system })
  }
};

// 菜单被选中
const handleMenuClick = (key) => {
  app.value.system_menu.find(item => item.key === app.value.current_system).current_menu_item = key
  console.log(key)
}

// 菜单项被选中
const handleMenuItemClick = (key) => {
  app.value.system_menu.find(item => item.key === app.value.current_system).current_menu_item = key
  router.push({ name: key })
  console.log(route)
}

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

const currentMenus = computed(() => {
  return app.value.system_menu.find(item => item.key === app.value.current_system)?.menus || []
})

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
  border-bottom: 1px solid var(--color-border);
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
  border-right: 1px solid var(--color-border);
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
  padding: 0px 20px 0px 40px;
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
