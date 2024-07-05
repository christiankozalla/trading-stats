<script setup lang="ts">
import { ref, computed, nextTick } from 'vue';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { useRouter } from 'vue-router';
import { useLoaderStore } from '@/stores/loader';
import { pb, type User } from '@/api-client';
import { ClientResponseError } from 'pocketbase';
import { useToast } from 'primevue/usetoast';
import { useI18nStore } from '@/stores/i18n';

const router = useRouter();
const loaderStore = useLoaderStore();
const toast = useToast();
const { t } = useI18nStore();

type AuthTypes = 'login' | 'signup';

const formErrors = ref<Record<string, string>>({});
const authType = ref<AuthTypes>('signup');
const validateEmail = (email: string | null): email is string =>
  email ? /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/.test(email) : false;

async function signup(event: Event) {
  event.preventDefault();
  formErrors.value = {};

  const formData = new FormData(event.target as HTMLFormElement);
  const email = formData.get('email') as string | null;
  const password = formData.get('password') as string | null;
  const passwordConfirm = formData.get('passwordConfirm') as string | null;
  const validated = validateSignupData(email, password, passwordConfirm);

  if (validated.email && validated.password && validated.passwordConfirm) {
    if (event.target) {
      try {
        loaderStore.startLoading();

        const userDataResponse = await pb
          .collection('users')
          .create<User>(new FormData(event.target as HTMLFormElement));

        if (userDataResponse) {
          const hasBeenSent = await pb.collection('users').requestVerification(validated.email);
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

  const validated = validateLoginData(email, password);

  if (validated.email && validated.password) {
    if (event.target) {
      try {
        loaderStore.startLoading();

        const userDataResponse = await pb
          .collection('users')
          .authWithPassword<User>(validated.email, validated.password);

        if (userDataResponse && redirectOnSuccess) {
          await router.push({ name: 'overview' });
        }
      } catch (e) {
        if (e instanceof ClientResponseError) {
          toast.add({
            severity: 'error',
            summary: 'Login error',
            detail: `${e.data.message} Check email and password.`,
            life: 5000
          });
        }
      } finally {
        loaderStore.stopLoading();
      }
    }
  }
}

function validateSignupData(
  email: string | null,
  password: string | null,
  passwordConfirm: string | null
) {
  const validated: Record<string, string> = {};
  const isValidEmail = validateEmail(email);

  if (!email) {
    formErrors.value.email = 'Please enter your email address.';
  } else if (!isValidEmail) {
    formErrors.value.email = 'Please enter a valid email address.';
  } else {
    validated.email = email;
  }
  if (!password) {
    formErrors.value.password = 'Please enter your password.';
  } else if (password.length < 8) {
    formErrors.value.password = 'Your password is too short. Please enter 8 characters.';
  } else if (password.length > 72) {
    formErrors.value.password = 'Your password is too long.';
  } else {
    validated.password = password;
  }
  if (!passwordConfirm) {
    formErrors.value.passwordConfirm = 'Please confirm your password.';
  } else if (password !== passwordConfirm) {
    formErrors.value.passwordConfirm = 'Passwords do not match.';
  } else {
    validated.passwordConfirm = passwordConfirm;
  }
  return validated;
}

function validateLoginData(email: string | null, password: string | null) {
  const validated: Record<string, string> = {};
  const isValidEmail = validateEmail(email);

  if (!email) {
    formErrors.value.email = 'Please enter your email address.';
  } else if (!isValidEmail) {
    formErrors.value.email = 'Please enter a valid email address.';
  } else {
    validated.email = email;
  }
  if (!password) {
    formErrors.value.password = 'Please enter your password.';
  } else if (password.length < 8) {
    formErrors.value.password = 'Your password is too short. Please enter 8 characters.';
  } else if (password.length > 72) {
    formErrors.value.password = 'Your password is too long.';
  } else {
    validated.password = password;
  }
  return validated;
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
    return t('authentication.headline-login');
  } else if (authType.value === 'signup') {
    return t('authentication.headline-signup');
  }
  return '';
});

const passwordContextOpen = ref<boolean>(false);
async function requestPasswordResetEmail(event: Event) {
  const formData = new FormData(event.target as HTMLFormElement);
  for (const entries of formData.entries()) {
    console.log('entries', entries);
  }
  const email = formData.get('email') as string | null;
  if (!validateEmail(email)) {
    toast.add({
      severity: 'error',
      summary: 'Please enter a valid email address.',
      life: 10000
    });
    return;
  }
  try {
    loaderStore.startLoading();
    const emailSent = await pb.collection('users').requestPasswordReset(email);
    if (emailSent) {
      toast.add({
        severity: 'success',
        summary: 'Password reset requested successfully',
        detail: 'Check your inbox to complete password reset process.',
        life: 5000
      });
    }
  } catch (e) {
    if (e instanceof ClientResponseError)
      toast.add({
        severity: 'error',
        summary: 'Failed to request password reset email',
        detail: `${e.data.message} Check email and password.`,
        life: 5000
      });
  } finally {
    loaderStore.stopLoading();
  }
}

const openPasswordResetContext = () => {
  passwordContextOpen.value = !passwordContextOpen.value;
  passwordContextOpen.value &&
    nextTick().then(() => {
      document.getElementById('password-reset-email')?.focus();
    });
};
</script>

<template>
  <div class="container">
    <h2>{{ headline }}</h2>
    <Tabs v-model:value="authType">
      <TabList>
        <Tab value="signup">
          <i class="icon icon-person" />
          Signup
        </Tab>
        <Tab value="login">
          <i class="icon icon-login" />
          Login
        </Tab>
      </TabList>
    </Tabs>
    <form @submit.prevent="onSubmit" novalidate>
      <div v-if="authType === 'signup'">
        <label for="name">Name:</label>
        <InputText id="name" name="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <InputText id="email" name="email" required />
        <small
          ><i>{{ formErrors.email }}</i></small
        >
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
        <small
          ><i>{{ formErrors.password }}</i></small
        >
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
        <small
          ><i>{{ formErrors.passwordConfirm }}</i></small
        >
      </div>
      <Button label="Submit" type="submit" style="width: 100%"></Button>
    </form>

    <Button
      v-if="authType === 'login'"
      label="Forgot password?"
      link
      @click="openPasswordResetContext"
      class="password-reset-button"
    />
    <form v-if="passwordContextOpen" @submit.prevent="requestPasswordResetEmail" novalidate>
      <label for="password-reset-email">Request password reset email</label>
      <InputText
        type="email"
        name="email"
        id="password-reset-email"
        placeholder="Your account's email"
        required
      />
      <Button label="Request password reset" type="submit" severity="secondary" />
    </form>
  </div>
</template>

<style scoped>
.container {
  max-width: 500px;
  margin: 0 auto 120px;
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

small {
  color: var(--p-red-400);
}

.password-reset-button {
  display: block;
  margin-left: auto;
  margin-right: 0;
  font-size: small;
}

:deep(.p-tablist-tab-list > *) {
  flex: 1 1 0;
  display: flex;
  align-items: center;
  gap: var(--inline-spacing);
  font-size: 15px;
}
</style>
