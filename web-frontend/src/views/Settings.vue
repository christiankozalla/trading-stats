<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';
import UserProfile from '@/components/UserProfile.vue';
import TradingAccounts from '@/components/TradingAccounts.vue';

const route = useRoute();
const router = useRouter()

function updateRouteHash(newValue: string) {
  router.replace({ ...route, hash: `#${newValue}` });
}
</script>

<template>
  <div class="settings">
    <Tabs :value="route.hash.slice(1) || 'user'" @update:value="updateRouteHash">
      <TabList class="tablist">
        <Tab value="user">
          <i class="icon icon-person" />
          User
        </Tab>
        <Tab value="trading-accounts">
          <i class="icon icon-folder" />
          Trading Accounts
        </Tab>
      </TabList>
      <TabPanels>
        <TabPanel value="user">
          <UserProfile />
        </TabPanel>
        <TabPanel value="trading-accounts">
          <TradingAccounts />
        </TabPanel>
      </TabPanels>
    </Tabs>
  </div>
</template>

<style scoped>
:deep(.p-tablist-tab-list > *) {
  flex: 1 1 0;
  display: flex;
  align-items: center;
  gap: var(--inline-spacing);
  font-size: 15px;
}
</style>
