<template>
  <Dialog
    :visible="isOpen"
    header="Screenshot Viewer"
    style="width: max-content; max-width: 90vw; max-height: 90vh; margin: 24px"
    :pt="{ content: { class: 'screenshot-viewer__dialog-content' } }"
    @update:visible="(visible) => $emit('update:visible', visible)"
    modal
  >
    <Screenshot v-if="activeRecord" :record="activeRecord" />
  </Dialog>
</template>

<script setup lang="ts">
import Screenshot from './Screenshot.vue';
import Dialog from 'primevue/dialog';
import { type Screenshot as ScreenshotType } from '@/api-client';

const { isOpen, activeRecord } = defineProps<{
  isOpen: boolean;
  activeRecord?: ScreenshotType | null;
}>();
</script>

<style scoped>
:deep(.screenshot-viewer__dialog-content) {
  margin: var(--inline-spacing);
  margin-top: 0;
}
</style>
