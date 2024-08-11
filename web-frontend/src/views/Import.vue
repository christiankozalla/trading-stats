<script setup lang="ts">
import { ref } from "vue";
import { useToast } from "primevue/usetoast";
import { pb } from "@/api-client";
import { useTradingAccountsStore } from "@/stores/tradingAccounts";
import { useI18nStore } from "@/stores/i18n";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import DataPanel from "@/components/DataPanel.vue";
import UploadScreenshot from "@/components/UploadScreenshot.vue";

const toast = useToast();
const tradingAccountsStore = useTradingAccountsStore();
const tradeLogFileNames = ref<string[]>([]);
const loading = ref(false);

const { t } = useI18nStore();

function handleTradeLogFiles(event: Event) {
  const files = (event.target as HTMLInputElement).files;
  if (files?.length) {
    tradeLogFileNames.value = Array.from(files).map((f) => f.name);
  }
}

async function uploadLogFile(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);
  if (!tradingAccountsStore.selected) {
    toast.add({
      severity: "info",
      summary: "Trading account required",
      detail: "Please select/create a trading account before importing.",
      life: 5000
    });
    return;
  } else if (!pb.authStore.model) {
    toast.add({
      severity: "warn",
      summary: "Authentication error",
      detail: "You are not logged in properly.",
      life: 5000
    });
    return;
  }
  loading.value = true;
  formData.append("user", pb.authStore.model.id);
  formData.append("account", tradingAccountsStore.selected);
  const response = await pb
    .collection("trade_log_files")
    .create(formData)
    .catch((e) => {
      toast.add({
        severity: "error",
        summary: t("generic.import-error"),
        detail: e.data.message,
        life: 5000
      });
    })
    .finally(() => {
      loading.value = false;
    });

  if (response) {
    toast.add({
      severity: "success",
      summary: t("generic.import-success"),
      life: 5000
    });
    (event.target as HTMLFormElement).reset();
    tradeLogFileNames.value = [];
  }
}
</script>

<template>
  <div class="container">
    <DataPanel :header="t('import.log-files.header')">
      <div>
        <div v-if="tradeLogFileNames?.length">
          <ul>
            <li v-for="fileName in tradeLogFileNames" :key="fileName">
              <small>{{ fileName }}</small>
            </li>
          </ul>
        </div>
        <div v-else>
          <p v-html="t('import.log-files.description')" />
        </div>

        <form @submit.prevent="uploadLogFile" id="log-files">
          <label for="tradeLogFile" class="p-button p-component p-button-secondary"
            ><span class="p-button-icon p-button-icon-left icon icon-upload"></span
            ><span class="p-button-label">{{
              tradeLogFileNames.length > 0
                ? t('generic.upload-button-change', { type: t('generic.log-files') })
                : t('generic.upload-button-choose', { type: t('generic.log-files') })
            }}</span></label
          >
          <InputText
            type="file"
            name="file"
            id="tradeLogFile"
            class="screen-reader-only"
            accept="text/plain"
            @change="handleTradeLogFiles"
            multiple
            :oninvalid="`this.setCustomValidity('${t('generic.file-input-required')}')`"
            required
          />
          <Button type="submit" label="Submit" :loading="loading" />
        </form>
      </div>
    </DataPanel>
    <DataPanel :header="t('import.screenshot.header')">
      <div>
        <UploadScreenshot />
      </div>
    </DataPanel>
  </div>
</template>

<style scoped>
.container > * {
  margin: var(--content-padding) 0;
}

form#log-files {
  display: flex;
  gap: var(--inline-spacing);
  margin-top: var(--content-padding);
}

form#log-files > * {
  width: 50%;
}
</style>
