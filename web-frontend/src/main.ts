import "./assets/main.css";
import "./assets/icons.css";

import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import PrimeVue from "primevue/config";
import { definePreset } from "@primevue/themes";
import Aura from "@primevue/themes/aura";
import ToastService from "primevue/toastservice";
import Toast from "primevue/toast";

const app = createApp(App);

app.use(createPinia());
app.use(router);

const PrimevuePreset = definePreset(Aura, {
  semantic: {
    primary: {
      50: "{indigo.50}",
      100: "{indigo.100}",
      200: "{indigo.200}",
      300: "{indigo.300}",
      400: "{indigo.400}",
      500: "{indigo.500}",
      600: "{indigo.600}",
      700: "{indigo.700}",
      800: "{indigo.800}",
      900: "{indigo.900}",
      950: "{indigo.950}"
    }
  }
});

app.use(PrimeVue, {
  theme: {
    preset: PrimevuePreset,
    options: {
      cssLayer: {
        name: "primevue"
      }
    }
  }
});
app.use(ToastService);
app.component("Toast", Toast);

app.provide("host", window.location.host);

app.mount("#app");
