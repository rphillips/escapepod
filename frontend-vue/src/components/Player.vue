<template>
  <v-bottom-navigation fixed app height="100" v-if="episode">
    <v-row ma-2 pa-2>
      <v-col>
        <v-card flat height="100%">
          <v-list-item two-line>
            <v-list-item-avatar size="40">
              <v-img :src="`/api/v1/podcasts/${episode.podcast.id}/image`"></v-img>
            </v-list-item-avatar>
            <v-list-item-content pa-2>
              <div class="overline mb-4">{{ episode.podcast.title }}</div>
              <div class="overflow">{{ episode.title }}</div>
            </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>

      <v-col cols="4">
        <v-container>
          <v-spacer></v-spacer>
          <vue-plyr ref="plyr">
            <audio></audio>
          </vue-plyr>
        </v-container>
      </v-col>

      <v-spacer></v-spacer>
    </v-row>
  </v-bottom-navigation>
</template>

<script>
import { VuePlyr } from 'vue-plyr';
import 'vue-plyr/dist/vue-plyr.css';

export default {
  components: {
    VuePlyr,
  },
  props: {
    episode: Object,
  },
  watch: {
    episode: {
      handler: function(ep) {
        this.playEpisode(ep);
      },
    },
  },
  computed: {
    player() {
      return this.$refs.plyr.player;
    },
  },
  methods: {
    playEpisode(ep) {
      this.player.source = {
        type: 'audio',
        sources: [{ src: ep.url, type: ep.type }],
      };
      this.player.play();
    },
  },
};
</script>

<style>
.overflow {
  max-width: 350px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>