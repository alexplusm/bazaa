import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const store =  new Vuex.Store({
    state: {
        auth : {
            authorized: false,
        },
    },
    mutations: {
        authorize(state) {
            state.auth.authorized = true;
        }
    },
    actions: {},
    modules: {}
});
