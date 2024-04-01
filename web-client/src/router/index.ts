import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from 'vue-router';
import Overview from '../views/Overview.vue';
import { pb } from '@/api-client';
import { useI18nStore, SUPPORTED_LOCALES, type Locale } from '@/stores/i18n';

declare module 'vue-router' {
  interface RouteMeta {
    // original RouteMeta allows every prop Record<string ..., unknown>
    requiresAuth?: boolean;
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/:locale',
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

function getSupportedBrowserLocale(): Locale | undefined {
  const browserLocale = navigator.language.split('-')[0];
  if (isSupportedLocale(browserLocale)) {
    return browserLocale;
  }
}

function isSupportedLocale(locale: string): locale is Locale {
  return (SUPPORTED_LOCALES as readonly string[]).includes(locale);
}

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
    const supportedBrowserLocale = getSupportedBrowserLocale();
    if (supportedBrowserLocale) {
      // Redirect to the same route after setting the supported browser locale
      to.params.locale = supportedBrowserLocale;
    } else {
      // If the browser locale is not supported, redirect to the default locale
      to.params.locale = 'en';
    }
    next(to);
  } else if (isSupportedLocale(localeParam)) {
    if (!i18n.currentLocale || localeParam !== i18n.currentLocale) {
      await i18n.setLocale(localeParam);
    }
    next();
  }
});

router.beforeEach((to, _from, next) => {
  if (protectedRoutes.some((route) => route.name === to.name) && !pb.authStore.isValid) {
    return next({ name: 'login-signup', params: { ...to.params, locale: 'en' } });
  } else next();
});

export default router;
