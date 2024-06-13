<template>
  <Dialog
    :visible="isOpen"
    header="Screenshot Viewer"
    style="width: max-content; max-width: 90vw; max-height: 90vh; margin: 24px"
    :pt="{ content: { class: 'screenshot-viewer__dialog-content' } }"
    @update:visible="(visible) => $emit('update:visible', visible)"
    modal
  >
    <Screenshot :record="activeRecord" />
  </Dialog>
</template>

<script setup lang="ts">
import { type Screenshot } from '@/api-client';

const { isOpen, activeRecord } = defineProps<{
  isOpen: boolean;
  activeRecord?: Screenshot | null;
}>();
</script>

<style>
.screenshot-viewer__dialog-content {
  margin: var(--inline-spacing);
  margin-top: 0;
}
</style>
