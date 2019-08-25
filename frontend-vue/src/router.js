import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('./views/Podcasts.vue'),
    },
    {
      path: '/podcast/create',
      name: 'PodcastCreate',
      component: () => import('./views/PodcastCreate.vue'),
    },
    {
      path: '/podcast/import',
      name: 'PodcastImport',
      component: () => import('./views/PodcastImport.vue'),
    },
    {
      path: '/podcast/:id',
      name: 'PodcastDetail',
      component: () => import('./views/PodcastDetail.vue'),
    },
    {
      path: '/podcasts',
      name: 'podcasts',
      component: () => import('./views/Podcasts.vue'),
    },
  ],
});
