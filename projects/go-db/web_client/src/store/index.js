import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

export const store =  new Vuex.Store({
    state: {
        auth : {
            authorized: false,
        },
        extSystems: [],
    },
    mutations: {
        authorize(state) {
            state.auth.authorized = true;
        },
        SET_EXT_SYSTEMS(state, extSystems) {
            state.extSystems = extSystems;
        }
    },
    actions: {
        GET_EXT_SYSTEMS_FROM_API({commit}) {
            return axios.get('/api/v1/ext-system')
                .then(({data}) => data)
                .then(data => {
                    if (!data.success) {
                        // todo: throw error
                    }
                    return data.data;
                })
                .then(data => commit('SET_EXT_SYSTEMS', data.extSystems))
        },

        // GET_GAMES_FROM_API(inst) {
        //     return axios.get('/api/v1/game')
        //         .then(resp => {
        //             console.log("inst", inst);
        //             console.log("resp: ", resp);
        //         })
        // }
    },
    getters: {
        EXT_SYSTEMS(state) {
            return state.extSystems;
        }
    },
    modules: {}
});
