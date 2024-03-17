<script setup lang="ts">
import { pb, type TradingAccount } from '@/api-client';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const tradingAccountsStore = useTradingAccountsStore();

async function createTradingAccount(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);
  formData.append('user', pb.authStore.model?.id);
  pb.collection('trading_accounts')
    .create<TradingAccount>(formData)
    .then((newTradingAcc) => {
      tradingAccountsStore.add(newTradingAcc);
    })
    .catch((e) => {
      //   toast.add({
      //     severity: 'error',
      //     summary: 'CRUD error',
      //     detail: e.data.message,
      //     life: 5000
      //   });
      console.log('createTradingAccount error', e.data);
    });
}
</script>

<template>
  <h2>Trading Accounts</h2>
  <ul v-if="tradingAccountsStore.accounts.length">
    <li v-for="tradingAcc in tradingAccountsStore.accounts" :key="tradingAcc.id">
      {{ tradingAcc.name }}
      {{ tradingAcc.created }}
    </li>
  </ul>
  <p v-else>No trading accounts...</p>
  <form @submit.prevent="createTradingAccount">
    <div>
      <InputText id="name" name="name" required />
      <!-- {{ formErrors.name }} -->
    </div>
    <Button type="submit" label="Create Account"></Button>
  </form>
</template>
