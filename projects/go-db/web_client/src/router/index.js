import VueRouter from 'vue-router'
import GamePage from '../pages/GamePage'
import GamesPage from '../pages/GameListPage'
import ExtSystemCreatePage from '../pages/ExtSystemCreatePage'
import ExtSystemListPage from '../pages/ExtSystemListPage'

const routes = [
    {
        path: '/games',
        component: GamesPage,
    },
    {
        path: '/game',
        component: GamePage,
    },
    {
        path: '/ext-systems',
        component: ExtSystemListPage,
    },
    {
        path: '/ext-system/create',
        component: ExtSystemCreatePage,
    },
];

export const router = new VueRouter({
    mode: 'history',
    routes
});
