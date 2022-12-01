import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css'

// import VueRouter from 'vue-router';
import router from '../src/router/index.js'

import App from './App.vue'


const app = createApp(App)
app.config.productionTip = false

app.use(router)
app.use(ElementPlus)
app.mount('#app')
