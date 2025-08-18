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
    // 项目管理
    {
      path: '/project',
      name: 'ProjectSystem',
      redirect: { name: 'ProjectList' },
      component: MenuLayout,
      meta: {
        title: '项目管理',
      },
      children: [
        {
          path: 'list',
          name: 'ProjectList',
          component: () => import('@/pages/project/ProjectList.vue'),
          meta: {
            title: '项目空间',
          },
        },
      ],
    },
    // 资源管理
    {
      path: '/resource',
      name: 'ResourceSystem',
      redirect: { name: 'ResourceSearch' },
      component: MenuLayout,
      meta: {
        title: '资源管理',
      },
      children: [
        {
          path: 'search',
          name: 'ResourceSearch',
          component: () => import('@/pages/resource/ResourceSearch.vue'),
          meta: {
            title: '资源检索',
          },
        },
      ],
    },
    // 研发交付
    {
      path: '/develop',
      name: 'DevelopSystem',
      redirect: { name: 'AppPage' },
      component: MenuLayout,
      meta: {
        title: '研发交付',
      },
      children: [
        {
          path: 'app',
          name: 'AppPage',
          component: () => import('@/pages/develop/AppPage.vue'),
          meta: {
            title: '应用管理',
          },
        },
        {
          path: 'version_iteration',
          name: 'VersionIteration',
          component: () => import('@/pages/develop/VersionIteration.vue'),
          meta: {
            title: '版本迭代',
          },
        },
        {
          path: 'pipeline_template',
          name: 'PipelineTemplate',
          component: () => import('@/pages/develop/PipelineTemplate.vue'),
          meta: {
            title: '流水线模板',
          },
        },
      ],
    },
    // 制品库
    {
      path: '/artifact',
      name: 'ArtifactSystem',
      redirect: { name: 'RegistryPage' },
      component: MenuLayout,
      meta: {
        title: '制品库',
      },
      children: [
        {
          path: 'registry',
          name: 'RegistryPage',
          component: () => import('@/pages/artifact/RegistryPage.vue'),
          meta: {
            title: '制品仓库',
          },
        },
        {
          path: 'asset',
          name: 'AssetPage',
          component: () => import('@/pages/artifact/AssetPage.vue'),
          meta: {
            title: '制品管理',
          },
        },
      ],
    },
  ],
})

// 导航守卫
router.beforeEach(beforeEach)

export default router
