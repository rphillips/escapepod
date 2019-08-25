import Vue from 'vue';
import Vuex from 'vuex';

import podcasts from './modules/podcasts';

Vue.use(Vuex);

export const store = new Vuex.Store({
  modules: {
    podcasts,
  },
});
