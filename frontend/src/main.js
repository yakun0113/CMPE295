import { createApp } from 'vue'
import App from './App.vue'
import store from './Store/store'
import router from './router/index'
const app = createApp(App).use(router)

app.use(store)
app.mount('#app')