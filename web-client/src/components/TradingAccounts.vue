<script setup lang="ts">
import TradingAccountCreator from './TradingAccountCreator.vue';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const tradingAccountsStore = useTradingAccountsStore();

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
  <TradingAccountCreator />
</template>
