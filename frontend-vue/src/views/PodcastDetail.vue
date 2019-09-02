<template>
  <v-container fluid grid-list-sm>
    <v-layout wrap v-if="podcast.id">
      <v-flex lg2>
        <v-card outlined>
          <v-img contains v-bind:src="`/api/v1/podcasts/${podcast.id}/image`" class="white--text" />
          <v-card-title>
            <v-btn text icon v-bind:href="podcast.link" target="_blank">
              <v-icon>mdi-link</v-icon>
            </v-btn>
          </v-card-title>
        </v-card>
      </v-flex>
      <v-flex lg10>
        <v-card outlined>
          <v-list-item two-line>
            <v-list-item-content class="align-self-start">
              <v-list-item-title class="headline mb-2" v-text="podcast.title"></v-list-item-title>
              <v-list-item-subtitle v-html="podcast.description"></v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-data-table
          dense
          v-bind:items-per-page="itemsPerPage"
          :headers="headers"
          :items="episodes"
        >
          <template v-slot:body="{ items }">
            <tbody>
              <tr v-on:dblclick="onDoubleClick(ep)" v-for="ep in items" :key="ep.id">
                <td width="15">{{ ep.id }}</td>
                <td class="overflow">{{ ep.title }}</td>
                <td width="120">{{ formatDuration(ep.duration) }}</td>
                <td width="125">{{ formatDate(ep.published) }}</td>
                <td>
                  <v-layout>
                    <span class="text-truncate" style="max-width: 300px">{{ ep.description }}</span>
                    <v-menu v-model="menu" :position-x="x" :position-y="y" absolute offset-y>
                      <template v-slot:activator="{ on }">
                        <v-icon right small v-on="on" class="right-img">mdi-information-variant</v-icon>
                      </template>
                      <v-card>
                        <v-card-title>{{ep.title}}</v-card-title>
                        <v-card-text v-html="ep.description"></v-card-text>
                      </v-card>
                    </v-menu>
                  </v-layout>
                </td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>
import { mapGetters } from 'vuex';
import format from 'date-fns/format';
import humanizeDuration from 'humanize-duration';

export default {
  data() {
    return {
      itemsPerPage: 15,
      headers: [
        { text: 'ID', value: 'id' },
        { text: 'Name', value: 'title' },
        { text: 'Time', value: 'time' },
        { text: 'Released', value: 'released' },
        { text: 'Description', value: 'description' },
      ],
    };
  },
  computed: {
    ...mapGetters(['podcast', 'episodes']),
  },
  methods: {
    formatDate(dateStr) {
      return format(dateStr, 'MMM DD, YYYY');
    },
    formatDuration(durationStr) {
      let a = durationStr.split(':');
      while (a.length < 3) {
        a.unshift(0);
      }
      const ms = (parseInt(a[0]) * 60 * 60 + parseInt(a[1]) * 60 + parseInt(a[2])) * 1000;
      const shortEnglishHumanizer = humanizeDuration.humanizer({
        language: 'shortEn',
        languages: {
          shortEn: {
            y: () => 'year',
            mo: () => 'month',
            w: () => 'week',
            d: () => 'day',
            h: () => 'hr',
            m: () => 'min',
            s: () => 'sec',
            ms: () => 'ms',
          },
        },
      });
      return shortEnglishHumanizer(ms, {
        units: ['h', 'm'],
        round: true,
        largest: 2,
        delimiter: ' ',
      });
    },
    onDoubleClick: function(item) {
      item.podcast = this.podcast;
      this.$emit('playEpisode', item);
    },
  },
  mounted() {
    this.$store.dispatch('listEpisodes', this.$route.params.id);
    this.$store.dispatch('getPodcast', this.$route.params.id);
  },
};
</script>

<style>
.overflow {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>