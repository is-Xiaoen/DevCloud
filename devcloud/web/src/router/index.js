import { createRouter, createWebHistory } from 'vue-router'

import MenuLayout from '@/layout/MenuLayout.vue'
import DashboardLayout from '@/layout/DashboardLayout.vue'
import FrontendLayout from '@/layout/FrontendLayout.vue'
import { beforeEach } from './interceptor'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Frontend',
      redirect: { name: 'ProductPage' },
      component: () => FrontendLayout,
      children: [
        {
          path: 'product',
          name: 'ProductPage',
          component: () => import('@/pages/frontend/ProductPage.vue'),
        },
      ],
    },
    {
      path: '/login',
      name: 'LoginPage',
      component: () => import('@/pages/LoginPage.vue'),
    },
    {
      path: '/person',
      name: 'DashBoard',
      redirect: { name: 'DashboardPage' },
      component: DashboardLayout,
      children: [
        {
          path: 'dashboard',
          name: 'DashboardPage',
          component: () => import('@/pages/dashboard/DashboardPage.vue'),
        },
      ],
    },
    {
      path: '/project',
      name: 'ProjectSystem',
      redirect: { name: 'AppPage' },
      component: MenuLayout,
      children: [
        {
          path: 'app',
          name: 'AppPage',
          component: () => import('@/pages/project/AppPage.vue'),
        },
      ],
    },
    {
      path: '/develop',
      name: 'DevelopSystem',
      redirect: { name: 'SprintPage' },
      component: MenuLayout,
      children: [
        {
          path: 'sprint',
          name: 'SprintPage',
          component: () => import('@/pages/develop/SprintPage.vue'),
        },
      ],
    },
  ],
})

// 导航守卫
router.beforeEach(beforeEach)

export default router
