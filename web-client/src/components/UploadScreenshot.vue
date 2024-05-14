<script setup lang="ts">
import { ref, computed } from 'vue';
import { useToast } from 'primevue/usetoast';
import { pb } from '@/api-client';
import { useTradingAccountsStore } from '@/stores/tradingAccounts';
import ImagePreview from '@/components/ImagePreview.vue';
import { useI18nStore } from '@/stores/i18n';

const { t } = useI18nStore();
const toast = useToast();
const tradingAccountsStore = useTradingAccountsStore();

const screenshots = ref<FileList | undefined>();
const screenshot = computed(() => {
  return screenshots.value?.item(0) ?? undefined;
});

const props = defineProps<{ date?: string | null }>();

function handleScreenshotFile(event: Event) {
  const files = (event.target as HTMLInputElement).files;
  if (files?.length) {
    screenshots.value = files;
  } else if (screenshots.value) {
    // File selection was cancelled
    (event.target as HTMLInputElement).files = screenshots.value; // add stored FileList back into file-input-element
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
</script>

<template>
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
        :pt="{ root: { value: props.date } }"
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
</template>

<style scoped>
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
