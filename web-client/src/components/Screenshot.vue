<script setup lang="ts">
import { pb, type Screenshot } from '@/api-client';

const props = defineProps<{
  record?: Screenshot;
}>();
</script>

<template>
  <div v-if="props.record">
    <img :src="pb.files.getUrl(props.record, props.record.image, { thumb: '0x40' })" alt="" />
  </div>
  <span v-else>Feature request: Click to upload a screenshot</span>
</template>

<style scoped>
img {
  max-height: 20px;
  margin: 0 auto;
  display: block;
}
</style>