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
            <v-toolbar-title>Add Podcast</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <v-form class="pa-2">
            <v-text-field
              :error-messages="errors.collect('feedURL')"
              refs="feedURL"
              prepend-icon="mdi-rss"
              required
              v-validate="'required|min:6|url'"
              v-model="feedURL"
              data-vv-name="feedURL"
              label="Feed URL"
              type="text"
            ></v-text-field>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" @click="add">Add</v-btn>
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
    feedURL: '',

    snackbar: false,
    snackbarColor: '',
    snackbarYTopPosition: true,
    snackbarText: '',
  }),
  methods: {
    async add() {
      let res = await this.$validator.validateAll();
      if (!res) {
        return;
      }
      this.$store
        .dispatch('createPodcast', this.feedURL)
        .then(() => {
          this.activeSnackbar('success', 'Added podcast');
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
      this.feedURL = '';
      this.$validator.reset();
    },
  },
};
</script>