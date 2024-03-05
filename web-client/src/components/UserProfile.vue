<script setup lang="ts">
import { pb } from '@/api-client';
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();

console.log('user profile api client isValid', pb.authStore.isValid);
console.log('user profile api client isAuthRecord', pb.authStore.isAuthRecord);
console.log('user profile api client record/record', pb.authStore.model);
</script>

<template>
  <div v-if="authStore.isAuthenticated">
    {{ JSON.stringify(authStore.model) }}
  </div>
</template>
