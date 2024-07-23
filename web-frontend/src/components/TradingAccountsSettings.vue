<script setup lang="ts">
import { ref } from 'vue';
import LoaderInline from './LoaderInline.vue';
import DataPanel from './DataPanel.vue';
import TradingAccountCreator from './TradingAccountCreator.vue';
import TradingAccountDetails from './TradingAccountDetails.vue';
import Button from 'primevue/button';
import Drawer from 'primevue/drawer';
import { pb, type TradingAccount } from '@/api-client';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { useI18nStore } from '@/stores/i18n';
import { useRoute } from 'vue-router';

const tradingAccountsStore = useTradingAccountsStore();
const route = useRoute();
const { t } = useI18nStore();
const loading = ref(false);
const drawerVisible = ref(false);
const selectedAccount = ref<TradingAccount | null>(null);

function openAccountDetailsDrawer(account: TradingAccount) {
  // there is a ref on the page that can have an account id (or a whole account object) or be null (empty)
  // then there is a drawer / sidebar primevue component, that opens if an account-id is in the ref
  // data needs not the be fetched, because it should already be available
  drawerVisible.value = true;
  selectedAccount.value = account;
}

// Deletes all records / files associated with this account
// Deletes account-related config aswell (to keep config, CLEAR an account)
// eslint-disable-next-line
async function deleteAccount(accountId: string) {
  try {
    await pb.collection('trading_accounts').delete(accountId);
  } catch (err) {
    console.error('deleteAccount', err);
  }
}

async function clearAccount(accountId: string) {
  // delete trade_log_files where accoundId === account
  // delete raw_trades where ...
  // delete trades where ...
  // delete public_dashboard_permissions ...

  // PocketBase JS SDK only provides a method to delete a single record by id with a single request (not in batches)
  // solution: use a custom endpoint with auth-guard that takes an account id and runs raw sql or uses the Dao()

  // better solution: delete trade_log_files one by one and use an ON DELETE CASCADE for connected records in raw_trades, trades
  try {
    loading.value = true;
    const allLogFiles = await pb
      .collection('trade_log_files')
      .getFullList({ filter: pb.filter('account = {:accountId}', { accountId }) });
    console.log('total # log files', allLogFiles.length, allLogFiles);
    for (const { id } of allLogFiles) {
      await pb.collection('trade_log_files').delete(id);
    }
  } catch (err) {
    console.error('clearAccount', err);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <h2>Trading Accounts</h2>
  <!-- List of all trading accounts -> Name (fat) + Created date + Link that opens right
   a drawer with  "Quick Info" (like Commissions, Fees),
   Account Details (List of Log-Files and Trades derived from them),
   Public Dashboard Switch, 
   "Export Data" Btn, Clear Account Btn, Delete Account Btn) -->
  <ul>
    <li v-for="tradingAcc in tradingAccountsStore.accounts || []" :key="tradingAcc.id">
      <DataPanel>
        <div>
          <h4>{{ tradingAcc.name }}</h4>
          <small>
            {{
              t('settings.accounts.created-date', {
                date: Intl.DateTimeFormat(route.params.locale).format(new Date(tradingAcc.created))
              })
            }}
          </small>

          <div class="controls">
            <Button
              :loading="loading"
              :label="t('settings.accounts.clear-account-btn')"
              @click="() => clearAccount(tradingAcc.id)"
              severity="secondary"
            />

            <Button
              :label="t('settings.accounts.view-details-btn')"
              @click="() => openAccountDetailsDrawer(tradingAcc)"
              severity="secondary"
            />
          </div>
        </div>
      </DataPanel>
    </li>
  </ul>
  <TradingAccountCreator style="max-width: 300px" />
  <Drawer
    header="Trading Account Details"
    v-model:visible="drawerVisible"
    style="width: clamp(400px, 60vw, 1000px)"
  >
    <div v-if="selectedAccount">
      <Suspense>
        <template #default>
          <TradingAccountDetails :account="selectedAccount" />
        </template>
        <template #fallback>
          <LoaderInline />
        </template>
      </Suspense>
    </div>
  </Drawer>
</template>

<style scoped>
h4 {
  margin: 0 0 0.7rem 0;
}

ul {
  padding: 0;
}

li {
  list-style: none;
}

.controls {
  display: flex;
  justify-content: flex-end;
  gap: var(--inline-spacing);
}
</style>
