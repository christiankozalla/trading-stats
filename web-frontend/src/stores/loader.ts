import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

const useLoaderStore = defineStore('loader', () => {
  const state = ref(0);

  function startLoading() {
    state.value++;
  }

  function stopLoading() {
    state.value--;
  }

  const isLoading = computed(() => {
    return state.value > 0;
  });

  return {
    startLoading,
    stopLoading,
    isLoading
  };
});

export { useLoaderStore };
