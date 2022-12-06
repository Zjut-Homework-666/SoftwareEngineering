import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css'

// import VueRouter from 'vue-router';
import router from './router/index.js'
import axios from 'axios'
import VueAxios from 'vue-axios'
import App from './App.vue'


const app = createApp(App)

app.use(router)
app.use(VueAxios, axios)
app.use(ElementPlus)

app.mount('#app')
