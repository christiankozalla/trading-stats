<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { pb } from '@/api-client';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';

const tradingAccountsStore = useTradingAccountsStore();
const filePath = ref();
const selectedTradingAccount = ref<string | null>(null);

onMounted(async () => {
  await tradingAccountsStore.get();
});

async function uploadLogFile(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);
  if (!selectedTradingAccount.value) {
    return;
    // toast.add()
  } else if (!pb.authStore.model) {
    return;
    // toast.add()
  }

  formData.append("user", pb.authStore.model.id)
  formData.append("account", selectedTradingAccount.value);
  const response = await pb
    .collection('trade_log_files')
    .create(formData)
    .catch((e) => {
      console.log('uploadLogFile error', e.data);
    });

  console.log('success', response);
}
</script>

<template>
  <form @submit.prevent="uploadLogFile">
    <Dropdown
      v-model="selectedTradingAccount"
      optionLabel="name"
      optionValue="id"
      :options="tradingAccountsStore.accounts"
      placeholder="Select a trading account"
    />
    <InputText type="date" name="date" />
    <label for="file" class="p-button p-component"
      ><span style="color: white" class="p-button-icon p-button-icon-left icon icon-upload"></span
      ><span class="p-button-label">{{
        filePath ? filePath.split('\\').at(-1) : 'Import File'
      }}</span></label
    >
    <InputText
      type="file"
      name="file"
      id="file"
      class="file"
      v-model="filePath"
      @change="() => console.log(filePath)"
    />
    <Button type="submit" label="submit"></Button>
  </form>
</template>

<style scoped>
form {
  display: flex;
  flex-direction: column;
  gap: var(--inline-spacing);
  max-width: 300px;
}

.p-button-label {
  overflow-x: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.file {
  display: none;
}
</style>
