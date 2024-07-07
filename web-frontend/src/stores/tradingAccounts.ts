import { defineStore } from 'pinia';
import { computed, ref, watch } from 'vue';
import { pb, isAuthenticated, type TradingAccount } from '@/api-client';

type AccountId = string;

const accountIdKey = 'selectedTradingAccountId';

const useTradingAccountsStore = defineStore('tradingAccounts', () => {
  const state = ref<TradingAccount[]>([]);
  const accounts = computed(() => state.value);
  const selected = ref<AccountId>();
  const scopedKey = `${accountIdKey}-${pb.authStore.model?.id}`;

  watch(accounts, (newAccounts) => {
    if (newAccounts.length === 1) {
      // if there is only one account, select it
      selected.value = newAccounts[0].id;
    }
  });

  watch(selected, (newValue) => {
    if (newValue) {
      localStorage.setItem(scopedKey, newValue);
    } else {
      localStorage.removeItem(scopedKey);
    }
  });

  watch(
    isAuthenticated,
    async (isAuthenticated) => {
      if (isAuthenticated) {
        const storedSelectedId = localStorage.getItem(scopedKey);
        if (storedSelectedId) {
          selected.value = storedSelectedId;
        }
        await get();
      } else {
        resetState();
      }
    },
    { immediate: true }
  );

  async function get() {
    pb.collection('trading_accounts')
      .getFullList()
      .then((res) => {
        state.value.push(...res);
      })
      .catch((e) => {
        console.log('tradingAccounts error', e.data);
      });
  }

  function add(...tradingAccounts: TradingAccount[]) {
    state.value.push(...tradingAccounts);
  }

  function resetState() {
    state.value = [];
  }

  return {
    accounts,
    selected,
    add
  };
});

export { useTradingAccountsStore };
