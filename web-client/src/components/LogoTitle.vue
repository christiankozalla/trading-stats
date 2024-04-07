<script setup lang="ts">
import { useI18nStore } from '@/stores/i18n';
const i18n = useI18nStore();
</script>

<template>
  <RouterLink
    class="logo-container"
    :to="{ name: 'overview', params: { locale: i18n.currentLocale || 'en' } }"
  >
    <!-- the fallback locale 'en' above is necessary, because RouterLink throws an error if locale is undefined,
      and the i18n.currentLocale f is undefined at first when / is requested and changes later after the navigation guards have resolved -->
    <img class="logo-image" src="@/assets/inloopo-logo.png" alt="Inloopo logo" />
    <h1 class="logo-text">Trade&nbsp;Stats</h1>
  </RouterLink>
</template>

<style scoped>
.logo-container {
  display: flex;
  height: 100%;

  align-items: center;
  gap: var(--inline-spacing);
  margin-right: 24px;
  text-decoration: none;
  color: var(--text-color);
}

.logo-image {
  aspect-ratio: 1;
  width: 32px;
}

.logo-text {
  display: inline-block;
  font-size: large;
}

@media screen and (max-width: 320px) {
  .logo-text {
    display: none;
  }
}
</style>
