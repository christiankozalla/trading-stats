<script setup lang="ts">
import { ref } from 'vue';
import type { MenuItem } from 'primevue/menuitem';
import LogoTitle from '@/components/LogoTitle.vue';
import { pb } from '@/api-client';
import { useAuthStore } from '@/stores/auth';

const menu = ref();
const authStore = useAuthStore();

const items = ref<MenuItem[]>([
  {
    separator: true
  },
  {
    label: 'Home',
    url: '/',
    icon: 'icon icon-house'
  },
  {
    label: 'Settings',
    url: '/settings',
    icon: 'icon icon-gear'
  },
  {
    label: 'Logout',
    icon: 'icon icon-logout',
    command: () => {
      pb.authStore.clear();
    }
  }
]);

function toggleMenu(event: MouseEvent) {
  menu.value.toggle(event);
}
</script>

<template>
  <header>
    <LogoTitle />
    <!-- <div v-if="userStore.isAuthenticated"> -->
    <div v-if="authStore.isAuthenticated">
      <Button
        @click="toggleMenu"
        :pt="{ root: { style: { 'aspect-ratio': 1, padding: '8px' } } }"
        icon="icon icon-person"
        outlined
        aria-haspopup="true"
        aria-label="Access user profile and settings"
        aria-controls="overlay_menu"
      />
      <Menu ref="menu" id="overlay_menu" :model="items" :popup="true">
        <template #start>
          <!-- <p class="font-bold">{{ userStore.record?.name }}</p> -->
          <!-- <p v-if="userStore.record?.email">{{ userStore.record?.email }}</p> -->
          <div v-if="pb.authStore.model?.name" class="font-bold menuitem-padding">
            {{ pb.authStore.model?.name }}
          </div>
          <div v-if="pb.authStore.model?.email" class="menuitem-padding">
            {{ pb.authStore.model?.email }}
          </div>
        </template>
      </Menu>
    </div>
  </header>
</template>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 42px;
  margin-left: var(--inline-spacing);
  margin-right: var(--inline-spacing);
}

.menuitem-padding {
  padding: 0.5rem 0.75rem;
}
</style>
