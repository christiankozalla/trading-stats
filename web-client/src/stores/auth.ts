import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { type User } from '@/api-client';

// this store is mainly for UI reactivity
const useAuthStore = defineStore('auth', () => {
  const model = ref<User | null>(null);

  const isAuthenticated = computed(() => {
    return !!model.value;
  });

  return {
    model,
    isAuthenticated
  };
});

export { useAuthStore };
