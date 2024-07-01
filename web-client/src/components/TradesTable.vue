<script setup lang="ts">
import { ref, type Ref, computed, onMounted, onUpdated } from 'vue';
import { pb, type Screenshot as ScreenshotType } from '@/api-client';
import { useI18nStore } from '@/stores/i18n';
import Screenshot from '@/components/Screenshot.vue';
import UploadScreenshot from '@/components/UploadScreenshot.vue';
import { type usePaginatedCollection } from '@/composables/usePaginatedCollection';
import { useScreenshotViewer } from '@/composables/useScreenshotViewer';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const { t } = useI18nStore();
const tradingAccountsStore = useTradingAccountsStore();

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

const { viewer } = useScreenshotViewer();

function handleScreenshotViewer(emittedRecord: ScreenshotType) {
  viewer.setActiveRecord(emittedRecord);
  viewer.setIsOpen(true);
}

// this could be implemented with a watcher as well
onMounted(fetchAndUpdateScreenshots(screenshots));
onUpdated(fetchAndUpdateScreenshots(screenshots));

function fetchAndUpdateScreenshots(ref: Ref<ScreenshotsByDate>) {
  return () => {
    if (dateRange.value.min && dateRange.value.max) {
      pb.collection('screenshots')
        .getList(undefined, undefined, {
          skipTotal: true,
          filter: pb.filter('account = {:accountId} && date >= {:startDate} && date <= {:endDate}', {
            accountId: tradingAccountsStore.selected,
            startDate: new Date(dateRange.value.min),
            endDate: new Date(dateRange.value.max)
          })
        })
        .then((list) => {
          for (const record of list.items) ref.value[record.date.slice(0, 10)] = record;
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

// for uploading a new screenshot
const screenshotDate = ref<string | null | undefined>();
const isScreenshotDialogOpen = ref<boolean>(false);
function openAddScreenshotDialog(date: string) {
  screenshotDate.value = date;
  isScreenshotDialogOpen.value = true;
}

function handleScreenshotUploadSuccess() {
  setTimeout(() => {
    isScreenshotDialogOpen.value = false;
  }, 100);
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
          <Screenshot
            v-if="screenshots[slotProps.data.DateTime_close]"
            :record="screenshots[slotProps.data.DateTime_close]"
            thumb
            @emitActiveScreenshot="handleScreenshotViewer"
          />
          <Button
            v-else
            icon="icon icon-upload"
            severity="secondary"
            rounded
            :aria-label="t('generic.upload', { type: 'Screenshot' })"
            @click="openAddScreenshotDialog(slotProps.data.DateTime_close)"
          />
        </template>
      </Column>
    </DataTable>
  </section>
  <Dialog
    v-model:visible="isScreenshotDialogOpen"
    :header="t('import.screenshot.header')"
    style="width: 90vw; max-width: 1100px"
    modal
  >
    <UploadScreenshot :date="screenshotDate" @uploadSuccess="handleScreenshotUploadSuccess" />
  </Dialog>
</template>
