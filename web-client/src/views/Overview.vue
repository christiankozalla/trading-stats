<script setup lang="ts">
import { type RecordModel } from 'pocketbase';
import { useCollection } from '@/composables/useCollection';
import TradesTable from '@/components/TradesTable.vue';
import DataPanel from '@/components/DataPanel.vue';
import { type Trade } from '@/api-client';

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
  <section>
    Three cards side by side Total PnL Avg. Vol/Day
    <DataPanel>
      <template #left>
        <p>Avg. PnL/Day</p>
      </template>
      <template #right> <p>Win Days 10 Loss Days 1</p></template>
    </DataPanel>

    <DataPanel>
      <div>
        <p>Total PnL</p>
        <p>$40.923 from $13</p>
      </div>
    </DataPanel>
    <DataPanel>
      <template #left> <p>Avg. Vol/Day</p> </template>
      <template #right> <p>PnL/Vol $75.47 Ratio</p></template>
    </DataPanel>
  </section>

  <section>
    Three charts side by side Aggregate PnL vs Date Cumulative PnL vs Date Aggregate Volume vs Date
  </section>

  <TradesTable :trades="trades" :mapper-fn="tradesMapper" />
</template>
