<script setup lang="ts">
import TradingAccountCreator from './TradingAccountCreator.vue';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const tradingAccountsStore = useTradingAccountsStore();
</script>

<template>
  <h2>Trading Accounts</h2>
  <!-- List of all trading accounts -> Name (fat) + Created date + Link that opens right drawer with "Quick Info" (like Commissions, Fees, Account Details (List of Log-Files and Trades derived from them), Public Dashboard Switch, "Export Data" Btn, Clear Account Btn, Delete Account Bntn) -->
  <ul v-if="tradingAccountsStore.accounts.length">
    <li v-for="tradingAcc in tradingAccountsStore.accounts" :key="tradingAcc.id">
      {{ tradingAcc.name }}
      {{ tradingAcc.created }}
    </li>
  </ul>
  <p v-else>No trading accounts...</p>
  <TradingAccountCreator />
</template>
