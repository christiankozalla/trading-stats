<script setup lang="ts">
import { ref } from 'vue';
import { pb } from '@/api-client';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { useToast } from 'primevue/usetoast';

const toast = useToast();
const tradingAccountsStore = useTradingAccountsStore();
const fileNames = ref<string[]>([]);

async function uploadLogFile(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);
  if (!tradingAccountsStore.selected) {
    toast.add({
      severity: 'info',
      summary: 'Trading account required',
      detail: 'Please select/create a trading account before importing.',
      life: 5000
    });
    return;
  } else if (!pb.authStore.model) {
    toast.add({
      severity: 'warn',
      summary: 'Authentication error',
      detail: 'You are not logged in properly.',
      life: 5000
    });
    return;
  }

  formData.append('user', pb.authStore.model.id);
  formData.append('account', tradingAccountsStore.selected);
  const response = await pb
    .collection('trade_log_files')
    .create(formData)
    .catch((e) => {
      toast.add({
        severity: 'error',
        summary: 'Import error',
        detail: e.data.message,
        life: 5000
      });
    });
  if (response) {
    toast.add({
      severity: 'success',
      summary: 'Import successful',
      life: 5000
    });
    fileNames.value = [];
  }
}

function onFileInputChange(event: Event) {
  const files = (event.target as HTMLInputElement).files;
  if (files?.length) {
    fileNames.value = Array.from(files).map((f) => f.name);
  }
}
</script>

<template>
  <form @submit.prevent="uploadLogFile">
    <div v-if="fileNames?.length">
      Files to be uploaded
      <ul>
        <li v-for="fileName in fileNames" :key="fileName">{{ fileName }}</li>
      </ul>
    </div>
    <label for="file" class="p-button p-component"
      ><span style="color: white" class="p-button-icon p-button-icon-left icon icon-upload"></span
      ><span class="p-button-label">{{
        fileNames.length > 0 ? 'Change Files' : 'Choose Files'
      }}</span></label
    >
    <InputText
      type="file"
      name="file"
      id="file"
      class="file"
      @change="onFileInputChange"
      multiple
    />
    <Button type="submit" label="Submit"></Button>
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
