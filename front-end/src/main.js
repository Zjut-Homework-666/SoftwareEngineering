import { createApp } from 'vue'
import App from './App.vue' // 引入组件入口

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css' //引入element-ui plus

import VueAxios from 'vue-axios' //引入vue-axios

import Router from 'vue-router'  //引入vue-router



const app = createApp(App)

app.use(ElementPlus) // 全局挂载ElementuiPlus

app.use(VueAxios, axios) // 全局挂载Axios

app.use(Router) // 全局挂载Router

app.mount('#app')