import { ref } from 'vue';
import { type Screenshot } from '@/api-client';

const isOpen = ref<boolean>(false);
const screenshotRecords = ref<Screenshot[]>([]);
const activeRecord = ref<Screenshot | null>(null);

export const useScreenshotViewer = () => ({
  viewer: {
    get isOpen() {
      return isOpen.value;
    },
    setIsOpen(state: boolean) {
      isOpen.value = state;
    },
    get activeRecord() {
      return activeRecord.value;
    },
    setActiveRecord(record: Screenshot) {
      activeRecord.value = record;
    }
  },
  screenshotRecords
});
