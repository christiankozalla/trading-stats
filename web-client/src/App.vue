<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
import AppHeader from '@/components/AppHeader.vue';
import Loader from '@/components/Loader.vue';
import { useAuthStore } from '@/stores/auth';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { pb, type User } from '@/api-client';
import { onBeforeUnmount } from 'vue';

const router = useRouter();
const authStore = useAuthStore();
const tradingAccountsStore = useTradingAccountsStore();

const fireImmediately = true;
const removeAuthStoreOnChangeListener = pb.authStore.onChange((token, model) => {
  authStore.model = model as User;
  if (!authStore.isAuthenticated) {
    router.push('/login-signup');
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
