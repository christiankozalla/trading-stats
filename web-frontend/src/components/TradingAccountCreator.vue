<script setup lang="ts">
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import { pb, type TradingAccount } from '@/api-client';
import { useI18nStore } from '@/stores/i18n';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const tradingAccountsStore = useTradingAccountsStore();
const { t } = useI18nStore();

async function createTradingAccount(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);
  formData.append('user', pb.authStore.model?.id);
  // TODO: stop here and add a toast if the user is not authenticated
  pb.collection('trading_accounts')
    .create<TradingAccount>(formData)
    .then((newTradingAcc) => {
      tradingAccountsStore.add(newTradingAcc);
    })
    .catch((e) => {
      console.log('createTradingAccount error', e.data);
    });
}
</script>

<template>
  <form @submit.prevent="createTradingAccount" style="padding: var(--inline-spacing)">
    <InputText
      id="name"
      name="name"
      required
      :placeholder="t('trading-accounts.name')"
      autocomplete="off"
    />
    <Button
      type="submit"
      :label="t('trading-accounts.create-cta')"
      style="display: block; width: 100%; margin-top: 8px"
    ></Button>
  </form>
</template>
