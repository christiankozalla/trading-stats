<script setup lang="ts">
import { ref } from 'vue';
import { pb, type TradingAccount } from '@/api-client';
import { useI18nStore } from '@/stores/i18n';
import { useRoute } from 'vue-router';
import ToggleSwitch from 'primevue/toggleswitch';
import { ClientResponseError } from 'pocketbase';

const { t } = useI18nStore();
const route = useRoute();

const props = defineProps<{
  account: TradingAccount;
}>();

const publicDashboardPermissions = await pb
  .collection('public_dashboard_permissions')
  .getFirstListItem(pb.filter(`account = {:accountId}`, { accountId: props.account.id }));

const logFiles = await pb
  .collection('trade_log_files')
  .getFullList({ filter: pb.filter('account = {:accountId}', { accountId: props.account.id }) });

const localTradesTablePublic = ref(publicDashboardPermissions.is_trades_table_public);

async function toggleDashboardPermission(event: Event) {
  try {
    const response = await pb
      .collection('public_dashboard_permissions')
      .update(publicDashboardPermissions.id, {
        is_trades_table_public: (event.target as HTMLInputElement).checked
      });

    publicDashboardPermissions.is_trades_table_public = response.is_trades_table_public;
    localTradesTablePublic.value = response.is_trades_table_public;
    // TODO: make useTradingAccountsStore revalidate
    // or make the store reactive - like returning the ref and being able to update it
  } catch (err) {
    if (err instanceof ClientResponseError) {
      console.error('Error updating public dashboard permissions', err.data);
    }
  }
}
</script>

<template>
  <h3>{{ props.account.name }}</h3>
  <small>
    {{
      t('settings.accounts.created-date', {
        date: Intl.DateTimeFormat(route.params.locale).format(new Date(props.account.created))
      })
    }}
  </small>
  <h4>Public Dashboard</h4>
  <p v-if="localTradesTablePublic">Your Dashboard is public.</p>
  <p v-else>Your Dashboard is private.</p>
  <p class="flex justify-between align-center">
    <span>Click to make your Dashboard {{ localTradesTablePublic ? 'private' : 'public' }}</span
    ><ToggleSwitch v-model="localTradesTablePublic" @change="toggleDashboardPermission" />
  </p>
  <h4>Uploaded Log-Files</h4>
  <ul>
    <li v-for="logFile in logFiles" :key="logFile.id">
      <p>
        {{
          t('settings.accounts.log-file-created-date', {
            date: Intl.DateTimeFormat(route.params.locale).format(new Date(logFile.created))
          })
        }}
      </p>
      <div v-for="filename in logFile.file" :key="filename" style="margin: 12px;">
        <h5 style="word-break: break-all; margin: 0">&#x2022; {{ filename }}</h5>
      </div>
    </li>
  </ul>
</template>

<style scoped>
h3 {
  margin: 0 0 1.2rem 0;
}

h4 {
  margin: 2rem 0 0.7rem 0;
}

ul {
  padding: 0;
}

li {
  list-style: none;
}
</style>
