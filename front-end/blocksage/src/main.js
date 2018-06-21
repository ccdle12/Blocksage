import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import { routes } from './routes'

Vue.config.productionTip = process.env.NODE_ENV === 'development'
console.log('Main js: ', process.env.API_BASE_URL)
Vue.use(VueRouter)
Vue.use(VueAxios, axios)

const router = new VueRouter({
  routes,
  mode: 'history'
})

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')