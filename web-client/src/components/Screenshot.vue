<script setup lang="ts">
import { computed } from 'vue';
import { pb, type Screenshot } from '@/api-client';

const props = defineProps<{
  record: Screenshot;
  loading?: 'eager' | 'lazy';
  thumb?: boolean;
}>();

const thumbSrc = pb.files.getUrl(props.record, props.record.image, { thumb: '0x40' });
const imageSrc = pb.files.getUrl(props.record, props.record.image);

const src = computed(() => (props.thumb ? thumbSrc : imageSrc));
</script>

<template>
  <img
    :src="src"
    :class="{ clickable: props.thumb, viewer: !props.thumb }"
    alt=""
    :loading="props.loading || 'eager'"
    @click="$emit('emitActiveScreenshot', props.record)"
  />
</template>

<style scoped>
.clickable {
  cursor: pointer;
}

.viewer {
  display: block;
  object-fit: contain;
  height: 100%;
  margin: 0 auto;
}
</style>
