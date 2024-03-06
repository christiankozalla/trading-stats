<script setup lang="ts">
import { useLoaderStore } from '@/stores/loader';

const loaderStore = useLoaderStore();
</script>

<template>
  <div v-if="loaderStore.isLoading" class="loader">
    <ProgressSpinner
      style="width: 50px; height: 50px"
      strokeWidth="6"
      animationDuration=".8s"
      aria-label="the app is loading"
    />
  </div>
</template>

<style>
.loader {
  display: grid;
  place-items: center;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.05);
  z-index: 10;
}
</style>
