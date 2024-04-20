import { onMounted } from 'vue';
import type { Collections } from '@/api-client';
import { useCollectionsStore } from '@/stores/collections';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

export function usePaginatedCollection(collectionId: Collections) {
  const collectionsStore = useCollectionsStore();
  const tradingAccountStore = useTradingAccountsStore();

  onMounted(async () => {
    await collectionsStore.get(collectionId);
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
