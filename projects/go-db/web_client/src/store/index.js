import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
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

// const auth = {
//     authorized: false
// }
