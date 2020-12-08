import VueRouter from 'vue-router';

import {store} from '../store/index';

// import GamePage from '../pages/GamePage';
import GamesPage from '../pages/GameListPage';
// import ExtSystemCreatePage from '../pages/ExtSystemCreatePage';
// import ExtSystemListPage from '../pages/ExtSystemListPage';
import AuthPage from '../pages/AuthPage';
import HomePage from "../pages/HomePage";

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
            {path: '', redirect: 'game'},
            {
                path: 'game',
                component: GamesPage,

            },
            { path: '*', redirect: '/home' }
        ],
    },
    { path: '*', redirect: '/' }
    // {
    //     path: '/game/:id/',
    //     component: GamePage,
    // },
    // {
    //     path: '/ext-systems',
    //     component: ExtSystemListPage,
    // },
    // {
    //     path: '/ext-system/create',
    //     component: ExtSystemCreatePage,
    // },
];

export const router = new VueRouter({
    mode: 'history',
    routes,
});

const authGuard = (to, from, next) => {
    const loginRequired = !!to.matched.find(({meta}) => meta.loginRequired);
    const {authorized} = store.state.auth;

    if (loginRequired && !authorized) {
        next('/');
    } else {
        next();
    }
}

router.beforeEach(authGuard)
