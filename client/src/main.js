import Vue from 'vue'
import App from './App.vue'
import VueAnalytics from 'vue-analytics'
Vue.use(VueAnalytics, {
  id: 'UA-126589970-2',
  linkers: ['rey.sh', 'reynaldo.dev']
})

new Vue({
  render: h => h(App),
}).$mount('#app')


