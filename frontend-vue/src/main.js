import Vue from 'vue';
import './plugins/axios';
import App from './App.vue';
import vuetify from './plugins/vuetify';
import router from './router';
import VeeValidate from 'vee-validate';
import { store } from './store/store';

Vue.config.productionTip = false;
Vue.use(VeeValidate);

new Vue({
  store,
  vuetify,
  router,
  render: (h) => h(App),
}).$mount('#app');
