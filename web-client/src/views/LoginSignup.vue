<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useLoaderStore } from '@/stores/loader';
import type { TabMenuChangeEvent } from 'primevue/tabmenu';
import { pb, type User } from '@/api-client';
import { ClientResponseError } from 'pocketbase';
import { useToast } from 'primevue/usetoast';

const router = useRouter();
const loaderStore = useLoaderStore();
const toast = useToast();

type AuthTypes = 'login' | 'signup';

const formErrors = ref<Record<string, string>>({});
const authType = ref<AuthTypes>('signup');
const authTypes: AuthTypes[] = ['signup', 'login'] as const;

async function signup(event: Event) {
  event.preventDefault();
  formErrors.value = {};

  const formData = new FormData(event.target as HTMLFormElement);
  const email = formData.get('email') as string | null;
  const password = formData.get('password') as string | null;
  const passwordConfirm = formData.get('passwordConfirm') as string | null;

  if (!email) {
    formErrors.value.email = 'Please enter your email address.';
  } else if (!password) {
    formErrors.value.password = 'Please enter your password.';
  } else if (password.length < 8) {
    formErrors.value.password = 'Your password is too short. Please enter 8 characters.';
  } else if (password.length > 72) {
    formErrors.value.password = 'Your password is too long.';
  } else if (!passwordConfirm) {
    formErrors.value.passwordConfirm = 'Please confirm your password.';
  } else if (password !== passwordConfirm) {
    formErrors.value.passwordConfirm = 'Passwords do not match.';
  } else {
    if (event.target) {
      try {
        loaderStore.startLoading();

        const userDataResponse = await pb
          .collection('users')
          .create<User>(new FormData(event.target as HTMLFormElement));

        if (userDataResponse) {
          const hasBeenSent = await pb.collection('users').requestVerification(email);
          if (hasBeenSent)
            toast.add({
              severity: 'success',
              summary: 'Signup Successful',
              detail: 'Please verify your email address!',
              life: 15000
            });
        }

        // log the user in after signup
        // verification can be done within 7 days - unverified older users will be deleted
        await login(event), { redirectOnSuccess: false };
      } catch (e) {
        if (e instanceof ClientResponseError) {
          toast.add({
            severity: 'error',
            summary: 'Signup error',
            detail: e.data.message,
            life: 5000
          });
        }
      } finally {
        loaderStore.stopLoading();
      }
    }
  }
}

async function login(event: Event, { redirectOnSuccess = true } = {}) {
  event.preventDefault();
  formErrors.value = {};

  const formData = new FormData(event.target as HTMLFormElement);
  const email = formData.get('email') as string | null;
  const password = formData.get('password') as string | null;

  if (!email) {
    formErrors.value.email = 'Please enter your email address.';
  } else if (!password) {
    formErrors.value.password = 'Please enter your password.';
  } else if (password.length < 8) {
    formErrors.value.password = 'Your password is too short. Please enter 8 characters.';
  } else if (password.length > 72) {
    formErrors.value.password = 'Your password is too long.';
  } else {
    if (event.target) {
      try {
        loaderStore.startLoading();

        const userDataResponse = await pb
          .collection('users')
          .authWithPassword<User>(email, password);

        if (userDataResponse && redirectOnSuccess) {
          await router.push({ name: 'overview' });
        }
      } catch (e) {
        if (e instanceof ClientResponseError) {
          toast.add({
            severity: 'error',
            summary: 'Login error',
            detail: e.data.message,
            life: 5000
          });
        }
      } finally {
        loaderStore.stopLoading();
      }
    }
  }
}

function onSubmit(event: Event) {
  if (authType.value === 'login') {
    login(event);
  } else if (authType.value === 'signup') {
    signup(event);
  } else {
    throw Error('unknown auth type, neither "login" nor "signup"');
  }
}

const headline = computed(() => {
  if (authType.value === 'login') {
    return 'Login to inloopo trading stats';
  } else if (authType.value === 'signup') {
    return 'Sign up to inloopo trading stats';
  }
  return '';
});

function changeAuthType(event: TabMenuChangeEvent) {
  authType.value = authTypes[event.index];
}
</script>

<template>
  <div class="container">
    <h2>{{ headline }}</h2>
    <TabMenu
      :pt="{
        menuitem: {
          style: {
            'max-width': '50%',
            width: '50%'
          }
        }
      }"
      :model="[
        { label: 'Signup', ariaLabel: 'Signup', icon: 'icon icon-person' },
        { label: 'Login', ariaLabel: 'Login', icon: 'icon icon-login' }
      ]"
      @tab-change="changeAuthType"
    />
    <form @submit.prevent="onSubmit" novalidate>
      <div v-if="authType === 'signup'">
        <label for="name">Name:</label>
        <InputText id="name" name="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <InputText id="email" name="email" required />
        {{ formErrors.email }}
      </div>
      <div>
        <label for="password">Password:</label>
        <InputText
          type="password"
          id="password"
          name="password"
          required
          minlength="8"
          maxlength="72"
        />
        {{ formErrors.password }}
      </div>
      <div v-if="authType === 'signup'">
        <label for="passwordConfirm">Confirm Password:</label>
        <InputText
          type="password"
          id="passwordConfirm"
          name="passwordConfirm"
          required
          minlength="8"
          maxlength="72"
        />
        {{ formErrors.passwordConfirm }}
      </div>
      <Button label="Submit" type="submit" style="width: 100%"></Button>
    </form>
  </div>
</template>

<style>
.container {
  max-width: 500px;
  margin: 0 auto;
}

form {
  display: block;
  margin-top: calc(2 * var(--inline-spacing));
}

form label,
form input,
form button {
  margin-top: var(--inline-spacing);
  margin-bottom: var(--inline-spacing);
}

label {
  display: block;
}

label + input {
  width: 100%;
  max-width: 100%;
}
</style>
