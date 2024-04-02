<script setup lang="ts">
import { RouterView, useRouter, useRoute } from 'vue-router';
import AppHeader from '@/components/AppHeader.vue';
import Loader from '@/components/Loader.vue';
import { useAuthStore } from '@/stores/auth';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { pb, type User } from '@/api-client';
import { onBeforeUnmount } from 'vue';
import { supportedOrFallbackLocale } from '@/router/helpers';

const router = useRouter();
const authStore = useAuthStore();
const tradingAccountsStore = useTradingAccountsStore();

// set to false for not, because if fires immediately then it always fires before navigation guards for root route /, even if /en/someting-else was requested
const fireImmediately = false;
const removeAuthStoreOnChangeListener = pb.authStore.onChange(async (token, model) => {
  authStore.model = model as User;
  if (!authStore.isAuthenticated) {
    await router.push({
      name: 'login-signup',
      params: { locale: supportedOrFallbackLocale(useRoute().params.locale) }
    });
  }
}, fireImmediately);

pb.beforeSend = function (originalUrl, options) {
  const url = new URL(originalUrl);

  url.searchParams.set('accountId', tradingAccountsStore.selected || '');

  return { url: url.toString(), options };
};

onBeforeUnmount(() => {
  removeAuthStoreOnChangeListener();
});
</script>

<template>
  <main>
    <AppHeader />
    <Loader />

    <RouterView :key="tradingAccountsStore.selected" />
    <Toast position="bottom-center" />
  </main>
</template>

<style scoped></style>
