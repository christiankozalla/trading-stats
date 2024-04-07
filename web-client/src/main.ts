import 'primevue/resources/themes/aura-light-indigo/theme.css';
import './assets/main.css';
import './assets/icons.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';
import PrimeVue from 'primevue/config';

// Manually importing because usage of unplugin-vue-components infinitely triggers file-watcher pattern in modd.conf
import Menu from 'primevue/menu';
import ProgressSpinner from 'primevue/progressspinner';
import ToastService from 'primevue/toastservice';
import Column from 'primevue/column';
import InputText from 'primevue/inputtext';
import TabMenu from 'primevue/tabmenu';
import Toast from 'primevue/toast';
import DataTable from 'primevue/datatable';
import Button from 'primevue/button';
import Dropdown from 'primevue/dropdown';
import Panel from 'primevue/panel';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(PrimeVue);
app.use(ToastService);

app.component('InputText', InputText);
app.component('TabMenu', TabMenu);
app.component('Toast', Toast);
app.component('DataTable', DataTable);
app.component('Button', Button);
app.component('Dropdown', Dropdown);
app.component('Panel', Panel);
app.component('Menu', Menu);
app.component('Column', Column);
app.component('ProgressSpinner', ProgressSpinner);

app.mount('#app');
