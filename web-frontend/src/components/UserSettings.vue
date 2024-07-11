<script setup lang="ts">
import { ref } from 'vue';
import { pb } from '@/api-client';
import { ClientResponseError } from 'pocketbase';
import { useToast } from 'primevue/usetoast';
import { useI18nStore } from '@/stores/i18n';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';

const toast = useToast();
const { t } = useI18nStore();
const loading = ref<boolean>(false);

async function requestPasswordReset() {
  try {
    if (pb.authStore.model?.email) {
      loading.value = true;
      await pb.collection('users').requestPasswordReset(pb.authStore.model.email);
      toast.add({
        severity: 'success',
        summary: t('authentication.password-reset.started'),
        detail: t('authentication.password-reset.started-detail'),
        life: 5000
      });
    }
  } catch (e) {
    if (e instanceof ClientResponseError)
      toast.add({
        severity: 'error',
        summary: 'Passwort-Reset Fehler',
        detail: e.data.message,
        life: 5000
      });
  } finally {
    loading.value = false;
  }
}

async function updateUser(event: Event, field: string) {
  try {
    loading.value = true;
    const formData = new FormData(event.target as HTMLFormElement);
    await pb.collection('users').update(pb.authStore.model?.id as string, formData);
    toast.add({
      severity: 'success',
      summary: t('settings.update-success'),
      detail: t(`settings.user.update-${field}-message`),
      life: 5000
    });
  } catch (err) {
    if (err instanceof ClientResponseError) {
      toast.add({
        severity: 'error',
        summary: t('settings.update-fail'),
        detail: err.data.message,
        life: 5000
      });
    }
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <!-- Display name RU (read update) -->
  <!-- Email RU -->
  <!-- Password U + Reset Password dialog -->
  <!-- Default date style format (can be overwritten by date style format in trading accounts settings) -->
  <!-- Default base currency (can be overwritten by currency in trading accounts settings)  -->
  <!-- Later: Subscription / Plan / Payment info (if not freemium) OR may get its own tab "Subscription" -->
  <!-- Delete account -->
  <form @submit.prevent="(event) => updateUser(event, 'name')">
    <label for="name">
      <h4>
        {{ t('authentication.user-name') }}
      </h4>
    </label>
    <InputText
      id="name"
      name="name"
      :value="pb.authStore.model?.name || ''"
      :placeholder="!pb.authStore.model?.name && t('settings.user.no-name-yet')"
    />
    <Button
      type="submit"
      severity="secondary"
      :label="
        pb.authStore.model?.name ? t('settings.user.update-name') : t('settings.user.create-name')
      "
      :loading="loading"
    />
  </form>

  <form class="mt" @submit.prevent="(event) => updateUser(event, 'email')">
    <label for="email">
      <h4>Email</h4>
    </label>
    <InputText id="email" name="email" :value="pb.authStore.model?.email" />
    :loading="loading"
    <Button
      type="submit"
      severity="secondary"
      :label="t('settings.user.update-email')"
      :loading="loading"
    />
  </form>

  <div class="mt">
    <h4>{{ t('authentication.forgot-password') }}</h4>
    <p>{{ t('authentication.password-reset.description') }}</p>
    <Button
      @click="requestPasswordReset"
      severity="secondary"
      :label="t('authentication.password-reset.btn')"
      :loading="loading"
    />
  </div>
</template>

<style scoped>
.mt {
  margin-top: 3rem;
}

h4 {
  margin-bottom: 0.7rem;
}

label {
  display: block;
}

input {
  width: 300px;
  max-width: 100%;
  margin-right: var(--inline-spacing);
  margin-bottom: var(--inline-spacing);
}

p {
  margin: 1rem 0;
}
</style>
