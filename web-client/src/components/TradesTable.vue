<script setup lang="ts">
import { ref, type Ref, computed, onMounted, onUpdated } from 'vue';
import { pb, type Screenshot as ScreenshotType } from '@/api-client';
import Screenshot from '@/components/Screenshot.vue';
import { type usePaginatedCollection } from '@/composables/usePaginatedCollection';

const props = defineProps<{
  trades: ReturnType<typeof usePaginatedCollection>;
  mapperFn: (trade: any) => any;
}>();

type ScreenshotsByDate = { [date: string]: ScreenshotType };

const items = computed(() => {
  return (props.trades.collection?.items || []).map(props.mapperFn);
});
const dateRange = computed(() => getDateRange(items.value));
const screenshots = ref<ScreenshotsByDate>({});

// this could be implemented with a watcher as well
onMounted(fetchAndUpdateScreenshots(screenshots));
onUpdated(fetchAndUpdateScreenshots(screenshots));

function fetchAndUpdateScreenshots(ref: Ref<ScreenshotsByDate>) {
  return () => {
    if (dateRange.value.min && dateRange.value.max) {
      pb.collection('screenshots')
        .getList(undefined, undefined, {
          skipTotal: true,
          filter: pb.filter('date >= {:startDate} && date <= {:endDate}', {
            startDate: new Date(dateRange.value.min),
            endDate: new Date(dateRange.value.max)
          })
        })
        .then((list) => {
          for (const record of list.items) ref.value[record.date.slice(0, 10)] = record; // refactor task: each date should hold a list of records - i.e. { [date: string]: ScreenshotType[] }
        });
    }
  };
}

// https://stackoverflow.com/questions/14781153/how-to-compare-two-string-dates-in-javascript
function isLater(date1: string, date2: string) {
  return date1 > date2; // date1 is later than date2 - e.g. 2024-02-10 is later than 2023-11-30
}
function isEarlier(date1: string, date2: string) {
  return date1 < date2;
}

function getDateRange(list: { DateTime_close: string }[]): { min?: string; max?: string } {
  let min = list[0]?.DateTime_close.slice(0, 10);
  let max = list[0]?.DateTime_close.slice(0, 10);

  for (const { DateTime_close } of list) {
    const currentDate = DateTime_close.slice(0, 10);
    if (isEarlier(currentDate, min)) {
      min = currentDate;
    } else if (isLater(currentDate, max)) {
      max = currentDate;
    }
  }
  return { min, max };
}
</script>

<template>
  <section v-if="props.trades.collection">
    <DataTable
      :value="items"
      sortField="DateTime_close"
      :sortOrder="-1"
      paginator
      :rows="props.trades.collection.perPage"
    >
      <template #paginatorstart>
        <Button
          type="button"
          icon="icon icon-refresh"
          @click="props.trades.next"
          style="padding: 4px 8px"
          label="Load&nbsp;more"
        />
      </template>
      <Column field="DateTime_close" header="Date" sortable>
        <template #body="slotProps">
          {{
            new Date(slotProps.data.DateTime_close).toLocaleDateString('de-DE', {
              year: 'numeric',
              month: '2-digit',
              day: '2-digit'
            })
          }}
        </template>
      </Column>
      <Column field="ShortSymbol" header="Symbol" />
      <Column field="ProfitLoss" header="PnL" sortable />
      <Column field="DateTime_close" header="Screenshot">
        <template #body="slotProps">
          <Screenshot :record="screenshots[slotProps.data.DateTime_close.slice(0, 10)]" />
        </template>
      </Column>
    </DataTable>
  </section>
</template>
