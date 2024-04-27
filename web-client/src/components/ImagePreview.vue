<script setup lang="ts">
import { ref } from 'vue';
const props = defineProps<{ file: File }>();
const dataURI = ref<string>();

if (props.file.type.startsWith('image/')) {
  dataURI.value = URL.createObjectURL(props.file);
} else {
  console.warn('No file or not an image.');
}
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
