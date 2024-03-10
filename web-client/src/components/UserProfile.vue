<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();
</script>

<template>
  <div v-if="authStore.isAuthenticated && authStore.model">
    <ul>
      <li>Name: {{ authStore.model.name }}</li>
      <li>Username: {{ authStore.model.username }}</li>
      <li>Avatar: {{ authStore.model.avatar }}</li>
      <li>Created: {{ authStore.model.created }}</li>
      <li>Email: {{ authStore.model.email }}</li>
      <li>Verified: {{ authStore.model.verified }}</li>
    </ul>
  </div>
</template>
