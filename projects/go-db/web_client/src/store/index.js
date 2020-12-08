import Vue from 'vue';
import Vuex from 'vuex';
import {api} from '../api'

Vue.use(Vuex);

export const store = new Vuex.Store({
    state: {
        auth : {
            authorized: false,
        },
        extSystems: [],
        currentExtSystem: null,
    },
    mutations: {
        authorize(state) {
            state.auth.authorized = true;
        },
        setExtSystemList(state, extSystems) {
            state.extSystems = extSystems;
        },
        setCurrentExtSystem(state, extSystem) {
            state.currentExtSystem = extSystem;
        }
    },
    actions: {
        getExtSystemList({commit}) {
            return api.extSystem.list()
                .then(data => commit('setExtSystemList', data.extSystems))
        },
        createExtSystem(context, extSystem) {
            return api.extSystem.create(extSystem)
        },
        setCurrentExtSystem({commit, dispatch}, extSystem) {
            commit('setCurrentExtSystem', extSystem);
            dispatch('getGameList');
        },
        getGameList({state}) {
            // todo check: state.currentExtSystem.extSystemId
            // todo: commit

            if (state.currentExtSystem && state.currentExtSystem.extSystemId) {
                return api.game.list(state.currentExtSystem.extSystemId);
            }
            return;
        }
    },
    getters: {
        extSystems(state) {
            return state.extSystems;
        },
        currentExtSystem(state) {
            return state.currentExtSystem;
        }
    },
    modules: {}
});
