<script setup lang="ts">
import { ref, computed } from 'vue';
import Button from 'primevue/button';
import Menu from 'primevue/menu';
import type { MenuItem } from 'primevue/menuitem';
import { RouterLink } from 'vue-router';
import LogoTitle from '@/components/LogoTitle.vue';
import { pb, isAuthenticated } from '@/api-client';
import TradingAccountSelector from '@/components/TradingAccountSelector.vue';
import { useI18nStore } from '@/stores/i18n';
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();
const menu = ref();
const i18n = useI18nStore();

const items = computed<MenuItem[]>(() => [
  {
    separator: true
  },
  {
    label: 'Home',
    url: `/${i18n.currentLocale}`,
    icon: 'icon icon-house'
  },
  {
    label: 'Import',
    url: `/${i18n.currentLocale}/import`,
    icon: 'icon icon-upload'
  },
  {
    label: i18n.currentLocale === 'de' ? 'English language' : 'Deutsche Sprache',
    command: changeLocale,
    icon: 'icon icon-language'
  },
  {
    label: 'Settings',
    url: `/${i18n.currentLocale}/settings`,
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

async function changeLocale() {
  const otherLocale = i18n.currentLocale === 'de' ? 'en' : 'de';
  const newPath = route.path.replace(/^(\/de|\/en)/, `/${otherLocale}`);
  await router.push(newPath);
}
</script>

<template>
  <header>
    <LogoTitle />
    <div v-if="isAuthenticated" class="auth-section">
      <TradingAccountSelector />
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
        <template #item="{ item, props }">
          <div v-if="item.url" class="p-menuitem-content" data-pc-section="content">
            <RouterLink :to="item.url as string" v-bind="props.action" class="p-menuitem-link">
              <span class="p-menuitem-icon" :class="item.icon" data-pc-section="icon" />
              <span class="p-menuitem-text" data-pc-section="label">{{ item.label }}</span>
            </RouterLink>
          </div>
          <div v-else class="p-menuitem-content" data-pc-section="content">
            <div v-bind="props.action">
              <span class="p-menuitem-icon" :class="item.icon" data-pc-section="icon" />
              <span class="p-menuitem-text" data-pc-section="label">{{ item.label }}</span>
            </div>
          </div>
        </template>
      </Menu>
    </div>
  </header>
</template>

<style scoped>
header {
  border: none;
  border-bottom: 1px solid var(--p-gray-200);
  border-radius: 0;
  padding: 4px 12px;

  display: flex;
  justify-content: space-between;
  align-items: center;
}

.auth-section {
  display: flex;
  align-items: center;
}

.menuitem-padding {
  padding: 0.5rem 0.75rem;
}

.menu-button {
  aspect-ratio: 1;
  padding: 5px;
  border-radius: 50%;
  border: none;
  margin-left: 12px;
}
.background-color-hover {
  transition: background-color 300ms ease;
  background-color: var(--p-primary-100);
}

.background-color-hover:hover {
  background-color: var(--p-primary-200);
}

.background-color-hover:active {
  background-color: var(--p-primary-300);
}
</style>
