import { defineStore } from "pinia";
import { computed, ref, watch } from "vue";
import { pb, isAuthenticated, type TradingAccount } from "@/api-client";

type AccountId = string;

const accountIdKey = "selectedTradingAccountId";

const useTradingAccountsStore = defineStore("tradingAccounts", () => {
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
        await revalidate();
      } else {
        resetState();
      }
    },
    { immediate: true }
  );

  async function revalidate() {
    pb.collection("trading_accounts")
      .getFullList()
      .then((res) => {
        state.value = res;
      })
      .catch((e) => {
        console.log("tradingAccounts error", e.data);
      });
  }

  function add(...tradingAccounts: TradingAccount[]) {
    state.value.push(...tradingAccounts);
  }

  function remove(accountId: string) {
    state.value = state.value.filter((ta) => ta.id !== accountId);
    if (selected.value === accountId) {
      selected.value = undefined;
    }
  }

  function resetState() {
    state.value = [];
    selected.value = undefined;
  }

  return {
    accounts,
    selected,
    add,
    remove,
    revalidate
  };
});

export { useTradingAccountsStore };
