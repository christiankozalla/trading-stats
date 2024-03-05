import 'primevue/resources/themes/aura-light-indigo/theme.css';
import './assets/main.css';
import './assets/icons.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';
import PrimeVue from 'primevue/config';

// Manually importing because usage of unplugin-vue-components infinitely triggers file-watcher pattern in modd.conf
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Menu from 'primevue/menu';
import TabMenu from 'primevue/tabmenu';
import ProgressSpinner from 'primevue/progressspinner';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';


const app = createApp(App);

app.use(createPinia());
app.use(PrimeVue);
app.use(ToastService);

app.component('InputText', InputText);
app.component('Button', Button);
app.component('Menu', Menu);
app.component('TabMenu', TabMenu);
app.component('ProgressSpinner', ProgressSpinner);
app.component('Toast', Toast);


app.use(router);

app.mount('#app');
