<script setup lang="ts">
import { ref, computed, inject } from 'vue';
import { pb, type TradingAccount, type TradeLogFileRecord } from '@/api-client';
import { useI18nStore } from '@/stores/i18n';
import { useRoute } from 'vue-router';
import DataPanel from './DataPanel.vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import ToggleSwitch from 'primevue/toggleswitch';
import { useToast } from 'primevue/usetoast';
import { ClientResponseError } from 'pocketbase';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const { t } = useI18nStore();
const tradingAccountsStore = useTradingAccountsStore();
const toast = useToast();
const route = useRoute();

const props = defineProps<{
  account: TradingAccount;
}>();

const accountName = ref(props.account.name);

const publicDashboardPermissions = await pb
  .collection('public_dashboard_permissions')
  .getFirstListItem(pb.filter(`account = {:accountId}`, { accountId: props.account.id }));

const rawLogFiles = await pb
  .collection('trade_log_files')
  .getFullList({ filter: pb.filter('account = {:accountId}', { accountId: props.account.id }) });

const logFiles = await Promise.all(
  rawLogFiles.map(async (record) => {
    const logFilesMeta = await pb.collection('raw_trades_count').getOne(record.id);
    record.tradeCount = logFilesMeta.trade_count;
    console.log('updated record', record);
    return record as TradeLogFileRecord & { tradeCount: number };
  })
);

const localTradesTablePublic = ref(publicDashboardPermissions.is_trades_table_public);
const status = computed(() =>
  localTradesTablePublic.value
    ? t('public-dashboard.settings.status-public')
    : t('public-dashboard.settings.status-private')
);

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

async function updateAccountName(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);

  try {
    await pb.collection('trading_accounts').update(props.account.id, formData);
    await tradingAccountsStore.revalidate();
    toast.add({
      severity: 'success',
      summary: t('settings.accounts.update-name-success'),
      life: 5000
    });
  } catch (err) {
    if (err instanceof ClientResponseError) {
      console.error('Error updating trading account name', err.data);
    }
  }
}
</script>

<template>
  <h3>
    <form @submit.prevent="updateAccountName">
      <label for="accountName" class="screen-reader-only">{{ t('settings.accounts.name') }}</label
      ><InputText id="accountName" v-model="accountName" name="name" />
      <Button type="submit" :label="t('settings.accounts.update-name')" class="ml-1" />
    </form>
  </h3>
  <small class="block mb-2">
    {{
      t('settings.accounts.created-date', {
        date: Intl.DateTimeFormat(route.params.locale).format(new Date(props.account.created))
      })
    }}
  </small>
  <DataPanel header="Public Dashboard" class="panel">
    <template #left>
      <p
        v-html="
          t('public-dashboard.settings.description', {
            host: inject('host')!,
            locale: route.params.locale as string,
            accountId: props.account.id
          })
        "
      ></p>
    </template>
    <template #right>
      <div>
        <p class="flex justify-between align-center">
          <strong
            ><em>Status: {{ status }}</em></strong
          >
          <ToggleSwitch v-model="localTradesTablePublic" @change="toggleDashboardPermission" />
        </p>
      </div>
    </template>
  </DataPanel>
  <DataPanel :header="t('settings.accounts.log-file-headline')" class="panel">
    <ul v-if="logFiles.length">
      <li v-for="logFile in logFiles" :key="logFile.id">
        <p>
          {{
            t('settings.accounts.log-file-created-date', {
              date: Intl.DateTimeFormat(route.params.locale).format(new Date(logFile.created))
            })
          }}
        </p>
        <small>{{ t('settings.accounts.log-file-count', { count: logFile.tradeCount }) }}</small>
        <div v-for="filename in logFile.file" :key="filename" style="margin: 12px">
          <h5 style="word-break: break-all; margin: 0">&#x2022; {{ filename }}</h5>
        </div>
      </li>
    </ul>
    <p v-else>{{ t('settings.accounts.no-log-files') }}</p>
  </DataPanel>
</template>

<style scoped>
h3 {
  margin: var(--inline-spacing) 0 1.2rem 0;
}

ul {
  padding: 0;
}

li {
  list-style: none;
}

:deep(.i18n-link) {
  padding: 0 !important;
}
</style>
