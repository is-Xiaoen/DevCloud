import { createRouter, createWebHistory } from 'vue-router'

import BackendLayout from '@/layout/BackendLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'HomePage',
      component: BackendLayout,
    },
    {
      path: '/login',
      name: 'LoginPage',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../pages/LoginPage.vue'),
    },
    {
      path: '/cmdb',
      component: () => import('@/layout/BackendLayout.vue'),
      children: [
        {
          path: 'secret',
          name: 'SecretPage',
          component: () => import('@/pages/cmdb/SecretPage.vue'),
        },
      ],
    },
  ],
})

export default router
