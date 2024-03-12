import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { pb, type TradingAccount } from '@/api-client';

const useTradingAccountsStore = defineStore('tradingAccounts', () => {
  const state = ref<TradingAccount[]>([]);
  const accounts = computed(() => state.value);

  async function get(page?: number) {
    pb.collection<TradingAccount>('trading_accounts')
      .getList(page)
      .then((res) => {
        state.value.push(...res.items);
        if (state.value.length === res.totalItems) return;
        if (res.page < res.totalPages) {
          return get(res.page + 1);
        }
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
    get,
    add
  };
});

export { useTradingAccountsStore };
