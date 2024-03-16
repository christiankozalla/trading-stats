<script setup lang="ts">
import { useCollection } from '@/composables/useCollection';
import TradesTable from '@/components/TradesTable.vue';
import { type Trade } from '@/api-client';
import { type RecordModel } from 'pocketbase';

const trades = useCollection('view_trades');

const tradesMapper = (trade: Trade & RecordModel) => ({
  ...trade,
  ProfitLoss: formatCurrency(
    trade.ProfitLoss * getMultiplier(trade.Symbol as keyof typeof ticksToCurrency)
  )
});

function getMultiplier(symbol: keyof typeof ticksToCurrency): number {
  return ticksToCurrency[symbol] || 0;
}

function formatCurrency(value: number) {
  return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
}

const ticksToCurrency = {
  'CLF4.NYMEX': 1000,
  'CLG4.NYMEX': 1000,
  'CLH4.NYMEX': 1000,
  'CLJ4.NYMEX': 1000,
  'CLK4.NYMEX': 1000,
  'CLM4.NYMEX': 1000,
  'CLN4.NYMEX': 1000,
  'CLQ4.NYMEX': 1000,
  'CLU4.NYMEX': 1000,
  'CLV4.NYMEX': 1000,
  'CLX4.NYMEX': 1000,
  'CLZ4.NYMEX': 1000,
  'NQF4.CME': 2,
  'NQG4.CME': 2,
  'NQH4.CME': 2,
  'NQJ4.CME': 2,
  'NQK4.CME': 2,
  'NQM4.CME': 2,
  'NQN4.CME': 2,
  'NQQ4.CME': 2,
  'NQU4.CME': 2,
  'NQV4.CME': 2,
  'NQX4.CME': 2,
  'NQZ4.CME': 2
};
</script>

<template>
  <section>Three cards side by side Avg. PnL/Day Total PnL Avg. Vol/Day</section>

  <section>
    Three charts side by side Aggregate PnL vs Date Cumulative PnL vs Date Aggregate Volume vs Date
  </section>

  <TradesTable :trades="trades" :mapper-fn="tradesMapper" />
</template>
