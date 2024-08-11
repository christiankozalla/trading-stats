<script setup lang="ts">
import Panel, { type PanelProps } from "primevue/panel";
const props = defineProps<PanelProps>();
</script>

<template>
  <Panel
    v-bind="props"
    :pt="{
      ...props.pt, // props.pt.header and props.pt.content are overwritten, thus cannot be set from outside
      header: { class: 'datapanel-header' },
      content: { class: 'datapanel-content' }
    }"
  >
    <slot></slot>
    <slot name="left"></slot>
    <slot name="right"></slot>
  </Panel>
</template>

<style scoped>
:deep(.datapanel-header) {
  padding: 9px 18px;
}

:deep(.datapanel-content) > * {
  display: inline-block;
  width: 50%;
  vertical-align: top;
}

:deep(.datapanel-content) > *:nth-child(1) {
  padding-right: 9px;
}

:deep(.datapanel-content) > *:nth-child(2) {
  padding-left: 9px;
}
:deep(.datapanel-content) > *:only-child {
  width: 100%;
  padding: 0;
}

:deep(.p-panel-title) {
  margin: var(--inline-spacing) 0;
}
</style>
