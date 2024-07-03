<script setup lang="ts">
import Button from 'primevue/button';
import { pb, isAuthenticated } from '@/api-client';
import { ClientResponseError } from 'pocketbase';
import { useToast } from 'primevue/usetoast';

const toast = useToast();
async function requestPasswordReset() {
  try {
    if (pb.authStore.model?.email) {
      const success = await pb.collection('users').requestPasswordReset(pb.authStore.model.email);
      if (success) {
        toast.add({
          severity: 'success',
          summary: 'Passwort-Reset gestartet',
          detail: 'Du hast eine Email erhalten.'
        });
      }
    }
  } catch (e) {
    if (e instanceof ClientResponseError)
      toast.add({
        severity: 'error',
        summary: 'Passwort-Reset Fehler',
        detail: e.data.message
      });
  }
}
</script>

<template>
  <div v-if="isAuthenticated && pb.authStore.model">
    <ul>
      <li>Name: {{ pb.authStore.model.name }}</li>
      <li>Username: {{ pb.authStore.model.username }}</li>
      <li>Avatar: {{ pb.authStore.model.avatar }}</li>
      <li>Created: {{ pb.authStore.model.created }}</li>
      <li>Email: {{ pb.authStore.model.email }}</li>
      <li><Button @click="requestPasswordReset" label="Passwort zurücksetzen" /></li>
      <li>Verified: {{ pb.authStore.model.verified }}</li>
    </ul>
  </div>
  <div v-else>Not authenticated</div>
</template>
