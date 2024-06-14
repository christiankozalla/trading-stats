<script setup lang="ts">
import { watch } from 'vue';
import { RouterView, useRouter, useRoute } from 'vue-router';
import AppHeader from '@/components/AppHeader.vue';
import Loader from '@/components/Loader.vue';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { pb } from '@/api-client';
import { useScreenshotViewer } from './composables/useScreenshotViewer';
import { isAuthenticated } from '@/api-client';
import { supportedOrFallbackLocale } from './router/helpers';

const router = useRouter();
const route = useRoute();

// this works, but overrides any 404s if the user is un-authenticated
watch(
  isAuthenticated,
  (newAuthStatus) => {
    if (newAuthStatus === false && route.meta.requiresAuth) {
      const chosenLocale = supportedOrFallbackLocale(route.params.locale);
      router.push({ name: 'login-signup', params: { ...route.params, locale: chosenLocale } });
    }
  },
  { immediate: true }
);

const tradingAccountsStore = useTradingAccountsStore();

pb.beforeSend = function (originalUrl, options) {
  if (tradingAccountsStore.selected) {
    const url = new URL(originalUrl);
    url.searchParams.set('accountId', tradingAccountsStore.selected);
    return { url: url.toString(), options };
  } else {
    return { url: originalUrl };
  }
};

const { viewer, screenshotRecords } = useScreenshotViewer();
</script>

<template>
  <main>
    <AppHeader />
    <Loader />

    <RouterView :key="tradingAccountsStore.selected" />
    <Toast position="bottom-center" />
  </main>
  <ScreenshotViewer
    :isOpen="viewer.isOpen"
    :screenshotRecords="screenshotRecords"
    :activeRecord="viewer.activeRecord"
    @update:visible="(state: boolean) => viewer.setIsOpen(state)"
  />
</template>

<style scoped></style>
