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
  <div class="flex flex-row flex-gap h-100">
  <img
    :src="src"
    :class="{ clickable: props.thumb, viewer: !props.thumb, 'm-auto': !props.thumb && !record.comment }"
    alt=""
    :loading="props.loading || 'eager'"
    @click="$emit('emitActiveScreenshot', props.record)"
  />
  <p v-if="!thumb && record.comment">{{  record.comment }}</p>
</div>
</template>

<style scoped>
.clickable {
  cursor: pointer;
}

.viewer {
  display: block;
  object-fit: contain;
  height: 100%;
}
</style>
