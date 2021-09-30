import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/components/Home.vue'
import SignIn from '@/components/SignIn.vue'
import SearchResult from '@/components/SearchResult.vue'
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/sign-in',
    name: 'signin',
    component: SignIn
  },
  {
    path: '/search-result',
    name: 'search',
    component: SearchResult
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
