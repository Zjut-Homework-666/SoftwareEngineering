import {createApp, getCurrentInstance} from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css'
import * as echarts from 'echarts'
// import VueRouter from 'vue-router';
import router from './router/index.js'
import axios from 'axios'
// import VueAxios from 'vue-axios'
// import './plugins/axios'
import App from './App.vue'
const app = createApp(App)

// 全局方法挂载全局变量
app.config.globalProperties.$BackendPort = "8089";
app.config.globalProperties.$url = "http://localhost:";
//调用方法
// const proxy :any = getCurrentInstance().appContext.config.globalProperties
app.use(router)
// app.use(VueAxios, axios)
app.use(ElementPlus)
app.config.globalProperties.$echarts = echarts
app.config.globalProperties.$axios = axios

app.mount('#app')
