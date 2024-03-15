import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router';
import Overview from '../views/Overview.vue';
import { pb } from '@/api-client';

declare module 'vue-router' {
  interface RouteMeta {
    // original RouteMeta allows every prop Record<string ..., unknown>
    requiresAuth?: boolean;
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'overview',
    component: Overview,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/import',
    name: 'import',
    component: () => import('../views/Import.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/settings',
    name: 'settings',
    component: () => import('../views/Settings.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/login-signup',
    name: 'login-signup',
    component: () => import('../views/LoginSignup.vue'),
    meta: {
      requiresAuth: false
    }
  }
];

const protectedRoutes = routes.filter((r) => r.meta?.requiresAuth === true);

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

router.beforeEach((to) => {
  if (protectedRoutes.some((route) => route.name === to.name) && !pb.authStore.isValid) {
    return { name: 'login-signup' };
  }
});

export default router;
