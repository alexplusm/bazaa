import Vue from 'vue';
import Vuex from 'vuex';
import {api} from '../api'

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
            return api.extSystem.list()
                .then(data => commit('SET_EXT_SYSTEMS', data.extSystems))
        },

        createExtSystem(context, extSystem) {
            return api.extSystem.create(extSystem)
        }
    },
    getters: {
        EXT_SYSTEMS(state) {
            return state.extSystems;
        }
    },
    modules: {}
});
