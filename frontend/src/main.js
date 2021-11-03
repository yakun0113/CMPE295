import { createApp } from 'vue'
import App from './App.vue'
import store from './Store/store'
import router from './router/index'
import Unicon from 'vue-unicons'
import { uniAngleRightB, uniAngleLeftB,uniAngleDoubleRight, uniAngleDoubleLeft,uniFilter, uniSortAmountDown, uniSortAmountUp } from 'vue-unicons/dist/icons'
import VueGoogleMaps from '@fawmi/vue-google-maps'
//import Toaster from '@meforma/vue-toaster'
Unicon.add([uniAngleRightB, uniAngleLeftB,uniAngleDoubleRight, uniAngleDoubleLeft, uniFilter, uniSortAmountDown, uniSortAmountUp])
const app = createApp(App).use(router)
app.use(Unicon)
app.use(store)
app.use(VueGoogleMaps, {
    load: {
        key: 'AIzaSyAQouS0RxRryOqSLk58s3ogACeMFrg1THQ',
    }})
//app.use(Toaster);

app.mount('#app')