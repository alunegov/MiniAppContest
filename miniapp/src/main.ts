import './assets/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';
import VWave from 'v-wave';

const app = createApp(App);

app.use(createPinia());
app.use(router);
// ripple effect on button click
app.use(VWave, {});

app.mount('#app');
