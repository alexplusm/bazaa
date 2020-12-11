import Vue from 'vue';
import { Settings } from 'luxon'
import VueRouter from 'vue-router';
import { Datetime } from 'vue-datetime';

// You need a specific loader for CSS files
import 'vue-datetime/dist/vue-datetime.css';

import App from './App.vue';
import { router } from './router';
import { store } from './store';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;

Vue.use(VueRouter);

Vue.component('Datetime', Datetime);
Settings.defaultLocale = 'ru';

new Vue({
	render: (h) => h(App),
	vuetify,
	store,
	router,
}).$mount('#app');
