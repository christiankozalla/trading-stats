import { onMounted } from 'vue';
import type { Collections } from '@/api-client';
import { useCollectionsStore } from '@/stores/collections';

export function useCollection(collection: Collections) {
  const collectionsStore = useCollectionsStore();

  onMounted(async () => {
    await collectionsStore.get(collection);
  });

  const next = () => collectionsStore.get(collection, { nextPage: true });

  return {
    get collection() {
      return collectionsStore.state[collection];
    },
    next
  };
}
