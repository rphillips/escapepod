<template>
  <v-container fluid fill-height>
    <v-snackbar
      :color="snackbarColor"
      :top="snackbarYTopPosition"
      v-model="snackbar"
    >{{ snackbarText }}</v-snackbar>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-11">
          <v-toolbar color="primary" dense dark flat class="indigo darken-2">
            <v-toolbar-title>Import OPML</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-form class="pa-2">
            <v-file-input v-model="opmlFile" prepend-icon="mdi-rss" label="OPML File"></v-file-input>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" @click="add">Upload</v-btn>
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>
export default {
  data: () => ({
    opmlFile: null,

    snackbar: false,
    snackbarColor: '',
    snackbarYTopPosition: true,
    snackbarText: '',
  }),
  methods: {
    add() {
      if (!this.opmlFile) {
        return;
      }
      let formData = new FormData();
      formData.append('file', this.opmlFile);
      this.$store
        .dispatch('importPodcasts', formData)
        .then(() => {
          this.activeSnackbar('success', 'Import opml');
        })
        .catch((err) => {
          this.activeSnackbar('error', err);
        });
      this.clear();
    },
    activeSnackbar(snackbarColor, snackbarText) {
      this.snackbar = true;
      this.snackbarColor = snackbarColor;
      this.snackbarText = snackbarText;
    },
    clear() {
      this.opmlFile = '';
    },
  },
};
</script>