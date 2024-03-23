<script setup lang="ts">
import { usePaginatedCollection } from '@/composables/usePaginatedCollection';
import TradesTable from '@/components/TradesTable.vue';
import DataPanel from '@/components/DataPanel.vue';
import { pb, type ProfitLoss, type Trade } from '@/api-client';
import { computed, onMounted, ref } from 'vue';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  PointElement,
  Legend,
  LineElement,
  CategoryScale,
  LinearScale,
  BarElement
} from 'chart.js';
import { Line, Bar, type ChartProps } from 'vue-chartjs';
import { useLoaderStore } from '@/stores/loader';

const loaderStore = useLoaderStore();
ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  LineElement,
  PointElement,
  Title,
  Tooltip,
  Legend
);
const profitLoss = ref<(ProfitLoss & { Date_close: string })[]>();

onMounted(() => {
  loaderStore.startLoading();
  pb.collection('profit_loss')
    .getFullList({ sort: '-DateTime_close' })
    .then(
      (records) =>
        (profitLoss.value = records.map(
          (record) =>
            ({
              ...record,
              Date_close: new Date(record.DateTime_close).toISOString().substring(0, 10)
            }) as ProfitLoss & { Date_close: string }
        ))
    )
    .finally(() => loaderStore.stopLoading());
});

type DateString = string;

const pnlData = computed(() => {
  let total = 0;
  const cumulative: ChartProps<'line'>['data'] = {
    labels: [],
    datasets: []
  };
  const saldo: ChartProps<'bar'>['data'] = {
    labels: [],
    datasets: []
  };
  if (profitLoss.value) {
    const cumulativeData: Record<DateString, number> = {};
    const saldoData: Record<DateString, number[]> = {};

    for (const record of profitLoss.value) {
      total += record.ProfitLoss_dollar;
      cumulativeData[record.Date_close] = total;
      if (!saldoData[record.Date_close]) {
        saldoData[record.Date_close] = [record.ProfitLoss_dollar];
      } else {
        saldoData[record.Date_close].push(record.ProfitLoss_dollar);
      }
    }
    cumulative.labels = Object.keys(cumulativeData);
    cumulative.datasets.push({ label: 'Cumulative PnL', data: Object.values(cumulativeData) });
    saldo.labels = Object.keys(saldoData);
    saldo.datasets.push({
      label: 'Change PnL',
      data: Object.values(saldoData).map((changes) => changes.reduce((sum, curr) => sum + curr, 0))
    });
  }
  return {
    cumulative,
    saldo
  };
});

const trades = usePaginatedCollection('trades');

const tradesMapper = (trade: Trade) => ({
  ...trade,
  ProfitLoss: formatCurrency(trade.ProfitLoss * trade.Multiplier)
});

function formatCurrency(value: number) {
  return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
}
</script>

<template>
  <section>
    <Line v-if="!loaderStore.isLoading" :data="pnlData.cumulative" />
    <Bar v-if="!loaderStore.isLoading" :data="pnlData.saldo" />
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
