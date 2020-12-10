import VueRouter from 'vue-router';

import { store } from '../store/index';

import AuthPage from '../pages/AuthPage';
import HomePage from '../pages/HomePage';
import GamePage from '../pages/GamePage';
import GamesPage from '../pages/GameListPage';
import GameCreationPage from '../pages/GameCreationPage';
import ExtSystemCreatePage from '../pages/ExtSystemCreatePage';

const routes = [
	{
		path: '/',
		component: AuthPage,
	},
	{
		path: '/home',
		component: HomePage,
		meta: { loginRequired: true },
		children: [
			{ path: '', redirect: 'game' },
			{
				path: 'game',
				component: GamesPage,
			},
			{
				path: 'game-create', // todo: -> game/create
				component: GameCreationPage,
			},
			{
				path: 'game/:id',
				component: GamePage,
			},
			{
				path: 'ext-system/create',
				component: ExtSystemCreatePage,
			},
			{ path: '*', redirect: '/home' },
		],
	},
	{ path: '*', redirect: '/' },
];

export const router = new VueRouter({
	mode: 'history',
	routes,
});

const authGuard = (to, from, next) => {
	const loginRequired = !!to.matched.find(({ meta }) => meta.loginRequired);
	const { authorized } = store.state.auth;

	// TODO: "/:id" !!!! don't work

	// console.log("from: ", from);
	// console.log("to: ", to);

	if (loginRequired && !authorized) {
		next('/');
	} else {
		next();
	}
};

router.beforeEach(authGuard);
