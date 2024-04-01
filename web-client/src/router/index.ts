import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router';
import Overview from '../views/Overview.vue';
import { pb } from '@/api-client';
import { useI18nStore } from '@/stores/i18n';

declare module 'vue-router' {
  interface RouteMeta {
    // original RouteMeta allows every prop Record<string ..., unknown>
    requiresAuth?: boolean;
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/:locale/',
    name: 'overview',
    component: Overview,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/:locale/import',
    name: 'import',
    component: () => import('../views/Import.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/:locale/settings',
    name: 'settings',
    component: () => import('../views/Settings.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/:locale/login-signup',
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

router.beforeEach(async (to) => {
  const i18n = useI18nStore();

  const currrentLocaleFromPath = to.path.slice(1, 3);
  if (!i18n.currentLocale || currrentLocaleFromPath !== i18n.currentLocale) {
    // TODO: what happens if another locale is present in path? -> 404 before requesting locale
    return await i18n.setLocale(currrentLocaleFromPath as 'en' | 'de');
  }
});

router.beforeEach((to) => {
  if (protectedRoutes.some((route) => route.name === to.name) && !pb.authStore.isValid) {
    return { name: 'login-signup' };
  }
});

export default router;
