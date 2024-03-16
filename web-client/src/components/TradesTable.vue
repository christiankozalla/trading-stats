<script setup lang="ts">
import { computed } from 'vue';
import { type useCollection } from '@/composables/useCollection';

const props = defineProps<{
  trades: ReturnType<typeof useCollection>;
  mapperFn: (trade: any) => any;
}>();

const items = computed(() => {
  return (props.trades.collection?.items || []).map(props.mapperFn);
});
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
      <Column field="Symbol" header="Symbol" />
      <Column field="ProfitLoss" header="PnL" sortable />
    </DataTable>
  </section>
</template>
