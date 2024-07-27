import { onMounted } from 'vue';
import type { Collections } from '@/api-client';
import { useCollectionsStore } from '@/stores/collections';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { useLoaderStore } from '@/stores/loader';

export function usePaginatedCollection(collectionId: Collections) {
  const collectionsStore = useCollectionsStore();
  const tradingAccountStore = useTradingAccountsStore();
  const loaderStore = useLoaderStore();

  onMounted(async () => {
    loaderStore.startLoading();
    await collectionsStore.get(collectionId).finally(() => {
      loaderStore.stopLoading();
    });
  });

  const next = () => collectionsStore.get(collectionId, { nextPage: true });

  return {
    get collection() {
      if (tradingAccountStore.selected) {
        return collectionsStore.state[tradingAccountStore.selected]?.[collectionId];
      } else {
        return null;
      }
    },
    next
  };
}
