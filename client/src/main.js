import Vue from 'vue'
import App from './App.vue'
import VueAnalytics from 'vue-analytics'
Vue.use(VueAnalytics, {
  id: ['UA-126589970-2', 'UA-126589970-3'],
})

new Vue({
  render: h => h(App),
}).$mount('#app')


