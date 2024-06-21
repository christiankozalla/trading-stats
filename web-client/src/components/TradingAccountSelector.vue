<script setup lang="ts">
import { computed, defineAsyncComponent } from 'vue';
import LoaderInline from './LoaderInline.vue';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { useI18nStore } from '@/stores/i18n';

const TradingAccountCreator = defineAsyncComponent({
  loader: () => import('@/components/TradingAccountCreator.vue'),
  loadingComponent: LoaderInline
});

const tradingAccountsStore = useTradingAccountsStore();
const { t } = useI18nStore();

const options = computed(() => {
  return tradingAccountsStore.accounts?.length
    ? tradingAccountsStore.accounts
    : [{ component: 'internal-trading-account-creator' }];
});

const tradingAccountCreatorProps = computed(() =>
  tradingAccountsStore.accounts?.length
    ? undefined
    : {
        item: {
          props: {
            'aria-selected': false,
            'aria-disabled': true,
            'aria-button': false
          },
          style: 'background-color: transparent; padding: 0'
        }
      }
);

function stopPropagation(e: Event) {
  e.stopPropagation();
}
</script>

<template>
  <Dropdown
    v-model="tradingAccountsStore.selected"
    optionLabel="name"
    optionValue="id"
    :options="options"
    placeholder="Select a trading account"
    :pt="tradingAccountCreatorProps"
  >
    <template #option="slotProps">
      <div
        v-if="slotProps.option.component === 'internal-trading-account-creator'"
        @click="stopPropagation"
        style="width: 100%"
      >
        <p style="padding-left: var(--inline-spacing); padding-top: var(--inline-spacing)">
          {{ t('trading-accounts.create') }}
        </p>
        <TradingAccountCreator />
      </div>
    </template>
  </Dropdown>
</template>
