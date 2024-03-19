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
    trade.ProfitLoss * trade.Multiplier
  )
});

function formatCurrency(value: number) {
  return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
}
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
