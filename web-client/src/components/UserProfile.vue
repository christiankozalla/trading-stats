<script setup lang="ts">
import { pb, isAuthenticated } from '@/api-client';
</script>

<template>
  <div v-if="isAuthenticated && pb.authStore.model">
    <ul>
      <li>Name: {{ pb.authStore.model.name }}</li>
      <li>Username: {{ pb.authStore.model.username }}</li>
      <li>Avatar: {{ pb.authStore.model.avatar }}</li>
      <li>Created: {{ pb.authStore.model.created }}</li>
      <li>Email: {{ pb.authStore.model.email }}</li>
      <li>Verified: {{ pb.authStore.model.verified }}</li>
    </ul>
  </div>
  <div v-else>
    Not authenticated
  </div>
</template>
