import Vue from 'vue'
import VueRouter from 'vue-router'

import App from './App.vue'
import {router} from './router/index'
import vuetify from './plugins/vuetify';
import store from './store'

Vue.config.productionTip = false

Vue.use(VueRouter)

new Vue({
    render: h => h(App),
    vuetify,
    store,
    router
}).$mount('#app')
