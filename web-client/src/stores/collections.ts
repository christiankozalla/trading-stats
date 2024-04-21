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

  const state = ref<
    Partial<{ [TradingAccountId: string]: Record<Collections, ListResult<RecordModel>> }>
  >({});

  async function getList(
    collectionId: Collections,
    page?: number,
    perPage?: number,
    options?: RecordListOptions
  ) {
    const tradingAcccountId = tradingAccountsStore.selected;
    if (!tradingAcccountId) {
      console.warn(
        `Cannot fetch collection: ${collectionId} because no trading account is selected`
      );
    } else {
      const existing = state.value[tradingAcccountId]?.[collectionId];
      if (existing && (page || 0) >= existing?.totalPages) {
        return existing;
      }

      loaderStore.startLoading();
      try {
        const list = await pb.collection(collectionId).getList(page, perPage, options);
        // make sure existing store collection data is from same account
        if (existing) {
          existing.items.push(...list.items);
          existing.page = list.page;
          existing.totalItems = list.totalItems;
          existing.totalPages = list.totalPages;
        } else {
          const existingTradingAccountNamespace = state.value[tradingAcccountId];
          if (existingTradingAccountNamespace) {
            existingTradingAccountNamespace[collectionId] = list;
          } else {
            state.value[tradingAcccountId] = {
              [collectionId]: list
            } as Record<Collections, ListResult<RecordModel>>;
          }
        }

        return state.value[collectionId];
      } catch (e) {
        console.error('error fetching collection ', collectionId);
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
  }

  async function get(collection: Collections, { nextPage } = { nextPage: false }) {
    const tradingAcccountId = tradingAccountsStore.selected;

    if (!tradingAcccountId) {
      return null;
    }

    if (nextPage) {
      return getList(collection, (state.value[tradingAcccountId]?.[collection]?.page || 0) + 1);
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
