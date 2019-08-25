<template>
  <v-app>
    <v-content>
      <AppBar @toggleDrawer="drawer = !drawer" />
      <Nav v-bind:drawer="drawer" />
      <router-view @playEpisode="playEpisode" />
      <Player v-bind:episode="episode"></Player>
    </v-content>
  </v-app>
</template>

<script>
import AppBar from './components/AppBar.vue';
import Nav from './components/Nav.vue';
import Player from './components/Player.vue';

export default {
  components: {
    AppBar,
    Nav,
    Player,
  },
  data: () => ({
    drawer: false,
    episode: null,
  }),
  methods: {
    playEpisode(episode) {
      this.episode = episode;
    },
  },
};
</script>
