<script setup lang="ts">
import UserProfile from '@/components/UserProfile.vue';
</script>

<template>
  <div class="settings">
    <h1>Settings Page</h1>
    <UserProfile />
  </div>
</template>

<style scoped>
@media (min-width: 1024px) {
  .settings {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }
}
</style>
