<script setup lang="ts">
import { computed } from 'vue';
const props = defineProps<{ file?: File }>();

const dataURI = computed(() => {
  if (props.file?.type.startsWith('image/')) {
    return URL.createObjectURL(props.file);
  } else {
    console.warn('No file or not an image.');
    return '';
  }
});
</script>

<template>
  <img v-if="dataURI" :src="dataURI" />
</template>

<style scoped>
img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
</style>
