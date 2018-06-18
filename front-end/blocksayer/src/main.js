import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VeeValidate from 'vee-validate'
import VueResource from 'vue-resource'
import { routes } from './routes'

Vue.config.productionTip = false
Vue.use(VeeValidate)
Vue.use(VueRouter)
Vue.use(VueResource)

const router = new VueRouter({
  routes,
  mode: 'history'
})

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')