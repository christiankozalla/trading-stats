import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router';
import { pb } from '@/api-client';
import { useI18nStore } from '@/stores/i18n';
import { isSupportedLocale, supportedOrFallbackLocale } from './helpers';

declare module 'vue-router' {
  interface RouteMeta {
    // original RouteMeta allows every prop Record<string ..., unknown>
    requiresAuth?: boolean;
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/:locale/', // could use regex pattern like: '/:locale(de|en)'
    name: 'overview',
    component: () => import('../views/Overview.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/:locale/import',
    component: () => import('../views/Import.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/:locale/settings',
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
  },
  {
    path: '/:locale/public/:accountId/overview',
    component: () => import('../views/PublicOverview.vue'),
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/:locale/:pathMatch(.*)', // Catch all 404
    component: () => import('../views/404.vue'),
    meta: {
      requiresAuth: false
    }
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

router.beforeEach(async (to, _from, next) => {
  // if the user is not requesting a path with a locale
  // a.) get the users standard language, see whether we support it and send the user there
  // a.2) if we do not support the users favorite locale, send them to /en
  const i18n = useI18nStore();
  const localeParam = to.params.locale;

  if (Array.isArray(localeParam)) {
    throw new Error(`Multiple locale params in route: ${localeParam.join()}`);
  }

  if (!localeParam || !isSupportedLocale(localeParam)) {
    const chosenLocale = supportedOrFallbackLocale(localeParam);
    if (!i18n.currentLocale || chosenLocale !== i18n.currentLocale) {
      await i18n.setLocale(chosenLocale);
    }
    next(`/${chosenLocale}${to.fullPath}`);
  } else if (isSupportedLocale(localeParam)) {
    if (!i18n.currentLocale || localeParam !== i18n.currentLocale) {
      await i18n.setLocale(localeParam);
    }
    next();
  } else {
    console.log('should not happen!', localeParam);
    next();
  }
});

router.beforeEach((to, _from, next) => {
  if (Array.isArray(to.params.locale)) {
    throw new Error(`Multiple locale params in route: ${to.params.locale.join()}`);
  }
  if (to.meta.requiresAuth && !pb.authStore.isValid) {
    const locale = supportedOrFallbackLocale(to.params.locale);
    next(`/${locale}/login-signup`); // or next({ name: 'login-signup', params: { locale }})
  } else next();
});

export default router;
