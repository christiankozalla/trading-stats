import { defineStore } from 'pinia';
import { computed, ref, onMounted, watch } from 'vue';
import { pb, type TradingAccount } from '@/api-client';

type AccountId = string;

const accountIdKey = 'selectedTradingAccountId';

const useTradingAccountsStore = defineStore('tradingAccounts', () => {
  const state = ref<TradingAccount[]>([]);
  const accounts = computed(() => state.value);
  const selected = ref<AccountId>();

  onMounted(async () => {
    const storedSelectedId = localStorage.getItem(accountIdKey);
    if (storedSelectedId) {
      selected.value = storedSelectedId;
    }

    await get();
  });

  watch(selected, (newValue) => {
    if (newValue) {
      localStorage.setItem(accountIdKey, newValue);
    } else {
      localStorage.removeItem(accountIdKey);
    }
  });

  async function get() {
    pb.collection('trading_accounts')
      .getFullList()
      .then((res) => {
        state.value.push(...res);
      })
      .catch((e) => {
        console.log('createTradingAccount error', e.data);
      });
  }

  function add(...tradingAccounts: TradingAccount[]) {
    state.value.push(...tradingAccounts);
  }

  return {
    accounts,
    selected,
    add
  };
});

export { useTradingAccountsStore };
