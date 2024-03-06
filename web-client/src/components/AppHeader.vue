<script setup lang="ts">
import { ref } from 'vue';
import type { MenuItem } from 'primevue/menuitem';
import LogoTitle from '@/components/LogoTitle.vue';
import { pb } from '@/api-client';
import { useAuthStore } from '@/stores/auth';

const menu = ref();
const authStore = useAuthStore();

const mainNav: MenuItem[] = [
  {
    label: 'Notes',
    url: '/notes',
    icon: ''
  },
  {
    label: 'Import/Export',
    url: '/import-export',
    icon: ''
  }
];

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
    <Menubar
      :model="mainNav"
      :pt="{
        root: {
          class: ['menubar']
        },
        menuitem: {
          style: {
            margin: 0,
            'margin-right': '12px',
            'margin-left': '12px'
          }
        }
      }"
    >
      <template #start>
        <LogoTitle />
      </template>
      <template #end>
        <div v-if="authStore.isAuthenticated">
          <Button
            @click="toggleMenu"
            :pt="{
              root: {
                class: ['menu-button background-color-hover']
              }
            }"
            icon="icon icon-person"
            text
            aria-haspopup="true"
            aria-label="Access user profile and settings"
            aria-controls="overlay_menu"
          />
          <Menu ref="menu" id="overlay_menu" :model="items" :popup="true">
            <template #start>
              <div v-if="pb.authStore.model?.name" class="font-bold menuitem-padding">
                {{ pb.authStore.model?.name }}
              </div>
              <div v-if="pb.authStore.model?.email" class="menuitem-padding">
                {{ pb.authStore.model?.email }}
              </div>
            </template>
          </Menu>
        </div>
      </template>
    </Menubar>
  </header>
</template>

<style scoped>
.menubar {
  border: none;
  border-bottom: 1px solid var(--surface-border);
  border-radius: 0;
}

.menuitem-padding {
  padding: 0.5rem 0.75rem;
}

.menu-button {
  aspect-ratio: 1;
  padding: 5px;
  border-radius: 50%;
  border: none;
}
.background-color-hover {
  transition: background-color 300ms ease;
  background-color: var(--primary-100);
}

.background-color-hover:hover {
  background-color: var(--primary-200);
}

.background-color-hover:active {
  background-color: var(--primary-300);
}
</style>
