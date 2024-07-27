<script setup lang="ts">
import { reactive, computed } from 'vue';
import { pb, type Screenshot } from '@/api-client';

const props = defineProps<{
  record: Screenshot;
  loading?: 'eager' | 'lazy';
  thumb?: boolean;
}>();

const thumbSrc = computed(() =>
  pb.files.getUrl(props.record, props.record.image, { thumb: '0x40' })
);
const imageSrc = computed(() => pb.files.getUrl(props.record, props.record.image));
const src = computed(() => (props.thumb ? thumbSrc.value : imageSrc.value));

const hasComment = computed(() => !props.thumb && props.record?.comment);
const imageClasses = reactive({
  clickable: props.thumb,
  viewer: !props.thumb,
  'm-auto': !props.thumb && !props.record.comment
});
const imageStyles = reactive({
  width: !props.thumb && props.record?.imageWidth ? `${props.record?.imageWidth}px` : undefined,
  'max-height': '90vh',
  'aspect-ratio':
    props.record?.imageHeight && props.record?.imageWidth
      ? props.record?.imageWidth / props.record?.imageHeight
      : undefined
});
</script>

<template>
  <div
    :class="{
      screenshot__container: !props.thumb,
      'screenshot__container--with-comment': hasComment
    }"
  >
    <img
      :src="src"
      :class="imageClasses"
      alt=""
      :style="imageStyles"
      :loading="props.loading || 'eager'"
      @click="$emit('emitActiveScreenshot', props.record)"
    />
    <p v-if="!props.thumb && props.record.comment">{{ props.record.comment }}</p>
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

.screenshot__container {
  max-height: 90vh;
}

.screenshot__container p {
  margin-top: var(--inline-spacing);
}
</style>
