<script setup lang="ts">
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const tradingAccountsStore = useTradingAccountsStore();
</script>

<template>
  <Dropdown
    v-model="tradingAccountsStore.selected"
    optionLabel="name"
    optionValue="id"
    :options="tradingAccountsStore.accounts"
    placeholder="Select a trading account"
  />
</template>
