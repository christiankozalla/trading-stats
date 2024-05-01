<script setup lang="ts">
import { ref, computed, type Ref } from 'vue';
import { useToast } from 'primevue/usetoast';
import Textarea from 'primevue/textarea';
import { pb } from '@/api-client';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import { useI18nStore } from '@/stores/i18n';
import DataPanel from '@/components/DataPanel.vue';
import ImagePreview from '@/components/ImagePreview.vue';

const toast = useToast();
const tradingAccountsStore = useTradingAccountsStore();
const tradeLogFileNames = ref<string[]>([]);
const screenshots = ref<FileList | undefined>();
const screenshot = computed(() => {
  return screenshots.value?.item(0) ?? undefined;
});
const { t } = useI18nStore();

const handleTradeLogFiles = onFileInputChange<string[]>(tradeLogFileNames, (files) =>
  Array.from(files).map((f) => f.name)
);
const handleScreenshotFile = onFileInputChange<FileList | undefined>(screenshots, (files) => files);

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
        summary: t('generic.import-error'),
        detail: e.data.message,
        life: 5000
      });
    });

  if (response) {
    toast.add({
      severity: 'success',
      summary: t('generic.import-success'),
      life: 5000
    });

    (event.target as HTMLFormElement).reset();
    tradeLogFileNames.value = [];
  }
}

async function uploadScreenshot(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);

  if (!tradingAccountsStore.selected) {
    toast.add({
      severity: 'info',
      summary: t('generic.trading-account-required'),
      detail: t('generic.trading-account-required-description'),
      life: 5000
    });
    return;
  }

  formData.append('account', tradingAccountsStore.selected);
  const response = await pb
    .collection('screenshots')
    .create(formData)
    .catch((e) => {
      toast.add({
        severity: 'error',
        summary: 'Upload error',
        detail: e.data.data?.image?.message || e.data.message,
        life: 10000
      });
    });

  if (response) {
    toast.add({
      severity: 'success',
      summary: 'Upload successful',
      detail: 'A screenshot record has been created.',
      life: 5000
    });

    (event.target as HTMLFormElement).reset();
    screenshots.value = undefined;
  }
}

function onFileInputChange<T>(destination: Ref<T>, cb: (files: FileList) => T) {
  return function (event: Event) {
    const files = (event.target as HTMLInputElement).files;
    if (files?.length) {
      destination.value = cb(files);
    } else {
      // File selection was cancelled
      if (screenshots.value) {
        (event.target as HTMLInputElement).files = screenshots.value; // add stored FileList back into file-input-element
      }
    }
  };
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
          <p v-html="t('import.log-files.description')" clas="log-files-description" />
        </div>

        <form @submit.prevent="uploadLogFile" id="log-files">
          <label for="tradeLogFile" class="p-button p-component"
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
            class="p-sr-only"
            accept="text/plain"
            @change="handleTradeLogFiles"
            multiple
            :oninvalid="`this.setCustomValidity('${t('generic.file-input-required')}')`"
            required
          />
          <Button type="submit" label="Submit"></Button>
        </form>
      </div>
    </DataPanel>
    <DataPanel :header="t('import.screenshot.header')">
      <div>
        <form @submit.prevent="uploadScreenshot" id="screenshots">
          <div class="screenshot-upload-wrapper p-panel">
            <label for="screenshot" class="screenshot-upload-button p-button p-component"
              ><span class="p-button-icon p-button-icon-left icon icon-upload"></span
              ><span class="p-button-label">{{
                screenshot
                  ? t('generic.upload-button-change', { type: t('generic.image') })
                  : t('generic.upload-button-choose', { type: t('generic.image') })
              }}</span></label
            >
            <InputText
              type="file"
              name="image"
              id="screenshot"
              class="p-sr-only"
              accept="image/png, image/jpeg, image/webp"
              :oninvalid="`this.setCustomValidity('${t('generic.file-input-required')}')`"
              required
              @change="handleScreenshotFile"
            />

            <ImagePreview :file="screenshot" />
          </div>
          <div class="screenshot-metadata-wrapper">
            <label for="screenshot-date" class="p-sr-only">Date</label>
            <InputText
              id="screenshot-date"
              type="date"
              name="date"
              :oninvalid="`this.setCustomValidity('${t('generic.date-input-required')}')`"
              required
            />
            <label for="screenshot-comment" class="p-sr-only">Comment</label>
            <Textarea
              autoResize
              id="screenshot-comment"
              name="comment"
              :placeholder="t('import.screenshot.comment-placeholder')"
            />
            <Button type="submit" label="Submit"></Button>
          </div>
        </form>
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

form#screenshots {
  display: flex;
  gap: var(--inline-spacing);
  height: 300px;
}

.p-button-label {
  overflow-x: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.screenshot-upload-wrapper {
  position: relative;
  flex-basis: 60%;
}

.screenshot-upload-button {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.screenshot-metadata-wrapper {
  display: flex;
  flex-direction: column;
  gap: var(--inline-spacing);
  flex-basis: 40%;
}

#screenshot-date,
#screenshot-comment {
  width: 100%;
}

#screenshot-comment {
  flex-grow: 1;
}
</style>
