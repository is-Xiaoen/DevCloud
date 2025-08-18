<template>
  <div class="fixed-header">
    <div class="header-content">
      <div class="logo-section">
        <h1 class="platform-name">研发交付平台</h1>
      </div>
      <div class="main-nav-section">
        <a-menu mode="horizontal" :default-selected-keys="[app.current_system]">
          <a-menu-item v-for="item in menuItems" :key="item.key" @click="handleMenuClick(item)">
            {{ item.label }}
          </a-menu-item>
        </a-menu>
      </div>
      <div class="user-section">
        <a-dropdown position="bottom">
          <a-avatar :size="32" class="user-avatar">
            <icon-user />
          </a-avatar>
          <template #content>
            <a-doption v-for="option in userOptions" :key="option.key" @click="option.handler">
              {{ option.label }}
            </a-doption>
          </template>
        </a-dropdown>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import token from '@/storage/token'
import app from '@/storage/app'
import { useRouter } from 'vue-router';


const router = useRouter()
const emit = defineEmits(['system-change', 'user-option-click']);

const menuItems = computed(() => {
  const systemMenu = app.value.system_menu || []
  return systemMenu.map(item => ({
    key: item.key,
    label: item.title,
  }))
})

const userOptions = ref([
  { key: 'profile', label: '个人中心', handler: () => emit('user-option-click', 'profile') },
  { key: 'settings', label: '系统设置', handler: () => emit('user-option-click', 'settings') },
  {
    key: 'logout', label: '退出登录', handler: () => {
      // 登录的状态去掉
      token.value = null
      router.push({ name: 'LoginPage' })
    }
  }
]);

const handleMenuClick = (menuItem) => {
  app.value.current_system = menuItem.key
  emit('system-change', menuItem.key);
};
</script>

<style lang="less" scoped>
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
</style>
