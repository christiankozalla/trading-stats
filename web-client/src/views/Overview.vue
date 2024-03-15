<script setup lang="ts">
import { computed } from 'vue';
import { useCollection } from '@/composables/useCollection';

const trades = useCollection('view_trades');

const tradeItems = computed(() => {
  return (trades.collection?.items || []).map((trade) => ({
    ...trade,
    ProfitLoss: formatCurrency(trade.ProfitLoss * getMultiplier(trade.Symbol))
  }));
});

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
  'NQH4.CME': 2
};

function getMultiplier(symbol: keyof typeof ticksToCurrency): number {
  return ticksToCurrency[symbol] || 0;
}

function formatCurrency(value: number) {
  return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
}
</script>

<template>
  <section>Three cards side by side Avg. PnL/Day Total PnL Avg. Vol/Day</section>

  <section>
    Three charts side by side Aggregate PnL vs Date Cumulative PnL vs Date Aggregate Volume vs Date
  </section>

  <section v-if="trades.collection?.items">
    <DataTable
      :value="tradeItems"
      sortField="DateTime"
      :sortOrder="-1"
      paginator
      :rows="trades.collection.perPage"
    >
      <template #paginatorstart>
        <Button
          type="button"
          icon="icon icon-refresh"
          @click="trades.next"
          style="padding: 4px 8px"
          label="Load&nbsp;more"
        />
      </template>
      <Column field="t1_DateTime" header="Date" sortable>
        <template #body="props">
          {{
            new Date(props.data.t1_DateTime).toLocaleDateString('de-DE', {
              year: 'numeric',
              month: '2-digit',
              day: '2-digit'
            })
          }}
        </template>
      </Column>
      <Column field="Symbol" header="Symbol" />
      <Column field="ProfitLoss" header="PnL" sortable>
        <template #body="props">
          {{ formatCurrency(props.data.ProfitLoss) }}
        </template>
      </Column>
    </DataTable>
  </section>
</template>
