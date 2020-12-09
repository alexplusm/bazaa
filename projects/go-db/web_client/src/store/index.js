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
        currentGame: null,
        games: [],
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
        },
        setGames(state, games) {
            state.games = games;
        },
        setCurrentGame(state, game) {
            state.currentGame = game;
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
        getGameList({commit, state}) {
            const {currentExtSystem} = state;

            if (currentExtSystem && currentExtSystem.extSystemId) {
                return api.game.list(state.currentExtSystem.extSystemId)
                    .then(data => commit('setGames', data.games));
            }
        },
        getGameDetails({commit, state}, gameId) {
            const {currentExtSystem} = state;

            if (currentExtSystem && currentExtSystem.extSystemId) {
                return api.game.details(gameId, currentExtSystem.extSystemId)
                    .then(data => commit('setCurrentGame', data));
            }
        },
        updateGameWithArchive({dispatch}, {gameId, file}) {
            return api.game.updateWithFile(gameId, file)
                .finally(() => dispatch('getGameDetails', gameId))
        }
    },
    getters: {
        extSystems(state) {
            return state.extSystems;
        },
        currentExtSystem(state) {
            return state.currentExtSystem;
        },
        games(state) {
            return state.games;
        },
        currentGame(state) {
            return state.currentGame;
        }
    },
    modules: {}
});
