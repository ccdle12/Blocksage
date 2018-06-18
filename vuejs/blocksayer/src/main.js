import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VeeValidate from 'vee-validate'
import { routes } from './routes'

Vue.config.productionTip = false
Vue.use(VeeValidate)
Vue.use(VueRouter)

const router = new VueRouter({
  routes,
  mode: 'history'
})

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')