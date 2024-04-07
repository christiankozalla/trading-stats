<script setup lang="ts">
import { RouterView } from 'vue-router';
import AppHeader from '@/components/AppHeader.vue';
import Loader from '@/components/Loader.vue';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { pb } from '@/api-client';

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
