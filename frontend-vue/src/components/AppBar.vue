<template>
  <v-app-bar color="indigo darken-2" dark dense>
    <v-app-bar-nav-icon @click="toggleDrawer" />
    <v-toolbar-title>Escapepod</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn icon @click="addPodcast()">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import router from '../router';

export default {
  methods: {
    toggleDrawer() {
      this.$emit('toggleDrawer');
    },
    addPodcast: () => {
      router.push({ name: 'PodcastCreate' });
    },
  },
};
</script>
