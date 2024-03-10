<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { pb } from '@/api-client';
import { ClientResponseError, type ListResult, type RecordModel } from 'pocketbase';

type TradingAccount = {
  name: string
} & RecordModel

const tradingAccounts = ref<ListResult<TradingAccount>>();

onMounted(() => {
  pb.collection<TradingAccount>('trading_accounts')
    .getList()
    .then((res) => {
      tradingAccounts.value = res;
    })
    .catch((e) => {
      if (e instanceof ClientResponseError) {
        console.log('createTradingAccount error', e.data);
      }
    });
});

async function createTradingAccount(event: Event) {
  try {
    const formData = new FormData(event.target as HTMLFormElement);
    formData.append('user', pb.authStore.model?.id);
    const newTradingAcc = await pb.collection('trading_accounts').create<TradingAccount>(formData);

    tradingAccounts.value?.items.push(newTradingAcc);
  } catch (e) {
    if (e instanceof ClientResponseError) {
      //   toast.add({
      //     severity: 'error',
      //     summary: 'CRUD error',
      //     detail: e.data.message,
      //     life: 5000
      //   });
      console.log('createTradingAccount error', e.data);
    }
  }
}
</script>

<template>
  <h2>Trading Accounts</h2>
  <ul v-if="tradingAccounts && tradingAccounts.items.length">
    <li v-for="tradingAcc in tradingAccounts.items" :key="tradingAcc.id">
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
