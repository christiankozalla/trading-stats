import 'primevue/resources/themes/aura-light-indigo/theme.css';
import './assets/main.css';
import './assets/icons.css';

import { createApp, defineAsyncComponent } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';
import PrimeVue from 'primevue/config';

// Manually importing because usage of unplugin-vue-components infinitely triggers file-watcher pattern in modd.conf
import Menu from 'primevue/menu';
import Menubar from 'primevue/menubar';
import ProgressSpinner from 'primevue/progressspinner';
import ToastService from 'primevue/toastservice';
import Column from 'primevue/column';

const app = createApp(App);

app.use(createPinia());
app.use(PrimeVue);
app.use(ToastService);

app.component(
  'InputText',
  defineAsyncComponent(() => import('primevue/inputtext'))
);
app.component(
  'TabMenu',
  defineAsyncComponent(() => import('primevue/tabmenu'))
);
app.component(
  'Toast',
  defineAsyncComponent(() => import('primevue/toast'))
);
app.component(
  'DataTable',
  defineAsyncComponent(() => import('primevue/datatable'))
);
app.component(
  'Button',
  defineAsyncComponent(() => import('primevue/button'))
);
app.component(
  'Dropdown',
  defineAsyncComponent(() => import('primevue/dropdown'))
);
app.component(
  'Panel',
  defineAsyncComponent(() => import('primevue/panel'))
);

app.component('Menu', Menu);
app.component('Menubar', Menubar);
app.component('Column', Column);
app.component('ProgressSpinner', ProgressSpinner);

app.use(router);

app.mount('#app');
