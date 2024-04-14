import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useToast } from 'primevue/usetoast';
import {
  type RecordListOptions,
  type ListResult,
  type RecordModel,
  ClientResponseError
} from 'pocketbase';
import { pb, type Collections } from '@/api-client';
import { useLoaderStore } from './loader';
import { useTradingAccountsStore } from './tradingAccounts';

const useCollectionsStore = defineStore('collections', () => {
  const tradingAccountsStore = useTradingAccountsStore();
  const loaderStore = useLoaderStore();
  const toast = useToast();

  const state = ref<Partial<Record<Collections, ListResult<RecordModel>>>>({});

  async function getList(
    collection: Collections,
    page?: number,
    perPage?: number,
    options?: RecordListOptions
  ) {
    const existing = state.value[collection];
    if (existing && (page || 0) >= existing?.totalPages) {
      return existing;
    }

    loaderStore.startLoading();
    try {
      const list = await pb.collection(collection).getList(page, perPage, options);
      // make sure existing store collection data is from same account
      if (existing && existing.items[0]?.accountId === tradingAccountsStore.selected) {
        existing.items.push(...list.items);
        existing.page = list.page;
        existing.totalItems = list.totalItems;
        existing.totalPages = list.totalPages;
      } else {
        state.value[collection] = list;
      }

      return state.value[collection];
    } catch (e) {
      console.error('error fetching collection ', collection);
      if (e instanceof ClientResponseError) {
        toast.add({
          severity: 'error',
          summary: 'Failed to fetch collection',
          detail: e.data.message,
          life: 5000
        });
      }
    } finally {
      loaderStore.stopLoading();
    }
  }

  async function get(collection: Collections, { nextPage } = { nextPage: false }) {
    if (nextPage) {
      return getList(collection, (state.value[collection]?.page || 0) + 1);
    } else {
      return getList(collection);
    }
  }

  return {
    state,
    get
  };
});

export { useCollectionsStore };
