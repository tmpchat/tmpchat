import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import TopPage from './pages/TopPage'
import ChatPage from './pages/ChatPage'

Vue.config.productionTip = false
Vue.use(VueRouter)

const routes = [
  { path: '/', component: TopPage },
  { path: '/chat/:id', component: ChatPage }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
