import Vue from 'vue';
import Vuex from 'vuex';
import { api } from '../api';

Vue.use(Vuex);

export const store = new Vuex.Store({
	state: {
		auth: {
			username: null,
			password: null,
			authorized: false,
		},
		extSystems: [],
		currentExtSystem: null,
		currentGame: null,
		games: [],
	},
	mutations: {
		authorize(state, credentials) {
			state.auth.username = credentials.username;
			state.auth.password = credentials.password;
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
		},
	},
	actions: {
		authorize({ commit }, credentials) {
			return api.auth.check(credentials).then((result) => {
				if (result) {
					commit('authorize', credentials);
				}
				return result;
			});
		},
		getExtSystemList({ commit, state }) {
			return api.extSystem
				.list(state.auth)
				.then((data) => commit('setExtSystemList', data.extSystems));
		},
		createExtSystem({ state }, extSystem) {
			return api.extSystem.create(extSystem, state.auth);
		},
		setCurrentExtSystem({ commit, dispatch }, extSystem) {
			commit('setCurrentExtSystem', extSystem);
			dispatch('getGameList');
		},
		getGameList({ commit, state }) {
			const { currentExtSystem } = state;

			if (currentExtSystem && currentExtSystem.extSystemId) {
				return api.game
					.list(state.currentExtSystem.extSystemId, state.auth)
					.then((data) => commit('setGames', data.games));
			}
		},
		getGameDetails({ commit, state }, gameId) {
			const { currentExtSystem } = state;

			if (currentExtSystem && currentExtSystem.extSystemId) {
				return api.game
					.details(gameId, currentExtSystem.extSystemId, state.auth)
					.then((data) => commit('setCurrentGame', data));
			}
		},
		attachArchivesToGame(
			{ dispatch, state },
			{ gameId, file, progressCallback }
		) {
			return api.game
				.attachArchives(gameId, file, progressCallback, state.auth)
				.finally(() => dispatch('getGameDetails', gameId));
		},
		attachAnotherGameResultsToGame({dispatch, state }, {gameId, formValue}) {
			return api.game.attachAnotherGameResults(gameId, formValue, state.auth)
				.finally(() => dispatch('getGameDetails', gameId));
		},
		createGame({ state }, game) {
			return api.game.create(game, state.auth);
		},
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
		},
	},
	modules: {},
});
