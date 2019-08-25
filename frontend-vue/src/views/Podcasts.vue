<template>
  <v-container v-if="podcasts.length > 0" fluid grid-list-sm>
    <v-layout wrap>
      <v-flex xs4 lg2 v-for="pod of podcasts" :key="pod.id">
        <v-card @click="onPodcastClick(pod)" elevation="8" tile class="ma-1">
          <v-img v-bind:src="getImageURL(pod)"></v-img>
          <v-card-text class="d-none d-md-block">
            <p style class="overflow my-0 py-0 display-11 text--primary">{{ pod.title }}</p>
            <p class="overflow my-0 py-0">{{ pod.author || "Unknown"}}</p>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
  <v-container v-else fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs4>
        <v-card dark color="indigo lighten-1">
          <v-card-text justify-center class="text-center headline">Not subscribed to any podcasts...</v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>
import { mapGetters } from 'vuex';

export default {
  created() {
    this.$store.dispatch('listPodcasts');
  },
  computed: {
    ...mapGetters(['podcasts']),
  },
  methods: {
    getImageURL: function(podcast) {
      return `/api/v1/podcasts/${podcast.id}/image`;
    },
    onPodcastClick: function(podcast) {
      this.$router.push({
        name: 'PodcastDetail',
        params: { id: podcast.id },
      });
    },
  },
};
</script>
<style>
.overflow {
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
</style>
