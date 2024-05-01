<script setup lang="ts">
import { ref, type Ref, computed, onMounted } from 'vue';
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
const screenshots = ref<ScreenshotsByDate>({});
onMounted(fetchAndUpdateScreenshots(screenshots));

function fetchAndUpdateScreenshots(ref: Ref<ScreenshotsByDate>) {
  return () => {
    pb.collection('screenshots')
      .getList(undefined, undefined, {
        skipTotal: true,
        filter: pb.filter('date >= {:startDate} && date <= {:endDate}', {
          startDate: new Date('2024-02-01'),
          endDate: new Date('2024-03-01')
        })
      })
      .then((list) => {
        for (const record of list.items) ref.value[record.date.slice(0, 10)] = record; // refactor task: each date should hold a list of records - i.e. { [date: string]: ScreenshotType[] }
      });
  };
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
