import VueRouter from 'vue-router'
import GamePage from '../pages/GamePage'
import GamesPage from '../pages/GamesPage'

const routes = [
    {
        path: '/games',
        component: GamesPage,
    },
    {
        path: '/game',
        component: GamePage
    }
]

export default new VueRouter({
    mode: 'history',
    routes
})
