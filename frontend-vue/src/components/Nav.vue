<template>
  <v-navigation-drawer app v-model="drawer" class="indigo darken-2">
    <v-list nav>
      <v-list-item v-for="link in links" :key="link.text" router :to="link.route" link>
        <v-list-item-icon>
          <v-icon class="white--text">{{ link.icon }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title class="white--text">{{link.text}}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  props: {
    drawer: Boolean,
  },
  data: () => ({
    links: [
      { icon: 'mdi-podcast', text: 'Podcasts', route: '/' },
      { icon: 'mdi-plus', text: 'Add Podcasts', route: '/podcast/create' },
      { icon: 'mdi-plus', text: 'Import OPML', route: '/podcast/import' },
    ],
  }),
};
</script>