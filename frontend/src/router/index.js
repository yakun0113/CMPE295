import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/components/Home.vue'
import SignIn from '@/components/SignIn.vue'
import SearchResult from '@/components/SearchResult.vue'
import Account from '@/components/Account.vue'
import Watchlist from '@/components/Watchlist.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    props: true
  },
  {
    path: '/sign-in',
    name: 'signin',
    component: SignIn
  },
  {
    path: '/search-result',
    name: 'search-result',
    component: SearchResult,
    props: true
  },
  {
    path: '/account',
    name: 'account',
    component: Account,
    props: true
  },
  {
    path: '/watchlist',
    name: 'watchlist',
    component: Watchlist,
    props: true
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})



export default router
