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
import { useI18nStore } from '@/stores/i18n';

const { t } = useI18nStore();

// helpers
const sum = (array: number[]) => array.reduce((sum, curr) => sum + curr, 0);
const displayMoney = (n: number) => `${n.toFixed(1)} $`;
const toISODate = (date: string | Date) => {
  if (typeof date === 'string') {
    return new Date(date).toISOString().substring(0, 10);
  } else return date.toISOString().substring(0, 10);
};

function startOfWeek(dateInput: string | Date): string {
  const date = new Date(dateInput);
  const dayOfWeek = date.getDay(); // 0 for Sunday, 1 for Monday, etc.
  const daysSinceMonday = (dayOfWeek + 6) % 7; // Number of days to subtract to get to Monday
  date.setDate(date.getDate() - daysSinceMonday);
  return toISODate(date);
}

function groupDataByWeek(data: Record<string, number[]>) {
  const groupedData: Record<string, { data: number[]; sum: number }> = {};

  for (const dateString in data) {
    const weekStart = toISODate(startOfWeek(dateString));
    if (!groupedData[weekStart]) {
      groupedData[weekStart] = { data: [], sum: 0 };
    }
    groupedData[weekStart].data.push(...data[dateString]);
    groupedData[weekStart].sum = sum(data[dateString]);
  }
  // example for accessing the latest week
  // const latestWeekString = DateTime.now().startOf('week')
  // const latestWeekData = groupedData[latestWeekString]
  return groupedData;
}
// end helpers

type ProfitLossExtended = ProfitLoss & { Date_close: string };

const sevenDaysInMilliseconds = 7 * 24 * 60 * 60 * 1000;
const currentWeekStart = toISODate(startOfWeek(new Date()));
const previousWeekStart = toISODate(
  startOfWeek(new Date(new Date().getTime() - sevenDaysInMilliseconds))
);

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
const profitLoss = ref<ProfitLossExtended[]>();

onMounted(() => {
  loaderStore.startLoading();
  pb.collection('profit_loss')
    .getFullList({ sort: 'DateTime_close' })
    .then(
      (records) =>
        (profitLoss.value = records.map(
          (record) =>
            ({
              ...record,
              Date_close: toISODate(record.DateTime_close)
            }) as ProfitLossExtended
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
  let weekly: ReturnType<typeof groupDataByWeek> = {};
  let avgPerDay = 0;
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
    cumulative.datasets.push({ label: 'Cumulative PnL $', data: Object.values(cumulativeData) });
    saldo.labels = Object.keys(saldoData);
    const saldoBars = Object.values(saldoData).map(sum);
    saldo.datasets.push({
      label: 'Daily Change PnL $',
      data: saldoBars,
      backgroundColor: saldoBars.map((bar) => (bar >= 0 ? 'green' : 'red'))
    });
    weekly = groupDataByWeek(saldoData);
    avgPerDay = total / saldoBars.length;
  }
  return {
    cumulative,
    saldo,
    weekly,
    total,
    avgPerDay
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
  <section v-if="!loaderStore.isLoading" class="data-panels">
    <DataPanel>
      <template #left>
        <p>
          <strong>{{ t('pnl.total') }}</strong> <br />
          {{ pnlData.total ? displayMoney(pnlData.total) : 'No data' }}
        </p>
      </template>
      <template #right>
        <p>
          <strong>{{ t('pnl.win-days') }}</strong>
          {{
            pnlData.saldo.datasets[0]?.data?.filter((n) => typeof n === 'number' && n > 0).length ||
            'No data'
          }}
          <br />
          <strong>{{ t('pnl.loss-days') }}</strong>
          {{
            pnlData.saldo.datasets[0]?.data?.filter((n) => typeof n === 'number' && n < 0).length ||
            'No data'
          }}
        </p>
      </template>
    </DataPanel>
    <DataPanel>
      <template #left>
        <p>
          <strong>{{ t('pnl.avg-per-day') }}</strong> <br />
          {{ displayMoney(pnlData.avgPerDay) }}
        </p>
      </template>
      <template #right>
        <p>
          <strong>{{ t('pnl.avg-per-week') }}</strong> <br />
          {{ displayMoney(Object.values(pnlData.weekly).reduce((sum, curr) => sum + curr.sum, 0)) }}
        </p>
      </template>
    </DataPanel>
    <DataPanel>
      <template #left>
        <p>
          <strong> {{ t('pnl.current-week') }}</strong
          ><br />
          {{ currentWeekStart }}:
          {{ pnlData.weekly[currentWeekStart]?.sum || 'No data' }}
        </p>
      </template>
      <template #right>
        <p>
          <strong>{{ t('pnl.last-week') }}</strong
          ><br />
          {{ previousWeekStart }}:
          {{ pnlData.weekly[previousWeekStart]?.sum || 'No data' }}
        </p>
      </template>
    </DataPanel>
  </section>

  <section v-if="!loaderStore.isLoading" class="charts">
    <div>
      <Line :data="pnlData.cumulative" />
    </div>
    <div>
      <Bar :data="pnlData.saldo" />
    </div>
  </section>

  <section>
    <TradesTable :trades="trades" :mapper-fn="tradesMapper" />
  </section>
</template>

<style scoped>
section {
  margin: 24px;
}

.data-panels {
  display: grid;
  gap: 12px;
  grid-template-columns: 3fr 3fr 5fr;
  grid-auto-rows: auto;
}

:deep(.p-component) {
  font-size: 0.8rem;
}

.charts {
  display: flex;
  gap: 12px;
}

.charts div {
  width: calc(50% - 12px);
}

@media (max-width: 500px) {
  .data-panels {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 960px) {
  .data-panels {
    grid-template-columns: 1fr 1fr;
  }

  .data-panels > *:last-child {
    grid-column: span 2;
  }

  .charts {
    display: block;
  }

  .charts div {
    width: 100%;
  }

  .charts div canvas {
    width: 100% !important;
    height: auto !important;
  }
}
</style>
