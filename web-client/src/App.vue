<script setup lang="ts">
import { RouterView } from 'vue-router';
import AppHeader from '@/components/AppHeader.vue';
import Loader from '@/components/Loader.vue';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { pb } from '@/api-client';
import { useScreenshotViewer } from './composables/useScreenshotViewer';

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
