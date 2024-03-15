<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
import AppHeader from '@/components/AppHeader.vue';
import Loader from '@/components/Loader.vue';
import { useAuthStore } from '@/stores/auth';
import { pb, type User } from '@/api-client';

const router = useRouter();
const authStore = useAuthStore();
const fireImmediately = true;
pb.authStore.onChange((token, model) => {
  authStore.model = model as User;
  if (!authStore.isAuthenticated) {
    router.push('/login-signup');
  }
}, fireImmediately);
</script>

<template>
  <main>
    <AppHeader />
    <Loader />

    <RouterView />
    <Toast position="bottom-center" />
  </main>
</template>

<style scoped></style>
