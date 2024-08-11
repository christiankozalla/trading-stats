<script setup lang="ts">
import { watch } from "vue";
import AppHeader from "@/components/AppHeader.vue";
import Loader from "@/components/Loader.vue";
import ScreenshotViewer from "./components/ScreenshotViewer.vue";
import { useTradingAccountsStore } from "@/stores/tradingAccounts";
import { RouterView, useRouter } from "vue-router";
import { useScreenshotViewer } from "./composables/useScreenshotViewer";
import { isAuthenticated } from "@/api-client";
import { supportedOrFallbackLocale } from "./router/helpers";

const router = useRouter();

// this works, but overrides any 404s if the user is un-authenticated
watch(isAuthenticated, async (newAuthStatus) => {
  await router.isReady();
  if (newAuthStatus === false && router.currentRoute.value.meta.requiresAuth === true) {
    const chosenLocale = supportedOrFallbackLocale(router.currentRoute.value.params.locale);
    router.push({
      name: "login-signup",
      params: { ...router.currentRoute.value.params, locale: chosenLocale }
    });
  }
});

const tradingAccountsStore = useTradingAccountsStore();
const { viewer, screenshotRecords } = useScreenshotViewer();
</script>

<template>
  <main>
    <AppHeader />
    <Loader />

    <RouterView :key="tradingAccountsStore.selected" />
    <Toast position="top-right" />
  </main>
  <ScreenshotViewer
    :isOpen="viewer.isOpen"
    :screenshotRecords="screenshotRecords"
    :activeRecord="viewer.activeRecord"
    @update:visible="(state: boolean) => viewer.setIsOpen(state)"
  />
</template>

<style scoped></style>
